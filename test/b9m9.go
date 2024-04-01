package test

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/bang9ming9/bm-governance/abis"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	simulated "github.com/bang9ming9/bm-governance/packages/simulted"
	"github.com/bang9ming9/bm-governance/packages/utils"
)

var (
	ctx      = context.Background()
	callOpts = new(bind.CallOpts)
)

type BmGovernor struct {
	erc20    *abis.BmERC20
	erc1155  *abis.BmERC1155
	governor *abis.BmGovernor
	target   *abis.TargetContract
	address  struct {
		erc20    common.Address
		erc1155  common.Address
		governor common.Address
		target   common.Address
	}
	abis struct {
		erc20    *abi.ABI
		erc1155  *abi.ABI
		governor *abi.ABI
		target   *abi.ABI
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

	erc20, tx, ERC20, err := abis.DeployBmERC20(owner, backend, "Bang9Ming9 Governance TOKEN", "BM20", erc1155Address)
	require.NoError(t, err)
	require.Equal(t, erc20Address, erc20)
	txs = append(txs, tx)

	erc1155, tx, ERC1155, err := abis.DeployBmERC1155(owner, backend, "Bang9Ming9 Governance VOTE", "1", "https://github.com/bang9ming9", erc20Address, governorAddress)
	require.NoError(t, err)
	require.Equal(t, erc1155Address, erc1155)
	txs = append(txs, tx)

	governor, tx, GOVERNOR, err := abis.DeployBmGovernor(owner, backend, "Bang9Ming9 Governance", erc1155Address)
	require.NoError(t, err)
	require.Equal(t, governorAddress, governor)
	txs = append(txs, tx)

	target, tx, TARGET, err := abis.DeployTargetContract(owner, backend, governor)
	require.NoError(t, err)
	txs = append(txs, tx)

	for _, tx := range txs {
		ctx2, cancel := context.WithTimeout(ctx, 5e9)
		receipt, err := bind.WaitMined(ctx2, backend, tx)
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		cancel()
	}

	erc20ABI, err := abis.BmERC20MetaData.GetAbi()
	require.NoError(t, err)
	erc1155ABI, err := abis.BmERC1155MetaData.GetAbi()
	require.NoError(t, err)
	governorABI, err := abis.BmGovernorMetaData.GetAbi()
	require.NoError(t, err)
	targetABI, err := abis.TargetContractMetaData.GetAbi()
	require.NoError(t, err)

	simulated.ClearErrors()
	simulated.EnrollErrors(erc20ABI, erc1155ABI, governorABI, targetABI)

	return &BmGovernor{
		erc20:    ERC20,
		erc1155:  ERC1155,
		governor: GOVERNOR,
		target:   TARGET,
		address: struct {
			erc20    common.Address
			erc1155  common.Address
			governor common.Address
			target   common.Address
		}{erc20: erc20, erc1155: erc1155, governor: governor, target: target},
		abis: struct {
			erc20    *abi.ABI
			erc1155  *abi.ABI
			governor *abi.ABI
			target   *abi.ABI
		}{erc20ABI, erc1155ABI, governorABI, targetABI},
	}
}

type Proposal struct {
	Targets     []common.Address `json:"targets"`
	Values      []*big.Int       `json:"values"`
	CallDatas   [][]byte         `json:"calldatas"`
	Description string           `json:"description"`
}

func (bm *BmGovernor) NewProposalToTarget(t *testing.T, desc string, values ...interface{}) *Proposal {
	length := len(values)
	require.True(t, length > 0)
	targets := make([]common.Address, length)
	zeroValues := make([]*big.Int, length)
	datas := make([][]byte, length)
	targetABI, err := abis.TargetContractMetaData.GetAbi()
	require.NoError(t, err)

	for i, value := range values {
		targets[i] = bm.address.target
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

func (bm *BmGovernor) ChargeERC20(t *testing.T, backend *simulated.Backend, eoas []*bind.TransactOpts) []*bind.TransactOpts {
	length := len(eoas)
	txs := make([]*types.Transaction, length)
	owner := backend.Owner
	owner.Value = utils.ToWei(1)
	for i, eoa := range eoas {
		tx, err := bm.erc20.Mint(backend.Owner, eoa.From)
		require.NoError(t, err)
		txs[i] = tx
	}
	owner.Value = common.Big0

	for _, tx := range txs {
		receipt, err := bind.WaitMined(ctx, backend, tx)
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	}
	return eoas
}

func (bm *BmGovernor) UnpackProposalCreated(t *testing.T, logs []*types.Log) (event abis.BmGovernorProposalCreated) {
	return UnpackLog[abis.BmGovernorProposalCreated](t, bm.abis.governor, "ProposalCreated", logs)
}

func UnpackLog[T any](t *testing.T, aBI *abi.ABI, name string, logs []*types.Log) T {
	event := new(T)
	e, ok := aBI.Events[name]
	require.True(t, ok)

	for _, log := range logs {
		if log.Topics[0] == e.ID {
			require.NoError(t, aBI.UnpackIntoInterface(event, name, log.Data))
			var indexed abi.Arguments
			for _, arg := range e.Inputs {
				if arg.Indexed {
					indexed = append(indexed, arg)
				}
			}
			require.NoError(t, abi.ParseTopics(event, indexed, log.Topics[1:]))
		}
	}

	return *event
}

const (
	ONE_WEEK      int64 = 86400 * 7
	VOTE_DAY      int64 = 86400 * 4
	MIN_DELAY     int64 = 86400 * 1
	VOTING_PERIOD int64 = ONE_WEEK - VOTE_DAY
)

func (bm *BmGovernor) NextProposalTime(t *testing.T, backend interface {
	AdjustTime(adjustment time.Duration) error
	Commit() common.Hash
}) {
	callClock, err := bm.erc1155.Clock(callOpts)
	require.NoError(t, err)
	clock := callClock.Int64()
	clock = (clock - (clock % ONE_WEEK)) + ONE_WEEK
	require.NoError(t, backend.AdjustTime(time.Duration(clock-callClock.Int64())))
	backend.Commit()
}

func (bm *BmGovernor) ProposalVoteTime(t *testing.T, backend interface {
	AdjustTime(adjustment time.Duration) error
	Commit() common.Hash
}) {
	callClock, err := bm.erc1155.Clock(callOpts)
	require.NoError(t, err)
	clock := callClock.Int64()
	clock = (clock - (clock % ONE_WEEK)) + VOTE_DAY
	if diff := clock - callClock.Int64(); diff > 0 {
		require.NoError(t, backend.AdjustTime(time.Duration(diff)))
	}
	backend.Commit()
}

type voteType struct {
	Against uint8
	For     uint8
	Abstain uint8
}

var VoteType = voteType{0, 1, 2}
