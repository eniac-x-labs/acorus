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

// IL1PoolManagerKeyValuePair is an auto generated low-level Go binding around an user-defined struct.
type IL1PoolManagerKeyValuePair struct {
	Key   common.Address
	Value *big.Int
}

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

// IL1PoolManagerUser is an auto generated low-level Go binding around an user-defined struct.
type IL1PoolManagerUser struct {
	IsWithdrawed bool
	Token        common.Address
	StartPoolId  *big.Int
	EndPoolId    *big.Int
	Amount       *big.Int
}

// L1PoolManagerMetaData contains all meta data concerning the L1PoolManager contract.
var L1PoolManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ChainIdIsNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"ChainIdNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorBlockChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"providedAmount\",\"type\":\"uint256\"}],\"name\":\"LessThanMinStakeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"MinTransferAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LessThanMinTransferAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LessThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MantaNotWETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MantleNotWETH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PoolIndex\",\"type\":\"uint256\"}],\"name\":\"NewPoolIsNotCreate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReward\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughETH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"NotEnoughToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"PoolLength\",\"type\":\"uint256\"}],\"name\":\"OutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolIndex\",\"type\":\"uint256\"}],\"name\":\"PoolIsCompleted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"name\":\"TokenIsAlreadySupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"TokenIsNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"sourceChainIdError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIDETH.BatchMint[]\",\"name\":\"batcher\",\"type\":\"tuple[]\"}],\"name\":\"BridgeFinalizeETHForStakingEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endPoolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Reward\",\"type\":\"uint256\"}],\"name\":\"ClaimReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolIndex\",\"type\":\"uint256\"}],\"name\":\"CompletePoolEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"shareAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeMessageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stakeMessageHash\",\"type\":\"bytes32\"}],\"name\":\"FinalizeStakingMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeMessageNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"stakeMessageHash\",\"type\":\"bytes32\"}],\"name\":\"InitiateStakingMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetMinStakeAmountEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSupport\",\"type\":\"bool\"}],\"name\":\"SetSupportTokenEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingETHEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingWETHEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StarkingERC20Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Blockchain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferAssertTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Reward\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIDETH.BatchMint[]\",\"name\":\"batcher\",\"type\":\"tuple[]\"}],\"name\":\"BridgeFinalizeETHForStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"shareAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeStakingMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"BridgeInitiateETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateStakingMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ClaimAllReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"ClaimbyID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"IsCompleted\",\"type\":\"bool\"}],\"internalType\":\"structIL1PoolManager.Pool[]\",\"name\":\"CompletePools\",\"type\":\"tuple[]\"}],\"name\":\"CompletePoolAndNew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DepositAndStaking\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DepositAndStakingERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DepositAndStakingETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositAndStakingWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FeePoolValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FundingPoolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"IsSupportChainId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"IsSupportToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MinStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinTransferAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Pools\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"IsCompleted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"QuickSendAssertToUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ReLayer\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SupportTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"Blockchain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TransferAssertToBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UpdateFundingPoolBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Users\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isWithdrawed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"StartPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WithdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"WithdrawByID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getPool\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"IsCompleted\",\"type\":\"bool\"}],\"internalType\":\"structIL1PoolManager.Pool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getPoolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrincipal\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structIL1PoolManager.KeyValuePair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structIL1PoolManager.KeyValuePair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isWithdrawed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"StartPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"}],\"internalType\":\"structIL1PoolManager.User[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isWithdrawed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"StartPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"}],\"internalType\":\"structIL1PoolManager.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MultisigWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageManager\",\"outputs\":[{\"internalType\":\"contractIMessageManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodTime\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setMinStakeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_MinTransferAmount\",\"type\":\"uint256\"}],\"name\":\"setMinTransferAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_PerFee\",\"type\":\"uint256\"}],\"name\":\"setPerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setSupportERC20Token\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isSupport\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"startTimes\",\"type\":\"uint32\"}],\"name\":\"setSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setValidChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingMessageNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60808060405234620000bd577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c16620000ae57506001600160401b036002600160401b03198282160162000068575b604051616e1d9081620000c38239f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a138808062000058565b63f92ee8a960e01b8152600490fd5b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a7146103d757806309dc480c146103d25780630cd727df146103cd57806310875a13146103c857806313e8544e146103c35780631ca2f173146103be5780631d31fac0146103b9578063248a9ca3146103b45780632c37388b146103af5780632f2ff15d146103aa57806331d98b3b146103a5578063325e5d43146103a05780633338562c1461039b57806334d34ef51461039657806336568abe1461039157806338731c471461038c5780633b0230f0146103875780633d18b912146103825780633f4ba83a1461037d578063415ade1f146103785780634663cdc814610373578063485cc9551461036e5780634dfaef841461036957806354dc027e146103645780635765a4eb1461035f5780635a648bc51461035a5780635b5b9ea2146103555780635c975abb14610350578063626417b51461034b578063650c2276146103465780636f77926b1461034157806372fe6a7e1461033c5780638456cb591461033757806391d148541461033257806392d847381461032d57806396f984e6146103285780639ab0b652146103235780639b4423801461031e578063a217fddf14610319578063a69bdf1614610314578063ab0f9e191461030f578063b1887e9b1461030a578063b92e639614610305578063bc42493d14610300578063cb314fab146102fb578063cb4f04ad146102f6578063cd01c665146102f1578063d547741f146102ec578063d73f14e5146102e7578063dac29568146102e2578063dbb0481f146102dd578063dd0c3460146102d8578063e1c9e844146102d3578063e2070893146102ce578063f363e52d146102c9578063f62a89cf146102c4578063f8fcc2aa146102bf578063f91fa9a8146102ba578063fa861848146102b55763ff2bf64f146102b057600080fd5b61430c565b6142e9565b614030565b613ff6565b613c82565b613c1e565b613a0a565b613959565b61393b565b613901565b6138c2565b613879565b61382d565b613535565b6134fb565b613490565b6131a7565b613183565b613156565b6130f3565b612fad565b612f91565b612dac565b612c91565b6129f6565b61281c565b6127bd565b612756565b612725565b61269b565b612600565b6125b7565b612562565b612420565b612050565b611e86565b611e4c565b611bf1565b611abb565b611a04565b61186d565b6117f6565b6115b6565b611265565b6110a3565b611026565b610d38565b610c45565b610b0f565b610abc565b610a6e565b6106a9565b61066d565b610649565b61060f565b61056e565b61050e565b6104c7565b61043d565b3461042d57602036600319011261042d5760043563ffffffff60e01b811680910361042d57602090637965db0b60e01b811490811561041c575b506040519015158152f35b6301ffc9a760e01b14905038610411565b600080fd5b600091031261042d57565b3461042d57600036600319011261042d576020600154604051908152f35b600435906001600160a01b038216820361042d57565b602435906001600160a01b038216820361042d57565b604435906001600160a01b038216820361042d57565b606435906001600160a01b038216820361042d57565b35906001600160a01b038216820361042d57565b3461042d57604036600319011261042d576104e061045b565b6104f06104eb6142a7565b614410565b6001600160a01b031660009081526005602052604090206024359055005b3461042d57600036600319011261042d576000546040516001600160a01b039091168152602090f35b634e487b7160e01b600052603260045260246000fd5b80548210156105695760005260206000209060021b0190600090565b610537565b3461042d57604036600319011261042d5761058761045b565b6024359060018060a01b03809116600052600b602052604060002091825481101561042d576105bb61060b9160ff9461054d565b50805490600181015460036002830154920154926040519687968260081c169116869192608093969594919660a08401971515845260018060a01b03166020840152604083015260608201520152565b0390f35b3461042d57602036600319011261042d576001600160a01b0361063061045b565b16600052600c6020526020604060002054604051908152f35b3461042d57600036600319011261042d57602063ffffffff60095416604051908152f35b3461042d57602036600319011261042d57600435600052600080516020616d688339815191526020526020600160406000200154604051908152f35b3461042d57600080600319360112610a6b576106c361480e565b6106cb61472b565b805b60078054821015610a55578252600080516020616ce88339815191528101546001600160a01b039081166000818152600a602052604090209091905415610a3c57835b336000908152600b6020526040902054811015610a2b5782826107626107518461074c3360018060a01b0316600052600b602052604060002090565b61054d565b505460081c6001600160a01b031690565b1614610777575b61077290614ec5565b610710565b336000908152600b6020526040902061079d9061079590839061054d565b505460ff1690565b610769576001600160a01b0383166000908152600a602052604090206107c4905b54614286565b336000908152600b602052604090206107de90839061054d565b506003908101546001600160a01b0386166000908152600a6020526040902061080d9084906119e8565b6119e8565b509061081f600180930191825461429a565b9055336000908152600b60205260409020889190839061084090879061054d565b500154336000908152600b60205260409020829061085f90889061054d565b50015493858511610a1957845b86811061092157505082857f7e6fedb6f824b2112ac0626ced32f6be6324d554cc03efd112aa021ee4d740ac95936108b783610917956108b26107729c9b991515614ddb565b614801565b50336000908152600b602052604090206108d290899061054d565b5001556108e081338b616b2f565b506040805133815260208101949094528301939093526001600160a01b0388166060830152608082019290925290819060a0820190565b0390a19050610769565b909192936109446107be8b60018060a01b0316600052600a602052604060002090565b82116109fc57816109f7916109f0856109e6848f9a9998976109ca6109c38d8d6109ba858f61099e6108089160026109956108089b6108088a60018060a01b0316600052600a602052604060002090565b50015490614da8565b6001600160a01b039095166000908152600a6020526040902090565b50015490614dbb565b8097614801565b6001600160a01b03909c166000908152600a6020526040902090565b5001918254614801565b9055614d99565b61086c565b604051637d58ebb960e01b815260048101839052602490fd5b0390fd5b60405163374c934360e11b8152600490fd5b505050610a3790614d99565b6106cd565b604051637d58ebb960e01b815260006004820152602490fd5b826001600080516020616da88339815191525580f35b80fd5b3461042d57604036600319011261042d57610aba600435610a8d610471565b9080600052600080516020616d68833981519152602052610ab5600160406000200154614410565b614515565b005b3461042d57600036600319011261042d576020600854604051908152f35b606090600319011261042d576001600160a01b0390600435828116810361042d5791602435908116810361042d579060443590565b3461042d57610aba610b2036610ada565b91610b2c6104eb6142a7565b616b2f565b634e487b7160e01b600052604160045260246000fd5b60e0810190811067ffffffffffffffff821117610b6357604052565b610b31565b67ffffffffffffffff8111610b6357604052565b60a0810190811067ffffffffffffffff821117610b6357604052565b6040810190811067ffffffffffffffff821117610b6357604052565b90601f8019910116810190811067ffffffffffffffff821117610b6357604052565b60405190610be382610b98565b565b60405190610be382610b7c565b60405190610be382610b47565b67ffffffffffffffff8111610b635760051b60200190565b6044359063ffffffff8216820361042d57565b359063ffffffff8216820361042d57565b8015150361042d57565b60208060031936011261042d576004359067ffffffffffffffff821161042d573660238301121561042d57816004013591610c7f83610bff565b916040610c8e81519485610bb4565b84845281840190602460e08097028501019336851161042d57602401915b848310610cbc57610aba86614eeb565b868336031261042d578387918351610cd381610b47565b610cdc86610c2a565b8152610ce9838701610c2a565b83820152610cf88587016104b3565b85820152606080870135908201526080808701359082015260a0808701359082015260c08087013590610d2a82610c3b565b820152815201920191610cac565b3461042d57602036600319011261042d5760048035610d5561480e565b610d5d61472b565b33600052600b6020526040918260002054808310156110055750336000908152600b60205260409020610d959061075190849061054d565b90610db56107be8360018060a01b0316600052600a602052604060002090565b336000908152600b60205260409020909290610dd290859061054d565b509460038096015494610dfb856108088560018060a01b0316600052600a602052604060002090565b5095610e0d600180980191825461429a565b9055336000908152600b602052604081209097908190610e2e90849061054d565b5001549087610e538461074c3360018060a01b0316600052600b602052604060002090565b50015495878711610ff757865b888110610f4357505050918587610ee9610f2e9694610ea88c80986108b27f7e6fedb6f824b2112ac0626ced32f6be6324d554cc03efd112aa021ee4d740ac9e9f1515614ddb565b50336000908152600b60205260409020610ed590610ec790839061054d565b50805460ff19166001179055565b336000908152600b6020526040902061054d565b500155610ef7833384616b2f565b5051338152602081019490945260408401949094526001600160a01b039093166060830152608082019290925290819060a0820190565b0390a1600080516020616da883398151915255005b6001600160a01b0387166000908152600a60205260409020610f64906107be565b8111610fdc5780610fb2610fab898d6109ba8561080861099e6002610fa2610fd79b6108088960018060a01b0316600052600a602052604060002090565b5001548d614da8565b809d614801565b9b6109f0856109e6846108088d60018060a01b0316600052600a602052604060002090565b610e60565b8551637d58ebb960e01b815291820190815281906020010390fd5b845163374c934360e11b8152fd5b835163abe5c32f60e01b815291820192835260208301529081906040010390fd5b3461042d57604036600319011261042d5761103f610471565b336001600160a01b0382160361105b57610aba906004356145b6565b60405163334bd91960e11b8152600490fd5b60c090600319011261042d5760043590602435906044356001600160a01b038116810361042d5790606435906084359060a43590565b6110ac3661106d565b6110bc6104eb96959294966142a7565b468203611253576110e86110e46110dd876000526003602052604060002090565b5460ff1690565b1590565b61123a576001600160a01b03861695600084888115611230575b600092839283928392f11561122b57600080516020616d488339815191526000526005602052600080516020616c8883398151915261114285825461429a565b9055600054611161906001600160a01b03165b6001600160a01b031690565b803b1561042d57604051631eb65ffb60e01b815260048101889052602481018590526001600160a01b0392909216604483015260648201959095526084810184905260a481019190915292600090849060c490829084905af192831561122b577f61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f3293611212575b50604080519485526020850191909152830152309180606081015b0390a360405160018152602090f35b8061121f61122592610b68565b80610432565b386111e8565b614d8d565b6108fc9250611102565b604051632ef6974160e11b815260048101869052602490fd5b604051631a26aa4d60e21b8152600490fd5b3461042d5760a036600319011261042d576004803560243591611286610487565b61128e61049d565b608435468503611552576112b26110e46110dd886000526003602052604060002090565b611536576001600160a01b03821660009081526004602052604090206112db906110e4906110dd565b61150e576040516370a0823160e01b808252308287019081526001600160a01b03858116979095939092602092839083908190830103818c5afa91821561122b576000926114ef575b506113318630338c61483e565b6040519081523084820190815283908290819060200103818c5afa90811561122b57611366936000926114c2575b505061429a565b9261138e6113868460018060a01b03166000526005602052604060002090565b918254614801565b90556113d16113b56113ae6113a560025487614da8565b620f4240900490565b809561429a565b6001600160a01b03909316600090815260066020526040902090565b6113dc848254614801565b90556000546113f3906001600160a01b0316611155565b92833b1561042d576040805163305f478560e21b8152928301898152602081018b90526001600160a01b0388169181019190915260608101849052608081019190915290949260009186919082908490829060a00103925af190811561122b577fece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a4946114a0926114af575b5060405193849316973397846040919493926060820195825260208201520152565b0390a460405160018152602090f35b8061121f6114bc92610b68565b3861147e565b6114e19250803d106114e8575b6114d98183610bb4565b8101906158cf565b388061135f565b503d6114cf565b611507919250833d85116114e8576114d98183610bb4565b9038611324565b506040516305fd61ad60e01b81526001600160a01b0390911681840190815281906020010390fd5b604051632ef6974160e11b815280850187815281906020010390fd5b604051631a26aa4d60e21b81528490fd5b60208082019080835283518092528060408094019401926000905b83821061158d57505050505090565b845180516001600160a01b0316875283015186840152948501949382019360019091019061157e565b3461042d57600080600319360112610a6b576007546115d481616927565b9082905b8082106115ed576040518061060b8582611563565b83805b336000908152600b60205260409020548110156117a157336000908152600b602052604090206116259061075190839061054d565b61164961115561163487612592565b905460039190911b1c6001600160a01b031690565b6001600160a01b0390911614611668575b61166390614d99565b6115f0565b336000908152600b602052604090206116869061079590839061054d565b61165a576116b66107be61169c61163487612592565b6001600160a01b03166000908152600a6020526040902090565b336000908152600b602052604090206003906116d390849061054d565b500154336000908152600b602052604090206116f090849061054d565b5091600180930154818111611794579291905b808410611713575050505061165a565b909192986117296107be61169c6116348b612592565b8a1161177b5761176d611773916117678a876109ba8f61080861169c6116346117618d60026109958761080861169c6116348e612592565b96612592565b90614801565b99614d99565b929190611703565b604051637d58ebb960e01b8152600481018b9052602490fd5b5050505061166390614d99565b5093916117f091926117b561163483612592565b906117d06117c1610bd6565b6001600160a01b039093168352565b60208201526117df8286615115565b526117ea8185615115565b50614d99565b906115d8565b3461042d57600036600319011261042d5761180f6143b8565b600080516020616d88833981519152805460ff81161561185b5760ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b604051638dfc202b60e01b8152600490fd5b3461042d5760c036600319011261042d5761188661045b565b61188e610471565b90611897610487565b60643560843591846040519360208501856118d883878787869092606092959493608083019660018060a01b03809216845216602083015260408201520152565b03956118ec601f1997888101835282610bb4565b519020604051635bf55fe360e01b602082019081526001600160a01b0394851660248301529385166044820152606481018690526084810183905260a496870181529590919061193c9087610bb4565b622673c060a43560061b01603f5a02106119b957611987600080611203957fac6725dae59a3a2ceb3b61453d910067b411ec59d487e1d72382b9cdb9ea723e995190828c5af16169d0565b604080516001600160a01b03988916815260208101969096528501526060840152908416949093169281906080820190565b6308c379a06000526020805278185361666543616c6c3a204e6f7420656e6f756768206761736058526064601cfd5b8054821015610569576000526005602060002091020190600090565b3461042d57604036600319011261042d57611a1d61045b565b6001600160a01b039081166000908152600a6020526040902080546024359081101561042d57611a4c916119e8565b5080546001820154600283015460038401546004909401546040805163ffffffff8087168252602087811c9091169082015294811c969096166001600160a01b0316958401959095526060830191909152608082015260a081019190915260ff909116151560c082015260e090f35b3461042d57604036600319011261042d57611ad461045b565b611adc610471565b90600080516020616dc8833981519152549167ffffffffffffffff60ff8460401c1615931680159081611be9575b6001149081611bdf575b159081611bd6575b50611bc457600080516020616dc8833981519152805467ffffffffffffffff19166001179055611b509183611b9f5761464e565b611b5657005b600080516020616dc8833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b600080516020616dc8833981519152805460ff60401b1916600160401b17905561464e565b60405163f92ee8a960e01b8152600490fd5b90501538611b1c565b303b159150611b14565b849150611b0a565b3461042d5760408060031936011261042d57611c0b61045b565b9060243590611c1861480e565b611c2061472b565b6001600160a01b0383166000908152600460205260409020611c45906110e4906110dd565b611e2d576001600160a01b0383166000908152600c60205260409020548210611de9576001600160a01b03831692611c7f8330338761483e565b6001600160a01b0381166000908152600a60205260409020611ca0906107be565b611cce611cc3826108088560018060a01b0316600052600a602052604060002090565b505463ffffffff1690565b63ffffffff42911611600014611dd0577f3ac7d3823c11677fba9479ed26f696a3f17e16a5a5c39162fa6d183905aa673592916001611d79611da293611d5c611d293360018060a01b0316600052600b602052604060002090565b611d31610be5565b600081526001600160a01b0387166020820152908389830152600060608301528a608083015261476c565b6001600160a01b0384166000908152600a602052604090206119e8565b5001611d86868254614801565b90556001600160a01b0316600090815260056020526040902090565b611dad848254614801565b9055519182523391602090a3610aba6001600080516020616da883398151915255565b8251637d58ebb960e01b81526004810191909152602490fd5b6001600160a01b0383166000908152600c60205260409020610a15905b5491516327500c6d60e21b8152600481019290925260248201929092529081906044820190565b516305fd61ad60e01b81526001600160a01b0383166004820152602490fd5b3461042d57602036600319011261042d576001600160a01b03611e6d61045b565b1660005260066020526020604060002054604051908152f35b3461042d57611e943661106d565b611ea56104eb9695969492946142a7565b46860361125357611ec66110e46110dd876000526003602052604060002090565b61123a57611ed8611155611155616a3a565b60405163a9059cbb60e01b81526001600160a01b038416600482015260248101859052600092916020908290604490829087905af1801561122b57612022575b50600080516020616c688339815191526000526005602052600080516020616ca8833981519152611f4a85825461429a565b90558154611f60906001600160a01b0316611155565b90813b1561201e57604051631eb65ffb60e01b815260048101889052602481018990526001600160a01b038516604482015260648101969096526084860185905260a4860152849060c490829084905af192831561122b577f7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc149361200b575b50604080519485526020850195909552938301526001600160a01b039092169130918060608101611203565b8061121f61201892610b68565b38611fdf565b8280fd5b6120429060203d8111612049575b61203a8183610bb4565b81019061488f565b5038611f18565b503d612030565b3461042d57600080600319360112610a6b5761206a61480e565b61207261472b565b805b60078054821015610a55578252600080516020616ce88339815191528101546001600160a01b039081166000818152600a60205260409020909391905415610a3c57815b336000908152600b602052604090205481101561240e5784826120f46107518461074c3360018060a01b0316600052600b602052604060002090565b1614612109575b61210490614ec5565b6120b8565b939291906121306107958661074c3360018060a01b0316600052600b602052604060002090565b612401576001600160a01b0384166000908152600a60205260409020612155906107be565b336000908152600b6020526040902090959061217290829061054d565b50956003809701549061219b816108088960018060a01b0316600052600a602052604060002090565b50916121ad600180940191825461429a565b90558497806121d28561074c3360018060a01b0316600052600b602052604060002090565b50015498836121f78661074c3360018060a01b0316600052600b602052604060002090565b50015491838311610a1957825b848110612371575050612224816121049798999a9b6108b2821515614ddb565b93612248610ec78761074c3360018060a01b0316600052600b602052604060002090565b61225385338d616b2f565b50336000908152600b602052604090205411612275575b5050505090506120fb565b80612327612365926123217f0ad30059d4072ae2525ebaaa5b6878f1c50df5f3b921615031c7d0f12dc4a77b97986123046122e36122c53360018060a01b0316600052600b602052604060002090565b336000908152600b602052604090206122dd906107be565b9061054d565b50336000908152600b602052604090206122fe90849061054d565b90614e13565b336000908152600b6020526040902061231c90614e7b565b614ed9565b9761429a565b6040805133815260208101959095528401949094526001600160a01b038b166060840152608083019390935260a082019290925290819060c0820190565b0390a13880808061226a565b90916123926107be8c60018060a01b0316600052600a602052604060002090565b82116109fc57816123fc916109f0858f8f97966108086123e06109c38b8f6109ba8661080861099e6109e69a6002610995856108088a60018060a01b0316600052600a602052604060002090565b6001600160a01b03909a166000908152600a6020526040902090565b612204565b9091929361210490614ec5565b5050915061241b90614d99565b612074565b3461042d5760408060031936011261042d5761060b9061243e61045b565b906125036124fa600461249a60c0855161245781610b47565b600091818380935282602082015282898201528260608201528260808201528260a0820152015260018060a01b038097168152600a6020528560243591206119e8565b506124d48551966124aa88610b47565b825463ffffffff8082168a52602082811c909116908a0152871c166001600160a01b031687870152565b6001810154606087015260028101546080870152600381015460a0870152015460ff1690565b151560c0840152565b519182918291909160c08060e083019463ffffffff808251168552602082015116602085015260018060a01b036040820151166040850152606081015160608501526080810151608085015260a081015160a085015201511515910152565b3461042d57600036600319011261042d57602060ff600080516020616d8883398151915254166040519015158152f35b600754811015610569576007600052600080516020616ce88339815191520190600090565b3461042d57602036600319011261042d5760043560075481101561042d576007600052600080516020616ce883398151915201546040516001600160a01b039091168152602090f35b3461042d57602036600319011261042d576126196143b8565b600435600255005b6020908160408183019282815285518094520193019160005b828110612648575050505090565b909192938260a08261268f60019489516080809180511515845260018060a01b03602082015116602085015260408101516040850152606081015160608501520151910152565b0195019392910161263a565b3461042d5760208060031936011261042d576001600160a01b036126bd61045b565b16600052600b8152604060002080546126d581610bff565b916126e36040519384610bb4565b8183526000908152838120938084015b838310612708576040518061060b8782612621565b600482600192612717896166c2565b8152019601920191946126f3565b3461042d57602036600319011261042d576004356000526003602052602060ff604060002054166040519015158152f35b3461042d57600036600319011261042d5761276f6143b8565b61277761472b565b600080516020616d88833981519152600160ff198254161790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b3461042d57604036600319011261042d57602060ff6128106127dd610471565b600435600052600080516020616d68833981519152845260406000209060018060a01b0316600052602052604060002090565b54166040519015158152f35b3461042d5760e036600319011261042d5760243560043561283b610487565b61284361049d565b90608435906128536104eb6142a7565b468503611253576128746110e46110dd866000526003602052604060002090565b6129dd576001600160a01b038316600090815260046020526040902061289d906110e4906110dd565b6129bc576001600160a01b03838116939091906128d9906128bf858488616997565b6001600160a01b0316600090815260056020526040902090565b6128e484825461429a565b90556000546128fb906001600160a01b0316611155565b92833b1561042d57604051631eb65ffb60e01b815260048101879052602481018890526001600160a01b038316604482015260a4803560648301526084820183905260c480359183019190915290946000918691829084905af190811561122b577f0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe946114a0926129a9575b5060405193849316973097846040919493926060820195825260208201520152565b8061121f6129b692610b68565b38612987565b6040516305fd61ad60e01b81526001600160a01b0384166004820152602490fd5b604051632ef6974160e11b815260048101859052602490fd5b600080600319360112610a6b57612a0b61480e565b612a1361472b565b600080516020616d48833981519152600052600c6020527f36dda7fac5d19118789d83982ba89ab2ea0a81d18d22af2919c731e22f75f227543410612c3857600080516020616d48833981519152600052600a602052612a80600080516020616d288339815191526107be565b600080516020616d48833981519152600052600a602052600080516020616d288339815191525415612c1f57600080516020616d48833981519152600052600a602052612ade611cc382600080516020616d288339815191526119e8565b63ffffffff42911611600014612bf457336000908152600b60205260409020600191612b6a91612b3f90612b10610be5565b868152600080516020616d4883398151915260208201529083604083015286606083015234608083015261476c565b600080516020616d48833981519152600052600a602052600080516020616d288339815191526119e8565b5001612b77348254614801565b9055600080516020616d488339815191526000526005602052600080516020616c88833981519152612baa348254614801565b905560405134815233907fe7466ea83435490635fc76a5f33da4815758ab48b1d45858f0452ca64655693790602090a2612bf16001600080516020616da883398151915255565b80f35b612c00610a15916147f3565b604051637d58ebb960e01b815260048101919091529081906024820190565b604051637d58ebb960e01b815260016004820152602490fd5b600080516020616d48833981519152600052600c6020527f36dda7fac5d19118789d83982ba89ab2ea0a81d18d22af2919c731e22f75f227546040516327500c6d60e21b81526004810191909152346024820152604490fd5b60408060031936011261042d57612ca661045b565b9060243590612cb361472b565b3415612cc457505050610aba6149cf565b6001600160a01b03831692600080516020616c688339815191528403612cf0575050610aba9150614b99565b83600052600460205260ff826000205416612d0757005b612d0f61480e565b612d1761472b565b6001600160a01b0381166000908152600460205260409020612d3c906110e4906110dd565b612d8b576001600160a01b0381166000908152600c60205260409020548310612d6b57611c7f8330338761483e565b6001600160a01b03166000908152600c60205260409020610a1590611e06565b90516305fd61ad60e01b81526001600160a01b039091166004820152602490fd5b606036600319011261042d57602435600435612dc6610487565b46820361125357612de76110e46110dd856000526003602052604060002090565b612f7857600154803410612f585750600080516020616d488339815191526000526005602052600080516020616c88833981519152612e27348254614801565b9055612e386113a560025434614da8565b90612e43823461429a565b600080516020616d488339815191526000526006602052907fa2e5aefc6e2cbe2917a296f0fd89c5f915c487c803db1d98eccb43f14012d711612e87848254614801565b9055600054612e9e906001600160a01b0316611155565b803b1561042d5760405163305f478560e21b8152466004820152602481018790526001600160a01b0383166044820152606481018490526084810194909452600090849060a490829084905af192831561122b577f0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc593612f45575b50604080519485526020850195909552938301526001600160a01b039092169133918060608101611203565b8061121f612f5292610b68565b38612f19565b6040516358f8331360e11b81526004810191909152346024820152604490fd5b604051632ef6974160e11b815260048101849052602490fd5b3461042d57600036600319011261042d57602060405160008152f35b3461042d57600080600319360112610a6b5760075490612fcc82616927565b91815b818310612fe4576040518061060b8682611563565b8091815b336000908152600b60205260409020548110156130b357336000908152600b6020526040902061301d9061075190839061054d565b61302c61115561163488612592565b6001600160a01b039091161461304b575b61304690614d99565b612fe8565b9261306f6107958561074c3360018060a01b0316600052600b602052604060002090565b6130a957336000908152600b60205260409020613046916130a19160039061309890889061054d565b50015490614801565b93905061303d565b9261304690614d99565b50926130ec91926130c661163483612592565b906130d26117c1610bd6565b60208201526130e18287615115565b526117ea8186615115565b9190612fcf565b3461042d57604036600319011261042d5761310c61045b565b60243561311881610c3b565b6131206143b8565b6001600160a01b0382166000908152600460205260409020805460ff191660ff8315151617905561314d57005b610aba90616670565b3461042d57600036600319011261042d576020613171616a3a565b6040516001600160a01b039091168152f35b3461042d57602036600319011261042d5761319f6104eb6142a7565b600435600155005b3461042d57602036600319011261042d576004356131c361480e565b6131cb61472b565b600080516020616c68833981519152600052600c6020527fc87f3dae70d3ac23e82424b4132d779583ae145fc334a52160c2d1a246c2bc2b548110613443576040516323b872dd60e01b8152336004820152306024820152604481018290526020816064816000600080516020616c688339815191525af1801561122b57613425575b50600080516020616c68833981519152600052600a60205261327d600080516020616d088339815191526107be565b600080516020616c68833981519152600052600a6020526132ba60046132b183600080516020616d088339815191526119e8565b50015460ff1690565b61340b57600080516020616c68833981519152600052600a6020526132f0611cc382600080516020616d088339815191526119e8565b63ffffffff42911611613377575b50600080516020616c688339815191526000526005602052600080516020616ca8833981519152613330828254614801565b905560405190815233907fc138e3bd6d13eefa5cb01c0d35c5794001141efaf4e5ad888cad059935f8338390602090a2610aba6001600080516020616da883398151915255565b336000908152600b602052604090206001916133f6916133cb905b61339a610be5565b60008152600080516020616c688339815191526020820152908360408301526000606083015286608083015261476c565b600080516020616c68833981519152600052600a602052600080516020616d088339815191526119e8565b5001613403828254614801565b9055386132fe565b6040516311cf1b0760e31b81526004810191909152602490fd5b61343c9060203d81116120495761203a8183610bb4565b503861324e565b60008052600c6020527f13649b2456f1b42fef0f0040b3aaeabcd21a76a0f3f5defd4f583839455116e8546040516327500c6d60e21b815260048101919091526024810191909152604490fd5b3461042d57604036600319011261042d576134a961045b565b7ff54d3b756d286b6b08e5d4eda6dfe5b135664abf029e58e637cbf013c442c9506020602435926134d86143b8565b6001600160a01b03166000818152600c83526040908190208590555193845292a2005b3461042d57602036600319011261042d576001600160a01b0361351c61045b565b16600052600b6020526020604060002054604051908152f35b3461042d57608036600319011261042d5760046024358135613555610487565b468203611552576135766110e46110dd856000526003602052604060002090565b61381157613588611155611155616a3a565b604080516370a0823160e01b8082523082890190815292979360209390919084908490819083010381855afa92831561122b576000936137f2575b5088516323b872dd60e01b815233868201908152306020820152606435604082015290919085908390819060600103816000875af191821561122b5785926137d5575b50895190815230868201908152909283918290819060200103915afa90811561122b5761363b936000926114c257505061429a565b6001548082106137b25750600080516020616c688339815191526000526005602052600080516020616ca8833981519152613677828254614801565b905561369261368b6113a560025484614da8565b809261429a565b600080516020616c688339815191526000526006602052917f5e5777fab7622aff3c042c1ece74307c2e9d699a9da444f416c35f2e1def28a56136d6838254614801565b90556000546136ed906001600160a01b0316611155565b91823b1561042d57875163305f478560e21b8152918201868152602081018890526001600160a01b038616604082015260608101859052608081019190915290939160009185919082908490829060a00103925af192831561122b577fcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a899361379f575b508551938452602084019490945260408301526001600160a01b03909216913391606090a35160018152602090f35b8061121f6137ac92610b68565b38613770565b86516358f8331360e11b8152808401918252602082018390529081906040010390fd5b6137eb90833d85116120495761203a8183610bb4565b5038613606565b61380a919350843d86116114e8576114d98183610bb4565b91386135c3565b6040516363b5c9db60e01b815280850184815281906020010390fd5b3461042d57604036600319011261042d57610aba60043561384c610471565b9080600052600080516020616d68833981519152602052613874600160406000200154614410565b6145b6565b3461042d57604036600319011261042d57610aba60243561389981610c3b565b6138a16143b8565b600435600052600360205260406000209060ff801983541691151516179055565b3461042d57602036600319011261042d576001600160a01b036138e361045b565b166000526004602052602060ff604060002054166040519015158152f35b3461042d57602036600319011261042d576001600160a01b0361392261045b565b1660005260056020526020604060002054604051908152f35b3461042d57600036600319011261042d576020600254604051908152f35b3461042d5761396736610ada565b600854604080516001600160a01b03868116602083019081529086169282019290925260608101849052608080820184905281529194937f5e2d465404712c588a7dffc8622a4fad4f4522708eaf138e2b70910cd36992a09290916139da916139d160a082610bb4565b51902095614d99565b60088190556040805194855260208501919091526001600160a01b039182169490911692a4602060405160018152f35b3461042d57606036600319011261042d57613a2361045b565b602435613a2f81610c3b565b613a37610c17565b91613a406143b8565b6001600160a01b0381166000908152600460205260409020613a61906110dd565b613bf357613bed81613abb84613aaa7fc8c34f23fafb34e68119c1d231ef03d0d47225b15e2c4de3efbefa14b0181d869560018060a01b03166000526004602052604060002090565b9060ff801983541691151516179055565b6001600160a01b0381166000908152600a60205260409020613bc89095613b45613af3613aed60095463ffffffff1690565b8361665a565b613b0a613afe610bf2565b63ffffffff9092168252565b63ffffffff831660208201526001600160a01b03851660408201526000988960608301528960808301528960a08301528960c0830152615141565b6001600160a01b0383166000908152600a60205260409020613b9c613b78613b7260095463ffffffff1690565b84615129565b613b8f613b83610bf2565b63ffffffff9095168552565b63ffffffff166020840152565b6001600160a01b03841660408301528760608301528760808301528760a08301528760c0830152615141565b613bd181616670565b60405193151584526001600160a01b0316929081906020820190565b0390a280f35b60405163411befff60e11b81526001600160a01b039190911660048201529015156024820152604490fd5b3461042d57606036600319011261042d57613c37610471565b6044359067ffffffffffffffff9081831161042d573660238401121561042d57826004013591821161042d573660248360061b8501011161042d576024610aba930190600435615209565b3461042d57602036600319011261042d576004803590613ca061480e565b613ca861472b565b33600052600b602052604091826000205480821015613fd25750336000908152600b60205260409020613ce09061075190839061054d565b91613d006107be8460018060a01b0316600052600a602052604060002090565b336000908152600b60205260409020909290613d1d90829061054d565b5060038091015495613d45856108088860018060a01b0316600052600a602052604060002090565b5096613d57600180990191825461429a565b9055336000908152600b602052604081209092908190613d7890869061054d565b5001549088613d9d8661074c3360018060a01b0316600052600b602052604060002090565b50015495878711613fc457865b888110613f105750505082613dc4916108b2821515614ddb565b95613de8610ec78561074c3360018060a01b0316600052600b602052604060002090565b336000908152600b60205260409020613e0690610ec790869061054d565b613e11873383616b2f565b50336000908152600b6020526040902088905411613e3e575b600080516020616da8833981519152889055005b613ec2837f0ad30059d4072ae2525ebaaa5b6878f1c50df5f3b921615031c7d0f12dc4a77b98613ea5613f01976122fe613e8d6122c53360018060a01b0316600052600b602052604060002090565b50336000908152600b6020526040902090929061054d565b336000908152600b60205260409020613ebd90614e7b565b61429a565b9151338152602081019590955260408501959095526001600160a01b039094166060840152608083019390935260a082019290925290819060c0820190565b0390a138808080808080613e2a565b9091929394613f346107be8b60018060a01b0316600052600a602052604060002090565b8211613fa85781613fa3916109f0866109e68f968f9b9a999897613f876109c38e610808936109ba8f61080861099e89926002610995856108088a60018060a01b0316600052600a602052604060002090565b6001600160a01b03909d166000908152600a6020526040902090565b613daa565b508451637d58ebb960e01b815291820190815281906020010390fd5b835163374c934360e11b8152fd5b925163abe5c32f60e01b815291820190815260208101929092529081906040010390fd5b3461042d57602036600319011261042d576001600160a01b0361401761045b565b16600052600a6020526020604060002054604051908152f35b3461042d57608036600319011261042d5760043561404c610471565b90614055610487565b6064356140636104eb6142a7565b6001600160a01b0384166000908152600460205260409020614088906110e4906110dd565b61424f576208275083036140f657816140b182600080516020616cc883398151915294876158de565b6001600160a01b03851660009081526005602052604090206140d483825461429a565b90556040805194855260208501929092526001600160a01b03908116941692a3005b61044d830361411f578161411a82600080516020616cc88339815191529487615c4f565b6140b1565b600a8303614142578161411a82600080516020616cc88339815191529487615f7b565b61a4b18303614166578161411a82600080516020616cc883398151915294876154d8565b61a4ba830361418a578161411a82600080516020616cc883398151915294876156e2565b61014483036141ae578161411a82600080516020616cc88339815191529487616209565b61138883036141d2578161411a82600080516020616cc883398151915294876163a1565b60a983036141f5578161411a82600080516020616cc883398151915294876160c2565b61a70e8303614219578161411a82600080516020616cc88339815191529487615dc3565b612105830361423d578161411a82600080516020616cc88339815191529487616513565b604051639474dee160e01b8152600490fd5b6040516305fd61ad60e01b81526001600160a01b0385166004820152602490fd5b634e487b7160e01b600052601160045260246000fd5b60001981019190821161429557565b614270565b9190820391821161429557565b604051602081017f33fe247d78fee2e7fd135c405eda4bd2911c0a73c0a81b36c3bcc967dd06e5ad8152602082526142de82610b98565b9051902060ff191690565b3461042d57600036600319011261042d5760206143046142a7565b604051908152f35b3461042d57604036600319011261042d5760a0614377614371608061432f61045b565b60405161433b81610b7c565b60009281848093528260208201528260408201528260608201520152600180861b03168152600b6020526040602435912061054d565b506166c2565b6143b660405180926080809180511515845260018060a01b03602082015116602085015260408101516040850152606081015160608501520151910152565bf35b3360009081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604090205460ff16156143f157565b60405163e2517d3f60e01b815233600482015260006024820152604490fd5b6000818152600080516020616d688339815191526020908152604080832033845290915290205460ff16156144425750565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b6001600160a01b03811660009081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d6020526040812054600080516020616d688339815191529060ff1661450f57818052602090815260408083206001600160a01b038516600090815292529020805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50905090565b6000818152600080516020616d68833981519152602081815260408084206001600160a01b038716855290915282205491929160ff166145af57818352602090815260408084206001600160a01b038616600090815292529020805460ff1916600117905533926001600160a01b0316917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b5050905090565b6000818152600080516020616d68833981519152602081815260408084206001600160a01b038716855290915282205491929160ff16156145af57818352602090815260408084206001600160a01b038616600090815292529020805460ff1916905533926001600160a01b0316917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b6146c29061465a6146fc565b6146626146fc565b61466a6146fc565b6001600080516020616da8833981519152556146846146fc565b61468c6146fc565b600080516020616d88833981519152805460ff191690556146ab6146fc565b67016345785d8a0000600155612710600255614460565b5060018060a01b03166bffffffffffffffffffffffff60a01b60005416176000556001600855621baf8063ffffffff196009541617600955565b60ff600080516020616dc88339815191525460401c161561471957565b604051631afcd79f60e31b8152600490fd5b60ff600080516020616d88833981519152541661474457565b60405163d93c066560e01b8152600490fd5b634e487b7160e01b600052600060045260246000fd5b8054600160401b811015610b63576147899160018201815561054d565b9190916147ee576080816147ae600393511515859060ff801983541691151516179055565b60208101518454610100600160a81b03191660089190911b610100600160a81b031617845560408101516001850155606081015160028501550151910155565b614756565b906001820180921161429557565b9190820180921161429557565b600080516020616da8833981519152600281541461482c5760029055565b604051633ee5aeb560e01b8152600490fd5b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064820192909252610be39161488a82608481015b03601f198101845283610bb4565b6148a7565b9081602091031261042d57516148a481610c3b565b90565b6000806148f19260018060a01b03169360208151910182865af13d15614948573d906148d282614950565b916148e06040519384610bb4565b82523d6000602084013e5b8361496c565b805190811515918261492a575b50506149075750565b604051635274afe760e01b81526001600160a01b03919091166004820152602490fd5b6149419250906020806110e493830101910161488f565b38806148fe565b6060906148eb565b67ffffffffffffffff8111610b6357601f01601f191660200190565b90614993575080511561498157805190602001fd5b604051630a12f52160e11b8152600490fd5b815115806149c6575b6149a4575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561499c565b6149d761480e565b6149df61472b565b600080516020616d48833981519152600052600c6020527f36dda7fac5d19118789d83982ba89ab2ea0a81d18d22af2919c731e22f75f227543410612c3857600080516020616d48833981519152600052600a602052614a4c600080516020616d288339815191526107be565b600080516020616d48833981519152600052600a602052600080516020616d288339815191525415612c1f57600080516020616d48833981519152600052600a602052614aaa611cc382600080516020616d288339815191526119e8565b63ffffffff42911611600014612bf457336000908152600b60205260409020600191614b0d91612b3f90614adc610be5565b60008152600080516020616d488339815191526020820152908360408301526000606083015234608083015261476c565b5001614b1a348254614801565b9055600080516020616d488339815191526000526005602052600080516020616c88833981519152614b4d348254614801565b905560405134815233907fe7466ea83435490635fc76a5f33da4815758ab48b1d45858f0452ca6465569379080602081015b0390a2610be36001600080516020616da883398151915255565b614ba161480e565b614ba961472b565b600080516020616c68833981519152600052600c6020527fc87f3dae70d3ac23e82424b4132d779583ae145fc334a52160c2d1a246c2bc2b548110613443576040516323b872dd60e01b8152336004820152306024820152604481018290526020816064816000600080516020616c688339815191525af1801561122b57614d6f575b50600080516020616c68833981519152600052600a602052614c5b600080516020616d088339815191526107be565b600080516020616c68833981519152600052600a602052614c8f60046132b183600080516020616d088339815191526119e8565b61340b57600080516020616c68833981519152600052600a602052614cc5611cc382600080516020616d088339815191526119e8565b63ffffffff42911611614d3b575b50600080516020616c688339815191526000526005602052600080516020616ca8833981519152614d05828254614801565b905560405190815233907fc138e3bd6d13eefa5cb01c0d35c5794001141efaf4e5ad888cad059935f83383908060208101614b7f565b336000908152600b60205260409020600191614d5a916133cb90613392565b5001614d67828254614801565b905538614cd3565b614d869060203d81116120495761203a8183610bb4565b5038614c2c565b6040513d6000823e3d90fd5b60001981146142955760010190565b8181029291811591840414171561429557565b8115614dc5570490565b634e487b7160e01b600052601260045260246000fd5b15614de257565b60405162461bcd60e51b8152602060048201526009602482015268139bc814995dd85c9960ba1b6044820152606490fd5b91906147ee57808203614e24575050565b600381614e4260ff83945416859060ff801983541691151516179055565b80548454610100600160a81b031916610100600160a81b0390911617845560018101546001850155600281015460028501550154910155565b80548015614eaf576000190190614e92828261054d565b6147ee576003816000809355826001820155826002820155015555565b634e487b7160e01b600052603160045260246000fd5b6001600160ff1b0381146142955760010190565b600160ff1b8114614295576000190190565b90614ef76104eb6142a7565b60005b8251811015615110576150ba9060407fb6f449f07ceaf55392c9899e0797c6529908ae827c2d498c682e90d42c2411676150b2614f4a83614f3b868a615115565b5101516001600160a01b031690565b6001600160a01b0381166000908152600a60205260409020614f6b906107be565b93614fae6004614f9f614f908560018060a01b0316600052600a602052604060002090565b614f9989614286565b906119e8565b5001805460ff19166001179055565b614fb785614286565b6150bf575b61509a614ff0614fe2876108088660018060a01b0316600052600a602052604060002090565b505460201c63ffffffff1690565b6001600160a01b0384166000908152600a6020526040902061502061501a60095463ffffffff1690565b83615129565b61506d60016150458b6108088a60018060a01b0316600052600a602052604060002090565b50015491615060615054610bf2565b63ffffffff9096168652565b63ffffffff166020850152565b6001600160a01b03861683860152606083015260006080830152600060a0830152600060c0830152615141565b519384526001600160a01b0316929081906020820190565b0390a2614d99565b614efa565b6001600160a01b038216600090815260066020908152604080832054600a9092529091206002906150ef90614f90565b5001556001600160a01b038216600090815260066020526040812055614fbc565b509050565b80518210156105695760209160051b010190565b91909163ffffffff8080941691160191821161429557565b8054600160401b811015610b635761515e916001820181556119e8565b6147ee5760046151f660c08461518e63ffffffff610be3975116869063ffffffff1663ffffffff19825416179055565b60208181015186546040808501519290931b67ffffffff0000000016640100000000600160e01b031990911617911b68010000000000000000600160e01b0316178555606081015160018601556080810151600286015560a081015160038601550151151590565b91019060ff801983541691151516179055565b90916152166104eb6142a7565b6801bc16d674ec8000008204156152ea576001600160a01b03831693843b1561042d57600060405180966337a6c88160e01b825281868161525c87898460048501615382565b03925af191821561122b577f33fa5868f7aec03a4bae17d5ca8d0869adac2237298253519ebf80c25ddf7fe3956152d2936152d7575b50600080516020616d488339815191526000526005602052600080516020616c888339815191526152c485825461429a565b905560405194859485615399565b0390a1565b8061121f6152e492610b68565b38615292565b60405162461bcd60e51b815260206004820152601760248201527f457468206e6f7420656e6f75676820746f207374616b650000000000000000006044820152606490fd5b9190808252602080920192916000905b82821061534d575050505090565b9293919290916001906001600160a01b03615367876104b3565b1681528583013583820152604090810195019392019061533f565b6040906148a494928152816020820152019161532f565b9081526001600160a01b0390911660208201526060604082018190526148a49391019161532f565b9190602092838183031261042d5780519067ffffffffffffffff821161042d57019281601f8501121561042d5783516153f981614950565b926154076040519485610bb4565b81845282828701011161042d5760005b81811061542c57508260009394955001015290565b8581018301518482018401528201615417565b9261010094929160018060a01b0392838092168652166020850152166040830152606082015260006080820152600060a082015260e060c0820152600060e08201520190565b600080516020616d4883398151915281526001600160a01b03918216602082015291166040820152606081019190915260006080820181905260a0820181905260e060c083018190528201526101000190565b6001600160a01b03811690600080516020616d4883398151915282036155625750508161551f92600092604051809581948293634fb1a07b60e01b84523060048501615485565b03917372ce9c846789fdb6fc1f34ac4ad25dd9ef7031ef5af1801561122b57615546575b50565b615543903d806000833e61555a8183610bb4565b8101906153c1565b600080516020616c688339815191528293921460001461562f5760405163095ea7b360e01b815273d92023e9d9911199a6711321d1277285e6d4e2db60048201526024810185905292602090849060449082906000905af192831561122b57600093615611575b506155ea6040519485938493634fb1a07b60e01b855230906004860161543f565b03818373d92023e9d9911199a6711321d1277285e6d4e2db5af1801561122b576155465750565b6156289060203d81116120495761203a8183610bb4565b50386155c9565b60405163095ea7b360e01b815273a3a7b6f88361f48403514059f1f16c8e78d60eec60048201526024810185905292602090849060449082906000905af192831561122b576000936156c4575b5061569d6040519485938493634fb1a07b60e01b855230906004860161543f565b03818373a3a7b6f88361f48403514059f1f16c8e78d60eec5af1801561122b576155465750565b6156db9060203d81116120495761203a8183610bb4565b503861567c565b6001600160a01b03811690600080516020616d48833981519152820361574f5750508161572992600092604051809581948293634fb1a07b60e01b84523060048501615485565b039173c840838bc438d73c16c2f8b22d2ce3669963cd485af1801561122b576155465750565b600080516020616c688339815191528293921460001461581c5760405163095ea7b360e01b815273e4e2121b479017955be0b175305b35f312330bae60048201526024810185905292602090849060449082906000905af192831561122b576000936157fe575b506157d76040519485938493634fb1a07b60e01b855230906004860161543f565b03818373e4e2121b479017955be0b175305b35f312330bae5af1801561122b576155465750565b6158159060203d81116120495761203a8183610bb4565b50386157b6565b60405163095ea7b360e01b815273b2535b988dce19f9d71dfb22db6da744acac21bf60048201526024810185905292602090849060449082906000905af192831561122b576000936158b1575b5061588a6040519485938493634fb1a07b60e01b855230906004860161543f565b03818373b2535b988dce19f9d71dfb22db6da744acac21bf5af1801561122b576155465750565b6158c89060203d81116120495761203a8183610bb4565b5038615869565b9081602091031261042d575190565b90916001600160a01b038216600080516020616d4883398151915281036159e45750604051636bb825d760e11b8152620298106004820152909150602081602481730d7e906bd9cafa154b048cfa766cc1e54e39af9b5afa801561122b5761594f916000916159c6575b5082614801565b90737f2b8c31f88b6006c382775eea88297ec1e3e905803b1561042d57604051636705b1e760e11b81526001600160a01b0390941660048501526024840191909152620298106044840152600091839190829081606481015b03925af1801561122b576159b95750565b8061121f610be392610b68565b6159de915060203d81116114e8576114d98183610bb4565b38615948565b600080516020616c688339815191528103615b1f57604051636bb825d760e11b8152614e206004820152602091908281602481730d7e906bd9cafa154b048cfa766cc1e54e39af9b5afa801561122b57615b02575b5060405163095ea7b360e01b8152737ac440cae8eb6328de4fa621163a792c1ea9d4fe600482015260248101849052908290829060449082906000905af1801561122b57615ae4575b5050737ac440cae8eb6328de4fa621163a792c1ea9d4fe803b1561042d5760405163790cfd3360e11b81526001600160a01b0393841660048201529390921660248401526044830152614e2060648301526000908290818381608481016159a8565b81615afa92903d106120495761203a8183610bb4565b503880615a82565b615b1890833d85116114e8576114d98183610bb4565b5038615a39565b604051636bb825d760e11b8152614e206004820152602091908281602481730d7e906bd9cafa154b048cfa766cc1e54e39af9b5afa801561122b57615c32575b5060405163095ea7b360e01b8152737ac440cae8eb6328de4fa621163a792c1ea9d4fe600482015260248101849052908290829060449082906000905af1801561122b57615c14575b505073d8a791fe2be73eb6e6cf1eb0cb3f36adc9b3f8f9803b1561042d5760405163790cfd3360e11b81526001600160a01b0393841660048201529390921660248401526044830152614e2060648301526000908290608490829084905af1801561122b576159b95750565b81615c2a92903d106120495761203a8183610bb4565b503880615ba8565b615c4890833d85116114e8576114d98183610bb4565b5038615b5f565b6001600160a01b038116600080516020616d488339815191528103615ce6575050732a3dd3eb832af982ec71669e178424b10dca2ede91823b1561042d5760405163cd58657960e01b8152600160048201526001600160a01b039092166024830152604482018190526000606483018190526084830181905260c060a484015260c483018190529192839182908160e481016159a8565b60405163095ea7b360e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede60048201526024810185905290602090829060449082906000905af1801561122b57615da5575b50732a3dd3eb832af982ec71669e178424b10dca2ede803b1561042d5760405163cd58657960e01b8152600160048201526001600160a01b03938416602482015260448101949094529116606483015260006084830181905260c060a484015260c4830181905290829081838160e481016159a8565b615dbc9060203d81116120495761203a8183610bb4565b5038615d2f565b6001600160a01b038116600080516020616d488339815191528103615e5a575050739cb4706e20a18e59a48ffa7616d700a3891e186191823b1561042d5760405163cd58657960e01b8152600160048201526001600160a01b039092166024830152604482018190526000606483018190526084830181905260c060a484015260c483018190529192839182908160e481016159a8565b60405163095ea7b360e01b8152739cb4706e20a18e59a48ffa7616d700a3891e186160048201526024810185905290602090829060449082906000905af1801561122b57615f19575b50739cb4706e20a18e59a48ffa7616d700a3891e1861803b1561042d5760405163cd58657960e01b8152600160048201526001600160a01b03938416602482015260448101949094529116606483015260006084830181905260c060a484015260c4830181905290829081838160e481016159a8565b615f309060203d81116120495761203a8183610bb4565b5038615ea3565b93909163ffffffff9360e0969360018060a01b0392838092168852166020870152166040850152606084015216608082015260c060a0820152600060c08201520190565b906001600160a01b038216600080516020616d48833981519152810361600057507399c9fc46f92e8a1c0dec1b1747d010903e884be19150813b1561042d57604051639a2ac6d560e01b81526001600160a01b0390911660048201526000602482018190526060604483015260648201819052909290918391829081608481016159a8565b6160098361670c565b60405163095ea7b360e01b81527399c9fc46f92e8a1c0dec1b1747d010903e884be16004820152602481018690529091602090829060449082906000905af1801561122b576160a4575b506160615a63ffffffff1690565b907399c9fc46f92e8a1c0dec1b1747d010903e884be191823b1561042d576000946159a886926040519889978896879563041c592960e51b875260048701615f37565b6160bb9060203d81116120495761203a8183610bb4565b5038616053565b906001600160a01b038216600080516020616d4883398151915281036161475750733b95bc951ee0f553ba487327278cac44f29715e59150813b1561042d57604051639a2ac6d560e01b81526001600160a01b0390911660048201526000602482018190526060604483015260648201819052909290918391829081608481016159a8565b61615083616882565b60405163095ea7b360e01b8152733b95bc951ee0f553ba487327278cac44f29715e56004820152602481018690529091602090829060449082906000905af1801561122b576161eb575b506161a85a63ffffffff1690565b90733b95bc951ee0f553ba487327278cac44f29715e591823b1561042d576000946159a886926040519889978896879563041c592960e51b875260048701615f37565b6162029060203d81116120495761203a8183610bb4565b503861619a565b6001600160a01b038116600080516020616d4883398151915281036162b057505060405163e8b99b1b60e01b81526001600160a01b03919091166004820152600060248201819052604482018390526064820181905260848201523060a482015290602090829060c49082907357891966931eb4bb6fb81430e6ce0a03aabde0635af1801561122b576162995750565b6155439060203d81116114e8576114d98183610bb4565b60405163095ea7b360e01b81527357891966931eb4bb6fb81430e6ce0a03aabde063600482015260248101859052602094909290918590849060449082906000905af192831561122b578593616384575b5060405163e8b99b1b60e01b81526001600160a01b03948516600482015293166024840152604483015260006064830181905260848301523060a4830152818060c48101038160007357891966931eb4bb6fb81430e6ce0a03aabde0635af1801561122b5761636e575050565b8161554392903d106114e8576114d98183610bb4565b61639a90843d86116120495761203a8183610bb4565b5038616301565b9091906001600160a01b038116600080516020616d488339815191528103616425575050507395fc37a27a2f68e3a647cdc081f0a89bb47c3012803b1561042d57604051639a2ac6d560e01b81526001600160a01b0390921660048301526000602483018190526060604484015260648301819052908290818381608481016159a8565b60405163095ea7b360e01b81527395fc37a27a2f68e3a647cdc081f0a89bb47c30126004820152602481018490529293919290602090829060449082906000905af1801561122b576164f5575b5061647c826167f9565b7395fc37a27a2f68e3a647cdc081f0a89bb47c301290813b1561042d5760405163041c592960e51b81526001600160a01b0394851660048201529084166024820152919092166044820152606481019290925260006084830181905260c060a484015260c4830181905290829081838160e481016159a8565b61650c9060203d81116120495761203a8183610bb4565b5038616472565b906001600160a01b038216600080516020616d4883398151915281036165985750733154cf16ccdb4c6d922629664174b904d80f2c359150813b1561042d57604051639a2ac6d560e01b81526001600160a01b0390911660048201526000602482018190526060604483015260648201819052909290918391829081608481016159a8565b6165a18361670c565b60405163095ea7b360e01b8152733154cf16ccdb4c6d922629664174b904d80f2c356004820152602481018690529091602090829060449082906000905af1801561122b5761663c575b506165f95a63ffffffff1690565b90733154cf16ccdb4c6d922629664174b904d80f2c3591823b1561042d576000946159a886926040519889978896879563041c592960e51b875260048701615f37565b6166539060203d81116120495761203a8183610bb4565b50386165eb565b63ffffffff918216908216039190821161429557565b600754600160401b811015610b63576001810180600755811015610569576007600052600080516020616ce88339815191520180546001600160a01b0319166001600160a01b03909216919091179055565b906040516166cf81610b7c565b825460ff81161515825260081c6001600160a01b031660208201526001830154604082015260028301546060820152600392909201546080830152565b6001600160a01b038116600080516020616c68833981519152810361673957506006602160991b01919050565b73dac17f958d2ee523a2206206994597c13d831ec7810361676e5750507394b008aa00579c1307b0ef2c499ad98a8ce58e5890565b73a0b86991c6218b36c1d19d4a2e9eb0ce3606eb4881036167a3575050730b2c639c533813f4aa9d7837caf62653d097ff8590565b736b175474e89094c44da98b954eedeac495271d0f036167d6575073da10009cbd5d07dd0cecc66161fc93d7c9000da190565b6040516305fd61ad60e01b81526001600160a01b03919091166004820152602490fd5b6001600160a01b031673dac17f958d2ee523a2206206994597c13d831ec78103616836575073201eba5cc46d216ce6dc03f6a759e8e766e956ae90565b73a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48810361686a575073201eba5cc46d216ce6dc03f6a759e8e766e956ae90565b602490604051906305fd61ad60e01b82526004820152fd5b6001600160a01b031673dac17f958d2ee523a2206206994597c13d831ec781036168bf575073f417f5a458ec102b90352f697d6e2ac3a3d2851f90565b73a0b86991c6218b36c1d19d4a2e9eb0ce3606eb4881036168f3575073b73603c5d87fa094b7314c74ace2e64d165016fb90565b736b175474e89094c44da98b954eedeac495271d0f810361686a5750731c466b9371f8aba0d7c458be10a62192fcb8aa7190565b9061693182610bff565b604061693f81519283610bb4565b8382528193616950601f1991610bff565b01906000805b838110616964575050505050565b8251908382019180831067ffffffffffffffff841117610b63576020928552838152828481830152828801015201616956565b60405163a9059cbb60e01b60208201526001600160a01b0390921660248301526044820192909252610be39161488a826064810161487c565b156169d757565b60405162461bcd60e51b815260206004820152603560248201527f546f6b656e4272696467652e42726964676546696e616c697a655374616b696e60448201527419d3595cdcd859d94e8818d85b1b0819985a5b1959605a1b6064820152608490fd5b466208275003616a4f576004605360981b0190565b4661044d03616a7057734f9a0e7fd2bf6067db6994cf12e4495df938e6e990565b46600a03616a83576006602160991b0190565b4661a4b103616aa4577382af49447d8a07e3bd95bd0d56f35241523fbab190565b4661a4ba03616ac55773722e8bdd2ce80a4422e880164f2079488e11536590565b4661014403616ae657735aea5775959fbc2557cc8789bc1bf90a239d9a9190565b4661138803616b015760405163e31d668960e01b8152600490fd5b4660a903616b1b576040516322c20c6960e11b8152600490fd5b466121050361423d576006602160991b0190565b9190616b536110e46110dd8560018060a01b03166000526004602052604060002090565b6129bc576001600160a01b0383166000908152600560205260409020616b7a83825461429a565b90556001600160a01b0392808416600080516020616d488339815191528103616be1575050814710616bcf5760009283928392839283918315616bc5575b1690f11561122b57600190565b6108fc9250616bb8565b604051632c1d501360e11b8152600490fd5b6040516370a0823160e01b81523060048201529195945090602081602481855afa801561122b578491600091616c49575b5010616c2857616c23939450616997565b600190565b6040516311745b6160e21b81526001600160a01b0386166004820152602490fd5b616c61915060203d81116114e8576114d98183610bb4565b38616c1256fe000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2a1829a9003092132f585b6ccdd167c19fe9774dbdea4260287e8a8e8ca8185d7a550ba85c46b24b567d2e17cd597f2283877afab43603f46d5de7858f1bdb73108ffba656e4ac665e1aa4bb6a342e1a69d4cc12fd751ac862fa3fe27b0c7524aa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688d0bcf4df132c65dad73803c5e5e1c826f151a3342680034a8a4c8e5f8eb0c13c3e65e660a0d3d61f62bb0309259c5a3ded6558e90d0e8aff997e553e7b030a75000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220d58241d44d92b93b36cd06959a897dc4c0da1fea269905c4852f2a1bb804bce964736f6c63430008140033",
}

// L1PoolManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use L1PoolManagerMetaData.ABI instead.
var L1PoolManagerABI = L1PoolManagerMetaData.ABI

// L1PoolManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1PoolManagerMetaData.Bin instead.
var L1PoolManagerBin = L1PoolManagerMetaData.Bin

// DeployL1PoolManager deploys a new Ethereum contract, binding an instance of L1PoolManager to it.
func DeployL1PoolManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L1PoolManager, error) {
	parsed, err := L1PoolManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1PoolManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1PoolManager{L1PoolManagerCaller: L1PoolManagerCaller{contract: contract}, L1PoolManagerTransactor: L1PoolManagerTransactor{contract: contract}, L1PoolManagerFilterer: L1PoolManagerFilterer{contract: contract}}, nil
}

// L1PoolManager is an auto generated Go binding around an Ethereum contract.
type L1PoolManager struct {
	L1PoolManagerCaller     // Read-only binding to the contract
	L1PoolManagerTransactor // Write-only binding to the contract
	L1PoolManagerFilterer   // Log filterer for contract events
}

// L1PoolManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1PoolManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1PoolManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1PoolManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1PoolManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1PoolManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1PoolManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1PoolManagerSession struct {
	Contract     *L1PoolManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1PoolManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1PoolManagerCallerSession struct {
	Contract *L1PoolManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L1PoolManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1PoolManagerTransactorSession struct {
	Contract     *L1PoolManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L1PoolManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1PoolManagerRaw struct {
	Contract *L1PoolManager // Generic contract binding to access the raw methods on
}

// L1PoolManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1PoolManagerCallerRaw struct {
	Contract *L1PoolManagerCaller // Generic read-only contract binding to access the raw methods on
}

// L1PoolManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1PoolManagerTransactorRaw struct {
	Contract *L1PoolManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1PoolManager creates a new instance of L1PoolManager, bound to a specific deployed contract.
func NewL1PoolManager(address common.Address, backend bind.ContractBackend) (*L1PoolManager, error) {
	contract, err := bindL1PoolManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1PoolManager{L1PoolManagerCaller: L1PoolManagerCaller{contract: contract}, L1PoolManagerTransactor: L1PoolManagerTransactor{contract: contract}, L1PoolManagerFilterer: L1PoolManagerFilterer{contract: contract}}, nil
}

// NewL1PoolManagerCaller creates a new read-only instance of L1PoolManager, bound to a specific deployed contract.
func NewL1PoolManagerCaller(address common.Address, caller bind.ContractCaller) (*L1PoolManagerCaller, error) {
	contract, err := bindL1PoolManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerCaller{contract: contract}, nil
}

// NewL1PoolManagerTransactor creates a new write-only instance of L1PoolManager, bound to a specific deployed contract.
func NewL1PoolManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*L1PoolManagerTransactor, error) {
	contract, err := bindL1PoolManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerTransactor{contract: contract}, nil
}

// NewL1PoolManagerFilterer creates a new log filterer instance of L1PoolManager, bound to a specific deployed contract.
func NewL1PoolManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*L1PoolManagerFilterer, error) {
	contract, err := bindL1PoolManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerFilterer{contract: contract}, nil
}

// bindL1PoolManager binds a generic wrapper to an already deployed contract.
func bindL1PoolManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1PoolManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1PoolManager *L1PoolManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1PoolManager.Contract.L1PoolManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1PoolManager *L1PoolManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PoolManager.Contract.L1PoolManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1PoolManager *L1PoolManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1PoolManager.Contract.L1PoolManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1PoolManager *L1PoolManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1PoolManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1PoolManager *L1PoolManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PoolManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1PoolManager *L1PoolManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1PoolManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L1PoolManager *L1PoolManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L1PoolManager *L1PoolManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _L1PoolManager.Contract.DEFAULTADMINROLE(&_L1PoolManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L1PoolManager *L1PoolManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _L1PoolManager.Contract.DEFAULTADMINROLE(&_L1PoolManager.CallOpts)
}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) FeePoolValue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "FeePoolValue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) FeePoolValue(arg0 common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.FeePoolValue(&_L1PoolManager.CallOpts, arg0)
}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) FeePoolValue(arg0 common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.FeePoolValue(&_L1PoolManager.CallOpts, arg0)
}

// FundingPoolBalance is a free data retrieval call binding the contract method 0xdbb0481f.
//
// Solidity: function FundingPoolBalance(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) FundingPoolBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "FundingPoolBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingPoolBalance is a free data retrieval call binding the contract method 0xdbb0481f.
//
// Solidity: function FundingPoolBalance(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) FundingPoolBalance(arg0 common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.FundingPoolBalance(&_L1PoolManager.CallOpts, arg0)
}

// FundingPoolBalance is a free data retrieval call binding the contract method 0xdbb0481f.
//
// Solidity: function FundingPoolBalance(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) FundingPoolBalance(arg0 common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.FundingPoolBalance(&_L1PoolManager.CallOpts, arg0)
}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_L1PoolManager *L1PoolManagerCaller) IsSupportChainId(opts *bind.CallOpts, chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "IsSupportChainId", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_L1PoolManager *L1PoolManagerSession) IsSupportChainId(chainId *big.Int) (bool, error) {
	return _L1PoolManager.Contract.IsSupportChainId(&_L1PoolManager.CallOpts, chainId)
}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_L1PoolManager *L1PoolManagerCallerSession) IsSupportChainId(chainId *big.Int) (bool, error) {
	return _L1PoolManager.Contract.IsSupportChainId(&_L1PoolManager.CallOpts, chainId)
}

// IsSupportToken is a free data retrieval call binding the contract method 0xdac29568.
//
// Solidity: function IsSupportToken(address ) view returns(bool)
func (_L1PoolManager *L1PoolManagerCaller) IsSupportToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "IsSupportToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportToken is a free data retrieval call binding the contract method 0xdac29568.
//
// Solidity: function IsSupportToken(address ) view returns(bool)
func (_L1PoolManager *L1PoolManagerSession) IsSupportToken(arg0 common.Address) (bool, error) {
	return _L1PoolManager.Contract.IsSupportToken(&_L1PoolManager.CallOpts, arg0)
}

// IsSupportToken is a free data retrieval call binding the contract method 0xdac29568.
//
// Solidity: function IsSupportToken(address ) view returns(bool)
func (_L1PoolManager *L1PoolManagerCallerSession) IsSupportToken(arg0 common.Address) (bool, error) {
	return _L1PoolManager.Contract.IsSupportToken(&_L1PoolManager.CallOpts, arg0)
}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_L1PoolManager *L1PoolManagerCaller) L2WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "L2WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_L1PoolManager *L1PoolManagerSession) L2WETH() (common.Address, error) {
	return _L1PoolManager.Contract.L2WETH(&_L1PoolManager.CallOpts)
}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_L1PoolManager *L1PoolManagerCallerSession) L2WETH() (common.Address, error) {
	return _L1PoolManager.Contract.L2WETH(&_L1PoolManager.CallOpts)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0x1ca2f173.
//
// Solidity: function MinStakeAmount(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) MinStakeAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "MinStakeAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeAmount is a free data retrieval call binding the contract method 0x1ca2f173.
//
// Solidity: function MinStakeAmount(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) MinStakeAmount(arg0 common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.MinStakeAmount(&_L1PoolManager.CallOpts, arg0)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0x1ca2f173.
//
// Solidity: function MinStakeAmount(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) MinStakeAmount(arg0 common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.MinStakeAmount(&_L1PoolManager.CallOpts, arg0)
}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) MinTransferAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "MinTransferAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) MinTransferAmount() (*big.Int, error) {
	return _L1PoolManager.Contract.MinTransferAmount(&_L1PoolManager.CallOpts)
}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) MinTransferAmount() (*big.Int, error) {
	return _L1PoolManager.Contract.MinTransferAmount(&_L1PoolManager.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) PerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "PerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) PerFee() (*big.Int, error) {
	return _L1PoolManager.Contract.PerFee(&_L1PoolManager.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) PerFee() (*big.Int, error) {
	return _L1PoolManager.Contract.PerFee(&_L1PoolManager.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0x4663cdc8.
//
// Solidity: function Pools(address , uint256 ) view returns(uint32 startTimestamp, uint32 endTimestamp, address token, uint256 TotalAmount, uint256 TotalFee, uint256 TotalFeeClaimed, bool IsCompleted)
func (_L1PoolManager *L1PoolManagerCaller) Pools(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StartTimestamp  uint32
	EndTimestamp    uint32
	Token           common.Address
	TotalAmount     *big.Int
	TotalFee        *big.Int
	TotalFeeClaimed *big.Int
	IsCompleted     bool
}, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "Pools", arg0, arg1)

	outstruct := new(struct {
		StartTimestamp  uint32
		EndTimestamp    uint32
		Token           common.Address
		TotalAmount     *big.Int
		TotalFee        *big.Int
		TotalFeeClaimed *big.Int
		IsCompleted     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTimestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.EndTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Token = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TotalAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalFee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalFeeClaimed = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.IsCompleted = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Pools is a free data retrieval call binding the contract method 0x4663cdc8.
//
// Solidity: function Pools(address , uint256 ) view returns(uint32 startTimestamp, uint32 endTimestamp, address token, uint256 TotalAmount, uint256 TotalFee, uint256 TotalFeeClaimed, bool IsCompleted)
func (_L1PoolManager *L1PoolManagerSession) Pools(arg0 common.Address, arg1 *big.Int) (struct {
	StartTimestamp  uint32
	EndTimestamp    uint32
	Token           common.Address
	TotalAmount     *big.Int
	TotalFee        *big.Int
	TotalFeeClaimed *big.Int
	IsCompleted     bool
}, error) {
	return _L1PoolManager.Contract.Pools(&_L1PoolManager.CallOpts, arg0, arg1)
}

// Pools is a free data retrieval call binding the contract method 0x4663cdc8.
//
// Solidity: function Pools(address , uint256 ) view returns(uint32 startTimestamp, uint32 endTimestamp, address token, uint256 TotalAmount, uint256 TotalFee, uint256 TotalFeeClaimed, bool IsCompleted)
func (_L1PoolManager *L1PoolManagerCallerSession) Pools(arg0 common.Address, arg1 *big.Int) (struct {
	StartTimestamp  uint32
	EndTimestamp    uint32
	Token           common.Address
	TotalAmount     *big.Int
	TotalFee        *big.Int
	TotalFeeClaimed *big.Int
	IsCompleted     bool
}, error) {
	return _L1PoolManager.Contract.Pools(&_L1PoolManager.CallOpts, arg0, arg1)
}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_L1PoolManager *L1PoolManagerCaller) ReLayer(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "ReLayer")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_L1PoolManager *L1PoolManagerSession) ReLayer() ([32]byte, error) {
	return _L1PoolManager.Contract.ReLayer(&_L1PoolManager.CallOpts)
}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_L1PoolManager *L1PoolManagerCallerSession) ReLayer() ([32]byte, error) {
	return _L1PoolManager.Contract.ReLayer(&_L1PoolManager.CallOpts)
}

// SupportTokens is a free data retrieval call binding the contract method 0x626417b5.
//
// Solidity: function SupportTokens(uint256 ) view returns(address)
func (_L1PoolManager *L1PoolManagerCaller) SupportTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "SupportTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportTokens is a free data retrieval call binding the contract method 0x626417b5.
//
// Solidity: function SupportTokens(uint256 ) view returns(address)
func (_L1PoolManager *L1PoolManagerSession) SupportTokens(arg0 *big.Int) (common.Address, error) {
	return _L1PoolManager.Contract.SupportTokens(&_L1PoolManager.CallOpts, arg0)
}

// SupportTokens is a free data retrieval call binding the contract method 0x626417b5.
//
// Solidity: function SupportTokens(uint256 ) view returns(address)
func (_L1PoolManager *L1PoolManagerCallerSession) SupportTokens(arg0 *big.Int) (common.Address, error) {
	return _L1PoolManager.Contract.SupportTokens(&_L1PoolManager.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x13e8544e.
//
// Solidity: function Users(address , uint256 ) view returns(bool isWithdrawed, address token, uint256 StartPoolId, uint256 EndPoolId, uint256 Amount)
func (_L1PoolManager *L1PoolManagerCaller) Users(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	IsWithdrawed bool
	Token        common.Address
	StartPoolId  *big.Int
	EndPoolId    *big.Int
	Amount       *big.Int
}, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "Users", arg0, arg1)

	outstruct := new(struct {
		IsWithdrawed bool
		Token        common.Address
		StartPoolId  *big.Int
		EndPoolId    *big.Int
		Amount       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsWithdrawed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StartPoolId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndPoolId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0x13e8544e.
//
// Solidity: function Users(address , uint256 ) view returns(bool isWithdrawed, address token, uint256 StartPoolId, uint256 EndPoolId, uint256 Amount)
func (_L1PoolManager *L1PoolManagerSession) Users(arg0 common.Address, arg1 *big.Int) (struct {
	IsWithdrawed bool
	Token        common.Address
	StartPoolId  *big.Int
	EndPoolId    *big.Int
	Amount       *big.Int
}, error) {
	return _L1PoolManager.Contract.Users(&_L1PoolManager.CallOpts, arg0, arg1)
}

// Users is a free data retrieval call binding the contract method 0x13e8544e.
//
// Solidity: function Users(address , uint256 ) view returns(bool isWithdrawed, address token, uint256 StartPoolId, uint256 EndPoolId, uint256 Amount)
func (_L1PoolManager *L1PoolManagerCallerSession) Users(arg0 common.Address, arg1 *big.Int) (struct {
	IsWithdrawed bool
	Token        common.Address
	StartPoolId  *big.Int
	EndPoolId    *big.Int
	Amount       *big.Int
}, error) {
	return _L1PoolManager.Contract.Users(&_L1PoolManager.CallOpts, arg0, arg1)
}

// GetPool is a free data retrieval call binding the contract method 0x5b5b9ea2.
//
// Solidity: function getPool(address _token, uint256 _index) view returns((uint32,uint32,address,uint256,uint256,uint256,bool))
func (_L1PoolManager *L1PoolManagerCaller) GetPool(opts *bind.CallOpts, _token common.Address, _index *big.Int) (IL1PoolManagerPool, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getPool", _token, _index)

	if err != nil {
		return *new(IL1PoolManagerPool), err
	}

	out0 := *abi.ConvertType(out[0], new(IL1PoolManagerPool)).(*IL1PoolManagerPool)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x5b5b9ea2.
//
// Solidity: function getPool(address _token, uint256 _index) view returns((uint32,uint32,address,uint256,uint256,uint256,bool))
func (_L1PoolManager *L1PoolManagerSession) GetPool(_token common.Address, _index *big.Int) (IL1PoolManagerPool, error) {
	return _L1PoolManager.Contract.GetPool(&_L1PoolManager.CallOpts, _token, _index)
}

// GetPool is a free data retrieval call binding the contract method 0x5b5b9ea2.
//
// Solidity: function getPool(address _token, uint256 _index) view returns((uint32,uint32,address,uint256,uint256,uint256,bool))
func (_L1PoolManager *L1PoolManagerCallerSession) GetPool(_token common.Address, _index *big.Int) (IL1PoolManagerPool, error) {
	return _L1PoolManager.Contract.GetPool(&_L1PoolManager.CallOpts, _token, _index)
}

// GetPoolLength is a free data retrieval call binding the contract method 0xf8fcc2aa.
//
// Solidity: function getPoolLength(address _token) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) GetPoolLength(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getPoolLength", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolLength is a free data retrieval call binding the contract method 0xf8fcc2aa.
//
// Solidity: function getPoolLength(address _token) view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) GetPoolLength(_token common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.GetPoolLength(&_L1PoolManager.CallOpts, _token)
}

// GetPoolLength is a free data retrieval call binding the contract method 0xf8fcc2aa.
//
// Solidity: function getPoolLength(address _token) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) GetPoolLength(_token common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.GetPoolLength(&_L1PoolManager.CallOpts, _token)
}

// GetPrincipal is a free data retrieval call binding the contract method 0xa69bdf16.
//
// Solidity: function getPrincipal() view returns((address,uint256)[])
func (_L1PoolManager *L1PoolManagerCaller) GetPrincipal(opts *bind.CallOpts) ([]IL1PoolManagerKeyValuePair, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getPrincipal")

	if err != nil {
		return *new([]IL1PoolManagerKeyValuePair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IL1PoolManagerKeyValuePair)).(*[]IL1PoolManagerKeyValuePair)

	return out0, err

}

// GetPrincipal is a free data retrieval call binding the contract method 0xa69bdf16.
//
// Solidity: function getPrincipal() view returns((address,uint256)[])
func (_L1PoolManager *L1PoolManagerSession) GetPrincipal() ([]IL1PoolManagerKeyValuePair, error) {
	return _L1PoolManager.Contract.GetPrincipal(&_L1PoolManager.CallOpts)
}

// GetPrincipal is a free data retrieval call binding the contract method 0xa69bdf16.
//
// Solidity: function getPrincipal() view returns((address,uint256)[])
func (_L1PoolManager *L1PoolManagerCallerSession) GetPrincipal() ([]IL1PoolManagerKeyValuePair, error) {
	return _L1PoolManager.Contract.GetPrincipal(&_L1PoolManager.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0x3d18b912.
//
// Solidity: function getReward() view returns((address,uint256)[])
func (_L1PoolManager *L1PoolManagerCaller) GetReward(opts *bind.CallOpts) ([]IL1PoolManagerKeyValuePair, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getReward")

	if err != nil {
		return *new([]IL1PoolManagerKeyValuePair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IL1PoolManagerKeyValuePair)).(*[]IL1PoolManagerKeyValuePair)

	return out0, err

}

// GetReward is a free data retrieval call binding the contract method 0x3d18b912.
//
// Solidity: function getReward() view returns((address,uint256)[])
func (_L1PoolManager *L1PoolManagerSession) GetReward() ([]IL1PoolManagerKeyValuePair, error) {
	return _L1PoolManager.Contract.GetReward(&_L1PoolManager.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0x3d18b912.
//
// Solidity: function getReward() view returns((address,uint256)[])
func (_L1PoolManager *L1PoolManagerCallerSession) GetReward() ([]IL1PoolManagerKeyValuePair, error) {
	return _L1PoolManager.Contract.GetReward(&_L1PoolManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L1PoolManager *L1PoolManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L1PoolManager *L1PoolManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _L1PoolManager.Contract.GetRoleAdmin(&_L1PoolManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L1PoolManager *L1PoolManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _L1PoolManager.Contract.GetRoleAdmin(&_L1PoolManager.CallOpts, role)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _user) view returns((bool,address,uint256,uint256,uint256)[])
func (_L1PoolManager *L1PoolManagerCaller) GetUser(opts *bind.CallOpts, _user common.Address) ([]IL1PoolManagerUser, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getUser", _user)

	if err != nil {
		return *new([]IL1PoolManagerUser), err
	}

	out0 := *abi.ConvertType(out[0], new([]IL1PoolManagerUser)).(*[]IL1PoolManagerUser)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _user) view returns((bool,address,uint256,uint256,uint256)[])
func (_L1PoolManager *L1PoolManagerSession) GetUser(_user common.Address) ([]IL1PoolManagerUser, error) {
	return _L1PoolManager.Contract.GetUser(&_L1PoolManager.CallOpts, _user)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _user) view returns((bool,address,uint256,uint256,uint256)[])
func (_L1PoolManager *L1PoolManagerCallerSession) GetUser(_user common.Address) ([]IL1PoolManagerUser, error) {
	return _L1PoolManager.Contract.GetUser(&_L1PoolManager.CallOpts, _user)
}

// GetUser0 is a free data retrieval call binding the contract method 0xff2bf64f.
//
// Solidity: function getUser(address _user, uint256 _index) view returns((bool,address,uint256,uint256,uint256))
func (_L1PoolManager *L1PoolManagerCaller) GetUser0(opts *bind.CallOpts, _user common.Address, _index *big.Int) (IL1PoolManagerUser, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getUser0", _user, _index)

	if err != nil {
		return *new(IL1PoolManagerUser), err
	}

	out0 := *abi.ConvertType(out[0], new(IL1PoolManagerUser)).(*IL1PoolManagerUser)

	return out0, err

}

// GetUser0 is a free data retrieval call binding the contract method 0xff2bf64f.
//
// Solidity: function getUser(address _user, uint256 _index) view returns((bool,address,uint256,uint256,uint256))
func (_L1PoolManager *L1PoolManagerSession) GetUser0(_user common.Address, _index *big.Int) (IL1PoolManagerUser, error) {
	return _L1PoolManager.Contract.GetUser0(&_L1PoolManager.CallOpts, _user, _index)
}

// GetUser0 is a free data retrieval call binding the contract method 0xff2bf64f.
//
// Solidity: function getUser(address _user, uint256 _index) view returns((bool,address,uint256,uint256,uint256))
func (_L1PoolManager *L1PoolManagerCallerSession) GetUser0(_user common.Address, _index *big.Int) (IL1PoolManagerUser, error) {
	return _L1PoolManager.Contract.GetUser0(&_L1PoolManager.CallOpts, _user, _index)
}

// GetUserLength is a free data retrieval call binding the contract method 0xcb4f04ad.
//
// Solidity: function getUserLength(address _user) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) GetUserLength(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getUserLength", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLength is a free data retrieval call binding the contract method 0xcb4f04ad.
//
// Solidity: function getUserLength(address _user) view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) GetUserLength(_user common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.GetUserLength(&_L1PoolManager.CallOpts, _user)
}

// GetUserLength is a free data retrieval call binding the contract method 0xcb4f04ad.
//
// Solidity: function getUserLength(address _user) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) GetUserLength(_user common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.GetUserLength(&_L1PoolManager.CallOpts, _user)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L1PoolManager *L1PoolManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L1PoolManager *L1PoolManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _L1PoolManager.Contract.HasRole(&_L1PoolManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L1PoolManager *L1PoolManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _L1PoolManager.Contract.HasRole(&_L1PoolManager.CallOpts, role, account)
}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_L1PoolManager *L1PoolManagerCaller) MessageManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "messageManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_L1PoolManager *L1PoolManagerSession) MessageManager() (common.Address, error) {
	return _L1PoolManager.Contract.MessageManager(&_L1PoolManager.CallOpts)
}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_L1PoolManager *L1PoolManagerCallerSession) MessageManager() (common.Address, error) {
	return _L1PoolManager.Contract.MessageManager(&_L1PoolManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1PoolManager *L1PoolManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1PoolManager *L1PoolManagerSession) Paused() (bool, error) {
	return _L1PoolManager.Contract.Paused(&_L1PoolManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1PoolManager *L1PoolManagerCallerSession) Paused() (bool, error) {
	return _L1PoolManager.Contract.Paused(&_L1PoolManager.CallOpts)
}

// PeriodTime is a free data retrieval call binding the contract method 0x1d31fac0.
//
// Solidity: function periodTime() view returns(uint32)
func (_L1PoolManager *L1PoolManagerCaller) PeriodTime(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "periodTime")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PeriodTime is a free data retrieval call binding the contract method 0x1d31fac0.
//
// Solidity: function periodTime() view returns(uint32)
func (_L1PoolManager *L1PoolManagerSession) PeriodTime() (uint32, error) {
	return _L1PoolManager.Contract.PeriodTime(&_L1PoolManager.CallOpts)
}

// PeriodTime is a free data retrieval call binding the contract method 0x1d31fac0.
//
// Solidity: function periodTime() view returns(uint32)
func (_L1PoolManager *L1PoolManagerCallerSession) PeriodTime() (uint32, error) {
	return _L1PoolManager.Contract.PeriodTime(&_L1PoolManager.CallOpts)
}

// StakingMessageNumber is a free data retrieval call binding the contract method 0x31d98b3b.
//
// Solidity: function stakingMessageNumber() view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) StakingMessageNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "stakingMessageNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingMessageNumber is a free data retrieval call binding the contract method 0x31d98b3b.
//
// Solidity: function stakingMessageNumber() view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) StakingMessageNumber() (*big.Int, error) {
	return _L1PoolManager.Contract.StakingMessageNumber(&_L1PoolManager.CallOpts)
}

// StakingMessageNumber is a free data retrieval call binding the contract method 0x31d98b3b.
//
// Solidity: function stakingMessageNumber() view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) StakingMessageNumber() (*big.Int, error) {
	return _L1PoolManager.Contract.StakingMessageNumber(&_L1PoolManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1PoolManager *L1PoolManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1PoolManager *L1PoolManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L1PoolManager.Contract.SupportsInterface(&_L1PoolManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1PoolManager *L1PoolManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L1PoolManager.Contract.SupportsInterface(&_L1PoolManager.CallOpts, interfaceId)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x92d84738.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactor) BridgeFinalizeERC20(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeFinalizeERC20", sourceChainId, destChainId, to, ERC20Address, amount, _fee, _nonce)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x92d84738.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L1PoolManager *L1PoolManagerSession) BridgeFinalizeERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeERC20(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, amount, _fee, _nonce)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x92d84738.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeFinalizeERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeERC20(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x38731c47.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_L1PoolManager *L1PoolManagerTransactor) BridgeFinalizeETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeFinalizeETH", sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x38731c47.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_L1PoolManager *L1PoolManagerSession) BridgeFinalizeETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x38731c47.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeFinalizeETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeETHForStaking is a paid mutator transaction binding the contract method 0xf363e52d.
//
// Solidity: function BridgeFinalizeETHForStaking(uint256 amount, address stakingManager, (address,uint256)[] batcher) returns()
func (_L1PoolManager *L1PoolManagerTransactor) BridgeFinalizeETHForStaking(opts *bind.TransactOpts, amount *big.Int, stakingManager common.Address, batcher []IDETHBatchMint) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeFinalizeETHForStaking", amount, stakingManager, batcher)
}

// BridgeFinalizeETHForStaking is a paid mutator transaction binding the contract method 0xf363e52d.
//
// Solidity: function BridgeFinalizeETHForStaking(uint256 amount, address stakingManager, (address,uint256)[] batcher) returns()
func (_L1PoolManager *L1PoolManagerSession) BridgeFinalizeETHForStaking(amount *big.Int, stakingManager common.Address, batcher []IDETHBatchMint) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeETHForStaking(&_L1PoolManager.TransactOpts, amount, stakingManager, batcher)
}

// BridgeFinalizeETHForStaking is a paid mutator transaction binding the contract method 0xf363e52d.
//
// Solidity: function BridgeFinalizeETHForStaking(uint256 amount, address stakingManager, (address,uint256)[] batcher) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeFinalizeETHForStaking(amount *big.Int, stakingManager common.Address, batcher []IDETHBatchMint) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeETHForStaking(&_L1PoolManager.TransactOpts, amount, stakingManager, batcher)
}

// BridgeFinalizeStakingMessage is a paid mutator transaction binding the contract method 0x415ade1f.
//
// Solidity: function BridgeFinalizeStakingMessage(address shareAddress, address from, address to, uint256 shares, uint256 stakeMessageNonce, uint256 gasLimit) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactor) BridgeFinalizeStakingMessage(opts *bind.TransactOpts, shareAddress common.Address, from common.Address, to common.Address, shares *big.Int, stakeMessageNonce *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeFinalizeStakingMessage", shareAddress, from, to, shares, stakeMessageNonce, gasLimit)
}

// BridgeFinalizeStakingMessage is a paid mutator transaction binding the contract method 0x415ade1f.
//
// Solidity: function BridgeFinalizeStakingMessage(address shareAddress, address from, address to, uint256 shares, uint256 stakeMessageNonce, uint256 gasLimit) returns(bool)
func (_L1PoolManager *L1PoolManagerSession) BridgeFinalizeStakingMessage(shareAddress common.Address, from common.Address, to common.Address, shares *big.Int, stakeMessageNonce *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeStakingMessage(&_L1PoolManager.TransactOpts, shareAddress, from, to, shares, stakeMessageNonce, gasLimit)
}

// BridgeFinalizeStakingMessage is a paid mutator transaction binding the contract method 0x415ade1f.
//
// Solidity: function BridgeFinalizeStakingMessage(address shareAddress, address from, address to, uint256 shares, uint256 stakeMessageNonce, uint256 gasLimit) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeFinalizeStakingMessage(shareAddress common.Address, from common.Address, to common.Address, shares *big.Int, stakeMessageNonce *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeStakingMessage(&_L1PoolManager.TransactOpts, shareAddress, from, to, shares, stakeMessageNonce, gasLimit)
}

// BridgeFinalizeWETH is a paid mutator transaction binding the contract method 0x5765a4eb.
//
// Solidity: function BridgeFinalizeWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactor) BridgeFinalizeWETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeFinalizeWETH", sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeWETH is a paid mutator transaction binding the contract method 0x5765a4eb.
//
// Solidity: function BridgeFinalizeWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L1PoolManager *L1PoolManagerSession) BridgeFinalizeWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeWETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeWETH is a paid mutator transaction binding the contract method 0x5765a4eb.
//
// Solidity: function BridgeFinalizeWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeFinalizeWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeWETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x3b0230f0.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactor) BridgeInitiateERC20(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeInitiateERC20", sourceChainId, destChainId, to, ERC20Address, value)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x3b0230f0.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_L1PoolManager *L1PoolManagerSession) BridgeInitiateERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeInitiateERC20(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, value)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x3b0230f0.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeInitiateERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeInitiateERC20(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, value)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x9b442380.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address to) payable returns(bool)
func (_L1PoolManager *L1PoolManagerTransactor) BridgeInitiateETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeInitiateETH", sourceChainId, destChainId, to)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x9b442380.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address to) payable returns(bool)
func (_L1PoolManager *L1PoolManagerSession) BridgeInitiateETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeInitiateETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x9b442380.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address to) payable returns(bool)
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeInitiateETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeInitiateETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to)
}

// BridgeInitiateStakingMessage is a paid mutator transaction binding the contract method 0xe1c9e844.
//
// Solidity: function BridgeInitiateStakingMessage(address from, address to, uint256 shares) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactor) BridgeInitiateStakingMessage(opts *bind.TransactOpts, from common.Address, to common.Address, shares *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeInitiateStakingMessage", from, to, shares)
}

// BridgeInitiateStakingMessage is a paid mutator transaction binding the contract method 0xe1c9e844.
//
// Solidity: function BridgeInitiateStakingMessage(address from, address to, uint256 shares) returns(bool)
func (_L1PoolManager *L1PoolManagerSession) BridgeInitiateStakingMessage(from common.Address, to common.Address, shares *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeInitiateStakingMessage(&_L1PoolManager.TransactOpts, from, to, shares)
}

// BridgeInitiateStakingMessage is a paid mutator transaction binding the contract method 0xe1c9e844.
//
// Solidity: function BridgeInitiateStakingMessage(address from, address to, uint256 shares) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeInitiateStakingMessage(from common.Address, to common.Address, shares *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeInitiateStakingMessage(&_L1PoolManager.TransactOpts, from, to, shares)
}

// BridgeInitiateWETH is a paid mutator transaction binding the contract method 0xcd01c665.
//
// Solidity: function BridgeInitiateWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 value) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactor) BridgeInitiateWETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeInitiateWETH", sourceChainId, destChainId, to, value)
}

// BridgeInitiateWETH is a paid mutator transaction binding the contract method 0xcd01c665.
//
// Solidity: function BridgeInitiateWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 value) returns(bool)
func (_L1PoolManager *L1PoolManagerSession) BridgeInitiateWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeInitiateWETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, value)
}

// BridgeInitiateWETH is a paid mutator transaction binding the contract method 0xcd01c665.
//
// Solidity: function BridgeInitiateWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 value) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeInitiateWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeInitiateWETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, value)
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x2c37388b.
//
// Solidity: function ClaimAllReward() returns()
func (_L1PoolManager *L1PoolManagerTransactor) ClaimAllReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "ClaimAllReward")
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x2c37388b.
//
// Solidity: function ClaimAllReward() returns()
func (_L1PoolManager *L1PoolManagerSession) ClaimAllReward() (*types.Transaction, error) {
	return _L1PoolManager.Contract.ClaimAllReward(&_L1PoolManager.TransactOpts)
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x2c37388b.
//
// Solidity: function ClaimAllReward() returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) ClaimAllReward() (*types.Transaction, error) {
	return _L1PoolManager.Contract.ClaimAllReward(&_L1PoolManager.TransactOpts)
}

// ClaimbyID is a paid mutator transaction binding the contract method 0x34d34ef5.
//
// Solidity: function ClaimbyID(uint256 i) returns()
func (_L1PoolManager *L1PoolManagerTransactor) ClaimbyID(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "ClaimbyID", i)
}

// ClaimbyID is a paid mutator transaction binding the contract method 0x34d34ef5.
//
// Solidity: function ClaimbyID(uint256 i) returns()
func (_L1PoolManager *L1PoolManagerSession) ClaimbyID(i *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.ClaimbyID(&_L1PoolManager.TransactOpts, i)
}

// ClaimbyID is a paid mutator transaction binding the contract method 0x34d34ef5.
//
// Solidity: function ClaimbyID(uint256 i) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) ClaimbyID(i *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.ClaimbyID(&_L1PoolManager.TransactOpts, i)
}

// CompletePoolAndNew is a paid mutator transaction binding the contract method 0x3338562c.
//
// Solidity: function CompletePoolAndNew((uint32,uint32,address,uint256,uint256,uint256,bool)[] CompletePools) payable returns()
func (_L1PoolManager *L1PoolManagerTransactor) CompletePoolAndNew(opts *bind.TransactOpts, CompletePools []IL1PoolManagerPool) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "CompletePoolAndNew", CompletePools)
}

// CompletePoolAndNew is a paid mutator transaction binding the contract method 0x3338562c.
//
// Solidity: function CompletePoolAndNew((uint32,uint32,address,uint256,uint256,uint256,bool)[] CompletePools) payable returns()
func (_L1PoolManager *L1PoolManagerSession) CompletePoolAndNew(CompletePools []IL1PoolManagerPool) (*types.Transaction, error) {
	return _L1PoolManager.Contract.CompletePoolAndNew(&_L1PoolManager.TransactOpts, CompletePools)
}

// CompletePoolAndNew is a paid mutator transaction binding the contract method 0x3338562c.
//
// Solidity: function CompletePoolAndNew((uint32,uint32,address,uint256,uint256,uint256,bool)[] CompletePools) payable returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) CompletePoolAndNew(CompletePools []IL1PoolManagerPool) (*types.Transaction, error) {
	return _L1PoolManager.Contract.CompletePoolAndNew(&_L1PoolManager.TransactOpts, CompletePools)
}

// DepositAndStaking is a paid mutator transaction binding the contract method 0x9ab0b652.
//
// Solidity: function DepositAndStaking(address _token, uint256 _amount) payable returns()
func (_L1PoolManager *L1PoolManagerTransactor) DepositAndStaking(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "DepositAndStaking", _token, _amount)
}

// DepositAndStaking is a paid mutator transaction binding the contract method 0x9ab0b652.
//
// Solidity: function DepositAndStaking(address _token, uint256 _amount) payable returns()
func (_L1PoolManager *L1PoolManagerSession) DepositAndStaking(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStaking(&_L1PoolManager.TransactOpts, _token, _amount)
}

// DepositAndStaking is a paid mutator transaction binding the contract method 0x9ab0b652.
//
// Solidity: function DepositAndStaking(address _token, uint256 _amount) payable returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) DepositAndStaking(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStaking(&_L1PoolManager.TransactOpts, _token, _amount)
}

// DepositAndStakingERC20 is a paid mutator transaction binding the contract method 0x4dfaef84.
//
// Solidity: function DepositAndStakingERC20(address _token, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactor) DepositAndStakingERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "DepositAndStakingERC20", _token, _amount)
}

// DepositAndStakingERC20 is a paid mutator transaction binding the contract method 0x4dfaef84.
//
// Solidity: function DepositAndStakingERC20(address _token, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerSession) DepositAndStakingERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStakingERC20(&_L1PoolManager.TransactOpts, _token, _amount)
}

// DepositAndStakingERC20 is a paid mutator transaction binding the contract method 0x4dfaef84.
//
// Solidity: function DepositAndStakingERC20(address _token, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) DepositAndStakingERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStakingERC20(&_L1PoolManager.TransactOpts, _token, _amount)
}

// DepositAndStakingETH is a paid mutator transaction binding the contract method 0x96f984e6.
//
// Solidity: function DepositAndStakingETH() payable returns()
func (_L1PoolManager *L1PoolManagerTransactor) DepositAndStakingETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "DepositAndStakingETH")
}

// DepositAndStakingETH is a paid mutator transaction binding the contract method 0x96f984e6.
//
// Solidity: function DepositAndStakingETH() payable returns()
func (_L1PoolManager *L1PoolManagerSession) DepositAndStakingETH() (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStakingETH(&_L1PoolManager.TransactOpts)
}

// DepositAndStakingETH is a paid mutator transaction binding the contract method 0x96f984e6.
//
// Solidity: function DepositAndStakingETH() payable returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) DepositAndStakingETH() (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStakingETH(&_L1PoolManager.TransactOpts)
}

// DepositAndStakingWETH is a paid mutator transaction binding the contract method 0xbc42493d.
//
// Solidity: function DepositAndStakingWETH(uint256 amount) returns()
func (_L1PoolManager *L1PoolManagerTransactor) DepositAndStakingWETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "DepositAndStakingWETH", amount)
}

// DepositAndStakingWETH is a paid mutator transaction binding the contract method 0xbc42493d.
//
// Solidity: function DepositAndStakingWETH(uint256 amount) returns()
func (_L1PoolManager *L1PoolManagerSession) DepositAndStakingWETH(amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStakingWETH(&_L1PoolManager.TransactOpts, amount)
}

// DepositAndStakingWETH is a paid mutator transaction binding the contract method 0xbc42493d.
//
// Solidity: function DepositAndStakingWETH(uint256 amount) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) DepositAndStakingWETH(amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStakingWETH(&_L1PoolManager.TransactOpts, amount)
}

// QuickSendAssertToUser is a paid mutator transaction binding the contract method 0x325e5d43.
//
// Solidity: function QuickSendAssertToUser(address _token, address to, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactor) QuickSendAssertToUser(opts *bind.TransactOpts, _token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "QuickSendAssertToUser", _token, to, _amount)
}

// QuickSendAssertToUser is a paid mutator transaction binding the contract method 0x325e5d43.
//
// Solidity: function QuickSendAssertToUser(address _token, address to, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerSession) QuickSendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.QuickSendAssertToUser(&_L1PoolManager.TransactOpts, _token, to, _amount)
}

// QuickSendAssertToUser is a paid mutator transaction binding the contract method 0x325e5d43.
//
// Solidity: function QuickSendAssertToUser(address _token, address to, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) QuickSendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.QuickSendAssertToUser(&_L1PoolManager.TransactOpts, _token, to, _amount)
}

// TransferAssertToBridge is a paid mutator transaction binding the contract method 0xf91fa9a8.
//
// Solidity: function TransferAssertToBridge(uint256 Blockchain, address _token, address _to, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactor) TransferAssertToBridge(opts *bind.TransactOpts, Blockchain *big.Int, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "TransferAssertToBridge", Blockchain, _token, _to, _amount)
}

// TransferAssertToBridge is a paid mutator transaction binding the contract method 0xf91fa9a8.
//
// Solidity: function TransferAssertToBridge(uint256 Blockchain, address _token, address _to, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerSession) TransferAssertToBridge(Blockchain *big.Int, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.TransferAssertToBridge(&_L1PoolManager.TransactOpts, Blockchain, _token, _to, _amount)
}

// TransferAssertToBridge is a paid mutator transaction binding the contract method 0xf91fa9a8.
//
// Solidity: function TransferAssertToBridge(uint256 Blockchain, address _token, address _to, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) TransferAssertToBridge(Blockchain *big.Int, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.TransferAssertToBridge(&_L1PoolManager.TransactOpts, Blockchain, _token, _to, _amount)
}

// UpdateFundingPoolBalance is a paid mutator transaction binding the contract method 0x0cd727df.
//
// Solidity: function UpdateFundingPoolBalance(address token, uint256 amount) returns()
func (_L1PoolManager *L1PoolManagerTransactor) UpdateFundingPoolBalance(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "UpdateFundingPoolBalance", token, amount)
}

// UpdateFundingPoolBalance is a paid mutator transaction binding the contract method 0x0cd727df.
//
// Solidity: function UpdateFundingPoolBalance(address token, uint256 amount) returns()
func (_L1PoolManager *L1PoolManagerSession) UpdateFundingPoolBalance(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.UpdateFundingPoolBalance(&_L1PoolManager.TransactOpts, token, amount)
}

// UpdateFundingPoolBalance is a paid mutator transaction binding the contract method 0x0cd727df.
//
// Solidity: function UpdateFundingPoolBalance(address token, uint256 amount) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) UpdateFundingPoolBalance(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.UpdateFundingPoolBalance(&_L1PoolManager.TransactOpts, token, amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x5a648bc5.
//
// Solidity: function WithdrawAll() returns()
func (_L1PoolManager *L1PoolManagerTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "WithdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x5a648bc5.
//
// Solidity: function WithdrawAll() returns()
func (_L1PoolManager *L1PoolManagerSession) WithdrawAll() (*types.Transaction, error) {
	return _L1PoolManager.Contract.WithdrawAll(&_L1PoolManager.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x5a648bc5.
//
// Solidity: function WithdrawAll() returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _L1PoolManager.Contract.WithdrawAll(&_L1PoolManager.TransactOpts)
}

// WithdrawByID is a paid mutator transaction binding the contract method 0xf62a89cf.
//
// Solidity: function WithdrawByID(uint256 i) returns()
func (_L1PoolManager *L1PoolManagerTransactor) WithdrawByID(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "WithdrawByID", i)
}

// WithdrawByID is a paid mutator transaction binding the contract method 0xf62a89cf.
//
// Solidity: function WithdrawByID(uint256 i) returns()
func (_L1PoolManager *L1PoolManagerSession) WithdrawByID(i *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.WithdrawByID(&_L1PoolManager.TransactOpts, i)
}

// WithdrawByID is a paid mutator transaction binding the contract method 0xf62a89cf.
//
// Solidity: function WithdrawByID(uint256 i) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) WithdrawByID(i *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.WithdrawByID(&_L1PoolManager.TransactOpts, i)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L1PoolManager *L1PoolManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L1PoolManager *L1PoolManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.GrantRole(&_L1PoolManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.GrantRole(&_L1PoolManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _MultisigWallet, address _messageManager) returns()
func (_L1PoolManager *L1PoolManagerTransactor) Initialize(opts *bind.TransactOpts, _MultisigWallet common.Address, _messageManager common.Address) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "initialize", _MultisigWallet, _messageManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _MultisigWallet, address _messageManager) returns()
func (_L1PoolManager *L1PoolManagerSession) Initialize(_MultisigWallet common.Address, _messageManager common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.Initialize(&_L1PoolManager.TransactOpts, _MultisigWallet, _messageManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _MultisigWallet, address _messageManager) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) Initialize(_MultisigWallet common.Address, _messageManager common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.Initialize(&_L1PoolManager.TransactOpts, _MultisigWallet, _messageManager)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1PoolManager *L1PoolManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1PoolManager *L1PoolManagerSession) Pause() (*types.Transaction, error) {
	return _L1PoolManager.Contract.Pause(&_L1PoolManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _L1PoolManager.Contract.Pause(&_L1PoolManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L1PoolManager *L1PoolManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L1PoolManager *L1PoolManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.RenounceRole(&_L1PoolManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.RenounceRole(&_L1PoolManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L1PoolManager *L1PoolManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L1PoolManager *L1PoolManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.RevokeRole(&_L1PoolManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.RevokeRole(&_L1PoolManager.TransactOpts, role, account)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xcb314fab.
//
// Solidity: function setMinStakeAmount(address _token, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactor) SetMinStakeAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "setMinStakeAmount", _token, _amount)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xcb314fab.
//
// Solidity: function setMinStakeAmount(address _token, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerSession) SetMinStakeAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetMinStakeAmount(&_L1PoolManager.TransactOpts, _token, _amount)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xcb314fab.
//
// Solidity: function setMinStakeAmount(address _token, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) SetMinStakeAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetMinStakeAmount(&_L1PoolManager.TransactOpts, _token, _amount)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_L1PoolManager *L1PoolManagerTransactor) SetMinTransferAmount(opts *bind.TransactOpts, _MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "setMinTransferAmount", _MinTransferAmount)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_L1PoolManager *L1PoolManagerSession) SetMinTransferAmount(_MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetMinTransferAmount(&_L1PoolManager.TransactOpts, _MinTransferAmount)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) SetMinTransferAmount(_MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetMinTransferAmount(&_L1PoolManager.TransactOpts, _MinTransferAmount)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_L1PoolManager *L1PoolManagerTransactor) SetPerFee(opts *bind.TransactOpts, _PerFee *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "setPerFee", _PerFee)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_L1PoolManager *L1PoolManagerSession) SetPerFee(_PerFee *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetPerFee(&_L1PoolManager.TransactOpts, _PerFee)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) SetPerFee(_PerFee *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetPerFee(&_L1PoolManager.TransactOpts, _PerFee)
}

// SetSupportERC20Token is a paid mutator transaction binding the contract method 0xab0f9e19.
//
// Solidity: function setSupportERC20Token(address ERC20Address, bool isValid) returns()
func (_L1PoolManager *L1PoolManagerTransactor) SetSupportERC20Token(opts *bind.TransactOpts, ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "setSupportERC20Token", ERC20Address, isValid)
}

// SetSupportERC20Token is a paid mutator transaction binding the contract method 0xab0f9e19.
//
// Solidity: function setSupportERC20Token(address ERC20Address, bool isValid) returns()
func (_L1PoolManager *L1PoolManagerSession) SetSupportERC20Token(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetSupportERC20Token(&_L1PoolManager.TransactOpts, ERC20Address, isValid)
}

// SetSupportERC20Token is a paid mutator transaction binding the contract method 0xab0f9e19.
//
// Solidity: function setSupportERC20Token(address ERC20Address, bool isValid) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) SetSupportERC20Token(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetSupportERC20Token(&_L1PoolManager.TransactOpts, ERC20Address, isValid)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0xe2070893.
//
// Solidity: function setSupportToken(address _token, bool _isSupport, uint32 startTimes) returns()
func (_L1PoolManager *L1PoolManagerTransactor) SetSupportToken(opts *bind.TransactOpts, _token common.Address, _isSupport bool, startTimes uint32) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "setSupportToken", _token, _isSupport, startTimes)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0xe2070893.
//
// Solidity: function setSupportToken(address _token, bool _isSupport, uint32 startTimes) returns()
func (_L1PoolManager *L1PoolManagerSession) SetSupportToken(_token common.Address, _isSupport bool, startTimes uint32) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetSupportToken(&_L1PoolManager.TransactOpts, _token, _isSupport, startTimes)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0xe2070893.
//
// Solidity: function setSupportToken(address _token, bool _isSupport, uint32 startTimes) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) SetSupportToken(_token common.Address, _isSupport bool, startTimes uint32) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetSupportToken(&_L1PoolManager.TransactOpts, _token, _isSupport, startTimes)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_L1PoolManager *L1PoolManagerTransactor) SetValidChainId(opts *bind.TransactOpts, chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "setValidChainId", chainId, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_L1PoolManager *L1PoolManagerSession) SetValidChainId(chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetValidChainId(&_L1PoolManager.TransactOpts, chainId, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) SetValidChainId(chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetValidChainId(&_L1PoolManager.TransactOpts, chainId, isValid)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1PoolManager *L1PoolManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1PoolManager *L1PoolManagerSession) Unpause() (*types.Transaction, error) {
	return _L1PoolManager.Contract.Unpause(&_L1PoolManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _L1PoolManager.Contract.Unpause(&_L1PoolManager.TransactOpts)
}

// L1PoolManagerBridgeFinalizeETHForStakingEventIterator is returned from FilterBridgeFinalizeETHForStakingEvent and is used to iterate over the raw logs and unpacked data for BridgeFinalizeETHForStakingEvent events raised by the L1PoolManager contract.
type L1PoolManagerBridgeFinalizeETHForStakingEventIterator struct {
	Event *L1PoolManagerBridgeFinalizeETHForStakingEvent // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerBridgeFinalizeETHForStakingEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerBridgeFinalizeETHForStakingEvent)
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
		it.Event = new(L1PoolManagerBridgeFinalizeETHForStakingEvent)
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
func (it *L1PoolManagerBridgeFinalizeETHForStakingEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerBridgeFinalizeETHForStakingEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerBridgeFinalizeETHForStakingEvent represents a BridgeFinalizeETHForStakingEvent event raised by the L1PoolManager contract.
type L1PoolManagerBridgeFinalizeETHForStakingEvent struct {
	Amount         *big.Int
	StakingManager common.Address
	Batcher        []IDETHBatchMint
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalizeETHForStakingEvent is a free log retrieval operation binding the contract event 0x33fa5868f7aec03a4bae17d5ca8d0869adac2237298253519ebf80c25ddf7fe3.
//
// Solidity: event BridgeFinalizeETHForStakingEvent(uint256 amount, address stakingManager, (address,uint256)[] batcher)
func (_L1PoolManager *L1PoolManagerFilterer) FilterBridgeFinalizeETHForStakingEvent(opts *bind.FilterOpts) (*L1PoolManagerBridgeFinalizeETHForStakingEventIterator, error) {

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "BridgeFinalizeETHForStakingEvent")
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerBridgeFinalizeETHForStakingEventIterator{contract: _L1PoolManager.contract, event: "BridgeFinalizeETHForStakingEvent", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalizeETHForStakingEvent is a free log subscription operation binding the contract event 0x33fa5868f7aec03a4bae17d5ca8d0869adac2237298253519ebf80c25ddf7fe3.
//
// Solidity: event BridgeFinalizeETHForStakingEvent(uint256 amount, address stakingManager, (address,uint256)[] batcher)
func (_L1PoolManager *L1PoolManagerFilterer) WatchBridgeFinalizeETHForStakingEvent(opts *bind.WatchOpts, sink chan<- *L1PoolManagerBridgeFinalizeETHForStakingEvent) (event.Subscription, error) {

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "BridgeFinalizeETHForStakingEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerBridgeFinalizeETHForStakingEvent)
				if err := _L1PoolManager.contract.UnpackLog(event, "BridgeFinalizeETHForStakingEvent", log); err != nil {
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

// ParseBridgeFinalizeETHForStakingEvent is a log parse operation binding the contract event 0x33fa5868f7aec03a4bae17d5ca8d0869adac2237298253519ebf80c25ddf7fe3.
//
// Solidity: event BridgeFinalizeETHForStakingEvent(uint256 amount, address stakingManager, (address,uint256)[] batcher)
func (_L1PoolManager *L1PoolManagerFilterer) ParseBridgeFinalizeETHForStakingEvent(log types.Log) (*L1PoolManagerBridgeFinalizeETHForStakingEvent, error) {
	event := new(L1PoolManagerBridgeFinalizeETHForStakingEvent)
	if err := _L1PoolManager.contract.UnpackLog(event, "BridgeFinalizeETHForStakingEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerClaimEventIterator is returned from FilterClaimEvent and is used to iterate over the raw logs and unpacked data for ClaimEvent events raised by the L1PoolManager contract.
type L1PoolManagerClaimEventIterator struct {
	Event *L1PoolManagerClaimEvent // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerClaimEvent)
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
		it.Event = new(L1PoolManagerClaimEvent)
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
func (it *L1PoolManagerClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerClaimEvent represents a ClaimEvent event raised by the L1PoolManager contract.
type L1PoolManagerClaimEvent struct {
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
func (_L1PoolManager *L1PoolManagerFilterer) FilterClaimEvent(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*L1PoolManagerClaimEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "ClaimEvent", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerClaimEventIterator{contract: _L1PoolManager.contract, event: "ClaimEvent", logs: logs, sub: sub}, nil
}

// WatchClaimEvent is a free log subscription operation binding the contract event 0x991c329471ab5230e2301ee30d0d6fba5906e32411d5d154d5a8c278d021a2ab.
//
// Solidity: event ClaimEvent(address indexed user, uint256 startPoolId, uint256 endPoolId, address indexed token, uint256 amount, uint256 fee)
func (_L1PoolManager *L1PoolManagerFilterer) WatchClaimEvent(opts *bind.WatchOpts, sink chan<- *L1PoolManagerClaimEvent, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "ClaimEvent", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerClaimEvent)
				if err := _L1PoolManager.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseClaimEvent(log types.Log) (*L1PoolManagerClaimEvent, error) {
	event := new(L1PoolManagerClaimEvent)
	if err := _L1PoolManager.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerClaimRewardIterator is returned from FilterClaimReward and is used to iterate over the raw logs and unpacked data for ClaimReward events raised by the L1PoolManager contract.
type L1PoolManagerClaimRewardIterator struct {
	Event *L1PoolManagerClaimReward // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerClaimRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerClaimReward)
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
		it.Event = new(L1PoolManagerClaimReward)
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
func (it *L1PoolManagerClaimRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerClaimRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerClaimReward represents a ClaimReward event raised by the L1PoolManager contract.
type L1PoolManagerClaimReward struct {
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
func (_L1PoolManager *L1PoolManagerFilterer) FilterClaimReward(opts *bind.FilterOpts) (*L1PoolManagerClaimRewardIterator, error) {

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "ClaimReward")
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerClaimRewardIterator{contract: _L1PoolManager.contract, event: "ClaimReward", logs: logs, sub: sub}, nil
}

// WatchClaimReward is a free log subscription operation binding the contract event 0x7e6fedb6f824b2112ac0626ced32f6be6324d554cc03efd112aa021ee4d740ac.
//
// Solidity: event ClaimReward(address _user, uint256 startPoolId, uint256 EndPoolId, address _token, uint256 Reward)
func (_L1PoolManager *L1PoolManagerFilterer) WatchClaimReward(opts *bind.WatchOpts, sink chan<- *L1PoolManagerClaimReward) (event.Subscription, error) {

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "ClaimReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerClaimReward)
				if err := _L1PoolManager.contract.UnpackLog(event, "ClaimReward", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseClaimReward(log types.Log) (*L1PoolManagerClaimReward, error) {
	event := new(L1PoolManagerClaimReward)
	if err := _L1PoolManager.contract.UnpackLog(event, "ClaimReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerCompletePoolEventIterator is returned from FilterCompletePoolEvent and is used to iterate over the raw logs and unpacked data for CompletePoolEvent events raised by the L1PoolManager contract.
type L1PoolManagerCompletePoolEventIterator struct {
	Event *L1PoolManagerCompletePoolEvent // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerCompletePoolEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerCompletePoolEvent)
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
		it.Event = new(L1PoolManagerCompletePoolEvent)
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
func (it *L1PoolManagerCompletePoolEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerCompletePoolEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerCompletePoolEvent represents a CompletePoolEvent event raised by the L1PoolManager contract.
type L1PoolManagerCompletePoolEvent struct {
	Token     common.Address
	PoolIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCompletePoolEvent is a free log retrieval operation binding the contract event 0xb6f449f07ceaf55392c9899e0797c6529908ae827c2d498c682e90d42c241167.
//
// Solidity: event CompletePoolEvent(address indexed token, uint256 poolIndex)
func (_L1PoolManager *L1PoolManagerFilterer) FilterCompletePoolEvent(opts *bind.FilterOpts, token []common.Address) (*L1PoolManagerCompletePoolEventIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "CompletePoolEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerCompletePoolEventIterator{contract: _L1PoolManager.contract, event: "CompletePoolEvent", logs: logs, sub: sub}, nil
}

// WatchCompletePoolEvent is a free log subscription operation binding the contract event 0xb6f449f07ceaf55392c9899e0797c6529908ae827c2d498c682e90d42c241167.
//
// Solidity: event CompletePoolEvent(address indexed token, uint256 poolIndex)
func (_L1PoolManager *L1PoolManagerFilterer) WatchCompletePoolEvent(opts *bind.WatchOpts, sink chan<- *L1PoolManagerCompletePoolEvent, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "CompletePoolEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerCompletePoolEvent)
				if err := _L1PoolManager.contract.UnpackLog(event, "CompletePoolEvent", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseCompletePoolEvent(log types.Log) (*L1PoolManagerCompletePoolEvent, error) {
	event := new(L1PoolManagerCompletePoolEvent)
	if err := _L1PoolManager.contract.UnpackLog(event, "CompletePoolEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerFinalizeERC20Iterator is returned from FilterFinalizeERC20 and is used to iterate over the raw logs and unpacked data for FinalizeERC20 events raised by the L1PoolManager contract.
type L1PoolManagerFinalizeERC20Iterator struct {
	Event *L1PoolManagerFinalizeERC20 // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerFinalizeERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerFinalizeERC20)
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
		it.Event = new(L1PoolManagerFinalizeERC20)
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
func (it *L1PoolManagerFinalizeERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerFinalizeERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerFinalizeERC20 represents a FinalizeERC20 event raised by the L1PoolManager contract.
type L1PoolManagerFinalizeERC20 struct {
	SourceChainId *big.Int
	DestChainId   *big.Int
	ERC20Address  common.Address
	From          common.Address
	To            common.Address
	Value         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFinalizeERC20 is a free log retrieval operation binding the contract event 0x0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe.
//
// Solidity: event FinalizeERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) FilterFinalizeERC20(opts *bind.FilterOpts, ERC20Address []common.Address, from []common.Address, to []common.Address) (*L1PoolManagerFinalizeERC20Iterator, error) {

	var ERC20AddressRule []interface{}
	for _, ERC20AddressItem := range ERC20Address {
		ERC20AddressRule = append(ERC20AddressRule, ERC20AddressItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "FinalizeERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerFinalizeERC20Iterator{contract: _L1PoolManager.contract, event: "FinalizeERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeERC20 is a free log subscription operation binding the contract event 0x0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe.
//
// Solidity: event FinalizeERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) WatchFinalizeERC20(opts *bind.WatchOpts, sink chan<- *L1PoolManagerFinalizeERC20, ERC20Address []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var ERC20AddressRule []interface{}
	for _, ERC20AddressItem := range ERC20Address {
		ERC20AddressRule = append(ERC20AddressRule, ERC20AddressItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "FinalizeERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerFinalizeERC20)
				if err := _L1PoolManager.contract.UnpackLog(event, "FinalizeERC20", log); err != nil {
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

// ParseFinalizeERC20 is a log parse operation binding the contract event 0x0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe.
//
// Solidity: event FinalizeERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) ParseFinalizeERC20(log types.Log) (*L1PoolManagerFinalizeERC20, error) {
	event := new(L1PoolManagerFinalizeERC20)
	if err := _L1PoolManager.contract.UnpackLog(event, "FinalizeERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerFinalizeETHIterator is returned from FilterFinalizeETH and is used to iterate over the raw logs and unpacked data for FinalizeETH events raised by the L1PoolManager contract.
type L1PoolManagerFinalizeETHIterator struct {
	Event *L1PoolManagerFinalizeETH // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerFinalizeETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerFinalizeETH)
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
		it.Event = new(L1PoolManagerFinalizeETH)
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
func (it *L1PoolManagerFinalizeETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerFinalizeETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerFinalizeETH represents a FinalizeETH event raised by the L1PoolManager contract.
type L1PoolManagerFinalizeETH struct {
	SourceChainId *big.Int
	DestChainId   *big.Int
	From          common.Address
	To            common.Address
	Value         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFinalizeETH is a free log retrieval operation binding the contract event 0x61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f32.
//
// Solidity: event FinalizeETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) FilterFinalizeETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1PoolManagerFinalizeETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "FinalizeETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerFinalizeETHIterator{contract: _L1PoolManager.contract, event: "FinalizeETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeETH is a free log subscription operation binding the contract event 0x61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f32.
//
// Solidity: event FinalizeETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) WatchFinalizeETH(opts *bind.WatchOpts, sink chan<- *L1PoolManagerFinalizeETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "FinalizeETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerFinalizeETH)
				if err := _L1PoolManager.contract.UnpackLog(event, "FinalizeETH", log); err != nil {
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

// ParseFinalizeETH is a log parse operation binding the contract event 0x61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f32.
//
// Solidity: event FinalizeETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) ParseFinalizeETH(log types.Log) (*L1PoolManagerFinalizeETH, error) {
	event := new(L1PoolManagerFinalizeETH)
	if err := _L1PoolManager.contract.UnpackLog(event, "FinalizeETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerFinalizeStakingMessageIterator is returned from FilterFinalizeStakingMessage and is used to iterate over the raw logs and unpacked data for FinalizeStakingMessage events raised by the L1PoolManager contract.
type L1PoolManagerFinalizeStakingMessageIterator struct {
	Event *L1PoolManagerFinalizeStakingMessage // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerFinalizeStakingMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerFinalizeStakingMessage)
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
		it.Event = new(L1PoolManagerFinalizeStakingMessage)
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
func (it *L1PoolManagerFinalizeStakingMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerFinalizeStakingMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerFinalizeStakingMessage represents a FinalizeStakingMessage event raised by the L1PoolManager contract.
type L1PoolManagerFinalizeStakingMessage struct {
	From              common.Address
	To                common.Address
	ShareAddress      common.Address
	Shares            *big.Int
	StakeMessageNonce *big.Int
	StakeMessageHash  [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFinalizeStakingMessage is a free log retrieval operation binding the contract event 0xac6725dae59a3a2ceb3b61453d910067b411ec59d487e1d72382b9cdb9ea723e.
//
// Solidity: event FinalizeStakingMessage(address indexed from, address indexed to, address shareAddress, uint256 shares, uint256 stakeMessageNonce, bytes32 stakeMessageHash)
func (_L1PoolManager *L1PoolManagerFilterer) FilterFinalizeStakingMessage(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1PoolManagerFinalizeStakingMessageIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "FinalizeStakingMessage", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerFinalizeStakingMessageIterator{contract: _L1PoolManager.contract, event: "FinalizeStakingMessage", logs: logs, sub: sub}, nil
}

// WatchFinalizeStakingMessage is a free log subscription operation binding the contract event 0xac6725dae59a3a2ceb3b61453d910067b411ec59d487e1d72382b9cdb9ea723e.
//
// Solidity: event FinalizeStakingMessage(address indexed from, address indexed to, address shareAddress, uint256 shares, uint256 stakeMessageNonce, bytes32 stakeMessageHash)
func (_L1PoolManager *L1PoolManagerFilterer) WatchFinalizeStakingMessage(opts *bind.WatchOpts, sink chan<- *L1PoolManagerFinalizeStakingMessage, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "FinalizeStakingMessage", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerFinalizeStakingMessage)
				if err := _L1PoolManager.contract.UnpackLog(event, "FinalizeStakingMessage", log); err != nil {
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

// ParseFinalizeStakingMessage is a log parse operation binding the contract event 0xac6725dae59a3a2ceb3b61453d910067b411ec59d487e1d72382b9cdb9ea723e.
//
// Solidity: event FinalizeStakingMessage(address indexed from, address indexed to, address shareAddress, uint256 shares, uint256 stakeMessageNonce, bytes32 stakeMessageHash)
func (_L1PoolManager *L1PoolManagerFilterer) ParseFinalizeStakingMessage(log types.Log) (*L1PoolManagerFinalizeStakingMessage, error) {
	event := new(L1PoolManagerFinalizeStakingMessage)
	if err := _L1PoolManager.contract.UnpackLog(event, "FinalizeStakingMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerFinalizeWETHIterator is returned from FilterFinalizeWETH and is used to iterate over the raw logs and unpacked data for FinalizeWETH events raised by the L1PoolManager contract.
type L1PoolManagerFinalizeWETHIterator struct {
	Event *L1PoolManagerFinalizeWETH // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerFinalizeWETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerFinalizeWETH)
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
		it.Event = new(L1PoolManagerFinalizeWETH)
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
func (it *L1PoolManagerFinalizeWETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerFinalizeWETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerFinalizeWETH represents a FinalizeWETH event raised by the L1PoolManager contract.
type L1PoolManagerFinalizeWETH struct {
	SourceChainId *big.Int
	DestChainId   *big.Int
	From          common.Address
	To            common.Address
	Value         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFinalizeWETH is a free log retrieval operation binding the contract event 0x7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc14.
//
// Solidity: event FinalizeWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) FilterFinalizeWETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1PoolManagerFinalizeWETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "FinalizeWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerFinalizeWETHIterator{contract: _L1PoolManager.contract, event: "FinalizeWETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeWETH is a free log subscription operation binding the contract event 0x7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc14.
//
// Solidity: event FinalizeWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) WatchFinalizeWETH(opts *bind.WatchOpts, sink chan<- *L1PoolManagerFinalizeWETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "FinalizeWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerFinalizeWETH)
				if err := _L1PoolManager.contract.UnpackLog(event, "FinalizeWETH", log); err != nil {
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

// ParseFinalizeWETH is a log parse operation binding the contract event 0x7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc14.
//
// Solidity: event FinalizeWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) ParseFinalizeWETH(log types.Log) (*L1PoolManagerFinalizeWETH, error) {
	event := new(L1PoolManagerFinalizeWETH)
	if err := _L1PoolManager.contract.UnpackLog(event, "FinalizeWETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1PoolManager contract.
type L1PoolManagerInitializedIterator struct {
	Event *L1PoolManagerInitialized // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerInitialized)
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
		it.Event = new(L1PoolManagerInitialized)
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
func (it *L1PoolManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerInitialized represents a Initialized event raised by the L1PoolManager contract.
type L1PoolManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L1PoolManager *L1PoolManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1PoolManagerInitializedIterator, error) {

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerInitializedIterator{contract: _L1PoolManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L1PoolManager *L1PoolManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1PoolManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerInitialized)
				if err := _L1PoolManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseInitialized(log types.Log) (*L1PoolManagerInitialized, error) {
	event := new(L1PoolManagerInitialized)
	if err := _L1PoolManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerInitiateERC20Iterator is returned from FilterInitiateERC20 and is used to iterate over the raw logs and unpacked data for InitiateERC20 events raised by the L1PoolManager contract.
type L1PoolManagerInitiateERC20Iterator struct {
	Event *L1PoolManagerInitiateERC20 // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerInitiateERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerInitiateERC20)
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
		it.Event = new(L1PoolManagerInitiateERC20)
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
func (it *L1PoolManagerInitiateERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerInitiateERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerInitiateERC20 represents a InitiateERC20 event raised by the L1PoolManager contract.
type L1PoolManagerInitiateERC20 struct {
	SourceChainId *big.Int
	DestChainId   *big.Int
	ERC20Address  common.Address
	From          common.Address
	To            common.Address
	Value         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiateERC20 is a free log retrieval operation binding the contract event 0xece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a4.
//
// Solidity: event InitiateERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) FilterInitiateERC20(opts *bind.FilterOpts, ERC20Address []common.Address, from []common.Address, to []common.Address) (*L1PoolManagerInitiateERC20Iterator, error) {

	var ERC20AddressRule []interface{}
	for _, ERC20AddressItem := range ERC20Address {
		ERC20AddressRule = append(ERC20AddressRule, ERC20AddressItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "InitiateERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerInitiateERC20Iterator{contract: _L1PoolManager.contract, event: "InitiateERC20", logs: logs, sub: sub}, nil
}

// WatchInitiateERC20 is a free log subscription operation binding the contract event 0xece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a4.
//
// Solidity: event InitiateERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) WatchInitiateERC20(opts *bind.WatchOpts, sink chan<- *L1PoolManagerInitiateERC20, ERC20Address []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var ERC20AddressRule []interface{}
	for _, ERC20AddressItem := range ERC20Address {
		ERC20AddressRule = append(ERC20AddressRule, ERC20AddressItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "InitiateERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerInitiateERC20)
				if err := _L1PoolManager.contract.UnpackLog(event, "InitiateERC20", log); err != nil {
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

// ParseInitiateERC20 is a log parse operation binding the contract event 0xece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a4.
//
// Solidity: event InitiateERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) ParseInitiateERC20(log types.Log) (*L1PoolManagerInitiateERC20, error) {
	event := new(L1PoolManagerInitiateERC20)
	if err := _L1PoolManager.contract.UnpackLog(event, "InitiateERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerInitiateETHIterator is returned from FilterInitiateETH and is used to iterate over the raw logs and unpacked data for InitiateETH events raised by the L1PoolManager contract.
type L1PoolManagerInitiateETHIterator struct {
	Event *L1PoolManagerInitiateETH // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerInitiateETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerInitiateETH)
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
		it.Event = new(L1PoolManagerInitiateETH)
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
func (it *L1PoolManagerInitiateETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerInitiateETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerInitiateETH represents a InitiateETH event raised by the L1PoolManager contract.
type L1PoolManagerInitiateETH struct {
	SourceChainId *big.Int
	DestChainId   *big.Int
	From          common.Address
	To            common.Address
	Value         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiateETH is a free log retrieval operation binding the contract event 0x0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc5.
//
// Solidity: event InitiateETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) FilterInitiateETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1PoolManagerInitiateETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "InitiateETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerInitiateETHIterator{contract: _L1PoolManager.contract, event: "InitiateETH", logs: logs, sub: sub}, nil
}

// WatchInitiateETH is a free log subscription operation binding the contract event 0x0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc5.
//
// Solidity: event InitiateETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) WatchInitiateETH(opts *bind.WatchOpts, sink chan<- *L1PoolManagerInitiateETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "InitiateETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerInitiateETH)
				if err := _L1PoolManager.contract.UnpackLog(event, "InitiateETH", log); err != nil {
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

// ParseInitiateETH is a log parse operation binding the contract event 0x0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc5.
//
// Solidity: event InitiateETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) ParseInitiateETH(log types.Log) (*L1PoolManagerInitiateETH, error) {
	event := new(L1PoolManagerInitiateETH)
	if err := _L1PoolManager.contract.UnpackLog(event, "InitiateETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerInitiateStakingMessageIterator is returned from FilterInitiateStakingMessage and is used to iterate over the raw logs and unpacked data for InitiateStakingMessage events raised by the L1PoolManager contract.
type L1PoolManagerInitiateStakingMessageIterator struct {
	Event *L1PoolManagerInitiateStakingMessage // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerInitiateStakingMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerInitiateStakingMessage)
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
		it.Event = new(L1PoolManagerInitiateStakingMessage)
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
func (it *L1PoolManagerInitiateStakingMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerInitiateStakingMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerInitiateStakingMessage represents a InitiateStakingMessage event raised by the L1PoolManager contract.
type L1PoolManagerInitiateStakingMessage struct {
	From              common.Address
	To                common.Address
	Shares            *big.Int
	StakeMessageNonce *big.Int
	StakeMessageHash  [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInitiateStakingMessage is a free log retrieval operation binding the contract event 0x5e2d465404712c588a7dffc8622a4fad4f4522708eaf138e2b70910cd36992a0.
//
// Solidity: event InitiateStakingMessage(address indexed from, address indexed to, uint256 shares, uint256 stakeMessageNonce, bytes32 indexed stakeMessageHash)
func (_L1PoolManager *L1PoolManagerFilterer) FilterInitiateStakingMessage(opts *bind.FilterOpts, from []common.Address, to []common.Address, stakeMessageHash [][32]byte) (*L1PoolManagerInitiateStakingMessageIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var stakeMessageHashRule []interface{}
	for _, stakeMessageHashItem := range stakeMessageHash {
		stakeMessageHashRule = append(stakeMessageHashRule, stakeMessageHashItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "InitiateStakingMessage", fromRule, toRule, stakeMessageHashRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerInitiateStakingMessageIterator{contract: _L1PoolManager.contract, event: "InitiateStakingMessage", logs: logs, sub: sub}, nil
}

// WatchInitiateStakingMessage is a free log subscription operation binding the contract event 0x5e2d465404712c588a7dffc8622a4fad4f4522708eaf138e2b70910cd36992a0.
//
// Solidity: event InitiateStakingMessage(address indexed from, address indexed to, uint256 shares, uint256 stakeMessageNonce, bytes32 indexed stakeMessageHash)
func (_L1PoolManager *L1PoolManagerFilterer) WatchInitiateStakingMessage(opts *bind.WatchOpts, sink chan<- *L1PoolManagerInitiateStakingMessage, from []common.Address, to []common.Address, stakeMessageHash [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var stakeMessageHashRule []interface{}
	for _, stakeMessageHashItem := range stakeMessageHash {
		stakeMessageHashRule = append(stakeMessageHashRule, stakeMessageHashItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "InitiateStakingMessage", fromRule, toRule, stakeMessageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerInitiateStakingMessage)
				if err := _L1PoolManager.contract.UnpackLog(event, "InitiateStakingMessage", log); err != nil {
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

// ParseInitiateStakingMessage is a log parse operation binding the contract event 0x5e2d465404712c588a7dffc8622a4fad4f4522708eaf138e2b70910cd36992a0.
//
// Solidity: event InitiateStakingMessage(address indexed from, address indexed to, uint256 shares, uint256 stakeMessageNonce, bytes32 indexed stakeMessageHash)
func (_L1PoolManager *L1PoolManagerFilterer) ParseInitiateStakingMessage(log types.Log) (*L1PoolManagerInitiateStakingMessage, error) {
	event := new(L1PoolManagerInitiateStakingMessage)
	if err := _L1PoolManager.contract.UnpackLog(event, "InitiateStakingMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerInitiateWETHIterator is returned from FilterInitiateWETH and is used to iterate over the raw logs and unpacked data for InitiateWETH events raised by the L1PoolManager contract.
type L1PoolManagerInitiateWETHIterator struct {
	Event *L1PoolManagerInitiateWETH // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerInitiateWETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerInitiateWETH)
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
		it.Event = new(L1PoolManagerInitiateWETH)
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
func (it *L1PoolManagerInitiateWETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerInitiateWETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerInitiateWETH represents a InitiateWETH event raised by the L1PoolManager contract.
type L1PoolManagerInitiateWETH struct {
	SourceChainId *big.Int
	DestChainId   *big.Int
	From          common.Address
	To            common.Address
	Value         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiateWETH is a free log retrieval operation binding the contract event 0xcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a89.
//
// Solidity: event InitiateWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) FilterInitiateWETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1PoolManagerInitiateWETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "InitiateWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerInitiateWETHIterator{contract: _L1PoolManager.contract, event: "InitiateWETH", logs: logs, sub: sub}, nil
}

// WatchInitiateWETH is a free log subscription operation binding the contract event 0xcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a89.
//
// Solidity: event InitiateWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) WatchInitiateWETH(opts *bind.WatchOpts, sink chan<- *L1PoolManagerInitiateWETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "InitiateWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerInitiateWETH)
				if err := _L1PoolManager.contract.UnpackLog(event, "InitiateWETH", log); err != nil {
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

// ParseInitiateWETH is a log parse operation binding the contract event 0xcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a89.
//
// Solidity: event InitiateWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) ParseInitiateWETH(log types.Log) (*L1PoolManagerInitiateWETH, error) {
	event := new(L1PoolManagerInitiateWETH)
	if err := _L1PoolManager.contract.UnpackLog(event, "InitiateWETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the L1PoolManager contract.
type L1PoolManagerPausedIterator struct {
	Event *L1PoolManagerPaused // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerPaused)
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
		it.Event = new(L1PoolManagerPaused)
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
func (it *L1PoolManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerPaused represents a Paused event raised by the L1PoolManager contract.
type L1PoolManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1PoolManager *L1PoolManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*L1PoolManagerPausedIterator, error) {

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerPausedIterator{contract: _L1PoolManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1PoolManager *L1PoolManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *L1PoolManagerPaused) (event.Subscription, error) {

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerPaused)
				if err := _L1PoolManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParsePaused(log types.Log) (*L1PoolManagerPaused, error) {
	event := new(L1PoolManagerPaused)
	if err := _L1PoolManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the L1PoolManager contract.
type L1PoolManagerRoleAdminChangedIterator struct {
	Event *L1PoolManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerRoleAdminChanged)
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
		it.Event = new(L1PoolManagerRoleAdminChanged)
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
func (it *L1PoolManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerRoleAdminChanged represents a RoleAdminChanged event raised by the L1PoolManager contract.
type L1PoolManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_L1PoolManager *L1PoolManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*L1PoolManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerRoleAdminChangedIterator{contract: _L1PoolManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_L1PoolManager *L1PoolManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *L1PoolManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerRoleAdminChanged)
				if err := _L1PoolManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseRoleAdminChanged(log types.Log) (*L1PoolManagerRoleAdminChanged, error) {
	event := new(L1PoolManagerRoleAdminChanged)
	if err := _L1PoolManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the L1PoolManager contract.
type L1PoolManagerRoleGrantedIterator struct {
	Event *L1PoolManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerRoleGranted)
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
		it.Event = new(L1PoolManagerRoleGranted)
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
func (it *L1PoolManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerRoleGranted represents a RoleGranted event raised by the L1PoolManager contract.
type L1PoolManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L1PoolManager *L1PoolManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L1PoolManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerRoleGrantedIterator{contract: _L1PoolManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L1PoolManager *L1PoolManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *L1PoolManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerRoleGranted)
				if err := _L1PoolManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseRoleGranted(log types.Log) (*L1PoolManagerRoleGranted, error) {
	event := new(L1PoolManagerRoleGranted)
	if err := _L1PoolManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the L1PoolManager contract.
type L1PoolManagerRoleRevokedIterator struct {
	Event *L1PoolManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerRoleRevoked)
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
		it.Event = new(L1PoolManagerRoleRevoked)
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
func (it *L1PoolManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerRoleRevoked represents a RoleRevoked event raised by the L1PoolManager contract.
type L1PoolManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L1PoolManager *L1PoolManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L1PoolManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerRoleRevokedIterator{contract: _L1PoolManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L1PoolManager *L1PoolManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *L1PoolManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerRoleRevoked)
				if err := _L1PoolManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseRoleRevoked(log types.Log) (*L1PoolManagerRoleRevoked, error) {
	event := new(L1PoolManagerRoleRevoked)
	if err := _L1PoolManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerSetMinStakeAmountEventIterator is returned from FilterSetMinStakeAmountEvent and is used to iterate over the raw logs and unpacked data for SetMinStakeAmountEvent events raised by the L1PoolManager contract.
type L1PoolManagerSetMinStakeAmountEventIterator struct {
	Event *L1PoolManagerSetMinStakeAmountEvent // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerSetMinStakeAmountEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerSetMinStakeAmountEvent)
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
		it.Event = new(L1PoolManagerSetMinStakeAmountEvent)
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
func (it *L1PoolManagerSetMinStakeAmountEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerSetMinStakeAmountEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerSetMinStakeAmountEvent represents a SetMinStakeAmountEvent event raised by the L1PoolManager contract.
type L1PoolManagerSetMinStakeAmountEvent struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMinStakeAmountEvent is a free log retrieval operation binding the contract event 0xf54d3b756d286b6b08e5d4eda6dfe5b135664abf029e58e637cbf013c442c950.
//
// Solidity: event SetMinStakeAmountEvent(address indexed token, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) FilterSetMinStakeAmountEvent(opts *bind.FilterOpts, token []common.Address) (*L1PoolManagerSetMinStakeAmountEventIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "SetMinStakeAmountEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerSetMinStakeAmountEventIterator{contract: _L1PoolManager.contract, event: "SetMinStakeAmountEvent", logs: logs, sub: sub}, nil
}

// WatchSetMinStakeAmountEvent is a free log subscription operation binding the contract event 0xf54d3b756d286b6b08e5d4eda6dfe5b135664abf029e58e637cbf013c442c950.
//
// Solidity: event SetMinStakeAmountEvent(address indexed token, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) WatchSetMinStakeAmountEvent(opts *bind.WatchOpts, sink chan<- *L1PoolManagerSetMinStakeAmountEvent, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "SetMinStakeAmountEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerSetMinStakeAmountEvent)
				if err := _L1PoolManager.contract.UnpackLog(event, "SetMinStakeAmountEvent", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseSetMinStakeAmountEvent(log types.Log) (*L1PoolManagerSetMinStakeAmountEvent, error) {
	event := new(L1PoolManagerSetMinStakeAmountEvent)
	if err := _L1PoolManager.contract.UnpackLog(event, "SetMinStakeAmountEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerSetSupportTokenEventIterator is returned from FilterSetSupportTokenEvent and is used to iterate over the raw logs and unpacked data for SetSupportTokenEvent events raised by the L1PoolManager contract.
type L1PoolManagerSetSupportTokenEventIterator struct {
	Event *L1PoolManagerSetSupportTokenEvent // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerSetSupportTokenEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerSetSupportTokenEvent)
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
		it.Event = new(L1PoolManagerSetSupportTokenEvent)
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
func (it *L1PoolManagerSetSupportTokenEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerSetSupportTokenEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerSetSupportTokenEvent represents a SetSupportTokenEvent event raised by the L1PoolManager contract.
type L1PoolManagerSetSupportTokenEvent struct {
	Token     common.Address
	IsSupport bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetSupportTokenEvent is a free log retrieval operation binding the contract event 0xc8c34f23fafb34e68119c1d231ef03d0d47225b15e2c4de3efbefa14b0181d86.
//
// Solidity: event SetSupportTokenEvent(address indexed token, bool isSupport)
func (_L1PoolManager *L1PoolManagerFilterer) FilterSetSupportTokenEvent(opts *bind.FilterOpts, token []common.Address) (*L1PoolManagerSetSupportTokenEventIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "SetSupportTokenEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerSetSupportTokenEventIterator{contract: _L1PoolManager.contract, event: "SetSupportTokenEvent", logs: logs, sub: sub}, nil
}

// WatchSetSupportTokenEvent is a free log subscription operation binding the contract event 0xc8c34f23fafb34e68119c1d231ef03d0d47225b15e2c4de3efbefa14b0181d86.
//
// Solidity: event SetSupportTokenEvent(address indexed token, bool isSupport)
func (_L1PoolManager *L1PoolManagerFilterer) WatchSetSupportTokenEvent(opts *bind.WatchOpts, sink chan<- *L1PoolManagerSetSupportTokenEvent, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "SetSupportTokenEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerSetSupportTokenEvent)
				if err := _L1PoolManager.contract.UnpackLog(event, "SetSupportTokenEvent", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseSetSupportTokenEvent(log types.Log) (*L1PoolManagerSetSupportTokenEvent, error) {
	event := new(L1PoolManagerSetSupportTokenEvent)
	if err := _L1PoolManager.contract.UnpackLog(event, "SetSupportTokenEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerStakingETHEventIterator is returned from FilterStakingETHEvent and is used to iterate over the raw logs and unpacked data for StakingETHEvent events raised by the L1PoolManager contract.
type L1PoolManagerStakingETHEventIterator struct {
	Event *L1PoolManagerStakingETHEvent // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerStakingETHEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerStakingETHEvent)
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
		it.Event = new(L1PoolManagerStakingETHEvent)
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
func (it *L1PoolManagerStakingETHEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerStakingETHEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerStakingETHEvent represents a StakingETHEvent event raised by the L1PoolManager contract.
type L1PoolManagerStakingETHEvent struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakingETHEvent is a free log retrieval operation binding the contract event 0xe7466ea83435490635fc76a5f33da4815758ab48b1d45858f0452ca646556937.
//
// Solidity: event StakingETHEvent(address indexed user, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) FilterStakingETHEvent(opts *bind.FilterOpts, user []common.Address) (*L1PoolManagerStakingETHEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "StakingETHEvent", userRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerStakingETHEventIterator{contract: _L1PoolManager.contract, event: "StakingETHEvent", logs: logs, sub: sub}, nil
}

// WatchStakingETHEvent is a free log subscription operation binding the contract event 0xe7466ea83435490635fc76a5f33da4815758ab48b1d45858f0452ca646556937.
//
// Solidity: event StakingETHEvent(address indexed user, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) WatchStakingETHEvent(opts *bind.WatchOpts, sink chan<- *L1PoolManagerStakingETHEvent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "StakingETHEvent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerStakingETHEvent)
				if err := _L1PoolManager.contract.UnpackLog(event, "StakingETHEvent", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseStakingETHEvent(log types.Log) (*L1PoolManagerStakingETHEvent, error) {
	event := new(L1PoolManagerStakingETHEvent)
	if err := _L1PoolManager.contract.UnpackLog(event, "StakingETHEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerStakingWETHEventIterator is returned from FilterStakingWETHEvent and is used to iterate over the raw logs and unpacked data for StakingWETHEvent events raised by the L1PoolManager contract.
type L1PoolManagerStakingWETHEventIterator struct {
	Event *L1PoolManagerStakingWETHEvent // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerStakingWETHEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerStakingWETHEvent)
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
		it.Event = new(L1PoolManagerStakingWETHEvent)
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
func (it *L1PoolManagerStakingWETHEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerStakingWETHEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerStakingWETHEvent represents a StakingWETHEvent event raised by the L1PoolManager contract.
type L1PoolManagerStakingWETHEvent struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakingWETHEvent is a free log retrieval operation binding the contract event 0xc138e3bd6d13eefa5cb01c0d35c5794001141efaf4e5ad888cad059935f83383.
//
// Solidity: event StakingWETHEvent(address indexed user, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) FilterStakingWETHEvent(opts *bind.FilterOpts, user []common.Address) (*L1PoolManagerStakingWETHEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "StakingWETHEvent", userRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerStakingWETHEventIterator{contract: _L1PoolManager.contract, event: "StakingWETHEvent", logs: logs, sub: sub}, nil
}

// WatchStakingWETHEvent is a free log subscription operation binding the contract event 0xc138e3bd6d13eefa5cb01c0d35c5794001141efaf4e5ad888cad059935f83383.
//
// Solidity: event StakingWETHEvent(address indexed user, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) WatchStakingWETHEvent(opts *bind.WatchOpts, sink chan<- *L1PoolManagerStakingWETHEvent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "StakingWETHEvent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerStakingWETHEvent)
				if err := _L1PoolManager.contract.UnpackLog(event, "StakingWETHEvent", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseStakingWETHEvent(log types.Log) (*L1PoolManagerStakingWETHEvent, error) {
	event := new(L1PoolManagerStakingWETHEvent)
	if err := _L1PoolManager.contract.UnpackLog(event, "StakingWETHEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerStarkingERC20EventIterator is returned from FilterStarkingERC20Event and is used to iterate over the raw logs and unpacked data for StarkingERC20Event events raised by the L1PoolManager contract.
type L1PoolManagerStarkingERC20EventIterator struct {
	Event *L1PoolManagerStarkingERC20Event // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerStarkingERC20EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerStarkingERC20Event)
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
		it.Event = new(L1PoolManagerStarkingERC20Event)
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
func (it *L1PoolManagerStarkingERC20EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerStarkingERC20EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerStarkingERC20Event represents a StarkingERC20Event event raised by the L1PoolManager contract.
type L1PoolManagerStarkingERC20Event struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStarkingERC20Event is a free log retrieval operation binding the contract event 0x3ac7d3823c11677fba9479ed26f696a3f17e16a5a5c39162fa6d183905aa6735.
//
// Solidity: event StarkingERC20Event(address indexed user, address indexed token, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) FilterStarkingERC20Event(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*L1PoolManagerStarkingERC20EventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "StarkingERC20Event", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerStarkingERC20EventIterator{contract: _L1PoolManager.contract, event: "StarkingERC20Event", logs: logs, sub: sub}, nil
}

// WatchStarkingERC20Event is a free log subscription operation binding the contract event 0x3ac7d3823c11677fba9479ed26f696a3f17e16a5a5c39162fa6d183905aa6735.
//
// Solidity: event StarkingERC20Event(address indexed user, address indexed token, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) WatchStarkingERC20Event(opts *bind.WatchOpts, sink chan<- *L1PoolManagerStarkingERC20Event, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "StarkingERC20Event", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerStarkingERC20Event)
				if err := _L1PoolManager.contract.UnpackLog(event, "StarkingERC20Event", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseStarkingERC20Event(log types.Log) (*L1PoolManagerStarkingERC20Event, error) {
	event := new(L1PoolManagerStarkingERC20Event)
	if err := _L1PoolManager.contract.UnpackLog(event, "StarkingERC20Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerTransferAssertToIterator is returned from FilterTransferAssertTo and is used to iterate over the raw logs and unpacked data for TransferAssertTo events raised by the L1PoolManager contract.
type L1PoolManagerTransferAssertToIterator struct {
	Event *L1PoolManagerTransferAssertTo // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerTransferAssertToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerTransferAssertTo)
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
		it.Event = new(L1PoolManagerTransferAssertTo)
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
func (it *L1PoolManagerTransferAssertToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerTransferAssertToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerTransferAssertTo represents a TransferAssertTo event raised by the L1PoolManager contract.
type L1PoolManagerTransferAssertTo struct {
	Blockchain *big.Int
	Token      common.Address
	To         common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferAssertTo is a free log retrieval operation binding the contract event 0x08ffba656e4ac665e1aa4bb6a342e1a69d4cc12fd751ac862fa3fe27b0c7524a.
//
// Solidity: event TransferAssertTo(uint256 Blockchain, address indexed token, address indexed to, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) FilterTransferAssertTo(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*L1PoolManagerTransferAssertToIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "TransferAssertTo", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerTransferAssertToIterator{contract: _L1PoolManager.contract, event: "TransferAssertTo", logs: logs, sub: sub}, nil
}

// WatchTransferAssertTo is a free log subscription operation binding the contract event 0x08ffba656e4ac665e1aa4bb6a342e1a69d4cc12fd751ac862fa3fe27b0c7524a.
//
// Solidity: event TransferAssertTo(uint256 Blockchain, address indexed token, address indexed to, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) WatchTransferAssertTo(opts *bind.WatchOpts, sink chan<- *L1PoolManagerTransferAssertTo, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "TransferAssertTo", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerTransferAssertTo)
				if err := _L1PoolManager.contract.UnpackLog(event, "TransferAssertTo", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseTransferAssertTo(log types.Log) (*L1PoolManagerTransferAssertTo, error) {
	event := new(L1PoolManagerTransferAssertTo)
	if err := _L1PoolManager.contract.UnpackLog(event, "TransferAssertTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the L1PoolManager contract.
type L1PoolManagerUnpausedIterator struct {
	Event *L1PoolManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerUnpaused)
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
		it.Event = new(L1PoolManagerUnpaused)
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
func (it *L1PoolManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerUnpaused represents a Unpaused event raised by the L1PoolManager contract.
type L1PoolManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1PoolManager *L1PoolManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*L1PoolManagerUnpausedIterator, error) {

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerUnpausedIterator{contract: _L1PoolManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1PoolManager *L1PoolManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *L1PoolManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerUnpaused)
				if err := _L1PoolManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseUnpaused(log types.Log) (*L1PoolManagerUnpaused, error) {
	event := new(L1PoolManagerUnpaused)
	if err := _L1PoolManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the L1PoolManager contract.
type L1PoolManagerWithdrawIterator struct {
	Event *L1PoolManagerWithdraw // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerWithdraw)
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
		it.Event = new(L1PoolManagerWithdraw)
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
func (it *L1PoolManagerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerWithdraw represents a Withdraw event raised by the L1PoolManager contract.
type L1PoolManagerWithdraw struct {
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
func (_L1PoolManager *L1PoolManagerFilterer) FilterWithdraw(opts *bind.FilterOpts) (*L1PoolManagerWithdrawIterator, error) {

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerWithdrawIterator{contract: _L1PoolManager.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x0ad30059d4072ae2525ebaaa5b6878f1c50df5f3b921615031c7d0f12dc4a77b.
//
// Solidity: event Withdraw(address _user, uint256 startPoolId, uint256 EndPoolId, address _token, uint256 Amount, uint256 Reward)
func (_L1PoolManager *L1PoolManagerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *L1PoolManagerWithdraw) (event.Subscription, error) {

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerWithdraw)
				if err := _L1PoolManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseWithdraw(log types.Log) (*L1PoolManagerWithdraw, error) {
	event := new(L1PoolManagerWithdraw)
	if err := _L1PoolManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
