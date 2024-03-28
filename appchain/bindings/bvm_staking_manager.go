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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ALLOCATOR_SERVICE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INITIATOR_SERVICE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_ALLOWLIST_MANAGER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_ALLOWLIST_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_MANAGER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOP_UP_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocateETH\",\"inputs\":[{\"name\":\"allocateToUnstakeRequestsManager\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allocateToDeposits\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocatedETHForDeposits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchDethAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimUnstakeRequest\",\"inputs\":[{\"name\":\"unstakeRequestID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDETH\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dETHToETH\",\"inputs\":[{\"name\":\"dETHAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dapplinkBridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDepositContract\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethToDETH\",\"inputs\":[{\"name\":\"ethAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"exchangeAdjustmentRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializationBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"init\",\"type\":\"tuple\",\"internalType\":\"structStakingManager.Init\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"manager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocatorService\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initiatorService\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"returnsAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawalWallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dapplinkBridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dETH\",\"type\":\"address\",\"internalType\":\"contractIDETH\"},{\"name\":\"depositContract\",\"type\":\"address\",\"internalType\":\"contractIDepositContract\"},{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"contractIOracleReadRecord\"},{\"name\":\"pauser\",\"type\":\"address\",\"internalType\":\"contractIL1Pauser\"},{\"name\":\"unstakeRequestsManager\",\"type\":\"address\",\"internalType\":\"contractIUnstakeRequestsManager\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateValidatorsWithDeposits\",\"inputs\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structStakingManagerStorage.ValidatorParams[]\",\"components\":[{\"name\":\"operatorID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"withdrawalCredentials\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"expectedDepositRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isStakingAllowlist\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"laveDethAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maximumDETHSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maximumDepositAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumDepositAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumStakeBound\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumUnstakeBound\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numInitiatedValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOracleReadRecord\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauser\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIL1Pauser\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveFromUnstakeRequestsManager\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"receiveReturns\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"reclaimAllocatedETHSurplus\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"returnsAggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExchangeAdjustmentRate\",\"inputs\":[{\"name\":\"exchangeAdjustmentRate_\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaximumDETHSupply\",\"inputs\":[{\"name\":\"maximumDETHSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaximumDepositAmount\",\"inputs\":[{\"name\":\"maximumDepositAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumDepositAmount\",\"inputs\":[{\"name\":\"minimumDepositAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumStakeBound\",\"inputs\":[{\"name\":\"minimumStakeBound_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumUnstakeBound\",\"inputs\":[{\"name\":\"minimumUnstakeBound_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStakingAllowlist\",\"inputs\":[{\"name\":\"isStakingAllowlist_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWithdrawalWallet\",\"inputs\":[{\"name\":\"withdrawalWallet_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"topUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"totalControlled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalDepositedInValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unallocatedETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unstakeRequest\",\"inputs\":[{\"name\":\"dethAmount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"l2Strategy\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakeRequestInfo\",\"inputs\":[{\"name\":\"unstakeRequestID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unstakeRequestsManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIUnstakeRequestsManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usedValidators\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawalWallet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AllocatedETHToDeposits\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocatedETHToUnstakeRequestsManager\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtocolConfigChanged\",\"inputs\":[{\"name\":\"setterSelector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"setterSignature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReturnsReceived\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ethAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"dETHAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeBatchRequest\",\"inputs\":[{\"name\":\"batchId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"batchEthAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"batchDETHLocked\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeLaveAmount\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dETHLocked\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeRequestClaimed\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bridge\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeSingle\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dETHLocked\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"l2Strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorInitiated\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"operatorID\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"amountDeposited\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DoesNotReceiveETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfiguration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDepositRoot\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWithdrawalCredentialsNotETH1\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}]},{\"type\":\"error\",\"name\":\"InvalidWithdrawalCredentialsWrongAddress\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWithdrawalCredentialsWrongLength\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MathOverflowedMulDiv\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaximumDETHSupplyExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaximumValidatorDepositExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinimumDepositAmountNotSatisfied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinimumStakeBoundNotSatisfied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinimumUnstakeBoundNotSatisfied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinimumValidatorDepositNotSatisfied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotDappLinkBridge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughDepositETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughUnallocatedETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotReturnsAggregator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotUnstakeRequestsManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Paused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PreviouslyUsedValidator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnstakeBelowMinimudETHAmount\",\"inputs\":[{\"name\":\"ethAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedMinimum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001c62000022565b620000d6565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000735760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d35780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6139b980620000e66000396000f3fe60806040526004361061039a5760003560e01c806375796f76116101dc578063b12de58611610102578063d938e492116100a0578063ea452b6d1161006f578063ea452b6d14610aba578063ed9daafb14610ad0578063ee210d8714610af0578063f1ec1e9714610b10576103b8565b8063d938e49214610a3e578063dc29f1de14610a5e578063e55d6cc014610a66578063e94ad65b14610a9a576103b8565b8063c151aa72116100dc578063c151aa72146109d6578063ca15c873146109de578063d547741f146109fe578063d70a6f3114610a1e576103b8565b8063b12de58614610994578063b91590b2146109aa578063bb635c65146109c0576103b8565b80639010d07c1161017a578063a217fddf11610149578063a217fddf1461092c578063a694fc3a14610941578063aab483d614610954578063aaf5993114610974576103b8565b80639010d07c146108ac57806391d14854146108cc57806399dd1deb146108ec5780639fd0506d1461090c576103b8565b80637dfcdd29116101b65780637dfcdd291461083a578063808d663f14610850578063854a63f61461085857806389e80ed314610878576103b8565b806375796f76146107e457806378abb49b146108045780637dc0d1d01461081a576103b8565b806335ead2a4116102c15780634d014f7f1161025f57806360a0f6281161022e57806360a0f62814610764578063670186791461077a5780636daa01a2146107905780636fce8ab2146107b0576103b8565b80634d014f7f146106de57806354e5f761146106fe5780635915ded1146107145780635940d90b1461074f576103b8565b80633f550b3a1161029b5780633f550b3a1461066757806342d3915d1461068757806349336f0f146106a85780634a7d80b3146106be576103b8565b806335ead2a41461060f57806336568abe146106255780633937c0b314610645576103b8565b80631943190d11610339578063277f175811610308578063277f17581461056357806329d487041461059b5780632f2ff15d146105bb5780633101d910146105db576103b8565b80631943190d146104da57806319efd5c7146104ef5780631d2d35ce14610523578063248a9ca314610543576103b8565b806304f36cc21161037557806304f36cc2146104485780630633af7614610468578063080c279a14610496578063147d36d5146104ba576103b8565b80628db05b146103d157806301ffc9a7146103f35780630208e4b514610428576103b8565b366103b8576040516334352c7360e01b815260040160405180910390fd5b6040516334352c7360e01b815260040160405180910390fd5b3480156103dd57600080fd5b506103f16103ec366004612e61565b610b47565b005b3480156103ff57600080fd5b5061041361040e366004612e7a565b610bae565b60405190151581526020015b60405180910390f35b34801561043457600080fd5b506103f1610443366004612ea4565b610bd9565b34801561045457600080fd5b506103f1610463366004612f2d565b611031565b34801561047457600080fd5b50603b546104839061ffff1681565b60405161ffff909116815260200161041f565b3480156104a257600080fd5b506104ac603c5481565b60405190815260200161041f565b3480156104c657600080fd5b506104ac6104d5366004612e61565b6110ac565b3480156104e657600080fd5b506103f16111dd565b3480156104fb57600080fd5b506104ac7fe30bb2df90b65284acd0e8b5ebe3483bb2bbe65a08e43f0f9e8300fd8607ee1181565b34801561052f57600080fd5b506103f161053e366004612e61565b611260565b34801561054f57600080fd5b506104ac61055e366004612e61565b6112bc565b34801561056f57600080fd5b50603f54610583906001600160a01b031681565b6040516001600160a01b03909116815260200161041f565b3480156105a757600080fd5b506103f16105b6366004612f4a565b6112de565b3480156105c757600080fd5b506103f16105d6366004612f93565b611395565b3480156105e757600080fd5b506104ac7fe6ef7125bfa79685f3bd2e4c4cea243c1e988ebbc0801ab7641ae36b9e2c529181565b34801561061b57600080fd5b506104ac603a5481565b34801561063157600080fd5b506103f1610640366004612f93565b6113b1565b34801561065157600080fd5b506104ac60008051602061396483398151915281565b34801561067357600080fd5b50604454610583906001600160a01b031681565b34801561069357600080fd5b5060445461041390600160a01b900460ff1681565b3480156106b457600080fd5b506104ac60465481565b3480156106ca57600080fd5b50604354610583906001600160a01b031681565b3480156106ea57600080fd5b506103f16106f9366004613003565b6113e9565b34801561070a57600080fd5b506104ac60015481565b34801561072057600080fd5b5061041361072f3660046130ef565b805160208183018101805160348252928201919093012091525460ff1681565b34801561075b57600080fd5b506104ac611702565b34801561077057600080fd5b506104ac60355481565b34801561078657600080fd5b506104ac60005481565b34801561079c57600080fd5b506103f16107ab3660046131a0565b61186b565b3480156107bc57600080fd5b506104ac7f5e4bd437d29fad01c10cdcfff414f0d6b0e84b96d2dade88d780d45b5630696b81565b3480156107f057600080fd5b506103f16107ff3660046131c2565b611a6e565b34801561081057600080fd5b506104ac603d5481565b34801561082657600080fd5b50604054610583906001600160a01b031681565b34801561084657600080fd5b506104ac60375481565b6103f1611b19565b34801561086457600080fd5b50604254610583906001600160a01b031681565b34801561088457600080fd5b506104ac7fdec9d30de0821ad67aa5b141b13a539f584a19f99319e6041698a892b0e7955981565b3480156108b857600080fd5b506105836108c73660046131a0565b611b8f565b3480156108d857600080fd5b506104136108e7366004612f93565b611bbe565b3480156108f857600080fd5b506103f1610907366004612e61565b611bf6565b34801561091857600080fd5b50604154610583906001600160a01b031681565b34801561093857600080fd5b506104ac600081565b6103f161094f366004612e61565b611c52565b34801561096057600080fd5b506103f161096f366004612e61565b611e7c565b34801561098057600080fd5b506103f161098f3660046131f4565b611ed8565b3480156109a057600080fd5b506104ac60395481565b3480156109b657600080fd5b506104ac60455481565b3480156109cc57600080fd5b506104ac60365481565b6103f1611ee6565b3480156109ea57600080fd5b506104ac6109f9366004612e61565b611f23565b348015610a0a57600080fd5b506103f1610a19366004612f93565b611f49565b348015610a2a57600080fd5b506103f1610a39366004612e61565b611f65565b348015610a4a57600080fd5b50604754610583906001600160a01b031681565b6103f1611fc1565b348015610a7257600080fd5b506104ac7f8ea5b4dbd68db0bf23bf4cda958b61a749f8c5aec6f2912d75a03246753ddd1681565b348015610aa657600080fd5b50603e54610583906001600160a01b031681565b348015610ac657600080fd5b506104ac60385481565b348015610adc57600080fd5b506104ac610aeb366004612e61565b612005565b348015610afc57600080fd5b506103f1610b0b366004613222565b612101565b348015610b1c57600080fd5b50610b30610b2b366004612e61565b61229b565b60408051921515835260208301919091520161041f565b600080516020613964833981519152610b5f81612315565b60398290556040805160208101849052628db05b60e01b91600080516020613944833981519152910160408051601f1981840301815290829052610ba2916132ba565b60405180910390a25050565b60006001600160e01b03198216635a05180f60e01b1480610bd35750610bd382612322565b92915050565b7fe30bb2df90b65284acd0e8b5ebe3483bb2bbe65a08e43f0f9e8300fd8607ee11610c0381612315565b604160009054906101000a90046001600160a01b03166001600160a01b031663eea12ab66040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c56573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c7a9190613301565b15610c98576040516313d0ff5960e31b815260040160405180910390fd5b821561102b57603e546040805163c5f2892f60e01b815290516000926001600160a01b03169163c5f2892f9160048083019260209291908290030181865afa158015610ce8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d0c919061331e565b9050808314610d3657604051631497ae9360e01b8152600481018290526024015b60405180910390fd5b6000805b85811015610ee35736878783818110610d5557610d55613337565b9050602002810190610d67919061334d565b90506034610d78604083018361336d565b604051610d869291906133bb565b9081526040519081900360200190205460ff1615610db75760405163932c5b0d60e01b815260040160405180910390fd5b603c5481602001351015610dde576040516305cacc5560e21b815260040160405180910390fd5b603d5481602001351115610e0557604051630c25396f60e11b815260040160405180910390fd5b610e1a610e15606083018361336d565b612357565b60016034610e2b604084018461336d565b604051610e399291906133bb565b90815260405160209181900382019020805460ff191692151592909217909155610e6690820135846133e1565b92508035610e77604083018361336d565b604051610e859291906133bb565b60405180910390207f15f16c2e13e50235799a97b981bf4a66c8cd86051f06aca745c5ff26f39b330e838060400190610ebe919061336d565b8560200135604051610ed29392919061341d565b60405180910390a350600101610d3a565b50603854811115610f07576040516307bb2bd760e21b815260040160405180910390fd5b8060386000828254610f199190613441565b925050819055508060356000828254610f3291906133e1565b909155505060368054869190600090610f4c9084906133e1565b90915550600090505b858110156110275736878783818110610f7057610f70613337565b9050602002810190610f82919061334d565b603e549091506001600160a01b031663228951186020830135610fa8604085018561336d565b610fb5606087018761336d565b610fc2608089018961336d565b8960a001356040518963ffffffff1660e01b8152600401610fe99796959493929190613454565b6000604051808303818588803b15801561100257600080fd5b505af1158015611016573d6000803e3d6000fd5b505050505050806001019050610f55565b5050505b50505050565b60008051602061396483398151915261104981612315565b60448054831515600160a01b0260ff60a01b19909116179055604051630279b66160e11b9060008051602061394483398151915290611092908590602001901515815260200190565b60408051601f1981840301815290829052610ba2916134a5565b603f54604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd9160048083019260209291908290030181865afa1580156110f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061111a919061331e565b600003611125575090565b603b54610bd390839061113e9061ffff166127106134ec565b61ffff16603f60009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611195573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111b9919061331e565b6111c39190613507565b6127106111ce611702565b6111d89190613507565b612426565b6000805160206139648339815191526111f581612315565b604260009054906101000a90046001600160a01b03166001600160a01b031663b2d42b466040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561124557600080fd5b505af1158015611259573d6000803e3d6000fd5b5050505050565b60008051602061396483398151915261127881612315565b60468290556040805160208101849052630e969ae760e11b91600080516020613944833981519152910160408051601f1981840301815290829052610ba29161351e565b6000908152600080516020613924833981519152602052604090206001015490565b6000805160206139648339815191526112f681612315565b611303600a61271061357b565b61ffff168261ffff16111561132b5760405163c52a9bd360e01b815260040160405180910390fd5b61271061ffff831611156113415761134161359c565b603b805461ffff191661ffff8416908117909155604080516020810192909252630a7521c160e21b91600080516020613944833981519152910160408051601f1981840301815290829052610ba2916135b2565b61139e826112bc565b6113a781612315565b61102b83836124ea565b6001600160a01b03811633146113da5760405163334bd91960e11b815260040160405180910390fd5b6113e4828261252f565b505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff1660008115801561142f5750825b905060008267ffffffffffffffff16600114801561144c5750303b155b90508115801561145a575080155b156114785760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156114a257845460ff60401b1916600160401b1785555b6114aa61256b565b85516114b8906000906124ea565b506114d560008051602061396483398151915287602001516124ea565b506115047fe6ef7125bfa79685f3bd2e4c4cea243c1e988ebbc0801ab7641ae36b9e2c529187604001516124ea565b506115337fe30bb2df90b65284acd0e8b5ebe3483bb2bbe65a08e43f0f9e8300fd8607ee1187606001516124ea565b5061156c7f8ea5b4dbd68db0bf23bf4cda958b61a749f8c5aec6f2912d75a03246753ddd16600080516020613964833981519152612575565b6115b67fdec9d30de0821ad67aa5b141b13a539f584a19f99319e6041698a892b0e795597f8ea5b4dbd68db0bf23bf4cda958b61a749f8c5aec6f2912d75a03246753ddd16612575565b60e0860151603f80546001600160a01b03199081166001600160a01b0393841617909155610100880151603e805483169184169190911790556101208801516040805483169184169190911790556101408801516041805483169184169190911790556080880151604480546101608b015160428054861691871691909117905560c08b015160478054861691871691909117905560a08b0151604380549095169086161790935567016345785d8a0000603955662386f26fc10000603a556801bc16d674ec800000603c819055603d55600160a01b919093166001600160a81b03199092169190911717905543604555683782dace9d9000000060465583156116fa57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b600080604060009054906101000a90046001600160a01b03166001600160a01b031663f3a009a06040518163ffffffff1660e01b815260040161010060405180830381865afa158015611759573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061177d9190613626565b905060006037548161178f91906133e1565b90506038548161179f91906133e1565b90508160e001516001600160801b03166035546117bc9190613441565b6117c690826133e1565b90508160c001516001600160801b0316816117e191906133e1565b9050604260009054906101000a90046001600160a01b03166001600160a01b031663b69ef8a86040518163ffffffff1660e01b8152600401602060405180830381865afa158015611836573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061185a919061331e565b61186490826133e1565b9392505050565b7fe6ef7125bfa79685f3bd2e4c4cea243c1e988ebbc0801ab7641ae36b9e2c529161189581612315565b604160009054906101000a90046001600160a01b03166001600160a01b031663fdcada5e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061190c9190613301565b1561192a576040516313d0ff5960e31b815260040160405180910390fd5b60375461193783856133e1565b1115611956576040516396b0c75160e01b815260040160405180910390fd5b61196082846133e1565b603760008282546119719190613441565b909155505081156119c757816038600082825461198e91906133e1565b90915550506040518281527f9d04ecb465d2c8754acb798a22293dd26215a1c2f7a2a56607afa215c1aadc779060200160405180910390a15b82156113e4576040518381527ffe89805cf5299ef9fbd1d1ddefb8dcc3fa9408064d2ea31e3fca6565768f52179060200160405180910390a1604260009054906101000a90046001600160a01b03166001600160a01b0316634d13bfa6846040518263ffffffff1660e01b81526004016000604051808303818588803b158015611a5057600080fd5b505af1158015611a64573d6000803e3d6000fd5b5050505050505050565b600080516020613964833981519152611a8681612315565b816001600160a01b038116611aae5760405163d92e233d60e01b815260040160405180910390fd5b604380546001600160a01b0319166001600160a01b038516908117909155604080516020810192909252633abcb7bb60e11b91600080516020613944833981519152910160408051601f1981840301815290829052611b0c916136f0565b60405180910390a2505050565b6044546001600160a01b03163314611b4357604051626310df60e31b815260040160405180910390fd5b6040513481527f4cbb9d73b003a252cee3f2ee51d8d65a562af35eebb23730dd4a76d68127b3709060200160405180910390a13460376000828254611b8891906133e1565b9091555050565b6000828152600080516020613904833981519152602081905260408220611bb690846125d8565b949350505050565b6000918252600080516020613924833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b600080516020613964833981519152611c0e81612315565b603a82905560408051602081018490526399dd1deb60e01b91600080516020613944833981519152910160408051601f1981840301815290829052610ba291613737565b6047546001600160a01b03163314611c7d5760405163c68e02cf60e01b815260040160405180910390fd5b604160009054906101000a90046001600160a01b03166001600160a01b0316631ea7ca896040518163ffffffff1660e01b8152600401602060405180830381865afa158015611cd0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cf49190613301565b15611d12576040516313d0ff5960e31b815260040160405180910390fd5b604454600160a01b900460ff1615611d4d57611d4d7fdec9d30de0821ad67aa5b141b13a539f584a19f99319e6041698a892b0e79559612315565b603c54811015611d7057604051630a3287a960e21b815260040160405180910390fd5b6000611d7b826110ac565b9050604654603f60009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611dd3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611df7919061331e565b611e0190836133e1565b1115611e2057604051630935f98160e31b815260040160405180910390fd5b8160376000828254611e3291906133e1565b909155505060475460408051848152602081018490526001600160a01b03909216917f1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee909101610ba2565b600080516020613964833981519152611e9481612315565b603c829055604080516020810184905263555a41eb60e11b91600080516020613944833981519152910160408051601f1981840301815290829052610ba29161377e565b611ee282826125e4565b5050565b6042546001600160a01b03163314611f1157604051637154fc4360e01b815260040160405180910390fd5b3460376000828254611b8891906133e1565b60008181526000805160206139048339815191526020819052604082206118649061295d565b611f52826112bc565b611f5b81612315565b61102b838361252f565b600080516020613964833981519152611f7d81612315565b603d829055604080516020810184905263d70a6f3160e01b91600080516020613944833981519152910160408051601f1981840301815290829052610ba2916137c5565b7f5e4bd437d29fad01c10cdcfff414f0d6b0e84b96d2dade88d780d45b5630696b611feb81612315565b3460376000828254611ffd91906133e1565b909155505050565b603f54604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd9160048083019260209291908290030181865afa15801561204f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612073919061331e565b60000361207e575090565b610bd38261208a611702565b603f60009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156120dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111d8919061331e565b6047546001600160a01b0316331461212c5760405163c68e02cf60e01b815260040160405180910390fd5b604160009054906101000a90046001600160a01b03166001600160a01b0316638b61391a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561217f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121a39190613301565b156121c1576040516313d0ff5960e31b815260040160405180910390fd5b60408051848152602081018490526001600160a01b03861691339188917fefe8bd77c6737d081d3f8cc6b2fd7c02c4a8342d837993ea6201ed133c9d22b0910160405180910390a460425460405163019a076b60e61b8152600481018790523360248201526001600160a01b038681166044830152606482018690526084820185905260a4820184905290911690636681dac09060c4016020604051808303816000875af1158015612277573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116fa9190613301565b604254604051634f588bf160e01b81526004810183905260009182916001600160a01b0390911690634f588bf1906024016040805180830381865afa1580156122e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061230c919061380c565b91509150915091565b61231f8133612967565b50565b60006001600160e01b03198216637965db0b60e01b1480610bd357506301ffc9a760e01b6001600160e01b0319831614610bd3565b6020811461237b57604051639b0ec52760e01b815260048101829052602401610d2d565b600061238a600c82848661383a565b61239391613864565b9050600160f81b6001600160a01b03198216146123cf576040516308ebf56560e01b81526001600160a01b031982166004820152602401610d2d565b60006123df6020600c858761383a565b6123e891613894565b60435460609190911c91506001600160a01b0316811461102b57604051631b4d561960e01b81526001600160a01b0382166004820152602401610d2d565b600083830281600019858709828110838203039150508060000361245d5783828161245357612453613565565b0492505050611864565b80841161247d5760405163227bc15360e01b815260040160405180910390fd5b6000848688096000868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b60006000805160206139048339815191528161250685856129a0565b90508015611bb65760008581526020839052604090206125269085612a4c565b50949350505050565b60006000805160206139048339815191528161254b8585612a61565b90508015611bb65760008581526020839052604090206125269085612add565b612573612af2565b565b600080516020613924833981519152600061258f846112bc565b600085815260208490526040808220600101869055519192508491839187917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a450505050565b60006118648383612b3b565b604160009054906101000a90046001600160a01b03166001600160a01b0316638b61391a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612637573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061265b9190613301565b15612679576040516313d0ff5960e31b815260040160405180910390fd5b603a54826001600160801b031610156126a45760405162771ba760e71b815260040160405180910390fd5b6001541561272557603f546042546001546126cf926001600160a01b03908116923392911690612b65565b60015460405190815233907f1f4bdd7f902a9da109e5ee8424af475df1f64b7b7a8d36e07d5f97caaa5f19ec9060200160405180910390a260015460008082825461271a91906133e1565b909155505060006001555b816001600160801b031660008082825461273f91906133e1565b9091555050603d54600080549091612756916138c7565b11156128ec576000603d5460005461276e91906138c7565b90506000603d54826127809190613507565b9050806000546127909190613441565b600155603f546042546127b2916001600160a01b039081169133911684612b65565b604080518281526001600160a01b038516602082015233917f50fe46c173651ad8e673839ce6e2e0a8d2871d35c0637f796e7b915031afc76f910160405180910390a26000612802603d54612005565b604254604754600080546040516305165da360e41b81526001600160a01b03938416600482015260248101919091526001600160801b038516604482015293945092911690635165da30906064016020604051808303816000875af115801561286f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612893919061331e565b600054604080516001600160801b0386168152602081019290925291925082917f8b9951bb8391228a7668a21d43327f1814aa2d30534e5cf769cfdecc131de9ff910160405180910390a250506000805550611ee29050565b603f54604254612914916001600160a01b03908116913391166001600160801b038616612b65565b604080516001600160801b03841681526001600160a01b038316602082015233917f50fe46c173651ad8e673839ce6e2e0a8d2871d35c0637f796e7b915031afc76f9101610ba2565b6000610bd3825490565b6129718282611bbe565b611ee25760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610d2d565b60006000805160206139248339815191526129bb8484611bbe565b612a3b576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556129f13390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610bd3565b6000915050610bd3565b5092915050565b6000611864836001600160a01b038416612bbf565b6000600080516020613924833981519152612a7c8484611bbe565b15612a3b576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610bd3565b6000611864836001600160a01b038416612c0e565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661257357604051631afcd79f60e31b815260040160405180910390fd5b6000826000018281548110612b5257612b52613337565b9060005260206000200154905092915050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b17905261102b908590612cf7565b6000818152600183016020526040812054612c0657508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610bd3565b506000610bd3565b60008181526001830160205260408120548015612a3b576000612c32600183613441565b8554909150600090612c4690600190613441565b9050808214612cab576000866000018281548110612c6657612c66613337565b9060005260206000200154905080876000018481548110612c8957612c89613337565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080612cbc57612cbc6138db565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610bd3565b6000612d0c6001600160a01b03841683612d5a565b90508051600014158015612d31575080806020019051810190612d2f9190613301565b155b156113e457604051635274afe760e01b81526001600160a01b0384166004820152602401610d2d565b60606118648383600084600080856001600160a01b03168486604051612d8091906138f1565b60006040518083038185875af1925050503d8060008114612dbd576040519150601f19603f3d011682016040523d82523d6000602084013e612dc2565b606091505b5091509150612dd2868383612ddc565b9695505050505050565b606082612df157612dec82612e38565b611864565b8151158015612e0857506001600160a01b0384163b155b15612e3157604051639996b31560e01b81526001600160a01b0385166004820152602401610d2d565b5080611864565b805115612e485780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b600060208284031215612e7357600080fd5b5035919050565b600060208284031215612e8c57600080fd5b81356001600160e01b03198116811461186457600080fd5b600080600060408486031215612eb957600080fd5b833567ffffffffffffffff80821115612ed157600080fd5b818601915086601f830112612ee557600080fd5b813581811115612ef457600080fd5b8760208260051b8501011115612f0957600080fd5b6020928301989097509590910135949350505050565b801515811461231f57600080fd5b600060208284031215612f3f57600080fd5b813561186481612f1f565b600060208284031215612f5c57600080fd5b813561ffff8116811461186457600080fd5b6001600160a01b038116811461231f57600080fd5b8035612f8e81612f6e565b919050565b60008060408385031215612fa657600080fd5b823591506020830135612fb881612f6e565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b604051610180810167ffffffffffffffff81118282101715612ffd57612ffd612fc3565b60405290565b6000610180828403121561301657600080fd5b61301e612fd9565b61302783612f83565b815261303560208401612f83565b602082015261304660408401612f83565b604082015261305760608401612f83565b606082015261306860808401612f83565b608082015261307960a08401612f83565b60a082015261308a60c08401612f83565b60c082015261309b60e08401612f83565b60e08201526101006130ae818501612f83565b908201526101206130c0848201612f83565b908201526101406130d2848201612f83565b908201526101606130e4848201612f83565b908201529392505050565b60006020828403121561310157600080fd5b813567ffffffffffffffff8082111561311957600080fd5b818401915084601f83011261312d57600080fd5b81358181111561313f5761313f612fc3565b604051601f8201601f19908116603f0116810190838211818310171561316757613167612fc3565b8160405282815287602084870101111561318057600080fd5b826020860160208301376000928101602001929092525095945050505050565b600080604083850312156131b357600080fd5b50508035926020909101359150565b6000602082840312156131d457600080fd5b813561186481612f6e565b6001600160801b038116811461231f57600080fd5b6000806040838503121561320757600080fd5b8235613212816131df565b91506020830135612fb881612f6e565b600080600080600060a0868803121561323a57600080fd5b85359450602086013561324c81612f6e565b94979496505050506040830135926060810135926080909101359150565b60005b8381101561328557818101518382015260200161326d565b50506000910152565b600081518084526132a681602086016020860161326a565b601f01601f19169290920160200192915050565b60408152601d60408201527f7365744d696e696d756d5374616b65426f756e642875696e74323536290000006060820152608060208201526000611864608083018461328e565b60006020828403121561331357600080fd5b815161186481612f1f565b60006020828403121561333057600080fd5b5051919050565b634e487b7160e01b600052603260045260246000fd5b6000823560be1983360301811261336357600080fd5b9190910192915050565b6000808335601e1984360301811261338457600080fd5b83018035915067ffffffffffffffff82111561339f57600080fd5b6020019150368190038213156133b457600080fd5b9250929050565b8183823760009101908152919050565b634e487b7160e01b600052601160045260246000fd5b80820180821115610bd357610bd36133cb565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6040815260006134316040830185876133f4565b9050826020830152949350505050565b81810381811115610bd357610bd36133cb565b60808152600061346860808301898b6133f4565b828103602084015261347b81888a6133f4565b905082810360408401526134908186886133f4565b91505082606083015298975050505050505050565b60408152601960408201527f7365745374616b696e67416c6c6f776c69737428626f6f6c29000000000000006060820152608060208201526000611864608083018461328e565b61ffff828116828216039080821115612a4557612a456133cb565b8082028115828204841417610bd357610bd36133cb565b60408152601d60408201527f7365744d6178696d756d44455448537570706c792875696e74323536290000006060820152608060208201526000611864608083018461328e565b634e487b7160e01b600052601260045260246000fd5b600061ffff8084168061359057613590613565565b92169190910492915050565b634e487b7160e01b600052600160045260246000fd5b60408152602160408201527f73657445786368616e676541646a7573746d656e74526174652875696e7431366060820152602960f81b608082015260a06020820152600061186460a083018461328e565b805167ffffffffffffffff81168114612f8e57600080fd5b8051612f8e816131df565b600061010080838503121561363a57600080fd5b6040519081019067ffffffffffffffff8211818310171561365d5761365d612fc3565b8160405261366a84613603565b815261367860208501613603565b602082015261368960408501613603565b604082015261369a60608501613603565b6060820152608084015191506136af826131df565b8160808201526136c160a0850161361b565b60a08201526136d260c0850161361b565b60c08201526136e360e0850161361b565b60e0820152949350505050565b60408152601c60408201527f7365745769746864726177616c57616c6c6574286164647265737329000000006060820152608060208201526000611864608083018461328e565b60408152601f60408201527f7365744d696e696d756d556e7374616b65426f756e642875696e7432353629006060820152608060208201526000611864608083018461328e565b60408152602060408201527f7365744d696e696d756d4465706f736974416d6f756e742875696e74323536296060820152608060208201526000611864608083018461328e565b60408152602060408201527f7365744d6178696d756d4465706f736974416d6f756e742875696e74323536296060820152608060208201526000611864608083018461328e565b6000806040838503121561381f57600080fd5b825161382a81612f1f565b6020939093015192949293505050565b6000808585111561384a57600080fd5b8386111561385757600080fd5b5050820193919092039150565b6001600160a01b0319813581811691600c85101561388c57808186600c0360031b1b83161692505b505092915050565b6bffffffffffffffffffffffff19813581811691601485101561388c5760149490940360031b84901b1690921692915050565b6000826138d6576138d6613565565b500490565b634e487b7160e01b600052603160045260246000fd5b6000825161336381846020870161326a56fec1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680001d854e8dde9402801a4c6f2840193465752abfad61e0bb7c4258d526ae42e74a6b5d83d32632203555cb9b2c2f68a8d94da48cadd9266ac0d17babedb52ea5ba26469706673582212201f27aec4c0399721078ac82ac96581d2eb7ec081508954d89f438438d44b575d64736f6c63430008180033",
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

// BatchDethAmount is a free data retrieval call binding the contract method 0x67018679.
//
// Solidity: function batchDethAmount() view returns(uint256)
func (_StakingManager *StakingManagerCaller) BatchDethAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "batchDethAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchDethAmount is a free data retrieval call binding the contract method 0x67018679.
//
// Solidity: function batchDethAmount() view returns(uint256)
func (_StakingManager *StakingManagerSession) BatchDethAmount() (*big.Int, error) {
	return _StakingManager.Contract.BatchDethAmount(&_StakingManager.CallOpts)
}

// BatchDethAmount is a free data retrieval call binding the contract method 0x67018679.
//
// Solidity: function batchDethAmount() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) BatchDethAmount() (*big.Int, error) {
	return _StakingManager.Contract.BatchDethAmount(&_StakingManager.CallOpts)
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

// LaveDethAmount is a free data retrieval call binding the contract method 0x54e5f761.
//
// Solidity: function laveDethAmount() view returns(uint256)
func (_StakingManager *StakingManagerCaller) LaveDethAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "laveDethAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LaveDethAmount is a free data retrieval call binding the contract method 0x54e5f761.
//
// Solidity: function laveDethAmount() view returns(uint256)
func (_StakingManager *StakingManagerSession) LaveDethAmount() (*big.Int, error) {
	return _StakingManager.Contract.LaveDethAmount(&_StakingManager.CallOpts)
}

// LaveDethAmount is a free data retrieval call binding the contract method 0x54e5f761.
//
// Solidity: function laveDethAmount() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) LaveDethAmount() (*big.Int, error) {
	return _StakingManager.Contract.LaveDethAmount(&_StakingManager.CallOpts)
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

// UnstakeRequestInfo is a free data retrieval call binding the contract method 0xf1ec1e97.
//
// Solidity: function unstakeRequestInfo(uint256 unstakeRequestID) view returns(bool, uint256)
func (_StakingManager *StakingManagerCaller) UnstakeRequestInfo(opts *bind.CallOpts, unstakeRequestID *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "unstakeRequestInfo", unstakeRequestID)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// UnstakeRequestInfo is a free data retrieval call binding the contract method 0xf1ec1e97.
//
// Solidity: function unstakeRequestInfo(uint256 unstakeRequestID) view returns(bool, uint256)
func (_StakingManager *StakingManagerSession) UnstakeRequestInfo(unstakeRequestID *big.Int) (bool, *big.Int, error) {
	return _StakingManager.Contract.UnstakeRequestInfo(&_StakingManager.CallOpts, unstakeRequestID)
}

// UnstakeRequestInfo is a free data retrieval call binding the contract method 0xf1ec1e97.
//
// Solidity: function unstakeRequestInfo(uint256 unstakeRequestID) view returns(bool, uint256)
func (_StakingManager *StakingManagerCallerSession) UnstakeRequestInfo(unstakeRequestID *big.Int) (bool, *big.Int, error) {
	return _StakingManager.Contract.UnstakeRequestInfo(&_StakingManager.CallOpts, unstakeRequestID)
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

// ClaimUnstakeRequest is a paid mutator transaction binding the contract method 0xee210d87.
//
// Solidity: function claimUnstakeRequest(uint256 unstakeRequestID, address bridge, uint256 sourceChainId, uint256 destChainId, uint256 gasLimit) returns()
func (_StakingManager *StakingManagerTransactor) ClaimUnstakeRequest(opts *bind.TransactOpts, unstakeRequestID *big.Int, bridge common.Address, sourceChainId *big.Int, destChainId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "claimUnstakeRequest", unstakeRequestID, bridge, sourceChainId, destChainId, gasLimit)
}

// ClaimUnstakeRequest is a paid mutator transaction binding the contract method 0xee210d87.
//
// Solidity: function claimUnstakeRequest(uint256 unstakeRequestID, address bridge, uint256 sourceChainId, uint256 destChainId, uint256 gasLimit) returns()
func (_StakingManager *StakingManagerSession) ClaimUnstakeRequest(unstakeRequestID *big.Int, bridge common.Address, sourceChainId *big.Int, destChainId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.ClaimUnstakeRequest(&_StakingManager.TransactOpts, unstakeRequestID, bridge, sourceChainId, destChainId, gasLimit)
}

// ClaimUnstakeRequest is a paid mutator transaction binding the contract method 0xee210d87.
//
// Solidity: function claimUnstakeRequest(uint256 unstakeRequestID, address bridge, uint256 sourceChainId, uint256 destChainId, uint256 gasLimit) returns()
func (_StakingManager *StakingManagerTransactorSession) ClaimUnstakeRequest(unstakeRequestID *big.Int, bridge common.Address, sourceChainId *big.Int, destChainId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.ClaimUnstakeRequest(&_StakingManager.TransactOpts, unstakeRequestID, bridge, sourceChainId, destChainId, gasLimit)
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

// UnstakeRequest is a paid mutator transaction binding the contract method 0xaaf59931.
//
// Solidity: function unstakeRequest(uint128 dethAmount, address l2Strategy) returns()
func (_StakingManager *StakingManagerTransactor) UnstakeRequest(opts *bind.TransactOpts, dethAmount *big.Int, l2Strategy common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "unstakeRequest", dethAmount, l2Strategy)
}

// UnstakeRequest is a paid mutator transaction binding the contract method 0xaaf59931.
//
// Solidity: function unstakeRequest(uint128 dethAmount, address l2Strategy) returns()
func (_StakingManager *StakingManagerSession) UnstakeRequest(dethAmount *big.Int, l2Strategy common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.UnstakeRequest(&_StakingManager.TransactOpts, dethAmount, l2Strategy)
}

// UnstakeRequest is a paid mutator transaction binding the contract method 0xaaf59931.
//
// Solidity: function unstakeRequest(uint128 dethAmount, address l2Strategy) returns()
func (_StakingManager *StakingManagerTransactorSession) UnstakeRequest(dethAmount *big.Int, l2Strategy common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.UnstakeRequest(&_StakingManager.TransactOpts, dethAmount, l2Strategy)
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

// StakingManagerUnstakeBatchRequestIterator is returned from FilterUnstakeBatchRequest and is used to iterate over the raw logs and unpacked data for UnstakeBatchRequest events raised by the StakingManager contract.
type StakingManagerUnstakeBatchRequestIterator struct {
	Event *StakingManagerUnstakeBatchRequest // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnstakeBatchRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnstakeBatchRequest)
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
		it.Event = new(StakingManagerUnstakeBatchRequest)
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
func (it *StakingManagerUnstakeBatchRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnstakeBatchRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnstakeBatchRequest represents a UnstakeBatchRequest event raised by the StakingManager contract.
type StakingManagerUnstakeBatchRequest struct {
	BatchId         *big.Int
	BatchEthAmount  *big.Int
	BatchDETHLocked *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnstakeBatchRequest is a free log retrieval operation binding the contract event 0x8b9951bb8391228a7668a21d43327f1814aa2d30534e5cf769cfdecc131de9ff.
//
// Solidity: event UnstakeBatchRequest(uint256 indexed batchId, uint256 batchEthAmount, uint256 batchDETHLocked)
func (_StakingManager *StakingManagerFilterer) FilterUnstakeBatchRequest(opts *bind.FilterOpts, batchId []*big.Int) (*StakingManagerUnstakeBatchRequestIterator, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnstakeBatchRequest", batchIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnstakeBatchRequestIterator{contract: _StakingManager.contract, event: "UnstakeBatchRequest", logs: logs, sub: sub}, nil
}

// WatchUnstakeBatchRequest is a free log subscription operation binding the contract event 0x8b9951bb8391228a7668a21d43327f1814aa2d30534e5cf769cfdecc131de9ff.
//
// Solidity: event UnstakeBatchRequest(uint256 indexed batchId, uint256 batchEthAmount, uint256 batchDETHLocked)
func (_StakingManager *StakingManagerFilterer) WatchUnstakeBatchRequest(opts *bind.WatchOpts, sink chan<- *StakingManagerUnstakeBatchRequest, batchId []*big.Int) (event.Subscription, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnstakeBatchRequest", batchIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnstakeBatchRequest)
				if err := _StakingManager.contract.UnpackLog(event, "UnstakeBatchRequest", log); err != nil {
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

// ParseUnstakeBatchRequest is a log parse operation binding the contract event 0x8b9951bb8391228a7668a21d43327f1814aa2d30534e5cf769cfdecc131de9ff.
//
// Solidity: event UnstakeBatchRequest(uint256 indexed batchId, uint256 batchEthAmount, uint256 batchDETHLocked)
func (_StakingManager *StakingManagerFilterer) ParseUnstakeBatchRequest(log types.Log) (*StakingManagerUnstakeBatchRequest, error) {
	event := new(StakingManagerUnstakeBatchRequest)
	if err := _StakingManager.contract.UnpackLog(event, "UnstakeBatchRequest", log); err != nil {
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
	Id            *big.Int
	Staker        common.Address
	Bridge        common.Address
	SourceChainId *big.Int
	DestChainId   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnstakeRequestClaimed is a free log retrieval operation binding the contract event 0xefe8bd77c6737d081d3f8cc6b2fd7c02c4a8342d837993ea6201ed133c9d22b0.
//
// Solidity: event UnstakeRequestClaimed(uint256 indexed id, address indexed staker, address indexed bridge, uint256 sourceChainId, uint256 destChainId)
func (_StakingManager *StakingManagerFilterer) FilterUnstakeRequestClaimed(opts *bind.FilterOpts, id []*big.Int, staker []common.Address, bridge []common.Address) (*StakingManagerUnstakeRequestClaimedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var bridgeRule []interface{}
	for _, bridgeItem := range bridge {
		bridgeRule = append(bridgeRule, bridgeItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnstakeRequestClaimed", idRule, stakerRule, bridgeRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnstakeRequestClaimedIterator{contract: _StakingManager.contract, event: "UnstakeRequestClaimed", logs: logs, sub: sub}, nil
}

// WatchUnstakeRequestClaimed is a free log subscription operation binding the contract event 0xefe8bd77c6737d081d3f8cc6b2fd7c02c4a8342d837993ea6201ed133c9d22b0.
//
// Solidity: event UnstakeRequestClaimed(uint256 indexed id, address indexed staker, address indexed bridge, uint256 sourceChainId, uint256 destChainId)
func (_StakingManager *StakingManagerFilterer) WatchUnstakeRequestClaimed(opts *bind.WatchOpts, sink chan<- *StakingManagerUnstakeRequestClaimed, id []*big.Int, staker []common.Address, bridge []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var bridgeRule []interface{}
	for _, bridgeItem := range bridge {
		bridgeRule = append(bridgeRule, bridgeItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnstakeRequestClaimed", idRule, stakerRule, bridgeRule)
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

// ParseUnstakeRequestClaimed is a log parse operation binding the contract event 0xefe8bd77c6737d081d3f8cc6b2fd7c02c4a8342d837993ea6201ed133c9d22b0.
//
// Solidity: event UnstakeRequestClaimed(uint256 indexed id, address indexed staker, address indexed bridge, uint256 sourceChainId, uint256 destChainId)
func (_StakingManager *StakingManagerFilterer) ParseUnstakeRequestClaimed(log types.Log) (*StakingManagerUnstakeRequestClaimed, error) {
	event := new(StakingManagerUnstakeRequestClaimed)
	if err := _StakingManager.contract.UnpackLog(event, "UnstakeRequestClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUnstakeSingleIterator is returned from FilterUnstakeSingle and is used to iterate over the raw logs and unpacked data for UnstakeSingle events raised by the StakingManager contract.
type StakingManagerUnstakeSingleIterator struct {
	Event *StakingManagerUnstakeSingle // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnstakeSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnstakeSingle)
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
		it.Event = new(StakingManagerUnstakeSingle)
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
func (it *StakingManagerUnstakeSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnstakeSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnstakeSingle represents a UnstakeSingle event raised by the StakingManager contract.
type StakingManagerUnstakeSingle struct {
	Staker     common.Address
	DETHLocked *big.Int
	L2Strategy common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnstakeSingle is a free log retrieval operation binding the contract event 0x50fe46c173651ad8e673839ce6e2e0a8d2871d35c0637f796e7b915031afc76f.
//
// Solidity: event UnstakeSingle(address indexed staker, uint256 dETHLocked, address l2Strategy)
func (_StakingManager *StakingManagerFilterer) FilterUnstakeSingle(opts *bind.FilterOpts, staker []common.Address) (*StakingManagerUnstakeSingleIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnstakeSingle", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnstakeSingleIterator{contract: _StakingManager.contract, event: "UnstakeSingle", logs: logs, sub: sub}, nil
}

// WatchUnstakeSingle is a free log subscription operation binding the contract event 0x50fe46c173651ad8e673839ce6e2e0a8d2871d35c0637f796e7b915031afc76f.
//
// Solidity: event UnstakeSingle(address indexed staker, uint256 dETHLocked, address l2Strategy)
func (_StakingManager *StakingManagerFilterer) WatchUnstakeSingle(opts *bind.WatchOpts, sink chan<- *StakingManagerUnstakeSingle, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnstakeSingle", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnstakeSingle)
				if err := _StakingManager.contract.UnpackLog(event, "UnstakeSingle", log); err != nil {
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

// ParseUnstakeSingle is a log parse operation binding the contract event 0x50fe46c173651ad8e673839ce6e2e0a8d2871d35c0637f796e7b915031afc76f.
//
// Solidity: event UnstakeSingle(address indexed staker, uint256 dETHLocked, address l2Strategy)
func (_StakingManager *StakingManagerFilterer) ParseUnstakeSingle(log types.Log) (*StakingManagerUnstakeSingle, error) {
	event := new(StakingManagerUnstakeSingle)
	if err := _StakingManager.contract.UnpackLog(event, "UnstakeSingle", log); err != nil {
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
