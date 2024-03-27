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

// BMerc20MetaData contains all meta data concerning the BMerc20 contract.
var BMerc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"erc1155\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20ExceededCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20InvalidCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"BM_ERC1155\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"minted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c060405234801562000010575f80fd5b5060405162000eea38038062000eea83398101604081905262000033916200015d565b6001600160d01b03838360036200004b83826200026e565b5060046200005a82826200026e565b505050805f03620000845760405163392e1e2760e01b81525f600482015260240160405180910390fd5b6080526001600160a01b031660a052506200033a9050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112620000c0575f80fd5b81516001600160401b0380821115620000dd57620000dd6200009c565b604051601f8301601f19908116603f011681019082821181831017156200010857620001086200009c565b816040528381526020925086602085880101111562000125575f80fd5b5f91505b8382101562000148578582018301518183018401529082019062000129565b5f602085830101528094505050505092915050565b5f805f6060848603121562000170575f80fd5b83516001600160401b038082111562000187575f80fd5b6200019587838801620000b0565b94506020860151915080821115620001ab575f80fd5b50620001ba86828701620000b0565b604086015190935090506001600160a01b0381168114620001d9575f80fd5b809150509250925092565b600181811c90821680620001f957607f821691505b6020821081036200021857634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156200026957805f5260205f20601f840160051c81016020851015620002455750805b601f840160051c820191505b8181101562000266575f815560010162000251565b50505b505050565b81516001600160401b038111156200028a576200028a6200009c565b620002a2816200029b8454620001e4565b846200021e565b602080601f831160018114620002d8575f8415620002c05750858301515b5f19600386901b1c1916600185901b17855562000332565b5f85815260208120601f198616915b828110156200030857888601518255948401946001909101908401620002e7565b50858210156200032657878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b60805160a051610b806200036a5f395f8181610315015261040d01525f818161021201526108370152610b805ff3fe6080604052600436106100f6575f3560e01c80636a62784211610089578063a9059cbb11610058578063a9059cbb146102ca578063bf8fbbd2146102e9578063c766c11414610304578063dd62ed3e1461034f5761010b565b80636a6278421461025557806370a082311461026357806379cc67901461029757806395d89b41146102b65761010b565b806323b872dd116100c557806323b872dd146101ca578063313ce567146101e9578063355274ea1461020457806342966c68146102365761010b565b806306fdde0314610125578063095ea7b31461014f57806318160ddd1461017e5780631e7269c51461019c5761010b565b3661010b5761010933610393565b610393565b005b348015610116575f80fd5b504715610109576101096109af565b348015610130575f80fd5b50610139610450565b60405161014691906109c3565b60405180910390f35b34801561015a575f80fd5b5061016e610169366004610a2a565b6104e0565b6040519015158152602001610146565b348015610189575f80fd5b506002545b604051908152602001610146565b3480156101a7575f80fd5b5061016e6101b6366004610a52565b60056020525f908152604090205460ff1681565b3480156101d5575f80fd5b5061016e6101e4366004610a72565b6104f9565b3480156101f4575f80fd5b5060405160128152602001610146565b34801561020f575f80fd5b507f000000000000000000000000000000000000000000000000000000000000000061018e565b348015610241575f80fd5b50610109610250366004610aab565b61051c565b610109610104366004610a52565b34801561026e575f80fd5b5061018e61027d366004610a52565b6001600160a01b03165f9081526020819052604090205490565b3480156102a2575f80fd5b506101096102b1366004610a2a565b610526565b3480156102c1575f80fd5b5061013961053f565b3480156102d5575f80fd5b5061016e6102e4366004610a2a565b61054e565b3480156102f4575f80fd5b5061018e670de0b6b3a764000081565b34801561030f575f80fd5b506103377f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610146565b34801561035a575f80fd5b5061018e610369366004610ac2565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6001600160a01b0381165f9081526005602052604090205460ff16156103e45760405162461bcd60e51b81526020600482015260016024820152603160f81b60448201526064015b60405180910390fd5b6001600160a01b038082165f908152600560205260409020805460ff1916600117905561043b907f000000000000000000000000000000000000000000000000000000000000000016670de0b6b3a764000061055b565b61044d81670de0b6b3a76400006105f3565b50565b60606003805461045f90610af3565b80601f016020809104026020016040519081016040528092919081815260200182805461048b90610af3565b80156104d65780601f106104ad576101008083540402835291602001916104d6565b820191905f5260205f20905b8154815290600101906020018083116104b957829003601f168201915b5050505050905090565b5f336104ed818585610627565b60019150505b92915050565b5f33610506858285610634565b6105118585856106af565b506001949350505050565b61044d338261070c565b610531823383610634565b61053b828261070c565b5050565b60606004805461045f90610af3565b5f336104ed8185856106af565b8047101561057e5760405163cd78605960e01b81523060048201526024016103db565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f81146105c7576040519150601f19603f3d011682016040523d82523d5f602084013e6105cc565b606091505b50509050806105ee57604051630a12f52160e11b815260040160405180910390fd5b505050565b6001600160a01b03821661061c5760405163ec442f0560e01b81525f60048201526024016103db565b61053b5f838361073c565b6105ee8383836001610747565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f1981146106a9578181101561069b57604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016103db565b6106a984848484035f610747565b50505050565b6001600160a01b0383166106d857604051634b637e8f60e11b81525f60048201526024016103db565b6001600160a01b0382166107015760405163ec442f0560e01b81525f60048201526024016103db565b6105ee83838361073c565b6001600160a01b03821661073557604051634b637e8f60e11b81525f60048201526024016103db565b61053b825f835b6105ee838383610819565b6001600160a01b0384166107705760405163e602df0560e01b81525f60048201526024016103db565b6001600160a01b03831661079957604051634a1406b160e11b81525f60048201526024016103db565b6001600160a01b038085165f90815260016020908152604080832093871683529290522082905580156106a957826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161080b91815260200190565b60405180910390a350505050565b610824838383610889565b6001600160a01b0383166105ee576002547f000000000000000000000000000000000000000000000000000000000000000090818111156108825760405163279e7e1560e21b815260048101829052602481018390526044016103db565b5050505050565b6001600160a01b0383166108b3578060025f8282546108a89190610b2b565b909155506109239050565b6001600160a01b0383165f90815260208190526040902054818110156109055760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016103db565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661093f5760028054829003905561095d565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516109a291815260200190565b60405180910390a3505050565b634e487b7160e01b5f52600160045260245ffd5b5f602080835283518060208501525f5b818110156109ef578581018301518582016040015282016109d3565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610a25575f80fd5b919050565b5f8060408385031215610a3b575f80fd5b610a4483610a0f565b946020939093013593505050565b5f60208284031215610a62575f80fd5b610a6b82610a0f565b9392505050565b5f805f60608486031215610a84575f80fd5b610a8d84610a0f565b9250610a9b60208501610a0f565b9150604084013590509250925092565b5f60208284031215610abb575f80fd5b5035919050565b5f8060408385031215610ad3575f80fd5b610adc83610a0f565b9150610aea60208401610a0f565b90509250929050565b600181811c90821680610b0757607f821691505b602082108103610b2557634e487b7160e01b5f52602260045260245ffd5b50919050565b808201808211156104f357634e487b7160e01b5f52601160045260245ffdfea264697066735822122014afd89cac3c55111373ff65e8fcff08f83a279ebf5bf1f678e551a31028188d64736f6c63430008180033",
}

// BMerc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use BMerc20MetaData.ABI instead.
var BMerc20ABI = BMerc20MetaData.ABI

// BMerc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BMerc20MetaData.Bin instead.
var BMerc20Bin = BMerc20MetaData.Bin

// DeployBMerc20 deploys a new Ethereum contract, binding an instance of BMerc20 to it.
func DeployBMerc20(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, erc1155 common.Address) (common.Address, *types.Transaction, *BMerc20, error) {
	parsed, err := BMerc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BMerc20Bin), backend, name, symbol, erc1155)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BMerc20{BMerc20Caller: BMerc20Caller{contract: contract}, BMerc20Transactor: BMerc20Transactor{contract: contract}, BMerc20Filterer: BMerc20Filterer{contract: contract}}, nil
}

// BMerc20 is an auto generated Go binding around an Ethereum contract.
type BMerc20 struct {
	BMerc20Caller     // Read-only binding to the contract
	BMerc20Transactor // Write-only binding to the contract
	BMerc20Filterer   // Log filterer for contract events
}

// BMerc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type BMerc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BMerc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BMerc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BMerc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BMerc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BMerc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BMerc20Session struct {
	Contract     *BMerc20          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BMerc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BMerc20CallerSession struct {
	Contract *BMerc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BMerc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BMerc20TransactorSession struct {
	Contract     *BMerc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BMerc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type BMerc20Raw struct {
	Contract *BMerc20 // Generic contract binding to access the raw methods on
}

// BMerc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BMerc20CallerRaw struct {
	Contract *BMerc20Caller // Generic read-only contract binding to access the raw methods on
}

// BMerc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BMerc20TransactorRaw struct {
	Contract *BMerc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBMerc20 creates a new instance of BMerc20, bound to a specific deployed contract.
func NewBMerc20(address common.Address, backend bind.ContractBackend) (*BMerc20, error) {
	contract, err := bindBMerc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BMerc20{BMerc20Caller: BMerc20Caller{contract: contract}, BMerc20Transactor: BMerc20Transactor{contract: contract}, BMerc20Filterer: BMerc20Filterer{contract: contract}}, nil
}

// NewBMerc20Caller creates a new read-only instance of BMerc20, bound to a specific deployed contract.
func NewBMerc20Caller(address common.Address, caller bind.ContractCaller) (*BMerc20Caller, error) {
	contract, err := bindBMerc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BMerc20Caller{contract: contract}, nil
}

// NewBMerc20Transactor creates a new write-only instance of BMerc20, bound to a specific deployed contract.
func NewBMerc20Transactor(address common.Address, transactor bind.ContractTransactor) (*BMerc20Transactor, error) {
	contract, err := bindBMerc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BMerc20Transactor{contract: contract}, nil
}

// NewBMerc20Filterer creates a new log filterer instance of BMerc20, bound to a specific deployed contract.
func NewBMerc20Filterer(address common.Address, filterer bind.ContractFilterer) (*BMerc20Filterer, error) {
	contract, err := bindBMerc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BMerc20Filterer{contract: contract}, nil
}

// bindBMerc20 binds a generic wrapper to an already deployed contract.
func bindBMerc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BMerc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BMerc20 *BMerc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BMerc20.Contract.BMerc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BMerc20 *BMerc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BMerc20.Contract.BMerc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BMerc20 *BMerc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BMerc20.Contract.BMerc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BMerc20 *BMerc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BMerc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BMerc20 *BMerc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BMerc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BMerc20 *BMerc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BMerc20.Contract.contract.Transact(opts, method, params...)
}

// BMERC1155 is a free data retrieval call binding the contract method 0xc766c114.
//
// Solidity: function BM_ERC1155() view returns(address)
func (_BMerc20 *BMerc20Caller) BMERC1155(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BMerc20.contract.Call(opts, &out, "BM_ERC1155")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BMERC1155 is a free data retrieval call binding the contract method 0xc766c114.
//
// Solidity: function BM_ERC1155() view returns(address)
func (_BMerc20 *BMerc20Session) BMERC1155() (common.Address, error) {
	return _BMerc20.Contract.BMERC1155(&_BMerc20.CallOpts)
}

// BMERC1155 is a free data retrieval call binding the contract method 0xc766c114.
//
// Solidity: function BM_ERC1155() view returns(address)
func (_BMerc20 *BMerc20CallerSession) BMERC1155() (common.Address, error) {
	return _BMerc20.Contract.BMERC1155(&_BMerc20.CallOpts)
}

// COST is a free data retrieval call binding the contract method 0xbf8fbbd2.
//
// Solidity: function COST() view returns(uint256)
func (_BMerc20 *BMerc20Caller) COST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMerc20.contract.Call(opts, &out, "COST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COST is a free data retrieval call binding the contract method 0xbf8fbbd2.
//
// Solidity: function COST() view returns(uint256)
func (_BMerc20 *BMerc20Session) COST() (*big.Int, error) {
	return _BMerc20.Contract.COST(&_BMerc20.CallOpts)
}

// COST is a free data retrieval call binding the contract method 0xbf8fbbd2.
//
// Solidity: function COST() view returns(uint256)
func (_BMerc20 *BMerc20CallerSession) COST() (*big.Int, error) {
	return _BMerc20.Contract.COST(&_BMerc20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BMerc20 *BMerc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BMerc20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BMerc20 *BMerc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BMerc20.Contract.Allowance(&_BMerc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BMerc20 *BMerc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BMerc20.Contract.Allowance(&_BMerc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BMerc20 *BMerc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BMerc20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BMerc20 *BMerc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _BMerc20.Contract.BalanceOf(&_BMerc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BMerc20 *BMerc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BMerc20.Contract.BalanceOf(&_BMerc20.CallOpts, account)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_BMerc20 *BMerc20Caller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMerc20.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_BMerc20 *BMerc20Session) Cap() (*big.Int, error) {
	return _BMerc20.Contract.Cap(&_BMerc20.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_BMerc20 *BMerc20CallerSession) Cap() (*big.Int, error) {
	return _BMerc20.Contract.Cap(&_BMerc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BMerc20 *BMerc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BMerc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BMerc20 *BMerc20Session) Decimals() (uint8, error) {
	return _BMerc20.Contract.Decimals(&_BMerc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BMerc20 *BMerc20CallerSession) Decimals() (uint8, error) {
	return _BMerc20.Contract.Decimals(&_BMerc20.CallOpts)
}

// Minted is a free data retrieval call binding the contract method 0x1e7269c5.
//
// Solidity: function minted(address account) view returns(bool)
func (_BMerc20 *BMerc20Caller) Minted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _BMerc20.contract.Call(opts, &out, "minted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Minted is a free data retrieval call binding the contract method 0x1e7269c5.
//
// Solidity: function minted(address account) view returns(bool)
func (_BMerc20 *BMerc20Session) Minted(account common.Address) (bool, error) {
	return _BMerc20.Contract.Minted(&_BMerc20.CallOpts, account)
}

// Minted is a free data retrieval call binding the contract method 0x1e7269c5.
//
// Solidity: function minted(address account) view returns(bool)
func (_BMerc20 *BMerc20CallerSession) Minted(account common.Address) (bool, error) {
	return _BMerc20.Contract.Minted(&_BMerc20.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BMerc20 *BMerc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BMerc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BMerc20 *BMerc20Session) Name() (string, error) {
	return _BMerc20.Contract.Name(&_BMerc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BMerc20 *BMerc20CallerSession) Name() (string, error) {
	return _BMerc20.Contract.Name(&_BMerc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BMerc20 *BMerc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BMerc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BMerc20 *BMerc20Session) Symbol() (string, error) {
	return _BMerc20.Contract.Symbol(&_BMerc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BMerc20 *BMerc20CallerSession) Symbol() (string, error) {
	return _BMerc20.Contract.Symbol(&_BMerc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BMerc20 *BMerc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMerc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BMerc20 *BMerc20Session) TotalSupply() (*big.Int, error) {
	return _BMerc20.Contract.TotalSupply(&_BMerc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BMerc20 *BMerc20CallerSession) TotalSupply() (*big.Int, error) {
	return _BMerc20.Contract.TotalSupply(&_BMerc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BMerc20 *BMerc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BMerc20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BMerc20 *BMerc20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BMerc20.Contract.Approve(&_BMerc20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BMerc20 *BMerc20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BMerc20.Contract.Approve(&_BMerc20.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_BMerc20 *BMerc20Transactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _BMerc20.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_BMerc20 *BMerc20Session) Burn(value *big.Int) (*types.Transaction, error) {
	return _BMerc20.Contract.Burn(&_BMerc20.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_BMerc20 *BMerc20TransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _BMerc20.Contract.Burn(&_BMerc20.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_BMerc20 *BMerc20Transactor) BurnFrom(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _BMerc20.contract.Transact(opts, "burnFrom", account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_BMerc20 *BMerc20Session) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _BMerc20.Contract.BurnFrom(&_BMerc20.TransactOpts, account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_BMerc20 *BMerc20TransactorSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _BMerc20.Contract.BurnFrom(&_BMerc20.TransactOpts, account, value)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address account) payable returns()
func (_BMerc20 *BMerc20Transactor) Mint(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BMerc20.contract.Transact(opts, "mint", account)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address account) payable returns()
func (_BMerc20 *BMerc20Session) Mint(account common.Address) (*types.Transaction, error) {
	return _BMerc20.Contract.Mint(&_BMerc20.TransactOpts, account)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address account) payable returns()
func (_BMerc20 *BMerc20TransactorSession) Mint(account common.Address) (*types.Transaction, error) {
	return _BMerc20.Contract.Mint(&_BMerc20.TransactOpts, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BMerc20 *BMerc20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BMerc20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BMerc20 *BMerc20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BMerc20.Contract.Transfer(&_BMerc20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BMerc20 *BMerc20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BMerc20.Contract.Transfer(&_BMerc20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BMerc20 *BMerc20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BMerc20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BMerc20 *BMerc20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BMerc20.Contract.TransferFrom(&_BMerc20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BMerc20 *BMerc20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BMerc20.Contract.TransferFrom(&_BMerc20.TransactOpts, from, to, value)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_BMerc20 *BMerc20Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BMerc20.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_BMerc20 *BMerc20Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BMerc20.Contract.Fallback(&_BMerc20.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_BMerc20 *BMerc20TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BMerc20.Contract.Fallback(&_BMerc20.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BMerc20 *BMerc20Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BMerc20.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BMerc20 *BMerc20Session) Receive() (*types.Transaction, error) {
	return _BMerc20.Contract.Receive(&_BMerc20.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BMerc20 *BMerc20TransactorSession) Receive() (*types.Transaction, error) {
	return _BMerc20.Contract.Receive(&_BMerc20.TransactOpts)
}

// BMerc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BMerc20 contract.
type BMerc20ApprovalIterator struct {
	Event *BMerc20Approval // Event containing the contract specifics and raw log

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
func (it *BMerc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BMerc20Approval)
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
		it.Event = new(BMerc20Approval)
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
func (it *BMerc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BMerc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BMerc20Approval represents a Approval event raised by the BMerc20 contract.
type BMerc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BMerc20 *BMerc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BMerc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BMerc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BMerc20ApprovalIterator{contract: _BMerc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BMerc20 *BMerc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BMerc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BMerc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BMerc20Approval)
				if err := _BMerc20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BMerc20 *BMerc20Filterer) ParseApproval(log types.Log) (*BMerc20Approval, error) {
	event := new(BMerc20Approval)
	if err := _BMerc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BMerc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BMerc20 contract.
type BMerc20TransferIterator struct {
	Event *BMerc20Transfer // Event containing the contract specifics and raw log

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
func (it *BMerc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BMerc20Transfer)
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
		it.Event = new(BMerc20Transfer)
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
func (it *BMerc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BMerc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BMerc20Transfer represents a Transfer event raised by the BMerc20 contract.
type BMerc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BMerc20 *BMerc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BMerc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BMerc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BMerc20TransferIterator{contract: _BMerc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BMerc20 *BMerc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BMerc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BMerc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BMerc20Transfer)
				if err := _BMerc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BMerc20 *BMerc20Filterer) ParseTransfer(log types.Log) (*BMerc20Transfer, error) {
	event := new(BMerc20Transfer)
	if err := _BMerc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
