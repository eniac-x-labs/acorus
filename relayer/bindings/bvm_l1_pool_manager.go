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

// IL1PoolPool is an auto generated low-level Go binding around an user-defined struct.
type IL1PoolPool struct {
	StartTimestamp  uint32
	EndTimestamp    uint32
	Token           common.Address
	TotalAmount     *big.Int
	TotalFee        *big.Int
	TotalFeeClaimed *big.Int
	IsCompleted     bool
}

// IL1PoolUser is an auto generated low-level Go binding around an user-defined struct.
type IL1PoolUser struct {
	IsClaimed   bool
	Token       common.Address
	StartPoolId *big.Int
	EndPoolId   *big.Int
	Amount      *big.Int
}

// L1PoolManagerMetaData contains all meta data concerning the L1PoolManager contract.
var L1PoolManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ChainIdIsNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"ChainIdNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorBlockChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"providedAmount\",\"type\":\"uint256\"}],\"name\":\"LessThanMinStakeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"MinTransferAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LessThanMinTransferAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LessThanZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PoolIndex\",\"type\":\"uint256\"}],\"name\":\"NewPoolIsNotCreate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReward\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughETH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"NotEnoughToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolIndex\",\"type\":\"uint256\"}],\"name\":\"PoolIsCompleted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"StableCoinNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"name\":\"TokenIsAlreadySupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenIsNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endPoolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolIndex\",\"type\":\"uint256\"}],\"name\":\"CompletePoolEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetMinStakeAmountEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSupport\",\"type\":\"bool\"}],\"name\":\"SetSupportTokenEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingETHEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingWETHEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StarkingERC20Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Blockchain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferAssertTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"BridgeInitiateETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ClaimAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"ClaimSimpleAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"IsCompleted\",\"type\":\"bool\"}],\"internalType\":\"structIL1Pool.Pool[]\",\"name\":\"CompletePools\",\"type\":\"tuple[]\"}],\"name\":\"CompletePoolAndNew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DepositAndStaking\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DepositAndStakingERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DepositAndStakingETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositAndStakingWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FeePoolValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"IsSupportChainId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"IsSupportStableCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"IsSupportToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MinStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinTransferAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Pools\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"IsCompleted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ReLayer\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SendAssertToUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isSupport\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"startTimes\",\"type\":\"uint32\"}],\"name\":\"SetSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SupportTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"Blockchain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TransferAssertToBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Users\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isClaimed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"StartPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getPool\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"IsCompleted\",\"type\":\"bool\"}],\"internalType\":\"structIL1Pool.Pool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getPoolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isClaimed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"StartPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"}],\"internalType\":\"structIL1Pool.User[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isClaimed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"StartPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"}],\"internalType\":\"structIL1Pool.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MultisigWallet\",\"type\":\"address\"},{\"internalType\":\"contractIMessageManager\",\"name\":\"_messageManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageManager\",\"outputs\":[{\"internalType\":\"contractIMessageManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodTime\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setMinStakeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_MinTransferAmount\",\"type\":\"uint256\"}],\"name\":\"setMinTransferAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_PerFee\",\"type\":\"uint256\"}],\"name\":\"setPerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setSupportStableCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setValidChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506200001b62000021565b620000d5565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d25780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b61508680620000e35f395ff3fe6080604052600436106102f3575f3560e01c806373c23be911610189578063cb314fab116100d8578063dd0c346011610092578063f91fa9a81161006d578063f91fa9a8146109f1578063fa86184814610a10578063faeb96a514610a24578063ff2bf64f14610a43575f80fd5b8063dd0c346014610971578063ed06da8714610986578063f8fcc2aa146109bd575f80fd5b8063cb314fab14610893578063cb4f04ad146108b2578063cd01c665146108e6578063d547741f14610905578063d73f14e514610924578063dac2956814610943575f80fd5b80639ab0b65211610143578063b1887e9b1161011e578063b1887e9b14610822578063b33280dd14610836578063b92e639614610855578063bc42493d14610874575f80fd5b80639ab0b652146107e95780639b442380146107fc578063a217fddf1461080f575f80fd5b806373c23be91461075d5780638456cb591461077057806391d148541461078457806392d84738146107a357806395455eef146107c257806396f984e6146107e1575f80fd5b80633f4ba83a116102455780635b5b9ea2116101ff578063650c2276116101da578063650c2276146106d05780636d6a78bf146106ef5780636f77926b1461070357806372fe6a7e1461072f575f80fd5b80635b5b9ea2146106085780635c975abb1461068e578063626417b5146106b1575f80fd5b80633f4ba83a146105005780634663cdc814610514578063485cc955146105805780634dfaef841461059f57806354dc027e146105be5780635765a4eb146105e9575f80fd5b8063248a9ca3116102b05780633338562c1161028b5780633338562c1461049c57806336568abe146104af57806338731c47146104ce5780633b0230f0146104e1575f80fd5b8063248a9ca31461043157806327e235e3146104505780632f2ff15d1461047b575f80fd5b806301ffc9a7146102f757806309dc480c1461032b57806310875a131461034e57806313e8544e146103845780631ca2f173146103d55780631d31fac014610400575b5f80fd5b348015610302575f80fd5b50610316610311366004614792565b610a6f565b60405190151581526020015b60405180910390f35b348015610336575f80fd5b5061034060015481565b604051908152602001610322565b348015610359575f80fd5b505f5461036c906001600160a01b031681565b6040516001600160a01b039091168152602001610322565b34801561038f575f80fd5b506103a361039e3660046147cd565b610aa5565b6040805195151586526001600160a01b039094166020860152928401919091526060830152608082015260a001610322565b3480156103e0575f80fd5b506103406103ef3660046147f7565b600c6020525f908152604090205481565b34801561040b575f80fd5b5060075461041c9063ffffffff1681565b60405163ffffffff9091168152602001610322565b34801561043c575f80fd5b5061034061044b366004614812565b610afb565b34801561045b575f80fd5b5061034061046a3660046147f7565b60096020525f908152604090205481565b348015610486575f80fd5b5061049a610495366004614829565b610b1b565b005b61049a6104aa3660046148e5565b610b3d565b3480156104ba575f80fd5b5061049a6104c9366004614829565b610ea3565b6103166104dc3660046149fb565b610ed6565b3480156104ec575f80fd5b506103166104fb366004614a4a565b61106a565b34801561050b575f80fd5b5061049a61136a565b34801561051f575f80fd5b5061053361052e3660046147cd565b61137f565b6040805163ffffffff98891681529790961660208801526001600160a01b03909416948601949094526060850191909152608084015260a0830191909152151560c082015260e001610322565b34801561058b575f80fd5b5061049a61059a366004614a98565b6113f0565b3480156105aa575f80fd5b5061049a6105b93660046147cd565b611522565b3480156105c9575f80fd5b506103406105d83660046147f7565b60066020525f908152604090205481565b3480156105f4575f80fd5b506103166106033660046149fb565b6117b5565b348015610613575f80fd5b506106276106223660046147cd565b6119aa565b6040516103229190815163ffffffff9081168252602080840151909116908201526040808301516001600160a01b031690820152606080830151908201526080808301519082015260a0808301519082015260c09182015115159181019190915260e00190565b348015610699575f80fd5b505f805160206150118339815191525460ff16610316565b3480156106bc575f80fd5b5061036c6106cb366004614812565b611a8c565b3480156106db575f80fd5b5061049a6106ea366004614812565b611ab4565b3480156106fa575f80fd5b5061049a611ac4565b34801561070e575f80fd5b5061072261071d3660046147f7565b611b3a565b6040516103229190614afc565b34801561073a575f80fd5b50610316610749366004614812565b5f9081526003602052604090205460ff1690565b61031661076b366004614b49565b611be8565b34801561077b575f80fd5b5061049a611da9565b34801561078f575f80fd5b5061031661079e366004614829565b611dbb565b3480156107ae575f80fd5b506103166107bd366004614b87565b611df1565b3480156107cd575f80fd5b5061049a6107dc3660046147f7565b611fb2565b61049a6126fd565b61049a6107f73660046147cd565b612a10565b61031661080a366004614be7565b612aa1565b34801561081a575f80fd5b506103405f81565b34801561082d575f80fd5b5061036c612cb8565b348015610841575f80fd5b5061049a610850366004614c1d565b612d26565b348015610860575f80fd5b5061049a61086f366004614812565b612d5b565b34801561087f575f80fd5b5061049a61088e366004614812565b612dae565b34801561089e575f80fd5b5061049a6108ad3660046147cd565b613127565b3480156108bd575f80fd5b506103406108cc3660046147f7565b6001600160a01b03165f908152600b602052604090205490565b3480156108f1575f80fd5b50610316610900366004614c49565b61318a565b348015610910575f80fd5b5061049a61091f366004614829565b613515565b34801561092f575f80fd5b5061049a61093e366004614c85565b613531565b34801561094e575f80fd5b5061031661095d3660046147f7565b60086020525f908152604090205460ff1681565b34801561097c575f80fd5b5061034060025481565b348015610991575f80fd5b506103166109a03660046147f7565b6001600160a01b03165f9081526004602052604090205460ff1690565b3480156109c8575f80fd5b506103406109d73660046147f7565b6001600160a01b03165f908152600a602052604090205490565b3480156109fc575f80fd5b5061049a610a0b366004614ca8565b61355b565b348015610a1b575f80fd5b50610340613707565b348015610a2f575f80fd5b5061049a610a3e366004614cdd565b61374e565b348015610a4e575f80fd5b50610a62610a5d3660046147cd565b613a44565b6040516103229190614d21565b5f6001600160e01b03198216637965db0b60e01b1480610a9f57506301ffc9a760e01b6001600160e01b03198316145b92915050565b600b602052815f5260405f208181548110610abe575f80fd5b5f918252602090912060049091020180546001820154600283015460039093015460ff831695506101009092046001600160a01b03169350919085565b5f9081525f80516020614ff1833981519152602052604090206001015490565b610b2482610afb565b610b2d81613b06565b610b378383613b10565b50505050565b60ff19610b5860015f80516020614fd1833981519152614d43565b604051602001610b6a91815260200190565b6040516020818303038152906040528051906020012016610b8a81613b06565b5f5b8251811015610e9e575f838281518110610ba857610ba8614d56565b60200260200101516040015190505f6001600a5f846001600160a01b03166001600160a01b031681526020019081526020015f2080549050610bea9190614d43565b6001600160a01b0383165f908152600a60205260409020805491925060019183908110610c1957610c19614d56565b905f5260205f2090600502016004015f6101000a81548160ff021916908315150217905550848381518110610c5057610c50614d56565b602002602001015160800151600a5f846001600160a01b03166001600160a01b031681526020019081526020015f208281548110610c9057610c90614d56565b905f5260205f209060050201600201819055505f600a5f846001600160a01b03166001600160a01b031681526020019081526020015f208281548110610cd857610cd8614d56565b5f91825260208083206005909202909101546001600160a01b0386168352600a8252604092839020835160e0810190945263ffffffff6401000000009092048216808552600754909550909392830191610d33911685614d6a565b63ffffffff168152602001856001600160a01b03168152602001600a5f876001600160a01b03166001600160a01b031681526020019081526020015f208581548110610d8157610d81614d56565b5f918252602080832060059283020160019081015485528482018490526040808601859052606095860185905287548083018955978552938290208651979093029092018054868301518786015163ffffffff998a1667ffffffffffffffff1990931692909217640100000000999091169890980297909717600160401b600160e01b031916600160401b6001600160a01b039889160217815593850151918401919091556080840151600284015560a0840151600384015560c0909301516004909201805460ff19169215159290921790915551848152918516917fb6f449f07ceaf55392c9899e0797c6529908ae827c2d498c682e90d42c241167910160405180910390a25050508080610e9690614d87565b915050610b8c565b505050565b6001600160a01b0381163314610ecc5760405163334bd91960e11b815260040160405180910390fd5b610e9e8282613bb8565b5f8681526003602052604081205460ff16610f0c57604051632ef6974160e11b8152600481018890526024015b60405180910390fd5b5f8681526003602052604090205460ff16610f3d57604051632ef6974160e11b815260048101879052602401610f03565b604051309085156108fc029086905f818181858888f19350505050158015610f67573d5f803e3d5ffd5b5073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f90815260056020525f80516020614f718339815191528054869290610fa4908490614d43565b90915550505f54604051631eb65ffb60e01b81526001600160a01b0390911690631eb65ffb90610fe2908a908a908a908a908a908a90600401614d9f565b5f604051808303815f87803b158015610ff9575f80fd5b505af115801561100b573d5f803e3d5ffd5b5050604080518a8152602081018a90529081018790526001600160a01b03881692503091507f61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f329060600160405180910390a35060019695505050505050565b5f8581526003602052604081205460ff1661109b57604051632ef6974160e11b815260048101879052602401610f03565b5f8581526003602052604090205460ff166110cc57604051632ef6974160e11b815260048101869052602401610f03565b6001600160a01b0383165f9081526004602052604090205460ff1661110f57604051631e859bc960e01b81526001600160a01b0384166004820152602401610f03565b6040516370a0823160e01b81523060048201525f906001600160a01b038516906370a0823190602401602060405180830381865afa158015611153573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111779190614dd0565b905061118e6001600160a01b038516333086613c31565b6040516370a0823160e01b81523060048201525f906001600160a01b038616906370a0823190602401602060405180830381865afa1580156111d2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111f69190614dd0565b90505f6112038383614d43565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f90815260056020525f80516020614f7183398151915280549293508792909190611244908490614de7565b90915550506002545f90620f42409061125d9084614dfa565b6112679190614e11565b90506112738183614d43565b6001600160a01b0388165f9081526006602052604081208054929450839290919061129f908490614de7565b90915550505f5460405163305f478560e21b81526001600160a01b039091169063c17d1e14906112db908d908d908d908c908890600401614e30565b5f604051808303815f87803b1580156112f2575f80fd5b505af1158015611304573d5f803e3d5ffd5b5050604080518d8152602081018d90529081018590526001600160a01b03808c1693503392508a16907fece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a49060600160405180910390a45060019998505050505050505050565b5f61137481613b06565b61137c613c8b565b50565b600a602052815f5260405f208181548110611398575f80fd5b5f9182526020909120600590910201805460018201546002830154600384015460049094015463ffffffff80851697506401000000008504169550600160401b9093046001600160a01b031693919290919060ff1687565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f811580156114355750825b90505f8267ffffffffffffffff1660011480156114515750303b155b90508115801561145f575080155b1561147d5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156114a757845460ff60401b1916600160401b1785555b6114af613cea565b6114b7613cf2565b6114c18787613d02565b6007805463ffffffff1916621baf80179055831561151957845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b61152a613d48565b611532613d7f565b6001600160a01b0382165f9081526008602052604090205460ff16611575576040516305fd61ad60e01b81526001600160a01b0383166004820152602401610f03565b6001600160a01b0382165f908152600c60205260409020548110156115d1576001600160a01b0382165f908152600c6020526040908190205490516327500c6d60e21b8152600481019190915260248101829052604401610f03565b6115e66001600160a01b038316333084613c31565b6001600160a01b0382165f908152600a602052604081205461160a90600190614d43565b6001600160a01b0384165f908152600a60205260409020805491925042918390811061163857611638614d56565b5f91825260209091206005909102015463ffffffff16111561173e57335f908152600b60209081526040808320815160a0810183528481526001600160a01b0388811682860181815283860189815260608501898152608086018c8152875460018082018a55988c528a8c2097516004909102909701805494516001600160a81b0319909516971515610100600160a81b03191697909717610100949096169390930294909417855551948401949094559051600283015591516003909101558352600a909152902080548391908390811061171657611716614d56565b905f5260205f2090600502016001015f8282546117339190614de7565b9091555061175a9050565b604051637d58ebb960e01b815260048101829052602401610f03565b6040518281526001600160a01b0384169033907f3ac7d3823c11677fba9479ed26f696a3f17e16a5a5c39162fa6d183905aa67359060200160405180910390a3506117b160015f8051602061503183398151915255565b5050565b5f8681526003602052604081205460ff166117e657604051632ef6974160e11b815260048101889052602401610f03565b5f8681526003602052604090205460ff1661181757604051632ef6974160e11b815260048101879052602401610f03565b5f611820612cb8565b6040516323b872dd60e01b81529091506001600160a01b038216906323b872dd906118539030908a908a90600401614e5c565b6020604051808303815f875af115801561186f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118939190614e80565b5073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25f90815260056020527fa550ba85c46b24b567d2e17cd597f2283877afab43603f46d5de7858f1bdb73180548792906118e3908490614d43565b90915550505f54604051631eb65ffb60e01b81526001600160a01b0390911690631eb65ffb90611921908b908b908b908b908b908b90600401614d9f565b5f604051808303815f87803b158015611938575f80fd5b505af115801561194a573d5f803e3d5ffd5b5050604080518b8152602081018b90529081018890526001600160a01b03891692503091507f7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc149060600160405180910390a3506001979650505050505050565b6040805160e0810182525f8082526020808301829052828401829052606083018290526080830182905260a0830182905260c083018290526001600160a01b0386168252600a905291909120805483908110611a0857611a08614d56565b5f9182526020918290206040805160e081018252600593909302909101805463ffffffff808216855264010000000082041694840194909452600160401b9093046001600160a01b0316908201526001820154606082015260028201546080820152600382015460a082015260049091015460ff16151560c0820152905092915050565b600d8181548110611a9b575f80fd5b5f918252602090912001546001600160a01b0316905081565b5f611abe81613b06565b50600255565b611acc613d48565b611ad4613d7f565b5f5b600d54811015611b2157611b0f600d8281548110611af657611af6614d56565b5f918252602090912001546001600160a01b0316611fb2565b80611b1981614d87565b915050611ad6565b50611b3860015f8051602061503183398151915255565b565b6001600160a01b0381165f908152600b60209081526040808320805482518185028101850190935280835260609492939192909184015b82821015611bdd575f8481526020908190206040805160a08101825260048602909201805460ff81161515845261010090046001600160a01b03168385015260018082015492840192909252600281015460608401526003015460808301529083529092019101611b71565b505050509050919050565b5f60ff19611c0460015f80516020614fd1833981519152614d43565b604051602001611c1691815260200190565b6040516020818303038152906040528051906020012016611c3681613b06565b6001600160a01b0385165f9081526004602052604090205460ff16611c7957604051631e859bc960e01b81526001600160a01b0386166004820152602401610f03565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeed196001600160a01b03861601611cf85782471015611cbf57604051632c1d501360e11b815260040160405180910390fd5b6040516001600160a01b0385169084156108fc029085905f818181858888f19350505050158015611cf2573d5f803e3d5ffd5b50611d9e565b6040516370a0823160e01b815230600482015283906001600160a01b038716906370a0823190602401602060405180830381865afa158015611d3c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d609190614dd0565b1015611d8a576040516311745b6160e21b81526001600160a01b0386166004820152602401610f03565b611d9e6001600160a01b0386168585613daf565b506001949350505050565b5f611db381613b06565b61137c613de0565b5f9182525f80516020614ff1833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b5f8781526003602052604081205460ff16611e2257604051632ef6974160e11b815260048101899052602401610f03565b5f8781526003602052604090205460ff16611e5357604051632ef6974160e11b815260048101889052602401610f03565b6001600160a01b0385165f9081526004602052604090205460ff16611e9657604051631e859bc960e01b81526001600160a01b0386166004820152602401610f03565b611eab6001600160a01b038616308887613c31565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f90815260056020525f80516020614f718339815191528054869290611ee7908490614d43565b90915550505f54604051631eb65ffb60e01b81526001600160a01b0390911690631eb65ffb90611f25908b908b908b908a908a908a90600401614d9f565b5f604051808303815f87803b158015611f3c575f80fd5b505af1158015611f4e573d5f803e3d5ffd5b5050604080518b8152602081018b90529081018790526001600160a01b03808a1693503092508816907f0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe9060600160405180910390a4506001979650505050505050565b611fba613d48565b611fc2613d7f565b6001600160a01b0381165f9081526008602052604090205460ff16612005576040516305fd61ad60e01b81526001600160a01b0382166004820152602401610f03565b6001600160a01b0381165f908152600a6020526040812054900361203e57604051637d58ebb960e01b81525f6004820152602401610f03565b5f5b335f908152600b60205260409020548110156126e657335f908152600b60205260409020805482916001600160a01b038516918390811061208357612083614d56565b5f91825260209091206004909102015461010090046001600160a01b03160361257657335f908152600b602052604090208054829081106120c6576120c6614d56565b5f91825260209091206004909102015460ff16156120e457506126d4565b6001600160a01b0383165f908152600a602052604081205461210890600190614d43565b335f908152600b602052604090208054919250908390811061212c5761212c614d56565b905f5260205f20906004020160030154600a5f866001600160a01b03166001600160a01b031681526020019081526020015f20828154811061217057612170614d56565b905f5260205f2090600502016001015f82825461218d9190614d43565b9091555050335f908152600b602052604081208054829190859081106121b5576121b5614d56565b905f5260205f2090600402016003015490505f600b5f336001600160a01b03166001600160a01b031681526020019081526020015f2085815481106121fc576121fc614d56565b905f5260205f2090600402016001015490508381111561222f5760405163374c934360e11b815260040160405180910390fd5b805b84811015612383576001600160a01b0388165f908152600a602052604090205461225d90600190614d43565b81111561228057604051637d58ebb960e01b815260048101829052602401610f03565b6001600160a01b0388165f908152600a602052604081208054839081106122a9576122a9614d56565b905f5260205f20906005020160010154600a5f8b6001600160a01b03166001600160a01b031681526020019081526020015f2083815481106122ed576122ed614d56565b905f5260205f20906005020160020154856123089190614dfa565b6123129190614e11565b905061231e8186614de7565b6001600160a01b038a165f908152600a60205260409020805491965082918490811061234c5761234c614d56565b905f5260205f2090600502016003015f8282546123699190614de7565b9091555082915061237b905081614d87565b915050612231565b505f83116123bf5760405162461bcd60e51b8152602060048201526009602482015268139bc814995dd85c9960ba1b6044820152606401610f03565b6123c98383614de7565b335f908152600b602052604090208054919350600191879081106123ef576123ef614d56565b5f9182526020909120600490910201805460ff19169115159190911790556001600160a01b03871673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeed190161246257604051339083156108fc029084905f818181858888f1935050505015801561245c573d5f803e3d5ffd5b50612514565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc1196001600160a01b038816016125005760405163a9059cbb60e01b81523360048201526024810183905273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc29063a9059cbb906044016020604051808303815f875af11580156124dc573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061245c9190614e80565b6125146001600160a01b0388163384613daf565b6001600160a01b038716337f991c329471ab5230e2301ee30d0d6fba5906e32411d5d154d5a8c278d021a2ab838761254c8888614d43565b604080519384526020840192909252908201526060810187905260800160405180910390a3505050505b335f908152600b6020526040902054600110156126d257335f908152600b6020526040902080546125a990600190614d43565b815481106125b9576125b9614d56565b905f5260205f209060040201600b5f336001600160a01b03166001600160a01b031681526020019081526020015f2082815481106125f9576125f9614d56565b5f91825260208083208454600490930201805460ff909316151560ff1984168117825585546001600160a01b03610100918290041602610100600160a81b03199091166001600160a81b03199094169390931792909217825560018085015490830155600280850154908301556003938401549390910192909255338152600b9091526040902080548061268f5761268f614e9b565b5f8281526020812060045f199093019283020180546001600160a81b03191681556001810182905560028101829055600301559055816126ce81614eaf565b9250505b505b806126de81614eca565b915050612040565b5061137c60015f8051602061503183398151915255565b612705613d48565b61270d613d7f565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f52600c6020527f36dda7fac5d19118789d83982ba89ab2ea0a81d18d22af2919c731e22f75f227543410156127b35773eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f52600c6020527f36dda7fac5d19118789d83982ba89ab2ea0a81d18d22af2919c731e22f75f227546040516327500c6d60e21b81526004810191909152346024820152604401610f03565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f908152600a6020525f80516020614fb1833981519152546127ec90600190614d43565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f908152600a6020525f80516020614fb1833981519152549192500361283d57604051637d58ebb960e01b815260016004820152602401610f03565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f52600a6020525f80516020614fb183398151915280544291908390811061287c5761287c614d56565b5f91825260209091206005909102015463ffffffff16111561299d57335f908152600b60209081526040808320815160a08101835284815273eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee818501818152938201878152606083018781523460808501818152865460018082018955978b52898b2096516004909102909601805498516001600160a81b0319909916961515610100600160a81b031916969096176101006001600160a01b03909916989098029790971785559151948401949094559251600283015592516003909101559252600a90525f80516020614fb183398151915280548390811061297557612975614d56565b905f5260205f2090600502016001015f8282546129929190614de7565b909155506129c49050565b6129a8816001614de7565b604051637d58ebb960e01b8152600401610f0391815260200190565b60405134815233907fe7466ea83435490635fc76a5f33da4815758ab48b1d45858f0452ca6465569379060200160405180910390a250611b3860015f8051602061503183398151915255565b612a18613d48565b612a20613d7f565b3415612a3357612a2e6126fd565b612a8b565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc1196001600160a01b03831601612a6157612a2e81612dae565b6001600160a01b0382165f9081526008602052604090205460ff1615612a8b57612a8b8282611522565b6117b160015f8051602061503183398151915255565b5f8381526003602052604081205460ff16612ad257604051632ef6974160e11b815260048101859052602401610f03565b5f8381526003602052604090205460ff16612b0357604051632ef6974160e11b815260048101849052602401610f03565b600154341015612b33576001546040516358f8331360e11b81526004810191909152346024820152604401610f03565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f90815260056020525f80516020614f718339815191528054349290612b6f908490614de7565b90915550506002545f90620f424090612b889034614dfa565b612b929190614e11565b90505f612b9f8234614d43565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f90815260066020527fa2e5aefc6e2cbe2917a296f0fd89c5f915c487c803db1d98eccb43f14012d71180549293508492909190612bf3908490614de7565b90915550505f5460405163305f478560e21b81526001600160a01b039091169063c17d1e1490612c2f9089908990899034908990600401614e30565b5f604051808303815f87803b158015612c46575f80fd5b505af1158015612c58573d5f803e3d5ffd5b505060408051898152602081018990529081018490526001600160a01b03871692503391507f0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc59060600160405180910390a36001925050505b9392505050565b5f4662082750819003612cd3576004605360981b0191505090565b8061044d03612cf757734f9a0e7fd2bf6067db6994cf12e4495df938e6e991505090565b80600a03612d0d576006602160991b0191505090565b604051639474dee160e01b815260040160405180910390fd5b5f612d3081613b06565b506001600160a01b03919091165f908152600460205260409020805460ff1916911515919091179055565b60ff19612d7660015f80516020614fd1833981519152614d43565b604051602001612d8891815260200190565b6040516020818303038152906040528051906020012016612da881613b06565b50600155565b612db6613d48565b612dbe613d7f565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25f52600c6020527fc87f3dae70d3ac23e82424b4132d779583ae145fc334a52160c2d1a246c2bc2b54811015612e51575f8052600c6020527f13649b2456f1b42fef0f0040b3aaeabcd21a76a0f3f5defd4f583839455116e8546040516327500c6d60e21b8152600481019190915260248101829052604401610f03565b6040516323b872dd60e01b815273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2906323b872dd90612e8c90339030908690600401614e5c565b6020604051808303815f875af1158015612ea8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ecc9190614e80565b5073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25f908152600a6020525f80516020614f9183398151915254612f0690600190614d43565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25f52600a6020525f80516020614f9183398151915280549192509082908110612f4657612f46614d56565b5f91825260209091206004600590920201015460ff1615612f7d576040516311cf1b0760e31b815260048101829052602401610f03565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25f52600a6020525f80516020614f91833981519152805442919083908110612fbc57612fbc614d56565b5f91825260209091206005909102015463ffffffff1611156130db57335f908152600b60209081526040808320815160a08101835284815273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc281850181815293820187815260608301878152608084018a8152855460018181018855968a52888a2095516004909102909501805497516001600160a01b031661010002610100600160a81b0319961515969096166001600160a81b0319909816979097179490941786559051938501939093559151600284015551600390920191909155909152600a90525f80516020614f918339815191528054839190839081106130b8576130b8614d56565b905f5260205f2090600502016001015f8282546130d59190614de7565b90915550505b60405182815233907fc138e3bd6d13eefa5cb01c0d35c5794001141efaf4e5ad888cad059935f833839060200160405180910390a25061137c60015f8051602061503183398151915255565b5f61313181613b06565b6001600160a01b0383165f818152600c602052604090819020849055517ff54d3b756d286b6b08e5d4eda6dfe5b135664abf029e58e637cbf013c442c9509061317d9085815260200190565b60405180910390a2505050565b5f8481526003602052604081205460ff166131bb576040516363b5c9db60e01b815260048101869052602401610f03565b5f8481526003602052604090205460ff166131ec576040516363b5c9db60e01b815260048101859052602401610f03565b5f6131f5612cb8565b6040516370a0823160e01b81523060048201529091505f906001600160a01b038316906370a0823190602401602060405180830381865afa15801561323c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906132609190614dd0565b6040516323b872dd60e01b81529091506001600160a01b038316906323b872dd9061329390339030908990600401614e5c565b6020604051808303815f875af11580156132af573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906132d39190614e80565b506040516370a0823160e01b81523060048201525f906001600160a01b038416906370a0823190602401602060405180830381865afa158015613318573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061333c9190614dd0565b90505f6133498383614d43565b905060015481101561337c576001546040516358f8331360e11b8152600481019190915260248101829052604401610f03565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25f90815260056020527fa550ba85c46b24b567d2e17cd597f2283877afab43603f46d5de7858f1bdb73180548392906133cb908490614de7565b90915550506002545f90620f4240906133e49084614dfa565b6133ee9190614e11565b90506133fa8183614d43565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25f90815260066020527f5e5777fab7622aff3c042c1ece74307c2e9d699a9da444f416c35f2e1def28a58054929450839290919061344e908490614de7565b90915550505f5460405163305f478560e21b81526001600160a01b039091169063c17d1e149061348a908d908d908d908d908890600401614e30565b5f604051808303815f87803b1580156134a1575f80fd5b505af11580156134b3573d5f803e3d5ffd5b5050604080518d8152602081018d90529081018590526001600160a01b038b1692503391507fcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a899060600160405180910390a35060019998505050505050505050565b61351e82610afb565b61352781613b06565b610b378383613bb8565b5f61353b81613b06565b505f91825260036020526040909120805460ff1916911515919091179055565b60ff1961357660015f80516020614fd1833981519152614d43565b60405160200161358891815260200190565b60405160208183030381529060405280519060200120166135a881613b06565b6001600160a01b0384165f9081526008602052604090205460ff166135eb576040516305fd61ad60e01b81526001600160a01b0385166004820152602401610f03565b84620827500361367e57604051636bb825d760e11b815261753060048201525f90730d7e906bd9cafa154b048cfa766cc1e54e39af9b9063d7704bae90602401602060405180830381865afa158015613646573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061366a9190614dd0565b905061367885858584613e28565b506136aa565b8461044d036136975761369284848461411c565b6136aa565b84600a03612d0d57613692848484614286565b826001600160a01b0316846001600160a01b03167f08ffba656e4ac665e1aa4bb6a342e1a69d4cc12fd751ac862fa3fe27b0c7524a87856040516136f8929190918252602082015260400190565b60405180910390a35050505050565b60ff1961372260015f80516020614fd1833981519152614d43565b60405160200161373491815260200190565b604051602081830303815290604052805190602001201681565b5f61375881613b06565b6001600160a01b0384165f9081526008602052604090205460ff16156137a45760405163411befff60e11b81526001600160a01b03851660048201528315156024820152604401610f03565b6001600160a01b0384165f908152600860209081526040808320805460ff1916871515179055600a90915290819020815160e0810190925260075490919081906137f49063ffffffff1686614ee1565b63ffffffff908116825285811660208084018290526001600160a01b03808b1660408087018290525f60608089018290526080808a0183905260a0808b0184905260c09a8b018490528c5460018181018f559d85528885208d5160059092020180548e8b01518f890151909a16600160401b02600160401b600160e01b03199a8e166401000000000267ffffffffffffffff19909216938e16939093171798909816178755918b01519b86019b909b5599890151600285015598880151600384015596909501516004909101805491151560ff19909216919091179055928552600a835293839020835160e08101909452938352600754918301916138fa911686614d6a565b63ffffffff90811682526001600160a01b0388811660208085018290525f60408087018290526060808801839052608080890184905260a09889018490528a5460018082018d559b85528585208b5160059092020180548c8801518d870151938c1667ffffffffffffffff199092169190911764010000000091909b160299909917600160401b600160e01b031916600160401b9190981602969096178755880151868a015593870151600286015594860151600385015560c0909501516004909301805460ff191693151593909317909255600d805495860181559092527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb590930180546001600160a01b031916841790555185151581527fc8c34f23fafb34e68119c1d231ef03d0d47225b15e2c4de3efbefa14b0181d86910160405180910390a250505050565b613a7c6040518060a001604052805f151581526020015f6001600160a01b031681526020015f81526020015f81526020015f81525090565b6001600160a01b0383165f908152600b60205260409020805483908110613aa557613aa5614d56565b5f9182526020918290206040805160a081018252600493909302909101805460ff81161515845261010090046001600160a01b0316938301939093526001830154908201526002820154606082015260039091015460808201529392505050565b61137c8133614441565b5f5f80516020614ff1833981519152613b298484611dbb565b613ba8575f848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055613b5e3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610a9f565b5f915050610a9f565b5092915050565b5f5f80516020614ff1833981519152613bd18484611dbb565b15613ba8575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610a9f565b610b3784856001600160a01b03166323b872dd868686604051602401613c5993929190614e5c565b604051602081830303815290604052915060e01b6020820180516001600160e01b03838183161783525050505061447a565b613c936144db565b5f80516020615011833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b611b3861450a565b613cfa61450a565b611b38614553565b613d0a61450a565b67016345785d8a0000600155612710600255613d265f83610b1b565b5f80546001600160a01b0319166001600160a01b039290921691909117905550565b5f80516020615031833981519152805460011901613d7957604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f805160206150118339815191525460ff1615611b385760405163d93c066560e01b815260040160405180910390fd5b6040516001600160a01b03838116602483015260448201839052610e9e91859182169063a9059cbb90606401613c59565b613de8613d7f565b5f80516020615011833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833613ccc565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeed196001600160a01b03851601613edb57736ea73e05adc79974b931123675ea8f78ffdacdf063ce0b63ce613e718385614de7565b85855a6040516001600160e01b031960e087901b1681526001600160a01b039093166004840152602483019190915260448201526064015f604051808303818588803b158015613ebf575f80fd5b505af1158015613ed1573d5f803e3d5ffd5b5050505050610b37565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc1196001600160a01b0385160161400d5760405163095ea7b360e01b8152737ac440cae8eb6328de4fa621163a792c1ea9d4fe6004820152602481018390526001600160a01b0385169063095ea7b3906044016020604051808303815f875af1158015613f5e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613f829190614e80565b50737ac440cae8eb6328de4fa621163a792c1ea9d4fe63f219fa668585855a6040516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152604483015260648201526084015f604051808303815f87803b158015613ff2575f80fd5b505af1158015614004573d5f803e3d5ffd5b50505050610b37565b60405163095ea7b360e01b8152737ac440cae8eb6328de4fa621163a792c1ea9d4fe6004820152602481018390526001600160a01b0385169063095ea7b3906044016020604051808303815f875af115801561406b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061408f9190614e80565b5073d8a791fe2be73eb6e6cf1eb0cb3f36adc9b3f8f963f219fa668585855a6040516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152604483015260648201526084015b5f604051808303815f87803b158015614100575f80fd5b505af1158015614112573d5f803e3d5ffd5b5050505050505050565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeed196001600160a01b0384160161419a5760405163cd58657960e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede9063cd58657990839061418390600190879084908a905f90600401614efe565b5f604051808303818588803b158015614100575f80fd5b60405163095ea7b360e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede6004820152602481018290526001600160a01b0384169063095ea7b3906044016020604051808303815f875af11580156141f8573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061421c9190614e80565b5060405163cd58657960e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede9063cd5865799061425d906001908690869089905f90600401614efe565b5f604051808303815f87803b158015614274575f80fd5b505af1158015611519573d5f803e3d5ffd5b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeed196001600160a01b03841601614342577399c9fc46f92e8a1c0dec1b1747d010903e884be1639a2ac6d5650a064586222083855a6040516001600160e01b031960e087901b1681526001600160a01b03909216600483015263ffffffff166024820152606060448201525f60648201526084015f604051808303818589803b158015614325575f80fd5b5088f1158015614337573d5f803e3d5ffd5b505050505050505050565b5f61434c84614573565b60405163095ea7b360e01b81526010602160991b016004820152602481018490529091506001600160a01b0385169063095ea7b3906044016020604051808303815f875af11580156143a0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906143c49190614e80565b507399c9fc46f92e8a1c0dec1b1747d010903e884be163838b2520858386865a6040516001600160e01b031960e088901b1681526001600160a01b0395861660048201529385166024850152939091166044830152606482015263ffffffff91909116608482015260c060a48201525f60c482015260e4016140e9565b61444b8282611dbb565b6117b15760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610f03565b5f61448e6001600160a01b03841683614690565b905080515f141580156144b25750808060200190518101906144b09190614e80565b155b15610e9e57604051635274afe760e01b81526001600160a01b0384166004820152602401610f03565b5f805160206150118339815191525460ff16611b3857604051638dfc202b60e01b815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16611b3857604051631afcd79f60e31b815260040160405180910390fd5b61455b61450a565b5f80516020615011833981519152805460ff19169055565b5f73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc1196001600160a01b038316016145a757506006602160991b01919050565b73dac17f958d2ee523a2206206994597c13d831ec6196001600160a01b038316016145e757507394b008aa00579c1307b0ef2c499ad98a8ce58e58919050565b73a0b86991c6218b36c1d19d4a2e9eb0ce3606eb47196001600160a01b038316016146275750730b2c639c533813f4aa9d7837caf62653d097ff85919050565b736b175474e89094c44da98b954eedeac495271d0e196001600160a01b03831601614667575073da10009cbd5d07dd0cecc66161fc93d7c9000da1919050565b6040516305fd61ad60e01b81526001600160a01b0383166004820152602401610f03565b919050565b6060612cb183835f845f80856001600160a01b031684866040516146b49190614f44565b5f6040518083038185875af1925050503d805f81146146ee576040519150601f19603f3d011682016040523d82523d5f602084013e6146f3565b606091505b509150915061470386838361470d565b9695505050505050565b6060826147225761471d82614769565b612cb1565b815115801561473957506001600160a01b0384163b155b1561476257604051639996b31560e01b81526001600160a01b0385166004820152602401610f03565b5080612cb1565b8051156147795780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f602082840312156147a2575f80fd5b81356001600160e01b031981168114612cb1575f80fd5b6001600160a01b038116811461137c575f80fd5b5f80604083850312156147de575f80fd5b82356147e9816147b9565b946020939093013593505050565b5f60208284031215614807575f80fd5b8135612cb1816147b9565b5f60208284031215614822575f80fd5b5035919050565b5f806040838503121561483a575f80fd5b82359150602083013561484c816147b9565b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b60405160e0810167ffffffffffffffff8111828210171561488e5761488e614857565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156148bd576148bd614857565b604052919050565b803563ffffffff8116811461468b575f80fd5b801515811461137c575f80fd5b5f60208083850312156148f6575f80fd5b823567ffffffffffffffff8082111561490d575f80fd5b818501915085601f830112614920575f80fd5b81358181111561493257614932614857565b614940848260051b01614894565b818152848101925060e091820284018501918883111561495e575f80fd5b938501935b828510156149ef5780858a03121561497a575f8081fd5b61498261486b565b61498b866148c5565b81526149988787016148c5565b878201526040808701356149ab816147b9565b90820152606086810135908201526080808701359082015260a0808701359082015260c0808701356149dc816148d8565b9082015284529384019392850192614963565b50979650505050505050565b5f805f805f8060c08789031215614a10575f80fd5b86359550602087013594506040870135614a29816147b9565b959894975094956060810135955060808101359460a0909101359350915050565b5f805f805f60a08688031215614a5e575f80fd5b85359450602086013593506040860135614a77816147b9565b92506060860135614a87816147b9565b949793965091946080013592915050565b5f8060408385031215614aa9575f80fd5b8235614ab4816147b9565b9150602083013561484c816147b9565b8051151582526020808201516001600160a01b0316908301526040808201519083015260608082015190830152608090810151910152565b602080825282518282018190525f9190848201906040850190845b81811015614b3d57614b2a838551614ac4565b9284019260a09290920191600101614b17565b50909695505050505050565b5f805f60608486031215614b5b575f80fd5b8335614b66816147b9565b92506020840135614b76816147b9565b929592945050506040919091013590565b5f805f805f805f60e0888a031215614b9d575f80fd5b87359650602088013595506040880135614bb6816147b9565b94506060880135614bc6816147b9565b9699959850939660808101359560a0820135955060c0909101359350915050565b5f805f60608486031215614bf9575f80fd5b83359250602084013591506040840135614c12816147b9565b809150509250925092565b5f8060408385031215614c2e575f80fd5b8235614c39816147b9565b9150602083013561484c816148d8565b5f805f8060808587031215614c5c575f80fd5b84359350602085013592506040850135614c75816147b9565b9396929550929360600135925050565b5f8060408385031215614c96575f80fd5b82359150602083013561484c816148d8565b5f805f8060808587031215614cbb575f80fd5b843593506020850135614ccd816147b9565b92506040850135614c75816147b9565b5f805f60608486031215614cef575f80fd5b8335614cfa816147b9565b92506020840135614d0a816148d8565b9150614d18604085016148c5565b90509250925092565b60a08101610a9f8284614ac4565b634e487b7160e01b5f52601160045260245ffd5b81810381811115610a9f57610a9f614d2f565b634e487b7160e01b5f52603260045260245ffd5b63ffffffff818116838216019080821115613bb157613bb1614d2f565b5f60018201614d9857614d98614d2f565b5060010190565b95865260208601949094526001600160a01b039290921660408501526060840152608083015260a082015260c00190565b5f60208284031215614de0575f80fd5b5051919050565b80820180821115610a9f57610a9f614d2f565b8082028115828204841417610a9f57610a9f614d2f565b5f82614e2b57634e487b7160e01b5f52601260045260245ffd5b500490565b94855260208501939093526001600160a01b039190911660408401526060830152608082015260a00190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b5f60208284031215614e90575f80fd5b8151612cb1816148d8565b634e487b7160e01b5f52603160045260245ffd5b5f600160ff1b8201614ec357614ec3614d2f565b505f190190565b5f6001600160ff1b018201614d9857614d98614d2f565b63ffffffff828116828216039080821115613bb157613bb1614d2f565b63ffffffff9590951685526001600160a01b039384166020860152604085019290925290911660608301521515608082015260c060a082018190525f9082015260e00190565b5f82515f5b81811015614f635760208186018101518583015201614f49565b505f92019182525091905056fea1829a9003092132f585b6ccdd167c19fe9774dbdea4260287e8a8e8ca8185d7d0bcf4df132c65dad73803c5e5e1c826f151a3342680034a8a4c8e5f8eb0c13c3e65e660a0d3d61f62bb0309259c5a3ded6558e90d0e8aff997e553e7b030a7533fe247d78fee2e7fd135c405eda4bd2911c0a73c0a81b36c3bcc967dd06e5ae02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220b5e943765b0cd576f8f2b66b5b1818a1083e41a4a5669b468cf344e2e9e5895164736f6c63430008140033",
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

// IsSupportStableCoin is a free data retrieval call binding the contract method 0xed06da87.
//
// Solidity: function IsSupportStableCoin(address ERC20Address) view returns(bool)
func (_L1PoolManager *L1PoolManagerCaller) IsSupportStableCoin(opts *bind.CallOpts, ERC20Address common.Address) (bool, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "IsSupportStableCoin", ERC20Address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportStableCoin is a free data retrieval call binding the contract method 0xed06da87.
//
// Solidity: function IsSupportStableCoin(address ERC20Address) view returns(bool)
func (_L1PoolManager *L1PoolManagerSession) IsSupportStableCoin(ERC20Address common.Address) (bool, error) {
	return _L1PoolManager.Contract.IsSupportStableCoin(&_L1PoolManager.CallOpts, ERC20Address)
}

// IsSupportStableCoin is a free data retrieval call binding the contract method 0xed06da87.
//
// Solidity: function IsSupportStableCoin(address ERC20Address) view returns(bool)
func (_L1PoolManager *L1PoolManagerCallerSession) IsSupportStableCoin(ERC20Address common.Address) (bool, error) {
	return _L1PoolManager.Contract.IsSupportStableCoin(&_L1PoolManager.CallOpts, ERC20Address)
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
// Solidity: function Users(address , uint256 ) view returns(bool isClaimed, address token, uint256 StartPoolId, uint256 EndPoolId, uint256 Amount)
func (_L1PoolManager *L1PoolManagerCaller) Users(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	IsClaimed   bool
	Token       common.Address
	StartPoolId *big.Int
	EndPoolId   *big.Int
	Amount      *big.Int
}, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "Users", arg0, arg1)

	outstruct := new(struct {
		IsClaimed   bool
		Token       common.Address
		StartPoolId *big.Int
		EndPoolId   *big.Int
		Amount      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsClaimed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StartPoolId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndPoolId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0x13e8544e.
//
// Solidity: function Users(address , uint256 ) view returns(bool isClaimed, address token, uint256 StartPoolId, uint256 EndPoolId, uint256 Amount)
func (_L1PoolManager *L1PoolManagerSession) Users(arg0 common.Address, arg1 *big.Int) (struct {
	IsClaimed   bool
	Token       common.Address
	StartPoolId *big.Int
	EndPoolId   *big.Int
	Amount      *big.Int
}, error) {
	return _L1PoolManager.Contract.Users(&_L1PoolManager.CallOpts, arg0, arg1)
}

// Users is a free data retrieval call binding the contract method 0x13e8544e.
//
// Solidity: function Users(address , uint256 ) view returns(bool isClaimed, address token, uint256 StartPoolId, uint256 EndPoolId, uint256 Amount)
func (_L1PoolManager *L1PoolManagerCallerSession) Users(arg0 common.Address, arg1 *big.Int) (struct {
	IsClaimed   bool
	Token       common.Address
	StartPoolId *big.Int
	EndPoolId   *big.Int
	Amount      *big.Int
}, error) {
	return _L1PoolManager.Contract.Users(&_L1PoolManager.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.Balances(&_L1PoolManager.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.Balances(&_L1PoolManager.CallOpts, arg0)
}

// GetPool is a free data retrieval call binding the contract method 0x5b5b9ea2.
//
// Solidity: function getPool(address _token, uint256 _index) view returns((uint32,uint32,address,uint256,uint256,uint256,bool))
func (_L1PoolManager *L1PoolManagerCaller) GetPool(opts *bind.CallOpts, _token common.Address, _index *big.Int) (IL1PoolPool, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getPool", _token, _index)

	if err != nil {
		return *new(IL1PoolPool), err
	}

	out0 := *abi.ConvertType(out[0], new(IL1PoolPool)).(*IL1PoolPool)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x5b5b9ea2.
//
// Solidity: function getPool(address _token, uint256 _index) view returns((uint32,uint32,address,uint256,uint256,uint256,bool))
func (_L1PoolManager *L1PoolManagerSession) GetPool(_token common.Address, _index *big.Int) (IL1PoolPool, error) {
	return _L1PoolManager.Contract.GetPool(&_L1PoolManager.CallOpts, _token, _index)
}

// GetPool is a free data retrieval call binding the contract method 0x5b5b9ea2.
//
// Solidity: function getPool(address _token, uint256 _index) view returns((uint32,uint32,address,uint256,uint256,uint256,bool))
func (_L1PoolManager *L1PoolManagerCallerSession) GetPool(_token common.Address, _index *big.Int) (IL1PoolPool, error) {
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
func (_L1PoolManager *L1PoolManagerCaller) GetUser(opts *bind.CallOpts, _user common.Address) ([]IL1PoolUser, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getUser", _user)

	if err != nil {
		return *new([]IL1PoolUser), err
	}

	out0 := *abi.ConvertType(out[0], new([]IL1PoolUser)).(*[]IL1PoolUser)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _user) view returns((bool,address,uint256,uint256,uint256)[])
func (_L1PoolManager *L1PoolManagerSession) GetUser(_user common.Address) ([]IL1PoolUser, error) {
	return _L1PoolManager.Contract.GetUser(&_L1PoolManager.CallOpts, _user)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _user) view returns((bool,address,uint256,uint256,uint256)[])
func (_L1PoolManager *L1PoolManagerCallerSession) GetUser(_user common.Address) ([]IL1PoolUser, error) {
	return _L1PoolManager.Contract.GetUser(&_L1PoolManager.CallOpts, _user)
}

// GetUser0 is a free data retrieval call binding the contract method 0xff2bf64f.
//
// Solidity: function getUser(address _user, uint256 _index) view returns((bool,address,uint256,uint256,uint256))
func (_L1PoolManager *L1PoolManagerCaller) GetUser0(opts *bind.CallOpts, _user common.Address, _index *big.Int) (IL1PoolUser, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getUser0", _user, _index)

	if err != nil {
		return *new(IL1PoolUser), err
	}

	out0 := *abi.ConvertType(out[0], new(IL1PoolUser)).(*IL1PoolUser)

	return out0, err

}

// GetUser0 is a free data retrieval call binding the contract method 0xff2bf64f.
//
// Solidity: function getUser(address _user, uint256 _index) view returns((bool,address,uint256,uint256,uint256))
func (_L1PoolManager *L1PoolManagerSession) GetUser0(_user common.Address, _index *big.Int) (IL1PoolUser, error) {
	return _L1PoolManager.Contract.GetUser0(&_L1PoolManager.CallOpts, _user, _index)
}

// GetUser0 is a free data retrieval call binding the contract method 0xff2bf64f.
//
// Solidity: function getUser(address _user, uint256 _index) view returns((bool,address,uint256,uint256,uint256))
func (_L1PoolManager *L1PoolManagerCallerSession) GetUser0(_user common.Address, _index *big.Int) (IL1PoolUser, error) {
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

// ClaimAll is a paid mutator transaction binding the contract method 0x6d6a78bf.
//
// Solidity: function ClaimAll() returns()
func (_L1PoolManager *L1PoolManagerTransactor) ClaimAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "ClaimAll")
}

// ClaimAll is a paid mutator transaction binding the contract method 0x6d6a78bf.
//
// Solidity: function ClaimAll() returns()
func (_L1PoolManager *L1PoolManagerSession) ClaimAll() (*types.Transaction, error) {
	return _L1PoolManager.Contract.ClaimAll(&_L1PoolManager.TransactOpts)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x6d6a78bf.
//
// Solidity: function ClaimAll() returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) ClaimAll() (*types.Transaction, error) {
	return _L1PoolManager.Contract.ClaimAll(&_L1PoolManager.TransactOpts)
}

// ClaimSimpleAsset is a paid mutator transaction binding the contract method 0x95455eef.
//
// Solidity: function ClaimSimpleAsset(address _token) returns()
func (_L1PoolManager *L1PoolManagerTransactor) ClaimSimpleAsset(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "ClaimSimpleAsset", _token)
}

// ClaimSimpleAsset is a paid mutator transaction binding the contract method 0x95455eef.
//
// Solidity: function ClaimSimpleAsset(address _token) returns()
func (_L1PoolManager *L1PoolManagerSession) ClaimSimpleAsset(_token common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.ClaimSimpleAsset(&_L1PoolManager.TransactOpts, _token)
}

// ClaimSimpleAsset is a paid mutator transaction binding the contract method 0x95455eef.
//
// Solidity: function ClaimSimpleAsset(address _token) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) ClaimSimpleAsset(_token common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.ClaimSimpleAsset(&_L1PoolManager.TransactOpts, _token)
}

// CompletePoolAndNew is a paid mutator transaction binding the contract method 0x3338562c.
//
// Solidity: function CompletePoolAndNew((uint32,uint32,address,uint256,uint256,uint256,bool)[] CompletePools) payable returns()
func (_L1PoolManager *L1PoolManagerTransactor) CompletePoolAndNew(opts *bind.TransactOpts, CompletePools []IL1PoolPool) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "CompletePoolAndNew", CompletePools)
}

// CompletePoolAndNew is a paid mutator transaction binding the contract method 0x3338562c.
//
// Solidity: function CompletePoolAndNew((uint32,uint32,address,uint256,uint256,uint256,bool)[] CompletePools) payable returns()
func (_L1PoolManager *L1PoolManagerSession) CompletePoolAndNew(CompletePools []IL1PoolPool) (*types.Transaction, error) {
	return _L1PoolManager.Contract.CompletePoolAndNew(&_L1PoolManager.TransactOpts, CompletePools)
}

// CompletePoolAndNew is a paid mutator transaction binding the contract method 0x3338562c.
//
// Solidity: function CompletePoolAndNew((uint32,uint32,address,uint256,uint256,uint256,bool)[] CompletePools) payable returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) CompletePoolAndNew(CompletePools []IL1PoolPool) (*types.Transaction, error) {
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

// SendAssertToUser is a paid mutator transaction binding the contract method 0x73c23be9.
//
// Solidity: function SendAssertToUser(address _token, address to, uint256 _amount) payable returns(bool)
func (_L1PoolManager *L1PoolManagerTransactor) SendAssertToUser(opts *bind.TransactOpts, _token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "SendAssertToUser", _token, to, _amount)
}

// SendAssertToUser is a paid mutator transaction binding the contract method 0x73c23be9.
//
// Solidity: function SendAssertToUser(address _token, address to, uint256 _amount) payable returns(bool)
func (_L1PoolManager *L1PoolManagerSession) SendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SendAssertToUser(&_L1PoolManager.TransactOpts, _token, to, _amount)
}

// SendAssertToUser is a paid mutator transaction binding the contract method 0x73c23be9.
//
// Solidity: function SendAssertToUser(address _token, address to, uint256 _amount) payable returns(bool)
func (_L1PoolManager *L1PoolManagerTransactorSession) SendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SendAssertToUser(&_L1PoolManager.TransactOpts, _token, to, _amount)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0xfaeb96a5.
//
// Solidity: function SetSupportToken(address _token, bool _isSupport, uint32 startTimes) returns()
func (_L1PoolManager *L1PoolManagerTransactor) SetSupportToken(opts *bind.TransactOpts, _token common.Address, _isSupport bool, startTimes uint32) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "SetSupportToken", _token, _isSupport, startTimes)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0xfaeb96a5.
//
// Solidity: function SetSupportToken(address _token, bool _isSupport, uint32 startTimes) returns()
func (_L1PoolManager *L1PoolManagerSession) SetSupportToken(_token common.Address, _isSupport bool, startTimes uint32) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetSupportToken(&_L1PoolManager.TransactOpts, _token, _isSupport, startTimes)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0xfaeb96a5.
//
// Solidity: function SetSupportToken(address _token, bool _isSupport, uint32 startTimes) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) SetSupportToken(_token common.Address, _isSupport bool, startTimes uint32) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetSupportToken(&_L1PoolManager.TransactOpts, _token, _isSupport, startTimes)
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

// SetSupportStableCoin is a paid mutator transaction binding the contract method 0xb33280dd.
//
// Solidity: function setSupportStableCoin(address ERC20Address, bool isValid) returns()
func (_L1PoolManager *L1PoolManagerTransactor) SetSupportStableCoin(opts *bind.TransactOpts, ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "setSupportStableCoin", ERC20Address, isValid)
}

// SetSupportStableCoin is a paid mutator transaction binding the contract method 0xb33280dd.
//
// Solidity: function setSupportStableCoin(address ERC20Address, bool isValid) returns()
func (_L1PoolManager *L1PoolManagerSession) SetSupportStableCoin(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetSupportStableCoin(&_L1PoolManager.TransactOpts, ERC20Address, isValid)
}

// SetSupportStableCoin is a paid mutator transaction binding the contract method 0xb33280dd.
//
// Solidity: function setSupportStableCoin(address ERC20Address, bool isValid) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) SetSupportStableCoin(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetSupportStableCoin(&_L1PoolManager.TransactOpts, ERC20Address, isValid)
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
