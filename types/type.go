package types

import (
	"context"
	"time"

	"github.com/bang9ming9/bm-governance/abis"
	utils "github.com/bang9ming9/go-hardhat/bms/bmsutils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

type BMGovernor struct {
	Erc20    *utils.Contract[abis.BmErc20]
	Erc1155  *utils.Contract[abis.BmErc1155]
	Governor *utils.Contract[abis.BmGovernor]
}

func DeployBMGovernor(
	ctx context.Context,
	owner *bind.TransactOpts,
	backend utils.Backend, timeout time.Duration,
	erc20Config struct {
		Name, Symbol string
	},
	erc1155Config struct {
		Name, Version, Uri string
	},
	governanceConfig struct {
		Name string
	},
) (*BMGovernor, error) {
	nonce, err := backend.PendingNonceAt(ctx, owner.From)
	if err != nil {
		return nil, errors.Wrap(err, "PendingNonceAt")
	}

	erc20Address := crypto.CreateAddress(owner.From, nonce)
	erc1155Address := crypto.CreateAddress(owner.From, nonce+1)
	governorAddress := crypto.CreateAddress(owner.From, nonce+2)

	txpool := utils.NewTxPool(backend)
	if timeout != 0 {
		txpool.SetTimeout(timeout)
	}

	contracts := new(BMGovernor)
	erc20, tx, err := utils.DeployContract(abis.DeployBmErc20(owner, backend, erc20Config.Name, erc20Config.Symbol, erc1155Address, governorAddress))
	if err != nil {
		return nil, errors.Wrap(err, "DeployBmErc20")
	}
	txpool.Append(tx)
	erc1155, tx, err := utils.DeployContract(abis.DeployBmErc1155(owner, backend, erc1155Config.Name, erc1155Config.Version, erc1155Config.Uri, erc20Address, governorAddress))
	if err != nil {
		return nil, errors.Wrap(err, "DeployBmErc1155")
	}
	txpool.Append(tx)
	governor, tx, err := utils.DeployContract(abis.DeployBmGovernor(owner, backend, governanceConfig.Name, erc1155Address))
	if err != nil {
		return nil, errors.Wrap(err, "DeployBmGovernor")
	}
	txpool.Append(tx)

	// check contract address
	if erc20Address != erc20.Address() {
		return nil, errors.New("invalid erc20 address")
	} else if erc1155Address != erc1155.Address() {
		return nil, errors.New("invalid erc1155 address")
	} else if governorAddress != governor.Address() {
		return nil, errors.New("invalid governor address")
	}

	// check transaction receipts
	if receipts, err := txpool.WaitMined(ctx); err != nil {
		return nil, errors.New("WaitMined")
	} else {
		for _, receipt := range receipts {
			if receipt.Status != types.ReceiptStatusSuccessful {
				return nil, errors.New("deploy failed")
			}
		}
	}

	// set abis
	if contracts.Erc20, err = erc20.SetABIWithError(abis.BmErc20MetaData.GetAbi()); err != nil {
		return nil, errors.Wrap(err, "abis.BmErc20MetaData.GetAbi")
	} else if contracts.Erc1155, err = erc1155.SetABIWithError(abis.BmErc1155MetaData.GetAbi()); err != nil {
		return nil, errors.Wrap(err, "abis.BmErc1155MetaData.GetAbi")
	} else if contracts.Governor, err = governor.SetABIWithError(abis.BmGovernorMetaData.GetAbi()); err != nil {
		return nil, errors.Wrap(err, "abis.BmGovernorMetaData.GetAbi")
	} else {
		return contracts, nil
	}
}
