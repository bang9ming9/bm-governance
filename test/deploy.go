package test

import (
	"context"
	"testing"

	"github.com/bang9ming9/bm-governance/abis"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	simulated "github.com/bang9ming9/bm-governance/packages/simulted"
)

var (
	ctx      = context.Background()
	callOpts = new(bind.CallOpts)
)

type BmGovernor struct {
	erc20    *abis.BMerc20
	erc1155  *abis.BMerc1155
	governor *abis.BmGovernor
	address  struct {
		erc20    common.Address
		erc1155  common.Address
		governor common.Address
	}
}

type Backend interface {
	bind.DeployBackend
	bind.ContractBackend
}

func DeployBMGovernorWithBackend(t *testing.T) (*simulated.Backend, *BmGovernor) {
	backend := simulated.NewBacked(t)
	return backend, DeployBMGovernor(t, backend)
}

func DeployBMGovernor(t *testing.T, backend Backend) *BmGovernor {
	owner := simulated.GetOwner(t)
	nonce, err := backend.PendingNonceAt(ctx, owner.From)
	require.NoError(t, err)

	erc20Address := crypto.CreateAddress(owner.From, nonce)
	erc1155Address := crypto.CreateAddress(owner.From, nonce+1)
	governorAddress := crypto.CreateAddress(owner.From, nonce+2)

	txs := make([]*types.Transaction, 0)

	erc20, tx, ERC20, err := abis.DeployBMerc20(owner, backend, "Bang9Ming9 Governance TOKEN", "BM20", erc1155Address)
	require.NoError(t, err)
	require.Equal(t, erc20Address, erc20)
	txs = append(txs, tx)

	erc1155, tx, ERC1155, err := abis.DeployBMerc1155(owner, backend, "Bang9Ming9 Governance VOTE", "1", "https://github.com/bang9ming9", erc20Address, governorAddress)
	require.NoError(t, err)
	require.Equal(t, erc1155Address, erc1155)
	txs = append(txs, tx)

	governor, tx, GOVERNOR, err := abis.DeployBmGovernor(owner, backend, "Bang9Ming9 Governance", erc1155Address)
	require.NoError(t, err)
	require.Equal(t, governorAddress, governor)
	txs = append(txs, tx)

	for _, tx := range txs {
		ctx2, cancel := context.WithTimeout(ctx, 5e9)
		receipt, err := bind.WaitMined(ctx2, backend, tx)
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		cancel()
	}

	return &BmGovernor{
		erc20:    ERC20,
		erc1155:  ERC1155,
		governor: GOVERNOR,
		address: struct {
			erc20    common.Address
			erc1155  common.Address
			governor common.Address
		}{erc20: erc20, erc1155: erc1155, governor: governor},
	}
}
