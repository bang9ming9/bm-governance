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

// BmGovernorMetaData contains all meta data concerning the BmGovernor contract.
var BmGovernorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"erc1155_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorAlreadyCastVote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorAlreadyQueuedProposal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorDisabledDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"GovernorInsufficientProposerVotes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldatas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"values\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidProposalLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorInvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorInvalidVoteType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"votingPeriod\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidVotingPeriod\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNonexistentProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNotQueuedProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyExecutor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyProposer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorQueueNotImplemented\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"GovernorRestrictedProposer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"current\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"expectedStates\",\"type\":\"bytes32\"}],\"name\":\"GovernorUnexpectedProposalState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueFull\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"etaSeconds\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"VoteCastWithParams\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BM_ERC1155\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLOCK_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTING_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castVoteWithReason\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParamsBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clock\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"getVotesWithParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalEta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalNeedsQueuing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC5805\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101a060405234801562000011575f80fd5b50604051620037c0380380620037c083398101604081905262000034916200020b565b808280620000566040805180820190915260018152603160f81b602082015290565b62000062825f62000134565b610120526200007381600162000134565b61014052815160208084019190912060e052815190820120610100524660a0526200010060e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c05260036200011782826200035d565b50506001600160a01b039081166101605216610180525062000481565b5f60208351101562000153576200014b836200016c565b905062000166565b816200016084826200035d565b5060ff90505b92915050565b5f80829050601f81511115620001a2578260405163305a27a960e01b815260040162000199919062000429565b60405180910390fd5b8051620001af826200045d565b179392505050565b634e487b7160e01b5f52604160045260245ffd5b5f5b83811015620001e7578181015183820152602001620001cd565b50505f910152565b80516001600160a01b038116811462000206575f80fd5b919050565b5f80604083850312156200021d575f80fd5b82516001600160401b038082111562000234575f80fd5b818501915085601f83011262000248575f80fd5b8151818111156200025d576200025d620001b7565b604051601f8201601f19908116603f01168101908382118183101715620002885762000288620001b7565b81604052828152886020848701011115620002a1575f80fd5b620002b4836020830160208801620001cb565b8096505050505050620002ca60208401620001ef565b90509250929050565b600181811c90821680620002e857607f821691505b6020821081036200030757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156200035857805f5260205f20601f840160051c81016020851015620003345750805b601f840160051c820191505b8181101562000355575f815560010162000340565b50505b505050565b81516001600160401b03811115620003795762000379620001b7565b62000391816200038a8454620002d3565b846200030d565b602080601f831160018114620003c7575f8415620003af5750858301515b5f19600386901b1c1916600185901b17855562000421565b5f85815260208120601f198616915b82811015620003f757888601518255948401946001909101908401620003d6565b50858210156200041557878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b602081525f825180602084015262000449816040850160208701620001cb565b601f01601f19169190910160400192915050565b8051602080830151919081101562000307575f1960209190910360031b1b16919050565b60805160a05160c05160e05161010051610120516101405161016051610180516132b66200050a5f395f8181610736015281816111be015261163001525f8181610816015281816114ad015281816117aa0152611f8e01525f61178101525f61175501525f61199c01525f61197401525f6118cf01525f6118f901525f61192301526132b65ff3fe608060405260043610610236575f3560e01c80637d5e81e211610129578063c01f9e37116100a8578063deaaa7cc1161006d578063deaaa7cc1461076c578063eb9019d41461079f578063f23a6e61146107be578063f8ce560a146107e9578063fc0c546a14610808575f80fd5b8063c01f9e37146106d4578063c28bc2fa146106f3578063c59057e414610706578063c766c11414610725578063dd4e2ba514610758575f80fd5b80639a802a6d116100ee5780639a802a6d1461061b578063a9a952941461063a578063ab58fb8e14610659578063b58131b01461068f578063bc197c81146106a9575f80fd5b80637d5e81e2146105575780637ecebe001461057657806384b0196e146105aa5780638ff262e3146105d157806391ddadf4146105f0575f80fd5b80633e4f49e6116101b557806354fd4d501161017a57806354fd4d50146104b257806356781388146104db5780635b8d0e0d146104fa5780635f398a14146105195780637b3c71d314610538575f80fd5b80633e4f49e6146103e1578063438596321461040d578063452115d61461042c5780634bf5d7e91461044b578063544ffc9c1461045f575f80fd5b8063160cbed7116101fb578063160cbed7146103495780632656227d146103685780632d63f6931461037b5780632fe3e2611461039a5780633932abb1146103cd575f80fd5b806301ffc9a71461024357806302a251a31461027757806306fdde0314610299578063143489d0146102ba578063150b7a0214610306575f80fd5b3661023f57005b005b5f80fd5b34801561024e575f80fd5b5061026261025d3660046123f6565b61083a565b60405190151581526020015b60405180910390f35b348015610282575f80fd5b5061028b61088b565b60405190815260200161026e565b3480156102a4575f80fd5b506102ad6108a1565b60405161026e919061246a565b3480156102c5575f80fd5b506102ee6102d436600461247c565b5f908152600460205260409020546001600160a01b031690565b6040516001600160a01b03909116815260200161026e565b348015610311575f80fd5b5061033061032036600461256c565b630a85bd0160e11b949350505050565b6040516001600160e01b0319909116815260200161026e565b348015610354575f80fd5b5061028b610363366004612735565b610931565b61028b610376366004612735565b610979565b348015610386575f80fd5b5061028b61039536600461247c565b610aa1565b3480156103a5575f80fd5b5061028b7f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a81181565b3480156103d8575f80fd5b5061028b610ac1565b3480156103ec575f80fd5b506104006103fb36600461247c565b610ae7565b60405161026e91906127f2565b348015610418575f80fd5b50610262610427366004612800565b610c10565b348015610437575f80fd5b5061028b610446366004612735565b610c40565b348015610456575f80fd5b506102ad610cac565b34801561046a575f80fd5b5061049761047936600461247c565b5f908152600760205260409020805460018201546002909201549092565b6040805193845260208401929092529082015260600161026e565b3480156104bd575f80fd5b506040805180820190915260018152603160f81b60208201526102ad565b3480156104e6575f80fd5b5061028b6104f536600461283a565b610cb6565b348015610505575f80fd5b5061028b61051436600461289f565b610cdd565b348015610524575f80fd5b5061028b61053336600461294f565b610e39565b348015610543575f80fd5b5061028b6105523660046129cc565b610e8c565b348015610562575f80fd5b5061028b610571366004612a21565b610ed2565b348015610581575f80fd5b5061028b610590366004612acd565b6001600160a01b03165f9081526002602052604090205490565b3480156105b5575f80fd5b506105be610f23565b60405161026e9796959493929190612b20565b3480156105dc575f80fd5b5061028b6105eb366004612b8f565b610f65565b3480156105fb575f80fd5b50610604611034565b60405165ffffffffffff909116815260200161026e565b348015610626575f80fd5b5061028b610635366004612bda565b61103d565b348015610645575f80fd5b5061026261065436600461247c565b505f90565b348015610664575f80fd5b5061028b61067336600461247c565b5f9081526004602052604090206001015465ffffffffffff1690565b34801561069a575f80fd5b50670de0b6b3a764000061028b565b3480156106b4575f80fd5b506103306106c3366004612c2c565b63bc197c8160e01b95945050505050565b3480156106df575f80fd5b5061028b6106ee36600461247c565b611049565b61023d610701366004612cb4565b61108b565b348015610711575f80fd5b5061028b610720366004612735565b611107565b348015610730575f80fd5b506102ee7f000000000000000000000000000000000000000000000000000000000000000081565b348015610763575f80fd5b506102ad611140565b348015610777575f80fd5b5061028b7ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d781565b3480156107aa575f80fd5b5061028b6107b9366004612cf1565b61117c565b3480156107c9575f80fd5b506103306107d8366004612d19565b63f23a6e6160e01b95945050505050565b3480156107f4575f80fd5b5061028b61080336600461247c565b61119b565b348015610813575f80fd5b507f00000000000000000000000000000000000000000000000000000000000000006102ee565b5f6001600160e01b031982166332a2ad4360e11b148061086a57506001600160e01b03198216630271189760e51b145b8061088557506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f61089c6205460062093a80612d8c565b905090565b6060600380546108b090612d9f565b80601f01602080910402602001604051908101604052809291908181526020018280546108dc90612d9f565b80156109275780601f106108fe57610100808354040283529160200191610927565b820191905f5260205f20905b81548152906001019060200180831161090a57829003601f168201915b5050505050905090565b5f8061093f86868686611107565b90506109548161094f6004611231565b611253565b505f604051634844252360e11b815260040160405180910390fd5b5095945050505050565b5f8061098786868686611107565b90506109a7816109976005611231565b6109a16004611231565b17611253565b505f818152600460205260409020805460ff60f01b1916600160f01b179055306109ce3090565b6001600160a01b031614610a57575f5b8651811015610a5557306001600160a01b0316878281518110610a0357610a03612dd7565b60200260200101516001600160a01b031603610a4d57610a4d858281518110610a2e57610a2e612dd7565b602002602001015180519060200120600561129090919063ffffffff16565b6001016109de565b505b610a648187878787611300565b6040518181527f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f906020015b60405180910390a195945050505050565b5f90815260046020526040902054600160a01b900465ffffffffffff1690565b5f4262093a80810681036205460001620151808201811015610ae1575f80fd5b03919050565b5f818152600460205260408120805460ff600160f01b8204811691600160f81b9004168115610b1b57506007949350505050565b8015610b2c57506002949350505050565b5f610b3686610aa1565b9050805f03610b6057604051636ad0607560e01b8152600481018790526024015b60405180910390fd5b5f610b69611034565b65ffffffffffff169050808210610b8657505f9695505050505050565b5f610b9088611049565b9050818110610ba757506001979650505050505050565b610bb0886113d5565b1580610bc25750610bc0886113df565b155b15610bd557506003979650505050505050565b5f8881526004602052604090206001015465ffffffffffff165f03610c0257506004979650505050505050565b506005979650505050505050565b5f8281526007602090815260408083206001600160a01b038516845260030190915281205460ff165b9392505050565b5f80610c4e86868686611107565b9050610c5d8161094f5f611231565b505f818152600460205260409020546001600160a01b03163314610c965760405163233d98e360e01b8152336004820152602401610b57565b610ca2868686866113fa565b9695505050505050565b606061089c6114a9565b5f80339050610cd584828560405180602001604052805f815250611569565b949350505050565b5f80610dbe87610db87f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a8118c8c8c610d308e6001600160a01b03165f90815260026020526040902080546001810190915590565b8d8d604051610d40929190612deb565b60405180910390208c80519060200120604051602001610d9d9796959493929190968752602087019590955260ff9390931660408601526001600160a01b03919091166060850152608084015260a083015260c082015260e00190565b6040516020818303038152906040528051906020012061158e565b856115ba565b905080610de9576040516394ab6c0760e01b81526001600160a01b0388166004820152602401610b57565b610e2c89888a89898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508b925061160f915050565b9998505050505050505050565b5f80339050610e8187828888888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508a925061160f915050565b979650505050505050565b5f80339050610ca286828787878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061156992505050565b5f80610ee086868686611696565b905061096f8133600260405180604001604052806008815260200167383937b837b9b2b960c11b815250610f1e60408051602081019091525f815290565b61160f565b5f6060805f805f6060610f3461174e565b610f3c61177a565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b5f80610fef84610db87ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d7898989610fb88b6001600160a01b03165f90815260026020526040902080546001810190915590565b60408051602081019690965285019390935260ff90911660608401526001600160a01b0316608083015260a082015260c001610d9d565b90508061101a576040516394ab6c0760e01b81526001600160a01b0385166004820152602401610b57565b610ca286858760405180602001604052805f815250611569565b5f61089c6117a7565b5f610cd584848461182e565b5f8181526004602052604081205461107d90600160d01b810463ffffffff1690600160a01b900465ffffffffffff16612dfa565b65ffffffffffff1692915050565b61109361183a565b5f80856001600160a01b03168585856040516110b0929190612deb565b5f6040518083038185875af1925050503d805f81146110ea576040519150601f19603f3d011682016040523d82523d5f602084013e6110ef565b606091505b50915091506110fe8282611871565b50505050505050565b5f8484848460405160200161111f9493929190612eb0565b60408051601f19818403018152919052805160209091012095945050505050565b606061089c6040805180820190915260208082527f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e9082015290565b5f610c39838361119660408051602081019091525f815290565b61182e565b604051632394e7a360e21b8152600481018290525f90600a906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638e539e8c90602401602060405180830381865afa158015611203573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112279190612efa565b6108859190612f11565b5f816007811115611244576112446127be565b600160ff919091161b92915050565b5f8061125e84610ae7565b90505f8361126b83611231565b1603610c39578381846040516331b75e4d60e01b8152600401610b5793929190612f30565b81546001600160801b03600160801b8204811691811660018301909116036112cb57604051638acb5f2760e01b815260040160405180910390fd5b6001600160801b038082165f90815260018086016020526040909120939093558354919092018216600160801b029116179055565b5f5b84518110156113cd575f8086838151811061131f5761131f612dd7565b60200260200101516001600160a01b031686848151811061134257611342612dd7565b602002602001015186858151811061135c5761135c612dd7565b60200260200101516040516113719190612f52565b5f6040518083038185875af1925050503d805f81146113ab576040519150601f19603f3d011682016040523d82523d5f602084013e6113b0565b606091505b50915091506113bf8282611871565b505050806001019050611302565b505050505050565b5f6108858261188d565b5f818152600760205260408120805460019091015411610885565b5f8061140886868686611107565b9050611456816114186007611231565b6114226006611231565b61142c6002611231565b6001611439600782612f6d565b611444906002613066565b61144e9190612d8c565b181818611253565b505f818152600460205260409081902080546001600160f81b0316600160f81b179055517f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c90610a909083815260200190565b60607f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634bf5d7e96040518163ffffffff1660e01b81526004015f60405180830381865afa92505050801561152857506040513d5f823e601f3d908101601f191682016040526115259190810190613074565b60015b611564575060408051808201909152601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000602082015290565b919050565b5f61158585858585610f1e60408051602081019091525f815290565b95945050505050565b5f61088561159a6118c3565b8360405161190160f01b8152600281019290925260228201526042902090565b5f805f6115c785856119ec565b5090925090505f8160038111156115e0576115e06127be565b1480156115fe5750856001600160a01b0316826001600160a01b0316145b80610ca25750610ca2868686611a35565b6040516306044ae960e41b81526001600160a01b0385811660048301525f917f000000000000000000000000000000000000000000000000000000000000000090911690636044ae90906024015f604051808303815f87803b158015611673575f80fd5b505af1158015611685573d5f803e3d5ffd5b50505050610ca28686868686611b0b565b5f336116a28184611be5565b6116ca5760405163d9b3955760e01b81526001600160a01b0382166004820152602401610b57565b5f6116f08260016116d9611034565b6116e391906130dc565b65ffffffffffff1661117c565b9050670de0b6b3a76400008082101561173557604051636121770b60e11b81526001600160a01b03841660048201526024810183905260448101829052606401610b57565b6117428888888887611ccd565b98975050505050505050565b606061089c7f00000000000000000000000000000000000000000000000000000000000000005f611ed8565b606061089c7f00000000000000000000000000000000000000000000000000000000000000006001611ed8565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611822575060408051601f3d908101601f1916820190925261181f918101906130fb565b60015b6115645761089c611f81565b5f610cd5848484611f8b565b30331461185c576040516347096e4760e01b8152336004820152602401610b57565b565b80611869600561201e565b0361185e5750565b606082611886576118818261209a565b610885565b5080610885565b5f818152600760205260408120600281015460018201546118ae9190613120565b6118ba61080385610aa1565b11159392505050565b5f306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614801561191b57507f000000000000000000000000000000000000000000000000000000000000000046145b1561194557507f000000000000000000000000000000000000000000000000000000000000000090565b61089c604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f805f8351604103611a23576020840151604085015160608601515f1a611a15888285856120c3565b955095509550505050611a2e565b505081515f91506002905b9250925092565b5f805f856001600160a01b03168585604051602401611a55929190613133565b60408051601f198184030181529181526020820180516001600160e01b0316630b135d3f60e11b17905251611a8a9190612f52565b5f60405180830381855afa9150503d805f8114611ac2576040519150601f19603f3d011682016040523d82523d5f602084013e611ac7565b606091505b5091509150818015611adb57506020815110155b8015610ca257508051630b135d3f60e11b90611b009083016020908101908401612efa565b149695505050505050565b5f611b1a8661094f6001611231565b505f611b2f86611b2989610aa1565b8561182e565b9050611b3e878787848761218b565b82515f03611b9257856001600160a01b03167fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda488878488604051611b85949392919061314b565b60405180910390a2610ca2565b856001600160a01b03167fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb87128887848888604051611bd3959493929190613172565b60405180910390a29695505050505050565b80515f906034811015611bfc576001915050610885565b82810160131901516001600160a01b031981166b046e0e4dee0dee6cae47a60f60a31b14611c2f57600192505050610885565b5f80611c3c602885612d8c565b90505b83811015611cac575f80611c72888481518110611c5e57611c5e612dd7565b01602001516001600160f81b0319166121aa565b9150915081611c8a5760019650505050505050610885565b8060ff166004856001600160a01b0316901b1793505050806001019050611c3f565b50856001600160a01b0316816001600160a01b031614935050505092915050565b5f611ce18686868680519060200120611107565b905084518651141580611cf657508351865114155b80611d0057508551155b15611d3557855184518651604051630447b05d60e41b8152600481019390935260248301919091526044820152606401610b57565b5f81815260046020526040902054600160a01b900465ffffffffffff1615611d7e5780611d6182610ae7565b6040516331b75e4d60e01b8152610b579291905f90600401612f30565b5f611d87610ac1565b611d8f611034565b65ffffffffffff16611da19190613120565b90505f611dac61088b565b5f84815260046020526040902080546001600160a01b0319166001600160a01b038716178155909150611dde8361223a565b815465ffffffffffff91909116600160a01b0265ffffffffffff60a01b19909116178155611e0b82612270565b815463ffffffff91909116600160d01b0263ffffffff60d01b1990911617815588517f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e090859087908c908c906001600160401b03811115611e6e57611e6e6124a9565b604051908082528060200260200182016040528015611ea157816020015b6060815260200190600190039081611e8c5790505b508c89611eae8a82613120565b8e604051611ec4999897969594939291906131ab565b60405180910390a150505095945050505050565b606060ff8314611ef257611eeb836122a0565b9050610885565b818054611efe90612d9f565b80601f0160208091040260200160405190810160405280929190818152602001828054611f2a90612d9f565b8015611f755780601f10611f4c57610100808354040283529160200191611f75565b820191905f5260205f20905b815481529060010190602001808311611f5857829003601f168201915b50505050509050610885565b5f61089c4361223a565b5f7f0000000000000000000000000000000000000000000000000000000000000000604051630748d63560e31b81526001600160a01b038681166004830152602482018690529190911690633a46b1a890604401602060405180830381865afa158015611ffa573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cd59190612efa565b80545f906001600160801b0380821691600160801b9004168103612055576040516375e52f4f60e01b815260040160405180910390fd5b6001600160801b038181165f908152600185810160205260408220805492905585546fffffffffffffffffffffffffffffffff19169301909116919091179092555090565b8051156120aa5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156120fc57505f91506003905082612181565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa15801561214d573d5f803e3d5ffd5b5050604051601f1901519150506001600160a01b03811661217857505f925060019150829050612181565b92505f91508190505b9450945094915050565b815f03612196575f80fd5b6121a385858585856122dd565b5050505050565b5f8060f883901c602f811180156121c45750603a8160ff16105b156121d957600194602f199091019350915050565b8060ff1660401080156121ef575060478160ff16105b15612204576001946036199091019350915050565b8060ff16606010801561221a575060678160ff16105b1561222f576001946056199091019350915050565b505f93849350915050565b5f65ffffffffffff82111561226c576040516306dfcc6560e41b81526030600482015260248101839052604401610b57565b5090565b5f63ffffffff82111561226c576040516306dfcc6560e41b81526020600482015260248101839052604401610b57565b60605f6122ac836123cf565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f8581526007602090815260408083206001600160a01b0388168452600381019092529091205460ff1615612330576040516371c6af4960e01b81526001600160a01b0386166004820152602401610b57565b6001600160a01b0385165f9081526003820160205260409020805460ff1916600117905560ff84166123795782815f015f82825461236e9190613120565b909155506113cd9050565b5f1960ff8516016123975782816001015f82825461236e9190613120565b60011960ff8516016123b65782816002015f82825461236e9190613120565b6040516303599be160e11b815260040160405180910390fd5b5f60ff8216601f81111561088557604051632cd44ac360e21b815260040160405180910390fd5b5f60208284031215612406575f80fd5b81356001600160e01b031981168114610c39575f80fd5b5f5b8381101561243757818101518382015260200161241f565b50505f910152565b5f815180845261245681602086016020860161241d565b601f01601f19169290920160200192915050565b602081525f610c39602083018461243f565b5f6020828403121561248c575f80fd5b5035919050565b80356001600160a01b0381168114611564575f80fd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156124e5576124e56124a9565b604052919050565b5f6001600160401b03821115612505576125056124a9565b50601f01601f191660200190565b5f612525612520846124ed565b6124bd565b9050828152838383011115612538575f80fd5b828260208301375f602084830101529392505050565b5f82601f83011261255d575f80fd5b610c3983833560208501612513565b5f805f806080858703121561257f575f80fd5b61258885612493565b935061259660208601612493565b92506040850135915060608501356001600160401b038111156125b7575f80fd5b6125c38782880161254e565b91505092959194509250565b5f6001600160401b038211156125e7576125e76124a9565b5060051b60200190565b5f82601f830112612600575f80fd5b81356020612610612520836125cf565b8083825260208201915060208460051b870101935086841115612631575f80fd5b602086015b848110156126545761264781612493565b8352918301918301612636565b509695505050505050565b5f82601f83011261266e575f80fd5b8135602061267e612520836125cf565b8083825260208201915060208460051b87010193508684111561269f575f80fd5b602086015b8481101561265457803583529183019183016126a4565b5f82601f8301126126ca575f80fd5b813560206126da612520836125cf565b82815260059290921b840181019181810190868411156126f8575f80fd5b8286015b848110156126545780356001600160401b03811115612719575f80fd5b6127278986838b010161254e565b8452509183019183016126fc565b5f805f8060808587031215612748575f80fd5b84356001600160401b038082111561275e575f80fd5b61276a888389016125f1565b9550602087013591508082111561277f575f80fd5b61278b8883890161265f565b945060408701359150808211156127a0575f80fd5b506127ad878288016126bb565b949793965093946060013593505050565b634e487b7160e01b5f52602160045260245ffd5b600881106127ee57634e487b7160e01b5f52602160045260245ffd5b9052565b6020810161088582846127d2565b5f8060408385031215612811575f80fd5b8235915061282160208401612493565b90509250929050565b803560ff81168114611564575f80fd5b5f806040838503121561284b575f80fd5b823591506128216020840161282a565b5f8083601f84011261286b575f80fd5b5081356001600160401b03811115612881575f80fd5b602083019150836020828501011115612898575f80fd5b9250929050565b5f805f805f805f60c0888a0312156128b5575f80fd5b873596506128c56020890161282a565b95506128d360408901612493565b945060608801356001600160401b03808211156128ee575f80fd5b6128fa8b838c0161285b565b909650945060808a0135915080821115612912575f80fd5b61291e8b838c0161254e565b935060a08a0135915080821115612933575f80fd5b506129408a828b0161254e565b91505092959891949750929550565b5f805f805f60808688031215612963575f80fd5b853594506129736020870161282a565b935060408601356001600160401b038082111561298e575f80fd5b61299a89838a0161285b565b909550935060608801359150808211156129b2575f80fd5b506129bf8882890161254e565b9150509295509295909350565b5f805f80606085870312156129df575f80fd5b843593506129ef6020860161282a565b925060408501356001600160401b03811115612a09575f80fd5b612a158782880161285b565b95989497509550505050565b5f805f8060808587031215612a34575f80fd5b84356001600160401b0380821115612a4a575f80fd5b612a56888389016125f1565b95506020870135915080821115612a6b575f80fd5b612a778883890161265f565b94506040870135915080821115612a8c575f80fd5b612a98888389016126bb565b93506060870135915080821115612aad575f80fd5b508501601f81018713612abe575f80fd5b6125c387823560208401612513565b5f60208284031215612add575f80fd5b610c3982612493565b5f815180845260208085019450602084015f5b83811015612b1557815187529582019590820190600101612af9565b509495945050505050565b60ff60f81b8816815260e060208201525f612b3e60e083018961243f565b8281036040840152612b50818961243f565b606084018890526001600160a01b038716608085015260a0840186905283810360c08501529050612b818185612ae6565b9a9950505050505050505050565b5f805f8060808587031215612ba2575f80fd5b84359350612bb26020860161282a565b9250612bc060408601612493565b915060608501356001600160401b038111156125b7575f80fd5b5f805f60608486031215612bec575f80fd5b612bf584612493565b92506020840135915060408401356001600160401b03811115612c16575f80fd5b612c228682870161254e565b9150509250925092565b5f805f805f60a08688031215612c40575f80fd5b612c4986612493565b9450612c5760208701612493565b935060408601356001600160401b0380821115612c72575f80fd5b612c7e89838a0161265f565b94506060880135915080821115612c93575f80fd5b612c9f89838a0161265f565b935060808801359150808211156129b2575f80fd5b5f805f8060608587031215612cc7575f80fd5b612cd085612493565b93506020850135925060408501356001600160401b03811115612a09575f80fd5b5f8060408385031215612d02575f80fd5b612d0b83612493565b946020939093013593505050565b5f805f805f60a08688031215612d2d575f80fd5b612d3686612493565b9450612d4460208701612493565b9350604086013592506060860135915060808601356001600160401b03811115612d6c575f80fd5b6129bf8882890161254e565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561088557610885612d78565b600181811c90821680612db357607f821691505b602082108103612dd157634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b818382375f9101908152919050565b65ffffffffffff818116838216019080821115612e1957612e19612d78565b5092915050565b5f815180845260208085019450602084015f5b83811015612b155781516001600160a01b031687529582019590820190600101612e33565b5f8282518085526020808601955060208260051b840101602086015f5b84811015612ea357601f19868403018952612e9183835161243f565b98840198925090830190600101612e75565b5090979650505050505050565b608081525f612ec26080830187612e20565b8281036020840152612ed48187612ae6565b90508281036040840152612ee88186612e58565b91505082606083015295945050505050565b5f60208284031215612f0a575f80fd5b5051919050565b5f82612f2b57634e487b7160e01b5f52601260045260245ffd5b500490565b83815260608101612f4460208301856127d2565b826040830152949350505050565b5f8251612f6381846020870161241d565b9190910192915050565b60ff818116838216019081111561088557610885612d78565b600181815b80851115612fc057815f1904821115612fa657612fa6612d78565b80851615612fb357918102915b93841c9390800290612f8b565b509250929050565b5f82612fd657506001610885565b81612fe257505f610885565b8160018114612ff857600281146130025761301e565b6001915050610885565b60ff84111561301357613013612d78565b50506001821b610885565b5060208310610133831016604e8410600b8410161715613041575081810a610885565b61304b8383612f86565b805f190482111561305e5761305e612d78565b029392505050565b5f610c3960ff841683612fc8565b5f60208284031215613084575f80fd5b81516001600160401b03811115613099575f80fd5b8201601f810184136130a9575f80fd5b80516130b7612520826124ed565b8181528560208385010111156130cb575f80fd5b61158582602083016020860161241d565b65ffffffffffff828116828216039080821115612e1957612e19612d78565b5f6020828403121561310b575f80fd5b815165ffffffffffff81168114610c39575f80fd5b8082018082111561088557610885612d78565b828152604060208201525f610cd5604083018461243f565b84815260ff84166020820152826040820152608060608201525f610ca2608083018461243f565b85815260ff8516602082015283604082015260a060608201525f61319960a083018561243f565b8281036080840152611742818561243f565b5f6101208b8352602060018060a01b038c16818501528160408501526131d38285018c612e20565b915083820360608501526131e7828b612ae6565b915083820360808501528189518084528284019150828160051b850101838c015f5b8381101561323757601f1987840301855261322583835161243f565b94860194925090850190600101613209565b505086810360a088015261324b818c612e58565b9450505050508560c08401528460e0840152828103610100840152613270818561243f565b9c9b50505050505050505050505056fea2646970667358221220700822bd739eba5e0ce90ad82e439f311c2b859b45ed045e566742dc9dff1ac264736f6c63430008180033",
}

// BmGovernorABI is the input ABI used to generate the binding from.
// Deprecated: Use BmGovernorMetaData.ABI instead.
var BmGovernorABI = BmGovernorMetaData.ABI

// BmGovernorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BmGovernorMetaData.Bin instead.
var BmGovernorBin = BmGovernorMetaData.Bin

// DeployBmGovernor deploys a new Ethereum contract, binding an instance of BmGovernor to it.
func DeployBmGovernor(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, erc1155_ common.Address) (common.Address, *types.Transaction, *BmGovernor, error) {
	parsed, err := BmGovernorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BmGovernorBin), backend, name_, erc1155_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BmGovernor{BmGovernorCaller: BmGovernorCaller{contract: contract}, BmGovernorTransactor: BmGovernorTransactor{contract: contract}, BmGovernorFilterer: BmGovernorFilterer{contract: contract}}, nil
}

// BmGovernor is an auto generated Go binding around an Ethereum contract.
type BmGovernor struct {
	BmGovernorCaller     // Read-only binding to the contract
	BmGovernorTransactor // Write-only binding to the contract
	BmGovernorFilterer   // Log filterer for contract events
}

// BmGovernorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BmGovernorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BmGovernorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BmGovernorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BmGovernorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BmGovernorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BmGovernorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BmGovernorSession struct {
	Contract     *BmGovernor       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BmGovernorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BmGovernorCallerSession struct {
	Contract *BmGovernorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BmGovernorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BmGovernorTransactorSession struct {
	Contract     *BmGovernorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BmGovernorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BmGovernorRaw struct {
	Contract *BmGovernor // Generic contract binding to access the raw methods on
}

// BmGovernorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BmGovernorCallerRaw struct {
	Contract *BmGovernorCaller // Generic read-only contract binding to access the raw methods on
}

// BmGovernorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BmGovernorTransactorRaw struct {
	Contract *BmGovernorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBmGovernor creates a new instance of BmGovernor, bound to a specific deployed contract.
func NewBmGovernor(address common.Address, backend bind.ContractBackend) (*BmGovernor, error) {
	contract, err := bindBmGovernor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BmGovernor{BmGovernorCaller: BmGovernorCaller{contract: contract}, BmGovernorTransactor: BmGovernorTransactor{contract: contract}, BmGovernorFilterer: BmGovernorFilterer{contract: contract}}, nil
}

// NewBmGovernorCaller creates a new read-only instance of BmGovernor, bound to a specific deployed contract.
func NewBmGovernorCaller(address common.Address, caller bind.ContractCaller) (*BmGovernorCaller, error) {
	contract, err := bindBmGovernor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BmGovernorCaller{contract: contract}, nil
}

// NewBmGovernorTransactor creates a new write-only instance of BmGovernor, bound to a specific deployed contract.
func NewBmGovernorTransactor(address common.Address, transactor bind.ContractTransactor) (*BmGovernorTransactor, error) {
	contract, err := bindBmGovernor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BmGovernorTransactor{contract: contract}, nil
}

// NewBmGovernorFilterer creates a new log filterer instance of BmGovernor, bound to a specific deployed contract.
func NewBmGovernorFilterer(address common.Address, filterer bind.ContractFilterer) (*BmGovernorFilterer, error) {
	contract, err := bindBmGovernor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BmGovernorFilterer{contract: contract}, nil
}

// bindBmGovernor binds a generic wrapper to an already deployed contract.
func bindBmGovernor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BmGovernorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BmGovernor *BmGovernorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BmGovernor.Contract.BmGovernorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BmGovernor *BmGovernorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BmGovernor.Contract.BmGovernorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BmGovernor *BmGovernorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BmGovernor.Contract.BmGovernorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BmGovernor *BmGovernorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BmGovernor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BmGovernor *BmGovernorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BmGovernor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BmGovernor *BmGovernorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BmGovernor.Contract.contract.Transact(opts, method, params...)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_BmGovernor *BmGovernorCaller) BALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_BmGovernor *BmGovernorSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _BmGovernor.Contract.BALLOTTYPEHASH(&_BmGovernor.CallOpts)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_BmGovernor *BmGovernorCallerSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _BmGovernor.Contract.BALLOTTYPEHASH(&_BmGovernor.CallOpts)
}

// BMERC1155 is a free data retrieval call binding the contract method 0xc766c114.
//
// Solidity: function BM_ERC1155() view returns(address)
func (_BmGovernor *BmGovernorCaller) BMERC1155(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "BM_ERC1155")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BMERC1155 is a free data retrieval call binding the contract method 0xc766c114.
//
// Solidity: function BM_ERC1155() view returns(address)
func (_BmGovernor *BmGovernorSession) BMERC1155() (common.Address, error) {
	return _BmGovernor.Contract.BMERC1155(&_BmGovernor.CallOpts)
}

// BMERC1155 is a free data retrieval call binding the contract method 0xc766c114.
//
// Solidity: function BM_ERC1155() view returns(address)
func (_BmGovernor *BmGovernorCallerSession) BMERC1155() (common.Address, error) {
	return _BmGovernor.Contract.BMERC1155(&_BmGovernor.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_BmGovernor *BmGovernorCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_BmGovernor *BmGovernorSession) CLOCKMODE() (string, error) {
	return _BmGovernor.Contract.CLOCKMODE(&_BmGovernor.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_BmGovernor *BmGovernorCallerSession) CLOCKMODE() (string, error) {
	return _BmGovernor.Contract.CLOCKMODE(&_BmGovernor.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_BmGovernor *BmGovernorCaller) COUNTINGMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "COUNTING_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_BmGovernor *BmGovernorSession) COUNTINGMODE() (string, error) {
	return _BmGovernor.Contract.COUNTINGMODE(&_BmGovernor.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_BmGovernor *BmGovernorCallerSession) COUNTINGMODE() (string, error) {
	return _BmGovernor.Contract.COUNTINGMODE(&_BmGovernor.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_BmGovernor *BmGovernorCaller) EXTENDEDBALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "EXTENDED_BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_BmGovernor *BmGovernorSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _BmGovernor.Contract.EXTENDEDBALLOTTYPEHASH(&_BmGovernor.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_BmGovernor *BmGovernorCallerSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _BmGovernor.Contract.EXTENDEDBALLOTTYPEHASH(&_BmGovernor.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_BmGovernor *BmGovernorCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_BmGovernor *BmGovernorSession) Clock() (*big.Int, error) {
	return _BmGovernor.Contract.Clock(&_BmGovernor.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_BmGovernor *BmGovernorCallerSession) Clock() (*big.Int, error) {
	return _BmGovernor.Contract.Clock(&_BmGovernor.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BmGovernor *BmGovernorCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BmGovernor *BmGovernorSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BmGovernor.Contract.Eip712Domain(&_BmGovernor.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BmGovernor *BmGovernorCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BmGovernor.Contract.Eip712Domain(&_BmGovernor.CallOpts)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_BmGovernor *BmGovernorCaller) GetVotes(opts *bind.CallOpts, account common.Address, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "getVotes", account, timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_BmGovernor *BmGovernorSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _BmGovernor.Contract.GetVotes(&_BmGovernor.CallOpts, account, timepoint)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_BmGovernor *BmGovernorCallerSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _BmGovernor.Contract.GetVotes(&_BmGovernor.CallOpts, account, timepoint)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_BmGovernor *BmGovernorCaller) GetVotesWithParams(opts *bind.CallOpts, account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "getVotesWithParams", account, timepoint, params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_BmGovernor *BmGovernorSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _BmGovernor.Contract.GetVotesWithParams(&_BmGovernor.CallOpts, account, timepoint, params)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_BmGovernor *BmGovernorCallerSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _BmGovernor.Contract.GetVotesWithParams(&_BmGovernor.CallOpts, account, timepoint, params)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_BmGovernor *BmGovernorCaller) HasVoted(opts *bind.CallOpts, proposalId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "hasVoted", proposalId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_BmGovernor *BmGovernorSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _BmGovernor.Contract.HasVoted(&_BmGovernor.CallOpts, proposalId, account)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_BmGovernor *BmGovernorCallerSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _BmGovernor.Contract.HasVoted(&_BmGovernor.CallOpts, proposalId, account)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_BmGovernor *BmGovernorCaller) HashProposal(opts *bind.CallOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "hashProposal", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_BmGovernor *BmGovernorSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _BmGovernor.Contract.HashProposal(&_BmGovernor.CallOpts, targets, values, calldatas, descriptionHash)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_BmGovernor *BmGovernorCallerSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _BmGovernor.Contract.HashProposal(&_BmGovernor.CallOpts, targets, values, calldatas, descriptionHash)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BmGovernor *BmGovernorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BmGovernor *BmGovernorSession) Name() (string, error) {
	return _BmGovernor.Contract.Name(&_BmGovernor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BmGovernor *BmGovernorCallerSession) Name() (string, error) {
	return _BmGovernor.Contract.Name(&_BmGovernor.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BmGovernor *BmGovernorCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BmGovernor *BmGovernorSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BmGovernor.Contract.Nonces(&_BmGovernor.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BmGovernor *BmGovernorCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BmGovernor.Contract.Nonces(&_BmGovernor.CallOpts, owner)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_BmGovernor *BmGovernorCaller) ProposalDeadline(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "proposalDeadline", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_BmGovernor *BmGovernorSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _BmGovernor.Contract.ProposalDeadline(&_BmGovernor.CallOpts, proposalId)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_BmGovernor *BmGovernorCallerSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _BmGovernor.Contract.ProposalDeadline(&_BmGovernor.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_BmGovernor *BmGovernorCaller) ProposalEta(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "proposalEta", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_BmGovernor *BmGovernorSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _BmGovernor.Contract.ProposalEta(&_BmGovernor.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_BmGovernor *BmGovernorCallerSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _BmGovernor.Contract.ProposalEta(&_BmGovernor.CallOpts, proposalId)
}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 ) view returns(bool)
func (_BmGovernor *BmGovernorCaller) ProposalNeedsQueuing(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "proposalNeedsQueuing", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 ) view returns(bool)
func (_BmGovernor *BmGovernorSession) ProposalNeedsQueuing(arg0 *big.Int) (bool, error) {
	return _BmGovernor.Contract.ProposalNeedsQueuing(&_BmGovernor.CallOpts, arg0)
}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 ) view returns(bool)
func (_BmGovernor *BmGovernorCallerSession) ProposalNeedsQueuing(arg0 *big.Int) (bool, error) {
	return _BmGovernor.Contract.ProposalNeedsQueuing(&_BmGovernor.CallOpts, arg0)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_BmGovernor *BmGovernorCaller) ProposalProposer(opts *bind.CallOpts, proposalId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "proposalProposer", proposalId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_BmGovernor *BmGovernorSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _BmGovernor.Contract.ProposalProposer(&_BmGovernor.CallOpts, proposalId)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_BmGovernor *BmGovernorCallerSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _BmGovernor.Contract.ProposalProposer(&_BmGovernor.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_BmGovernor *BmGovernorCaller) ProposalSnapshot(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "proposalSnapshot", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_BmGovernor *BmGovernorSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _BmGovernor.Contract.ProposalSnapshot(&_BmGovernor.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_BmGovernor *BmGovernorCallerSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _BmGovernor.Contract.ProposalSnapshot(&_BmGovernor.CallOpts, proposalId)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (_BmGovernor *BmGovernorCaller) ProposalThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "proposalThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (_BmGovernor *BmGovernorSession) ProposalThreshold() (*big.Int, error) {
	return _BmGovernor.Contract.ProposalThreshold(&_BmGovernor.CallOpts)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (_BmGovernor *BmGovernorCallerSession) ProposalThreshold() (*big.Int, error) {
	return _BmGovernor.Contract.ProposalThreshold(&_BmGovernor.CallOpts)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_BmGovernor *BmGovernorCaller) ProposalVotes(opts *bind.CallOpts, proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "proposalVotes", proposalId)

	outstruct := new(struct {
		AgainstVotes *big.Int
		ForVotes     *big.Int
		AbstainVotes *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AgainstVotes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ForVotes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AbstainVotes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_BmGovernor *BmGovernorSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _BmGovernor.Contract.ProposalVotes(&_BmGovernor.CallOpts, proposalId)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_BmGovernor *BmGovernorCallerSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _BmGovernor.Contract.ProposalVotes(&_BmGovernor.CallOpts, proposalId)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (_BmGovernor *BmGovernorCaller) Quorum(opts *bind.CallOpts, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "quorum", timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (_BmGovernor *BmGovernorSession) Quorum(timepoint *big.Int) (*big.Int, error) {
	return _BmGovernor.Contract.Quorum(&_BmGovernor.CallOpts, timepoint)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (_BmGovernor *BmGovernorCallerSession) Quorum(timepoint *big.Int) (*big.Int, error) {
	return _BmGovernor.Contract.Quorum(&_BmGovernor.CallOpts, timepoint)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_BmGovernor *BmGovernorCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_BmGovernor *BmGovernorSession) State(proposalId *big.Int) (uint8, error) {
	return _BmGovernor.Contract.State(&_BmGovernor.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_BmGovernor *BmGovernorCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _BmGovernor.Contract.State(&_BmGovernor.CallOpts, proposalId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BmGovernor *BmGovernorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BmGovernor *BmGovernorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BmGovernor.Contract.SupportsInterface(&_BmGovernor.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BmGovernor *BmGovernorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BmGovernor.Contract.SupportsInterface(&_BmGovernor.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BmGovernor *BmGovernorCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BmGovernor *BmGovernorSession) Token() (common.Address, error) {
	return _BmGovernor.Contract.Token(&_BmGovernor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BmGovernor *BmGovernorCallerSession) Token() (common.Address, error) {
	return _BmGovernor.Contract.Token(&_BmGovernor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BmGovernor *BmGovernorCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BmGovernor *BmGovernorSession) Version() (string, error) {
	return _BmGovernor.Contract.Version(&_BmGovernor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BmGovernor *BmGovernorCallerSession) Version() (string, error) {
	return _BmGovernor.Contract.Version(&_BmGovernor.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_BmGovernor *BmGovernorCaller) VotingDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "votingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_BmGovernor *BmGovernorSession) VotingDelay() (*big.Int, error) {
	return _BmGovernor.Contract.VotingDelay(&_BmGovernor.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_BmGovernor *BmGovernorCallerSession) VotingDelay() (*big.Int, error) {
	return _BmGovernor.Contract.VotingDelay(&_BmGovernor.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (_BmGovernor *BmGovernorCaller) VotingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (_BmGovernor *BmGovernorSession) VotingPeriod() (*big.Int, error) {
	return _BmGovernor.Contract.VotingPeriod(&_BmGovernor.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (_BmGovernor *BmGovernorCallerSession) VotingPeriod() (*big.Int, error) {
	return _BmGovernor.Contract.VotingPeriod(&_BmGovernor.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_BmGovernor *BmGovernorTransactor) Cancel(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _BmGovernor.contract.Transact(opts, "cancel", targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_BmGovernor *BmGovernorSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.Cancel(&_BmGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_BmGovernor *BmGovernorTransactorSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.Cancel(&_BmGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_BmGovernor *BmGovernorTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _BmGovernor.contract.Transact(opts, "castVote", proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_BmGovernor *BmGovernorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _BmGovernor.Contract.CastVote(&_BmGovernor.TransactOpts, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_BmGovernor *BmGovernorTransactorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _BmGovernor.Contract.CastVote(&_BmGovernor.TransactOpts, proposalId, support)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_BmGovernor *BmGovernorTransactor) CastVoteBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _BmGovernor.contract.Transact(opts, "castVoteBySig", proposalId, support, voter, signature)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_BmGovernor *BmGovernorSession) CastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.CastVoteBySig(&_BmGovernor.TransactOpts, proposalId, support, voter, signature)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_BmGovernor *BmGovernorTransactorSession) CastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.CastVoteBySig(&_BmGovernor.TransactOpts, proposalId, support, voter, signature)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_BmGovernor *BmGovernorTransactor) CastVoteWithReason(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _BmGovernor.contract.Transact(opts, "castVoteWithReason", proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_BmGovernor *BmGovernorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _BmGovernor.Contract.CastVoteWithReason(&_BmGovernor.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_BmGovernor *BmGovernorTransactorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _BmGovernor.Contract.CastVoteWithReason(&_BmGovernor.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_BmGovernor *BmGovernorTransactor) CastVoteWithReasonAndParams(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _BmGovernor.contract.Transact(opts, "castVoteWithReasonAndParams", proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_BmGovernor *BmGovernorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.CastVoteWithReasonAndParams(&_BmGovernor.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_BmGovernor *BmGovernorTransactorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.CastVoteWithReasonAndParams(&_BmGovernor.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_BmGovernor *BmGovernorTransactor) CastVoteWithReasonAndParamsBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _BmGovernor.contract.Transact(opts, "castVoteWithReasonAndParamsBySig", proposalId, support, voter, reason, params, signature)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_BmGovernor *BmGovernorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.CastVoteWithReasonAndParamsBySig(&_BmGovernor.TransactOpts, proposalId, support, voter, reason, params, signature)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_BmGovernor *BmGovernorTransactorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.CastVoteWithReasonAndParamsBySig(&_BmGovernor.TransactOpts, proposalId, support, voter, reason, params, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_BmGovernor *BmGovernorTransactor) Execute(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _BmGovernor.contract.Transact(opts, "execute", targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_BmGovernor *BmGovernorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.Execute(&_BmGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_BmGovernor *BmGovernorTransactorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.Execute(&_BmGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_BmGovernor *BmGovernorTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BmGovernor.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_BmGovernor *BmGovernorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.OnERC1155BatchReceived(&_BmGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_BmGovernor *BmGovernorTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.OnERC1155BatchReceived(&_BmGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_BmGovernor *BmGovernorTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BmGovernor.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_BmGovernor *BmGovernorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.OnERC1155Received(&_BmGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_BmGovernor *BmGovernorTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.OnERC1155Received(&_BmGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_BmGovernor *BmGovernorTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _BmGovernor.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_BmGovernor *BmGovernorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.OnERC721Received(&_BmGovernor.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_BmGovernor *BmGovernorTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.OnERC721Received(&_BmGovernor.TransactOpts, arg0, arg1, arg2, arg3)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_BmGovernor *BmGovernorTransactor) Propose(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _BmGovernor.contract.Transact(opts, "propose", targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_BmGovernor *BmGovernorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _BmGovernor.Contract.Propose(&_BmGovernor.TransactOpts, targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_BmGovernor *BmGovernorTransactorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _BmGovernor.Contract.Propose(&_BmGovernor.TransactOpts, targets, values, calldatas, description)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_BmGovernor *BmGovernorTransactor) Queue(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _BmGovernor.contract.Transact(opts, "queue", targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_BmGovernor *BmGovernorSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.Queue(&_BmGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_BmGovernor *BmGovernorTransactorSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.Queue(&_BmGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_BmGovernor *BmGovernorTransactor) Relay(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BmGovernor.contract.Transact(opts, "relay", target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_BmGovernor *BmGovernorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.Relay(&_BmGovernor.TransactOpts, target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_BmGovernor *BmGovernorTransactorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BmGovernor.Contract.Relay(&_BmGovernor.TransactOpts, target, value, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BmGovernor *BmGovernorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BmGovernor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BmGovernor *BmGovernorSession) Receive() (*types.Transaction, error) {
	return _BmGovernor.Contract.Receive(&_BmGovernor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BmGovernor *BmGovernorTransactorSession) Receive() (*types.Transaction, error) {
	return _BmGovernor.Contract.Receive(&_BmGovernor.TransactOpts)
}

// BmGovernorEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the BmGovernor contract.
type BmGovernorEIP712DomainChangedIterator struct {
	Event *BmGovernorEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *BmGovernorEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BmGovernorEIP712DomainChanged)
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
		it.Event = new(BmGovernorEIP712DomainChanged)
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
func (it *BmGovernorEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BmGovernorEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BmGovernorEIP712DomainChanged represents a EIP712DomainChanged event raised by the BmGovernor contract.
type BmGovernorEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BmGovernor *BmGovernorFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BmGovernorEIP712DomainChangedIterator, error) {

	logs, sub, err := _BmGovernor.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BmGovernorEIP712DomainChangedIterator{contract: _BmGovernor.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BmGovernor *BmGovernorFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BmGovernorEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _BmGovernor.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BmGovernorEIP712DomainChanged)
				if err := _BmGovernor.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BmGovernor *BmGovernorFilterer) ParseEIP712DomainChanged(log types.Log) (*BmGovernorEIP712DomainChanged, error) {
	event := new(BmGovernorEIP712DomainChanged)
	if err := _BmGovernor.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BmGovernorProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the BmGovernor contract.
type BmGovernorProposalCanceledIterator struct {
	Event *BmGovernorProposalCanceled // Event containing the contract specifics and raw log

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
func (it *BmGovernorProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BmGovernorProposalCanceled)
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
		it.Event = new(BmGovernorProposalCanceled)
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
func (it *BmGovernorProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BmGovernorProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BmGovernorProposalCanceled represents a ProposalCanceled event raised by the BmGovernor contract.
type BmGovernorProposalCanceled struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_BmGovernor *BmGovernorFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*BmGovernorProposalCanceledIterator, error) {

	logs, sub, err := _BmGovernor.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &BmGovernorProposalCanceledIterator{contract: _BmGovernor.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_BmGovernor *BmGovernorFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *BmGovernorProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _BmGovernor.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BmGovernorProposalCanceled)
				if err := _BmGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
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

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_BmGovernor *BmGovernorFilterer) ParseProposalCanceled(log types.Log) (*BmGovernorProposalCanceled, error) {
	event := new(BmGovernorProposalCanceled)
	if err := _BmGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BmGovernorProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the BmGovernor contract.
type BmGovernorProposalCreatedIterator struct {
	Event *BmGovernorProposalCreated // Event containing the contract specifics and raw log

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
func (it *BmGovernorProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BmGovernorProposalCreated)
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
		it.Event = new(BmGovernorProposalCreated)
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
func (it *BmGovernorProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BmGovernorProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BmGovernorProposalCreated represents a ProposalCreated event raised by the BmGovernor contract.
type BmGovernorProposalCreated struct {
	ProposalId  *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	VoteStart   *big.Int
	VoteEnd     *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_BmGovernor *BmGovernorFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*BmGovernorProposalCreatedIterator, error) {

	logs, sub, err := _BmGovernor.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &BmGovernorProposalCreatedIterator{contract: _BmGovernor.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_BmGovernor *BmGovernorFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *BmGovernorProposalCreated) (event.Subscription, error) {

	logs, sub, err := _BmGovernor.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BmGovernorProposalCreated)
				if err := _BmGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_BmGovernor *BmGovernorFilterer) ParseProposalCreated(log types.Log) (*BmGovernorProposalCreated, error) {
	event := new(BmGovernorProposalCreated)
	if err := _BmGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BmGovernorProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the BmGovernor contract.
type BmGovernorProposalExecutedIterator struct {
	Event *BmGovernorProposalExecuted // Event containing the contract specifics and raw log

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
func (it *BmGovernorProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BmGovernorProposalExecuted)
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
		it.Event = new(BmGovernorProposalExecuted)
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
func (it *BmGovernorProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BmGovernorProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BmGovernorProposalExecuted represents a ProposalExecuted event raised by the BmGovernor contract.
type BmGovernorProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_BmGovernor *BmGovernorFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*BmGovernorProposalExecutedIterator, error) {

	logs, sub, err := _BmGovernor.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &BmGovernorProposalExecutedIterator{contract: _BmGovernor.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_BmGovernor *BmGovernorFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *BmGovernorProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _BmGovernor.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BmGovernorProposalExecuted)
				if err := _BmGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_BmGovernor *BmGovernorFilterer) ParseProposalExecuted(log types.Log) (*BmGovernorProposalExecuted, error) {
	event := new(BmGovernorProposalExecuted)
	if err := _BmGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BmGovernorProposalQueuedIterator is returned from FilterProposalQueued and is used to iterate over the raw logs and unpacked data for ProposalQueued events raised by the BmGovernor contract.
type BmGovernorProposalQueuedIterator struct {
	Event *BmGovernorProposalQueued // Event containing the contract specifics and raw log

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
func (it *BmGovernorProposalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BmGovernorProposalQueued)
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
		it.Event = new(BmGovernorProposalQueued)
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
func (it *BmGovernorProposalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BmGovernorProposalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BmGovernorProposalQueued represents a ProposalQueued event raised by the BmGovernor contract.
type BmGovernorProposalQueued struct {
	ProposalId *big.Int
	EtaSeconds *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalQueued is a free log retrieval operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_BmGovernor *BmGovernorFilterer) FilterProposalQueued(opts *bind.FilterOpts) (*BmGovernorProposalQueuedIterator, error) {

	logs, sub, err := _BmGovernor.contract.FilterLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return &BmGovernorProposalQueuedIterator{contract: _BmGovernor.contract, event: "ProposalQueued", logs: logs, sub: sub}, nil
}

// WatchProposalQueued is a free log subscription operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_BmGovernor *BmGovernorFilterer) WatchProposalQueued(opts *bind.WatchOpts, sink chan<- *BmGovernorProposalQueued) (event.Subscription, error) {

	logs, sub, err := _BmGovernor.contract.WatchLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BmGovernorProposalQueued)
				if err := _BmGovernor.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
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

// ParseProposalQueued is a log parse operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_BmGovernor *BmGovernorFilterer) ParseProposalQueued(log types.Log) (*BmGovernorProposalQueued, error) {
	event := new(BmGovernorProposalQueued)
	if err := _BmGovernor.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BmGovernorVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the BmGovernor contract.
type BmGovernorVoteCastIterator struct {
	Event *BmGovernorVoteCast // Event containing the contract specifics and raw log

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
func (it *BmGovernorVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BmGovernorVoteCast)
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
		it.Event = new(BmGovernorVoteCast)
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
func (it *BmGovernorVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BmGovernorVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BmGovernorVoteCast represents a VoteCast event raised by the BmGovernor contract.
type BmGovernorVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_BmGovernor *BmGovernorFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address) (*BmGovernorVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _BmGovernor.contract.FilterLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return &BmGovernorVoteCastIterator{contract: _BmGovernor.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_BmGovernor *BmGovernorFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *BmGovernorVoteCast, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _BmGovernor.contract.WatchLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BmGovernorVoteCast)
				if err := _BmGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_BmGovernor *BmGovernorFilterer) ParseVoteCast(log types.Log) (*BmGovernorVoteCast, error) {
	event := new(BmGovernorVoteCast)
	if err := _BmGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BmGovernorVoteCastWithParamsIterator is returned from FilterVoteCastWithParams and is used to iterate over the raw logs and unpacked data for VoteCastWithParams events raised by the BmGovernor contract.
type BmGovernorVoteCastWithParamsIterator struct {
	Event *BmGovernorVoteCastWithParams // Event containing the contract specifics and raw log

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
func (it *BmGovernorVoteCastWithParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BmGovernorVoteCastWithParams)
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
		it.Event = new(BmGovernorVoteCastWithParams)
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
func (it *BmGovernorVoteCastWithParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BmGovernorVoteCastWithParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BmGovernorVoteCastWithParams represents a VoteCastWithParams event raised by the BmGovernor contract.
type BmGovernorVoteCastWithParams struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Params     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCastWithParams is a free log retrieval operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_BmGovernor *BmGovernorFilterer) FilterVoteCastWithParams(opts *bind.FilterOpts, voter []common.Address) (*BmGovernorVoteCastWithParamsIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _BmGovernor.contract.FilterLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return &BmGovernorVoteCastWithParamsIterator{contract: _BmGovernor.contract, event: "VoteCastWithParams", logs: logs, sub: sub}, nil
}

// WatchVoteCastWithParams is a free log subscription operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_BmGovernor *BmGovernorFilterer) WatchVoteCastWithParams(opts *bind.WatchOpts, sink chan<- *BmGovernorVoteCastWithParams, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _BmGovernor.contract.WatchLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BmGovernorVoteCastWithParams)
				if err := _BmGovernor.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
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

// ParseVoteCastWithParams is a log parse operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_BmGovernor *BmGovernorFilterer) ParseVoteCastWithParams(log types.Log) (*BmGovernorVoteCastWithParams, error) {
	event := new(BmGovernorVoteCastWithParams)
	if err := _BmGovernor.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
