// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Bridge

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BridgeProposal is an auto generated low-level Go binding around an user-defined struct.
type BridgeProposal struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	YesVotes      []common.Address
	NoVotes       []common.Address
	Status        uint8
	ProposedBlock *big.Int
}

// BridgeABI is the input ABI used to generate the binding from.
// const BridgeABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"domainID\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"initialRelayers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"initialRelayerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"destinationDomainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"handlerResponse\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelData\",\"type\":\"bytes\"}],\"name\":\"FailedHandlerExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"originDomainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enum Bridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"originDomainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enum Bridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"RelayerThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_RELAYERS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"_depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_domainID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_expiry\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fee\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint72\",\"name\":\"destNonce\",\"type\":\"uint72\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"_hasVotedOnProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_relayerThreshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalRelayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"adminAddRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"adminChangeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"adminChangeRelayerThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"adminRemoveRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"domainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"adminSetDepositNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"adminSetForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"depositFunctionDepositerOffset\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"adminSetGenericResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminUnpauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"adminWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"domainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"cancelProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationDomainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"domainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"revertOnFail\",\"type\":\"bool\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"originDomainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"enum Bridge.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint200\",\"name\":\"_yesVotes\",\"type\":\"uint200\"},{\"internalType\":\"uint8\",\"name\":\"_yesVotesTotal\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"_proposedBlock\",\"type\":\"uint40\"}],\"internalType\":\"struct Bridge.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getRoleMemberIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isValidForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address payable[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"transferFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"domainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"voteProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
const BridgeABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"initialRelayers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"initialRelayerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destinationChainID\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"originChainID\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"enum Bridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"originChainID\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"enum Bridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"ProposalVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"RelayerThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_chainID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"_depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_expiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint72\",\"name\":\"\",\"type\":\"uint72\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_hasVotedOnProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint72\",\"name\":\"\",\"type\":\"uint72\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_proposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"enum Bridge.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_relayerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalProposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalRelayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminUnpauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"adminChangeRelayerThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"adminAddRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"adminRemoveRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"adminSetGenericResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"originChainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_yesVotes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_noVotes\",\"type\":\"address[]\"},{\"internalType\":\"enum Bridge.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"}],\"internalType\":\"struct Bridge.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"adminChangeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOrTokenID\",\"type\":\"uint256\"}],\"name\":\"adminWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"voteProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"cancelProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address payable[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"transferFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x60806040523480156200001157600080fd5b5060405162004d8938038062004d898339818101604052810190620000379190620004bc565b60008060006101000a81548160ff02191690831515021790555084600260006101000a81548160ff021916908360ff160217905550826003819055508160068190555080600781905550620000966000801b336200013460201b60201c565b620000c0604051620000a89062000601565b60405180910390206000801b6200014a60201b60201c565b60005b8451811015620001285762000108604051620000df9062000601565b6040518091039020868381518110620000f457fe5b60200260200101516200016960201b60201c565b6004600081548092919060010191905055508080600101915050620000c3565b50505050505062000746565b620001468282620001f860201b60201c565b5050565b8060016000848152602001908152602001600020600201819055505050565b620001a06001600084815260200190815260200160002060020154620001946200029c60201b60201c565b620002a460201b60201c565b620001e2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001d99062000618565b60405180910390fd5b620001f48282620001f860201b60201c565b5050565b620002278160016000858152602001908152602001600020600001620002dd60201b6200285b1790919060201c565b1562000298576200023d6200029c60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b6000620002d582600160008681526020019081526020016000206000016200031560201b620027601790919060201c565b905092915050565b60006200030d836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200034d60201b60201c565b905092915050565b600062000345836000018373ffffffffffffffffffffffffffffffffffffffff1660001b620003c760201b60201c565b905092915050565b6000620003618383620003c760201b60201c565b620003bc578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050620003c1565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081519050620003fb81620006f8565b92915050565b600082601f8301126200041357600080fd5b81516200042a620004248262000668565b6200063a565b915081818352602084019350602081019050838560208402820111156200045057600080fd5b60005b83811015620004845781620004698882620003ea565b84526020840193506020830192505060018101905062000453565b5050505092915050565b6000815190506200049f8162000712565b92915050565b600081519050620004b6816200072c565b92915050565b600080600080600060a08688031215620004d557600080fd5b6000620004e588828901620004a5565b955050602086015167ffffffffffffffff8111156200050357600080fd5b620005118882890162000401565b945050604062000524888289016200048e565b935050606062000537888289016200048e565b92505060806200054a888289016200048e565b9150509295509295909350565b600062000566602f8362000691565b91507f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008301527f2061646d696e20746f206772616e7400000000000000000000000000000000006020830152604082019050919050565b6000620005ce600c83620006a2565b91507f52454c415945525f524f4c4500000000000000000000000000000000000000006000830152600c82019050919050565b60006200060e82620005bf565b9150819050919050565b60006020820190508181036000830152620006338162000557565b9050919050565b6000604051905081810181811067ffffffffffffffff821117156200065e57600080fd5b8060405250919050565b600067ffffffffffffffff8211156200068057600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b600081905092915050565b6000620006ba82620006c1565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6200070381620006ad565b81146200070f57600080fd5b50565b6200071d81620006e1565b81146200072957600080fd5b50565b6200073781620006eb565b81146200074357600080fd5b50565b61463380620007566000396000f3fe6080604052600436106102305760003560e01c806384db809f1161012e578063beab7131116100ab578063cdb0f73a1161006f578063cdb0f73a14610877578063d547741f146108a0578063d7a9cd79146108c9578063e8437ee7146108f4578063ffaac0eb1461091d57610230565b8063beab713114610790578063c5b37c22146107bb578063c5ec8970146107e6578063ca15c87314610811578063cb10f2151461084e57610230565b8063926d7d7f116100f2578063926d7d7f146106a95780639d5773e0146106d45780639d82dd63146106ff578063a217fddf14610728578063a9cf69fa1461075357610230565b806384db809f146105a05780638c0c2631146105dd5780639010d07c1461060657806391c404ac1461064357806391d148541461066c57610230565b80634b0b919d116101bc5780635e1fab0f116101805780635e1fab0f146104cf578063780cf004146104f85780637febe63f14610521578063802aabe81461055e57806380ae1c281461058957610230565b80634b0b919d146103c15780634e056005146103fe5780635059871914610427578063541d5548146104675780635c975abb146104a457610230565b80632f2ff15d116102035780632f2ff15d146102e057806336568abe146103095780633ee7094a146103325780634454b20d1461036f5780634603ae381461039857610230565b806305e2ca171461023557806317f03ce5146102515780631ff013f11461027a578063248a9ca3146102a3575b600080fd5b61024f600480360381019061024a9190613229565b610934565b005b34801561025d57600080fd5b5061027860048036038101906102739190613295565b610b91565b005b34801561028657600080fd5b506102a1600480360381019061029c91906132e4565b610d3e565b005b3480156102af57600080fd5b506102ca60048036038101906102c5919061306f565b6114bf565b6040516102d79190613df0565b60405180910390f35b3480156102ec57600080fd5b5061030760048036038101906103029190613098565b6114df565b005b34801561031557600080fd5b50610330600480360381019061032b9190613098565b611553565b005b34801561033e57600080fd5b5061035960048036038101906103549190613139565b6115d6565b6040516103669190613f75565b60405180910390f35b34801561037b57600080fd5b5061039660048036038101906103919190613347565b611693565b005b3480156103a457600080fd5b506103bf60048036038101906103ba9190612ffa565b6119d4565b005b3480156103cd57600080fd5b506103e860048036038101906103e39190613200565b611a7a565b6040516103f591906142d6565b60405180910390f35b34801561040a57600080fd5b5061042560048036038101906104209190613110565b611aa1565b005b34801561043357600080fd5b5061044e60048036038101906104499190613175565b611ae0565b60405161045e9493929190613ea2565b60405180910390f35b34801561047357600080fd5b5061048e60048036038101906104899190612e43565b611b2a565b60405161049b9190613dd5565b60405180910390f35b3480156104b057600080fd5b506104b9611b50565b6040516104c69190613dd5565b60405180910390f35b3480156104db57600080fd5b506104f660048036038101906104f19190612e43565b611b66565b005b34801561050457600080fd5b5061051f600480360381019061051a9190612ed1565b611b8b565b005b34801561052d57600080fd5b50610548600480360381019061054391906131b1565b611c0e565b6040516105559190613dd5565b60405180910390f35b34801561056a57600080fd5b50610573611c4a565b60405161058091906142bb565b60405180910390f35b34801561059557600080fd5b5061059e611c50565b005b3480156105ac57600080fd5b506105c760048036038101906105c2919061306f565b611c62565b6040516105d49190613d68565b60405180910390f35b3480156105e957600080fd5b5061060460048036038101906105ff9190612e95565b611c95565b005b34801561061257600080fd5b5061062d600480360381019061062891906130d4565b611d12565b60405161063a9190613d68565b60405180910390f35b34801561064f57600080fd5b5061066a60048036038101906106659190613110565b611d44565b005b34801561067857600080fd5b50610693600480360381019061068e9190613098565b611d9b565b6040516106a09190613dd5565b60405180910390f35b3480156106b557600080fd5b506106be611dcd565b6040516106cb9190613df0565b60405180910390f35b3480156106e057600080fd5b506106e9611de4565b6040516106f691906142bb565b60405180910390f35b34801561070b57600080fd5b5061072660048036038101906107219190612e43565b611dea565b005b34801561073457600080fd5b5061073d611ec4565b60405161074a9190613df0565b60405180910390f35b34801561075f57600080fd5b5061077a60048036038101906107759190613295565b611ecb565b6040516107879190614299565b60405180910390f35b34801561079c57600080fd5b506107a56120ac565b6040516107b291906142f1565b60405180910390f35b3480156107c757600080fd5b506107d06120bf565b6040516107dd91906142bb565b60405180910390f35b3480156107f257600080fd5b506107fb6120c5565b60405161080891906142bb565b60405180910390f35b34801561081d57600080fd5b506108386004803603810190610833919061306f565b6120cb565b60405161084591906142bb565b60405180910390f35b34801561085a57600080fd5b5061087560048036038101906108709190612f34565b6120f2565b005b34801561088357600080fd5b5061089e60048036038101906108999190612e43565b6121c4565b005b3480156108ac57600080fd5b506108c760048036038101906108c29190613098565b61229e565b005b3480156108d557600080fd5b506108de612312565b6040516108eb91906142bb565b60405180910390f35b34801561090057600080fd5b5061091b60048036038101906109169190612f83565b612318565b005b34801561092957600080fd5b506109326123f0565b005b61093c612402565b6006543414610980576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610977906140b9565b60405180910390fd5b60006009600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610a28576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1f90614119565b60405180910390fd5b6000600860008760ff1660ff168152602001908152602001600020600081819054906101000a900467ffffffffffffffff1660010191906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905590508383600a60008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008960ff1660ff1681526020019081526020016000209190610ad1929190612b04565b5060008290508073ffffffffffffffffffffffffffffffffffffffff166338995da9878985338a8a6040518763ffffffff1660e01b8152600401610b1a96959493929190613f19565b600060405180830381600087803b158015610b3457600080fd5b505af1158015610b48573d6000803e3d6000fd5b505050508167ffffffffffffffff16868860ff167fdbb69440df8433824a026ef190652f29929eb64b4d1d5d2a69be8afe3e6eaed860405160405180910390a450505050505050565b610b99612453565b60008360ff1660088467ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000209050600480811115610c0657fe5b8160040160009054906101000a900460ff166004811115610c2357fe5b1415610c64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c5b906140f9565b60405180910390fd5b600754610c754383600501546124c5565b11610cb5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cac90614159565b60405180910390fd5b60048160040160006101000a81548160ff02191690836004811115610cd657fe5b0217905550600480811115610ce757fe5b8467ffffffffffffffff168660ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041784600001548560010154604051610d2f929190613e79565b60405180910390a45050505050565b610d4661250f565b610d4e612402565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000209050600073ffffffffffffffffffffffffffffffffffffffff166009600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610e52576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4990614259565b60405180910390fd5b60018160040160009054906101000a900460ff166004811115610e7157fe5b1115610eb2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ea9906141b9565b60405180910390fd5b600c60008368ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610f77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6e90614059565b60405180910390fd5b60008160040160009054906101000a900460ff166004811115610f9657fe5b14156111b657600560008154600101919050819055506040518060c001604052808581526020018481526020016001604051908082528060200260200182016040528015610ff35781602001602082028036833780820191505090505b50815260200160006040519080825280602002602001820160405280156110295781602001602082028036833780820191505090505b5081526020016001600481111561103c57fe5b815260200143815250600b60008468ffffffffffffffffff1668ffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020600082015181600001556020820151816001015560408201518160020190805190602001906110af929190612b84565b5060608201518160030190805190602001906110cc929190612b84565b5060808201518160040160006101000a81548160ff021916908360048111156110f157fe5b021790555060a08201518160050155905050338160020160008154811061111457fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600481111561116957fe5b8567ffffffffffffffff168760ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041787876040516111a9929190613e79565b60405180910390a46112f9565b6007546111c74383600501546124c5565b111561124c5760048160040160006101000a81548160ff021916908360048111156111ee57fe5b02179055506004808111156111ff57fe5b8567ffffffffffffffff168760ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab650417878760405161123f929190613e79565b60405180910390a46112f8565b80600101548314611292576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161128990614239565b60405180910390fd5b80600201339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b60048081111561130557fe5b8160040160009054906101000a900460ff16600481111561132257fe5b146114b7576001600c60008468ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508060040160009054906101000a900460ff1660048111156113d457fe5b8567ffffffffffffffff168760ff167f25f8daaa4635a7729927ba3f5b3d59cc3320aca7c32c9db4e7ca7b9574343640876040516114129190613df0565b60405180910390a460016003541115806114355750600354816002018054905010155b156114b65760028160040160006101000a81548160ff0219169083600481111561145b57fe5b02179055506002600481111561146d57fe5b8567ffffffffffffffff168760ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041787876040516114ad929190613e79565b60405180910390a45b5b505050505050565b600060016000838152602001908152602001600020600201549050919050565b611506600160008481526020019081526020016000206002015461150161256d565b611d9b565b611545576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161153c90614039565b60405180910390fd5b61154f8282612575565b5050565b61155b61256d565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146115c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115bf90614279565b60405180910390fd5b6115d28282612609565b5050565b600a602052816000526040600020602052806000526040600020600091509150508054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561168b5780601f106116605761010080835404028352916020019161168b565b820191906000526020600020905b81548152906001019060200180831161166e57829003601f168201915b505050505081565b61169b61250f565b6116a3612402565b60006009600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008660ff1660088767ffffffffffffffff1668ffffffffffffffffff16901b179050600082868660405160200161171593929190613d29565b6040516020818303038152906040528051906020012090506000600b60008468ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002090506000600481111561177857fe5b8160040160009054906101000a900460ff16600481111561179557fe5b14156117d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117cd90613fd9565b60405180910390fd5b600260048111156117e357fe5b8160040160009054906101000a900460ff16600481111561180057fe5b14611840576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161183790613fb9565b60405180910390fd5b80600101548214611886576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161187d906140d9565b60405180910390fd5b60038160040160006101000a81548160ff021916908360048111156118a757fe5b02179055506000600960008360000154815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663e248cff283600001548a8a6040518463ffffffff1660e01b815260040161192993929190613ee7565b600060405180830381600087803b15801561194357600080fd5b505af1158015611957573d6000803e3d6000fd5b505050508160040160009054906101000a900460ff16600481111561197857fe5b8967ffffffffffffffff168b60ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab650417856000015486600101546040516119c0929190613e79565b60405180910390a450505050505050505050565b6119dc61269d565b60008090505b84849050811015611a73578484828181106119f957fe5b9050602002016020810190611a0e9190612e6c565b73ffffffffffffffffffffffffffffffffffffffff166108fc848484818110611a3357fe5b905060200201359081150290604051600060405180830381858888f19350505050158015611a65573d6000803e3d6000fd5b5080806001019150506119e2565b5050505050565b60086020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b611aa961269d565b80600381905550807fa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c860405160405180910390a250565b600b602052816000526040600020602052806000526040600020600091509150508060000154908060010154908060040160009054906101000a900460ff16908060050154905084565b6000611b49604051611b3b90613d53565b604051809103902083611d9b565b9050919050565b60008060009054906101000a900460ff16905090565b611b6e61269d565b611b7b6000801b826114df565b611b886000801b33611553565b50565b611b9361269d565b60008490508073ffffffffffffffffffffffffffffffffffffffff1663d9caed128585856040518463ffffffff1660e01b8152600401611bd593929190613d9e565b600060405180830381600087803b158015611bef57600080fd5b505af1158015611c03573d6000803e3d6000fd5b505050505050505050565b600c602052826000526040600020602052816000526040600020602052806000526040600020600092509250509054906101000a900460ff1681565b60045481565b611c5861269d565b611c606126eb565b565b60096020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611c9d61269d565b60008290508073ffffffffffffffffffffffffffffffffffffffff166307b7ed99836040518263ffffffff1660e01b8152600401611cdb9190613d68565b600060405180830381600087803b158015611cf557600080fd5b505af1158015611d09573d6000803e3d6000fd5b50505050505050565b6000611d3c826001600086815260200190815260200160002060000161274690919063ffffffff16565b905092915050565b611d4c61269d565b806006541415611d91576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d8890614219565b60405180910390fd5b8060068190555050565b6000611dc5826001600086815260200190815260200160002060000161276090919063ffffffff16565b905092915050565b604051611dd990613d53565b604051809103902081565b60055481565b611df261269d565b611e0f604051611e0190613d53565b604051809103902082611d9b565b611e4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e4590614099565b60405180910390fd5b611e6b604051611e5d90613d53565b60405180910390208261229e565b8073ffffffffffffffffffffffffffffffffffffffff167f10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b60405160405180910390a26004600081548092919060019003919050555050565b6000801b81565b611ed3612c0e565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b179050600b60008268ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206040518060c0016040529081600082015481526020016001820154815260200160028201805480602002602001604051908101604052809291908181526020018280548015611fd457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611f8a575b505050505081526020016003820180548060200260200160405190810160405280929190818152602001828054801561206257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612018575b505050505081526020016004820160009054906101000a900460ff16600481111561208957fe5b600481111561209457fe5b81526020016005820154815250509150509392505050565b600260009054906101000a900460ff1681565b60065481565b60075481565b60006120eb60016000848152602001908152602001600020600001612790565b9050919050565b6120fa61269d565b826009600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008390508073ffffffffffffffffffffffffffffffffffffffff1663b8fa373684846040518363ffffffff1660e01b815260040161218c929190613e0b565b600060405180830381600087803b1580156121a657600080fd5b505af11580156121ba573d6000803e3d6000fd5b5050505050505050565b6121cc61269d565b6121e96040516121db90613d53565b604051809103902082611d9b565b15612229576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161222090614199565b60405180910390fd5b61224660405161223890613d53565b6040518091039020826114df565b8073ffffffffffffffffffffffffffffffffffffffff167f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c560405160405180910390a260046000815480929190600101919050555050565b6122c560016000848152602001908152602001600020600201546122c061256d565b611d9b565b612304576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122fb90614139565b60405180910390fd5b61230e8282612609565b5050565b60035481565b61232061269d565b846009600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008590508073ffffffffffffffffffffffffffffffffffffffff1663bba8185a868686866040518563ffffffff1660e01b81526004016123b69493929190613e34565b600060405180830381600087803b1580156123d057600080fd5b505af11580156123e4573d6000803e3d6000fd5b50505050505050505050565b6123f861269d565b6124006127a5565b565b6000809054906101000a900460ff1615612451576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161244890614179565b60405180910390fd5b565b6124606000801b33611d9b565b80612484575061248360405161247590613d53565b604051809103902033611d9b565b5b6124c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124ba90614019565b60405180910390fd5b565b600061250783836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250612800565b905092915050565b61252c60405161251e90613d53565b604051809103902033611d9b565b61256b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612562906141d9565b60405180910390fd5b565b600033905090565b61259d816001600085815260200190815260200160002060000161285b90919063ffffffff16565b15612605576125aa61256d565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b612631816001600085815260200190815260200160002060000161288b90919063ffffffff16565b156126995761263e61256d565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6126aa6000801b33611d9b565b6126e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126e0906141f9565b60405180910390fd5b565b6126f3612402565b60016000806101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2583360405161273c9190613d83565b60405180910390a1565b600061275583600001836128bb565b60001c905092915050565b6000612788836000018373ffffffffffffffffffffffffffffffffffffffff1660001b612928565b905092915050565b600061279e8260000161294b565b9050919050565b6127ad61295c565b60008060006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa336040516127f69190613d83565b60405180910390a1565b6000838311158290612848576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161283f9190613f97565b60405180910390fd5b5060008385039050809150509392505050565b6000612883836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6129ac565b905092915050565b60006128b3836000018373ffffffffffffffffffffffffffffffffffffffff1660001b612a1c565b905092915050565b600081836000018054905011612906576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128fd90613ff9565b60405180910390fd5b82600001828154811061291557fe5b9060005260206000200154905092915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081600001805490509050919050565b6000809054906101000a900460ff166129aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129a190614079565b60405180910390fd5b565b60006129b88383612928565b612a11578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050612a16565b600090505b92915050565b60008083600101600084815260200190815260200160002054905060008114612af85760006001820390506000600186600001805490500390506000866000018281548110612a6757fe5b9060005260206000200154905080876000018481548110612a8457fe5b9060005260206000200181905550600183018760010160008381526020019081526020016000208190555086600001805480612abc57fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050612afe565b60009150505b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10612b4557803560ff1916838001178555612b73565b82800160010185558215612b73579182015b82811115612b72578235825591602001919060010190612b57565b5b509050612b809190612c55565b5090565b828054828255906000526020600020908101928215612bfd579160200282015b82811115612bfc5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190612ba4565b5b509050612c0a9190612c7a565b5090565b6040518060c001604052806000801916815260200160008019168152602001606081526020016060815260200160006004811115612c4857fe5b8152602001600081525090565b612c7791905b80821115612c73576000816000905550600101612c5b565b5090565b90565b612cba91905b80821115612cb657600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101612c80565b5090565b90565b600081359050612ccc81614545565b92915050565b600081359050612ce18161455c565b92915050565b60008083601f840112612cf957600080fd5b8235905067ffffffffffffffff811115612d1257600080fd5b602083019150836020820283011115612d2a57600080fd5b9250929050565b60008083601f840112612d4357600080fd5b8235905067ffffffffffffffff811115612d5c57600080fd5b602083019150836020820283011115612d7457600080fd5b9250929050565b600081359050612d8a81614573565b92915050565b600081359050612d9f8161458a565b92915050565b60008083601f840112612db757600080fd5b8235905067ffffffffffffffff811115612dd057600080fd5b602083019150836001820283011115612de857600080fd5b9250929050565b600081359050612dfe816145a1565b92915050565b600081359050612e13816145b8565b92915050565b600081359050612e28816145cf565b92915050565b600081359050612e3d816145e6565b92915050565b600060208284031215612e5557600080fd5b6000612e6384828501612cbd565b91505092915050565b600060208284031215612e7e57600080fd5b6000612e8c84828501612cd2565b91505092915050565b60008060408385031215612ea857600080fd5b6000612eb685828601612cbd565b9250506020612ec785828601612cbd565b9150509250929050565b60008060008060808587031215612ee757600080fd5b6000612ef587828801612cbd565b9450506020612f0687828801612cbd565b9350506040612f1787828801612cbd565b9250506060612f2887828801612def565b91505092959194509250565b600080600060608486031215612f4957600080fd5b6000612f5786828701612cbd565b9350506020612f6886828701612d7b565b9250506040612f7986828701612cbd565b9150509250925092565b600080600080600060a08688031215612f9b57600080fd5b6000612fa988828901612cbd565b9550506020612fba88828901612d7b565b9450506040612fcb88828901612cbd565b9350506060612fdc88828901612d90565b9250506080612fed88828901612d90565b9150509295509295909350565b6000806000806040858703121561301057600080fd5b600085013567ffffffffffffffff81111561302a57600080fd5b61303687828801612ce7565b9450945050602085013567ffffffffffffffff81111561305557600080fd5b61306187828801612d31565b925092505092959194509250565b60006020828403121561308157600080fd5b600061308f84828501612d7b565b91505092915050565b600080604083850312156130ab57600080fd5b60006130b985828601612d7b565b92505060206130ca85828601612cbd565b9150509250929050565b600080604083850312156130e757600080fd5b60006130f585828601612d7b565b925050602061310685828601612def565b9150509250929050565b60006020828403121561312257600080fd5b600061313084828501612def565b91505092915050565b6000806040838503121561314c57600080fd5b600061315a85828601612e04565b925050602061316b85828601612e2e565b9150509250929050565b6000806040838503121561318857600080fd5b600061319685828601612e19565b92505060206131a785828601612d7b565b9150509250929050565b6000806000606084860312156131c657600080fd5b60006131d486828701612e19565b93505060206131e586828701612d7b565b92505060406131f686828701612cbd565b9150509250925092565b60006020828403121561321257600080fd5b600061322084828501612e2e565b91505092915050565b6000806000806060858703121561323f57600080fd5b600061324d87828801612e2e565b945050602061325e87828801612d7b565b935050604085013567ffffffffffffffff81111561327b57600080fd5b61328787828801612da5565b925092505092959194509250565b6000806000606084860312156132aa57600080fd5b60006132b886828701612e2e565b93505060206132c986828701612e04565b92505060406132da86828701612d7b565b9150509250925092565b600080600080608085870312156132fa57600080fd5b600061330887828801612e2e565b945050602061331987828801612e04565b935050604061332a87828801612d7b565b925050606061333b87828801612d7b565b91505092959194509250565b60008060008060006080868803121561335f57600080fd5b600061336d88828901612e2e565b955050602061337e88828901612e04565b945050604086013567ffffffffffffffff81111561339b57600080fd5b6133a788828901612da5565b935093505060606133ba88828901612d7b565b9150509295509295909350565b60006133d383836133ee565b60208301905092915050565b6133e88161446c565b82525050565b6133f781614393565b82525050565b61340681614393565b82525050565b61341d61341882614393565b6144f6565b82525050565b600061342e8261431c565b613438818561434a565b93506134438361430c565b8060005b8381101561347457815161345b88826133c7565b97506134668361433d565b925050600181019050613447565b5085935050505092915050565b61348a816143b7565b82525050565b613499816143c3565b82525050565b6134a8816143c3565b82525050565b6134b7816143cd565b82525050565b60006134c9838561435b565b93506134d68385846144b4565b6134df8361451a565b840190509392505050565b60006134f6838561436c565b93506135038385846144b4565b82840190509392505050565b600061351a82614327565b613524818561435b565b93506135348185602086016144c3565b61353d8161451a565b840191505092915050565b6135518161447e565b82525050565b6135608161447e565b82525050565b600061357182614332565b61357b8185614377565b935061358b8185602086016144c3565b6135948161451a565b840191505092915050565b60006135ac601c83614377565b91507f70726f706f73616c20616c7265616479207472616e73666572726564000000006000830152602082019050919050565b60006135ec601683614377565b91507f70726f706f73616c206973206e6f7420616374697665000000000000000000006000830152602082019050919050565b600061362c602283614377565b91507f456e756d657261626c655365743a20696e646578206f7574206f6620626f756e60008301527f64730000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000613692601e83614377565b91507f73656e646572206973206e6f742072656c61796572206f722061646d696e00006000830152602082019050919050565b60006136d2602f83614377565b91507f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008301527f2061646d696e20746f206772616e7400000000000000000000000000000000006020830152604082019050919050565b6000613738601583614377565b91507f72656c6179657220616c726561647920766f74656400000000000000000000006000830152602082019050919050565b6000613778601483614377565b91507f5061757361626c653a206e6f74207061757365640000000000000000000000006000830152602082019050919050565b60006137b8601f83614377565b91507f6164647220646f65736e277420686176652072656c6179657220726f6c6521006000830152602082019050919050565b60006137f8601683614377565b91507f496e636f72726563742066656520737570706c696564000000000000000000006000830152602082019050919050565b6000613838601b83614377565b91507f6461746120646f65736e2774206d6174636820646174616861736800000000006000830152602082019050919050565b6000613878601a83614377565b91507f50726f706f73616c20616c72656164792063616e63656c6c65640000000000006000830152602082019050919050565b60006138b8602083614377565b91507f7265736f757263654944206e6f74206d617070656420746f2068616e646c65726000830152602082019050919050565b60006138f8603083614377565b91507f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008301527f2061646d696e20746f207265766f6b65000000000000000000000000000000006020830152604082019050919050565b600061395e602083614377565b91507f50726f706f73616c206e6f7420617420657870697279207468726573686f6c646000830152602082019050919050565b600061399e601083614377565b91507f5061757361626c653a20706175736564000000000000000000000000000000006000830152602082019050919050565b60006139de601e83614377565b91507f6164647220616c7265616479206861732072656c6179657220726f6c652100006000830152602082019050919050565b6000613a1e602a83614377565b91507f70726f706f73616c20616c7265616479207061737365642f657865637574656460008301527f2f63616e63656c6c6564000000000000000000000000000000000000000000006020830152604082019050919050565b6000613a84602083614377565b91507f73656e64657220646f65736e277420686176652072656c6179657220726f6c656000830152602082019050919050565b6000613ac4601e83614377565b91507f73656e64657220646f65736e277420686176652061646d696e20726f6c6500006000830152602082019050919050565b6000613b04601f83614377565b91507f43757272656e742066656520697320657175616c20746f206e657720666565006000830152602082019050919050565b6000613b44600c83614388565b91507f52454c415945525f524f4c4500000000000000000000000000000000000000006000830152600c82019050919050565b6000613b84601183614377565b91507f6461746168617368206d69736d617463680000000000000000000000000000006000830152602082019050919050565b6000613bc4601983614377565b91507f6e6f2068616e646c657220666f72207265736f757263654944000000000000006000830152602082019050919050565b6000613c04602f83614377565b91507f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008301527f20726f6c657320666f722073656c6600000000000000000000000000000000006020830152604082019050919050565b600060c083016000830151613c756000860182613490565b506020830151613c886020860182613490565b5060408301518482036040860152613ca08282613423565b91505060608301518482036060860152613cba8282613423565b9150506080830151613ccf6080860182613548565b5060a0830151613ce260a0860182613ced565b508091505092915050565b613cf68161442c565b82525050565b613d058161442c565b82525050565b613d1481614436565b82525050565b613d238161445f565b82525050565b6000613d35828661340c565b601482019150613d468284866134ea565b9150819050949350505050565b6000613d5e82613b37565b9150819050919050565b6000602082019050613d7d60008301846133fd565b92915050565b6000602082019050613d9860008301846133df565b92915050565b6000606082019050613db360008301866133fd565b613dc060208301856133fd565b613dcd6040830184613cfc565b949350505050565b6000602082019050613dea6000830184613481565b92915050565b6000602082019050613e05600083018461349f565b92915050565b6000604082019050613e20600083018561349f565b613e2d60208301846133fd565b9392505050565b6000608082019050613e49600083018761349f565b613e5660208301866133fd565b613e6360408301856134ae565b613e7060608301846134ae565b95945050505050565b6000604082019050613e8e600083018561349f565b613e9b602083018461349f565b9392505050565b6000608082019050613eb7600083018761349f565b613ec4602083018661349f565b613ed16040830185613557565b613ede6060830184613cfc565b95945050505050565b6000604082019050613efc600083018661349f565b8181036020830152613f0f8184866134bd565b9050949350505050565b600060a082019050613f2e600083018961349f565b613f3b6020830188613d1a565b613f486040830187613d0b565b613f5560608301866133df565b8181036080830152613f688184866134bd565b9050979650505050505050565b60006020820190508181036000830152613f8f818461350f565b905092915050565b60006020820190508181036000830152613fb18184613566565b905092915050565b60006020820190508181036000830152613fd28161359f565b9050919050565b60006020820190508181036000830152613ff2816135df565b9050919050565b600060208201905081810360008301526140128161361f565b9050919050565b6000602082019050818103600083015261403281613685565b9050919050565b60006020820190508181036000830152614052816136c5565b9050919050565b600060208201905081810360008301526140728161372b565b9050919050565b600060208201905081810360008301526140928161376b565b9050919050565b600060208201905081810360008301526140b2816137ab565b9050919050565b600060208201905081810360008301526140d2816137eb565b9050919050565b600060208201905081810360008301526140f28161382b565b9050919050565b600060208201905081810360008301526141128161386b565b9050919050565b60006020820190508181036000830152614132816138ab565b9050919050565b60006020820190508181036000830152614152816138eb565b9050919050565b6000602082019050818103600083015261417281613951565b9050919050565b6000602082019050818103600083015261419281613991565b9050919050565b600060208201905081810360008301526141b2816139d1565b9050919050565b600060208201905081810360008301526141d281613a11565b9050919050565b600060208201905081810360008301526141f281613a77565b9050919050565b6000602082019050818103600083015261421281613ab7565b9050919050565b6000602082019050818103600083015261423281613af7565b9050919050565b6000602082019050818103600083015261425281613b77565b9050919050565b6000602082019050818103600083015261427281613bb7565b9050919050565b6000602082019050818103600083015261429281613bf7565b9050919050565b600060208201905081810360008301526142b38184613c5d565b905092915050565b60006020820190506142d06000830184613cfc565b92915050565b60006020820190506142eb6000830184613d0b565b92915050565b60006020820190506143066000830184613d1a565b92915050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b600061439e8261440c565b9050919050565b60006143b08261440c565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600081905061440782614538565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600068ffffffffffffffffff82169050919050565b600060ff82169050919050565b600061447782614490565b9050919050565b6000614489826143f9565b9050919050565b600061449b826144a2565b9050919050565b60006144ad8261440c565b9050919050565b82818337600083830152505050565b60005b838110156144e15780820151818401526020810190506144c6565b838111156144f0576000848401525b50505050565b600061450182614508565b9050919050565b60006145138261452b565b9050919050565b6000601f19601f8301169050919050565b60008160601b9050919050565b6005811061454257fe5b50565b61454e81614393565b811461455957600080fd5b50565b614565816143a5565b811461457057600080fd5b50565b61457c816143c3565b811461458757600080fd5b50565b614593816143cd565b811461459e57600080fd5b50565b6145aa8161442c565b81146145b557600080fd5b50565b6145c181614436565b81146145cc57600080fd5b50565b6145d88161444a565b81146145e357600080fd5b50565b6145ef8161445f565b81146145fa57600080fd5b5056fea2646970667358221220a0d717a9570e262fc40ab8ab73141bec66b06cc29fdcb86332c1e2d5f8bc2fba64736f6c63430006040033"

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, chainID uint64, initialRelayers []common.Address, initialRelayerThreshold *big.Int, fee *big.Int, expiry *big.Int) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBin), backend, chainID, initialRelayers, initialRelayerThreshold, fee, expiry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) RELAYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "RELAYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) RELAYERROLE() ([32]byte, error) {
	return _Bridge.Contract.RELAYERROLE(&_Bridge.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) RELAYERROLE() ([32]byte, error) {
	return _Bridge.Contract.RELAYERROLE(&_Bridge.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _domainID() view returns(uint64)
func (_Bridge *BridgeCaller) ChainID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_chainID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _domainID() view returns(uint64)
func (_Bridge *BridgeSession) ChainID() (uint64, error) {
	return _Bridge.Contract.ChainID(&_Bridge.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _domainID() view returns(uint64)
func (_Bridge *BridgeCallerSession) ChainID() (uint64, error) {
	return _Bridge.Contract.ChainID(&_Bridge.CallOpts)
}

// DepositCounts is a free data retrieval call binding the contract method 0x4b0b919d.
//
// Solidity: function _depositCounts(uint8 ) view returns(uint64)
func (_Bridge *BridgeCaller) DepositCounts(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_depositCounts", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DepositCounts is a free data retrieval call binding the contract method 0x4b0b919d.
//
// Solidity: function _depositCounts(uint8 ) view returns(uint64)
func (_Bridge *BridgeSession) DepositCounts(arg0 uint8) (uint64, error) {
	return _Bridge.Contract.DepositCounts(&_Bridge.CallOpts, arg0)
}

// DepositCounts is a free data retrieval call binding the contract method 0x4b0b919d.
//
// Solidity: function _depositCounts(uint8 ) view returns(uint64)
func (_Bridge *BridgeCallerSession) DepositCounts(arg0 uint8) (uint64, error) {
	return _Bridge.Contract.DepositCounts(&_Bridge.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0x3ee7094a.
//
// Solidity: function _depositRecords(uint64 , uint8 ) view returns(bytes)
func (_Bridge *BridgeCaller) DepositRecords(opts *bind.CallOpts, arg0 uint64, arg1 uint8) ([]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DepositRecords is a free data retrieval call binding the contract method 0x3ee7094a.
//
// Solidity: function _depositRecords(uint64 , uint8 ) view returns(bytes)
func (_Bridge *BridgeSession) DepositRecords(arg0 uint64, arg1 uint8) ([]byte, error) {
	return _Bridge.Contract.DepositRecords(&_Bridge.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0x3ee7094a.
//
// Solidity: function _depositRecords(uint64 , uint8 ) view returns(bytes)
func (_Bridge *BridgeCallerSession) DepositRecords(arg0 uint64, arg1 uint8) ([]byte, error) {
	return _Bridge.Contract.DepositRecords(&_Bridge.CallOpts, arg0, arg1)
}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Bridge *BridgeCaller) Expiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_expiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Bridge *BridgeSession) Expiry() (*big.Int, error) {
	return _Bridge.Contract.Expiry(&_Bridge.CallOpts)
}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Bridge *BridgeCallerSession) Expiry() (*big.Int, error) {
	return _Bridge.Contract.Expiry(&_Bridge.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Bridge *BridgeCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Bridge *BridgeSession) Fee() (*big.Int, error) {
	return _Bridge.Contract.Fee(&_Bridge.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Bridge *BridgeCallerSession) Fee() (*big.Int, error) {
	return _Bridge.Contract.Fee(&_Bridge.CallOpts)
}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x7febe63f.
//
// Solidity: function _hasVotedOnProposal(uint72 , bytes32 , address ) view returns(bool)
func (_Bridge *BridgeCaller) HasVotedOnProposal(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_hasVotedOnProposal", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x7febe63f.
//
// Solidity: function _hasVotedOnProposal(uint72 , bytes32 , address ) view returns(bool)
func (_Bridge *BridgeSession) HasVotedOnProposal(arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (bool, error) {
	return _Bridge.Contract.HasVotedOnProposal(&_Bridge.CallOpts, arg0, arg1, arg2)
}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x7febe63f.
//
// Solidity: function _hasVotedOnProposal(uint72 , bytes32 , address ) view returns(bool)
func (_Bridge *BridgeCallerSession) HasVotedOnProposal(arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (bool, error) {
	return _Bridge.Contract.HasVotedOnProposal(&_Bridge.CallOpts, arg0, arg1, arg2)
}

// Proposals is a free data retrieval call binding the contract method 0x50598719.
//
// Solidity: function _proposals(uint72 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock)
func (_Bridge *BridgeCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_proposals", arg0, arg1)

	outstruct := new(struct {
		ResourceID    [32]byte
		DataHash      [32]byte
		Status        uint8
		ProposedBlock *big.Int
	})

	outstruct.ResourceID = out[0].([32]byte)
	outstruct.DataHash = out[1].([32]byte)
	outstruct.Status = out[2].(uint8)
	outstruct.ProposedBlock = out[3].(*big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x50598719.
//
// Solidity: function _proposals(uint72 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock)
func (_Bridge *BridgeSession) Proposals(arg0 *big.Int, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
}, error) {
	return _Bridge.Contract.Proposals(&_Bridge.CallOpts, arg0, arg1)
}

// Proposals is a free data retrieval call binding the contract method 0x50598719.
//
// Solidity: function _proposals(uint72 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock)
func (_Bridge *BridgeCallerSession) Proposals(arg0 *big.Int, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
}, error) {
	return _Bridge.Contract.Proposals(&_Bridge.CallOpts, arg0, arg1)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Bridge *BridgeCaller) RelayerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_relayerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Bridge *BridgeSession) RelayerThreshold() (*big.Int, error) {
	return _Bridge.Contract.RelayerThreshold(&_Bridge.CallOpts)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Bridge *BridgeCallerSession) RelayerThreshold() (*big.Int, error) {
	return _Bridge.Contract.RelayerThreshold(&_Bridge.CallOpts)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Bridge *BridgeCaller) ResourceIDToHandlerAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_resourceIDToHandlerAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Bridge *BridgeSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _Bridge.Contract.ResourceIDToHandlerAddress(&_Bridge.CallOpts, arg0)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Bridge *BridgeCallerSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _Bridge.Contract.ResourceIDToHandlerAddress(&_Bridge.CallOpts, arg0)
}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Bridge *BridgeCaller) TotalProposals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_totalProposals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Bridge *BridgeSession) TotalProposals() (*big.Int, error) {
	return _Bridge.Contract.TotalProposals(&_Bridge.CallOpts)
}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Bridge *BridgeCallerSession) TotalProposals() (*big.Int, error) {
	return _Bridge.Contract.TotalProposals(&_Bridge.CallOpts)
}

// TotalRelayers is a free data retrieval call binding the contract method 0x802aabe8.
//
// Solidity: function _totalRelayers() view returns(uint256)
func (_Bridge *BridgeCaller) TotalRelayers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_totalRelayers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRelayers is a free data retrieval call binding the contract method 0x802aabe8.
//
// Solidity: function _totalRelayers() view returns(uint256)
func (_Bridge *BridgeSession) TotalRelayers() (*big.Int, error) {
	return _Bridge.Contract.TotalRelayers(&_Bridge.CallOpts)
}

// TotalRelayers is a free data retrieval call binding the contract method 0x802aabe8.
//
// Solidity: function _totalRelayers() view returns(uint256)
func (_Bridge *BridgeCallerSession) TotalRelayers() (*big.Int, error) {
	return _Bridge.Contract.TotalRelayers(&_Bridge.CallOpts)
}

// GetProposal is a free data retrieval call binding the contract method 0xa9cf69fa.
//
// Solidity: function getProposal(uint64 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Bridge *BridgeCaller) GetProposal(opts *bind.CallOpts, originChainID uint64, depositNonce uint64, dataHash [32]byte) (BridgeProposal, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getProposal", originChainID, depositNonce, dataHash)

	if err != nil {
		return *new(BridgeProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeProposal)).(*BridgeProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xa9cf69fa.
//
// Solidity: function getProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Bridge *BridgeSession) GetProposal(originChainID uint64, depositNonce uint64, dataHash [32]byte) (BridgeProposal, error) {
	return _Bridge.Contract.GetProposal(&_Bridge.CallOpts, originChainID, depositNonce, dataHash)
}

// GetProposal is a free data retrieval call binding the contract method 0xa9cf69fa.
//
// Solidity: function getProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Bridge *BridgeCallerSession) GetProposal(originChainID uint64, depositNonce uint64, dataHash [32]byte) (BridgeProposal, error) {
	return _Bridge.Contract.GetProposal(&_Bridge.CallOpts, originChainID, depositNonce, dataHash)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bridge *BridgeCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bridge *BridgeSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bridge.Contract.GetRoleMember(&_Bridge.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bridge *BridgeCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bridge.Contract.GetRoleMember(&_Bridge.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bridge *BridgeCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bridge *BridgeSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bridge.Contract.GetRoleMemberCount(&_Bridge.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bridge *BridgeCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bridge.Contract.GetRoleMemberCount(&_Bridge.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayer) view returns(bool)
func (_Bridge *BridgeCaller) IsRelayer(opts *bind.CallOpts, relayer common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isRelayer", relayer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayer) view returns(bool)
func (_Bridge *BridgeSession) IsRelayer(relayer common.Address) (bool, error) {
	return _Bridge.Contract.IsRelayer(&_Bridge.CallOpts, relayer)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayer) view returns(bool)
func (_Bridge *BridgeCallerSession) IsRelayer(relayer common.Address) (bool, error) {
	return _Bridge.Contract.IsRelayer(&_Bridge.CallOpts, relayer)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCallerSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// AdminAddRelayer is a paid mutator transaction binding the contract method 0xcdb0f73a.
//
// Solidity: function adminAddRelayer(address relayerAddress) returns()
func (_Bridge *BridgeTransactor) AdminAddRelayer(opts *bind.TransactOpts, relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminAddRelayer", relayerAddress)
}

// AdminAddRelayer is a paid mutator transaction binding the contract method 0xcdb0f73a.
//
// Solidity: function adminAddRelayer(address relayerAddress) returns()
func (_Bridge *BridgeSession) AdminAddRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminAddRelayer(&_Bridge.TransactOpts, relayerAddress)
}

// AdminAddRelayer is a paid mutator transaction binding the contract method 0xcdb0f73a.
//
// Solidity: function adminAddRelayer(address relayerAddress) returns()
func (_Bridge *BridgeTransactorSession) AdminAddRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminAddRelayer(&_Bridge.TransactOpts, relayerAddress)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_Bridge *BridgeTransactor) AdminChangeFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminChangeFee", newFee)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_Bridge *BridgeSession) AdminChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminChangeFee(&_Bridge.TransactOpts, newFee)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_Bridge *BridgeTransactorSession) AdminChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminChangeFee(&_Bridge.TransactOpts, newFee)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_Bridge *BridgeTransactor) AdminChangeRelayerThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminChangeRelayerThreshold", newThreshold)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_Bridge *BridgeSession) AdminChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminChangeRelayerThreshold(&_Bridge.TransactOpts, newThreshold)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_Bridge *BridgeTransactorSession) AdminChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminChangeRelayerThreshold(&_Bridge.TransactOpts, newThreshold)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Bridge *BridgeTransactor) AdminPauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminPauseTransfers")
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Bridge *BridgeSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminPauseTransfers(&_Bridge.TransactOpts)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Bridge *BridgeTransactorSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminPauseTransfers(&_Bridge.TransactOpts)
}

// AdminRemoveRelayer is a paid mutator transaction binding the contract method 0x9d82dd63.
//
// Solidity: function adminRemoveRelayer(address relayerAddress) returns()
func (_Bridge *BridgeTransactor) AdminRemoveRelayer(opts *bind.TransactOpts, relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminRemoveRelayer", relayerAddress)
}

// AdminRemoveRelayer is a paid mutator transaction binding the contract method 0x9d82dd63.
//
// Solidity: function adminRemoveRelayer(address relayerAddress) returns()
func (_Bridge *BridgeSession) AdminRemoveRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminRemoveRelayer(&_Bridge.TransactOpts, relayerAddress)
}

// AdminRemoveRelayer is a paid mutator transaction binding the contract method 0x9d82dd63.
//
// Solidity: function adminRemoveRelayer(address relayerAddress) returns()
func (_Bridge *BridgeTransactorSession) AdminRemoveRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminRemoveRelayer(&_Bridge.TransactOpts, relayerAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_Bridge *BridgeTransactor) AdminSetBurnable(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetBurnable", handlerAddress, tokenAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_Bridge *BridgeSession) AdminSetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetBurnable(&_Bridge.TransactOpts, handlerAddress, tokenAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_Bridge *BridgeTransactorSession) AdminSetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetBurnable(&_Bridge.TransactOpts, handlerAddress, tokenAddress)
}

// AdminSetGenericResource is a paid mutator transaction binding the contract method 0xe8437ee7.
//
// Solidity: function adminSetGenericResource(address handlerAddress, bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_Bridge *BridgeTransactor) AdminSetGenericResource(opts *bind.TransactOpts, handlerAddress common.Address, resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetGenericResource", handlerAddress, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// AdminSetGenericResource is a paid mutator transaction binding the contract method 0xe8437ee7.
//
// Solidity: function adminSetGenericResource(address handlerAddress, bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_Bridge *BridgeSession) AdminSetGenericResource(handlerAddress common.Address, resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetGenericResource(&_Bridge.TransactOpts, handlerAddress, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// AdminSetGenericResource is a paid mutator transaction binding the contract method 0xe8437ee7.
//
// Solidity: function adminSetGenericResource(address handlerAddress, bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_Bridge *BridgeTransactorSession) AdminSetGenericResource(handlerAddress common.Address, resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetGenericResource(&_Bridge.TransactOpts, handlerAddress, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Bridge *BridgeTransactor) AdminSetResource(opts *bind.TransactOpts, handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetResource", handlerAddress, resourceID, tokenAddress)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Bridge *BridgeSession) AdminSetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetResource(&_Bridge.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Bridge *BridgeTransactorSession) AdminSetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetResource(&_Bridge.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Bridge *BridgeTransactor) AdminUnpauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminUnpauseTransfers")
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Bridge *BridgeSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminUnpauseTransfers(&_Bridge.TransactOpts)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Bridge *BridgeTransactorSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminUnpauseTransfers(&_Bridge.TransactOpts)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Bridge *BridgeTransactor) AdminWithdraw(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminWithdraw", handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Bridge *BridgeSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminWithdraw(&_Bridge.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Bridge *BridgeTransactorSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminWithdraw(&_Bridge.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x17f03ce5.
//
// Solidity: function cancelProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactor) CancelProposal(opts *bind.TransactOpts, chainID uint64, depositNonce uint64, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "cancelProposal", chainID, depositNonce, dataHash)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x17f03ce5.
//
// Solidity: function cancelProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash) returns()
func (_Bridge *BridgeSession) CancelProposal(chainID uint64, depositNonce uint64, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.CancelProposal(&_Bridge.TransactOpts, chainID, depositNonce, dataHash)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x17f03ce5.
//
// Solidity: function cancelProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactorSession) CancelProposal(chainID uint64, depositNonce uint64, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.CancelProposal(&_Bridge.TransactOpts, chainID, depositNonce, dataHash)
}

// Deposit is a paid mutator transaction binding the contract method 0x05e2ca17.
//
// Solidity: function deposit(uint8 destinationChainID, bytes32 resourceID, bytes data) payable returns()
func (_Bridge *BridgeTransactor) Deposit(opts *bind.TransactOpts, destinationChainID uint64, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deposit", destinationChainID, resourceID, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x05e2ca17.
//
// Solidity: function deposit(uint8 destinationChainID, bytes32 resourceID, bytes data) payable returns()
func (_Bridge *BridgeSession) Deposit(destinationChainID uint64, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, destinationChainID, resourceID, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x05e2ca17.
//
// Solidity: function deposit(uint8 destinationChainID, bytes32 resourceID, bytes data) payable returns()
func (_Bridge *BridgeTransactorSession) Deposit(destinationChainID uint64, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, destinationChainID, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x4454b20d.
//
// Solidity: function executeProposal(uint8 chainID, uint64 depositNonce, bytes data, bytes32 resourceID) returns()
func (_Bridge *BridgeTransactor) ExecuteProposal(opts *bind.TransactOpts, chainID uint64, depositNonce uint64, data []byte, resourceID [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "executeProposal", chainID, depositNonce, data, resourceID)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x4454b20d.
//
// Solidity: function executeProposal(uint8 chainID, uint64 depositNonce, bytes data, bytes32 resourceID) returns()
func (_Bridge *BridgeSession) ExecuteProposal(chainID uint64, depositNonce uint64, data []byte, resourceID [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteProposal(&_Bridge.TransactOpts, chainID, depositNonce, data, resourceID)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x4454b20d.
//
// Solidity: function executeProposal(uint8 chainID, uint64 depositNonce, bytes data, bytes32 resourceID) returns()
func (_Bridge *BridgeTransactorSession) ExecuteProposal(chainID uint64, depositNonce uint64, data []byte, resourceID [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteProposal(&_Bridge.TransactOpts, chainID, depositNonce, data, resourceID)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_Bridge *BridgeTransactor) RenounceAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceAdmin", newAdmin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_Bridge *BridgeSession) RenounceAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceAdmin(&_Bridge.TransactOpts, newAdmin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_Bridge *BridgeTransactorSession) RenounceAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceAdmin(&_Bridge.TransactOpts, newAdmin)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// TransferFunds is a paid mutator transaction binding the contract method 0x4603ae38.
//
// Solidity: function transferFunds(address[] addrs, uint256[] amounts) returns()
func (_Bridge *BridgeTransactor) TransferFunds(opts *bind.TransactOpts, addrs []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferFunds", addrs, amounts)
}

// TransferFunds is a paid mutator transaction binding the contract method 0x4603ae38.
//
// Solidity: function transferFunds(address[] addrs, uint256[] amounts) returns()
func (_Bridge *BridgeSession) TransferFunds(addrs []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TransferFunds(&_Bridge.TransactOpts, addrs, amounts)
}

// TransferFunds is a paid mutator transaction binding the contract method 0x4603ae38.
//
// Solidity: function transferFunds(address[] addrs, uint256[] amounts) returns()
func (_Bridge *BridgeTransactorSession) TransferFunds(addrs []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TransferFunds(&_Bridge.TransactOpts, addrs, amounts)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x1ff013f1.
//
// Solidity: function voteProposal(uint8 chainID, uint64 depositNonce, bytes32 resourceID, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactor) VoteProposal(opts *bind.TransactOpts, chainID uint64, depositNonce uint64, resourceID [32]byte, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "voteProposal", chainID, depositNonce, resourceID, dataHash)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x1ff013f1.
//
// Solidity: function voteProposal(uint8 chainID, uint64 depositNonce, bytes32 resourceID, bytes32 dataHash) returns()
func (_Bridge *BridgeSession) VoteProposal(chainID uint64, depositNonce uint64, resourceID [32]byte, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.VoteProposal(&_Bridge.TransactOpts, chainID, depositNonce, resourceID, dataHash)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x1ff013f1.
//
// Solidity: function voteProposal(uint8 chainID, uint64 depositNonce, bytes32 resourceID, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactorSession) VoteProposal(chainID uint64, depositNonce uint64, resourceID [32]byte, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.VoteProposal(&_Bridge.TransactOpts, chainID, depositNonce, resourceID, dataHash)
}

// BridgeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Bridge contract.
type BridgeDepositIterator struct {
	Event *BridgeDeposit // Event containing the contract specifics and raw log

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
func (it *BridgeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDeposit)
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
		it.Event = new(BridgeDeposit)
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
func (it *BridgeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDeposit represents a Deposit event raised by the Bridge contract.
type BridgeDeposit struct {
	DestinationChainID uint64
	ResourceID         [32]byte
	DepositNonce       uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdbb69440df8433824a026ef190652f29929eb64b4d1d5d2a69be8afe3e6eaed8.
//
// Solidity: event Deposit(uint8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce)
func (_Bridge *BridgeFilterer) FilterDeposit(opts *bind.FilterOpts, destinationChainID []uint8, resourceID [][32]byte, depositNonce []uint64) (*BridgeDepositIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var resourceIDRule []interface{}
	for _, resourceIDItem := range resourceID {
		resourceIDRule = append(resourceIDRule, resourceIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Deposit", destinationChainIDRule, resourceIDRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return &BridgeDepositIterator{contract: _Bridge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdbb69440df8433824a026ef190652f29929eb64b4d1d5d2a69be8afe3e6eaed8.
//
// Solidity: event Deposit(uint8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce)
func (_Bridge *BridgeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BridgeDeposit, destinationChainID []uint8, resourceID [][32]byte, depositNonce []uint64) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var resourceIDRule []interface{}
	for _, resourceIDItem := range resourceID {
		resourceIDRule = append(resourceIDRule, resourceIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Deposit", destinationChainIDRule, resourceIDRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDeposit)
				if err := _Bridge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdbb69440df8433824a026ef190652f29929eb64b4d1d5d2a69be8afe3e6eaed8.
//
// Solidity: event Deposit(uint8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce)
func (_Bridge *BridgeFilterer) ParseDeposit(log types.Log) (*BridgeDeposit, error) {
	event := new(BridgeDeposit)
	if err := _Bridge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bridge contract.
type BridgePausedIterator struct {
	Event *BridgePaused // Event containing the contract specifics and raw log

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
func (it *BridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePaused)
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
		it.Event = new(BridgePaused)
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
func (it *BridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePaused represents a Paused event raised by the Bridge contract.
type BridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgePausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgePausedIterator{contract: _Bridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgePaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePaused)
				if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) ParsePaused(log types.Log) (*BridgePaused, error) {
	event := new(BridgePaused)
	if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeProposalEventIterator is returned from FilterProposalEvent and is used to iterate over the raw logs and unpacked data for ProposalEvent events raised by the Bridge contract.
type BridgeProposalEventIterator struct {
	Event *BridgeProposalEvent // Event containing the contract specifics and raw log

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
func (it *BridgeProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeProposalEvent)
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
		it.Event = new(BridgeProposalEvent)
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
func (it *BridgeProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeProposalEvent represents a ProposalEvent event raised by the Bridge contract.
type BridgeProposalEvent struct {
	OriginChainID uint64
	DepositNonce  uint64
	Status        uint8
	ResourceID    [32]byte
	DataHash      [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProposalEvent is a free log retrieval operation binding the contract event 0x803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab650417.
//
// Solidity: event ProposalEvent(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID, bytes32 dataHash)
func (_Bridge *BridgeFilterer) FilterProposalEvent(opts *bind.FilterOpts, originChainID []uint8, depositNonce []uint64, status []uint8) (*BridgeProposalEventIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ProposalEvent", originChainIDRule, depositNonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &BridgeProposalEventIterator{contract: _Bridge.contract, event: "ProposalEvent", logs: logs, sub: sub}, nil
}

// WatchProposalEvent is a free log subscription operation binding the contract event 0x803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab650417.
//
// Solidity: event ProposalEvent(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID, bytes32 dataHash)
func (_Bridge *BridgeFilterer) WatchProposalEvent(opts *bind.WatchOpts, sink chan<- *BridgeProposalEvent, originChainID []uint8, depositNonce []uint64, status []uint8) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ProposalEvent", originChainIDRule, depositNonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeProposalEvent)
				if err := _Bridge.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
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

// ParseProposalEvent is a log parse operation binding the contract event 0x803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab650417.
//
// Solidity: event ProposalEvent(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID, bytes32 dataHash)
func (_Bridge *BridgeFilterer) ParseProposalEvent(log types.Log) (*BridgeProposalEvent, error) {
	event := new(BridgeProposalEvent)
	if err := _Bridge.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeProposalVoteIterator is returned from FilterProposalVote and is used to iterate over the raw logs and unpacked data for ProposalVote events raised by the Bridge contract.
type BridgeProposalVoteIterator struct {
	Event *BridgeProposalVote // Event containing the contract specifics and raw log

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
func (it *BridgeProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeProposalVote)
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
		it.Event = new(BridgeProposalVote)
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
func (it *BridgeProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeProposalVote represents a ProposalVote event raised by the Bridge contract.
type BridgeProposalVote struct {
	OriginChainID uint64
	DepositNonce  uint64
	Status        uint8
	ResourceID    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProposalVote is a free log retrieval operation binding the contract event 0x25f8daaa4635a7729927ba3f5b3d59cc3320aca7c32c9db4e7ca7b9574343640.
//
// Solidity: event ProposalVote(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID)
func (_Bridge *BridgeFilterer) FilterProposalVote(opts *bind.FilterOpts, originChainID []uint8, depositNonce []uint64, status []uint8) (*BridgeProposalVoteIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ProposalVote", originChainIDRule, depositNonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &BridgeProposalVoteIterator{contract: _Bridge.contract, event: "ProposalVote", logs: logs, sub: sub}, nil
}

// WatchProposalVote is a free log subscription operation binding the contract event 0x25f8daaa4635a7729927ba3f5b3d59cc3320aca7c32c9db4e7ca7b9574343640.
//
// Solidity: event ProposalVote(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID)
func (_Bridge *BridgeFilterer) WatchProposalVote(opts *bind.WatchOpts, sink chan<- *BridgeProposalVote, originChainID []uint8, depositNonce []uint64, status []uint8) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ProposalVote", originChainIDRule, depositNonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeProposalVote)
				if err := _Bridge.contract.UnpackLog(event, "ProposalVote", log); err != nil {
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

// ParseProposalVote is a log parse operation binding the contract event 0x25f8daaa4635a7729927ba3f5b3d59cc3320aca7c32c9db4e7ca7b9574343640.
//
// Solidity: event ProposalVote(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID)
func (_Bridge *BridgeFilterer) ParseProposalVote(log types.Log) (*BridgeProposalVote, error) {
	event := new(BridgeProposalVote)
	if err := _Bridge.contract.UnpackLog(event, "ProposalVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRelayerAddedIterator is returned from FilterRelayerAdded and is used to iterate over the raw logs and unpacked data for RelayerAdded events raised by the Bridge contract.
type BridgeRelayerAddedIterator struct {
	Event *BridgeRelayerAdded // Event containing the contract specifics and raw log

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
func (it *BridgeRelayerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRelayerAdded)
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
		it.Event = new(BridgeRelayerAdded)
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
func (it *BridgeRelayerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRelayerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRelayerAdded represents a RelayerAdded event raised by the Bridge contract.
type BridgeRelayerAdded struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerAdded is a free log retrieval operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address indexed relayer)
func (_Bridge *BridgeFilterer) FilterRelayerAdded(opts *bind.FilterOpts, relayer []common.Address) (*BridgeRelayerAddedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RelayerAdded", relayerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRelayerAddedIterator{contract: _Bridge.contract, event: "RelayerAdded", logs: logs, sub: sub}, nil
}

// WatchRelayerAdded is a free log subscription operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address indexed relayer)
func (_Bridge *BridgeFilterer) WatchRelayerAdded(opts *bind.WatchOpts, sink chan<- *BridgeRelayerAdded, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RelayerAdded", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRelayerAdded)
				if err := _Bridge.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
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

// ParseRelayerAdded is a log parse operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address indexed relayer)
func (_Bridge *BridgeFilterer) ParseRelayerAdded(log types.Log) (*BridgeRelayerAdded, error) {
	event := new(BridgeRelayerAdded)
	if err := _Bridge.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRelayerRemovedIterator is returned from FilterRelayerRemoved and is used to iterate over the raw logs and unpacked data for RelayerRemoved events raised by the Bridge contract.
type BridgeRelayerRemovedIterator struct {
	Event *BridgeRelayerRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeRelayerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRelayerRemoved)
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
		it.Event = new(BridgeRelayerRemoved)
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
func (it *BridgeRelayerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRelayerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRelayerRemoved represents a RelayerRemoved event raised by the Bridge contract.
type BridgeRelayerRemoved struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerRemoved is a free log retrieval operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address indexed relayer)
func (_Bridge *BridgeFilterer) FilterRelayerRemoved(opts *bind.FilterOpts, relayer []common.Address) (*BridgeRelayerRemovedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RelayerRemoved", relayerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRelayerRemovedIterator{contract: _Bridge.contract, event: "RelayerRemoved", logs: logs, sub: sub}, nil
}

// WatchRelayerRemoved is a free log subscription operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address indexed relayer)
func (_Bridge *BridgeFilterer) WatchRelayerRemoved(opts *bind.WatchOpts, sink chan<- *BridgeRelayerRemoved, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RelayerRemoved", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRelayerRemoved)
				if err := _Bridge.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
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

// ParseRelayerRemoved is a log parse operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address indexed relayer)
func (_Bridge *BridgeFilterer) ParseRelayerRemoved(log types.Log) (*BridgeRelayerRemoved, error) {
	event := new(BridgeRelayerRemoved)
	if err := _Bridge.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRelayerThresholdChangedIterator is returned from FilterRelayerThresholdChanged and is used to iterate over the raw logs and unpacked data for RelayerThresholdChanged events raised by the Bridge contract.
type BridgeRelayerThresholdChangedIterator struct {
	Event *BridgeRelayerThresholdChanged // Event containing the contract specifics and raw log

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
func (it *BridgeRelayerThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRelayerThresholdChanged)
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
		it.Event = new(BridgeRelayerThresholdChanged)
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
func (it *BridgeRelayerThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRelayerThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRelayerThresholdChanged represents a RelayerThresholdChanged event raised by the Bridge contract.
type BridgeRelayerThresholdChanged struct {
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRelayerThresholdChanged is a free log retrieval operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 indexed newThreshold)
func (_Bridge *BridgeFilterer) FilterRelayerThresholdChanged(opts *bind.FilterOpts, newThreshold []*big.Int) (*BridgeRelayerThresholdChangedIterator, error) {

	var newThresholdRule []interface{}
	for _, newThresholdItem := range newThreshold {
		newThresholdRule = append(newThresholdRule, newThresholdItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RelayerThresholdChanged", newThresholdRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRelayerThresholdChangedIterator{contract: _Bridge.contract, event: "RelayerThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchRelayerThresholdChanged is a free log subscription operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 indexed newThreshold)
func (_Bridge *BridgeFilterer) WatchRelayerThresholdChanged(opts *bind.WatchOpts, sink chan<- *BridgeRelayerThresholdChanged, newThreshold []*big.Int) (event.Subscription, error) {

	var newThresholdRule []interface{}
	for _, newThresholdItem := range newThreshold {
		newThresholdRule = append(newThresholdRule, newThresholdItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RelayerThresholdChanged", newThresholdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRelayerThresholdChanged)
				if err := _Bridge.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
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

// ParseRelayerThresholdChanged is a log parse operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 indexed newThreshold)
func (_Bridge *BridgeFilterer) ParseRelayerThresholdChanged(log types.Log) (*BridgeRelayerThresholdChanged, error) {
	event := new(BridgeRelayerThresholdChanged)
	if err := _Bridge.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bridge contract.
type BridgeRoleGrantedIterator struct {
	Event *BridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *BridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleGranted)
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
		it.Event = new(BridgeRoleGranted)
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
func (it *BridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleGranted represents a RoleGranted event raised by the Bridge contract.
type BridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleGrantedIterator{contract: _Bridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleGranted)
				if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) ParseRoleGranted(log types.Log) (*BridgeRoleGranted, error) {
	event := new(BridgeRoleGranted)
	if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bridge contract.
type BridgeRoleRevokedIterator struct {
	Event *BridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleRevoked)
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
		it.Event = new(BridgeRoleRevoked)
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
func (it *BridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleRevoked represents a RoleRevoked event raised by the Bridge contract.
type BridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleRevokedIterator{contract: _Bridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleRevoked)
				if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) ParseRoleRevoked(log types.Log) (*BridgeRoleRevoked, error) {
	event := new(BridgeRoleRevoked)
	if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bridge contract.
type BridgeUnpausedIterator struct {
	Event *BridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnpaused)
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
		it.Event = new(BridgeUnpaused)
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
func (it *BridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnpaused represents a Unpaused event raised by the Bridge contract.
type BridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeUnpausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeUnpausedIterator{contract: _Bridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnpaused)
				if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) ParseUnpaused(log types.Log) (*BridgeUnpaused, error) {
	event := new(BridgeUnpaused)
	if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
