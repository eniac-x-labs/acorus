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

// L2RewardManagerMetaData contains all meta data concerning the L2RewardManager contract.
var L2RewardManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateFee\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"baseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositDappLinkToken\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_rewardToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorClaimReward\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorRewards\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerClaimReward\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakerPercent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerRewards\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerRewardsAmount\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorAndStakerShareFee\",\"inputs\":[{\"name\":\"_stakerPercent\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DepositDappLinkToken\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorClaimReward\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorStakerReward\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"stakerFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"operatorFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerClaimReward\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608080604052346100bd57605c6005557ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166100ae57506001600160401b036002600160401b031982821601610069575b604051610bbc90816100c38239f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a138808061005a565b63f92ee8a960e01b8152600490fd5b600080fdfe608060408181526004918236101561001657600080fd5b600090813560e01c9081631546e9801461078d575080632cb082721461073257806339b70e381461070a57806341a2b8d6146106d25780635adbc180146106aa578063715018a6146106405780637ca87cb6146104705780638da5cb5b1461043a578063df5cf72314610411578063f0f0b228146103da578063f2fde38b14610384578063f4d7152b1461036a578063f7c618c114610341578063f8401008146102d8578063f8c8765e1461014a5763fb1bca50146100d457600080fd5b3461014757602036600319011261014757507f08fba1e7a1091d4e4242ae25b8f8291ec7423fba91244b59db8030ee503dfbe861013d602093356101268160018060a01b03600254163090339061095f565b835133815260208101919091529081906040820190565b0390a15160018152f35b80fd5b5091346102d45760803660031901126102d4576101656107a9565b6001600160a01b03602435818116908190036102d057604435938285168095036102cc576064359283168093036102cc577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009485549460ff86891c16159567ffffffffffffffff8116801590816102c4575b60011490816102ba575b1590816102b1575b506102a15767ffffffffffffffff1981166001178855610210919087610282575b50610aad565b6bffffffffffffffffffffffff60a01b92836003541617600355828254161790556002541617600255610241578280f35b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a138808280f35b68ffffffffffffffffff1916680100000000000000011788553861020a565b885163f92ee8a960e01b81528490fd5b905015386101e9565b303b1591506101e1565b8891506101d7565b8680fd5b8580fd5b8280fd5b50903461033d57602036600319011261033d576020907f7246707380751b547f467fe94a754399171213cf43250351e02327307bc7f40c61013d61032261031d6107a9565b610845565b600254610126908290339030906001600160a01b031661095f565b5080fd5b50903461033d578160031936011261033d5760025490516001600160a01b039091168152602090f35b50823461033d57602036600319011261033d573560055580f35b5091346102d45760203660031901126102d45761039f6107a9565b916103a8610a74565b6001600160a01b038316156103c457836103c184610aad565b80f35b51631e4fbdf760e01b8152908101839052602490fd5b50903461033d57602036600319011261033d5760209181906001600160a01b036104026107a9565b16815280845220549051908152f35b50903461033d578160031936011261033d5760035490516001600160a01b039091168152602090f35b50903461033d578160031936011261033d57600080516020610b678339815191525490516001600160a01b039091168152602090f35b50903461033d57606036600319011261033d5761048b6107a9565b90602435906001600160a01b0380831691908284036102d0578151633a98ef3960e01b81528582169560209586838b818b5afa92831561063657908792918a946105ff575b50600354865163778e55f360e01b81526001600160a01b03928316818e019081529290931660208301529394919384921690829081906040015b03915afa9081156105f55787916105c6575b506105329161052a916107fc565b6044356107fc565b916105426064600554048461081c565b928587528685528383882055600554606403606481116105b35791610592608096949260647f6ceab34dc8550787796af13d9da5dfd30e73c264921e6d874a04ebce40a352c1999795049061081c565b938189526001815284838a205582519586528501528301526060820152a180f35b634e487b7160e01b885260118952602488fd5b90508481813d83116105ee575b6105dd81836107c4565b810103126102cc575161053261051c565b503d6105d3565b83513d89823e3d90fd5b9193509181813d811161062f575b61061781836107c4565b8101031261062b575191869161050a6104d0565b8880fd5b503d61060d565b85513d8b823e3d90fd5b5034610147578060031936011261014757610659610a74565b600080516020610b6783398151915280546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b50903461033d57602036600319011261033d576020906106cb61031d6107a9565b9051908152f35b50903461033d57602036600319011261033d5760209181906001600160a01b036106fa6107a9565b1681526001845220549051908152f35b5091346102d457826003193601126102d4575490516001600160a01b03909116815260209150f35b50903461033d578160031936011261033d577ff358f13bbeaf315167e16e69fa1b06ca23f46b7fcc119b65b735a427539b332b61013d826020943381526001865220546101268160018060a01b03600254163390309061095f565b90503461033d578160031936011261033d576020906005548152f35b600435906001600160a01b03821682036107bf57565b600080fd5b90601f8019910116810190811067ffffffffffffffff8211176107e657604052565b634e487b7160e01b600052604160045260246000fd5b8115610806570490565b634e487b7160e01b600052601260045260246000fd5b8181029291811591840414171561082f57565b634e487b7160e01b600052601160045260246000fd5b60048054604051633d3f06c960e11b815233928101929092526001600160a01b038381166024840152602093909184908490604490829086165afa92831561092457600093610930575b501691604051633a98ef3960e01b81528181600481875afa908115610924576000916108f7575b50821580156108ef575b6108e65760006108dd926108e395825252604060002054926107fc565b9061081c565b90565b50505050600090565b5080156108c0565b90508181813d831161091d575b61090e81836107c4565b810103126107bf5751386108b6565b503d610904565b6040513d6000823e3d90fd5b9092508381813d8311610958575b61094881836107c4565b810103126107bf5751913861088f565b503d61093e565b6040516323b872dd60e01b602082019081526001600160a01b039384166024830152938316604482015260648082019590955293845267ffffffffffffffff919060a08501838111868210176107e6576040521692600080938192519082875af13d15610a67573d918211610a5357906109fb91604051916109eb6020601f19601f84011601846107c4565b82523d84602084013e5b84610b03565b908151918215159283610a2b575b505050610a135750565b60249060405190635274afe760e01b82526004820152fd5b81929350906020918101031261033d5760200151908115918215036101475750388080610a09565b634e487b7160e01b83526041600452602483fd5b6109fb91506060906109f5565b600080516020610b67833981519152546001600160a01b03163303610a9557565b60405163118cdaa760e01b8152336004820152602490fd5b600080516020610b6783398151915280546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b90610b2a5750805115610b1857805190602001fd5b604051630a12f52160e11b8152600490fd5b81511580610b5d575b610b3b575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b15610b3356fe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300a2646970667358221220fede1540f8fd11466c753797286acb51abd6cf5ff141f91478941b53a54e4ead64736f6c63430008180033",
}

// L2RewardManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use L2RewardManagerMetaData.ABI instead.
var L2RewardManagerABI = L2RewardManagerMetaData.ABI

// L2RewardManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2RewardManagerMetaData.Bin instead.
var L2RewardManagerBin = L2RewardManagerMetaData.Bin

// DeployL2RewardManager deploys a new Ethereum contract, binding an instance of L2RewardManager to it.
func DeployL2RewardManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2RewardManager, error) {
	parsed, err := L2RewardManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2RewardManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2RewardManager{L2RewardManagerCaller: L2RewardManagerCaller{contract: contract}, L2RewardManagerTransactor: L2RewardManagerTransactor{contract: contract}, L2RewardManagerFilterer: L2RewardManagerFilterer{contract: contract}}, nil
}

// L2RewardManager is an auto generated Go binding around an Ethereum contract.
type L2RewardManager struct {
	L2RewardManagerCaller     // Read-only binding to the contract
	L2RewardManagerTransactor // Write-only binding to the contract
	L2RewardManagerFilterer   // Log filterer for contract events
}

// L2RewardManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2RewardManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2RewardManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2RewardManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2RewardManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2RewardManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2RewardManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2RewardManagerSession struct {
	Contract     *L2RewardManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2RewardManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2RewardManagerCallerSession struct {
	Contract *L2RewardManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L2RewardManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2RewardManagerTransactorSession struct {
	Contract     *L2RewardManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L2RewardManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2RewardManagerRaw struct {
	Contract *L2RewardManager // Generic contract binding to access the raw methods on
}

// L2RewardManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2RewardManagerCallerRaw struct {
	Contract *L2RewardManagerCaller // Generic read-only contract binding to access the raw methods on
}

// L2RewardManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2RewardManagerTransactorRaw struct {
	Contract *L2RewardManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2RewardManager creates a new instance of L2RewardManager, bound to a specific deployed contract.
func NewL2RewardManager(address common.Address, backend bind.ContractBackend) (*L2RewardManager, error) {
	contract, err := bindL2RewardManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2RewardManager{L2RewardManagerCaller: L2RewardManagerCaller{contract: contract}, L2RewardManagerTransactor: L2RewardManagerTransactor{contract: contract}, L2RewardManagerFilterer: L2RewardManagerFilterer{contract: contract}}, nil
}

// NewL2RewardManagerCaller creates a new read-only instance of L2RewardManager, bound to a specific deployed contract.
func NewL2RewardManagerCaller(address common.Address, caller bind.ContractCaller) (*L2RewardManagerCaller, error) {
	contract, err := bindL2RewardManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2RewardManagerCaller{contract: contract}, nil
}

// NewL2RewardManagerTransactor creates a new write-only instance of L2RewardManager, bound to a specific deployed contract.
func NewL2RewardManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*L2RewardManagerTransactor, error) {
	contract, err := bindL2RewardManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2RewardManagerTransactor{contract: contract}, nil
}

// NewL2RewardManagerFilterer creates a new log filterer instance of L2RewardManager, bound to a specific deployed contract.
func NewL2RewardManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*L2RewardManagerFilterer, error) {
	contract, err := bindL2RewardManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2RewardManagerFilterer{contract: contract}, nil
}

// bindL2RewardManager binds a generic wrapper to an already deployed contract.
func bindL2RewardManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2RewardManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2RewardManager *L2RewardManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2RewardManager.Contract.L2RewardManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2RewardManager *L2RewardManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2RewardManager.Contract.L2RewardManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2RewardManager *L2RewardManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2RewardManager.Contract.L2RewardManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2RewardManager *L2RewardManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2RewardManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2RewardManager *L2RewardManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2RewardManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2RewardManager *L2RewardManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2RewardManager.Contract.contract.Transact(opts, method, params...)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_L2RewardManager *L2RewardManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2RewardManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_L2RewardManager *L2RewardManagerSession) Delegation() (common.Address, error) {
	return _L2RewardManager.Contract.Delegation(&_L2RewardManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_L2RewardManager *L2RewardManagerCallerSession) Delegation() (common.Address, error) {
	return _L2RewardManager.Contract.Delegation(&_L2RewardManager.CallOpts)
}

// OperatorRewards is a free data retrieval call binding the contract method 0x41a2b8d6.
//
// Solidity: function operatorRewards(address ) view returns(uint256)
func (_L2RewardManager *L2RewardManagerCaller) OperatorRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2RewardManager.contract.Call(opts, &out, "operatorRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorRewards is a free data retrieval call binding the contract method 0x41a2b8d6.
//
// Solidity: function operatorRewards(address ) view returns(uint256)
func (_L2RewardManager *L2RewardManagerSession) OperatorRewards(arg0 common.Address) (*big.Int, error) {
	return _L2RewardManager.Contract.OperatorRewards(&_L2RewardManager.CallOpts, arg0)
}

// OperatorRewards is a free data retrieval call binding the contract method 0x41a2b8d6.
//
// Solidity: function operatorRewards(address ) view returns(uint256)
func (_L2RewardManager *L2RewardManagerCallerSession) OperatorRewards(arg0 common.Address) (*big.Int, error) {
	return _L2RewardManager.Contract.OperatorRewards(&_L2RewardManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2RewardManager *L2RewardManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2RewardManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2RewardManager *L2RewardManagerSession) Owner() (common.Address, error) {
	return _L2RewardManager.Contract.Owner(&_L2RewardManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2RewardManager *L2RewardManagerCallerSession) Owner() (common.Address, error) {
	return _L2RewardManager.Contract.Owner(&_L2RewardManager.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_L2RewardManager *L2RewardManagerCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2RewardManager.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_L2RewardManager *L2RewardManagerSession) RewardToken() (common.Address, error) {
	return _L2RewardManager.Contract.RewardToken(&_L2RewardManager.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_L2RewardManager *L2RewardManagerCallerSession) RewardToken() (common.Address, error) {
	return _L2RewardManager.Contract.RewardToken(&_L2RewardManager.CallOpts)
}

// StakerPercent is a free data retrieval call binding the contract method 0x1546e980.
//
// Solidity: function stakerPercent() view returns(uint256)
func (_L2RewardManager *L2RewardManagerCaller) StakerPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2RewardManager.contract.Call(opts, &out, "stakerPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerPercent is a free data retrieval call binding the contract method 0x1546e980.
//
// Solidity: function stakerPercent() view returns(uint256)
func (_L2RewardManager *L2RewardManagerSession) StakerPercent() (*big.Int, error) {
	return _L2RewardManager.Contract.StakerPercent(&_L2RewardManager.CallOpts)
}

// StakerPercent is a free data retrieval call binding the contract method 0x1546e980.
//
// Solidity: function stakerPercent() view returns(uint256)
func (_L2RewardManager *L2RewardManagerCallerSession) StakerPercent() (*big.Int, error) {
	return _L2RewardManager.Contract.StakerPercent(&_L2RewardManager.CallOpts)
}

// StakerRewards is a free data retrieval call binding the contract method 0xf0f0b228.
//
// Solidity: function stakerRewards(address ) view returns(uint256)
func (_L2RewardManager *L2RewardManagerCaller) StakerRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2RewardManager.contract.Call(opts, &out, "stakerRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerRewards is a free data retrieval call binding the contract method 0xf0f0b228.
//
// Solidity: function stakerRewards(address ) view returns(uint256)
func (_L2RewardManager *L2RewardManagerSession) StakerRewards(arg0 common.Address) (*big.Int, error) {
	return _L2RewardManager.Contract.StakerRewards(&_L2RewardManager.CallOpts, arg0)
}

// StakerRewards is a free data retrieval call binding the contract method 0xf0f0b228.
//
// Solidity: function stakerRewards(address ) view returns(uint256)
func (_L2RewardManager *L2RewardManagerCallerSession) StakerRewards(arg0 common.Address) (*big.Int, error) {
	return _L2RewardManager.Contract.StakerRewards(&_L2RewardManager.CallOpts, arg0)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_L2RewardManager *L2RewardManagerCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2RewardManager.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_L2RewardManager *L2RewardManagerSession) StrategyManager() (common.Address, error) {
	return _L2RewardManager.Contract.StrategyManager(&_L2RewardManager.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_L2RewardManager *L2RewardManagerCallerSession) StrategyManager() (common.Address, error) {
	return _L2RewardManager.Contract.StrategyManager(&_L2RewardManager.CallOpts)
}

// CalculateFee is a paid mutator transaction binding the contract method 0x7ca87cb6.
//
// Solidity: function calculateFee(address strategy, address operator, uint256 baseFee) returns()
func (_L2RewardManager *L2RewardManagerTransactor) CalculateFee(opts *bind.TransactOpts, strategy common.Address, operator common.Address, baseFee *big.Int) (*types.Transaction, error) {
	return _L2RewardManager.contract.Transact(opts, "calculateFee", strategy, operator, baseFee)
}

// CalculateFee is a paid mutator transaction binding the contract method 0x7ca87cb6.
//
// Solidity: function calculateFee(address strategy, address operator, uint256 baseFee) returns()
func (_L2RewardManager *L2RewardManagerSession) CalculateFee(strategy common.Address, operator common.Address, baseFee *big.Int) (*types.Transaction, error) {
	return _L2RewardManager.Contract.CalculateFee(&_L2RewardManager.TransactOpts, strategy, operator, baseFee)
}

// CalculateFee is a paid mutator transaction binding the contract method 0x7ca87cb6.
//
// Solidity: function calculateFee(address strategy, address operator, uint256 baseFee) returns()
func (_L2RewardManager *L2RewardManagerTransactorSession) CalculateFee(strategy common.Address, operator common.Address, baseFee *big.Int) (*types.Transaction, error) {
	return _L2RewardManager.Contract.CalculateFee(&_L2RewardManager.TransactOpts, strategy, operator, baseFee)
}

// DepositDappLinkToken is a paid mutator transaction binding the contract method 0xfb1bca50.
//
// Solidity: function depositDappLinkToken(uint256 amount) returns(bool)
func (_L2RewardManager *L2RewardManagerTransactor) DepositDappLinkToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _L2RewardManager.contract.Transact(opts, "depositDappLinkToken", amount)
}

// DepositDappLinkToken is a paid mutator transaction binding the contract method 0xfb1bca50.
//
// Solidity: function depositDappLinkToken(uint256 amount) returns(bool)
func (_L2RewardManager *L2RewardManagerSession) DepositDappLinkToken(amount *big.Int) (*types.Transaction, error) {
	return _L2RewardManager.Contract.DepositDappLinkToken(&_L2RewardManager.TransactOpts, amount)
}

// DepositDappLinkToken is a paid mutator transaction binding the contract method 0xfb1bca50.
//
// Solidity: function depositDappLinkToken(uint256 amount) returns(bool)
func (_L2RewardManager *L2RewardManagerTransactorSession) DepositDappLinkToken(amount *big.Int) (*types.Transaction, error) {
	return _L2RewardManager.Contract.DepositDappLinkToken(&_L2RewardManager.TransactOpts, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address initialOwner, address _delegation, address _strategyManager, address _rewardToken) returns()
func (_L2RewardManager *L2RewardManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _delegation common.Address, _strategyManager common.Address, _rewardToken common.Address) (*types.Transaction, error) {
	return _L2RewardManager.contract.Transact(opts, "initialize", initialOwner, _delegation, _strategyManager, _rewardToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address initialOwner, address _delegation, address _strategyManager, address _rewardToken) returns()
func (_L2RewardManager *L2RewardManagerSession) Initialize(initialOwner common.Address, _delegation common.Address, _strategyManager common.Address, _rewardToken common.Address) (*types.Transaction, error) {
	return _L2RewardManager.Contract.Initialize(&_L2RewardManager.TransactOpts, initialOwner, _delegation, _strategyManager, _rewardToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address initialOwner, address _delegation, address _strategyManager, address _rewardToken) returns()
func (_L2RewardManager *L2RewardManagerTransactorSession) Initialize(initialOwner common.Address, _delegation common.Address, _strategyManager common.Address, _rewardToken common.Address) (*types.Transaction, error) {
	return _L2RewardManager.Contract.Initialize(&_L2RewardManager.TransactOpts, initialOwner, _delegation, _strategyManager, _rewardToken)
}

// OperatorClaimReward is a paid mutator transaction binding the contract method 0x2cb08272.
//
// Solidity: function operatorClaimReward() returns(bool)
func (_L2RewardManager *L2RewardManagerTransactor) OperatorClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2RewardManager.contract.Transact(opts, "operatorClaimReward")
}

// OperatorClaimReward is a paid mutator transaction binding the contract method 0x2cb08272.
//
// Solidity: function operatorClaimReward() returns(bool)
func (_L2RewardManager *L2RewardManagerSession) OperatorClaimReward() (*types.Transaction, error) {
	return _L2RewardManager.Contract.OperatorClaimReward(&_L2RewardManager.TransactOpts)
}

// OperatorClaimReward is a paid mutator transaction binding the contract method 0x2cb08272.
//
// Solidity: function operatorClaimReward() returns(bool)
func (_L2RewardManager *L2RewardManagerTransactorSession) OperatorClaimReward() (*types.Transaction, error) {
	return _L2RewardManager.Contract.OperatorClaimReward(&_L2RewardManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2RewardManager *L2RewardManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2RewardManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2RewardManager *L2RewardManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2RewardManager.Contract.RenounceOwnership(&_L2RewardManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2RewardManager *L2RewardManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2RewardManager.Contract.RenounceOwnership(&_L2RewardManager.TransactOpts)
}

// StakerClaimReward is a paid mutator transaction binding the contract method 0xf8401008.
//
// Solidity: function stakerClaimReward(address strategy) returns(bool)
func (_L2RewardManager *L2RewardManagerTransactor) StakerClaimReward(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _L2RewardManager.contract.Transact(opts, "stakerClaimReward", strategy)
}

// StakerClaimReward is a paid mutator transaction binding the contract method 0xf8401008.
//
// Solidity: function stakerClaimReward(address strategy) returns(bool)
func (_L2RewardManager *L2RewardManagerSession) StakerClaimReward(strategy common.Address) (*types.Transaction, error) {
	return _L2RewardManager.Contract.StakerClaimReward(&_L2RewardManager.TransactOpts, strategy)
}

// StakerClaimReward is a paid mutator transaction binding the contract method 0xf8401008.
//
// Solidity: function stakerClaimReward(address strategy) returns(bool)
func (_L2RewardManager *L2RewardManagerTransactorSession) StakerClaimReward(strategy common.Address) (*types.Transaction, error) {
	return _L2RewardManager.Contract.StakerClaimReward(&_L2RewardManager.TransactOpts, strategy)
}

// StakerRewardsAmount is a paid mutator transaction binding the contract method 0x5adbc180.
//
// Solidity: function stakerRewardsAmount(address strategy) returns(uint256)
func (_L2RewardManager *L2RewardManagerTransactor) StakerRewardsAmount(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _L2RewardManager.contract.Transact(opts, "stakerRewardsAmount", strategy)
}

// StakerRewardsAmount is a paid mutator transaction binding the contract method 0x5adbc180.
//
// Solidity: function stakerRewardsAmount(address strategy) returns(uint256)
func (_L2RewardManager *L2RewardManagerSession) StakerRewardsAmount(strategy common.Address) (*types.Transaction, error) {
	return _L2RewardManager.Contract.StakerRewardsAmount(&_L2RewardManager.TransactOpts, strategy)
}

// StakerRewardsAmount is a paid mutator transaction binding the contract method 0x5adbc180.
//
// Solidity: function stakerRewardsAmount(address strategy) returns(uint256)
func (_L2RewardManager *L2RewardManagerTransactorSession) StakerRewardsAmount(strategy common.Address) (*types.Transaction, error) {
	return _L2RewardManager.Contract.StakerRewardsAmount(&_L2RewardManager.TransactOpts, strategy)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2RewardManager *L2RewardManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2RewardManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2RewardManager *L2RewardManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2RewardManager.Contract.TransferOwnership(&_L2RewardManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2RewardManager *L2RewardManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2RewardManager.Contract.TransferOwnership(&_L2RewardManager.TransactOpts, newOwner)
}

// UpdateOperatorAndStakerShareFee is a paid mutator transaction binding the contract method 0xf4d7152b.
//
// Solidity: function updateOperatorAndStakerShareFee(uint256 _stakerPercent) returns()
func (_L2RewardManager *L2RewardManagerTransactor) UpdateOperatorAndStakerShareFee(opts *bind.TransactOpts, _stakerPercent *big.Int) (*types.Transaction, error) {
	return _L2RewardManager.contract.Transact(opts, "updateOperatorAndStakerShareFee", _stakerPercent)
}

// UpdateOperatorAndStakerShareFee is a paid mutator transaction binding the contract method 0xf4d7152b.
//
// Solidity: function updateOperatorAndStakerShareFee(uint256 _stakerPercent) returns()
func (_L2RewardManager *L2RewardManagerSession) UpdateOperatorAndStakerShareFee(_stakerPercent *big.Int) (*types.Transaction, error) {
	return _L2RewardManager.Contract.UpdateOperatorAndStakerShareFee(&_L2RewardManager.TransactOpts, _stakerPercent)
}

// UpdateOperatorAndStakerShareFee is a paid mutator transaction binding the contract method 0xf4d7152b.
//
// Solidity: function updateOperatorAndStakerShareFee(uint256 _stakerPercent) returns()
func (_L2RewardManager *L2RewardManagerTransactorSession) UpdateOperatorAndStakerShareFee(_stakerPercent *big.Int) (*types.Transaction, error) {
	return _L2RewardManager.Contract.UpdateOperatorAndStakerShareFee(&_L2RewardManager.TransactOpts, _stakerPercent)
}

// L2RewardManagerDepositDappLinkTokenIterator is returned from FilterDepositDappLinkToken and is used to iterate over the raw logs and unpacked data for DepositDappLinkToken events raised by the L2RewardManager contract.
type L2RewardManagerDepositDappLinkTokenIterator struct {
	Event *L2RewardManagerDepositDappLinkToken // Event containing the contract specifics and raw log

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
func (it *L2RewardManagerDepositDappLinkTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2RewardManagerDepositDappLinkToken)
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
		it.Event = new(L2RewardManagerDepositDappLinkToken)
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
func (it *L2RewardManagerDepositDappLinkTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2RewardManagerDepositDappLinkTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2RewardManagerDepositDappLinkToken represents a DepositDappLinkToken event raised by the L2RewardManager contract.
type L2RewardManagerDepositDappLinkToken struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositDappLinkToken is a free log retrieval operation binding the contract event 0x08fba1e7a1091d4e4242ae25b8f8291ec7423fba91244b59db8030ee503dfbe8.
//
// Solidity: event DepositDappLinkToken(address sender, uint256 amount)
func (_L2RewardManager *L2RewardManagerFilterer) FilterDepositDappLinkToken(opts *bind.FilterOpts) (*L2RewardManagerDepositDappLinkTokenIterator, error) {

	logs, sub, err := _L2RewardManager.contract.FilterLogs(opts, "DepositDappLinkToken")
	if err != nil {
		return nil, err
	}
	return &L2RewardManagerDepositDappLinkTokenIterator{contract: _L2RewardManager.contract, event: "DepositDappLinkToken", logs: logs, sub: sub}, nil
}

// WatchDepositDappLinkToken is a free log subscription operation binding the contract event 0x08fba1e7a1091d4e4242ae25b8f8291ec7423fba91244b59db8030ee503dfbe8.
//
// Solidity: event DepositDappLinkToken(address sender, uint256 amount)
func (_L2RewardManager *L2RewardManagerFilterer) WatchDepositDappLinkToken(opts *bind.WatchOpts, sink chan<- *L2RewardManagerDepositDappLinkToken) (event.Subscription, error) {

	logs, sub, err := _L2RewardManager.contract.WatchLogs(opts, "DepositDappLinkToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2RewardManagerDepositDappLinkToken)
				if err := _L2RewardManager.contract.UnpackLog(event, "DepositDappLinkToken", log); err != nil {
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

// ParseDepositDappLinkToken is a log parse operation binding the contract event 0x08fba1e7a1091d4e4242ae25b8f8291ec7423fba91244b59db8030ee503dfbe8.
//
// Solidity: event DepositDappLinkToken(address sender, uint256 amount)
func (_L2RewardManager *L2RewardManagerFilterer) ParseDepositDappLinkToken(log types.Log) (*L2RewardManagerDepositDappLinkToken, error) {
	event := new(L2RewardManagerDepositDappLinkToken)
	if err := _L2RewardManager.contract.UnpackLog(event, "DepositDappLinkToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2RewardManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2RewardManager contract.
type L2RewardManagerInitializedIterator struct {
	Event *L2RewardManagerInitialized // Event containing the contract specifics and raw log

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
func (it *L2RewardManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2RewardManagerInitialized)
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
		it.Event = new(L2RewardManagerInitialized)
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
func (it *L2RewardManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2RewardManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2RewardManagerInitialized represents a Initialized event raised by the L2RewardManager contract.
type L2RewardManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L2RewardManager *L2RewardManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2RewardManagerInitializedIterator, error) {

	logs, sub, err := _L2RewardManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2RewardManagerInitializedIterator{contract: _L2RewardManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L2RewardManager *L2RewardManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2RewardManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _L2RewardManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2RewardManagerInitialized)
				if err := _L2RewardManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2RewardManager *L2RewardManagerFilterer) ParseInitialized(log types.Log) (*L2RewardManagerInitialized, error) {
	event := new(L2RewardManagerInitialized)
	if err := _L2RewardManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2RewardManagerOperatorClaimRewardIterator is returned from FilterOperatorClaimReward and is used to iterate over the raw logs and unpacked data for OperatorClaimReward events raised by the L2RewardManager contract.
type L2RewardManagerOperatorClaimRewardIterator struct {
	Event *L2RewardManagerOperatorClaimReward // Event containing the contract specifics and raw log

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
func (it *L2RewardManagerOperatorClaimRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2RewardManagerOperatorClaimReward)
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
		it.Event = new(L2RewardManagerOperatorClaimReward)
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
func (it *L2RewardManagerOperatorClaimRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2RewardManagerOperatorClaimRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2RewardManagerOperatorClaimReward represents a OperatorClaimReward event raised by the L2RewardManager contract.
type L2RewardManagerOperatorClaimReward struct {
	Operator common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorClaimReward is a free log retrieval operation binding the contract event 0xf358f13bbeaf315167e16e69fa1b06ca23f46b7fcc119b65b735a427539b332b.
//
// Solidity: event OperatorClaimReward(address operator, uint256 amount)
func (_L2RewardManager *L2RewardManagerFilterer) FilterOperatorClaimReward(opts *bind.FilterOpts) (*L2RewardManagerOperatorClaimRewardIterator, error) {

	logs, sub, err := _L2RewardManager.contract.FilterLogs(opts, "OperatorClaimReward")
	if err != nil {
		return nil, err
	}
	return &L2RewardManagerOperatorClaimRewardIterator{contract: _L2RewardManager.contract, event: "OperatorClaimReward", logs: logs, sub: sub}, nil
}

// WatchOperatorClaimReward is a free log subscription operation binding the contract event 0xf358f13bbeaf315167e16e69fa1b06ca23f46b7fcc119b65b735a427539b332b.
//
// Solidity: event OperatorClaimReward(address operator, uint256 amount)
func (_L2RewardManager *L2RewardManagerFilterer) WatchOperatorClaimReward(opts *bind.WatchOpts, sink chan<- *L2RewardManagerOperatorClaimReward) (event.Subscription, error) {

	logs, sub, err := _L2RewardManager.contract.WatchLogs(opts, "OperatorClaimReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2RewardManagerOperatorClaimReward)
				if err := _L2RewardManager.contract.UnpackLog(event, "OperatorClaimReward", log); err != nil {
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

// ParseOperatorClaimReward is a log parse operation binding the contract event 0xf358f13bbeaf315167e16e69fa1b06ca23f46b7fcc119b65b735a427539b332b.
//
// Solidity: event OperatorClaimReward(address operator, uint256 amount)
func (_L2RewardManager *L2RewardManagerFilterer) ParseOperatorClaimReward(log types.Log) (*L2RewardManagerOperatorClaimReward, error) {
	event := new(L2RewardManagerOperatorClaimReward)
	if err := _L2RewardManager.contract.UnpackLog(event, "OperatorClaimReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2RewardManagerOperatorStakerRewardIterator is returned from FilterOperatorStakerReward and is used to iterate over the raw logs and unpacked data for OperatorStakerReward events raised by the L2RewardManager contract.
type L2RewardManagerOperatorStakerRewardIterator struct {
	Event *L2RewardManagerOperatorStakerReward // Event containing the contract specifics and raw log

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
func (it *L2RewardManagerOperatorStakerRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2RewardManagerOperatorStakerReward)
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
		it.Event = new(L2RewardManagerOperatorStakerReward)
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
func (it *L2RewardManagerOperatorStakerRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2RewardManagerOperatorStakerRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2RewardManagerOperatorStakerReward represents a OperatorStakerReward event raised by the L2RewardManager contract.
type L2RewardManagerOperatorStakerReward struct {
	Strategy    common.Address
	Operator    common.Address
	StakerFee   *big.Int
	OperatorFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorStakerReward is a free log retrieval operation binding the contract event 0x6ceab34dc8550787796af13d9da5dfd30e73c264921e6d874a04ebce40a352c1.
//
// Solidity: event OperatorStakerReward(address strategy, address operator, uint256 stakerFee, uint256 operatorFee)
func (_L2RewardManager *L2RewardManagerFilterer) FilterOperatorStakerReward(opts *bind.FilterOpts) (*L2RewardManagerOperatorStakerRewardIterator, error) {

	logs, sub, err := _L2RewardManager.contract.FilterLogs(opts, "OperatorStakerReward")
	if err != nil {
		return nil, err
	}
	return &L2RewardManagerOperatorStakerRewardIterator{contract: _L2RewardManager.contract, event: "OperatorStakerReward", logs: logs, sub: sub}, nil
}

// WatchOperatorStakerReward is a free log subscription operation binding the contract event 0x6ceab34dc8550787796af13d9da5dfd30e73c264921e6d874a04ebce40a352c1.
//
// Solidity: event OperatorStakerReward(address strategy, address operator, uint256 stakerFee, uint256 operatorFee)
func (_L2RewardManager *L2RewardManagerFilterer) WatchOperatorStakerReward(opts *bind.WatchOpts, sink chan<- *L2RewardManagerOperatorStakerReward) (event.Subscription, error) {

	logs, sub, err := _L2RewardManager.contract.WatchLogs(opts, "OperatorStakerReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2RewardManagerOperatorStakerReward)
				if err := _L2RewardManager.contract.UnpackLog(event, "OperatorStakerReward", log); err != nil {
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

// ParseOperatorStakerReward is a log parse operation binding the contract event 0x6ceab34dc8550787796af13d9da5dfd30e73c264921e6d874a04ebce40a352c1.
//
// Solidity: event OperatorStakerReward(address strategy, address operator, uint256 stakerFee, uint256 operatorFee)
func (_L2RewardManager *L2RewardManagerFilterer) ParseOperatorStakerReward(log types.Log) (*L2RewardManagerOperatorStakerReward, error) {
	event := new(L2RewardManagerOperatorStakerReward)
	if err := _L2RewardManager.contract.UnpackLog(event, "OperatorStakerReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2RewardManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2RewardManager contract.
type L2RewardManagerOwnershipTransferredIterator struct {
	Event *L2RewardManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2RewardManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2RewardManagerOwnershipTransferred)
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
		it.Event = new(L2RewardManagerOwnershipTransferred)
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
func (it *L2RewardManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2RewardManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2RewardManagerOwnershipTransferred represents a OwnershipTransferred event raised by the L2RewardManager contract.
type L2RewardManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2RewardManager *L2RewardManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2RewardManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2RewardManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2RewardManagerOwnershipTransferredIterator{contract: _L2RewardManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2RewardManager *L2RewardManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2RewardManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2RewardManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2RewardManagerOwnershipTransferred)
				if err := _L2RewardManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2RewardManager *L2RewardManagerFilterer) ParseOwnershipTransferred(log types.Log) (*L2RewardManagerOwnershipTransferred, error) {
	event := new(L2RewardManagerOwnershipTransferred)
	if err := _L2RewardManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2RewardManagerStakerClaimRewardIterator is returned from FilterStakerClaimReward and is used to iterate over the raw logs and unpacked data for StakerClaimReward events raised by the L2RewardManager contract.
type L2RewardManagerStakerClaimRewardIterator struct {
	Event *L2RewardManagerStakerClaimReward // Event containing the contract specifics and raw log

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
func (it *L2RewardManagerStakerClaimRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2RewardManagerStakerClaimReward)
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
		it.Event = new(L2RewardManagerStakerClaimReward)
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
func (it *L2RewardManagerStakerClaimRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2RewardManagerStakerClaimRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2RewardManagerStakerClaimReward represents a StakerClaimReward event raised by the L2RewardManager contract.
type L2RewardManagerStakerClaimReward struct {
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakerClaimReward is a free log retrieval operation binding the contract event 0x7246707380751b547f467fe94a754399171213cf43250351e02327307bc7f40c.
//
// Solidity: event StakerClaimReward(address staker, uint256 amount)
func (_L2RewardManager *L2RewardManagerFilterer) FilterStakerClaimReward(opts *bind.FilterOpts) (*L2RewardManagerStakerClaimRewardIterator, error) {

	logs, sub, err := _L2RewardManager.contract.FilterLogs(opts, "StakerClaimReward")
	if err != nil {
		return nil, err
	}
	return &L2RewardManagerStakerClaimRewardIterator{contract: _L2RewardManager.contract, event: "StakerClaimReward", logs: logs, sub: sub}, nil
}

// WatchStakerClaimReward is a free log subscription operation binding the contract event 0x7246707380751b547f467fe94a754399171213cf43250351e02327307bc7f40c.
//
// Solidity: event StakerClaimReward(address staker, uint256 amount)
func (_L2RewardManager *L2RewardManagerFilterer) WatchStakerClaimReward(opts *bind.WatchOpts, sink chan<- *L2RewardManagerStakerClaimReward) (event.Subscription, error) {

	logs, sub, err := _L2RewardManager.contract.WatchLogs(opts, "StakerClaimReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2RewardManagerStakerClaimReward)
				if err := _L2RewardManager.contract.UnpackLog(event, "StakerClaimReward", log); err != nil {
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

// ParseStakerClaimReward is a log parse operation binding the contract event 0x7246707380751b547f467fe94a754399171213cf43250351e02327307bc7f40c.
//
// Solidity: event StakerClaimReward(address staker, uint256 amount)
func (_L2RewardManager *L2RewardManagerFilterer) ParseStakerClaimReward(log types.Log) (*L2RewardManagerStakerClaimReward, error) {
	event := new(L2RewardManagerStakerClaimReward)
	if err := _L2RewardManager.contract.UnpackLog(event, "StakerClaimReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
