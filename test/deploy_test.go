package test

import (
	"math/big"
	"testing"
	"time"

	simulated "github.com/bang9ming9/bm-governance/packages/simulted"
	"github.com/bang9ming9/bm-governance/packages/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestDeployBMGovernor(t *testing.T) {
	backend, contracts := DeployBMGovernorWithBackend(t)
	eoas := simulated.GetEOAs(t, 10)
	// TC1 : ERC20 을 정상적으로 얻을 수 있는지 확인
	t.Run("TC1", func(t *testing.T) {
		// 1. eoas 에게 코인들 넉넉하게 전송한다.
		owner := backend.Owner
		owner.Value = utils.ToWei(2)
		txs := make([]*types.Transaction, 0)
		for _, eoa := range eoas {
			tx, err := utils.SendDynamicTx(backend, owner, &eoa.From, []byte{})
			require.NoError(t, err)
			require.NotNil(t, tx)
			txs = append(txs, tx)
		}
		for _, tx := range txs {
			receipt, err := bind.WaitMined(ctx, backend, tx)
			require.NoError(t, err)
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		}
		owner.Value = common.Big0

		// 2. eoas 의 반은 mint 실행
		txs = make([]*types.Transaction, 0)
		for _, eoa := range eoas[:len(eoas)/2] {
			eoa.Value = utils.ToWei(1)
			tx, err := contracts.erc20.Mint(eoa, eoa.From)
			require.NoError(t, err)
			require.NotNil(t, tx)
			txs = append(txs, tx)
			eoa.Value = common.Big0
		}
		for _, tx := range txs {
			receipt, err := bind.WaitMined(ctx, backend, tx)
			require.NoError(t, err)
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		}

		// 3. 나머지 eoas 는 erc20 으로 Send value
		txs = make([]*types.Transaction, 0)
		for _, eoa := range eoas[len(eoas)/2:] {
			eoa.Value = utils.ToWei(1)
			tx, err := utils.SendDynamicTx(backend, eoa, &contracts.address.erc20, []byte{})
			require.NoError(t, err)
			require.NotNil(t, tx)
			txs = append(txs, tx)
			eoa.Value = common.Big0
		}
		for _, tx := range txs {
			receipt, err := bind.WaitMined(ctx, backend, tx)
			require.NoError(t, err)
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		}

		// 4. eoas 의 erc20 balance 확인
		for _, eoa := range eoas {
			balance, err := contracts.erc20.BalanceOf(callOpts, eoa.From)
			require.NoError(t, err)
			require.Equal(t, utils.ToWei(1), balance)
		}

		// 5. 이미 받아갔기 때문에 mint 를 더 실행하지는 못한다.
		// 5-1. MINT
		txs = make([]*types.Transaction, 0)
		eoa := eoas[0]
		tx, err := contracts.erc20.Mint(owner, eoa.From)
		if err == nil {
			txs = append(txs, tx)
		}
		require.Error(t, err)
		// 5-2. SEND VALUE
		tx, err = utils.SendDynamicTx(backend, eoa, &contracts.address.erc20, []byte{})
		if err == nil {
			txs = append(txs, tx)
		}
		for _, tx := range txs {
			receipt, err := bind.WaitMined(ctx, backend, tx)
			require.NoError(t, err)
			require.NotEqual(t, types.ReceiptStatusSuccessful, receipt.Status)
		}
	})
	// TC2 : ERC20 으로 ERC1155 가 정상적으로 mint 되는지 확인
	t.Run("TC2", func(t *testing.T) {
		emptyEOA := simulated.GetEOA(t)
		blen := len(eoas)

		currentID, err := contracts.erc1155.CurrentID(callOpts)
		require.NoError(t, err)
		require.False(t, currentID.Sign() == 0)
		// 1. mint 확인
		txs := make([]*types.Transaction, 0)
		for _, eoa := range append(eoas, emptyEOA) {
			tx, err := contracts.erc1155.Mint(eoa, eoa.From)
			require.NoError(t, err)
			require.NotNil(t, tx)
			txs = append(txs, tx)
		}
		require.Equal(t, blen, len(eoas))
		for _, tx := range txs {
			receipt, err := bind.WaitMined(ctx, backend, tx)
			require.NoError(t, err)
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		}

		onewei := utils.ToWei(1)
		for _, eoa := range eoas {
			balance20, err := contracts.erc20.BalanceOf(callOpts, eoa.From)
			require.NoError(t, err)
			balance, err := contracts.erc1155.BalanceOf(callOpts, eoa.From, currentID)
			require.NoError(t, err)
			require.Equal(t, balance20, balance)
			require.Equal(t, onewei, balance)
		}

		// 2. 같은 id 라면 추가 민팅은 안된다
		txs = make([]*types.Transaction, 0)
		for _, eoa := range append(eoas, emptyEOA) {
			tx, err := contracts.erc1155.Mint(eoa, eoa.From)
			require.NoError(t, err)
			require.NotNil(t, tx)
			txs = append(txs, tx)
		}
		require.Equal(t, blen, len(eoas))
		for _, tx := range txs {
			receipt, err := bind.WaitMined(ctx, backend, tx)
			require.NoError(t, err)
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		}
		for _, eoa := range eoas {
			balance20, err := contracts.erc20.BalanceOf(callOpts, eoa.From)
			require.NoError(t, err)
			balance, err := contracts.erc1155.BalanceOf(callOpts, eoa.From, currentID)
			require.NoError(t, err)
			require.Equal(t, balance20, balance)
			require.Equal(t, onewei, balance) // 어차피 1 ether
		}

		// 3. id가 바뀌면 민팅을 한다.
		backend.AdjustTime(time.Duration(86400 * 7)) // 1 주일을 보낸다.
		nextCurrentId, err := contracts.erc1155.CurrentID(callOpts)
		require.NoError(t, err)
		require.Equal(t, new(big.Int).Add(currentID, common.Big1), nextCurrentId) // id가 1 증가하였는지 확인
		txs = make([]*types.Transaction, 0)
		for _, eoa := range append(eoas, emptyEOA) {
			tx, err := contracts.erc1155.Mint(eoa, eoa.From)
			require.NoError(t, err)
			require.NotNil(t, tx)
			txs = append(txs, tx)
		}
		require.Equal(t, blen, len(eoas))
		for _, tx := range txs {
			receipt, err := bind.WaitMined(ctx, backend, tx)
			require.NoError(t, err)
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		}
		for _, eoa := range eoas {
			balance20, err := contracts.erc20.BalanceOf(callOpts, eoa.From)
			require.NoError(t, err)
			// 이전 id 도 1 ether
			curBalance, err := contracts.erc1155.BalanceOf(callOpts, eoa.From, currentID)
			require.NoError(t, err)
			require.Equal(t, balance20, curBalance)
			require.Equal(t, onewei, curBalance)
			// 현재 id 도 1 ether
			nextCurBalance, err := contracts.erc1155.BalanceOf(callOpts, eoa.From, nextCurrentId)
			require.NoError(t, err)
			require.Equal(t, balance20, nextCurBalance)
			require.Equal(t, onewei, nextCurBalance)
		}
	})
}
