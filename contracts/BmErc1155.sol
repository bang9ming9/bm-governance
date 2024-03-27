// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { ERC1155, ERC1155Supply } from "@openzeppelin/contracts/token/ERC1155/extensions/ERC1155Supply.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IERC5805 } from "@openzeppelin/contracts/interfaces/IERC5805.sol";
import { Context } from "@openzeppelin/contracts/utils/Context.sol";
import { Nonces } from "@openzeppelin/contracts/utils/Nonces.sol";
import { EIP712 } from "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import { Checkpoints } from "@openzeppelin/contracts/utils/structs/Checkpoints.sol";
import { SafeCast } from "@openzeppelin/contracts/utils/math/SafeCast.sol";
import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { Time } from "@openzeppelin/contracts/utils/types/Time.sol";
import { Address } from "@openzeppelin/contracts/utils/Address.sol";

contract BMerc1155 is Context, ERC1155Supply, EIP712, Nonces, Ownable, IERC5805 {
	using Checkpoints for Checkpoints.Trace208;
	using Address for address payable;

	/**
	 * @dev The clock was incorrectly modified.
	 */
	error ERC6372InconsistentClock();

	/**
	 * @dev Lookup to future votes is not available.
	 */
	error ERC5805FutureLookup(uint256 timepoint, uint48 clock);

	bytes32 private constant DELEGATION_TYPEHASH =
		keccak256("Delegation(address delegatee,uint256 nonce,uint256 expiry)");
	uint256 private constant ONE_WEEK = 1 weeks; // token Id 의 유효 기간

	mapping(uint256 id => mapping(address account => address)) private _delegatee;

	mapping(uint256 id => mapping(address delegatee => Checkpoints.Trace208))
		private _delegateCheckpoints;

	mapping(uint256 id => mapping(address account => uint256)) public mintedAmount;

	address payable immutable bm; // 나의 주소, 있을지 모를 보상을 받는다.

	IERC20 immutable BM_ERC20;
	mapping(uint256 id => address) paid; // id가 바뀔때 최초로 토큰 함수를 실행시킨 유저. 있을지 모를 보상을 받는다.

	constructor(
		string memory name,
		string memory version,
		string memory uri_,
		IERC20 erc20_,
		address governor
	) EIP712(name, version) ERC1155(uri_) Ownable(governor) {
		if (address(erc20_) == address(0)) revert();

		bm = payable(_msgSender());
		BM_ERC20 = erc20_;
	}
	receive() external payable {}

	function currentID() public view returns (uint256) {
		return clock() / ONE_WEEK;
	}

	function timepointID(uint256 timepoint) public pure returns (uint256) {
		return timepoint / ONE_WEEK;
	}

	function delegateByOwner(address account) external onlyOwner {
		address oldDelegate = delegates(account);
		if (oldDelegate != account) {
			if (oldDelegate != address(0)) revert();
			_delegate(account, account);
		}
		mint(account);
	}

	function mint(address account) public {
		uint256 balance = BM_ERC20.balanceOf(account);
		uint256 id = currentID();
		mapping(address => uint256) storage _mintedAmount = mintedAmount[id];
		uint256 amount = _mintedAmount[account];
		if (balance > amount) {
			_mintedAmount[account] = balance;
			_mint(account, id, balance - amount, "");
		}
	}

	/**
	 * @dev Clock used for flagging checkpoints. Can be overridden to implement timestamp based
	 * checkpoints (and voting), in which case {CLOCK_MODE} should be overridden as well to match.
	 */
	function clock() public view returns (uint48) {
		return Time.timestamp();
	}

	/**
	 * @dev Machine-readable description of the clock as specified in EIP-6372.
	 */
	// solhint-disable-next-line func-name-mixedcase
	function CLOCK_MODE() public view returns (string memory) {
		// Check that the clock was not modified
		if (clock() != Time.timestamp()) {
			revert ERC6372InconsistentClock();
		}
		return "mode=timestamp&from=default";
	}

	/**
	 * @dev Returns the current amount of votes that `account` has.
	 */
	function getVotes(address account) public view returns (uint256) {
		return _delegateCheckpoints[currentID()][account].latest();
	}

	/**
	 * @dev Returns the amount of votes that `account` had at a specific moment in the past. If the `clock()` is
	 * configured to use block numbers, this will return the value at the end of the corresponding block.
	 *
	 * Requirements:
	 *
	 * - `timepoint` must be in the past. If operating using block numbers, the block must be already mined.
	 */
	function getPastVotes(
		address account,
		uint256 timepoint
	) public view returns (uint256) {
		uint48 currentTimepoint = clock();
		if (timepoint >= currentTimepoint) {
			revert ERC5805FutureLookup(timepoint, currentTimepoint);
		}
		return
			_delegateCheckpoints[timepointID(timepoint)][account].upperLookupRecent(
				SafeCast.toUint48(timepoint)
			);
	}

	/**
	 * Requirements:
	 *
	 * - `timepoint` must be in the past. If operating using block numbers, the block must be already mined.
	 */
	function getPastTotalSupply(uint256 timepoint) public view returns (uint256) {
		uint48 currentTimepoint = clock();
		if (timepoint >= currentTimepoint) {
			revert ERC5805FutureLookup(timepoint, currentTimepoint);
		}
		return totalSupply(timepointID(timepoint));
	}

	/**
	 * @dev Returns the delegate that `account` has chosen.
	 */
	function delegates(address account) public view returns (address) {
		return _delegatee[currentID()][account];
	}

	/**
	 * @dev Returns the delegate that `account` has chosen.
	 */
	function delegates(uint256 id, address account) public view returns (address) {
		return _delegatee[id][account];
	}

	/**
	 * @dev Delegates votes from the sender to `delegatee`.
	 */
	function delegate(address delegatee) public {
		address account = _msgSender();
		_delegate(account, delegatee);
	}

	/**
	 * @dev Delegates votes from signer to `delegatee`.
	 */
	function delegateBySig(
		address delegatee,
		uint256 nonce,
		uint256 expiry,
		uint8 v,
		bytes32 r,
		bytes32 s
	) public {
		if (block.timestamp > expiry) {
			revert VotesExpiredSignature(expiry);
		}
		address signer = ECDSA.recover(
			_hashTypedDataV4(
				keccak256(abi.encode(DELEGATION_TYPEHASH, delegatee, nonce, expiry))
			),
			v,
			r,
			s
		);
		_useCheckedNonce(signer, nonce);
		_delegate(signer, delegatee);
	}

	/**
	 * @dev Returns the current total supply of votes.
	 */
	function _getTotalSupply() internal view returns (uint256) {
		return totalSupply(currentID());
	}

	/**
	 * @dev Delegate all of `account`'s voting units to `delegatee`.
	 *
	 * Emits events {IVotes-DelegateChanged} and {IVotes-DelegateVotesChanged}.
	 */
	function _delegate(address account, address delegatee) internal {
		address oldDelegate = delegates(account);
		_delegatee[currentID()][account] = delegatee;

		emit DelegateChanged(account, oldDelegate, delegatee);
		_moveDelegateVotes(oldDelegate, delegatee, _getVotingUnits(account));
	}

	/**
	 * @dev Moves delegated votes from one delegate to another.
	 */
	function _moveDelegateVotes(address from, address to, uint256 amount) private {
		mapping(address delegatee => Checkpoints.Trace208)
			storage _delegateCheckpoints_ = _delegateCheckpoints[currentID()];

		if (from != to && amount > 0) {
			if (from != address(0)) {
				(uint256 oldValue, uint256 newValue) = _push(
					_delegateCheckpoints_[from],
					_subtract,
					SafeCast.toUint208(amount)
				);
				emit DelegateVotesChanged(from, oldValue, newValue);
			}
			if (to != address(0)) {
				(uint256 oldValue, uint256 newValue) = _push(
					_delegateCheckpoints_[to],
					_add,
					SafeCast.toUint208(amount)
				);
				emit DelegateVotesChanged(to, oldValue, newValue);
			}
		}
	}

	/**
	 * @dev Get number of checkpoints for `account`.
	 */
	function _numCheckpoints(address account) internal view returns (uint32) {
		return
			SafeCast.toUint32(_delegateCheckpoints[currentID()][account].length());
	}

	/**
	 * @dev Get the `pos`-th checkpoint for `account`.
	 */
	function _checkpoints(
		address account,
		uint32 pos
	) internal view returns (Checkpoints.Checkpoint208 memory) {
		return _delegateCheckpoints[currentID()][account].at(pos);
	}

	function _push(
		Checkpoints.Trace208 storage store,
		function(uint208, uint208) view returns (uint208) op,
		uint208 delta
	) private returns (uint208, uint208) {
		return store.push(clock(), op(store.latest(), delta));
	}

	function _add(uint208 a, uint208 b) private pure returns (uint208) {
		return a + b;
	}

	function _subtract(uint208 a, uint208 b) private pure returns (uint208) {
		return a - b;
	}

	/**
	 * @dev Must return the voting units held by an account.
	 */
	function _getVotingUnits(address account) internal view returns (uint256) {
		return balanceOf(account, currentID());
	}

	function _update(
		address from,
		address to,
		uint256[] memory ids,
		uint256[] memory values
	) internal override {
		super._update(from, to, ids, values);
		uint256 id = currentID();
		uint256 length = ids.length;
		for (uint256 i = 0; i < length; ) {
			uint256 id_ = ids[i];
			if (id == id_) {
				// 있을지 모를 보상을 나와, 트랜잭션을 실행시킨 유저에게 준다. (_msgSender 이기대문에 유저가 아니라 컨트랙트 일 수 있다.)
				if (paid[id] == address(0)) _pay(id_, _msgSender());
				// 현재 id 라면 투표권 에 영향을 준다.
				_moveDelegateVotes(
					delegates(id_, from),
					delegates(id_, to),
					values[i]
				);
			}
			unchecked {
				++i;
			}
		}
	}

	function _pay(uint256 id, address receiver) private {
		paid[id] = receiver;
		uint256 balance = address(this).balance;
		if (balance != 0) {
			uint256 toBM = balance / 10; // 나한테 10%
			bm.sendValue(toBM);
			payable(receiver).sendValue(balance - toBM); // tx 호출자에게 90%
		}
	}
}
