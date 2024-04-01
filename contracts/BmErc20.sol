// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { ERC20Capped } from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Capped.sol";
import { ERC20Burnable } from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import { Address } from "@openzeppelin/contracts/utils/Address.sol";

contract BmErc20 is ERC20, ERC20Capped, ERC20Burnable {
	using Address for address payable;

	uint256 public constant COST = 1 ether;
	address payable public immutable BM_ERC1155;
	mapping(address account => bool) public minted;

	constructor(
		string memory name,
		string memory symbol,
		address payable erc1155
	) ERC20(name, symbol) ERC20Capped(type(uint208).max) {
		BM_ERC1155 = erc1155;
	}

	receive() external payable {
		mint(_msgSender());
	}

	fallback() external {
		assert(address(this).balance == 0);
	}

	function mint(address account) public payable {
		if (minted[account]) revert("1");
		minted[account] = true;

		BM_ERC1155.sendValue(COST);
		_mint(account, COST);
	}

	function _update(
		address from,
		address to,
		uint256 value
	) internal virtual override(ERC20, ERC20Capped) {
		super._update(from, to, value);
	}
}
