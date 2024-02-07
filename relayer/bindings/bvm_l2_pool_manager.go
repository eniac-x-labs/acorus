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

// L2PoolManagerMetaData contains all meta data concerning the L2PoolManager contract.
var L2PoolManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ChainIdIsNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"ChainIdNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorBlockChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"MinTransferAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LessThanMinTransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughETH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"NotEnoughToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"StableCoinNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20toL1Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawETHtoL1Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawWETHtoL1Success\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"BridgeInitiateETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FeePoolValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"IsSupportChainId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"IsSupportStableCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_GAS_Limit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinTransferAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ReLayer\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SendAssertToUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20ToL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawETHtoL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawWETHToL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MultisigWallet\",\"type\":\"address\"},{\"internalType\":\"contractIMessageManager\",\"name\":\"_messageManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageManager\",\"outputs\":[{\"internalType\":\"contractIMessageManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_MinTransferAmount\",\"type\":\"uint256\"}],\"name\":\"setMinTransferAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_PerFee\",\"type\":\"uint256\"}],\"name\":\"setPerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setSupportStableCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setValidChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506200001b62000021565b620000d5565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d25780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b612d4080620000e35f395ff3fe6080604052600436106101f1575f3560e01c806373c23be911610108578063b33280dd1161009d578063d547741f1161006d578063d547741f14610566578063d73f14e514610585578063dd0c3460146105a4578063ed06da87146105b9578063fa861848146105d8575f80fd5b8063b33280dd146104f6578063b92e639614610515578063cd01c66514610534578063d28597f414610553575f80fd5b80639b442380116100d85780639b4423801461048b578063a217fddf1461049e578063a525ae2b146104b1578063b1887e9b146104e2575f80fd5b806373c23be9146104265780638456cb591461043957806391d148541461044d57806392d847381461046c575f80fd5b806339891dd71161018957806354dc027e1161015957806354dc027e1461036c5780635765a4eb146103975780635c975abb146103b6578063650c2276146103d957806372fe6a7e146103f8575f80fd5b806339891dd7146103075780633b0230f01461031a5780633f4ba83a14610339578063485cc9551461034d575f80fd5b8063248a9ca3116101c4578063248a9ca3146102955780632f2ff15d146102b457806336568abe146102d557806338731c47146102f4575f80fd5b806301ffc9a7146101f55780630227047a1461022957806309dc480c1461023c57806310875a131461025f575b5f80fd5b348015610200575f80fd5b5061021461020f3660046127e3565b6105ec565b60405190151581526020015b60405180910390f35b61021461023736600461281e565b610622565b348015610247575f80fd5b5061025160015481565b604051908152602001610220565b34801561026a575f80fd5b505f5461027d906001600160a01b031681565b6040516001600160a01b039091168152602001610220565b3480156102a0575f80fd5b506102516102af36600461285c565b610945565b3480156102bf575f80fd5b506102d36102ce366004612873565b610965565b005b3480156102e0575f80fd5b506102d36102ef366004612873565b610987565b6102146103023660046128a1565b6109bf565b6102146103153660046128f0565b610b4e565b348015610325575f80fd5b5061021461033436600461291a565b610d7f565b348015610344575f80fd5b506102d361106d565b348015610358575f80fd5b506102d3610367366004612968565b611082565b348015610377575f80fd5b50610251610386366004612994565b60066020525f908152604090205481565b3480156103a2575f80fd5b506102146103b13660046128a1565b6111b4565b3480156103c1575f80fd5b505f80516020612ceb8339815191525460ff16610214565b3480156103e4575f80fd5b506102d36103f336600461285c565b6113b0565b348015610403575f80fd5b5061021461041236600461285c565b5f9081526003602052604090205460ff1690565b61021461043436600461281e565b6113c0565b348015610444575f80fd5b506102d361156f565b348015610458575f80fd5b50610214610467366004612873565b611581565b348015610477575f80fd5b506102146104863660046129af565b6115b7565b610214610499366004612a0f565b611766565b3480156104a9575f80fd5b506102515f81565b3480156104bc575f80fd5b506007546104cd9063ffffffff1681565b60405163ffffffff9091168152602001610220565b3480156104ed575f80fd5b5061027d61197d565b348015610501575f80fd5b506102d3610510366004612a52565b6119d2565b348015610520575f80fd5b506102d361052f36600461285c565b611a07565b34801561053f575f80fd5b5061021461054e366004612a7e565b611a5a565b6102146105613660046128f0565b611de9565b348015610571575f80fd5b506102d3610580366004612873565b61220e565b348015610590575f80fd5b506102d361059f366004612aba565b61222a565b3480156105af575f80fd5b5061025160025481565b3480156105c4575f80fd5b506102146105d3366004612994565b612254565b3480156105e3575f80fd5b50610251612271565b5f6001600160e01b03198216637965db0b60e01b148061061c57506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f60ff1961063e60015f80516020612cab833981519152612af1565b60405160200161065091815260200190565b6040516020818303038152906040528051906020012016610670816122b8565b4661067a86612254565b6106a757604051631e859bc960e01b81526001600160a01b03871660048201526024015b60405180910390fd5b8062082750036107c45760405163095ea7b360e01b815273d8a791fe2be73eb6e6cf1eb0cb3f36adc9b3f8f96004820152602481018590526001600160a01b0387169063095ea7b3906044016020604051808303815f875af115801561070f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107339190612b04565b5060075460405163a93a4af960e01b81526001600160a01b038089166004830152871660248201526044810186905263ffffffff9091166064820181905273d8a791fe2be73eb6e6cf1eb0cb3f36adc9b3f8f99163a93a4af991906084015f604051808303815f88803b1580156107a8575f80fd5b5087f11580156107ba573d5f803e3d5ffd5b50505050506108e3565b8061044d036108c15760405163095ea7b360e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede6004820152602481018590526001600160a01b0387169063095ea7b3906044016020604051808303815f875af115801561082b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061084f9190612b04565b5060405163cd58657960e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede9063cd5865799061088f905f90899089908c908490600401612b1f565b5f604051808303815f87803b1580156108a6575f80fd5b505af11580156108b8573d5f803e3d5ffd5b505050506108e3565b600a81146108e3575b604051639474dee160e01b815260040160405180910390fd5b604080514681524260208201526001600160a01b0388811682840152871660608201526080810186905290517f1bc1b1207e0c319b7baad0317c09f30575218af6f5a59f1f0f15de8d217672cc9181900360a00190a150600195945050505050565b5f9081525f80516020612ccb833981519152602052604090206001015490565b61096e82610945565b610977816122b8565b61098183836122c2565b50505050565b6001600160a01b03811633146109b05760405163334bd91960e11b815260040160405180910390fd5b6109ba8282612363565b505050565b5f8681526003602052604081205460ff166109f057604051632ef6974160e11b81526004810188905260240161069e565b5f8681526003602052604090205460ff16610a2157604051632ef6974160e11b81526004810187905260240161069e565b604051309085156108fc029086905f818181858888f19350505050158015610a4b573d5f803e3d5ffd5b5073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f90815260056020525f80516020612c8b8339815191528054869290610a88908490612af1565b90915550505f54604051631eb65ffb60e01b81526001600160a01b0390911690631eb65ffb90610ac6908a908a908a908a908a908a90600401612b65565b5f604051808303815f87803b158015610add575f80fd5b505af1158015610aef573d5f803e3d5ffd5b5050604080518a8152602081018a90529081018790526001600160a01b03881692503091507f61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f329060600160405180910390a35060019695505050505050565b5f60ff19610b6a60015f80516020612cab833981519152612af1565b604051602001610b7c91815260200190565b6040516020818303038152906040528051906020012016610b9c816122b8565b4647841115610bbe57604051632c1d501360e11b815260040160405180910390fd5b806208275003610c5357600754604051636705b1e760e11b81526001600160a01b03871660048201526024810186905263ffffffff90911660448201819052737003e7b7186f0e6601203b99f7b8decbfa391cf99163ce0b63ce919087906064015f604051808303818589803b158015610c36575f80fd5b5088f1158015610c48573d5f803e3d5ffd5b505050505050610d26565b8061044d03610cd05760405163cd58657960e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede9063cd586579908690610c9d905f908a90849083908190600401612b1f565b5f604051808303818588803b158015610cb4575f80fd5b505af1158015610cc6573d5f803e3d5ffd5b5050505050610d26565b80600a036108ca57600754604051631474f2a960e31b81526010602160991b019163a3a79548918791610c9d9173deaddeaddeaddeaddeaddeaddeaddeaddead0000918b91859163ffffffff1690600401612b96565b604080514681524260208201526001600160a01b038716818301526060810186905290517ff1fa6dbbb3e74f1bef249f22545b238726c1cc79c5571a1416575674c7e63f399181900360800190a1506001949350505050565b5f8581526003602052604081205460ff16610db057604051632ef6974160e11b81526004810187905260240161069e565b5f8581526003602052604090205460ff16610de157604051632ef6974160e11b81526004810186905260240161069e565b610dea83612254565b610e1257604051631e859bc960e01b81526001600160a01b038416600482015260240161069e565b6040516370a0823160e01b81523060048201525f906001600160a01b038516906370a0823190602401602060405180830381865afa158015610e56573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e7a9190612bd2565b9050610e916001600160a01b0385163330866123dc565b6040516370a0823160e01b81523060048201525f906001600160a01b038616906370a0823190602401602060405180830381865afa158015610ed5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ef99190612bd2565b90505f610f068383612af1565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f90815260056020525f80516020612c8b83398151915280549293508792909190610f47908490612be9565b90915550506002545f90620f424090610f609084612bfc565b610f6a9190612c13565b9050610f768183612af1565b6001600160a01b0388165f90815260066020526040812080549294508392909190610fa2908490612be9565b90915550505f5460405163305f478560e21b81526001600160a01b039091169063c17d1e1490610fde908d908d908d908c908890600401612c32565b5f604051808303815f87803b158015610ff5575f80fd5b505af1158015611007573d5f803e3d5ffd5b5050604080518d8152602081018d90529081018590526001600160a01b03808c1693503392508a16907fece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a49060600160405180910390a45060019998505050505050505050565b5f611077816122b8565b61107f612443565b50565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f811580156110c75750825b90505f8267ffffffffffffffff1660011480156110e35750303b155b9050811580156110f1575080155b1561110f5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561113957845460ff60401b1916600160401b1785555b6111416124a2565b6111496124ac565b61115387876124bc565b6007805463ffffffff1916620493e017905583156111ab57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b5f8681526003602052604081205460ff166111e557604051632ef6974160e11b81526004810188905260240161069e565b5f8681526003602052604090205460ff1661121657604051632ef6974160e11b81526004810187905260240161069e565b5f61121f61197d565b6040516323b872dd60e01b81523060048201526001600160a01b03888116602483015260448201889052919250908216906323b872dd906064016020604051808303815f875af1158015611275573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112999190612b04565b5073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25f90815260056020527fa550ba85c46b24b567d2e17cd597f2283877afab43603f46d5de7858f1bdb73180548792906112e9908490612af1565b90915550505f54604051631eb65ffb60e01b81526001600160a01b0390911690631eb65ffb90611327908b908b908b908b908b908b90600401612b65565b5f604051808303815f87803b15801561133e575f80fd5b505af1158015611350573d5f803e3d5ffd5b5050604080518b8152602081018b90529081018890526001600160a01b03891692503091507f7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc149060600160405180910390a3506001979650505050505050565b5f6113ba816122b8565b50600255565b5f60ff196113dc60015f80516020612cab833981519152612af1565b6040516020016113ee91815260200190565b604051602081830303815290604052805190602001201661140e816122b8565b61141785612254565b61143f57604051631e859bc960e01b81526001600160a01b038616600482015260240161069e565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeed196001600160a01b038616016114be578247101561148557604051632c1d501360e11b815260040160405180910390fd5b6040516001600160a01b0385169084156108fc029085905f818181858888f193505050501580156114b8573d5f803e3d5ffd5b50611564565b6040516370a0823160e01b815230600482015283906001600160a01b038716906370a0823190602401602060405180830381865afa158015611502573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115269190612bd2565b1015611550576040516311745b6160e21b81526001600160a01b038616600482015260240161069e565b6115646001600160a01b0386168585612502565b506001949350505050565b5f611579816122b8565b61107f612533565b5f9182525f80516020612ccb833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b5f8781526003602052604081205460ff166115e857604051632ef6974160e11b81526004810189905260240161069e565b5f8781526003602052604090205460ff1661161957604051632ef6974160e11b81526004810188905260240161069e565b61162285612254565b61164a57604051631e859bc960e01b81526001600160a01b038616600482015260240161069e565b61165f6001600160a01b0386163088876123dc565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f90815260056020525f80516020612c8b833981519152805486929061169b908490612af1565b90915550505f54604051631eb65ffb60e01b81526001600160a01b0390911690631eb65ffb906116d9908b908b908b908a908a908a90600401612b65565b5f604051808303815f87803b1580156116f0575f80fd5b505af1158015611702573d5f803e3d5ffd5b5050604080518b8152602081018b90529081018790526001600160a01b03808a1693503092508816907f0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe9060600160405180910390a4506001979650505050505050565b5f8381526003602052604081205460ff1661179757604051632ef6974160e11b81526004810185905260240161069e565b5f8381526003602052604090205460ff166117c857604051632ef6974160e11b81526004810184905260240161069e565b6001543410156117f8576001546040516358f8331360e11b8152600481019190915234602482015260440161069e565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f90815260056020525f80516020612c8b8339815191528054349290611834908490612be9565b90915550506002545f90620f42409061184d9034612bfc565b6118579190612c13565b90505f6118648234612af1565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f90815260066020527fa2e5aefc6e2cbe2917a296f0fd89c5f915c487c803db1d98eccb43f14012d711805492935084929091906118b8908490612be9565b90915550505f5460405163305f478560e21b81526001600160a01b039091169063c17d1e14906118f49089908990899034908990600401612c32565b5f604051808303815f87803b15801561190b575f80fd5b505af115801561191d573d5f803e3d5ffd5b505060408051898152602081018990529081018490526001600160a01b03871692503391507f0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc59060600160405180910390a36001925050505b9392505050565b5f4662082750819003611998576004605360981b0191505090565b8061044d036119bc57734f9a0e7fd2bf6067db6994cf12e4495df938e6e991505090565b80600a036108ca576006602160991b0191505090565b5f6119dc816122b8565b506001600160a01b03919091165f908152600460205260409020805460ff1916911515919091179055565b60ff19611a2260015f80516020612cab833981519152612af1565b604051602001611a3491815260200190565b6040516020818303038152906040528051906020012016611a54816122b8565b50600155565b5f8481526003602052604081205460ff16611a8b576040516363b5c9db60e01b81526004810186905260240161069e565b5f8481526003602052604090205460ff16611abc576040516363b5c9db60e01b81526004810185905260240161069e565b5f611ac561197d565b6040516370a0823160e01b81523060048201529091505f906001600160a01b038316906370a0823190602401602060405180830381865afa158015611b0c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b309190612bd2565b6040516323b872dd60e01b8152336004820152306024820152604481018690529091506001600160a01b038316906323b872dd906064016020604051808303815f875af1158015611b83573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ba79190612b04565b506040516370a0823160e01b81523060048201525f906001600160a01b038416906370a0823190602401602060405180830381865afa158015611bec573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c109190612bd2565b90505f611c1d8383612af1565b9050600154811015611c50576001546040516358f8331360e11b815260048101919091526024810182905260440161069e565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25f90815260056020527fa550ba85c46b24b567d2e17cd597f2283877afab43603f46d5de7858f1bdb7318054839290611c9f908490612be9565b90915550506002545f90620f424090611cb89084612bfc565b611cc29190612c13565b9050611cce8183612af1565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25f90815260066020527f5e5777fab7622aff3c042c1ece74307c2e9d699a9da444f416c35f2e1def28a580549294508392909190611d22908490612be9565b90915550505f5460405163305f478560e21b81526001600160a01b039091169063c17d1e1490611d5e908d908d908d908d908890600401612c32565b5f604051808303815f87803b158015611d75575f80fd5b505af1158015611d87573d5f803e3d5ffd5b5050604080518d8152602081018d90529081018590526001600160a01b038b1692503391507fcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a899060600160405180910390a35060019998505050505050505050565b5f60ff19611e0560015f80516020612cab833981519152612af1565b604051602001611e1791815260200190565b6040516020818303038152906040528051906020012016611e37816122b8565b465f611e4161197d565b6040516370a0823160e01b81523060048201529091506001600160a01b038216906370a0823190602401602060405180830381865afa158015611e86573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611eaa9190612bd2565b851115611ed5576040516311745b6160e21b81526001600160a01b038216600482015260240161069e565b816208275003611ff25760405163095ea7b360e01b8152737003e7b7186f0e6601203b99f7b8decbfa391cf96004820152602481018690526001600160a01b0382169063095ea7b3906044016020604051808303815f875af1158015611f3d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f619190612b04565b5060075460405163a93a4af960e01b81526001600160a01b038084166004830152881660248201526044810187905263ffffffff90911660648201819052737003e7b7186f0e6601203b99f7b8decbfa391cf99163a93a4af991906084015f604051808303815f88803b158015611fd6575f80fd5b5087f1158015611fe8573d5f803e3d5ffd5b50505050506121b4565b8161044d036120e85760405163095ea7b360e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede6004820152602481018690526001600160a01b0382169063095ea7b3906044016020604051808303815f875af1158015612059573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061207d9190612b04565b5060405163cd58657960e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede9063cd5865799087906120bf905f908b90849083908190600401612b1f565b5f604051808303818588803b1580156120d6575f80fd5b505af1158015611fe8573d5f803e3d5ffd5b81600a036108ca5760405163095ea7b360e01b81526010602160991b016004820152602481018690526001600160a01b0382169063095ea7b3906044016020604051808303815f875af1158015612141573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121659190612b04565b50600754604051631474f2a960e31b81526010602160991b019163a3a795489188916120bf9173deaddeaddeaddeaddeaddeaddeaddeaddead0000918c91859163ffffffff1690600401612b96565b604080514681524260208201526001600160a01b038816818301526060810187905290517fcb1a101498aadad4159168d719f955e559e768912e6f5f5c57b094eb047d000f9181900360800190a150600195945050505050565b61221782610945565b612220816122b8565b6109818383612363565b5f612234816122b8565b505f91825260036020526040909120805460ff1916911515919091179055565b6001600160a01b03165f9081526004602052604090205460ff1690565b60ff1961228c60015f80516020612cab833981519152612af1565b60405160200161229e91815260200190565b604051602081830303815290604052805190602001201681565b61107f813361257b565b5f5f80516020612ccb8339815191526122db8484611581565b61235a575f848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556123103390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061061c565b5f91505061061c565b5f5f80516020612ccb83398151915261237c8484611581565b1561235a575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061061c565b6040516001600160a01b0384811660248301528381166044830152606482018390526109819186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506125b8565b61244b612619565b5f80516020612ceb833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b6124aa612648565b565b6124b4612648565b6124aa612691565b6124c4612648565b67016345785d8a00006001556127106002556124e05f83610965565b5f80546001600160a01b0319166001600160a01b039290921691909117905550565b6040516001600160a01b038381166024830152604482018390526109ba91859182169063a9059cbb90606401612411565b61253b6126b1565b5f80516020612ceb833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612484565b6125858282611581565b6125b45760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440161069e565b5050565b5f6125cc6001600160a01b038416836126e1565b905080515f141580156125f05750808060200190518101906125ee9190612b04565b155b156109ba57604051635274afe760e01b81526001600160a01b038416600482015260240161069e565b5f80516020612ceb8339815191525460ff166124aa57604051638dfc202b60e01b815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166124aa57604051631afcd79f60e31b815260040160405180910390fd5b612699612648565b5f80516020612ceb833981519152805460ff19169055565b5f80516020612ceb8339815191525460ff16156124aa5760405163d93c066560e01b815260040160405180910390fd5b606061197683835f845f80856001600160a01b031684866040516127059190612c5e565b5f6040518083038185875af1925050503d805f811461273f576040519150601f19603f3d011682016040523d82523d5f602084013e612744565b606091505b509150915061275486838361275e565b9695505050505050565b6060826127735761276e826127ba565b611976565b815115801561278a57506001600160a01b0384163b155b156127b357604051639996b31560e01b81526001600160a01b038516600482015260240161069e565b5080611976565b8051156127ca5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f602082840312156127f3575f80fd5b81356001600160e01b031981168114611976575f80fd5b6001600160a01b038116811461107f575f80fd5b5f805f60608486031215612830575f80fd5b833561283b8161280a565b9250602084013561284b8161280a565b929592945050506040919091013590565b5f6020828403121561286c575f80fd5b5035919050565b5f8060408385031215612884575f80fd5b8235915060208301356128968161280a565b809150509250929050565b5f805f805f8060c087890312156128b6575f80fd5b863595506020870135945060408701356128cf8161280a565b959894975094956060810135955060808101359460a0909101359350915050565b5f8060408385031215612901575f80fd5b823561290c8161280a565b946020939093013593505050565b5f805f805f60a0868803121561292e575f80fd5b853594506020860135935060408601356129478161280a565b925060608601356129578161280a565b949793965091946080013592915050565b5f8060408385031215612979575f80fd5b82356129848161280a565b915060208301356128968161280a565b5f602082840312156129a4575f80fd5b81356119768161280a565b5f805f805f805f60e0888a0312156129c5575f80fd5b873596506020880135955060408801356129de8161280a565b945060608801356129ee8161280a565b9699959850939660808101359560a0820135955060c0909101359350915050565b5f805f60608486031215612a21575f80fd5b83359250602084013591506040840135612a3a8161280a565b809150509250925092565b801515811461107f575f80fd5b5f8060408385031215612a63575f80fd5b8235612a6e8161280a565b9150602083013561289681612a45565b5f805f8060808587031215612a91575f80fd5b84359350602085013592506040850135612aaa8161280a565b9396929550929360600135925050565b5f8060408385031215612acb575f80fd5b82359150602083013561289681612a45565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561061c5761061c612add565b5f60208284031215612b14575f80fd5b815161197681612a45565b63ffffffff9590951685526001600160a01b039384166020860152604085019290925290911660608301521515608082015260c060a082018190525f9082015260e00190565b95865260208601949094526001600160a01b039290921660408501526060840152608083015260a082015260c00190565b6001600160a01b039485168152929093166020830152604082015263ffffffff909116606082015260a0608082018190525f9082015260c00190565b5f60208284031215612be2575f80fd5b5051919050565b8082018082111561061c5761061c612add565b808202811582820484141761061c5761061c612add565b5f82612c2d57634e487b7160e01b5f52601260045260245ffd5b500490565b94855260208501939093526001600160a01b039190911660408401526060830152608082015260a00190565b5f82515f5b81811015612c7d5760208186018101518583015201612c63565b505f92019182525091905056fea1829a9003092132f585b6ccdd167c19fe9774dbdea4260287e8a8e8ca8185d733fe247d78fee2e7fd135c405eda4bd2911c0a73c0a81b36c3bcc967dd06e5ae02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300a2646970667358221220a717caba8890bfd4ed6ed69397e36375a72f25470560d98bb8a351c7bd598cf364736f6c63430008140033",
}

// L2PoolManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use L2PoolManagerMetaData.ABI instead.
var L2PoolManagerABI = L2PoolManagerMetaData.ABI

// L2PoolManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2PoolManagerMetaData.Bin instead.
var L2PoolManagerBin = L2PoolManagerMetaData.Bin

// DeployL2PoolManager deploys a new Ethereum contract, binding an instance of L2PoolManager to it.
func DeployL2PoolManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2PoolManager, error) {
	parsed, err := L2PoolManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2PoolManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2PoolManager{L2PoolManagerCaller: L2PoolManagerCaller{contract: contract}, L2PoolManagerTransactor: L2PoolManagerTransactor{contract: contract}, L2PoolManagerFilterer: L2PoolManagerFilterer{contract: contract}}, nil
}

// L2PoolManager is an auto generated Go binding around an Ethereum contract.
type L2PoolManager struct {
	L2PoolManagerCaller     // Read-only binding to the contract
	L2PoolManagerTransactor // Write-only binding to the contract
	L2PoolManagerFilterer   // Log filterer for contract events
}

// L2PoolManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2PoolManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2PoolManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2PoolManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2PoolManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2PoolManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2PoolManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2PoolManagerSession struct {
	Contract     *L2PoolManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2PoolManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2PoolManagerCallerSession struct {
	Contract *L2PoolManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L2PoolManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2PoolManagerTransactorSession struct {
	Contract     *L2PoolManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L2PoolManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2PoolManagerRaw struct {
	Contract *L2PoolManager // Generic contract binding to access the raw methods on
}

// L2PoolManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2PoolManagerCallerRaw struct {
	Contract *L2PoolManagerCaller // Generic read-only contract binding to access the raw methods on
}

// L2PoolManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2PoolManagerTransactorRaw struct {
	Contract *L2PoolManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2PoolManager creates a new instance of L2PoolManager, bound to a specific deployed contract.
func NewL2PoolManager(address common.Address, backend bind.ContractBackend) (*L2PoolManager, error) {
	contract, err := bindL2PoolManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2PoolManager{L2PoolManagerCaller: L2PoolManagerCaller{contract: contract}, L2PoolManagerTransactor: L2PoolManagerTransactor{contract: contract}, L2PoolManagerFilterer: L2PoolManagerFilterer{contract: contract}}, nil
}

// NewL2PoolManagerCaller creates a new read-only instance of L2PoolManager, bound to a specific deployed contract.
func NewL2PoolManagerCaller(address common.Address, caller bind.ContractCaller) (*L2PoolManagerCaller, error) {
	contract, err := bindL2PoolManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerCaller{contract: contract}, nil
}

// NewL2PoolManagerTransactor creates a new write-only instance of L2PoolManager, bound to a specific deployed contract.
func NewL2PoolManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*L2PoolManagerTransactor, error) {
	contract, err := bindL2PoolManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerTransactor{contract: contract}, nil
}

// NewL2PoolManagerFilterer creates a new log filterer instance of L2PoolManager, bound to a specific deployed contract.
func NewL2PoolManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*L2PoolManagerFilterer, error) {
	contract, err := bindL2PoolManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerFilterer{contract: contract}, nil
}

// bindL2PoolManager binds a generic wrapper to an already deployed contract.
func bindL2PoolManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2PoolManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2PoolManager *L2PoolManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2PoolManager.Contract.L2PoolManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2PoolManager *L2PoolManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2PoolManager.Contract.L2PoolManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2PoolManager *L2PoolManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2PoolManager.Contract.L2PoolManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2PoolManager *L2PoolManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2PoolManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2PoolManager *L2PoolManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2PoolManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2PoolManager *L2PoolManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2PoolManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L2PoolManager *L2PoolManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L2PoolManager *L2PoolManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _L2PoolManager.Contract.DEFAULTADMINROLE(&_L2PoolManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L2PoolManager *L2PoolManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _L2PoolManager.Contract.DEFAULTADMINROLE(&_L2PoolManager.CallOpts)
}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_L2PoolManager *L2PoolManagerCaller) FeePoolValue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "FeePoolValue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_L2PoolManager *L2PoolManagerSession) FeePoolValue(arg0 common.Address) (*big.Int, error) {
	return _L2PoolManager.Contract.FeePoolValue(&_L2PoolManager.CallOpts, arg0)
}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_L2PoolManager *L2PoolManagerCallerSession) FeePoolValue(arg0 common.Address) (*big.Int, error) {
	return _L2PoolManager.Contract.FeePoolValue(&_L2PoolManager.CallOpts, arg0)
}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_L2PoolManager *L2PoolManagerCaller) IsSupportChainId(opts *bind.CallOpts, chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "IsSupportChainId", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_L2PoolManager *L2PoolManagerSession) IsSupportChainId(chainId *big.Int) (bool, error) {
	return _L2PoolManager.Contract.IsSupportChainId(&_L2PoolManager.CallOpts, chainId)
}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_L2PoolManager *L2PoolManagerCallerSession) IsSupportChainId(chainId *big.Int) (bool, error) {
	return _L2PoolManager.Contract.IsSupportChainId(&_L2PoolManager.CallOpts, chainId)
}

// IsSupportStableCoin is a free data retrieval call binding the contract method 0xed06da87.
//
// Solidity: function IsSupportStableCoin(address ERC20Address) view returns(bool)
func (_L2PoolManager *L2PoolManagerCaller) IsSupportStableCoin(opts *bind.CallOpts, ERC20Address common.Address) (bool, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "IsSupportStableCoin", ERC20Address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportStableCoin is a free data retrieval call binding the contract method 0xed06da87.
//
// Solidity: function IsSupportStableCoin(address ERC20Address) view returns(bool)
func (_L2PoolManager *L2PoolManagerSession) IsSupportStableCoin(ERC20Address common.Address) (bool, error) {
	return _L2PoolManager.Contract.IsSupportStableCoin(&_L2PoolManager.CallOpts, ERC20Address)
}

// IsSupportStableCoin is a free data retrieval call binding the contract method 0xed06da87.
//
// Solidity: function IsSupportStableCoin(address ERC20Address) view returns(bool)
func (_L2PoolManager *L2PoolManagerCallerSession) IsSupportStableCoin(ERC20Address common.Address) (bool, error) {
	return _L2PoolManager.Contract.IsSupportStableCoin(&_L2PoolManager.CallOpts, ERC20Address)
}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_L2PoolManager *L2PoolManagerCaller) L2WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "L2WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_L2PoolManager *L2PoolManagerSession) L2WETH() (common.Address, error) {
	return _L2PoolManager.Contract.L2WETH(&_L2PoolManager.CallOpts)
}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_L2PoolManager *L2PoolManagerCallerSession) L2WETH() (common.Address, error) {
	return _L2PoolManager.Contract.L2WETH(&_L2PoolManager.CallOpts)
}

// MAXGASLimit is a free data retrieval call binding the contract method 0xa525ae2b.
//
// Solidity: function MAX_GAS_Limit() view returns(uint32)
func (_L2PoolManager *L2PoolManagerCaller) MAXGASLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "MAX_GAS_Limit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXGASLimit is a free data retrieval call binding the contract method 0xa525ae2b.
//
// Solidity: function MAX_GAS_Limit() view returns(uint32)
func (_L2PoolManager *L2PoolManagerSession) MAXGASLimit() (uint32, error) {
	return _L2PoolManager.Contract.MAXGASLimit(&_L2PoolManager.CallOpts)
}

// MAXGASLimit is a free data retrieval call binding the contract method 0xa525ae2b.
//
// Solidity: function MAX_GAS_Limit() view returns(uint32)
func (_L2PoolManager *L2PoolManagerCallerSession) MAXGASLimit() (uint32, error) {
	return _L2PoolManager.Contract.MAXGASLimit(&_L2PoolManager.CallOpts)
}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_L2PoolManager *L2PoolManagerCaller) MinTransferAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "MinTransferAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_L2PoolManager *L2PoolManagerSession) MinTransferAmount() (*big.Int, error) {
	return _L2PoolManager.Contract.MinTransferAmount(&_L2PoolManager.CallOpts)
}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_L2PoolManager *L2PoolManagerCallerSession) MinTransferAmount() (*big.Int, error) {
	return _L2PoolManager.Contract.MinTransferAmount(&_L2PoolManager.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L2PoolManager *L2PoolManagerCaller) PerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "PerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L2PoolManager *L2PoolManagerSession) PerFee() (*big.Int, error) {
	return _L2PoolManager.Contract.PerFee(&_L2PoolManager.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L2PoolManager *L2PoolManagerCallerSession) PerFee() (*big.Int, error) {
	return _L2PoolManager.Contract.PerFee(&_L2PoolManager.CallOpts)
}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_L2PoolManager *L2PoolManagerCaller) ReLayer(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "ReLayer")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_L2PoolManager *L2PoolManagerSession) ReLayer() ([32]byte, error) {
	return _L2PoolManager.Contract.ReLayer(&_L2PoolManager.CallOpts)
}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_L2PoolManager *L2PoolManagerCallerSession) ReLayer() ([32]byte, error) {
	return _L2PoolManager.Contract.ReLayer(&_L2PoolManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L2PoolManager *L2PoolManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L2PoolManager *L2PoolManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _L2PoolManager.Contract.GetRoleAdmin(&_L2PoolManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L2PoolManager *L2PoolManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _L2PoolManager.Contract.GetRoleAdmin(&_L2PoolManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L2PoolManager *L2PoolManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L2PoolManager *L2PoolManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _L2PoolManager.Contract.HasRole(&_L2PoolManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L2PoolManager *L2PoolManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _L2PoolManager.Contract.HasRole(&_L2PoolManager.CallOpts, role, account)
}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_L2PoolManager *L2PoolManagerCaller) MessageManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "messageManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_L2PoolManager *L2PoolManagerSession) MessageManager() (common.Address, error) {
	return _L2PoolManager.Contract.MessageManager(&_L2PoolManager.CallOpts)
}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_L2PoolManager *L2PoolManagerCallerSession) MessageManager() (common.Address, error) {
	return _L2PoolManager.Contract.MessageManager(&_L2PoolManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2PoolManager *L2PoolManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2PoolManager *L2PoolManagerSession) Paused() (bool, error) {
	return _L2PoolManager.Contract.Paused(&_L2PoolManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2PoolManager *L2PoolManagerCallerSession) Paused() (bool, error) {
	return _L2PoolManager.Contract.Paused(&_L2PoolManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2PoolManager *L2PoolManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2PoolManager *L2PoolManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2PoolManager.Contract.SupportsInterface(&_L2PoolManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2PoolManager *L2PoolManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2PoolManager.Contract.SupportsInterface(&_L2PoolManager.CallOpts, interfaceId)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x92d84738.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeFinalizeERC20(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeFinalizeERC20", sourceChainId, destChainId, to, ERC20Address, amount, _fee, _nonce)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x92d84738.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeFinalizeERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeFinalizeERC20(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, amount, _fee, _nonce)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x92d84738.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeFinalizeERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeFinalizeERC20(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x38731c47.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeFinalizeETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeFinalizeETH", sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x38731c47.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeFinalizeETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeFinalizeETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x38731c47.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeFinalizeETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeFinalizeETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeWETH is a paid mutator transaction binding the contract method 0x5765a4eb.
//
// Solidity: function BridgeFinalizeWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeFinalizeWETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeFinalizeWETH", sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeWETH is a paid mutator transaction binding the contract method 0x5765a4eb.
//
// Solidity: function BridgeFinalizeWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeFinalizeWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeFinalizeWETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeWETH is a paid mutator transaction binding the contract method 0x5765a4eb.
//
// Solidity: function BridgeFinalizeWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeFinalizeWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeFinalizeWETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x3b0230f0.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeInitiateERC20(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeInitiateERC20", sourceChainId, destChainId, to, ERC20Address, value)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x3b0230f0.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeInitiateERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateERC20(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, value)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x3b0230f0.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeInitiateERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateERC20(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, value)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x9b442380.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address to) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeInitiateETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeInitiateETH", sourceChainId, destChainId, to)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x9b442380.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address to) payable returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeInitiateETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x9b442380.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address to) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeInitiateETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to)
}

// BridgeInitiateWETH is a paid mutator transaction binding the contract method 0xcd01c665.
//
// Solidity: function BridgeInitiateWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 value) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeInitiateWETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeInitiateWETH", sourceChainId, destChainId, to, value)
}

// BridgeInitiateWETH is a paid mutator transaction binding the contract method 0xcd01c665.
//
// Solidity: function BridgeInitiateWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 value) returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeInitiateWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateWETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, value)
}

// BridgeInitiateWETH is a paid mutator transaction binding the contract method 0xcd01c665.
//
// Solidity: function BridgeInitiateWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 value) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeInitiateWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateWETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, value)
}

// SendAssertToUser is a paid mutator transaction binding the contract method 0x73c23be9.
//
// Solidity: function SendAssertToUser(address _token, address to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) SendAssertToUser(opts *bind.TransactOpts, _token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "SendAssertToUser", _token, to, _amount)
}

// SendAssertToUser is a paid mutator transaction binding the contract method 0x73c23be9.
//
// Solidity: function SendAssertToUser(address _token, address to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerSession) SendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SendAssertToUser(&_L2PoolManager.TransactOpts, _token, to, _amount)
}

// SendAssertToUser is a paid mutator transaction binding the contract method 0x73c23be9.
//
// Solidity: function SendAssertToUser(address _token, address to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) SendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SendAssertToUser(&_L2PoolManager.TransactOpts, _token, to, _amount)
}

// WithdrawERC20ToL1 is a paid mutator transaction binding the contract method 0x0227047a.
//
// Solidity: function WithdrawERC20ToL1(address _token, address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) WithdrawERC20ToL1(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "WithdrawERC20ToL1", _token, _to, _amount)
}

// WithdrawERC20ToL1 is a paid mutator transaction binding the contract method 0x0227047a.
//
// Solidity: function WithdrawERC20ToL1(address _token, address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerSession) WithdrawERC20ToL1(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.WithdrawERC20ToL1(&_L2PoolManager.TransactOpts, _token, _to, _amount)
}

// WithdrawERC20ToL1 is a paid mutator transaction binding the contract method 0x0227047a.
//
// Solidity: function WithdrawERC20ToL1(address _token, address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) WithdrawERC20ToL1(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.WithdrawERC20ToL1(&_L2PoolManager.TransactOpts, _token, _to, _amount)
}

// WithdrawETHtoL1 is a paid mutator transaction binding the contract method 0x39891dd7.
//
// Solidity: function WithdrawETHtoL1(address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) WithdrawETHtoL1(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "WithdrawETHtoL1", _to, _amount)
}

// WithdrawETHtoL1 is a paid mutator transaction binding the contract method 0x39891dd7.
//
// Solidity: function WithdrawETHtoL1(address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerSession) WithdrawETHtoL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.WithdrawETHtoL1(&_L2PoolManager.TransactOpts, _to, _amount)
}

// WithdrawETHtoL1 is a paid mutator transaction binding the contract method 0x39891dd7.
//
// Solidity: function WithdrawETHtoL1(address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) WithdrawETHtoL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.WithdrawETHtoL1(&_L2PoolManager.TransactOpts, _to, _amount)
}

// WithdrawWETHToL1 is a paid mutator transaction binding the contract method 0xd28597f4.
//
// Solidity: function WithdrawWETHToL1(address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) WithdrawWETHToL1(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "WithdrawWETHToL1", _to, _amount)
}

// WithdrawWETHToL1 is a paid mutator transaction binding the contract method 0xd28597f4.
//
// Solidity: function WithdrawWETHToL1(address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerSession) WithdrawWETHToL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.WithdrawWETHToL1(&_L2PoolManager.TransactOpts, _to, _amount)
}

// WithdrawWETHToL1 is a paid mutator transaction binding the contract method 0xd28597f4.
//
// Solidity: function WithdrawWETHToL1(address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) WithdrawWETHToL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.WithdrawWETHToL1(&_L2PoolManager.TransactOpts, _to, _amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L2PoolManager *L2PoolManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L2PoolManager *L2PoolManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.GrantRole(&_L2PoolManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.GrantRole(&_L2PoolManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _MultisigWallet, address _messageManager) returns()
func (_L2PoolManager *L2PoolManagerTransactor) Initialize(opts *bind.TransactOpts, _MultisigWallet common.Address, _messageManager common.Address) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "initialize", _MultisigWallet, _messageManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _MultisigWallet, address _messageManager) returns()
func (_L2PoolManager *L2PoolManagerSession) Initialize(_MultisigWallet common.Address, _messageManager common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.Initialize(&_L2PoolManager.TransactOpts, _MultisigWallet, _messageManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _MultisigWallet, address _messageManager) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) Initialize(_MultisigWallet common.Address, _messageManager common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.Initialize(&_L2PoolManager.TransactOpts, _MultisigWallet, _messageManager)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L2PoolManager *L2PoolManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L2PoolManager *L2PoolManagerSession) Pause() (*types.Transaction, error) {
	return _L2PoolManager.Contract.Pause(&_L2PoolManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _L2PoolManager.Contract.Pause(&_L2PoolManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L2PoolManager *L2PoolManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L2PoolManager *L2PoolManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.RenounceRole(&_L2PoolManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.RenounceRole(&_L2PoolManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L2PoolManager *L2PoolManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L2PoolManager *L2PoolManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.RevokeRole(&_L2PoolManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.RevokeRole(&_L2PoolManager.TransactOpts, role, account)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_L2PoolManager *L2PoolManagerTransactor) SetMinTransferAmount(opts *bind.TransactOpts, _MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "setMinTransferAmount", _MinTransferAmount)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_L2PoolManager *L2PoolManagerSession) SetMinTransferAmount(_MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetMinTransferAmount(&_L2PoolManager.TransactOpts, _MinTransferAmount)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) SetMinTransferAmount(_MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetMinTransferAmount(&_L2PoolManager.TransactOpts, _MinTransferAmount)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_L2PoolManager *L2PoolManagerTransactor) SetPerFee(opts *bind.TransactOpts, _PerFee *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "setPerFee", _PerFee)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_L2PoolManager *L2PoolManagerSession) SetPerFee(_PerFee *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetPerFee(&_L2PoolManager.TransactOpts, _PerFee)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) SetPerFee(_PerFee *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetPerFee(&_L2PoolManager.TransactOpts, _PerFee)
}

// SetSupportStableCoin is a paid mutator transaction binding the contract method 0xb33280dd.
//
// Solidity: function setSupportStableCoin(address ERC20Address, bool isValid) returns()
func (_L2PoolManager *L2PoolManagerTransactor) SetSupportStableCoin(opts *bind.TransactOpts, ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "setSupportStableCoin", ERC20Address, isValid)
}

// SetSupportStableCoin is a paid mutator transaction binding the contract method 0xb33280dd.
//
// Solidity: function setSupportStableCoin(address ERC20Address, bool isValid) returns()
func (_L2PoolManager *L2PoolManagerSession) SetSupportStableCoin(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetSupportStableCoin(&_L2PoolManager.TransactOpts, ERC20Address, isValid)
}

// SetSupportStableCoin is a paid mutator transaction binding the contract method 0xb33280dd.
//
// Solidity: function setSupportStableCoin(address ERC20Address, bool isValid) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) SetSupportStableCoin(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetSupportStableCoin(&_L2PoolManager.TransactOpts, ERC20Address, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_L2PoolManager *L2PoolManagerTransactor) SetValidChainId(opts *bind.TransactOpts, chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "setValidChainId", chainId, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_L2PoolManager *L2PoolManagerSession) SetValidChainId(chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetValidChainId(&_L2PoolManager.TransactOpts, chainId, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) SetValidChainId(chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetValidChainId(&_L2PoolManager.TransactOpts, chainId, isValid)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L2PoolManager *L2PoolManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L2PoolManager *L2PoolManagerSession) Unpause() (*types.Transaction, error) {
	return _L2PoolManager.Contract.Unpause(&_L2PoolManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _L2PoolManager.Contract.Unpause(&_L2PoolManager.TransactOpts)
}

// L2PoolManagerFinalizeERC20Iterator is returned from FilterFinalizeERC20 and is used to iterate over the raw logs and unpacked data for FinalizeERC20 events raised by the L2PoolManager contract.
type L2PoolManagerFinalizeERC20Iterator struct {
	Event *L2PoolManagerFinalizeERC20 // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerFinalizeERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerFinalizeERC20)
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
		it.Event = new(L2PoolManagerFinalizeERC20)
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
func (it *L2PoolManagerFinalizeERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerFinalizeERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerFinalizeERC20 represents a FinalizeERC20 event raised by the L2PoolManager contract.
type L2PoolManagerFinalizeERC20 struct {
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
func (_L2PoolManager *L2PoolManagerFilterer) FilterFinalizeERC20(opts *bind.FilterOpts, ERC20Address []common.Address, from []common.Address, to []common.Address) (*L2PoolManagerFinalizeERC20Iterator, error) {

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

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "FinalizeERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerFinalizeERC20Iterator{contract: _L2PoolManager.contract, event: "FinalizeERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeERC20 is a free log subscription operation binding the contract event 0x0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe.
//
// Solidity: event FinalizeERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchFinalizeERC20(opts *bind.WatchOpts, sink chan<- *L2PoolManagerFinalizeERC20, ERC20Address []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "FinalizeERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerFinalizeERC20)
				if err := _L2PoolManager.contract.UnpackLog(event, "FinalizeERC20", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseFinalizeERC20(log types.Log) (*L2PoolManagerFinalizeERC20, error) {
	event := new(L2PoolManagerFinalizeERC20)
	if err := _L2PoolManager.contract.UnpackLog(event, "FinalizeERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerFinalizeETHIterator is returned from FilterFinalizeETH and is used to iterate over the raw logs and unpacked data for FinalizeETH events raised by the L2PoolManager contract.
type L2PoolManagerFinalizeETHIterator struct {
	Event *L2PoolManagerFinalizeETH // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerFinalizeETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerFinalizeETH)
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
		it.Event = new(L2PoolManagerFinalizeETH)
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
func (it *L2PoolManagerFinalizeETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerFinalizeETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerFinalizeETH represents a FinalizeETH event raised by the L2PoolManager contract.
type L2PoolManagerFinalizeETH struct {
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
func (_L2PoolManager *L2PoolManagerFilterer) FilterFinalizeETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2PoolManagerFinalizeETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "FinalizeETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerFinalizeETHIterator{contract: _L2PoolManager.contract, event: "FinalizeETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeETH is a free log subscription operation binding the contract event 0x61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f32.
//
// Solidity: event FinalizeETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchFinalizeETH(opts *bind.WatchOpts, sink chan<- *L2PoolManagerFinalizeETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "FinalizeETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerFinalizeETH)
				if err := _L2PoolManager.contract.UnpackLog(event, "FinalizeETH", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseFinalizeETH(log types.Log) (*L2PoolManagerFinalizeETH, error) {
	event := new(L2PoolManagerFinalizeETH)
	if err := _L2PoolManager.contract.UnpackLog(event, "FinalizeETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerFinalizeWETHIterator is returned from FilterFinalizeWETH and is used to iterate over the raw logs and unpacked data for FinalizeWETH events raised by the L2PoolManager contract.
type L2PoolManagerFinalizeWETHIterator struct {
	Event *L2PoolManagerFinalizeWETH // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerFinalizeWETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerFinalizeWETH)
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
		it.Event = new(L2PoolManagerFinalizeWETH)
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
func (it *L2PoolManagerFinalizeWETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerFinalizeWETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerFinalizeWETH represents a FinalizeWETH event raised by the L2PoolManager contract.
type L2PoolManagerFinalizeWETH struct {
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
func (_L2PoolManager *L2PoolManagerFilterer) FilterFinalizeWETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2PoolManagerFinalizeWETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "FinalizeWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerFinalizeWETHIterator{contract: _L2PoolManager.contract, event: "FinalizeWETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeWETH is a free log subscription operation binding the contract event 0x7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc14.
//
// Solidity: event FinalizeWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchFinalizeWETH(opts *bind.WatchOpts, sink chan<- *L2PoolManagerFinalizeWETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "FinalizeWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerFinalizeWETH)
				if err := _L2PoolManager.contract.UnpackLog(event, "FinalizeWETH", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseFinalizeWETH(log types.Log) (*L2PoolManagerFinalizeWETH, error) {
	event := new(L2PoolManagerFinalizeWETH)
	if err := _L2PoolManager.contract.UnpackLog(event, "FinalizeWETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2PoolManager contract.
type L2PoolManagerInitializedIterator struct {
	Event *L2PoolManagerInitialized // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerInitialized)
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
		it.Event = new(L2PoolManagerInitialized)
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
func (it *L2PoolManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerInitialized represents a Initialized event raised by the L2PoolManager contract.
type L2PoolManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L2PoolManager *L2PoolManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2PoolManagerInitializedIterator, error) {

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerInitializedIterator{contract: _L2PoolManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L2PoolManager *L2PoolManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2PoolManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerInitialized)
				if err := _L2PoolManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseInitialized(log types.Log) (*L2PoolManagerInitialized, error) {
	event := new(L2PoolManagerInitialized)
	if err := _L2PoolManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerInitiateERC20Iterator is returned from FilterInitiateERC20 and is used to iterate over the raw logs and unpacked data for InitiateERC20 events raised by the L2PoolManager contract.
type L2PoolManagerInitiateERC20Iterator struct {
	Event *L2PoolManagerInitiateERC20 // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerInitiateERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerInitiateERC20)
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
		it.Event = new(L2PoolManagerInitiateERC20)
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
func (it *L2PoolManagerInitiateERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerInitiateERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerInitiateERC20 represents a InitiateERC20 event raised by the L2PoolManager contract.
type L2PoolManagerInitiateERC20 struct {
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
func (_L2PoolManager *L2PoolManagerFilterer) FilterInitiateERC20(opts *bind.FilterOpts, ERC20Address []common.Address, from []common.Address, to []common.Address) (*L2PoolManagerInitiateERC20Iterator, error) {

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

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "InitiateERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerInitiateERC20Iterator{contract: _L2PoolManager.contract, event: "InitiateERC20", logs: logs, sub: sub}, nil
}

// WatchInitiateERC20 is a free log subscription operation binding the contract event 0xece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a4.
//
// Solidity: event InitiateERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchInitiateERC20(opts *bind.WatchOpts, sink chan<- *L2PoolManagerInitiateERC20, ERC20Address []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "InitiateERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerInitiateERC20)
				if err := _L2PoolManager.contract.UnpackLog(event, "InitiateERC20", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseInitiateERC20(log types.Log) (*L2PoolManagerInitiateERC20, error) {
	event := new(L2PoolManagerInitiateERC20)
	if err := _L2PoolManager.contract.UnpackLog(event, "InitiateERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerInitiateETHIterator is returned from FilterInitiateETH and is used to iterate over the raw logs and unpacked data for InitiateETH events raised by the L2PoolManager contract.
type L2PoolManagerInitiateETHIterator struct {
	Event *L2PoolManagerInitiateETH // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerInitiateETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerInitiateETH)
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
		it.Event = new(L2PoolManagerInitiateETH)
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
func (it *L2PoolManagerInitiateETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerInitiateETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerInitiateETH represents a InitiateETH event raised by the L2PoolManager contract.
type L2PoolManagerInitiateETH struct {
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
func (_L2PoolManager *L2PoolManagerFilterer) FilterInitiateETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2PoolManagerInitiateETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "InitiateETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerInitiateETHIterator{contract: _L2PoolManager.contract, event: "InitiateETH", logs: logs, sub: sub}, nil
}

// WatchInitiateETH is a free log subscription operation binding the contract event 0x0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc5.
//
// Solidity: event InitiateETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchInitiateETH(opts *bind.WatchOpts, sink chan<- *L2PoolManagerInitiateETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "InitiateETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerInitiateETH)
				if err := _L2PoolManager.contract.UnpackLog(event, "InitiateETH", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseInitiateETH(log types.Log) (*L2PoolManagerInitiateETH, error) {
	event := new(L2PoolManagerInitiateETH)
	if err := _L2PoolManager.contract.UnpackLog(event, "InitiateETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerInitiateWETHIterator is returned from FilterInitiateWETH and is used to iterate over the raw logs and unpacked data for InitiateWETH events raised by the L2PoolManager contract.
type L2PoolManagerInitiateWETHIterator struct {
	Event *L2PoolManagerInitiateWETH // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerInitiateWETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerInitiateWETH)
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
		it.Event = new(L2PoolManagerInitiateWETH)
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
func (it *L2PoolManagerInitiateWETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerInitiateWETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerInitiateWETH represents a InitiateWETH event raised by the L2PoolManager contract.
type L2PoolManagerInitiateWETH struct {
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
func (_L2PoolManager *L2PoolManagerFilterer) FilterInitiateWETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2PoolManagerInitiateWETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "InitiateWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerInitiateWETHIterator{contract: _L2PoolManager.contract, event: "InitiateWETH", logs: logs, sub: sub}, nil
}

// WatchInitiateWETH is a free log subscription operation binding the contract event 0xcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a89.
//
// Solidity: event InitiateWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchInitiateWETH(opts *bind.WatchOpts, sink chan<- *L2PoolManagerInitiateWETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "InitiateWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerInitiateWETH)
				if err := _L2PoolManager.contract.UnpackLog(event, "InitiateWETH", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseInitiateWETH(log types.Log) (*L2PoolManagerInitiateWETH, error) {
	event := new(L2PoolManagerInitiateWETH)
	if err := _L2PoolManager.contract.UnpackLog(event, "InitiateWETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the L2PoolManager contract.
type L2PoolManagerPausedIterator struct {
	Event *L2PoolManagerPaused // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerPaused)
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
		it.Event = new(L2PoolManagerPaused)
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
func (it *L2PoolManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerPaused represents a Paused event raised by the L2PoolManager contract.
type L2PoolManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2PoolManager *L2PoolManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*L2PoolManagerPausedIterator, error) {

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerPausedIterator{contract: _L2PoolManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2PoolManager *L2PoolManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *L2PoolManagerPaused) (event.Subscription, error) {

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerPaused)
				if err := _L2PoolManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParsePaused(log types.Log) (*L2PoolManagerPaused, error) {
	event := new(L2PoolManagerPaused)
	if err := _L2PoolManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the L2PoolManager contract.
type L2PoolManagerRoleAdminChangedIterator struct {
	Event *L2PoolManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerRoleAdminChanged)
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
		it.Event = new(L2PoolManagerRoleAdminChanged)
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
func (it *L2PoolManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerRoleAdminChanged represents a RoleAdminChanged event raised by the L2PoolManager contract.
type L2PoolManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_L2PoolManager *L2PoolManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*L2PoolManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerRoleAdminChangedIterator{contract: _L2PoolManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_L2PoolManager *L2PoolManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *L2PoolManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerRoleAdminChanged)
				if err := _L2PoolManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseRoleAdminChanged(log types.Log) (*L2PoolManagerRoleAdminChanged, error) {
	event := new(L2PoolManagerRoleAdminChanged)
	if err := _L2PoolManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the L2PoolManager contract.
type L2PoolManagerRoleGrantedIterator struct {
	Event *L2PoolManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerRoleGranted)
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
		it.Event = new(L2PoolManagerRoleGranted)
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
func (it *L2PoolManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerRoleGranted represents a RoleGranted event raised by the L2PoolManager contract.
type L2PoolManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2PoolManager *L2PoolManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L2PoolManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerRoleGrantedIterator{contract: _L2PoolManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2PoolManager *L2PoolManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *L2PoolManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerRoleGranted)
				if err := _L2PoolManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseRoleGranted(log types.Log) (*L2PoolManagerRoleGranted, error) {
	event := new(L2PoolManagerRoleGranted)
	if err := _L2PoolManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the L2PoolManager contract.
type L2PoolManagerRoleRevokedIterator struct {
	Event *L2PoolManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerRoleRevoked)
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
		it.Event = new(L2PoolManagerRoleRevoked)
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
func (it *L2PoolManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerRoleRevoked represents a RoleRevoked event raised by the L2PoolManager contract.
type L2PoolManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2PoolManager *L2PoolManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L2PoolManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerRoleRevokedIterator{contract: _L2PoolManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2PoolManager *L2PoolManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *L2PoolManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerRoleRevoked)
				if err := _L2PoolManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseRoleRevoked(log types.Log) (*L2PoolManagerRoleRevoked, error) {
	event := new(L2PoolManagerRoleRevoked)
	if err := _L2PoolManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the L2PoolManager contract.
type L2PoolManagerUnpausedIterator struct {
	Event *L2PoolManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerUnpaused)
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
		it.Event = new(L2PoolManagerUnpaused)
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
func (it *L2PoolManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerUnpaused represents a Unpaused event raised by the L2PoolManager contract.
type L2PoolManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2PoolManager *L2PoolManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*L2PoolManagerUnpausedIterator, error) {

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerUnpausedIterator{contract: _L2PoolManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2PoolManager *L2PoolManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *L2PoolManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerUnpaused)
				if err := _L2PoolManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseUnpaused(log types.Log) (*L2PoolManagerUnpaused, error) {
	event := new(L2PoolManagerUnpaused)
	if err := _L2PoolManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerWithdrawERC20toL1SuccessIterator is returned from FilterWithdrawERC20toL1Success and is used to iterate over the raw logs and unpacked data for WithdrawERC20toL1Success events raised by the L2PoolManager contract.
type L2PoolManagerWithdrawERC20toL1SuccessIterator struct {
	Event *L2PoolManagerWithdrawERC20toL1Success // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerWithdrawERC20toL1SuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerWithdrawERC20toL1Success)
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
		it.Event = new(L2PoolManagerWithdrawERC20toL1Success)
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
func (it *L2PoolManagerWithdrawERC20toL1SuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerWithdrawERC20toL1SuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerWithdrawERC20toL1Success represents a WithdrawERC20toL1Success event raised by the L2PoolManager contract.
type L2PoolManagerWithdrawERC20toL1Success struct {
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
func (_L2PoolManager *L2PoolManagerFilterer) FilterWithdrawERC20toL1Success(opts *bind.FilterOpts) (*L2PoolManagerWithdrawERC20toL1SuccessIterator, error) {

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "WithdrawERC20toL1Success")
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerWithdrawERC20toL1SuccessIterator{contract: _L2PoolManager.contract, event: "WithdrawERC20toL1Success", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20toL1Success is a free log subscription operation binding the contract event 0x1bc1b1207e0c319b7baad0317c09f30575218af6f5a59f1f0f15de8d217672cc.
//
// Solidity: event WithdrawERC20toL1Success(uint256 chainId, uint256 timestamp, address token, address to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchWithdrawERC20toL1Success(opts *bind.WatchOpts, sink chan<- *L2PoolManagerWithdrawERC20toL1Success) (event.Subscription, error) {

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "WithdrawERC20toL1Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerWithdrawERC20toL1Success)
				if err := _L2PoolManager.contract.UnpackLog(event, "WithdrawERC20toL1Success", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseWithdrawERC20toL1Success(log types.Log) (*L2PoolManagerWithdrawERC20toL1Success, error) {
	event := new(L2PoolManagerWithdrawERC20toL1Success)
	if err := _L2PoolManager.contract.UnpackLog(event, "WithdrawERC20toL1Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerWithdrawETHtoL1SuccessIterator is returned from FilterWithdrawETHtoL1Success and is used to iterate over the raw logs and unpacked data for WithdrawETHtoL1Success events raised by the L2PoolManager contract.
type L2PoolManagerWithdrawETHtoL1SuccessIterator struct {
	Event *L2PoolManagerWithdrawETHtoL1Success // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerWithdrawETHtoL1SuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerWithdrawETHtoL1Success)
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
		it.Event = new(L2PoolManagerWithdrawETHtoL1Success)
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
func (it *L2PoolManagerWithdrawETHtoL1SuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerWithdrawETHtoL1SuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerWithdrawETHtoL1Success represents a WithdrawETHtoL1Success event raised by the L2PoolManager contract.
type L2PoolManagerWithdrawETHtoL1Success struct {
	ChainId   *big.Int
	Timestamp *big.Int
	To        common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETHtoL1Success is a free log retrieval operation binding the contract event 0xf1fa6dbbb3e74f1bef249f22545b238726c1cc79c5571a1416575674c7e63f39.
//
// Solidity: event WithdrawETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) FilterWithdrawETHtoL1Success(opts *bind.FilterOpts) (*L2PoolManagerWithdrawETHtoL1SuccessIterator, error) {

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "WithdrawETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerWithdrawETHtoL1SuccessIterator{contract: _L2PoolManager.contract, event: "WithdrawETHtoL1Success", logs: logs, sub: sub}, nil
}

// WatchWithdrawETHtoL1Success is a free log subscription operation binding the contract event 0xf1fa6dbbb3e74f1bef249f22545b238726c1cc79c5571a1416575674c7e63f39.
//
// Solidity: event WithdrawETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchWithdrawETHtoL1Success(opts *bind.WatchOpts, sink chan<- *L2PoolManagerWithdrawETHtoL1Success) (event.Subscription, error) {

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "WithdrawETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerWithdrawETHtoL1Success)
				if err := _L2PoolManager.contract.UnpackLog(event, "WithdrawETHtoL1Success", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseWithdrawETHtoL1Success(log types.Log) (*L2PoolManagerWithdrawETHtoL1Success, error) {
	event := new(L2PoolManagerWithdrawETHtoL1Success)
	if err := _L2PoolManager.contract.UnpackLog(event, "WithdrawETHtoL1Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerWithdrawWETHtoL1SuccessIterator is returned from FilterWithdrawWETHtoL1Success and is used to iterate over the raw logs and unpacked data for WithdrawWETHtoL1Success events raised by the L2PoolManager contract.
type L2PoolManagerWithdrawWETHtoL1SuccessIterator struct {
	Event *L2PoolManagerWithdrawWETHtoL1Success // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerWithdrawWETHtoL1SuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerWithdrawWETHtoL1Success)
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
		it.Event = new(L2PoolManagerWithdrawWETHtoL1Success)
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
func (it *L2PoolManagerWithdrawWETHtoL1SuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerWithdrawWETHtoL1SuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerWithdrawWETHtoL1Success represents a WithdrawWETHtoL1Success event raised by the L2PoolManager contract.
type L2PoolManagerWithdrawWETHtoL1Success struct {
	ChainId   *big.Int
	Timestamp *big.Int
	To        common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawWETHtoL1Success is a free log retrieval operation binding the contract event 0xcb1a101498aadad4159168d719f955e559e768912e6f5f5c57b094eb047d000f.
//
// Solidity: event WithdrawWETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) FilterWithdrawWETHtoL1Success(opts *bind.FilterOpts) (*L2PoolManagerWithdrawWETHtoL1SuccessIterator, error) {

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "WithdrawWETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerWithdrawWETHtoL1SuccessIterator{contract: _L2PoolManager.contract, event: "WithdrawWETHtoL1Success", logs: logs, sub: sub}, nil
}

// WatchWithdrawWETHtoL1Success is a free log subscription operation binding the contract event 0xcb1a101498aadad4159168d719f955e559e768912e6f5f5c57b094eb047d000f.
//
// Solidity: event WithdrawWETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchWithdrawWETHtoL1Success(opts *bind.WatchOpts, sink chan<- *L2PoolManagerWithdrawWETHtoL1Success) (event.Subscription, error) {

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "WithdrawWETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerWithdrawWETHtoL1Success)
				if err := _L2PoolManager.contract.UnpackLog(event, "WithdrawWETHtoL1Success", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseWithdrawWETHtoL1Success(log types.Log) (*L2PoolManagerWithdrawWETHtoL1Success, error) {
	event := new(L2PoolManagerWithdrawWETHtoL1Success)
	if err := _L2PoolManager.contract.UnpackLog(event, "WithdrawWETHtoL1Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
