// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// IStrategyManagerDeprecatedStructQueuedWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IStrategyManagerDeprecatedStructQueuedWithdrawal struct {
	Strategies           []common.Address
	Shares               []*big.Int
	Staker               common.Address
	WithdrawerAndNonce   IStrategyManagerDeprecatedStructWithdrawerAndNonce
	WithdrawalStartBlock uint32
	DelegatedAddress     common.Address
}

// IStrategyManagerDeprecatedStructWithdrawerAndNonce is an auto generated low-level Go binding around an user-defined struct.
type IStrategyManagerDeprecatedStructWithdrawerAndNonce struct {
	Withdrawer common.Address
	Nonce      *big.Int
}

// StrategyManagerMetaData contains all meta data concerning the StrategyManager contract.
var StrategyManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEPOSIT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weth\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStrategiesToDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"thirdPartyTransfersForbiddenValues\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateWithdrawalRoot\",\"inputs\":[{\"name\":\"queuedWithdrawal\",\"type\":\"tuple\",\"internalType\":\"structIStrategyManager.DeprecatedStruct_QueuedWithdrawal\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawerAndNonce\",\"type\":\"tuple\",\"internalType\":\"structIStrategyManager.DeprecatedStruct_WithdrawerAndNonce\",\"components\":[{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"withdrawalStartBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"delegatedAddress\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositETHIntoStrategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositETHIntoStrategyWithSignature\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositWETHIntoStrategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"weth\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositWETHIntoStrategyWithSignature\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"weth\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakerStrategyL1BackShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_slasher\",\"type\":\"address\",\"internalType\":\"contractISlashManager\"},{\"name\":\"_pauser\",\"type\":\"address\",\"internalType\":\"contractIL2Pauser\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"migrateQueuedWithdrawal\",\"inputs\":[{\"name\":\"queuedWithdrawal\",\"type\":\"tuple\",\"internalType\":\"structIStrategyManager.DeprecatedStruct_QueuedWithdrawal\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawerAndNonce\",\"type\":\"tuple\",\"internalType\":\"structIStrategyManager.DeprecatedStruct_WithdrawerAndNonce\",\"components\":[{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"withdrawalStartBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"delegatedAddress\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"migrateRelatedL1StakerShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l1UnStakeMessageNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauser\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIL2Pauser\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"relayer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyWhitelister\",\"inputs\":[{\"name\":\"newStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlashManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyL1BackShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyList\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyListLength\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyIsWhitelistedForDeposit\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyWhitelister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"thirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateStakerStrategyL1BackShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsWeth\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weth\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawalRootPending\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"weth\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MigrateRelatedL1StakerShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"l1UnStakeMessageNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyWhitelisterChanged\",\"inputs\":[{\"name\":\"previousAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedThirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// StrategyManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyManagerMetaData.ABI instead.
var StrategyManagerABI = StrategyManagerMetaData.ABI

// StrategyManager is an auto generated Go binding around an Ethereum contract.
type StrategyManager struct {
	StrategyManagerCaller     // Read-only binding to the contract
	StrategyManagerTransactor // Write-only binding to the contract
	StrategyManagerFilterer   // Log filterer for contract events
}

// StrategyManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategyManagerSession struct {
	Contract     *StrategyManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrategyManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyManagerCallerSession struct {
	Contract *StrategyManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StrategyManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyManagerTransactorSession struct {
	Contract     *StrategyManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StrategyManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyManagerRaw struct {
	Contract *StrategyManager // Generic contract binding to access the raw methods on
}

// StrategyManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyManagerCallerRaw struct {
	Contract *StrategyManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyManagerTransactorRaw struct {
	Contract *StrategyManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategyManager creates a new instance of StrategyManager, bound to a specific deployed contract.
func NewStrategyManager(address common.Address, backend bind.ContractBackend) (*StrategyManager, error) {
	contract, err := bindStrategyManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrategyManager{StrategyManagerCaller: StrategyManagerCaller{contract: contract}, StrategyManagerTransactor: StrategyManagerTransactor{contract: contract}, StrategyManagerFilterer: StrategyManagerFilterer{contract: contract}}, nil
}

// NewStrategyManagerCaller creates a new read-only instance of StrategyManager, bound to a specific deployed contract.
func NewStrategyManagerCaller(address common.Address, caller bind.ContractCaller) (*StrategyManagerCaller, error) {
	contract, err := bindStrategyManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerCaller{contract: contract}, nil
}

// NewStrategyManagerTransactor creates a new write-only instance of StrategyManager, bound to a specific deployed contract.
func NewStrategyManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyManagerTransactor, error) {
	contract, err := bindStrategyManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerTransactor{contract: contract}, nil
}

// NewStrategyManagerFilterer creates a new log filterer instance of StrategyManager, bound to a specific deployed contract.
func NewStrategyManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyManagerFilterer, error) {
	contract, err := bindStrategyManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerFilterer{contract: contract}, nil
}

// bindStrategyManager binds a generic wrapper to an already deployed contract.
func bindStrategyManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrategyManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyManager *StrategyManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyManager.Contract.StrategyManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyManager *StrategyManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyManager.Contract.StrategyManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyManager *StrategyManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyManager.Contract.StrategyManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyManager *StrategyManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyManager *StrategyManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyManager *StrategyManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyManager.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_StrategyManager *StrategyManagerCaller) DEPOSITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "DEPOSIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_StrategyManager *StrategyManagerSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _StrategyManager.Contract.DEPOSITTYPEHASH(&_StrategyManager.CallOpts)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_StrategyManager *StrategyManagerCallerSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _StrategyManager.Contract.DEPOSITTYPEHASH(&_StrategyManager.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_StrategyManager *StrategyManagerCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_StrategyManager *StrategyManagerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _StrategyManager.Contract.DOMAINTYPEHASH(&_StrategyManager.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_StrategyManager *StrategyManagerCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _StrategyManager.Contract.DOMAINTYPEHASH(&_StrategyManager.CallOpts)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0xb43b514b.
//
// Solidity: function calculateWithdrawalRoot((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) pure returns(bytes32)
func (_StrategyManager *StrategyManagerCaller) CalculateWithdrawalRoot(opts *bind.CallOpts, queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) ([32]byte, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "calculateWithdrawalRoot", queuedWithdrawal)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0xb43b514b.
//
// Solidity: function calculateWithdrawalRoot((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) pure returns(bytes32)
func (_StrategyManager *StrategyManagerSession) CalculateWithdrawalRoot(queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) ([32]byte, error) {
	return _StrategyManager.Contract.CalculateWithdrawalRoot(&_StrategyManager.CallOpts, queuedWithdrawal)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0xb43b514b.
//
// Solidity: function calculateWithdrawalRoot((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) pure returns(bytes32)
func (_StrategyManager *StrategyManagerCallerSession) CalculateWithdrawalRoot(queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) ([32]byte, error) {
	return _StrategyManager.Contract.CalculateWithdrawalRoot(&_StrategyManager.CallOpts, queuedWithdrawal)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_StrategyManager *StrategyManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_StrategyManager *StrategyManagerSession) Delegation() (common.Address, error) {
	return _StrategyManager.Contract.Delegation(&_StrategyManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_StrategyManager *StrategyManagerCallerSession) Delegation() (common.Address, error) {
	return _StrategyManager.Contract.Delegation(&_StrategyManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_StrategyManager *StrategyManagerCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_StrategyManager *StrategyManagerSession) DomainSeparator() ([32]byte, error) {
	return _StrategyManager.Contract.DomainSeparator(&_StrategyManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_StrategyManager *StrategyManagerCallerSession) DomainSeparator() ([32]byte, error) {
	return _StrategyManager.Contract.DomainSeparator(&_StrategyManager.CallOpts)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_StrategyManager *StrategyManagerCaller) GetDeposits(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "getDeposits", staker)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_StrategyManager *StrategyManagerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _StrategyManager.Contract.GetDeposits(&_StrategyManager.CallOpts, staker)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_StrategyManager *StrategyManagerCallerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _StrategyManager.Contract.GetDeposits(&_StrategyManager.CallOpts, staker)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_StrategyManager *StrategyManagerCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_StrategyManager *StrategyManagerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.Nonces(&_StrategyManager.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_StrategyManager *StrategyManagerCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.Nonces(&_StrategyManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrategyManager *StrategyManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrategyManager *StrategyManagerSession) Owner() (common.Address, error) {
	return _StrategyManager.Contract.Owner(&_StrategyManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrategyManager *StrategyManagerCallerSession) Owner() (common.Address, error) {
	return _StrategyManager.Contract.Owner(&_StrategyManager.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_StrategyManager *StrategyManagerCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_StrategyManager *StrategyManagerSession) Pauser() (common.Address, error) {
	return _StrategyManager.Contract.Pauser(&_StrategyManager.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_StrategyManager *StrategyManagerCallerSession) Pauser() (common.Address, error) {
	return _StrategyManager.Contract.Pauser(&_StrategyManager.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_StrategyManager *StrategyManagerCaller) Relayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "relayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_StrategyManager *StrategyManagerSession) Relayer() (common.Address, error) {
	return _StrategyManager.Contract.Relayer(&_StrategyManager.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_StrategyManager *StrategyManagerCallerSession) Relayer() (common.Address, error) {
	return _StrategyManager.Contract.Relayer(&_StrategyManager.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_StrategyManager *StrategyManagerCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_StrategyManager *StrategyManagerSession) Slasher() (common.Address, error) {
	return _StrategyManager.Contract.Slasher(&_StrategyManager.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_StrategyManager *StrategyManagerCallerSession) Slasher() (common.Address, error) {
	return _StrategyManager.Contract.Slasher(&_StrategyManager.CallOpts)
}

// StakerStrategyL1BackShares is a free data retrieval call binding the contract method 0xfd769186.
//
// Solidity: function stakerStrategyL1BackShares(address , address ) view returns(uint256)
func (_StrategyManager *StrategyManagerCaller) StakerStrategyL1BackShares(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "stakerStrategyL1BackShares", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyL1BackShares is a free data retrieval call binding the contract method 0xfd769186.
//
// Solidity: function stakerStrategyL1BackShares(address , address ) view returns(uint256)
func (_StrategyManager *StrategyManagerSession) StakerStrategyL1BackShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.StakerStrategyL1BackShares(&_StrategyManager.CallOpts, arg0, arg1)
}

// StakerStrategyL1BackShares is a free data retrieval call binding the contract method 0xfd769186.
//
// Solidity: function stakerStrategyL1BackShares(address , address ) view returns(uint256)
func (_StrategyManager *StrategyManagerCallerSession) StakerStrategyL1BackShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.StakerStrategyL1BackShares(&_StrategyManager.CallOpts, arg0, arg1)
}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address , uint256 ) view returns(address)
func (_StrategyManager *StrategyManagerCaller) StakerStrategyList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "stakerStrategyList", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address , uint256 ) view returns(address)
func (_StrategyManager *StrategyManagerSession) StakerStrategyList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _StrategyManager.Contract.StakerStrategyList(&_StrategyManager.CallOpts, arg0, arg1)
}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address , uint256 ) view returns(address)
func (_StrategyManager *StrategyManagerCallerSession) StakerStrategyList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _StrategyManager.Contract.StakerStrategyList(&_StrategyManager.CallOpts, arg0, arg1)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_StrategyManager *StrategyManagerCaller) StakerStrategyListLength(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "stakerStrategyListLength", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_StrategyManager *StrategyManagerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.StakerStrategyListLength(&_StrategyManager.CallOpts, staker)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_StrategyManager *StrategyManagerCallerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.StakerStrategyListLength(&_StrategyManager.CallOpts, staker)
}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address , address ) view returns(uint256)
func (_StrategyManager *StrategyManagerCaller) StakerStrategyShares(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "stakerStrategyShares", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address , address ) view returns(uint256)
func (_StrategyManager *StrategyManagerSession) StakerStrategyShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.StakerStrategyShares(&_StrategyManager.CallOpts, arg0, arg1)
}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address , address ) view returns(uint256)
func (_StrategyManager *StrategyManagerCallerSession) StakerStrategyShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.StakerStrategyShares(&_StrategyManager.CallOpts, arg0, arg1)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address ) view returns(bool)
func (_StrategyManager *StrategyManagerCaller) StrategyIsWhitelistedForDeposit(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "strategyIsWhitelistedForDeposit", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address ) view returns(bool)
func (_StrategyManager *StrategyManagerSession) StrategyIsWhitelistedForDeposit(arg0 common.Address) (bool, error) {
	return _StrategyManager.Contract.StrategyIsWhitelistedForDeposit(&_StrategyManager.CallOpts, arg0)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address ) view returns(bool)
func (_StrategyManager *StrategyManagerCallerSession) StrategyIsWhitelistedForDeposit(arg0 common.Address) (bool, error) {
	return _StrategyManager.Contract.StrategyIsWhitelistedForDeposit(&_StrategyManager.CallOpts, arg0)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_StrategyManager *StrategyManagerCaller) StrategyWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "strategyWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_StrategyManager *StrategyManagerSession) StrategyWhitelister() (common.Address, error) {
	return _StrategyManager.Contract.StrategyWhitelister(&_StrategyManager.CallOpts)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_StrategyManager *StrategyManagerCallerSession) StrategyWhitelister() (common.Address, error) {
	return _StrategyManager.Contract.StrategyWhitelister(&_StrategyManager.CallOpts)
}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address ) view returns(bool)
func (_StrategyManager *StrategyManagerCaller) ThirdPartyTransfersForbidden(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "thirdPartyTransfersForbidden", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address ) view returns(bool)
func (_StrategyManager *StrategyManagerSession) ThirdPartyTransfersForbidden(arg0 common.Address) (bool, error) {
	return _StrategyManager.Contract.ThirdPartyTransfersForbidden(&_StrategyManager.CallOpts, arg0)
}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address ) view returns(bool)
func (_StrategyManager *StrategyManagerCallerSession) ThirdPartyTransfersForbidden(arg0 common.Address) (bool, error) {
	return _StrategyManager.Contract.ThirdPartyTransfersForbidden(&_StrategyManager.CallOpts, arg0)
}

// WithdrawalRootPending is a free data retrieval call binding the contract method 0xc3c6b3a9.
//
// Solidity: function withdrawalRootPending(bytes32 ) view returns(bool)
func (_StrategyManager *StrategyManagerCaller) WithdrawalRootPending(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "withdrawalRootPending", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalRootPending is a free data retrieval call binding the contract method 0xc3c6b3a9.
//
// Solidity: function withdrawalRootPending(bytes32 ) view returns(bool)
func (_StrategyManager *StrategyManagerSession) WithdrawalRootPending(arg0 [32]byte) (bool, error) {
	return _StrategyManager.Contract.WithdrawalRootPending(&_StrategyManager.CallOpts, arg0)
}

// WithdrawalRootPending is a free data retrieval call binding the contract method 0xc3c6b3a9.
//
// Solidity: function withdrawalRootPending(bytes32 ) view returns(bool)
func (_StrategyManager *StrategyManagerCallerSession) WithdrawalRootPending(arg0 [32]byte) (bool, error) {
	return _StrategyManager.Contract.WithdrawalRootPending(&_StrategyManager.CallOpts, arg0)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address weth, address strategy, uint256 shares) returns()
func (_StrategyManager *StrategyManagerTransactor) AddShares(opts *bind.TransactOpts, staker common.Address, weth common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "addShares", staker, weth, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address weth, address strategy, uint256 shares) returns()
func (_StrategyManager *StrategyManagerSession) AddShares(staker common.Address, weth common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.AddShares(&_StrategyManager.TransactOpts, staker, weth, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address weth, address strategy, uint256 shares) returns()
func (_StrategyManager *StrategyManagerTransactorSession) AddShares(staker common.Address, weth common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.AddShares(&_StrategyManager.TransactOpts, staker, weth, strategy, shares)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0xdf5b3547.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_StrategyManager *StrategyManagerTransactor) AddStrategiesToDepositWhitelist(opts *bind.TransactOpts, strategiesToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "addStrategiesToDepositWhitelist", strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0xdf5b3547.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_StrategyManager *StrategyManagerSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _StrategyManager.Contract.AddStrategiesToDepositWhitelist(&_StrategyManager.TransactOpts, strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0xdf5b3547.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_StrategyManager *StrategyManagerTransactorSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _StrategyManager.Contract.AddStrategiesToDepositWhitelist(&_StrategyManager.TransactOpts, strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
}

// DepositETHIntoStrategy is a paid mutator transaction binding the contract method 0x9573ddbb.
//
// Solidity: function depositETHIntoStrategy(address strategy) payable returns(uint256 shares)
func (_StrategyManager *StrategyManagerTransactor) DepositETHIntoStrategy(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "depositETHIntoStrategy", strategy)
}

// DepositETHIntoStrategy is a paid mutator transaction binding the contract method 0x9573ddbb.
//
// Solidity: function depositETHIntoStrategy(address strategy) payable returns(uint256 shares)
func (_StrategyManager *StrategyManagerSession) DepositETHIntoStrategy(strategy common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.DepositETHIntoStrategy(&_StrategyManager.TransactOpts, strategy)
}

// DepositETHIntoStrategy is a paid mutator transaction binding the contract method 0x9573ddbb.
//
// Solidity: function depositETHIntoStrategy(address strategy) payable returns(uint256 shares)
func (_StrategyManager *StrategyManagerTransactorSession) DepositETHIntoStrategy(strategy common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.DepositETHIntoStrategy(&_StrategyManager.TransactOpts, strategy)
}

// DepositETHIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0xcbc336c2.
//
// Solidity: function depositETHIntoStrategyWithSignature(address strategy, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_StrategyManager *StrategyManagerTransactor) DepositETHIntoStrategyWithSignature(opts *bind.TransactOpts, strategy common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "depositETHIntoStrategyWithSignature", strategy, amount, staker, expiry, signature)
}

// DepositETHIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0xcbc336c2.
//
// Solidity: function depositETHIntoStrategyWithSignature(address strategy, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_StrategyManager *StrategyManagerSession) DepositETHIntoStrategyWithSignature(strategy common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrategyManager.Contract.DepositETHIntoStrategyWithSignature(&_StrategyManager.TransactOpts, strategy, amount, staker, expiry, signature)
}

// DepositETHIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0xcbc336c2.
//
// Solidity: function depositETHIntoStrategyWithSignature(address strategy, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_StrategyManager *StrategyManagerTransactorSession) DepositETHIntoStrategyWithSignature(strategy common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrategyManager.Contract.DepositETHIntoStrategyWithSignature(&_StrategyManager.TransactOpts, strategy, amount, staker, expiry, signature)
}

// DepositWETHIntoStrategy is a paid mutator transaction binding the contract method 0x72e80be6.
//
// Solidity: function depositWETHIntoStrategy(address strategy, address weth, uint256 amount) returns(uint256 shares)
func (_StrategyManager *StrategyManagerTransactor) DepositWETHIntoStrategy(opts *bind.TransactOpts, strategy common.Address, weth common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "depositWETHIntoStrategy", strategy, weth, amount)
}

// DepositWETHIntoStrategy is a paid mutator transaction binding the contract method 0x72e80be6.
//
// Solidity: function depositWETHIntoStrategy(address strategy, address weth, uint256 amount) returns(uint256 shares)
func (_StrategyManager *StrategyManagerSession) DepositWETHIntoStrategy(strategy common.Address, weth common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.DepositWETHIntoStrategy(&_StrategyManager.TransactOpts, strategy, weth, amount)
}

// DepositWETHIntoStrategy is a paid mutator transaction binding the contract method 0x72e80be6.
//
// Solidity: function depositWETHIntoStrategy(address strategy, address weth, uint256 amount) returns(uint256 shares)
func (_StrategyManager *StrategyManagerTransactorSession) DepositWETHIntoStrategy(strategy common.Address, weth common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.DepositWETHIntoStrategy(&_StrategyManager.TransactOpts, strategy, weth, amount)
}

// DepositWETHIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x0891e0ed.
//
// Solidity: function depositWETHIntoStrategyWithSignature(address strategy, address weth, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_StrategyManager *StrategyManagerTransactor) DepositWETHIntoStrategyWithSignature(opts *bind.TransactOpts, strategy common.Address, weth common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "depositWETHIntoStrategyWithSignature", strategy, weth, amount, staker, expiry, signature)
}

// DepositWETHIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x0891e0ed.
//
// Solidity: function depositWETHIntoStrategyWithSignature(address strategy, address weth, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_StrategyManager *StrategyManagerSession) DepositWETHIntoStrategyWithSignature(strategy common.Address, weth common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrategyManager.Contract.DepositWETHIntoStrategyWithSignature(&_StrategyManager.TransactOpts, strategy, weth, amount, staker, expiry, signature)
}

// DepositWETHIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x0891e0ed.
//
// Solidity: function depositWETHIntoStrategyWithSignature(address strategy, address weth, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_StrategyManager *StrategyManagerTransactorSession) DepositWETHIntoStrategyWithSignature(strategy common.Address, weth common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrategyManager.Contract.DepositWETHIntoStrategyWithSignature(&_StrategyManager.TransactOpts, strategy, weth, amount, staker, expiry, signature)
}

// GetStakerStrategyL1BackShares is a paid mutator transaction binding the contract method 0xf0249f19.
//
// Solidity: function getStakerStrategyL1BackShares(address staker, address strategy) returns(uint256)
func (_StrategyManager *StrategyManagerTransactor) GetStakerStrategyL1BackShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "getStakerStrategyL1BackShares", staker, strategy)
}

// GetStakerStrategyL1BackShares is a paid mutator transaction binding the contract method 0xf0249f19.
//
// Solidity: function getStakerStrategyL1BackShares(address staker, address strategy) returns(uint256)
func (_StrategyManager *StrategyManagerSession) GetStakerStrategyL1BackShares(staker common.Address, strategy common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.GetStakerStrategyL1BackShares(&_StrategyManager.TransactOpts, staker, strategy)
}

// GetStakerStrategyL1BackShares is a paid mutator transaction binding the contract method 0xf0249f19.
//
// Solidity: function getStakerStrategyL1BackShares(address staker, address strategy) returns(uint256)
func (_StrategyManager *StrategyManagerTransactorSession) GetStakerStrategyL1BackShares(staker common.Address, strategy common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.GetStakerStrategyL1BackShares(&_StrategyManager.TransactOpts, staker, strategy)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, address _relayer, address _delegation, address _slasher, address _pauser) returns()
func (_StrategyManager *StrategyManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialStrategyWhitelister common.Address, _relayer common.Address, _delegation common.Address, _slasher common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "initialize", initialOwner, initialStrategyWhitelister, _relayer, _delegation, _slasher, _pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, address _relayer, address _delegation, address _slasher, address _pauser) returns()
func (_StrategyManager *StrategyManagerSession) Initialize(initialOwner common.Address, initialStrategyWhitelister common.Address, _relayer common.Address, _delegation common.Address, _slasher common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.Initialize(&_StrategyManager.TransactOpts, initialOwner, initialStrategyWhitelister, _relayer, _delegation, _slasher, _pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, address _relayer, address _delegation, address _slasher, address _pauser) returns()
func (_StrategyManager *StrategyManagerTransactorSession) Initialize(initialOwner common.Address, initialStrategyWhitelister common.Address, _relayer common.Address, _delegation common.Address, _slasher common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.Initialize(&_StrategyManager.TransactOpts, initialOwner, initialStrategyWhitelister, _relayer, _delegation, _slasher, _pauser)
}

// MigrateQueuedWithdrawal is a paid mutator transaction binding the contract method 0xcd293f6f.
//
// Solidity: function migrateQueuedWithdrawal((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) returns(bool, bytes32)
func (_StrategyManager *StrategyManagerTransactor) MigrateQueuedWithdrawal(opts *bind.TransactOpts, queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "migrateQueuedWithdrawal", queuedWithdrawal)
}

// MigrateQueuedWithdrawal is a paid mutator transaction binding the contract method 0xcd293f6f.
//
// Solidity: function migrateQueuedWithdrawal((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) returns(bool, bytes32)
func (_StrategyManager *StrategyManagerSession) MigrateQueuedWithdrawal(queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) (*types.Transaction, error) {
	return _StrategyManager.Contract.MigrateQueuedWithdrawal(&_StrategyManager.TransactOpts, queuedWithdrawal)
}

// MigrateQueuedWithdrawal is a paid mutator transaction binding the contract method 0xcd293f6f.
//
// Solidity: function migrateQueuedWithdrawal((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) returns(bool, bytes32)
func (_StrategyManager *StrategyManagerTransactorSession) MigrateQueuedWithdrawal(queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) (*types.Transaction, error) {
	return _StrategyManager.Contract.MigrateQueuedWithdrawal(&_StrategyManager.TransactOpts, queuedWithdrawal)
}

// MigrateRelatedL1StakerShares is a paid mutator transaction binding the contract method 0x78dab7a8.
//
// Solidity: function migrateRelatedL1StakerShares(address staker, address strategy, uint256 shares, uint256 l1UnStakeMessageNonce) returns(bool)
func (_StrategyManager *StrategyManagerTransactor) MigrateRelatedL1StakerShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int, l1UnStakeMessageNonce *big.Int) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "migrateRelatedL1StakerShares", staker, strategy, shares, l1UnStakeMessageNonce)
}

// MigrateRelatedL1StakerShares is a paid mutator transaction binding the contract method 0x78dab7a8.
//
// Solidity: function migrateRelatedL1StakerShares(address staker, address strategy, uint256 shares, uint256 l1UnStakeMessageNonce) returns(bool)
func (_StrategyManager *StrategyManagerSession) MigrateRelatedL1StakerShares(staker common.Address, strategy common.Address, shares *big.Int, l1UnStakeMessageNonce *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.MigrateRelatedL1StakerShares(&_StrategyManager.TransactOpts, staker, strategy, shares, l1UnStakeMessageNonce)
}

// MigrateRelatedL1StakerShares is a paid mutator transaction binding the contract method 0x78dab7a8.
//
// Solidity: function migrateRelatedL1StakerShares(address staker, address strategy, uint256 shares, uint256 l1UnStakeMessageNonce) returns(bool)
func (_StrategyManager *StrategyManagerTransactorSession) MigrateRelatedL1StakerShares(staker common.Address, strategy common.Address, shares *big.Int, l1UnStakeMessageNonce *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.MigrateRelatedL1StakerShares(&_StrategyManager.TransactOpts, staker, strategy, shares, l1UnStakeMessageNonce)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_StrategyManager *StrategyManagerTransactor) RemoveShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "removeShares", staker, strategy, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_StrategyManager *StrategyManagerSession) RemoveShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.RemoveShares(&_StrategyManager.TransactOpts, staker, strategy, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_StrategyManager *StrategyManagerTransactorSession) RemoveShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.RemoveShares(&_StrategyManager.TransactOpts, staker, strategy, shares)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyManager *StrategyManagerTransactor) RemoveStrategiesFromDepositWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "removeStrategiesFromDepositWhitelist", strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyManager *StrategyManagerSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_StrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyManager *StrategyManagerTransactorSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_StrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrategyManager *StrategyManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrategyManager *StrategyManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _StrategyManager.Contract.RenounceOwnership(&_StrategyManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrategyManager *StrategyManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StrategyManager.Contract.RenounceOwnership(&_StrategyManager.TransactOpts)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_StrategyManager *StrategyManagerTransactor) SetStrategyWhitelister(opts *bind.TransactOpts, newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "setStrategyWhitelister", newStrategyWhitelister)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_StrategyManager *StrategyManagerSession) SetStrategyWhitelister(newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.SetStrategyWhitelister(&_StrategyManager.TransactOpts, newStrategyWhitelister)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_StrategyManager *StrategyManagerTransactorSession) SetStrategyWhitelister(newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.SetStrategyWhitelister(&_StrategyManager.TransactOpts, newStrategyWhitelister)
}

// SetThirdPartyTransfersForbidden is a paid mutator transaction binding the contract method 0x4e5a4263.
//
// Solidity: function setThirdPartyTransfersForbidden(address strategy, bool value) returns()
func (_StrategyManager *StrategyManagerTransactor) SetThirdPartyTransfersForbidden(opts *bind.TransactOpts, strategy common.Address, value bool) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "setThirdPartyTransfersForbidden", strategy, value)
}

// SetThirdPartyTransfersForbidden is a paid mutator transaction binding the contract method 0x4e5a4263.
//
// Solidity: function setThirdPartyTransfersForbidden(address strategy, bool value) returns()
func (_StrategyManager *StrategyManagerSession) SetThirdPartyTransfersForbidden(strategy common.Address, value bool) (*types.Transaction, error) {
	return _StrategyManager.Contract.SetThirdPartyTransfersForbidden(&_StrategyManager.TransactOpts, strategy, value)
}

// SetThirdPartyTransfersForbidden is a paid mutator transaction binding the contract method 0x4e5a4263.
//
// Solidity: function setThirdPartyTransfersForbidden(address strategy, bool value) returns()
func (_StrategyManager *StrategyManagerTransactorSession) SetThirdPartyTransfersForbidden(strategy common.Address, value bool) (*types.Transaction, error) {
	return _StrategyManager.Contract.SetThirdPartyTransfersForbidden(&_StrategyManager.TransactOpts, strategy, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrategyManager *StrategyManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrategyManager *StrategyManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.TransferOwnership(&_StrategyManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrategyManager *StrategyManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.TransferOwnership(&_StrategyManager.TransactOpts, newOwner)
}

// UpdateStakerStrategyL1BackShares is a paid mutator transaction binding the contract method 0xc75a04b7.
//
// Solidity: function updateStakerStrategyL1BackShares(address staker, address strategy, uint256 shares) returns()
func (_StrategyManager *StrategyManagerTransactor) UpdateStakerStrategyL1BackShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "updateStakerStrategyL1BackShares", staker, strategy, shares)
}

// UpdateStakerStrategyL1BackShares is a paid mutator transaction binding the contract method 0xc75a04b7.
//
// Solidity: function updateStakerStrategyL1BackShares(address staker, address strategy, uint256 shares) returns()
func (_StrategyManager *StrategyManagerSession) UpdateStakerStrategyL1BackShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.UpdateStakerStrategyL1BackShares(&_StrategyManager.TransactOpts, staker, strategy, shares)
}

// UpdateStakerStrategyL1BackShares is a paid mutator transaction binding the contract method 0xc75a04b7.
//
// Solidity: function updateStakerStrategyL1BackShares(address staker, address strategy, uint256 shares) returns()
func (_StrategyManager *StrategyManagerTransactorSession) UpdateStakerStrategyL1BackShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.UpdateStakerStrategyL1BackShares(&_StrategyManager.TransactOpts, staker, strategy, shares)
}

// WithdrawSharesAsWeth is a paid mutator transaction binding the contract method 0xe10456fe.
//
// Solidity: function withdrawSharesAsWeth(address recipient, address strategy, uint256 shares, address weth) returns()
func (_StrategyManager *StrategyManagerTransactor) WithdrawSharesAsWeth(opts *bind.TransactOpts, recipient common.Address, strategy common.Address, shares *big.Int, weth common.Address) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "withdrawSharesAsWeth", recipient, strategy, shares, weth)
}

// WithdrawSharesAsWeth is a paid mutator transaction binding the contract method 0xe10456fe.
//
// Solidity: function withdrawSharesAsWeth(address recipient, address strategy, uint256 shares, address weth) returns()
func (_StrategyManager *StrategyManagerSession) WithdrawSharesAsWeth(recipient common.Address, strategy common.Address, shares *big.Int, weth common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.WithdrawSharesAsWeth(&_StrategyManager.TransactOpts, recipient, strategy, shares, weth)
}

// WithdrawSharesAsWeth is a paid mutator transaction binding the contract method 0xe10456fe.
//
// Solidity: function withdrawSharesAsWeth(address recipient, address strategy, uint256 shares, address weth) returns()
func (_StrategyManager *StrategyManagerTransactorSession) WithdrawSharesAsWeth(recipient common.Address, strategy common.Address, shares *big.Int, weth common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.WithdrawSharesAsWeth(&_StrategyManager.TransactOpts, recipient, strategy, shares, weth)
}

// StrategyManagerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the StrategyManager contract.
type StrategyManagerDepositIterator struct {
	Event *StrategyManagerDeposit // Event containing the contract specifics and raw log

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
func (it *StrategyManagerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerDeposit)
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
		it.Event = new(StrategyManagerDeposit)
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
func (it *StrategyManagerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerDeposit represents a Deposit event raised by the StrategyManager contract.
type StrategyManagerDeposit struct {
	Staker   common.Address
	Weth     common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address weth, address strategy, uint256 shares)
func (_StrategyManager *StrategyManagerFilterer) FilterDeposit(opts *bind.FilterOpts) (*StrategyManagerDepositIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerDepositIterator{contract: _StrategyManager.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address weth, address strategy, uint256 shares)
func (_StrategyManager *StrategyManagerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *StrategyManagerDeposit) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerDeposit)
				if err := _StrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address weth, address strategy, uint256 shares)
func (_StrategyManager *StrategyManagerFilterer) ParseDeposit(log types.Log) (*StrategyManagerDeposit, error) {
	event := new(StrategyManagerDeposit)
	if err := _StrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StrategyManager contract.
type StrategyManagerInitializedIterator struct {
	Event *StrategyManagerInitialized // Event containing the contract specifics and raw log

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
func (it *StrategyManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerInitialized)
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
		it.Event = new(StrategyManagerInitialized)
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
func (it *StrategyManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerInitialized represents a Initialized event raised by the StrategyManager contract.
type StrategyManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StrategyManager *StrategyManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*StrategyManagerInitializedIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerInitializedIterator{contract: _StrategyManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StrategyManager *StrategyManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StrategyManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerInitialized)
				if err := _StrategyManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StrategyManager *StrategyManagerFilterer) ParseInitialized(log types.Log) (*StrategyManagerInitialized, error) {
	event := new(StrategyManagerInitialized)
	if err := _StrategyManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerMigrateRelatedL1StakerSharesIterator is returned from FilterMigrateRelatedL1StakerShares and is used to iterate over the raw logs and unpacked data for MigrateRelatedL1StakerShares events raised by the StrategyManager contract.
type StrategyManagerMigrateRelatedL1StakerSharesIterator struct {
	Event *StrategyManagerMigrateRelatedL1StakerShares // Event containing the contract specifics and raw log

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
func (it *StrategyManagerMigrateRelatedL1StakerSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerMigrateRelatedL1StakerShares)
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
		it.Event = new(StrategyManagerMigrateRelatedL1StakerShares)
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
func (it *StrategyManagerMigrateRelatedL1StakerSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerMigrateRelatedL1StakerSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerMigrateRelatedL1StakerShares represents a MigrateRelatedL1StakerShares event raised by the StrategyManager contract.
type StrategyManagerMigrateRelatedL1StakerShares struct {
	Staker                common.Address
	Strategy              common.Address
	Shares                *big.Int
	L1UnStakeMessageNonce *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMigrateRelatedL1StakerShares is a free log retrieval operation binding the contract event 0xce34b7dd67090a34b625a73404430e44a695f90160883cd318b40858e856a2b6.
//
// Solidity: event MigrateRelatedL1StakerShares(address staker, address strategy, uint256 shares, uint256 l1UnStakeMessageNonce)
func (_StrategyManager *StrategyManagerFilterer) FilterMigrateRelatedL1StakerShares(opts *bind.FilterOpts) (*StrategyManagerMigrateRelatedL1StakerSharesIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "MigrateRelatedL1StakerShares")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerMigrateRelatedL1StakerSharesIterator{contract: _StrategyManager.contract, event: "MigrateRelatedL1StakerShares", logs: logs, sub: sub}, nil
}

// WatchMigrateRelatedL1StakerShares is a free log subscription operation binding the contract event 0xce34b7dd67090a34b625a73404430e44a695f90160883cd318b40858e856a2b6.
//
// Solidity: event MigrateRelatedL1StakerShares(address staker, address strategy, uint256 shares, uint256 l1UnStakeMessageNonce)
func (_StrategyManager *StrategyManagerFilterer) WatchMigrateRelatedL1StakerShares(opts *bind.WatchOpts, sink chan<- *StrategyManagerMigrateRelatedL1StakerShares) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "MigrateRelatedL1StakerShares")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerMigrateRelatedL1StakerShares)
				if err := _StrategyManager.contract.UnpackLog(event, "MigrateRelatedL1StakerShares", log); err != nil {
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

// ParseMigrateRelatedL1StakerShares is a log parse operation binding the contract event 0xce34b7dd67090a34b625a73404430e44a695f90160883cd318b40858e856a2b6.
//
// Solidity: event MigrateRelatedL1StakerShares(address staker, address strategy, uint256 shares, uint256 l1UnStakeMessageNonce)
func (_StrategyManager *StrategyManagerFilterer) ParseMigrateRelatedL1StakerShares(log types.Log) (*StrategyManagerMigrateRelatedL1StakerShares, error) {
	event := new(StrategyManagerMigrateRelatedL1StakerShares)
	if err := _StrategyManager.contract.UnpackLog(event, "MigrateRelatedL1StakerShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StrategyManager contract.
type StrategyManagerOwnershipTransferredIterator struct {
	Event *StrategyManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StrategyManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerOwnershipTransferred)
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
		it.Event = new(StrategyManagerOwnershipTransferred)
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
func (it *StrategyManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerOwnershipTransferred represents a OwnershipTransferred event raised by the StrategyManager contract.
type StrategyManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrategyManager *StrategyManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StrategyManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerOwnershipTransferredIterator{contract: _StrategyManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrategyManager *StrategyManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StrategyManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerOwnershipTransferred)
				if err := _StrategyManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StrategyManager *StrategyManagerFilterer) ParseOwnershipTransferred(log types.Log) (*StrategyManagerOwnershipTransferred, error) {
	event := new(StrategyManagerOwnershipTransferred)
	if err := _StrategyManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerStrategyAddedToDepositWhitelistIterator is returned from FilterStrategyAddedToDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyAddedToDepositWhitelist events raised by the StrategyManager contract.
type StrategyManagerStrategyAddedToDepositWhitelistIterator struct {
	Event *StrategyManagerStrategyAddedToDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStrategyAddedToDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStrategyAddedToDepositWhitelist)
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
		it.Event = new(StrategyManagerStrategyAddedToDepositWhitelist)
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
func (it *StrategyManagerStrategyAddedToDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStrategyAddedToDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStrategyAddedToDepositWhitelist represents a StrategyAddedToDepositWhitelist event raised by the StrategyManager contract.
type StrategyManagerStrategyAddedToDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToDepositWhitelist is a free log retrieval operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_StrategyManager *StrategyManagerFilterer) FilterStrategyAddedToDepositWhitelist(opts *bind.FilterOpts) (*StrategyManagerStrategyAddedToDepositWhitelistIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStrategyAddedToDepositWhitelistIterator{contract: _StrategyManager.contract, event: "StrategyAddedToDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToDepositWhitelist is a free log subscription operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_StrategyManager *StrategyManagerFilterer) WatchStrategyAddedToDepositWhitelist(opts *bind.WatchOpts, sink chan<- *StrategyManagerStrategyAddedToDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStrategyAddedToDepositWhitelist)
				if err := _StrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
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

// ParseStrategyAddedToDepositWhitelist is a log parse operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_StrategyManager *StrategyManagerFilterer) ParseStrategyAddedToDepositWhitelist(log types.Log) (*StrategyManagerStrategyAddedToDepositWhitelist, error) {
	event := new(StrategyManagerStrategyAddedToDepositWhitelist)
	if err := _StrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerStrategyRemovedFromDepositWhitelistIterator is returned from FilterStrategyRemovedFromDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromDepositWhitelist events raised by the StrategyManager contract.
type StrategyManagerStrategyRemovedFromDepositWhitelistIterator struct {
	Event *StrategyManagerStrategyRemovedFromDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStrategyRemovedFromDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStrategyRemovedFromDepositWhitelist)
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
		it.Event = new(StrategyManagerStrategyRemovedFromDepositWhitelist)
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
func (it *StrategyManagerStrategyRemovedFromDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStrategyRemovedFromDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStrategyRemovedFromDepositWhitelist represents a StrategyRemovedFromDepositWhitelist event raised by the StrategyManager contract.
type StrategyManagerStrategyRemovedFromDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromDepositWhitelist is a free log retrieval operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_StrategyManager *StrategyManagerFilterer) FilterStrategyRemovedFromDepositWhitelist(opts *bind.FilterOpts) (*StrategyManagerStrategyRemovedFromDepositWhitelistIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStrategyRemovedFromDepositWhitelistIterator{contract: _StrategyManager.contract, event: "StrategyRemovedFromDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromDepositWhitelist is a free log subscription operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_StrategyManager *StrategyManagerFilterer) WatchStrategyRemovedFromDepositWhitelist(opts *bind.WatchOpts, sink chan<- *StrategyManagerStrategyRemovedFromDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStrategyRemovedFromDepositWhitelist)
				if err := _StrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
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

// ParseStrategyRemovedFromDepositWhitelist is a log parse operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_StrategyManager *StrategyManagerFilterer) ParseStrategyRemovedFromDepositWhitelist(log types.Log) (*StrategyManagerStrategyRemovedFromDepositWhitelist, error) {
	event := new(StrategyManagerStrategyRemovedFromDepositWhitelist)
	if err := _StrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerStrategyWhitelisterChangedIterator is returned from FilterStrategyWhitelisterChanged and is used to iterate over the raw logs and unpacked data for StrategyWhitelisterChanged events raised by the StrategyManager contract.
type StrategyManagerStrategyWhitelisterChangedIterator struct {
	Event *StrategyManagerStrategyWhitelisterChanged // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStrategyWhitelisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStrategyWhitelisterChanged)
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
		it.Event = new(StrategyManagerStrategyWhitelisterChanged)
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
func (it *StrategyManagerStrategyWhitelisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStrategyWhitelisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStrategyWhitelisterChanged represents a StrategyWhitelisterChanged event raised by the StrategyManager contract.
type StrategyManagerStrategyWhitelisterChanged struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStrategyWhitelisterChanged is a free log retrieval operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_StrategyManager *StrategyManagerFilterer) FilterStrategyWhitelisterChanged(opts *bind.FilterOpts) (*StrategyManagerStrategyWhitelisterChangedIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStrategyWhitelisterChangedIterator{contract: _StrategyManager.contract, event: "StrategyWhitelisterChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyWhitelisterChanged is a free log subscription operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_StrategyManager *StrategyManagerFilterer) WatchStrategyWhitelisterChanged(opts *bind.WatchOpts, sink chan<- *StrategyManagerStrategyWhitelisterChanged) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStrategyWhitelisterChanged)
				if err := _StrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
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

// ParseStrategyWhitelisterChanged is a log parse operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_StrategyManager *StrategyManagerFilterer) ParseStrategyWhitelisterChanged(log types.Log) (*StrategyManagerStrategyWhitelisterChanged, error) {
	event := new(StrategyManagerStrategyWhitelisterChanged)
	if err := _StrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerUpdatedThirdPartyTransfersForbiddenIterator is returned from FilterUpdatedThirdPartyTransfersForbidden and is used to iterate over the raw logs and unpacked data for UpdatedThirdPartyTransfersForbidden events raised by the StrategyManager contract.
type StrategyManagerUpdatedThirdPartyTransfersForbiddenIterator struct {
	Event *StrategyManagerUpdatedThirdPartyTransfersForbidden // Event containing the contract specifics and raw log

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
func (it *StrategyManagerUpdatedThirdPartyTransfersForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerUpdatedThirdPartyTransfersForbidden)
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
		it.Event = new(StrategyManagerUpdatedThirdPartyTransfersForbidden)
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
func (it *StrategyManagerUpdatedThirdPartyTransfersForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerUpdatedThirdPartyTransfersForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerUpdatedThirdPartyTransfersForbidden represents a UpdatedThirdPartyTransfersForbidden event raised by the StrategyManager contract.
type StrategyManagerUpdatedThirdPartyTransfersForbidden struct {
	Strategy common.Address
	Value    bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedThirdPartyTransfersForbidden is a free log retrieval operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address strategy, bool value)
func (_StrategyManager *StrategyManagerFilterer) FilterUpdatedThirdPartyTransfersForbidden(opts *bind.FilterOpts) (*StrategyManagerUpdatedThirdPartyTransfersForbiddenIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "UpdatedThirdPartyTransfersForbidden")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerUpdatedThirdPartyTransfersForbiddenIterator{contract: _StrategyManager.contract, event: "UpdatedThirdPartyTransfersForbidden", logs: logs, sub: sub}, nil
}

// WatchUpdatedThirdPartyTransfersForbidden is a free log subscription operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address strategy, bool value)
func (_StrategyManager *StrategyManagerFilterer) WatchUpdatedThirdPartyTransfersForbidden(opts *bind.WatchOpts, sink chan<- *StrategyManagerUpdatedThirdPartyTransfersForbidden) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "UpdatedThirdPartyTransfersForbidden")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerUpdatedThirdPartyTransfersForbidden)
				if err := _StrategyManager.contract.UnpackLog(event, "UpdatedThirdPartyTransfersForbidden", log); err != nil {
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

// ParseUpdatedThirdPartyTransfersForbidden is a log parse operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address strategy, bool value)
func (_StrategyManager *StrategyManagerFilterer) ParseUpdatedThirdPartyTransfersForbidden(log types.Log) (*StrategyManagerUpdatedThirdPartyTransfersForbidden, error) {
	event := new(StrategyManagerUpdatedThirdPartyTransfersForbidden)
	if err := _StrategyManager.contract.UnpackLog(event, "UpdatedThirdPartyTransfersForbidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
