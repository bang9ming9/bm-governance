package test

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/bang9ming9/bm-governance/abis"
	govTypes "github.com/bang9ming9/bm-governance/types"
	"github.com/bang9ming9/go-hardhat/bms"
	utils "github.com/bang9ming9/go-hardhat/bms/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

var (
	ctx      = context.Background()
	callOpts = new(bind.CallOpts)
)

type BMGovernor struct {
	*govTypes.BMGovernor
	target *utils.Contract[abis.TargetContract]
}

func DeployBMGovernorWithBackend(t *testing.T) (*bms.Backend, *BMGovernor) {
	var err error
	contracts := new(BMGovernor)

	backend := bms.NewBacked(t)
	contracts.BMGovernor, err = govTypes.DeployBMGovernor(context.Background(), backend.Owner, backend,
		struct {
			Name   string
			Symbol string
		}{"", ""},
		struct {
			Name    string
			Version string
			Uri     string
		}{"", "", ""}, struct{ Name string }{""},
	)
	require.NoError(t, err)

	target, _, err := utils.DeployContract(abis.DeployTargetContract(backend.Owner, backend, contracts.BMGovernor.Governor.Address()))
	require.NoError(t, err)
	backend.Commit()

	contracts.target, err = target.SetABIWithError(abis.TargetContractMetaData.GetAbi())
	require.NoError(t, err)

	return backend, contracts
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

func (bm *BMGovernor) ChargeERC20(t *testing.T, backend utils.Backend, eoas []*bind.TransactOpts) []*bind.TransactOpts {
	txpool := utils.NewTxPool(backend)
	owner := bms.GetOwner(t)
	owner.Value = utils.ToWei(1)
	for _, eoa := range eoas {
		require.NoError(t, txpool.Exec(bm.Erc20.Funcs().Mint(owner, eoa.From)))
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
	events, err := utils.UnpackEvents[abis.BmGovernorProposalCreated](bm.Governor.ABI(), "ProposalCreated", receipt)
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
	callClock, err := bm.Erc1155.Funcs().Clock(callOpts)
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
	clock, err := bm.Erc1155.Funcs().Clock(callOpts)
	require.NoError(t, err)
	start, err := bm.Governor.Funcs().ProposalSnapshot(callOpts, proposalID)
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
