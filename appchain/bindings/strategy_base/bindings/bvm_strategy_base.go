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

// StrategyBaseMetaData contains all meta data concerning the StrategyBase contract.
var StrategyBaseMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ETHBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WETHBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"weth\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"newShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_stakingWeth\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_pauser\",\"type\":\"address\",\"internalType\":\"contractIL2Pauser\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauser\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIL2Pauser\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"relayer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToStaking\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToStakingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakingToShares\",\"inputs\":[{\"name\":\"amountStaking\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakingToSharesView\",\"inputs\":[{\"name\":\"amountStaking\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakingWeth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferETHToL2DappLinkBridge\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1StakingManagerAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"batchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transferWETHToL2DappLinkBridge\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1StakingManagerAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wethAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"batchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"userStaking\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userStakingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weth\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferETHToL2DappLinkBridge\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bridge\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"l1StakingManagerAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"bridgeEthAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"batchId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608080604052346100b8577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166100a957506001600160401b036002600160401b031982821601610064575b60405161141f90816100be8239f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1388080610055565b63f92ee8a960e01b8152600490fd5b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806310ed193a1461015d57806314d1441f146101585780633055a78c1461015357806339b70e381461014e5780633a98ef391461014957806347e7ef2414610144578063548b95601461013f57806356a7117b14610130578063821460f51461010d5780638406c0791461013a578063893d4d9e146101355780639910a665146101305780639d9cc41a1461012b5780639fd0506d14610126578063ab5921e114610121578063ce7c2ac21461011c578063d69c3d3014610117578063d9caed1214610112578063f5f1b29c1461010d578063f8c8765e146101085763fc1d98fb1461010357600080fd5b610888565b61072c565b610422565b610638565b61061a565b6105f8565b610527565b6104fe565b61047a565b610404565b61015d565b610451565b6103b2565b610296565b610278565b61024f565b610233565b6101cb565b346101b55760203660031901126101b5576003546103e8908181018091116101b05761018761102f565b9182018092116101b0576020916101a36101a892600435610901565b610914565b604051908152f35b6108b1565b600080fd5b6001600160a01b038116036101b557565b60e03660031901126101b55760206102296044356101e8816101ba565b606435906101f5826101ba565b608435610201816101ba565b61021660018060a01b03600254163314610934565b60c4359260a43592602435600435610a1f565b6040519015158152f35b346101b55760003660031901126101b557602047604051908152f35b346101b55760003660031901126101b5576000546040516001600160a01b039091168152602090f35b346101b55760003660031901126101b5576020600354604051908152f35b346101b55760403660031901126101b557600480356102b4816101ba565b6024359060206102e260018060a01b036102d381600054163314610c22565b8554166001600160a01b031690565b604051631c187a5960e11b815294859182905afa9081156103ad5761031661031b9261037a9560009161037e575b50610c85565b6110ea565b61036a610365610353600354936101a3610334866108c7565b61034d8361034861034361102f565b6108c7565b610cd1565b92610901565b8093610360821515610cde565b6108d6565b600355565b6040519081529081906020820190565b0390f35b6103a0915060203d6020116103a6575b61039881836109d3565b810190610c6d565b38610310565b503d61038e565b610a04565b60c03660031901126101b55760206102296044356103cf816101ba565b606435906103dc826101ba565b6103f160018060a01b03600254163314610934565b60a4359160843591602435600435610d41565b346101b55760203660031901126101b55760206101a8600435610e34565b346101b55760203660031901126101b55760206101a861044c600435610447816101ba565b610e82565b610e34565b346101b55760003660031901126101b5576002546040516001600160a01b039091168152602090f35b346101b55760003660031901126101b5576001546040516370a0823160e01b815230600482015290602090829060249082906001600160a01b03165afa80156103ad576020916000916104d1575b50604051908152f35b6104f19150823d84116104f7575b6104e981836109d3565b8101906109f5565b386104c8565b503d6104df565b346101b55760003660031901126101b5576004546040516001600160a01b039091168152602090f35b346101b55760003660031901126101b557604080519061054682610996565b604d825260207f4261736520537472617465677920696d706c656d656e746174696f6e20746f2060208401527f696e68657269742066726f6d20666f72206d6f726520636f6d706c657820696d828401526c706c656d656e746174696f6e7360981b606084015281519283916020835281519182602085015260005b8381106105e2575050600083830185015250601f01601f19168101030190f35b81810183015187820187015286945082016105c2565b346101b55760203660031901126101b55760206101a8600435610447816101ba565b346101b55760003660031901126101b5576020600554604051908152f35b346101b55760603660031901126101b55760048035610656816101ba565b602435610662816101ba565b60443591602061069060018060a01b0361068181600054163314610c22565b8654166001600160a01b031690565b604051632602edb160e01b815295869182905afa9384156103ad5761070b946106c19160009161070d575b50610ee1565b6106ca82611195565b610706610365600354946106e086821115610f2d565b6107006106ec876108c7565b6101a3836106fb61034361102f565b610901565b95610cd1565b61124b565b005b610726915060203d6020116103a65761039881836109d3565b386106bb565b346101b55760803660031901126101b557600435610749816101ba565b60243590610756826101ba565b604435610762816101ba565b6064359061076f826101ba565b6000805160206113ca833981519152549367ffffffffffffffff60ff8660401c1615951680159081610880575b6001149081610876575b15908161086d575b5061085b576000805160206113ca833981519152805467ffffffffffffffff191660011790556107e2938561083157610fb5565b6107e857005b6000805160206113ca833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b6000805160206113ca833981519152805460ff60401b191668010000000000000000179055610fb5565b60405163f92ee8a960e01b8152600490fd5b905015386107ae565b303b1591506107a6565b86915061079c565b346101b55760003660031901126101b5576001546040516001600160a01b039091168152602090f35b634e487b7160e01b600052601160045260246000fd5b906103e882018092116101b057565b919082018092116101b057565b906801bc16d674ec800000918281029281840414901517156101b057565b818102929181159184041417156101b057565b811561091e570490565b634e487b7160e01b600052601260045260246000fd5b1561093b57565b60405162461bcd60e51b815260206004820152601860248201527f5374726174656779426173652e6f6e6c7952656c6179657200000000000000006044820152606490fd5b634e487b7160e01b600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176109b257604052565b610980565b60c0810190811067ffffffffffffffff8211176109b257604052565b90601f8019910116810190811067ffffffffffffffff8211176109b257604052565b908160209103126101b5575190565b6040513d6000823e3d90fd5b60001981146101b05760010190565b929593919093610a45610a3960015460018060a01b031690565b6001600160a01b031690565b6040516370a0823160e01b8082523060048301529194916020919082826024818a5afa9182156103ad57600092610c03575b506801bc16d674ec80000080921015610a9a575050505050505050505050600090565b604051908152306004820152958290879060249082905afa80156103ad577f1e41494416713921a8fa17cffca7814892d9c5c3be42174a8a5d95114635bce899610b72610b7894610b648c8f958d97610b0091610be09e600091610be6575b50046108e3565b80998197610b17610b12600554610a10565b600555565b600554936040519a8b97635d4bf3b760e11b90890152602488019290969594919360a09460c085019885526020850152600180861b03809216604085015216606083015260808201520152565b03601f1981018552846109d3565b86611097565b600154909890610b90906001600160a01b0316610a39565b906005549460405198899889949098979693919260e0969361010087019a8752602087015260018060a01b039283809216604088015216606086015216608084015260a083015260c08201520152565b0390a190565b610bfd9150863d88116104f7576104e981836109d3565b38610af9565b610c1b919250833d85116104f7576104e981836109d3565b9038610a77565b15610c2957565b606460405162461bcd60e51b815260206004820152602060248201527f5374726174656779426173652e6f6e6c7953747261746567794d616e616765726044820152fd5b908160209103126101b5575180151581036101b55790565b15610c8c57565b60405162461bcd60e51b815260206004820152601b60248201527f5374726174656779426173653a6465706f7369742070617573656400000000006044820152606490fd5b919082039182116101b057565b15610ce557565b60405162461bcd60e51b815260206004820152602e60248201527f5374726174656779426173652e6465706f7369743a206e65775368617265732060448201526d63616e6e6f74206265207a65726f60901b6064820152608490fd5b91946801bc16d674ec80000093844711610d615750505050505050600090565b7f1e41494416713921a8fa17cffca7814892d9c5c3be42174a8a5d95114635bce895610d916101009647046108e3565b91610dec610da0600554610a10565b92836005558460405191633b90b1f960e01b602084015289602484015288604484015260018060a01b0380951695866064850152608484015260848352610de6836109b7565b8c611097565b98600554956040519788526020880152166040860152606085015273eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee608085015260a084015260c083015260e0820152a190565b6003546103e88082018092116101b057610e4c61102f565b9081018091116101b057610e63926101a391610901565b90565b67ffffffffffffffff81116109b257601f01601f191660200190565b600054604051633d3f06c960e11b81526001600160a01b0392831660048201523060248201529160209183916044918391165afa9081156103ad57600091610ec8575090565b610e63915060203d6020116104f7576104e981836109d3565b15610ee857565b60405162461bcd60e51b815260206004820152601c60248201527f5374726174656779426173653a776974686472617720706175736564000000006044820152606490fd5b15610f3457565b60405162461bcd60e51b815260206004820152604d60248201527f5374726174656779426173652e77697468647261773a20616d6f756e7453686160448201527f726573206d757374206265206c657373207468616e206f7220657175616c207460648201526c6f20746f74616c53686172657360981b608482015260a490fd5b9260ff6000805160206113ca8339815191525460401c161561101d5760018060a01b0392838092816bffffffffffffffffffffffff60a01b97168760015416176001551685600454161760045516836000541617600055169060025416176002556001600555565b604051631afcd79f60e31b8152600490fd5b6001546040516370a0823160e01b815230600482015290602090829060249082906001600160a01b03165afa9081156103ad57600091611078575b504781018091116101b05790565b611091915060203d6020116104f7576104e981836109d3565b3861106a565b929060061b622673c001603f5a02106110bb576000928392602083519301915af190565b6308c379a06000526020805278185361666543616c6c3a204e6f7420656e6f756768206761736058526064601cfd5b6001546001600160a01b0391821691168114908115611177575b501561110c57565b60405162461bcd60e51b815260206004820152603a60248201527f5374726174656779426173652e6465706f7369743a2043616e206f6e6c79206460448201527f65706f736974207374616b696e675765746820616e64206574680000000000006064820152608490fd5b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee91501438611104565b6001546001600160a01b039182169116811490811561122d575b50156111b757565b60405162461bcd60e51b815260206004820152604260248201527f5374726174656779426173652e77697468647261773a2043616e206f6e6c792060448201527f776974686472617720746865207374726174656779207765746820616e6420656064820152610e8d60f31b608482015260a490fd5b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee915014386111af565b906001600160a01b0390811673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee8103611298575060008093819382938391831561128e575b1690f1156103ad57565b6108fc9250611284565b9261130392600092839260405191602083019363a9059cbb60e01b85521660248301526044820152604481526112cd81610996565b519082865af13d1561135e573d906112e482610e66565b916112f260405193846109d3565b82523d6000602084013e5b83611366565b805190811515918261133c575b50506113195750565b604051635274afe760e01b81526001600160a01b03919091166004820152602490fd5b611357925090602080611353938301019101610c6d565b1590565b3880611310565b6060906112fd565b9061138d575080511561137b57805190602001fd5b604051630a12f52160e11b8152600490fd5b815115806113c0575b61139e575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561139656fef0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220a2fc885d26b05d24766497244739ffff20c9b950ed4f68d0542a05708c64b07764736f6c63430008180033",
}

// StrategyBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyBaseMetaData.ABI instead.
var StrategyBaseABI = StrategyBaseMetaData.ABI

// StrategyBaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StrategyBaseMetaData.Bin instead.
var StrategyBaseBin = StrategyBaseMetaData.Bin

// DeployStrategyBase deploys a new Ethereum contract, binding an instance of StrategyBase to it.
func DeployStrategyBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StrategyBase, error) {
	parsed, err := StrategyBaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StrategyBaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StrategyBase{StrategyBaseCaller: StrategyBaseCaller{contract: contract}, StrategyBaseTransactor: StrategyBaseTransactor{contract: contract}, StrategyBaseFilterer: StrategyBaseFilterer{contract: contract}}, nil
}

// StrategyBase is an auto generated Go binding around an Ethereum contract.
type StrategyBase struct {
	StrategyBaseCaller     // Read-only binding to the contract
	StrategyBaseTransactor // Write-only binding to the contract
	StrategyBaseFilterer   // Log filterer for contract events
}

// StrategyBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategyBaseSession struct {
	Contract     *StrategyBase     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrategyBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyBaseCallerSession struct {
	Contract *StrategyBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StrategyBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyBaseTransactorSession struct {
	Contract     *StrategyBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StrategyBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyBaseRaw struct {
	Contract *StrategyBase // Generic contract binding to access the raw methods on
}

// StrategyBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyBaseCallerRaw struct {
	Contract *StrategyBaseCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyBaseTransactorRaw struct {
	Contract *StrategyBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategyBase creates a new instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBase(address common.Address, backend bind.ContractBackend) (*StrategyBase, error) {
	contract, err := bindStrategyBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrategyBase{StrategyBaseCaller: StrategyBaseCaller{contract: contract}, StrategyBaseTransactor: StrategyBaseTransactor{contract: contract}, StrategyBaseFilterer: StrategyBaseFilterer{contract: contract}}, nil
}

// NewStrategyBaseCaller creates a new read-only instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBaseCaller(address common.Address, caller bind.ContractCaller) (*StrategyBaseCaller, error) {
	contract, err := bindStrategyBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseCaller{contract: contract}, nil
}

// NewStrategyBaseTransactor creates a new write-only instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyBaseTransactor, error) {
	contract, err := bindStrategyBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTransactor{contract: contract}, nil
}

// NewStrategyBaseFilterer creates a new log filterer instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyBaseFilterer, error) {
	contract, err := bindStrategyBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseFilterer{contract: contract}, nil
}

// bindStrategyBase binds a generic wrapper to an already deployed contract.
func bindStrategyBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrategyBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyBase *StrategyBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyBase.Contract.StrategyBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyBase *StrategyBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBase.Contract.StrategyBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyBase *StrategyBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyBase.Contract.StrategyBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyBase *StrategyBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyBase *StrategyBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyBase *StrategyBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyBase.Contract.contract.Transact(opts, method, params...)
}

// ETHBalance is a free data retrieval call binding the contract method 0x3055a78c.
//
// Solidity: function ETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) ETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "ETHBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ETHBalance is a free data retrieval call binding the contract method 0x3055a78c.
//
// Solidity: function ETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) ETHBalance() (*big.Int, error) {
	return _StrategyBase.Contract.ETHBalance(&_StrategyBase.CallOpts)
}

// ETHBalance is a free data retrieval call binding the contract method 0x3055a78c.
//
// Solidity: function ETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) ETHBalance() (*big.Int, error) {
	return _StrategyBase.Contract.ETHBalance(&_StrategyBase.CallOpts)
}

// WETHBalance is a free data retrieval call binding the contract method 0x9d9cc41a.
//
// Solidity: function WETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) WETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "WETHBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WETHBalance is a free data retrieval call binding the contract method 0x9d9cc41a.
//
// Solidity: function WETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) WETHBalance() (*big.Int, error) {
	return _StrategyBase.Contract.WETHBalance(&_StrategyBase.CallOpts)
}

// WETHBalance is a free data retrieval call binding the contract method 0x9d9cc41a.
//
// Solidity: function WETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) WETHBalance() (*big.Int, error) {
	return _StrategyBase.Contract.WETHBalance(&_StrategyBase.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_StrategyBase *StrategyBaseCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_StrategyBase *StrategyBaseSession) Explanation() (string, error) {
	return _StrategyBase.Contract.Explanation(&_StrategyBase.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_StrategyBase *StrategyBaseCallerSession) Explanation() (string, error) {
	return _StrategyBase.Contract.Explanation(&_StrategyBase.CallOpts)
}

// NextNonce is a free data retrieval call binding the contract method 0xd69c3d30.
//
// Solidity: function nextNonce() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) NextNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "nextNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextNonce is a free data retrieval call binding the contract method 0xd69c3d30.
//
// Solidity: function nextNonce() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) NextNonce() (*big.Int, error) {
	return _StrategyBase.Contract.NextNonce(&_StrategyBase.CallOpts)
}

// NextNonce is a free data retrieval call binding the contract method 0xd69c3d30.
//
// Solidity: function nextNonce() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) NextNonce() (*big.Int, error) {
	return _StrategyBase.Contract.NextNonce(&_StrategyBase.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_StrategyBase *StrategyBaseCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_StrategyBase *StrategyBaseSession) Pauser() (common.Address, error) {
	return _StrategyBase.Contract.Pauser(&_StrategyBase.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) Pauser() (common.Address, error) {
	return _StrategyBase.Contract.Pauser(&_StrategyBase.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_StrategyBase *StrategyBaseCaller) Relayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "relayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_StrategyBase *StrategyBaseSession) Relayer() (common.Address, error) {
	return _StrategyBase.Contract.Relayer(&_StrategyBase.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) Relayer() (common.Address, error) {
	return _StrategyBase.Contract.Relayer(&_StrategyBase.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) Shares(user common.Address) (*big.Int, error) {
	return _StrategyBase.Contract.Shares(&_StrategyBase.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _StrategyBase.Contract.Shares(&_StrategyBase.CallOpts, user)
}

// SharesToStaking is a free data retrieval call binding the contract method 0x56a7117b.
//
// Solidity: function sharesToStaking(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) SharesToStaking(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "sharesToStaking", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToStaking is a free data retrieval call binding the contract method 0x56a7117b.
//
// Solidity: function sharesToStaking(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) SharesToStaking(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToStaking(&_StrategyBase.CallOpts, amountShares)
}

// SharesToStaking is a free data retrieval call binding the contract method 0x56a7117b.
//
// Solidity: function sharesToStaking(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) SharesToStaking(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToStaking(&_StrategyBase.CallOpts, amountShares)
}

// SharesToStakingView is a free data retrieval call binding the contract method 0x9910a665.
//
// Solidity: function sharesToStakingView(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) SharesToStakingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "sharesToStakingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToStakingView is a free data retrieval call binding the contract method 0x9910a665.
//
// Solidity: function sharesToStakingView(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) SharesToStakingView(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToStakingView(&_StrategyBase.CallOpts, amountShares)
}

// SharesToStakingView is a free data retrieval call binding the contract method 0x9910a665.
//
// Solidity: function sharesToStakingView(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) SharesToStakingView(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToStakingView(&_StrategyBase.CallOpts, amountShares)
}

// StakingToShares is a free data retrieval call binding the contract method 0x893d4d9e.
//
// Solidity: function stakingToShares(uint256 amountStaking) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) StakingToShares(opts *bind.CallOpts, amountStaking *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "stakingToShares", amountStaking)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingToShares is a free data retrieval call binding the contract method 0x893d4d9e.
//
// Solidity: function stakingToShares(uint256 amountStaking) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) StakingToShares(amountStaking *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.StakingToShares(&_StrategyBase.CallOpts, amountStaking)
}

// StakingToShares is a free data retrieval call binding the contract method 0x893d4d9e.
//
// Solidity: function stakingToShares(uint256 amountStaking) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) StakingToShares(amountStaking *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.StakingToShares(&_StrategyBase.CallOpts, amountStaking)
}

// StakingToSharesView is a free data retrieval call binding the contract method 0x10ed193a.
//
// Solidity: function stakingToSharesView(uint256 amountStaking) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) StakingToSharesView(opts *bind.CallOpts, amountStaking *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "stakingToSharesView", amountStaking)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingToSharesView is a free data retrieval call binding the contract method 0x10ed193a.
//
// Solidity: function stakingToSharesView(uint256 amountStaking) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) StakingToSharesView(amountStaking *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.StakingToSharesView(&_StrategyBase.CallOpts, amountStaking)
}

// StakingToSharesView is a free data retrieval call binding the contract method 0x10ed193a.
//
// Solidity: function stakingToSharesView(uint256 amountStaking) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) StakingToSharesView(amountStaking *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.StakingToSharesView(&_StrategyBase.CallOpts, amountStaking)
}

// StakingWeth is a free data retrieval call binding the contract method 0xfc1d98fb.
//
// Solidity: function stakingWeth() view returns(address)
func (_StrategyBase *StrategyBaseCaller) StakingWeth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "stakingWeth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingWeth is a free data retrieval call binding the contract method 0xfc1d98fb.
//
// Solidity: function stakingWeth() view returns(address)
func (_StrategyBase *StrategyBaseSession) StakingWeth() (common.Address, error) {
	return _StrategyBase.Contract.StakingWeth(&_StrategyBase.CallOpts)
}

// StakingWeth is a free data retrieval call binding the contract method 0xfc1d98fb.
//
// Solidity: function stakingWeth() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) StakingWeth() (common.Address, error) {
	return _StrategyBase.Contract.StakingWeth(&_StrategyBase.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyBase *StrategyBaseCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyBase *StrategyBaseSession) StrategyManager() (common.Address, error) {
	return _StrategyBase.Contract.StrategyManager(&_StrategyBase.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) StrategyManager() (common.Address, error) {
	return _StrategyBase.Contract.StrategyManager(&_StrategyBase.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) TotalShares() (*big.Int, error) {
	return _StrategyBase.Contract.TotalShares(&_StrategyBase.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) TotalShares() (*big.Int, error) {
	return _StrategyBase.Contract.TotalShares(&_StrategyBase.CallOpts)
}

// UserStakingView is a free data retrieval call binding the contract method 0xf5f1b29c.
//
// Solidity: function userStakingView(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) UserStakingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "userStakingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserStakingView is a free data retrieval call binding the contract method 0xf5f1b29c.
//
// Solidity: function userStakingView(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) UserStakingView(user common.Address) (*big.Int, error) {
	return _StrategyBase.Contract.UserStakingView(&_StrategyBase.CallOpts, user)
}

// UserStakingView is a free data retrieval call binding the contract method 0xf5f1b29c.
//
// Solidity: function userStakingView(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) UserStakingView(user common.Address) (*big.Int, error) {
	return _StrategyBase.Contract.UserStakingView(&_StrategyBase.CallOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address weth, uint256 amount) returns(uint256 newShares)
func (_StrategyBase *StrategyBaseTransactor) Deposit(opts *bind.TransactOpts, weth common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "deposit", weth, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address weth, uint256 amount) returns(uint256 newShares)
func (_StrategyBase *StrategyBaseSession) Deposit(weth common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Deposit(&_StrategyBase.TransactOpts, weth, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address weth, uint256 amount) returns(uint256 newShares)
func (_StrategyBase *StrategyBaseTransactorSession) Deposit(weth common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Deposit(&_StrategyBase.TransactOpts, weth, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _stakingWeth, address _relayer, address _strategyManager, address _pauser) returns()
func (_StrategyBase *StrategyBaseTransactor) Initialize(opts *bind.TransactOpts, _stakingWeth common.Address, _relayer common.Address, _strategyManager common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "initialize", _stakingWeth, _relayer, _strategyManager, _pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _stakingWeth, address _relayer, address _strategyManager, address _pauser) returns()
func (_StrategyBase *StrategyBaseSession) Initialize(_stakingWeth common.Address, _relayer common.Address, _strategyManager common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.Initialize(&_StrategyBase.TransactOpts, _stakingWeth, _relayer, _strategyManager, _pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _stakingWeth, address _relayer, address _strategyManager, address _pauser) returns()
func (_StrategyBase *StrategyBaseTransactorSession) Initialize(_stakingWeth common.Address, _relayer common.Address, _strategyManager common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.Initialize(&_StrategyBase.TransactOpts, _stakingWeth, _relayer, _strategyManager, _pauser)
}

// TransferETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0x548b9560.
//
// Solidity: function transferETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, address l1StakingManagerAddr, uint256 gasLimit, uint256 batchId) payable returns(bool)
func (_StrategyBase *StrategyBaseTransactor) TransferETHToL2DappLinkBridge(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, l1StakingManagerAddr common.Address, gasLimit *big.Int, batchId *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "transferETHToL2DappLinkBridge", sourceChainId, destChainId, bridge, l1StakingManagerAddr, gasLimit, batchId)
}

// TransferETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0x548b9560.
//
// Solidity: function transferETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, address l1StakingManagerAddr, uint256 gasLimit, uint256 batchId) payable returns(bool)
func (_StrategyBase *StrategyBaseSession) TransferETHToL2DappLinkBridge(sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, l1StakingManagerAddr common.Address, gasLimit *big.Int, batchId *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.TransferETHToL2DappLinkBridge(&_StrategyBase.TransactOpts, sourceChainId, destChainId, bridge, l1StakingManagerAddr, gasLimit, batchId)
}

// TransferETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0x548b9560.
//
// Solidity: function transferETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, address l1StakingManagerAddr, uint256 gasLimit, uint256 batchId) payable returns(bool)
func (_StrategyBase *StrategyBaseTransactorSession) TransferETHToL2DappLinkBridge(sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, l1StakingManagerAddr common.Address, gasLimit *big.Int, batchId *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.TransferETHToL2DappLinkBridge(&_StrategyBase.TransactOpts, sourceChainId, destChainId, bridge, l1StakingManagerAddr, gasLimit, batchId)
}

// TransferWETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0x14d1441f.
//
// Solidity: function transferWETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, address l1StakingManagerAddr, address wethAddress, uint256 gasLimit, uint256 batchId) payable returns(bool)
func (_StrategyBase *StrategyBaseTransactor) TransferWETHToL2DappLinkBridge(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, l1StakingManagerAddr common.Address, wethAddress common.Address, gasLimit *big.Int, batchId *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "transferWETHToL2DappLinkBridge", sourceChainId, destChainId, bridge, l1StakingManagerAddr, wethAddress, gasLimit, batchId)
}

// TransferWETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0x14d1441f.
//
// Solidity: function transferWETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, address l1StakingManagerAddr, address wethAddress, uint256 gasLimit, uint256 batchId) payable returns(bool)
func (_StrategyBase *StrategyBaseSession) TransferWETHToL2DappLinkBridge(sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, l1StakingManagerAddr common.Address, wethAddress common.Address, gasLimit *big.Int, batchId *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.TransferWETHToL2DappLinkBridge(&_StrategyBase.TransactOpts, sourceChainId, destChainId, bridge, l1StakingManagerAddr, wethAddress, gasLimit, batchId)
}

// TransferWETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0x14d1441f.
//
// Solidity: function transferWETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, address l1StakingManagerAddr, address wethAddress, uint256 gasLimit, uint256 batchId) payable returns(bool)
func (_StrategyBase *StrategyBaseTransactorSession) TransferWETHToL2DappLinkBridge(sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, l1StakingManagerAddr common.Address, wethAddress common.Address, gasLimit *big.Int, batchId *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.TransferWETHToL2DappLinkBridge(&_StrategyBase.TransactOpts, sourceChainId, destChainId, bridge, l1StakingManagerAddr, wethAddress, gasLimit, batchId)
}

// UserStaking is a paid mutator transaction binding the contract method 0x821460f5.
//
// Solidity: function userStaking(address user) returns(uint256)
func (_StrategyBase *StrategyBaseTransactor) UserStaking(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "userStaking", user)
}

// UserStaking is a paid mutator transaction binding the contract method 0x821460f5.
//
// Solidity: function userStaking(address user) returns(uint256)
func (_StrategyBase *StrategyBaseSession) UserStaking(user common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.UserStaking(&_StrategyBase.TransactOpts, user)
}

// UserStaking is a paid mutator transaction binding the contract method 0x821460f5.
//
// Solidity: function userStaking(address user) returns(uint256)
func (_StrategyBase *StrategyBaseTransactorSession) UserStaking(user common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.UserStaking(&_StrategyBase.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address weth, uint256 amountShares) returns()
func (_StrategyBase *StrategyBaseTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, weth common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "withdraw", recipient, weth, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address weth, uint256 amountShares) returns()
func (_StrategyBase *StrategyBaseSession) Withdraw(recipient common.Address, weth common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Withdraw(&_StrategyBase.TransactOpts, recipient, weth, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address weth, uint256 amountShares) returns()
func (_StrategyBase *StrategyBaseTransactorSession) Withdraw(recipient common.Address, weth common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Withdraw(&_StrategyBase.TransactOpts, recipient, weth, amountShares)
}

// StrategyBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StrategyBase contract.
type StrategyBaseInitializedIterator struct {
	Event *StrategyBaseInitialized // Event containing the contract specifics and raw log

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
func (it *StrategyBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseInitialized)
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
		it.Event = new(StrategyBaseInitialized)
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
func (it *StrategyBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseInitialized represents a Initialized event raised by the StrategyBase contract.
type StrategyBaseInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StrategyBase *StrategyBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*StrategyBaseInitializedIterator, error) {

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseInitializedIterator{contract: _StrategyBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StrategyBase *StrategyBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StrategyBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseInitialized)
				if err := _StrategyBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StrategyBase *StrategyBaseFilterer) ParseInitialized(log types.Log) (*StrategyBaseInitialized, error) {
	event := new(StrategyBaseInitialized)
	if err := _StrategyBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseTransferETHToL2DappLinkBridgeIterator is returned from FilterTransferETHToL2DappLinkBridge and is used to iterate over the raw logs and unpacked data for TransferETHToL2DappLinkBridge events raised by the StrategyBase contract.
type StrategyBaseTransferETHToL2DappLinkBridgeIterator struct {
	Event *StrategyBaseTransferETHToL2DappLinkBridge // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTransferETHToL2DappLinkBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTransferETHToL2DappLinkBridge)
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
		it.Event = new(StrategyBaseTransferETHToL2DappLinkBridge)
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
func (it *StrategyBaseTransferETHToL2DappLinkBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTransferETHToL2DappLinkBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTransferETHToL2DappLinkBridge represents a TransferETHToL2DappLinkBridge event raised by the StrategyBase contract.
type StrategyBaseTransferETHToL2DappLinkBridge struct {
	SourceChainId        *big.Int
	DestChainId          *big.Int
	Bridge               common.Address
	L1StakingManagerAddr common.Address
	TokenAddress         common.Address
	BridgeEthAmount      *big.Int
	BatchId              *big.Int
	Nonce                *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTransferETHToL2DappLinkBridge is a free log retrieval operation binding the contract event 0x1e41494416713921a8fa17cffca7814892d9c5c3be42174a8a5d95114635bce8.
//
// Solidity: event TransferETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, address l1StakingManagerAddr, address tokenAddress, uint256 bridgeEthAmount, uint256 batchId, uint256 nonce)
func (_StrategyBase *StrategyBaseFilterer) FilterTransferETHToL2DappLinkBridge(opts *bind.FilterOpts) (*StrategyBaseTransferETHToL2DappLinkBridgeIterator, error) {

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "TransferETHToL2DappLinkBridge")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTransferETHToL2DappLinkBridgeIterator{contract: _StrategyBase.contract, event: "TransferETHToL2DappLinkBridge", logs: logs, sub: sub}, nil
}

// WatchTransferETHToL2DappLinkBridge is a free log subscription operation binding the contract event 0x1e41494416713921a8fa17cffca7814892d9c5c3be42174a8a5d95114635bce8.
//
// Solidity: event TransferETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, address l1StakingManagerAddr, address tokenAddress, uint256 bridgeEthAmount, uint256 batchId, uint256 nonce)
func (_StrategyBase *StrategyBaseFilterer) WatchTransferETHToL2DappLinkBridge(opts *bind.WatchOpts, sink chan<- *StrategyBaseTransferETHToL2DappLinkBridge) (event.Subscription, error) {

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "TransferETHToL2DappLinkBridge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTransferETHToL2DappLinkBridge)
				if err := _StrategyBase.contract.UnpackLog(event, "TransferETHToL2DappLinkBridge", log); err != nil {
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

// ParseTransferETHToL2DappLinkBridge is a log parse operation binding the contract event 0x1e41494416713921a8fa17cffca7814892d9c5c3be42174a8a5d95114635bce8.
//
// Solidity: event TransferETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, address l1StakingManagerAddr, address tokenAddress, uint256 bridgeEthAmount, uint256 batchId, uint256 nonce)
func (_StrategyBase *StrategyBaseFilterer) ParseTransferETHToL2DappLinkBridge(log types.Log) (*StrategyBaseTransferETHToL2DappLinkBridge, error) {
	event := new(StrategyBaseTransferETHToL2DappLinkBridge)
	if err := _StrategyBase.contract.UnpackLog(event, "TransferETHToL2DappLinkBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
