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
	Bin: "0x6101a060405234801562000011575f80fd5b50604051620038ec380380620038ec83398101604081905262000034916200020b565b808280620000566040805180820190915260018152603160f81b602082015290565b62000062825f62000134565b610120526200007381600162000134565b61014052815160208084019190912060e052815190820120610100524660a0526200010060e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c05260036200011782826200035d565b50506001600160a01b039081166101605216610180525062000481565b5f60208351101562000153576200014b836200016c565b905062000166565b816200016084826200035d565b5060ff90505b92915050565b5f80829050601f81511115620001a2578260405163305a27a960e01b815260040162000199919062000429565b60405180910390fd5b8051620001af826200045d565b179392505050565b634e487b7160e01b5f52604160045260245ffd5b5f5b83811015620001e7578181015183820152602001620001cd565b50505f910152565b80516001600160a01b038116811462000206575f80fd5b919050565b5f80604083850312156200021d575f80fd5b82516001600160401b038082111562000234575f80fd5b818501915085601f83011262000248575f80fd5b8151818111156200025d576200025d620001b7565b604051601f8201601f19908116603f01168101908382118183101715620002885762000288620001b7565b81604052828152886020848701011115620002a1575f80fd5b620002b4836020830160208801620001cb565b8096505050505050620002ca60208401620001ef565b90509250929050565b600181811c90821680620002e857607f821691505b6020821081036200030757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156200035857805f5260205f20601f840160051c81016020851015620003345750805b601f840160051c820191505b8181101562000355575f815560010162000340565b50505b505050565b81516001600160401b03811115620003795762000379620001b7565b62000391816200038a8454620002d3565b846200030d565b602080601f831160018114620003c7575f8415620003af5750858301515b5f19600386901b1c1916600185901b17855562000421565b5f85815260208120601f198616915b82811015620003f757888601518255948401946001909101908401620003d6565b50858210156200041557878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b602081525f825180602084015262000449816040850160208701620001cb565b601f01601f19169190910160400192915050565b8051602080830151919081101562000307575f1960209190910360031b1b16919050565b60805160a05160c05160e05161010051610120516101405161016051610180516133d4620005185f395f818161073601528181610eed01528181610f7b015281816112a5015261171c01525f818161081601528181611594015281816118e701526121bd01525f6118be01525f61189201525f611ad901525f611ab101525f611a0c01525f611a3601525f611a6001526133d45ff3fe608060405260043610610236575f3560e01c80637d5e81e211610129578063c01f9e37116100a8578063deaaa7cc1161006d578063deaaa7cc1461076c578063eb9019d41461079f578063f23a6e61146107be578063f8ce560a146107e9578063fc0c546a14610808575f80fd5b8063c01f9e37146106d4578063c28bc2fa146106f3578063c59057e414610706578063c766c11414610725578063dd4e2ba514610758575f80fd5b80639a802a6d116100ee5780639a802a6d1461061b578063a9a952941461063a578063ab58fb8e14610659578063b58131b01461068f578063bc197c81146106a9575f80fd5b80637d5e81e2146105575780637ecebe001461057657806384b0196e146105aa5780638ff262e3146105d157806391ddadf4146105f0575f80fd5b80633e4f49e6116101b557806354fd4d501161017a57806354fd4d50146104b257806356781388146104db5780635b8d0e0d146104fa5780635f398a14146105195780637b3c71d314610538575f80fd5b80633e4f49e6146103e1578063438596321461040d578063452115d61461042c5780634bf5d7e91461044b578063544ffc9c1461045f575f80fd5b8063160cbed7116101fb578063160cbed7146103495780632656227d146103685780632d63f6931461037b5780632fe3e2611461039a5780633932abb1146103cd575f80fd5b806301ffc9a71461024357806302a251a31461027757806306fdde0314610299578063143489d0146102ba578063150b7a0214610306575f80fd5b3661023f57005b005b5f80fd5b34801561024e575f80fd5b5061026261025d366004612514565b61083a565b60405190151581526020015b60405180910390f35b348015610282575f80fd5b5061028b61088b565b60405190815260200161026e565b3480156102a4575f80fd5b506102ad6108a1565b60405161026e9190612588565b3480156102c5575f80fd5b506102ee6102d436600461259a565b5f908152600460205260409020546001600160a01b031690565b6040516001600160a01b03909116815260200161026e565b348015610311575f80fd5b5061033061032036600461268a565b630a85bd0160e11b949350505050565b6040516001600160e01b0319909116815260200161026e565b348015610354575f80fd5b5061028b610363366004612853565b610931565b61028b610376366004612853565b61096f565b348015610386575f80fd5b5061028b61039536600461259a565b610a97565b3480156103a5575f80fd5b5061028b7f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a81181565b3480156103d8575f80fd5b5061028b610ab7565b3480156103ec575f80fd5b506104006103fb36600461259a565b610add565b60405161026e9190612910565b348015610418575f80fd5b5061026261042736600461291e565b610c06565b348015610437575f80fd5b5061028b610446366004612853565b610c36565b348015610456575f80fd5b506102ad610ca2565b34801561046a575f80fd5b5061049761047936600461259a565b5f908152600760205260409020805460018201546002909201549092565b6040805193845260208401929092529082015260600161026e565b3480156104bd575f80fd5b506040805180820190915260018152603160f81b60208201526102ad565b3480156104e6575f80fd5b5061028b6104f5366004612958565b610cac565b348015610505575f80fd5b5061028b6105143660046129bd565b610cd3565b348015610524575f80fd5b5061028b610533366004612a6d565b610e2f565b348015610543575f80fd5b5061028b610552366004612aea565b610e82565b348015610562575f80fd5b5061028b610571366004612b3f565b610ec8565b348015610581575f80fd5b5061028b610590366004612beb565b6001600160a01b03165f9081526002602052604090205490565b3480156105b5575f80fd5b506105be61100a565b60405161026e9796959493929190612c3e565b3480156105dc575f80fd5b5061028b6105eb366004612cad565b61104c565b3480156105fb575f80fd5b5061060461111b565b60405165ffffffffffff909116815260200161026e565b348015610626575f80fd5b5061028b610635366004612cf8565b611124565b348015610645575f80fd5b5061026261065436600461259a565b505f90565b348015610664575f80fd5b5061028b61067336600461259a565b5f9081526004602052604090206001015465ffffffffffff1690565b34801561069a575f80fd5b50670de0b6b3a764000061028b565b3480156106b4575f80fd5b506103306106c3366004612d4a565b63bc197c8160e01b95945050505050565b3480156106df575f80fd5b5061028b6106ee36600461259a565b611130565b61023d610701366004612dd2565b611172565b348015610711575f80fd5b5061028b610720366004612853565b6111ee565b348015610730575f80fd5b506102ee7f000000000000000000000000000000000000000000000000000000000000000081565b348015610763575f80fd5b506102ad611227565b348015610777575f80fd5b5061028b7ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d781565b3480156107aa575f80fd5b5061028b6107b9366004612e0f565b611263565b3480156107c9575f80fd5b506103306107d8366004612e37565b63f23a6e6160e01b95945050505050565b3480156107f4575f80fd5b5061028b61080336600461259a565b611282565b348015610813575f80fd5b507f00000000000000000000000000000000000000000000000000000000000000006102ee565b5f6001600160e01b031982166332a2ad4360e11b148061086a57506001600160e01b03198216630271189760e51b145b8061088557506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f61089c6205460062093a80612eaa565b905090565b6060600380546108b090612ebd565b80601f01602080910402602001604051908101604052809291908181526020018280546108dc90612ebd565b80156109275780601f106108fe57610100808354040283529160200191610927565b820191905f5260205f20905b81548152906001019060200180831161090a57829003601f168201915b5050505050905090565b5f8061093f868686866111ee565b90506109548161094f6004611318565b61133a565b505f604051634844252360e11b815260040160405180910390fd5b5f8061097d868686866111ee565b905061099d8161098d6005611318565b6109976004611318565b1761133a565b505f818152600460205260409020805460ff60f01b1916600160f01b179055306109c43090565b6001600160a01b031614610a4d575f5b8651811015610a4b57306001600160a01b03168782815181106109f9576109f9612ef5565b60200260200101516001600160a01b031603610a4357610a43858281518110610a2457610a24612ef5565b602002602001015180519060200120600561137790919063ffffffff16565b6001016109d4565b505b610a5a81878787876113e7565b6040518181527f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f906020015b60405180910390a195945050505050565b5f90815260046020526040902054600160a01b900465ffffffffffff1690565b5f4262093a80810681036205460001620151808201811015610ad7575f80fd5b03919050565b5f818152600460205260408120805460ff600160f01b8204811691600160f81b9004168115610b1157506007949350505050565b8015610b2257506002949350505050565b5f610b2c86610a97565b9050805f03610b5657604051636ad0607560e01b8152600481018790526024015b60405180910390fd5b5f610b5f61111b565b65ffffffffffff169050808210610b7c57505f9695505050505050565b5f610b8688611130565b9050818110610b9d57506001979650505050505050565b610ba6886114bc565b1580610bb85750610bb6886114c6565b155b15610bcb57506003979650505050505050565b5f8881526004602052604090206001015465ffffffffffff165f03610bf857506004979650505050505050565b506005979650505050505050565b5f8281526007602090815260408083206001600160a01b038516845260030190915281205460ff165b9392505050565b5f80610c44868686866111ee565b9050610c538161094f5f611318565b505f818152600460205260409020546001600160a01b03163314610c8c5760405163233d98e360e01b8152336004820152602401610b4d565b610c98868686866114e1565b9695505050505050565b606061089c611590565b5f80339050610ccb84828560405180602001604052805f815250611650565b949350505050565b5f80610db487610dae7f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a8118c8c8c610d268e6001600160a01b03165f90815260026020526040902080546001810190915590565b8d8d604051610d36929190612f09565b60405180910390208c80519060200120604051602001610d939796959493929190968752602087019590955260ff9390931660408601526001600160a01b03919091166060850152608084015260a083015260c082015260e00190565b6040516020818303038152906040528051906020012061167a565b856116a6565b905080610ddf576040516394ab6c0760e01b81526001600160a01b0388166004820152602401610b4d565b610e2289888a89898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508b92506116fb915050565b9998505050505050505050565b5f80339050610e7787828888888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508a92506116fb915050565b979650505050505050565b5f80339050610c9886828787878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061165092505050565b5f80336040516306044ae960e41b81526001600160a01b0380831660048301529192507f000000000000000000000000000000000000000000000000000000000000000090911690636044ae90906024015f604051808303815f87803b158015610f30575f80fd5b505af1158015610f42573d5f803e3d5ffd5b505050505f610f5387878787611782565b9050610c98818360026040516309ab24eb60e41b81526001600160a01b0387811660048301527f00000000000000000000000000000000000000000000000000000000000000001690639ab24eb090602401602060405180830381865afa158015610fc0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fe49190612f18565b60405180604001604052806008815260200167383937b837b9b2b960c11b81525061183a565b5f6060805f805f606061101b61188b565b6110236118b7565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b5f806110d684610dae7ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d789898961109f8b6001600160a01b03165f90815260026020526040902080546001810190915590565b60408051602081019690965285019390935260ff90911660608401526001600160a01b0316608083015260a082015260c001610d93565b905080611101576040516394ab6c0760e01b81526001600160a01b0385166004820152602401610b4d565b610c9886858760405180602001604052805f815250611650565b5f61089c6118e4565b5f610ccb84848461196b565b5f8181526004602052604081205461116490600160d01b810463ffffffff1690600160a01b900465ffffffffffff16612f2f565b65ffffffffffff1692915050565b61117a611977565b5f80856001600160a01b0316858585604051611197929190612f09565b5f6040518083038185875af1925050503d805f81146111d1576040519150601f19603f3d011682016040523d82523d5f602084013e6111d6565b606091505b50915091506111e582826119ae565b50505050505050565b5f848484846040516020016112069493929190612fe5565b60408051601f19818403018152919052805160209091012095945050505050565b606061089c6040805180820190915260208082527f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e9082015290565b5f610c2f838361127d60408051602081019091525f815290565b61196b565b604051632394e7a360e21b8152600481018290525f90600a906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638e539e8c90602401602060405180830381865afa1580156112ea573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061130e9190612f18565b610885919061302f565b5f81600781111561132b5761132b6128dc565b600160ff919091161b92915050565b5f8061134584610add565b90505f8361135283611318565b1603610c2f578381846040516331b75e4d60e01b8152600401610b4d9392919061304e565b81546001600160801b03600160801b8204811691811660018301909116036113b257604051638acb5f2760e01b815260040160405180910390fd5b6001600160801b038082165f90815260018086016020526040909120939093558354919092018216600160801b029116179055565b5f5b84518110156114b4575f8086838151811061140657611406612ef5565b60200260200101516001600160a01b031686848151811061142957611429612ef5565b602002602001015186858151811061144357611443612ef5565b60200260200101516040516114589190613070565b5f6040518083038185875af1925050503d805f8114611492576040519150601f19603f3d011682016040523d82523d5f602084013e611497565b606091505b50915091506114a682826119ae565b5050508060010190506113e9565b505050505050565b5f610885826119ca565b5f818152600760205260408120805460019091015411610885565b5f806114ef868686866111ee565b905061153d816114ff6007611318565b6115096006611318565b6115136002611318565b600161152060078261308b565b61152b906002613184565b6115359190612eaa565b18181861133a565b505f818152600460205260409081902080546001600160f81b0316600160f81b179055517f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c90610a869083815260200190565b60607f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634bf5d7e96040518163ffffffff1660e01b81526004015f60405180830381865afa92505050801561160f57506040513d5f823e601f3d908101601f1916820160405261160c9190810190613192565b60015b61164b575060408051808201909152601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000602082015290565b919050565b5f6116718585858561166c60408051602081019091525f815290565b6116fb565b95945050505050565b5f610885611686611a00565b8360405161190160f01b8152600281019290925260228201526042902090565b5f805f6116b38585611b29565b5090925090505f8160038111156116cc576116cc6128dc565b1480156116ea5750856001600160a01b0316826001600160a01b0316145b80610c985750610c98868686611b72565b6040516306044ae960e41b81526001600160a01b0385811660048301525f917f000000000000000000000000000000000000000000000000000000000000000090911690636044ae90906024015f604051808303815f87803b15801561175f575f80fd5b505af1158015611771573d5f803e3d5ffd5b50505050610c988686868686611c48565b5f3361178e8184611d22565b6117b65760405163d9b3955760e01b81526001600160a01b0382166004820152602401610b4d565b5f6117dc8260016117c561111b565b6117cf91906131fa565b65ffffffffffff16611263565b9050670de0b6b3a76400008082101561182157604051636121770b60e11b81526001600160a01b03841660048201526024810183905260448101829052606401610b4d565b61182e8888888887611e0a565b98975050505050505050565b815f036118775760405162461bcd60e51b815260206004820152600b60248201526a1e995c9bc81dd95a59da1d60aa1b6044820152606401610b4d565b6118848585858585612015565b5050505050565b606061089c7f00000000000000000000000000000000000000000000000000000000000000005f612107565b606061089c7f00000000000000000000000000000000000000000000000000000000000000006001612107565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561195f575060408051601f3d908101601f1916820190925261195c91810190613219565b60015b61164b5761089c6121b0565b5f610ccb8484846121ba565b303314611999576040516347096e4760e01b8152336004820152602401610b4d565b565b806119a6600561224d565b0361199b5750565b6060826119c3576119be826122c9565b610885565b5080610885565b5f818152600760205260408120600281015460018201546119eb919061323e565b6119f761080385610a97565b11159392505050565b5f306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015611a5857507f000000000000000000000000000000000000000000000000000000000000000046145b15611a8257507f000000000000000000000000000000000000000000000000000000000000000090565b61089c604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f805f8351604103611b60576020840151604085015160608601515f1a611b52888285856122f2565b955095509550505050611b6b565b505081515f91506002905b9250925092565b5f805f856001600160a01b03168585604051602401611b92929190613251565b60408051601f198184030181529181526020820180516001600160e01b0316630b135d3f60e11b17905251611bc79190613070565b5f60405180830381855afa9150503d805f8114611bff576040519150601f19603f3d011682016040523d82523d5f602084013e611c04565b606091505b5091509150818015611c1857506020815110155b8015610c9857508051630b135d3f60e11b90611c3d9083016020908101908401612f18565b149695505050505050565b5f611c578661094f6001611318565b505f611c6c86611c6689610a97565b8561196b565b9050611c7b878787848761183a565b82515f03611ccf57856001600160a01b03167fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda488878488604051611cc29493929190613269565b60405180910390a2610c98565b856001600160a01b03167fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb87128887848888604051611d10959493929190613290565b60405180910390a29695505050505050565b80515f906034811015611d39576001915050610885565b82810160131901516001600160a01b031981166b046e0e4dee0dee6cae47a60f60a31b14611d6c57600192505050610885565b5f80611d79602885612eaa565b90505b83811015611de9575f80611daf888481518110611d9b57611d9b612ef5565b01602001516001600160f81b0319166123ba565b9150915081611dc75760019650505050505050610885565b8060ff166004856001600160a01b0316901b1793505050806001019050611d7c565b50856001600160a01b0316816001600160a01b031614935050505092915050565b5f611e1e86868686805190602001206111ee565b905084518651141580611e3357508351865114155b80611e3d57508551155b15611e7257855184518651604051630447b05d60e41b8152600481019390935260248301919091526044820152606401610b4d565b5f81815260046020526040902054600160a01b900465ffffffffffff1615611ebb5780611e9e82610add565b6040516331b75e4d60e01b8152610b4d9291905f9060040161304e565b5f611ec4610ab7565b611ecc61111b565b65ffffffffffff16611ede919061323e565b90505f611ee961088b565b5f84815260046020526040902080546001600160a01b0319166001600160a01b038716178155909150611f1b8361244a565b815465ffffffffffff91909116600160a01b0265ffffffffffff60a01b19909116178155611f4882612480565b815463ffffffff91909116600160d01b0263ffffffff60d01b1990911617815588517f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e090859087908c908c906001600160401b03811115611fab57611fab6125c7565b604051908082528060200260200182016040528015611fde57816020015b6060815260200190600190039081611fc95790505b508c89611feb8a8261323e565b8e604051612001999897969594939291906132c9565b60405180910390a150505095945050505050565b5f8581526007602090815260408083206001600160a01b0388168452600381019092529091205460ff1615612068576040516371c6af4960e01b81526001600160a01b0386166004820152602401610b4d565b6001600160a01b0385165f9081526003820160205260409020805460ff1916600117905560ff84166120b15782815f015f8282546120a6919061323e565b909155506114b49050565b5f1960ff8516016120cf5782816001015f8282546120a6919061323e565b60011960ff8516016120ee5782816002015f8282546120a6919061323e565b6040516303599be160e11b815260040160405180910390fd5b606060ff83146121215761211a836124b0565b9050610885565b81805461212d90612ebd565b80601f016020809104026020016040519081016040528092919081815260200182805461215990612ebd565b80156121a45780601f1061217b576101008083540402835291602001916121a4565b820191905f5260205f20905b81548152906001019060200180831161218757829003601f168201915b50505050509050610885565b5f61089c4361244a565b5f7f0000000000000000000000000000000000000000000000000000000000000000604051630748d63560e31b81526001600160a01b038681166004830152602482018690529190911690633a46b1a890604401602060405180830381865afa158015612229573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ccb9190612f18565b80545f906001600160801b0380821691600160801b9004168103612284576040516375e52f4f60e01b815260040160405180910390fd5b6001600160801b038181165f908152600185810160205260408220805492905585546fffffffffffffffffffffffffffffffff19169301909116919091179092555090565b8051156122d95780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561232b57505f915060039050826123b0565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa15801561237c573d5f803e3d5ffd5b5050604051601f1901519150506001600160a01b0381166123a757505f9250600191508290506123b0565b92505f91508190505b9450945094915050565b5f8060f883901c602f811180156123d45750603a8160ff16105b156123e957600194602f199091019350915050565b8060ff1660401080156123ff575060478160ff16105b15612414576001946036199091019350915050565b8060ff16606010801561242a575060678160ff16105b1561243f576001946056199091019350915050565b505f93849350915050565b5f65ffffffffffff82111561247c576040516306dfcc6560e41b81526030600482015260248101839052604401610b4d565b5090565b5f63ffffffff82111561247c576040516306dfcc6560e41b81526020600482015260248101839052604401610b4d565b60605f6124bc836124ed565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f60ff8216601f81111561088557604051632cd44ac360e21b815260040160405180910390fd5b5f60208284031215612524575f80fd5b81356001600160e01b031981168114610c2f575f80fd5b5f5b8381101561255557818101518382015260200161253d565b50505f910152565b5f815180845261257481602086016020860161253b565b601f01601f19169290920160200192915050565b602081525f610c2f602083018461255d565b5f602082840312156125aa575f80fd5b5035919050565b80356001600160a01b038116811461164b575f80fd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715612603576126036125c7565b604052919050565b5f6001600160401b03821115612623576126236125c7565b50601f01601f191660200190565b5f61264361263e8461260b565b6125db565b9050828152838383011115612656575f80fd5b828260208301375f602084830101529392505050565b5f82601f83011261267b575f80fd5b610c2f83833560208501612631565b5f805f806080858703121561269d575f80fd5b6126a6856125b1565b93506126b4602086016125b1565b92506040850135915060608501356001600160401b038111156126d5575f80fd5b6126e18782880161266c565b91505092959194509250565b5f6001600160401b03821115612705576127056125c7565b5060051b60200190565b5f82601f83011261271e575f80fd5b8135602061272e61263e836126ed565b8083825260208201915060208460051b87010193508684111561274f575f80fd5b602086015b8481101561277257612765816125b1565b8352918301918301612754565b509695505050505050565b5f82601f83011261278c575f80fd5b8135602061279c61263e836126ed565b8083825260208201915060208460051b8701019350868411156127bd575f80fd5b602086015b8481101561277257803583529183019183016127c2565b5f82601f8301126127e8575f80fd5b813560206127f861263e836126ed565b82815260059290921b84018101918181019086841115612816575f80fd5b8286015b848110156127725780356001600160401b03811115612837575f80fd5b6128458986838b010161266c565b84525091830191830161281a565b5f805f8060808587031215612866575f80fd5b84356001600160401b038082111561287c575f80fd5b6128888883890161270f565b9550602087013591508082111561289d575f80fd5b6128a98883890161277d565b945060408701359150808211156128be575f80fd5b506128cb878288016127d9565b949793965093946060013593505050565b634e487b7160e01b5f52602160045260245ffd5b6008811061290c57634e487b7160e01b5f52602160045260245ffd5b9052565b6020810161088582846128f0565b5f806040838503121561292f575f80fd5b8235915061293f602084016125b1565b90509250929050565b803560ff8116811461164b575f80fd5b5f8060408385031215612969575f80fd5b8235915061293f60208401612948565b5f8083601f840112612989575f80fd5b5081356001600160401b0381111561299f575f80fd5b6020830191508360208285010111156129b6575f80fd5b9250929050565b5f805f805f805f60c0888a0312156129d3575f80fd5b873596506129e360208901612948565b95506129f1604089016125b1565b945060608801356001600160401b0380821115612a0c575f80fd5b612a188b838c01612979565b909650945060808a0135915080821115612a30575f80fd5b612a3c8b838c0161266c565b935060a08a0135915080821115612a51575f80fd5b50612a5e8a828b0161266c565b91505092959891949750929550565b5f805f805f60808688031215612a81575f80fd5b85359450612a9160208701612948565b935060408601356001600160401b0380821115612aac575f80fd5b612ab889838a01612979565b90955093506060880135915080821115612ad0575f80fd5b50612add8882890161266c565b9150509295509295909350565b5f805f8060608587031215612afd575f80fd5b84359350612b0d60208601612948565b925060408501356001600160401b03811115612b27575f80fd5b612b3387828801612979565b95989497509550505050565b5f805f8060808587031215612b52575f80fd5b84356001600160401b0380821115612b68575f80fd5b612b748883890161270f565b95506020870135915080821115612b89575f80fd5b612b958883890161277d565b94506040870135915080821115612baa575f80fd5b612bb6888389016127d9565b93506060870135915080821115612bcb575f80fd5b508501601f81018713612bdc575f80fd5b6126e187823560208401612631565b5f60208284031215612bfb575f80fd5b610c2f826125b1565b5f815180845260208085019450602084015f5b83811015612c3357815187529582019590820190600101612c17565b509495945050505050565b60ff60f81b8816815260e060208201525f612c5c60e083018961255d565b8281036040840152612c6e818961255d565b606084018890526001600160a01b038716608085015260a0840186905283810360c08501529050612c9f8185612c04565b9a9950505050505050505050565b5f805f8060808587031215612cc0575f80fd5b84359350612cd060208601612948565b9250612cde604086016125b1565b915060608501356001600160401b038111156126d5575f80fd5b5f805f60608486031215612d0a575f80fd5b612d13846125b1565b92506020840135915060408401356001600160401b03811115612d34575f80fd5b612d408682870161266c565b9150509250925092565b5f805f805f60a08688031215612d5e575f80fd5b612d67866125b1565b9450612d75602087016125b1565b935060408601356001600160401b0380821115612d90575f80fd5b612d9c89838a0161277d565b94506060880135915080821115612db1575f80fd5b612dbd89838a0161277d565b93506080880135915080821115612ad0575f80fd5b5f805f8060608587031215612de5575f80fd5b612dee856125b1565b93506020850135925060408501356001600160401b03811115612b27575f80fd5b5f8060408385031215612e20575f80fd5b612e29836125b1565b946020939093013593505050565b5f805f805f60a08688031215612e4b575f80fd5b612e54866125b1565b9450612e62602087016125b1565b9350604086013592506060860135915060808601356001600160401b03811115612e8a575f80fd5b612add8882890161266c565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561088557610885612e96565b600181811c90821680612ed157607f821691505b602082108103612eef57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b818382375f9101908152919050565b5f60208284031215612f28575f80fd5b5051919050565b65ffffffffffff818116838216019080821115612f4e57612f4e612e96565b5092915050565b5f815180845260208085019450602084015f5b83811015612c335781516001600160a01b031687529582019590820190600101612f68565b5f8282518085526020808601955060208260051b840101602086015f5b84811015612fd857601f19868403018952612fc683835161255d565b98840198925090830190600101612faa565b5090979650505050505050565b608081525f612ff76080830187612f55565b82810360208401526130098187612c04565b9050828103604084015261301d8186612f8d565b91505082606083015295945050505050565b5f8261304957634e487b7160e01b5f52601260045260245ffd5b500490565b8381526060810161306260208301856128f0565b826040830152949350505050565b5f825161308181846020870161253b565b9190910192915050565b60ff818116838216019081111561088557610885612e96565b600181815b808511156130de57815f19048211156130c4576130c4612e96565b808516156130d157918102915b93841c93908002906130a9565b509250929050565b5f826130f457506001610885565b8161310057505f610885565b816001811461311657600281146131205761313c565b6001915050610885565b60ff84111561313157613131612e96565b50506001821b610885565b5060208310610133831016604e8410600b841016171561315f575081810a610885565b61316983836130a4565b805f190482111561317c5761317c612e96565b029392505050565b5f610c2f60ff8416836130e6565b5f602082840312156131a2575f80fd5b81516001600160401b038111156131b7575f80fd5b8201601f810184136131c7575f80fd5b80516131d561263e8261260b565b8181528560208385010111156131e9575f80fd5b61167182602083016020860161253b565b65ffffffffffff828116828216039080821115612f4e57612f4e612e96565b5f60208284031215613229575f80fd5b815165ffffffffffff81168114610c2f575f80fd5b8082018082111561088557610885612e96565b828152604060208201525f610ccb604083018461255d565b84815260ff84166020820152826040820152608060608201525f610c98608083018461255d565b85815260ff8516602082015283604082015260a060608201525f6132b760a083018561255d565b828103608084015261182e818561255d565b5f6101208b8352602060018060a01b038c16818501528160408501526132f18285018c612f55565b91508382036060850152613305828b612c04565b915083820360808501528189518084528284019150828160051b850101838c015f5b8381101561335557601f1987840301855261334383835161255d565b94860194925090850190600101613327565b505086810360a0880152613369818c612f8d565b9450505050508560c08401528460e084015282810361010084015261338e818561255d565b9c9b50505050505050505050505056fea264697066735822122058ab54541c3810e027414222047fc047d139c81c14b3d5cfee027c164387040064736f6c63430008180033",
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
