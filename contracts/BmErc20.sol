// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { ERC20Capped } from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Capped.sol";
import { ERC20Burnable } from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import { Address } from "@openzeppelin/contracts/utils/Address.sol";
import { IGovernor as OZIGovernor } from "@openzeppelin/contracts/governance/IGovernor.sol";

interface IGovernor is OZIGovernor {
	function proposalStateToClaim(
		uint256 proposalID,
		address account
	)
		external
		view
		returns (bool hasVoted, uint8 supported, OZIGovernor.ProposalState state);
}

contract BmErc20 is ERC20, ERC20Capped, ERC20Burnable {
	using Address for address payable;

	error BmErc20NilInput(string arg);
	error BmErc20AlreadyMintedUser(address account);
	error BmErc20AlreadyClaimed(uint256 proposalID, address account);
	error BmErc20NotVoteUser(uint256 proposalID, address account);
	error BmErc20InvalidProposalState(uint256 proposalID, uint8 state);

	event Minted(address indexed account);
	event Claimed(
		uint256 indexed proposalID,
		address indexed account,
		uint256 indexed amount // WIN_CLAIM,LOSE_CLAIM,ABSTAIN_CLAIM 중에 1개이기 때문에 검색 가능 하다
	);

	uint256 public constant COST = 1 ether;
	uint256 public constant WIN_CLAIM = 0.5 ether;
	uint256 public constant LOSE_CLAIM = 0.1 ether;
	uint256 public constant ABSTAIN_CLAIM = 0.2 ether;

	address payable public immutable BM_ERC1155;
	IGovernor public immutable BM_GOVERNOR;

	mapping(address account => bool) public minted;
	mapping(uint256 proposalID => mapping(address account => bool)) public claimed;

	constructor(
		string memory name,
		string memory symbol,
		address payable erc1155,
		address governor
	) ERC20(name, symbol) ERC20Capped(type(uint208).max) {
		if (erc1155 == address(0)) revert BmErc20NilInput("erc1155");
		if (governor == address(0)) revert BmErc20NilInput("governor");
		BM_ERC1155 = erc1155;
		BM_GOVERNOR = IGovernor(governor);
	}

	receive() external payable {
		mint(_msgSender());
	}

	fallback() external {
		assert(address(this).balance == 0);
	}

	function mint(address account) public payable {
		if (minted[account]) revert BmErc20AlreadyMintedUser(account);
		minted[account] = true;

		BM_ERC1155.sendValue(COST);
		_mint(account, COST);
		emit Minted(account);
	}

	function claim(uint256 proposalID, address account) external {
		// 1. 중복 보상 확인
		if (claimed[proposalID][account])
			revert BmErc20AlreadyClaimed(proposalID, account);
		(
			bool hasVoted,
			uint8 supported,
			OZIGovernor.ProposalState state
		) = BM_GOVERNOR.proposalStateToClaim(proposalID, account);
		// 2. 투표한 유저인지 확인
		if (!hasVoted) revert BmErc20NotVoteUser(proposalID, account);
		// 3. 결과에 따른 지급 보상
		uint256 claimAmount;
		/*
		 *   enum VoteType {
		 *		Against, 	= 0
		 *		For, 		= 1
		 *		Abstain		= 2
		 *	}
		 */

		// 3-1. 기권에 투표한 경우
		if (supported == 2) {
			claimAmount = ABSTAIN_CLAIM;
		} else {
			if (
				state == OZIGovernor.ProposalState.Executed ||
				state == OZIGovernor.ProposalState.Succeeded
			) {
				// 3-2 제안이 성공한 경우 For 가 위너
				claimAmount = supported == 1 ? WIN_CLAIM : LOSE_CLAIM;
			} else if (state == OZIGovernor.ProposalState.Defeated) {
				// 3-3. 제안이 실패한 경우 Against 가 위너
				claimAmount = supported == 0 ? WIN_CLAIM : LOSE_CLAIM;
			} else {
				// 이외의 상태는 정료 처리
				revert BmErc20InvalidProposalState(proposalID, uint8(state));
			}
		}
		// 4. 보상 지급
		_mint(account, claimAmount);
		claimed[proposalID][account] = true;

		emit Claimed(proposalID, account, claimAmount);
	}

	function _update(
		address from,
		address to,
		uint256 value
	) internal virtual override(ERC20, ERC20Capped) {
		super._update(from, to, value);
	}
}
