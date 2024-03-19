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

// PolygonzkevmbridgeMetaData contains all meta data concerning the Polygonzkevmbridge contract.
var PolygonzkevmbridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ChainIdIsNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"ChainIdNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorBlockChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"MinTransferAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LessThanMinTransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MantaNotWETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MantleNotWETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughETH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"NotEnoughToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"TokenIsNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"sourceChainIdError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20toL1Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawETHtoL1Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawWETHtoL1Success\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"BridgeInitiateETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FeePoolValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FundingPoolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"IsSupportChainId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"IsSupportToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_GAS_Limit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinTransferAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"QuickSendAssertToUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ReLayer\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SupportTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UpdateFundingPoolBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20ToL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawETHtoL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawWETHToL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MultisigWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageManager\",\"outputs\":[{\"internalType\":\"contractIMessageManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_MinTransferAmount\",\"type\":\"uint256\"}],\"name\":\"setMinTransferAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_PerFee\",\"type\":\"uint256\"}],\"name\":\"setPerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setSupportERC20Token\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setValidChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615c1c806100206000396000f3fe608060405260043610620001db5760003560e01c8063787e05da11620000ff578063be5831c71162000095578063dbc16976116200006c578063dbc169761462000713578063ee25560b146200072b578063f851a440146200075c578063fb570834146200078b57600080fd5b8063be5831c71462000688578063cd58657914620006c4578063d02103ca14620006db57600080fd5b806382b12dd711620000d657806382b12dd714620005cc5780639e34070f14620005e4578063aaa13cc21462000629578063bab161bf146200064e57600080fd5b8063787e05da146200054857806379e2cf97146200056d57806381b1c174146200058557600080fd5b806334ac9cf211620001755780633e197043116200014c5780633e19704314620003e85780634127535814620004dc5780635a278b3c146200050b57806363a9bf82146200052357600080fd5b806334ac9cf214620003725780633ae0504714620003a15780633c351e1014620003b957600080fd5b806327aef4e811620001b657806327aef4e814620002765780632cffd02e146200029d5780632dfdf0b514620002c2578063318aee3d14620002e957600080fd5b806315064c9614620001e05780632072f6c5146200021157806322e95f2c146200022b575b600080fd5b348015620001ed57600080fd5b50606854620001fc9060ff1681565b60405190151581526020015b60405180910390f35b3480156200021e57600080fd5b5062000229620007b0565b005b3480156200023857600080fd5b50620002506200024a3660046200323a565b6200080e565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200162000208565b3480156200028357600080fd5b506200028e620008b1565b604051620002089190620032e8565b348015620002aa57600080fd5b5062000229620002bc36600462003363565b62000947565b348015620002cf57600080fd5b50620002da60535481565b60405190815260200162000208565b348015620002f657600080fd5b50620003406200030836600462003449565b606b6020526000908152604090205463ffffffff811690640100000000900473ffffffffffffffffffffffffffffffffffffffff1682565b6040805163ffffffff909316835273ffffffffffffffffffffffffffffffffffffffff90911660208301520162000208565b3480156200037f57600080fd5b50606c54620002509073ffffffffffffffffffffffffffffffffffffffff1681565b348015620003ae57600080fd5b50620002da62000eeb565b348015620003c657600080fd5b50607054620002509073ffffffffffffffffffffffffffffffffffffffff1681565b348015620003f557600080fd5b50620002da6200040736600462003479565b6040517fff0000000000000000000000000000000000000000000000000000000000000060f889901b1660208201527fffffffff0000000000000000000000000000000000000000000000000000000060e088811b821660218401527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606089811b821660258601529188901b909216603984015285901b16603d8201526051810183905260718101829052600090609101604051602081830303815290604052805190602001209050979650505050505050565b348015620004e957600080fd5b50606f54620002509073ffffffffffffffffffffffffffffffffffffffff1681565b3480156200051857600080fd5b50620002da60725481565b3480156200053057600080fd5b50620002296200054236600462003629565b62000fd1565b3480156200055557600080fd5b506200022962000567366004620036f8565b620014ec565b3480156200057a57600080fd5b5062000229620015ad565b3480156200059257600080fd5b5062000250620005a436600462003727565b606a6020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b348015620005d957600080fd5b50620002da606e5481565b348015620005f157600080fd5b50620001fc6200060336600462003727565b600881901c600090815260696020526040902054600160ff9092169190911b9081161490565b3480156200063657600080fd5b50620002506200064836600462003741565b620015e7565b3480156200065b57600080fd5b506068546200067290610100900463ffffffff1681565b60405163ffffffff909116815260200162000208565b3480156200069557600080fd5b506068546200067290790100000000000000000000000000000000000000000000000000900463ffffffff1681565b62000229620006d536600462003804565b620017d3565b348015620006e857600080fd5b50606854620002509065010000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b3480156200072057600080fd5b506200022962001fca565b3480156200073857600080fd5b50620002da6200074a36600462003727565b60696020526000908152604090205481565b3480156200076957600080fd5b50606d54620002509073ffffffffffffffffffffffffffffffffffffffff1681565b3480156200079857600080fd5b50620001fc620007aa366004620038a6565b62002026565b606c5473ffffffffffffffffffffffffffffffffffffffff16331462000802576040517fe2e8106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6200080c6200210f565b565b6040805160e084901b7fffffffff0000000000000000000000000000000000000000000000000000000016602080830191909152606084901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602483015282516018818403018152603890920183528151918101919091206000908152606a909152205473ffffffffffffffffffffffffffffffffffffffff165b92915050565b60718054620008c090620038f3565b80601f0160208091040260200160405190810160405280929190818152602001828054620008ee90620038f3565b80156200093f5780601f1062000913576101008083540402835291602001916200093f565b820191906000526020600020905b8154815290600101906020018083116200092157829003601f168201915b505050505081565b60685460ff161562000985576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6200099c8b8b8b8b8b8b8b8b8b8b8b6000620021a3565b73ffffffffffffffffffffffffffffffffffffffff861662000a78576040805160008082526020820190925273ffffffffffffffffffffffffffffffffffffffff8616908590604051620009f1919062003948565b60006040518083038185875af1925050503d806000811462000a30576040519150601f19603f3d011682016040523d82523d6000602084013e62000a35565b606091505b505090508062000a71576040517f6747a28800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5062000e73565b60685463ffffffff61010090910481169088160362000aba5762000ab473ffffffffffffffffffffffffffffffffffffffff8716858562002471565b62000e73565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e089901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606088901b166024820152600090603801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152606a90935291205490915073ffffffffffffffffffffffffffffffffffffffff168062000de5576000808062000b8f8688018862003966565b92509250925060008584848460405162000ba990620031f1565b62000bb793929190620039e7565b8190604051809103906000f590508015801562000bd8573d6000803e3d6000fd5b506040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c81166004830152602482018c9052919250908216906340c10f1990604401600060405180830381600087803b15801562000c4e57600080fd5b505af115801562000c63573d6000803e3d6000fd5b5050505080606a600088815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405180604001604052808e63ffffffff1681526020018d73ffffffffffffffffffffffffffffffffffffffff16815250606b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548163ffffffff021916908363ffffffff16021790555060208201518160000160046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050507f490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a398d8d838b8b60405162000dd395949392919062003a6d565b60405180910390a15050505062000e70565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152602482018790528216906340c10f1990604401600060405180830381600087803b15801562000e5657600080fd5b505af115801562000e6b573d6000803e3d6000fd5b505050505b50505b6040805163ffffffff8c811682528916602082015273ffffffffffffffffffffffffffffffffffffffff88811682840152861660608201526080810185905290517f25308c93ceeed162da955b3f7ce3e3f93606579e40fb92029faa9efe275459839181900360a00190a15050505050505050505050565b605354600090819081805b602081101562000fc8578083901c60011660010362000f59576033816020811062000f255762000f2562003ab5565b0154604080516020810192909252810185905260600160405160208183030381529060405280519060200120935062000f86565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b6040805160208101849052908101839052606001604051602081830303815290604052805190602001209150808062000fbf9062003b13565b91505062000ef6565b50919392505050565b8673ffffffffffffffffffffffffffffffffffffffff811662001055576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f496c6c6567616c2061646472657373000000000000000000000000000000000060448201526064015b60405180910390fd5b8673ffffffffffffffffffffffffffffffffffffffff8116620010d5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f496c6c6567616c2061646472657373000000000000000000000000000000000060448201526064016200104c565b8573ffffffffffffffffffffffffffffffffffffffff811662001155576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f496c6c6567616c2061646472657373000000000000000000000000000000000060448201526064016200104c565b600054610100900460ff1615808015620011765750600054600160ff909116105b80620011925750303b15801562001192575060005460ff166001145b62001220576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016200104c565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156200127f57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b60008511620012eb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600360248201527f494446000000000000000000000000000000000000000000000000000000000060448201526064016200104c565b8c606860016101000a81548163ffffffff021916908363ffffffff1602179055508b606860056101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508a606c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555089606d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555088606e8190555087606f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555086607060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550856071908162001469919062003ba0565b506072859055620014796200254c565b8015620014dd57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050505050505050565b606d5473ffffffffffffffffffffffffffffffffffffffff1633146200153e576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216156200159c57606f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b8015620015a957606e8190555b5050565b605354606854790100000000000000000000000000000000000000000000000000900463ffffffff1610156200080c576200080c620025ef565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e089901b1660208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606088901b1660248201526000908190603801604051602081830303815290604052805190602001209050600060ff60f81b3083604051806020016200167d90620031f1565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820381018352601f909101166040819052620016c8908d908d908d908d908d9060200162003cc7565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905262001706929160200162003d08565b604051602081830303815290604052805190602001206040516020016200178f94939291907fff0000000000000000000000000000000000000000000000000000000000000094909416845260609290921b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660018401526015830152603582015260550190565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815291905280516020909101209a9950505050505050505050565b60685460ff161562001811576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6200181b620026ca565b60685463ffffffff888116610100909204161480620018415750600263ffffffff881610155b1562001879576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008060608773ffffffffffffffffffffffffffffffffffffffff8816620018ee5788606e5434620018ac919062003d3b565b14620018e4576040517fb89240f500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000925062001bd6565b606e5434146200192a576040517fb89240f500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8089166000908152606b602090815260409182902082518084019093525463ffffffff811683526401000000009004909216918101829052901562001a1a576040517f9dc29fac000000000000000000000000000000000000000000000000000000008152336004820152602481018b905273ffffffffffffffffffffffffffffffffffffffff8a1690639dc29fac90604401600060405180830381600087803b158015620019ed57600080fd5b505af115801562001a02573d6000803e3d6000fd5b50505050806020015194508060000151935062001bd4565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009073ffffffffffffffffffffffffffffffffffffffff8b16906370a0823190602401602060405180830381865afa15801562001a88573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062001aae919062003d51565b905062001ad473ffffffffffffffffffffffffffffffffffffffff8b1633308e6200273f565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009073ffffffffffffffffffffffffffffffffffffffff8c16906370a0823190602401602060405180830381865afa15801562001b42573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062001b68919062003d51565b905062001b76828262003d3b565b6068548c9850610100900463ffffffff169650935062001b96876200279f565b62001ba18c620028b7565b62001bac8d620029c4565b60405160200162001bc093929190620039e7565b604051602081830303815290604052945050505b505b60705473ffffffffffffffffffffffffffffffffffffffff161562001d785773ffffffffffffffffffffffffffffffffffffffff881662001d2b576070546071805473ffffffffffffffffffffffffffffffffffffffff90921695509062001c3e90620038f3565b80601f016020809104026020016040519081016040528092919081815260200182805462001c6c90620038f3565b801562001cbd5780601f1062001c915761010080835404028352916020019162001cbd565b820191906000526020600020905b81548152906001019060200180831162001c9f57829003601f168201915b5050606854939550505050610100900463ffffffff161562001d255760725462001ce8908262003d6b565b90508060000362001d25576040517fc2f5625a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62001d78565b60705473ffffffffffffffffffffffffffffffffffffffff9081169085160362001d785760685460009450610100900463ffffffff1662001d785760725462001d75908262003da7565b90505b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b600084868e8e868860535460405162001dba98979695949392919062003dc1565b60405180910390a162001eb462001eae600085878f8f8789805190602001206040517fff0000000000000000000000000000000000000000000000000000000000000060f889901b1660208201527fffffffff0000000000000000000000000000000000000000000000000000000060e088811b821660218401527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606089811b821660258601529188901b909216603984015285901b16603d8201526051810183905260718101829052600090609101604051602081830303815290604052805190602001209050979650505050505050565b62002ac0565b606f5473ffffffffffffffffffffffffffffffffffffffff161580159062001ede57506000606e54115b1562001fa257606f54606e5460408051600080825260208201928390529373ffffffffffffffffffffffffffffffffffffffff16929162001f20919062003948565b60006040518083038185875af1925050503d806000811462001f5f576040519150601f19603f3d011682016040523d82523d6000602084013e62001f64565b606091505b505090508062001fa0576040517f6747a28800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b861562001fb35762001fb3620025ef565b5050505062001fc160018055565b50505050505050565b606c5473ffffffffffffffffffffffffffffffffffffffff1633146200201c576040517fe2e8106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6200080c62002bd8565b600084815b60208110156200210157600163ffffffff8616821c811690036200209d578581602081106200205e576200205e62003ab5565b6020020135826040516020016200207f929190918252602082015260400190565b604051602081830303815290604052805190602001209150620020ec565b81868260208110620020b357620020b362003ab5565b6020020135604051602001620020d3929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b80620020f88162003b13565b9150506200202b565b50821490505b949350505050565b60685460ff16156200214d576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a549790600090a1565b620021b48b63ffffffff1662002c68565b6068546040805160208082018e90528183018d9052825180830384018152606083019384905280519101207f257b363200000000000000000000000000000000000000000000000000000000909252606481019190915260009165010000000000900473ffffffffffffffffffffffffffffffffffffffff169063257b3632906084016020604051808303816000875af115801562002257573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200227d919062003d51565b905080600003620022b9576040517e2f6fad00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60685463ffffffff888116610100909204161462002303576040517f0595ea2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606854600090610100900463ffffffff166200232157508962002324565b508a5b6200242a62002421848c8c8c8c8c8c8c6040516200234492919062003e3a565b60405180910390206040517fff0000000000000000000000000000000000000000000000000000000000000060f889901b1660208201527fffffffff0000000000000000000000000000000000000000000000000000000060e088811b821660218401527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606089811b821660258601529188901b909216603984015285901b16603d8201526051810183905260718101829052600090609101604051602081830303815290604052805190602001209050979650505050505050565b8f8f8462002026565b62002461576040517fe0417cec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050505050505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052620025479084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915262002cd4565b505050565b600054610100900460ff16620025e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016200104c565b6200080c62002de7565b6053546068805463ffffffff909216790100000000000000000000000000000000000000000000000000027fffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffffff909216919091179081905573ffffffffffffffffffffffffffffffffffffffff65010000000000909104166333d6247d6200267562000eeb565b6040518263ffffffff1660e01b81526004016200269491815260200190565b600060405180830381600087803b158015620026af57600080fd5b505af1158015620026c4573d6000803e3d6000fd5b50505050565b60026001540362002738576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016200104c565b6002600155565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052620026c49085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401620024c4565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f06fdde03000000000000000000000000000000000000000000000000000000001790529051606091600091829173ffffffffffffffffffffffffffffffffffffffff86169162002823919062003948565b600060405180830381855afa9150503d806000811462002860576040519150601f19603f3d011682016040523d82523d6000602084013e62002865565b606091505b509150915081620028ac576040518060400160405280600781526020017f4e4f5f4e414d450000000000000000000000000000000000000000000000000081525062002107565b620021078162002e80565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f95d89b41000000000000000000000000000000000000000000000000000000001790529051606091600091829173ffffffffffffffffffffffffffffffffffffffff8616916200293b919062003948565b600060405180830381855afa9150503d806000811462002978576040519150601f19603f3d011682016040523d82523d6000602084013e6200297d565b606091505b509150915081620028ac576040518060400160405280600981526020017f4e4f5f53594d424f4c000000000000000000000000000000000000000000000081525062002107565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f313ce5670000000000000000000000000000000000000000000000000000000017905290516000918291829173ffffffffffffffffffffffffffffffffffffffff86169162002a47919062003948565b600060405180830381855afa9150503d806000811462002a84576040519150601f19603f3d011682016040523d82523d6000602084013e62002a89565b606091505b509150915081801562002a9d575080516020145b62002aaa57601262002107565b8080602001905181019062002107919062003e4a565b80600162002ad16020600262003fa3565b62002add919062003d3b565b6053541062002b18576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060536000815462002b2b9062003b13565b9182905550905060005b602081101562002bc7578082901c60011660010362002b6d57826033826020811062002b655762002b6562003ab5565b015550505050565b6033816020811062002b835762002b8362003ab5565b01546040805160208101929092528101849052606001604051602081830303815290604052805190602001209250808062002bbe9062003b13565b91505062002b35565b506200254762003fb1565b60018055565b60685460ff1662002c15576040517f5386698100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b390600090a1565b600881901c60008181526069602052604081208054600160ff861690811b9182189283905592909190818316900362002ccd576040517f646cf55800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b600062002d38826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16620030709092919063ffffffff16565b80519091501562002547578080602001905181019062002d59919062003fe0565b62002547576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016200104c565b600054610100900460ff1662002bd2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016200104c565b6060604082511062002ea25781806020019051810190620008ab919062004000565b8151602003620030325760005b60208110801562002efa575082818151811062002ed05762002ed062003ab5565b01602001517fff000000000000000000000000000000000000000000000000000000000000001615155b1562002f15578062002f0c8162003b13565b91505062002eaf565b8060000362002f5957505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e4700000000000000000000000000006020820152919050565b60008167ffffffffffffffff81111562002f775762002f7762003503565b6040519080825280601f01601f19166020018201604052801562002fa2576020820181803683370190505b50905060005b828110156200302a5784818151811062002fc65762002fc662003ab5565b602001015160f81c60f81b82828151811062002fe65762002fe662003ab5565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080620030218162003b13565b91505062002fa8565b509392505050565b505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e470000000000000000000000000000602082015290565b919050565b6060620021078484600085856000808673ffffffffffffffffffffffffffffffffffffffff168587604051620030a7919062003948565b60006040518083038185875af1925050503d8060008114620030e6576040519150601f19603f3d011682016040523d82523d6000602084013e620030eb565b606091505b5091509150620030fe8783838762003109565b979650505050505050565b60608315620031a45782516000036200319c5773ffffffffffffffffffffffffffffffffffffffff85163b6200319c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016200104c565b508162002107565b620021078383815115620031bb5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200104c9190620032e8565b611b66806200408183390190565b803563ffffffff811681146200306b57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811681146200323757600080fd5b50565b600080604083850312156200324e57600080fd5b6200325983620031ff565b915060208301356200326b8162003214565b809150509250929050565b60005b838110156200329357818101518382015260200162003279565b50506000910152565b60008151808452620032b681602086016020860162003276565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000620032fd60208301846200329c565b9392505050565b806104008101831015620008ab57600080fd5b60008083601f8401126200332a57600080fd5b50813567ffffffffffffffff8111156200334357600080fd5b6020830191508360208285010111156200335c57600080fd5b9250929050565b60008060008060008060008060008060006105208c8e0312156200338657600080fd5b620033928d8d62003304565b9a50620033a36104008d01620031ff565b99506104208c013598506104408c01359750620033c46104608d01620031ff565b96506104808c0135620033d78162003214565b9550620033e86104a08d01620031ff565b94506104c08c0135620033fb8162003214565b93506104e08c013592506105008c013567ffffffffffffffff8111156200342157600080fd5b6200342f8e828f0162003317565b915080935050809150509295989b509295989b9093969950565b6000602082840312156200345c57600080fd5b8135620032fd8162003214565b60ff811681146200323757600080fd5b600080600080600080600060e0888a0312156200349557600080fd5b8735620034a28162003469565b9650620034b260208901620031ff565b95506040880135620034c48162003214565b9450620034d460608901620031ff565b93506080880135620034e68162003214565b9699959850939692959460a0840135945060c09093013592915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156200357c576200357c62003503565b604052919050565b600067ffffffffffffffff821115620035a157620035a162003503565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112620035df57600080fd5b8135620035f6620035f08262003584565b62003532565b8181528460208386010111156200360c57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060008060006101208a8c0312156200364957600080fd5b620036548a620031ff565b985060208a0135620036668162003214565b975060408a0135620036788162003214565b965060608a01356200368a8162003214565b955060808a0135945060a08a0135620036a38162003214565b935060c08a0135620036b58162003214565b925060e08a013567ffffffffffffffff811115620036d257600080fd5b620036e08c828d01620035cd565b9250506101008a013590509295985092959850929598565b600080604083850312156200370c57600080fd5b8235620037198162003214565b946020939093013593505050565b6000602082840312156200373a57600080fd5b5035919050565b600080600080600080600060a0888a0312156200375d57600080fd5b6200376888620031ff565b965060208801356200377a8162003214565b9550604088013567ffffffffffffffff808211156200379857600080fd5b620037a68b838c0162003317565b909750955060608a0135915080821115620037c057600080fd5b50620037cf8a828b0162003317565b9094509250506080880135620037e58162003469565b8091505092959891949750929550565b80151581146200323757600080fd5b600080600080600080600060c0888a0312156200382057600080fd5b6200382b88620031ff565b965060208801356200383d8162003214565b9550604088013594506060880135620038568162003214565b935060808801356200386881620037f5565b925060a088013567ffffffffffffffff8111156200388557600080fd5b620038938a828b0162003317565b989b979a50959850939692959293505050565b6000806000806104608587031215620038be57600080fd5b84359350620038d1866020870162003304565b9250620038e26104208601620031ff565b939692955092936104400135925050565b600181811c908216806200390857607f821691505b60208210810362003942577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b600082516200395c81846020870162003276565b9190910192915050565b6000806000606084860312156200397c57600080fd5b833567ffffffffffffffff808211156200399557600080fd5b620039a387838801620035cd565b94506020860135915080821115620039ba57600080fd5b50620039c986828701620035cd565b9250506040840135620039dc8162003469565b809150509250925092565b606081526000620039fc60608301866200329c565b828103602084015262003a1081866200329c565b91505060ff83166040830152949350505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b63ffffffff86168152600073ffffffffffffffffffffffffffffffffffffffff808716602084015280861660408401525060806060830152620030fe60808301848662003a24565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820362003b475762003b4762003ae4565b5060010190565b601f8211156200254757600081815260208120601f850160051c8101602086101562003b775750805b601f850160051c820191505b8181101562003b985782815560010162003b83565b505050505050565b815167ffffffffffffffff81111562003bbd5762003bbd62003503565b62003bd58162003bce8454620038f3565b8462003b4e565b602080601f83116001811462003c2b576000841562003bf45750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b17855562003b98565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101562003c7a5788860151825594840194600190910190840162003c59565b508582101562003cb757878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b60608152600062003cdd60608301878962003a24565b828103602084015262003cf281868862003a24565b91505060ff831660408301529695505050505050565b6000835162003d1c81846020880162003276565b83519083019062003d3281836020880162003276565b01949350505050565b81810381811115620008ab57620008ab62003ae4565b60006020828403121562003d6457600080fd5b5051919050565b60008262003da2577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b8082028115828204841417620008ab57620008ab62003ae4565b600061010060ff8b16835263ffffffff808b16602085015273ffffffffffffffffffffffffffffffffffffffff808b166040860152818a1660608601528089166080860152508660a08501528160c085015262003e21828501876200329c565b925080851660e085015250509998505050505050505050565b8183823760009101908152919050565b60006020828403121562003e5d57600080fd5b8151620032fd8162003469565b600181815b8085111562003ec957817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0482111562003ead5762003ead62003ae4565b8085161562003ebb57918102915b93841c939080029062003e6f565b509250929050565b60008262003ee257506001620008ab565b8162003ef157506000620008ab565b816001811462003f0a576002811462003f155762003f35565b6001915050620008ab565b60ff84111562003f295762003f2962003ae4565b50506001821b620008ab565b5060208310610133831016604e8410600b841016171562003f5a575081810a620008ab565b62003f66838362003e6a565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0482111562003f9b5762003f9b62003ae4565b029392505050565b6000620032fd838362003ed1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60006020828403121562003ff357600080fd5b8151620032fd81620037f5565b6000602082840312156200401357600080fd5b815167ffffffffffffffff8111156200402b57600080fd5b8201601f810184136200403d57600080fd5b80516200404e620035f08262003584565b8181528560208385010111156200406457600080fd5b6200407782602083016020860162003276565b9594505050505056fe6101006040523480156200001257600080fd5b5060405162001b6638038062001b6683398101604081905262000035916200028d565b82826003620000458382620003a1565b506004620000548282620003a1565b50503360c0525060ff811660e052466080819052620000739062000080565b60a052506200046d915050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f620000ad6200012e565b805160209182012060408051808201825260018152603160f81b90840152805192830193909352918101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc66060820152608081018390523060a082015260c001604051602081830303815290604052805190602001209050919050565b6060600380546200013f9062000312565b80601f01602080910402602001604051908101604052809291908181526020018280546200016d9062000312565b8015620001be5780601f106200019257610100808354040283529160200191620001be565b820191906000526020600020905b815481529060010190602001808311620001a057829003601f168201915b5050505050905090565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620001f057600080fd5b81516001600160401b03808211156200020d576200020d620001c8565b604051601f8301601f19908116603f01168101908282118183101715620002385762000238620001c8565b816040528381526020925086838588010111156200025557600080fd5b600091505b838210156200027957858201830151818301840152908201906200025a565b600093810190920192909252949350505050565b600080600060608486031215620002a357600080fd5b83516001600160401b0380821115620002bb57600080fd5b620002c987838801620001de565b94506020860151915080821115620002e057600080fd5b50620002ef86828701620001de565b925050604084015160ff811681146200030757600080fd5b809150509250925092565b600181811c908216806200032757607f821691505b6020821081036200034857634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200039c57600081815260208120601f850160051c81016020861015620003775750805b601f850160051c820191505b81811015620003985782815560010162000383565b5050505b505050565b81516001600160401b03811115620003bd57620003bd620001c8565b620003d581620003ce845462000312565b846200034e565b602080601f8311600181146200040d5760008415620003f45750858301515b600019600386901b1c1916600185901b17855562000398565b600085815260208120601f198616915b828110156200043e578886015182559484019460019091019084016200041d565b50858210156200045d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c05160e0516116aa620004bc6000396000610237015260008181610307015281816105c001526106a70152600061053a015260008181610379015261050401526116aa6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063a457c2d71161008c578063d505accf11610066578063d505accf1461039b578063dd62ed3e146103ae578063ffa1ad74146103f457600080fd5b8063a457c2d71461034e578063a9059cbb14610361578063cd0d00961461037457600080fd5b806395d89b41116100bd57806395d89b41146102e75780639dc29fac146102ef578063a3c573eb1461030257600080fd5b806370a08231146102915780637ecebe00146102c757600080fd5b806330adf81f1161012f5780633644e515116101145780633644e51514610261578063395093511461026957806340c10f191461027c57600080fd5b806330adf81f14610209578063313ce5671461023057600080fd5b806318160ddd1161016057806318160ddd146101bd57806320606b70146101cf57806323b872dd146101f657600080fd5b806306fdde031461017c578063095ea7b31461019a575b600080fd5b610184610430565b60405161019191906113e4565b60405180910390f35b6101ad6101a8366004611479565b6104c2565b6040519015158152602001610191565b6002545b604051908152602001610191565b6101c17f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81565b6101ad6102043660046114a3565b6104dc565b6101c17f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610191565b6101c1610500565b6101ad610277366004611479565b61055c565b61028f61028a366004611479565b6105a8565b005b6101c161029f3660046114df565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101c16102d53660046114df565b60056020526000908152604090205481565b610184610680565b61028f6102fd366004611479565b61068f565b6103297f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610191565b6101ad61035c366004611479565b61075e565b6101ad61036f366004611479565b61082f565b6101c17f000000000000000000000000000000000000000000000000000000000000000081565b61028f6103a9366004611501565b61083d565b6101c16103bc366004611574565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101846040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525081565b60606003805461043f906115a7565b80601f016020809104026020016040519081016040528092919081815260200182805461046b906115a7565b80156104b85780601f1061048d576101008083540402835291602001916104b8565b820191906000526020600020905b81548152906001019060200180831161049b57829003601f168201915b5050505050905090565b6000336104d0818585610b73565b60019150505b92915050565b6000336104ea858285610d27565b6104f5858585610dfe565b506001949350505050565b60007f00000000000000000000000000000000000000000000000000000000000000004614610537576105324661106d565b905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906104d090829086906105a3908790611629565b610b73565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610672576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f546f6b656e577261707065643a3a6f6e6c794272696467653a204e6f7420506f60448201527f6c79676f6e5a6b45564d4272696467650000000000000000000000000000000060648201526084015b60405180910390fd5b61067c8282611135565b5050565b60606004805461043f906115a7565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610754576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f546f6b656e577261707065643a3a6f6e6c794272696467653a204e6f7420506f60448201527f6c79676f6e5a6b45564d427269646765000000000000000000000000000000006064820152608401610669565b61067c8282611228565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610822576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610669565b6104f58286868403610b73565b6000336104d0818585610dfe565b834211156108cc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f546f6b656e577261707065643a3a7065726d69743a204578706972656420706560448201527f726d6974000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8716600090815260056020526040812080547f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9918a918a918a9190866109268361163c565b9091555060408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610991610500565b6040517f19010000000000000000000000000000000000000000000000000000000000006020820152602281019190915260428101839052606201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600080855291840180845281905260ff89169284019290925260608301879052608083018690529092509060019060a0016020604051602081039080840390855afa158015610a55573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811615801590610ad057508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610b5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f546f6b656e577261707065643a3a7065726d69743a20496e76616c696420736960448201527f676e6174757265000000000000000000000000000000000000000000000000006064820152608401610669565b610b678a8a8a610b73565b50505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8216610cb8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610df85781811015610deb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610669565b610df88484848403610b73565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610ea1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8216610f44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610ffa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610df8565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f611098610430565b8051602091820120604080518082018252600181527f310000000000000000000000000000000000000000000000000000000000000090840152805192830193909352918101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc66060820152608081018390523060a082015260c001604051602081830303815290604052805190602001209050919050565b73ffffffffffffffffffffffffffffffffffffffff82166111b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610669565b80600260008282546111c49190611629565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff82166112cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090205481811015611381576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610d1a565b600060208083528351808285015260005b81811015611411578581018301518582016040015282016113f5565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461147457600080fd5b919050565b6000806040838503121561148c57600080fd5b61149583611450565b946020939093013593505050565b6000806000606084860312156114b857600080fd5b6114c184611450565b92506114cf60208501611450565b9150604084013590509250925092565b6000602082840312156114f157600080fd5b6114fa82611450565b9392505050565b600080600080600080600060e0888a03121561151c57600080fd5b61152588611450565b965061153360208901611450565b95506040880135945060608801359350608088013560ff8116811461155757600080fd5b9699959850939692959460a0840135945060c09093013592915050565b6000806040838503121561158757600080fd5b61159083611450565b915061159e60208401611450565b90509250929050565b600181811c908216806115bb57607f821691505b6020821081036115f4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156104d6576104d66115fa565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361166d5761166d6115fa565b506001019056fea26469706673582212206fe049be388a1d04319b9913fe6d100c4c8e272fe0f5748eb056bd0622beabcc64736f6c63430008140033a2646970667358221220c177217974a307e507e4eacb4176abda9e213fe4662f7a3524711e087806547864736f6c63430008140033",
}

// PolygonzkevmbridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonzkevmbridgeMetaData.ABI instead.
var PolygonzkevmbridgeABI = PolygonzkevmbridgeMetaData.ABI

// PolygonzkevmbridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonzkevmbridgeMetaData.Bin instead.
var PolygonzkevmbridgeBin = PolygonzkevmbridgeMetaData.Bin

// DeployPolygonzkevmbridge deploys a new Ethereum contract, binding an instance of Polygonzkevmbridge to it.
func DeployPolygonzkevmbridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Polygonzkevmbridge, error) {
	parsed, err := PolygonzkevmbridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonzkevmbridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevmbridge{PolygonzkevmbridgeCaller: PolygonzkevmbridgeCaller{contract: contract}, PolygonzkevmbridgeTransactor: PolygonzkevmbridgeTransactor{contract: contract}, PolygonzkevmbridgeFilterer: PolygonzkevmbridgeFilterer{contract: contract}}, nil
}

// Polygonzkevmbridge is an auto generated Go binding around an Ethereum contract.
type Polygonzkevmbridge struct {
	PolygonzkevmbridgeCaller     // Read-only binding to the contract
	PolygonzkevmbridgeTransactor // Write-only binding to the contract
	PolygonzkevmbridgeFilterer   // Log filterer for contract events
}

// PolygonzkevmbridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonzkevmbridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmbridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonzkevmbridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmbridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonzkevmbridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmbridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonzkevmbridgeSession struct {
	Contract     *Polygonzkevmbridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PolygonzkevmbridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonzkevmbridgeCallerSession struct {
	Contract *PolygonzkevmbridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PolygonzkevmbridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonzkevmbridgeTransactorSession struct {
	Contract     *PolygonzkevmbridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PolygonzkevmbridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonzkevmbridgeRaw struct {
	Contract *Polygonzkevmbridge // Generic contract binding to access the raw methods on
}

// PolygonzkevmbridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonzkevmbridgeCallerRaw struct {
	Contract *PolygonzkevmbridgeCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonzkevmbridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonzkevmbridgeTransactorRaw struct {
	Contract *PolygonzkevmbridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevmbridge creates a new instance of Polygonzkevmbridge, bound to a specific deployed contract.
func NewPolygonzkevmbridge(address common.Address, backend bind.ContractBackend) (*Polygonzkevmbridge, error) {
	contract, err := bindPolygonzkevmbridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmbridge{PolygonzkevmbridgeCaller: PolygonzkevmbridgeCaller{contract: contract}, PolygonzkevmbridgeTransactor: PolygonzkevmbridgeTransactor{contract: contract}, PolygonzkevmbridgeFilterer: PolygonzkevmbridgeFilterer{contract: contract}}, nil
}

// NewPolygonzkevmbridgeCaller creates a new read-only instance of Polygonzkevmbridge, bound to a specific deployed contract.
func NewPolygonzkevmbridgeCaller(address common.Address, caller bind.ContractCaller) (*PolygonzkevmbridgeCaller, error) {
	contract, err := bindPolygonzkevmbridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeCaller{contract: contract}, nil
}

// NewPolygonzkevmbridgeTransactor creates a new write-only instance of Polygonzkevmbridge, bound to a specific deployed contract.
func NewPolygonzkevmbridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonzkevmbridgeTransactor, error) {
	contract, err := bindPolygonzkevmbridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeTransactor{contract: contract}, nil
}

// NewPolygonzkevmbridgeFilterer creates a new log filterer instance of Polygonzkevmbridge, bound to a specific deployed contract.
func NewPolygonzkevmbridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonzkevmbridgeFilterer, error) {
	contract, err := bindPolygonzkevmbridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeFilterer{contract: contract}, nil
}

// bindPolygonzkevmbridge binds a generic wrapper to an already deployed contract.
func bindPolygonzkevmbridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonzkevmbridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmbridge *PolygonzkevmbridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmbridge.Contract.PolygonzkevmbridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmbridge *PolygonzkevmbridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.PolygonzkevmbridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmbridge *PolygonzkevmbridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.PolygonzkevmbridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmbridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonzkevmbridge.Contract.DEFAULTADMINROLE(&_Polygonzkevmbridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonzkevmbridge.Contract.DEFAULTADMINROLE(&_Polygonzkevmbridge.CallOpts)
}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) FeePoolValue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "FeePoolValue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) FeePoolValue(arg0 common.Address) (*big.Int, error) {
	return _Polygonzkevmbridge.Contract.FeePoolValue(&_Polygonzkevmbridge.CallOpts, arg0)
}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) FeePoolValue(arg0 common.Address) (*big.Int, error) {
	return _Polygonzkevmbridge.Contract.FeePoolValue(&_Polygonzkevmbridge.CallOpts, arg0)
}

// FundingPoolBalance is a free data retrieval call binding the contract method 0xdbb0481f.
//
// Solidity: function FundingPoolBalance(address ) view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) FundingPoolBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "FundingPoolBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingPoolBalance is a free data retrieval call binding the contract method 0xdbb0481f.
//
// Solidity: function FundingPoolBalance(address ) view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) FundingPoolBalance(arg0 common.Address) (*big.Int, error) {
	return _Polygonzkevmbridge.Contract.FundingPoolBalance(&_Polygonzkevmbridge.CallOpts, arg0)
}

// FundingPoolBalance is a free data retrieval call binding the contract method 0xdbb0481f.
//
// Solidity: function FundingPoolBalance(address ) view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) FundingPoolBalance(arg0 common.Address) (*big.Int, error) {
	return _Polygonzkevmbridge.Contract.FundingPoolBalance(&_Polygonzkevmbridge.CallOpts, arg0)
}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) IsSupportChainId(opts *bind.CallOpts, chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "IsSupportChainId", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) IsSupportChainId(chainId *big.Int) (bool, error) {
	return _Polygonzkevmbridge.Contract.IsSupportChainId(&_Polygonzkevmbridge.CallOpts, chainId)
}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) IsSupportChainId(chainId *big.Int) (bool, error) {
	return _Polygonzkevmbridge.Contract.IsSupportChainId(&_Polygonzkevmbridge.CallOpts, chainId)
}

// IsSupportToken is a free data retrieval call binding the contract method 0xdac29568.
//
// Solidity: function IsSupportToken(address ) view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) IsSupportToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "IsSupportToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportToken is a free data retrieval call binding the contract method 0xdac29568.
//
// Solidity: function IsSupportToken(address ) view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) IsSupportToken(arg0 common.Address) (bool, error) {
	return _Polygonzkevmbridge.Contract.IsSupportToken(&_Polygonzkevmbridge.CallOpts, arg0)
}

// IsSupportToken is a free data retrieval call binding the contract method 0xdac29568.
//
// Solidity: function IsSupportToken(address ) view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) IsSupportToken(arg0 common.Address) (bool, error) {
	return _Polygonzkevmbridge.Contract.IsSupportToken(&_Polygonzkevmbridge.CallOpts, arg0)
}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) L2WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "L2WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) L2WETH() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.L2WETH(&_Polygonzkevmbridge.CallOpts)
}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) L2WETH() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.L2WETH(&_Polygonzkevmbridge.CallOpts)
}

// MAXGASLimit is a free data retrieval call binding the contract method 0xa525ae2b.
//
// Solidity: function MAX_GAS_Limit() view returns(uint32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) MAXGASLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "MAX_GAS_Limit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXGASLimit is a free data retrieval call binding the contract method 0xa525ae2b.
//
// Solidity: function MAX_GAS_Limit() view returns(uint32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) MAXGASLimit() (uint32, error) {
	return _Polygonzkevmbridge.Contract.MAXGASLimit(&_Polygonzkevmbridge.CallOpts)
}

// MAXGASLimit is a free data retrieval call binding the contract method 0xa525ae2b.
//
// Solidity: function MAX_GAS_Limit() view returns(uint32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) MAXGASLimit() (uint32, error) {
	return _Polygonzkevmbridge.Contract.MAXGASLimit(&_Polygonzkevmbridge.CallOpts)
}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) MinTransferAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "MinTransferAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) MinTransferAmount() (*big.Int, error) {
	return _Polygonzkevmbridge.Contract.MinTransferAmount(&_Polygonzkevmbridge.CallOpts)
}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) MinTransferAmount() (*big.Int, error) {
	return _Polygonzkevmbridge.Contract.MinTransferAmount(&_Polygonzkevmbridge.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) PerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "PerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) PerFee() (*big.Int, error) {
	return _Polygonzkevmbridge.Contract.PerFee(&_Polygonzkevmbridge.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) PerFee() (*big.Int, error) {
	return _Polygonzkevmbridge.Contract.PerFee(&_Polygonzkevmbridge.CallOpts)
}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) ReLayer(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "ReLayer")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) ReLayer() ([32]byte, error) {
	return _Polygonzkevmbridge.Contract.ReLayer(&_Polygonzkevmbridge.CallOpts)
}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) ReLayer() ([32]byte, error) {
	return _Polygonzkevmbridge.Contract.ReLayer(&_Polygonzkevmbridge.CallOpts)
}

// SupportTokens is a free data retrieval call binding the contract method 0x626417b5.
//
// Solidity: function SupportTokens(uint256 ) view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) SupportTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "SupportTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportTokens is a free data retrieval call binding the contract method 0x626417b5.
//
// Solidity: function SupportTokens(uint256 ) view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) SupportTokens(arg0 *big.Int) (common.Address, error) {
	return _Polygonzkevmbridge.Contract.SupportTokens(&_Polygonzkevmbridge.CallOpts, arg0)
}

// SupportTokens is a free data retrieval call binding the contract method 0x626417b5.
//
// Solidity: function SupportTokens(uint256 ) view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) SupportTokens(arg0 *big.Int) (common.Address, error) {
	return _Polygonzkevmbridge.Contract.SupportTokens(&_Polygonzkevmbridge.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonzkevmbridge.Contract.GetRoleAdmin(&_Polygonzkevmbridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonzkevmbridge.Contract.GetRoleAdmin(&_Polygonzkevmbridge.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonzkevmbridge.Contract.HasRole(&_Polygonzkevmbridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonzkevmbridge.Contract.HasRole(&_Polygonzkevmbridge.CallOpts, role, account)
}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) MessageManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "messageManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) MessageManager() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.MessageManager(&_Polygonzkevmbridge.CallOpts)
}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) MessageManager() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.MessageManager(&_Polygonzkevmbridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) Paused() (bool, error) {
	return _Polygonzkevmbridge.Contract.Paused(&_Polygonzkevmbridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) Paused() (bool, error) {
	return _Polygonzkevmbridge.Contract.Paused(&_Polygonzkevmbridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Polygonzkevmbridge.Contract.SupportsInterface(&_Polygonzkevmbridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Polygonzkevmbridge.Contract.SupportsInterface(&_Polygonzkevmbridge.CallOpts, interfaceId)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x92d84738.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) BridgeFinalizeERC20(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "BridgeFinalizeERC20", sourceChainId, destChainId, to, ERC20Address, amount, _fee, _nonce)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x92d84738.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) BridgeFinalizeERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeFinalizeERC20(&_Polygonzkevmbridge.TransactOpts, sourceChainId, destChainId, to, ERC20Address, amount, _fee, _nonce)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x92d84738.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) BridgeFinalizeERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeFinalizeERC20(&_Polygonzkevmbridge.TransactOpts, sourceChainId, destChainId, to, ERC20Address, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x38731c47.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) BridgeFinalizeETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "BridgeFinalizeETH", sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x38731c47.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) BridgeFinalizeETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeFinalizeETH(&_Polygonzkevmbridge.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x38731c47.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) BridgeFinalizeETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeFinalizeETH(&_Polygonzkevmbridge.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeWETH is a paid mutator transaction binding the contract method 0x5765a4eb.
//
// Solidity: function BridgeFinalizeWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) BridgeFinalizeWETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "BridgeFinalizeWETH", sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeWETH is a paid mutator transaction binding the contract method 0x5765a4eb.
//
// Solidity: function BridgeFinalizeWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) BridgeFinalizeWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeFinalizeWETH(&_Polygonzkevmbridge.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeWETH is a paid mutator transaction binding the contract method 0x5765a4eb.
//
// Solidity: function BridgeFinalizeWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) BridgeFinalizeWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeFinalizeWETH(&_Polygonzkevmbridge.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x3b0230f0.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) BridgeInitiateERC20(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "BridgeInitiateERC20", sourceChainId, destChainId, to, ERC20Address, value)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x3b0230f0.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) BridgeInitiateERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeInitiateERC20(&_Polygonzkevmbridge.TransactOpts, sourceChainId, destChainId, to, ERC20Address, value)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x3b0230f0.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) BridgeInitiateERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeInitiateERC20(&_Polygonzkevmbridge.TransactOpts, sourceChainId, destChainId, to, ERC20Address, value)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x9b442380.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address to) payable returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) BridgeInitiateETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "BridgeInitiateETH", sourceChainId, destChainId, to)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x9b442380.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address to) payable returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) BridgeInitiateETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeInitiateETH(&_Polygonzkevmbridge.TransactOpts, sourceChainId, destChainId, to)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x9b442380.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address to) payable returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) BridgeInitiateETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeInitiateETH(&_Polygonzkevmbridge.TransactOpts, sourceChainId, destChainId, to)
}

// BridgeInitiateWETH is a paid mutator transaction binding the contract method 0xcd01c665.
//
// Solidity: function BridgeInitiateWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 value) returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) BridgeInitiateWETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "BridgeInitiateWETH", sourceChainId, destChainId, to, value)
}

// BridgeInitiateWETH is a paid mutator transaction binding the contract method 0xcd01c665.
//
// Solidity: function BridgeInitiateWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 value) returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) BridgeInitiateWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeInitiateWETH(&_Polygonzkevmbridge.TransactOpts, sourceChainId, destChainId, to, value)
}

// BridgeInitiateWETH is a paid mutator transaction binding the contract method 0xcd01c665.
//
// Solidity: function BridgeInitiateWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 value) returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) BridgeInitiateWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeInitiateWETH(&_Polygonzkevmbridge.TransactOpts, sourceChainId, destChainId, to, value)
}

// QuickSendAssertToUser is a paid mutator transaction binding the contract method 0x325e5d43.
//
// Solidity: function QuickSendAssertToUser(address _token, address to, uint256 _amount) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) QuickSendAssertToUser(opts *bind.TransactOpts, _token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "QuickSendAssertToUser", _token, to, _amount)
}

// QuickSendAssertToUser is a paid mutator transaction binding the contract method 0x325e5d43.
//
// Solidity: function QuickSendAssertToUser(address _token, address to, uint256 _amount) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) QuickSendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.QuickSendAssertToUser(&_Polygonzkevmbridge.TransactOpts, _token, to, _amount)
}

// QuickSendAssertToUser is a paid mutator transaction binding the contract method 0x325e5d43.
//
// Solidity: function QuickSendAssertToUser(address _token, address to, uint256 _amount) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) QuickSendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.QuickSendAssertToUser(&_Polygonzkevmbridge.TransactOpts, _token, to, _amount)
}

// UpdateFundingPoolBalance is a paid mutator transaction binding the contract method 0x0cd727df.
//
// Solidity: function UpdateFundingPoolBalance(address token, uint256 amount) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) UpdateFundingPoolBalance(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "UpdateFundingPoolBalance", token, amount)
}

// UpdateFundingPoolBalance is a paid mutator transaction binding the contract method 0x0cd727df.
//
// Solidity: function UpdateFundingPoolBalance(address token, uint256 amount) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) UpdateFundingPoolBalance(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.UpdateFundingPoolBalance(&_Polygonzkevmbridge.TransactOpts, token, amount)
}

// UpdateFundingPoolBalance is a paid mutator transaction binding the contract method 0x0cd727df.
//
// Solidity: function UpdateFundingPoolBalance(address token, uint256 amount) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) UpdateFundingPoolBalance(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.UpdateFundingPoolBalance(&_Polygonzkevmbridge.TransactOpts, token, amount)
}

// WithdrawERC20ToL1 is a paid mutator transaction binding the contract method 0x0227047a.
//
// Solidity: function WithdrawERC20ToL1(address _token, address _to, uint256 _amount) payable returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) WithdrawERC20ToL1(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "WithdrawERC20ToL1", _token, _to, _amount)
}

// WithdrawERC20ToL1 is a paid mutator transaction binding the contract method 0x0227047a.
//
// Solidity: function WithdrawERC20ToL1(address _token, address _to, uint256 _amount) payable returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) WithdrawERC20ToL1(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.WithdrawERC20ToL1(&_Polygonzkevmbridge.TransactOpts, _token, _to, _amount)
}

// WithdrawERC20ToL1 is a paid mutator transaction binding the contract method 0x0227047a.
//
// Solidity: function WithdrawERC20ToL1(address _token, address _to, uint256 _amount) payable returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) WithdrawERC20ToL1(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.WithdrawERC20ToL1(&_Polygonzkevmbridge.TransactOpts, _token, _to, _amount)
}

// WithdrawETHtoL1 is a paid mutator transaction binding the contract method 0x39891dd7.
//
// Solidity: function WithdrawETHtoL1(address _to, uint256 _amount) payable returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) WithdrawETHtoL1(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "WithdrawETHtoL1", _to, _amount)
}

// WithdrawETHtoL1 is a paid mutator transaction binding the contract method 0x39891dd7.
//
// Solidity: function WithdrawETHtoL1(address _to, uint256 _amount) payable returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) WithdrawETHtoL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.WithdrawETHtoL1(&_Polygonzkevmbridge.TransactOpts, _to, _amount)
}

// WithdrawETHtoL1 is a paid mutator transaction binding the contract method 0x39891dd7.
//
// Solidity: function WithdrawETHtoL1(address _to, uint256 _amount) payable returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) WithdrawETHtoL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.WithdrawETHtoL1(&_Polygonzkevmbridge.TransactOpts, _to, _amount)
}

// WithdrawWETHToL1 is a paid mutator transaction binding the contract method 0xd28597f4.
//
// Solidity: function WithdrawWETHToL1(address _to, uint256 _amount) payable returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) WithdrawWETHToL1(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "WithdrawWETHToL1", _to, _amount)
}

// WithdrawWETHToL1 is a paid mutator transaction binding the contract method 0xd28597f4.
//
// Solidity: function WithdrawWETHToL1(address _to, uint256 _amount) payable returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) WithdrawWETHToL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.WithdrawWETHToL1(&_Polygonzkevmbridge.TransactOpts, _to, _amount)
}

// WithdrawWETHToL1 is a paid mutator transaction binding the contract method 0xd28597f4.
//
// Solidity: function WithdrawWETHToL1(address _to, uint256 _amount) payable returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) WithdrawWETHToL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.WithdrawWETHToL1(&_Polygonzkevmbridge.TransactOpts, _to, _amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.GrantRole(&_Polygonzkevmbridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.GrantRole(&_Polygonzkevmbridge.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _MultisigWallet, address _messageManager) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) Initialize(opts *bind.TransactOpts, _MultisigWallet common.Address, _messageManager common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "initialize", _MultisigWallet, _messageManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _MultisigWallet, address _messageManager) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) Initialize(_MultisigWallet common.Address, _messageManager common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.Initialize(&_Polygonzkevmbridge.TransactOpts, _MultisigWallet, _messageManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _MultisigWallet, address _messageManager) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) Initialize(_MultisigWallet common.Address, _messageManager common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.Initialize(&_Polygonzkevmbridge.TransactOpts, _MultisigWallet, _messageManager)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) Pause() (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.Pause(&_Polygonzkevmbridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.Pause(&_Polygonzkevmbridge.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.RenounceRole(&_Polygonzkevmbridge.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.RenounceRole(&_Polygonzkevmbridge.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.RevokeRole(&_Polygonzkevmbridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.RevokeRole(&_Polygonzkevmbridge.TransactOpts, role, account)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) SetMinTransferAmount(opts *bind.TransactOpts, _MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "setMinTransferAmount", _MinTransferAmount)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) SetMinTransferAmount(_MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.SetMinTransferAmount(&_Polygonzkevmbridge.TransactOpts, _MinTransferAmount)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) SetMinTransferAmount(_MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.SetMinTransferAmount(&_Polygonzkevmbridge.TransactOpts, _MinTransferAmount)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) SetPerFee(opts *bind.TransactOpts, _PerFee *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "setPerFee", _PerFee)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) SetPerFee(_PerFee *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.SetPerFee(&_Polygonzkevmbridge.TransactOpts, _PerFee)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) SetPerFee(_PerFee *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.SetPerFee(&_Polygonzkevmbridge.TransactOpts, _PerFee)
}

// SetSupportERC20Token is a paid mutator transaction binding the contract method 0xab0f9e19.
//
// Solidity: function setSupportERC20Token(address ERC20Address, bool isValid) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) SetSupportERC20Token(opts *bind.TransactOpts, ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "setSupportERC20Token", ERC20Address, isValid)
}

// SetSupportERC20Token is a paid mutator transaction binding the contract method 0xab0f9e19.
//
// Solidity: function setSupportERC20Token(address ERC20Address, bool isValid) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) SetSupportERC20Token(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.SetSupportERC20Token(&_Polygonzkevmbridge.TransactOpts, ERC20Address, isValid)
}

// SetSupportERC20Token is a paid mutator transaction binding the contract method 0xab0f9e19.
//
// Solidity: function setSupportERC20Token(address ERC20Address, bool isValid) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) SetSupportERC20Token(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.SetSupportERC20Token(&_Polygonzkevmbridge.TransactOpts, ERC20Address, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) SetValidChainId(opts *bind.TransactOpts, chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "setValidChainId", chainId, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) SetValidChainId(chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.SetValidChainId(&_Polygonzkevmbridge.TransactOpts, chainId, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) SetValidChainId(chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.SetValidChainId(&_Polygonzkevmbridge.TransactOpts, chainId, isValid)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) Unpause() (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.Unpause(&_Polygonzkevmbridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.Unpause(&_Polygonzkevmbridge.TransactOpts)
}

// PolygonzkevmbridgeFinalizeERC20Iterator is returned from FilterFinalizeERC20 and is used to iterate over the raw logs and unpacked data for FinalizeERC20 events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeFinalizeERC20Iterator struct {
	Event *PolygonzkevmbridgeFinalizeERC20 // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeFinalizeERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeFinalizeERC20)
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
		it.Event = new(PolygonzkevmbridgeFinalizeERC20)
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
func (it *PolygonzkevmbridgeFinalizeERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeFinalizeERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeFinalizeERC20 represents a FinalizeERC20 event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeFinalizeERC20 struct {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterFinalizeERC20(opts *bind.FilterOpts, ERC20Address []common.Address, from []common.Address, to []common.Address) (*PolygonzkevmbridgeFinalizeERC20Iterator, error) {

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

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "FinalizeERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeFinalizeERC20Iterator{contract: _Polygonzkevmbridge.contract, event: "FinalizeERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeERC20 is a free log subscription operation binding the contract event 0x0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe.
//
// Solidity: event FinalizeERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchFinalizeERC20(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeFinalizeERC20, ERC20Address []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "FinalizeERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeFinalizeERC20)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "FinalizeERC20", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseFinalizeERC20(log types.Log) (*PolygonzkevmbridgeFinalizeERC20, error) {
	event := new(PolygonzkevmbridgeFinalizeERC20)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "FinalizeERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeFinalizeETHIterator is returned from FilterFinalizeETH and is used to iterate over the raw logs and unpacked data for FinalizeETH events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeFinalizeETHIterator struct {
	Event *PolygonzkevmbridgeFinalizeETH // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeFinalizeETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeFinalizeETH)
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
		it.Event = new(PolygonzkevmbridgeFinalizeETH)
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
func (it *PolygonzkevmbridgeFinalizeETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeFinalizeETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeFinalizeETH represents a FinalizeETH event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeFinalizeETH struct {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterFinalizeETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PolygonzkevmbridgeFinalizeETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "FinalizeETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeFinalizeETHIterator{contract: _Polygonzkevmbridge.contract, event: "FinalizeETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeETH is a free log subscription operation binding the contract event 0x61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f32.
//
// Solidity: event FinalizeETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchFinalizeETH(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeFinalizeETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "FinalizeETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeFinalizeETH)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "FinalizeETH", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseFinalizeETH(log types.Log) (*PolygonzkevmbridgeFinalizeETH, error) {
	event := new(PolygonzkevmbridgeFinalizeETH)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "FinalizeETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeFinalizeWETHIterator is returned from FilterFinalizeWETH and is used to iterate over the raw logs and unpacked data for FinalizeWETH events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeFinalizeWETHIterator struct {
	Event *PolygonzkevmbridgeFinalizeWETH // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeFinalizeWETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeFinalizeWETH)
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
		it.Event = new(PolygonzkevmbridgeFinalizeWETH)
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
func (it *PolygonzkevmbridgeFinalizeWETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeFinalizeWETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeFinalizeWETH represents a FinalizeWETH event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeFinalizeWETH struct {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterFinalizeWETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PolygonzkevmbridgeFinalizeWETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "FinalizeWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeFinalizeWETHIterator{contract: _Polygonzkevmbridge.contract, event: "FinalizeWETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeWETH is a free log subscription operation binding the contract event 0x7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc14.
//
// Solidity: event FinalizeWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchFinalizeWETH(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeFinalizeWETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "FinalizeWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeFinalizeWETH)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "FinalizeWETH", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseFinalizeWETH(log types.Log) (*PolygonzkevmbridgeFinalizeWETH, error) {
	event := new(PolygonzkevmbridgeFinalizeWETH)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "FinalizeWETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeInitializedIterator struct {
	Event *PolygonzkevmbridgeInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeInitialized)
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
		it.Event = new(PolygonzkevmbridgeInitialized)
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
func (it *PolygonzkevmbridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeInitialized represents a Initialized event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonzkevmbridgeInitializedIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeInitializedIterator{contract: _Polygonzkevmbridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeInitialized)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseInitialized(log types.Log) (*PolygonzkevmbridgeInitialized, error) {
	event := new(PolygonzkevmbridgeInitialized)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeInitiateERC20Iterator is returned from FilterInitiateERC20 and is used to iterate over the raw logs and unpacked data for InitiateERC20 events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeInitiateERC20Iterator struct {
	Event *PolygonzkevmbridgeInitiateERC20 // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeInitiateERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeInitiateERC20)
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
		it.Event = new(PolygonzkevmbridgeInitiateERC20)
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
func (it *PolygonzkevmbridgeInitiateERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeInitiateERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeInitiateERC20 represents a InitiateERC20 event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeInitiateERC20 struct {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterInitiateERC20(opts *bind.FilterOpts, ERC20Address []common.Address, from []common.Address, to []common.Address) (*PolygonzkevmbridgeInitiateERC20Iterator, error) {

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

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "InitiateERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeInitiateERC20Iterator{contract: _Polygonzkevmbridge.contract, event: "InitiateERC20", logs: logs, sub: sub}, nil
}

// WatchInitiateERC20 is a free log subscription operation binding the contract event 0xece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a4.
//
// Solidity: event InitiateERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchInitiateERC20(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeInitiateERC20, ERC20Address []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "InitiateERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeInitiateERC20)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "InitiateERC20", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseInitiateERC20(log types.Log) (*PolygonzkevmbridgeInitiateERC20, error) {
	event := new(PolygonzkevmbridgeInitiateERC20)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "InitiateERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeInitiateETHIterator is returned from FilterInitiateETH and is used to iterate over the raw logs and unpacked data for InitiateETH events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeInitiateETHIterator struct {
	Event *PolygonzkevmbridgeInitiateETH // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeInitiateETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeInitiateETH)
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
		it.Event = new(PolygonzkevmbridgeInitiateETH)
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
func (it *PolygonzkevmbridgeInitiateETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeInitiateETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeInitiateETH represents a InitiateETH event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeInitiateETH struct {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterInitiateETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PolygonzkevmbridgeInitiateETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "InitiateETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeInitiateETHIterator{contract: _Polygonzkevmbridge.contract, event: "InitiateETH", logs: logs, sub: sub}, nil
}

// WatchInitiateETH is a free log subscription operation binding the contract event 0x0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc5.
//
// Solidity: event InitiateETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchInitiateETH(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeInitiateETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "InitiateETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeInitiateETH)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "InitiateETH", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseInitiateETH(log types.Log) (*PolygonzkevmbridgeInitiateETH, error) {
	event := new(PolygonzkevmbridgeInitiateETH)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "InitiateETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeInitiateWETHIterator is returned from FilterInitiateWETH and is used to iterate over the raw logs and unpacked data for InitiateWETH events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeInitiateWETHIterator struct {
	Event *PolygonzkevmbridgeInitiateWETH // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeInitiateWETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeInitiateWETH)
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
		it.Event = new(PolygonzkevmbridgeInitiateWETH)
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
func (it *PolygonzkevmbridgeInitiateWETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeInitiateWETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeInitiateWETH represents a InitiateWETH event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeInitiateWETH struct {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterInitiateWETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PolygonzkevmbridgeInitiateWETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "InitiateWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeInitiateWETHIterator{contract: _Polygonzkevmbridge.contract, event: "InitiateWETH", logs: logs, sub: sub}, nil
}

// WatchInitiateWETH is a free log subscription operation binding the contract event 0xcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a89.
//
// Solidity: event InitiateWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchInitiateWETH(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeInitiateWETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "InitiateWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeInitiateWETH)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "InitiateWETH", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseInitiateWETH(log types.Log) (*PolygonzkevmbridgeInitiateWETH, error) {
	event := new(PolygonzkevmbridgeInitiateWETH)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "InitiateWETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgePausedIterator struct {
	Event *PolygonzkevmbridgePaused // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgePaused)
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
		it.Event = new(PolygonzkevmbridgePaused)
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
func (it *PolygonzkevmbridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgePaused represents a Paused event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*PolygonzkevmbridgePausedIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgePausedIterator{contract: _Polygonzkevmbridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgePaused) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgePaused)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParsePaused(log types.Log) (*PolygonzkevmbridgePaused, error) {
	event := new(PolygonzkevmbridgePaused)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeRoleAdminChangedIterator struct {
	Event *PolygonzkevmbridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeRoleAdminChanged)
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
		it.Event = new(PolygonzkevmbridgeRoleAdminChanged)
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
func (it *PolygonzkevmbridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeRoleAdminChanged represents a RoleAdminChanged event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PolygonzkevmbridgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeRoleAdminChangedIterator{contract: _Polygonzkevmbridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeRoleAdminChanged)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseRoleAdminChanged(log types.Log) (*PolygonzkevmbridgeRoleAdminChanged, error) {
	event := new(PolygonzkevmbridgeRoleAdminChanged)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeRoleGrantedIterator struct {
	Event *PolygonzkevmbridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeRoleGranted)
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
		it.Event = new(PolygonzkevmbridgeRoleGranted)
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
func (it *PolygonzkevmbridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeRoleGranted represents a RoleGranted event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolygonzkevmbridgeRoleGrantedIterator, error) {

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

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeRoleGrantedIterator{contract: _Polygonzkevmbridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeRoleGranted)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseRoleGranted(log types.Log) (*PolygonzkevmbridgeRoleGranted, error) {
	event := new(PolygonzkevmbridgeRoleGranted)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeRoleRevokedIterator struct {
	Event *PolygonzkevmbridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeRoleRevoked)
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
		it.Event = new(PolygonzkevmbridgeRoleRevoked)
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
func (it *PolygonzkevmbridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeRoleRevoked represents a RoleRevoked event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolygonzkevmbridgeRoleRevokedIterator, error) {

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

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeRoleRevokedIterator{contract: _Polygonzkevmbridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeRoleRevoked)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseRoleRevoked(log types.Log) (*PolygonzkevmbridgeRoleRevoked, error) {
	event := new(PolygonzkevmbridgeRoleRevoked)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeUnpausedIterator struct {
	Event *PolygonzkevmbridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeUnpaused)
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
		it.Event = new(PolygonzkevmbridgeUnpaused)
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
func (it *PolygonzkevmbridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeUnpaused represents a Unpaused event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PolygonzkevmbridgeUnpausedIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeUnpausedIterator{contract: _Polygonzkevmbridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeUnpaused)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseUnpaused(log types.Log) (*PolygonzkevmbridgeUnpaused, error) {
	event := new(PolygonzkevmbridgeUnpaused)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeWithdrawERC20toL1SuccessIterator is returned from FilterWithdrawERC20toL1Success and is used to iterate over the raw logs and unpacked data for WithdrawERC20toL1Success events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeWithdrawERC20toL1SuccessIterator struct {
	Event *PolygonzkevmbridgeWithdrawERC20toL1Success // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeWithdrawERC20toL1SuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeWithdrawERC20toL1Success)
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
		it.Event = new(PolygonzkevmbridgeWithdrawERC20toL1Success)
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
func (it *PolygonzkevmbridgeWithdrawERC20toL1SuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeWithdrawERC20toL1SuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeWithdrawERC20toL1Success represents a WithdrawERC20toL1Success event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeWithdrawERC20toL1Success struct {
	ChainId   *big.Int
	Timestamp *big.Int
	Token     common.Address
	To        common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC20toL1Success is a free log retrieval operation binding the contract event 0x1bc1b1207e0c319b7baad0317c09f30575218af6f5a59f1f0f15de8d217672cc.
//
// Solidity: event WithdrawERC20toL1Success(uint256 chainId, uint256 timestamp, address token, address to, uint256 value)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterWithdrawERC20toL1Success(opts *bind.FilterOpts) (*PolygonzkevmbridgeWithdrawERC20toL1SuccessIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "WithdrawERC20toL1Success")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeWithdrawERC20toL1SuccessIterator{contract: _Polygonzkevmbridge.contract, event: "WithdrawERC20toL1Success", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20toL1Success is a free log subscription operation binding the contract event 0x1bc1b1207e0c319b7baad0317c09f30575218af6f5a59f1f0f15de8d217672cc.
//
// Solidity: event WithdrawERC20toL1Success(uint256 chainId, uint256 timestamp, address token, address to, uint256 value)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchWithdrawERC20toL1Success(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeWithdrawERC20toL1Success) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "WithdrawERC20toL1Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeWithdrawERC20toL1Success)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "WithdrawERC20toL1Success", log); err != nil {
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

// ParseWithdrawERC20toL1Success is a log parse operation binding the contract event 0x1bc1b1207e0c319b7baad0317c09f30575218af6f5a59f1f0f15de8d217672cc.
//
// Solidity: event WithdrawERC20toL1Success(uint256 chainId, uint256 timestamp, address token, address to, uint256 value)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseWithdrawERC20toL1Success(log types.Log) (*PolygonzkevmbridgeWithdrawERC20toL1Success, error) {
	event := new(PolygonzkevmbridgeWithdrawERC20toL1Success)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "WithdrawERC20toL1Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeWithdrawETHtoL1SuccessIterator is returned from FilterWithdrawETHtoL1Success and is used to iterate over the raw logs and unpacked data for WithdrawETHtoL1Success events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeWithdrawETHtoL1SuccessIterator struct {
	Event *PolygonzkevmbridgeWithdrawETHtoL1Success // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeWithdrawETHtoL1SuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeWithdrawETHtoL1Success)
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
		it.Event = new(PolygonzkevmbridgeWithdrawETHtoL1Success)
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
func (it *PolygonzkevmbridgeWithdrawETHtoL1SuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeWithdrawETHtoL1SuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeWithdrawETHtoL1Success represents a WithdrawETHtoL1Success event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeWithdrawETHtoL1Success struct {
	ChainId   *big.Int
	Timestamp *big.Int
	To        common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETHtoL1Success is a free log retrieval operation binding the contract event 0xf1fa6dbbb3e74f1bef249f22545b238726c1cc79c5571a1416575674c7e63f39.
//
// Solidity: event WithdrawETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterWithdrawETHtoL1Success(opts *bind.FilterOpts) (*PolygonzkevmbridgeWithdrawETHtoL1SuccessIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "WithdrawETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeWithdrawETHtoL1SuccessIterator{contract: _Polygonzkevmbridge.contract, event: "WithdrawETHtoL1Success", logs: logs, sub: sub}, nil
}

// WatchWithdrawETHtoL1Success is a free log subscription operation binding the contract event 0xf1fa6dbbb3e74f1bef249f22545b238726c1cc79c5571a1416575674c7e63f39.
//
// Solidity: event WithdrawETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchWithdrawETHtoL1Success(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeWithdrawETHtoL1Success) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "WithdrawETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeWithdrawETHtoL1Success)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "WithdrawETHtoL1Success", log); err != nil {
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

// ParseWithdrawETHtoL1Success is a log parse operation binding the contract event 0xf1fa6dbbb3e74f1bef249f22545b238726c1cc79c5571a1416575674c7e63f39.
//
// Solidity: event WithdrawETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseWithdrawETHtoL1Success(log types.Log) (*PolygonzkevmbridgeWithdrawETHtoL1Success, error) {
	event := new(PolygonzkevmbridgeWithdrawETHtoL1Success)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "WithdrawETHtoL1Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeWithdrawWETHtoL1SuccessIterator is returned from FilterWithdrawWETHtoL1Success and is used to iterate over the raw logs and unpacked data for WithdrawWETHtoL1Success events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeWithdrawWETHtoL1SuccessIterator struct {
	Event *PolygonzkevmbridgeWithdrawWETHtoL1Success // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeWithdrawWETHtoL1SuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeWithdrawWETHtoL1Success)
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
		it.Event = new(PolygonzkevmbridgeWithdrawWETHtoL1Success)
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
func (it *PolygonzkevmbridgeWithdrawWETHtoL1SuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeWithdrawWETHtoL1SuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeWithdrawWETHtoL1Success represents a WithdrawWETHtoL1Success event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeWithdrawWETHtoL1Success struct {
	ChainId   *big.Int
	Timestamp *big.Int
	To        common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawWETHtoL1Success is a free log retrieval operation binding the contract event 0xcb1a101498aadad4159168d719f955e559e768912e6f5f5c57b094eb047d000f.
//
// Solidity: event WithdrawWETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterWithdrawWETHtoL1Success(opts *bind.FilterOpts) (*PolygonzkevmbridgeWithdrawWETHtoL1SuccessIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "WithdrawWETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeWithdrawWETHtoL1SuccessIterator{contract: _Polygonzkevmbridge.contract, event: "WithdrawWETHtoL1Success", logs: logs, sub: sub}, nil
}

// WatchWithdrawWETHtoL1Success is a free log subscription operation binding the contract event 0xcb1a101498aadad4159168d719f955e559e768912e6f5f5c57b094eb047d000f.
//
// Solidity: event WithdrawWETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchWithdrawWETHtoL1Success(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeWithdrawWETHtoL1Success) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "WithdrawWETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeWithdrawWETHtoL1Success)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "WithdrawWETHtoL1Success", log); err != nil {
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

// ParseWithdrawWETHtoL1Success is a log parse operation binding the contract event 0xcb1a101498aadad4159168d719f955e559e768912e6f5f5c57b094eb047d000f.
//
// Solidity: event WithdrawWETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseWithdrawWETHtoL1Success(log types.Log) (*PolygonzkevmbridgeWithdrawWETHtoL1Success, error) {
	event := new(PolygonzkevmbridgeWithdrawWETHtoL1Success)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "WithdrawWETHtoL1Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
