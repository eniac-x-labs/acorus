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

// L1RewardManagerMetaData contains all meta data concerning the L1RewardManager contract.
var L1RewardManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"L1RewardBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimL1Reward\",\"inputs\":[{\"name\":\"_strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositETHRewardTo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakerRewardsAmount\",\"inputs\":[{\"name\":\"_strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ClaimL1Reward\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositETHRewardTo\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x608080604052346100b8577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166100a957506001600160401b036002600160401b031982821601610064575b6040516107a290816100be8239f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1388080610055565b63f92ee8a960e01b8152600490fd5b600080fdfe604060808152600436101561001357600080fd5b600090813560e01c80631fbe54531461041957806339b70e38146103f15780633d9dd75714610375578063485cc9551461021c578063715018a6146101af57806380fd8a7a146101935780638da5cb5b1461015e578063a90d9de2146100db5763f2fde38b1461008257600080fd5b346100d75760203660031901126100d75761009b610487565b906100a4610713565b6001600160a01b038216156100c057506100bd906106bd565b80f35b51631e4fbdf760e01b815260048101839052602490fd5b5080fd5b50816003193601126100d75781803415610155575b8180809234903090f11561014b5760209161010c34825461050f565b905580513381523460208201527fe1803cf89e002ffcc9db1d627a63ae2358933f2d1ab32929935c4db3b9edce149080604081015b0390a15160018152f35b51903d90823e3d90fd5b506108fc6100f0565b50346100d757816003193601126100d75760008051602061074d8339815191525490516001600160a01b039091168152602090f35b50346100d757816003193601126100d757602091549051908152f35b82346102195780600319360112610219576101c8610713565b60008051602061074d83398151915280546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b80fd5b50346100d757806003193601126100d757610235610487565b6024356001600160a01b0381169190829003610371577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009182549160ff83861c16159267ffffffffffffffff811680159081610369575b600114908161035f575b159081610356575b506103455767ffffffffffffffff19811660011785556102c5919084610326575b506106bd565b6bffffffffffffffffffffffff60a01b60015416176001556102e5578280f35b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a138808280f35b68ffffffffffffffffff191668010000000000000001178555386102bf565b855163f92ee8a960e01b8152600490fd5b9050153861029e565b303b159150610296565b85915061028c565b8380fd5b5061038861038236610435565b90610532565b9180808480156103e7575b8280929181923390f1156103dc575080513381526020818101939093527f131faee866ce16b47522e4e23f5ec64a21817b14c056b57cf6f187d6551561bd908060408101610141565b9051903d90823e3d90fd5b6108fc9150610393565b50346100d757816003193601126100d75760015490516001600160a01b039091168152602090f35b50346100d75760209061042e61038236610435565b9051908152f35b9060206003198301126104825760043567ffffffffffffffff9283821161048257806023830112156104825781600401359384116104825760248460051b83010111610482576024019190565b600080fd5b600435906001600160a01b038216820361048257565b91908110156104ad5760051b0190565b634e487b7160e01b600052603260045260246000fd5b356001600160a01b03811681036104825790565b90601f8019910116810190811067ffffffffffffffff8211176104f957604052565b634e487b7160e01b600052604160045260246000fd5b9190820180921161051c57565b634e487b7160e01b600052601160045260246000fd5b600090819282915b8183106105965750505080159182801561058e575b610586576000549261057057049081810291818304149015171561051c5790565b634e487b7160e01b600052601260045260246000fd5b505050600090565b50801561054f565b91936001600160a01b0393919291846105b86105b388878761049d565b6104c3565b1690604091825191828092633a98ef3960e01b8252602093849160049687915afa9081156106b257908391600091610682575b506105f6919261050f565b976106056105b38b8a8a61049d565b1692602485518095819363673e156160e11b835233908301525afa9283156106785750600092610647575b505061063e9060019261050f565b9401919061053a565b81819392933d8311610671575b61065e81836104d7565b810103126102195750518161063e610630565b503d610654565b513d6000823e3d90fd5b9182813d83116106ab575b61069781836104d7565b8101031261021957505182906105f66105eb565b503d61068d565b85513d6000823e3d90fd5b60008051602061074d83398151915280546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b60008051602061074d833981519152546001600160a01b0316330361073457565b60405163118cdaa760e01b8152336004820152602490fdfe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300a26469706673582212205ccf5f4fc55b25af35aba97f7e0d0892305f9d459a269a39056ed2ac9119fed264736f6c63430008180033",
}

// L1RewardManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use L1RewardManagerMetaData.ABI instead.
var L1RewardManagerABI = L1RewardManagerMetaData.ABI

// L1RewardManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1RewardManagerMetaData.Bin instead.
var L1RewardManagerBin = L1RewardManagerMetaData.Bin

// DeployL1RewardManager deploys a new Ethereum contract, binding an instance of L1RewardManager to it.
func DeployL1RewardManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L1RewardManager, error) {
	parsed, err := L1RewardManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1RewardManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1RewardManager{L1RewardManagerCaller: L1RewardManagerCaller{contract: contract}, L1RewardManagerTransactor: L1RewardManagerTransactor{contract: contract}, L1RewardManagerFilterer: L1RewardManagerFilterer{contract: contract}}, nil
}

// L1RewardManager is an auto generated Go binding around an Ethereum contract.
type L1RewardManager struct {
	L1RewardManagerCaller     // Read-only binding to the contract
	L1RewardManagerTransactor // Write-only binding to the contract
	L1RewardManagerFilterer   // Log filterer for contract events
}

// L1RewardManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1RewardManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1RewardManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1RewardManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1RewardManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1RewardManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1RewardManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1RewardManagerSession struct {
	Contract     *L1RewardManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1RewardManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1RewardManagerCallerSession struct {
	Contract *L1RewardManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L1RewardManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1RewardManagerTransactorSession struct {
	Contract     *L1RewardManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L1RewardManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1RewardManagerRaw struct {
	Contract *L1RewardManager // Generic contract binding to access the raw methods on
}

// L1RewardManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1RewardManagerCallerRaw struct {
	Contract *L1RewardManagerCaller // Generic read-only contract binding to access the raw methods on
}

// L1RewardManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1RewardManagerTransactorRaw struct {
	Contract *L1RewardManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1RewardManager creates a new instance of L1RewardManager, bound to a specific deployed contract.
func NewL1RewardManager(address common.Address, backend bind.ContractBackend) (*L1RewardManager, error) {
	contract, err := bindL1RewardManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1RewardManager{L1RewardManagerCaller: L1RewardManagerCaller{contract: contract}, L1RewardManagerTransactor: L1RewardManagerTransactor{contract: contract}, L1RewardManagerFilterer: L1RewardManagerFilterer{contract: contract}}, nil
}

// NewL1RewardManagerCaller creates a new read-only instance of L1RewardManager, bound to a specific deployed contract.
func NewL1RewardManagerCaller(address common.Address, caller bind.ContractCaller) (*L1RewardManagerCaller, error) {
	contract, err := bindL1RewardManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1RewardManagerCaller{contract: contract}, nil
}

// NewL1RewardManagerTransactor creates a new write-only instance of L1RewardManager, bound to a specific deployed contract.
func NewL1RewardManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*L1RewardManagerTransactor, error) {
	contract, err := bindL1RewardManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1RewardManagerTransactor{contract: contract}, nil
}

// NewL1RewardManagerFilterer creates a new log filterer instance of L1RewardManager, bound to a specific deployed contract.
func NewL1RewardManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*L1RewardManagerFilterer, error) {
	contract, err := bindL1RewardManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1RewardManagerFilterer{contract: contract}, nil
}

// bindL1RewardManager binds a generic wrapper to an already deployed contract.
func bindL1RewardManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1RewardManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1RewardManager *L1RewardManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1RewardManager.Contract.L1RewardManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1RewardManager *L1RewardManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1RewardManager.Contract.L1RewardManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1RewardManager *L1RewardManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1RewardManager.Contract.L1RewardManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1RewardManager *L1RewardManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1RewardManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1RewardManager *L1RewardManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1RewardManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1RewardManager *L1RewardManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1RewardManager.Contract.contract.Transact(opts, method, params...)
}

// L1RewardBalance is a free data retrieval call binding the contract method 0x80fd8a7a.
//
// Solidity: function L1RewardBalance() view returns(uint256)
func (_L1RewardManager *L1RewardManagerCaller) L1RewardBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1RewardManager.contract.Call(opts, &out, "L1RewardBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1RewardBalance is a free data retrieval call binding the contract method 0x80fd8a7a.
//
// Solidity: function L1RewardBalance() view returns(uint256)
func (_L1RewardManager *L1RewardManagerSession) L1RewardBalance() (*big.Int, error) {
	return _L1RewardManager.Contract.L1RewardBalance(&_L1RewardManager.CallOpts)
}

// L1RewardBalance is a free data retrieval call binding the contract method 0x80fd8a7a.
//
// Solidity: function L1RewardBalance() view returns(uint256)
func (_L1RewardManager *L1RewardManagerCallerSession) L1RewardBalance() (*big.Int, error) {
	return _L1RewardManager.Contract.L1RewardBalance(&_L1RewardManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1RewardManager *L1RewardManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1RewardManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1RewardManager *L1RewardManagerSession) Owner() (common.Address, error) {
	return _L1RewardManager.Contract.Owner(&_L1RewardManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1RewardManager *L1RewardManagerCallerSession) Owner() (common.Address, error) {
	return _L1RewardManager.Contract.Owner(&_L1RewardManager.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_L1RewardManager *L1RewardManagerCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1RewardManager.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_L1RewardManager *L1RewardManagerSession) StrategyManager() (common.Address, error) {
	return _L1RewardManager.Contract.StrategyManager(&_L1RewardManager.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_L1RewardManager *L1RewardManagerCallerSession) StrategyManager() (common.Address, error) {
	return _L1RewardManager.Contract.StrategyManager(&_L1RewardManager.CallOpts)
}

// ClaimL1Reward is a paid mutator transaction binding the contract method 0x3d9dd757.
//
// Solidity: function claimL1Reward(address[] _strategies) payable returns(bool)
func (_L1RewardManager *L1RewardManagerTransactor) ClaimL1Reward(opts *bind.TransactOpts, _strategies []common.Address) (*types.Transaction, error) {
	return _L1RewardManager.contract.Transact(opts, "claimL1Reward", _strategies)
}

// ClaimL1Reward is a paid mutator transaction binding the contract method 0x3d9dd757.
//
// Solidity: function claimL1Reward(address[] _strategies) payable returns(bool)
func (_L1RewardManager *L1RewardManagerSession) ClaimL1Reward(_strategies []common.Address) (*types.Transaction, error) {
	return _L1RewardManager.Contract.ClaimL1Reward(&_L1RewardManager.TransactOpts, _strategies)
}

// ClaimL1Reward is a paid mutator transaction binding the contract method 0x3d9dd757.
//
// Solidity: function claimL1Reward(address[] _strategies) payable returns(bool)
func (_L1RewardManager *L1RewardManagerTransactorSession) ClaimL1Reward(_strategies []common.Address) (*types.Transaction, error) {
	return _L1RewardManager.Contract.ClaimL1Reward(&_L1RewardManager.TransactOpts, _strategies)
}

// DepositETHRewardTo is a paid mutator transaction binding the contract method 0xa90d9de2.
//
// Solidity: function depositETHRewardTo() payable returns(bool)
func (_L1RewardManager *L1RewardManagerTransactor) DepositETHRewardTo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1RewardManager.contract.Transact(opts, "depositETHRewardTo")
}

// DepositETHRewardTo is a paid mutator transaction binding the contract method 0xa90d9de2.
//
// Solidity: function depositETHRewardTo() payable returns(bool)
func (_L1RewardManager *L1RewardManagerSession) DepositETHRewardTo() (*types.Transaction, error) {
	return _L1RewardManager.Contract.DepositETHRewardTo(&_L1RewardManager.TransactOpts)
}

// DepositETHRewardTo is a paid mutator transaction binding the contract method 0xa90d9de2.
//
// Solidity: function depositETHRewardTo() payable returns(bool)
func (_L1RewardManager *L1RewardManagerTransactorSession) DepositETHRewardTo() (*types.Transaction, error) {
	return _L1RewardManager.Contract.DepositETHRewardTo(&_L1RewardManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _strategyManager) returns()
func (_L1RewardManager *L1RewardManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _strategyManager common.Address) (*types.Transaction, error) {
	return _L1RewardManager.contract.Transact(opts, "initialize", initialOwner, _strategyManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _strategyManager) returns()
func (_L1RewardManager *L1RewardManagerSession) Initialize(initialOwner common.Address, _strategyManager common.Address) (*types.Transaction, error) {
	return _L1RewardManager.Contract.Initialize(&_L1RewardManager.TransactOpts, initialOwner, _strategyManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _strategyManager) returns()
func (_L1RewardManager *L1RewardManagerTransactorSession) Initialize(initialOwner common.Address, _strategyManager common.Address) (*types.Transaction, error) {
	return _L1RewardManager.Contract.Initialize(&_L1RewardManager.TransactOpts, initialOwner, _strategyManager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1RewardManager *L1RewardManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1RewardManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1RewardManager *L1RewardManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1RewardManager.Contract.RenounceOwnership(&_L1RewardManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1RewardManager *L1RewardManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1RewardManager.Contract.RenounceOwnership(&_L1RewardManager.TransactOpts)
}

// StakerRewardsAmount is a paid mutator transaction binding the contract method 0x1fbe5453.
//
// Solidity: function stakerRewardsAmount(address[] _strategies) returns(uint256)
func (_L1RewardManager *L1RewardManagerTransactor) StakerRewardsAmount(opts *bind.TransactOpts, _strategies []common.Address) (*types.Transaction, error) {
	return _L1RewardManager.contract.Transact(opts, "stakerRewardsAmount", _strategies)
}

// StakerRewardsAmount is a paid mutator transaction binding the contract method 0x1fbe5453.
//
// Solidity: function stakerRewardsAmount(address[] _strategies) returns(uint256)
func (_L1RewardManager *L1RewardManagerSession) StakerRewardsAmount(_strategies []common.Address) (*types.Transaction, error) {
	return _L1RewardManager.Contract.StakerRewardsAmount(&_L1RewardManager.TransactOpts, _strategies)
}

// StakerRewardsAmount is a paid mutator transaction binding the contract method 0x1fbe5453.
//
// Solidity: function stakerRewardsAmount(address[] _strategies) returns(uint256)
func (_L1RewardManager *L1RewardManagerTransactorSession) StakerRewardsAmount(_strategies []common.Address) (*types.Transaction, error) {
	return _L1RewardManager.Contract.StakerRewardsAmount(&_L1RewardManager.TransactOpts, _strategies)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1RewardManager *L1RewardManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1RewardManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1RewardManager *L1RewardManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1RewardManager.Contract.TransferOwnership(&_L1RewardManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1RewardManager *L1RewardManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1RewardManager.Contract.TransferOwnership(&_L1RewardManager.TransactOpts, newOwner)
}

// L1RewardManagerClaimL1RewardIterator is returned from FilterClaimL1Reward and is used to iterate over the raw logs and unpacked data for ClaimL1Reward events raised by the L1RewardManager contract.
type L1RewardManagerClaimL1RewardIterator struct {
	Event *L1RewardManagerClaimL1Reward // Event containing the contract specifics and raw log

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
func (it *L1RewardManagerClaimL1RewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1RewardManagerClaimL1Reward)
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
		it.Event = new(L1RewardManagerClaimL1Reward)
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
func (it *L1RewardManagerClaimL1RewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1RewardManagerClaimL1RewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1RewardManagerClaimL1Reward represents a ClaimL1Reward event raised by the L1RewardManager contract.
type L1RewardManagerClaimL1Reward struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimL1Reward is a free log retrieval operation binding the contract event 0x131faee866ce16b47522e4e23f5ec64a21817b14c056b57cf6f187d6551561bd.
//
// Solidity: event ClaimL1Reward(address receiver, uint256 amount)
func (_L1RewardManager *L1RewardManagerFilterer) FilterClaimL1Reward(opts *bind.FilterOpts) (*L1RewardManagerClaimL1RewardIterator, error) {

	logs, sub, err := _L1RewardManager.contract.FilterLogs(opts, "ClaimL1Reward")
	if err != nil {
		return nil, err
	}
	return &L1RewardManagerClaimL1RewardIterator{contract: _L1RewardManager.contract, event: "ClaimL1Reward", logs: logs, sub: sub}, nil
}

// WatchClaimL1Reward is a free log subscription operation binding the contract event 0x131faee866ce16b47522e4e23f5ec64a21817b14c056b57cf6f187d6551561bd.
//
// Solidity: event ClaimL1Reward(address receiver, uint256 amount)
func (_L1RewardManager *L1RewardManagerFilterer) WatchClaimL1Reward(opts *bind.WatchOpts, sink chan<- *L1RewardManagerClaimL1Reward) (event.Subscription, error) {

	logs, sub, err := _L1RewardManager.contract.WatchLogs(opts, "ClaimL1Reward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1RewardManagerClaimL1Reward)
				if err := _L1RewardManager.contract.UnpackLog(event, "ClaimL1Reward", log); err != nil {
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

// ParseClaimL1Reward is a log parse operation binding the contract event 0x131faee866ce16b47522e4e23f5ec64a21817b14c056b57cf6f187d6551561bd.
//
// Solidity: event ClaimL1Reward(address receiver, uint256 amount)
func (_L1RewardManager *L1RewardManagerFilterer) ParseClaimL1Reward(log types.Log) (*L1RewardManagerClaimL1Reward, error) {
	event := new(L1RewardManagerClaimL1Reward)
	if err := _L1RewardManager.contract.UnpackLog(event, "ClaimL1Reward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1RewardManagerDepositETHRewardToIterator is returned from FilterDepositETHRewardTo and is used to iterate over the raw logs and unpacked data for DepositETHRewardTo events raised by the L1RewardManager contract.
type L1RewardManagerDepositETHRewardToIterator struct {
	Event *L1RewardManagerDepositETHRewardTo // Event containing the contract specifics and raw log

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
func (it *L1RewardManagerDepositETHRewardToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1RewardManagerDepositETHRewardTo)
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
		it.Event = new(L1RewardManagerDepositETHRewardTo)
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
func (it *L1RewardManagerDepositETHRewardToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1RewardManagerDepositETHRewardToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1RewardManagerDepositETHRewardTo represents a DepositETHRewardTo event raised by the L1RewardManager contract.
type L1RewardManagerDepositETHRewardTo struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositETHRewardTo is a free log retrieval operation binding the contract event 0xe1803cf89e002ffcc9db1d627a63ae2358933f2d1ab32929935c4db3b9edce14.
//
// Solidity: event DepositETHRewardTo(address sender, uint256 amount)
func (_L1RewardManager *L1RewardManagerFilterer) FilterDepositETHRewardTo(opts *bind.FilterOpts) (*L1RewardManagerDepositETHRewardToIterator, error) {

	logs, sub, err := _L1RewardManager.contract.FilterLogs(opts, "DepositETHRewardTo")
	if err != nil {
		return nil, err
	}
	return &L1RewardManagerDepositETHRewardToIterator{contract: _L1RewardManager.contract, event: "DepositETHRewardTo", logs: logs, sub: sub}, nil
}

// WatchDepositETHRewardTo is a free log subscription operation binding the contract event 0xe1803cf89e002ffcc9db1d627a63ae2358933f2d1ab32929935c4db3b9edce14.
//
// Solidity: event DepositETHRewardTo(address sender, uint256 amount)
func (_L1RewardManager *L1RewardManagerFilterer) WatchDepositETHRewardTo(opts *bind.WatchOpts, sink chan<- *L1RewardManagerDepositETHRewardTo) (event.Subscription, error) {

	logs, sub, err := _L1RewardManager.contract.WatchLogs(opts, "DepositETHRewardTo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1RewardManagerDepositETHRewardTo)
				if err := _L1RewardManager.contract.UnpackLog(event, "DepositETHRewardTo", log); err != nil {
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

// ParseDepositETHRewardTo is a log parse operation binding the contract event 0xe1803cf89e002ffcc9db1d627a63ae2358933f2d1ab32929935c4db3b9edce14.
//
// Solidity: event DepositETHRewardTo(address sender, uint256 amount)
func (_L1RewardManager *L1RewardManagerFilterer) ParseDepositETHRewardTo(log types.Log) (*L1RewardManagerDepositETHRewardTo, error) {
	event := new(L1RewardManagerDepositETHRewardTo)
	if err := _L1RewardManager.contract.UnpackLog(event, "DepositETHRewardTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1RewardManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1RewardManager contract.
type L1RewardManagerInitializedIterator struct {
	Event *L1RewardManagerInitialized // Event containing the contract specifics and raw log

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
func (it *L1RewardManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1RewardManagerInitialized)
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
		it.Event = new(L1RewardManagerInitialized)
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
func (it *L1RewardManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1RewardManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1RewardManagerInitialized represents a Initialized event raised by the L1RewardManager contract.
type L1RewardManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L1RewardManager *L1RewardManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1RewardManagerInitializedIterator, error) {

	logs, sub, err := _L1RewardManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1RewardManagerInitializedIterator{contract: _L1RewardManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L1RewardManager *L1RewardManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1RewardManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _L1RewardManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1RewardManagerInitialized)
				if err := _L1RewardManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1RewardManager *L1RewardManagerFilterer) ParseInitialized(log types.Log) (*L1RewardManagerInitialized, error) {
	event := new(L1RewardManagerInitialized)
	if err := _L1RewardManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1RewardManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1RewardManager contract.
type L1RewardManagerOwnershipTransferredIterator struct {
	Event *L1RewardManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1RewardManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1RewardManagerOwnershipTransferred)
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
		it.Event = new(L1RewardManagerOwnershipTransferred)
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
func (it *L1RewardManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1RewardManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1RewardManagerOwnershipTransferred represents a OwnershipTransferred event raised by the L1RewardManager contract.
type L1RewardManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1RewardManager *L1RewardManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1RewardManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1RewardManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1RewardManagerOwnershipTransferredIterator{contract: _L1RewardManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1RewardManager *L1RewardManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1RewardManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1RewardManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1RewardManagerOwnershipTransferred)
				if err := _L1RewardManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1RewardManager *L1RewardManagerFilterer) ParseOwnershipTransferred(log types.Log) (*L1RewardManagerOwnershipTransferred, error) {
	event := new(L1RewardManagerOwnershipTransferred)
	if err := _L1RewardManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
