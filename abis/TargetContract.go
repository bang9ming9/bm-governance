// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TargetContractMetaData contains all meta data concerning the TargetContract contract.
var TargetContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"AddressWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"Bytes32Written\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"StringWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256Written\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addrValue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"b32Value\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strValue\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uintValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_value\",\"type\":\"address\"}],\"name\":\"writeAddrValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"writeB32Value\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"writeStrValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"writeUintValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516107df3803806107df83398101604081905261002e916100bb565b806001600160a01b03811661005c57604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b6100658161006c565b50506100e8565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602082840312156100cb575f80fd5b81516001600160a01b03811681146100e1575f80fd5b9392505050565b6106ea806100f55f395ff3fe608060405234801561000f575f80fd5b50600436106100a6575f3560e01c8063988817b01161006e578063988817b01461011c578063a5f1038914610133578063ba1434e714610146578063c567fecd14610159578063eb27590d1461016c578063f2fde38b14610175575f80fd5b80634ccf37a0146100aa57806368322be7146100da578063715018a6146100ef5780638b26be0d146100f95780638da5cb5b1461010c575b5f80fd5b6002546100bd906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100e2610188565b6040516100d19190610420565b6100f7610214565b005b6100f7610107366004610466565b610227565b5f546001600160a01b03166100bd565b61012560015481565b6040519081526020016100d1565b6100f7610141366004610511565b61027c565b6100f761015436600461053e565b6102cd565b6100f761016736600461053e565b610307565b61012560035481565b6100f7610183366004610511565b610341565b6004805461019590610555565b80601f01602080910402602001604051908101604052809291908181526020018280546101c190610555565b801561020c5780601f106101e35761010080835404028352916020019161020c565b820191905f5260205f20905b8154815290600101906020018083116101ef57829003601f168201915b505050505081565b61021c610383565b6102255f6103af565b565b61022f610383565b600461023b82826105d9565b508060405161024a9190610699565b604051908190038120907f6ac728ecdca9efb5f0dec54e77bef32812b52dde84f40fb0d0bbe0cb2cf49a5a905f90a250565b610284610383565b600280546001600160a01b0319166001600160a01b0383169081179091556040517f22b522eb67a8c1256002bbca0d50bcecf337e109bb750f63c301bade4bc198f0905f90a250565b6102d5610383565b600381905560405181907fb20eff94cd1e4024682891d638cc33ad15661f4b9de7ef517dce69ff7392b952905f90a250565b61030f610383565b600181905560405181907f6370c32f0eb3cd69b66de366ebcfd36200df04ca19510c44f8159824ce32d159905f90a250565b610349610383565b6001600160a01b03811661037757604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b610380816103af565b50565b5f546001600160a01b031633146102255760405163118cdaa760e01b815233600482015260240161036e565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f5b83811015610418578181015183820152602001610400565b50505f910152565b602081525f825180602084015261043e8160408501602087016103fe565b601f01601f19169190910160400192915050565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215610476575f80fd5b813567ffffffffffffffff8082111561048d575f80fd5b818401915084601f8301126104a0575f80fd5b8135818111156104b2576104b2610452565b604051601f8201601f19908116603f011681019083821181831017156104da576104da610452565b816040528281528760208487010111156104f2575f80fd5b826020860160208301375f928101602001929092525095945050505050565b5f60208284031215610521575f80fd5b81356001600160a01b0381168114610537575f80fd5b9392505050565b5f6020828403121561054e575f80fd5b5035919050565b600181811c9082168061056957607f821691505b60208210810361058757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156105d457805f5260205f20601f840160051c810160208510156105b25750805b601f840160051c820191505b818110156105d1575f81556001016105be565b50505b505050565b815167ffffffffffffffff8111156105f3576105f3610452565b610607816106018454610555565b8461058d565b602080601f83116001811461063a575f84156106235750858301515b5f19600386901b1c1916600185901b178555610691565b5f85815260208120601f198616915b8281101561066857888601518255948401946001909101908401610649565b508582101561068557878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f82516106aa8184602087016103fe565b919091019291505056fea26469706673582212209a9eaf6c21184317210cd62fe3029839d7d1f7de3dcb12b5277bb5e9dddfbac764736f6c63430008180033",
}

// TargetContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TargetContractMetaData.ABI instead.
var TargetContractABI = TargetContractMetaData.ABI

// TargetContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TargetContractMetaData.Bin instead.
var TargetContractBin = TargetContractMetaData.Bin

// DeployTargetContract deploys a new Ethereum contract, binding an instance of TargetContract to it.
func DeployTargetContract(auth *bind.TransactOpts, backend bind.ContractBackend, governor common.Address) (common.Address, *types.Transaction, *TargetContract, error) {
	parsed, err := TargetContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TargetContractBin), backend, governor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TargetContract{TargetContractCaller: TargetContractCaller{contract: contract}, TargetContractTransactor: TargetContractTransactor{contract: contract}, TargetContractFilterer: TargetContractFilterer{contract: contract}}, nil
}

// TargetContract is an auto generated Go binding around an Ethereum contract.
type TargetContract struct {
	TargetContractCaller     // Read-only binding to the contract
	TargetContractTransactor // Write-only binding to the contract
	TargetContractFilterer   // Log filterer for contract events
}

// TargetContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TargetContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TargetContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TargetContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TargetContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TargetContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TargetContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TargetContractSession struct {
	Contract     *TargetContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TargetContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TargetContractCallerSession struct {
	Contract *TargetContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TargetContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TargetContractTransactorSession struct {
	Contract     *TargetContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TargetContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TargetContractRaw struct {
	Contract *TargetContract // Generic contract binding to access the raw methods on
}

// TargetContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TargetContractCallerRaw struct {
	Contract *TargetContractCaller // Generic read-only contract binding to access the raw methods on
}

// TargetContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TargetContractTransactorRaw struct {
	Contract *TargetContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTargetContract creates a new instance of TargetContract, bound to a specific deployed contract.
func NewTargetContract(address common.Address, backend bind.ContractBackend) (*TargetContract, error) {
	contract, err := bindTargetContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TargetContract{TargetContractCaller: TargetContractCaller{contract: contract}, TargetContractTransactor: TargetContractTransactor{contract: contract}, TargetContractFilterer: TargetContractFilterer{contract: contract}}, nil
}

// NewTargetContractCaller creates a new read-only instance of TargetContract, bound to a specific deployed contract.
func NewTargetContractCaller(address common.Address, caller bind.ContractCaller) (*TargetContractCaller, error) {
	contract, err := bindTargetContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TargetContractCaller{contract: contract}, nil
}

// NewTargetContractTransactor creates a new write-only instance of TargetContract, bound to a specific deployed contract.
func NewTargetContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TargetContractTransactor, error) {
	contract, err := bindTargetContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TargetContractTransactor{contract: contract}, nil
}

// NewTargetContractFilterer creates a new log filterer instance of TargetContract, bound to a specific deployed contract.
func NewTargetContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TargetContractFilterer, error) {
	contract, err := bindTargetContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TargetContractFilterer{contract: contract}, nil
}

// bindTargetContract binds a generic wrapper to an already deployed contract.
func bindTargetContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TargetContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TargetContract *TargetContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TargetContract.Contract.TargetContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TargetContract *TargetContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TargetContract.Contract.TargetContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TargetContract *TargetContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TargetContract.Contract.TargetContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TargetContract *TargetContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TargetContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TargetContract *TargetContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TargetContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TargetContract *TargetContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TargetContract.Contract.contract.Transact(opts, method, params...)
}

// AddrValue is a free data retrieval call binding the contract method 0x4ccf37a0.
//
// Solidity: function addrValue() view returns(address)
func (_TargetContract *TargetContractCaller) AddrValue(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TargetContract.contract.Call(opts, &out, "addrValue")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddrValue is a free data retrieval call binding the contract method 0x4ccf37a0.
//
// Solidity: function addrValue() view returns(address)
func (_TargetContract *TargetContractSession) AddrValue() (common.Address, error) {
	return _TargetContract.Contract.AddrValue(&_TargetContract.CallOpts)
}

// AddrValue is a free data retrieval call binding the contract method 0x4ccf37a0.
//
// Solidity: function addrValue() view returns(address)
func (_TargetContract *TargetContractCallerSession) AddrValue() (common.Address, error) {
	return _TargetContract.Contract.AddrValue(&_TargetContract.CallOpts)
}

// B32Value is a free data retrieval call binding the contract method 0xeb27590d.
//
// Solidity: function b32Value() view returns(bytes32)
func (_TargetContract *TargetContractCaller) B32Value(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TargetContract.contract.Call(opts, &out, "b32Value")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// B32Value is a free data retrieval call binding the contract method 0xeb27590d.
//
// Solidity: function b32Value() view returns(bytes32)
func (_TargetContract *TargetContractSession) B32Value() ([32]byte, error) {
	return _TargetContract.Contract.B32Value(&_TargetContract.CallOpts)
}

// B32Value is a free data retrieval call binding the contract method 0xeb27590d.
//
// Solidity: function b32Value() view returns(bytes32)
func (_TargetContract *TargetContractCallerSession) B32Value() ([32]byte, error) {
	return _TargetContract.Contract.B32Value(&_TargetContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TargetContract *TargetContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TargetContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TargetContract *TargetContractSession) Owner() (common.Address, error) {
	return _TargetContract.Contract.Owner(&_TargetContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TargetContract *TargetContractCallerSession) Owner() (common.Address, error) {
	return _TargetContract.Contract.Owner(&_TargetContract.CallOpts)
}

// StrValue is a free data retrieval call binding the contract method 0x68322be7.
//
// Solidity: function strValue() view returns(string)
func (_TargetContract *TargetContractCaller) StrValue(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TargetContract.contract.Call(opts, &out, "strValue")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StrValue is a free data retrieval call binding the contract method 0x68322be7.
//
// Solidity: function strValue() view returns(string)
func (_TargetContract *TargetContractSession) StrValue() (string, error) {
	return _TargetContract.Contract.StrValue(&_TargetContract.CallOpts)
}

// StrValue is a free data retrieval call binding the contract method 0x68322be7.
//
// Solidity: function strValue() view returns(string)
func (_TargetContract *TargetContractCallerSession) StrValue() (string, error) {
	return _TargetContract.Contract.StrValue(&_TargetContract.CallOpts)
}

// UintValue is a free data retrieval call binding the contract method 0x988817b0.
//
// Solidity: function uintValue() view returns(uint256)
func (_TargetContract *TargetContractCaller) UintValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TargetContract.contract.Call(opts, &out, "uintValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UintValue is a free data retrieval call binding the contract method 0x988817b0.
//
// Solidity: function uintValue() view returns(uint256)
func (_TargetContract *TargetContractSession) UintValue() (*big.Int, error) {
	return _TargetContract.Contract.UintValue(&_TargetContract.CallOpts)
}

// UintValue is a free data retrieval call binding the contract method 0x988817b0.
//
// Solidity: function uintValue() view returns(uint256)
func (_TargetContract *TargetContractCallerSession) UintValue() (*big.Int, error) {
	return _TargetContract.Contract.UintValue(&_TargetContract.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TargetContract *TargetContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TargetContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TargetContract *TargetContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _TargetContract.Contract.RenounceOwnership(&_TargetContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TargetContract *TargetContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TargetContract.Contract.RenounceOwnership(&_TargetContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TargetContract *TargetContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TargetContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TargetContract *TargetContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TargetContract.Contract.TransferOwnership(&_TargetContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TargetContract *TargetContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TargetContract.Contract.TransferOwnership(&_TargetContract.TransactOpts, newOwner)
}

// WriteAddrValue is a paid mutator transaction binding the contract method 0xa5f10389.
//
// Solidity: function writeAddrValue(address _value) returns()
func (_TargetContract *TargetContractTransactor) WriteAddrValue(opts *bind.TransactOpts, _value common.Address) (*types.Transaction, error) {
	return _TargetContract.contract.Transact(opts, "writeAddrValue", _value)
}

// WriteAddrValue is a paid mutator transaction binding the contract method 0xa5f10389.
//
// Solidity: function writeAddrValue(address _value) returns()
func (_TargetContract *TargetContractSession) WriteAddrValue(_value common.Address) (*types.Transaction, error) {
	return _TargetContract.Contract.WriteAddrValue(&_TargetContract.TransactOpts, _value)
}

// WriteAddrValue is a paid mutator transaction binding the contract method 0xa5f10389.
//
// Solidity: function writeAddrValue(address _value) returns()
func (_TargetContract *TargetContractTransactorSession) WriteAddrValue(_value common.Address) (*types.Transaction, error) {
	return _TargetContract.Contract.WriteAddrValue(&_TargetContract.TransactOpts, _value)
}

// WriteB32Value is a paid mutator transaction binding the contract method 0xba1434e7.
//
// Solidity: function writeB32Value(bytes32 _value) returns()
func (_TargetContract *TargetContractTransactor) WriteB32Value(opts *bind.TransactOpts, _value [32]byte) (*types.Transaction, error) {
	return _TargetContract.contract.Transact(opts, "writeB32Value", _value)
}

// WriteB32Value is a paid mutator transaction binding the contract method 0xba1434e7.
//
// Solidity: function writeB32Value(bytes32 _value) returns()
func (_TargetContract *TargetContractSession) WriteB32Value(_value [32]byte) (*types.Transaction, error) {
	return _TargetContract.Contract.WriteB32Value(&_TargetContract.TransactOpts, _value)
}

// WriteB32Value is a paid mutator transaction binding the contract method 0xba1434e7.
//
// Solidity: function writeB32Value(bytes32 _value) returns()
func (_TargetContract *TargetContractTransactorSession) WriteB32Value(_value [32]byte) (*types.Transaction, error) {
	return _TargetContract.Contract.WriteB32Value(&_TargetContract.TransactOpts, _value)
}

// WriteStrValue is a paid mutator transaction binding the contract method 0x8b26be0d.
//
// Solidity: function writeStrValue(string _value) returns()
func (_TargetContract *TargetContractTransactor) WriteStrValue(opts *bind.TransactOpts, _value string) (*types.Transaction, error) {
	return _TargetContract.contract.Transact(opts, "writeStrValue", _value)
}

// WriteStrValue is a paid mutator transaction binding the contract method 0x8b26be0d.
//
// Solidity: function writeStrValue(string _value) returns()
func (_TargetContract *TargetContractSession) WriteStrValue(_value string) (*types.Transaction, error) {
	return _TargetContract.Contract.WriteStrValue(&_TargetContract.TransactOpts, _value)
}

// WriteStrValue is a paid mutator transaction binding the contract method 0x8b26be0d.
//
// Solidity: function writeStrValue(string _value) returns()
func (_TargetContract *TargetContractTransactorSession) WriteStrValue(_value string) (*types.Transaction, error) {
	return _TargetContract.Contract.WriteStrValue(&_TargetContract.TransactOpts, _value)
}

// WriteUintValue is a paid mutator transaction binding the contract method 0xc567fecd.
//
// Solidity: function writeUintValue(uint256 _value) returns()
func (_TargetContract *TargetContractTransactor) WriteUintValue(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _TargetContract.contract.Transact(opts, "writeUintValue", _value)
}

// WriteUintValue is a paid mutator transaction binding the contract method 0xc567fecd.
//
// Solidity: function writeUintValue(uint256 _value) returns()
func (_TargetContract *TargetContractSession) WriteUintValue(_value *big.Int) (*types.Transaction, error) {
	return _TargetContract.Contract.WriteUintValue(&_TargetContract.TransactOpts, _value)
}

// WriteUintValue is a paid mutator transaction binding the contract method 0xc567fecd.
//
// Solidity: function writeUintValue(uint256 _value) returns()
func (_TargetContract *TargetContractTransactorSession) WriteUintValue(_value *big.Int) (*types.Transaction, error) {
	return _TargetContract.Contract.WriteUintValue(&_TargetContract.TransactOpts, _value)
}

// TargetContractAddressWrittenIterator is returned from FilterAddressWritten and is used to iterate over the raw logs and unpacked data for AddressWritten events raised by the TargetContract contract.
type TargetContractAddressWrittenIterator struct {
	Event *TargetContractAddressWritten // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TargetContractAddressWrittenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TargetContractAddressWritten)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TargetContractAddressWritten)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TargetContractAddressWrittenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TargetContractAddressWrittenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TargetContractAddressWritten represents a AddressWritten event raised by the TargetContract contract.
type TargetContractAddressWritten struct {
	Value common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddressWritten is a free log retrieval operation binding the contract event 0x22b522eb67a8c1256002bbca0d50bcecf337e109bb750f63c301bade4bc198f0.
//
// Solidity: event AddressWritten(address indexed value)
func (_TargetContract *TargetContractFilterer) FilterAddressWritten(opts *bind.FilterOpts, value []common.Address) (*TargetContractAddressWrittenIterator, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _TargetContract.contract.FilterLogs(opts, "AddressWritten", valueRule)
	if err != nil {
		return nil, err
	}
	return &TargetContractAddressWrittenIterator{contract: _TargetContract.contract, event: "AddressWritten", logs: logs, sub: sub}, nil
}

// WatchAddressWritten is a free log subscription operation binding the contract event 0x22b522eb67a8c1256002bbca0d50bcecf337e109bb750f63c301bade4bc198f0.
//
// Solidity: event AddressWritten(address indexed value)
func (_TargetContract *TargetContractFilterer) WatchAddressWritten(opts *bind.WatchOpts, sink chan<- *TargetContractAddressWritten, value []common.Address) (event.Subscription, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _TargetContract.contract.WatchLogs(opts, "AddressWritten", valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TargetContractAddressWritten)
				if err := _TargetContract.contract.UnpackLog(event, "AddressWritten", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddressWritten is a log parse operation binding the contract event 0x22b522eb67a8c1256002bbca0d50bcecf337e109bb750f63c301bade4bc198f0.
//
// Solidity: event AddressWritten(address indexed value)
func (_TargetContract *TargetContractFilterer) ParseAddressWritten(log types.Log) (*TargetContractAddressWritten, error) {
	event := new(TargetContractAddressWritten)
	if err := _TargetContract.contract.UnpackLog(event, "AddressWritten", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TargetContractBytes32WrittenIterator is returned from FilterBytes32Written and is used to iterate over the raw logs and unpacked data for Bytes32Written events raised by the TargetContract contract.
type TargetContractBytes32WrittenIterator struct {
	Event *TargetContractBytes32Written // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TargetContractBytes32WrittenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TargetContractBytes32Written)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TargetContractBytes32Written)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TargetContractBytes32WrittenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TargetContractBytes32WrittenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TargetContractBytes32Written represents a Bytes32Written event raised by the TargetContract contract.
type TargetContractBytes32Written struct {
	Value [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBytes32Written is a free log retrieval operation binding the contract event 0xb20eff94cd1e4024682891d638cc33ad15661f4b9de7ef517dce69ff7392b952.
//
// Solidity: event Bytes32Written(bytes32 indexed value)
func (_TargetContract *TargetContractFilterer) FilterBytes32Written(opts *bind.FilterOpts, value [][32]byte) (*TargetContractBytes32WrittenIterator, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _TargetContract.contract.FilterLogs(opts, "Bytes32Written", valueRule)
	if err != nil {
		return nil, err
	}
	return &TargetContractBytes32WrittenIterator{contract: _TargetContract.contract, event: "Bytes32Written", logs: logs, sub: sub}, nil
}

// WatchBytes32Written is a free log subscription operation binding the contract event 0xb20eff94cd1e4024682891d638cc33ad15661f4b9de7ef517dce69ff7392b952.
//
// Solidity: event Bytes32Written(bytes32 indexed value)
func (_TargetContract *TargetContractFilterer) WatchBytes32Written(opts *bind.WatchOpts, sink chan<- *TargetContractBytes32Written, value [][32]byte) (event.Subscription, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _TargetContract.contract.WatchLogs(opts, "Bytes32Written", valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TargetContractBytes32Written)
				if err := _TargetContract.contract.UnpackLog(event, "Bytes32Written", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBytes32Written is a log parse operation binding the contract event 0xb20eff94cd1e4024682891d638cc33ad15661f4b9de7ef517dce69ff7392b952.
//
// Solidity: event Bytes32Written(bytes32 indexed value)
func (_TargetContract *TargetContractFilterer) ParseBytes32Written(log types.Log) (*TargetContractBytes32Written, error) {
	event := new(TargetContractBytes32Written)
	if err := _TargetContract.contract.UnpackLog(event, "Bytes32Written", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TargetContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TargetContract contract.
type TargetContractOwnershipTransferredIterator struct {
	Event *TargetContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TargetContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TargetContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TargetContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TargetContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TargetContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TargetContractOwnershipTransferred represents a OwnershipTransferred event raised by the TargetContract contract.
type TargetContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TargetContract *TargetContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TargetContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TargetContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TargetContractOwnershipTransferredIterator{contract: _TargetContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TargetContract *TargetContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TargetContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TargetContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TargetContractOwnershipTransferred)
				if err := _TargetContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TargetContract *TargetContractFilterer) ParseOwnershipTransferred(log types.Log) (*TargetContractOwnershipTransferred, error) {
	event := new(TargetContractOwnershipTransferred)
	if err := _TargetContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TargetContractStringWrittenIterator is returned from FilterStringWritten and is used to iterate over the raw logs and unpacked data for StringWritten events raised by the TargetContract contract.
type TargetContractStringWrittenIterator struct {
	Event *TargetContractStringWritten // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TargetContractStringWrittenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TargetContractStringWritten)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TargetContractStringWritten)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TargetContractStringWrittenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TargetContractStringWrittenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TargetContractStringWritten represents a StringWritten event raised by the TargetContract contract.
type TargetContractStringWritten struct {
	Value common.Hash
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStringWritten is a free log retrieval operation binding the contract event 0x6ac728ecdca9efb5f0dec54e77bef32812b52dde84f40fb0d0bbe0cb2cf49a5a.
//
// Solidity: event StringWritten(string indexed value)
func (_TargetContract *TargetContractFilterer) FilterStringWritten(opts *bind.FilterOpts, value []string) (*TargetContractStringWrittenIterator, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _TargetContract.contract.FilterLogs(opts, "StringWritten", valueRule)
	if err != nil {
		return nil, err
	}
	return &TargetContractStringWrittenIterator{contract: _TargetContract.contract, event: "StringWritten", logs: logs, sub: sub}, nil
}

// WatchStringWritten is a free log subscription operation binding the contract event 0x6ac728ecdca9efb5f0dec54e77bef32812b52dde84f40fb0d0bbe0cb2cf49a5a.
//
// Solidity: event StringWritten(string indexed value)
func (_TargetContract *TargetContractFilterer) WatchStringWritten(opts *bind.WatchOpts, sink chan<- *TargetContractStringWritten, value []string) (event.Subscription, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _TargetContract.contract.WatchLogs(opts, "StringWritten", valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TargetContractStringWritten)
				if err := _TargetContract.contract.UnpackLog(event, "StringWritten", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStringWritten is a log parse operation binding the contract event 0x6ac728ecdca9efb5f0dec54e77bef32812b52dde84f40fb0d0bbe0cb2cf49a5a.
//
// Solidity: event StringWritten(string indexed value)
func (_TargetContract *TargetContractFilterer) ParseStringWritten(log types.Log) (*TargetContractStringWritten, error) {
	event := new(TargetContractStringWritten)
	if err := _TargetContract.contract.UnpackLog(event, "StringWritten", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TargetContractUint256WrittenIterator is returned from FilterUint256Written and is used to iterate over the raw logs and unpacked data for Uint256Written events raised by the TargetContract contract.
type TargetContractUint256WrittenIterator struct {
	Event *TargetContractUint256Written // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TargetContractUint256WrittenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TargetContractUint256Written)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TargetContractUint256Written)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TargetContractUint256WrittenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TargetContractUint256WrittenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TargetContractUint256Written represents a Uint256Written event raised by the TargetContract contract.
type TargetContractUint256Written struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUint256Written is a free log retrieval operation binding the contract event 0x6370c32f0eb3cd69b66de366ebcfd36200df04ca19510c44f8159824ce32d159.
//
// Solidity: event Uint256Written(uint256 indexed value)
func (_TargetContract *TargetContractFilterer) FilterUint256Written(opts *bind.FilterOpts, value []*big.Int) (*TargetContractUint256WrittenIterator, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _TargetContract.contract.FilterLogs(opts, "Uint256Written", valueRule)
	if err != nil {
		return nil, err
	}
	return &TargetContractUint256WrittenIterator{contract: _TargetContract.contract, event: "Uint256Written", logs: logs, sub: sub}, nil
}

// WatchUint256Written is a free log subscription operation binding the contract event 0x6370c32f0eb3cd69b66de366ebcfd36200df04ca19510c44f8159824ce32d159.
//
// Solidity: event Uint256Written(uint256 indexed value)
func (_TargetContract *TargetContractFilterer) WatchUint256Written(opts *bind.WatchOpts, sink chan<- *TargetContractUint256Written, value []*big.Int) (event.Subscription, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _TargetContract.contract.WatchLogs(opts, "Uint256Written", valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TargetContractUint256Written)
				if err := _TargetContract.contract.UnpackLog(event, "Uint256Written", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUint256Written is a log parse operation binding the contract event 0x6370c32f0eb3cd69b66de366ebcfd36200df04ca19510c44f8159824ce32d159.
//
// Solidity: event Uint256Written(uint256 indexed value)
func (_TargetContract *TargetContractFilterer) ParseUint256Written(log types.Log) (*TargetContractUint256Written, error) {
	event := new(TargetContractUint256Written)
	if err := _TargetContract.contract.UnpackLog(event, "Uint256Written", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
