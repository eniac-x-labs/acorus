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

// StakingManagerInit is an auto generated low-level Go binding around an user-defined struct.
type StakingManagerInit struct {
	Admin                  common.Address
	Manager                common.Address
	AllocatorService       common.Address
	InitiatorService       common.Address
	ReturnsAggregator      common.Address
	WithdrawalWallet       common.Address
	DapplinkBridge         common.Address
	DETH                   common.Address
	DepositContract        common.Address
	Oracle                 common.Address
	Pauser                 common.Address
	UnstakeRequestsManager common.Address
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ALLOCATOR_SERVICE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INITIATOR_SERVICE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_ALLOWLIST_MANAGER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_ALLOWLIST_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_MANAGER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOP_UP_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocateETH\",\"inputs\":[{\"name\":\"allocateToUnstakeRequestsManager\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allocateToDeposits\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocatedETHForDeposits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimUnstakeRequest\",\"inputs\":[{\"name\":\"l2Strategy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDETH\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dETHToETH\",\"inputs\":[{\"name\":\"dETHAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dapplinkBridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDepositContract\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethToDETH\",\"inputs\":[{\"name\":\"ethAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"exchangeAdjustmentRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializationBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"init\",\"type\":\"tuple\",\"internalType\":\"structStakingManager.Init\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"manager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocatorService\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initiatorService\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"returnsAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawalWallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dapplinkBridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dETH\",\"type\":\"address\",\"internalType\":\"contractIDETH\"},{\"name\":\"depositContract\",\"type\":\"address\",\"internalType\":\"contractIDepositContract\"},{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"contractIOracleReadRecord\"},{\"name\":\"pauser\",\"type\":\"address\",\"internalType\":\"contractIL1Pauser\"},{\"name\":\"unstakeRequestsManager\",\"type\":\"address\",\"internalType\":\"contractIUnstakeRequestsManager\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateValidatorsWithDeposits\",\"inputs\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structStakingManagerStorage.ValidatorParams[]\",\"components\":[{\"name\":\"operatorID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"withdrawalCredentials\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"expectedDepositRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isStakingAllowlist\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maximumDETHSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maximumDepositAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumDepositAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumStakeBound\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumUnstakeBound\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numInitiatedValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOracleReadRecord\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauser\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIL1Pauser\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveFromUnstakeRequestsManager\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"receiveReturns\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"reclaimAllocatedETHSurplus\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"returnsAggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExchangeAdjustmentRate\",\"inputs\":[{\"name\":\"exchangeAdjustmentRate_\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaximumDETHSupply\",\"inputs\":[{\"name\":\"maximumDETHSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaximumDepositAmount\",\"inputs\":[{\"name\":\"maximumDepositAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumDepositAmount\",\"inputs\":[{\"name\":\"minimumDepositAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumStakeBound\",\"inputs\":[{\"name\":\"minimumStakeBound_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumUnstakeBound\",\"inputs\":[{\"name\":\"minimumUnstakeBound_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStakingAllowlist\",\"inputs\":[{\"name\":\"isStakingAllowlist_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWithdrawalWallet\",\"inputs\":[{\"name\":\"withdrawalWallet_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"topUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"totalControlled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalDepositedInValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unallocatedETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unstakeRequest\",\"inputs\":[{\"name\":\"dethAmount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"minETHAmount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"l2Strategy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakeRequestInfo\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2strategy\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unstakeRequestsManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIUnstakeRequestsManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usedValidators\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawalWallet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AllocatedETHToDeposits\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocatedETHToUnstakeRequestsManager\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtocolConfigChanged\",\"inputs\":[{\"name\":\"setterSelector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"setterSignature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReturnsReceived\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ethAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"dETHAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeLaveAmount\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dETHLocked\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeRequestClaimed\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"l2Strategy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bridge\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeRequested\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"l2Strategy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ethAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"dETHLocked\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorInitiated\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"operatorID\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"amountDeposited\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DoesNotReceiveETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfiguration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDepositRoot\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWithdrawalCredentialsNotETH1\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}]},{\"type\":\"error\",\"name\":\"InvalidWithdrawalCredentialsWrongAddress\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWithdrawalCredentialsWrongLength\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MathOverflowedMulDiv\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaximumDETHSupplyExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaximumValidatorDepositExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinimumDepositAmountNotSatisfied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinimumStakeBoundNotSatisfied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinimumUnstakeBoundNotSatisfied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinimumValidatorDepositNotSatisfied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotDappLinkBridge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughDepositETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughUnallocatedETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotReturnsAggregator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotUnstakeRequestsManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Paused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PreviouslyUsedValidator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnstakeBelowMinimudETHAmount\",\"inputs\":[{\"name\":\"ethAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedMinimum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x608080604052346100b9577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166100aa57506001600160401b036002600160401b031982821601610065575b6040516132639081620000bf8239f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1388080610055565b63f92ee8a960e01b8152600490fd5b600080fdfe60806040526004361015610038575b3615610026576040516334352c7360e01b8152600490fd5b6040516334352c7360e01b8152600490fd5b6000803560e01c80628db05b14611fec57806301ffc9a714611f7b5780630208e4b514611ec257806304f36cc214611e095780630633af7614611de7578063080c279a14611dc957806312e9ead614611b21578063147d36d514611b025780631943190d14611aac57806319efd5c714611a835780631d2d35ce146119eb578063248a9ca3146119b1578063277f17581461198857806329d48704146118995780632f2ff15d146118165780633101d910146117ed57806335ead2a4146117cf57806336568abe146117875780633937c0b31461175e5780633f550b3a1461173557806342d3915d1461170f57806349336f0f146116f15780634a7d80b3146116c85780634bd6fc08146115535780634d014f7f14610f825780635915ded114610f105780635940d90b14610ef557806360a0f62814610ed75780636daa01a214610d065780636fce8ab214610ccb57806375796f7614610bf857806378abb49b14610bda5780637dc0d1d014610bb15780637dfcdd2914610b93578063808d663f14610b2e578063854a63f614610b0857806389e80ed314610acd5780639010d07c14610a7a57806391d1485414610a2157806399dd1deb146109895780639fd0506d14610960578063a217fddf14610944578063a694fc3a1461073e578063aab483d6146106a6578063ac1e2257146105e5578063b12de586146105c7578063b91590b2146105a9578063bb635c651461058b578063c151aa721461054f578063ca15c87314610518578063d547741f146104c9578063d70a6f311461042b578063d938e49214610402578063dc29f1de1461036c578063e55d6cc014610331578063e94ad65b14610308578063ea452b6d146102ea5763ed9daafb146102c1575061000e565b346102e75760203660031901126102e75760206102df6004356129a4565b604051908152f35b80fd5b50346102e757806003193601126102e7576020603654604051908152f35b50346102e757806003193601126102e757603c546040516001600160a01b039091168152602090f35b50346102e757806003193601126102e75760206040517f8ea5b4dbd68db0bf23bf4cda958b61a749f8c5aec6f2912d75a03246753ddd168152f35b50806003193601126102e7573360009081527f4d8059bacdb67d1f09126effca401868f01ce1cbafb88f03ac8f26295425f9e560205260409020547f5e4bd437d29fad01c10cdcfff414f0d6b0e84b96d2dade88d780d45b5630696b9060ff16156103e457506103de34603554612225565b60355580f35b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b50346102e757806003193601126102e7576045546040516001600160a01b039091168152602090f35b50346102e75760203660031901126102e757600435610448612a79565b80603b5560405190602082015260208152610462816120d8565b6000805160206131ce83398151915260405160408152602060408201527f7365744d6178696d756d4465706f736974416d6f756e742875696e7432353629606082015260806020820152806104c363d70a6f3160e01b946080830190612189565b0390a280f35b50346102e75760403660031901126102e7576105146004356104e9612099565b9080845260008051602061318e83398151915260205261050f6001604086200154612ac3565b612ba2565b5080f35b50346102e75760203660031901126102e7576040602091600435815260008051602061314e83398151915283522054604051908152f35b50806003193601126102e7576040546001600160a01b03163303610579576103de34603554612225565b604051637154fc4360e01b8152600490fd5b50346102e757806003193601126102e7576020603454604051908152f35b50346102e757806003193601126102e7576020604354604051908152f35b50346102e757806003193601126102e7576020603754604051908152f35b50346102e75760403660031901126102e7576105ff612099565b60408054815163d4be074f60e01b815260048035908201526001600160a01b0393841660248201529283916044918391165afa801561069b5782918391610653575b60408383825191151582526020820152f35b9150506040813d604011610693575b8161066f604093836120f3565b8101031261068f57604091506020610686826121ae565b91015138610641565b5080fd5b3d9150610662565b6040513d84823e3d90fd5b50346102e75760203660031901126102e7576004356106c3612a79565b80603a55604051906020820152602081526106dd816120d8565b6000805160206131ce83398151915260405160408152602060408201527f7365744d696e696d756d4465706f736974416d6f756e742875696e7432353629606082015260806020820152806104c363555a41eb60e11b946080830190612189565b5060208060031936011261068f576045546001600160a01b0390811691600435338490036109325760048284603f541660405192838092631ea7ca8960e01b82525afa9081156108755786916108f9575b506108e75760ff60425460a01c16610892575b603a548110610880576004826107b783612680565b94603d5416604051928380926318160ddd60e01b82525afa8015610875578690610841575b6107e7915084612225565b6044541061082f577f1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee909260409261082083603554612225565b6035558351928352820152a280f35b604051630935f98160e31b8152600490fd5b508281813d831161086e575b61085781836120f3565b81010312610869576107e790516107dc565b600080fd5b503d61084d565b6040513d88823e3d90fd5b604051630a3287a960e21b8152600490fd5b7fdec9d30de0821ad67aa5b141b13a539f584a19f99319e6041698a892b0e795598060005260008051602061318e8339815191528352604060002033600052835260ff60406000205416156103e457506107a2565b6040516313d0ff5960e31b8152600490fd5b90508281813d831161092b575b61091081836120f3565b8101031261092757610921906121ae565b3861078f565b8580fd5b503d610906565b60405163c68e02cf60e01b8152600490fd5b50346102e757806003193601126102e757602090604051908152f35b50346102e757806003193601126102e757603f546040516001600160a01b039091168152602090f35b50346102e75760203660031901126102e7576004356109a6612a79565b80603855604051906020820152602081526109c0816120d8565b6000805160206131ce83398151915260405160408152601f60408201527f7365744d696e696d756d556e7374616b65426f756e642875696e743235362900606082015260806020820152806104c36399dd1deb60e01b946080830190612189565b50346102e75760403660031901126102e7576040610a3d612099565b91600435815260008051602061318e833981519152602052209060018060a01b0316600052602052602060ff604060002054166040519015158152f35b50346102e75760403660031901126102e757610ab4602091600435815260008051602061314e833981519152835260406024359120612f38565b905460405160039290921b1c6001600160a01b03168152f35b50346102e757806003193601126102e75760206040517fdec9d30de0821ad67aa5b141b13a539f584a19f99319e6041698a892b0e795598152f35b50346102e757806003193601126102e757602060018060a01b0360405416604051908152f35b50806003193601126102e7576042546001600160a01b03163303610b82577f4cbb9d73b003a252cee3f2ee51d8d65a562af35eebb23730dd4a76d68127b3706020604051348152a16103de34603554612225565b604051626310df60e31b8152600490fd5b50346102e757806003193601126102e7576020603554604051908152f35b50346102e757806003193601126102e757603e546040516001600160a01b039091168152602090f35b50346102e757806003193601126102e7576020603b54604051908152f35b50346102e75760203660031901126102e757610c126120af565b610c1a612a79565b6001600160a01b03168015610cb957806bffffffffffffffffffffffff60a01b604154161760415560405190602082015260208152610c58816120d8565b6000805160206131ce83398151915260405160408152601c60408201527f7365745769746864726177616c57616c6c657428616464726573732900000000606082015260806020820152806104c3633abcb7bb60e11b946080830190612189565b60405163d92e233d60e01b8152600490fd5b50346102e757806003193601126102e75760206040517f5e4bd437d29fad01c10cdcfff414f0d6b0e84b96d2dade88d780d45b5630696b8152f35b50346102e75760403660031901126102e7573360009081527ff79084f49a5c4fa4c48f70bba1e67b61c2b9ca8b3d302dc944c028fdea010b82602090815260409091205460243591600435916000805160206131ae8339815191529060ff16156103e45750603f54604051637ee56d2f60e11b81526001600160a01b03929182908290600490829087165afa908115610875578691610ea2575b506108e757610daf8484612225565b93603554809511610e9057610dcf8695610dc98387612225565b90612269565b6035558181610e51575b505082610de557505050f35b7ffe89805cf5299ef9fbd1d1ddefb8dcc3fa9408064d2ea31e3fca6565768f521790604051848152a160405416803b15610e4d578290600460405180948193632689dfd360e11b83525af1801561069b57610e3d5750f35b610e46906120c5565b6102e75780f35b5050fd5b81610e7f7f9d04ecb465d2c8754acb798a22293dd26215a1c2f7a2a56607afa215c1aadc7793603654612225565b603655604051908152a13881610dd9565b6040516396b0c75160e01b8152600490fd5b90508181813d8311610ed0575b610eb981836120f3565b8101031261092757610eca906121ae565b38610da0565b503d610eaf565b50346102e757806003193601126102e7576020603354604051908152f35b50346102e757806003193601126102e75760206102df6127bc565b50346102e75760203660031901126102e757600435906001600160401b0382116102e757366023830112156102e757602060ff610f6c82610f593660048801356024890161212f565b8160405193828580945193849201612166565b8101603281520301902054166040519015158152f35b50346102e7576101808060031936011261068f576040519081018181106001600160401b0382111761153d57604052610fb96120af565b8152610fc3612099565b6020820152610fd0612083565b60408201526064356001600160a01b03811681036108695760608201526084356001600160a01b038116810361086957608082015260a4356001600160a01b03811681036108695760a082015260c4356001600160a01b03811681036108695760c082015260e4356001600160a01b03811681036115395760e0820152610104356001600160a01b038116810361153957610100820152610124356001600160a01b038116810361153957610120820152610144356001600160a01b038116810361153957610140820152610164356001600160a01b03811681036115395761016082015260008051602061320e833981519152546001600160401b038116158061152b575b60016001600160401b038316149081611521575b159081611518575b506115065760016001600160401b031982161760008051602061320e8339815191525560ff8160401c16156114d9575b60ff60008051602061320e8339815191525460401c16156114c757815160ff92906001600160a01b031661115581612bea565b61149e575b5060208101516001600160a01b031661117281612c78565b611467575b5060408101516001600160a01b031661118f81612d10565b611430575b5060608101516001600160a01b03166111ac81612da8565b6113f9575b507f8ea5b4dbd68db0bf23bf4cda958b61a749f8c5aec6f2912d75a03246753ddd16600081815260008051602061318e8339815191526020527f0cf65ee5f25796078ddc4621a40807239bac87a14b48afca00334599cfbf78c580546000805160206131ee833981519152918290557fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff928490849080a47fdec9d30de0821ad67aa5b141b13a539f584a19f99319e6041698a892b0e79559806000526001604060002001918383549355600080a460018060a01b0360e0820151166bffffffffffffffffffffffff60a01b9081603d541617603d5560018060a01b036101008301511681603c541617603c5560018060a01b036101208301511681603e541617603e5560018060a01b036101408301511681603f541617603f5560018060a01b03608083015116906042549260018060a01b036101608201511682604054161760405560018060a01b0360c08201511682604554161760455560a0600180821b039101511690604154161760415567016345785d8a0000603755662386f26fc100006038556801bc16d674ec80000080603a55603b55600160a01b916affffffffffffffffffffff60a81b16171760425543604355683782dace9d9000000060445560401c161561139f5780f35b68ff00000000000000001960008051602061320e833981519152541660008051602061320e833981519152557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b6114299060008051602061316e833981519152865260008051602061314e83398151915260205260408620612f50565b50386111b1565b611460906000805160206131ae833981519152865260008051602061314e83398151915260205260408620612f50565b5038611194565b611497906000805160206131ee833981519152865260008051602061314e83398151915260205260408620612f50565b5038611177565b6114c09085805260008051602061314e83398151915260205260408620612f50565b503861115a565b604051631afcd79f60e31b8152600490fd5b68ffffffffffffffffff198116680100000000000000011760008051602061320e83398151915255611122565b60405163f92ee8a960e01b8152600490fd5b905015386110f2565b303b1591506110ea565b5060ff8160401c16156110d6565b8280fd5b634e487b7160e01b600052604160045260246000fd5b50346102e75760a03660031901126102e75761156d6120af565b611575612099565b60455460443592906001600160a01b039060643590821633036109325781603f54169360405180956345b09c8d60e11b825281600460209889935afa9081156116bd578891611684575b506108e7578660a492848088971693169485846040518b8152848a8201527fdefb23229f98ca9823b290c7a4ca8908d69e2a25a6fa9e0b65b3ffad35798eb360403392a460405416906040519889968795633c82491960e01b8752600487015260248601526044850152606484015260843560848401525af1801561167957611646578280f35b81813d8311611672575b61165a81836120f3565b8101031261068f5761166b906121ae565b5038808280f35b503d611650565b6040513d85823e3d90fd5b90508581813d83116116b6575b61169b81836120f3565b810103126116b2576116ac906121ae565b386115bf565b8780fd5b503d611691565b6040513d8a823e3d90fd5b50346102e757806003193601126102e7576041546040516001600160a01b039091168152602090f35b50346102e757806003193601126102e7576020604454604051908152f35b50346102e757806003193601126102e757602060ff60425460a01c166040519015158152f35b50346102e757806003193601126102e7576042546040516001600160a01b039091168152602090f35b50346102e757806003193601126102e75760206040516000805160206131ee8339815191528152f35b50346102e75760403660031901126102e7576117a1612099565b336001600160a01b038216036117bd5761051490600435612ba2565b60405163334bd91960e11b8152600490fd5b50346102e757806003193601126102e7576020603854604051908152f35b50346102e757806003193601126102e75760206040516000805160206131ae8339815191528152f35b50346102e75760403660031901126102e757600435611833612099565b81835260008051602061318e8339815191526020526118586001604085200154612ac3565b6118628183612e40565b61186a578280f35b600091825260008051602061314e833981519152602052604090912061166b916001600160a01b031690612f50565b50346102e75760203660031901126102e75760043561ffff811680910361068f576118c2612a79565b6103e88111611976576127108111611962578061ffff196039541617603955604051906020820152602081526118f7816120d8565b6000805160206131ce83398151915260405160408152602160408201527f73657445786368616e676541646a7573746d656e74526174652875696e7431366060820152602960f81b608082015260a06020820152806104c3630a7521c160e21b9460a0830190612189565b634e487b7160e01b82526001600452602482fd5b60405163c52a9bd360e01b8152600490fd5b50346102e757806003193601126102e757603d546040516001600160a01b039091168152602090f35b50346102e75760203660031901126102e75760016040602092600435815260008051602061318e8339815191528452200154604051908152f35b50346102e75760203660031901126102e757600435611a08612a79565b8060445560405190602082015260208152611a22816120d8565b6000805160206131ce83398151915260405160408152601d60408201527f7365744d6178696d756d44455448537570706c792875696e7432353629000000606082015260806020820152806104c3630e969ae760e11b946080830190612189565b50346102e757806003193601126102e757602060405160008051602061316e8339815191528152f35b50346102e757806003193601126102e757611ac5612a79565b60405481906001600160a01b0316803b15611aff5781809160046040518094819363596a15a360e11b83525af1801561069b57610e3d5750f35b50fd5b50346102e75760203660031901126102e75760206102df600435612680565b50346102e75760803660031901126102e7576004356001600160801b039081811680910361153957602491823591818316809303611d1857611b61612083565b603f546040516345b09c8d60e11b81526001600160a01b0395602095909260643592908790829060049082908c165afa908115611dbe578a91611d85575b506108e7576038548510611d7457611bb6856129a4565b1691808310611d575750856040541688813b156102e7578760a48a9683604051958694859362448d8d60e91b855233600486015216809a8401528a60448401528860648401528760848401525af18015611d4c57611d39575b50604051918252838583015260408201527f7fd59a2c98b4d595be25d7976e588f7b6189ace6d41522f11868d951c293436760603392a382603d541692604054169060405190838201926323b872dd60e01b84523387840152604483015260648201526064815260a081018181106001600160401b03821117611d245760405251611cca918691829182875af13d15611d1c573d90611cad82612114565b91611cbb60405193846120f3565b82523d878584013e5b846130ca565b8051918215159283611cf7575b505050611ce2578280f35b60405190635274afe760e01b82526004820152fd5b82935091819281010312611d1857611d0f91016121ae565b15388080611cd7565b8480fd5b606090611cc4565b85634e487b7160e01b60005260416004526000fd5b611d45909891986120c5565b9638611c0f565b6040513d8b823e3d90fd5b8260449189604051926347f961e960e11b84526004840152820152fd5b60405162771ba760e71b8152600490fd5b90508681813d8311611db7575b611d9c81836120f3565b81010312611db357611dad906121ae565b38611b9f565b8980fd5b503d611d92565b6040513d8c823e3d90fd5b50346102e757806003193601126102e7576020603a54604051908152f35b50346102e757806003193601126102e757602061ffff60395416604051908152f35b50346102e75760203660031901126102e75760043580151580910361068f57611e30612a79565b6042805460ff60a01b191660a083901b60ff60a01b16179055604051602080820192909252908152611e61816120d8565b6000805160206131ce83398151915260405160408152601960408201527f7365745374616b696e67416c6c6f776c69737428626f6f6c2900000000000000606082015260806020820152806104c3630279b66160e11b946080830190612189565b50346102e75760403660031901126102e7576001600160401b039060043582811161068f573660238201121561068f57806004013592831161068f573660248460051b8301011161068f5760008051602061316e8339815191529283835260008051602061318e8339815191526020526040832033845260205260ff60408420541615611f5c5790611f5991602480359201612276565b80f35b60405163e2517d3f60e01b815233600482015260248101859052604490fd5b50346102e75760203660031901126102e75760043563ffffffff60e01b811680910361068f57602090635a05180f60e01b8114908115611fc1575b506040519015158152f35b637965db0b60e01b811491508115611fdb575b5082611fb6565b6301ffc9a760e01b14905082611fd4565b50346102e75760203660031901126102e757600435612009612a79565b8060375560405190602082015260208152612023816120d8565b6000805160206131ce83398151915260405160408152601d60408201527f7365744d696e696d756d5374616b65426f756e642875696e7432353629000000606082015260806020820152806104c3628db05b60e01b946080830190612189565b604435906001600160a01b038216820361086957565b602435906001600160a01b038216820361086957565b600435906001600160a01b038216820361086957565b6001600160401b03811161153d57604052565b604081019081106001600160401b0382111761153d57604052565b90601f801991011681019081106001600160401b0382111761153d57604052565b6001600160401b03811161153d57601f01601f191660200190565b92919261213b82612114565b9161214960405193846120f3565b829481845281830111610869578281602093846000960137010152565b60005b8381106121795750506000910152565b8181015183820152602001612169565b906020916121a281518092818552858086019101612166565b601f01601f1916010190565b5190811515820361086957565b91908110156121dd5760051b8101359060be1981360301821215610869570190565b634e487b7160e01b600052603260045260246000fd5b903590601e198136030182121561086957018035906001600160401b0382116108695760200191813603831361086957565b9190820180921161223257565b634e487b7160e01b600052601160045260246000fd5b908060209392818452848401376000828201840152601f01601f1916010190565b9190820391821161223257565b603f5460408051637750955b60e11b8152909492936001600160a01b03936004936020939084908290879082908a165afa9081156124175760009161264b575b5061263b57811561263257838386603c541689519283809263c5f2892f60e01b82525afa90811561241757600091612605575b508091036125ef57506000805b828110612432575060365480821161242257906123168161232193612269565b603655603354612225565b60335561233081603454612225565b60345560005b8181106123465750505050505050565b6123518183886121bb565b85603c5416612362898301836121f3565b909261237160608201826121f3565b9093608091612382838501856121f3565b9091833b15610869578f978d928d946123d16123e0946123bd60009d519e8f9d8e9c8d9a6304512a2360e31b8c528b015260848a0191612248565b9060031995868984030160248a0152612248565b92858403016044860152612248565b60a0860135606483015203930135905af18015612417579060019291612408575b5001612336565b612411906120c5565b38612401565b88513d6000823e3d90fd5b87516307bb2bd760e21b81528590fd5b60ff61243f82858a6121bb565b928984019061244e82866121f3565b8c5194918190863784019388816032968781520301902054166125df578685013592603a5484106125cf57603b5484116125bf576060612490818801886121f3565b908a82036125a857600c918083116108695781356001600160a01b0319166001600160f81b0319810161258d57508b11610869570135901c8a60415416810361257757506125678c6125308686957f15f16c2e13e50235799a97b981bf4a66c8cd86051f06aca745c5ff26f39b330e958d61250f60019c9b9a8e6121f3565b928388519485938437820190815203019020805460ff19168b179055612225565b9761255561254861254187846121f3565b369161212f565b8c815191012095826121f3565b93908351948486958652850191612248565b958b83015235940390a3016122f6565b896024918e5191631b4d561960e01b8352820152fd5b90508f602493508d925051916308ebf56560e01b8352820152fd5b508a6024918f5191639b0ec52760e01b8352820152fd5b8b51630c25396f60e11b81528990fd5b8b516305cacc5560e21b81528990fd5b8a5163932c5b0d60e01b81528890fd5b83602491885191631497ae9360e01b8352820152fd5b90508381813d831161262b575b61261c81836120f3565b810103126108695751386122e9565b503d612612565b50505050505050565b86516313d0ff5960e31b81528490fd5b90508381813d8311612679575b61266281836120f3565b8101031261086957612673906121ae565b386122b6565b503d612658565b603d546040516318160ddd60e01b8082526001600160a01b03909216916020918281600481875afa90811561275557600091612767575b501561276157819060046040518095819382525afa90811561275557600091612729575b50603954612710925061ffff90811683038181116122325716908181029181830414901517156122325761270d6127bc565b8281029281840414901517156122325761272692612af6565b90565b82813d831161274e575b61273d81836120f3565b810103126102e757505180386126db565b503d612733565b6040513d6000823e3d90fd5b50505090565b90508281813d831161278d575b61277e81836120f3565b810103126108695751386126b7565b503d612774565b51906001600160401b038216820361086957565b51906001600160801b038216820361086957565b603e546040805163079d004d60e51b815292916001600160a01b03906101009081908690600490829086165afa948515612999576000956128b4575b505060206128446004949561281260355460365490612225565b60c061283a603354926128346001600160801b03948560e08801511690612269565b90612225565b9201511690612225565b918354168351948580926316d3df1560e31b82525afa9182156128aa5750600091612874575b6127269250612225565b90506020823d6020116128a2575b8161288f602093836120f3565b810103126108695761272691519061286a565b3d9150612882565b513d6000823e3d90fd5b90948582813d8311612992575b6128cb81836120f3565b810103126102e7578351958601908682106001600160401b0383111761297e57506004949561297160e060209461284494885261290781612794565b8452612914868201612794565b86850152612923888201612794565b8885015261293360608201612794565b6060850152612944608082016127a8565b608085015261295560a082016127a8565b60a085015261296660c082016127a8565b60c0850152016127a8565b60e08201529594506127f8565b634e487b7160e01b81526041600452602490fd5b503d6128c1565b83513d6000823e3d90fd5b603d546040516318160ddd60e01b8082526001600160a01b0390921692602092918381600481885afa90811561275557600091612a4c575b5015612a4557826129eb6127bc565b9260046040518097819382525afa92831561275557600093612a13575b506127269350612af6565b90925083813d8311612a3e575b612a2a81836120f3565b810103126108695761272692519138612a08565b503d612a20565b9250505090565b90508381813d8311612a72575b612a6381836120f3565b810103126108695751386129dc565b503d612a59565b3360009081527fc5abc2d863e77a76627fe1702320d3afc3f93a9ad30aebcafab5d12854da5c0f60205260409020546000805160206131ee8339815191529060ff16156103e45750565b8060005260008051602061318e83398151915260205260406000203360005260205260ff60406000205416156103e45750565b90918282029160001984820993838086109503948086039514612b7e5784831115612b6c5782910981600003821680920460028082600302188083028203028083028203028083028203028083028203028083028203028092029003029360018380600003040190848311900302920304170290565b60405163227bc15360e01b8152600490fd5b505080925015612b8c570490565b634e487b7160e01b600052601260045260246000fd5b612bac8282612eb5565b9182612bb757505090565b600091825260008051602061314e8339815191526020526040909120612be6916001600160a01b031690612fda565b5090565b6001600160a01b031660008181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604081205490919060008051602061318e8339815191529060ff16612c73578280526020526040822081835260205260408220600160ff19825416179055339160008051602061312e8339815191528180a4600190565b505090565b6001600160a01b031660008181527fc5abc2d863e77a76627fe1702320d3afc3f93a9ad30aebcafab5d12854da5c0f60205260408120549091906000805160206131ee8339815191529060008051602061318e8339815191529060ff16612761578184526020526040832082845260205260408320600160ff1982541617905560008051602061312e833981519152339380a4600190565b6001600160a01b031660008181527ff79084f49a5c4fa4c48f70bba1e67b61c2b9ca8b3d302dc944c028fdea010b8260205260408120549091906000805160206131ae8339815191529060008051602061318e8339815191529060ff16612761578184526020526040832082845260205260408320600160ff1982541617905560008051602061312e833981519152339380a4600190565b6001600160a01b031660008181527ff1e23661530d14d05c9291333c54312223931d3f1ab2285de8cf548f5a18240d602052604081205490919060008051602061316e8339815191529060008051602061318e8339815191529060ff16612761578184526020526040832082845260205260408320600160ff1982541617905560008051602061312e833981519152339380a4600190565b9060009180835260008051602061318e83398151915280602052604084209260018060a01b03169283855260205260ff60408520541615600014612761578184526020526040832082845260205260408320600160ff1982541617905560008051602061312e833981519152339380a4600190565b9060009180835260008051602061318e83398151915280602052604084209260018060a01b03169283855260205260ff60408520541660001461276157818452602052604083208284526020526040832060ff1981541690557ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b339380a4600190565b80548210156121dd5760005260206000200190600090565b91906001830160009082825280602052604082205415600014612fd45784549468010000000000000000861015612fc05783612fb0612f99886001604098999a01855584612f38565b819391549060031b91821b91600019901b19161790565b9055549382526020522055600190565b634e487b7160e01b83526041600452602483fd5b50925050565b906001820190600092818452826020526040842054908115156000146130c357600019918083018181116130af5782549084820191821161309b57818103613066575b50505080548015613052578201916130358383612f38565b909182549160031b1b191690555582526020526040812055600190565b634e487b7160e01b86526031600452602486fd5b613086613076612f999386612f38565b90549060031b1c92839286612f38565b9055865284602052604086205538808061301d565b634e487b7160e01b88526011600452602488fd5b634e487b7160e01b87526011600452602487fd5b5050505090565b906130f157508051156130df57805190602001fd5b604051630a12f52160e11b8152600490fd5b81511580613124575b613102575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b156130fa56fe2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0dc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000e30bb2df90b65284acd0e8b5ebe3483bb2bbe65a08e43f0f9e8300fd8607ee1102dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800e6ef7125bfa79685f3bd2e4c4cea243c1e988ebbc0801ab7641ae36b9e2c529101d854e8dde9402801a4c6f2840193465752abfad61e0bb7c4258d526ae42e74a6b5d83d32632203555cb9b2c2f68a8d94da48cadd9266ac0d17babedb52ea5bf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220ddd83b82d457db3f9a65b2375019ec76f04a0db3f64f210e2fbc4c6c6453e49e64736f6c63430008180033",
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

// DETH is a free data retrieval call binding the contract method 0x277f1758.
//
// Solidity: function dETH() view returns(address)
func (_StakingManager *StakingManagerCaller) DETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "dETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DETH is a free data retrieval call binding the contract method 0x277f1758.
//
// Solidity: function dETH() view returns(address)
func (_StakingManager *StakingManagerSession) DETH() (common.Address, error) {
	return _StakingManager.Contract.DETH(&_StakingManager.CallOpts)
}

// DETH is a free data retrieval call binding the contract method 0x277f1758.
//
// Solidity: function dETH() view returns(address)
func (_StakingManager *StakingManagerCallerSession) DETH() (common.Address, error) {
	return _StakingManager.Contract.DETH(&_StakingManager.CallOpts)
}

// DETHToETH is a free data retrieval call binding the contract method 0xed9daafb.
//
// Solidity: function dETHToETH(uint256 dETHAmount) view returns(uint256)
func (_StakingManager *StakingManagerCaller) DETHToETH(opts *bind.CallOpts, dETHAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "dETHToETH", dETHAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DETHToETH is a free data retrieval call binding the contract method 0xed9daafb.
//
// Solidity: function dETHToETH(uint256 dETHAmount) view returns(uint256)
func (_StakingManager *StakingManagerSession) DETHToETH(dETHAmount *big.Int) (*big.Int, error) {
	return _StakingManager.Contract.DETHToETH(&_StakingManager.CallOpts, dETHAmount)
}

// DETHToETH is a free data retrieval call binding the contract method 0xed9daafb.
//
// Solidity: function dETHToETH(uint256 dETHAmount) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) DETHToETH(dETHAmount *big.Int) (*big.Int, error) {
	return _StakingManager.Contract.DETHToETH(&_StakingManager.CallOpts, dETHAmount)
}

// DapplinkBridge is a free data retrieval call binding the contract method 0xd938e492.
//
// Solidity: function dapplinkBridge() view returns(address)
func (_StakingManager *StakingManagerCaller) DapplinkBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "dapplinkBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DapplinkBridge is a free data retrieval call binding the contract method 0xd938e492.
//
// Solidity: function dapplinkBridge() view returns(address)
func (_StakingManager *StakingManagerSession) DapplinkBridge() (common.Address, error) {
	return _StakingManager.Contract.DapplinkBridge(&_StakingManager.CallOpts)
}

// DapplinkBridge is a free data retrieval call binding the contract method 0xd938e492.
//
// Solidity: function dapplinkBridge() view returns(address)
func (_StakingManager *StakingManagerCallerSession) DapplinkBridge() (common.Address, error) {
	return _StakingManager.Contract.DapplinkBridge(&_StakingManager.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_StakingManager *StakingManagerCaller) DepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "depositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_StakingManager *StakingManagerSession) DepositContract() (common.Address, error) {
	return _StakingManager.Contract.DepositContract(&_StakingManager.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_StakingManager *StakingManagerCallerSession) DepositContract() (common.Address, error) {
	return _StakingManager.Contract.DepositContract(&_StakingManager.CallOpts)
}

// EthToDETH is a free data retrieval call binding the contract method 0x147d36d5.
//
// Solidity: function ethToDETH(uint256 ethAmount) view returns(uint256)
func (_StakingManager *StakingManagerCaller) EthToDETH(opts *bind.CallOpts, ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "ethToDETH", ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthToDETH is a free data retrieval call binding the contract method 0x147d36d5.
//
// Solidity: function ethToDETH(uint256 ethAmount) view returns(uint256)
func (_StakingManager *StakingManagerSession) EthToDETH(ethAmount *big.Int) (*big.Int, error) {
	return _StakingManager.Contract.EthToDETH(&_StakingManager.CallOpts, ethAmount)
}

// EthToDETH is a free data retrieval call binding the contract method 0x147d36d5.
//
// Solidity: function ethToDETH(uint256 ethAmount) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) EthToDETH(ethAmount *big.Int) (*big.Int, error) {
	return _StakingManager.Contract.EthToDETH(&_StakingManager.CallOpts, ethAmount)
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

// MinimumStakeBound is a free data retrieval call binding the contract method 0xb12de586.
//
// Solidity: function minimumStakeBound() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MinimumStakeBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "minimumStakeBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStakeBound is a free data retrieval call binding the contract method 0xb12de586.
//
// Solidity: function minimumStakeBound() view returns(uint256)
func (_StakingManager *StakingManagerSession) MinimumStakeBound() (*big.Int, error) {
	return _StakingManager.Contract.MinimumStakeBound(&_StakingManager.CallOpts)
}

// MinimumStakeBound is a free data retrieval call binding the contract method 0xb12de586.
//
// Solidity: function minimumStakeBound() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MinimumStakeBound() (*big.Int, error) {
	return _StakingManager.Contract.MinimumStakeBound(&_StakingManager.CallOpts)
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

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_StakingManager *StakingManagerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_StakingManager *StakingManagerSession) Oracle() (common.Address, error) {
	return _StakingManager.Contract.Oracle(&_StakingManager.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_StakingManager *StakingManagerCallerSession) Oracle() (common.Address, error) {
	return _StakingManager.Contract.Oracle(&_StakingManager.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_StakingManager *StakingManagerCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_StakingManager *StakingManagerSession) Pauser() (common.Address, error) {
	return _StakingManager.Contract.Pauser(&_StakingManager.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_StakingManager *StakingManagerCallerSession) Pauser() (common.Address, error) {
	return _StakingManager.Contract.Pauser(&_StakingManager.CallOpts)
}

// ReturnsAggregator is a free data retrieval call binding the contract method 0x3f550b3a.
//
// Solidity: function returnsAggregator() view returns(address)
func (_StakingManager *StakingManagerCaller) ReturnsAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "returnsAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReturnsAggregator is a free data retrieval call binding the contract method 0x3f550b3a.
//
// Solidity: function returnsAggregator() view returns(address)
func (_StakingManager *StakingManagerSession) ReturnsAggregator() (common.Address, error) {
	return _StakingManager.Contract.ReturnsAggregator(&_StakingManager.CallOpts)
}

// ReturnsAggregator is a free data retrieval call binding the contract method 0x3f550b3a.
//
// Solidity: function returnsAggregator() view returns(address)
func (_StakingManager *StakingManagerCallerSession) ReturnsAggregator() (common.Address, error) {
	return _StakingManager.Contract.ReturnsAggregator(&_StakingManager.CallOpts)
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

// TotalControlled is a free data retrieval call binding the contract method 0x5940d90b.
//
// Solidity: function totalControlled() view returns(uint256)
func (_StakingManager *StakingManagerCaller) TotalControlled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "totalControlled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalControlled is a free data retrieval call binding the contract method 0x5940d90b.
//
// Solidity: function totalControlled() view returns(uint256)
func (_StakingManager *StakingManagerSession) TotalControlled() (*big.Int, error) {
	return _StakingManager.Contract.TotalControlled(&_StakingManager.CallOpts)
}

// TotalControlled is a free data retrieval call binding the contract method 0x5940d90b.
//
// Solidity: function totalControlled() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) TotalControlled() (*big.Int, error) {
	return _StakingManager.Contract.TotalControlled(&_StakingManager.CallOpts)
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

// UnstakeRequestsManager is a free data retrieval call binding the contract method 0x854a63f6.
//
// Solidity: function unstakeRequestsManager() view returns(address)
func (_StakingManager *StakingManagerCaller) UnstakeRequestsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "unstakeRequestsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnstakeRequestsManager is a free data retrieval call binding the contract method 0x854a63f6.
//
// Solidity: function unstakeRequestsManager() view returns(address)
func (_StakingManager *StakingManagerSession) UnstakeRequestsManager() (common.Address, error) {
	return _StakingManager.Contract.UnstakeRequestsManager(&_StakingManager.CallOpts)
}

// UnstakeRequestsManager is a free data retrieval call binding the contract method 0x854a63f6.
//
// Solidity: function unstakeRequestsManager() view returns(address)
func (_StakingManager *StakingManagerCallerSession) UnstakeRequestsManager() (common.Address, error) {
	return _StakingManager.Contract.UnstakeRequestsManager(&_StakingManager.CallOpts)
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

// ClaimUnstakeRequest is a paid mutator transaction binding the contract method 0x4bd6fc08.
//
// Solidity: function claimUnstakeRequest(address l2Strategy, address bridge, uint256 sourceChainId, uint256 destChainId, uint256 gasLimit) returns()
func (_StakingManager *StakingManagerTransactor) ClaimUnstakeRequest(opts *bind.TransactOpts, l2Strategy common.Address, bridge common.Address, sourceChainId *big.Int, destChainId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "claimUnstakeRequest", l2Strategy, bridge, sourceChainId, destChainId, gasLimit)
}

// ClaimUnstakeRequest is a paid mutator transaction binding the contract method 0x4bd6fc08.
//
// Solidity: function claimUnstakeRequest(address l2Strategy, address bridge, uint256 sourceChainId, uint256 destChainId, uint256 gasLimit) returns()
func (_StakingManager *StakingManagerSession) ClaimUnstakeRequest(l2Strategy common.Address, bridge common.Address, sourceChainId *big.Int, destChainId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.ClaimUnstakeRequest(&_StakingManager.TransactOpts, l2Strategy, bridge, sourceChainId, destChainId, gasLimit)
}

// ClaimUnstakeRequest is a paid mutator transaction binding the contract method 0x4bd6fc08.
//
// Solidity: function claimUnstakeRequest(address l2Strategy, address bridge, uint256 sourceChainId, uint256 destChainId, uint256 gasLimit) returns()
func (_StakingManager *StakingManagerTransactorSession) ClaimUnstakeRequest(l2Strategy common.Address, bridge common.Address, sourceChainId *big.Int, destChainId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.ClaimUnstakeRequest(&_StakingManager.TransactOpts, l2Strategy, bridge, sourceChainId, destChainId, gasLimit)
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

// Initialize is a paid mutator transaction binding the contract method 0x4d014f7f.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,address,address,address,address) init) returns()
func (_StakingManager *StakingManagerTransactor) Initialize(opts *bind.TransactOpts, init StakingManagerInit) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "initialize", init)
}

// Initialize is a paid mutator transaction binding the contract method 0x4d014f7f.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,address,address,address,address) init) returns()
func (_StakingManager *StakingManagerSession) Initialize(init StakingManagerInit) (*types.Transaction, error) {
	return _StakingManager.Contract.Initialize(&_StakingManager.TransactOpts, init)
}

// Initialize is a paid mutator transaction binding the contract method 0x4d014f7f.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,address,address,address,address) init) returns()
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

// SetMaximumDepositAmount is a paid mutator transaction binding the contract method 0xd70a6f31.
//
// Solidity: function setMaximumDepositAmount(uint256 maximumDepositAmount_) returns()
func (_StakingManager *StakingManagerTransactor) SetMaximumDepositAmount(opts *bind.TransactOpts, maximumDepositAmount_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setMaximumDepositAmount", maximumDepositAmount_)
}

// SetMaximumDepositAmount is a paid mutator transaction binding the contract method 0xd70a6f31.
//
// Solidity: function setMaximumDepositAmount(uint256 maximumDepositAmount_) returns()
func (_StakingManager *StakingManagerSession) SetMaximumDepositAmount(maximumDepositAmount_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMaximumDepositAmount(&_StakingManager.TransactOpts, maximumDepositAmount_)
}

// SetMaximumDepositAmount is a paid mutator transaction binding the contract method 0xd70a6f31.
//
// Solidity: function setMaximumDepositAmount(uint256 maximumDepositAmount_) returns()
func (_StakingManager *StakingManagerTransactorSession) SetMaximumDepositAmount(maximumDepositAmount_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMaximumDepositAmount(&_StakingManager.TransactOpts, maximumDepositAmount_)
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

// SetMinimumStakeBound is a paid mutator transaction binding the contract method 0x008db05b.
//
// Solidity: function setMinimumStakeBound(uint256 minimumStakeBound_) returns()
func (_StakingManager *StakingManagerTransactor) SetMinimumStakeBound(opts *bind.TransactOpts, minimumStakeBound_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setMinimumStakeBound", minimumStakeBound_)
}

// SetMinimumStakeBound is a paid mutator transaction binding the contract method 0x008db05b.
//
// Solidity: function setMinimumStakeBound(uint256 minimumStakeBound_) returns()
func (_StakingManager *StakingManagerSession) SetMinimumStakeBound(minimumStakeBound_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMinimumStakeBound(&_StakingManager.TransactOpts, minimumStakeBound_)
}

// SetMinimumStakeBound is a paid mutator transaction binding the contract method 0x008db05b.
//
// Solidity: function setMinimumStakeBound(uint256 minimumStakeBound_) returns()
func (_StakingManager *StakingManagerTransactorSession) SetMinimumStakeBound(minimumStakeBound_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMinimumStakeBound(&_StakingManager.TransactOpts, minimumStakeBound_)
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

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 stakeAmount) payable returns()
func (_StakingManager *StakingManagerTransactor) Stake(opts *bind.TransactOpts, stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "stake", stakeAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 stakeAmount) payable returns()
func (_StakingManager *StakingManagerSession) Stake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.Stake(&_StakingManager.TransactOpts, stakeAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 stakeAmount) payable returns()
func (_StakingManager *StakingManagerTransactorSession) Stake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.Stake(&_StakingManager.TransactOpts, stakeAmount)
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

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakingManager *StakingManagerTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _StakingManager.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakingManager *StakingManagerSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StakingManager.Contract.Fallback(&_StakingManager.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakingManager *StakingManagerTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StakingManager.Contract.Fallback(&_StakingManager.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingManager *StakingManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingManager *StakingManagerSession) Receive() (*types.Transaction, error) {
	return _StakingManager.Contract.Receive(&_StakingManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingManager *StakingManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _StakingManager.Contract.Receive(&_StakingManager.TransactOpts)
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
	L2Strategy    common.Address
	Bridge        common.Address
	SourceChainId *big.Int
	DestChainId   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnstakeRequestClaimed is a free log retrieval operation binding the contract event 0xdefb23229f98ca9823b290c7a4ca8908d69e2a25a6fa9e0b65b3ffad35798eb3.
//
// Solidity: event UnstakeRequestClaimed(address indexed staker, address indexed l2Strategy, address indexed bridge, uint256 sourceChainId, uint256 destChainId)
func (_StakingManager *StakingManagerFilterer) FilterUnstakeRequestClaimed(opts *bind.FilterOpts, staker []common.Address, l2Strategy []common.Address, bridge []common.Address) (*StakingManagerUnstakeRequestClaimedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var l2StrategyRule []interface{}
	for _, l2StrategyItem := range l2Strategy {
		l2StrategyRule = append(l2StrategyRule, l2StrategyItem)
	}
	var bridgeRule []interface{}
	for _, bridgeItem := range bridge {
		bridgeRule = append(bridgeRule, bridgeItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnstakeRequestClaimed", stakerRule, l2StrategyRule, bridgeRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnstakeRequestClaimedIterator{contract: _StakingManager.contract, event: "UnstakeRequestClaimed", logs: logs, sub: sub}, nil
}

// WatchUnstakeRequestClaimed is a free log subscription operation binding the contract event 0xdefb23229f98ca9823b290c7a4ca8908d69e2a25a6fa9e0b65b3ffad35798eb3.
//
// Solidity: event UnstakeRequestClaimed(address indexed staker, address indexed l2Strategy, address indexed bridge, uint256 sourceChainId, uint256 destChainId)
func (_StakingManager *StakingManagerFilterer) WatchUnstakeRequestClaimed(opts *bind.WatchOpts, sink chan<- *StakingManagerUnstakeRequestClaimed, staker []common.Address, l2Strategy []common.Address, bridge []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var l2StrategyRule []interface{}
	for _, l2StrategyItem := range l2Strategy {
		l2StrategyRule = append(l2StrategyRule, l2StrategyItem)
	}
	var bridgeRule []interface{}
	for _, bridgeItem := range bridge {
		bridgeRule = append(bridgeRule, bridgeItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnstakeRequestClaimed", stakerRule, l2StrategyRule, bridgeRule)
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

// ParseUnstakeRequestClaimed is a log parse operation binding the contract event 0xdefb23229f98ca9823b290c7a4ca8908d69e2a25a6fa9e0b65b3ffad35798eb3.
//
// Solidity: event UnstakeRequestClaimed(address indexed staker, address indexed l2Strategy, address indexed bridge, uint256 sourceChainId, uint256 destChainId)
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
	Staker      common.Address
	L2Strategy  common.Address
	EthAmount   *big.Int
	DETHLocked  *big.Int
	DestChainId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnstakeRequested is a free log retrieval operation binding the contract event 0x7fd59a2c98b4d595be25d7976e588f7b6189ace6d41522f11868d951c2934367.
//
// Solidity: event UnstakeRequested(address indexed staker, address indexed l2Strategy, uint256 ethAmount, uint256 dETHLocked, uint256 destChainId)
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

// WatchUnstakeRequested is a free log subscription operation binding the contract event 0x7fd59a2c98b4d595be25d7976e588f7b6189ace6d41522f11868d951c2934367.
//
// Solidity: event UnstakeRequested(address indexed staker, address indexed l2Strategy, uint256 ethAmount, uint256 dETHLocked, uint256 destChainId)
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

// ParseUnstakeRequested is a log parse operation binding the contract event 0x7fd59a2c98b4d595be25d7976e588f7b6189ace6d41522f11868d951c2934367.
//
// Solidity: event UnstakeRequested(address indexed staker, address indexed l2Strategy, uint256 ethAmount, uint256 dETHLocked, uint256 destChainId)
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
