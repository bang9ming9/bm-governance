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

// BmErc20MetaData contains all meta data concerning the BmErc20 contract.
var BmErc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"erc1155\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"BmErc20AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"BmErc20AlreadyMintedUser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"BmErc20InvalidProposalState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"arg\",\"type\":\"string\"}],\"name\":\"BmErc20NilInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"BmErc20NotVoteUser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20ExceededCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20InvalidCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"ABSTAIN_CLAIM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BM_ERC1155\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BM_GOVERNOR\",\"outputs\":[{\"internalType\":\"contractIGovernor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOSE_CLAIM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WIN_CLAIM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"minted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e060405234801562000010575f80fd5b50604051620013f7380380620013f7833981016040819052620000339162000203565b6001600160d01b03848460036200004b83826200031c565b5060046200005a82826200031c565b505050805f03620000855760405163392e1e2760e01b81525f60048201526024015b60405180910390fd5b6080526001600160a01b038216620000cb57604051631f15b64760e31b81526020600482015260076024820152666572633131353560c81b60448201526064016200007c565b6001600160a01b0381166200010f57604051631f15b64760e31b815260206004820152600860248201526733b7bb32b93737b960c11b60448201526064016200007c565b6001600160a01b0391821660a0521660c05250620003e89050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126200014e575f80fd5b81516001600160401b03808211156200016b576200016b6200012a565b604051601f8301601f19908116603f011681019082821181831017156200019657620001966200012a565b8160405283815260209250866020858801011115620001b3575f80fd5b5f91505b83821015620001d65785820183015181830184015290820190620001b7565b5f602085830101528094505050505092915050565b6001600160a01b038116811462000200575f80fd5b50565b5f805f806080858703121562000217575f80fd5b84516001600160401b03808211156200022e575f80fd5b6200023c888389016200013e565b9550602087015191508082111562000252575f80fd5b5062000261878288016200013e565b93505060408501516200027481620001eb565b60608601519092506200028781620001eb565b939692955090935050565b600181811c90821680620002a757607f821691505b602082108103620002c657634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156200031757805f5260205f20601f840160051c81016020851015620002f35750805b601f840160051c820191505b8181101562000314575f8155600101620002ff565b50505b505050565b81516001600160401b038111156200033857620003386200012a565b620003508162000349845462000292565b84620002cc565b602080601f83116001811462000386575f84156200036e5750858301515b5f19600386901b1c1916600185901b178555620003e0565b5f85815260208120601f198616915b82811015620003b65788860151825594840194600190910190840162000395565b5085821015620003d457878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b60805160a05160c051610fcf620004285f395f818161048f015261070101525f81816103e1015261052301525f81816102a80152610bfb0152610fcf5ff3fe608060405260043610610138575f3560e01c806370a08231116100aa578063ad1267be1161006e578063ad1267be1461039a578063bf8fbbd2146103b5578063c766c114146103d0578063dd62ed3e1461041b578063ddd5e1b21461045f578063f0732fd31461047e5761014d565b806370a08231146102f957806379cc67901461032d57806386a3199c1461034c57806395d89b4114610367578063a9059cbb1461037b5761014d565b80631e7269c5116100fc5780631e7269c51461023257806323b872dd14610260578063313ce5671461027f578063355274ea1461029a57806342966c68146102cc5780636a627842146102eb5761014d565b806305dc76da1461016757806306fdde0314610195578063095ea7b3146101b6578063120aa877146101e557806318160ddd1461021e5761014d565b3661014d5761014b336104b1565b6104b1565b005b348015610158575f80fd5b50471561014b5761014b610d73565b348015610172575f80fd5b506101826702c68af0bb14000081565b6040519081526020015b60405180910390f35b3480156101a0575f80fd5b506101a9610599565b60405161018c9190610d87565b3480156101c1575f80fd5b506101d56101d0366004610dee565b610629565b604051901515815260200161018c565b3480156101f0575f80fd5b506101d56101ff366004610e16565b600660209081525f928352604080842090915290825290205460ff1681565b348015610229575f80fd5b50600254610182565b34801561023d575f80fd5b506101d561024c366004610e40565b60056020525f908152604090205460ff1681565b34801561026b575f80fd5b506101d561027a366004610e60565b610642565b34801561028a575f80fd5b506040516012815260200161018c565b3480156102a5575f80fd5b507f0000000000000000000000000000000000000000000000000000000000000000610182565b3480156102d7575f80fd5b5061014b6102e6366004610e99565b610665565b61014b610146366004610e40565b348015610304575f80fd5b50610182610313366004610e40565b6001600160a01b03165f9081526020819052604090205490565b348015610338575f80fd5b5061014b610347366004610dee565b610672565b348015610357575f80fd5b506101826706f05b59d3b2000081565b348015610372575f80fd5b506101a961068b565b348015610386575f80fd5b506101d5610395366004610dee565b61069a565b3480156103a5575f80fd5b5061018267016345785d8a000081565b3480156103c0575f80fd5b50610182670de0b6b3a764000081565b3480156103db575f80fd5b506104037f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161018c565b348015610426575f80fd5b50610182610435366004610eb0565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b34801561046a575f80fd5b5061014b610479366004610e16565b6106a7565b348015610489575f80fd5b506104037f000000000000000000000000000000000000000000000000000000000000000081565b6001600160a01b0381165f9081526005602052604090205460ff16156104fa5760405163699e6f2f60e01b81526001600160a01b03821660048201526024015b60405180910390fd5b6001600160a01b038082165f908152600560205260409020805460ff19166001179055610551907f000000000000000000000000000000000000000000000000000000000000000016670de0b6b3a764000061091f565b61056381670de0b6b3a76400006109b7565b6040516001600160a01b038216907f90ddedd5a25821bba11fbb98de02ec1f75c1be90ae147d6450ce873e7b78b5d8905f90a250565b6060600380546105a890610ed8565b80601f01602080910402602001604051908101604052809291908181526020018280546105d490610ed8565b801561061f5780601f106105f65761010080835404028352916020019161061f565b820191905f5260205f20905b81548152906001019060200180831161060257829003601f168201915b5050505050905090565b5f336106368185856109eb565b60019150505b92915050565b5f3361064f8582856109f8565b61065a858585610a73565b506001949350505050565b61066f3382610ad0565b50565b61067d8233836109f8565b6106878282610ad0565b5050565b6060600480546105a890610ed8565b5f33610636818585610a73565b5f8281526006602090815260408083206001600160a01b038516845290915290205460ff16156106fc5760405163109f4dc760e01b8152600481018390526001600160a01b03821660248201526044016104f1565b5f805f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a47516e186866040518363ffffffff1660e01b815260040161075f9291909182526001600160a01b0316602082015260400190565b606060405180830381865afa15801561077a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061079e9190610f10565b925092509250826107d457604051630aa10b2560e31b8152600481018690526001600160a01b03851660248201526044016104f1565b5f8260ff166002036107ef57506702c68af0bb1400006108b7565b600782600781111561080357610803610f66565b14806108205750600482600781111561081e5761081e610f66565b145b1561084f578260ff1660011461083e5767016345785d8a0000610848565b6706f05b59d3b200005b90506108b7565b600382600781111561086357610863610f66565b0361087f5760ff83161561083e5767016345785d8a0000610848565b8582600781111561089257610892610f66565b604051636204fd5560e01b8152600481019290925260ff1660248201526044016104f1565b6108c185826109b7565b5f8681526006602090815260408083206001600160a01b0389168085529252808320805460ff1916600117905551839289917f4ec90e965519d92681267467f775ada5bd214aa92c0dc93d90a5e880ce9ed0269190a4505050505050565b804710156109425760405163cd78605960e01b81523060048201526024016104f1565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f811461098b576040519150601f19603f3d011682016040523d82523d5f602084013e610990565b606091505b50509050806109b257604051630a12f52160e11b815260040160405180910390fd5b505050565b6001600160a01b0382166109e05760405163ec442f0560e01b81525f60048201526024016104f1565b6106875f8383610b00565b6109b28383836001610b0b565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f198114610a6d5781811015610a5f57604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016104f1565b610a6d84848484035f610b0b565b50505050565b6001600160a01b038316610a9c57604051634b637e8f60e11b81525f60048201526024016104f1565b6001600160a01b038216610ac55760405163ec442f0560e01b81525f60048201526024016104f1565b6109b2838383610b00565b6001600160a01b038216610af957604051634b637e8f60e11b81525f60048201526024016104f1565b610687825f835b6109b2838383610bdd565b6001600160a01b038416610b345760405163e602df0560e01b81525f60048201526024016104f1565b6001600160a01b038316610b5d57604051634a1406b160e11b81525f60048201526024016104f1565b6001600160a01b038085165f9081526001602090815260408083209387168352929052208290558015610a6d57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bcf91815260200190565b60405180910390a350505050565b610be8838383610c4d565b6001600160a01b0383166109b2576002547f00000000000000000000000000000000000000000000000000000000000000009081811115610c465760405163279e7e1560e21b815260048101829052602481018390526044016104f1565b5050505050565b6001600160a01b038316610c77578060025f828254610c6c9190610f7a565b90915550610ce79050565b6001600160a01b0383165f9081526020819052604090205481811015610cc95760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016104f1565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b038216610d0357600280548290039055610d21565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610d6691815260200190565b60405180910390a3505050565b634e487b7160e01b5f52600160045260245ffd5b5f602080835283518060208501525f5b81811015610db357858101830151858201604001528201610d97565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610de9575f80fd5b919050565b5f8060408385031215610dff575f80fd5b610e0883610dd3565b946020939093013593505050565b5f8060408385031215610e27575f80fd5b82359150610e3760208401610dd3565b90509250929050565b5f60208284031215610e50575f80fd5b610e5982610dd3565b9392505050565b5f805f60608486031215610e72575f80fd5b610e7b84610dd3565b9250610e8960208501610dd3565b9150604084013590509250925092565b5f60208284031215610ea9575f80fd5b5035919050565b5f8060408385031215610ec1575f80fd5b610eca83610dd3565b9150610e3760208401610dd3565b600181811c90821680610eec57607f821691505b602082108103610f0a57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f805f60608486031215610f22575f80fd5b83518015158114610f31575f80fd5b602085015190935060ff81168114610f47575f80fd5b604085015190925060088110610f5b575f80fd5b809150509250925092565b634e487b7160e01b5f52602160045260245ffd5b8082018082111561063c57634e487b7160e01b5f52601160045260245ffdfea2646970667358221220bce87e3c58c7ca0c73d96e5c2bf87f532620fc8b9974727f0a0345e409f6663464736f6c63430008180033",
}

// BmErc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use BmErc20MetaData.ABI instead.
var BmErc20ABI = BmErc20MetaData.ABI

// BmErc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BmErc20MetaData.Bin instead.
var BmErc20Bin = BmErc20MetaData.Bin

// DeployBmErc20 deploys a new Ethereum contract, binding an instance of BmErc20 to it.
func DeployBmErc20(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, erc1155 common.Address, governor common.Address) (common.Address, *types.Transaction, *BmErc20, error) {
	parsed, err := BmErc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BmErc20Bin), backend, name, symbol, erc1155, governor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BmErc20{BmErc20Caller: BmErc20Caller{contract: contract}, BmErc20Transactor: BmErc20Transactor{contract: contract}, BmErc20Filterer: BmErc20Filterer{contract: contract}}, nil
}

// BmErc20 is an auto generated Go binding around an Ethereum contract.
type BmErc20 struct {
	BmErc20Caller     // Read-only binding to the contract
	BmErc20Transactor // Write-only binding to the contract
	BmErc20Filterer   // Log filterer for contract events
}

// BmErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type BmErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BmErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BmErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BmErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BmErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BmErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BmErc20Session struct {
	Contract     *BmErc20          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BmErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BmErc20CallerSession struct {
	Contract *BmErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BmErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BmErc20TransactorSession struct {
	Contract     *BmErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BmErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type BmErc20Raw struct {
	Contract *BmErc20 // Generic contract binding to access the raw methods on
}

// BmErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BmErc20CallerRaw struct {
	Contract *BmErc20Caller // Generic read-only contract binding to access the raw methods on
}

// BmErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BmErc20TransactorRaw struct {
	Contract *BmErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBmErc20 creates a new instance of BmErc20, bound to a specific deployed contract.
func NewBmErc20(address common.Address, backend bind.ContractBackend) (*BmErc20, error) {
	contract, err := bindBmErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BmErc20{BmErc20Caller: BmErc20Caller{contract: contract}, BmErc20Transactor: BmErc20Transactor{contract: contract}, BmErc20Filterer: BmErc20Filterer{contract: contract}}, nil
}

// NewBmErc20Caller creates a new read-only instance of BmErc20, bound to a specific deployed contract.
func NewBmErc20Caller(address common.Address, caller bind.ContractCaller) (*BmErc20Caller, error) {
	contract, err := bindBmErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BmErc20Caller{contract: contract}, nil
}

// NewBmErc20Transactor creates a new write-only instance of BmErc20, bound to a specific deployed contract.
func NewBmErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*BmErc20Transactor, error) {
	contract, err := bindBmErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BmErc20Transactor{contract: contract}, nil
}

// NewBmErc20Filterer creates a new log filterer instance of BmErc20, bound to a specific deployed contract.
func NewBmErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*BmErc20Filterer, error) {
	contract, err := bindBmErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BmErc20Filterer{contract: contract}, nil
}

// bindBmErc20 binds a generic wrapper to an already deployed contract.
func bindBmErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BmErc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BmErc20 *BmErc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BmErc20.Contract.BmErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BmErc20 *BmErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BmErc20.Contract.BmErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BmErc20 *BmErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BmErc20.Contract.BmErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BmErc20 *BmErc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BmErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BmErc20 *BmErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BmErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BmErc20 *BmErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BmErc20.Contract.contract.Transact(opts, method, params...)
}

// ABSTAINCLAIM is a free data retrieval call binding the contract method 0x05dc76da.
//
// Solidity: function ABSTAIN_CLAIM() view returns(uint256)
func (_BmErc20 *BmErc20Caller) ABSTAINCLAIM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BmErc20.contract.Call(opts, &out, "ABSTAIN_CLAIM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ABSTAINCLAIM is a free data retrieval call binding the contract method 0x05dc76da.
//
// Solidity: function ABSTAIN_CLAIM() view returns(uint256)
func (_BmErc20 *BmErc20Session) ABSTAINCLAIM() (*big.Int, error) {
	return _BmErc20.Contract.ABSTAINCLAIM(&_BmErc20.CallOpts)
}

// ABSTAINCLAIM is a free data retrieval call binding the contract method 0x05dc76da.
//
// Solidity: function ABSTAIN_CLAIM() view returns(uint256)
func (_BmErc20 *BmErc20CallerSession) ABSTAINCLAIM() (*big.Int, error) {
	return _BmErc20.Contract.ABSTAINCLAIM(&_BmErc20.CallOpts)
}

// BMERC1155 is a free data retrieval call binding the contract method 0xc766c114.
//
// Solidity: function BM_ERC1155() view returns(address)
func (_BmErc20 *BmErc20Caller) BMERC1155(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BmErc20.contract.Call(opts, &out, "BM_ERC1155")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BMERC1155 is a free data retrieval call binding the contract method 0xc766c114.
//
// Solidity: function BM_ERC1155() view returns(address)
func (_BmErc20 *BmErc20Session) BMERC1155() (common.Address, error) {
	return _BmErc20.Contract.BMERC1155(&_BmErc20.CallOpts)
}

// BMERC1155 is a free data retrieval call binding the contract method 0xc766c114.
//
// Solidity: function BM_ERC1155() view returns(address)
func (_BmErc20 *BmErc20CallerSession) BMERC1155() (common.Address, error) {
	return _BmErc20.Contract.BMERC1155(&_BmErc20.CallOpts)
}

// BMGOVERNOR is a free data retrieval call binding the contract method 0xf0732fd3.
//
// Solidity: function BM_GOVERNOR() view returns(address)
func (_BmErc20 *BmErc20Caller) BMGOVERNOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BmErc20.contract.Call(opts, &out, "BM_GOVERNOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BMGOVERNOR is a free data retrieval call binding the contract method 0xf0732fd3.
//
// Solidity: function BM_GOVERNOR() view returns(address)
func (_BmErc20 *BmErc20Session) BMGOVERNOR() (common.Address, error) {
	return _BmErc20.Contract.BMGOVERNOR(&_BmErc20.CallOpts)
}

// BMGOVERNOR is a free data retrieval call binding the contract method 0xf0732fd3.
//
// Solidity: function BM_GOVERNOR() view returns(address)
func (_BmErc20 *BmErc20CallerSession) BMGOVERNOR() (common.Address, error) {
	return _BmErc20.Contract.BMGOVERNOR(&_BmErc20.CallOpts)
}

// COST is a free data retrieval call binding the contract method 0xbf8fbbd2.
//
// Solidity: function COST() view returns(uint256)
func (_BmErc20 *BmErc20Caller) COST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BmErc20.contract.Call(opts, &out, "COST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COST is a free data retrieval call binding the contract method 0xbf8fbbd2.
//
// Solidity: function COST() view returns(uint256)
func (_BmErc20 *BmErc20Session) COST() (*big.Int, error) {
	return _BmErc20.Contract.COST(&_BmErc20.CallOpts)
}

// COST is a free data retrieval call binding the contract method 0xbf8fbbd2.
//
// Solidity: function COST() view returns(uint256)
func (_BmErc20 *BmErc20CallerSession) COST() (*big.Int, error) {
	return _BmErc20.Contract.COST(&_BmErc20.CallOpts)
}

// LOSECLAIM is a free data retrieval call binding the contract method 0xad1267be.
//
// Solidity: function LOSE_CLAIM() view returns(uint256)
func (_BmErc20 *BmErc20Caller) LOSECLAIM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BmErc20.contract.Call(opts, &out, "LOSE_CLAIM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOSECLAIM is a free data retrieval call binding the contract method 0xad1267be.
//
// Solidity: function LOSE_CLAIM() view returns(uint256)
func (_BmErc20 *BmErc20Session) LOSECLAIM() (*big.Int, error) {
	return _BmErc20.Contract.LOSECLAIM(&_BmErc20.CallOpts)
}

// LOSECLAIM is a free data retrieval call binding the contract method 0xad1267be.
//
// Solidity: function LOSE_CLAIM() view returns(uint256)
func (_BmErc20 *BmErc20CallerSession) LOSECLAIM() (*big.Int, error) {
	return _BmErc20.Contract.LOSECLAIM(&_BmErc20.CallOpts)
}

// WINCLAIM is a free data retrieval call binding the contract method 0x86a3199c.
//
// Solidity: function WIN_CLAIM() view returns(uint256)
func (_BmErc20 *BmErc20Caller) WINCLAIM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BmErc20.contract.Call(opts, &out, "WIN_CLAIM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WINCLAIM is a free data retrieval call binding the contract method 0x86a3199c.
//
// Solidity: function WIN_CLAIM() view returns(uint256)
func (_BmErc20 *BmErc20Session) WINCLAIM() (*big.Int, error) {
	return _BmErc20.Contract.WINCLAIM(&_BmErc20.CallOpts)
}

// WINCLAIM is a free data retrieval call binding the contract method 0x86a3199c.
//
// Solidity: function WIN_CLAIM() view returns(uint256)
func (_BmErc20 *BmErc20CallerSession) WINCLAIM() (*big.Int, error) {
	return _BmErc20.Contract.WINCLAIM(&_BmErc20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BmErc20 *BmErc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BmErc20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BmErc20 *BmErc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BmErc20.Contract.Allowance(&_BmErc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BmErc20 *BmErc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BmErc20.Contract.Allowance(&_BmErc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BmErc20 *BmErc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BmErc20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BmErc20 *BmErc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _BmErc20.Contract.BalanceOf(&_BmErc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BmErc20 *BmErc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BmErc20.Contract.BalanceOf(&_BmErc20.CallOpts, account)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_BmErc20 *BmErc20Caller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BmErc20.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_BmErc20 *BmErc20Session) Cap() (*big.Int, error) {
	return _BmErc20.Contract.Cap(&_BmErc20.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_BmErc20 *BmErc20CallerSession) Cap() (*big.Int, error) {
	return _BmErc20.Contract.Cap(&_BmErc20.CallOpts)
}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 proposalID, address account) view returns(bool)
func (_BmErc20 *BmErc20Caller) Claimed(opts *bind.CallOpts, proposalID *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _BmErc20.contract.Call(opts, &out, "claimed", proposalID, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 proposalID, address account) view returns(bool)
func (_BmErc20 *BmErc20Session) Claimed(proposalID *big.Int, account common.Address) (bool, error) {
	return _BmErc20.Contract.Claimed(&_BmErc20.CallOpts, proposalID, account)
}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 proposalID, address account) view returns(bool)
func (_BmErc20 *BmErc20CallerSession) Claimed(proposalID *big.Int, account common.Address) (bool, error) {
	return _BmErc20.Contract.Claimed(&_BmErc20.CallOpts, proposalID, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BmErc20 *BmErc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BmErc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BmErc20 *BmErc20Session) Decimals() (uint8, error) {
	return _BmErc20.Contract.Decimals(&_BmErc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BmErc20 *BmErc20CallerSession) Decimals() (uint8, error) {
	return _BmErc20.Contract.Decimals(&_BmErc20.CallOpts)
}

// Minted is a free data retrieval call binding the contract method 0x1e7269c5.
//
// Solidity: function minted(address account) view returns(bool)
func (_BmErc20 *BmErc20Caller) Minted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _BmErc20.contract.Call(opts, &out, "minted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Minted is a free data retrieval call binding the contract method 0x1e7269c5.
//
// Solidity: function minted(address account) view returns(bool)
func (_BmErc20 *BmErc20Session) Minted(account common.Address) (bool, error) {
	return _BmErc20.Contract.Minted(&_BmErc20.CallOpts, account)
}

// Minted is a free data retrieval call binding the contract method 0x1e7269c5.
//
// Solidity: function minted(address account) view returns(bool)
func (_BmErc20 *BmErc20CallerSession) Minted(account common.Address) (bool, error) {
	return _BmErc20.Contract.Minted(&_BmErc20.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BmErc20 *BmErc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BmErc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BmErc20 *BmErc20Session) Name() (string, error) {
	return _BmErc20.Contract.Name(&_BmErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BmErc20 *BmErc20CallerSession) Name() (string, error) {
	return _BmErc20.Contract.Name(&_BmErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BmErc20 *BmErc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BmErc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BmErc20 *BmErc20Session) Symbol() (string, error) {
	return _BmErc20.Contract.Symbol(&_BmErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BmErc20 *BmErc20CallerSession) Symbol() (string, error) {
	return _BmErc20.Contract.Symbol(&_BmErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BmErc20 *BmErc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BmErc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BmErc20 *BmErc20Session) TotalSupply() (*big.Int, error) {
	return _BmErc20.Contract.TotalSupply(&_BmErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BmErc20 *BmErc20CallerSession) TotalSupply() (*big.Int, error) {
	return _BmErc20.Contract.TotalSupply(&_BmErc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BmErc20 *BmErc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BmErc20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BmErc20 *BmErc20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BmErc20.Contract.Approve(&_BmErc20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BmErc20 *BmErc20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BmErc20.Contract.Approve(&_BmErc20.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_BmErc20 *BmErc20Transactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _BmErc20.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_BmErc20 *BmErc20Session) Burn(value *big.Int) (*types.Transaction, error) {
	return _BmErc20.Contract.Burn(&_BmErc20.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_BmErc20 *BmErc20TransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _BmErc20.Contract.Burn(&_BmErc20.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_BmErc20 *BmErc20Transactor) BurnFrom(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _BmErc20.contract.Transact(opts, "burnFrom", account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_BmErc20 *BmErc20Session) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _BmErc20.Contract.BurnFrom(&_BmErc20.TransactOpts, account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_BmErc20 *BmErc20TransactorSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _BmErc20.Contract.BurnFrom(&_BmErc20.TransactOpts, account, value)
}

// Claim is a paid mutator transaction binding the contract method 0xddd5e1b2.
//
// Solidity: function claim(uint256 proposalID, address account) returns()
func (_BmErc20 *BmErc20Transactor) Claim(opts *bind.TransactOpts, proposalID *big.Int, account common.Address) (*types.Transaction, error) {
	return _BmErc20.contract.Transact(opts, "claim", proposalID, account)
}

// Claim is a paid mutator transaction binding the contract method 0xddd5e1b2.
//
// Solidity: function claim(uint256 proposalID, address account) returns()
func (_BmErc20 *BmErc20Session) Claim(proposalID *big.Int, account common.Address) (*types.Transaction, error) {
	return _BmErc20.Contract.Claim(&_BmErc20.TransactOpts, proposalID, account)
}

// Claim is a paid mutator transaction binding the contract method 0xddd5e1b2.
//
// Solidity: function claim(uint256 proposalID, address account) returns()
func (_BmErc20 *BmErc20TransactorSession) Claim(proposalID *big.Int, account common.Address) (*types.Transaction, error) {
	return _BmErc20.Contract.Claim(&_BmErc20.TransactOpts, proposalID, account)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address account) payable returns()
func (_BmErc20 *BmErc20Transactor) Mint(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BmErc20.contract.Transact(opts, "mint", account)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address account) payable returns()
func (_BmErc20 *BmErc20Session) Mint(account common.Address) (*types.Transaction, error) {
	return _BmErc20.Contract.Mint(&_BmErc20.TransactOpts, account)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address account) payable returns()
func (_BmErc20 *BmErc20TransactorSession) Mint(account common.Address) (*types.Transaction, error) {
	return _BmErc20.Contract.Mint(&_BmErc20.TransactOpts, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BmErc20 *BmErc20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BmErc20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BmErc20 *BmErc20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BmErc20.Contract.Transfer(&_BmErc20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BmErc20 *BmErc20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BmErc20.Contract.Transfer(&_BmErc20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BmErc20 *BmErc20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BmErc20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BmErc20 *BmErc20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BmErc20.Contract.TransferFrom(&_BmErc20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BmErc20 *BmErc20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BmErc20.Contract.TransferFrom(&_BmErc20.TransactOpts, from, to, value)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_BmErc20 *BmErc20Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BmErc20.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_BmErc20 *BmErc20Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BmErc20.Contract.Fallback(&_BmErc20.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_BmErc20 *BmErc20TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BmErc20.Contract.Fallback(&_BmErc20.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BmErc20 *BmErc20Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BmErc20.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BmErc20 *BmErc20Session) Receive() (*types.Transaction, error) {
	return _BmErc20.Contract.Receive(&_BmErc20.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BmErc20 *BmErc20TransactorSession) Receive() (*types.Transaction, error) {
	return _BmErc20.Contract.Receive(&_BmErc20.TransactOpts)
}

// BmErc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BmErc20 contract.
type BmErc20ApprovalIterator struct {
	Event *BmErc20Approval // Event containing the contract specifics and raw log

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
func (it *BmErc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BmErc20Approval)
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
		it.Event = new(BmErc20Approval)
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
func (it *BmErc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BmErc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BmErc20Approval represents a Approval event raised by the BmErc20 contract.
type BmErc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BmErc20 *BmErc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BmErc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BmErc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BmErc20ApprovalIterator{contract: _BmErc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BmErc20 *BmErc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BmErc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BmErc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BmErc20Approval)
				if err := _BmErc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BmErc20 *BmErc20Filterer) ParseApproval(log types.Log) (*BmErc20Approval, error) {
	event := new(BmErc20Approval)
	if err := _BmErc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BmErc20ClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the BmErc20 contract.
type BmErc20ClaimedIterator struct {
	Event *BmErc20Claimed // Event containing the contract specifics and raw log

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
func (it *BmErc20ClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BmErc20Claimed)
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
		it.Event = new(BmErc20Claimed)
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
func (it *BmErc20ClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BmErc20ClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BmErc20Claimed represents a Claimed event raised by the BmErc20 contract.
type BmErc20Claimed struct {
	ProposalID *big.Int
	Account    common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x4ec90e965519d92681267467f775ada5bd214aa92c0dc93d90a5e880ce9ed026.
//
// Solidity: event Claimed(uint256 indexed proposalID, address indexed account, uint256 indexed amount)
func (_BmErc20 *BmErc20Filterer) FilterClaimed(opts *bind.FilterOpts, proposalID []*big.Int, account []common.Address, amount []*big.Int) (*BmErc20ClaimedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _BmErc20.contract.FilterLogs(opts, "Claimed", proposalIDRule, accountRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &BmErc20ClaimedIterator{contract: _BmErc20.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x4ec90e965519d92681267467f775ada5bd214aa92c0dc93d90a5e880ce9ed026.
//
// Solidity: event Claimed(uint256 indexed proposalID, address indexed account, uint256 indexed amount)
func (_BmErc20 *BmErc20Filterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *BmErc20Claimed, proposalID []*big.Int, account []common.Address, amount []*big.Int) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _BmErc20.contract.WatchLogs(opts, "Claimed", proposalIDRule, accountRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BmErc20Claimed)
				if err := _BmErc20.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x4ec90e965519d92681267467f775ada5bd214aa92c0dc93d90a5e880ce9ed026.
//
// Solidity: event Claimed(uint256 indexed proposalID, address indexed account, uint256 indexed amount)
func (_BmErc20 *BmErc20Filterer) ParseClaimed(log types.Log) (*BmErc20Claimed, error) {
	event := new(BmErc20Claimed)
	if err := _BmErc20.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BmErc20MintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the BmErc20 contract.
type BmErc20MintedIterator struct {
	Event *BmErc20Minted // Event containing the contract specifics and raw log

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
func (it *BmErc20MintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BmErc20Minted)
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
		it.Event = new(BmErc20Minted)
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
func (it *BmErc20MintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BmErc20MintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BmErc20Minted represents a Minted event raised by the BmErc20 contract.
type BmErc20Minted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x90ddedd5a25821bba11fbb98de02ec1f75c1be90ae147d6450ce873e7b78b5d8.
//
// Solidity: event Minted(address indexed account)
func (_BmErc20 *BmErc20Filterer) FilterMinted(opts *bind.FilterOpts, account []common.Address) (*BmErc20MintedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BmErc20.contract.FilterLogs(opts, "Minted", accountRule)
	if err != nil {
		return nil, err
	}
	return &BmErc20MintedIterator{contract: _BmErc20.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x90ddedd5a25821bba11fbb98de02ec1f75c1be90ae147d6450ce873e7b78b5d8.
//
// Solidity: event Minted(address indexed account)
func (_BmErc20 *BmErc20Filterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *BmErc20Minted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BmErc20.contract.WatchLogs(opts, "Minted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BmErc20Minted)
				if err := _BmErc20.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x90ddedd5a25821bba11fbb98de02ec1f75c1be90ae147d6450ce873e7b78b5d8.
//
// Solidity: event Minted(address indexed account)
func (_BmErc20 *BmErc20Filterer) ParseMinted(log types.Log) (*BmErc20Minted, error) {
	event := new(BmErc20Minted)
	if err := _BmErc20.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BmErc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BmErc20 contract.
type BmErc20TransferIterator struct {
	Event *BmErc20Transfer // Event containing the contract specifics and raw log

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
func (it *BmErc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BmErc20Transfer)
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
		it.Event = new(BmErc20Transfer)
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
func (it *BmErc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BmErc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BmErc20Transfer represents a Transfer event raised by the BmErc20 contract.
type BmErc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BmErc20 *BmErc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BmErc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BmErc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BmErc20TransferIterator{contract: _BmErc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BmErc20 *BmErc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BmErc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BmErc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BmErc20Transfer)
				if err := _BmErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BmErc20 *BmErc20Filterer) ParseTransfer(log types.Log) (*BmErc20Transfer, error) {
	event := new(BmErc20Transfer)
	if err := _BmErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
