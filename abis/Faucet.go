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

// FaucetMetaData contains all meta data concerning the Faucet contract.
var FaucetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FAUCET_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061028c8061001d5f395ff3fe608060405260043610610036575f3560e01c80631e83409a146100415780636df62cef14610062578063d17891761461008a575f80fd5b3661003d57005b5f80fd5b34801561004c575f80fd5b5061006061005b36600461020a565b6100a5565b005b34801561006d575f80fd5b506100786201518081565b60405190815260200160405180910390f35b348015610095575f80fd5b50610078678ac7230489e8000081565b6001600160a01b0381166100b65750335b5f6100c46201518042610237565b6001600160a01b0383165f9081526020819052604090205490915081900361010f57604051632058b6db60e01b81526001600160a01b03831660048201526024015b60405180910390fd5b6001600160a01b0382165f81815260208190526040902082905561013b90678ac7230489e80000610172565b6040516001600160a01b038316907fb449c24d261a59627b537c8c41c57ab559f4205c56bea745ff61c5521bece214905f90a25050565b804710156101955760405163cd78605960e01b8152306004820152602401610106565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f81146101de576040519150601f19603f3d011682016040523d82523d5f602084013e6101e3565b606091505b505090508061020557604051630a12f52160e11b815260040160405180910390fd5b505050565b5f6020828403121561021a575f80fd5b81356001600160a01b0381168114610230575f80fd5b9392505050565b5f8261025157634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220f36399db1829d3478a75796846041ae8d3dae3a7b05bfa8abf54fe0ecbe013ee64736f6c63430008180033",
}

// FaucetABI is the input ABI used to generate the binding from.
// Deprecated: Use FaucetMetaData.ABI instead.
var FaucetABI = FaucetMetaData.ABI

// FaucetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FaucetMetaData.Bin instead.
var FaucetBin = FaucetMetaData.Bin

// DeployFaucet deploys a new Ethereum contract, binding an instance of Faucet to it.
func DeployFaucet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Faucet, error) {
	parsed, err := FaucetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FaucetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Faucet{FaucetCaller: FaucetCaller{contract: contract}, FaucetTransactor: FaucetTransactor{contract: contract}, FaucetFilterer: FaucetFilterer{contract: contract}}, nil
}

// Faucet is an auto generated Go binding around an Ethereum contract.
type Faucet struct {
	FaucetCaller     // Read-only binding to the contract
	FaucetTransactor // Write-only binding to the contract
	FaucetFilterer   // Log filterer for contract events
}

// FaucetCaller is an auto generated read-only Go binding around an Ethereum contract.
type FaucetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FaucetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FaucetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FaucetSession struct {
	Contract     *Faucet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaucetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FaucetCallerSession struct {
	Contract *FaucetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FaucetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FaucetTransactorSession struct {
	Contract     *FaucetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaucetRaw is an auto generated low-level Go binding around an Ethereum contract.
type FaucetRaw struct {
	Contract *Faucet // Generic contract binding to access the raw methods on
}

// FaucetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FaucetCallerRaw struct {
	Contract *FaucetCaller // Generic read-only contract binding to access the raw methods on
}

// FaucetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FaucetTransactorRaw struct {
	Contract *FaucetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFaucet creates a new instance of Faucet, bound to a specific deployed contract.
func NewFaucet(address common.Address, backend bind.ContractBackend) (*Faucet, error) {
	contract, err := bindFaucet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Faucet{FaucetCaller: FaucetCaller{contract: contract}, FaucetTransactor: FaucetTransactor{contract: contract}, FaucetFilterer: FaucetFilterer{contract: contract}}, nil
}

// NewFaucetCaller creates a new read-only instance of Faucet, bound to a specific deployed contract.
func NewFaucetCaller(address common.Address, caller bind.ContractCaller) (*FaucetCaller, error) {
	contract, err := bindFaucet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetCaller{contract: contract}, nil
}

// NewFaucetTransactor creates a new write-only instance of Faucet, bound to a specific deployed contract.
func NewFaucetTransactor(address common.Address, transactor bind.ContractTransactor) (*FaucetTransactor, error) {
	contract, err := bindFaucet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetTransactor{contract: contract}, nil
}

// NewFaucetFilterer creates a new log filterer instance of Faucet, bound to a specific deployed contract.
func NewFaucetFilterer(address common.Address, filterer bind.ContractFilterer) (*FaucetFilterer, error) {
	contract, err := bindFaucet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FaucetFilterer{contract: contract}, nil
}

// bindFaucet binds a generic wrapper to an already deployed contract.
func bindFaucet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FaucetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Faucet *FaucetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Faucet.Contract.FaucetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Faucet *FaucetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.Contract.FaucetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Faucet *FaucetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Faucet.Contract.FaucetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Faucet *FaucetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Faucet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Faucet *FaucetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Faucet *FaucetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Faucet.Contract.contract.Transact(opts, method, params...)
}

// AMOUNT is a free data retrieval call binding the contract method 0xd1789176.
//
// Solidity: function AMOUNT() view returns(uint256)
func (_Faucet *FaucetCaller) AMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AMOUNT is a free data retrieval call binding the contract method 0xd1789176.
//
// Solidity: function AMOUNT() view returns(uint256)
func (_Faucet *FaucetSession) AMOUNT() (*big.Int, error) {
	return _Faucet.Contract.AMOUNT(&_Faucet.CallOpts)
}

// AMOUNT is a free data retrieval call binding the contract method 0xd1789176.
//
// Solidity: function AMOUNT() view returns(uint256)
func (_Faucet *FaucetCallerSession) AMOUNT() (*big.Int, error) {
	return _Faucet.Contract.AMOUNT(&_Faucet.CallOpts)
}

// FAUCETINTERVAL is a free data retrieval call binding the contract method 0x6df62cef.
//
// Solidity: function FAUCET_INTERVAL() view returns(uint256)
func (_Faucet *FaucetCaller) FAUCETINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "FAUCET_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FAUCETINTERVAL is a free data retrieval call binding the contract method 0x6df62cef.
//
// Solidity: function FAUCET_INTERVAL() view returns(uint256)
func (_Faucet *FaucetSession) FAUCETINTERVAL() (*big.Int, error) {
	return _Faucet.Contract.FAUCETINTERVAL(&_Faucet.CallOpts)
}

// FAUCETINTERVAL is a free data retrieval call binding the contract method 0x6df62cef.
//
// Solidity: function FAUCET_INTERVAL() view returns(uint256)
func (_Faucet *FaucetCallerSession) FAUCETINTERVAL() (*big.Int, error) {
	return _Faucet.Contract.FAUCETINTERVAL(&_Faucet.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address account) returns()
func (_Faucet *FaucetTransactor) Claim(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "claim", account)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address account) returns()
func (_Faucet *FaucetSession) Claim(account common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.Claim(&_Faucet.TransactOpts, account)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address account) returns()
func (_Faucet *FaucetTransactorSession) Claim(account common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.Claim(&_Faucet.TransactOpts, account)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Faucet *FaucetTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Faucet *FaucetSession) Receive() (*types.Transaction, error) {
	return _Faucet.Contract.Receive(&_Faucet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Faucet *FaucetTransactorSession) Receive() (*types.Transaction, error) {
	return _Faucet.Contract.Receive(&_Faucet.TransactOpts)
}

// FaucetClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Faucet contract.
type FaucetClaimedIterator struct {
	Event *FaucetClaimed // Event containing the contract specifics and raw log

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
func (it *FaucetClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaucetClaimed)
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
		it.Event = new(FaucetClaimed)
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
func (it *FaucetClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaucetClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaucetClaimed represents a Claimed event raised by the Faucet contract.
type FaucetClaimed struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xb449c24d261a59627b537c8c41c57ab559f4205c56bea745ff61c5521bece214.
//
// Solidity: event Claimed(address indexed account)
func (_Faucet *FaucetFilterer) FilterClaimed(opts *bind.FilterOpts, account []common.Address) (*FaucetClaimedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Faucet.contract.FilterLogs(opts, "Claimed", accountRule)
	if err != nil {
		return nil, err
	}
	return &FaucetClaimedIterator{contract: _Faucet.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xb449c24d261a59627b537c8c41c57ab559f4205c56bea745ff61c5521bece214.
//
// Solidity: event Claimed(address indexed account)
func (_Faucet *FaucetFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *FaucetClaimed, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Faucet.contract.WatchLogs(opts, "Claimed", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaucetClaimed)
				if err := _Faucet.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xb449c24d261a59627b537c8c41c57ab559f4205c56bea745ff61c5521bece214.
//
// Solidity: event Claimed(address indexed account)
func (_Faucet *FaucetFilterer) ParseClaimed(log types.Log) (*FaucetClaimed, error) {
	event := new(FaucetClaimed)
	if err := _Faucet.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
