package simulated

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
)

func ChainID() *big.Int {
	return params.AllDevChainProtocolChanges.ChainID
}

type Backend struct {
	simulated.Backend
	simulated.Client
	Owner common.Address
}

func NewBacked(t *testing.T) *Backend {
	owner := GetOwner(t).From

	backend := simulated.NewBackend(
		core.GenesisAlloc{
			owner: core.GenesisAccount{Balance: common.Big256},
		},
		func(nodeConf *node.Config, ethConf *ethconfig.Config) {
			ethConf.Genesis.BaseFee = common.Big0
			ethConf.Miner.GasPrice = common.Big0
		},
	)

	return &Backend{
		Backend: *backend,
		Client:  backend.Client(),
		Owner:   owner,
	}
}

func (ec *Backend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return common.Big0, nil
}

func (ec *Backend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return common.Big0, nil
}
