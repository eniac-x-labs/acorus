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

// IUnstakeRequestsManagerWriterequestsInfo is an auto generated low-level Go binding around an user-defined struct.
type IUnstakeRequestsManagerWriterequestsInfo struct {
	RequestAddress      common.Address
	UnStakeMessageNonce *big.Int
}

// UnstakeRequestsManagerInit is an auto generated low-level Go binding around an user-defined struct.
type UnstakeRequestsManagerInit struct {
	Admin                    common.Address
	Manager                  common.Address
	RequestCanceller         common.Address
	NumberOfBlocksToFinalize *big.Int
}

// UnstakeRequestsManagerMetaData contains all meta data concerning the UnstakeRequestsManager contract.
var UnstakeRequestsManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MANAGER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REQUEST_CANCELLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocateETH\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"allocatedETHDeficit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocatedETHForClaims\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocatedETHSurplus\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"requests\",\"type\":\"tuple[]\",\"internalType\":\"structIUnstakeRequestsManagerWrite.requestsInfo[]\",\"components\":[{\"name\":\"requestAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"unStakeMessageNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"create\",\"inputs\":[{\"name\":\"requester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l2Strategy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dETHLocked\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ethRequested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"currentRequestedCumulativeETH\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dEthLockedAmount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLocator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractL1ILocator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"init\",\"type\":\"tuple\",\"internalType\":\"structUnstakeRequestsManager.Init\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"manager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestCanceller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"numberOfBlocksToFinalize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l2ChainStrategyAmount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2ChainStrategyBlockNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestCumulativeETHRequested\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"locator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numberOfBlocksToFinalize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestByID\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Strategy\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestInfo\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Strategy\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLocator\",\"inputs\":[{\"name\":\"_locator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNumberOfBlocksToFinalize\",\"inputs\":[{\"name\":\"numberOfBlocksToFinalize_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalClaimed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawAllocatedETHSurplus\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtocolConfigChanged\",\"inputs\":[{\"name\":\"setterSelector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"setterSignature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeRequestCancelled\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"requester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dETHLocked\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"ethRequested\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"cumulativeETHRequested\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeRequestClaimed\",\"inputs\":[{\"name\":\"l2strategy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ethRequested\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"dETHLocked\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"csBlockNumber\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"bridgeAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"unStakeMessageNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeRequestCreated\",\"inputs\":[{\"name\":\"requester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dETHLocked\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"ethRequested\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"cumulativeETHRequested\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DoesNotReceiveETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoRequests\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughFunds\",\"inputs\":[{\"name\":\"cumulativeETHOnRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allocatedETHForClaims\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRequester\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotStakingManagerContract\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// UnstakeRequestsManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use UnstakeRequestsManagerMetaData.ABI instead.
var UnstakeRequestsManagerABI = UnstakeRequestsManagerMetaData.ABI

// UnstakeRequestsManager is an auto generated Go binding around an Ethereum contract.
type UnstakeRequestsManager struct {
	UnstakeRequestsManagerCaller     // Read-only binding to the contract
	UnstakeRequestsManagerTransactor // Write-only binding to the contract
	UnstakeRequestsManagerFilterer   // Log filterer for contract events
}

// UnstakeRequestsManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnstakeRequestsManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnstakeRequestsManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnstakeRequestsManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnstakeRequestsManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnstakeRequestsManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnstakeRequestsManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnstakeRequestsManagerSession struct {
	Contract     *UnstakeRequestsManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UnstakeRequestsManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnstakeRequestsManagerCallerSession struct {
	Contract *UnstakeRequestsManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// UnstakeRequestsManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnstakeRequestsManagerTransactorSession struct {
	Contract     *UnstakeRequestsManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// UnstakeRequestsManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnstakeRequestsManagerRaw struct {
	Contract *UnstakeRequestsManager // Generic contract binding to access the raw methods on
}

// UnstakeRequestsManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnstakeRequestsManagerCallerRaw struct {
	Contract *UnstakeRequestsManagerCaller // Generic read-only contract binding to access the raw methods on
}

// UnstakeRequestsManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnstakeRequestsManagerTransactorRaw struct {
	Contract *UnstakeRequestsManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnstakeRequestsManager creates a new instance of UnstakeRequestsManager, bound to a specific deployed contract.
func NewUnstakeRequestsManager(address common.Address, backend bind.ContractBackend) (*UnstakeRequestsManager, error) {
	contract, err := bindUnstakeRequestsManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnstakeRequestsManager{UnstakeRequestsManagerCaller: UnstakeRequestsManagerCaller{contract: contract}, UnstakeRequestsManagerTransactor: UnstakeRequestsManagerTransactor{contract: contract}, UnstakeRequestsManagerFilterer: UnstakeRequestsManagerFilterer{contract: contract}}, nil
}

// NewUnstakeRequestsManagerCaller creates a new read-only instance of UnstakeRequestsManager, bound to a specific deployed contract.
func NewUnstakeRequestsManagerCaller(address common.Address, caller bind.ContractCaller) (*UnstakeRequestsManagerCaller, error) {
	contract, err := bindUnstakeRequestsManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnstakeRequestsManagerCaller{contract: contract}, nil
}

// NewUnstakeRequestsManagerTransactor creates a new write-only instance of UnstakeRequestsManager, bound to a specific deployed contract.
func NewUnstakeRequestsManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*UnstakeRequestsManagerTransactor, error) {
	contract, err := bindUnstakeRequestsManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnstakeRequestsManagerTransactor{contract: contract}, nil
}

// NewUnstakeRequestsManagerFilterer creates a new log filterer instance of UnstakeRequestsManager, bound to a specific deployed contract.
func NewUnstakeRequestsManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*UnstakeRequestsManagerFilterer, error) {
	contract, err := bindUnstakeRequestsManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnstakeRequestsManagerFilterer{contract: contract}, nil
}

// bindUnstakeRequestsManager binds a generic wrapper to an already deployed contract.
func bindUnstakeRequestsManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UnstakeRequestsManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnstakeRequestsManager *UnstakeRequestsManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnstakeRequestsManager.Contract.UnstakeRequestsManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnstakeRequestsManager *UnstakeRequestsManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.UnstakeRequestsManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnstakeRequestsManager *UnstakeRequestsManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.UnstakeRequestsManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnstakeRequestsManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UnstakeRequestsManager.Contract.DEFAULTADMINROLE(&_UnstakeRequestsManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UnstakeRequestsManager.Contract.DEFAULTADMINROLE(&_UnstakeRequestsManager.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) MANAGERROLE() ([32]byte, error) {
	return _UnstakeRequestsManager.Contract.MANAGERROLE(&_UnstakeRequestsManager.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) MANAGERROLE() ([32]byte, error) {
	return _UnstakeRequestsManager.Contract.MANAGERROLE(&_UnstakeRequestsManager.CallOpts)
}

// REQUESTCANCELLERROLE is a free data retrieval call binding the contract method 0xfe3af1c1.
//
// Solidity: function REQUEST_CANCELLER_ROLE() view returns(bytes32)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) REQUESTCANCELLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "REQUEST_CANCELLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REQUESTCANCELLERROLE is a free data retrieval call binding the contract method 0xfe3af1c1.
//
// Solidity: function REQUEST_CANCELLER_ROLE() view returns(bytes32)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) REQUESTCANCELLERROLE() ([32]byte, error) {
	return _UnstakeRequestsManager.Contract.REQUESTCANCELLERROLE(&_UnstakeRequestsManager.CallOpts)
}

// REQUESTCANCELLERROLE is a free data retrieval call binding the contract method 0xfe3af1c1.
//
// Solidity: function REQUEST_CANCELLER_ROLE() view returns(bytes32)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) REQUESTCANCELLERROLE() ([32]byte, error) {
	return _UnstakeRequestsManager.Contract.REQUESTCANCELLERROLE(&_UnstakeRequestsManager.CallOpts)
}

// AllocatedETHDeficit is a free data retrieval call binding the contract method 0x106b263c.
//
// Solidity: function allocatedETHDeficit() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) AllocatedETHDeficit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "allocatedETHDeficit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllocatedETHDeficit is a free data retrieval call binding the contract method 0x106b263c.
//
// Solidity: function allocatedETHDeficit() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) AllocatedETHDeficit() (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.AllocatedETHDeficit(&_UnstakeRequestsManager.CallOpts)
}

// AllocatedETHDeficit is a free data retrieval call binding the contract method 0x106b263c.
//
// Solidity: function allocatedETHDeficit() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) AllocatedETHDeficit() (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.AllocatedETHDeficit(&_UnstakeRequestsManager.CallOpts)
}

// AllocatedETHForClaims is a free data retrieval call binding the contract method 0x1453444d.
//
// Solidity: function allocatedETHForClaims() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) AllocatedETHForClaims(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "allocatedETHForClaims")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllocatedETHForClaims is a free data retrieval call binding the contract method 0x1453444d.
//
// Solidity: function allocatedETHForClaims() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) AllocatedETHForClaims() (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.AllocatedETHForClaims(&_UnstakeRequestsManager.CallOpts)
}

// AllocatedETHForClaims is a free data retrieval call binding the contract method 0x1453444d.
//
// Solidity: function allocatedETHForClaims() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) AllocatedETHForClaims() (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.AllocatedETHForClaims(&_UnstakeRequestsManager.CallOpts)
}

// AllocatedETHSurplus is a free data retrieval call binding the contract method 0xc02d4c8e.
//
// Solidity: function allocatedETHSurplus() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) AllocatedETHSurplus(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "allocatedETHSurplus")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllocatedETHSurplus is a free data retrieval call binding the contract method 0xc02d4c8e.
//
// Solidity: function allocatedETHSurplus() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) AllocatedETHSurplus() (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.AllocatedETHSurplus(&_UnstakeRequestsManager.CallOpts)
}

// AllocatedETHSurplus is a free data retrieval call binding the contract method 0xc02d4c8e.
//
// Solidity: function allocatedETHSurplus() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) AllocatedETHSurplus() (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.AllocatedETHSurplus(&_UnstakeRequestsManager.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) Balance() (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.Balance(&_UnstakeRequestsManager.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) Balance() (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.Balance(&_UnstakeRequestsManager.CallOpts)
}

// CurrentRequestedCumulativeETH is a free data retrieval call binding the contract method 0xf48811a5.
//
// Solidity: function currentRequestedCumulativeETH(uint256 , address ) view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) CurrentRequestedCumulativeETH(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "currentRequestedCumulativeETH", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRequestedCumulativeETH is a free data retrieval call binding the contract method 0xf48811a5.
//
// Solidity: function currentRequestedCumulativeETH(uint256 , address ) view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) CurrentRequestedCumulativeETH(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.CurrentRequestedCumulativeETH(&_UnstakeRequestsManager.CallOpts, arg0, arg1)
}

// CurrentRequestedCumulativeETH is a free data retrieval call binding the contract method 0xf48811a5.
//
// Solidity: function currentRequestedCumulativeETH(uint256 , address ) view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) CurrentRequestedCumulativeETH(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.CurrentRequestedCumulativeETH(&_UnstakeRequestsManager.CallOpts, arg0, arg1)
}

// DEthLockedAmount is a free data retrieval call binding the contract method 0xa75c3e17.
//
// Solidity: function dEthLockedAmount(uint256 , address ) view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) DEthLockedAmount(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "dEthLockedAmount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEthLockedAmount is a free data retrieval call binding the contract method 0xa75c3e17.
//
// Solidity: function dEthLockedAmount(uint256 , address ) view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) DEthLockedAmount(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.DEthLockedAmount(&_UnstakeRequestsManager.CallOpts, arg0, arg1)
}

// DEthLockedAmount is a free data retrieval call binding the contract method 0xa75c3e17.
//
// Solidity: function dEthLockedAmount(uint256 , address ) view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) DEthLockedAmount(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.DEthLockedAmount(&_UnstakeRequestsManager.CallOpts, arg0, arg1)
}

// GetLocator is a free data retrieval call binding the contract method 0xd8343dcb.
//
// Solidity: function getLocator() view returns(address)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) GetLocator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "getLocator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLocator is a free data retrieval call binding the contract method 0xd8343dcb.
//
// Solidity: function getLocator() view returns(address)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) GetLocator() (common.Address, error) {
	return _UnstakeRequestsManager.Contract.GetLocator(&_UnstakeRequestsManager.CallOpts)
}

// GetLocator is a free data retrieval call binding the contract method 0xd8343dcb.
//
// Solidity: function getLocator() view returns(address)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) GetLocator() (common.Address, error) {
	return _UnstakeRequestsManager.Contract.GetLocator(&_UnstakeRequestsManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UnstakeRequestsManager.Contract.GetRoleAdmin(&_UnstakeRequestsManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UnstakeRequestsManager.Contract.GetRoleAdmin(&_UnstakeRequestsManager.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _UnstakeRequestsManager.Contract.GetRoleMember(&_UnstakeRequestsManager.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _UnstakeRequestsManager.Contract.GetRoleMember(&_UnstakeRequestsManager.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.GetRoleMemberCount(&_UnstakeRequestsManager.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.GetRoleMemberCount(&_UnstakeRequestsManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UnstakeRequestsManager.Contract.HasRole(&_UnstakeRequestsManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UnstakeRequestsManager.Contract.HasRole(&_UnstakeRequestsManager.CallOpts, role, account)
}

// L2ChainStrategyAmount is a free data retrieval call binding the contract method 0xa5a05aee.
//
// Solidity: function l2ChainStrategyAmount(uint256 , address ) view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) L2ChainStrategyAmount(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "l2ChainStrategyAmount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ChainStrategyAmount is a free data retrieval call binding the contract method 0xa5a05aee.
//
// Solidity: function l2ChainStrategyAmount(uint256 , address ) view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) L2ChainStrategyAmount(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.L2ChainStrategyAmount(&_UnstakeRequestsManager.CallOpts, arg0, arg1)
}

// L2ChainStrategyAmount is a free data retrieval call binding the contract method 0xa5a05aee.
//
// Solidity: function l2ChainStrategyAmount(uint256 , address ) view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) L2ChainStrategyAmount(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.L2ChainStrategyAmount(&_UnstakeRequestsManager.CallOpts, arg0, arg1)
}

// L2ChainStrategyBlockNumber is a free data retrieval call binding the contract method 0x88a3e1af.
//
// Solidity: function l2ChainStrategyBlockNumber(uint256 , address ) view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) L2ChainStrategyBlockNumber(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "l2ChainStrategyBlockNumber", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ChainStrategyBlockNumber is a free data retrieval call binding the contract method 0x88a3e1af.
//
// Solidity: function l2ChainStrategyBlockNumber(uint256 , address ) view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) L2ChainStrategyBlockNumber(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.L2ChainStrategyBlockNumber(&_UnstakeRequestsManager.CallOpts, arg0, arg1)
}

// L2ChainStrategyBlockNumber is a free data retrieval call binding the contract method 0x88a3e1af.
//
// Solidity: function l2ChainStrategyBlockNumber(uint256 , address ) view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) L2ChainStrategyBlockNumber(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.L2ChainStrategyBlockNumber(&_UnstakeRequestsManager.CallOpts, arg0, arg1)
}

// LatestCumulativeETHRequested is a free data retrieval call binding the contract method 0x278c5acd.
//
// Solidity: function latestCumulativeETHRequested() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) LatestCumulativeETHRequested(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "latestCumulativeETHRequested")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestCumulativeETHRequested is a free data retrieval call binding the contract method 0x278c5acd.
//
// Solidity: function latestCumulativeETHRequested() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) LatestCumulativeETHRequested() (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.LatestCumulativeETHRequested(&_UnstakeRequestsManager.CallOpts)
}

// LatestCumulativeETHRequested is a free data retrieval call binding the contract method 0x278c5acd.
//
// Solidity: function latestCumulativeETHRequested() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) LatestCumulativeETHRequested() (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.LatestCumulativeETHRequested(&_UnstakeRequestsManager.CallOpts)
}

// Locator is a free data retrieval call binding the contract method 0x7c957fc8.
//
// Solidity: function locator() view returns(address)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) Locator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "locator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Locator is a free data retrieval call binding the contract method 0x7c957fc8.
//
// Solidity: function locator() view returns(address)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) Locator() (common.Address, error) {
	return _UnstakeRequestsManager.Contract.Locator(&_UnstakeRequestsManager.CallOpts)
}

// Locator is a free data retrieval call binding the contract method 0x7c957fc8.
//
// Solidity: function locator() view returns(address)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) Locator() (common.Address, error) {
	return _UnstakeRequestsManager.Contract.Locator(&_UnstakeRequestsManager.CallOpts)
}

// NumberOfBlocksToFinalize is a free data retrieval call binding the contract method 0xae5cf272.
//
// Solidity: function numberOfBlocksToFinalize() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) NumberOfBlocksToFinalize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "numberOfBlocksToFinalize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfBlocksToFinalize is a free data retrieval call binding the contract method 0xae5cf272.
//
// Solidity: function numberOfBlocksToFinalize() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) NumberOfBlocksToFinalize() (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.NumberOfBlocksToFinalize(&_UnstakeRequestsManager.CallOpts)
}

// NumberOfBlocksToFinalize is a free data retrieval call binding the contract method 0xae5cf272.
//
// Solidity: function numberOfBlocksToFinalize() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) NumberOfBlocksToFinalize() (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.NumberOfBlocksToFinalize(&_UnstakeRequestsManager.CallOpts)
}

// RequestByID is a free data retrieval call binding the contract method 0x7b2071c6.
//
// Solidity: function requestByID(uint256 destChainId, address l2Strategy) view returns(uint256, uint256, uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) RequestByID(opts *bind.CallOpts, destChainId *big.Int, l2Strategy common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "requestByID", destChainId, l2Strategy)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// RequestByID is a free data retrieval call binding the contract method 0x7b2071c6.
//
// Solidity: function requestByID(uint256 destChainId, address l2Strategy) view returns(uint256, uint256, uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) RequestByID(destChainId *big.Int, l2Strategy common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _UnstakeRequestsManager.Contract.RequestByID(&_UnstakeRequestsManager.CallOpts, destChainId, l2Strategy)
}

// RequestByID is a free data retrieval call binding the contract method 0x7b2071c6.
//
// Solidity: function requestByID(uint256 destChainId, address l2Strategy) view returns(uint256, uint256, uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) RequestByID(destChainId *big.Int, l2Strategy common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _UnstakeRequestsManager.Contract.RequestByID(&_UnstakeRequestsManager.CallOpts, destChainId, l2Strategy)
}

// RequestInfo is a free data retrieval call binding the contract method 0xd4be074f.
//
// Solidity: function requestInfo(uint256 destChainId, address l2Strategy) view returns(bool, uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) RequestInfo(opts *bind.CallOpts, destChainId *big.Int, l2Strategy common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "requestInfo", destChainId, l2Strategy)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RequestInfo is a free data retrieval call binding the contract method 0xd4be074f.
//
// Solidity: function requestInfo(uint256 destChainId, address l2Strategy) view returns(bool, uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) RequestInfo(destChainId *big.Int, l2Strategy common.Address) (bool, *big.Int, error) {
	return _UnstakeRequestsManager.Contract.RequestInfo(&_UnstakeRequestsManager.CallOpts, destChainId, l2Strategy)
}

// RequestInfo is a free data retrieval call binding the contract method 0xd4be074f.
//
// Solidity: function requestInfo(uint256 destChainId, address l2Strategy) view returns(bool, uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) RequestInfo(destChainId *big.Int, l2Strategy common.Address) (bool, *big.Int, error) {
	return _UnstakeRequestsManager.Contract.RequestInfo(&_UnstakeRequestsManager.CallOpts, destChainId, l2Strategy)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UnstakeRequestsManager.Contract.SupportsInterface(&_UnstakeRequestsManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UnstakeRequestsManager.Contract.SupportsInterface(&_UnstakeRequestsManager.CallOpts, interfaceId)
}

// TotalClaimed is a free data retrieval call binding the contract method 0xd54ad2a1.
//
// Solidity: function totalClaimed() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCaller) TotalClaimed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnstakeRequestsManager.contract.Call(opts, &out, "totalClaimed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimed is a free data retrieval call binding the contract method 0xd54ad2a1.
//
// Solidity: function totalClaimed() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) TotalClaimed() (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.TotalClaimed(&_UnstakeRequestsManager.CallOpts)
}

// TotalClaimed is a free data retrieval call binding the contract method 0xd54ad2a1.
//
// Solidity: function totalClaimed() view returns(uint256)
func (_UnstakeRequestsManager *UnstakeRequestsManagerCallerSession) TotalClaimed() (*big.Int, error) {
	return _UnstakeRequestsManager.Contract.TotalClaimed(&_UnstakeRequestsManager.CallOpts)
}

// AllocateETH is a paid mutator transaction binding the contract method 0x4d13bfa6.
//
// Solidity: function allocateETH() payable returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactor) AllocateETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnstakeRequestsManager.contract.Transact(opts, "allocateETH")
}

// AllocateETH is a paid mutator transaction binding the contract method 0x4d13bfa6.
//
// Solidity: function allocateETH() payable returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) AllocateETH() (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.AllocateETH(&_UnstakeRequestsManager.TransactOpts)
}

// AllocateETH is a paid mutator transaction binding the contract method 0x4d13bfa6.
//
// Solidity: function allocateETH() payable returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactorSession) AllocateETH() (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.AllocateETH(&_UnstakeRequestsManager.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x0597ca9c.
//
// Solidity: function claim((address,uint256)[] requests, uint256 sourceChainId, uint256 destChainId, uint256 gasLimit) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactor) Claim(opts *bind.TransactOpts, requests []IUnstakeRequestsManagerWriterequestsInfo, sourceChainId *big.Int, destChainId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _UnstakeRequestsManager.contract.Transact(opts, "claim", requests, sourceChainId, destChainId, gasLimit)
}

// Claim is a paid mutator transaction binding the contract method 0x0597ca9c.
//
// Solidity: function claim((address,uint256)[] requests, uint256 sourceChainId, uint256 destChainId, uint256 gasLimit) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) Claim(requests []IUnstakeRequestsManagerWriterequestsInfo, sourceChainId *big.Int, destChainId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.Claim(&_UnstakeRequestsManager.TransactOpts, requests, sourceChainId, destChainId, gasLimit)
}

// Claim is a paid mutator transaction binding the contract method 0x0597ca9c.
//
// Solidity: function claim((address,uint256)[] requests, uint256 sourceChainId, uint256 destChainId, uint256 gasLimit) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactorSession) Claim(requests []IUnstakeRequestsManagerWriterequestsInfo, sourceChainId *big.Int, destChainId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.Claim(&_UnstakeRequestsManager.TransactOpts, requests, sourceChainId, destChainId, gasLimit)
}

// Create is a paid mutator transaction binding the contract method 0x891b1a00.
//
// Solidity: function create(address requester, address l2Strategy, uint256 dETHLocked, uint256 ethRequested, uint256 destChainId) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactor) Create(opts *bind.TransactOpts, requester common.Address, l2Strategy common.Address, dETHLocked *big.Int, ethRequested *big.Int, destChainId *big.Int) (*types.Transaction, error) {
	return _UnstakeRequestsManager.contract.Transact(opts, "create", requester, l2Strategy, dETHLocked, ethRequested, destChainId)
}

// Create is a paid mutator transaction binding the contract method 0x891b1a00.
//
// Solidity: function create(address requester, address l2Strategy, uint256 dETHLocked, uint256 ethRequested, uint256 destChainId) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) Create(requester common.Address, l2Strategy common.Address, dETHLocked *big.Int, ethRequested *big.Int, destChainId *big.Int) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.Create(&_UnstakeRequestsManager.TransactOpts, requester, l2Strategy, dETHLocked, ethRequested, destChainId)
}

// Create is a paid mutator transaction binding the contract method 0x891b1a00.
//
// Solidity: function create(address requester, address l2Strategy, uint256 dETHLocked, uint256 ethRequested, uint256 destChainId) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactorSession) Create(requester common.Address, l2Strategy common.Address, dETHLocked *big.Int, ethRequested *big.Int, destChainId *big.Int) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.Create(&_UnstakeRequestsManager.TransactOpts, requester, l2Strategy, dETHLocked, ethRequested, destChainId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UnstakeRequestsManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.GrantRole(&_UnstakeRequestsManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.GrantRole(&_UnstakeRequestsManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x65e80414.
//
// Solidity: function initialize((address,address,address,uint256) init) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactor) Initialize(opts *bind.TransactOpts, init UnstakeRequestsManagerInit) (*types.Transaction, error) {
	return _UnstakeRequestsManager.contract.Transact(opts, "initialize", init)
}

// Initialize is a paid mutator transaction binding the contract method 0x65e80414.
//
// Solidity: function initialize((address,address,address,uint256) init) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) Initialize(init UnstakeRequestsManagerInit) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.Initialize(&_UnstakeRequestsManager.TransactOpts, init)
}

// Initialize is a paid mutator transaction binding the contract method 0x65e80414.
//
// Solidity: function initialize((address,address,address,uint256) init) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactorSession) Initialize(init UnstakeRequestsManagerInit) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.Initialize(&_UnstakeRequestsManager.TransactOpts, init)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _UnstakeRequestsManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.RenounceRole(&_UnstakeRequestsManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.RenounceRole(&_UnstakeRequestsManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UnstakeRequestsManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.RevokeRole(&_UnstakeRequestsManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.RevokeRole(&_UnstakeRequestsManager.TransactOpts, role, account)
}

// SetLocator is a paid mutator transaction binding the contract method 0xa5e84562.
//
// Solidity: function setLocator(address _locator) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactor) SetLocator(opts *bind.TransactOpts, _locator common.Address) (*types.Transaction, error) {
	return _UnstakeRequestsManager.contract.Transact(opts, "setLocator", _locator)
}

// SetLocator is a paid mutator transaction binding the contract method 0xa5e84562.
//
// Solidity: function setLocator(address _locator) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) SetLocator(_locator common.Address) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.SetLocator(&_UnstakeRequestsManager.TransactOpts, _locator)
}

// SetLocator is a paid mutator transaction binding the contract method 0xa5e84562.
//
// Solidity: function setLocator(address _locator) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactorSession) SetLocator(_locator common.Address) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.SetLocator(&_UnstakeRequestsManager.TransactOpts, _locator)
}

// SetNumberOfBlocksToFinalize is a paid mutator transaction binding the contract method 0x28ad3aac.
//
// Solidity: function setNumberOfBlocksToFinalize(uint256 numberOfBlocksToFinalize_) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactor) SetNumberOfBlocksToFinalize(opts *bind.TransactOpts, numberOfBlocksToFinalize_ *big.Int) (*types.Transaction, error) {
	return _UnstakeRequestsManager.contract.Transact(opts, "setNumberOfBlocksToFinalize", numberOfBlocksToFinalize_)
}

// SetNumberOfBlocksToFinalize is a paid mutator transaction binding the contract method 0x28ad3aac.
//
// Solidity: function setNumberOfBlocksToFinalize(uint256 numberOfBlocksToFinalize_) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) SetNumberOfBlocksToFinalize(numberOfBlocksToFinalize_ *big.Int) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.SetNumberOfBlocksToFinalize(&_UnstakeRequestsManager.TransactOpts, numberOfBlocksToFinalize_)
}

// SetNumberOfBlocksToFinalize is a paid mutator transaction binding the contract method 0x28ad3aac.
//
// Solidity: function setNumberOfBlocksToFinalize(uint256 numberOfBlocksToFinalize_) returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactorSession) SetNumberOfBlocksToFinalize(numberOfBlocksToFinalize_ *big.Int) (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.SetNumberOfBlocksToFinalize(&_UnstakeRequestsManager.TransactOpts, numberOfBlocksToFinalize_)
}

// WithdrawAllocatedETHSurplus is a paid mutator transaction binding the contract method 0xb2d42b46.
//
// Solidity: function withdrawAllocatedETHSurplus() returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactor) WithdrawAllocatedETHSurplus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnstakeRequestsManager.contract.Transact(opts, "withdrawAllocatedETHSurplus")
}

// WithdrawAllocatedETHSurplus is a paid mutator transaction binding the contract method 0xb2d42b46.
//
// Solidity: function withdrawAllocatedETHSurplus() returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerSession) WithdrawAllocatedETHSurplus() (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.WithdrawAllocatedETHSurplus(&_UnstakeRequestsManager.TransactOpts)
}

// WithdrawAllocatedETHSurplus is a paid mutator transaction binding the contract method 0xb2d42b46.
//
// Solidity: function withdrawAllocatedETHSurplus() returns()
func (_UnstakeRequestsManager *UnstakeRequestsManagerTransactorSession) WithdrawAllocatedETHSurplus() (*types.Transaction, error) {
	return _UnstakeRequestsManager.Contract.WithdrawAllocatedETHSurplus(&_UnstakeRequestsManager.TransactOpts)
}

// UnstakeRequestsManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerInitializedIterator struct {
	Event *UnstakeRequestsManagerInitialized // Event containing the contract specifics and raw log

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
func (it *UnstakeRequestsManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnstakeRequestsManagerInitialized)
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
		it.Event = new(UnstakeRequestsManagerInitialized)
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
func (it *UnstakeRequestsManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnstakeRequestsManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnstakeRequestsManagerInitialized represents a Initialized event raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*UnstakeRequestsManagerInitializedIterator, error) {

	logs, sub, err := _UnstakeRequestsManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UnstakeRequestsManagerInitializedIterator{contract: _UnstakeRequestsManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UnstakeRequestsManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _UnstakeRequestsManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnstakeRequestsManagerInitialized)
				if err := _UnstakeRequestsManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) ParseInitialized(log types.Log) (*UnstakeRequestsManagerInitialized, error) {
	event := new(UnstakeRequestsManagerInitialized)
	if err := _UnstakeRequestsManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnstakeRequestsManagerProtocolConfigChangedIterator is returned from FilterProtocolConfigChanged and is used to iterate over the raw logs and unpacked data for ProtocolConfigChanged events raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerProtocolConfigChangedIterator struct {
	Event *UnstakeRequestsManagerProtocolConfigChanged // Event containing the contract specifics and raw log

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
func (it *UnstakeRequestsManagerProtocolConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnstakeRequestsManagerProtocolConfigChanged)
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
		it.Event = new(UnstakeRequestsManagerProtocolConfigChanged)
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
func (it *UnstakeRequestsManagerProtocolConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnstakeRequestsManagerProtocolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnstakeRequestsManagerProtocolConfigChanged represents a ProtocolConfigChanged event raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerProtocolConfigChanged struct {
	SetterSelector  [4]byte
	SetterSignature string
	Value           []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProtocolConfigChanged is a free log retrieval operation binding the contract event 0x01d854e8dde9402801a4c6f2840193465752abfad61e0bb7c4258d526ae42e74.
//
// Solidity: event ProtocolConfigChanged(bytes4 indexed setterSelector, string setterSignature, bytes value)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) FilterProtocolConfigChanged(opts *bind.FilterOpts, setterSelector [][4]byte) (*UnstakeRequestsManagerProtocolConfigChangedIterator, error) {

	var setterSelectorRule []interface{}
	for _, setterSelectorItem := range setterSelector {
		setterSelectorRule = append(setterSelectorRule, setterSelectorItem)
	}

	logs, sub, err := _UnstakeRequestsManager.contract.FilterLogs(opts, "ProtocolConfigChanged", setterSelectorRule)
	if err != nil {
		return nil, err
	}
	return &UnstakeRequestsManagerProtocolConfigChangedIterator{contract: _UnstakeRequestsManager.contract, event: "ProtocolConfigChanged", logs: logs, sub: sub}, nil
}

// WatchProtocolConfigChanged is a free log subscription operation binding the contract event 0x01d854e8dde9402801a4c6f2840193465752abfad61e0bb7c4258d526ae42e74.
//
// Solidity: event ProtocolConfigChanged(bytes4 indexed setterSelector, string setterSignature, bytes value)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) WatchProtocolConfigChanged(opts *bind.WatchOpts, sink chan<- *UnstakeRequestsManagerProtocolConfigChanged, setterSelector [][4]byte) (event.Subscription, error) {

	var setterSelectorRule []interface{}
	for _, setterSelectorItem := range setterSelector {
		setterSelectorRule = append(setterSelectorRule, setterSelectorItem)
	}

	logs, sub, err := _UnstakeRequestsManager.contract.WatchLogs(opts, "ProtocolConfigChanged", setterSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnstakeRequestsManagerProtocolConfigChanged)
				if err := _UnstakeRequestsManager.contract.UnpackLog(event, "ProtocolConfigChanged", log); err != nil {
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

// ParseProtocolConfigChanged is a log parse operation binding the contract event 0x01d854e8dde9402801a4c6f2840193465752abfad61e0bb7c4258d526ae42e74.
//
// Solidity: event ProtocolConfigChanged(bytes4 indexed setterSelector, string setterSignature, bytes value)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) ParseProtocolConfigChanged(log types.Log) (*UnstakeRequestsManagerProtocolConfigChanged, error) {
	event := new(UnstakeRequestsManagerProtocolConfigChanged)
	if err := _UnstakeRequestsManager.contract.UnpackLog(event, "ProtocolConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnstakeRequestsManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerRoleAdminChangedIterator struct {
	Event *UnstakeRequestsManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *UnstakeRequestsManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnstakeRequestsManagerRoleAdminChanged)
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
		it.Event = new(UnstakeRequestsManagerRoleAdminChanged)
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
func (it *UnstakeRequestsManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnstakeRequestsManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnstakeRequestsManagerRoleAdminChanged represents a RoleAdminChanged event raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*UnstakeRequestsManagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _UnstakeRequestsManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &UnstakeRequestsManagerRoleAdminChangedIterator{contract: _UnstakeRequestsManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *UnstakeRequestsManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _UnstakeRequestsManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnstakeRequestsManagerRoleAdminChanged)
				if err := _UnstakeRequestsManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) ParseRoleAdminChanged(log types.Log) (*UnstakeRequestsManagerRoleAdminChanged, error) {
	event := new(UnstakeRequestsManagerRoleAdminChanged)
	if err := _UnstakeRequestsManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnstakeRequestsManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerRoleGrantedIterator struct {
	Event *UnstakeRequestsManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *UnstakeRequestsManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnstakeRequestsManagerRoleGranted)
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
		it.Event = new(UnstakeRequestsManagerRoleGranted)
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
func (it *UnstakeRequestsManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnstakeRequestsManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnstakeRequestsManagerRoleGranted represents a RoleGranted event raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UnstakeRequestsManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _UnstakeRequestsManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UnstakeRequestsManagerRoleGrantedIterator{contract: _UnstakeRequestsManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *UnstakeRequestsManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _UnstakeRequestsManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnstakeRequestsManagerRoleGranted)
				if err := _UnstakeRequestsManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) ParseRoleGranted(log types.Log) (*UnstakeRequestsManagerRoleGranted, error) {
	event := new(UnstakeRequestsManagerRoleGranted)
	if err := _UnstakeRequestsManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnstakeRequestsManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerRoleRevokedIterator struct {
	Event *UnstakeRequestsManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *UnstakeRequestsManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnstakeRequestsManagerRoleRevoked)
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
		it.Event = new(UnstakeRequestsManagerRoleRevoked)
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
func (it *UnstakeRequestsManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnstakeRequestsManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnstakeRequestsManagerRoleRevoked represents a RoleRevoked event raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UnstakeRequestsManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _UnstakeRequestsManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UnstakeRequestsManagerRoleRevokedIterator{contract: _UnstakeRequestsManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *UnstakeRequestsManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _UnstakeRequestsManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnstakeRequestsManagerRoleRevoked)
				if err := _UnstakeRequestsManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) ParseRoleRevoked(log types.Log) (*UnstakeRequestsManagerRoleRevoked, error) {
	event := new(UnstakeRequestsManagerRoleRevoked)
	if err := _UnstakeRequestsManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnstakeRequestsManagerUnstakeRequestCancelledIterator is returned from FilterUnstakeRequestCancelled and is used to iterate over the raw logs and unpacked data for UnstakeRequestCancelled events raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerUnstakeRequestCancelledIterator struct {
	Event *UnstakeRequestsManagerUnstakeRequestCancelled // Event containing the contract specifics and raw log

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
func (it *UnstakeRequestsManagerUnstakeRequestCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnstakeRequestsManagerUnstakeRequestCancelled)
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
		it.Event = new(UnstakeRequestsManagerUnstakeRequestCancelled)
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
func (it *UnstakeRequestsManagerUnstakeRequestCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnstakeRequestsManagerUnstakeRequestCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnstakeRequestsManagerUnstakeRequestCancelled represents a UnstakeRequestCancelled event raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerUnstakeRequestCancelled struct {
	Id                     *big.Int
	Requester              common.Address
	DETHLocked             *big.Int
	EthRequested           *big.Int
	CumulativeETHRequested *big.Int
	BlockNumber            *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUnstakeRequestCancelled is a free log retrieval operation binding the contract event 0xf8d5df096390c80d709314ba3a8e55a91854bab1fbf8cce78081bdd48eb49aed.
//
// Solidity: event UnstakeRequestCancelled(uint256 indexed id, address indexed requester, uint256 dETHLocked, uint256 ethRequested, uint256 cumulativeETHRequested, uint256 blockNumber)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) FilterUnstakeRequestCancelled(opts *bind.FilterOpts, id []*big.Int, requester []common.Address) (*UnstakeRequestsManagerUnstakeRequestCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _UnstakeRequestsManager.contract.FilterLogs(opts, "UnstakeRequestCancelled", idRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &UnstakeRequestsManagerUnstakeRequestCancelledIterator{contract: _UnstakeRequestsManager.contract, event: "UnstakeRequestCancelled", logs: logs, sub: sub}, nil
}

// WatchUnstakeRequestCancelled is a free log subscription operation binding the contract event 0xf8d5df096390c80d709314ba3a8e55a91854bab1fbf8cce78081bdd48eb49aed.
//
// Solidity: event UnstakeRequestCancelled(uint256 indexed id, address indexed requester, uint256 dETHLocked, uint256 ethRequested, uint256 cumulativeETHRequested, uint256 blockNumber)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) WatchUnstakeRequestCancelled(opts *bind.WatchOpts, sink chan<- *UnstakeRequestsManagerUnstakeRequestCancelled, id []*big.Int, requester []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _UnstakeRequestsManager.contract.WatchLogs(opts, "UnstakeRequestCancelled", idRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnstakeRequestsManagerUnstakeRequestCancelled)
				if err := _UnstakeRequestsManager.contract.UnpackLog(event, "UnstakeRequestCancelled", log); err != nil {
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

// ParseUnstakeRequestCancelled is a log parse operation binding the contract event 0xf8d5df096390c80d709314ba3a8e55a91854bab1fbf8cce78081bdd48eb49aed.
//
// Solidity: event UnstakeRequestCancelled(uint256 indexed id, address indexed requester, uint256 dETHLocked, uint256 ethRequested, uint256 cumulativeETHRequested, uint256 blockNumber)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) ParseUnstakeRequestCancelled(log types.Log) (*UnstakeRequestsManagerUnstakeRequestCancelled, error) {
	event := new(UnstakeRequestsManagerUnstakeRequestCancelled)
	if err := _UnstakeRequestsManager.contract.UnpackLog(event, "UnstakeRequestCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnstakeRequestsManagerUnstakeRequestClaimedIterator is returned from FilterUnstakeRequestClaimed and is used to iterate over the raw logs and unpacked data for UnstakeRequestClaimed events raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerUnstakeRequestClaimedIterator struct {
	Event *UnstakeRequestsManagerUnstakeRequestClaimed // Event containing the contract specifics and raw log

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
func (it *UnstakeRequestsManagerUnstakeRequestClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnstakeRequestsManagerUnstakeRequestClaimed)
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
		it.Event = new(UnstakeRequestsManagerUnstakeRequestClaimed)
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
func (it *UnstakeRequestsManagerUnstakeRequestClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnstakeRequestsManagerUnstakeRequestClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnstakeRequestsManagerUnstakeRequestClaimed represents a UnstakeRequestClaimed event raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerUnstakeRequestClaimed struct {
	L2strategy          common.Address
	EthRequested        *big.Int
	DETHLocked          *big.Int
	DestChainId         *big.Int
	CsBlockNumber       *big.Int
	BridgeAddress       common.Address
	UnStakeMessageNonce *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUnstakeRequestClaimed is a free log retrieval operation binding the contract event 0x4d88ab48fd3a514a018a38f136c8a4bf616fb79e3e2f25d70c4e76a181cee2cf.
//
// Solidity: event UnstakeRequestClaimed(address indexed l2strategy, uint256 ethRequested, uint256 dETHLocked, uint256 indexed destChainId, uint256 indexed csBlockNumber, address bridgeAddress, uint256 unStakeMessageNonce)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) FilterUnstakeRequestClaimed(opts *bind.FilterOpts, l2strategy []common.Address, destChainId []*big.Int, csBlockNumber []*big.Int) (*UnstakeRequestsManagerUnstakeRequestClaimedIterator, error) {

	var l2strategyRule []interface{}
	for _, l2strategyItem := range l2strategy {
		l2strategyRule = append(l2strategyRule, l2strategyItem)
	}

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var csBlockNumberRule []interface{}
	for _, csBlockNumberItem := range csBlockNumber {
		csBlockNumberRule = append(csBlockNumberRule, csBlockNumberItem)
	}

	logs, sub, err := _UnstakeRequestsManager.contract.FilterLogs(opts, "UnstakeRequestClaimed", l2strategyRule, destChainIdRule, csBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &UnstakeRequestsManagerUnstakeRequestClaimedIterator{contract: _UnstakeRequestsManager.contract, event: "UnstakeRequestClaimed", logs: logs, sub: sub}, nil
}

// WatchUnstakeRequestClaimed is a free log subscription operation binding the contract event 0x4d88ab48fd3a514a018a38f136c8a4bf616fb79e3e2f25d70c4e76a181cee2cf.
//
// Solidity: event UnstakeRequestClaimed(address indexed l2strategy, uint256 ethRequested, uint256 dETHLocked, uint256 indexed destChainId, uint256 indexed csBlockNumber, address bridgeAddress, uint256 unStakeMessageNonce)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) WatchUnstakeRequestClaimed(opts *bind.WatchOpts, sink chan<- *UnstakeRequestsManagerUnstakeRequestClaimed, l2strategy []common.Address, destChainId []*big.Int, csBlockNumber []*big.Int) (event.Subscription, error) {

	var l2strategyRule []interface{}
	for _, l2strategyItem := range l2strategy {
		l2strategyRule = append(l2strategyRule, l2strategyItem)
	}

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var csBlockNumberRule []interface{}
	for _, csBlockNumberItem := range csBlockNumber {
		csBlockNumberRule = append(csBlockNumberRule, csBlockNumberItem)
	}

	logs, sub, err := _UnstakeRequestsManager.contract.WatchLogs(opts, "UnstakeRequestClaimed", l2strategyRule, destChainIdRule, csBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnstakeRequestsManagerUnstakeRequestClaimed)
				if err := _UnstakeRequestsManager.contract.UnpackLog(event, "UnstakeRequestClaimed", log); err != nil {
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

// ParseUnstakeRequestClaimed is a log parse operation binding the contract event 0x4d88ab48fd3a514a018a38f136c8a4bf616fb79e3e2f25d70c4e76a181cee2cf.
//
// Solidity: event UnstakeRequestClaimed(address indexed l2strategy, uint256 ethRequested, uint256 dETHLocked, uint256 indexed destChainId, uint256 indexed csBlockNumber, address bridgeAddress, uint256 unStakeMessageNonce)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) ParseUnstakeRequestClaimed(log types.Log) (*UnstakeRequestsManagerUnstakeRequestClaimed, error) {
	event := new(UnstakeRequestsManagerUnstakeRequestClaimed)
	if err := _UnstakeRequestsManager.contract.UnpackLog(event, "UnstakeRequestClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnstakeRequestsManagerUnstakeRequestCreatedIterator is returned from FilterUnstakeRequestCreated and is used to iterate over the raw logs and unpacked data for UnstakeRequestCreated events raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerUnstakeRequestCreatedIterator struct {
	Event *UnstakeRequestsManagerUnstakeRequestCreated // Event containing the contract specifics and raw log

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
func (it *UnstakeRequestsManagerUnstakeRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnstakeRequestsManagerUnstakeRequestCreated)
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
		it.Event = new(UnstakeRequestsManagerUnstakeRequestCreated)
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
func (it *UnstakeRequestsManagerUnstakeRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnstakeRequestsManagerUnstakeRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnstakeRequestsManagerUnstakeRequestCreated represents a UnstakeRequestCreated event raised by the UnstakeRequestsManager contract.
type UnstakeRequestsManagerUnstakeRequestCreated struct {
	Requester              common.Address
	Strategy               common.Address
	DETHLocked             *big.Int
	EthRequested           *big.Int
	CumulativeETHRequested *big.Int
	BlockNumber            *big.Int
	DestChainId            *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUnstakeRequestCreated is a free log retrieval operation binding the contract event 0x319883689186488248969ba55a688029c4bb01b7e60c9b4b9914f475f60ad2e4.
//
// Solidity: event UnstakeRequestCreated(address indexed requester, address indexed strategy, uint256 dETHLocked, uint256 ethRequested, uint256 cumulativeETHRequested, uint256 blockNumber, uint256 destChainId)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) FilterUnstakeRequestCreated(opts *bind.FilterOpts, requester []common.Address, strategy []common.Address) (*UnstakeRequestsManagerUnstakeRequestCreatedIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _UnstakeRequestsManager.contract.FilterLogs(opts, "UnstakeRequestCreated", requesterRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &UnstakeRequestsManagerUnstakeRequestCreatedIterator{contract: _UnstakeRequestsManager.contract, event: "UnstakeRequestCreated", logs: logs, sub: sub}, nil
}

// WatchUnstakeRequestCreated is a free log subscription operation binding the contract event 0x319883689186488248969ba55a688029c4bb01b7e60c9b4b9914f475f60ad2e4.
//
// Solidity: event UnstakeRequestCreated(address indexed requester, address indexed strategy, uint256 dETHLocked, uint256 ethRequested, uint256 cumulativeETHRequested, uint256 blockNumber, uint256 destChainId)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) WatchUnstakeRequestCreated(opts *bind.WatchOpts, sink chan<- *UnstakeRequestsManagerUnstakeRequestCreated, requester []common.Address, strategy []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _UnstakeRequestsManager.contract.WatchLogs(opts, "UnstakeRequestCreated", requesterRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnstakeRequestsManagerUnstakeRequestCreated)
				if err := _UnstakeRequestsManager.contract.UnpackLog(event, "UnstakeRequestCreated", log); err != nil {
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

// ParseUnstakeRequestCreated is a log parse operation binding the contract event 0x319883689186488248969ba55a688029c4bb01b7e60c9b4b9914f475f60ad2e4.
//
// Solidity: event UnstakeRequestCreated(address indexed requester, address indexed strategy, uint256 dETHLocked, uint256 ethRequested, uint256 cumulativeETHRequested, uint256 blockNumber, uint256 destChainId)
func (_UnstakeRequestsManager *UnstakeRequestsManagerFilterer) ParseUnstakeRequestCreated(log types.Log) (*UnstakeRequestsManagerUnstakeRequestCreated, error) {
	event := new(UnstakeRequestsManagerUnstakeRequestCreated)
	if err := _UnstakeRequestsManager.contract.UnpackLog(event, "UnstakeRequestCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
