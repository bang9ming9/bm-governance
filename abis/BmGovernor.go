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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"erc1155_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorAlreadyCastVote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorAlreadyQueuedProposal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorDisabledDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"GovernorInsufficientProposerVotes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldatas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"values\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidProposalLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorInvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorInvalidVoteType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"votingPeriod\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidVotingPeriod\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNonexistentProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNotQueuedProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyExecutor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyProposer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorQueueNotImplemented\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"GovernorRestrictedProposer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"current\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"expectedStates\",\"type\":\"bytes32\"}],\"name\":\"GovernorUnexpectedProposalState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueFull\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"etaSeconds\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"VoteCastWithParams\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BM_ERC1155\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLOCK_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTING_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castVoteWithReason\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParamsBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clock\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"getVotesWithParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalEta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalNeedsQueuing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"proposalStateToClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"hasVoted_\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"supported\",\"type\":\"uint8\"},{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"state_\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC5805\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101a060405234801562000011575f80fd5b50604051620039bc380380620039bc83398101604081905262000034916200020b565b808280620000566040805180820190915260018152603160f81b602082015290565b62000062825f62000134565b610120526200007381600162000134565b61014052815160208084019190912060e052815190820120610100524660a0526200010060e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c05260036200011782826200035d565b50506001600160a01b039081166101605216610180525062000481565b5f60208351101562000153576200014b836200016c565b905062000166565b816200016084826200035d565b5060ff90505b92915050565b5f80829050601f81511115620001a2578260405163305a27a960e01b815260040162000199919062000429565b60405180910390fd5b8051620001af826200045d565b179392505050565b634e487b7160e01b5f52604160045260245ffd5b5f5b83811015620001e7578181015183820152602001620001cd565b50505f910152565b80516001600160a01b038116811462000206575f80fd5b919050565b5f80604083850312156200021d575f80fd5b82516001600160401b038082111562000234575f80fd5b818501915085601f83011262000248575f80fd5b8151818111156200025d576200025d620001b7565b604051601f8201601f19908116603f01168101908382118183101715620002885762000288620001b7565b81604052828152886020848701011115620002a1575f80fd5b620002b4836020830160208801620001cb565b8096505050505050620002ca60208401620001ef565b90509250929050565b600181811c90821680620002e857607f821691505b6020821081036200030757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156200035857805f5260205f20601f840160051c81016020851015620003345750805b601f840160051c820191505b8181101562000355575f815560010162000340565b50505b505050565b81516001600160401b03811115620003795762000379620001b7565b62000391816200038a8454620002d3565b846200030d565b602080601f831160018114620003c7575f8415620003af5750858301515b5f19600386901b1c1916600185901b17855562000421565b5f85815260208120601f198616915b82811015620003f757888601518255948401946001909101908401620003d6565b50858210156200041557878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b602081525f825180602084015262000449816040850160208701620001cb565b601f01601f19169190910160400192915050565b8051602080830151919081101562000307575f1960209190910360031b1b16919050565b60805160a05160c05160e05161010051610120516101405161016051610180516134a4620005185f395f818161076401528181610f2401528181610fb201528181611325015261179c01525f81816108440152818161161401528181611998015261226e01525f61196f01525f61194301525f611b8a01525f611b6201525f611abd01525f611ae701525f611b1101526134a45ff3fe608060405260043610610241575f3560e01c80637d5e81e211610134578063bc197c81116100b3578063dd4e2ba511610078578063dd4e2ba514610786578063deaaa7cc1461079a578063eb9019d4146107cd578063f23a6e61146107ec578063f8ce560a14610817578063fc0c546a14610836575f80fd5b8063bc197c81146106d7578063c01f9e3714610702578063c28bc2fa14610721578063c59057e414610734578063c766c11414610753575f80fd5b80639a802a6d116100f95780639a802a6d1461061b578063a47516e11461063a578063a9a9529414610668578063ab58fb8e14610687578063b58131b0146106bd575f80fd5b80637d5e81e2146105575780637ecebe001461057657806384b0196e146105aa5780638ff262e3146105d157806391ddadf4146105f0575f80fd5b80633e4f49e6116101c057806354fd4d501161018557806354fd4d50146104b257806356781388146104db5780635b8d0e0d146104fa5780635f398a14146105195780637b3c71d314610538575f80fd5b80633e4f49e6146103e1578063438596321461040d578063452115d61461042c5780634bf5d7e91461044b578063544ffc9c1461045f575f80fd5b8063160cbed711610206578063160cbed7146103495780632656227d146103685780632d63f6931461037b5780632fe3e2611461039a5780633932abb1146103cd575f80fd5b806301ffc9a71461024e57806302a251a31461028257806306fdde03146102a4578063143489d0146102c5578063150b7a0214610311575f80fd5b3661024a57005b005b5f80fd5b348015610259575f80fd5b5061026d6102683660046125c5565b610868565b60405190151581526020015b60405180910390f35b34801561028d575f80fd5b506102966108b9565b604051908152602001610279565b3480156102af575f80fd5b506102b86108cf565b6040516102799190612639565b3480156102d0575f80fd5b506102f96102df36600461264b565b5f908152600460205260409020546001600160a01b031690565b6040516001600160a01b039091168152602001610279565b34801561031c575f80fd5b5061033061032b36600461273b565b61095f565b6040516001600160e01b03199091168152602001610279565b348015610354575f80fd5b50610296610363366004612904565b610970565b610296610376366004612904565b6109ae565b348015610386575f80fd5b5061029661039536600461264b565b610ad6565b3480156103a5575f80fd5b506102967f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a81181565b3480156103d8575f80fd5b50610296610af6565b3480156103ec575f80fd5b506104006103fb36600461264b565b610b1c565b60405161027991906129c1565b348015610418575f80fd5b5061026d6104273660046129cf565b610c45565b348015610437575f80fd5b50610296610446366004612904565b610c75565b348015610456575f80fd5b506102b8610ce1565b34801561046a575f80fd5b5061049761047936600461264b565b5f908152600760205260409020805460018201546002909201549092565b60408051938452602084019290925290820152606001610279565b3480156104bd575f80fd5b506040805180820190915260018152603160f81b60208201526102b8565b3480156104e6575f80fd5b506102966104f5366004612a09565b610ceb565b348015610505575f80fd5b50610296610514366004612a6e565b610d0a565b348015610524575f80fd5b50610296610533366004612b1e565b610e66565b348015610543575f80fd5b50610296610552366004612b9b565b610eb9565b348015610562575f80fd5b50610296610571366004612bf0565b610eff565b348015610581575f80fd5b50610296610590366004612c9c565b6001600160a01b03165f9081526002602052604090205490565b3480156105b5575f80fd5b506105be611041565b6040516102799796959493929190612cef565b3480156105dc575f80fd5b506102966105eb366004612d5e565b611083565b3480156105fb575f80fd5b50610604611152565b60405165ffffffffffff9091168152602001610279565b348015610626575f80fd5b50610296610635366004612da9565b61115b565b348015610645575f80fd5b506106596106543660046129cf565b611167565b60405161027993929190612dfb565b348015610673575f80fd5b5061026d61068236600461264b565b505f90565b348015610692575f80fd5b506102966106a136600461264b565b5f9081526004602052604090206001015465ffffffffffff1690565b3480156106c8575f80fd5b50670de0b6b3a7640000610296565b3480156106e2575f80fd5b506103306106f1366004612e1a565b63bc197c8160e01b95945050505050565b34801561070d575f80fd5b5061029661071c36600461264b565b6111b0565b61024861072f366004612ea2565b6111f2565b34801561073f575f80fd5b5061029661074e366004612904565b61126e565b34801561075e575f80fd5b506102f97f000000000000000000000000000000000000000000000000000000000000000081565b348015610791575f80fd5b506102b86112a7565b3480156107a5575f80fd5b506102967ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d781565b3480156107d8575f80fd5b506102966107e7366004612edf565b6112e3565b3480156107f7575f80fd5b50610330610806366004612f07565b63f23a6e6160e01b95945050505050565b348015610822575f80fd5b5061029661083136600461264b565b611302565b348015610841575f80fd5b507f00000000000000000000000000000000000000000000000000000000000000006102f9565b5f6001600160e01b031982166332a2ad4360e11b148061089857506001600160e01b03198216630271189760e51b145b806108b357506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f6108ca6205460062093a80612f7a565b905090565b6060600380546108de90612f8d565b80601f016020809104026020016040519081016040528092919081815260200182805461090a90612f8d565b80156109555780601f1061092c57610100808354040283529160200191610955565b820191905f5260205f20905b81548152906001019060200180831161093857829003601f168201915b5050505050905090565b630a85bd0160e11b5b949350505050565b5f8061097e8686868661126e565b90506109938161098e6004611398565b6113ba565b505f604051634844252360e11b815260040160405180910390fd5b5f806109bc8686868661126e565b90506109dc816109cc6005611398565b6109d66004611398565b176113ba565b505f818152600460205260409020805460ff60f01b1916600160f01b17905530610a033090565b6001600160a01b031614610a8c575f5b8651811015610a8a57306001600160a01b0316878281518110610a3857610a38612fc5565b60200260200101516001600160a01b031603610a8257610a82858281518110610a6357610a63612fc5565b60200260200101518051906020012060056113f790919063ffffffff16565b600101610a13565b505b610a998187878787611467565b6040518181527f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f906020015b60405180910390a195945050505050565b5f90815260046020526040902054600160a01b900465ffffffffffff1690565b5f4262093a80810681036205460001620151808201811015610b16575f80fd5b03919050565b5f818152600460205260408120805460ff600160f01b8204811691600160f81b9004168115610b5057506007949350505050565b8015610b6157506002949350505050565b5f610b6b86610ad6565b9050805f03610b9557604051636ad0607560e01b8152600481018790526024015b60405180910390fd5b5f610b9e611152565b65ffffffffffff169050808210610bbb57505f9695505050505050565b5f610bc5886111b0565b9050818110610bdc57506001979650505050505050565b610be58861153c565b1580610bf75750610bf588611546565b155b15610c0a57506003979650505050505050565b5f8881526004602052604090206001015465ffffffffffff165f03610c3757506004979650505050505050565b506005979650505050505050565b5f8281526007602090815260408083206001600160a01b038516845260030190915281205460ff165b9392505050565b5f80610c838686868661126e565b9050610c928161098e5f611398565b505f818152600460205260409020546001600160a01b03163314610ccb5760405163233d98e360e01b8152336004820152602401610b8c565b610cd786868686611561565b9695505050505050565b60606108ca611610565b5f8033905061096884828560405180602001604052805f8152506116d0565b5f80610deb87610de57f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a8118c8c8c610d5d8e6001600160a01b03165f90815260026020526040902080546001810190915590565b8d8d604051610d6d929190612fd9565b60405180910390208c80519060200120604051602001610dca9796959493929190968752602087019590955260ff9390931660408601526001600160a01b03919091166060850152608084015260a083015260c082015260e00190565b604051602081830303815290604052805190602001206116fa565b85611726565b905080610e16576040516394ab6c0760e01b81526001600160a01b0388166004820152602401610b8c565b610e5989888a89898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508b925061177b915050565b9998505050505050505050565b5f80339050610eae87828888888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508a925061177b915050565b979650505050505050565b5f80339050610cd786828787878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506116d092505050565b5f80336040516306044ae960e41b81526001600160a01b0380831660048301529192507f000000000000000000000000000000000000000000000000000000000000000090911690636044ae90906024015f604051808303815f87803b158015610f67575f80fd5b505af1158015610f79573d5f803e3d5ffd5b505050505f610f8a87878787611802565b9050610cd7818360026040516309ab24eb60e41b81526001600160a01b0387811660048301527f00000000000000000000000000000000000000000000000000000000000000001690639ab24eb090602401602060405180830381865afa158015610ff7573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061101b9190612fe8565b60405180604001604052806008815260200167383937b837b9b2b960c11b8152506118ba565b5f6060805f805f606061105261193c565b61105a611968565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b5f8061110d84610de57ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d78989896110d68b6001600160a01b03165f90815260026020526040902080546001810190915590565b60408051602081019690965285019390935260ff90911660608401526001600160a01b0316608083015260a082015260c001610dca565b905080611138576040516394ab6c0760e01b81526001600160a01b0385166004820152602401610b8c565b610cd786858760405180602001604052805f8152506116d0565b5f6108ca611995565b5f610968848484611a1c565b5f805f6111748585610c45565b5f8681526008602090815260408083206001600160a01b038916845290915290205490935060ff1691506111a785610b1c565b90509250925092565b5f818152600460205260408120546111e490600160d01b810463ffffffff1690600160a01b900465ffffffffffff16612fff565b65ffffffffffff1692915050565b6111fa611a28565b5f80856001600160a01b0316858585604051611217929190612fd9565b5f6040518083038185875af1925050503d805f8114611251576040519150601f19603f3d011682016040523d82523d5f602084013e611256565b606091505b50915091506112658282611a5f565b50505050505050565b5f8484848460405160200161128694939291906130b5565b60408051601f19818403018152919052805160209091012095945050505050565b60606108ca6040805180820190915260208082527f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e9082015290565b5f610c6e83836112fd60408051602081019091525f815290565b611a1c565b604051632394e7a360e21b8152600481018290525f90600a906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638e539e8c90602401602060405180830381865afa15801561136a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061138e9190612fe8565b6108b391906130ff565b5f8160078111156113ab576113ab61298d565b600160ff919091161b92915050565b5f806113c584610b1c565b90505f836113d283611398565b1603610c6e578381846040516331b75e4d60e01b8152600401610b8c9392919061311e565b81546001600160801b03600160801b82048116918116600183019091160361143257604051638acb5f2760e01b815260040160405180910390fd5b6001600160801b038082165f90815260018086016020526040909120939093558354919092018216600160801b029116179055565b5f5b8451811015611534575f8086838151811061148657611486612fc5565b60200260200101516001600160a01b03168684815181106114a9576114a9612fc5565b60200260200101518685815181106114c3576114c3612fc5565b60200260200101516040516114d89190613140565b5f6040518083038185875af1925050503d805f8114611512576040519150601f19603f3d011682016040523d82523d5f602084013e611517565b606091505b50915091506115268282611a5f565b505050806001019050611469565b505050505050565b5f6108b382611a7b565b5f8181526007602052604081208054600190910154116108b3565b5f8061156f8686868661126e565b90506115bd8161157f6007611398565b6115896006611398565b6115936002611398565b60016115a060078261315b565b6115ab906002613254565b6115b59190612f7a565b1818186113ba565b505f818152600460205260409081902080546001600160f81b0316600160f81b179055517f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c90610ac59083815260200190565b60607f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634bf5d7e96040518163ffffffff1660e01b81526004015f60405180830381865afa92505050801561168f57506040513d5f823e601f3d908101601f1916820160405261168c9190810190613262565b60015b6116cb575060408051808201909152601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000602082015290565b919050565b5f6116f1858585856116ec60408051602081019091525f815290565b61177b565b95945050505050565b5f6108b3611706611ab1565b8360405161190160f01b8152600281019290925260228201526042902090565b5f805f6117338585611bda565b5090925090505f81600381111561174c5761174c61298d565b14801561176a5750856001600160a01b0316826001600160a01b0316145b80610cd75750610cd7868686611c23565b6040516306044ae960e41b81526001600160a01b0385811660048301525f917f000000000000000000000000000000000000000000000000000000000000000090911690636044ae90906024015f604051808303815f87803b1580156117df575f80fd5b505af11580156117f1573d5f803e3d5ffd5b50505050610cd78686868686611cf9565b5f3361180e8184611dd3565b6118365760405163d9b3955760e01b81526001600160a01b0382166004820152602401610b8c565b5f61185c826001611845611152565b61184f91906132ca565b65ffffffffffff166112e3565b9050670de0b6b3a7640000808210156118a157604051636121770b60e11b81526001600160a01b03841660048201526024810183905260448101829052606401610b8c565b6118ae8888888887611ebb565b98975050505050505050565b815f036118f75760405162461bcd60e51b815260206004820152600b60248201526a1e995c9bc81dd95a59da1d60aa1b6044820152606401610b8c565b61190485858585856120c6565b50505f9283526008602090815260408085206001600160a01b039490941685529290529120805460ff191660ff909216919091179055565b60606108ca7f00000000000000000000000000000000000000000000000000000000000000005f6121b8565b60606108ca7f000000000000000000000000000000000000000000000000000000000000000060016121b8565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611a10575060408051601f3d908101601f19168201909252611a0d918101906132e9565b60015b6116cb576108ca612261565b5f61096884848461226b565b303314611a4a576040516347096e4760e01b8152336004820152602401610b8c565b565b80611a5760056122fe565b03611a4c5750565b606082611a7457611a6f8261237a565b6108b3565b50806108b3565b5f81815260076020526040812060028101546001820154611a9c919061330e565b611aa861083185610ad6565b11159392505050565b5f306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015611b0957507f000000000000000000000000000000000000000000000000000000000000000046145b15611b3357507f000000000000000000000000000000000000000000000000000000000000000090565b6108ca604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f805f8351604103611c11576020840151604085015160608601515f1a611c03888285856123a3565b955095509550505050611c1c565b505081515f91506002905b9250925092565b5f805f856001600160a01b03168585604051602401611c43929190613321565b60408051601f198184030181529181526020820180516001600160e01b0316630b135d3f60e11b17905251611c789190613140565b5f60405180830381855afa9150503d805f8114611cb0576040519150601f19603f3d011682016040523d82523d5f602084013e611cb5565b606091505b5091509150818015611cc957506020815110155b8015610cd757508051630b135d3f60e11b90611cee9083016020908101908401612fe8565b149695505050505050565b5f611d088661098e6001611398565b505f611d1d86611d1789610ad6565b85611a1c565b9050611d2c87878784876118ba565b82515f03611d8057856001600160a01b03167fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda488878488604051611d739493929190613339565b60405180910390a2610cd7565b856001600160a01b03167fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb87128887848888604051611dc1959493929190613360565b60405180910390a29695505050505050565b80515f906034811015611dea5760019150506108b3565b82810160131901516001600160a01b031981166b046e0e4dee0dee6cae47a60f60a31b14611e1d576001925050506108b3565b5f80611e2a602885612f7a565b90505b83811015611e9a575f80611e60888481518110611e4c57611e4c612fc5565b01602001516001600160f81b03191661246b565b9150915081611e7857600196505050505050506108b3565b8060ff166004856001600160a01b0316901b1793505050806001019050611e2d565b50856001600160a01b0316816001600160a01b031614935050505092915050565b5f611ecf868686868051906020012061126e565b905084518651141580611ee457508351865114155b80611eee57508551155b15611f2357855184518651604051630447b05d60e41b8152600481019390935260248301919091526044820152606401610b8c565b5f81815260046020526040902054600160a01b900465ffffffffffff1615611f6c5780611f4f82610b1c565b6040516331b75e4d60e01b8152610b8c9291905f9060040161311e565b5f611f75610af6565b611f7d611152565b65ffffffffffff16611f8f919061330e565b90505f611f9a6108b9565b5f84815260046020526040902080546001600160a01b0319166001600160a01b038716178155909150611fcc836124fb565b815465ffffffffffff91909116600160a01b0265ffffffffffff60a01b19909116178155611ff982612531565b815463ffffffff91909116600160d01b0263ffffffff60d01b1990911617815588517f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e090859087908c908c906001600160401b0381111561205c5761205c612678565b60405190808252806020026020018201604052801561208f57816020015b606081526020019060019003908161207a5790505b508c8961209c8a8261330e565b8e6040516120b299989796959493929190613399565b60405180910390a150505095945050505050565b5f8581526007602090815260408083206001600160a01b0388168452600381019092529091205460ff1615612119576040516371c6af4960e01b81526001600160a01b0386166004820152602401610b8c565b6001600160a01b0385165f9081526003820160205260409020805460ff1916600117905560ff84166121625782815f015f828254612157919061330e565b909155506115349050565b5f1960ff8516016121805782816001015f828254612157919061330e565b60011960ff85160161219f5782816002015f828254612157919061330e565b6040516303599be160e11b815260040160405180910390fd5b606060ff83146121d2576121cb83612561565b90506108b3565b8180546121de90612f8d565b80601f016020809104026020016040519081016040528092919081815260200182805461220a90612f8d565b80156122555780601f1061222c57610100808354040283529160200191612255565b820191905f5260205f20905b81548152906001019060200180831161223857829003601f168201915b505050505090506108b3565b5f6108ca436124fb565b5f7f0000000000000000000000000000000000000000000000000000000000000000604051630748d63560e31b81526001600160a01b038681166004830152602482018690529190911690633a46b1a890604401602060405180830381865afa1580156122da573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109689190612fe8565b80545f906001600160801b0380821691600160801b9004168103612335576040516375e52f4f60e01b815260040160405180910390fd5b6001600160801b038181165f908152600185810160205260408220805492905585546fffffffffffffffffffffffffffffffff19169301909116919091179092555090565b80511561238a5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156123dc57505f91506003905082612461565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa15801561242d573d5f803e3d5ffd5b5050604051601f1901519150506001600160a01b03811661245857505f925060019150829050612461565b92505f91508190505b9450945094915050565b5f8060f883901c602f811180156124855750603a8160ff16105b1561249a57600194602f199091019350915050565b8060ff1660401080156124b0575060478160ff16105b156124c5576001946036199091019350915050565b8060ff1660601080156124db575060678160ff16105b156124f0576001946056199091019350915050565b505f93849350915050565b5f65ffffffffffff82111561252d576040516306dfcc6560e41b81526030600482015260248101839052604401610b8c565b5090565b5f63ffffffff82111561252d576040516306dfcc6560e41b81526020600482015260248101839052604401610b8c565b60605f61256d8361259e565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f60ff8216601f8111156108b357604051632cd44ac360e21b815260040160405180910390fd5b5f602082840312156125d5575f80fd5b81356001600160e01b031981168114610c6e575f80fd5b5f5b838110156126065781810151838201526020016125ee565b50505f910152565b5f81518084526126258160208601602086016125ec565b601f01601f19169290920160200192915050565b602081525f610c6e602083018461260e565b5f6020828403121561265b575f80fd5b5035919050565b80356001600160a01b03811681146116cb575f80fd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156126b4576126b4612678565b604052919050565b5f6001600160401b038211156126d4576126d4612678565b50601f01601f191660200190565b5f6126f46126ef846126bc565b61268c565b9050828152838383011115612707575f80fd5b828260208301375f602084830101529392505050565b5f82601f83011261272c575f80fd5b610c6e838335602085016126e2565b5f805f806080858703121561274e575f80fd5b61275785612662565b935061276560208601612662565b92506040850135915060608501356001600160401b03811115612786575f80fd5b6127928782880161271d565b91505092959194509250565b5f6001600160401b038211156127b6576127b6612678565b5060051b60200190565b5f82601f8301126127cf575f80fd5b813560206127df6126ef8361279e565b8083825260208201915060208460051b870101935086841115612800575f80fd5b602086015b848110156128235761281681612662565b8352918301918301612805565b509695505050505050565b5f82601f83011261283d575f80fd5b8135602061284d6126ef8361279e565b8083825260208201915060208460051b87010193508684111561286e575f80fd5b602086015b848110156128235780358352918301918301612873565b5f82601f830112612899575f80fd5b813560206128a96126ef8361279e565b82815260059290921b840181019181810190868411156128c7575f80fd5b8286015b848110156128235780356001600160401b038111156128e8575f80fd5b6128f68986838b010161271d565b8452509183019183016128cb565b5f805f8060808587031215612917575f80fd5b84356001600160401b038082111561292d575f80fd5b612939888389016127c0565b9550602087013591508082111561294e575f80fd5b61295a8883890161282e565b9450604087013591508082111561296f575f80fd5b5061297c8782880161288a565b949793965093946060013593505050565b634e487b7160e01b5f52602160045260245ffd5b600881106129bd57634e487b7160e01b5f52602160045260245ffd5b9052565b602081016108b382846129a1565b5f80604083850312156129e0575f80fd5b823591506129f060208401612662565b90509250929050565b803560ff811681146116cb575f80fd5b5f8060408385031215612a1a575f80fd5b823591506129f0602084016129f9565b5f8083601f840112612a3a575f80fd5b5081356001600160401b03811115612a50575f80fd5b602083019150836020828501011115612a67575f80fd5b9250929050565b5f805f805f805f60c0888a031215612a84575f80fd5b87359650612a94602089016129f9565b9550612aa260408901612662565b945060608801356001600160401b0380821115612abd575f80fd5b612ac98b838c01612a2a565b909650945060808a0135915080821115612ae1575f80fd5b612aed8b838c0161271d565b935060a08a0135915080821115612b02575f80fd5b50612b0f8a828b0161271d565b91505092959891949750929550565b5f805f805f60808688031215612b32575f80fd5b85359450612b42602087016129f9565b935060408601356001600160401b0380821115612b5d575f80fd5b612b6989838a01612a2a565b90955093506060880135915080821115612b81575f80fd5b50612b8e8882890161271d565b9150509295509295909350565b5f805f8060608587031215612bae575f80fd5b84359350612bbe602086016129f9565b925060408501356001600160401b03811115612bd8575f80fd5b612be487828801612a2a565b95989497509550505050565b5f805f8060808587031215612c03575f80fd5b84356001600160401b0380821115612c19575f80fd5b612c25888389016127c0565b95506020870135915080821115612c3a575f80fd5b612c468883890161282e565b94506040870135915080821115612c5b575f80fd5b612c678883890161288a565b93506060870135915080821115612c7c575f80fd5b508501601f81018713612c8d575f80fd5b612792878235602084016126e2565b5f60208284031215612cac575f80fd5b610c6e82612662565b5f815180845260208085019450602084015f5b83811015612ce457815187529582019590820190600101612cc8565b509495945050505050565b60ff60f81b8816815260e060208201525f612d0d60e083018961260e565b8281036040840152612d1f818961260e565b606084018890526001600160a01b038716608085015260a0840186905283810360c08501529050612d508185612cb5565b9a9950505050505050505050565b5f805f8060808587031215612d71575f80fd5b84359350612d81602086016129f9565b9250612d8f60408601612662565b915060608501356001600160401b03811115612786575f80fd5b5f805f60608486031215612dbb575f80fd5b612dc484612662565b92506020840135915060408401356001600160401b03811115612de5575f80fd5b612df18682870161271d565b9150509250925092565b831515815260ff831660208201526060810161096860408301846129a1565b5f805f805f60a08688031215612e2e575f80fd5b612e3786612662565b9450612e4560208701612662565b935060408601356001600160401b0380821115612e60575f80fd5b612e6c89838a0161282e565b94506060880135915080821115612e81575f80fd5b612e8d89838a0161282e565b93506080880135915080821115612b81575f80fd5b5f805f8060608587031215612eb5575f80fd5b612ebe85612662565b93506020850135925060408501356001600160401b03811115612bd8575f80fd5b5f8060408385031215612ef0575f80fd5b612ef983612662565b946020939093013593505050565b5f805f805f60a08688031215612f1b575f80fd5b612f2486612662565b9450612f3260208701612662565b9350604086013592506060860135915060808601356001600160401b03811115612f5a575f80fd5b612b8e8882890161271d565b634e487b7160e01b5f52601160045260245ffd5b818103818111156108b3576108b3612f66565b600181811c90821680612fa157607f821691505b602082108103612fbf57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b818382375f9101908152919050565b5f60208284031215612ff8575f80fd5b5051919050565b65ffffffffffff81811683821601908082111561301e5761301e612f66565b5092915050565b5f815180845260208085019450602084015f5b83811015612ce45781516001600160a01b031687529582019590820190600101613038565b5f8282518085526020808601955060208260051b840101602086015f5b848110156130a857601f1986840301895261309683835161260e565b9884019892509083019060010161307a565b5090979650505050505050565b608081525f6130c76080830187613025565b82810360208401526130d98187612cb5565b905082810360408401526130ed818661305d565b91505082606083015295945050505050565b5f8261311957634e487b7160e01b5f52601260045260245ffd5b500490565b8381526060810161313260208301856129a1565b826040830152949350505050565b5f82516131518184602087016125ec565b9190910192915050565b60ff81811683821601908111156108b3576108b3612f66565b600181815b808511156131ae57815f190482111561319457613194612f66565b808516156131a157918102915b93841c9390800290613179565b509250929050565b5f826131c4575060016108b3565b816131d057505f6108b3565b81600181146131e657600281146131f05761320c565b60019150506108b3565b60ff84111561320157613201612f66565b50506001821b6108b3565b5060208310610133831016604e8410600b841016171561322f575081810a6108b3565b6132398383613174565b805f190482111561324c5761324c612f66565b029392505050565b5f610c6e60ff8416836131b6565b5f60208284031215613272575f80fd5b81516001600160401b03811115613287575f80fd5b8201601f81018413613297575f80fd5b80516132a56126ef826126bc565b8181528560208385010111156132b9575f80fd5b6116f18260208301602086016125ec565b65ffffffffffff82811682821603908082111561301e5761301e612f66565b5f602082840312156132f9575f80fd5b815165ffffffffffff81168114610c6e575f80fd5b808201808211156108b3576108b3612f66565b828152604060208201525f610968604083018461260e565b84815260ff84166020820152826040820152608060608201525f610cd7608083018461260e565b85815260ff8516602082015283604082015260a060608201525f61338760a083018561260e565b82810360808401526118ae818561260e565b5f6101208b8352602060018060a01b038c16818501528160408501526133c18285018c613025565b915083820360608501526133d5828b612cb5565b915083820360808501528189518084528284019150828160051b850101838c015f5b8381101561342557601f1987840301855261341383835161260e565b948601949250908501906001016133f7565b505086810360a0880152613439818c61305d565b9450505050508560c08401528460e084015282810361010084015261345e818561260e565b9c9b50505050505050505050505056fea2646970667358221220b04032e144f6e2ae524298de4d5aa4adb1357323afa099b0b304d53a2e33234064736f6c63430008180033",
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

// ProposalStateToClaim is a free data retrieval call binding the contract method 0xa47516e1.
//
// Solidity: function proposalStateToClaim(uint256 proposalID, address account) view returns(bool hasVoted_, uint8 supported, uint8 state_)
func (_BmGovernor *BmGovernorCaller) ProposalStateToClaim(opts *bind.CallOpts, proposalID *big.Int, account common.Address) (struct {
	HasVoted  bool
	Supported uint8
	State     uint8
}, error) {
	var out []interface{}
	err := _BmGovernor.contract.Call(opts, &out, "proposalStateToClaim", proposalID, account)

	outstruct := new(struct {
		HasVoted  bool
		Supported uint8
		State     uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HasVoted = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Supported = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.State = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// ProposalStateToClaim is a free data retrieval call binding the contract method 0xa47516e1.
//
// Solidity: function proposalStateToClaim(uint256 proposalID, address account) view returns(bool hasVoted_, uint8 supported, uint8 state_)
func (_BmGovernor *BmGovernorSession) ProposalStateToClaim(proposalID *big.Int, account common.Address) (struct {
	HasVoted  bool
	Supported uint8
	State     uint8
}, error) {
	return _BmGovernor.Contract.ProposalStateToClaim(&_BmGovernor.CallOpts, proposalID, account)
}

// ProposalStateToClaim is a free data retrieval call binding the contract method 0xa47516e1.
//
// Solidity: function proposalStateToClaim(uint256 proposalID, address account) view returns(bool hasVoted_, uint8 supported, uint8 state_)
func (_BmGovernor *BmGovernorCallerSession) ProposalStateToClaim(proposalID *big.Int, account common.Address) (struct {
	HasVoted  bool
	Supported uint8
	State     uint8
}, error) {
	return _BmGovernor.Contract.ProposalStateToClaim(&_BmGovernor.CallOpts, proposalID, account)
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
