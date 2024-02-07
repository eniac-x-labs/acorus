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

// IL1PoolManagerPool is an auto generated low-level Go binding around an user-defined struct.
type IL1PoolManagerPool struct {
	StartTimestamp  uint32
	EndTimestamp    uint32
	Token           common.Address
	TotalAmount     *big.Int
	TotalFee        *big.Int
	TotalFeeClaimed *big.Int
	IsCompleted     bool
}

// IL1PoolManagerMetaData contains all meta data concerning the IL1PoolManager contract.
var IL1PoolManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"providedAmount\",\"type\":\"uint256\"}],\"name\":\"LessThanMinStakeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LessThanZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PoolIndex\",\"type\":\"uint256\"}],\"name\":\"NewPoolIsNotCreate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReward\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"PoolLength\",\"type\":\"uint256\"}],\"name\":\"OutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolIndex\",\"type\":\"uint256\"}],\"name\":\"PoolIsCompleted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"name\":\"TokenIsAlreadySupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenIsNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endPoolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Reward\",\"type\":\"uint256\"}],\"name\":\"ClaimReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolIndex\",\"type\":\"uint256\"}],\"name\":\"CompletePoolEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetMinStakeAmountEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSupport\",\"type\":\"bool\"}],\"name\":\"SetSupportTokenEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingETHEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingWETHEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StarkingERC20Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Blockchain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferAssertTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Reward\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ClaimAllReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"IsCompleted\",\"type\":\"bool\"}],\"internalType\":\"structIL1PoolManager.Pool[]\",\"name\":\"CompletePools\",\"type\":\"tuple[]\"}],\"name\":\"CompletePoolAndNew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DepositAndStaking\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DepositAndStakingERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DepositAndStakingETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositAndStakingWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isSupport\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"startTimes\",\"type\":\"uint32\"}],\"name\":\"SetSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"Blockchain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TransferAssertToBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setMinStakeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IL1PoolManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IL1PoolManagerMetaData.ABI instead.
var IL1PoolManagerABI = IL1PoolManagerMetaData.ABI

// IL1PoolManager is an auto generated Go binding around an Ethereum contract.
type IL1PoolManager struct {
	IL1PoolManagerCaller     // Read-only binding to the contract
	IL1PoolManagerTransactor // Write-only binding to the contract
	IL1PoolManagerFilterer   // Log filterer for contract events
}

// IL1PoolManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL1PoolManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1PoolManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL1PoolManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1PoolManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL1PoolManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1PoolManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL1PoolManagerSession struct {
	Contract     *IL1PoolManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL1PoolManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL1PoolManagerCallerSession struct {
	Contract *IL1PoolManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IL1PoolManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL1PoolManagerTransactorSession struct {
	Contract     *IL1PoolManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IL1PoolManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL1PoolManagerRaw struct {
	Contract *IL1PoolManager // Generic contract binding to access the raw methods on
}

// IL1PoolManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL1PoolManagerCallerRaw struct {
	Contract *IL1PoolManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IL1PoolManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL1PoolManagerTransactorRaw struct {
	Contract *IL1PoolManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL1PoolManager creates a new instance of IL1PoolManager, bound to a specific deployed contract.
func NewIL1PoolManager(address common.Address, backend bind.ContractBackend) (*IL1PoolManager, error) {
	contract, err := bindIL1PoolManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL1PoolManager{IL1PoolManagerCaller: IL1PoolManagerCaller{contract: contract}, IL1PoolManagerTransactor: IL1PoolManagerTransactor{contract: contract}, IL1PoolManagerFilterer: IL1PoolManagerFilterer{contract: contract}}, nil
}

// NewIL1PoolManagerCaller creates a new read-only instance of IL1PoolManager, bound to a specific deployed contract.
func NewIL1PoolManagerCaller(address common.Address, caller bind.ContractCaller) (*IL1PoolManagerCaller, error) {
	contract, err := bindIL1PoolManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL1PoolManagerCaller{contract: contract}, nil
}

// NewIL1PoolManagerTransactor creates a new write-only instance of IL1PoolManager, bound to a specific deployed contract.
func NewIL1PoolManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IL1PoolManagerTransactor, error) {
	contract, err := bindIL1PoolManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL1PoolManagerTransactor{contract: contract}, nil
}

// NewIL1PoolManagerFilterer creates a new log filterer instance of IL1PoolManager, bound to a specific deployed contract.
func NewIL1PoolManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IL1PoolManagerFilterer, error) {
	contract, err := bindIL1PoolManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL1PoolManagerFilterer{contract: contract}, nil
}

// bindIL1PoolManager binds a generic wrapper to an already deployed contract.
func bindIL1PoolManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IL1PoolManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1PoolManager *IL1PoolManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1PoolManager.Contract.IL1PoolManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1PoolManager *IL1PoolManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.IL1PoolManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1PoolManager *IL1PoolManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.IL1PoolManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1PoolManager *IL1PoolManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1PoolManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1PoolManager *IL1PoolManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1PoolManager *IL1PoolManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.contract.Transact(opts, method, params...)
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x2c37388b.
//
// Solidity: function ClaimAllReward() returns()
func (_IL1PoolManager *IL1PoolManagerTransactor) ClaimAllReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1PoolManager.contract.Transact(opts, "ClaimAllReward")
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x2c37388b.
//
// Solidity: function ClaimAllReward() returns()
func (_IL1PoolManager *IL1PoolManagerSession) ClaimAllReward() (*types.Transaction, error) {
	return _IL1PoolManager.Contract.ClaimAllReward(&_IL1PoolManager.TransactOpts)
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x2c37388b.
//
// Solidity: function ClaimAllReward() returns()
func (_IL1PoolManager *IL1PoolManagerTransactorSession) ClaimAllReward() (*types.Transaction, error) {
	return _IL1PoolManager.Contract.ClaimAllReward(&_IL1PoolManager.TransactOpts)
}

// CompletePoolAndNew is a paid mutator transaction binding the contract method 0x3338562c.
//
// Solidity: function CompletePoolAndNew((uint32,uint32,address,uint256,uint256,uint256,bool)[] CompletePools) payable returns()
func (_IL1PoolManager *IL1PoolManagerTransactor) CompletePoolAndNew(opts *bind.TransactOpts, CompletePools []IL1PoolManagerPool) (*types.Transaction, error) {
	return _IL1PoolManager.contract.Transact(opts, "CompletePoolAndNew", CompletePools)
}

// CompletePoolAndNew is a paid mutator transaction binding the contract method 0x3338562c.
//
// Solidity: function CompletePoolAndNew((uint32,uint32,address,uint256,uint256,uint256,bool)[] CompletePools) payable returns()
func (_IL1PoolManager *IL1PoolManagerSession) CompletePoolAndNew(CompletePools []IL1PoolManagerPool) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.CompletePoolAndNew(&_IL1PoolManager.TransactOpts, CompletePools)
}

// CompletePoolAndNew is a paid mutator transaction binding the contract method 0x3338562c.
//
// Solidity: function CompletePoolAndNew((uint32,uint32,address,uint256,uint256,uint256,bool)[] CompletePools) payable returns()
func (_IL1PoolManager *IL1PoolManagerTransactorSession) CompletePoolAndNew(CompletePools []IL1PoolManagerPool) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.CompletePoolAndNew(&_IL1PoolManager.TransactOpts, CompletePools)
}

// DepositAndStaking is a paid mutator transaction binding the contract method 0x9ab0b652.
//
// Solidity: function DepositAndStaking(address _token, uint256 _amount) payable returns()
func (_IL1PoolManager *IL1PoolManagerTransactor) DepositAndStaking(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1PoolManager.contract.Transact(opts, "DepositAndStaking", _token, _amount)
}

// DepositAndStaking is a paid mutator transaction binding the contract method 0x9ab0b652.
//
// Solidity: function DepositAndStaking(address _token, uint256 _amount) payable returns()
func (_IL1PoolManager *IL1PoolManagerSession) DepositAndStaking(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.DepositAndStaking(&_IL1PoolManager.TransactOpts, _token, _amount)
}

// DepositAndStaking is a paid mutator transaction binding the contract method 0x9ab0b652.
//
// Solidity: function DepositAndStaking(address _token, uint256 _amount) payable returns()
func (_IL1PoolManager *IL1PoolManagerTransactorSession) DepositAndStaking(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.DepositAndStaking(&_IL1PoolManager.TransactOpts, _token, _amount)
}

// DepositAndStakingERC20 is a paid mutator transaction binding the contract method 0x4dfaef84.
//
// Solidity: function DepositAndStakingERC20(address _token, uint256 _amount) returns()
func (_IL1PoolManager *IL1PoolManagerTransactor) DepositAndStakingERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1PoolManager.contract.Transact(opts, "DepositAndStakingERC20", _token, _amount)
}

// DepositAndStakingERC20 is a paid mutator transaction binding the contract method 0x4dfaef84.
//
// Solidity: function DepositAndStakingERC20(address _token, uint256 _amount) returns()
func (_IL1PoolManager *IL1PoolManagerSession) DepositAndStakingERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.DepositAndStakingERC20(&_IL1PoolManager.TransactOpts, _token, _amount)
}

// DepositAndStakingERC20 is a paid mutator transaction binding the contract method 0x4dfaef84.
//
// Solidity: function DepositAndStakingERC20(address _token, uint256 _amount) returns()
func (_IL1PoolManager *IL1PoolManagerTransactorSession) DepositAndStakingERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.DepositAndStakingERC20(&_IL1PoolManager.TransactOpts, _token, _amount)
}

// DepositAndStakingETH is a paid mutator transaction binding the contract method 0x96f984e6.
//
// Solidity: function DepositAndStakingETH() payable returns()
func (_IL1PoolManager *IL1PoolManagerTransactor) DepositAndStakingETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1PoolManager.contract.Transact(opts, "DepositAndStakingETH")
}

// DepositAndStakingETH is a paid mutator transaction binding the contract method 0x96f984e6.
//
// Solidity: function DepositAndStakingETH() payable returns()
func (_IL1PoolManager *IL1PoolManagerSession) DepositAndStakingETH() (*types.Transaction, error) {
	return _IL1PoolManager.Contract.DepositAndStakingETH(&_IL1PoolManager.TransactOpts)
}

// DepositAndStakingETH is a paid mutator transaction binding the contract method 0x96f984e6.
//
// Solidity: function DepositAndStakingETH() payable returns()
func (_IL1PoolManager *IL1PoolManagerTransactorSession) DepositAndStakingETH() (*types.Transaction, error) {
	return _IL1PoolManager.Contract.DepositAndStakingETH(&_IL1PoolManager.TransactOpts)
}

// DepositAndStakingWETH is a paid mutator transaction binding the contract method 0xbc42493d.
//
// Solidity: function DepositAndStakingWETH(uint256 amount) returns()
func (_IL1PoolManager *IL1PoolManagerTransactor) DepositAndStakingWETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IL1PoolManager.contract.Transact(opts, "DepositAndStakingWETH", amount)
}

// DepositAndStakingWETH is a paid mutator transaction binding the contract method 0xbc42493d.
//
// Solidity: function DepositAndStakingWETH(uint256 amount) returns()
func (_IL1PoolManager *IL1PoolManagerSession) DepositAndStakingWETH(amount *big.Int) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.DepositAndStakingWETH(&_IL1PoolManager.TransactOpts, amount)
}

// DepositAndStakingWETH is a paid mutator transaction binding the contract method 0xbc42493d.
//
// Solidity: function DepositAndStakingWETH(uint256 amount) returns()
func (_IL1PoolManager *IL1PoolManagerTransactorSession) DepositAndStakingWETH(amount *big.Int) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.DepositAndStakingWETH(&_IL1PoolManager.TransactOpts, amount)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0xfaeb96a5.
//
// Solidity: function SetSupportToken(address _token, bool _isSupport, uint32 startTimes) returns()
func (_IL1PoolManager *IL1PoolManagerTransactor) SetSupportToken(opts *bind.TransactOpts, _token common.Address, _isSupport bool, startTimes uint32) (*types.Transaction, error) {
	return _IL1PoolManager.contract.Transact(opts, "SetSupportToken", _token, _isSupport, startTimes)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0xfaeb96a5.
//
// Solidity: function SetSupportToken(address _token, bool _isSupport, uint32 startTimes) returns()
func (_IL1PoolManager *IL1PoolManagerSession) SetSupportToken(_token common.Address, _isSupport bool, startTimes uint32) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.SetSupportToken(&_IL1PoolManager.TransactOpts, _token, _isSupport, startTimes)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0xfaeb96a5.
//
// Solidity: function SetSupportToken(address _token, bool _isSupport, uint32 startTimes) returns()
func (_IL1PoolManager *IL1PoolManagerTransactorSession) SetSupportToken(_token common.Address, _isSupport bool, startTimes uint32) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.SetSupportToken(&_IL1PoolManager.TransactOpts, _token, _isSupport, startTimes)
}

// TransferAssertToBridge is a paid mutator transaction binding the contract method 0xf91fa9a8.
//
// Solidity: function TransferAssertToBridge(uint256 Blockchain, address _token, address _to, uint256 _amount) returns()
func (_IL1PoolManager *IL1PoolManagerTransactor) TransferAssertToBridge(opts *bind.TransactOpts, Blockchain *big.Int, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1PoolManager.contract.Transact(opts, "TransferAssertToBridge", Blockchain, _token, _to, _amount)
}

// TransferAssertToBridge is a paid mutator transaction binding the contract method 0xf91fa9a8.
//
// Solidity: function TransferAssertToBridge(uint256 Blockchain, address _token, address _to, uint256 _amount) returns()
func (_IL1PoolManager *IL1PoolManagerSession) TransferAssertToBridge(Blockchain *big.Int, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.TransferAssertToBridge(&_IL1PoolManager.TransactOpts, Blockchain, _token, _to, _amount)
}

// TransferAssertToBridge is a paid mutator transaction binding the contract method 0xf91fa9a8.
//
// Solidity: function TransferAssertToBridge(uint256 Blockchain, address _token, address _to, uint256 _amount) returns()
func (_IL1PoolManager *IL1PoolManagerTransactorSession) TransferAssertToBridge(Blockchain *big.Int, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.TransferAssertToBridge(&_IL1PoolManager.TransactOpts, Blockchain, _token, _to, _amount)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xcb314fab.
//
// Solidity: function setMinStakeAmount(address _token, uint256 _amount) returns()
func (_IL1PoolManager *IL1PoolManagerTransactor) SetMinStakeAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1PoolManager.contract.Transact(opts, "setMinStakeAmount", _token, _amount)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xcb314fab.
//
// Solidity: function setMinStakeAmount(address _token, uint256 _amount) returns()
func (_IL1PoolManager *IL1PoolManagerSession) SetMinStakeAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.SetMinStakeAmount(&_IL1PoolManager.TransactOpts, _token, _amount)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xcb314fab.
//
// Solidity: function setMinStakeAmount(address _token, uint256 _amount) returns()
func (_IL1PoolManager *IL1PoolManagerTransactorSession) SetMinStakeAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1PoolManager.Contract.SetMinStakeAmount(&_IL1PoolManager.TransactOpts, _token, _amount)
}

// IL1PoolManagerClaimEventIterator is returned from FilterClaimEvent and is used to iterate over the raw logs and unpacked data for ClaimEvent events raised by the IL1PoolManager contract.
type IL1PoolManagerClaimEventIterator struct {
	Event *IL1PoolManagerClaimEvent // Event containing the contract specifics and raw log

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
func (it *IL1PoolManagerClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1PoolManagerClaimEvent)
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
		it.Event = new(IL1PoolManagerClaimEvent)
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
func (it *IL1PoolManagerClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1PoolManagerClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1PoolManagerClaimEvent represents a ClaimEvent event raised by the IL1PoolManager contract.
type IL1PoolManagerClaimEvent struct {
	User        common.Address
	StartPoolId *big.Int
	EndPoolId   *big.Int
	Token       common.Address
	Amount      *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaimEvent is a free log retrieval operation binding the contract event 0x991c329471ab5230e2301ee30d0d6fba5906e32411d5d154d5a8c278d021a2ab.
//
// Solidity: event ClaimEvent(address indexed user, uint256 startPoolId, uint256 endPoolId, address indexed token, uint256 amount, uint256 fee)
func (_IL1PoolManager *IL1PoolManagerFilterer) FilterClaimEvent(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*IL1PoolManagerClaimEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IL1PoolManager.contract.FilterLogs(opts, "ClaimEvent", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IL1PoolManagerClaimEventIterator{contract: _IL1PoolManager.contract, event: "ClaimEvent", logs: logs, sub: sub}, nil
}

// WatchClaimEvent is a free log subscription operation binding the contract event 0x991c329471ab5230e2301ee30d0d6fba5906e32411d5d154d5a8c278d021a2ab.
//
// Solidity: event ClaimEvent(address indexed user, uint256 startPoolId, uint256 endPoolId, address indexed token, uint256 amount, uint256 fee)
func (_IL1PoolManager *IL1PoolManagerFilterer) WatchClaimEvent(opts *bind.WatchOpts, sink chan<- *IL1PoolManagerClaimEvent, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IL1PoolManager.contract.WatchLogs(opts, "ClaimEvent", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1PoolManagerClaimEvent)
				if err := _IL1PoolManager.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
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

// ParseClaimEvent is a log parse operation binding the contract event 0x991c329471ab5230e2301ee30d0d6fba5906e32411d5d154d5a8c278d021a2ab.
//
// Solidity: event ClaimEvent(address indexed user, uint256 startPoolId, uint256 endPoolId, address indexed token, uint256 amount, uint256 fee)
func (_IL1PoolManager *IL1PoolManagerFilterer) ParseClaimEvent(log types.Log) (*IL1PoolManagerClaimEvent, error) {
	event := new(IL1PoolManagerClaimEvent)
	if err := _IL1PoolManager.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1PoolManagerClaimRewardIterator is returned from FilterClaimReward and is used to iterate over the raw logs and unpacked data for ClaimReward events raised by the IL1PoolManager contract.
type IL1PoolManagerClaimRewardIterator struct {
	Event *IL1PoolManagerClaimReward // Event containing the contract specifics and raw log

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
func (it *IL1PoolManagerClaimRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1PoolManagerClaimReward)
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
		it.Event = new(IL1PoolManagerClaimReward)
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
func (it *IL1PoolManagerClaimRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1PoolManagerClaimRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1PoolManagerClaimReward represents a ClaimReward event raised by the IL1PoolManager contract.
type IL1PoolManagerClaimReward struct {
	User        common.Address
	StartPoolId *big.Int
	EndPoolId   *big.Int
	Token       common.Address
	Reward      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaimReward is a free log retrieval operation binding the contract event 0x7e6fedb6f824b2112ac0626ced32f6be6324d554cc03efd112aa021ee4d740ac.
//
// Solidity: event ClaimReward(address _user, uint256 startPoolId, uint256 EndPoolId, address _token, uint256 Reward)
func (_IL1PoolManager *IL1PoolManagerFilterer) FilterClaimReward(opts *bind.FilterOpts) (*IL1PoolManagerClaimRewardIterator, error) {

	logs, sub, err := _IL1PoolManager.contract.FilterLogs(opts, "ClaimReward")
	if err != nil {
		return nil, err
	}
	return &IL1PoolManagerClaimRewardIterator{contract: _IL1PoolManager.contract, event: "ClaimReward", logs: logs, sub: sub}, nil
}

// WatchClaimReward is a free log subscription operation binding the contract event 0x7e6fedb6f824b2112ac0626ced32f6be6324d554cc03efd112aa021ee4d740ac.
//
// Solidity: event ClaimReward(address _user, uint256 startPoolId, uint256 EndPoolId, address _token, uint256 Reward)
func (_IL1PoolManager *IL1PoolManagerFilterer) WatchClaimReward(opts *bind.WatchOpts, sink chan<- *IL1PoolManagerClaimReward) (event.Subscription, error) {

	logs, sub, err := _IL1PoolManager.contract.WatchLogs(opts, "ClaimReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1PoolManagerClaimReward)
				if err := _IL1PoolManager.contract.UnpackLog(event, "ClaimReward", log); err != nil {
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

// ParseClaimReward is a log parse operation binding the contract event 0x7e6fedb6f824b2112ac0626ced32f6be6324d554cc03efd112aa021ee4d740ac.
//
// Solidity: event ClaimReward(address _user, uint256 startPoolId, uint256 EndPoolId, address _token, uint256 Reward)
func (_IL1PoolManager *IL1PoolManagerFilterer) ParseClaimReward(log types.Log) (*IL1PoolManagerClaimReward, error) {
	event := new(IL1PoolManagerClaimReward)
	if err := _IL1PoolManager.contract.UnpackLog(event, "ClaimReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1PoolManagerCompletePoolEventIterator is returned from FilterCompletePoolEvent and is used to iterate over the raw logs and unpacked data for CompletePoolEvent events raised by the IL1PoolManager contract.
type IL1PoolManagerCompletePoolEventIterator struct {
	Event *IL1PoolManagerCompletePoolEvent // Event containing the contract specifics and raw log

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
func (it *IL1PoolManagerCompletePoolEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1PoolManagerCompletePoolEvent)
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
		it.Event = new(IL1PoolManagerCompletePoolEvent)
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
func (it *IL1PoolManagerCompletePoolEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1PoolManagerCompletePoolEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1PoolManagerCompletePoolEvent represents a CompletePoolEvent event raised by the IL1PoolManager contract.
type IL1PoolManagerCompletePoolEvent struct {
	Token     common.Address
	PoolIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCompletePoolEvent is a free log retrieval operation binding the contract event 0xb6f449f07ceaf55392c9899e0797c6529908ae827c2d498c682e90d42c241167.
//
// Solidity: event CompletePoolEvent(address indexed token, uint256 poolIndex)
func (_IL1PoolManager *IL1PoolManagerFilterer) FilterCompletePoolEvent(opts *bind.FilterOpts, token []common.Address) (*IL1PoolManagerCompletePoolEventIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IL1PoolManager.contract.FilterLogs(opts, "CompletePoolEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return &IL1PoolManagerCompletePoolEventIterator{contract: _IL1PoolManager.contract, event: "CompletePoolEvent", logs: logs, sub: sub}, nil
}

// WatchCompletePoolEvent is a free log subscription operation binding the contract event 0xb6f449f07ceaf55392c9899e0797c6529908ae827c2d498c682e90d42c241167.
//
// Solidity: event CompletePoolEvent(address indexed token, uint256 poolIndex)
func (_IL1PoolManager *IL1PoolManagerFilterer) WatchCompletePoolEvent(opts *bind.WatchOpts, sink chan<- *IL1PoolManagerCompletePoolEvent, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IL1PoolManager.contract.WatchLogs(opts, "CompletePoolEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1PoolManagerCompletePoolEvent)
				if err := _IL1PoolManager.contract.UnpackLog(event, "CompletePoolEvent", log); err != nil {
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

// ParseCompletePoolEvent is a log parse operation binding the contract event 0xb6f449f07ceaf55392c9899e0797c6529908ae827c2d498c682e90d42c241167.
//
// Solidity: event CompletePoolEvent(address indexed token, uint256 poolIndex)
func (_IL1PoolManager *IL1PoolManagerFilterer) ParseCompletePoolEvent(log types.Log) (*IL1PoolManagerCompletePoolEvent, error) {
	event := new(IL1PoolManagerCompletePoolEvent)
	if err := _IL1PoolManager.contract.UnpackLog(event, "CompletePoolEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1PoolManagerSetMinStakeAmountEventIterator is returned from FilterSetMinStakeAmountEvent and is used to iterate over the raw logs and unpacked data for SetMinStakeAmountEvent events raised by the IL1PoolManager contract.
type IL1PoolManagerSetMinStakeAmountEventIterator struct {
	Event *IL1PoolManagerSetMinStakeAmountEvent // Event containing the contract specifics and raw log

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
func (it *IL1PoolManagerSetMinStakeAmountEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1PoolManagerSetMinStakeAmountEvent)
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
		it.Event = new(IL1PoolManagerSetMinStakeAmountEvent)
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
func (it *IL1PoolManagerSetMinStakeAmountEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1PoolManagerSetMinStakeAmountEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1PoolManagerSetMinStakeAmountEvent represents a SetMinStakeAmountEvent event raised by the IL1PoolManager contract.
type IL1PoolManagerSetMinStakeAmountEvent struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMinStakeAmountEvent is a free log retrieval operation binding the contract event 0xf54d3b756d286b6b08e5d4eda6dfe5b135664abf029e58e637cbf013c442c950.
//
// Solidity: event SetMinStakeAmountEvent(address indexed token, uint256 amount)
func (_IL1PoolManager *IL1PoolManagerFilterer) FilterSetMinStakeAmountEvent(opts *bind.FilterOpts, token []common.Address) (*IL1PoolManagerSetMinStakeAmountEventIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IL1PoolManager.contract.FilterLogs(opts, "SetMinStakeAmountEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return &IL1PoolManagerSetMinStakeAmountEventIterator{contract: _IL1PoolManager.contract, event: "SetMinStakeAmountEvent", logs: logs, sub: sub}, nil
}

// WatchSetMinStakeAmountEvent is a free log subscription operation binding the contract event 0xf54d3b756d286b6b08e5d4eda6dfe5b135664abf029e58e637cbf013c442c950.
//
// Solidity: event SetMinStakeAmountEvent(address indexed token, uint256 amount)
func (_IL1PoolManager *IL1PoolManagerFilterer) WatchSetMinStakeAmountEvent(opts *bind.WatchOpts, sink chan<- *IL1PoolManagerSetMinStakeAmountEvent, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IL1PoolManager.contract.WatchLogs(opts, "SetMinStakeAmountEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1PoolManagerSetMinStakeAmountEvent)
				if err := _IL1PoolManager.contract.UnpackLog(event, "SetMinStakeAmountEvent", log); err != nil {
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

// ParseSetMinStakeAmountEvent is a log parse operation binding the contract event 0xf54d3b756d286b6b08e5d4eda6dfe5b135664abf029e58e637cbf013c442c950.
//
// Solidity: event SetMinStakeAmountEvent(address indexed token, uint256 amount)
func (_IL1PoolManager *IL1PoolManagerFilterer) ParseSetMinStakeAmountEvent(log types.Log) (*IL1PoolManagerSetMinStakeAmountEvent, error) {
	event := new(IL1PoolManagerSetMinStakeAmountEvent)
	if err := _IL1PoolManager.contract.UnpackLog(event, "SetMinStakeAmountEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1PoolManagerSetSupportTokenEventIterator is returned from FilterSetSupportTokenEvent and is used to iterate over the raw logs and unpacked data for SetSupportTokenEvent events raised by the IL1PoolManager contract.
type IL1PoolManagerSetSupportTokenEventIterator struct {
	Event *IL1PoolManagerSetSupportTokenEvent // Event containing the contract specifics and raw log

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
func (it *IL1PoolManagerSetSupportTokenEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1PoolManagerSetSupportTokenEvent)
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
		it.Event = new(IL1PoolManagerSetSupportTokenEvent)
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
func (it *IL1PoolManagerSetSupportTokenEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1PoolManagerSetSupportTokenEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1PoolManagerSetSupportTokenEvent represents a SetSupportTokenEvent event raised by the IL1PoolManager contract.
type IL1PoolManagerSetSupportTokenEvent struct {
	Token     common.Address
	IsSupport bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetSupportTokenEvent is a free log retrieval operation binding the contract event 0xc8c34f23fafb34e68119c1d231ef03d0d47225b15e2c4de3efbefa14b0181d86.
//
// Solidity: event SetSupportTokenEvent(address indexed token, bool isSupport)
func (_IL1PoolManager *IL1PoolManagerFilterer) FilterSetSupportTokenEvent(opts *bind.FilterOpts, token []common.Address) (*IL1PoolManagerSetSupportTokenEventIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IL1PoolManager.contract.FilterLogs(opts, "SetSupportTokenEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return &IL1PoolManagerSetSupportTokenEventIterator{contract: _IL1PoolManager.contract, event: "SetSupportTokenEvent", logs: logs, sub: sub}, nil
}

// WatchSetSupportTokenEvent is a free log subscription operation binding the contract event 0xc8c34f23fafb34e68119c1d231ef03d0d47225b15e2c4de3efbefa14b0181d86.
//
// Solidity: event SetSupportTokenEvent(address indexed token, bool isSupport)
func (_IL1PoolManager *IL1PoolManagerFilterer) WatchSetSupportTokenEvent(opts *bind.WatchOpts, sink chan<- *IL1PoolManagerSetSupportTokenEvent, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IL1PoolManager.contract.WatchLogs(opts, "SetSupportTokenEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1PoolManagerSetSupportTokenEvent)
				if err := _IL1PoolManager.contract.UnpackLog(event, "SetSupportTokenEvent", log); err != nil {
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

// ParseSetSupportTokenEvent is a log parse operation binding the contract event 0xc8c34f23fafb34e68119c1d231ef03d0d47225b15e2c4de3efbefa14b0181d86.
//
// Solidity: event SetSupportTokenEvent(address indexed token, bool isSupport)
func (_IL1PoolManager *IL1PoolManagerFilterer) ParseSetSupportTokenEvent(log types.Log) (*IL1PoolManagerSetSupportTokenEvent, error) {
	event := new(IL1PoolManagerSetSupportTokenEvent)
	if err := _IL1PoolManager.contract.UnpackLog(event, "SetSupportTokenEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1PoolManagerStakingETHEventIterator is returned from FilterStakingETHEvent and is used to iterate over the raw logs and unpacked data for StakingETHEvent events raised by the IL1PoolManager contract.
type IL1PoolManagerStakingETHEventIterator struct {
	Event *IL1PoolManagerStakingETHEvent // Event containing the contract specifics and raw log

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
func (it *IL1PoolManagerStakingETHEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1PoolManagerStakingETHEvent)
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
		it.Event = new(IL1PoolManagerStakingETHEvent)
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
func (it *IL1PoolManagerStakingETHEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1PoolManagerStakingETHEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1PoolManagerStakingETHEvent represents a StakingETHEvent event raised by the IL1PoolManager contract.
type IL1PoolManagerStakingETHEvent struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakingETHEvent is a free log retrieval operation binding the contract event 0xe7466ea83435490635fc76a5f33da4815758ab48b1d45858f0452ca646556937.
//
// Solidity: event StakingETHEvent(address indexed user, uint256 amount)
func (_IL1PoolManager *IL1PoolManagerFilterer) FilterStakingETHEvent(opts *bind.FilterOpts, user []common.Address) (*IL1PoolManagerStakingETHEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IL1PoolManager.contract.FilterLogs(opts, "StakingETHEvent", userRule)
	if err != nil {
		return nil, err
	}
	return &IL1PoolManagerStakingETHEventIterator{contract: _IL1PoolManager.contract, event: "StakingETHEvent", logs: logs, sub: sub}, nil
}

// WatchStakingETHEvent is a free log subscription operation binding the contract event 0xe7466ea83435490635fc76a5f33da4815758ab48b1d45858f0452ca646556937.
//
// Solidity: event StakingETHEvent(address indexed user, uint256 amount)
func (_IL1PoolManager *IL1PoolManagerFilterer) WatchStakingETHEvent(opts *bind.WatchOpts, sink chan<- *IL1PoolManagerStakingETHEvent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IL1PoolManager.contract.WatchLogs(opts, "StakingETHEvent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1PoolManagerStakingETHEvent)
				if err := _IL1PoolManager.contract.UnpackLog(event, "StakingETHEvent", log); err != nil {
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

// ParseStakingETHEvent is a log parse operation binding the contract event 0xe7466ea83435490635fc76a5f33da4815758ab48b1d45858f0452ca646556937.
//
// Solidity: event StakingETHEvent(address indexed user, uint256 amount)
func (_IL1PoolManager *IL1PoolManagerFilterer) ParseStakingETHEvent(log types.Log) (*IL1PoolManagerStakingETHEvent, error) {
	event := new(IL1PoolManagerStakingETHEvent)
	if err := _IL1PoolManager.contract.UnpackLog(event, "StakingETHEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1PoolManagerStakingWETHEventIterator is returned from FilterStakingWETHEvent and is used to iterate over the raw logs and unpacked data for StakingWETHEvent events raised by the IL1PoolManager contract.
type IL1PoolManagerStakingWETHEventIterator struct {
	Event *IL1PoolManagerStakingWETHEvent // Event containing the contract specifics and raw log

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
func (it *IL1PoolManagerStakingWETHEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1PoolManagerStakingWETHEvent)
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
		it.Event = new(IL1PoolManagerStakingWETHEvent)
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
func (it *IL1PoolManagerStakingWETHEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1PoolManagerStakingWETHEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1PoolManagerStakingWETHEvent represents a StakingWETHEvent event raised by the IL1PoolManager contract.
type IL1PoolManagerStakingWETHEvent struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakingWETHEvent is a free log retrieval operation binding the contract event 0xc138e3bd6d13eefa5cb01c0d35c5794001141efaf4e5ad888cad059935f83383.
//
// Solidity: event StakingWETHEvent(address indexed user, uint256 amount)
func (_IL1PoolManager *IL1PoolManagerFilterer) FilterStakingWETHEvent(opts *bind.FilterOpts, user []common.Address) (*IL1PoolManagerStakingWETHEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IL1PoolManager.contract.FilterLogs(opts, "StakingWETHEvent", userRule)
	if err != nil {
		return nil, err
	}
	return &IL1PoolManagerStakingWETHEventIterator{contract: _IL1PoolManager.contract, event: "StakingWETHEvent", logs: logs, sub: sub}, nil
}

// WatchStakingWETHEvent is a free log subscription operation binding the contract event 0xc138e3bd6d13eefa5cb01c0d35c5794001141efaf4e5ad888cad059935f83383.
//
// Solidity: event StakingWETHEvent(address indexed user, uint256 amount)
func (_IL1PoolManager *IL1PoolManagerFilterer) WatchStakingWETHEvent(opts *bind.WatchOpts, sink chan<- *IL1PoolManagerStakingWETHEvent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IL1PoolManager.contract.WatchLogs(opts, "StakingWETHEvent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1PoolManagerStakingWETHEvent)
				if err := _IL1PoolManager.contract.UnpackLog(event, "StakingWETHEvent", log); err != nil {
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

// ParseStakingWETHEvent is a log parse operation binding the contract event 0xc138e3bd6d13eefa5cb01c0d35c5794001141efaf4e5ad888cad059935f83383.
//
// Solidity: event StakingWETHEvent(address indexed user, uint256 amount)
func (_IL1PoolManager *IL1PoolManagerFilterer) ParseStakingWETHEvent(log types.Log) (*IL1PoolManagerStakingWETHEvent, error) {
	event := new(IL1PoolManagerStakingWETHEvent)
	if err := _IL1PoolManager.contract.UnpackLog(event, "StakingWETHEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1PoolManagerStarkingERC20EventIterator is returned from FilterStarkingERC20Event and is used to iterate over the raw logs and unpacked data for StarkingERC20Event events raised by the IL1PoolManager contract.
type IL1PoolManagerStarkingERC20EventIterator struct {
	Event *IL1PoolManagerStarkingERC20Event // Event containing the contract specifics and raw log

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
func (it *IL1PoolManagerStarkingERC20EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1PoolManagerStarkingERC20Event)
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
		it.Event = new(IL1PoolManagerStarkingERC20Event)
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
func (it *IL1PoolManagerStarkingERC20EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1PoolManagerStarkingERC20EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1PoolManagerStarkingERC20Event represents a StarkingERC20Event event raised by the IL1PoolManager contract.
type IL1PoolManagerStarkingERC20Event struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStarkingERC20Event is a free log retrieval operation binding the contract event 0x3ac7d3823c11677fba9479ed26f696a3f17e16a5a5c39162fa6d183905aa6735.
//
// Solidity: event StarkingERC20Event(address indexed user, address indexed token, uint256 amount)
func (_IL1PoolManager *IL1PoolManagerFilterer) FilterStarkingERC20Event(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*IL1PoolManagerStarkingERC20EventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IL1PoolManager.contract.FilterLogs(opts, "StarkingERC20Event", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IL1PoolManagerStarkingERC20EventIterator{contract: _IL1PoolManager.contract, event: "StarkingERC20Event", logs: logs, sub: sub}, nil
}

// WatchStarkingERC20Event is a free log subscription operation binding the contract event 0x3ac7d3823c11677fba9479ed26f696a3f17e16a5a5c39162fa6d183905aa6735.
//
// Solidity: event StarkingERC20Event(address indexed user, address indexed token, uint256 amount)
func (_IL1PoolManager *IL1PoolManagerFilterer) WatchStarkingERC20Event(opts *bind.WatchOpts, sink chan<- *IL1PoolManagerStarkingERC20Event, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IL1PoolManager.contract.WatchLogs(opts, "StarkingERC20Event", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1PoolManagerStarkingERC20Event)
				if err := _IL1PoolManager.contract.UnpackLog(event, "StarkingERC20Event", log); err != nil {
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

// ParseStarkingERC20Event is a log parse operation binding the contract event 0x3ac7d3823c11677fba9479ed26f696a3f17e16a5a5c39162fa6d183905aa6735.
//
// Solidity: event StarkingERC20Event(address indexed user, address indexed token, uint256 amount)
func (_IL1PoolManager *IL1PoolManagerFilterer) ParseStarkingERC20Event(log types.Log) (*IL1PoolManagerStarkingERC20Event, error) {
	event := new(IL1PoolManagerStarkingERC20Event)
	if err := _IL1PoolManager.contract.UnpackLog(event, "StarkingERC20Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1PoolManagerTransferAssertToIterator is returned from FilterTransferAssertTo and is used to iterate over the raw logs and unpacked data for TransferAssertTo events raised by the IL1PoolManager contract.
type IL1PoolManagerTransferAssertToIterator struct {
	Event *IL1PoolManagerTransferAssertTo // Event containing the contract specifics and raw log

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
func (it *IL1PoolManagerTransferAssertToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1PoolManagerTransferAssertTo)
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
		it.Event = new(IL1PoolManagerTransferAssertTo)
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
func (it *IL1PoolManagerTransferAssertToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1PoolManagerTransferAssertToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1PoolManagerTransferAssertTo represents a TransferAssertTo event raised by the IL1PoolManager contract.
type IL1PoolManagerTransferAssertTo struct {
	Blockchain *big.Int
	Token      common.Address
	To         common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferAssertTo is a free log retrieval operation binding the contract event 0x08ffba656e4ac665e1aa4bb6a342e1a69d4cc12fd751ac862fa3fe27b0c7524a.
//
// Solidity: event TransferAssertTo(uint256 Blockchain, address indexed token, address indexed to, uint256 amount)
func (_IL1PoolManager *IL1PoolManagerFilterer) FilterTransferAssertTo(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*IL1PoolManagerTransferAssertToIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IL1PoolManager.contract.FilterLogs(opts, "TransferAssertTo", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IL1PoolManagerTransferAssertToIterator{contract: _IL1PoolManager.contract, event: "TransferAssertTo", logs: logs, sub: sub}, nil
}

// WatchTransferAssertTo is a free log subscription operation binding the contract event 0x08ffba656e4ac665e1aa4bb6a342e1a69d4cc12fd751ac862fa3fe27b0c7524a.
//
// Solidity: event TransferAssertTo(uint256 Blockchain, address indexed token, address indexed to, uint256 amount)
func (_IL1PoolManager *IL1PoolManagerFilterer) WatchTransferAssertTo(opts *bind.WatchOpts, sink chan<- *IL1PoolManagerTransferAssertTo, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IL1PoolManager.contract.WatchLogs(opts, "TransferAssertTo", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1PoolManagerTransferAssertTo)
				if err := _IL1PoolManager.contract.UnpackLog(event, "TransferAssertTo", log); err != nil {
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

// ParseTransferAssertTo is a log parse operation binding the contract event 0x08ffba656e4ac665e1aa4bb6a342e1a69d4cc12fd751ac862fa3fe27b0c7524a.
//
// Solidity: event TransferAssertTo(uint256 Blockchain, address indexed token, address indexed to, uint256 amount)
func (_IL1PoolManager *IL1PoolManagerFilterer) ParseTransferAssertTo(log types.Log) (*IL1PoolManagerTransferAssertTo, error) {
	event := new(IL1PoolManagerTransferAssertTo)
	if err := _IL1PoolManager.contract.UnpackLog(event, "TransferAssertTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1PoolManagerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the IL1PoolManager contract.
type IL1PoolManagerWithdrawIterator struct {
	Event *IL1PoolManagerWithdraw // Event containing the contract specifics and raw log

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
func (it *IL1PoolManagerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1PoolManagerWithdraw)
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
		it.Event = new(IL1PoolManagerWithdraw)
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
func (it *IL1PoolManagerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1PoolManagerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1PoolManagerWithdraw represents a Withdraw event raised by the IL1PoolManager contract.
type IL1PoolManagerWithdraw struct {
	User        common.Address
	StartPoolId *big.Int
	EndPoolId   *big.Int
	Token       common.Address
	Amount      *big.Int
	Reward      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x0ad30059d4072ae2525ebaaa5b6878f1c50df5f3b921615031c7d0f12dc4a77b.
//
// Solidity: event Withdraw(address _user, uint256 startPoolId, uint256 EndPoolId, address _token, uint256 Amount, uint256 Reward)
func (_IL1PoolManager *IL1PoolManagerFilterer) FilterWithdraw(opts *bind.FilterOpts) (*IL1PoolManagerWithdrawIterator, error) {

	logs, sub, err := _IL1PoolManager.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &IL1PoolManagerWithdrawIterator{contract: _IL1PoolManager.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x0ad30059d4072ae2525ebaaa5b6878f1c50df5f3b921615031c7d0f12dc4a77b.
//
// Solidity: event Withdraw(address _user, uint256 startPoolId, uint256 EndPoolId, address _token, uint256 Amount, uint256 Reward)
func (_IL1PoolManager *IL1PoolManagerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IL1PoolManagerWithdraw) (event.Subscription, error) {

	logs, sub, err := _IL1PoolManager.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1PoolManagerWithdraw)
				if err := _IL1PoolManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x0ad30059d4072ae2525ebaaa5b6878f1c50df5f3b921615031c7d0f12dc4a77b.
//
// Solidity: event Withdraw(address _user, uint256 startPoolId, uint256 EndPoolId, address _token, uint256 Amount, uint256 Reward)
func (_IL1PoolManager *IL1PoolManagerFilterer) ParseWithdraw(log types.Log) (*IL1PoolManagerWithdraw, error) {
	event := new(IL1PoolManagerWithdraw)
	if err := _IL1PoolManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
