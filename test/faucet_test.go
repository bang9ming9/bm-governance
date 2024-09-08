package test

import (
	"context"
	"math/big"
	"testing"

	"github.com/bang9ming9/bm-governance/abis"
	"github.com/bang9ming9/go-hardhat/bms"
	utils "github.com/bang9ming9/go-hardhat/bms/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// TC1: FAUCET 에 보유한 코인이 없다면 지급할 수 없다.
// TC2: 일일 한번만 코인을 지급 받을 수 있다.
func TestFaucet(t *testing.T) {
	ctx := context.Background()
	callOpts := &bind.CallOpts{Context: ctx}

	backend := bms.NewBacked(t)
	txpool := utils.NewTxPool(backend)

	faucet, tx, FAUCET, err := abis.DeployFaucet(backend.Owner, backend.Client)
	require.NoError(t, txpool.Exec(tx, err))
	require.NoError(t, txpool.AllReceiptStatusSuccessful(ctx))

	require.NotNil(t, FAUCET)
	testers := bms.GetEOAs(t, 2)

	// TC1: FAUCET 에 보유한 코인이 없다면 지급할 수 없다.
	t.Run("TC1", func(t *testing.T) {
		for _, tester := range testers {
			require.Error(t, txpool.Exec(FAUCET.Claim(tester)))
		}
		require.NoError(t, txpool.Clear())
	})

	// FAUCET 에 토큰 지급
	backend.Owner.Value = utils.ToWei((len(testers) * 10) + 11)
	require.NoError(t, txpool.Exec(utils.SendDynamicTx(backend, backend.Owner, &faucet, []byte{})))
	backend.Owner.Value = nil
	require.NoError(t, txpool.AllReceiptStatusSuccessful(ctx))

	// TC2: 일일 한번만 코인을 지급 받을 수 있다.
	t.Run("TC2", func(t *testing.T) {
		// claim 신청
		for _, tester := range testers {
			require.NoError(t, txpool.Exec(FAUCET.Claim(tester)))
		}
		require.NoError(t, txpool.AllReceiptStatusSuccessful(ctx))

		// 잔액 확인
		// - 지급 받아야 하는 수량 컨트랙트에서 조회
		amount, err := FAUCET.AMOUNT(callOpts)
		require.True(t, amount.Sign() != 0)
		require.NoError(t, err)
		for _, tester := range testers {
			// - 테스터의 잔액 확인
			balance, err := backend.Client.PendingBalanceAt(ctx, tester.From)
			require.NoError(t, err)
			require.Equal(t, amount, balance)
		}

		// claim 재신청 (실패 확인)
		for _, tester := range testers {
			require.Error(t, txpool.Exec(FAUCET.Claim(tester)))
		}
		require.NoError(t, txpool.Clear())

		// 하루 시간 보낸뒤 claim 신청
		require.NoError(t, backend.AdjustTime(86400))
		// - tester[0] 은 지급을 받을 수 있따.
		require.NoError(t, txpool.Exec(FAUCET.Claim(testers[0])))
		require.NoError(t, txpool.AllReceiptStatusSuccessful(ctx))
		// - tester[1:] 은 컨트랙트에 코인이 부족하여 지급 받을 수 없다.
		for _, tester := range testers[1:] {
			require.Error(t, txpool.Exec(FAUCET.Claim(tester)))
		}
		require.NoError(t, txpool.Clear())
		for i, tester := range testers {
			// - 테스터의 잔액 확인
			balance, err := backend.Client.PendingBalanceAt(ctx, tester.From)
			require.NoError(t, err)
			if i == 0 {
				require.Equal(t, new(big.Int).Mul(amount, common.Big2), balance)
			} else {
				require.Equal(t, amount, balance)
			}
		}
	})
}
