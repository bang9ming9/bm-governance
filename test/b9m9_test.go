package test

import (
	"context"
	"math/big"
	"testing"
	"time"

	simulated "github.com/bang9ming9/bm-governance/packages/simulted"
	"github.com/bang9ming9/bm-governance/packages/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// TC1 : ERC20 을 정상적으로 얻을 수 있는지 확인
// TC2 : ERC20 으로 ERC1155 가 정상적으로 mint 되는지 확인
func TestDeployBMGovernorVote(t *testing.T) {
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

/* TC1: Vote(ERC1155) 가 없다면 제안(Propose)할 수 없다. */
/* TC2: TOKEN(ERC20) 가 있다면 제안 할 수 있다.
 *	TC2-2: 제안자를 제외한 나머지 인원은 자유롭게 투표할 수 있다. (찬성 투표)
 *		: CastVote 에서 ERC20 을 가지고 있다면 ERC1155 를 알아서 mint 한다.
 *	TC2-2-1: 중복 투표는 할수 없다.
 *	TC2-2-2: ERC20 이 없다면 투표할 수 없다
 *	TC2-3: 찬성된 제안은 투표 기간이 지나면 실행할 수 있다.
 *		: 투표 기간중에는 실행할 수 없다.
 *	TC2-3-1: 한번 실행된 제안은 두번이상 실행될 수 없다. */
/* TC3: 대표자(delegatee)를 설정하지 않은경우, 자신을 대표자로 설정한다. */
/* TC4: 대표자(delegatee)를 다른 유저로 설정하였다면 제안할 수 없다. */
func TestDeployBMGovernorProposal(t *testing.T) {
	backend, contracts := DeployBMGovernorWithBackend(t)
	eoas := contracts.ChargeERC20(t, backend, simulated.GetEOAs(t, 5))

	// TC1: Vote(ERC1155) 가 없다면 제안(Propose)할 수 없다.
	t.Run("TC1", func(t *testing.T) {
		proposal := contracts.NewProposalToTarget(t, "TC1", 1, common.Address{1}, common.Hash{1}, "1")
		tx, err := contracts.governor.Propose(backend.Owner, proposal.Targets, proposal.Values, proposal.CallDatas, proposal.Description)
		if err == nil {
			receipt, err := bind.WaitMined(ctx, backend, tx)
			require.NoError(t, err)
			require.NotEqual(t, types.ReceiptStatusSuccessful, receipt.Status)
		}
	})

	// TC2: TOKEN(ERC20) 가 있다면 제안 할 수 있다.
	// TC2-2: 제안자를 제외한 나머지 인원은 자유롭게 투표할 수 있다. (찬성 투표)
	// 		: CastVote 에서 ERC20 을 가지고 있다면 ERC1155 를 알아서 mint 한다.
	// TC2-2-1: 중복 투표는 할수 없다.
	// TC2-2-2: ERC20 이 없다면 투표할 수 없다
	// TC2-3: 찬성된 제안은 투표 기간이 지나면 실행할 수 있다.
	// 		: 투표 기간중에는 실행할 수 없다.
	// TC2-3-1: 한번 실행된 제안은 두번이상 실행될 수 없다.
	t.Run("TC2", func(t *testing.T) {
		contracts.NextProposalTime(t, backend)

		proposer := eoas[0]
		// 제안자 가 투표토큰이 있는지 확인
		balance20, err := contracts.erc20.BalanceOf(callOpts, proposer.From)
		require.NoError(t, err)
		require.Equal(t, utils.ToWei(1), balance20)
		// 투표권 획득
		tx, err := contracts.erc1155.Mint(proposer, proposer.From)
		require.NoError(t, err)
		receipt, err := bind.WaitMined(ctx, backend, tx)
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		currentID, err := contracts.erc1155.CurrentID(callOpts)
		require.NoError(t, err)
		balance1155, err := contracts.erc1155.BalanceOf(callOpts, proposer.From, currentID)
		require.NoError(t, err)
		require.Equal(t, balance20, balance1155)

		// 대표자 설정
		tx, err = contracts.erc1155.Delegate(proposer, proposer.From)
		require.NoError(t, err)
		receipt, err = bind.WaitMined(ctx, backend, tx)
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		votes, err := contracts.erc1155.GetVotes(callOpts, proposer.From)
		require.NoError(t, err)
		require.Equal(t, votes, balance1155)
		backend.Commit()

		// 제안 실행
		proposal := contracts.NewProposalToTarget(t, "TC2", 1, common.Address{1}, common.Hash{1}, "1")
		tx, err = contracts.governor.Propose(proposer, proposal.Targets, proposal.Values, proposal.CallDatas, proposal.Description)
		require.NoError(t, err)
		receipt, err = bind.WaitMined(ctx, backend, tx)
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

		proposalId := contracts.UnpackProposalCreated(t, receipt.Logs).ProposalId
		require.True(t, proposalId.Sign() != 0)

		// 투표 시간으로 이동
		contracts.ToProposalSnapshotTime(t, backend, proposalId)

		// TC2-1: 제안자는 투표할 수 없다. (자동 기권)
		t.Run("TC2-1", func(t *testing.T) {
			// 투표 실행 실패
			_, err := contracts.governor.CastVote(proposer, proposalId, VoteType.For)
			require.Error(t, err)

			// 투표 확인
			hasVotes, err := contracts.governor.HasVoted(callOpts, proposalId, proposer.From)
			require.NoError(t, err)
			require.True(t, hasVotes)
			cid, err := contracts.erc1155.CurrentID(callOpts)
			require.NoError(t, err)
			require.Equal(t, currentID, cid)
		})

		// TC2-2: 제안자를 제외한 나머지 인원은 자유롭게 투표할 수 있다. (찬성 투표)
		// 		: CastVote 에서 ERC20 을 가지고 있다면 ERC1155 를 알아서 mint 한다.
		t.Run("TC2-2", func(t *testing.T) {
			txs := make([]*types.Transaction, 0)
			for _, eoa := range eoas[1:] {
				tx, err := contracts.governor.CastVote(eoa, proposalId, VoteType.For)
				require.NoError(t, err)
				require.NotNil(t, tx)
				txs = append(txs, tx)
			}
			for _, tx := range txs {
				receipt, err := bind.WaitMined(ctx, backend, tx)
				require.NoError(t, err)
				require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
			}
			cid, err := contracts.erc1155.CurrentID(callOpts)
			require.NoError(t, err)
			require.Equal(t, currentID, cid)
			for _, eoa := range eoas {
				delegates, err := contracts.erc1155.Delegates(callOpts, eoa.From)
				require.NoError(t, err)
				require.Equal(t, eoa.From, delegates)
				votes, err := contracts.erc1155.GetVotes(callOpts, eoa.From)
				require.NoError(t, err)
				require.Equal(t, utils.ToWei(1), votes)
			}
			proposalVotes, err := contracts.governor.ProposalVotes(callOpts, proposalId)
			require.NoError(t, err)
			// 아무도 반대하지 않음
			require.True(t, proposalVotes.AgainstVotes.Sign() == 0)
			// 제안자만 기권
			require.Equal(t, utils.ToWei(1), proposalVotes.AbstainVotes)
			// 나머지는 찬성
			require.Equal(t, new(big.Int).Mul(utils.ToWei(1), big.NewInt(int64(len(eoas[1:])))), proposalVotes.ForVotes)

			// TC2-2-1: 중복 투표는 할수 없다.
			t.Run("TC2-2-1", func(t *testing.T) {
				for _, eoa := range eoas { // eoas 는 제안자를 포함하여 모두 투표를 했다.
					_, err := contracts.governor.CastVote(eoa, proposalId, VoteType.For)
					require.Error(t, err)
				}
			})
			// TC2-2-2: ERC20 이 없다면 투표할 수 없다
			t.Run("TC2-2-2", func(t *testing.T) {
				noHolder := simulated.GetEOA(t)
				_, err := contracts.governor.CastVote(noHolder, proposalId, VoteType.For)
				require.Error(t, err)
			})
			// TC2-3: 찬성된 제안은 투표 기간이 지나면 실행할 수 있다.
			// 		: 투표 기간중에는 실행할 수 없다.
			t.Run("TC2-3", func(t *testing.T) {
				// 투표 기간중에는 실행할 수 없다.
				_, err := contracts.governor.Execute(proposer, proposal.Targets, proposal.Values, proposal.CallDatas, crypto.Keccak256Hash([]byte(proposal.Description)))
				require.Error(t, err)
				// 실행 전 데이터 확인
				beforeUint, err := contracts.target.UintValue(callOpts)
				require.NoError(t, err)
				require.True(t, beforeUint.Sign() == 0)

				// 다음 투표 기간으로 이동
				contracts.NextProposalTime(t, backend)
				tx, err := contracts.governor.Execute(proposer, proposal.Targets, proposal.Values, proposal.CallDatas, crypto.Keccak256Hash([]byte(proposal.Description)))
				require.NoError(t, err)
				receipt, err := bind.WaitMined(ctx, backend, tx)
				require.NoError(t, err)
				require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

				// 실행 후 데이터 확인
				uintValue, err := contracts.target.UintValue(callOpts)
				require.NoError(t, err)
				require.Equal(t, common.Big1, uintValue)
				addrValue, err := contracts.target.AddrValue(callOpts)
				require.NoError(t, err)
				require.Equal(t, common.Address{1}, addrValue)
				b32Value, err := contracts.target.B32Value(callOpts)
				require.NoError(t, err)
				require.Equal(t, [32]byte{1}, b32Value)
				strValue, err := contracts.target.StrValue(callOpts)
				require.NoError(t, err)
				require.Equal(t, "1", strValue)

				// TC2-3-1: 한번 실행된 제안은 두번이상 실행될 수 없다.
				t.Run("TC2-3-1", func(t *testing.T) {
					_, err := contracts.governor.Execute(proposer, proposal.Targets, proposal.Values, proposal.CallDatas, crypto.Keccak256Hash([]byte(proposal.Description)))
					require.Error(t, err)
				})
			})
		})
	})

	// TC3: 대표자(delegatee)를 설정하지 않은경우, 자신을 대표자로 설정한다.
	t.Run("TC3", func(t *testing.T) {
		proposer := eoas[0]
		// 제안 실행
		proposal := contracts.NewProposalToTarget(t, "TC3", 1, common.Address{1}, common.Hash{1}, "1")
		tx, err := contracts.governor.Propose(proposer, proposal.Targets, proposal.Values, proposal.CallDatas, proposal.Description)
		require.NoError(t, err)
		receipt, err := bind.WaitMined(ctx, backend, tx)
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	})

	// TC4: 대표자(delegatee)를 다른 유저로 설정하였다면 제안할 수 없다.
	t.Run("TC4", func(t *testing.T) {
		proposer, delegatee := eoas[0], eoas[1]
		// 제안자 가 투표토큰이 있는지 확인
		balance20, err := contracts.erc20.BalanceOf(callOpts, proposer.From)
		require.NoError(t, err)
		require.Equal(t, utils.ToWei(1), balance20)
		// 투표권 획득
		tx, err := contracts.erc1155.Mint(proposer, proposer.From)
		require.NoError(t, err)
		receipt, err := bind.WaitMined(ctx, backend, tx)
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		currentID, err := contracts.erc1155.CurrentID(callOpts)
		require.NoError(t, err)
		balance1155, err := contracts.erc1155.BalanceOf(callOpts, proposer.From, currentID)
		require.NoError(t, err)
		require.Equal(t, balance20, balance1155)

		// 대표자 설정
		tx, err = contracts.erc1155.Delegate(proposer, delegatee.From)
		require.NoError(t, err)
		receipt, err = bind.WaitMined(ctx, backend, tx)
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		votes, err := contracts.erc1155.GetVotes(callOpts, proposer.From)
		require.NoError(t, err)
		require.True(t, votes.Sign() == 0)

		// 제안 실행
		proposal := contracts.NewProposalToTarget(t, "TC4", 1, common.Address{1}, common.Hash{1}, "1")
		_, err = contracts.governor.Propose(proposer, proposal.Targets, proposal.Values, proposal.CallDatas, proposal.Description)
		require.Error(t, err)
	})
}

/* TC1: 제안 찬성 후 supprot 별 잔액 확인
 *	TC1-1: 찬성 > 기권 > 반대 (state == Successed)
 *	TC1-2: 기권 > 찬성 > 반대 (state == Execute) */
/* TC2: 제안 반대 후 supprot 별 잔액 확인
 *	TC2-1: 반대 > 기권 > 찬성
 *	TC2-2: 기권 > 반대 > 찬성 */
/* TC3: 종적수 미달 (ERC1155 는 많지만 투표 참여를 안함) 은 실패
 * 	TC3-1: 투표에 참여하지 않은 유저는 claim 을 할 수 없다.*/
func TestDeployBMGovernorClaim(t *testing.T) {
	backend, contracts := DeployBMGovernorWithBackend(t)
	proposer := contracts.ChargeERC20(t, backend, []*bind.TransactOpts{backend.Owner})[0]
	proposeAndVote := func(proposal *Proposal, againsts, fors, abstains []*bind.TransactOpts) *big.Int {
		tx, err := contracts.governor.Propose(proposer, proposal.Targets, proposal.Values, proposal.CallDatas, proposal.Description)
		require.NoError(t, err)
		receipt, err := bind.WaitMined(ctx, backend, tx)
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		proposalId := contracts.UnpackProposalCreated(t, receipt.Logs).ProposalId
		require.True(t, proposalId.Sign() != 0)

		contracts.ToProposalSnapshotTime(t, backend, proposalId)

		txs := make([]*types.Transaction, 0)
		for _, eoa := range againsts {
			tx, err := contracts.governor.CastVote(eoa, proposalId, VoteType.Against)
			require.NoError(t, err)
			txs = append(txs, tx)
		}
		for _, eoa := range fors {
			tx, err := contracts.governor.CastVote(eoa, proposalId, VoteType.For)
			require.NoError(t, err)
			txs = append(txs, tx)
		}
		for _, eoa := range abstains {
			tx, err := contracts.governor.CastVote(eoa, proposalId, VoteType.Abstain)
			require.NoError(t, err)
			txs = append(txs, tx)
		}
		for _, tx := range txs {
			receipt, err := bind.WaitMined(ctx, backend, tx)
			require.NoError(t, err)
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		}
		contracts.NextProposalTime(t, backend)

		return proposalId
	}
	claim := func(proposalID *big.Int, eoas []*bind.TransactOpts) map[common.Address]*big.Int {
		txs := make([]*types.Transaction, len(eoas)+1)
		fromBlock, err := backend.BlockNumber(ctx)
		require.NoError(t, err)
		for i, eoa := range append(eoas, proposer) {
			tx, err := contracts.erc20.Claim(eoa, proposalID, eoa.From)
			require.NoError(t, err)
			txs[i] = tx
		}
		for _, tx := range txs {
			receipt, err := bind.WaitMined(ctx, backend, tx)
			require.NoError(t, err)
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		}
		toBlock, err := backend.BlockNumber(ctx)
		require.NoError(t, err)

		cctx, cancel := context.WithTimeout(ctx, 5e9)
		defer cancel()
		iter, err := contracts.erc20.FilterClaimed(&bind.FilterOpts{
			Start:   fromBlock,
			End:     &toBlock,
			Context: cctx,
		}, []*big.Int{proposalID}, nil, nil)
		require.NoError(t, err)

		result := make(map[common.Address]*big.Int)
		for iter.Next() {
			log := iter.Event
			if log.Account == proposer.From {
				require.Equal(t, log.Amount, ABSTAIN_CLAIM)
			} else {
				result[log.Account] = log.Amount
			}
		}

		return result
	}

	// TC1: 제안 찬성 후 supprot 별 잔액 확인
	t.Run("TC1", func(t *testing.T) {
		contracts.NextProposalTime(t, backend)
		// TC1-1: 찬성 > 기권 > 반대 (state == Successed)
		t.Run("TC1-1", func(t *testing.T) {
			eoas := contracts.ChargeERC20(t, backend, simulated.GetEOAs(t, 7))
			proposal := contracts.NewProposalToTarget(t, "TC1-1", "TC1-1")
			againsts, fors, abstains := eoas[:1], eoas[1:5], eoas[5:] // 1 4 2
			proposalID := proposeAndVote(proposal, againsts, fors, abstains)
			state, err := contracts.governor.State(callOpts, proposalID)
			require.NoError(t, err)
			require.Equal(t, ProposalState.Succeeded, state)
			result := claim(proposalID, eoas)
			for _, eoa := range againsts {
				require.Equal(t, LOSE_CLAIM, result[eoa.From])
			}
			for _, eoa := range fors {
				require.Equal(t, WIN_CLAIM, result[eoa.From])
			}
			for _, eoa := range abstains {
				require.Equal(t, ABSTAIN_CLAIM, result[eoa.From])
			}
		})
		// TC1-2: 기권 > 찬성 > 반대 (state == Execute)
		t.Run("TC1-2", func(t *testing.T) {
			eoas := contracts.ChargeERC20(t, backend, simulated.GetEOAs(t, 7))
			proposal := contracts.NewProposalToTarget(t, "TC1-2", "TC1-2")
			againsts, fors, abstains := eoas[:1], eoas[1:5], eoas[5:] // 1 4 2
			proposalID := proposeAndVote(proposal, againsts, fors, abstains)
			tx, err := contracts.governor.Execute(proposer, proposal.Targets, proposal.Values, proposal.CallDatas, crypto.Keccak256Hash([]byte(proposal.Description)))
			require.NoError(t, err)
			receipt, err := bind.WaitMined(ctx, backend, tx)
			require.NoError(t, err)
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

			state, err := contracts.governor.State(callOpts, proposalID)
			require.NoError(t, err)
			require.Equal(t, ProposalState.Execute, state)

			result := claim(proposalID, eoas)
			for _, eoa := range againsts {
				require.Equal(t, LOSE_CLAIM, result[eoa.From])
			}
			for _, eoa := range fors {
				require.Equal(t, WIN_CLAIM, result[eoa.From])
			}
			for _, eoa := range abstains {
				require.Equal(t, ABSTAIN_CLAIM, result[eoa.From])
			}
		})
	})

	// TC2: 제안 반대 후 supprot 별 잔액 확인
	t.Run("TC2", func(t *testing.T) {
		// TC2-1: 반대 > 기권 > 찬성
		t.Run("TC2-1", func(t *testing.T) {
			eoas := contracts.ChargeERC20(t, backend, simulated.GetEOAs(t, 7))
			proposal := contracts.NewProposalToTarget(t, "TC2-1", "TC2-1")
			againsts, fors, abstains := eoas[:4], eoas[4:5], eoas[5:] // 4 1 2
			proposalID := proposeAndVote(proposal, againsts, fors, abstains)
			state, err := contracts.governor.State(callOpts, proposalID)
			require.NoError(t, err)
			require.Equal(t, ProposalState.Defeated, state)
			result := claim(proposalID, eoas)
			for _, eoa := range againsts {
				require.Equal(t, WIN_CLAIM, result[eoa.From])
			}
			for _, eoa := range fors {
				require.Equal(t, LOSE_CLAIM, result[eoa.From])
			}
			for _, eoa := range abstains {
				require.Equal(t, ABSTAIN_CLAIM, result[eoa.From])
			}
		})
		// TC2-2: 기권 > 반대 > 찬성
		t.Run("TC2-2", func(t *testing.T) {
			eoas := contracts.ChargeERC20(t, backend, simulated.GetEOAs(t, 7))
			proposal := contracts.NewProposalToTarget(t, "TC2-2", "TC2-2")
			againsts, fors, abstains := eoas[:2], eoas[2:3], eoas[3:] // 2 1 4
			proposalID := proposeAndVote(proposal, againsts, fors, abstains)
			state, err := contracts.governor.State(callOpts, proposalID)
			require.NoError(t, err)
			require.Equal(t, ProposalState.Defeated, state)
			result := claim(proposalID, eoas)
			for _, eoa := range againsts {
				require.Equal(t, WIN_CLAIM, result[eoa.From])
			}
			for _, eoa := range fors {
				require.Equal(t, LOSE_CLAIM, result[eoa.From])
			}
			for _, eoa := range abstains {
				require.Equal(t, ABSTAIN_CLAIM, result[eoa.From])
			}
		})
	})
	// TC3: 종적수 미달 (ERC1155 는 많지만 투표 참여를 안함) 은 실패
	t.Run("TC3", func(t *testing.T) {
		eoas := contracts.ChargeERC20(t, backend, simulated.GetEOAs(t, 50))
		// mint erc1155
		txs := make([]*types.Transaction, len(eoas))
		for i, eoa := range eoas {
			tx, err := contracts.erc1155.Mint(eoa, eoa.From)
			require.NoError(t, err)
			txs[i] = tx
		}
		for _, tx := range txs {
			receipt, err := bind.WaitMined(ctx, backend, tx)
			require.NoError(t, err)
			require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		}

		proposal := contracts.NewProposalToTarget(t, "TC3", "TC3")
		// 50 명 중에 4(3 + proposer)명만 투표 (종적수는 10% 이기 때문에 5명 이상 투표를 했어야 한다.)
		againsts, fors, abstains := eoas[:1], eoas[1:2], eoas[2:3]
		proposalID := proposeAndVote(proposal, againsts, fors, abstains)
		state, err := contracts.governor.State(callOpts, proposalID)
		require.NoError(t, err)
		// 종적수 미달 ==> 제안 실패
		require.Equal(t, ProposalState.Defeated, state)
		result := claim(proposalID, eoas[:3])
		for _, eoa := range againsts {
			require.Equal(t, WIN_CLAIM, result[eoa.From])
		}
		for _, eoa := range fors {
			require.Equal(t, LOSE_CLAIM, result[eoa.From])
		}
		for _, eoa := range abstains {
			require.Equal(t, ABSTAIN_CLAIM, result[eoa.From])
		}

		// TC3-1: 투표에 참여하지 않은 유저는 claim 을 할 수 없다.
		t.Run("TC3-1", func(t *testing.T) {
			for _, eoa := range eoas[3:] {
				_, err := contracts.erc20.Claim(eoa, proposalID, eoa.From)
				require.Error(t, err)
			}
		})
	})
}
