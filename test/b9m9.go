package test

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/bang9ming9/bm-governance/abis"
	"github.com/bang9ming9/go-hardhat/bms"
	utils "github.com/bang9ming9/go-hardhat/bms/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

var (
	ctx      = context.Background()
	callOpts = new(bind.CallOpts)
)

type BMGovernor struct {
	erc20    *utils.Contract[abis.BmErc20]
	erc1155  *utils.Contract[abis.BmErc1155]
	governor *utils.Contract[abis.BmGovernor]
	target   *utils.Contract[abis.TargetContract]
}

type Backend interface {
	bind.DeployBackend
	bind.ContractBackend
}

func DeployBMGovernorWithBackend(t *testing.T) (*bms.Backend, *BMGovernor) {
	backend := bms.NewBacked(t)
	contracts, err := DeployBMGovernor(backend.Owner, backend)
	require.NoError(t, err)
	return backend, contracts
}

func DeployBMGovernor(owner *bind.TransactOpts, backend Backend) (*BMGovernor, error) {
	nonce, err := backend.PendingNonceAt(ctx, owner.From)
	if err != nil {
		return nil, errors.Wrap(err, "PendingNonceAt")
	}

	erc20Address := crypto.CreateAddress(owner.From, nonce)
	erc1155Address := crypto.CreateAddress(owner.From, nonce+1)
	governorAddress := crypto.CreateAddress(owner.From, nonce+2)

	txpool := utils.NewTxPool(backend)

	contracts := new(BMGovernor)
	erc20, tx, err := utils.DeployContract(abis.DeployBmErc20(owner, backend, "", "", erc1155Address, governorAddress))
	if err != nil {
		return nil, errors.Wrap(err, "DeployBmErc20")
	}
	txpool.Append(tx)
	erc1155, tx, err := utils.DeployContract(abis.DeployBmErc1155(owner, backend, "", "", "", erc20Address, governorAddress))
	if err != nil {
		return nil, errors.Wrap(err, "DeployBmErc1155")
	}
	txpool.Append(tx)
	governor, tx, err := utils.DeployContract(abis.DeployBmGovernor(owner, backend, "", erc1155Address))
	if err != nil {
		return nil, errors.Wrap(err, "DeployBmGovernor")
	}
	txpool.Append(tx)
	target, tx, err := utils.DeployContract(abis.DeployTargetContract(owner, backend, governorAddress))
	if err != nil {
		return nil, errors.Wrap(err, "DeployTargetContract")
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
	if contracts.erc20, err = erc20.SetABIWithError(abis.BmErc20MetaData.GetAbi()); err != nil {
		return nil, errors.Wrap(err, "abis.BmErc20MetaData.GetAbi")
	} else if contracts.erc1155, err = erc1155.SetABIWithError(abis.BmErc1155MetaData.GetAbi()); err != nil {
		return nil, errors.Wrap(err, "abis.BmErc1155MetaData.GetAbi")
	} else if contracts.governor, err = governor.SetABIWithError(abis.BmGovernorMetaData.GetAbi()); err != nil {
		return nil, errors.Wrap(err, "abis.BmGovernorMetaData.GetAbi")
	} else if contracts.target, err = target.SetABIWithError(abis.TargetContractMetaData.GetAbi()); err != nil {
		return nil, errors.Wrap(err, "abis.TargetContractMetaData.GetAbi")
	}

	return contracts, nil
}

type Proposal struct {
	Targets     []common.Address `json:"targets"`
	Values      []*big.Int       `json:"values"`
	CallDatas   [][]byte         `json:"calldatas"`
	Description string           `json:"description"`
}

func (bm *BMGovernor) NewProposalToTarget(t *testing.T, desc string, values ...interface{}) *Proposal {
	length := len(values)
	require.True(t, length > 0)
	targets := make([]common.Address, length)
	zeroValues := make([]*big.Int, length)
	datas := make([][]byte, length)
	targetABI, err := abis.TargetContractMetaData.GetAbi()
	require.NoError(t, err)

	for i, value := range values {
		targets[i] = bm.target.Address()
		zeroValues[i] = common.Big0
		var (
			data []byte
			err  error
		)
		switch v := value.(type) {
		case int, uint, int32, uint32, int64, uint64, *big.Int:
			d, _ := new(big.Int).SetString(fmt.Sprint(v), 10)
			data, err = targetABI.Pack("writeUintValue", d)
		case common.Address, [20]byte:
			data, err = targetABI.Pack("writeAddrValue", v)
		case common.Hash, [32]byte:
			data, err = targetABI.Pack("writeB32Value", v)
		case string:
			data, err = targetABI.Pack("writeStrValue", v)
		default:
			err = errors.New("invalid type")
		}
		require.NoError(t, err)
		datas[i] = data
	}

	return &Proposal{
		Targets:     targets,
		Values:      zeroValues,
		CallDatas:   datas,
		Description: desc,
	}
}

func (bm *BMGovernor) ChargeERC20(t *testing.T, backend Backend, eoas []*bind.TransactOpts) []*bind.TransactOpts {
	txpool := utils.NewTxPool(backend)
	owner := bms.GetOwner(t)
	owner.Value = utils.ToWei(1)
	for _, eoa := range eoas {
		require.NoError(t, txpool.Exec(bm.erc20.Funcs().Mint(owner, eoa.From)))
	}
	owner.Value = common.Big0

	receipts, err := txpool.WaitMined(ctx)
	require.NoError(t, err)
	for _, receipt := range receipts {
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	}

	return eoas
}

func (bm *BMGovernor) UnpackProposalCreated(t *testing.T, receipt *types.Receipt) *abis.BmGovernorProposalCreated {
	events, err := utils.UnpackEvents[abis.BmGovernorProposalCreated](bm.governor.ABI(), "ProposalCreated", receipt)
	require.NoError(t, err)
	require.Equal(t, 1, len(events))
	return events[0]
}

const (
	ONE_WEEK      int64 = 86400 * 7
	VOTE_DAY      int64 = 86400 * 4
	MIN_DELAY     int64 = 86400 * 1
	VOTING_PERIOD int64 = ONE_WEEK - VOTE_DAY
)

func (bm *BMGovernor) NextProposalTime(t *testing.T, backend interface {
	AdjustTime(adjustment time.Duration) error
	Commit() common.Hash
}) {
	callClock, err := bm.erc1155.Funcs().Clock(callOpts)
	require.NoError(t, err)
	clock := callClock.Int64()
	clock = (clock - (clock % ONE_WEEK)) + ONE_WEEK
	require.NoError(t, backend.AdjustTime(time.Duration(clock-callClock.Int64())))
	backend.Commit()
}

func (bm *BMGovernor) ToProposalSnapshotTime(t *testing.T, backend interface {
	AdjustTime(adjustment time.Duration) error
	Commit() common.Hash
}, proposalID *big.Int) {
	clock, err := bm.erc1155.Funcs().Clock(callOpts)
	require.NoError(t, err)
	start, err := bm.governor.Funcs().ProposalSnapshot(callOpts, proposalID)
	require.NoError(t, err)
	if diff := new(big.Int).Sub(start, clock).Int64(); diff > 0 {
		require.NoError(t, backend.AdjustTime(time.Duration(diff)))
		backend.Commit()
	}
}

type voteType struct {
	Against uint8
	For     uint8
	Abstain uint8
}

var VoteType = voteType{0, 1, 2}

type proposalState struct {
	Pending   uint8
	Active    uint8
	Canceled  uint8
	Defeated  uint8
	Succeeded uint8
	Queued    uint8
	Expired   uint8
	Execute   uint8
}

var ProposalState = proposalState{0, 1, 2, 3, 4, 5, 6, 7}
var (
	WIN_CLAIM     = utils.ToWei(0.5)
	LOSE_CLAIM    = utils.ToWei(0.1)
	ABSTAIN_CLAIM = utils.ToWei(0.2)
)
