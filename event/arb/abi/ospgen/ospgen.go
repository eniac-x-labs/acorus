// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ospgen

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

// ExecutionContext is an auto generated low-level Go binding around an user-defined struct.
type ExecutionContext struct {
	MaxInboxMessagesRead *big.Int
	Bridge               common.Address
}

// Instruction is an auto generated low-level Go binding around an user-defined struct.
type Instruction struct {
	Opcode       uint16
	ArgumentData *big.Int
}

// Machine is an auto generated low-level Go binding around an user-defined struct.
type Machine struct {
	Status          uint8
	ValueStack      ValueStack
	InternalStack   ValueStack
	FrameStack      StackFrameWindow
	GlobalStateHash [32]byte
	ModuleIdx       uint32
	FunctionIdx     uint32
	FunctionPc      uint32
	ModulesRoot     [32]byte
}

// Module is an auto generated low-level Go binding around an user-defined struct.
type Module struct {
	GlobalsMerkleRoot   [32]byte
	ModuleMemory        ModuleMemory
	TablesMerkleRoot    [32]byte
	FunctionsMerkleRoot [32]byte
	InternalsOffset     uint32
}

// ModuleMemory is an auto generated low-level Go binding around an user-defined struct.
type ModuleMemory struct {
	Size       uint64
	MaxSize    uint64
	MerkleRoot [32]byte
}

// StackFrame is an auto generated low-level Go binding around an user-defined struct.
type StackFrame struct {
	ReturnPc              Value
	LocalsMerkleRoot      [32]byte
	CallerModule          uint32
	CallerModuleInternals uint32
}

// StackFrameWindow is an auto generated low-level Go binding around an user-defined struct.
type StackFrameWindow struct {
	Proved        []StackFrame
	RemainingHash [32]byte
}

// Value is an auto generated low-level Go binding around an user-defined struct.
type Value struct {
	ValueType uint8
	Contents  *big.Int
}

// ValueArray is an auto generated low-level Go binding around an user-defined struct.
type ValueArray struct {
	Inner []Value
}

// ValueStack is an auto generated low-level Go binding around an user-defined struct.
type ValueStack struct {
	Proved        ValueArray
	RemainingHash [32]byte
}

// HashProofHelperMetaData contains all meta data concerning the HashProofHelper contract.
var HashProofHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"}],\"name\":\"NotProven\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"part\",\"type\":\"bytes\"}],\"name\":\"PreimagePartProven\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"clearSplitProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"}],\"name\":\"getPreimagePart\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keccakStates\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"part\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"}],\"name\":\"proveWithFullPreimage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"name\":\"proveWithSplitPreimage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611cea806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063740085d71461005c57806379754cba14610085578063ae364ac2146100a6578063b7465799146100b0578063d4e5dd2b146100d2575b600080fd5b61006f61006a3660046118e0565b6100e5565b60405161007c9190611959565b60405180910390f35b6100986100933660046119bb565b6101d8565b60405190815260200161007c565b6100ae6106cb565b005b6100c36100be366004611a16565b610713565b60405161007c93929190611a3f565b6100986100e0366004611a71565b6107c9565b6000828152602081815260408083206001600160401b0385168452909152902080546060919060ff16610142576040516309cb23c960e11b8152600481018590526001600160401b03841660248201526044015b60405180910390fd5b80600101805461015190611ac4565b80601f016020809104026020016040519081016040528092919081815260200182805461017d90611ac4565b80156101ca5780601f1061019f576101008083540402835291602001916101ca565b820191906000526020600020905b8154815290600101906020018083116101ad57829003601f168201915b505050505091505092915050565b60006001821615156002831615610230573360009081526001602081905260408220805467ffffffffffffffff191681559190610217908301826117a2565b6102256002830160006117df565b600982016000905550505b80806102445750610242608886611b15565b155b6102845760405162461bcd60e51b81526020600482015260116024820152701393d517d09313d0d2d7d0531251d39151607a1b6044820152606401610139565b3360009081526001602052604090206009810154806102bc57815467ffffffffffffffff19166001600160401b038716178255610306565b81546001600160401b038781169116146103065760405162461bcd60e51b815260206004820152600b60248201526a1112519197d3d19194d15560aa1b6044820152606401610139565b61031282898986610920565b8061032760206001600160401b038916611b3f565b11801561034057508160090154866001600160401b0316105b1561045a57600081876001600160401b0316111561036e5761036b826001600160401b038916611b57565b90505b60008261038560206001600160401b038b16611b3f565b61038f9190611b57565b90508881111561039c5750875b815b8181101561045657846001018b8b838181106103bc576103bc611b6e565b9050013560f81c60f81b90808054806103d490611ac4565b80601f81146103e2576103f8565b83600052602060002060ff1984168155603f9350505b506002919091019091558154600116156104215790600052602060002090602091828204019190065b909190919091601f036101000a81548160ff02191690600160f81b84040217905550808061044e90611b84565b91505061039e565b5050505b8261046c5750600092506106c3915050565b60005b602081101561053c576000610485600883611b9f565b9050610492600582611b9f565b61049d600583611b15565b6104a8906005611bb3565b6104b29190611b3f565b905060006104c1600884611b15565b6104cc906008611bb3565b8560020183601981106104e1576104e1611b6e565b60048104909101546001600160401b036008600390931683026101000a9091041690911c9150610512908490611bb3565b61051d9060f8611b57565b60ff909116901b9590951794508061053481611b84565b91505061046f565b50604051806040016040528060011515815260200183600101805461056090611ac4565b80601f016020809104026020016040519081016040528092919081815260200182805461058c90611ac4565b80156105d95780601f106105ae576101008083540402835291602001916105d9565b820191906000526020600020905b8154815290600101906020018083116105bc57829003601f168201915b50505091909252505060008581526020818152604080832086546001600160401b0316845282529091208251815460ff1916901515178155828201518051919261062b926001850192909101906117ee565b505082546040516001600160401b03909116915085907ff88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c90610671906001870190611bd2565b60405180910390a33360009081526001602081905260408220805467ffffffffffffffff1916815591906106a7908301826117a2565b6106b56002830160006117df565b600982016000905550505050505b949350505050565b3360009081526001602081905260408220805467ffffffffffffffff1916815591906106f9908301826117a2565b6107076002830160006117df565b60098201600090555050565b6001602081905260009182526040909120805491810180546001600160401b039093169261074090611ac4565b80601f016020809104026020016040519081016040528092919081815260200182805461076c90611ac4565b80156107b95780601f1061078e576101008083540402835291602001916107b9565b820191906000526020600020905b81548152906001019060200180831161079c57829003601f168201915b5050505050908060090154905083565b600083836040516107db929190611c7a565b604051908190039020905060606001600160401b03831684111561087957600061080e6001600160401b03851686611b57565b9050602081111561081d575060205b856001600160401b038516866108338483611b3f565b9261084093929190611c8a565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929450505050505b6040805180820182526001808252602080830185815260008781528083528581206001600160401b038a1682528352949094208351815460ff1916901515178155935180519394936108d29385019291909101906117ee565b50905050826001600160401b0316827ff88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c836040516109109190611959565b60405180910390a3509392505050565b828290508460090160008282546109379190611b3f565b90915550505b81158015610949575080155b1561095357610ba2565b60005b6088811015610a7a5760008382101561098c5784848381811061097b5761097b611b6e565b919091013560f81c91506109af9050565b81841415610998576001175b6109a460016088611b57565b8214156109af576080175b60006109bc600884611b9f565b90506109c9600582611b9f565b6109d4600583611b15565b6109df906005611bb3565b6109e99190611b3f565b90506109f6600884611b15565b610a01906008611bb3565b6001600160401b03168260ff166001600160401b0316901b876002018260198110610a2e57610a2e611b6e565b6004810490910180546001600160401b0360086003909416939093026101000a808204841690941883168402929093021990921617905550819050610a7281611b84565b915050610956565b50610a83611872565b60005b6019811015610af557856002018160198110610aa457610aa4611b6e565b600491828204019190066008029054906101000a90046001600160401b03166001600160401b0316828260198110610ade57610ade611b6e565b602002015280610aed81611b84565b915050610a86565b50610aff81610ba8565b905060005b6019811015610b7b57818160198110610b1f57610b1f611b6e565b6020020151866002018260198110610b3957610b39611b6e565b600491828204019190066008026101000a8154816001600160401b0302191690836001600160401b031602179055508080610b7390611b84565b915050610b04565b506088831015610b8b5750610ba2565b610b988360888187611c8a565b935093505061093d565b50505050565b610bb0611872565b610bb8611891565b610bc0611891565b610bc8611872565b600060405180610300016040528060018152602001618082815260200167800000000000808a8152602001678000000080008000815260200161808b81526020016380000001815260200167800000008000808181526020016780000000000080098152602001608a81526020016088815260200163800080098152602001638000000a8152602001638000808b815260200167800000000000008b8152602001678000000000008089815260200167800000000000800381526020016780000000000080028152602001678000000000000080815260200161800a815260200167800000008000000a81526020016780000000800080818152602001678000000000008080815260200163800000018152602001678000000080008008815250905060005b6018811015611797576080878101516060808a01516040808c01516020808e01518e511890911890921890931889526101208b01516101008c015160e08d015160c08e015160a08f0151181818189089018190526101c08b01516101a08c01516101808d01516101608e01516101408f0151181818189289019283526102608b01516102408c01516102208d01516102008e01516101e08f015118181818918901919091526103008a01516102e08b01516102c08c01516102a08d01516102808e0151181818189288018390526001600160401b0360028202166001603f1b91829004179092188652510485600260200201516002026001600160401b03161785600060200201511884600160200201526001603f1b856003602002015181610e1957610e19611aff565b0485600360200201516002026001600160401b03161785600160200201511884600260200201526001603f1b856004602002015181610e5a57610e5a611aff565b0485600460200201516002026001600160401b03161785600260058110610e8357610e83611b6e565b602002015118606085015284516001603f1b9086516060808901519390920460029091026001600160401b031617909118608086810191825286518a5118808b5287516020808d018051909218825289516040808f0180519092189091528a518e8801805190911890528a51948e0180519095189094528901805160a08e0180519091189052805160c08e0180519091189052805160e08e018051909118905280516101008e0180519091189052516101208d018051909118905291880180516101408d018051909118905280516101608d018051909118905280516101808d018051909118905280516101a08d0180519091189052516101c08c018051909118905292870180516101e08c018051909118905280516102008c018051909118905280516102208c018051909118905280516102408c0180519091189052516102608b018051909118905281516102808b018051909118905281516102a08b018051909118905281516102c08b018051909118905281516102e08b018051909118905290516103008a01805190911890529084525163100000009060208901516001600160401b03641000000000909102169190041761010084015260408701516001603d1b9060408901516001600160401b03600890910216919004176101608401526060870151628000009060608901516001600160401b036502000000000090910216919004176102608401526080870151654000000000009060808901516001600160401b036204000090910216919004176102c084015260a08701516001603f1b900487600560200201516002026001600160401b031617836002601981106110f3576110f3611b6e565b602002015260c08701516210000081046001602c1b9091026001600160401b039081169190911760a085015260e0880151664000000000000081046104009091028216176101a08501526101008801516208000081046520000000000090910282161761020085015261012088015160048082029092166001603e1b909104176103008501526101408801516101408901516001600160401b036001603e1b909102169190041760808401526101608701516001603a1b906101608901516001600160401b036040909102169190041760e084015261018087015162200000906101808901516001600160401b036001602b1b90910216919004176101408401526101a08701516602000000000000906101a08901516001600160401b0361800090910216919004176102408401526101c08701516008906101c08901516001600160401b036001603d1b90910216919004176102a08401526101e0870151641000000000906101e08901516001600160401b03631000000090910216919004176020840152610200808801516102008901516001600160401b0366800000000000009091021691900417610120840152610220870151648000000000906102208901516001600160401b03630200000090910216919004176101808401526102408701516001602b1b906102408901516001600160401b036220000090910216919004176101e0840152610260870151610100906102608901516001600160401b03600160381b90910216919004176102e0840152610280870151642000000000906102808901516001600160401b036308000000909102169190041760608401526102a08701516001602c1b906102a08901516001600160401b0362100000909102169190041760c08401526102c08701516302000000906102c08901516001600160401b0364800000000090910216919004176101c08401526102e0870151600160381b906102e08901516001600160401b036101009091021691900417610220840152610300870151660400000000000090048760186020020151614000026001600160401b031617836014602002015282600a602002015183600560200201511916836000602002015118876000602002015282600b602002015183600660200201511916836001602002015118876001602002015282600c602002015183600760200201511916836002602002015118876002602002015282600d602002015183600860200201511916836003602002015118876003602002015282600e602002015183600960200201511916836004602002015118876004602002015282600f602002015183600a602002015119168360056020020151188760056020020152826010602002015183600b602002015119168360066020020151188760066020020152826011602002015183600c602002015119168360076020020151188760076020020152826012602002015183600d602002015119168360086020020151188760086020020152826013602002015183600e602002015119168360096020020151188760096020020152826014602002015183600f6020020151191683600a60200201511887600a602002015282601560200201518360106020020151191683600b60200201511887600b602002015282601660200201518360116020020151191683600c60200201511887600c602002015282601760200201518360126020020151191683600d60200201511887600d602002015282601860200201518360136020020151191683600e60200201511887600e602002015282600060200201518360146020020151191683600f60200201511887600f602002015282600160200201518360156020020151191683601060200201511887601060200201528260026020020151836016602002015119168360116020020151188760116020020152826003602002015183601760200201511916836012602002015118876012602002015282600460200201518360186020020151191683601360200201511887601360200201528260056020020151836000602002015119168360146020020151188760146020020152826006602002015183600160200201511916836015602002015118876015602002015282600760200201518360026020020151191683601660200201511887601660200201528260086020020151836003602002015119168360176020020151188760176020020152826009602002015183600460200201511916836018602002015118876018602002015281816018811061178557611785611b6e565b60200201518751188752600101610cee565b509495945050505050565b5080546117ae90611ac4565b6000825580601f106117be575050565b601f0160209004906000526020600020908101906117dc91906118af565b50565b506117dc9060078101906118af565b8280546117fa90611ac4565b90600052602060002090601f01602090048101928261181c5760008555611862565b82601f1061183557805160ff1916838001178555611862565b82800160010185558215611862579182015b82811115611862578251825591602001919060010190611847565b5061186e9291506118af565b5090565b6040518061032001604052806019906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b5b8082111561186e57600081556001016118b0565b80356001600160401b03811681146118db57600080fd5b919050565b600080604083850312156118f357600080fd5b82359150611903602084016118c4565b90509250929050565b6000815180845260005b8181101561193257602081850181015186830182015201611916565b81811115611944576000602083870101525b50601f01601f19169290920160200192915050565b60208152600061196c602083018461190c565b9392505050565b60008083601f84011261198557600080fd5b5081356001600160401b0381111561199c57600080fd5b6020830191508360208285010111156119b457600080fd5b9250929050565b600080600080606085870312156119d157600080fd5b84356001600160401b038111156119e757600080fd5b6119f387828801611973565b9095509350611a069050602086016118c4565b9396929550929360400135925050565b600060208284031215611a2857600080fd5b81356001600160a01b038116811461196c57600080fd5b6001600160401b0384168152606060208201526000611a61606083018561190c565b9050826040830152949350505050565b600080600060408486031215611a8657600080fd5b83356001600160401b03811115611a9c57600080fd5b611aa886828701611973565b9094509250611abb9050602085016118c4565b90509250925092565b600181811c90821680611ad857607f821691505b60208210811415611af957634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601260045260246000fd5b600082611b2457611b24611aff565b500690565b634e487b7160e01b600052601160045260246000fd5b60008219821115611b5257611b52611b29565b500190565b600082821015611b6957611b69611b29565b500390565b634e487b7160e01b600052603260045260246000fd5b6000600019821415611b9857611b98611b29565b5060010190565b600082611bae57611bae611aff565b500490565b6000816000190483118215151615611bcd57611bcd611b29565b500290565b600060208083526000845481600182811c915080831680611bf457607f831692505b858310811415611c1257634e487b7160e01b85526022600452602485fd5b878601838152602001818015611c2f5760018114611c4057611c6b565b60ff19861682528782019650611c6b565b60008b81526020902060005b86811015611c6557815484820152908501908901611c4c565b83019750505b50949998505050505050505050565b8183823760009101908152919050565b60008085851115611c9a57600080fd5b83861115611ca757600080fd5b505082019391909203915056fea26469706673582212202f2dcb3af29934bcb41f7861970718df9ad65661f9f6cb26c88c963224a1e4c264736f6c63430008090033",
}

// HashProofHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use HashProofHelperMetaData.ABI instead.
var HashProofHelperABI = HashProofHelperMetaData.ABI

// HashProofHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HashProofHelperMetaData.Bin instead.
var HashProofHelperBin = HashProofHelperMetaData.Bin

// DeployHashProofHelper deploys a new Ethereum contract, binding an instance of HashProofHelper to it.
func DeployHashProofHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HashProofHelper, error) {
	parsed, err := HashProofHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HashProofHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HashProofHelper{HashProofHelperCaller: HashProofHelperCaller{contract: contract}, HashProofHelperTransactor: HashProofHelperTransactor{contract: contract}, HashProofHelperFilterer: HashProofHelperFilterer{contract: contract}}, nil
}

// HashProofHelper is an auto generated Go binding around an Ethereum contract.
type HashProofHelper struct {
	HashProofHelperCaller     // Read-only binding to the contract
	HashProofHelperTransactor // Write-only binding to the contract
	HashProofHelperFilterer   // Log filterer for contract events
}

// HashProofHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashProofHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashProofHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashProofHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashProofHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashProofHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashProofHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashProofHelperSession struct {
	Contract     *HashProofHelper  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashProofHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashProofHelperCallerSession struct {
	Contract *HashProofHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// HashProofHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashProofHelperTransactorSession struct {
	Contract     *HashProofHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// HashProofHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashProofHelperRaw struct {
	Contract *HashProofHelper // Generic contract binding to access the raw methods on
}

// HashProofHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashProofHelperCallerRaw struct {
	Contract *HashProofHelperCaller // Generic read-only contract binding to access the raw methods on
}

// HashProofHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashProofHelperTransactorRaw struct {
	Contract *HashProofHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashProofHelper creates a new instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelper(address common.Address, backend bind.ContractBackend) (*HashProofHelper, error) {
	contract, err := bindHashProofHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashProofHelper{HashProofHelperCaller: HashProofHelperCaller{contract: contract}, HashProofHelperTransactor: HashProofHelperTransactor{contract: contract}, HashProofHelperFilterer: HashProofHelperFilterer{contract: contract}}, nil
}

// NewHashProofHelperCaller creates a new read-only instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelperCaller(address common.Address, caller bind.ContractCaller) (*HashProofHelperCaller, error) {
	contract, err := bindHashProofHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperCaller{contract: contract}, nil
}

// NewHashProofHelperTransactor creates a new write-only instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*HashProofHelperTransactor, error) {
	contract, err := bindHashProofHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperTransactor{contract: contract}, nil
}

// NewHashProofHelperFilterer creates a new log filterer instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*HashProofHelperFilterer, error) {
	contract, err := bindHashProofHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperFilterer{contract: contract}, nil
}

// bindHashProofHelper binds a generic wrapper to an already deployed contract.
func bindHashProofHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HashProofHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashProofHelper *HashProofHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashProofHelper.Contract.HashProofHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashProofHelper *HashProofHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashProofHelper.Contract.HashProofHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashProofHelper *HashProofHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashProofHelper.Contract.HashProofHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashProofHelper *HashProofHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashProofHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashProofHelper *HashProofHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashProofHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashProofHelper *HashProofHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashProofHelper.Contract.contract.Transact(opts, method, params...)
}

// GetPreimagePart is a free data retrieval call binding the contract method 0x740085d7.
//
// Solidity: function getPreimagePart(bytes32 fullHash, uint64 offset) view returns(bytes)
func (_HashProofHelper *HashProofHelperCaller) GetPreimagePart(opts *bind.CallOpts, fullHash [32]byte, offset uint64) ([]byte, error) {
	var out []interface{}
	err := _HashProofHelper.contract.Call(opts, &out, "getPreimagePart", fullHash, offset)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPreimagePart is a free data retrieval call binding the contract method 0x740085d7.
//
// Solidity: function getPreimagePart(bytes32 fullHash, uint64 offset) view returns(bytes)
func (_HashProofHelper *HashProofHelperSession) GetPreimagePart(fullHash [32]byte, offset uint64) ([]byte, error) {
	return _HashProofHelper.Contract.GetPreimagePart(&_HashProofHelper.CallOpts, fullHash, offset)
}

// GetPreimagePart is a free data retrieval call binding the contract method 0x740085d7.
//
// Solidity: function getPreimagePart(bytes32 fullHash, uint64 offset) view returns(bytes)
func (_HashProofHelper *HashProofHelperCallerSession) GetPreimagePart(fullHash [32]byte, offset uint64) ([]byte, error) {
	return _HashProofHelper.Contract.GetPreimagePart(&_HashProofHelper.CallOpts, fullHash, offset)
}

// KeccakStates is a free data retrieval call binding the contract method 0xb7465799.
//
// Solidity: function keccakStates(address ) view returns(uint64 offset, bytes part, uint256 length)
func (_HashProofHelper *HashProofHelperCaller) KeccakStates(opts *bind.CallOpts, arg0 common.Address) (struct {
	Offset uint64
	Part   []byte
	Length *big.Int
}, error) {
	var out []interface{}
	err := _HashProofHelper.contract.Call(opts, &out, "keccakStates", arg0)

	outstruct := new(struct {
		Offset uint64
		Part   []byte
		Length *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Offset = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Part = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Length = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// KeccakStates is a free data retrieval call binding the contract method 0xb7465799.
//
// Solidity: function keccakStates(address ) view returns(uint64 offset, bytes part, uint256 length)
func (_HashProofHelper *HashProofHelperSession) KeccakStates(arg0 common.Address) (struct {
	Offset uint64
	Part   []byte
	Length *big.Int
}, error) {
	return _HashProofHelper.Contract.KeccakStates(&_HashProofHelper.CallOpts, arg0)
}

// KeccakStates is a free data retrieval call binding the contract method 0xb7465799.
//
// Solidity: function keccakStates(address ) view returns(uint64 offset, bytes part, uint256 length)
func (_HashProofHelper *HashProofHelperCallerSession) KeccakStates(arg0 common.Address) (struct {
	Offset uint64
	Part   []byte
	Length *big.Int
}, error) {
	return _HashProofHelper.Contract.KeccakStates(&_HashProofHelper.CallOpts, arg0)
}

// ClearSplitProof is a paid mutator transaction binding the contract method 0xae364ac2.
//
// Solidity: function clearSplitProof() returns()
func (_HashProofHelper *HashProofHelperTransactor) ClearSplitProof(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashProofHelper.contract.Transact(opts, "clearSplitProof")
}

// ClearSplitProof is a paid mutator transaction binding the contract method 0xae364ac2.
//
// Solidity: function clearSplitProof() returns()
func (_HashProofHelper *HashProofHelperSession) ClearSplitProof() (*types.Transaction, error) {
	return _HashProofHelper.Contract.ClearSplitProof(&_HashProofHelper.TransactOpts)
}

// ClearSplitProof is a paid mutator transaction binding the contract method 0xae364ac2.
//
// Solidity: function clearSplitProof() returns()
func (_HashProofHelper *HashProofHelperTransactorSession) ClearSplitProof() (*types.Transaction, error) {
	return _HashProofHelper.Contract.ClearSplitProof(&_HashProofHelper.TransactOpts)
}

// ProveWithFullPreimage is a paid mutator transaction binding the contract method 0xd4e5dd2b.
//
// Solidity: function proveWithFullPreimage(bytes data, uint64 offset) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactor) ProveWithFullPreimage(opts *bind.TransactOpts, data []byte, offset uint64) (*types.Transaction, error) {
	return _HashProofHelper.contract.Transact(opts, "proveWithFullPreimage", data, offset)
}

// ProveWithFullPreimage is a paid mutator transaction binding the contract method 0xd4e5dd2b.
//
// Solidity: function proveWithFullPreimage(bytes data, uint64 offset) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperSession) ProveWithFullPreimage(data []byte, offset uint64) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithFullPreimage(&_HashProofHelper.TransactOpts, data, offset)
}

// ProveWithFullPreimage is a paid mutator transaction binding the contract method 0xd4e5dd2b.
//
// Solidity: function proveWithFullPreimage(bytes data, uint64 offset) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactorSession) ProveWithFullPreimage(data []byte, offset uint64) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithFullPreimage(&_HashProofHelper.TransactOpts, data, offset)
}

// ProveWithSplitPreimage is a paid mutator transaction binding the contract method 0x79754cba.
//
// Solidity: function proveWithSplitPreimage(bytes data, uint64 offset, uint256 flags) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactor) ProveWithSplitPreimage(opts *bind.TransactOpts, data []byte, offset uint64, flags *big.Int) (*types.Transaction, error) {
	return _HashProofHelper.contract.Transact(opts, "proveWithSplitPreimage", data, offset, flags)
}

// ProveWithSplitPreimage is a paid mutator transaction binding the contract method 0x79754cba.
//
// Solidity: function proveWithSplitPreimage(bytes data, uint64 offset, uint256 flags) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperSession) ProveWithSplitPreimage(data []byte, offset uint64, flags *big.Int) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithSplitPreimage(&_HashProofHelper.TransactOpts, data, offset, flags)
}

// ProveWithSplitPreimage is a paid mutator transaction binding the contract method 0x79754cba.
//
// Solidity: function proveWithSplitPreimage(bytes data, uint64 offset, uint256 flags) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactorSession) ProveWithSplitPreimage(data []byte, offset uint64, flags *big.Int) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithSplitPreimage(&_HashProofHelper.TransactOpts, data, offset, flags)
}

// HashProofHelperPreimagePartProvenIterator is returned from FilterPreimagePartProven and is used to iterate over the raw logs and unpacked data for PreimagePartProven events raised by the HashProofHelper contract.
type HashProofHelperPreimagePartProvenIterator struct {
	Event *HashProofHelperPreimagePartProven // Event containing the contract specifics and raw log

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
func (it *HashProofHelperPreimagePartProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashProofHelperPreimagePartProven)
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
		it.Event = new(HashProofHelperPreimagePartProven)
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
func (it *HashProofHelperPreimagePartProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashProofHelperPreimagePartProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashProofHelperPreimagePartProven represents a PreimagePartProven event raised by the HashProofHelper contract.
type HashProofHelperPreimagePartProven struct {
	FullHash [32]byte
	Offset   uint64
	Part     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPreimagePartProven is a free log retrieval operation binding the contract event 0xf88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c.
//
// Solidity: event PreimagePartProven(bytes32 indexed fullHash, uint64 indexed offset, bytes part)
func (_HashProofHelper *HashProofHelperFilterer) FilterPreimagePartProven(opts *bind.FilterOpts, fullHash [][32]byte, offset []uint64) (*HashProofHelperPreimagePartProvenIterator, error) {

	var fullHashRule []interface{}
	for _, fullHashItem := range fullHash {
		fullHashRule = append(fullHashRule, fullHashItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _HashProofHelper.contract.FilterLogs(opts, "PreimagePartProven", fullHashRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperPreimagePartProvenIterator{contract: _HashProofHelper.contract, event: "PreimagePartProven", logs: logs, sub: sub}, nil
}

// WatchPreimagePartProven is a free log subscription operation binding the contract event 0xf88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c.
//
// Solidity: event PreimagePartProven(bytes32 indexed fullHash, uint64 indexed offset, bytes part)
func (_HashProofHelper *HashProofHelperFilterer) WatchPreimagePartProven(opts *bind.WatchOpts, sink chan<- *HashProofHelperPreimagePartProven, fullHash [][32]byte, offset []uint64) (event.Subscription, error) {

	var fullHashRule []interface{}
	for _, fullHashItem := range fullHash {
		fullHashRule = append(fullHashRule, fullHashItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _HashProofHelper.contract.WatchLogs(opts, "PreimagePartProven", fullHashRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashProofHelperPreimagePartProven)
				if err := _HashProofHelper.contract.UnpackLog(event, "PreimagePartProven", log); err != nil {
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

// ParsePreimagePartProven is a log parse operation binding the contract event 0xf88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c.
//
// Solidity: event PreimagePartProven(bytes32 indexed fullHash, uint64 indexed offset, bytes part)
func (_HashProofHelper *HashProofHelperFilterer) ParsePreimagePartProven(log types.Log) (*HashProofHelperPreimagePartProven, error) {
	event := new(HashProofHelperPreimagePartProven)
	if err := _HashProofHelper.contract.UnpackLog(event, "PreimagePartProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOneStepProofEntryMetaData contains all meta data concerning the IOneStepProofEntry contract.
var IOneStepProofEntryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"machineStep\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveOneStep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IOneStepProofEntryABI is the input ABI used to generate the binding from.
// Deprecated: Use IOneStepProofEntryMetaData.ABI instead.
var IOneStepProofEntryABI = IOneStepProofEntryMetaData.ABI

// IOneStepProofEntry is an auto generated Go binding around an Ethereum contract.
type IOneStepProofEntry struct {
	IOneStepProofEntryCaller     // Read-only binding to the contract
	IOneStepProofEntryTransactor // Write-only binding to the contract
	IOneStepProofEntryFilterer   // Log filterer for contract events
}

// IOneStepProofEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOneStepProofEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOneStepProofEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOneStepProofEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOneStepProofEntrySession struct {
	Contract     *IOneStepProofEntry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IOneStepProofEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOneStepProofEntryCallerSession struct {
	Contract *IOneStepProofEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IOneStepProofEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOneStepProofEntryTransactorSession struct {
	Contract     *IOneStepProofEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IOneStepProofEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOneStepProofEntryRaw struct {
	Contract *IOneStepProofEntry // Generic contract binding to access the raw methods on
}

// IOneStepProofEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOneStepProofEntryCallerRaw struct {
	Contract *IOneStepProofEntryCaller // Generic read-only contract binding to access the raw methods on
}

// IOneStepProofEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOneStepProofEntryTransactorRaw struct {
	Contract *IOneStepProofEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOneStepProofEntry creates a new instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntry(address common.Address, backend bind.ContractBackend) (*IOneStepProofEntry, error) {
	contract, err := bindIOneStepProofEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntry{IOneStepProofEntryCaller: IOneStepProofEntryCaller{contract: contract}, IOneStepProofEntryTransactor: IOneStepProofEntryTransactor{contract: contract}, IOneStepProofEntryFilterer: IOneStepProofEntryFilterer{contract: contract}}, nil
}

// NewIOneStepProofEntryCaller creates a new read-only instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntryCaller(address common.Address, caller bind.ContractCaller) (*IOneStepProofEntryCaller, error) {
	contract, err := bindIOneStepProofEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntryCaller{contract: contract}, nil
}

// NewIOneStepProofEntryTransactor creates a new write-only instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*IOneStepProofEntryTransactor, error) {
	contract, err := bindIOneStepProofEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntryTransactor{contract: contract}, nil
}

// NewIOneStepProofEntryFilterer creates a new log filterer instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*IOneStepProofEntryFilterer, error) {
	contract, err := bindIOneStepProofEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntryFilterer{contract: contract}, nil
}

// bindIOneStepProofEntry binds a generic wrapper to an already deployed contract.
func bindIOneStepProofEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProofEntry *IOneStepProofEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProofEntry.Contract.IOneStepProofEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProofEntry *IOneStepProofEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.IOneStepProofEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProofEntry *IOneStepProofEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.IOneStepProofEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProofEntry *IOneStepProofEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProofEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProofEntry *IOneStepProofEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProofEntry *IOneStepProofEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.contract.Transact(opts, method, params...)
}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntryCaller) ProveOneStep(opts *bind.CallOpts, execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	var out []interface{}
	err := _IOneStepProofEntry.contract.Call(opts, &out, "proveOneStep", execCtx, machineStep, beforeHash, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntrySession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.ProveOneStep(&_IOneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntryCallerSession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.ProveOneStep(&_IOneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// IOneStepProverMetaData contains all meta data concerning the IOneStepProver contract.
var IOneStepProverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"instruction\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"result\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"resultMod\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IOneStepProverABI is the input ABI used to generate the binding from.
// Deprecated: Use IOneStepProverMetaData.ABI instead.
var IOneStepProverABI = IOneStepProverMetaData.ABI

// IOneStepProver is an auto generated Go binding around an Ethereum contract.
type IOneStepProver struct {
	IOneStepProverCaller     // Read-only binding to the contract
	IOneStepProverTransactor // Write-only binding to the contract
	IOneStepProverFilterer   // Log filterer for contract events
}

// IOneStepProverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOneStepProverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOneStepProverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOneStepProverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOneStepProverSession struct {
	Contract     *IOneStepProver   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOneStepProverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOneStepProverCallerSession struct {
	Contract *IOneStepProverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IOneStepProverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOneStepProverTransactorSession struct {
	Contract     *IOneStepProverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IOneStepProverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOneStepProverRaw struct {
	Contract *IOneStepProver // Generic contract binding to access the raw methods on
}

// IOneStepProverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOneStepProverCallerRaw struct {
	Contract *IOneStepProverCaller // Generic read-only contract binding to access the raw methods on
}

// IOneStepProverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOneStepProverTransactorRaw struct {
	Contract *IOneStepProverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOneStepProver creates a new instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProver(address common.Address, backend bind.ContractBackend) (*IOneStepProver, error) {
	contract, err := bindIOneStepProver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOneStepProver{IOneStepProverCaller: IOneStepProverCaller{contract: contract}, IOneStepProverTransactor: IOneStepProverTransactor{contract: contract}, IOneStepProverFilterer: IOneStepProverFilterer{contract: contract}}, nil
}

// NewIOneStepProverCaller creates a new read-only instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProverCaller(address common.Address, caller bind.ContractCaller) (*IOneStepProverCaller, error) {
	contract, err := bindIOneStepProver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProverCaller{contract: contract}, nil
}

// NewIOneStepProverTransactor creates a new write-only instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProverTransactor(address common.Address, transactor bind.ContractTransactor) (*IOneStepProverTransactor, error) {
	contract, err := bindIOneStepProver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProverTransactor{contract: contract}, nil
}

// NewIOneStepProverFilterer creates a new log filterer instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProverFilterer(address common.Address, filterer bind.ContractFilterer) (*IOneStepProverFilterer, error) {
	contract, err := bindIOneStepProver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOneStepProverFilterer{contract: contract}, nil
}

// bindIOneStepProver binds a generic wrapper to an already deployed contract.
func bindIOneStepProver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOneStepProverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProver *IOneStepProverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProver.Contract.IOneStepProverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProver *IOneStepProverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProver.Contract.IOneStepProverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProver *IOneStepProverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProver.Contract.IOneStepProverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProver *IOneStepProverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProver *IOneStepProverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProver *IOneStepProverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProver.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xda78e7d1.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverCaller) ExecuteOneStep(opts *bind.CallOpts, execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	var out []interface{}
	err := _IOneStepProver.contract.Call(opts, &out, "executeOneStep", execCtx, mach, mod, instruction, proof)

	outstruct := new(struct {
		Result    Machine
		ResultMod Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Result = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.ResultMod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xda78e7d1.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverSession) ExecuteOneStep(execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	return _IOneStepProver.Contract.ExecuteOneStep(&_IOneStepProver.CallOpts, execCtx, mach, mod, instruction, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xda78e7d1.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverCallerSession) ExecuteOneStep(execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	return _IOneStepProver.Contract.ExecuteOneStep(&_IOneStepProver.CallOpts, execCtx, mach, mod, instruction, proof)
}

// OneStepProofEntryMetaData contains all meta data concerning the OneStepProofEntry contract.
var OneStepProofEntryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"prover0_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverMem_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverMath_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverHostIo_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"machineStep\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveOneStep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prover0\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverHostIo\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverMath\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverMem\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002378380380620023788339810160408190526200003491620000a5565b600080546001600160a01b039586166001600160a01b031991821617909155600180549486169482169490941790935560028054928516928416929092179091556003805491909316911617905562000102565b80516001600160a01b0381168114620000a057600080fd5b919050565b60008060008060808587031215620000bc57600080fd5b620000c78562000088565b9350620000d76020860162000088565b9250620000e76040860162000088565b9150620000f76060860162000088565b905092959194509250565b61226680620001126000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631f128bc01461005c57806330a5509f1461008c5780635d3adcfb1461009f5780635f52fd7c146100c057806366e5d9c3146100d3575b600080fd5b60015461006f906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b60005461006f906001600160a01b031681565b6100b26100ad3660046117fb565b6100e6565b604051908152602001610083565b60035461006f906001600160a01b031681565b60025461006f906001600160a01b031681565b60006100f06116f7565b6100f8611769565b6040805160208101909152606081526040805180820190915260008082526020820152600061012888888361069c565b9095509050886101378661086b565b1461017f5760405162461bcd60e51b815260206004820152601360248201527209a828690929c8abe848a8c9ea48abe9082a69606b1b60448201526064015b60405180910390fd5b60008551600381111561019457610194611896565b146101ae576101a28561086b565b95505050505050610693565b650800000000006101c08b60016118c2565b14156101d357600285526101a28561086b565b6101de888883610a5a565b90945090506101ee888883610b06565b809250819450505084610100015161021b8660a0015163ffffffff168686610be09092919063ffffffff16565b146102575760405162461bcd60e51b815260206004820152600c60248201526b1353d115531154d7d493d3d560a21b6044820152606401610176565b6040805160208101909152606081526040805160208101909152606081526102808a8a85610c29565b90945092506102908a8a85610b06565b9350915061029f8a8a85610b06565b809450819250505060006102c88860e0015163ffffffff168685610c839092919063ffffffff16565b905060006102eb8960c0015163ffffffff168385610cc99092919063ffffffff16565b9050876060015181146103355760405162461bcd60e51b815260206004820152601260248201527110905117d1955390d51253d394d7d493d3d560721b6044820152606401610176565b506103489250899150839050818b6118da565b975097505060008460a0015163ffffffff16905060018560e00181815161036f9190611904565b63ffffffff1690525081516000602861ffff8316108015906103965750603561ffff831611155b806103b65750603661ffff8316108015906103b65750603e61ffff831611155b806103c5575061ffff8216603f145b806103d4575061ffff82166040145b156103eb57506001546001600160a01b03166105e0565b61ffff821660451480610402575061ffff82166050145b806104305750604661ffff83161080159061043057506104246009604661192c565b61ffff168261ffff1611155b8061045e5750606761ffff83161080159061045e57506104526002606761192c565b61ffff168261ffff1611155b8061047e5750606a61ffff83161080159061047e5750607861ffff831611155b806104ac5750605161ffff8316108015906104ac57506104a06009605161192c565b61ffff168261ffff1611155b806104da5750607961ffff8316108015906104da57506104ce6002607961192c565b61ffff168261ffff1611155b806104fa5750607c61ffff8316108015906104fa5750608a61ffff831611155b80610509575061ffff821660a7145b80610526575061ffff821660ac1480610526575061ffff821660ad145b80610546575060c061ffff831610801590610546575060c461ffff831611155b80610566575060bc61ffff831610801590610566575060bf61ffff831611155b1561057d57506002546001600160a01b03166105e0565b61801061ffff831610801590610599575061801361ffff831611155b806105bb575061802061ffff8316108015906105bb575061802261ffff831611155b156105d257506003546001600160a01b03166105e0565b506000546001600160a01b03165b806001600160a01b031663da78e7d18e8989888f8f6040518763ffffffff1660e01b815260040161061696959493929190611a8b565b60006040518083038186803b15801561062e57600080fd5b505afa158015610642573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261066a9190810190612042565b909750955061067a858488610be0565b6101008801526106898761086b565b9750505050505050505b95945050505050565b6106a46116f7565b816000806106b3878785610d3e565b9350905060ff81166106c85760009150610745565b8060ff16600114156106dd5760019150610745565b8060ff16600214156106f25760029150610745565b8060ff16600314156107075760039150610745565b60405162461bcd60e51b8152602060048201526013602482015272554e4b4e4f574e5f4d4143485f53544154555360681b6044820152606401610176565b5061074e6117ba565b6107566117ba565b60008060008061077760408051808201909152606081526000602082015290565b60006107848e8e8c610d74565b9a5097506107938e8e8c610d74565b9a5096506107a28e8e8c610e73565b9a5091506107b18e8e8c610f9b565b9a5095506107c08e8e8c610fb7565b9a5094506107cf8e8e8c610fb7565b9a5093506107de8e8e8c610fb7565b9a5092506107ed8e8e8c610f9b565b809b5081925050506040518061012001604052808a600381111561081357610813611896565b81526020018981526020018881526020018381526020018781526020018663ffffffff1681526020018563ffffffff1681526020018463ffffffff168152602001828152509a50505050505050505050935093915050565b6000808251600381111561088157610881611896565b141561095057610894826020015161101b565b6108a1836040015161101b565b6108ae84606001516110a0565b608085015160a086015160c087015160e0808901516101008a01516040516f26b0b1b434b73290393ab73734b7339d60811b602082015260308101999099526050890197909752607088019590955260908701939093526001600160e01b031991831b821660b0870152821b811660b486015291901b1660b883015260bc82015260dc015b604051602081830303815290604052805190602001209050919050565b60018251600381111561096557610965611896565b141561099d5760808201516040517026b0b1b434b732903334b734b9b432b21d60791b60208201526031810191909152605101610933565b6002825160038111156109b2576109b2611896565b14156109dc576040516f26b0b1b434b7329032b93937b932b21d60811b6020820152603001610933565b6003825160038111156109f1576109f1611896565b1415610a1b576040516f26b0b1b434b732903a37b7903330b91d60811b6020820152603001610933565b60405162461bcd60e51b815260206004820152600f60248201526e4241445f4d4143485f53544154555360881b6044820152606401610176565b919050565b610a62611769565b604080516060810182526000808252602082018190529181018290528391906000806000610a918a8a88610f9b565b96509450610aa08a8a88611139565b96509350610aaf8a8a88610f9b565b96509250610abe8a8a88610f9b565b96509150610acd8a8a88610fb7565b6040805160a08101825297885260208801969096529486019390935250606084015263ffffffff16608083015290969095509350505050565b604080516020810190915260608152816000610b23868684610d3e565b92509050600060ff82166001600160401b03811115610b4457610b44611c19565b604051908082528060200260200182016040528015610b6d578160200160208202803683370190505b50905060005b8260ff168160ff161015610bc457610b8c888886610f9b565b838360ff1681518110610ba157610ba1612164565b602002602001018196508281525050508080610bbc9061217a565b915050610b73565b5060405180602001604052808281525093505050935093915050565b6000610c218484610bf0856111b4565b6040518060400160405280601381526020017226b7b23ab6329036b2b935b632903a3932b29d60691b815250611221565b949350505050565b604080518082019091526000808252602082015281600080610c4c8787856112f3565b93509150610c5b87878561134c565b6040805180820190915261ffff90941684526020840191909152919791965090945050505050565b6000610c218484610c93856113a1565b6040518060400160405280601881526020017724b739ba393ab1ba34b7b71036b2b935b632903a3932b29d60411b815250611221565b60405168233ab731ba34b7b71d60b91b602082015260298101829052600090819060490160405160208183030381529060405280519060200120905061069385858360405180604001604052806015815260200174233ab731ba34b7b71036b2b935b632903a3932b29d60591b815250611221565b600081848482818110610d5357610d53612164565b919091013560f81c9250819050610d698161219a565b915050935093915050565b610d7c6117ba565b816000610d8a868684610f9b565b925090506000610d9b87878561134c565b935090506000816001600160401b03811115610db957610db9611c19565b604051908082528060200260200182016040528015610dfe57816020015b6040805180820190915260008082526020820152815260200190600190039081610dd75790505b50905060005b8151811015610e4c57610e188989876113eb565b838381518110610e2a57610e2a612164565b6020026020010181975082905250508080610e449061219a565b915050610e04565b50604080516060810182529081019182529081526020810192909252509590945092505050565b604080518082019091526060815260006020820152816000610e96868684610f9b565b925090506060868684818110610eae57610eae612164565b909101356001600160f81b031916159050610f365782610ecd8161219a565b604080516001808252818301909252919550909150816020015b610eef6117d8565b815260200190600190039081610ee7579050509050610f0f8787856114e7565b82600081518110610f2257610f22612164565b602002602001018195508290525050610f7a565b82610f408161219a565b60408051600080825260208201909252919550909150610f76565b610f636117d8565b815260200190600190039081610f5b5790505b5090505b60405180604001604052808281526020018381525093505050935093915050565b60008181610faa86868461134c565b9097909650945050505050565b600081815b60048110156110125760088363ffffffff16901b9250858583818110610fe457610fe4612164565b919091013560f81c93909317925081610ffc8161219a565b925050808061100a9061219a565b915050610fbc565b50935093915050565b60208101518151515160005b818110156110995783516110449061103f9083611580565b6115b8565b6040516b2b30b63ab29039ba30b1b59d60a11b6020820152602c810191909152604c8101849052606c0160405160208183030381529060405280519060200120925080806110919061219a565b915050611027565b5050919050565b602081015160005b825151811015611133576110d8836000015182815181106110cb576110cb612164565b60200260200101516115d5565b6040517129ba30b1b590333930b6b29039ba30b1b59d60711b6020820152603281019190915260528101839052607201604051602081830303815290604052805190602001209150808061112b9061219a565b9150506110a8565b50919050565b60408051606081018252600080825260208201819052918101919091528160008080611166888886611645565b94509250611175888886611645565b94509150611184888886610f9b565b604080516060810182526001600160401b0396871681529490951660208501529383015250969095509350505050565b600081600001516111c883602001516116a3565b6040848101516060860151608087015192516626b7b23ab6329d60c91b6020820152602781019590955260478501939093526067840152608783019190915260e01b6001600160e01b03191660a782015260ab01610933565b8160005b8551518110156112ea57600185166112865782828760000151838151811061124f5761124f612164565b6020026020010151604051602001611269939291906121b5565b6040516020818303038152906040528051906020012091506112d1565b828660000151828151811061129d5761129d612164565b6020026020010151836040516020016112b8939291906121b5565b6040516020818303038152906040528051906020012091505b60019490941c93806112e28161219a565b915050611225565b50949350505050565b600081815b60028110156110125760088361ffff16901b925085858381811061131e5761131e612164565b919091013560f81c939093179250816113368161219a565b92505080806113449061219a565b9150506112f8565b600081815b602081101561101257600883901b925085858381811061137357611373612164565b919091013560f81c9390931792508161138b8161219a565b92505080806113999061219a565b915050611351565b6000816000015182602001516040516020016109339291906b24b739ba393ab1ba34b7b71d60a11b815260f09290921b6001600160f01b031916600c830152600e820152602e0190565b604080518082019091526000808252602082015281600085858381811061141457611414612164565b919091013560f81c915082905061142a8161219a565b925050611435600690565b600681111561144657611446611896565b60ff168160ff16111561148c5760405162461bcd60e51b815260206004820152600e60248201526d4241445f56414c55455f5459504560901b6044820152606401610176565b600061149987878561134c565b809450819250505060405180604001604052808360ff1660068111156114c1576114c1611896565b60068111156114d2576114d2611896565b81526020018281525093505050935093915050565b6114ef6117d8565b8161150a604080518082019091526000808252602082015290565b600080600061151a8989876113eb565b95509350611529898987610f9b565b95509250611538898987610fb7565b95509150611547898987610fb7565b60408051608081018252968752602087019590955263ffffffff9384169486019490945290911660608401525090969095509350505050565b604080518082019091526000808252602082015282518051839081106115a8576115a8612164565b6020026020010151905092915050565b6000816000015182602001516040516020016109339291906121fb565b60006115e482600001516115b8565b602080840151604080860151606087015191516b29ba30b1b590333930b6b29d60a11b94810194909452602c840194909452604c8301919091526001600160e01b031960e093841b8116606c840152921b9091166070820152607401610933565b600081815b6008811015611012576008836001600160401b0316901b925085858381811061167557611675612164565b919091013560f81c9390931792508161168d8161219a565b925050808061169b9061219a565b91505061164a565b805160208083015160408085015190516626b2b6b7b93c9d60c91b938101939093526001600160c01b031960c094851b811660278501529190931b16602f8201526037810191909152600090605701610933565b60408051610120810190915280600081526020016117136117ba565b81526020016117206117ba565b815260200161174060408051808201909152606081526000602082015290565b815260006020820181905260408201819052606082018190526080820181905260a09091015290565b6040805160a08101909152600081526020810161179f604080516060810182526000808252602082018190529181019190915290565b81526000602082018190526040820181905260609091015290565b60408051606080820183529181019182529081526000602082015290565b6040805160c0810190915260006080820181815260a0830191909152819061179f565b600080600080600085870360a081121561181457600080fd5b604081121561182257600080fd5b50859450604086013593506060860135925060808601356001600160401b038082111561184e57600080fd5b818801915088601f83011261186257600080fd5b81358181111561187157600080fd5b89602082850101111561188357600080fd5b9699959850939650602001949392505050565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082198211156118d5576118d56118ac565b500190565b600080858511156118ea57600080fd5b838611156118f757600080fd5b5050820193919092039150565b600063ffffffff808316818516808303821115611923576119236118ac565b01949350505050565b600061ffff808316818516808303821115611923576119236118ac565b6004811061195957611959611896565b9052565b80516007811061196f5761196f611896565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b808410156119cc576119b882865161195d565b9382019360019390930192908501906119a5565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b82811015611a4b578451611a1785825161195d565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a090930192600101611a02565b509687015197909601969096525093949350505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8635815260006101a060208901356001600160a01b038116808214611aaf57600080fd5b8060208601525050806040840152611aca8184018951611949565b5060208701516101206101c0840152611ae76102c084018261197c565b9050604088015161019f1980858403016101e0860152611b07838361197c565b925060608a01519150808584030161020086015250611b2682826119e0565b915050608088015161022084015260a0880151611b4c61024085018263ffffffff169052565b5060c088015163ffffffff81166102608501525060e088015163ffffffff8116610280850152506101008801516102a0840152611be160608401888051825260208101516001600160401b0380825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a083015263ffffffff60808201511660c08301525050565b855161ffff166101408401526020860151610160840152828103610180840152611c0c818587611a62565b9998505050505050505050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715611c5157611c51611c19565b60405290565b604051602081016001600160401b0381118282101715611c5157611c51611c19565b604051608081016001600160401b0381118282101715611c5157611c51611c19565b60405160a081016001600160401b0381118282101715611c5157611c51611c19565b604051606081016001600160401b0381118282101715611c5157611c51611c19565b60405161012081016001600160401b0381118282101715611c5157611c51611c19565b604051601f8201601f191681016001600160401b0381118282101715611d2a57611d2a611c19565b604052919050565b805160048110610a5557600080fd5b60006001600160401b03821115611d5a57611d5a611c19565b5060051b60200190565b600060408284031215611d7657600080fd5b611d7e611c2f565b9050815160078110611d8f57600080fd5b808252506020820151602082015292915050565b60006040808385031215611db657600080fd5b611dbe611c2f565b915082516001600160401b0380821115611dd757600080fd5b81850191506020808388031215611ded57600080fd5b611df5611c57565b835183811115611e0457600080fd5b80850194505087601f850112611e1957600080fd5b83519250611e2e611e2984611d41565b611d02565b83815260069390931b84018201928281019089851115611e4d57600080fd5b948301945b84861015611e7357611e648a87611d64565b82529486019490830190611e52565b8252508552948501519484019490945250909392505050565b805163ffffffff81168114610a5557600080fd5b60006040808385031215611eb357600080fd5b611ebb611c2f565b915082516001600160401b03811115611ed357600080fd5b8301601f81018513611ee457600080fd5b80516020611ef4611e2983611d41565b82815260a09283028401820192828201919089851115611f1357600080fd5b948301945b84861015611f7c5780868b031215611f305760008081fd5b611f38611c79565b611f428b88611d64565b815287870151858201526060611f59818901611e8c565b89830152611f6960808901611e8c565b9082015283529485019491830191611f18565b50808752505080860151818601525050505092915050565b80516001600160401b0381168114610a5557600080fd5b600081830360e0811215611fbe57600080fd5b611fc6611c9b565b8351815291506060601f1982011215611fde57600080fd5b50611fe7611cbd565b611ff360208401611f94565b815261200160408401611f94565b602082015260608301516040820152806020830152506080820151604082015260a0820151606082015261203760c08301611e8c565b608082015292915050565b60008061010080848603121561205757600080fd5b83516001600160401b038082111561206e57600080fd5b90850190610120828803121561208357600080fd5b61208b611cdf565b61209483611d32565b81526020830151828111156120a857600080fd5b6120b489828601611da3565b6020830152506040830151828111156120cc57600080fd5b6120d889828601611da3565b6040830152506060830151828111156120f057600080fd5b6120fc89828601611ea0565b6060830152506080830151608082015261211860a08401611e8c565b60a082015261212960c08401611e8c565b60c082015261213a60e08401611e8c565b60e082015283830151848201528095505050505061215b8460208501611fab565b90509250929050565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff811415612191576121916118ac565b60010192915050565b60006000198214156121ae576121ae6118ac565b5060010190565b6000845160005b818110156121d657602081880181015185830152016121bc565b818111156121e5576000828501525b5091909101928352506020820152604001919050565b652b30b63ab29d60d11b815260006007841061221957612219611896565b5060f89290921b600683015260078201526027019056fea2646970667358221220700823cc910eeee191650902368f52a36b4136c62bc727da6ef5233ea8d9d43e64736f6c63430008090033",
}

// OneStepProofEntryABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProofEntryMetaData.ABI instead.
var OneStepProofEntryABI = OneStepProofEntryMetaData.ABI

// OneStepProofEntryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProofEntryMetaData.Bin instead.
var OneStepProofEntryBin = OneStepProofEntryMetaData.Bin

// DeployOneStepProofEntry deploys a new Ethereum contract, binding an instance of OneStepProofEntry to it.
func DeployOneStepProofEntry(auth *bind.TransactOpts, backend bind.ContractBackend, prover0_ common.Address, proverMem_ common.Address, proverMath_ common.Address, proverHostIo_ common.Address) (common.Address, *types.Transaction, *OneStepProofEntry, error) {
	parsed, err := OneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProofEntryBin), backend, prover0_, proverMem_, proverMath_, proverHostIo_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProofEntry{OneStepProofEntryCaller: OneStepProofEntryCaller{contract: contract}, OneStepProofEntryTransactor: OneStepProofEntryTransactor{contract: contract}, OneStepProofEntryFilterer: OneStepProofEntryFilterer{contract: contract}}, nil
}

// OneStepProofEntry is an auto generated Go binding around an Ethereum contract.
type OneStepProofEntry struct {
	OneStepProofEntryCaller     // Read-only binding to the contract
	OneStepProofEntryTransactor // Write-only binding to the contract
	OneStepProofEntryFilterer   // Log filterer for contract events
}

// OneStepProofEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofEntrySession struct {
	Contract     *OneStepProofEntry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OneStepProofEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofEntryCallerSession struct {
	Contract *OneStepProofEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OneStepProofEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofEntryTransactorSession struct {
	Contract     *OneStepProofEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OneStepProofEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofEntryRaw struct {
	Contract *OneStepProofEntry // Generic contract binding to access the raw methods on
}

// OneStepProofEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofEntryCallerRaw struct {
	Contract *OneStepProofEntryCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofEntryTransactorRaw struct {
	Contract *OneStepProofEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProofEntry creates a new instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntry(address common.Address, backend bind.ContractBackend) (*OneStepProofEntry, error) {
	contract, err := bindOneStepProofEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntry{OneStepProofEntryCaller: OneStepProofEntryCaller{contract: contract}, OneStepProofEntryTransactor: OneStepProofEntryTransactor{contract: contract}, OneStepProofEntryFilterer: OneStepProofEntryFilterer{contract: contract}}, nil
}

// NewOneStepProofEntryCaller creates a new read-only instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntryCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofEntryCaller, error) {
	contract, err := bindOneStepProofEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryCaller{contract: contract}, nil
}

// NewOneStepProofEntryTransactor creates a new write-only instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofEntryTransactor, error) {
	contract, err := bindOneStepProofEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryTransactor{contract: contract}, nil
}

// NewOneStepProofEntryFilterer creates a new log filterer instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofEntryFilterer, error) {
	contract, err := bindOneStepProofEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryFilterer{contract: contract}, nil
}

// bindOneStepProofEntry binds a generic wrapper to an already deployed contract.
func bindOneStepProofEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntry *OneStepProofEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntry.Contract.OneStepProofEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntry *OneStepProofEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.OneStepProofEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntry *OneStepProofEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.OneStepProofEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntry *OneStepProofEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntry *OneStepProofEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntry *OneStepProofEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.contract.Transact(opts, method, params...)
}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProveOneStep(opts *bind.CallOpts, execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proveOneStep", execCtx, machineStep, beforeHash, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntrySession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.ProveOneStep(&_OneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.ProveOneStep(&_OneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// Prover0 is a free data retrieval call binding the contract method 0x30a5509f.
//
// Solidity: function prover0() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) Prover0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "prover0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Prover0 is a free data retrieval call binding the contract method 0x30a5509f.
//
// Solidity: function prover0() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) Prover0() (common.Address, error) {
	return _OneStepProofEntry.Contract.Prover0(&_OneStepProofEntry.CallOpts)
}

// Prover0 is a free data retrieval call binding the contract method 0x30a5509f.
//
// Solidity: function prover0() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) Prover0() (common.Address, error) {
	return _OneStepProofEntry.Contract.Prover0(&_OneStepProofEntry.CallOpts)
}

// ProverHostIo is a free data retrieval call binding the contract method 0x5f52fd7c.
//
// Solidity: function proverHostIo() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProverHostIo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proverHostIo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverHostIo is a free data retrieval call binding the contract method 0x5f52fd7c.
//
// Solidity: function proverHostIo() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) ProverHostIo() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverHostIo(&_OneStepProofEntry.CallOpts)
}

// ProverHostIo is a free data retrieval call binding the contract method 0x5f52fd7c.
//
// Solidity: function proverHostIo() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProverHostIo() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverHostIo(&_OneStepProofEntry.CallOpts)
}

// ProverMath is a free data retrieval call binding the contract method 0x66e5d9c3.
//
// Solidity: function proverMath() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProverMath(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proverMath")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverMath is a free data retrieval call binding the contract method 0x66e5d9c3.
//
// Solidity: function proverMath() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) ProverMath() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMath(&_OneStepProofEntry.CallOpts)
}

// ProverMath is a free data retrieval call binding the contract method 0x66e5d9c3.
//
// Solidity: function proverMath() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProverMath() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMath(&_OneStepProofEntry.CallOpts)
}

// ProverMem is a free data retrieval call binding the contract method 0x1f128bc0.
//
// Solidity: function proverMem() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProverMem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proverMem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverMem is a free data retrieval call binding the contract method 0x1f128bc0.
//
// Solidity: function proverMem() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) ProverMem() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMem(&_OneStepProofEntry.CallOpts)
}

// ProverMem is a free data retrieval call binding the contract method 0x1f128bc0.
//
// Solidity: function proverMem() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProverMem() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMem(&_OneStepProofEntry.CallOpts)
}

// OneStepProofEntryLibMetaData contains all meta data concerning the OneStepProofEntryLib contract.
var OneStepProofEntryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220995bc75fd3882c2ddc7f079b12262d41b3a96596d620dafd8f32afd385ba28e164736f6c63430008090033",
}

// OneStepProofEntryLibABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProofEntryLibMetaData.ABI instead.
var OneStepProofEntryLibABI = OneStepProofEntryLibMetaData.ABI

// OneStepProofEntryLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProofEntryLibMetaData.Bin instead.
var OneStepProofEntryLibBin = OneStepProofEntryLibMetaData.Bin

// DeployOneStepProofEntryLib deploys a new Ethereum contract, binding an instance of OneStepProofEntryLib to it.
func DeployOneStepProofEntryLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProofEntryLib, error) {
	parsed, err := OneStepProofEntryLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProofEntryLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProofEntryLib{OneStepProofEntryLibCaller: OneStepProofEntryLibCaller{contract: contract}, OneStepProofEntryLibTransactor: OneStepProofEntryLibTransactor{contract: contract}, OneStepProofEntryLibFilterer: OneStepProofEntryLibFilterer{contract: contract}}, nil
}

// OneStepProofEntryLib is an auto generated Go binding around an Ethereum contract.
type OneStepProofEntryLib struct {
	OneStepProofEntryLibCaller     // Read-only binding to the contract
	OneStepProofEntryLibTransactor // Write-only binding to the contract
	OneStepProofEntryLibFilterer   // Log filterer for contract events
}

// OneStepProofEntryLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofEntryLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofEntryLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofEntryLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofEntryLibSession struct {
	Contract     *OneStepProofEntryLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OneStepProofEntryLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofEntryLibCallerSession struct {
	Contract *OneStepProofEntryLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// OneStepProofEntryLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofEntryLibTransactorSession struct {
	Contract     *OneStepProofEntryLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// OneStepProofEntryLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofEntryLibRaw struct {
	Contract *OneStepProofEntryLib // Generic contract binding to access the raw methods on
}

// OneStepProofEntryLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofEntryLibCallerRaw struct {
	Contract *OneStepProofEntryLibCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofEntryLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofEntryLibTransactorRaw struct {
	Contract *OneStepProofEntryLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProofEntryLib creates a new instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLib(address common.Address, backend bind.ContractBackend) (*OneStepProofEntryLib, error) {
	contract, err := bindOneStepProofEntryLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLib{OneStepProofEntryLibCaller: OneStepProofEntryLibCaller{contract: contract}, OneStepProofEntryLibTransactor: OneStepProofEntryLibTransactor{contract: contract}, OneStepProofEntryLibFilterer: OneStepProofEntryLibFilterer{contract: contract}}, nil
}

// NewOneStepProofEntryLibCaller creates a new read-only instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLibCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofEntryLibCaller, error) {
	contract, err := bindOneStepProofEntryLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLibCaller{contract: contract}, nil
}

// NewOneStepProofEntryLibTransactor creates a new write-only instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLibTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofEntryLibTransactor, error) {
	contract, err := bindOneStepProofEntryLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLibTransactor{contract: contract}, nil
}

// NewOneStepProofEntryLibFilterer creates a new log filterer instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLibFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofEntryLibFilterer, error) {
	contract, err := bindOneStepProofEntryLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLibFilterer{contract: contract}, nil
}

// bindOneStepProofEntryLib binds a generic wrapper to an already deployed contract.
func bindOneStepProofEntryLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProofEntryLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntryLib *OneStepProofEntryLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntryLib.Contract.OneStepProofEntryLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntryLib *OneStepProofEntryLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.OneStepProofEntryLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntryLib *OneStepProofEntryLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.OneStepProofEntryLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntryLib *OneStepProofEntryLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntryLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntryLib *OneStepProofEntryLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntryLib *OneStepProofEntryLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.contract.Transact(opts, method, params...)
}

// OneStepProver0MetaData contains all meta data concerning the OneStepProver0 contract.
var OneStepProver0MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612640806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063da78e7d114610030575b600080fd5b61004361003e366004611bc6565b61005a565b604051610051929190611dd8565b60405180910390f35b610062611a8e565b61006a611b37565b61007387612280565b915061008436879003870187612383565b90506000610095602087018761241a565b9050611b8161ffff82166100ac57506102bd61029f565b61ffff8216600114156100c257506102c861029f565b61ffff8216600f14156100d857506102cf61029f565b61ffff8216601014156100ee57506103fb61029f565b61ffff82166180091415610105575061049561029f565b61ffff821661800a141561011c575061054161029f565b61ffff821660111415610132575061062e61029f565b61ffff821661800314156101495750610a0f61029f565b61ffff821661800414156101605750610a4e61029f565b61ffff8216602014156101765750610aac61029f565b61ffff82166021141561018c5750610aee61029f565b61ffff8216602314156101a25750610b3361029f565b61ffff8216602414156101b85750610b5b61029f565b61ffff821661800214156101cf5750610b8b61029f565b61ffff8216601a14156101e55750610c2861029f565b61ffff8216601b14156101fb5750610c3561029f565b604161ffff8316108015906102155750604461ffff831611155b156102235750610ca461029f565b61ffff8216618005148061023c575061ffff8216618006145b1561024a5750610d9861029f565b61ffff821661800814156102615750610e6b61029f565b60405162461bcd60e51b815260206004820152600e60248201526d494e56414c49445f4f50434f444560901b60448201526064015b60405180910390fd5b6102b084848989898663ffffffff16565b5050965096945050505050565b505060029092525050565b5050505050565b60006102de8660600151610e7a565b9050600481515160068111156102f6576102f6611ca9565b141561031d578560025b9081600381111561031357610313611ca9565b81525050506102c8565b6006815151600681111561033357610333611ca9565b146103795760405162461bcd60e51b8152602060048201526016602482015275494e56414c49445f52455455524e5f50435f5459504560501b6044820152606401610296565b805160209081015190819081901c604082901c606083901c156103d75760405162461bcd60e51b8152602060048201526016602482015275494e56414c49445f52455455524e5f50435f4441544160501b6044820152606401610296565b63ffffffff92831660e08b015290821660c08a01521660a088015250505050505050565b61041261040786610f1a565b602087015190610f7d565b60006104218660600151610f8d565b905061043e6104338260400151610fd9565b602088015190610f7d565b61044e6104338260600151610fd9565b602084013563ffffffff811681146104785760405162461bcd60e51b81526004016102969061243e565b63ffffffff1660c08701525050600060e090940193909352505050565b6104a161040786610f1a565b6104b16104078660a00151610fd9565b6104c16104078560800151610fd9565b6020808401359081901c604082901c1561051d5760405162461bcd60e51b815260206004820152601a60248201527f4241445f43524f53535f4d4f44554c455f43414c4c5f444154410000000000006044820152606401610296565b63ffffffff90811660a08801521660c08601525050600060e0909301929092525050565b61054d61040786610f1a565b61055d6104078660a00151610fd9565b61056d6104078560800151610fd9565b600061057c8660600151610f8d565b9050806060015163ffffffff166000141561059957856002610300565b602084013563ffffffff811681146105f35760405162461bcd60e51b815260206004820152601d60248201527f4241445f43414c4c45525f494e5445524e414c5f43414c4c5f444154410000006044820152606401610296565b604082015163ffffffff1660a0880152606082015161061390829061247b565b63ffffffff1660c08801525050600060e08601525050505050565b600080610646610641886020015161100c565b611031565b90506000806000808060006106676040518060200160405280606081525090565b6106728b8b876110c2565b955093506106818b8b87611129565b90965094506106918b8b87611145565b955092506106a08b8b876110c2565b955091506106af8b8b87611129565b90975094506106bf8b8b8761117b565b6040516d21b0b6361034b73234b932b1ba1d60911b60208201526001600160c01b031960c088901b16602e8201526036810189905290965090915060009060560160408051601f19818403018152919052805160209182012091508d013581146107645760405162461bcd60e51b81526020600482015260166024820152754241445f43414c4c5f494e4449524543545f4441544160501b6044820152606401610296565b61077a826001600160401b03871686868c611255565b90508d6040015181146107c15760405162461bcd60e51b815260206004820152600f60248201526e10905117d51050931154d7d493d3d5608a1b6044820152606401610296565b826001600160401b03168963ffffffff16106107eb57505060028d52506102c89650505050505050565b5050505050600061080c604080518082019091526000808252602082015290565b6040805160208101909152606081526108268a8a86611129565b945092506108358a8a866112f7565b945091506108448a8a8661117b565b9450905060006108618263ffffffff808b1690879087906113f316565b90508681146108a65760405162461bcd60e51b815260206004820152601160248201527010905117d153115351539514d7d493d3d5607a1b6044820152606401610296565b8584146108d6578d60025b908160038111156108c4576108c4611ca9565b815250505050505050505050506102c8565b6004835160068111156108eb576108eb611ca9565b14156108f9578d60026108b1565b60058351600681111561090e5761090e611ca9565b141561096d576020830151985063ffffffff891689146109685760405162461bcd60e51b81526020600482015260156024820152744241445f46554e435f5245465f434f4e54454e545360581b6044820152606401610296565b6109a5565b60405162461bcd60e51b815260206004820152600d60248201526c4241445f454c454d5f5459504560981b6044820152606401610296565b50505050505050506109b961043387610f1a565b60006109c88760600151610f8d565b90506109e56109da8260400151610fd9565b602089015190610f7d565b6109f56109da8260600151610fd9565b5063ffffffff1660c0860152600060e08601525050505050565b602083013563ffffffff81168114610a395760405162461bcd60e51b81526004016102969061243e565b63ffffffff1660e09095019490945250505050565b6000610a60610641876020015161100c565b905063ffffffff811615610aa457602084013563ffffffff81168114610a985760405162461bcd60e51b81526004016102969061243e565b63ffffffff1660e08701525b505050505050565b6000610abb8660600151610f8d565b90506000610ad382602001518660200135868661148d565b6020880151909150610ae59082610f7d565b50505050505050565b6000610afd866020015161100c565b90506000610b0e8760600151610f8d565b9050610b2581602001518660200135848787611525565b602090910152505050505050565b6000610b4985600001518560200135858561148d565b6020870151909150610aa49082610f7d565b6000610b6a866020015161100c565b9050610b8185600001518560200135838686611525565b9094525050505050565b6000610b9a866020015161100c565b90506000610bab876020015161100c565b90506000610bbc886020015161100c565b905060006040518060800160405280838152602001886020013560001b8152602001610be785611031565b63ffffffff168152602001610bfb86611031565b63ffffffff168152509050610c1d818a606001516115bf90919063ffffffff16565b505050505050505050565b610aa4856020015161100c565b6000610c47610641876020015161100c565b90506000610c58876020015161100c565b90506000610c69886020015161100c565b905063ffffffff831615610c8b576020880151610c869082610f7d565b610c9a565b6020880151610c9a9083610f7d565b5050505050505050565b6000610cb3602085018561241a565b9050600061ffff821660411415610ccc57506000610d4f565b61ffff821660421415610ce157506001610d4f565b61ffff821660431415610cf657506002610d4f565b61ffff821660441415610d0b57506003610d4f565b60405162461bcd60e51b8152602060048201526019602482015278434f4e53545f505553485f494e56414c49445f4f50434f444560381b6044820152606401610296565b610ae56040518060400160405280836006811115610d6f57610d6f611ca9565b815260200187602001356001600160401b03168152508860200151610f7d90919063ffffffff16565b6040805180820190915260008082526020820152618005610dbc602086018661241a565b61ffff161415610dea57610dd3866020015161100c565b6040870151909150610de59082610f7d565b610aa4565b618006610dfa602086018661241a565b61ffff161415610e2357610e11866040015161100c565b6020870151909150610de59082610f7d565b60405162461bcd60e51b815260206004820152601c60248201527f4d4f56455f494e5445524e414c5f494e56414c49445f4f50434f4445000000006044820152606401610296565b6000610b4986602001516116a6565b610e82611b8b565b815151600114610ea45760405162461bcd60e51b8152600401610296906124a3565b81518051600090610eb757610eb76124ce565b6020026020010151905060006001600160401b03811115610eda57610eda611f00565b604051908082528060200260200182016040528015610f1357816020015b610f00611b8b565b815260200190600190039081610ef85790505b5090915290565b604080518082018252600080825260209182015260e083015160c084015160a090940151835180850185526006815263ffffffff90921694831b67ffffffff0000000016949094179390921b63ffffffff60401b16929092179181019190915290565b8151610f8990826116db565b5050565b610f95611b8b565b815151600114610fb75760405162461bcd60e51b8152600401610296906124a3565b81518051600090610fca57610fca6124ce565b60200260200101519050919050565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b6040805180820190915260008082526020820152815161102b906117a4565b92915050565b6020810151600090818351600681111561104d5761104d611ca9565b146110845760405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b6044820152606401610296565b640100000000811061102b5760405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b6044820152606401610296565b600081815b6008811015611120576008836001600160401b0316901b92508585838181106110f2576110f26124ce565b919091013560f81c9390931792508161110a816124e4565b9250508080611118906124e4565b9150506110c7565b50935093915050565b600081816111388686846118ad565b9097909650945050505050565b60008184848281811061115a5761115a6124ce565b919091013560f81c9250819050611170816124e4565b915050935093915050565b604080516020810190915260608152816000611198868684611145565b92509050600060ff82166001600160401b038111156111b9576111b9611f00565b6040519080825280602002602001820160405280156111e2578160200160208202803683370190505b50905060005b8260ff168160ff16101561123957611201888886611129565b838360ff1681518110611216576112166124ce565b602002602001018196508281525050508080611231906124ff565b9150506111e8565b5060405180602001604052808281525093505050935093915050565b604051652a30b136329d60d11b60208201526001600160f81b031960f885901b1660268201526001600160c01b031960c084901b166027820152602f81018290526000908190604f016040516020818303038152906040528051906020012090506112ec878783604051806040016040528060128152602001712a30b136329036b2b935b632903a3932b29d60711b815250611902565b979650505050505050565b6040805180820190915260008082526020820152816000858583818110611320576113206124ce565b919091013560f81c9150829050611336816124e4565b925050611341600690565b600681111561135257611352611ca9565b60ff168160ff1611156113985760405162461bcd60e51b815260206004820152600e60248201526d4241445f56414c55455f5459504560901b6044820152606401610296565b60006113a58787856118ad565b809450819250505060405180604001604052808360ff1660068111156113cd576113cd611ca9565b60068111156113de576113de611ca9565b81526020018281525093505050935093915050565b60008083611400846119d4565b6040516d2a30b136329032b632b6b2b73a1d60911b6020820152602e810192909252604e820152606e016040516020818303038152906040528051906020012090506114838686836040518060400160405280601a81526020017f5461626c6520656c656d656e74206d65726b6c6520747265653a000000000000815250611902565b9695505050505050565b604080518082019091526000808252602082015260006114bd604080518082019091526000808252602082015290565b6040805160208101909152606081526114d78686856112f7565b935091506114e686868561117b565b9350905060006114f7828985611a0e565b90508881146115185760405162461bcd60e51b81526004016102969061251f565b5090979650505050505050565b6000611541604080518082019091526000808252602082015290565b60006115596040518060200160405280606081525090565b6115648686846112f7565b909350915061157486868461117b565b925090506000611585828a86611a0e565b90508981146115a65760405162461bcd60e51b81526004016102969061251f565b6115b1828a8a611a0e565b9a9950505050505050505050565b8151516000906115d090600161254a565b6001600160401b038111156115e7576115e7611f00565b60405190808252806020026020018201604052801561162057816020015b61160d611b8b565b8152602001906001900390816116055790505b50905060005b83515181101561167c578351805182908110611644576116446124ce565b602002602001015182828151811061165e5761165e6124ce565b60200260200101819052508080611674906124e4565b915050611626565b50818184600001515181518110611695576116956124ce565b602090810291909101015290915250565b6040805180820190915260008082526020820152815151516116d46116cc600183612562565b845190611a56565b9392505050565b8151516000906116ec90600161254a565b6001600160401b0381111561170357611703611f00565b60405190808252806020026020018201604052801561174857816020015b60408051808201909152600080825260208201528152602001906001900390816117215790505b50905060005b83515181101561167c57835180518290811061176c5761176c6124ce565b6020026020010151828281518110611786576117866124ce565b6020026020010181905250808061179c906124e4565b91505061174e565b6040805180820190915260008082526020820152815180516117c890600190612562565b815181106117d8576117d86124ce565b60200260200101519050600060018360000151516117f69190612562565b6001600160401b0381111561180d5761180d611f00565b60405190808252806020026020018201604052801561185257816020015b604080518082019091526000808252602082015281526020019060019003908161182b5790505b50905060005b8151811015610f13578351805182908110611875576118756124ce565b602002602001015182828151811061188f5761188f6124ce565b602002602001018190525080806118a5906124e4565b915050611858565b600081815b602081101561112057600883901b92508585838181106118d4576118d46124ce565b919091013560f81c939093179250816118ec816124e4565b92505080806118fa906124e4565b9150506118b2565b8160005b8551518110156119cb576001851661196757828287600001518381518110611930576119306124ce565b602002602001015160405160200161194a93929190612579565b6040516020818303038152906040528051906020012091506119b2565b828660000151828151811061197e5761197e6124ce565b60200260200101518360405160200161199993929190612579565b6040516020818303038152906040528051906020012091505b60019490941c93806119c3816124e4565b915050611906565b50949350505050565b6000816000015182602001516040516020016119f19291906125bf565b604051602081830303815290604052805190602001209050919050565b6000611a4e8484611a1e856119d4565b604051806040016040528060128152602001712b30b63ab29036b2b935b632903a3932b29d60711b815250611902565b949350505050565b60408051808201909152600080825260208201528251805183908110611a7e57611a7e6124ce565b6020026020010151905092915050565b6040805161012081019091528060008152602001611ac360408051606080820183529181019182529081526000602082015290565b8152602001611ae960408051606080820183529181019182529081526000602082015290565b8152602001611b0e604051806040016040528060608152602001600080191681525090565b815260006020820181905260408201819052606082018190526080820181905260a09091015290565b6040805160a0810182526000808252825160608101845281815260208181018390529381019190915290918201905b81526000602082018190526040820181905260609091015290565b611b896125f4565b565b6040805160c0810190915260006080820181815260a08301919091528190611b66565b600060408284031215611bc057600080fd5b50919050565b6000806000806000808688036101a0811215611be157600080fd5b611beb8989611bae565b965060408801356001600160401b0380821115611c0757600080fd5b90890190610120828c031215611c1c57600080fd5b81975060e0605f1984011215611c3157600080fd5b60608a019650611c458b6101408c01611bae565b95506101808a0135925080831115611c5c57600080fd5b828a0192508a601f840112611c7057600080fd5b8235915080821115611c8157600080fd5b50896020828401011115611c9457600080fd5b60208201935080925050509295509295509295565b634e487b7160e01b600052602160045260246000fd5b60048110611ccf57611ccf611ca9565b9052565b805160078110611ce557611ce5611ca9565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b80841015611d4257611d2e828651611cd3565b938201936001939093019290850190611d1b565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b82811015611dc1578451611d8d858251611cd3565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a090930192600101611d78565b509687015197909601969096525093949350505050565b6000610100808352611ded8184018651611cbf565b602085015161012084810152611e07610220850182611cf2565b9050604086015160ff198086840301610140870152611e268383611cf2565b925060608801519150808684030161016087015250611e458282611d56565b915050608086015161018085015260a0860151611e6b6101a086018263ffffffff169052565b5060c086015163ffffffff81166101c08601525060e086015163ffffffff81166101e0860152509085015161020084015290506116d460208301848051825260208101516001600160401b0380825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a083015263ffffffff60808201511660c08301525050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715611f3857611f38611f00565b60405290565b604051602081016001600160401b0381118282101715611f3857611f38611f00565b604051608081016001600160401b0381118282101715611f3857611f38611f00565b60405161012081016001600160401b0381118282101715611f3857611f38611f00565b60405160a081016001600160401b0381118282101715611f3857611f38611f00565b604051606081016001600160401b0381118282101715611f3857611f38611f00565b604051601f8201601f191681016001600160401b038111828210171561201157612011611f00565b604052919050565b80356004811061202857600080fd5b919050565b60006001600160401b0382111561204657612046611f00565b5060051b60200190565b60006040828403121561206257600080fd5b61206a611f16565b905081356007811061207b57600080fd5b808252506020820135602082015292915050565b600060408083850312156120a257600080fd5b6120aa611f16565b915082356001600160401b03808211156120c357600080fd5b818501915060208083880312156120d957600080fd5b6120e1611f3e565b8335838111156120f057600080fd5b80850194505087601f85011261210557600080fd5b8335925061211a6121158461202d565b611fe9565b83815260069390931b8401820192828101908985111561213957600080fd5b948301945b8486101561215f576121508a87612050565b8252948601949083019061213e565b8252508552948501359484019490945250909392505050565b803563ffffffff8116811461202857600080fd5b6000604080838503121561219f57600080fd5b6121a7611f16565b915082356001600160401b038111156121bf57600080fd5b8301601f810185136121d057600080fd5b803560206121e06121158361202d565b82815260a092830284018201928282019190898511156121ff57600080fd5b948301945b848610156122685780868b03121561221c5760008081fd5b612224611f60565b61222e8b88612050565b815287870135858201526060612245818901612178565b8983015261225560808901612178565b9082015283529485019491830191612204565b50808752505080860135818601525050505092915050565b6000610120823603121561229357600080fd5b61229b611f82565b6122a483612019565b815260208301356001600160401b03808211156122c057600080fd5b6122cc3683870161208f565b602084015260408501359150808211156122e557600080fd5b6122f13683870161208f565b6040840152606085013591508082111561230a57600080fd5b506123173682860161218c565b6060830152506080830135608082015261233360a08401612178565b60a082015261234460c08401612178565b60c082015261235560e08401612178565b60e082015261010092830135928101929092525090565b80356001600160401b038116811461202857600080fd5b600081830360e081121561239657600080fd5b61239e611fa5565b833581526060601f19830112156123b457600080fd5b6123bc611fc7565b91506123ca6020850161236c565b82526123d86040850161236c565b6020830152606084013560408301528160208201526080840135604082015260a0840135606082015261240d60c08501612178565b6080820152949350505050565b60006020828403121561242c57600080fd5b813561ffff811681146116d457600080fd5b6020808252600d908201526c4241445f43414c4c5f4441544160981b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff80831681851680830382111561249a5761249a612465565b01949350505050565b6020808252601190820152700848288beae929c889eaebe988a9c8ea89607b1b604082015260600190565b634e487b7160e01b600052603260045260246000fd5b60006000198214156124f8576124f8612465565b5060010190565b600060ff821660ff81141561251657612516612465565b60010192915050565b60208082526011908201527015d493d391d7d3515492d31157d493d3d5607a1b604082015260600190565b6000821982111561255d5761255d612465565b500190565b60008282101561257457612574612465565b500390565b6000845160005b8181101561259a5760208188018101518583015201612580565b818111156125a9576000828501525b5091909101928352506020820152604001919050565b652b30b63ab29d60d11b81526000600784106125dd576125dd611ca9565b5060f89290921b6006830152600782015260270190565b634e487b7160e01b600052605160045260246000fdfea2646970667358221220b902a6388a562f04c941d5ce1c3185c91f9ee83e272d461c79a1235e019d1fb464736f6c63430008090033",
}

// OneStepProver0ABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProver0MetaData.ABI instead.
var OneStepProver0ABI = OneStepProver0MetaData.ABI

// OneStepProver0Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProver0MetaData.Bin instead.
var OneStepProver0Bin = OneStepProver0MetaData.Bin

// DeployOneStepProver0 deploys a new Ethereum contract, binding an instance of OneStepProver0 to it.
func DeployOneStepProver0(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProver0, error) {
	parsed, err := OneStepProver0MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProver0Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProver0{OneStepProver0Caller: OneStepProver0Caller{contract: contract}, OneStepProver0Transactor: OneStepProver0Transactor{contract: contract}, OneStepProver0Filterer: OneStepProver0Filterer{contract: contract}}, nil
}

// OneStepProver0 is an auto generated Go binding around an Ethereum contract.
type OneStepProver0 struct {
	OneStepProver0Caller     // Read-only binding to the contract
	OneStepProver0Transactor // Write-only binding to the contract
	OneStepProver0Filterer   // Log filterer for contract events
}

// OneStepProver0Caller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProver0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProver0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProver0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProver0Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProver0Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProver0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProver0Session struct {
	Contract     *OneStepProver0   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneStepProver0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProver0CallerSession struct {
	Contract *OneStepProver0Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OneStepProver0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProver0TransactorSession struct {
	Contract     *OneStepProver0Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OneStepProver0Raw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProver0Raw struct {
	Contract *OneStepProver0 // Generic contract binding to access the raw methods on
}

// OneStepProver0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProver0CallerRaw struct {
	Contract *OneStepProver0Caller // Generic read-only contract binding to access the raw methods on
}

// OneStepProver0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProver0TransactorRaw struct {
	Contract *OneStepProver0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProver0 creates a new instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0(address common.Address, backend bind.ContractBackend) (*OneStepProver0, error) {
	contract, err := bindOneStepProver0(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0{OneStepProver0Caller: OneStepProver0Caller{contract: contract}, OneStepProver0Transactor: OneStepProver0Transactor{contract: contract}, OneStepProver0Filterer: OneStepProver0Filterer{contract: contract}}, nil
}

// NewOneStepProver0Caller creates a new read-only instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0Caller(address common.Address, caller bind.ContractCaller) (*OneStepProver0Caller, error) {
	contract, err := bindOneStepProver0(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0Caller{contract: contract}, nil
}

// NewOneStepProver0Transactor creates a new write-only instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0Transactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProver0Transactor, error) {
	contract, err := bindOneStepProver0(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0Transactor{contract: contract}, nil
}

// NewOneStepProver0Filterer creates a new log filterer instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0Filterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProver0Filterer, error) {
	contract, err := bindOneStepProver0(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0Filterer{contract: contract}, nil
}

// bindOneStepProver0 binds a generic wrapper to an already deployed contract.
func bindOneStepProver0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProver0MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProver0 *OneStepProver0Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProver0.Contract.OneStepProver0Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProver0 *OneStepProver0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProver0.Contract.OneStepProver0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProver0 *OneStepProver0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProver0.Contract.OneStepProver0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProver0 *OneStepProver0CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProver0.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProver0 *OneStepProver0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProver0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProver0 *OneStepProver0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProver0.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xda78e7d1.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0Caller) ExecuteOneStep(opts *bind.CallOpts, arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProver0.contract.Call(opts, &out, "executeOneStep", arg0, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xda78e7d1.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0Session) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProver0.Contract.ExecuteOneStep(&_OneStepProver0.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xda78e7d1.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0CallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProver0.Contract.ExecuteOneStep(&_OneStepProver0.CallOpts, arg0, startMach, startMod, inst, proof)
}

// OneStepProverHostIoMetaData contains all meta data concerning the OneStepProverHostIo contract.
var OneStepProverHostIoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612c97806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063da78e7d114610030575b600080fd5b61004361003e3660046120aa565b61005a565b6040516100519291906122bc565b60405180910390f35b610062611f9c565b6040805160a0810182526000808252825160608082018552828252602080830184905282860184905284019190915292820181905291810182905260808101919091526100ae87612764565b91506100bf36879003870187612867565b905060006100d060208701876128fe565b905061204561801061ffff8316108015906100f1575061801361ffff831611155b156100ff57506101a8610189565b61ffff8216618020141561011657506102f0610189565b61ffff8216618021141561012d57506109ad610189565b61ffff821661802214156101445750610cd2610189565b60405162461bcd60e51b8152602060048201526015602482015274494e56414c49445f4d454d4f52595f4f50434f444560581b60448201526064015b60405180910390fd5b61019b8a85858a8a8a8763ffffffff16565b5050965096945050505050565b60006101b760208501856128fe565b90506101c161204f565b60006101ce858583610cde565b60808a015191935091506101e183610db9565b146102215760405162461bcd60e51b815260206004820152601060248201526f4241445f474c4f42414c5f535441544560801b6044820152606401610180565b61ffff8316618010148061023a575061ffff8316618011145b1561025c57610257888884896102528987818d612922565b610e2d565b6102d4565b61ffff83166180121415610274576102578883610fd5565b61ffff8316618013141561028c57610257888361106c565b60405162461bcd60e51b815260206004820152601a60248201527f494e56414c49445f474c4f42414c53544154455f4f50434f44450000000000006044820152606401610180565b6102dd82610db9565b6080909801979097525050505050505050565b600061030761030287602001516110e2565b611107565b63ffffffff169050600061032161030288602001516110e2565b63ffffffff169050610334602083612962565b15158061035a5750602080870151516001600160401b03169061035890839061298c565b115b8061036e575061036b602082612962565b15155b15610395578660025b9081600381111561038a5761038a61218d565b8152505050506109a5565b60006103a26020836129a4565b90506000806103bd6040518060200160405280606081525090565b60208a01516103cf90858a8a87611198565b9094509092509050606060008989868181106103ed576103ed6129b8565b919091013560f81c9150859050610403816129ce565b95505060208b01356104de5760ff81166104c6573660006104268b88818f612922565b9150915085828260405161043b9291906129e9565b6040518091039020146104605760405162461bcd60e51b8152600401610180906129f9565b600061046d8b602061298c565b90508181111561047a5750805b610486818c8486612922565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092975061092995505050505050565b60405162461bcd60e51b815260040161018090612a1f565b8a602001356001141561058f5760ff81161561050c5760405162461bcd60e51b815260040161018090612a1f565b36600061051b8b88818f612922565b9150915085600283836040516105329291906129e9565b602060405180830381855afa15801561054f573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906105729190612a4f565b146104605760405162461bcd60e51b8152600401610180906129f9565b8a60200135600214156108e95760ff8116156105bd5760405162461bcd60e51b815260040161018090612a1f565b3660006105cc8b88818f612922565b9092509050856105e0602060008486612922565b6105e991612a68565b1461062d5760405162461bcd60e51b8152602060048201526014602482015273096b48ebea0a49e9e8cbeaea49e9c8ebe9082a6960631b6044820152606401610180565b600080600080600a6001600160a01b0316868660405161064e9291906129e9565b600060405180830381855afa9150503d8060008114610689576040519150601f19603f3d011682016040523d82523d6000602084013e61068e565b606091505b5091509150816106d45760405162461bcd60e51b815260206004820152601160248201527024a72b20a624a22fa5ad23afa82927a7a360791b6044820152606401610180565b600081511161071e5760405162461bcd60e51b81526020600482015260166024820152754b5a475f505245434f4d50494c455f4d495353494e4760501b6044820152606401610180565b808060200190518101906107329190612a86565b9094509250507f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff000000018214905061079f5760405162461bcd60e51b8152602060048201526013602482015272554e4b4e4f574e5f424c535f4d4f44554c555360681b6044820152606401610180565b6107aa826020612aaa565b8c10156108e0576000806107bf60208f6129a4565b905060015b848110156107ee57600192831b9282811614156107e2576001831792505b600191821c911b6107c4565b506000610800856401000000006129a4565b905061080c8382612aaa565b9050600061083b7f16a2a19edfe81f20d09b681922c813b4b63683508c2280b93829971f439f0d2b8387611232565b90508061084c604060208a8c612922565b61085591612a68565b146108965760405162461bcd60e51b815260206004820152601160248201527025ad23afa82927a7a32faba927a723afad60791b6044820152606401610180565b6108a460606040898b612922565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929c50505050505050505b50505050610929565b60405162461bcd60e51b8152602060048201526015602482015274554e4b4e4f574e5f505245494d4147455f5459504560581b6044820152606401610180565b60005b825181101561096d57610959858285848151811061094c5761094c6129b8565b016020015160f81c611366565b945080610965816129ce565b91505061092c565b506109798387866113eb565b60208d015160400152815161099c906109919061146a565b60208f01519061149d565b50505050505050505b505050505050565b60006109bf61030287602001516110e2565b63ffffffff16905060006109d961030288602001516110e2565b63ffffffff16905060006109f86109f389602001516110e2565b6114ad565b6001600160401b031690506020860135158015610a16575088358110155b15610a3e578760035b90816003811115610a3257610a3261218d565b815250505050506109a5565b602080880151516001600160401b031690610a5a90849061298c565b1180610a6f5750610a6c602083612962565b15155b15610a7c57876002610a1f565b6000610a896020846129a4565b9050600080610aa46040518060200160405280606081525090565b60208b0151610ab690858b8b87611198565b9094509092509050888884818110610ad057610ad06129b8565b909101356001600160f81b031916159050610b235760405162461bcd60e51b81526020600482015260136024820152722aa725a727aba72fa4a72127ac2fa82927a7a360691b6044820152606401610180565b82610b2d816129ce565b9350506120456000808c602001351415610b4b5761153e9150610b8b565b60018c602001351415610b625761183d9150610b8b565b8d60025b90816003811115610b7957610b7961218d565b815250505050505050505050506109a5565b610bab8f888d8d89908092610ba293929190612922565b8663ffffffff16565b905080610bba578d6002610b66565b505082881015610c005760405162461bcd60e51b81526020600482015260116024820152702120a22fa6a2a9a9a0a3a2afa82927a7a360791b6044820152606401610180565b6000610c0c848a612ac9565b905060005b60208163ffffffff16108015610c35575081610c3363ffffffff83168b61298c565b105b15610c8e57610c7a8463ffffffff83168d8d82610c528f8c61298c565b610c5c919061298c565b818110610c6b57610c6b6129b8565b919091013560f81c9050611366565b935080610c8681612ae0565b915050610c11565b610c998387866113eb565b60208e015160400152610cc1610cae8261146a565b8f6020015161149d90919063ffffffff16565b505050505050505050505050505050565b50506001909252505050565b610ce661204f565b81610cef612074565b610cf7612074565b60005b600260ff82161015610d4257610d11888886611adf565b848360ff1660028110610d2657610d266129b8565b6020020191909152935080610d3a81612b04565b915050610cfa565b5060005b600260ff82161015610d9c57610d5d888886611afb565b838360ff1660028110610d7257610d726129b8565b6001600160401b039093166020939093020191909152935080610d9481612b04565b915050610d46565b506040805180820190915291825260208201529590945092505050565b8051805160209182015192820151805190830151604080516c23b637b130b61039ba30ba329d60991b81870152602d810194909452604d8401959095526001600160c01b031960c092831b8116606d850152911b1660758201528251808203605d018152607d909101909252815191012090565b6000610e3f61030288602001516110e2565b63ffffffff1690506000610e5961030289602001516110e2565b9050600263ffffffff821610610e7157876002610377565b602080880151516001600160401b031690610e8d90849061298c565b1180610ea25750610e9f602083612962565b15155b15610eaf57876002610377565b6000610ebc6020846129a4565b9050600080610ed76040518060200160405280606081525090565b60208b0151610ee990858a8a87611198565b9094509092509050618010610f0160208b018b6128fe565b61ffff161415610f4657610f38848b600001518763ffffffff1660028110610f2b57610f2b6129b8565b60200201518391906113eb565b60208c015160400152610fc7565b618011610f5660208b018b6128fe565b61ffff161415610f85578951829063ffffffff871660028110610f7b57610f7b6129b8565b6020020152610fc7565b60405162461bcd60e51b81526020600482015260176024820152764241445f474c4f42414c5f53544154455f4f50434f444560481b6044820152606401610180565b505050505050505050505050565b6000610fe761030284602001516110e2565b9050600263ffffffff821610610fff57505060029052565b61106761105c83602001518363ffffffff1660028110611021576110216129b8565b602002015160408051808201909152600080825260208201525060408051808201909152600181526001600160401b03909116602082015290565b60208501519061149d565b505050565b600061107e6109f384602001516110e2565b9050600061109261030285602001516110e2565b9050600263ffffffff8216106110ac575050600290915250565b8183602001518263ffffffff16600281106110c9576110c96129b8565b6001600160401b03909216602092909202015250505050565b6040805180820190915260008082526020820152815161110190611b62565b92915050565b602081015160009081835160068111156111235761112361218d565b1461115a5760405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b6044820152606401610180565b64010000000081106111015760405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b6044820152606401610180565b6000806111b16040518060200160405280606081525090565b8391506111bf868684611adf565b90935091506111cf868684611c72565b9250905060006111e08289866113eb565b9050886040015181146112265760405162461bcd60e51b815260206004820152600e60248201526d15d493d391d7d3515357d493d3d560921b6044820152606401610180565b50955095509592505050565b60408051602080820181905281830181905260608201526080810185905260a0810184905260c08082018490528251808303909101815260e09091019182905260009182908190600590611287908590612b54565b600060405180830381855afa9150503d80600081146112c2576040519150601f19603f3d011682016040523d82523d6000602084013e6112c7565b606091505b5091509150816113095760405162461bcd60e51b815260206004820152600d60248201526c1353d111561417d19052531151609a1b6044820152606401610180565b80516020146113505760405162461bcd60e51b815260206004820152601360248201527209a9e888ab0a0beaea49e9c8ebe988a9c8ea89606b1b6044820152606401610180565b61135981612b70565b93505050505b9392505050565b6000602083106113b05760405162461bcd60e51b81526020600482015260156024820152740848288bea68aa8be988a828cbe84b2a88abe9288b605b1b6044820152606401610180565b6000836113bf60016020612ac9565b6113c99190612ac9565b6113d4906008612aaa565b60ff848116821b911b198616179150509392505050565b6040516b26b2b6b7b93c903632b0b31d60a11b6020820152602c81018290526000908190604c016040516020818303038152906040528051906020012090506114618585836040518060400160405280601381526020017226b2b6b7b93c9036b2b935b632903a3932b29d60691b815250611d4c565b95945050505050565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b81516114a99082611e1e565b5050565b60208101516000906001835160068111156114ca576114ca61218d565b146115015760405162461bcd60e51b81526020600482015260076024820152661393d517d24d8d60ca1b6044820152606401610180565b600160401b81106111015760405162461bcd60e51b815260206004820152600760248201526610905117d24d8d60ca1b6044820152606401610180565b600060288210156115865760405162461bcd60e51b81526020600482015260126024820152712120a22fa9a2a8a4a72127ac2fa82927a7a360711b6044820152606401610180565b600061159484846020611afb565b5080915050600084846040516115ab9291906129e9565b60405190819003902090506000806001600160401b0388161561166b576115d860408a0160208b01612b94565b6001600160a01b03166316bf55796115f160018b612bbd565b6040516001600160e01b031960e084901b1681526001600160401b03909116600482015260240160206040518083038186803b15801561163057600080fd5b505afa158015611644573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116689190612a4f565b91505b6001600160401b0384161561171d5761168a60408a0160208b01612b94565b6001600160a01b031663d5719dc26116a3600187612bbd565b6040516001600160e01b031960e084901b1681526001600160401b03909116600482015260240160206040518083038186803b1580156116e257600080fd5b505afa1580156116f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061171a9190612a4f565b90505b6040805160208101849052908101849052606081018290526000906080016040516020818303038152906040528051906020012090508960200160208101906117669190612b94565b6040516316bf557960e01b81526001600160401b038b1660048201526001600160a01b0391909116906316bf55799060240160206040518083038186803b1580156117b057600080fd5b505afa1580156117c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117e89190612a4f565b811461182d5760405162461bcd60e51b81526020600482015260146024820152734241445f534551494e424f585f4d45535341474560601b6044820152606401610180565b5060019998505050505050505050565b600060718210156118845760405162461bcd60e51b81526020600482015260116024820152702120a22fa222a620aca2a22fa82927a7a360791b6044820152606401610180565b60006001600160401b03851615611938576118a56040870160208801612b94565b6001600160a01b031663d5719dc26118be600188612bbd565b6040516001600160e01b031960e084901b1681526001600160401b03909116600482015260240160206040518083038186803b1580156118fd57600080fd5b505afa158015611911573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119359190612a4f565b90505b60006119478460718188612922565b6040516119559291906129e9565b60405180910390209050600085856000818110611974576119746129b8565b9050013560f81c60f81b9050600061198e87876001611f11565b509050600082826119a3607160218b8d612922565b876040516020016119b8959493929190612be5565b60408051601f19818403018152828252805160209182012083820189905283830181905282518085038401815260609094019092528251920191909120909150611a0860408c0160208d01612b94565b604051636ab8cee160e11b81526001600160401b038c1660048201526001600160a01b03919091169063d5719dc29060240160206040518083038186803b158015611a5257600080fd5b505afa158015611a66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a8a9190612a4f565b8114611ace5760405162461bcd60e51b81526020600482015260136024820152724241445f44454c415945445f4d45535341474560681b6044820152606401610180565b5060019a9950505050505050505050565b60008181611aee868684611f11565b9097909650945050505050565b600081815b6008811015611b59576008836001600160401b0316901b9250858583818110611b2b57611b2b6129b8565b919091013560f81c93909317925081611b43816129ce565b9250508080611b51906129ce565b915050611b00565b50935093915050565b604080518082019091526000808252602082015281518051611b8690600190612ac9565b81518110611b9657611b966129b8565b6020026020010151905060006001836000015151611bb49190612ac9565b6001600160401b03811115611bcb57611bcb6123e4565b604051908082528060200260200182016040528015611c1057816020015b6040805180820190915260008082526020820152815260200190600190039081611be95790505b50905060005b8151811015611c6b578351805182908110611c3357611c336129b8565b6020026020010151828281518110611c4d57611c4d6129b8565b60200260200101819052508080611c63906129ce565b915050611c16565b5090915290565b604080516020810190915260608152816000611c8f868684611f66565b92509050600060ff82166001600160401b03811115611cb057611cb06123e4565b604051908082528060200260200182016040528015611cd9578160200160208202803683370190505b50905060005b8260ff168160ff161015611d3057611cf8888886611adf565b838360ff1681518110611d0d57611d0d6129b8565b602002602001018196508281525050508080611d2890612b04565b915050611cdf565b5060405180602001604052808281525093505050935093915050565b8160005b855151811015611e155760018516611db157828287600001518381518110611d7a57611d7a6129b8565b6020026020010151604051602001611d9493929190612c24565b604051602081830303815290604052805190602001209150611dfc565b8286600001518281518110611dc857611dc86129b8565b602002602001015183604051602001611de393929190612c24565b6040516020818303038152906040528051906020012091505b60019490941c9380611e0d816129ce565b915050611d50565b50949350505050565b815151600090611e2f90600161298c565b6001600160401b03811115611e4657611e466123e4565b604051908082528060200260200182016040528015611e8b57816020015b6040805180820190915260008082526020820152815260200190600190039081611e645790505b50905060005b835151811015611ee7578351805182908110611eaf57611eaf6129b8565b6020026020010151828281518110611ec957611ec96129b8565b60200260200101819052508080611edf906129ce565b915050611e91565b50818184600001515181518110611f0057611f006129b8565b602090810291909101015290915250565b600081815b6020811015611b5957600883901b9250858583818110611f3857611f386129b8565b919091013560f81c93909317925081611f50816129ce565b9250508080611f5e906129ce565b915050611f16565b600081848482818110611f7b57611f7b6129b8565b919091013560f81c9250819050611f91816129ce565b915050935093915050565b6040805161012081019091528060008152602001611fd160408051606080820183529181019182529081526000602082015290565b8152602001611ff760408051606080820183529181019182529081526000602082015290565b815260200161201c604051806040016040528060608152602001600080191681525090565b815260006020820181905260408201819052606082018190526080820181905260a09091015290565b61204d612c4b565b565b6040518060400160405280612062612074565b815260200161206f612074565b905290565b60405180604001604052806002906020820280368337509192915050565b6000604082840312156120a457600080fd5b50919050565b6000806000806000808688036101a08112156120c557600080fd5b6120cf8989612092565b965060408801356001600160401b03808211156120eb57600080fd5b90890190610120828c03121561210057600080fd5b81975060e0605f198401121561211557600080fd5b60608a0196506121298b6101408c01612092565b95506101808a013592508083111561214057600080fd5b828a0192508a601f84011261215457600080fd5b823591508082111561216557600080fd5b5089602082840101111561217857600080fd5b60208201935080925050509295509295509295565b634e487b7160e01b600052602160045260246000fd5b600481106121b3576121b361218d565b9052565b8051600781106121c9576121c961218d565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b80841015612226576122128286516121b7565b9382019360019390930192908501906121ff565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b828110156122a55784516122718582516121b7565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a09093019260010161225c565b509687015197909601969096525093949350505050565b60006101008083526122d181840186516121a3565b6020850151610120848101526122eb6102208501826121d6565b9050604086015160ff19808684030161014087015261230a83836121d6565b925060608801519150808684030161016087015250612329828261223a565b915050608086015161018085015260a086015161234f6101a086018263ffffffff169052565b5060c086015163ffffffff81166101c08601525060e086015163ffffffff81166101e08601525090850151610200840152905061135f60208301848051825260208101516001600160401b0380825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a083015263ffffffff60808201511660c08301525050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561241c5761241c6123e4565b60405290565b604051602081016001600160401b038111828210171561241c5761241c6123e4565b604051608081016001600160401b038111828210171561241c5761241c6123e4565b60405161012081016001600160401b038111828210171561241c5761241c6123e4565b60405160a081016001600160401b038111828210171561241c5761241c6123e4565b604051606081016001600160401b038111828210171561241c5761241c6123e4565b604051601f8201601f191681016001600160401b03811182821017156124f5576124f56123e4565b604052919050565b80356004811061250c57600080fd5b919050565b60006001600160401b0382111561252a5761252a6123e4565b5060051b60200190565b60006040828403121561254657600080fd5b61254e6123fa565b905081356007811061255f57600080fd5b808252506020820135602082015292915050565b6000604080838503121561258657600080fd5b61258e6123fa565b915082356001600160401b03808211156125a757600080fd5b818501915060208083880312156125bd57600080fd5b6125c5612422565b8335838111156125d457600080fd5b80850194505087601f8501126125e957600080fd5b833592506125fe6125f984612511565b6124cd565b83815260069390931b8401820192828101908985111561261d57600080fd5b948301945b84861015612643576126348a87612534565b82529486019490830190612622565b8252508552948501359484019490945250909392505050565b803563ffffffff8116811461250c57600080fd5b6000604080838503121561268357600080fd5b61268b6123fa565b915082356001600160401b038111156126a357600080fd5b8301601f810185136126b457600080fd5b803560206126c46125f983612511565b82815260a092830284018201928282019190898511156126e357600080fd5b948301945b8486101561274c5780868b0312156127005760008081fd5b612708612444565b6127128b88612534565b81528787013585820152606061272981890161265c565b898301526127396080890161265c565b90820152835294850194918301916126e8565b50808752505080860135818601525050505092915050565b6000610120823603121561277757600080fd5b61277f612466565b612788836124fd565b815260208301356001600160401b03808211156127a457600080fd5b6127b036838701612573565b602084015260408501359150808211156127c957600080fd5b6127d536838701612573565b604084015260608501359150808211156127ee57600080fd5b506127fb36828601612670565b6060830152506080830135608082015261281760a0840161265c565b60a082015261282860c0840161265c565b60c082015261283960e0840161265c565b60e082015261010092830135928101929092525090565b80356001600160401b038116811461250c57600080fd5b600081830360e081121561287a57600080fd5b612882612489565b833581526060601f198301121561289857600080fd5b6128a06124ab565b91506128ae60208501612850565b82526128bc60408501612850565b6020830152606084013560408301528160208201526080840135604082015260a084013560608201526128f160c0850161265c565b6080820152949350505050565b60006020828403121561291057600080fd5b813561ffff8116811461135f57600080fd5b6000808585111561293257600080fd5b8386111561293f57600080fd5b5050820193919092039150565b634e487b7160e01b600052601260045260246000fd5b6000826129715761297161294c565b500690565b634e487b7160e01b600052601160045260246000fd5b6000821982111561299f5761299f612976565b500190565b6000826129b3576129b361294c565b500490565b634e487b7160e01b600052603260045260246000fd5b60006000198214156129e2576129e2612976565b5060010190565b8183823760009101908152919050565b6020808252600c908201526b4241445f505245494d41474560a01b604082015260600190565b6020808252601690820152752aa725a727aba72fa82922a4a6a0a3a2afa82927a7a360511b604082015260600190565b600060208284031215612a6157600080fd5b5051919050565b8035602083101561110157600019602084900360031b1b1692915050565b60008060408385031215612a9957600080fd5b505080516020909101519092909150565b6000816000190483118215151615612ac457612ac4612976565b500290565b600082821015612adb57612adb612976565b500390565b600063ffffffff80831681811415612afa57612afa612976565b6001019392505050565b600060ff821660ff811415612b1b57612b1b612976565b60010192915050565b60005b83811015612b3f578181015183820152602001612b27565b83811115612b4e576000848401525b50505050565b60008251612b66818460208701612b24565b9190910192915050565b805160208083015191908110156120a45760001960209190910360031b1b16919050565b600060208284031215612ba657600080fd5b81356001600160a01b038116811461135f57600080fd5b60006001600160401b0383811690831681811015612bdd57612bdd612976565b039392505050565b6001600160f81b031986168152606085901b6bffffffffffffffffffffffff191660018201528284601583013760159201918201526035019392505050565b60008451612c36818460208901612b24565b91909101928352506020820152604001919050565b634e487b7160e01b600052605160045260246000fdfea264697066735822122020f5be13dfffca96166bd579df0591b2568f6771d8323485ecf0a3f786c2195464736f6c63430008090033",
}

// OneStepProverHostIoABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProverHostIoMetaData.ABI instead.
var OneStepProverHostIoABI = OneStepProverHostIoMetaData.ABI

// OneStepProverHostIoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProverHostIoMetaData.Bin instead.
var OneStepProverHostIoBin = OneStepProverHostIoMetaData.Bin

// DeployOneStepProverHostIo deploys a new Ethereum contract, binding an instance of OneStepProverHostIo to it.
func DeployOneStepProverHostIo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProverHostIo, error) {
	parsed, err := OneStepProverHostIoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProverHostIoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProverHostIo{OneStepProverHostIoCaller: OneStepProverHostIoCaller{contract: contract}, OneStepProverHostIoTransactor: OneStepProverHostIoTransactor{contract: contract}, OneStepProverHostIoFilterer: OneStepProverHostIoFilterer{contract: contract}}, nil
}

// OneStepProverHostIo is an auto generated Go binding around an Ethereum contract.
type OneStepProverHostIo struct {
	OneStepProverHostIoCaller     // Read-only binding to the contract
	OneStepProverHostIoTransactor // Write-only binding to the contract
	OneStepProverHostIoFilterer   // Log filterer for contract events
}

// OneStepProverHostIoCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProverHostIoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverHostIoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProverHostIoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverHostIoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProverHostIoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverHostIoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProverHostIoSession struct {
	Contract     *OneStepProverHostIo // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OneStepProverHostIoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProverHostIoCallerSession struct {
	Contract *OneStepProverHostIoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OneStepProverHostIoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProverHostIoTransactorSession struct {
	Contract     *OneStepProverHostIoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OneStepProverHostIoRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProverHostIoRaw struct {
	Contract *OneStepProverHostIo // Generic contract binding to access the raw methods on
}

// OneStepProverHostIoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProverHostIoCallerRaw struct {
	Contract *OneStepProverHostIoCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProverHostIoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProverHostIoTransactorRaw struct {
	Contract *OneStepProverHostIoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProverHostIo creates a new instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIo(address common.Address, backend bind.ContractBackend) (*OneStepProverHostIo, error) {
	contract, err := bindOneStepProverHostIo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIo{OneStepProverHostIoCaller: OneStepProverHostIoCaller{contract: contract}, OneStepProverHostIoTransactor: OneStepProverHostIoTransactor{contract: contract}, OneStepProverHostIoFilterer: OneStepProverHostIoFilterer{contract: contract}}, nil
}

// NewOneStepProverHostIoCaller creates a new read-only instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIoCaller(address common.Address, caller bind.ContractCaller) (*OneStepProverHostIoCaller, error) {
	contract, err := bindOneStepProverHostIo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIoCaller{contract: contract}, nil
}

// NewOneStepProverHostIoTransactor creates a new write-only instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIoTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProverHostIoTransactor, error) {
	contract, err := bindOneStepProverHostIo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIoTransactor{contract: contract}, nil
}

// NewOneStepProverHostIoFilterer creates a new log filterer instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIoFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProverHostIoFilterer, error) {
	contract, err := bindOneStepProverHostIo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIoFilterer{contract: contract}, nil
}

// bindOneStepProverHostIo binds a generic wrapper to an already deployed contract.
func bindOneStepProverHostIo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProverHostIoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverHostIo *OneStepProverHostIoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverHostIo.Contract.OneStepProverHostIoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverHostIo *OneStepProverHostIoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.OneStepProverHostIoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverHostIo *OneStepProverHostIoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.OneStepProverHostIoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverHostIo *OneStepProverHostIoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverHostIo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverHostIo *OneStepProverHostIoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverHostIo *OneStepProverHostIoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xda78e7d1.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoCaller) ExecuteOneStep(opts *bind.CallOpts, execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProverHostIo.contract.Call(opts, &out, "executeOneStep", execCtx, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xda78e7d1.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoSession) ExecuteOneStep(execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverHostIo.Contract.ExecuteOneStep(&_OneStepProverHostIo.CallOpts, execCtx, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xda78e7d1.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoCallerSession) ExecuteOneStep(execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverHostIo.Contract.ExecuteOneStep(&_OneStepProverHostIo.CallOpts, execCtx, startMach, startMod, inst, proof)
}

// OneStepProverMathMetaData contains all meta data concerning the OneStepProverMath contract.
var OneStepProverMathMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612364806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063da78e7d114610030575b600080fd5b61004361003e3660046118ae565b61005a565b604051610051929190611ac0565b60405180910390f35b6100626117e3565b6040805160a0810182526000808252825160608082018552828252602080830184905282860184905284019190915292820181905291810182905260808101919091526100ae87611f63565b91506100bf36879003870187612066565b905060006100d060208701876120fd565b905061188c61ffff8216604514806100ec575061ffff82166050145b156100fa57506103096102eb565b604661ffff831610801590610122575061011660096046612137565b61ffff168261ffff1611155b1561013057506104226102eb565b606761ffff831610801590610158575061014c60026067612137565b61ffff168261ffff1611155b1561016657506105056102eb565b606a61ffff8316108015906101805750607861ffff831611155b1561018e575061056d6102eb565b605161ffff8316108015906101b657506101aa60096051612137565b61ffff168261ffff1611155b156101c4575061075a6102eb565b607961ffff8316108015906101ec57506101e060026079612137565b61ffff168261ffff1611155b156101fa57506107bf6102eb565b607c61ffff8316108015906102145750608a61ffff831611155b1561022257506108126102eb565b61ffff821660a7141561023857506109cd6102eb565b61ffff821660ac148061024f575061ffff821660ad145b1561025d57506109ee6102eb565b60c061ffff831610801590610277575060c461ffff831611155b156102855750610a426102eb565b60bc61ffff83161080159061029f575060bf61ffff831611155b156102ad5750610c576102eb565b60405162461bcd60e51b815260206004820152600e60248201526d494e56414c49445f4f50434f444560901b60448201526064015b60405180910390fd5b6102fc84848989898663ffffffff16565b5050965096945050505050565b60006103188660200151610de2565b9050604561032960208601866120fd565b61ffff16141561036a5760008151600681111561034857610348611991565b146103655760405162461bcd60e51b81526004016102e29061215d565b6103e7565b605061037960208601866120fd565b61ffff1614156103b55760018151600681111561039857610398611991565b146103655760405162461bcd60e51b81526004016102e29061217e565b60405162461bcd60e51b81526020600482015260076024820152662120a22fa2a8ad60c91b60448201526064016102e2565b60008160200151600014156103fe57506001610402565b5060005b61041961040e82610e07565b602089015190610e3a565b50505050505050565b60006104396104348760200151610de2565b610e4a565b9050600061044d6104348860200151610de2565b90506000604661046060208801886120fd565b61046a919061219f565b905060008061ffff831660021480610486575061ffff83166004145b80610495575061ffff83166006145b806104a4575061ffff83166008145b156104c4576104b284610ec1565b91506104bd85610ec1565b90506104d2565b505063ffffffff8083169084165b60006104df838386610eed565b90506104f86104ed82611096565b60208d015190610e3a565b5050505050505050505050565b60006105176104348760200151610de2565b90506000606761052a60208701876120fd565b610534919061219f565b9050600061054a8363ffffffff168360206110c9565b905061056361055882610e07565b60208a015190610e3a565b5050505050505050565b600061057f6104348760200151610de2565b905060006105936104348860200151610de2565b9050600080606a6105a760208901896120fd565b6105b1919061219f565b90508061ffff166003141561062f5763ffffffff841615806105e957508260030b637fffffff191480156105e957508360030b600019145b15610612578860025b9081600381111561060557610605611991565b8152505050505050610753565b8360030b8360030b81610627576106276121c2565b059150610737565b8061ffff166005141561066c5763ffffffff841661064f578860026105f2565b8360030b8360030b81610664576106646121c2565b079150610737565b8061ffff16600a141561068c5763ffffffff8316601f85161b9150610737565b8061ffff16600c14156106ac5763ffffffff8316601f85161c9150610737565b8061ffff16600b14156106ca57600383900b601f85161d9150610737565b8061ffff16600d14156106e8576106e1838561128d565b9150610737565b8061ffff16600e14156106ff576106e183856112cf565b6000806107198563ffffffff168763ffffffff1685611311565b915091508015610733575050600289525061075392505050565b5091505b61074e61074383610e07565b60208b015190610e3a565b505050505b5050505050565b600061077161076c8760200151610de2565b611497565b9050600061078561076c8860200151610de2565b90506000605161079860208801886120fd565b6107a2919061219f565b905060006107b1838584610eed565b905061074e61074382611096565b60006107d161076c8760200151610de2565b9050600060796107e460208701876120fd565b6107ee919061219f565b905060006107fe838360406110c9565b63ffffffff1690506105636105588261150e565b600061082461076c8760200151610de2565b9050600061083861076c8860200151610de2565b9050600080607c61084c60208901896120fd565b610856919061219f565b90508061ffff16600314156108bf576001600160401b038416158061089557508260070b677fffffffffffffff1914801561089557508360070b600019145b156108a2578860026105f2565b8360070b8360070b816108b7576108b76121c2565b0591506109c1565b8061ffff16600514156108ff576001600160401b0384166108e2578860026105f2565b8360070b8360070b816108f7576108f76121c2565b0791506109c1565b8061ffff16600a1415610922576001600160401b038316603f85161b91506109c1565b8061ffff16600c1415610945576001600160401b038316603f85161c91506109c1565b8061ffff16600b141561096357600783900b603f85161d91506109c1565b8061ffff16600d14156109815761097a8385611544565b91506109c1565b8061ffff16600e14156109985761097a8385611592565b60006109a5848684611311565b909350905080156109bf5750506002885250610753915050565b505b61074e6107438361150e565b60006109df61076c8760200151610de2565b90508061041961040e82610e07565b6000610a006104348760200151610de2565b9050600060ac610a1360208701876120fd565b61ffff161415610a2d57610a2682610ec1565b9050610a36565b5063ffffffff81165b61041961040e8261150e565b60008060c0610a5460208701876120fd565b61ffff161415610a6a5750600090506008610b41565b60c1610a7960208701876120fd565b61ffff161415610a8f5750600090506010610b41565b60c2610a9e60208701876120fd565b61ffff161415610ab45750600190506008610b41565b60c3610ac360208701876120fd565b61ffff161415610ad95750600190506010610b41565b60c4610ae860208701876120fd565b61ffff161415610afe5750600190506020610b41565b60405162461bcd60e51b8152602060048201526018602482015277494e56414c49445f455854454e445f53414d455f5459504560401b60448201526064016102e2565b600080836006811115610b5657610b56611991565b1415610b67575063ffffffff610b71565b506001600160401b035b6000610b808960200151610de2565b9050836006811115610b9457610b94611991565b81516006811115610ba757610ba7611991565b14610bf05760405162461bcd60e51b81526020600482015260196024820152784241445f455854454e445f53414d455f545950455f5459504560381b60448201526064016102e2565b6000610c03600160ff861681901b6121d8565b602083018051821690529050610c1a6001856121ef565b60ff166001901b826020015116600014610c3c57602082018051821985161790525b60208a0151610c4b9083610e3a565b50505050505050505050565b60008060bc610c6960208701876120fd565b61ffff161415610c7f5750600090506002610d2c565b60bd610c8e60208701876120fd565b61ffff161415610ca45750600190506003610d2c565b60be610cb360208701876120fd565b61ffff161415610cc95750600290506000610d2c565b60bf610cd860208701876120fd565b61ffff161415610cee5750600390506001610d2c565b60405162461bcd60e51b81526020600482015260136024820152721253959053125117d491525395115494149155606a1b60448201526064016102e2565b6000610d3b8860200151610de2565b9050816006811115610d4f57610d4f611991565b81516006811115610d6257610d62611991565b14610daa5760405162461bcd60e51b8152602060048201526018602482015277494e56414c49445f5245494e544552505245545f5459504560401b60448201526064016102e2565b80836006811115610dbd57610dbd611991565b90816006811115610dd057610dd0611991565b90525060208801516105639082610e3a565b60408051808201909152600080825260208201528151610e01906115e0565b92915050565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b8151610e4690826116f0565b5050565b60208101516000908183516006811115610e6657610e66611991565b14610e835760405162461bcd60e51b81526004016102e29061215d565b6401000000008110610e015760405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b60448201526064016102e2565b60006380000000821615610ee3575063ffffffff1667ffffffff000000001790565b5063ffffffff1690565b600061ffff8216610f1457826001600160401b0316846001600160401b031614905061108f565b61ffff821660011415610f3e57826001600160401b0316846001600160401b03161415905061108f565b61ffff821660021415610f5b578260070b8460070b12905061108f565b61ffff821660031415610f8457826001600160401b0316846001600160401b031610905061108f565b61ffff821660041415610fa1578260070b8460070b13905061108f565b61ffff821660051415610fca57826001600160401b0316846001600160401b031611905061108f565b61ffff821660061415610fe8578260070b8460070b1315905061108f565b61ffff82166007141561101257826001600160401b0316846001600160401b03161115905061108f565b61ffff821660081415611030578260070b8460070b1215905061108f565b61ffff82166009141561105a57826001600160401b0316846001600160401b03161015905061108f565b60405162461bcd60e51b815260206004820152600a6024820152690424144204952454c4f560b41b60448201526064016102e2565b9392505050565b604080518082019091526000808252602082015281156110ba57610e016001610e07565b610e016000610e07565b919050565b60008161ffff16602014806110e257508161ffff166040145b6111295760405162461bcd60e51b8152602060048201526018602482015277057524f4e4720555345204f462067656e65726963556e4f760441b60448201526064016102e2565b61ffff831661119a5761ffff82165b60008163ffffffff1611801561116d5750611154600182612212565b63ffffffff166001901b856001600160401b0316166000145b156111845761117d600182612212565b9050611138565b6111928161ffff8516612212565b91505061108f565b61ffff8316600114156111f35760005b8261ffff168163ffffffff161080156111d55750600163ffffffff82161b85166001600160401b0316155b156111ec576111e560018261222f565b90506111aa565b905061108f565b61ffff831660021415611259576000805b8361ffff168263ffffffff16101561125057600163ffffffff83161b86166001600160401b03161561123e5761123b60018261222f565b90505b816112488161224e565b925050611204565b915061108f9050565b60405162461bcd60e51b815260206004820152600960248201526804241442049556e4f760bc1b60448201526064016102e2565b600061129a602083612272565b91506112a7826020612212565b63ffffffff168363ffffffff16901c8263ffffffff168463ffffffff16901b17905092915050565b60006112dc602083612272565b91506112e9826020612212565b63ffffffff168363ffffffff16901b8263ffffffff168463ffffffff16901c17905092915050565b60008061ffff8316611329575050828201600061148f565b8261ffff1660011415611342575050818303600061148f565b8261ffff166002141561135b575050828202600061148f565b8261ffff16600414156113af576001600160401b038416611382575060009050600161148f565b836001600160401b0316856001600160401b0316816113a3576113a36121c2565b0460009150915061148f565b8261ffff1660061415611403576001600160401b0384166113d6575060009050600161148f565b836001600160401b0316856001600160401b0316816113f7576113f76121c2565b0660009150915061148f565b8261ffff166007141561141c575050828216600061148f565b8261ffff1660081415611435575050828217600061148f565b8261ffff166009141561144e575050828218600061148f565b60405162461bcd60e51b81526020600482015260166024820152750494e56414c49445f47454e455249435f42494e5f4f560541b60448201526064016102e2565b935093915050565b60208101516000906001835160068111156114b4576114b4611991565b146114d15760405162461bcd60e51b81526004016102e29061217e565b600160401b8110610e015760405162461bcd60e51b815260206004820152600760248201526610905117d24d8d60ca1b60448201526064016102e2565b60408051808201909152600080825260208201525060408051808201909152600181526001600160401b03909116602082015290565b6000611551604083612295565b915061155e8260406122af565b6001600160401b0316836001600160401b0316901c826001600160401b0316846001600160401b0316901b17905092915050565b600061159f604083612295565b91506115ac8260406122af565b6001600160401b0316836001600160401b0316901b826001600160401b0316846001600160401b0316901c17905092915050565b604080518082019091526000808252602082015281518051611604906001906121d8565b81518110611614576116146122cf565b602002602001015190506000600183600001515161163291906121d8565b6001600160401b0381111561164957611649611be8565b60405190808252806020026020018201604052801561168e57816020015b60408051808201909152600080825260208201528152602001906001900390816116675790505b50905060005b81518110156116e95783518051829081106116b1576116b16122cf565b60200260200101518282815181106116cb576116cb6122cf565b602002602001018190525080806116e1906122e5565b915050611694565b5090915290565b815151600090611701906001612300565b6001600160401b0381111561171857611718611be8565b60405190808252806020026020018201604052801561175d57816020015b60408051808201909152600080825260208201528152602001906001900390816117365790505b50905060005b8351518110156117b9578351805182908110611781576117816122cf565b602002602001015182828151811061179b5761179b6122cf565b602002602001018190525080806117b1906122e5565b915050611763565b508181846000015151815181106117d2576117d26122cf565b602090810291909101015290915250565b604080516101208101909152806000815260200161181860408051606080820183529181019182529081526000602082015290565b815260200161183e60408051606080820183529181019182529081526000602082015290565b8152602001611863604051806040016040528060608152602001600080191681525090565b815260006020820181905260408201819052606082018190526080820181905260a09091015290565b611894612318565b565b6000604082840312156118a857600080fd5b50919050565b6000806000806000808688036101a08112156118c957600080fd5b6118d38989611896565b965060408801356001600160401b03808211156118ef57600080fd5b90890190610120828c03121561190457600080fd5b81975060e0605f198401121561191957600080fd5b60608a01965061192d8b6101408c01611896565b95506101808a013592508083111561194457600080fd5b828a0192508a601f84011261195857600080fd5b823591508082111561196957600080fd5b5089602082840101111561197c57600080fd5b60208201935080925050509295509295509295565b634e487b7160e01b600052602160045260246000fd5b600481106119b7576119b7611991565b9052565b8051600781106119cd576119cd611991565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b80841015611a2a57611a168286516119bb565b938201936001939093019290850190611a03565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b82811015611aa9578451611a758582516119bb565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a090930192600101611a60565b509687015197909601969096525093949350505050565b6000610100808352611ad581840186516119a7565b602085015161012084810152611aef6102208501826119da565b9050604086015160ff198086840301610140870152611b0e83836119da565b925060608801519150808684030161016087015250611b2d8282611a3e565b915050608086015161018085015260a0860151611b536101a086018263ffffffff169052565b5060c086015163ffffffff81166101c08601525060e086015163ffffffff81166101e08601525090850151610200840152905061108f60208301848051825260208101516001600160401b0380825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a083015263ffffffff60808201511660c08301525050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715611c2057611c20611be8565b60405290565b604051602081016001600160401b0381118282101715611c2057611c20611be8565b604051608081016001600160401b0381118282101715611c2057611c20611be8565b60405161012081016001600160401b0381118282101715611c2057611c20611be8565b60405160a081016001600160401b0381118282101715611c2057611c20611be8565b604051606081016001600160401b0381118282101715611c2057611c20611be8565b604051601f8201601f191681016001600160401b0381118282101715611cf957611cf9611be8565b604052919050565b8035600481106110c457600080fd5b60006001600160401b03821115611d2957611d29611be8565b5060051b60200190565b600060408284031215611d4557600080fd5b611d4d611bfe565b9050813560078110611d5e57600080fd5b808252506020820135602082015292915050565b60006040808385031215611d8557600080fd5b611d8d611bfe565b915082356001600160401b0380821115611da657600080fd5b81850191506020808388031215611dbc57600080fd5b611dc4611c26565b833583811115611dd357600080fd5b80850194505087601f850112611de857600080fd5b83359250611dfd611df884611d10565b611cd1565b83815260069390931b84018201928281019089851115611e1c57600080fd5b948301945b84861015611e4257611e338a87611d33565b82529486019490830190611e21565b8252508552948501359484019490945250909392505050565b803563ffffffff811681146110c457600080fd5b60006040808385031215611e8257600080fd5b611e8a611bfe565b915082356001600160401b03811115611ea257600080fd5b8301601f81018513611eb357600080fd5b80356020611ec3611df883611d10565b82815260a09283028401820192828201919089851115611ee257600080fd5b948301945b84861015611f4b5780868b031215611eff5760008081fd5b611f07611c48565b611f118b88611d33565b815287870135858201526060611f28818901611e5b565b89830152611f3860808901611e5b565b9082015283529485019491830191611ee7565b50808752505080860135818601525050505092915050565b60006101208236031215611f7657600080fd5b611f7e611c6a565b611f8783611d01565b815260208301356001600160401b0380821115611fa357600080fd5b611faf36838701611d72565b60208401526040850135915080821115611fc857600080fd5b611fd436838701611d72565b60408401526060850135915080821115611fed57600080fd5b50611ffa36828601611e6f565b6060830152506080830135608082015261201660a08401611e5b565b60a082015261202760c08401611e5b565b60c082015261203860e08401611e5b565b60e082015261010092830135928101929092525090565b80356001600160401b03811681146110c457600080fd5b600081830360e081121561207957600080fd5b612081611c8d565b833581526060601f198301121561209757600080fd5b61209f611caf565b91506120ad6020850161204f565b82526120bb6040850161204f565b6020830152606084013560408301528160208201526080840135604082015260a084013560608201526120f060c08501611e5b565b6080820152949350505050565b60006020828403121561210f57600080fd5b813561ffff8116811461108f57600080fd5b634e487b7160e01b600052601160045260246000fd5b600061ffff80831681851680830382111561215457612154612121565b01949350505050565b6020808252600790820152662727aa2fa4999960c91b604082015260600190565b6020808252600790820152661393d517d24d8d60ca1b604082015260600190565b600061ffff838116908316818110156121ba576121ba612121565b039392505050565b634e487b7160e01b600052601260045260246000fd5b6000828210156121ea576121ea612121565b500390565b600060ff821660ff84168082101561220957612209612121565b90039392505050565b600063ffffffff838116908316818110156121ba576121ba612121565b600063ffffffff80831681851680830382111561215457612154612121565b600063ffffffff8083168181141561226857612268612121565b6001019392505050565b600063ffffffff80841680612289576122896121c2565b92169190910692915050565b60006001600160401b0380841680612289576122896121c2565b60006001600160401b03838116908316818110156121ba576121ba612121565b634e487b7160e01b600052603260045260246000fd5b60006000198214156122f9576122f9612121565b5060010190565b6000821982111561231357612313612121565b500190565b634e487b7160e01b600052605160045260246000fdfea26469706673582212206d20f5051fec8ca7db44e43ae1b109c681c63dbf478b3e6492752fd50681f0df64736f6c63430008090033",
}

// OneStepProverMathABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProverMathMetaData.ABI instead.
var OneStepProverMathABI = OneStepProverMathMetaData.ABI

// OneStepProverMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProverMathMetaData.Bin instead.
var OneStepProverMathBin = OneStepProverMathMetaData.Bin

// DeployOneStepProverMath deploys a new Ethereum contract, binding an instance of OneStepProverMath to it.
func DeployOneStepProverMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProverMath, error) {
	parsed, err := OneStepProverMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProverMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProverMath{OneStepProverMathCaller: OneStepProverMathCaller{contract: contract}, OneStepProverMathTransactor: OneStepProverMathTransactor{contract: contract}, OneStepProverMathFilterer: OneStepProverMathFilterer{contract: contract}}, nil
}

// OneStepProverMath is an auto generated Go binding around an Ethereum contract.
type OneStepProverMath struct {
	OneStepProverMathCaller     // Read-only binding to the contract
	OneStepProverMathTransactor // Write-only binding to the contract
	OneStepProverMathFilterer   // Log filterer for contract events
}

// OneStepProverMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProverMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProverMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProverMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProverMathSession struct {
	Contract     *OneStepProverMath // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OneStepProverMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProverMathCallerSession struct {
	Contract *OneStepProverMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OneStepProverMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProverMathTransactorSession struct {
	Contract     *OneStepProverMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OneStepProverMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProverMathRaw struct {
	Contract *OneStepProverMath // Generic contract binding to access the raw methods on
}

// OneStepProverMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProverMathCallerRaw struct {
	Contract *OneStepProverMathCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProverMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProverMathTransactorRaw struct {
	Contract *OneStepProverMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProverMath creates a new instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMath(address common.Address, backend bind.ContractBackend) (*OneStepProverMath, error) {
	contract, err := bindOneStepProverMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMath{OneStepProverMathCaller: OneStepProverMathCaller{contract: contract}, OneStepProverMathTransactor: OneStepProverMathTransactor{contract: contract}, OneStepProverMathFilterer: OneStepProverMathFilterer{contract: contract}}, nil
}

// NewOneStepProverMathCaller creates a new read-only instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMathCaller(address common.Address, caller bind.ContractCaller) (*OneStepProverMathCaller, error) {
	contract, err := bindOneStepProverMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMathCaller{contract: contract}, nil
}

// NewOneStepProverMathTransactor creates a new write-only instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMathTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProverMathTransactor, error) {
	contract, err := bindOneStepProverMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMathTransactor{contract: contract}, nil
}

// NewOneStepProverMathFilterer creates a new log filterer instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMathFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProverMathFilterer, error) {
	contract, err := bindOneStepProverMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMathFilterer{contract: contract}, nil
}

// bindOneStepProverMath binds a generic wrapper to an already deployed contract.
func bindOneStepProverMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProverMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMath *OneStepProverMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMath.Contract.OneStepProverMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMath *OneStepProverMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.OneStepProverMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMath *OneStepProverMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.OneStepProverMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMath *OneStepProverMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMath *OneStepProverMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMath *OneStepProverMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xda78e7d1.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathCaller) ExecuteOneStep(opts *bind.CallOpts, arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProverMath.contract.Call(opts, &out, "executeOneStep", arg0, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xda78e7d1.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMath.Contract.ExecuteOneStep(&_OneStepProverMath.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xda78e7d1.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathCallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMath.Contract.ExecuteOneStep(&_OneStepProverMath.CallOpts, arg0, startMach, startMod, inst, proof)
}

// OneStepProverMemoryMetaData contains all meta data concerning the OneStepProverMemory contract.
var OneStepProverMemoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611e07806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063da78e7d114610030575b600080fd5b61004361003e36600461138d565b61005a565b60405161005192919061159f565b60405180910390f35b6100626112c2565b6040805160a0810182526000808252825160608082018552828252602080830184905282860184905284019190915292820181905291810182905260808101919091526100ae87611a47565b91506100bf36879003870187611b4a565b905060006100d06020870187611be1565b905061136b602861ffff8316108015906100ef5750603561ffff831611155b156100fd57506101b4610196565b603661ffff8316108015906101175750603e61ffff831611155b1561012557506106bf610196565b61ffff8216603f141561013b5750610a69610196565b61ffff8216604014156101515750610aa1610196565b60405162461bcd60e51b8152602060048201526015602482015274494e56414c49445f4d454d4f52595f4f50434f444560581b60448201526064015b60405180910390fd5b6101a784848989898663ffffffff16565b5050965096945050505050565b6000808060286101c76020880188611be1565b61ffff1614156101e05750600091506004905081610435565b60296101ef6020880188611be1565b61ffff161415610209575060019150600890506000610435565b602a6102186020880188611be1565b61ffff161415610232575060029150600490506000610435565b602b6102416020880188611be1565b61ffff16141561025b575060039150600890506000610435565b602c61026a6020880188611be1565b61ffff1614156102835750600091506001905080610435565b602d6102926020880188611be1565b61ffff1614156102ab5750600091506001905081610435565b602e6102ba6020880188611be1565b61ffff1614156102d4575060009150600290506001610435565b602f6102e36020880188611be1565b61ffff1614156102fc5750600091506002905081610435565b603061030b6020880188611be1565b61ffff16141561032357506001915081905080610435565b60316103326020880188611be1565b61ffff16141561034b5750600191508190506000610435565b603261035a6020880188611be1565b61ffff1614156103735750600191506002905081610435565b60336103826020880188611be1565b61ffff16141561039c575060019150600290506000610435565b60346103ab6020880188611be1565b61ffff1614156103c45750600191506004905081610435565b60356103d36020880188611be1565b61ffff1614156103ed575060019150600490506000610435565b60405162461bcd60e51b815260206004820152601a60248201527f494e56414c49445f4d454d4f52595f4c4f41445f4f50434f4445000000000000604482015260640161018d565b600061044c6104478a60200151610b50565b610b75565b6104609063ffffffff166020890135611c1b565b6020890151519091506001600160401b031661047c8483611c1b565b111561049057505060028752506106b89050565b60006000198180805b878110156105355760006104ad8288611c1b565b905060006104bc602083611c49565b90508581146104e0576104d68f60200151828f8f8b610c06565b5097509095509350845b60006104ed602084611c5d565b90506104fa846008611c71565b6001600160401b031661050d8783610ca0565b60ff166001600160401b0316901b85179450505050808061052d90611c90565b915050610499565b5085156106745786600114801561055d5750600088600681111561055b5761055b611470565b145b15610573578060000b63ffffffff169050610674565b8660011480156105945750600188600681111561059257610592611470565b145b156105a15760000b610674565b8660021480156105c2575060008860068111156105c0576105c0611470565b145b156105d8578060010b63ffffffff169050610674565b8660021480156105f9575060018860068111156105f7576105f7611470565b145b156106065760010b610674565b8660041480156106275750600188600681111561062557610625611470565b145b156106345760030b610674565b60405162461bcd60e51b815260206004820152601560248201527410905117d491505117d096551154d7d4d251d39151605a1b604482015260640161018d565b6106af60405180604001604052808a600681111561069457610694611470565b81526001600160401b0384166020918201528f015190610d1a565b50505050505050505b5050505050565b6000808060366106d26020880188611be1565b61ffff1614156106e85750600491506000610857565b60376106f76020880188611be1565b61ffff16141561070d5750600891506001610857565b603861071c6020880188611be1565b61ffff1614156107325750600491506002610857565b60396107416020880188611be1565b61ffff1614156107575750600891506003610857565b603a6107666020880188611be1565b61ffff16141561077c5750600191506000610857565b603b61078b6020880188611be1565b61ffff1614156107a15750600291506000610857565b603c6107b06020880188611be1565b61ffff1614156107c557506001915081610857565b603d6107d46020880188611be1565b61ffff1614156107ea5750600291506001610857565b603e6107f96020880188611be1565b61ffff16141561080f5750600491506001610857565b60405162461bcd60e51b815260206004820152601b60248201527f494e56414c49445f4d454d4f52595f53544f52455f4f50434f44450000000000604482015260640161018d565b60006108668960200151610b50565b905081600681111561087a5761087a611470565b8151600681111561088d5761088d611470565b146108cb5760405162461bcd60e51b815260206004820152600e60248201526d4241445f53544f52455f5459504560901b604482015260640161018d565b806020015192506008846001600160401b031610156109165760016108f1856008611cab565b6001600160401b031660016001600160401b0316901b6109119190611cda565b831692505b5050600061092a6104478960200151610b50565b61093e9063ffffffff166020880135611c1b565b90508660200151600001516001600160401b0316836001600160401b0316826109679190611c1b565b111561097957505060028652506106b8565b604080516020810190915260608152600090600019906000805b876001600160401b0316811015610a465760006109b08288611c1b565b905060006109bf602083611c49565b9050858114610a045760001986146109e6576109dc858786610d2a565b60208f0151604001525b6109f78e60200151828e8e8b610c06565b9098509196509094509250845b6000610a11602084611c5d565b9050610a1e85828c610dab565b945060088a6001600160401b0316901c99505050508080610a3e90611c90565b915050610993565b50610a52828483610d2a565b60208c015160400152505050505050505050505050565b602084015151600090610a80906201000090611d02565b9050610a99610a8e82610e30565b602088015190610d1a565b505050505050565b602084015151600090610ab8906201000090611d02565b90506000610acc6104478860200151610b50565b90506000610ae363ffffffff808416908516611c1b565b90508660200151602001516001600160401b03168111610b3857610b0a6201000082611c71565b60208801516001600160401b039091169052610b33610b2884610e30565b60208a015190610d1a565b610b46565b610b46610b28600019610e30565b5050505050505050565b60408051808201909152600080825260208201528151610b6f90610e63565b92915050565b60208101516000908183516006811115610b9157610b91611470565b14610bc85760405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b604482015260640161018d565b6401000000008110610b6f5760405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b604482015260640161018d565b600080610c1f6040518060200160405280606081525090565b839150610c2d868684610f73565b9093509150610c3d868684610f8f565b925090506000610c4e828986610d2a565b905088604001518114610c945760405162461bcd60e51b815260206004820152600e60248201526d15d493d391d7d3515357d493d3d560921b604482015260640161018d565b50955095509592505050565b600060208210610ceb5760405162461bcd60e51b81526020600482015260166024820152750848288bea0aa9898be988a828cbe84b2a88abe9288b60531b604482015260640161018d565b600082610cfa60016020611d28565b610d049190611d28565b610d0f906008611c71565b9390931c9392505050565b8151610d269082611069565b5050565b6040516b26b2b6b7b93c903632b0b31d60a11b6020820152602c81018290526000908190604c01604051602081830303815290604052805190602001209050610da08585836040518060400160405280601381526020017226b2b6b7b93c9036b2b935b632903a3932b29d60691b81525061115c565b9150505b9392505050565b600060208310610df55760405162461bcd60e51b81526020600482015260156024820152740848288bea68aa8be988a828cbe84b2a88abe9288b605b1b604482015260640161018d565b600083610e0460016020611d28565b610e0e9190611d28565b610e19906008611c71565b60ff848116821b911b198616179150509392505050565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b604080518082019091526000808252602082015281518051610e8790600190611d28565b81518110610e9757610e97611d3f565b6020026020010151905060006001836000015151610eb59190611d28565b6001600160401b03811115610ecc57610ecc6116c7565b604051908082528060200260200182016040528015610f1157816020015b6040805180820190915260008082526020820152815260200190600190039081610eea5790505b50905060005b8151811015610f6c578351805182908110610f3457610f34611d3f565b6020026020010151828281518110610f4e57610f4e611d3f565b60200260200101819052508080610f6490611c90565b915050610f17565b5090915290565b60008181610f8286868461122e565b9097909650945050505050565b604080516020810190915260608152816000610fac86868461128c565b92509050600060ff82166001600160401b03811115610fcd57610fcd6116c7565b604051908082528060200260200182016040528015610ff6578160200160208202803683370190505b50905060005b8260ff168160ff16101561104d57611015888886610f73565b838360ff168151811061102a5761102a611d3f565b60200260200101819650828152505050808061104590611d55565b915050610ffc565b5060405180602001604052808281525093505050935093915050565b81515160009061107a906001611c1b565b6001600160401b03811115611091576110916116c7565b6040519080825280602002602001820160405280156110d657816020015b60408051808201909152600080825260208201528152602001906001900390816110af5790505b50905060005b8351518110156111325783518051829081106110fa576110fa611d3f565b602002602001015182828151811061111457611114611d3f565b6020026020010181905250808061112a90611c90565b9150506110dc565b5081818460000151518151811061114b5761114b611d3f565b602090810291909101015290915250565b8160005b85515181101561122557600185166111c15782828760000151838151811061118a5761118a611d3f565b60200260200101516040516020016111a493929190611d75565b60405160208183030381529060405280519060200120915061120c565b82866000015182815181106111d8576111d8611d3f565b6020026020010151836040516020016111f393929190611d75565b6040516020818303038152906040528051906020012091505b60019490941c938061121d81611c90565b915050611160565b50949350505050565b600081815b602081101561128357600883901b925085858381811061125557611255611d3f565b919091013560f81c9390931792508161126d81611c90565b925050808061127b90611c90565b915050611233565b50935093915050565b6000818484828181106112a1576112a1611d3f565b919091013560f81c92508190506112b781611c90565b915050935093915050565b60408051610120810190915280600081526020016112f760408051606080820183529181019182529081526000602082015290565b815260200161131d60408051606080820183529181019182529081526000602082015290565b8152602001611342604051806040016040528060608152602001600080191681525090565b815260006020820181905260408201819052606082018190526080820181905260a09091015290565b611373611dbb565b565b60006040828403121561138757600080fd5b50919050565b6000806000806000808688036101a08112156113a857600080fd5b6113b28989611375565b965060408801356001600160401b03808211156113ce57600080fd5b90890190610120828c0312156113e357600080fd5b81975060e0605f19840112156113f857600080fd5b60608a01965061140c8b6101408c01611375565b95506101808a013592508083111561142357600080fd5b828a0192508a601f84011261143757600080fd5b823591508082111561144857600080fd5b5089602082840101111561145b57600080fd5b60208201935080925050509295509295509295565b634e487b7160e01b600052602160045260246000fd5b6004811061149657611496611470565b9052565b8051600781106114ac576114ac611470565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b80841015611509576114f582865161149a565b9382019360019390930192908501906114e2565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b8281101561158857845161155485825161149a565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a09093019260010161153f565b509687015197909601969096525093949350505050565b60006101008083526115b48184018651611486565b6020850151610120848101526115ce6102208501826114b9565b9050604086015160ff1980868403016101408701526115ed83836114b9565b92506060880151915080868403016101608701525061160c828261151d565b915050608086015161018085015260a08601516116326101a086018263ffffffff169052565b5060c086015163ffffffff81166101c08601525060e086015163ffffffff81166101e086015250908501516102008401529050610da460208301848051825260208101516001600160401b0380825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a083015263ffffffff60808201511660c08301525050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156116ff576116ff6116c7565b60405290565b604051602081016001600160401b03811182821017156116ff576116ff6116c7565b604051608081016001600160401b03811182821017156116ff576116ff6116c7565b60405161012081016001600160401b03811182821017156116ff576116ff6116c7565b60405160a081016001600160401b03811182821017156116ff576116ff6116c7565b604051606081016001600160401b03811182821017156116ff576116ff6116c7565b604051601f8201601f191681016001600160401b03811182821017156117d8576117d86116c7565b604052919050565b8035600481106117ef57600080fd5b919050565b60006001600160401b0382111561180d5761180d6116c7565b5060051b60200190565b60006040828403121561182957600080fd5b6118316116dd565b905081356007811061184257600080fd5b808252506020820135602082015292915050565b6000604080838503121561186957600080fd5b6118716116dd565b915082356001600160401b038082111561188a57600080fd5b818501915060208083880312156118a057600080fd5b6118a8611705565b8335838111156118b757600080fd5b80850194505087601f8501126118cc57600080fd5b833592506118e16118dc846117f4565b6117b0565b83815260069390931b8401820192828101908985111561190057600080fd5b948301945b84861015611926576119178a87611817565b82529486019490830190611905565b8252508552948501359484019490945250909392505050565b803563ffffffff811681146117ef57600080fd5b6000604080838503121561196657600080fd5b61196e6116dd565b915082356001600160401b0381111561198657600080fd5b8301601f8101851361199757600080fd5b803560206119a76118dc836117f4565b82815260a092830284018201928282019190898511156119c657600080fd5b948301945b84861015611a2f5780868b0312156119e35760008081fd5b6119eb611727565b6119f58b88611817565b815287870135858201526060611a0c81890161193f565b89830152611a1c6080890161193f565b90820152835294850194918301916119cb565b50808752505080860135818601525050505092915050565b60006101208236031215611a5a57600080fd5b611a62611749565b611a6b836117e0565b815260208301356001600160401b0380821115611a8757600080fd5b611a9336838701611856565b60208401526040850135915080821115611aac57600080fd5b611ab836838701611856565b60408401526060850135915080821115611ad157600080fd5b50611ade36828601611953565b60608301525060808301356080820152611afa60a0840161193f565b60a0820152611b0b60c0840161193f565b60c0820152611b1c60e0840161193f565b60e082015261010092830135928101929092525090565b80356001600160401b03811681146117ef57600080fd5b600081830360e0811215611b5d57600080fd5b611b6561176c565b833581526060601f1983011215611b7b57600080fd5b611b8361178e565b9150611b9160208501611b33565b8252611b9f60408501611b33565b6020830152606084013560408301528160208201526080840135604082015260a08401356060820152611bd460c0850161193f565b6080820152949350505050565b600060208284031215611bf357600080fd5b813561ffff81168114610da457600080fd5b634e487b7160e01b600052601160045260246000fd5b60008219821115611c2e57611c2e611c05565b500190565b634e487b7160e01b600052601260045260246000fd5b600082611c5857611c58611c33565b500490565b600082611c6c57611c6c611c33565b500690565b6000816000190483118215151615611c8b57611c8b611c05565b500290565b6000600019821415611ca457611ca4611c05565b5060010190565b60006001600160401b0380831681851681830481118215151615611cd157611cd1611c05565b02949350505050565b60006001600160401b0383811690831681811015611cfa57611cfa611c05565b039392505050565b60006001600160401b0380841680611d1c57611d1c611c33565b92169190910492915050565b600082821015611d3a57611d3a611c05565b500390565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff811415611d6c57611d6c611c05565b60010192915050565b6000845160005b81811015611d965760208188018101518583015201611d7c565b81811115611da5576000828501525b5091909101928352506020820152604001919050565b634e487b7160e01b600052605160045260246000fdfea26469706673582212207c8daffac8659578f86b81f371304233c84a267ced7fcf82a02b3737d5481d9264736f6c63430008090033",
}

// OneStepProverMemoryABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProverMemoryMetaData.ABI instead.
var OneStepProverMemoryABI = OneStepProverMemoryMetaData.ABI

// OneStepProverMemoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProverMemoryMetaData.Bin instead.
var OneStepProverMemoryBin = OneStepProverMemoryMetaData.Bin

// DeployOneStepProverMemory deploys a new Ethereum contract, binding an instance of OneStepProverMemory to it.
func DeployOneStepProverMemory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProverMemory, error) {
	parsed, err := OneStepProverMemoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProverMemoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProverMemory{OneStepProverMemoryCaller: OneStepProverMemoryCaller{contract: contract}, OneStepProverMemoryTransactor: OneStepProverMemoryTransactor{contract: contract}, OneStepProverMemoryFilterer: OneStepProverMemoryFilterer{contract: contract}}, nil
}

// OneStepProverMemory is an auto generated Go binding around an Ethereum contract.
type OneStepProverMemory struct {
	OneStepProverMemoryCaller     // Read-only binding to the contract
	OneStepProverMemoryTransactor // Write-only binding to the contract
	OneStepProverMemoryFilterer   // Log filterer for contract events
}

// OneStepProverMemoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProverMemoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMemoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProverMemoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMemoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProverMemoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMemorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProverMemorySession struct {
	Contract     *OneStepProverMemory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OneStepProverMemoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProverMemoryCallerSession struct {
	Contract *OneStepProverMemoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OneStepProverMemoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProverMemoryTransactorSession struct {
	Contract     *OneStepProverMemoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OneStepProverMemoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProverMemoryRaw struct {
	Contract *OneStepProverMemory // Generic contract binding to access the raw methods on
}

// OneStepProverMemoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProverMemoryCallerRaw struct {
	Contract *OneStepProverMemoryCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProverMemoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProverMemoryTransactorRaw struct {
	Contract *OneStepProverMemoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProverMemory creates a new instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemory(address common.Address, backend bind.ContractBackend) (*OneStepProverMemory, error) {
	contract, err := bindOneStepProverMemory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemory{OneStepProverMemoryCaller: OneStepProverMemoryCaller{contract: contract}, OneStepProverMemoryTransactor: OneStepProverMemoryTransactor{contract: contract}, OneStepProverMemoryFilterer: OneStepProverMemoryFilterer{contract: contract}}, nil
}

// NewOneStepProverMemoryCaller creates a new read-only instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemoryCaller(address common.Address, caller bind.ContractCaller) (*OneStepProverMemoryCaller, error) {
	contract, err := bindOneStepProverMemory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemoryCaller{contract: contract}, nil
}

// NewOneStepProverMemoryTransactor creates a new write-only instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemoryTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProverMemoryTransactor, error) {
	contract, err := bindOneStepProverMemory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemoryTransactor{contract: contract}, nil
}

// NewOneStepProverMemoryFilterer creates a new log filterer instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemoryFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProverMemoryFilterer, error) {
	contract, err := bindOneStepProverMemory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemoryFilterer{contract: contract}, nil
}

// bindOneStepProverMemory binds a generic wrapper to an already deployed contract.
func bindOneStepProverMemory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProverMemoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMemory *OneStepProverMemoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMemory.Contract.OneStepProverMemoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMemory *OneStepProverMemoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.OneStepProverMemoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMemory *OneStepProverMemoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.OneStepProverMemoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMemory *OneStepProverMemoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMemory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMemory *OneStepProverMemoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMemory *OneStepProverMemoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xda78e7d1.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemoryCaller) ExecuteOneStep(opts *bind.CallOpts, arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProverMemory.contract.Call(opts, &out, "executeOneStep", arg0, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xda78e7d1.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemorySession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMemory.Contract.ExecuteOneStep(&_OneStepProverMemory.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xda78e7d1.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemoryCallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMemory.Contract.ExecuteOneStep(&_OneStepProverMemory.CallOpts, arg0, startMach, startMod, inst, proof)
}
