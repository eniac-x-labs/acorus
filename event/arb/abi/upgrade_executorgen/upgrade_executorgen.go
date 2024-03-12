// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package upgrade_executorgen

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

// UpgradeExecutorMetaData contains all meta data concerning the UpgradeExecutor contract.
var UpgradeExecutorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TargetCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"upgrade\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"UpgradeExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXECUTOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"upgrade\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"upgradeCallData\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetCallData\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"executors\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600160975561001e610023565b6100e3565b600054610100900460ff161561008f5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156100e1576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b61183c806100f26000396000f3fe6080604052600436106100c75760003560e01c806375b238fc11610074578063a217fddf1161004e578063a217fddf1461026f578063bca8c7b514610284578063d547741f1461029757600080fd5b806375b238fc146101c857806391d14854146101fc578063946d92041461024f57600080fd5b8063248a9ca3116100a5578063248a9ca3146101585780632f2ff15d1461018857806336568abe146101a857600080fd5b806301ffc9a7146100cc57806307bd0265146101015780631cff79cd14610143575b600080fd5b3480156100d857600080fd5b506100ec6100e736600461128f565b6102b7565b60405190151581526020015b60405180910390f35b34801561010d57600080fd5b506101357fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6381565b6040519081526020016100f8565b610156610151366004611378565b610350565b005b34801561016457600080fd5b5061013561017336600461143c565b60009081526065602052604090206001015490565b34801561019457600080fd5b506101566101a3366004611455565b610489565b3480156101b457600080fd5b506101566101c3366004611455565b6104b3565b3480156101d457600080fd5b506101357fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177581565b34801561020857600080fd5b506100ec610217366004611455565b600091825260656020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b34801561025b57600080fd5b5061015661026a366004611481565b610566565b34801561027b57600080fd5b50610135600081565b610156610292366004611378565b610874565b3480156102a357600080fd5b506101566102b2366004611455565b610998565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061034a57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b7fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6361037a816109bd565b6002609754036103eb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064015b60405180910390fd5b600260978190555061042e826040518060600160405280603a81526020016117cd603a913973ffffffffffffffffffffffffffffffffffffffff861691906109ca565b508273ffffffffffffffffffffffffffffffffffffffff167f49f6851d1cd01a518db5bdea5cffbbe90276baa2595f74250b7472b96806302e34846040516104779291906115af565b60405180910390a25050600160975550565b6000828152606560205260409020600101546104a4816109bd565b6104ae8383610af4565b505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610558576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201527f20726f6c657320666f722073656c66000000000000000000000000000000000060648201526084016103e2565b6105628282610be8565b5050565b600054610100900460ff16158080156105865750600054600160ff909116105b806105a05750303b1580156105a0575060005460ff166001145b61062c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016103e2565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561068a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8316610707576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f557067726164654578656375746f723a207a65726f2061646d696e000000000060448201526064016103e2565b61070f610ca3565b6107397fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177580610d3c565b6107837fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e637fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775610d3c565b6107ad7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177584610d87565b60005b825181101561080c576107fc7fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e638483815181106107ef576107ef6115d0565b6020026020010151610d87565b6108058161162e565b90506107b0565b5080156104ae57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b7fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6361089e816109bd565b60026097540361090a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016103e2565b600260978190555061094f823460405180606001604052806031815260200161179c6031913973ffffffffffffffffffffffffffffffffffffffff8716929190610d91565b508273ffffffffffffffffffffffffffffffffffffffff167f4d7dbdcc249630ec373f584267f10abf44938de920c32562f5aee93959c2525834846040516104779291906115af565b6000828152606560205260409020600101546109b3816109bd565b6104ae8383610be8565b6109c78133610f27565b50565b606073ffffffffffffffffffffffffffffffffffffffff84163b610a70576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60448201527f6e7472616374000000000000000000000000000000000000000000000000000060648201526084016103e2565b6000808573ffffffffffffffffffffffffffffffffffffffff1685604051610a989190611666565b600060405180830381855af49150503d8060008114610ad3576040519150601f19603f3d011682016040523d82523d6000602084013e610ad8565b606091505b5091509150610ae8828286610ff9565b925050505b9392505050565b600082815260656020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661056257600082815260656020908152604080832073ffffffffffffffffffffffffffffffffffffffff85168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610b8a3390565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b600082815260656020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff161561056257600082815260656020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b600054610100900460ff16610d3a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016103e2565b565b600082815260656020526040808220600101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b6105628282610af4565b606082471015610e23576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016103e2565b73ffffffffffffffffffffffffffffffffffffffff85163b610ea1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103e2565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610eca9190611666565b60006040518083038185875af1925050503d8060008114610f07576040519150601f19603f3d011682016040523d82523d6000602084013e610f0c565b606091505b5091509150610f1c828286610ff9565b979650505050505050565b600082815260656020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661056257610f7f8173ffffffffffffffffffffffffffffffffffffffff16601461104c565b610f8a83602061104c565b604051602001610f9b929190611682565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526103e291600401611703565b60608315611008575081610aed565b8251156110185782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e29190611703565b6060600061105b836002611716565b611066906002611753565b67ffffffffffffffff81111561107e5761107e6112fa565b6040519080825280601f01601f1916602001820160405280156110a8576020820181803683370190505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106110df576110df6115d0565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110611142576111426115d0565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600061117e846002611716565b611189906001611753565b90505b6001811115611226577f303132333435363738396162636465660000000000000000000000000000000085600f16601081106111ca576111ca6115d0565b1a60f81b8282815181106111e0576111e06115d0565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c9361121f81611766565b905061118c565b508315610aed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016103e2565b6000602082840312156112a157600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610aed57600080fd5b803573ffffffffffffffffffffffffffffffffffffffff811681146112f557600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611370576113706112fa565b604052919050565b6000806040838503121561138b57600080fd5b611394836112d1565b915060208084013567ffffffffffffffff808211156113b257600080fd5b818601915086601f8301126113c657600080fd5b8135818111156113d8576113d86112fa565b611408847fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611329565b9150808252878482850101111561141e57600080fd5b80848401858401376000848284010152508093505050509250929050565b60006020828403121561144e57600080fd5b5035919050565b6000806040838503121561146857600080fd5b82359150611478602084016112d1565b90509250929050565b6000806040838503121561149457600080fd5b61149d836112d1565b915060208084013567ffffffffffffffff808211156114bb57600080fd5b818601915086601f8301126114cf57600080fd5b8135818111156114e1576114e16112fa565b8060051b91506114f2848301611329565b818152918301840191848101908984111561150c57600080fd5b938501935b8385101561153157611522856112d1565b82529385019390850190611511565b8096505050505050509250929050565b60005b8381101561155c578181015183820152602001611544565b50506000910152565b6000815180845261157d816020860160208601611541565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b8281526040602082015260006115c86040830184611565565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361165f5761165f6115ff565b5060010190565b60008251611678818460208701611541565b9190910192915050565b7f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008152600083516116ba816017850160208801611541565b7f206973206d697373696e6720726f6c652000000000000000000000000000000060179184019182015283516116f7816028840160208801611541565b01602801949350505050565b602081526000610aed6020830184611565565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561174e5761174e6115ff565b500290565b8082018082111561034a5761034a6115ff565b600081611775576117756115ff565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019056fe557067726164654578656375746f723a20696e6e65722063616c6c206661696c656420776974686f757420726561736f6e557067726164654578656375746f723a20696e6e65722064656c65676174652063616c6c206661696c656420776974686f757420726561736f6ea26469706673582212202dd5382fdba08a29cf28575cf947a054897172a386b7932a48f84af2cd16c35c64736f6c63430008100033",
}

// UpgradeExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use UpgradeExecutorMetaData.ABI instead.
var UpgradeExecutorABI = UpgradeExecutorMetaData.ABI

// UpgradeExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UpgradeExecutorMetaData.Bin instead.
var UpgradeExecutorBin = UpgradeExecutorMetaData.Bin

// DeployUpgradeExecutor deploys a new Ethereum contract, binding an instance of UpgradeExecutor to it.
func DeployUpgradeExecutor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UpgradeExecutor, error) {
	parsed, err := UpgradeExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UpgradeExecutorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UpgradeExecutor{UpgradeExecutorCaller: UpgradeExecutorCaller{contract: contract}, UpgradeExecutorTransactor: UpgradeExecutorTransactor{contract: contract}, UpgradeExecutorFilterer: UpgradeExecutorFilterer{contract: contract}}, nil
}

// UpgradeExecutor is an auto generated Go binding around an Ethereum contract.
type UpgradeExecutor struct {
	UpgradeExecutorCaller     // Read-only binding to the contract
	UpgradeExecutorTransactor // Write-only binding to the contract
	UpgradeExecutorFilterer   // Log filterer for contract events
}

// UpgradeExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeExecutorSession struct {
	Contract     *UpgradeExecutor  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradeExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeExecutorCallerSession struct {
	Contract *UpgradeExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UpgradeExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeExecutorTransactorSession struct {
	Contract     *UpgradeExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UpgradeExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeExecutorRaw struct {
	Contract *UpgradeExecutor // Generic contract binding to access the raw methods on
}

// UpgradeExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeExecutorCallerRaw struct {
	Contract *UpgradeExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeExecutorTransactorRaw struct {
	Contract *UpgradeExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeExecutor creates a new instance of UpgradeExecutor, bound to a specific deployed contract.
func NewUpgradeExecutor(address common.Address, backend bind.ContractBackend) (*UpgradeExecutor, error) {
	contract, err := bindUpgradeExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutor{UpgradeExecutorCaller: UpgradeExecutorCaller{contract: contract}, UpgradeExecutorTransactor: UpgradeExecutorTransactor{contract: contract}, UpgradeExecutorFilterer: UpgradeExecutorFilterer{contract: contract}}, nil
}

// NewUpgradeExecutorCaller creates a new read-only instance of UpgradeExecutor, bound to a specific deployed contract.
func NewUpgradeExecutorCaller(address common.Address, caller bind.ContractCaller) (*UpgradeExecutorCaller, error) {
	contract, err := bindUpgradeExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorCaller{contract: contract}, nil
}

// NewUpgradeExecutorTransactor creates a new write-only instance of UpgradeExecutor, bound to a specific deployed contract.
func NewUpgradeExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeExecutorTransactor, error) {
	contract, err := bindUpgradeExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorTransactor{contract: contract}, nil
}

// NewUpgradeExecutorFilterer creates a new log filterer instance of UpgradeExecutor, bound to a specific deployed contract.
func NewUpgradeExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeExecutorFilterer, error) {
	contract, err := bindUpgradeExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorFilterer{contract: contract}, nil
}

// bindUpgradeExecutor binds a generic wrapper to an already deployed contract.
func bindUpgradeExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UpgradeExecutorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeExecutor *UpgradeExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeExecutor.Contract.UpgradeExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeExecutor *UpgradeExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.UpgradeExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeExecutor *UpgradeExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.UpgradeExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeExecutor *UpgradeExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeExecutor *UpgradeExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeExecutor *UpgradeExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeExecutor.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorSession) ADMINROLE() ([32]byte, error) {
	return _UpgradeExecutor.Contract.ADMINROLE(&_UpgradeExecutor.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCallerSession) ADMINROLE() ([32]byte, error) {
	return _UpgradeExecutor.Contract.ADMINROLE(&_UpgradeExecutor.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeExecutor.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UpgradeExecutor.Contract.DEFAULTADMINROLE(&_UpgradeExecutor.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UpgradeExecutor.Contract.DEFAULTADMINROLE(&_UpgradeExecutor.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCaller) EXECUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeExecutor.contract.Call(opts, &out, "EXECUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorSession) EXECUTORROLE() ([32]byte, error) {
	return _UpgradeExecutor.Contract.EXECUTORROLE(&_UpgradeExecutor.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCallerSession) EXECUTORROLE() ([32]byte, error) {
	return _UpgradeExecutor.Contract.EXECUTORROLE(&_UpgradeExecutor.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeExecutor.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UpgradeExecutor.Contract.GetRoleAdmin(&_UpgradeExecutor.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UpgradeExecutor.Contract.GetRoleAdmin(&_UpgradeExecutor.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UpgradeExecutor *UpgradeExecutorCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _UpgradeExecutor.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UpgradeExecutor *UpgradeExecutorSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UpgradeExecutor.Contract.HasRole(&_UpgradeExecutor.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UpgradeExecutor *UpgradeExecutorCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UpgradeExecutor.Contract.HasRole(&_UpgradeExecutor.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UpgradeExecutor *UpgradeExecutorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _UpgradeExecutor.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UpgradeExecutor *UpgradeExecutorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UpgradeExecutor.Contract.SupportsInterface(&_UpgradeExecutor.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UpgradeExecutor *UpgradeExecutorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UpgradeExecutor.Contract.SupportsInterface(&_UpgradeExecutor.CallOpts, interfaceId)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address upgrade, bytes upgradeCallData) payable returns()
func (_UpgradeExecutor *UpgradeExecutorTransactor) Execute(opts *bind.TransactOpts, upgrade common.Address, upgradeCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutor.contract.Transact(opts, "execute", upgrade, upgradeCallData)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address upgrade, bytes upgradeCallData) payable returns()
func (_UpgradeExecutor *UpgradeExecutorSession) Execute(upgrade common.Address, upgradeCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.Execute(&_UpgradeExecutor.TransactOpts, upgrade, upgradeCallData)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address upgrade, bytes upgradeCallData) payable returns()
func (_UpgradeExecutor *UpgradeExecutorTransactorSession) Execute(upgrade common.Address, upgradeCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.Execute(&_UpgradeExecutor.TransactOpts, upgrade, upgradeCallData)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0xbca8c7b5.
//
// Solidity: function executeCall(address target, bytes targetCallData) payable returns()
func (_UpgradeExecutor *UpgradeExecutorTransactor) ExecuteCall(opts *bind.TransactOpts, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutor.contract.Transact(opts, "executeCall", target, targetCallData)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0xbca8c7b5.
//
// Solidity: function executeCall(address target, bytes targetCallData) payable returns()
func (_UpgradeExecutor *UpgradeExecutorSession) ExecuteCall(target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.ExecuteCall(&_UpgradeExecutor.TransactOpts, target, targetCallData)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0xbca8c7b5.
//
// Solidity: function executeCall(address target, bytes targetCallData) payable returns()
func (_UpgradeExecutor *UpgradeExecutorTransactorSession) ExecuteCall(target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.ExecuteCall(&_UpgradeExecutor.TransactOpts, target, targetCallData)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.GrantRole(&_UpgradeExecutor.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.GrantRole(&_UpgradeExecutor.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address admin, address[] executors) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, executors []common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.contract.Transact(opts, "initialize", admin, executors)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address admin, address[] executors) returns()
func (_UpgradeExecutor *UpgradeExecutorSession) Initialize(admin common.Address, executors []common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.Initialize(&_UpgradeExecutor.TransactOpts, admin, executors)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address admin, address[] executors) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactorSession) Initialize(admin common.Address, executors []common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.Initialize(&_UpgradeExecutor.TransactOpts, admin, executors)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.RenounceRole(&_UpgradeExecutor.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.RenounceRole(&_UpgradeExecutor.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.RevokeRole(&_UpgradeExecutor.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.RevokeRole(&_UpgradeExecutor.TransactOpts, role, account)
}

// UpgradeExecutorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UpgradeExecutor contract.
type UpgradeExecutorInitializedIterator struct {
	Event *UpgradeExecutorInitialized // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorInitialized)
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
		it.Event = new(UpgradeExecutorInitialized)
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
func (it *UpgradeExecutorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorInitialized represents a Initialized event raised by the UpgradeExecutor contract.
type UpgradeExecutorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UpgradeExecutor *UpgradeExecutorFilterer) FilterInitialized(opts *bind.FilterOpts) (*UpgradeExecutorInitializedIterator, error) {

	logs, sub, err := _UpgradeExecutor.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorInitializedIterator{contract: _UpgradeExecutor.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UpgradeExecutor *UpgradeExecutorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorInitialized) (event.Subscription, error) {

	logs, sub, err := _UpgradeExecutor.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorInitialized)
				if err := _UpgradeExecutor.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UpgradeExecutor *UpgradeExecutorFilterer) ParseInitialized(log types.Log) (*UpgradeExecutorInitialized, error) {
	event := new(UpgradeExecutorInitialized)
	if err := _UpgradeExecutor.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the UpgradeExecutor contract.
type UpgradeExecutorRoleAdminChangedIterator struct {
	Event *UpgradeExecutorRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorRoleAdminChanged)
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
		it.Event = new(UpgradeExecutorRoleAdminChanged)
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
func (it *UpgradeExecutorRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorRoleAdminChanged represents a RoleAdminChanged event raised by the UpgradeExecutor contract.
type UpgradeExecutorRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UpgradeExecutor *UpgradeExecutorFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*UpgradeExecutorRoleAdminChangedIterator, error) {

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

	logs, sub, err := _UpgradeExecutor.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorRoleAdminChangedIterator{contract: _UpgradeExecutor.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UpgradeExecutor *UpgradeExecutorFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _UpgradeExecutor.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorRoleAdminChanged)
				if err := _UpgradeExecutor.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_UpgradeExecutor *UpgradeExecutorFilterer) ParseRoleAdminChanged(log types.Log) (*UpgradeExecutorRoleAdminChanged, error) {
	event := new(UpgradeExecutorRoleAdminChanged)
	if err := _UpgradeExecutor.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the UpgradeExecutor contract.
type UpgradeExecutorRoleGrantedIterator struct {
	Event *UpgradeExecutorRoleGranted // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorRoleGranted)
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
		it.Event = new(UpgradeExecutorRoleGranted)
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
func (it *UpgradeExecutorRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorRoleGranted represents a RoleGranted event raised by the UpgradeExecutor contract.
type UpgradeExecutorRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutor *UpgradeExecutorFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UpgradeExecutorRoleGrantedIterator, error) {

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

	logs, sub, err := _UpgradeExecutor.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorRoleGrantedIterator{contract: _UpgradeExecutor.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutor *UpgradeExecutorFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _UpgradeExecutor.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorRoleGranted)
				if err := _UpgradeExecutor.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_UpgradeExecutor *UpgradeExecutorFilterer) ParseRoleGranted(log types.Log) (*UpgradeExecutorRoleGranted, error) {
	event := new(UpgradeExecutorRoleGranted)
	if err := _UpgradeExecutor.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the UpgradeExecutor contract.
type UpgradeExecutorRoleRevokedIterator struct {
	Event *UpgradeExecutorRoleRevoked // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorRoleRevoked)
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
		it.Event = new(UpgradeExecutorRoleRevoked)
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
func (it *UpgradeExecutorRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorRoleRevoked represents a RoleRevoked event raised by the UpgradeExecutor contract.
type UpgradeExecutorRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutor *UpgradeExecutorFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UpgradeExecutorRoleRevokedIterator, error) {

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

	logs, sub, err := _UpgradeExecutor.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorRoleRevokedIterator{contract: _UpgradeExecutor.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutor *UpgradeExecutorFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _UpgradeExecutor.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorRoleRevoked)
				if err := _UpgradeExecutor.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_UpgradeExecutor *UpgradeExecutorFilterer) ParseRoleRevoked(log types.Log) (*UpgradeExecutorRoleRevoked, error) {
	event := new(UpgradeExecutorRoleRevoked)
	if err := _UpgradeExecutor.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorTargetCallExecutedIterator is returned from FilterTargetCallExecuted and is used to iterate over the raw logs and unpacked data for TargetCallExecuted events raised by the UpgradeExecutor contract.
type UpgradeExecutorTargetCallExecutedIterator struct {
	Event *UpgradeExecutorTargetCallExecuted // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorTargetCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorTargetCallExecuted)
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
		it.Event = new(UpgradeExecutorTargetCallExecuted)
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
func (it *UpgradeExecutorTargetCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorTargetCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorTargetCallExecuted represents a TargetCallExecuted event raised by the UpgradeExecutor contract.
type UpgradeExecutorTargetCallExecuted struct {
	Target common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTargetCallExecuted is a free log retrieval operation binding the contract event 0x4d7dbdcc249630ec373f584267f10abf44938de920c32562f5aee93959c25258.
//
// Solidity: event TargetCallExecuted(address indexed target, uint256 value, bytes data)
func (_UpgradeExecutor *UpgradeExecutorFilterer) FilterTargetCallExecuted(opts *bind.FilterOpts, target []common.Address) (*UpgradeExecutorTargetCallExecutedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _UpgradeExecutor.contract.FilterLogs(opts, "TargetCallExecuted", targetRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorTargetCallExecutedIterator{contract: _UpgradeExecutor.contract, event: "TargetCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTargetCallExecuted is a free log subscription operation binding the contract event 0x4d7dbdcc249630ec373f584267f10abf44938de920c32562f5aee93959c25258.
//
// Solidity: event TargetCallExecuted(address indexed target, uint256 value, bytes data)
func (_UpgradeExecutor *UpgradeExecutorFilterer) WatchTargetCallExecuted(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorTargetCallExecuted, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _UpgradeExecutor.contract.WatchLogs(opts, "TargetCallExecuted", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorTargetCallExecuted)
				if err := _UpgradeExecutor.contract.UnpackLog(event, "TargetCallExecuted", log); err != nil {
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

// ParseTargetCallExecuted is a log parse operation binding the contract event 0x4d7dbdcc249630ec373f584267f10abf44938de920c32562f5aee93959c25258.
//
// Solidity: event TargetCallExecuted(address indexed target, uint256 value, bytes data)
func (_UpgradeExecutor *UpgradeExecutorFilterer) ParseTargetCallExecuted(log types.Log) (*UpgradeExecutorTargetCallExecuted, error) {
	event := new(UpgradeExecutorTargetCallExecuted)
	if err := _UpgradeExecutor.contract.UnpackLog(event, "TargetCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorUpgradeExecutedIterator is returned from FilterUpgradeExecuted and is used to iterate over the raw logs and unpacked data for UpgradeExecuted events raised by the UpgradeExecutor contract.
type UpgradeExecutorUpgradeExecutedIterator struct {
	Event *UpgradeExecutorUpgradeExecuted // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorUpgradeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorUpgradeExecuted)
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
		it.Event = new(UpgradeExecutorUpgradeExecuted)
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
func (it *UpgradeExecutorUpgradeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorUpgradeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorUpgradeExecuted represents a UpgradeExecuted event raised by the UpgradeExecutor contract.
type UpgradeExecutorUpgradeExecuted struct {
	Upgrade common.Address
	Value   *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpgradeExecuted is a free log retrieval operation binding the contract event 0x49f6851d1cd01a518db5bdea5cffbbe90276baa2595f74250b7472b96806302e.
//
// Solidity: event UpgradeExecuted(address indexed upgrade, uint256 value, bytes data)
func (_UpgradeExecutor *UpgradeExecutorFilterer) FilterUpgradeExecuted(opts *bind.FilterOpts, upgrade []common.Address) (*UpgradeExecutorUpgradeExecutedIterator, error) {

	var upgradeRule []interface{}
	for _, upgradeItem := range upgrade {
		upgradeRule = append(upgradeRule, upgradeItem)
	}

	logs, sub, err := _UpgradeExecutor.contract.FilterLogs(opts, "UpgradeExecuted", upgradeRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorUpgradeExecutedIterator{contract: _UpgradeExecutor.contract, event: "UpgradeExecuted", logs: logs, sub: sub}, nil
}

// WatchUpgradeExecuted is a free log subscription operation binding the contract event 0x49f6851d1cd01a518db5bdea5cffbbe90276baa2595f74250b7472b96806302e.
//
// Solidity: event UpgradeExecuted(address indexed upgrade, uint256 value, bytes data)
func (_UpgradeExecutor *UpgradeExecutorFilterer) WatchUpgradeExecuted(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorUpgradeExecuted, upgrade []common.Address) (event.Subscription, error) {

	var upgradeRule []interface{}
	for _, upgradeItem := range upgrade {
		upgradeRule = append(upgradeRule, upgradeItem)
	}

	logs, sub, err := _UpgradeExecutor.contract.WatchLogs(opts, "UpgradeExecuted", upgradeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorUpgradeExecuted)
				if err := _UpgradeExecutor.contract.UnpackLog(event, "UpgradeExecuted", log); err != nil {
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

// ParseUpgradeExecuted is a log parse operation binding the contract event 0x49f6851d1cd01a518db5bdea5cffbbe90276baa2595f74250b7472b96806302e.
//
// Solidity: event UpgradeExecuted(address indexed upgrade, uint256 value, bytes data)
func (_UpgradeExecutor *UpgradeExecutorFilterer) ParseUpgradeExecuted(log types.Log) (*UpgradeExecutorUpgradeExecuted, error) {
	event := new(UpgradeExecutorUpgradeExecuted)
	if err := _UpgradeExecutor.contract.UnpackLog(event, "UpgradeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
