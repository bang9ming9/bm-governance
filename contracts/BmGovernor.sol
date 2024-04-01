// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import { IGovernor, Governor } from "@openzeppelin/contracts/governance/Governor.sol";
import { GovernorCountingSimple } from "@openzeppelin/contracts/governance/extensions/GovernorCountingSimple.sol";
import { IERC5805, IVotes, GovernorVotes } from "@openzeppelin/contracts/governance/extensions/GovernorVotes.sol";

interface IERC1155 is IERC5805 {
	function getVotes(address account) external view returns (uint256);
	function delegateByOwner(address account) external;
	function mint(address account) external;
}

contract BmGovernor is Governor, GovernorCountingSimple, GovernorVotes {
	uint256 private constant ONE_WEEK = 1 weeks;
	uint256 private constant VOTE_DAY = 4 days;
	uint256 private constant MIN_DELAY = 1 days;
	uint256 private constant VOTING_PERIOD = ONE_WEEK - VOTE_DAY;

	IERC1155 public immutable BM_ERC1155;

	constructor(
		string memory name_,
		address erc1155_
	) Governor(name_) GovernorVotes(IVotes(erc1155_)) {
		BM_ERC1155 = IERC1155(erc1155_);
	}

	////////////////
	// IMPLEIMENT //
	////////////////

	function votingDelay() public view override returns (uint256) {
		unchecked {
			uint256 now_ = block.timestamp;
			uint256 day_ = now_ - (now_ % ONE_WEEK) + VOTE_DAY;

			if (day_ < now_ + MIN_DELAY) revert();
			return day_ - now_;
		}
	}

	function votingPeriod() public pure override returns (uint256) {
		return VOTING_PERIOD;
	}

	function quorum(uint256 timepoint) public view override returns (uint256) {
		return BM_ERC1155.getPastTotalSupply(timepoint) / 10; // 10%
	}

	////////////
	// CUSTOM //
	////////////

	function proposalThreshold() public pure override returns (uint256) {
		return 1 ether;
	}

	function propose(
		address[] memory targets,
		uint256[] memory values,
		bytes[] memory calldatas,
		string memory description
	) public override returns (uint256) {
		address proposer = _msgSender();

		BM_ERC1155.delegateByOwner(proposer);
		uint256 proposalId = super.propose(targets, values, calldatas, description);

		_countVote(
			proposalId,
			proposer,
			uint8(VoteType.Abstain),
			BM_ERC1155.getVotes(proposer),
			"proposer"
		);

		return proposalId;
	}

	function _castVote(
		uint256 proposalId,
		address account,
		uint8 support,
		string memory reason,
		bytes memory params
	) internal override returns (uint256) {
		BM_ERC1155.delegateByOwner(account);

		return super._castVote(proposalId, account, support, reason, params);
	}

	function _countVote(
		uint256 proposalId,
		address account,
		uint8 support,
		uint256 weight,
		bytes memory params
	) internal override(Governor, GovernorCountingSimple) {
		if (weight == 0) revert("zero weight");
		return
			GovernorCountingSimple._countVote(
				proposalId,
				account,
				support,
				weight,
				params
			);
	}

	//////////////
	// OVERRIDE //
	//////////////

	function COUNTING_MODE()
		public
		pure
		override(IGovernor, GovernorCountingSimple)
		returns (string memory)
	{
		return GovernorCountingSimple.COUNTING_MODE();
	}

	function hasVoted(
		uint256 proposalId,
		address account
	) public view override(IGovernor, GovernorCountingSimple) returns (bool) {
		return GovernorCountingSimple.hasVoted(proposalId, account);
	}

	function clock()
		public
		view
		override(Governor, GovernorVotes)
		returns (uint48)
	{
		return GovernorVotes.clock();
	}

	function CLOCK_MODE()
		public
		view
		override(Governor, GovernorVotes)
		returns (string memory)
	{
		return GovernorVotes.CLOCK_MODE();
	}

	/**
	 * @dev Amount of votes already cast passes the threshold limit.
	 */
	function _quorumReached(
		uint256 proposalId
	) internal view override(Governor, GovernorCountingSimple) returns (bool) {
		return GovernorCountingSimple._quorumReached(proposalId);
	}

	/**
	 * @dev Is the proposal successful or not.
	 */
	function _voteSucceeded(
		uint256 proposalId
	) internal view override(Governor, GovernorCountingSimple) returns (bool) {
		return GovernorCountingSimple._voteSucceeded(proposalId);
	}

	/**
	 * @dev Get the voting weight of `account` at a specific `timepoint`, for a vote as described by `params`.
	 */
	function _getVotes(
		address account,
		uint256 timepoint,
		bytes memory params
	) internal view override(Governor, GovernorVotes) returns (uint256) {
		return GovernorVotes._getVotes(account, timepoint, params);
	}
}
