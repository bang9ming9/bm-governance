// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";

contract TargetContract is Ownable {
	uint256 public uintValue;
	address public addrValue;
	bytes32 public b32Value;
	string public strValue;

	event Uint256Written(uint256 indexed value);
	event AddressWritten(address indexed value);
	event Bytes32Written(bytes32 indexed value);
	event StringWritten(string indexed value);

	constructor(address governor) Ownable(governor) {}

	function writeUintValue(uint256 _value) external onlyOwner {
		uintValue = _value;
		emit Uint256Written(_value);
	}

	function writeAddrValue(address _value) external onlyOwner {
		addrValue = _value;
		emit AddressWritten(_value);
	}

	function writeB32Value(bytes32 _value) external onlyOwner {
		b32Value = _value;
		emit Bytes32Written(_value);
	}

	function writeStrValue(string memory _value) external onlyOwner {
		strValue = _value;
		emit StringWritten(_value);
	}
}
