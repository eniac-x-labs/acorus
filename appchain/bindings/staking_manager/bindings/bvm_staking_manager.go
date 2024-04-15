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

// IDETHBatchMint is an auto generated low-level Go binding around an user-defined struct.
type IDETHBatchMint struct {
	Staker common.Address
	Amount *big.Int
}

// IUnstakeRequestsManagerWriterequestsInfo is an auto generated low-level Go binding around an user-defined struct.
type IUnstakeRequestsManagerWriterequestsInfo struct {
	RequestAddress      common.Address
	UnStakeMessageNonce *big.Int
}

// StakingManagerInit is an auto generated low-level Go binding around an user-defined struct.
type StakingManagerInit struct {
	Admin            common.Address
	Manager          common.Address
	AllocatorService common.Address
	InitiatorService common.Address
	WithdrawalWallet common.Address
}

// StakingManagerStorageValidatorParams is an auto generated low-level Go binding around an user-defined struct.
type StakingManagerStorageValidatorParams struct {
	OperatorID            *big.Int
	DepositAmount         *big.Int
	Pubkey                []byte
	WithdrawalCredentials []byte
	Signature             []byte
	DepositDataRoot       [32]byte
}

// StakingManagerMetaData contains all meta data concerning the StakingManager contract.
var StakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOCATOR_SERVICE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INITIATOR_SERVICE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_ALLOWLIST_MANAGER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_ALLOWLIST_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_MANAGER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOP_UP_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocateETH\",\"inputs\":[{\"name\":\"allocateToUnstakeRequestsManager\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allocateToDeposits\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocatedETHForDeposits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimUnstakeRequest\",\"inputs\":[{\"name\":\"requests\",\"type\":\"tuple[]\",\"internalType\":\"structIUnstakeRequestsManagerWrite.requestsInfo[]\",\"components\":[{\"name\":\"requestAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"unStakeMessageNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dETHToETH\",\"inputs\":[{\"name\":\"dETHAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ethToDETH\",\"inputs\":[{\"name\":\"ethAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exchangeAdjustmentRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLocator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractL1ILocator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializationBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"init\",\"type\":\"tuple\",\"internalType\":\"structStakingManager.Init\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"manager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocatorService\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initiatorService\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawalWallet\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateValidatorsWithDeposits\",\"inputs\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structStakingManagerStorage.ValidatorParams[]\",\"components\":[{\"name\":\"operatorID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"withdrawalCredentials\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"expectedDepositRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isStakingAllowlist\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"locator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maximumDETHSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maximumDepositAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumDepositAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumUnstakeBound\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numInitiatedValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveFromUnstakeRequestsManager\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"receiveReturns\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"reclaimAllocatedETHSurplus\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExchangeAdjustmentRate\",\"inputs\":[{\"name\":\"exchangeAdjustmentRate_\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLocator\",\"inputs\":[{\"name\":\"_locator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaximumDETHSupply\",\"inputs\":[{\"name\":\"maximumDETHSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumDepositAmount\",\"inputs\":[{\"name\":\"minimumDepositAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumUnstakeBound\",\"inputs\":[{\"name\":\"minimumUnstakeBound_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStakingAllowlist\",\"inputs\":[{\"name\":\"isStakingAllowlist_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWithdrawalWallet\",\"inputs\":[{\"name\":\"withdrawalWallet_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"batchMints\",\"type\":\"tuple[]\",\"internalType\":\"structIDETH.BatchMint[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"topUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"totalControlled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalDepositedInValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unStakeMessageNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unallocatedETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unstakeRequest\",\"inputs\":[{\"name\":\"dethAmount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"minETHAmount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"l2Strategy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakeRequestInfo\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2strategy\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usedValidators\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawalWallet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AllocatedETHToDeposits\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocatedETHToUnstakeRequestsManager\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtocolConfigChanged\",\"inputs\":[{\"name\":\"setterSelector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"setterSignature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReturnsReceived\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ethAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"dETHAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeLaveAmount\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dETHLocked\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeRequestClaimed\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"l2Strategys\",\"type\":\"address[]\",\"indexed\":true,\"internalType\":\"address[]\"},{\"name\":\"bridge\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeRequested\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"l2Strategy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ethAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"dETHLocked\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"unStakeMessageNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorInitiated\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"operatorID\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"amountDeposited\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DoesNotReceiveETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfiguration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDepositRoot\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWithdrawalCredentialsNotETH1\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}]},{\"type\":\"error\",\"name\":\"InvalidWithdrawalCredentialsWrongAddress\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWithdrawalCredentialsWrongLength\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MathOverflowedMulDiv\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaximumDETHSupplyExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaximumValidatorDepositExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinimumDepositAmountNotSatisfied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinimumStakeBoundNotSatisfied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinimumUnstakeBoundNotSatisfied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinimumValidatorDepositNotSatisfied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotDappLinkBridge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughDepositETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughUnallocatedETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotReturnsAggregator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotUnstakeRequestsManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Paused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PreviouslyUsedValidator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnstakeBelowMinimudETHAmount\",\"inputs\":[{\"name\":\"ethAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedMinimum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x608080604052346100b9577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166100aa57506001600160401b036002600160401b031982821601610065575b6040516132cb9081620000bf8239f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1388080610055565b63f92ee8a960e01b8152600490fd5b600080fdfe608060408181526004918236101561001657600080fd5b600090813560e01c90816301ffc9a714611e8f575080630208e4b514611ddb57806304f36cc214611d255780630633af7614611d02578063080c279a14611ce357806312e9ead6146119ba578063147d36d51461199b5780631943190d1461193757806319efd5c71461190e5780631d2d35ce14611879578063248a9ca31461184157806329d48704146117575780632f2ff15d146116cf5780633101d910146116a657806335ead2a41461168757806336568abe1461164157806337a6c8811461130c5780633937c0b3146112e357806342d3915d146112bc57806349336f0f1461129d5780634a7d80b3146112745780635915ded1146112035780635940d90b146111e657806360a0f628146111c7578063646648df146111a85780636daa01a214610fa05780636fce8ab214610f6557806375796f7614610e9557806378abb49b14610e765780637c957fc8146103775780637dfcdd2914610e57578063808d663f14610da457806389e80ed314610d695780638f656d2214610a535780639010d07c14610a0557806391d14854146109b257806399dd1deb1461091d578063a217fddf14610902578063a5e84562146108c4578063aab483d614610829578063ac1e225714610772578063b91590b214610753578063bb635c6514610734578063c151aa72146106a9578063c2c3c18c14610421578063ca15c873146103ec578063d547741f1461039f578063d8343dcb14610377578063dc29f1de146102e6578063e55d6cc0146102ab578063ea452b6d146102885763ed9daafb1461026057600080fd5b34610285576020366003190112610285575061027e602092356128b8565b9051908152f35b80fd5b5090346102a757816003193601126102a7576020906037549051908152f35b5080fd5b5090346102a757816003193601126102a757602090517f8ea5b4dbd68db0bf23bf4cda958b61a749f8c5aec6f2912d75a03246753ddd168152f35b50919082600319360112610373577f5e4bd437d29fad01c10cdcfff414f0d6b0e84b96d2dade88d780d45b5630696b9081600052600080516020613216833981519152602052806000203360005260205260ff816000205416156103575783610351346036546120a4565b60365580f35b60449350519163e2517d3f60e01b835233908301526024820152fd5b8280fd5b5090346102a757816003193601126102a757905490516001600160a01b039091168152602090f35b509190346103735780600319360112610373576103e891356103e360016103c4611f13565b9383875260008051602061321683398151915260205286200154612a05565b612be9565b5080f35b5091346103735760203660031901126103735760209282913581526000805160206131d6833981519152845220549051908152f35b5090346102a75760803660031901126102a7576001600160401b039083358281116106a557366023820112156106a55780850135602484821161069457602092845195610473858560051b0188611f97565b8387528487019083829560061b84010192368411610690578401915b83831061063f57505087548651636c9c724960e11b81526001600160a01b039a9350915085908290849082908d165afa9081156105fa579089918991610612575b50163303610604578084896104e3612a38565b168751928380926345b09c8d60e11b82525afa9081156105fa5788916105c5575b506105b75787610512612baf565b1693843b156105b35790855196630165f2a760e21b8852608060848901928901525180915260a487019391885b8281106105945789808a8a82828c8183818f8f8035908301526044356044830152606435606483015203925af190811561058b575061057b5750f35b61058490611f53565b6102855780f35b513d84823e3d90fd5b835180518c16875282015186830152948701949281019260010161053f565b8780fd5b84516313d0ff5960e31b8152fd5b90508481813d83116105f3575b6105dc8183611f97565b810103126105b3576105ed9061202d565b38610504565b503d6105d2565b86513d8a823e3d90fd5b845163c68e02cf60e01b8152fd5b6106329150863d8811610638575b61062a8183611f97565b81019061263c565b386104d0565b503d610620565b87833603126106905787518881018181108482111761067c5789528891889161066786611f3f565b8152828601358382015281520192019161048f565b8660418e634e487b7160e01b600052526000fd5b8980fd5b634e487b7160e01b86526041875285fd5b8380fd5b50919082600319360112610373578254815163352a8adf60e21b81526001600160a01b039160209082908690829086165afa90811561072a57859161070b575b501633036106fe5782610351346036546120a4565b51637154fc4360e01b8152fd5b610724915060203d6020116106385761062a8183611f97565b386106e9565b83513d87823e3d90fd5b5090346102a757816003193601126102a7576020906035549051908152f35b5090346102a757816003193601126102a757602090603d549051908152f35b5082346102a757826003193601126102a7578261078d611f13565b6001600160a01b0392604490846107a2612baf565b168451958694859363d4be074f60e01b85528035908501521660248301525afa90811561081f578280926107e1575b5050825191151582526020820152f35b915091508282813d8311610818575b6107fa8183611f97565b810103126102855750602061080e8261202d565b91015183806107d1565b503d6107f0565b83513d84823e3d90fd5b5091346103735760203660031901126103735760008051602061325683398151915290359161085661299d565b82603a5580519260208401526020835261086f83611f7c565b60208151918083528201527f7365744d696e696d756d4465706f736974416d6f756e742875696e7432353629606082015260806020820152806108be63555a41eb60e11b9460808301906124f3565b0390a280f35b5034610285576020366003190112610285576001600160a01b036108e6611f29565b166bffffffffffffffffffffffff60a01b600054161760005580f35b5090346102a757816003193601126102a75751908152602090f35b5091346103735760203660031901126103735760008051602061325683398151915290359161094a61299d565b8260385580519260208401526020835261096383611f7c565b601f8151918083528201527f7365744d696e696d756d556e7374616b65426f756e642875696e743235362900606082015260806020820152806108be6399dd1deb60e01b9460808301906124f3565b509134610373578160031936011261037357816020936109d0611f13565b923581526000805160206132168339815191528552209060018060a01b0316600052825260ff81600020541690519015158152f35b509134610373578160031936011261037357602092610a3d913581526000805160206131d68339815191528452826024359120612fc0565b905491519160018060a01b039160031b1c168152f35b5091346103735760a03660031901126103735781516001600160401b039060a0810182811182821017610d54578452610a8a611f29565b8152610a94611f13565b9060208101918252610aa4611efd565b8186019081526001600160a01b039260643591908483168303610d4f5760608401928352608435938585168503610d4f57608081019485527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009788549060ff828c1c161598821680159081610d47575b6001149081610d3d575b159081610d34575b50610d26575086949391928b86809581948c8e60016001600160401b03198416179055610d07575b505116610b59612f7f565b610b61612f7f565b8c610b6b82612c31565b610cdf575b5050505116610b7e81612cbf565b610ca9575b505116610b8f81612d57565b610c73575b505116610ba081612def565b610c3d575b505116603c5490662386f26fc100006038556801bc16d674ec800000603a55600160a01b916affffffffffffffffffffff60a81b161717603c5543603d55683782dace9d90000000603e5583603f55610bfc578280f35b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a138808280f35b610c6c906000805160206131f683398151915288526000805160206131d6833981519152602052868820612fd8565b5038610ba5565b610ca2906000805160206132368339815191528a526000805160206131d6833981519152602052888a20612fd8565b5038610b94565b610cd8906000805160206132768339815191528c526000805160206131d68339815191526020528a8c20612fd8565b5038610b83565b8280610cfe94526000805160206131d683398151915260205220612fd8565b508b388c610b70565b68ffffffffffffffffff191668010000000000000001178d5538610b4e565b8a5163f92ee8a960e01b8152fd5b90501538610b26565b303b159150610b1e565b8a9150610b14565b600080fd5b604184634e487b7160e01b6000525260246000fd5b5090346102a757816003193601126102a757602090517fdec9d30de0821ad67aa5b141b13a539f584a19f99319e6041698a892b0e795598152f35b5091826003193601126103735782548251631faa859d60e11b81526001600160a01b039160209082908590829086165afa908115610e4d578591610e2e575b50163303610e21575060207f4cbb9d73b003a252cee3f2ee51d8d65a562af35eebb23730dd4a76d68127b3709151348152a1610351346036546120a4565b9051626310df60e31b8152fd5b610e47915060203d6020116106385761062a8183611f97565b38610de3565b84513d87823e3d90fd5b5090346102a757816003193601126102a7576020906036549051908152f35b5090346102a757816003193601126102a757602090603b549051908152f35b5091903461037357602036600319011261037357610eb1611f29565b610eb961299d565b6001600160a01b0316918215610f57575060008051602061325683398151915290826bffffffffffffffffffffffff60a01b603c541617603c55805192602084015260208352610f0883611f7c565b601c8151918083528201527f7365745769746864726177616c57616c6c657428616464726573732900000000606082015260806020820152806108be633abcb7bb60e11b9460808301906124f3565b905163d92e233d60e01b8152fd5b5090346102a757816003193601126102a757602090517f5e4bd437d29fad01c10cdcfff414f0d6b0e84b96d2dade88d780d45b5630696b8152f35b5090346102a757806003193601126102a757823590602435936000805160206132368339815191528060005260209060008051602061321683398151915282528360002033600052825260ff8460002054161561118b57506001600160a01b039082818361100c612a38565b16865192838092637ee56d2f60e11b82525afa908115611181578791611148575b506111385761103c87866120a4565b966036548098116111285761105c879861105683896120a4565b906120e8565b60365581816110ea575b505084611071578580f35b7ffe89805cf5299ef9fbd1d1ddefb8dcc3fa9408064d2ea31e3fca6565768f5217908451868152a16110a1612baf565b16803b156110e6578491835180958193632689dfd360e11b83525af190811561058b57506110d2575b808080808580f35b6110db90611f53565b6102855780386110ca565b8480fd5b816111187f9d04ecb465d2c8754acb798a22293dd26215a1c2f7a2a56607afa215c1aadc77936037546120a4565b6037558651908152a13881611066565b84516396b0c75160e01b81528490fd5b83516313d0ff5960e31b81528390fd5b90508181813d831161117a575b61115f8183611f97565b81010312611176576111709061202d565b3861102d565b8680fd5b503d611155565b85513d89823e3d90fd5b9050604492519163e2517d3f60e01b835233908301526024820152fd5b5090346102a757816003193601126102a757602090603f549051908152f35b5090346102a757816003193601126102a7576020906034549051908152f35b5090346102a757816003193601126102a75760209061027e612683565b5034610285576020366003190112610285578235906001600160401b0382116102855736602383011215610285575061125f602061124d819584602460ff96369301359101611fd3565b8185519382858094519384920161200a565b81016033815203019020541690519015158152f35b5090346102a757816003193601126102a757603c5490516001600160a01b039091168152602090f35b5090346102a757816003193601126102a757602090603e549051908152f35b5090346102a757816003193601126102a75760209060ff603c5460a01c1690519015158152f35b5090346102a757816003193601126102a757602090516000805160206132768339815191528152f35b50919080600319360112610373578135906024356001600160401b0380821161163d573660238301121561163d578185013590811161163d573660248260061b8401011161163d5760018060a01b03918287541684519283636c9c724960e11b92838252818a60209788935afa908115611606579086918b91611620575b501633036116105787848661139d612a38565b16885192838092631ea7ca8960e01b82525afa908115611606578a916115d1575b506115c157603a548034109081156115b7575b506115a7576113df87612518565b928885876113eb612ac9565b168951928380926318160ddd60e01b82525afa801561159d578b9061156e575b6114169150856120a4565b603e541061155e5761142a886036546120a4565b60365585611436612ac9565b1690813b1561155a57908a9291885192639386e19760e01b8452808c89602487019187015252602460448501920190855b8b82821061152a57505050508391838381809403925af18015611520579085939291611507575b50859054169786518099819382525afa9182156114fd577f1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90959688936114de575b5084519687528601521692a280f35b816114f69294503d85116106385761062a8183611f97565b91386114cf565b84513d89823e3d90fd5b61151391929350611f53565b6105b3578290883861148e565b87513d84823e3d90fd5b8395969750600192948c61153f839496611f3f565b1681528b8701358c8201520194019101918d95949392611467565b8a80fd5b8651630935f98160e31b81528990fd5b508581813d8311611596575b6115848183611f97565b81010312610d4f57611416905161140b565b503d61157a565b88513d8d823e3d90fd5b8551630a3287a960e21b81528890fd5b90508710386113d1565b85516313d0ff5960e31b81528890fd5b90508481813d83116115ff575b6115e88183611f97565b81010312610690576115f99061202d565b386113be565b503d6115de565b87513d8c823e3d90fd5b855163c68e02cf60e01b81528890fd5b6116379150863d88116106385761062a8183611f97565b3861138a565b8580fd5b5090346102a757806003193601126102a75761165b611f13565b90336001600160a01b0383160361167857506103e8919235612be9565b5163334bd91960e11b81528390fd5b5090346102a757816003193601126102a7576020906038549051908152f35b5090346102a757816003193601126102a757602090516000805160206132368339815191528152f35b50913461037357816003193601126103735735906116eb611f13565b9082845260008051602061321683398151915260205261171060018286200154612a05565b61171a8284612e87565b611722578380f35b60009283526000805160206131d683398151915260205290912061174f916001600160a01b031690612fd8565b503880808380f35b509190346103735760203660031901126103735781359161ffff83168093036106a55761178261299d565b6103e883116118335761271083116118205750600080516020613256833981519152908261ffff1960395416176039558051926020840152602083526117c783611f7c565b60218151918083528201527f73657445786368616e676541646a7573746d656e74526174652875696e7431366060820152602960f81b608082015260a06020820152806108be630a7521c160e21b9460a08301906124f3565b634e487b7160e01b845260019052602483fd5b905163c52a9bd360e01b8152fd5b509134610373576020366003190112610373578160209360019235815260008051602061321683398151915285522001549051908152f35b509134610373576020366003190112610373576000805160206132568339815191529035916118a661299d565b82603e558051926020840152602083526118bf83611f7c565b601d8151918083528201527f7365744d6178696d756d44455448537570706c792875696e7432353629000000606082015260806020820152806108be630e969ae760e11b9460808301906124f3565b5090346102a757816003193601126102a757602090516000805160206131f68339815191528152f35b509182913461199757826003193601126119975761195361299d565b6001600160a01b03611963612baf565b1691823b1561199257815163596a15a360e11b81529284918491829084905af190811561058b575061057b5750f35b505050fd5b5050fd5b5034610285576020366003190112610285575061027e60209235612518565b5091346103735760803660031901126103735780356001600160801b03928382168092036110e6576024938435818116809103611176576119f9611efd565b6001600160a01b039160643583611a0e612a38565b1694865180966345b09c8d60e11b8252818b6020998a935afa908115611cd9578c91611ca0575b50611c90576038548810611c8157611a4c886128b8565b1691808310611c66575083611a5f612baf565b168a813b15610285578a9460a48b8389958c51968795869462448d8d60e91b8652339086015216809a8401528d60448401528860648401528760848401525af18015611c5c57611c49575b50603f546000198114611c37579060018995949392019081603f55875192835288878401528783015260608201527fdd412332aa89c96943c00eeac315cfc5e887074571f52af163514c8c34cc9dc360803392a3611b06612ac9565b83828a541686519485809263352a8adf60e21b82525afa928315611c2d578993611c0c575b50811694845191848301936323b872dd60e01b8552338a85015216604483015260648201526064815260a081018181106001600160401b03821117611bf857845251611ba6918891829182885af13d15611bf0573d90611b8a82611fb8565b91611b9786519384611f97565b82523d898584013e5b85613152565b8051918215159283611bcf575b505050611bbe578480f35b51635274afe760e01b815291820152fd5b8293509181928101031261117657611be7910161202d565b15388080611bb3565b606090611ba0565b87604188634e487b7160e01b600052526000fd5b82919350611c2690853d87116106385761062a8183611f97565b9290611b2b565b85513d8b823e3d90fd5b634e487b7160e01b8b5260118952898bfd5b611c55909a919a611f53565b9838611aaa565b87513d8d823e3d90fd5b886044918b858a51936347f961e960e11b8552840152820152fd5b865162771ba760e71b81528990fd5b86516313d0ff5960e31b81528990fd5b90508681813d8311611cd2575b611cb78183611f97565b81010312611cce57611cc89061202d565b38611a35565b8b80fd5b503d611cad565b88513d8e823e3d90fd5b5090346102a757816003193601126102a757602090603a549051908152f35b5090346102a757816003193601126102a75760209061ffff603954169051908152f35b5091346103735760203660031901126103735735908115158092036103735760008051602061325683398151915290611d5c61299d565b603c805460ff60a01b191660a085901b60ff60a01b161790558051602080820194909452928352611d8c83611f7c565b60198151918083528201527f7365745374616b696e67416c6c6f776c69737428626f6f6c2900000000000000606082015260806020820152806108be630279b66160e11b9460808301906124f3565b5082346102a757826003193601126102a7578035906001600160401b03908183116106a557366023840112156106a557828101359182116106a5573660248360051b850101116106a5576000805160206131f683398151915280855260008051602061321683398151915260205285852033865260205260ff868620541615611e725784611e6f85856024803592016120f5565b80f35b6044925085519163e2517d3f60e01b835233908301526024820152fd5b90508334610373576020366003190112610373573563ffffffff60e01b81168091036103735760209250635a05180f60e01b8114908115611ed2575b5015158152f35b637965db0b60e01b811491508115611eec575b5083611ecb565b6301ffc9a760e01b14905083611ee5565b604435906001600160a01b0382168203610d4f57565b602435906001600160a01b0382168203610d4f57565b600435906001600160a01b0382168203610d4f57565b35906001600160a01b0382168203610d4f57565b6001600160401b038111611f6657604052565b634e487b7160e01b600052604160045260246000fd5b604081019081106001600160401b03821117611f6657604052565b90601f801991011681019081106001600160401b03821117611f6657604052565b6001600160401b038111611f6657601f01601f191660200190565b929192611fdf82611fb8565b91611fed6040519384611f97565b829481845281830111610d4f578281602093846000960137010152565b60005b83811061201d5750506000910152565b818101518382015260200161200d565b51908115158203610d4f57565b919081101561205c5760051b8101359060be1981360301821215610d4f570190565b634e487b7160e01b600052603260045260246000fd5b903590601e1981360301821215610d4f57018035906001600160401b038211610d4f57602001918136038313610d4f57565b919082018092116120b157565b634e487b7160e01b600052601160045260246000fd5b908060209392818452848401376000828201840152601f01601f1916010190565b919082039182116120b157565b90916001600160a01b039081612109612a38565b16604094855192838093637750955b60e11b8252602094859160049788915afa9081156122a3576000916124be575b506124ae5781156124a55783838661214e612a8f565b1689519283809263c5f2892f60e01b82525afa9081156122a357600091612478575b5080910361246257506000805b8281106122be57506037548082116122ae579061219d816121a8936120e8565b6037556034546120a4565b6034556121b7816035546120a4565b60355560005b8181106121cd5750505050505050565b6121d881838861203a565b856121e1612a8f565b166121ee89830183612072565b90926121fd6060820182612072565b909360809161220e83850185612072565b9091833b15610d4f578f978d928d9461225d61226c9461224960009d519e8f9d8e9c8d9a6304512a2360e31b8c528b015260848a01916120c7565b9060031995868984030160248a01526120c7565b928584030160448601526120c7565b60a0860135606483015203930135905af180156122a3579060019291612294575b50016121bd565b61229d90611f53565b3861228d565b88513d6000823e3d90fd5b87516307bb2bd760e21b81528590fd5b60ff6122cb82858a61203a565b92898401906122da8286612072565b8c519491819086378401938881603396878152030190205416612452578685013592603a54840361244257606061231381880188612072565b908a820361242b57600c91808311610d4f5781356001600160a01b0319166001600160f81b0319810161241057508b11610d4f570135901c8a603c541681036123fa57506123ea8c6123b38686957f15f16c2e13e50235799a97b981bf4a66c8cd86051f06aca745c5ff26f39b330e958d61239260019c9b9a8e612072565b928388519485938437820190815203019020805460ff19168b1790556120a4565b976123d86123cb6123c48784612072565b3691611fd3565b8c81519101209582612072565b939083519484869586528501916120c7565b958b83015235940390a30161217d565b896024918e5191631b4d561960e01b8352820152fd5b90508f602493508d925051916308ebf56560e01b8352820152fd5b508a6024918f5191639b0ec52760e01b8352820152fd5b8b516305cacc5560e21b81528990fd5b8a5163932c5b0d60e01b81528890fd5b83602491885191631497ae9360e01b8352820152fd5b90508381813d831161249e575b61248f8183611f97565b81010312610d4f575138612170565b503d612485565b50505050505050565b86516313d0ff5960e31b81528490fd5b90508381813d83116124ec575b6124d58183611f97565b81010312610d4f576124e69061202d565b38612138565b503d6124cb565b9060209161250c8151809281855285808601910161200a565b601f01601f1916010190565b6001600160a01b0380612529612ac9565b169060405190816318160ddd60e01b9384825281600460209586935afa9081156125fd5760009161260f575b5015612609578190612565612ac9565b169260046040518095819382525afa9081156125fd576000916125d1575b50603954612710925061ffff90811683038181116120b15716908181029181830414901517156120b1576125b5612683565b8281029281840414901517156120b1576125ce92612b03565b90565b82813d83116125f6575b6125e58183611f97565b810103126102855750518038612583565b503d6125db565b6040513d6000823e3d90fd5b50505090565b90508281813d8311612635575b6126268183611f97565b81010312610d4f575138612555565b503d61261c565b90816020910312610d4f57516001600160a01b0381168103610d4f5790565b51906001600160401b0382168203610d4f57565b51906001600160801b0382168203610d4f57565b600080546040805163159761e360e21b815292939290916001600160a01b039160209182908290600490829087165afa9081156128ae578691612891575b508284519163079d004d60e51b835282600481610100948594165afa9182156111815787926127ae575b50509080612736600493612704603654603754906120a4565b60c061272c603454926127266001600160801b03948560e088015116906120e8565b906120a4565b92015116906120a4565b9361273f612baf565b168551938480926316d3df1560e31b82525afa9384156127a55750859361276e575b50506125ce9293506120a4565b9080929350813d831161279e575b6127868183611f97565b810103126106a5576125ce9293505190839238612761565b503d61277c565b513d87823e3d90fd5b9080925081813d831161288a575b6127c68183611f97565b810103126111765784519182018281106001600160401b0382111761287657839261286960e060049694612736948a526127ff8161265b565b845261280c86820161265b565b8685015261281b8a820161265b565b8a85015261282b6060820161265b565b606085015261283c6080820161266f565b608085015261284d60a0820161266f565b60a085015261285e60c0820161266f565b60c08501520161266f565b60e08201529293506126eb565b634e487b7160e01b88526041600452602488fd5b503d6127bc565b6128a89150823d84116106385761062a8183611f97565b386126c1565b84513d88823e3d90fd5b6001600160a01b0390816128ca612ac9565b169160405191826318160ddd60e01b9485825281600460209687935afa9081156125fd57600091612970575b50156129695782612905612683565b9261290e612ac9565b169460046040518097819382525afa9283156125fd57600093612937575b506125ce9350612b03565b90925083813d8311612962575b61294e8183611f97565b81010312610d4f576125ce9251913861292c565b503d612944565b9250505090565b90508381813d8311612996575b6129878183611f97565b81010312610d4f5751386128f6565b503d61297d565b3360009081527fc5abc2d863e77a76627fe1702320d3afc3f93a9ad30aebcafab5d12854da5c0f60205260409020546000805160206132768339815191529060ff16156129e75750565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b8060005260008051602061321683398151915260205260406000203360005260205260ff60406000205416156129e75750565b600054604051639fd0506d60e01b81526001600160a01b03916020908290600490829086165afa9081156125fd57600091612a7257501690565b612a8b915060203d6020116106385761062a8183611f97565b1690565b60005460405163e94ad65b60e01b81526001600160a01b03916020908290600490829086165afa9081156125fd57600091612a7257501690565b6000546040516304efe2eb60e31b81526001600160a01b03916020908290600490829086165afa9081156125fd57600091612a7257501690565b90918282029160001984820993838086109503948086039514612b8b5784831115612b795782910981600003821680920460028082600302188083028203028083028203028083028203028083028203028083028203028092029003029360018380600003040190848311900302920304170290565b60405163227bc15360e01b8152600490fd5b505080925015612b99570490565b634e487b7160e01b600052601260045260246000fd5b60005460405163352a8adf60e21b81526001600160a01b03916020908290600490829086165afa9081156125fd57600091612a7257501690565b612bf38282612efc565b9182612bfe57505090565b60009182526000805160206131d68339815191526020526040909120612c2d916001600160a01b031690613062565b5090565b6001600160a01b031660008181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260408120549091906000805160206132168339815191529060ff16612cba578280526020526040822081835260205260408220600160ff1982541617905533916000805160206131b68339815191528180a4600190565b505090565b6001600160a01b031660008181527fc5abc2d863e77a76627fe1702320d3afc3f93a9ad30aebcafab5d12854da5c0f6020526040812054909190600080516020613276833981519152906000805160206132168339815191529060ff16612609578184526020526040832082845260205260408320600160ff198254161790556000805160206131b6833981519152339380a4600190565b6001600160a01b031660008181527ff79084f49a5c4fa4c48f70bba1e67b61c2b9ca8b3d302dc944c028fdea010b826020526040812054909190600080516020613236833981519152906000805160206132168339815191529060ff16612609578184526020526040832082845260205260408320600160ff198254161790556000805160206131b6833981519152339380a4600190565b6001600160a01b031660008181527ff1e23661530d14d05c9291333c54312223931d3f1ab2285de8cf548f5a18240d60205260408120549091906000805160206131f6833981519152906000805160206132168339815191529060ff16612609578184526020526040832082845260205260408320600160ff198254161790556000805160206131b6833981519152339380a4600190565b9060009180835260008051602061321683398151915280602052604084209260018060a01b03169283855260205260ff60408520541615600014612609578184526020526040832082845260205260408320600160ff198254161790556000805160206131b6833981519152339380a4600190565b9060009180835260008051602061321683398151915280602052604084209260018060a01b03169283855260205260ff60408520541660001461260957818452602052604083208284526020526040832060ff1981541690557ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b339380a4600190565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c1615612fae57565b604051631afcd79f60e31b8152600490fd5b805482101561205c5760005260206000200190600090565b9190600183016000908282528060205260408220541560001461305c57845494680100000000000000008610156130485783613038613021886001604098999a01855584612fc0565b819391549060031b91821b91600019901b19161790565b9055549382526020522055600190565b634e487b7160e01b83526041600452602483fd5b50925050565b9060018201906000928184528260205260408420549081151560001461314b576000199180830181811161313757825490848201918211613123578181036130ee575b505050805480156130da578201916130bd8383612fc0565b909182549160031b1b191690555582526020526040812055600190565b634e487b7160e01b86526031600452602486fd5b61310e6130fe6130219386612fc0565b90549060031b1c92839286612fc0565b905586528460205260408620553880806130a5565b634e487b7160e01b88526011600452602488fd5b634e487b7160e01b87526011600452602487fd5b5050505090565b90613179575080511561316757805190602001fd5b604051630a12f52160e11b8152600490fd5b815115806131ac575b61318a575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561318256fe2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0dc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000e30bb2df90b65284acd0e8b5ebe3483bb2bbe65a08e43f0f9e8300fd8607ee1102dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800e6ef7125bfa79685f3bd2e4c4cea243c1e988ebbc0801ab7641ae36b9e2c529101d854e8dde9402801a4c6f2840193465752abfad61e0bb7c4258d526ae42e74a6b5d83d32632203555cb9b2c2f68a8d94da48cadd9266ac0d17babedb52ea5ba264697066735822122058bf603d3553538fd8f82dcbafa10e5cb8b11a7da0f17de8ce19d62a9799157c64736f6c63430008180033",
}

// StakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingManagerMetaData.ABI instead.
var StakingManagerABI = StakingManagerMetaData.ABI

// StakingManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingManagerMetaData.Bin instead.
var StakingManagerBin = StakingManagerMetaData.Bin

// DeployStakingManager deploys a new Ethereum contract, binding an instance of StakingManager to it.
func DeployStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StakingManager, error) {
	parsed, err := StakingManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// StakingManager is an auto generated Go binding around an Ethereum contract.
type StakingManager struct {
	StakingManagerCaller     // Read-only binding to the contract
	StakingManagerTransactor // Write-only binding to the contract
	StakingManagerFilterer   // Log filterer for contract events
}

// StakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingManagerSession struct {
	Contract     *StakingManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingManagerCallerSession struct {
	Contract *StakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingManagerTransactorSession struct {
	Contract     *StakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingManagerRaw struct {
	Contract *StakingManager // Generic contract binding to access the raw methods on
}

// StakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingManagerCallerRaw struct {
	Contract *StakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingManagerTransactorRaw struct {
	Contract *StakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingManager creates a new instance of StakingManager, bound to a specific deployed contract.
func NewStakingManager(address common.Address, backend bind.ContractBackend) (*StakingManager, error) {
	contract, err := bindStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// NewStakingManagerCaller creates a new read-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*StakingManagerCaller, error) {
	contract, err := bindStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerCaller{contract: contract}, nil
}

// NewStakingManagerTransactor creates a new write-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingManagerTransactor, error) {
	contract, err := bindStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerTransactor{contract: contract}, nil
}

// NewStakingManagerFilterer creates a new log filterer instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingManagerFilterer, error) {
	contract, err := bindStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingManagerFilterer{contract: contract}, nil
}

// bindStakingManager binds a generic wrapper to an already deployed contract.
func bindStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.StakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transact(opts, method, params...)
}

// ALLOCATORSERVICEROLE is a free data retrieval call binding the contract method 0x3101d910.
//
// Solidity: function ALLOCATOR_SERVICE_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) ALLOCATORSERVICEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "ALLOCATOR_SERVICE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ALLOCATORSERVICEROLE is a free data retrieval call binding the contract method 0x3101d910.
//
// Solidity: function ALLOCATOR_SERVICE_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) ALLOCATORSERVICEROLE() ([32]byte, error) {
	return _StakingManager.Contract.ALLOCATORSERVICEROLE(&_StakingManager.CallOpts)
}

// ALLOCATORSERVICEROLE is a free data retrieval call binding the contract method 0x3101d910.
//
// Solidity: function ALLOCATOR_SERVICE_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) ALLOCATORSERVICEROLE() ([32]byte, error) {
	return _StakingManager.Contract.ALLOCATORSERVICEROLE(&_StakingManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingManager.Contract.DEFAULTADMINROLE(&_StakingManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingManager.Contract.DEFAULTADMINROLE(&_StakingManager.CallOpts)
}

// INITIATORSERVICEROLE is a free data retrieval call binding the contract method 0x19efd5c7.
//
// Solidity: function INITIATOR_SERVICE_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) INITIATORSERVICEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "INITIATOR_SERVICE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITIATORSERVICEROLE is a free data retrieval call binding the contract method 0x19efd5c7.
//
// Solidity: function INITIATOR_SERVICE_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) INITIATORSERVICEROLE() ([32]byte, error) {
	return _StakingManager.Contract.INITIATORSERVICEROLE(&_StakingManager.CallOpts)
}

// INITIATORSERVICEROLE is a free data retrieval call binding the contract method 0x19efd5c7.
//
// Solidity: function INITIATOR_SERVICE_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) INITIATORSERVICEROLE() ([32]byte, error) {
	return _StakingManager.Contract.INITIATORSERVICEROLE(&_StakingManager.CallOpts)
}

// STAKINGALLOWLISTMANAGERROLE is a free data retrieval call binding the contract method 0xe55d6cc0.
//
// Solidity: function STAKING_ALLOWLIST_MANAGER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) STAKINGALLOWLISTMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "STAKING_ALLOWLIST_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGALLOWLISTMANAGERROLE is a free data retrieval call binding the contract method 0xe55d6cc0.
//
// Solidity: function STAKING_ALLOWLIST_MANAGER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) STAKINGALLOWLISTMANAGERROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGALLOWLISTMANAGERROLE(&_StakingManager.CallOpts)
}

// STAKINGALLOWLISTMANAGERROLE is a free data retrieval call binding the contract method 0xe55d6cc0.
//
// Solidity: function STAKING_ALLOWLIST_MANAGER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) STAKINGALLOWLISTMANAGERROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGALLOWLISTMANAGERROLE(&_StakingManager.CallOpts)
}

// STAKINGALLOWLISTROLE is a free data retrieval call binding the contract method 0x89e80ed3.
//
// Solidity: function STAKING_ALLOWLIST_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) STAKINGALLOWLISTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "STAKING_ALLOWLIST_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGALLOWLISTROLE is a free data retrieval call binding the contract method 0x89e80ed3.
//
// Solidity: function STAKING_ALLOWLIST_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) STAKINGALLOWLISTROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGALLOWLISTROLE(&_StakingManager.CallOpts)
}

// STAKINGALLOWLISTROLE is a free data retrieval call binding the contract method 0x89e80ed3.
//
// Solidity: function STAKING_ALLOWLIST_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) STAKINGALLOWLISTROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGALLOWLISTROLE(&_StakingManager.CallOpts)
}

// STAKINGMANAGERROLE is a free data retrieval call binding the contract method 0x3937c0b3.
//
// Solidity: function STAKING_MANAGER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) STAKINGMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "STAKING_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMANAGERROLE is a free data retrieval call binding the contract method 0x3937c0b3.
//
// Solidity: function STAKING_MANAGER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) STAKINGMANAGERROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGMANAGERROLE(&_StakingManager.CallOpts)
}

// STAKINGMANAGERROLE is a free data retrieval call binding the contract method 0x3937c0b3.
//
// Solidity: function STAKING_MANAGER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) STAKINGMANAGERROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGMANAGERROLE(&_StakingManager.CallOpts)
}

// TOPUPROLE is a free data retrieval call binding the contract method 0x6fce8ab2.
//
// Solidity: function TOP_UP_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) TOPUPROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "TOP_UP_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOPUPROLE is a free data retrieval call binding the contract method 0x6fce8ab2.
//
// Solidity: function TOP_UP_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) TOPUPROLE() ([32]byte, error) {
	return _StakingManager.Contract.TOPUPROLE(&_StakingManager.CallOpts)
}

// TOPUPROLE is a free data retrieval call binding the contract method 0x6fce8ab2.
//
// Solidity: function TOP_UP_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) TOPUPROLE() ([32]byte, error) {
	return _StakingManager.Contract.TOPUPROLE(&_StakingManager.CallOpts)
}

// AllocatedETHForDeposits is a free data retrieval call binding the contract method 0xea452b6d.
//
// Solidity: function allocatedETHForDeposits() view returns(uint256)
func (_StakingManager *StakingManagerCaller) AllocatedETHForDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "allocatedETHForDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllocatedETHForDeposits is a free data retrieval call binding the contract method 0xea452b6d.
//
// Solidity: function allocatedETHForDeposits() view returns(uint256)
func (_StakingManager *StakingManagerSession) AllocatedETHForDeposits() (*big.Int, error) {
	return _StakingManager.Contract.AllocatedETHForDeposits(&_StakingManager.CallOpts)
}

// AllocatedETHForDeposits is a free data retrieval call binding the contract method 0xea452b6d.
//
// Solidity: function allocatedETHForDeposits() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) AllocatedETHForDeposits() (*big.Int, error) {
	return _StakingManager.Contract.AllocatedETHForDeposits(&_StakingManager.CallOpts)
}

// ExchangeAdjustmentRate is a free data retrieval call binding the contract method 0x0633af76.
//
// Solidity: function exchangeAdjustmentRate() view returns(uint16)
func (_StakingManager *StakingManagerCaller) ExchangeAdjustmentRate(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "exchangeAdjustmentRate")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ExchangeAdjustmentRate is a free data retrieval call binding the contract method 0x0633af76.
//
// Solidity: function exchangeAdjustmentRate() view returns(uint16)
func (_StakingManager *StakingManagerSession) ExchangeAdjustmentRate() (uint16, error) {
	return _StakingManager.Contract.ExchangeAdjustmentRate(&_StakingManager.CallOpts)
}

// ExchangeAdjustmentRate is a free data retrieval call binding the contract method 0x0633af76.
//
// Solidity: function exchangeAdjustmentRate() view returns(uint16)
func (_StakingManager *StakingManagerCallerSession) ExchangeAdjustmentRate() (uint16, error) {
	return _StakingManager.Contract.ExchangeAdjustmentRate(&_StakingManager.CallOpts)
}

// GetLocator is a free data retrieval call binding the contract method 0xd8343dcb.
//
// Solidity: function getLocator() view returns(address)
func (_StakingManager *StakingManagerCaller) GetLocator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getLocator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLocator is a free data retrieval call binding the contract method 0xd8343dcb.
//
// Solidity: function getLocator() view returns(address)
func (_StakingManager *StakingManagerSession) GetLocator() (common.Address, error) {
	return _StakingManager.Contract.GetLocator(&_StakingManager.CallOpts)
}

// GetLocator is a free data retrieval call binding the contract method 0xd8343dcb.
//
// Solidity: function getLocator() view returns(address)
func (_StakingManager *StakingManagerCallerSession) GetLocator() (common.Address, error) {
	return _StakingManager.Contract.GetLocator(&_StakingManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingManager *StakingManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingManager *StakingManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingManager.Contract.GetRoleAdmin(&_StakingManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingManager.Contract.GetRoleAdmin(&_StakingManager.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StakingManager *StakingManagerCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StakingManager *StakingManagerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _StakingManager.Contract.GetRoleMember(&_StakingManager.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StakingManager *StakingManagerCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _StakingManager.Contract.GetRoleMember(&_StakingManager.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _StakingManager.Contract.GetRoleMemberCount(&_StakingManager.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _StakingManager.Contract.GetRoleMemberCount(&_StakingManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingManager *StakingManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingManager *StakingManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingManager.Contract.HasRole(&_StakingManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingManager.Contract.HasRole(&_StakingManager.CallOpts, role, account)
}

// InitializationBlockNumber is a free data retrieval call binding the contract method 0xb91590b2.
//
// Solidity: function initializationBlockNumber() view returns(uint256)
func (_StakingManager *StakingManagerCaller) InitializationBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "initializationBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializationBlockNumber is a free data retrieval call binding the contract method 0xb91590b2.
//
// Solidity: function initializationBlockNumber() view returns(uint256)
func (_StakingManager *StakingManagerSession) InitializationBlockNumber() (*big.Int, error) {
	return _StakingManager.Contract.InitializationBlockNumber(&_StakingManager.CallOpts)
}

// InitializationBlockNumber is a free data retrieval call binding the contract method 0xb91590b2.
//
// Solidity: function initializationBlockNumber() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) InitializationBlockNumber() (*big.Int, error) {
	return _StakingManager.Contract.InitializationBlockNumber(&_StakingManager.CallOpts)
}

// IsStakingAllowlist is a free data retrieval call binding the contract method 0x42d3915d.
//
// Solidity: function isStakingAllowlist() view returns(bool)
func (_StakingManager *StakingManagerCaller) IsStakingAllowlist(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "isStakingAllowlist")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStakingAllowlist is a free data retrieval call binding the contract method 0x42d3915d.
//
// Solidity: function isStakingAllowlist() view returns(bool)
func (_StakingManager *StakingManagerSession) IsStakingAllowlist() (bool, error) {
	return _StakingManager.Contract.IsStakingAllowlist(&_StakingManager.CallOpts)
}

// IsStakingAllowlist is a free data retrieval call binding the contract method 0x42d3915d.
//
// Solidity: function isStakingAllowlist() view returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsStakingAllowlist() (bool, error) {
	return _StakingManager.Contract.IsStakingAllowlist(&_StakingManager.CallOpts)
}

// Locator is a free data retrieval call binding the contract method 0x7c957fc8.
//
// Solidity: function locator() view returns(address)
func (_StakingManager *StakingManagerCaller) Locator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "locator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Locator is a free data retrieval call binding the contract method 0x7c957fc8.
//
// Solidity: function locator() view returns(address)
func (_StakingManager *StakingManagerSession) Locator() (common.Address, error) {
	return _StakingManager.Contract.Locator(&_StakingManager.CallOpts)
}

// Locator is a free data retrieval call binding the contract method 0x7c957fc8.
//
// Solidity: function locator() view returns(address)
func (_StakingManager *StakingManagerCallerSession) Locator() (common.Address, error) {
	return _StakingManager.Contract.Locator(&_StakingManager.CallOpts)
}

// MaximumDETHSupply is a free data retrieval call binding the contract method 0x49336f0f.
//
// Solidity: function maximumDETHSupply() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MaximumDETHSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "maximumDETHSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumDETHSupply is a free data retrieval call binding the contract method 0x49336f0f.
//
// Solidity: function maximumDETHSupply() view returns(uint256)
func (_StakingManager *StakingManagerSession) MaximumDETHSupply() (*big.Int, error) {
	return _StakingManager.Contract.MaximumDETHSupply(&_StakingManager.CallOpts)
}

// MaximumDETHSupply is a free data retrieval call binding the contract method 0x49336f0f.
//
// Solidity: function maximumDETHSupply() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MaximumDETHSupply() (*big.Int, error) {
	return _StakingManager.Contract.MaximumDETHSupply(&_StakingManager.CallOpts)
}

// MaximumDepositAmount is a free data retrieval call binding the contract method 0x78abb49b.
//
// Solidity: function maximumDepositAmount() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MaximumDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "maximumDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumDepositAmount is a free data retrieval call binding the contract method 0x78abb49b.
//
// Solidity: function maximumDepositAmount() view returns(uint256)
func (_StakingManager *StakingManagerSession) MaximumDepositAmount() (*big.Int, error) {
	return _StakingManager.Contract.MaximumDepositAmount(&_StakingManager.CallOpts)
}

// MaximumDepositAmount is a free data retrieval call binding the contract method 0x78abb49b.
//
// Solidity: function maximumDepositAmount() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MaximumDepositAmount() (*big.Int, error) {
	return _StakingManager.Contract.MaximumDepositAmount(&_StakingManager.CallOpts)
}

// MinimumDepositAmount is a free data retrieval call binding the contract method 0x080c279a.
//
// Solidity: function minimumDepositAmount() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MinimumDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "minimumDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumDepositAmount is a free data retrieval call binding the contract method 0x080c279a.
//
// Solidity: function minimumDepositAmount() view returns(uint256)
func (_StakingManager *StakingManagerSession) MinimumDepositAmount() (*big.Int, error) {
	return _StakingManager.Contract.MinimumDepositAmount(&_StakingManager.CallOpts)
}

// MinimumDepositAmount is a free data retrieval call binding the contract method 0x080c279a.
//
// Solidity: function minimumDepositAmount() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MinimumDepositAmount() (*big.Int, error) {
	return _StakingManager.Contract.MinimumDepositAmount(&_StakingManager.CallOpts)
}

// MinimumUnstakeBound is a free data retrieval call binding the contract method 0x35ead2a4.
//
// Solidity: function minimumUnstakeBound() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MinimumUnstakeBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "minimumUnstakeBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumUnstakeBound is a free data retrieval call binding the contract method 0x35ead2a4.
//
// Solidity: function minimumUnstakeBound() view returns(uint256)
func (_StakingManager *StakingManagerSession) MinimumUnstakeBound() (*big.Int, error) {
	return _StakingManager.Contract.MinimumUnstakeBound(&_StakingManager.CallOpts)
}

// MinimumUnstakeBound is a free data retrieval call binding the contract method 0x35ead2a4.
//
// Solidity: function minimumUnstakeBound() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MinimumUnstakeBound() (*big.Int, error) {
	return _StakingManager.Contract.MinimumUnstakeBound(&_StakingManager.CallOpts)
}

// NumInitiatedValidators is a free data retrieval call binding the contract method 0xbb635c65.
//
// Solidity: function numInitiatedValidators() view returns(uint256)
func (_StakingManager *StakingManagerCaller) NumInitiatedValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "numInitiatedValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumInitiatedValidators is a free data retrieval call binding the contract method 0xbb635c65.
//
// Solidity: function numInitiatedValidators() view returns(uint256)
func (_StakingManager *StakingManagerSession) NumInitiatedValidators() (*big.Int, error) {
	return _StakingManager.Contract.NumInitiatedValidators(&_StakingManager.CallOpts)
}

// NumInitiatedValidators is a free data retrieval call binding the contract method 0xbb635c65.
//
// Solidity: function numInitiatedValidators() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) NumInitiatedValidators() (*big.Int, error) {
	return _StakingManager.Contract.NumInitiatedValidators(&_StakingManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingManager *StakingManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingManager *StakingManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingManager.Contract.SupportsInterface(&_StakingManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingManager.Contract.SupportsInterface(&_StakingManager.CallOpts, interfaceId)
}

// TotalDepositedInValidators is a free data retrieval call binding the contract method 0x60a0f628.
//
// Solidity: function totalDepositedInValidators() view returns(uint256)
func (_StakingManager *StakingManagerCaller) TotalDepositedInValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "totalDepositedInValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDepositedInValidators is a free data retrieval call binding the contract method 0x60a0f628.
//
// Solidity: function totalDepositedInValidators() view returns(uint256)
func (_StakingManager *StakingManagerSession) TotalDepositedInValidators() (*big.Int, error) {
	return _StakingManager.Contract.TotalDepositedInValidators(&_StakingManager.CallOpts)
}

// TotalDepositedInValidators is a free data retrieval call binding the contract method 0x60a0f628.
//
// Solidity: function totalDepositedInValidators() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) TotalDepositedInValidators() (*big.Int, error) {
	return _StakingManager.Contract.TotalDepositedInValidators(&_StakingManager.CallOpts)
}

// UnStakeMessageNonce is a free data retrieval call binding the contract method 0x646648df.
//
// Solidity: function unStakeMessageNonce() view returns(uint256)
func (_StakingManager *StakingManagerCaller) UnStakeMessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "unStakeMessageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnStakeMessageNonce is a free data retrieval call binding the contract method 0x646648df.
//
// Solidity: function unStakeMessageNonce() view returns(uint256)
func (_StakingManager *StakingManagerSession) UnStakeMessageNonce() (*big.Int, error) {
	return _StakingManager.Contract.UnStakeMessageNonce(&_StakingManager.CallOpts)
}

// UnStakeMessageNonce is a free data retrieval call binding the contract method 0x646648df.
//
// Solidity: function unStakeMessageNonce() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) UnStakeMessageNonce() (*big.Int, error) {
	return _StakingManager.Contract.UnStakeMessageNonce(&_StakingManager.CallOpts)
}

// UnallocatedETH is a free data retrieval call binding the contract method 0x7dfcdd29.
//
// Solidity: function unallocatedETH() view returns(uint256)
func (_StakingManager *StakingManagerCaller) UnallocatedETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "unallocatedETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnallocatedETH is a free data retrieval call binding the contract method 0x7dfcdd29.
//
// Solidity: function unallocatedETH() view returns(uint256)
func (_StakingManager *StakingManagerSession) UnallocatedETH() (*big.Int, error) {
	return _StakingManager.Contract.UnallocatedETH(&_StakingManager.CallOpts)
}

// UnallocatedETH is a free data retrieval call binding the contract method 0x7dfcdd29.
//
// Solidity: function unallocatedETH() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) UnallocatedETH() (*big.Int, error) {
	return _StakingManager.Contract.UnallocatedETH(&_StakingManager.CallOpts)
}

// UnstakeRequestInfo is a free data retrieval call binding the contract method 0xac1e2257.
//
// Solidity: function unstakeRequestInfo(uint256 destChainId, address l2strategy) view returns(bool, uint256)
func (_StakingManager *StakingManagerCaller) UnstakeRequestInfo(opts *bind.CallOpts, destChainId *big.Int, l2strategy common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "unstakeRequestInfo", destChainId, l2strategy)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// UnstakeRequestInfo is a free data retrieval call binding the contract method 0xac1e2257.
//
// Solidity: function unstakeRequestInfo(uint256 destChainId, address l2strategy) view returns(bool, uint256)
func (_StakingManager *StakingManagerSession) UnstakeRequestInfo(destChainId *big.Int, l2strategy common.Address) (bool, *big.Int, error) {
	return _StakingManager.Contract.UnstakeRequestInfo(&_StakingManager.CallOpts, destChainId, l2strategy)
}

// UnstakeRequestInfo is a free data retrieval call binding the contract method 0xac1e2257.
//
// Solidity: function unstakeRequestInfo(uint256 destChainId, address l2strategy) view returns(bool, uint256)
func (_StakingManager *StakingManagerCallerSession) UnstakeRequestInfo(destChainId *big.Int, l2strategy common.Address) (bool, *big.Int, error) {
	return _StakingManager.Contract.UnstakeRequestInfo(&_StakingManager.CallOpts, destChainId, l2strategy)
}

// UsedValidators is a free data retrieval call binding the contract method 0x5915ded1.
//
// Solidity: function usedValidators(bytes pubkey) view returns(bool exists)
func (_StakingManager *StakingManagerCaller) UsedValidators(opts *bind.CallOpts, pubkey []byte) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "usedValidators", pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedValidators is a free data retrieval call binding the contract method 0x5915ded1.
//
// Solidity: function usedValidators(bytes pubkey) view returns(bool exists)
func (_StakingManager *StakingManagerSession) UsedValidators(pubkey []byte) (bool, error) {
	return _StakingManager.Contract.UsedValidators(&_StakingManager.CallOpts, pubkey)
}

// UsedValidators is a free data retrieval call binding the contract method 0x5915ded1.
//
// Solidity: function usedValidators(bytes pubkey) view returns(bool exists)
func (_StakingManager *StakingManagerCallerSession) UsedValidators(pubkey []byte) (bool, error) {
	return _StakingManager.Contract.UsedValidators(&_StakingManager.CallOpts, pubkey)
}

// WithdrawalWallet is a free data retrieval call binding the contract method 0x4a7d80b3.
//
// Solidity: function withdrawalWallet() view returns(address)
func (_StakingManager *StakingManagerCaller) WithdrawalWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "withdrawalWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawalWallet is a free data retrieval call binding the contract method 0x4a7d80b3.
//
// Solidity: function withdrawalWallet() view returns(address)
func (_StakingManager *StakingManagerSession) WithdrawalWallet() (common.Address, error) {
	return _StakingManager.Contract.WithdrawalWallet(&_StakingManager.CallOpts)
}

// WithdrawalWallet is a free data retrieval call binding the contract method 0x4a7d80b3.
//
// Solidity: function withdrawalWallet() view returns(address)
func (_StakingManager *StakingManagerCallerSession) WithdrawalWallet() (common.Address, error) {
	return _StakingManager.Contract.WithdrawalWallet(&_StakingManager.CallOpts)
}

// AllocateETH is a paid mutator transaction binding the contract method 0x6daa01a2.
//
// Solidity: function allocateETH(uint256 allocateToUnstakeRequestsManager, uint256 allocateToDeposits) returns()
func (_StakingManager *StakingManagerTransactor) AllocateETH(opts *bind.TransactOpts, allocateToUnstakeRequestsManager *big.Int, allocateToDeposits *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "allocateETH", allocateToUnstakeRequestsManager, allocateToDeposits)
}

// AllocateETH is a paid mutator transaction binding the contract method 0x6daa01a2.
//
// Solidity: function allocateETH(uint256 allocateToUnstakeRequestsManager, uint256 allocateToDeposits) returns()
func (_StakingManager *StakingManagerSession) AllocateETH(allocateToUnstakeRequestsManager *big.Int, allocateToDeposits *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.AllocateETH(&_StakingManager.TransactOpts, allocateToUnstakeRequestsManager, allocateToDeposits)
}

// AllocateETH is a paid mutator transaction binding the contract method 0x6daa01a2.
//
// Solidity: function allocateETH(uint256 allocateToUnstakeRequestsManager, uint256 allocateToDeposits) returns()
func (_StakingManager *StakingManagerTransactorSession) AllocateETH(allocateToUnstakeRequestsManager *big.Int, allocateToDeposits *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.AllocateETH(&_StakingManager.TransactOpts, allocateToUnstakeRequestsManager, allocateToDeposits)
}

// ClaimUnstakeRequest is a paid mutator transaction binding the contract method 0xc2c3c18c.
//
// Solidity: function claimUnstakeRequest((address,uint256)[] requests, uint256 sourceChainId, uint256 destChainId, uint256 gasLimit) returns()
func (_StakingManager *StakingManagerTransactor) ClaimUnstakeRequest(opts *bind.TransactOpts, requests []IUnstakeRequestsManagerWriterequestsInfo, sourceChainId *big.Int, destChainId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "claimUnstakeRequest", requests, sourceChainId, destChainId, gasLimit)
}

// ClaimUnstakeRequest is a paid mutator transaction binding the contract method 0xc2c3c18c.
//
// Solidity: function claimUnstakeRequest((address,uint256)[] requests, uint256 sourceChainId, uint256 destChainId, uint256 gasLimit) returns()
func (_StakingManager *StakingManagerSession) ClaimUnstakeRequest(requests []IUnstakeRequestsManagerWriterequestsInfo, sourceChainId *big.Int, destChainId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.ClaimUnstakeRequest(&_StakingManager.TransactOpts, requests, sourceChainId, destChainId, gasLimit)
}

// ClaimUnstakeRequest is a paid mutator transaction binding the contract method 0xc2c3c18c.
//
// Solidity: function claimUnstakeRequest((address,uint256)[] requests, uint256 sourceChainId, uint256 destChainId, uint256 gasLimit) returns()
func (_StakingManager *StakingManagerTransactorSession) ClaimUnstakeRequest(requests []IUnstakeRequestsManagerWriterequestsInfo, sourceChainId *big.Int, destChainId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.ClaimUnstakeRequest(&_StakingManager.TransactOpts, requests, sourceChainId, destChainId, gasLimit)
}

// DETHToETH is a paid mutator transaction binding the contract method 0xed9daafb.
//
// Solidity: function dETHToETH(uint256 dETHAmount) returns(uint256)
func (_StakingManager *StakingManagerTransactor) DETHToETH(opts *bind.TransactOpts, dETHAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "dETHToETH", dETHAmount)
}

// DETHToETH is a paid mutator transaction binding the contract method 0xed9daafb.
//
// Solidity: function dETHToETH(uint256 dETHAmount) returns(uint256)
func (_StakingManager *StakingManagerSession) DETHToETH(dETHAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.DETHToETH(&_StakingManager.TransactOpts, dETHAmount)
}

// DETHToETH is a paid mutator transaction binding the contract method 0xed9daafb.
//
// Solidity: function dETHToETH(uint256 dETHAmount) returns(uint256)
func (_StakingManager *StakingManagerTransactorSession) DETHToETH(dETHAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.DETHToETH(&_StakingManager.TransactOpts, dETHAmount)
}

// EthToDETH is a paid mutator transaction binding the contract method 0x147d36d5.
//
// Solidity: function ethToDETH(uint256 ethAmount) returns(uint256)
func (_StakingManager *StakingManagerTransactor) EthToDETH(opts *bind.TransactOpts, ethAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "ethToDETH", ethAmount)
}

// EthToDETH is a paid mutator transaction binding the contract method 0x147d36d5.
//
// Solidity: function ethToDETH(uint256 ethAmount) returns(uint256)
func (_StakingManager *StakingManagerSession) EthToDETH(ethAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.EthToDETH(&_StakingManager.TransactOpts, ethAmount)
}

// EthToDETH is a paid mutator transaction binding the contract method 0x147d36d5.
//
// Solidity: function ethToDETH(uint256 ethAmount) returns(uint256)
func (_StakingManager *StakingManagerTransactorSession) EthToDETH(ethAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.EthToDETH(&_StakingManager.TransactOpts, ethAmount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.GrantRole(&_StakingManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.GrantRole(&_StakingManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f656d22.
//
// Solidity: function initialize((address,address,address,address,address) init) returns()
func (_StakingManager *StakingManagerTransactor) Initialize(opts *bind.TransactOpts, init StakingManagerInit) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "initialize", init)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f656d22.
//
// Solidity: function initialize((address,address,address,address,address) init) returns()
func (_StakingManager *StakingManagerSession) Initialize(init StakingManagerInit) (*types.Transaction, error) {
	return _StakingManager.Contract.Initialize(&_StakingManager.TransactOpts, init)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f656d22.
//
// Solidity: function initialize((address,address,address,address,address) init) returns()
func (_StakingManager *StakingManagerTransactorSession) Initialize(init StakingManagerInit) (*types.Transaction, error) {
	return _StakingManager.Contract.Initialize(&_StakingManager.TransactOpts, init)
}

// InitiateValidatorsWithDeposits is a paid mutator transaction binding the contract method 0x0208e4b5.
//
// Solidity: function initiateValidatorsWithDeposits((uint256,uint256,bytes,bytes,bytes,bytes32)[] validators, bytes32 expectedDepositRoot) returns()
func (_StakingManager *StakingManagerTransactor) InitiateValidatorsWithDeposits(opts *bind.TransactOpts, validators []StakingManagerStorageValidatorParams, expectedDepositRoot [32]byte) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "initiateValidatorsWithDeposits", validators, expectedDepositRoot)
}

// InitiateValidatorsWithDeposits is a paid mutator transaction binding the contract method 0x0208e4b5.
//
// Solidity: function initiateValidatorsWithDeposits((uint256,uint256,bytes,bytes,bytes,bytes32)[] validators, bytes32 expectedDepositRoot) returns()
func (_StakingManager *StakingManagerSession) InitiateValidatorsWithDeposits(validators []StakingManagerStorageValidatorParams, expectedDepositRoot [32]byte) (*types.Transaction, error) {
	return _StakingManager.Contract.InitiateValidatorsWithDeposits(&_StakingManager.TransactOpts, validators, expectedDepositRoot)
}

// InitiateValidatorsWithDeposits is a paid mutator transaction binding the contract method 0x0208e4b5.
//
// Solidity: function initiateValidatorsWithDeposits((uint256,uint256,bytes,bytes,bytes,bytes32)[] validators, bytes32 expectedDepositRoot) returns()
func (_StakingManager *StakingManagerTransactorSession) InitiateValidatorsWithDeposits(validators []StakingManagerStorageValidatorParams, expectedDepositRoot [32]byte) (*types.Transaction, error) {
	return _StakingManager.Contract.InitiateValidatorsWithDeposits(&_StakingManager.TransactOpts, validators, expectedDepositRoot)
}

// ReceiveFromUnstakeRequestsManager is a paid mutator transaction binding the contract method 0xc151aa72.
//
// Solidity: function receiveFromUnstakeRequestsManager() payable returns()
func (_StakingManager *StakingManagerTransactor) ReceiveFromUnstakeRequestsManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "receiveFromUnstakeRequestsManager")
}

// ReceiveFromUnstakeRequestsManager is a paid mutator transaction binding the contract method 0xc151aa72.
//
// Solidity: function receiveFromUnstakeRequestsManager() payable returns()
func (_StakingManager *StakingManagerSession) ReceiveFromUnstakeRequestsManager() (*types.Transaction, error) {
	return _StakingManager.Contract.ReceiveFromUnstakeRequestsManager(&_StakingManager.TransactOpts)
}

// ReceiveFromUnstakeRequestsManager is a paid mutator transaction binding the contract method 0xc151aa72.
//
// Solidity: function receiveFromUnstakeRequestsManager() payable returns()
func (_StakingManager *StakingManagerTransactorSession) ReceiveFromUnstakeRequestsManager() (*types.Transaction, error) {
	return _StakingManager.Contract.ReceiveFromUnstakeRequestsManager(&_StakingManager.TransactOpts)
}

// ReceiveReturns is a paid mutator transaction binding the contract method 0x808d663f.
//
// Solidity: function receiveReturns() payable returns()
func (_StakingManager *StakingManagerTransactor) ReceiveReturns(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "receiveReturns")
}

// ReceiveReturns is a paid mutator transaction binding the contract method 0x808d663f.
//
// Solidity: function receiveReturns() payable returns()
func (_StakingManager *StakingManagerSession) ReceiveReturns() (*types.Transaction, error) {
	return _StakingManager.Contract.ReceiveReturns(&_StakingManager.TransactOpts)
}

// ReceiveReturns is a paid mutator transaction binding the contract method 0x808d663f.
//
// Solidity: function receiveReturns() payable returns()
func (_StakingManager *StakingManagerTransactorSession) ReceiveReturns() (*types.Transaction, error) {
	return _StakingManager.Contract.ReceiveReturns(&_StakingManager.TransactOpts)
}

// ReclaimAllocatedETHSurplus is a paid mutator transaction binding the contract method 0x1943190d.
//
// Solidity: function reclaimAllocatedETHSurplus() returns()
func (_StakingManager *StakingManagerTransactor) ReclaimAllocatedETHSurplus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "reclaimAllocatedETHSurplus")
}

// ReclaimAllocatedETHSurplus is a paid mutator transaction binding the contract method 0x1943190d.
//
// Solidity: function reclaimAllocatedETHSurplus() returns()
func (_StakingManager *StakingManagerSession) ReclaimAllocatedETHSurplus() (*types.Transaction, error) {
	return _StakingManager.Contract.ReclaimAllocatedETHSurplus(&_StakingManager.TransactOpts)
}

// ReclaimAllocatedETHSurplus is a paid mutator transaction binding the contract method 0x1943190d.
//
// Solidity: function reclaimAllocatedETHSurplus() returns()
func (_StakingManager *StakingManagerTransactorSession) ReclaimAllocatedETHSurplus() (*types.Transaction, error) {
	return _StakingManager.Contract.ReclaimAllocatedETHSurplus(&_StakingManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingManager *StakingManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingManager *StakingManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceRole(&_StakingManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingManager *StakingManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceRole(&_StakingManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RevokeRole(&_StakingManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RevokeRole(&_StakingManager.TransactOpts, role, account)
}

// SetExchangeAdjustmentRate is a paid mutator transaction binding the contract method 0x29d48704.
//
// Solidity: function setExchangeAdjustmentRate(uint16 exchangeAdjustmentRate_) returns()
func (_StakingManager *StakingManagerTransactor) SetExchangeAdjustmentRate(opts *bind.TransactOpts, exchangeAdjustmentRate_ uint16) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setExchangeAdjustmentRate", exchangeAdjustmentRate_)
}

// SetExchangeAdjustmentRate is a paid mutator transaction binding the contract method 0x29d48704.
//
// Solidity: function setExchangeAdjustmentRate(uint16 exchangeAdjustmentRate_) returns()
func (_StakingManager *StakingManagerSession) SetExchangeAdjustmentRate(exchangeAdjustmentRate_ uint16) (*types.Transaction, error) {
	return _StakingManager.Contract.SetExchangeAdjustmentRate(&_StakingManager.TransactOpts, exchangeAdjustmentRate_)
}

// SetExchangeAdjustmentRate is a paid mutator transaction binding the contract method 0x29d48704.
//
// Solidity: function setExchangeAdjustmentRate(uint16 exchangeAdjustmentRate_) returns()
func (_StakingManager *StakingManagerTransactorSession) SetExchangeAdjustmentRate(exchangeAdjustmentRate_ uint16) (*types.Transaction, error) {
	return _StakingManager.Contract.SetExchangeAdjustmentRate(&_StakingManager.TransactOpts, exchangeAdjustmentRate_)
}

// SetLocator is a paid mutator transaction binding the contract method 0xa5e84562.
//
// Solidity: function setLocator(address _locator) returns()
func (_StakingManager *StakingManagerTransactor) SetLocator(opts *bind.TransactOpts, _locator common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setLocator", _locator)
}

// SetLocator is a paid mutator transaction binding the contract method 0xa5e84562.
//
// Solidity: function setLocator(address _locator) returns()
func (_StakingManager *StakingManagerSession) SetLocator(_locator common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetLocator(&_StakingManager.TransactOpts, _locator)
}

// SetLocator is a paid mutator transaction binding the contract method 0xa5e84562.
//
// Solidity: function setLocator(address _locator) returns()
func (_StakingManager *StakingManagerTransactorSession) SetLocator(_locator common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetLocator(&_StakingManager.TransactOpts, _locator)
}

// SetMaximumDETHSupply is a paid mutator transaction binding the contract method 0x1d2d35ce.
//
// Solidity: function setMaximumDETHSupply(uint256 maximumDETHSupply_) returns()
func (_StakingManager *StakingManagerTransactor) SetMaximumDETHSupply(opts *bind.TransactOpts, maximumDETHSupply_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setMaximumDETHSupply", maximumDETHSupply_)
}

// SetMaximumDETHSupply is a paid mutator transaction binding the contract method 0x1d2d35ce.
//
// Solidity: function setMaximumDETHSupply(uint256 maximumDETHSupply_) returns()
func (_StakingManager *StakingManagerSession) SetMaximumDETHSupply(maximumDETHSupply_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMaximumDETHSupply(&_StakingManager.TransactOpts, maximumDETHSupply_)
}

// SetMaximumDETHSupply is a paid mutator transaction binding the contract method 0x1d2d35ce.
//
// Solidity: function setMaximumDETHSupply(uint256 maximumDETHSupply_) returns()
func (_StakingManager *StakingManagerTransactorSession) SetMaximumDETHSupply(maximumDETHSupply_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMaximumDETHSupply(&_StakingManager.TransactOpts, maximumDETHSupply_)
}

// SetMinimumDepositAmount is a paid mutator transaction binding the contract method 0xaab483d6.
//
// Solidity: function setMinimumDepositAmount(uint256 minimumDepositAmount_) returns()
func (_StakingManager *StakingManagerTransactor) SetMinimumDepositAmount(opts *bind.TransactOpts, minimumDepositAmount_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setMinimumDepositAmount", minimumDepositAmount_)
}

// SetMinimumDepositAmount is a paid mutator transaction binding the contract method 0xaab483d6.
//
// Solidity: function setMinimumDepositAmount(uint256 minimumDepositAmount_) returns()
func (_StakingManager *StakingManagerSession) SetMinimumDepositAmount(minimumDepositAmount_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMinimumDepositAmount(&_StakingManager.TransactOpts, minimumDepositAmount_)
}

// SetMinimumDepositAmount is a paid mutator transaction binding the contract method 0xaab483d6.
//
// Solidity: function setMinimumDepositAmount(uint256 minimumDepositAmount_) returns()
func (_StakingManager *StakingManagerTransactorSession) SetMinimumDepositAmount(minimumDepositAmount_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMinimumDepositAmount(&_StakingManager.TransactOpts, minimumDepositAmount_)
}

// SetMinimumUnstakeBound is a paid mutator transaction binding the contract method 0x99dd1deb.
//
// Solidity: function setMinimumUnstakeBound(uint256 minimumUnstakeBound_) returns()
func (_StakingManager *StakingManagerTransactor) SetMinimumUnstakeBound(opts *bind.TransactOpts, minimumUnstakeBound_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setMinimumUnstakeBound", minimumUnstakeBound_)
}

// SetMinimumUnstakeBound is a paid mutator transaction binding the contract method 0x99dd1deb.
//
// Solidity: function setMinimumUnstakeBound(uint256 minimumUnstakeBound_) returns()
func (_StakingManager *StakingManagerSession) SetMinimumUnstakeBound(minimumUnstakeBound_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMinimumUnstakeBound(&_StakingManager.TransactOpts, minimumUnstakeBound_)
}

// SetMinimumUnstakeBound is a paid mutator transaction binding the contract method 0x99dd1deb.
//
// Solidity: function setMinimumUnstakeBound(uint256 minimumUnstakeBound_) returns()
func (_StakingManager *StakingManagerTransactorSession) SetMinimumUnstakeBound(minimumUnstakeBound_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMinimumUnstakeBound(&_StakingManager.TransactOpts, minimumUnstakeBound_)
}

// SetStakingAllowlist is a paid mutator transaction binding the contract method 0x04f36cc2.
//
// Solidity: function setStakingAllowlist(bool isStakingAllowlist_) returns()
func (_StakingManager *StakingManagerTransactor) SetStakingAllowlist(opts *bind.TransactOpts, isStakingAllowlist_ bool) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setStakingAllowlist", isStakingAllowlist_)
}

// SetStakingAllowlist is a paid mutator transaction binding the contract method 0x04f36cc2.
//
// Solidity: function setStakingAllowlist(bool isStakingAllowlist_) returns()
func (_StakingManager *StakingManagerSession) SetStakingAllowlist(isStakingAllowlist_ bool) (*types.Transaction, error) {
	return _StakingManager.Contract.SetStakingAllowlist(&_StakingManager.TransactOpts, isStakingAllowlist_)
}

// SetStakingAllowlist is a paid mutator transaction binding the contract method 0x04f36cc2.
//
// Solidity: function setStakingAllowlist(bool isStakingAllowlist_) returns()
func (_StakingManager *StakingManagerTransactorSession) SetStakingAllowlist(isStakingAllowlist_ bool) (*types.Transaction, error) {
	return _StakingManager.Contract.SetStakingAllowlist(&_StakingManager.TransactOpts, isStakingAllowlist_)
}

// SetWithdrawalWallet is a paid mutator transaction binding the contract method 0x75796f76.
//
// Solidity: function setWithdrawalWallet(address withdrawalWallet_) returns()
func (_StakingManager *StakingManagerTransactor) SetWithdrawalWallet(opts *bind.TransactOpts, withdrawalWallet_ common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setWithdrawalWallet", withdrawalWallet_)
}

// SetWithdrawalWallet is a paid mutator transaction binding the contract method 0x75796f76.
//
// Solidity: function setWithdrawalWallet(address withdrawalWallet_) returns()
func (_StakingManager *StakingManagerSession) SetWithdrawalWallet(withdrawalWallet_ common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetWithdrawalWallet(&_StakingManager.TransactOpts, withdrawalWallet_)
}

// SetWithdrawalWallet is a paid mutator transaction binding the contract method 0x75796f76.
//
// Solidity: function setWithdrawalWallet(address withdrawalWallet_) returns()
func (_StakingManager *StakingManagerTransactorSession) SetWithdrawalWallet(withdrawalWallet_ common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetWithdrawalWallet(&_StakingManager.TransactOpts, withdrawalWallet_)
}

// Stake is a paid mutator transaction binding the contract method 0x37a6c881.
//
// Solidity: function stake(uint256 stakeAmount, (address,uint256)[] batchMints) payable returns()
func (_StakingManager *StakingManagerTransactor) Stake(opts *bind.TransactOpts, stakeAmount *big.Int, batchMints []IDETHBatchMint) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "stake", stakeAmount, batchMints)
}

// Stake is a paid mutator transaction binding the contract method 0x37a6c881.
//
// Solidity: function stake(uint256 stakeAmount, (address,uint256)[] batchMints) payable returns()
func (_StakingManager *StakingManagerSession) Stake(stakeAmount *big.Int, batchMints []IDETHBatchMint) (*types.Transaction, error) {
	return _StakingManager.Contract.Stake(&_StakingManager.TransactOpts, stakeAmount, batchMints)
}

// Stake is a paid mutator transaction binding the contract method 0x37a6c881.
//
// Solidity: function stake(uint256 stakeAmount, (address,uint256)[] batchMints) payable returns()
func (_StakingManager *StakingManagerTransactorSession) Stake(stakeAmount *big.Int, batchMints []IDETHBatchMint) (*types.Transaction, error) {
	return _StakingManager.Contract.Stake(&_StakingManager.TransactOpts, stakeAmount, batchMints)
}

// TopUp is a paid mutator transaction binding the contract method 0xdc29f1de.
//
// Solidity: function topUp() payable returns()
func (_StakingManager *StakingManagerTransactor) TopUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "topUp")
}

// TopUp is a paid mutator transaction binding the contract method 0xdc29f1de.
//
// Solidity: function topUp() payable returns()
func (_StakingManager *StakingManagerSession) TopUp() (*types.Transaction, error) {
	return _StakingManager.Contract.TopUp(&_StakingManager.TransactOpts)
}

// TopUp is a paid mutator transaction binding the contract method 0xdc29f1de.
//
// Solidity: function topUp() payable returns()
func (_StakingManager *StakingManagerTransactorSession) TopUp() (*types.Transaction, error) {
	return _StakingManager.Contract.TopUp(&_StakingManager.TransactOpts)
}

// TotalControlled is a paid mutator transaction binding the contract method 0x5940d90b.
//
// Solidity: function totalControlled() returns(uint256)
func (_StakingManager *StakingManagerTransactor) TotalControlled(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "totalControlled")
}

// TotalControlled is a paid mutator transaction binding the contract method 0x5940d90b.
//
// Solidity: function totalControlled() returns(uint256)
func (_StakingManager *StakingManagerSession) TotalControlled() (*types.Transaction, error) {
	return _StakingManager.Contract.TotalControlled(&_StakingManager.TransactOpts)
}

// TotalControlled is a paid mutator transaction binding the contract method 0x5940d90b.
//
// Solidity: function totalControlled() returns(uint256)
func (_StakingManager *StakingManagerTransactorSession) TotalControlled() (*types.Transaction, error) {
	return _StakingManager.Contract.TotalControlled(&_StakingManager.TransactOpts)
}

// UnstakeRequest is a paid mutator transaction binding the contract method 0x12e9ead6.
//
// Solidity: function unstakeRequest(uint128 dethAmount, uint128 minETHAmount, address l2Strategy, uint256 destChainId) returns()
func (_StakingManager *StakingManagerTransactor) UnstakeRequest(opts *bind.TransactOpts, dethAmount *big.Int, minETHAmount *big.Int, l2Strategy common.Address, destChainId *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "unstakeRequest", dethAmount, minETHAmount, l2Strategy, destChainId)
}

// UnstakeRequest is a paid mutator transaction binding the contract method 0x12e9ead6.
//
// Solidity: function unstakeRequest(uint128 dethAmount, uint128 minETHAmount, address l2Strategy, uint256 destChainId) returns()
func (_StakingManager *StakingManagerSession) UnstakeRequest(dethAmount *big.Int, minETHAmount *big.Int, l2Strategy common.Address, destChainId *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.UnstakeRequest(&_StakingManager.TransactOpts, dethAmount, minETHAmount, l2Strategy, destChainId)
}

// UnstakeRequest is a paid mutator transaction binding the contract method 0x12e9ead6.
//
// Solidity: function unstakeRequest(uint128 dethAmount, uint128 minETHAmount, address l2Strategy, uint256 destChainId) returns()
func (_StakingManager *StakingManagerTransactorSession) UnstakeRequest(dethAmount *big.Int, minETHAmount *big.Int, l2Strategy common.Address, destChainId *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.UnstakeRequest(&_StakingManager.TransactOpts, dethAmount, minETHAmount, l2Strategy, destChainId)
}

// StakingManagerAllocatedETHToDepositsIterator is returned from FilterAllocatedETHToDeposits and is used to iterate over the raw logs and unpacked data for AllocatedETHToDeposits events raised by the StakingManager contract.
type StakingManagerAllocatedETHToDepositsIterator struct {
	Event *StakingManagerAllocatedETHToDeposits // Event containing the contract specifics and raw log

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
func (it *StakingManagerAllocatedETHToDepositsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerAllocatedETHToDeposits)
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
		it.Event = new(StakingManagerAllocatedETHToDeposits)
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
func (it *StakingManagerAllocatedETHToDepositsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerAllocatedETHToDepositsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerAllocatedETHToDeposits represents a AllocatedETHToDeposits event raised by the StakingManager contract.
type StakingManagerAllocatedETHToDeposits struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAllocatedETHToDeposits is a free log retrieval operation binding the contract event 0x9d04ecb465d2c8754acb798a22293dd26215a1c2f7a2a56607afa215c1aadc77.
//
// Solidity: event AllocatedETHToDeposits(uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterAllocatedETHToDeposits(opts *bind.FilterOpts) (*StakingManagerAllocatedETHToDepositsIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "AllocatedETHToDeposits")
	if err != nil {
		return nil, err
	}
	return &StakingManagerAllocatedETHToDepositsIterator{contract: _StakingManager.contract, event: "AllocatedETHToDeposits", logs: logs, sub: sub}, nil
}

// WatchAllocatedETHToDeposits is a free log subscription operation binding the contract event 0x9d04ecb465d2c8754acb798a22293dd26215a1c2f7a2a56607afa215c1aadc77.
//
// Solidity: event AllocatedETHToDeposits(uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchAllocatedETHToDeposits(opts *bind.WatchOpts, sink chan<- *StakingManagerAllocatedETHToDeposits) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "AllocatedETHToDeposits")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerAllocatedETHToDeposits)
				if err := _StakingManager.contract.UnpackLog(event, "AllocatedETHToDeposits", log); err != nil {
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

// ParseAllocatedETHToDeposits is a log parse operation binding the contract event 0x9d04ecb465d2c8754acb798a22293dd26215a1c2f7a2a56607afa215c1aadc77.
//
// Solidity: event AllocatedETHToDeposits(uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseAllocatedETHToDeposits(log types.Log) (*StakingManagerAllocatedETHToDeposits, error) {
	event := new(StakingManagerAllocatedETHToDeposits)
	if err := _StakingManager.contract.UnpackLog(event, "AllocatedETHToDeposits", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerAllocatedETHToUnstakeRequestsManagerIterator is returned from FilterAllocatedETHToUnstakeRequestsManager and is used to iterate over the raw logs and unpacked data for AllocatedETHToUnstakeRequestsManager events raised by the StakingManager contract.
type StakingManagerAllocatedETHToUnstakeRequestsManagerIterator struct {
	Event *StakingManagerAllocatedETHToUnstakeRequestsManager // Event containing the contract specifics and raw log

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
func (it *StakingManagerAllocatedETHToUnstakeRequestsManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerAllocatedETHToUnstakeRequestsManager)
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
		it.Event = new(StakingManagerAllocatedETHToUnstakeRequestsManager)
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
func (it *StakingManagerAllocatedETHToUnstakeRequestsManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerAllocatedETHToUnstakeRequestsManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerAllocatedETHToUnstakeRequestsManager represents a AllocatedETHToUnstakeRequestsManager event raised by the StakingManager contract.
type StakingManagerAllocatedETHToUnstakeRequestsManager struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAllocatedETHToUnstakeRequestsManager is a free log retrieval operation binding the contract event 0xfe89805cf5299ef9fbd1d1ddefb8dcc3fa9408064d2ea31e3fca6565768f5217.
//
// Solidity: event AllocatedETHToUnstakeRequestsManager(uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterAllocatedETHToUnstakeRequestsManager(opts *bind.FilterOpts) (*StakingManagerAllocatedETHToUnstakeRequestsManagerIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "AllocatedETHToUnstakeRequestsManager")
	if err != nil {
		return nil, err
	}
	return &StakingManagerAllocatedETHToUnstakeRequestsManagerIterator{contract: _StakingManager.contract, event: "AllocatedETHToUnstakeRequestsManager", logs: logs, sub: sub}, nil
}

// WatchAllocatedETHToUnstakeRequestsManager is a free log subscription operation binding the contract event 0xfe89805cf5299ef9fbd1d1ddefb8dcc3fa9408064d2ea31e3fca6565768f5217.
//
// Solidity: event AllocatedETHToUnstakeRequestsManager(uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchAllocatedETHToUnstakeRequestsManager(opts *bind.WatchOpts, sink chan<- *StakingManagerAllocatedETHToUnstakeRequestsManager) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "AllocatedETHToUnstakeRequestsManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerAllocatedETHToUnstakeRequestsManager)
				if err := _StakingManager.contract.UnpackLog(event, "AllocatedETHToUnstakeRequestsManager", log); err != nil {
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

// ParseAllocatedETHToUnstakeRequestsManager is a log parse operation binding the contract event 0xfe89805cf5299ef9fbd1d1ddefb8dcc3fa9408064d2ea31e3fca6565768f5217.
//
// Solidity: event AllocatedETHToUnstakeRequestsManager(uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseAllocatedETHToUnstakeRequestsManager(log types.Log) (*StakingManagerAllocatedETHToUnstakeRequestsManager, error) {
	event := new(StakingManagerAllocatedETHToUnstakeRequestsManager)
	if err := _StakingManager.contract.UnpackLog(event, "AllocatedETHToUnstakeRequestsManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakingManager contract.
type StakingManagerInitializedIterator struct {
	Event *StakingManagerInitialized // Event containing the contract specifics and raw log

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
func (it *StakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerInitialized)
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
		it.Event = new(StakingManagerInitialized)
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
func (it *StakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerInitialized represents a Initialized event raised by the StakingManager contract.
type StakingManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakingManager *StakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingManagerInitializedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingManagerInitializedIterator{contract: _StakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakingManager *StakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerInitialized)
				if err := _StakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseInitialized(log types.Log) (*StakingManagerInitialized, error) {
	event := new(StakingManagerInitialized)
	if err := _StakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerProtocolConfigChangedIterator is returned from FilterProtocolConfigChanged and is used to iterate over the raw logs and unpacked data for ProtocolConfigChanged events raised by the StakingManager contract.
type StakingManagerProtocolConfigChangedIterator struct {
	Event *StakingManagerProtocolConfigChanged // Event containing the contract specifics and raw log

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
func (it *StakingManagerProtocolConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerProtocolConfigChanged)
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
		it.Event = new(StakingManagerProtocolConfigChanged)
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
func (it *StakingManagerProtocolConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerProtocolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerProtocolConfigChanged represents a ProtocolConfigChanged event raised by the StakingManager contract.
type StakingManagerProtocolConfigChanged struct {
	SetterSelector  [4]byte
	SetterSignature string
	Value           []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProtocolConfigChanged is a free log retrieval operation binding the contract event 0x01d854e8dde9402801a4c6f2840193465752abfad61e0bb7c4258d526ae42e74.
//
// Solidity: event ProtocolConfigChanged(bytes4 indexed setterSelector, string setterSignature, bytes value)
func (_StakingManager *StakingManagerFilterer) FilterProtocolConfigChanged(opts *bind.FilterOpts, setterSelector [][4]byte) (*StakingManagerProtocolConfigChangedIterator, error) {

	var setterSelectorRule []interface{}
	for _, setterSelectorItem := range setterSelector {
		setterSelectorRule = append(setterSelectorRule, setterSelectorItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "ProtocolConfigChanged", setterSelectorRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerProtocolConfigChangedIterator{contract: _StakingManager.contract, event: "ProtocolConfigChanged", logs: logs, sub: sub}, nil
}

// WatchProtocolConfigChanged is a free log subscription operation binding the contract event 0x01d854e8dde9402801a4c6f2840193465752abfad61e0bb7c4258d526ae42e74.
//
// Solidity: event ProtocolConfigChanged(bytes4 indexed setterSelector, string setterSignature, bytes value)
func (_StakingManager *StakingManagerFilterer) WatchProtocolConfigChanged(opts *bind.WatchOpts, sink chan<- *StakingManagerProtocolConfigChanged, setterSelector [][4]byte) (event.Subscription, error) {

	var setterSelectorRule []interface{}
	for _, setterSelectorItem := range setterSelector {
		setterSelectorRule = append(setterSelectorRule, setterSelectorItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "ProtocolConfigChanged", setterSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerProtocolConfigChanged)
				if err := _StakingManager.contract.UnpackLog(event, "ProtocolConfigChanged", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseProtocolConfigChanged(log types.Log) (*StakingManagerProtocolConfigChanged, error) {
	event := new(StakingManagerProtocolConfigChanged)
	if err := _StakingManager.contract.UnpackLog(event, "ProtocolConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerReturnsReceivedIterator is returned from FilterReturnsReceived and is used to iterate over the raw logs and unpacked data for ReturnsReceived events raised by the StakingManager contract.
type StakingManagerReturnsReceivedIterator struct {
	Event *StakingManagerReturnsReceived // Event containing the contract specifics and raw log

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
func (it *StakingManagerReturnsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerReturnsReceived)
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
		it.Event = new(StakingManagerReturnsReceived)
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
func (it *StakingManagerReturnsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerReturnsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerReturnsReceived represents a ReturnsReceived event raised by the StakingManager contract.
type StakingManagerReturnsReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReturnsReceived is a free log retrieval operation binding the contract event 0x4cbb9d73b003a252cee3f2ee51d8d65a562af35eebb23730dd4a76d68127b370.
//
// Solidity: event ReturnsReceived(uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterReturnsReceived(opts *bind.FilterOpts) (*StakingManagerReturnsReceivedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "ReturnsReceived")
	if err != nil {
		return nil, err
	}
	return &StakingManagerReturnsReceivedIterator{contract: _StakingManager.contract, event: "ReturnsReceived", logs: logs, sub: sub}, nil
}

// WatchReturnsReceived is a free log subscription operation binding the contract event 0x4cbb9d73b003a252cee3f2ee51d8d65a562af35eebb23730dd4a76d68127b370.
//
// Solidity: event ReturnsReceived(uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchReturnsReceived(opts *bind.WatchOpts, sink chan<- *StakingManagerReturnsReceived) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "ReturnsReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerReturnsReceived)
				if err := _StakingManager.contract.UnpackLog(event, "ReturnsReceived", log); err != nil {
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

// ParseReturnsReceived is a log parse operation binding the contract event 0x4cbb9d73b003a252cee3f2ee51d8d65a562af35eebb23730dd4a76d68127b370.
//
// Solidity: event ReturnsReceived(uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseReturnsReceived(log types.Log) (*StakingManagerReturnsReceived, error) {
	event := new(StakingManagerReturnsReceived)
	if err := _StakingManager.contract.UnpackLog(event, "ReturnsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StakingManager contract.
type StakingManagerRoleAdminChangedIterator struct {
	Event *StakingManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerRoleAdminChanged)
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
		it.Event = new(StakingManagerRoleAdminChanged)
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
func (it *StakingManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerRoleAdminChanged represents a RoleAdminChanged event raised by the StakingManager contract.
type StakingManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingManager *StakingManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakingManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerRoleAdminChangedIterator{contract: _StakingManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingManager *StakingManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerRoleAdminChanged)
				if err := _StakingManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseRoleAdminChanged(log types.Log) (*StakingManagerRoleAdminChanged, error) {
	event := new(StakingManagerRoleAdminChanged)
	if err := _StakingManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StakingManager contract.
type StakingManagerRoleGrantedIterator struct {
	Event *StakingManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakingManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerRoleGranted)
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
		it.Event = new(StakingManagerRoleGranted)
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
func (it *StakingManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerRoleGranted represents a RoleGranted event raised by the StakingManager contract.
type StakingManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerRoleGrantedIterator{contract: _StakingManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakingManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerRoleGranted)
				if err := _StakingManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseRoleGranted(log types.Log) (*StakingManagerRoleGranted, error) {
	event := new(StakingManagerRoleGranted)
	if err := _StakingManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StakingManager contract.
type StakingManagerRoleRevokedIterator struct {
	Event *StakingManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakingManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerRoleRevoked)
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
		it.Event = new(StakingManagerRoleRevoked)
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
func (it *StakingManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerRoleRevoked represents a RoleRevoked event raised by the StakingManager contract.
type StakingManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerRoleRevokedIterator{contract: _StakingManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakingManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerRoleRevoked)
				if err := _StakingManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseRoleRevoked(log types.Log) (*StakingManagerRoleRevoked, error) {
	event := new(StakingManagerRoleRevoked)
	if err := _StakingManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the StakingManager contract.
type StakingManagerStakedIterator struct {
	Event *StakingManagerStaked // Event containing the contract specifics and raw log

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
func (it *StakingManagerStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerStaked)
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
		it.Event = new(StakingManagerStaked)
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
func (it *StakingManagerStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerStaked represents a Staked event raised by the StakingManager contract.
type StakingManagerStaked struct {
	Staker     common.Address
	EthAmount  *big.Int
	DETHAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed staker, uint256 ethAmount, uint256 dETHAmount)
func (_StakingManager *StakingManagerFilterer) FilterStaked(opts *bind.FilterOpts, staker []common.Address) (*StakingManagerStakedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerStakedIterator{contract: _StakingManager.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed staker, uint256 ethAmount, uint256 dETHAmount)
func (_StakingManager *StakingManagerFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingManagerStaked, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerStaked)
				if err := _StakingManager.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed staker, uint256 ethAmount, uint256 dETHAmount)
func (_StakingManager *StakingManagerFilterer) ParseStaked(log types.Log) (*StakingManagerStaked, error) {
	event := new(StakingManagerStaked)
	if err := _StakingManager.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUnstakeLaveAmountIterator is returned from FilterUnstakeLaveAmount and is used to iterate over the raw logs and unpacked data for UnstakeLaveAmount events raised by the StakingManager contract.
type StakingManagerUnstakeLaveAmountIterator struct {
	Event *StakingManagerUnstakeLaveAmount // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnstakeLaveAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnstakeLaveAmount)
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
		it.Event = new(StakingManagerUnstakeLaveAmount)
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
func (it *StakingManagerUnstakeLaveAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnstakeLaveAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnstakeLaveAmount represents a UnstakeLaveAmount event raised by the StakingManager contract.
type StakingManagerUnstakeLaveAmount struct {
	Staker     common.Address
	DETHLocked *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnstakeLaveAmount is a free log retrieval operation binding the contract event 0x1f4bdd7f902a9da109e5ee8424af475df1f64b7b7a8d36e07d5f97caaa5f19ec.
//
// Solidity: event UnstakeLaveAmount(address indexed staker, uint256 dETHLocked)
func (_StakingManager *StakingManagerFilterer) FilterUnstakeLaveAmount(opts *bind.FilterOpts, staker []common.Address) (*StakingManagerUnstakeLaveAmountIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnstakeLaveAmount", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnstakeLaveAmountIterator{contract: _StakingManager.contract, event: "UnstakeLaveAmount", logs: logs, sub: sub}, nil
}

// WatchUnstakeLaveAmount is a free log subscription operation binding the contract event 0x1f4bdd7f902a9da109e5ee8424af475df1f64b7b7a8d36e07d5f97caaa5f19ec.
//
// Solidity: event UnstakeLaveAmount(address indexed staker, uint256 dETHLocked)
func (_StakingManager *StakingManagerFilterer) WatchUnstakeLaveAmount(opts *bind.WatchOpts, sink chan<- *StakingManagerUnstakeLaveAmount, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnstakeLaveAmount", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnstakeLaveAmount)
				if err := _StakingManager.contract.UnpackLog(event, "UnstakeLaveAmount", log); err != nil {
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

// ParseUnstakeLaveAmount is a log parse operation binding the contract event 0x1f4bdd7f902a9da109e5ee8424af475df1f64b7b7a8d36e07d5f97caaa5f19ec.
//
// Solidity: event UnstakeLaveAmount(address indexed staker, uint256 dETHLocked)
func (_StakingManager *StakingManagerFilterer) ParseUnstakeLaveAmount(log types.Log) (*StakingManagerUnstakeLaveAmount, error) {
	event := new(StakingManagerUnstakeLaveAmount)
	if err := _StakingManager.contract.UnpackLog(event, "UnstakeLaveAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUnstakeRequestClaimedIterator is returned from FilterUnstakeRequestClaimed and is used to iterate over the raw logs and unpacked data for UnstakeRequestClaimed events raised by the StakingManager contract.
type StakingManagerUnstakeRequestClaimedIterator struct {
	Event *StakingManagerUnstakeRequestClaimed // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnstakeRequestClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnstakeRequestClaimed)
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
		it.Event = new(StakingManagerUnstakeRequestClaimed)
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
func (it *StakingManagerUnstakeRequestClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnstakeRequestClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnstakeRequestClaimed represents a UnstakeRequestClaimed event raised by the StakingManager contract.
type StakingManagerUnstakeRequestClaimed struct {
	Staker        common.Address
	L2Strategys   []common.Address
	Bridge        common.Address
	SourceChainId *big.Int
	DestChainId   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnstakeRequestClaimed is a free log retrieval operation binding the contract event 0xad6735aba298e7700470a993c5cbfd34a66ca15b17ac9b2b0984435d6f11f48d.
//
// Solidity: event UnstakeRequestClaimed(address indexed staker, address[] indexed l2Strategys, address indexed bridge, uint256 sourceChainId, uint256 destChainId)
func (_StakingManager *StakingManagerFilterer) FilterUnstakeRequestClaimed(opts *bind.FilterOpts, staker []common.Address, l2Strategys [][]common.Address, bridge []common.Address) (*StakingManagerUnstakeRequestClaimedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var l2StrategysRule []interface{}
	for _, l2StrategysItem := range l2Strategys {
		l2StrategysRule = append(l2StrategysRule, l2StrategysItem)
	}
	var bridgeRule []interface{}
	for _, bridgeItem := range bridge {
		bridgeRule = append(bridgeRule, bridgeItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnstakeRequestClaimed", stakerRule, l2StrategysRule, bridgeRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnstakeRequestClaimedIterator{contract: _StakingManager.contract, event: "UnstakeRequestClaimed", logs: logs, sub: sub}, nil
}

// WatchUnstakeRequestClaimed is a free log subscription operation binding the contract event 0xad6735aba298e7700470a993c5cbfd34a66ca15b17ac9b2b0984435d6f11f48d.
//
// Solidity: event UnstakeRequestClaimed(address indexed staker, address[] indexed l2Strategys, address indexed bridge, uint256 sourceChainId, uint256 destChainId)
func (_StakingManager *StakingManagerFilterer) WatchUnstakeRequestClaimed(opts *bind.WatchOpts, sink chan<- *StakingManagerUnstakeRequestClaimed, staker []common.Address, l2Strategys [][]common.Address, bridge []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var l2StrategysRule []interface{}
	for _, l2StrategysItem := range l2Strategys {
		l2StrategysRule = append(l2StrategysRule, l2StrategysItem)
	}
	var bridgeRule []interface{}
	for _, bridgeItem := range bridge {
		bridgeRule = append(bridgeRule, bridgeItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnstakeRequestClaimed", stakerRule, l2StrategysRule, bridgeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnstakeRequestClaimed)
				if err := _StakingManager.contract.UnpackLog(event, "UnstakeRequestClaimed", log); err != nil {
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

// ParseUnstakeRequestClaimed is a log parse operation binding the contract event 0xad6735aba298e7700470a993c5cbfd34a66ca15b17ac9b2b0984435d6f11f48d.
//
// Solidity: event UnstakeRequestClaimed(address indexed staker, address[] indexed l2Strategys, address indexed bridge, uint256 sourceChainId, uint256 destChainId)
func (_StakingManager *StakingManagerFilterer) ParseUnstakeRequestClaimed(log types.Log) (*StakingManagerUnstakeRequestClaimed, error) {
	event := new(StakingManagerUnstakeRequestClaimed)
	if err := _StakingManager.contract.UnpackLog(event, "UnstakeRequestClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUnstakeRequestedIterator is returned from FilterUnstakeRequested and is used to iterate over the raw logs and unpacked data for UnstakeRequested events raised by the StakingManager contract.
type StakingManagerUnstakeRequestedIterator struct {
	Event *StakingManagerUnstakeRequested // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnstakeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnstakeRequested)
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
		it.Event = new(StakingManagerUnstakeRequested)
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
func (it *StakingManagerUnstakeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnstakeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnstakeRequested represents a UnstakeRequested event raised by the StakingManager contract.
type StakingManagerUnstakeRequested struct {
	Staker              common.Address
	L2Strategy          common.Address
	EthAmount           *big.Int
	DETHLocked          *big.Int
	DestChainId         *big.Int
	UnStakeMessageNonce *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUnstakeRequested is a free log retrieval operation binding the contract event 0xdd412332aa89c96943c00eeac315cfc5e887074571f52af163514c8c34cc9dc3.
//
// Solidity: event UnstakeRequested(address indexed staker, address indexed l2Strategy, uint256 ethAmount, uint256 dETHLocked, uint256 destChainId, uint256 unStakeMessageNonce)
func (_StakingManager *StakingManagerFilterer) FilterUnstakeRequested(opts *bind.FilterOpts, staker []common.Address, l2Strategy []common.Address) (*StakingManagerUnstakeRequestedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var l2StrategyRule []interface{}
	for _, l2StrategyItem := range l2Strategy {
		l2StrategyRule = append(l2StrategyRule, l2StrategyItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnstakeRequested", stakerRule, l2StrategyRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnstakeRequestedIterator{contract: _StakingManager.contract, event: "UnstakeRequested", logs: logs, sub: sub}, nil
}

// WatchUnstakeRequested is a free log subscription operation binding the contract event 0xdd412332aa89c96943c00eeac315cfc5e887074571f52af163514c8c34cc9dc3.
//
// Solidity: event UnstakeRequested(address indexed staker, address indexed l2Strategy, uint256 ethAmount, uint256 dETHLocked, uint256 destChainId, uint256 unStakeMessageNonce)
func (_StakingManager *StakingManagerFilterer) WatchUnstakeRequested(opts *bind.WatchOpts, sink chan<- *StakingManagerUnstakeRequested, staker []common.Address, l2Strategy []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var l2StrategyRule []interface{}
	for _, l2StrategyItem := range l2Strategy {
		l2StrategyRule = append(l2StrategyRule, l2StrategyItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnstakeRequested", stakerRule, l2StrategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnstakeRequested)
				if err := _StakingManager.contract.UnpackLog(event, "UnstakeRequested", log); err != nil {
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

// ParseUnstakeRequested is a log parse operation binding the contract event 0xdd412332aa89c96943c00eeac315cfc5e887074571f52af163514c8c34cc9dc3.
//
// Solidity: event UnstakeRequested(address indexed staker, address indexed l2Strategy, uint256 ethAmount, uint256 dETHLocked, uint256 destChainId, uint256 unStakeMessageNonce)
func (_StakingManager *StakingManagerFilterer) ParseUnstakeRequested(log types.Log) (*StakingManagerUnstakeRequested, error) {
	event := new(StakingManagerUnstakeRequested)
	if err := _StakingManager.contract.UnpackLog(event, "UnstakeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerValidatorInitiatedIterator is returned from FilterValidatorInitiated and is used to iterate over the raw logs and unpacked data for ValidatorInitiated events raised by the StakingManager contract.
type StakingManagerValidatorInitiatedIterator struct {
	Event *StakingManagerValidatorInitiated // Event containing the contract specifics and raw log

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
func (it *StakingManagerValidatorInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerValidatorInitiated)
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
		it.Event = new(StakingManagerValidatorInitiated)
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
func (it *StakingManagerValidatorInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerValidatorInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerValidatorInitiated represents a ValidatorInitiated event raised by the StakingManager contract.
type StakingManagerValidatorInitiated struct {
	Id              [32]byte
	OperatorID      *big.Int
	Pubkey          []byte
	AmountDeposited *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorInitiated is a free log retrieval operation binding the contract event 0x15f16c2e13e50235799a97b981bf4a66c8cd86051f06aca745c5ff26f39b330e.
//
// Solidity: event ValidatorInitiated(bytes32 indexed id, uint256 indexed operatorID, bytes pubkey, uint256 amountDeposited)
func (_StakingManager *StakingManagerFilterer) FilterValidatorInitiated(opts *bind.FilterOpts, id [][32]byte, operatorID []*big.Int) (*StakingManagerValidatorInitiatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var operatorIDRule []interface{}
	for _, operatorIDItem := range operatorID {
		operatorIDRule = append(operatorIDRule, operatorIDItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "ValidatorInitiated", idRule, operatorIDRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerValidatorInitiatedIterator{contract: _StakingManager.contract, event: "ValidatorInitiated", logs: logs, sub: sub}, nil
}

// WatchValidatorInitiated is a free log subscription operation binding the contract event 0x15f16c2e13e50235799a97b981bf4a66c8cd86051f06aca745c5ff26f39b330e.
//
// Solidity: event ValidatorInitiated(bytes32 indexed id, uint256 indexed operatorID, bytes pubkey, uint256 amountDeposited)
func (_StakingManager *StakingManagerFilterer) WatchValidatorInitiated(opts *bind.WatchOpts, sink chan<- *StakingManagerValidatorInitiated, id [][32]byte, operatorID []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var operatorIDRule []interface{}
	for _, operatorIDItem := range operatorID {
		operatorIDRule = append(operatorIDRule, operatorIDItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "ValidatorInitiated", idRule, operatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerValidatorInitiated)
				if err := _StakingManager.contract.UnpackLog(event, "ValidatorInitiated", log); err != nil {
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

// ParseValidatorInitiated is a log parse operation binding the contract event 0x15f16c2e13e50235799a97b981bf4a66c8cd86051f06aca745c5ff26f39b330e.
//
// Solidity: event ValidatorInitiated(bytes32 indexed id, uint256 indexed operatorID, bytes pubkey, uint256 amountDeposited)
func (_StakingManager *StakingManagerFilterer) ParseValidatorInitiated(log types.Log) (*StakingManagerValidatorInitiated, error) {
	event := new(StakingManagerValidatorInitiated)
	if err := _StakingManager.contract.UnpackLog(event, "ValidatorInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
