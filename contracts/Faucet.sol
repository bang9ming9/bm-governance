// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import { Address } from "@openzeppelin/contracts/utils/Address.sol";

contract Faucet {
	using Address for address payable;

	uint256 public constant AMOUNT = 10 ether;
	uint256 public constant FAUCET_INTERVAL = 1 days;

	error AlreadyClaimed(address);
	event Claimed(address indexed account);

	mapping(address account => uint256) private _claimed;

	receive() external payable {}

	function claim() external {
		address account = msg.sender;
		uint256 time = block.timestamp / FAUCET_INTERVAL;

		if (_claimed[account] == time) revert AlreadyClaimed(account);
		_claimed[account] = time;

		payable(account).sendValue(AMOUNT);
		emit Claimed(account);
	}
}
