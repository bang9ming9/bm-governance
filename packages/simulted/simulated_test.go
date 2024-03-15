package simulated_test

import (
	"context"
	"testing"

	simulated "github.com/bang9ming9/bm-governance/packages/simulted"
	"github.com/stretchr/testify/require"
)

func TestNewBackend(t *testing.T) {
	backend := simulated.NewBacked(t)
	ctx := context.Background()
	balance, err := backend.BalanceAt(ctx, backend.Owner, nil)
	require.NoError(t, err)
	t.Log(backend.Owner, "balance", balance)

	eoa := simulated.GetEOA(t)
	balance, err = backend.BalanceAt(ctx, eoa.From, nil)
	require.NoError(t, err)
	t.Log(eoa.From, "balance", balance)

}
