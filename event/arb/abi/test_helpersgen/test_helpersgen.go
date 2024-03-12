// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package test_helpersgen

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

// BridgeTesterMetaData contains all meta data concerning the BridgeTester contract.
var BridgeTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotDelayedInbox\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotOutbox\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotRollupOrOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BridgeCallTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"InboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"OutboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"RollupUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSequencerInbox\",\"type\":\"address\"}],\"name\":\"SequencerInboxUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptFundsFromOldBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeOutbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedDelayedInboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"}],\"name\":\"allowedDelayedInboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedOutboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"}],\"name\":\"allowedOutboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delayedInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"enqueueDelayedMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"}],\"name\":\"enqueueSequencerMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seqMessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"acc\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"rollup_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerReportedSubMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setDelayedInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencerInbox\",\"type\":\"address\"}],\"name\":\"setSequencerInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"batchPoster\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"submitBatchSpendingReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"_rollup\",\"type\":\"address\"}],\"name\":\"updateRollupAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5060805161138e61002d6000396000505061138e6000f3fe6080604052600436106101385760003560e01c80639e5d4c49116100ab578063cee3d7281161006f578063cee3d728146103d6578063d5719dc2146103f6578063e76f5c8d14610416578063e77145f4146101f4578063eca067ad14610436578063ee35f3271461044b57600080fd5b80639e5d4c4914610316578063ab5d894314610344578063ae60bd1314610359578063c4d66de814610396578063cb23bcb5146103b657600080fd5b80635fca4a16116100fd5780635fca4a16146102165780637a88b1071461022c57806386598a561461024f5780638db5993b14610299578063919cc706146102ac578063945e1147146102e957600080fd5b806284120c1461014457806316bf557914610168578063413b35bd1461018857806347fb24c5146101d45780634f61f850146101f657600080fd5b3661013f57005b600080fd5b34801561015057600080fd5b506009545b6040519081526020015b60405180910390f35b34801561017457600080fd5b5061015561018336600461104d565b61046b565b34801561019457600080fd5b506101c46101a336600461107e565b6001600160a01b031660009081526002602052604090206001015460ff1690565b604051901515815260200161015f565b3480156101e057600080fd5b506101f46101ef3660046110a2565b61048c565b005b34801561020257600080fd5b506101f461021136600461107e565b610797565b34801561022257600080fd5b50610155600a5481565b34801561023857600080fd5b506101556102473660046110e0565b600092915050565b34801561025b57600080fd5b5061027961026a36600461110c565b50600093849350839250829150565b60408051948552602085019390935291830152606082015260800161015f565b6101556102a736600461113e565b6108bc565b3480156102b857600080fd5b506101f46102c736600461107e565b600680546001600160a01b0319166001600160a01b0392909216919091179055565b3480156102f557600080fd5b5061030961030436600461104d565b610907565b60405161015f9190611185565b34801561032257600080fd5b50610336610331366004611199565b610931565b60405161015f929190611222565b34801561035057600080fd5b50610309610a89565b34801561036557600080fd5b506101c461037436600461107e565b6001600160a01b03166000908152600160208190526040909120015460ff1690565b3480156103a257600080fd5b506101f46103b136600461107e565b610ab6565b3480156103c257600080fd5b50600654610309906001600160a01b031681565b3480156103e257600080fd5b506101f46103f13660046110a2565b610b9c565b34801561040257600080fd5b5061015561041136600461104d565b610e9e565b34801561042257600080fd5b5061030961043136600461104d565b610eae565b34801561044257600080fd5b50600854610155565b34801561045757600080fd5b50600754610309906001600160a01b031681565b6009818154811061047b57600080fd5b600091825260209091200154905081565b6006546001600160a01b031633146105645760065460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b1580156104e357600080fd5b505afa1580156104f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061051b9190611281565b9050336001600160a01b0382161461056257600654604051630739600760e01b81526105599133916001600160a01b0390911690849060040161129e565b60405180910390fd5b505b6001600160a01b0382166000818152600160208181526040928390209182015492518515158152919360ff90931692917f6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521910160405180910390a28080156105c95750825b806105db5750801580156105db575082155b156105e65750505050565b821561067457604080518082018252600380548252600160208084018281526001600160a01b038a166000818152928490529582209451855551938201805460ff1916941515949094179093558154908101825591527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319169091179055610790565b60038054610684906001906112c1565b81548110610694576106946112e6565b6000918252602090912001548254600380546001600160a01b039093169290919081106106c3576106c36112e6565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508160000154600160006003856000015481548110610711576107116112e6565b60009182526020808320909101546001600160a01b031683528201929092526040019020556003805480610747576107476112fc565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526001908190526040822091825501805460ff191690555b50505b5050565b6006546001600160a01b031633146108665760065460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b1580156107ee57600080fd5b505afa158015610802573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108269190611281565b9050336001600160a01b0382161461086457600654604051630739600760e01b81526105599133916001600160a01b0390911690849060040161129e565b505b600780546001600160a01b0319166001600160a01b0383161790556040517f8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a906108b1908390611185565b60405180910390a150565b3360009081526001602081905260408220015460ff166108f1573360405163b6c60ea360e01b81526004016105599190611185565b6108ff848443424887610ebe565b949350505050565b6004818154811061091757600080fd5b6000918252602090912001546001600160a01b0316905081565b3360009081526002602052604081206001015460609060ff1661096957336040516332ea82ab60e01b81526004016105599190611185565b821580159061098057506001600160a01b0386163b155b156109a0578560405163b5cf5b8f60e01b81526004016105599190611185565b600580546001600160a01b0319811633179091556040516001600160a01b039182169188169087906109d59088908890611312565b60006040518083038185875af1925050503d8060008114610a12576040519150601f19603f3d011682016040523d82523d6000602084013e610a17565b606091505b50600580546001600160a01b0319166001600160a01b038581169190911790915560405192955090935088169033907f2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d46690610a77908a908a908a90611322565b60405180910390a35094509492505050565b6005546000906001600160a01b039081161415610aa65750600090565b506005546001600160a01b031690565b600054610100900460ff16610ad15760005460ff1615610ad5565b303b155b610b385760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610559565b600054610100900460ff16158015610b5a576000805461ffff19166101011790555b600580546001600160a01b036001600160a01b03199182168117909255600680549091169184169190911790558015610793576000805461ff00191690555050565b6006546001600160a01b03163314610c6b5760065460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b158015610bf357600080fd5b505afa158015610c07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c2b9190611281565b9050336001600160a01b03821614610c6957600654604051630739600760e01b81526105599133916001600160a01b0390911690849060040161129e565b505b6001600160a01b038216600081815260026020908152604091829020600181015492518515158152909360ff90931692917f49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa910160405180910390a2808015610cd15750825b80610ce3575080158015610ce3575082155b15610cee5750505050565b8215610d7d57604080518082018252600480548252600160208084018281526001600160a01b038a16600081815260029093529582209451855551938201805460ff1916941515949094179093558154908101825591527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180546001600160a01b0319169091179055610790565b60048054610d8d906001906112c1565b81548110610d9d57610d9d6112e6565b6000918252602090912001548254600480546001600160a01b03909316929091908110610dcc57610dcc6112e6565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508160000154600260006004856000015481548110610e1a57610e1a6112e6565b60009182526020808320909101546001600160a01b031683528201929092526040019020556004805480610e5057610e506112fc565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526002905260408120908155600101805460ff1916905550505050565b6008818154811061047b57600080fd5b6003818154811061091757600080fd5b600854604080516001600160f81b031960f88a901b166020808301919091526bffffffffffffffffffffffff1960608a901b1660218301526001600160c01b031960c089811b8216603585015288901b16603d830152604582018490526065820186905260858083018690528351808403909101815260a590920190925280519101206000919060008215610f78576008610f5a6001856112c1565b81548110610f6a57610f6a6112e6565b906000526020600020015490505b6040805160208082018490528183018590528251808303840181526060830180855281519190920120600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3015533905260ff8c1660808201526001600160a01b038b1660a082015260c0810187905260e0810188905267ffffffffffffffff89166101008201529051829185917f5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1918190036101200190a3509098975050505050505050565b60006020828403121561105f57600080fd5b5035919050565b6001600160a01b038116811461107b57600080fd5b50565b60006020828403121561109057600080fd5b813561109b81611066565b9392505050565b600080604083850312156110b557600080fd5b82356110c081611066565b9150602083013580151581146110d557600080fd5b809150509250929050565b600080604083850312156110f357600080fd5b82356110fe81611066565b946020939093013593505050565b6000806000806080858703121561112257600080fd5b5050823594602084013594506040840135936060013592509050565b60008060006060848603121561115357600080fd5b833560ff8116811461116457600080fd5b9250602084013561117481611066565b929592945050506040919091013590565b6001600160a01b0391909116815260200190565b600080600080606085870312156111af57600080fd5b84356111ba81611066565b935060208501359250604085013567ffffffffffffffff808211156111de57600080fd5b818701915087601f8301126111f257600080fd5b81358181111561120157600080fd5b88602082850101111561121357600080fd5b95989497505060200194505050565b821515815260006020604081840152835180604085015260005b818110156112585785810183015185820160600152820161123c565b8181111561126a576000606083870101525b50601f01601f191692909201606001949350505050565b60006020828403121561129357600080fd5b815161109b81611066565b6001600160a01b0393841681529183166020830152909116604082015260600190565b6000828210156112e157634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b8183823760009101908152919050565b83815260406020820152816040820152818360608301376000818301606090810191909152601f909201601f191601019291505056fea26469706673582212200222d1d4c0f68ae17395bf00cd2960d3d1a4659dc906f1affe50de5af307abb664736f6c63430008090033",
}

// BridgeTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeTesterMetaData.ABI instead.
var BridgeTesterABI = BridgeTesterMetaData.ABI

// BridgeTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeTesterMetaData.Bin instead.
var BridgeTesterBin = BridgeTesterMetaData.Bin

// DeployBridgeTester deploys a new Ethereum contract, binding an instance of BridgeTester to it.
func DeployBridgeTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeTester, error) {
	parsed, err := BridgeTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeTester{BridgeTesterCaller: BridgeTesterCaller{contract: contract}, BridgeTesterTransactor: BridgeTesterTransactor{contract: contract}, BridgeTesterFilterer: BridgeTesterFilterer{contract: contract}}, nil
}

// BridgeTester is an auto generated Go binding around an Ethereum contract.
type BridgeTester struct {
	BridgeTesterCaller     // Read-only binding to the contract
	BridgeTesterTransactor // Write-only binding to the contract
	BridgeTesterFilterer   // Log filterer for contract events
}

// BridgeTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeTesterSession struct {
	Contract     *BridgeTester     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeTesterCallerSession struct {
	Contract *BridgeTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BridgeTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTesterTransactorSession struct {
	Contract     *BridgeTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BridgeTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeTesterRaw struct {
	Contract *BridgeTester // Generic contract binding to access the raw methods on
}

// BridgeTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeTesterCallerRaw struct {
	Contract *BridgeTesterCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTesterTransactorRaw struct {
	Contract *BridgeTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeTester creates a new instance of BridgeTester, bound to a specific deployed contract.
func NewBridgeTester(address common.Address, backend bind.ContractBackend) (*BridgeTester, error) {
	contract, err := bindBridgeTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeTester{BridgeTesterCaller: BridgeTesterCaller{contract: contract}, BridgeTesterTransactor: BridgeTesterTransactor{contract: contract}, BridgeTesterFilterer: BridgeTesterFilterer{contract: contract}}, nil
}

// NewBridgeTesterCaller creates a new read-only instance of BridgeTester, bound to a specific deployed contract.
func NewBridgeTesterCaller(address common.Address, caller bind.ContractCaller) (*BridgeTesterCaller, error) {
	contract, err := bindBridgeTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterCaller{contract: contract}, nil
}

// NewBridgeTesterTransactor creates a new write-only instance of BridgeTester, bound to a specific deployed contract.
func NewBridgeTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTesterTransactor, error) {
	contract, err := bindBridgeTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterTransactor{contract: contract}, nil
}

// NewBridgeTesterFilterer creates a new log filterer instance of BridgeTester, bound to a specific deployed contract.
func NewBridgeTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeTesterFilterer, error) {
	contract, err := bindBridgeTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterFilterer{contract: contract}, nil
}

// bindBridgeTester binds a generic wrapper to an already deployed contract.
func bindBridgeTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTester *BridgeTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTester.Contract.BridgeTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTester *BridgeTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTester.Contract.BridgeTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTester *BridgeTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTester.Contract.BridgeTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTester *BridgeTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTester *BridgeTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTester *BridgeTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTester.Contract.contract.Transact(opts, method, params...)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeTester *BridgeTesterCaller) ActiveOutbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "activeOutbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeTester *BridgeTesterSession) ActiveOutbox() (common.Address, error) {
	return _BridgeTester.Contract.ActiveOutbox(&_BridgeTester.CallOpts)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) ActiveOutbox() (common.Address, error) {
	return _BridgeTester.Contract.ActiveOutbox(&_BridgeTester.CallOpts)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterCaller) AllowedDelayedInboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "allowedDelayedInboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeTester.Contract.AllowedDelayedInboxList(&_BridgeTester.CallOpts, arg0)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeTester.Contract.AllowedDelayedInboxList(&_BridgeTester.CallOpts, arg0)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeTester *BridgeTesterCaller) AllowedDelayedInboxes(opts *bind.CallOpts, inbox common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "allowedDelayedInboxes", inbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeTester *BridgeTesterSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeTester.Contract.AllowedDelayedInboxes(&_BridgeTester.CallOpts, inbox)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeTester *BridgeTesterCallerSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeTester.Contract.AllowedDelayedInboxes(&_BridgeTester.CallOpts, inbox)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterCaller) AllowedOutboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "allowedOutboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeTester.Contract.AllowedOutboxList(&_BridgeTester.CallOpts, arg0)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeTester.Contract.AllowedOutboxList(&_BridgeTester.CallOpts, arg0)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_BridgeTester *BridgeTesterCaller) AllowedOutboxes(opts *bind.CallOpts, outbox common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "allowedOutboxes", outbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_BridgeTester *BridgeTesterSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _BridgeTester.Contract.AllowedOutboxes(&_BridgeTester.CallOpts, outbox)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_BridgeTester *BridgeTesterCallerSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _BridgeTester.Contract.AllowedOutboxes(&_BridgeTester.CallOpts, outbox)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterCaller) DelayedInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "delayedInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeTester.Contract.DelayedInboxAccs(&_BridgeTester.CallOpts, arg0)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterCallerSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeTester.Contract.DelayedInboxAccs(&_BridgeTester.CallOpts, arg0)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCaller) DelayedMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "delayedMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.DelayedMessageCount(&_BridgeTester.CallOpts)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCallerSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.DelayedMessageCount(&_BridgeTester.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_BridgeTester *BridgeTesterCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_BridgeTester *BridgeTesterSession) Rollup() (common.Address, error) {
	return _BridgeTester.Contract.Rollup(&_BridgeTester.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) Rollup() (common.Address, error) {
	return _BridgeTester.Contract.Rollup(&_BridgeTester.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeTester *BridgeTesterCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeTester *BridgeTesterSession) SequencerInbox() (common.Address, error) {
	return _BridgeTester.Contract.SequencerInbox(&_BridgeTester.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) SequencerInbox() (common.Address, error) {
	return _BridgeTester.Contract.SequencerInbox(&_BridgeTester.CallOpts)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterCaller) SequencerInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "sequencerInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeTester.Contract.SequencerInboxAccs(&_BridgeTester.CallOpts, arg0)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterCallerSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeTester.Contract.SequencerInboxAccs(&_BridgeTester.CallOpts, arg0)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCaller) SequencerMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "sequencerMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.SequencerMessageCount(&_BridgeTester.CallOpts)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCallerSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.SequencerMessageCount(&_BridgeTester.CallOpts)
}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCaller) SequencerReportedSubMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "sequencerReportedSubMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterSession) SequencerReportedSubMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.SequencerReportedSubMessageCount(&_BridgeTester.CallOpts)
}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCallerSession) SequencerReportedSubMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.SequencerReportedSubMessageCount(&_BridgeTester.CallOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeTester *BridgeTesterTransactor) AcceptFundsFromOldBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "acceptFundsFromOldBridge")
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeTester *BridgeTesterSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeTester.Contract.AcceptFundsFromOldBridge(&_BridgeTester.TransactOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeTester *BridgeTesterTransactorSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeTester.Contract.AcceptFundsFromOldBridge(&_BridgeTester.TransactOpts)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeTester *BridgeTesterTransactor) EnqueueDelayedMessage(opts *bind.TransactOpts, kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "enqueueDelayedMessage", kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeTester *BridgeTesterSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.EnqueueDelayedMessage(&_BridgeTester.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeTester *BridgeTesterTransactorSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.EnqueueDelayedMessage(&_BridgeTester.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeTester *BridgeTesterTransactor) EnqueueSequencerMessage(opts *bind.TransactOpts, dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "enqueueSequencerMessage", dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeTester *BridgeTesterSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeTester.Contract.EnqueueSequencerMessage(&_BridgeTester.TransactOpts, dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeTester *BridgeTesterTransactorSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeTester.Contract.EnqueueSequencerMessage(&_BridgeTester.TransactOpts, dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeTester *BridgeTesterTransactor) ExecuteCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "executeCall", to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeTester *BridgeTesterSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.ExecuteCall(&_BridgeTester.TransactOpts, to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeTester *BridgeTesterTransactorSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.ExecuteCall(&_BridgeTester.TransactOpts, to, value, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_BridgeTester *BridgeTesterTransactor) Initialize(opts *bind.TransactOpts, rollup_ common.Address) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "initialize", rollup_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_BridgeTester *BridgeTesterSession) Initialize(rollup_ common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.Initialize(&_BridgeTester.TransactOpts, rollup_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_BridgeTester *BridgeTesterTransactorSession) Initialize(rollup_ common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.Initialize(&_BridgeTester.TransactOpts, rollup_)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterTransactor) SetDelayedInbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "setDelayedInbox", inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetDelayedInbox(&_BridgeTester.TransactOpts, inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterTransactorSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetDelayedInbox(&_BridgeTester.TransactOpts, inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterTransactor) SetOutbox(opts *bind.TransactOpts, outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "setOutbox", outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetOutbox(&_BridgeTester.TransactOpts, outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterTransactorSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetOutbox(&_BridgeTester.TransactOpts, outbox, enabled)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeTester *BridgeTesterTransactor) SetSequencerInbox(opts *bind.TransactOpts, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "setSequencerInbox", _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeTester *BridgeTesterSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetSequencerInbox(&_BridgeTester.TransactOpts, _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeTester *BridgeTesterTransactorSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetSequencerInbox(&_BridgeTester.TransactOpts, _sequencerInbox)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeTester *BridgeTesterTransactor) SubmitBatchSpendingReport(opts *bind.TransactOpts, batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "submitBatchSpendingReport", batchPoster, dataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeTester *BridgeTesterSession) SubmitBatchSpendingReport(batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.SubmitBatchSpendingReport(&_BridgeTester.TransactOpts, batchPoster, dataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeTester *BridgeTesterTransactorSession) SubmitBatchSpendingReport(batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.SubmitBatchSpendingReport(&_BridgeTester.TransactOpts, batchPoster, dataHash)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x919cc706.
//
// Solidity: function updateRollupAddress(address _rollup) returns()
func (_BridgeTester *BridgeTesterTransactor) UpdateRollupAddress(opts *bind.TransactOpts, _rollup common.Address) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "updateRollupAddress", _rollup)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x919cc706.
//
// Solidity: function updateRollupAddress(address _rollup) returns()
func (_BridgeTester *BridgeTesterSession) UpdateRollupAddress(_rollup common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.UpdateRollupAddress(&_BridgeTester.TransactOpts, _rollup)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x919cc706.
//
// Solidity: function updateRollupAddress(address _rollup) returns()
func (_BridgeTester *BridgeTesterTransactorSession) UpdateRollupAddress(_rollup common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.UpdateRollupAddress(&_BridgeTester.TransactOpts, _rollup)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeTester *BridgeTesterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTester.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeTester *BridgeTesterSession) Receive() (*types.Transaction, error) {
	return _BridgeTester.Contract.Receive(&_BridgeTester.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeTester *BridgeTesterTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeTester.Contract.Receive(&_BridgeTester.TransactOpts)
}

// BridgeTesterBridgeCallTriggeredIterator is returned from FilterBridgeCallTriggered and is used to iterate over the raw logs and unpacked data for BridgeCallTriggered events raised by the BridgeTester contract.
type BridgeTesterBridgeCallTriggeredIterator struct {
	Event *BridgeTesterBridgeCallTriggered // Event containing the contract specifics and raw log

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
func (it *BridgeTesterBridgeCallTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterBridgeCallTriggered)
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
		it.Event = new(BridgeTesterBridgeCallTriggered)
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
func (it *BridgeTesterBridgeCallTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterBridgeCallTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterBridgeCallTriggered represents a BridgeCallTriggered event raised by the BridgeTester contract.
type BridgeTesterBridgeCallTriggered struct {
	Outbox common.Address
	To     common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgeCallTriggered is a free log retrieval operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeTester *BridgeTesterFilterer) FilterBridgeCallTriggered(opts *bind.FilterOpts, outbox []common.Address, to []common.Address) (*BridgeTesterBridgeCallTriggeredIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterBridgeCallTriggeredIterator{contract: _BridgeTester.contract, event: "BridgeCallTriggered", logs: logs, sub: sub}, nil
}

// WatchBridgeCallTriggered is a free log subscription operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeTester *BridgeTesterFilterer) WatchBridgeCallTriggered(opts *bind.WatchOpts, sink chan<- *BridgeTesterBridgeCallTriggered, outbox []common.Address, to []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterBridgeCallTriggered)
				if err := _BridgeTester.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
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

// ParseBridgeCallTriggered is a log parse operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeTester *BridgeTesterFilterer) ParseBridgeCallTriggered(log types.Log) (*BridgeTesterBridgeCallTriggered, error) {
	event := new(BridgeTesterBridgeCallTriggered)
	if err := _BridgeTester.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterInboxToggleIterator is returned from FilterInboxToggle and is used to iterate over the raw logs and unpacked data for InboxToggle events raised by the BridgeTester contract.
type BridgeTesterInboxToggleIterator struct {
	Event *BridgeTesterInboxToggle // Event containing the contract specifics and raw log

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
func (it *BridgeTesterInboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterInboxToggle)
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
		it.Event = new(BridgeTesterInboxToggle)
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
func (it *BridgeTesterInboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterInboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterInboxToggle represents a InboxToggle event raised by the BridgeTester contract.
type BridgeTesterInboxToggle struct {
	Inbox   common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInboxToggle is a free log retrieval operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) FilterInboxToggle(opts *bind.FilterOpts, inbox []common.Address) (*BridgeTesterInboxToggleIterator, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterInboxToggleIterator{contract: _BridgeTester.contract, event: "InboxToggle", logs: logs, sub: sub}, nil
}

// WatchInboxToggle is a free log subscription operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) WatchInboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeTesterInboxToggle, inbox []common.Address) (event.Subscription, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterInboxToggle)
				if err := _BridgeTester.contract.UnpackLog(event, "InboxToggle", log); err != nil {
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

// ParseInboxToggle is a log parse operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) ParseInboxToggle(log types.Log) (*BridgeTesterInboxToggle, error) {
	event := new(BridgeTesterInboxToggle)
	if err := _BridgeTester.contract.UnpackLog(event, "InboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the BridgeTester contract.
type BridgeTesterMessageDeliveredIterator struct {
	Event *BridgeTesterMessageDelivered // Event containing the contract specifics and raw log

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
func (it *BridgeTesterMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterMessageDelivered)
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
		it.Event = new(BridgeTesterMessageDelivered)
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
func (it *BridgeTesterMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterMessageDelivered represents a MessageDelivered event raised by the BridgeTester contract.
type BridgeTesterMessageDelivered struct {
	MessageIndex    *big.Int
	BeforeInboxAcc  [32]byte
	Inbox           common.Address
	Kind            uint8
	Sender          common.Address
	MessageDataHash [32]byte
	BaseFeeL1       *big.Int
	Timestamp       uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeTester *BridgeTesterFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (*BridgeTesterMessageDeliveredIterator, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterMessageDeliveredIterator{contract: _BridgeTester.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeTester *BridgeTesterFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *BridgeTesterMessageDelivered, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterMessageDelivered)
				if err := _BridgeTester.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
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

// ParseMessageDelivered is a log parse operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeTester *BridgeTesterFilterer) ParseMessageDelivered(log types.Log) (*BridgeTesterMessageDelivered, error) {
	event := new(BridgeTesterMessageDelivered)
	if err := _BridgeTester.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterOutboxToggleIterator is returned from FilterOutboxToggle and is used to iterate over the raw logs and unpacked data for OutboxToggle events raised by the BridgeTester contract.
type BridgeTesterOutboxToggleIterator struct {
	Event *BridgeTesterOutboxToggle // Event containing the contract specifics and raw log

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
func (it *BridgeTesterOutboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterOutboxToggle)
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
		it.Event = new(BridgeTesterOutboxToggle)
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
func (it *BridgeTesterOutboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterOutboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterOutboxToggle represents a OutboxToggle event raised by the BridgeTester contract.
type BridgeTesterOutboxToggle struct {
	Outbox  common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOutboxToggle is a free log retrieval operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) FilterOutboxToggle(opts *bind.FilterOpts, outbox []common.Address) (*BridgeTesterOutboxToggleIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterOutboxToggleIterator{contract: _BridgeTester.contract, event: "OutboxToggle", logs: logs, sub: sub}, nil
}

// WatchOutboxToggle is a free log subscription operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) WatchOutboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeTesterOutboxToggle, outbox []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterOutboxToggle)
				if err := _BridgeTester.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
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

// ParseOutboxToggle is a log parse operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) ParseOutboxToggle(log types.Log) (*BridgeTesterOutboxToggle, error) {
	event := new(BridgeTesterOutboxToggle)
	if err := _BridgeTester.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterRollupUpdatedIterator is returned from FilterRollupUpdated and is used to iterate over the raw logs and unpacked data for RollupUpdated events raised by the BridgeTester contract.
type BridgeTesterRollupUpdatedIterator struct {
	Event *BridgeTesterRollupUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeTesterRollupUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterRollupUpdated)
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
		it.Event = new(BridgeTesterRollupUpdated)
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
func (it *BridgeTesterRollupUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterRollupUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterRollupUpdated represents a RollupUpdated event raised by the BridgeTester contract.
type BridgeTesterRollupUpdated struct {
	Rollup common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRollupUpdated is a free log retrieval operation binding the contract event 0xae1f5aa15f6ff844896347ceca2a3c24c8d3a27785efdeacd581a0a95172784a.
//
// Solidity: event RollupUpdated(address rollup)
func (_BridgeTester *BridgeTesterFilterer) FilterRollupUpdated(opts *bind.FilterOpts) (*BridgeTesterRollupUpdatedIterator, error) {

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "RollupUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeTesterRollupUpdatedIterator{contract: _BridgeTester.contract, event: "RollupUpdated", logs: logs, sub: sub}, nil
}

// WatchRollupUpdated is a free log subscription operation binding the contract event 0xae1f5aa15f6ff844896347ceca2a3c24c8d3a27785efdeacd581a0a95172784a.
//
// Solidity: event RollupUpdated(address rollup)
func (_BridgeTester *BridgeTesterFilterer) WatchRollupUpdated(opts *bind.WatchOpts, sink chan<- *BridgeTesterRollupUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "RollupUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterRollupUpdated)
				if err := _BridgeTester.contract.UnpackLog(event, "RollupUpdated", log); err != nil {
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

// ParseRollupUpdated is a log parse operation binding the contract event 0xae1f5aa15f6ff844896347ceca2a3c24c8d3a27785efdeacd581a0a95172784a.
//
// Solidity: event RollupUpdated(address rollup)
func (_BridgeTester *BridgeTesterFilterer) ParseRollupUpdated(log types.Log) (*BridgeTesterRollupUpdated, error) {
	event := new(BridgeTesterRollupUpdated)
	if err := _BridgeTester.contract.UnpackLog(event, "RollupUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterSequencerInboxUpdatedIterator is returned from FilterSequencerInboxUpdated and is used to iterate over the raw logs and unpacked data for SequencerInboxUpdated events raised by the BridgeTester contract.
type BridgeTesterSequencerInboxUpdatedIterator struct {
	Event *BridgeTesterSequencerInboxUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeTesterSequencerInboxUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterSequencerInboxUpdated)
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
		it.Event = new(BridgeTesterSequencerInboxUpdated)
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
func (it *BridgeTesterSequencerInboxUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterSequencerInboxUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterSequencerInboxUpdated represents a SequencerInboxUpdated event raised by the BridgeTester contract.
type BridgeTesterSequencerInboxUpdated struct {
	NewSequencerInbox common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSequencerInboxUpdated is a free log retrieval operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeTester *BridgeTesterFilterer) FilterSequencerInboxUpdated(opts *bind.FilterOpts) (*BridgeTesterSequencerInboxUpdatedIterator, error) {

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeTesterSequencerInboxUpdatedIterator{contract: _BridgeTester.contract, event: "SequencerInboxUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerInboxUpdated is a free log subscription operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeTester *BridgeTesterFilterer) WatchSequencerInboxUpdated(opts *bind.WatchOpts, sink chan<- *BridgeTesterSequencerInboxUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterSequencerInboxUpdated)
				if err := _BridgeTester.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
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

// ParseSequencerInboxUpdated is a log parse operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeTester *BridgeTesterFilterer) ParseSequencerInboxUpdated(log types.Log) (*BridgeTesterSequencerInboxUpdated, error) {
	event := new(BridgeTesterSequencerInboxUpdated)
	if err := _BridgeTester.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptographyPrimitivesTesterMetaData contains all meta data concerning the CryptographyPrimitivesTester contract.
var CryptographyPrimitivesTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[25]\",\"name\":\"input\",\"type\":\"uint256[25]\"}],\"name\":\"keccakF\",\"outputs\":[{\"internalType\":\"uint256[25]\",\"name\":\"\",\"type\":\"uint256[25]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[2]\",\"name\":\"inputChunk\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32\",\"name\":\"hashState\",\"type\":\"bytes32\"}],\"name\":\"sha256Block\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6116e861003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c8063ac90ed4614610045578063e479f5321461006e575b600080fd5b610058610053366004611503565b61008f565b604051610065919061158f565b60405180910390f35b61008161007c3660046115c1565b6100a6565b604051908152602001610065565b610097611488565b6100a0826100e9565b92915050565b60006100e26040518060400160405280856000600281106100c9576100c9611653565b6020908102919091015182528681015191015283610ce3565b9392505050565b6100f1611488565b6100f96114a7565b6101016114a7565b610109611488565b600060405180610300016040528060018152602001618082815260200167800000000000808a8152602001678000000080008000815260200161808b81526020016380000001815260200167800000008000808181526020016780000000000080098152602001608a81526020016088815260200163800080098152602001638000000a8152602001638000808b815260200167800000000000008b8152602001678000000000008089815260200167800000000000800381526020016780000000000080028152602001678000000000000080815260200161800a815260200167800000008000000a81526020016780000000800080818152602001678000000000008080815260200163800000018152602001678000000080008008815250905060005b6018811015610cd8576080878101516060808a01516040808c01516020808e01518e511890911890921890931889526101208b01516101008c015160e08d015160c08e015160a08f0151181818189089018190526101c08b01516101a08c01516101808d01516101608e01516101408f0151181818189289019283526102608b01516102408c01516102208d01516102008e01516101e08f015118181818918901919091526103008a01516102e08b01516102c08c01516102a08d01516102808e0151181818189288018390526001600160401b0360028202166001603f1b91829004179092188652510485600260200201516002026001600160401b03161785600060200201511884600160200201526001603f1b85600360200201518161035a5761035a611669565b0485600360200201516002026001600160401b03161785600160200201511884600260200201526001603f1b85600460200201518161039b5761039b611669565b0485600460200201516002026001600160401b031617856002600581106103c4576103c4611653565b602002015118606085015284516001603f1b9086516060808901519390920460029091026001600160401b031617909118608086810191825286518a5118808b5287516020808d018051909218825289516040808f0180519092189091528a518e8801805190911890528a51948e0180519095189094528901805160a08e0180519091189052805160c08e0180519091189052805160e08e018051909118905280516101008e0180519091189052516101208d018051909118905291880180516101408d018051909118905280516101608d018051909118905280516101808d018051909118905280516101a08d0180519091189052516101c08c018051909118905292870180516101e08c018051909118905280516102008c018051909118905280516102208c018051909118905280516102408c0180519091189052516102608b018051909118905281516102808b018051909118905281516102a08b018051909118905281516102c08b018051909118905281516102e08b018051909118905290516103008a01805190911890529084525163100000009060208901516001600160401b03641000000000909102169190041761010084015260408701516001603d1b9060408901516001600160401b03600890910216919004176101608401526060870151628000009060608901516001600160401b036502000000000090910216919004176102608401526080870151654000000000009060808901516001600160401b036204000090910216919004176102c084015260a08701516001603f1b900487600560200201516002026001600160401b0316178360026019811061063457610634611653565b602002015260c08701516210000081046001602c1b9091026001600160401b039081169190911760a085015260e0880151664000000000000081046104009091028216176101a08501526101008801516208000081046520000000000090910282161761020085015261012088015160048082029092166001603e1b909104176103008501526101408801516101408901516001600160401b036001603e1b909102169190041760808401526101608701516001603a1b906101608901516001600160401b036040909102169190041760e084015261018087015162200000906101808901516001600160401b036001602b1b90910216919004176101408401526101a08701516602000000000000906101a08901516001600160401b0361800090910216919004176102408401526101c08701516008906101c08901516001600160401b036001603d1b90910216919004176102a08401526101e0870151641000000000906101e08901516001600160401b03631000000090910216919004176020840152610200808801516102008901516001600160401b0366800000000000009091021691900417610120840152610220870151648000000000906102208901516001600160401b03630200000090910216919004176101808401526102408701516001602b1b906102408901516001600160401b036220000090910216919004176101e0840152610260870151610100906102608901516001600160401b03600160381b90910216919004176102e0840152610280870151642000000000906102808901516001600160401b036308000000909102169190041760608401526102a08701516001602c1b906102a08901516001600160401b0362100000909102169190041760c08401526102c08701516302000000906102c08901516001600160401b0364800000000090910216919004176101c08401526102e0870151600160381b906102e08901516001600160401b036101009091021691900417610220840152610300870151660400000000000090048760186020020151614000026001600160401b031617836014602002015282600a602002015183600560200201511916836000602002015118876000602002015282600b602002015183600660200201511916836001602002015118876001602002015282600c602002015183600760200201511916836002602002015118876002602002015282600d602002015183600860200201511916836003602002015118876003602002015282600e602002015183600960200201511916836004602002015118876004602002015282600f602002015183600a602002015119168360056020020151188760056020020152826010602002015183600b602002015119168360066020020151188760066020020152826011602002015183600c602002015119168360076020020151188760076020020152826012602002015183600d602002015119168360086020020151188760086020020152826013602002015183600e602002015119168360096020020151188760096020020152826014602002015183600f6020020151191683600a60200201511887600a602002015282601560200201518360106020020151191683600b60200201511887600b602002015282601660200201518360116020020151191683600c60200201511887600c602002015282601760200201518360126020020151191683600d60200201511887600d602002015282601860200201518360136020020151191683600e60200201511887600e602002015282600060200201518360146020020151191683600f60200201511887600f6020020152826001602002015183601560200201511916836010602002015118876010602002015282600260200201518360166020020151191683601160200201511887601160200201528260036020020151836017602002015119168360126020020151188760126020020152826004602002015183601860200201511916836013602002015118876013602002015282600560200201518360006020020151191683601460200201511887601460200201528260066020020151836001602002015119168360156020020151188760156020020152826007602002015183600260200201511916836016602002015118876016602002015282600860200201518360036020020151191683601760200201511887601760200201528260096020020151836004602002015119168360186020020151188760186020020152818160188110610cc657610cc6611653565b6020020151875118875260010161022f565b509495945050505050565b604080516108008101825263428a2f9881526371374491602082015263b5c0fbcf9181019190915263e9b5dba56060820152633956c25b60808201526359f111f160a082015263923f82a460c082015263ab1c5ed560e082015263d807aa986101008201526312835b0161012082015263243185be61014082015263550c7dc36101608201526372be5d746101808201526380deb1fe6101a0820152639bdc06a76101c082015263c19bf1746101e082015263e49b69c161020082015263efbe4786610220820152630fc19dc661024082015263240ca1cc610260820152632de92c6f610280820152634a7484aa6102a0820152635cb0a9dc6102c08201526376f988da6102e082015263983e515261030082015263a831c66d61032082015263b00327c861034082015263bf597fc761036082015263c6e00bf361038082015263d5a791476103a08201526306ca63516103c082015263142929676103e08201526327b70a85610400820152632e1b2138610420820152634d2c6dfc6104408201526353380d1361046082015263650a735461048082015263766a0abb6104a08201526381c2c92e6104c08201526392722c856104e082015263a2bfe8a161050082015263a81a664b61052082015263c24b8b7061054082015263c76c51a361056082015263d192e81961058082015263d69906246105a082015263f40e35856105c082015263106aa0706105e08201526319a4c116610600820152631e376c08610620820152632748774c6106408201526334b0bcb561066082015263391c0cb3610680820152634ed8aa4a6106a0820152635b9cca4f6106c082015263682e6ff36106e082015263748f82ee6107008201526378a5636f6107208201526384c87814610740820152638cc702086107608201526390befffa61078082015263a4506ceb6107a082015263bef9a3f76107c082015263c67178f26107e0820152600090610fb06114c5565b60005b60088163ffffffff1610156110495763ffffffff6020820260e003168660006020020151901c828263ffffffff1660408110610ff157610ff1611653565b63ffffffff92831660209182029290920191909152820260e003168660016020020151901c828260080163ffffffff166040811061103157611031611653565b63ffffffff9092166020929092020152600101610fb3565b5060106000805b60408363ffffffff1610156111db57600384600f850363ffffffff166040811061107c5761107c611653565b602002015163ffffffff16901c6110b385600f860363ffffffff16604081106110a7576110a7611653565b60200201516012611453565b6110dd86600f870363ffffffff16604081106110d1576110d1611653565b60200201516007611453565b18189150600a846002850363ffffffff16604081106110fe576110fe611653565b602002015163ffffffff16901c611135856002860363ffffffff166040811061112957611129611653565b60200201516013611453565b61115f866002870363ffffffff166040811061115357611153611653565b60200201516011611453565b1818905080846007850363ffffffff166040811061117f5761117f611653565b602002015183866010870363ffffffff16604081106111a0576111a0611653565b6020020151010101848463ffffffff16604081106111c0576111c0611653565b63ffffffff9092166020929092020152600190920191611050565b6111e36114e4565b600093505b60088463ffffffff16101561123a578360200260e00363ffffffff1688901c818563ffffffff166008811061121f5761121f611653565b63ffffffff90921660209290920201526001909301926111e8565b60008060008096505b60408763ffffffff161015611396576080840151611262906019611453565b608085015161127290600b611453565b6080860151611282906006611453565b18189450878763ffffffff166040811061129e5761129e611653565b6020020151898863ffffffff16604081106112bb576112bb611653565b6020020151608086015160a087015160c08801518219169116188787600760200201510101010192506112f684600060200201516016611453565b845161130390600d611453565b8551611310906002611453565b6040870180516020890180518a5160c08c01805163ffffffff90811660e08f015260a08e018051821690925260808e018051821690925260608e0180518e01821690925280861690915280831690955284811690925280831891909116911618929091189290921881810186810190931687526001999099019897509092509050611243565b600096505b60088763ffffffff1610156113f0578660200260e00363ffffffff168b901c848863ffffffff16600881106113d2576113d2611653565b60200201805163ffffffff920191909116905260019096019561139b565b60008097505b60088863ffffffff161015611443578760200260e00363ffffffff16858963ffffffff166008811061142a5761142a611653565b602002015160019099019863ffffffff16901b176113f6565b9c9b505050505050505050505050565b600061146082602061167f565b63ffffffff168363ffffffff16901b8263ffffffff168463ffffffff16901c17905092915050565b6040518061032001604052806019906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b6040518061080001604052806040906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b600061032080838503121561151757600080fd5b83601f84011261152657600080fd5b6040518181018181106001600160401b038211171561155557634e487b7160e01b600052604160045260246000fd5b60405290830190808583111561156a57600080fd5b845b8381101561158457803582526020918201910161156c565b509095945050505050565b6103208101818360005b60198110156115b8578151835260209283019290910190600101611599565b50505092915050565b600080606083850312156115d457600080fd5b83601f8401126115e357600080fd5b604051604081018181106001600160401b038211171561161357634e487b7160e01b600052604160045260246000fd5b806040525080604085018681111561162a57600080fd5b855b8181101561164457803583526020928301920161162c565b50919691359550909350505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601260045260246000fd5b600063ffffffff838116908316818110156116aa57634e487b7160e01b600052601160045260246000fd5b03939250505056fea264697066735822122079db04e60449ed09bb6c0649b9629678bc5b45a89ed9095af66d3cf29bbf98fe64736f6c63430008090033",
}

// CryptographyPrimitivesTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use CryptographyPrimitivesTesterMetaData.ABI instead.
var CryptographyPrimitivesTesterABI = CryptographyPrimitivesTesterMetaData.ABI

// CryptographyPrimitivesTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CryptographyPrimitivesTesterMetaData.Bin instead.
var CryptographyPrimitivesTesterBin = CryptographyPrimitivesTesterMetaData.Bin

// DeployCryptographyPrimitivesTester deploys a new Ethereum contract, binding an instance of CryptographyPrimitivesTester to it.
func DeployCryptographyPrimitivesTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CryptographyPrimitivesTester, error) {
	parsed, err := CryptographyPrimitivesTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CryptographyPrimitivesTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CryptographyPrimitivesTester{CryptographyPrimitivesTesterCaller: CryptographyPrimitivesTesterCaller{contract: contract}, CryptographyPrimitivesTesterTransactor: CryptographyPrimitivesTesterTransactor{contract: contract}, CryptographyPrimitivesTesterFilterer: CryptographyPrimitivesTesterFilterer{contract: contract}}, nil
}

// CryptographyPrimitivesTester is an auto generated Go binding around an Ethereum contract.
type CryptographyPrimitivesTester struct {
	CryptographyPrimitivesTesterCaller     // Read-only binding to the contract
	CryptographyPrimitivesTesterTransactor // Write-only binding to the contract
	CryptographyPrimitivesTesterFilterer   // Log filterer for contract events
}

// CryptographyPrimitivesTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CryptographyPrimitivesTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptographyPrimitivesTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CryptographyPrimitivesTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptographyPrimitivesTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CryptographyPrimitivesTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptographyPrimitivesTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CryptographyPrimitivesTesterSession struct {
	Contract     *CryptographyPrimitivesTester // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CryptographyPrimitivesTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CryptographyPrimitivesTesterCallerSession struct {
	Contract *CryptographyPrimitivesTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// CryptographyPrimitivesTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CryptographyPrimitivesTesterTransactorSession struct {
	Contract     *CryptographyPrimitivesTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// CryptographyPrimitivesTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CryptographyPrimitivesTesterRaw struct {
	Contract *CryptographyPrimitivesTester // Generic contract binding to access the raw methods on
}

// CryptographyPrimitivesTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CryptographyPrimitivesTesterCallerRaw struct {
	Contract *CryptographyPrimitivesTesterCaller // Generic read-only contract binding to access the raw methods on
}

// CryptographyPrimitivesTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CryptographyPrimitivesTesterTransactorRaw struct {
	Contract *CryptographyPrimitivesTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCryptographyPrimitivesTester creates a new instance of CryptographyPrimitivesTester, bound to a specific deployed contract.
func NewCryptographyPrimitivesTester(address common.Address, backend bind.ContractBackend) (*CryptographyPrimitivesTester, error) {
	contract, err := bindCryptographyPrimitivesTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CryptographyPrimitivesTester{CryptographyPrimitivesTesterCaller: CryptographyPrimitivesTesterCaller{contract: contract}, CryptographyPrimitivesTesterTransactor: CryptographyPrimitivesTesterTransactor{contract: contract}, CryptographyPrimitivesTesterFilterer: CryptographyPrimitivesTesterFilterer{contract: contract}}, nil
}

// NewCryptographyPrimitivesTesterCaller creates a new read-only instance of CryptographyPrimitivesTester, bound to a specific deployed contract.
func NewCryptographyPrimitivesTesterCaller(address common.Address, caller bind.ContractCaller) (*CryptographyPrimitivesTesterCaller, error) {
	contract, err := bindCryptographyPrimitivesTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CryptographyPrimitivesTesterCaller{contract: contract}, nil
}

// NewCryptographyPrimitivesTesterTransactor creates a new write-only instance of CryptographyPrimitivesTester, bound to a specific deployed contract.
func NewCryptographyPrimitivesTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*CryptographyPrimitivesTesterTransactor, error) {
	contract, err := bindCryptographyPrimitivesTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CryptographyPrimitivesTesterTransactor{contract: contract}, nil
}

// NewCryptographyPrimitivesTesterFilterer creates a new log filterer instance of CryptographyPrimitivesTester, bound to a specific deployed contract.
func NewCryptographyPrimitivesTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*CryptographyPrimitivesTesterFilterer, error) {
	contract, err := bindCryptographyPrimitivesTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CryptographyPrimitivesTesterFilterer{contract: contract}, nil
}

// bindCryptographyPrimitivesTester binds a generic wrapper to an already deployed contract.
func bindCryptographyPrimitivesTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CryptographyPrimitivesTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CryptographyPrimitivesTester.Contract.CryptographyPrimitivesTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptographyPrimitivesTester.Contract.CryptographyPrimitivesTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptographyPrimitivesTester.Contract.CryptographyPrimitivesTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CryptographyPrimitivesTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptographyPrimitivesTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptographyPrimitivesTester.Contract.contract.Transact(opts, method, params...)
}

// KeccakF is a free data retrieval call binding the contract method 0xac90ed46.
//
// Solidity: function keccakF(uint256[25] input) pure returns(uint256[25])
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterCaller) KeccakF(opts *bind.CallOpts, input [25]*big.Int) ([25]*big.Int, error) {
	var out []interface{}
	err := _CryptographyPrimitivesTester.contract.Call(opts, &out, "keccakF", input)

	if err != nil {
		return *new([25]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([25]*big.Int)).(*[25]*big.Int)

	return out0, err

}

// KeccakF is a free data retrieval call binding the contract method 0xac90ed46.
//
// Solidity: function keccakF(uint256[25] input) pure returns(uint256[25])
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterSession) KeccakF(input [25]*big.Int) ([25]*big.Int, error) {
	return _CryptographyPrimitivesTester.Contract.KeccakF(&_CryptographyPrimitivesTester.CallOpts, input)
}

// KeccakF is a free data retrieval call binding the contract method 0xac90ed46.
//
// Solidity: function keccakF(uint256[25] input) pure returns(uint256[25])
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterCallerSession) KeccakF(input [25]*big.Int) ([25]*big.Int, error) {
	return _CryptographyPrimitivesTester.Contract.KeccakF(&_CryptographyPrimitivesTester.CallOpts, input)
}

// Sha256Block is a free data retrieval call binding the contract method 0xe479f532.
//
// Solidity: function sha256Block(bytes32[2] inputChunk, bytes32 hashState) pure returns(bytes32)
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterCaller) Sha256Block(opts *bind.CallOpts, inputChunk [2][32]byte, hashState [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CryptographyPrimitivesTester.contract.Call(opts, &out, "sha256Block", inputChunk, hashState)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Sha256Block is a free data retrieval call binding the contract method 0xe479f532.
//
// Solidity: function sha256Block(bytes32[2] inputChunk, bytes32 hashState) pure returns(bytes32)
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterSession) Sha256Block(inputChunk [2][32]byte, hashState [32]byte) ([32]byte, error) {
	return _CryptographyPrimitivesTester.Contract.Sha256Block(&_CryptographyPrimitivesTester.CallOpts, inputChunk, hashState)
}

// Sha256Block is a free data retrieval call binding the contract method 0xe479f532.
//
// Solidity: function sha256Block(bytes32[2] inputChunk, bytes32 hashState) pure returns(bytes32)
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterCallerSession) Sha256Block(inputChunk [2][32]byte, hashState [32]byte) ([32]byte, error) {
	return _CryptographyPrimitivesTester.Contract.Sha256Block(&_CryptographyPrimitivesTester.CallOpts, inputChunk, hashState)
}

// EthVaultMetaData contains all meta data concerning the EthVault contract.
var EthVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"justRevert\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_version\",\"type\":\"uint256\"}],\"name\":\"setVersion\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526000805534801561001457600080fd5b5060f3806100236000396000f3fe60806040526004361060305760003560e01c8063408def1e14603557806350b23fd214604757806354fd4d5014604d575b600080fd5b6045604036600460a5565b600055565b005b60456073565b348015605857600080fd5b50606160005481565b60405190815260200160405180910390f35b60405162461bcd60e51b815260206004820152600360248201526262796560e81b604482015260640160405180910390fd5b60006020828403121560b657600080fd5b503591905056fea26469706673582212207ebc17d77ebaabefad89161947016e1af2be3259ff0102b57344eb6ffc216b9064736f6c63430008090033",
}

// EthVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use EthVaultMetaData.ABI instead.
var EthVaultABI = EthVaultMetaData.ABI

// EthVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthVaultMetaData.Bin instead.
var EthVaultBin = EthVaultMetaData.Bin

// DeployEthVault deploys a new Ethereum contract, binding an instance of EthVault to it.
func DeployEthVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthVault, error) {
	parsed, err := EthVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthVaultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthVault{EthVaultCaller: EthVaultCaller{contract: contract}, EthVaultTransactor: EthVaultTransactor{contract: contract}, EthVaultFilterer: EthVaultFilterer{contract: contract}}, nil
}

// EthVault is an auto generated Go binding around an Ethereum contract.
type EthVault struct {
	EthVaultCaller     // Read-only binding to the contract
	EthVaultTransactor // Write-only binding to the contract
	EthVaultFilterer   // Log filterer for contract events
}

// EthVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthVaultSession struct {
	Contract     *EthVault         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthVaultCallerSession struct {
	Contract *EthVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EthVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthVaultTransactorSession struct {
	Contract     *EthVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthVaultRaw struct {
	Contract *EthVault // Generic contract binding to access the raw methods on
}

// EthVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthVaultCallerRaw struct {
	Contract *EthVaultCaller // Generic read-only contract binding to access the raw methods on
}

// EthVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthVaultTransactorRaw struct {
	Contract *EthVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthVault creates a new instance of EthVault, bound to a specific deployed contract.
func NewEthVault(address common.Address, backend bind.ContractBackend) (*EthVault, error) {
	contract, err := bindEthVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthVault{EthVaultCaller: EthVaultCaller{contract: contract}, EthVaultTransactor: EthVaultTransactor{contract: contract}, EthVaultFilterer: EthVaultFilterer{contract: contract}}, nil
}

// NewEthVaultCaller creates a new read-only instance of EthVault, bound to a specific deployed contract.
func NewEthVaultCaller(address common.Address, caller bind.ContractCaller) (*EthVaultCaller, error) {
	contract, err := bindEthVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthVaultCaller{contract: contract}, nil
}

// NewEthVaultTransactor creates a new write-only instance of EthVault, bound to a specific deployed contract.
func NewEthVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*EthVaultTransactor, error) {
	contract, err := bindEthVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthVaultTransactor{contract: contract}, nil
}

// NewEthVaultFilterer creates a new log filterer instance of EthVault, bound to a specific deployed contract.
func NewEthVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*EthVaultFilterer, error) {
	contract, err := bindEthVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthVaultFilterer{contract: contract}, nil
}

// bindEthVault binds a generic wrapper to an already deployed contract.
func bindEthVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthVault *EthVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthVault.Contract.EthVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthVault *EthVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthVault.Contract.EthVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthVault *EthVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthVault.Contract.EthVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthVault *EthVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthVault *EthVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthVault *EthVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthVault.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_EthVault *EthVaultCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthVault.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_EthVault *EthVaultSession) Version() (*big.Int, error) {
	return _EthVault.Contract.Version(&_EthVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_EthVault *EthVaultCallerSession) Version() (*big.Int, error) {
	return _EthVault.Contract.Version(&_EthVault.CallOpts)
}

// JustRevert is a paid mutator transaction binding the contract method 0x50b23fd2.
//
// Solidity: function justRevert() payable returns()
func (_EthVault *EthVaultTransactor) JustRevert(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthVault.contract.Transact(opts, "justRevert")
}

// JustRevert is a paid mutator transaction binding the contract method 0x50b23fd2.
//
// Solidity: function justRevert() payable returns()
func (_EthVault *EthVaultSession) JustRevert() (*types.Transaction, error) {
	return _EthVault.Contract.JustRevert(&_EthVault.TransactOpts)
}

// JustRevert is a paid mutator transaction binding the contract method 0x50b23fd2.
//
// Solidity: function justRevert() payable returns()
func (_EthVault *EthVaultTransactorSession) JustRevert() (*types.Transaction, error) {
	return _EthVault.Contract.JustRevert(&_EthVault.TransactOpts)
}

// SetVersion is a paid mutator transaction binding the contract method 0x408def1e.
//
// Solidity: function setVersion(uint256 _version) payable returns()
func (_EthVault *EthVaultTransactor) SetVersion(opts *bind.TransactOpts, _version *big.Int) (*types.Transaction, error) {
	return _EthVault.contract.Transact(opts, "setVersion", _version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x408def1e.
//
// Solidity: function setVersion(uint256 _version) payable returns()
func (_EthVault *EthVaultSession) SetVersion(_version *big.Int) (*types.Transaction, error) {
	return _EthVault.Contract.SetVersion(&_EthVault.TransactOpts, _version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x408def1e.
//
// Solidity: function setVersion(uint256 _version) payable returns()
func (_EthVault *EthVaultTransactorSession) SetVersion(_version *big.Int) (*types.Transaction, error) {
	return _EthVault.Contract.SetVersion(&_EthVault.TransactOpts, _version)
}

// MessageTesterMetaData contains all meta data concerning the MessageTester contract.
var MessageTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"inbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"accumulateInboxMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceL1\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"messageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610217806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80638f3c79c01461003b578063bf00905214610087575b600080fd5b61007561004936600461011d565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b60405190815260200160405180910390f35b61007561009536600461015c565b6040805160f89890981b6001600160f81b0319166020808a019190915260609790971b6bffffffffffffffffffffffff1916602189015260c095861b6001600160c01b031990811660358a01529490951b909316603d870152604586019190915260658501526085808501919091528151808503909101815260a59093019052815191012090565b6000806040838503121561013057600080fd5b50508035926020909101359150565b803567ffffffffffffffff8116811461015757600080fd5b919050565b600080600080600080600060e0888a03121561017757600080fd5b873560ff8116811461018857600080fd5b965060208801356001600160a01b03811681146101a457600080fd5b95506101b26040890161013f565b94506101c06060890161013f565b9699959850939660808101359560a0820135955060c090910135935091505056fea264697066735822122072a45102b60096b28934a2038ef6338a8b9f3891357da127e83b1dd321a42e5364736f6c63430008090033",
}

// MessageTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageTesterMetaData.ABI instead.
var MessageTesterABI = MessageTesterMetaData.ABI

// MessageTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageTesterMetaData.Bin instead.
var MessageTesterBin = MessageTesterMetaData.Bin

// DeployMessageTester deploys a new Ethereum contract, binding an instance of MessageTester to it.
func DeployMessageTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageTester, error) {
	parsed, err := MessageTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageTester{MessageTesterCaller: MessageTesterCaller{contract: contract}, MessageTesterTransactor: MessageTesterTransactor{contract: contract}, MessageTesterFilterer: MessageTesterFilterer{contract: contract}}, nil
}

// MessageTester is an auto generated Go binding around an Ethereum contract.
type MessageTester struct {
	MessageTesterCaller     // Read-only binding to the contract
	MessageTesterTransactor // Write-only binding to the contract
	MessageTesterFilterer   // Log filterer for contract events
}

// MessageTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageTesterSession struct {
	Contract     *MessageTester    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageTesterCallerSession struct {
	Contract *MessageTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MessageTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageTesterTransactorSession struct {
	Contract     *MessageTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MessageTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageTesterRaw struct {
	Contract *MessageTester // Generic contract binding to access the raw methods on
}

// MessageTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageTesterCallerRaw struct {
	Contract *MessageTesterCaller // Generic read-only contract binding to access the raw methods on
}

// MessageTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageTesterTransactorRaw struct {
	Contract *MessageTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageTester creates a new instance of MessageTester, bound to a specific deployed contract.
func NewMessageTester(address common.Address, backend bind.ContractBackend) (*MessageTester, error) {
	contract, err := bindMessageTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageTester{MessageTesterCaller: MessageTesterCaller{contract: contract}, MessageTesterTransactor: MessageTesterTransactor{contract: contract}, MessageTesterFilterer: MessageTesterFilterer{contract: contract}}, nil
}

// NewMessageTesterCaller creates a new read-only instance of MessageTester, bound to a specific deployed contract.
func NewMessageTesterCaller(address common.Address, caller bind.ContractCaller) (*MessageTesterCaller, error) {
	contract, err := bindMessageTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTesterCaller{contract: contract}, nil
}

// NewMessageTesterTransactor creates a new write-only instance of MessageTester, bound to a specific deployed contract.
func NewMessageTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageTesterTransactor, error) {
	contract, err := bindMessageTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTesterTransactor{contract: contract}, nil
}

// NewMessageTesterFilterer creates a new log filterer instance of MessageTester, bound to a specific deployed contract.
func NewMessageTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageTesterFilterer, error) {
	contract, err := bindMessageTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageTesterFilterer{contract: contract}, nil
}

// bindMessageTester binds a generic wrapper to an already deployed contract.
func bindMessageTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MessageTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTester *MessageTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageTester.Contract.MessageTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTester *MessageTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTester.Contract.MessageTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTester *MessageTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTester.Contract.MessageTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTester *MessageTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTester *MessageTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTester *MessageTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTester.Contract.contract.Transact(opts, method, params...)
}

// AccumulateInboxMessage is a free data retrieval call binding the contract method 0x8f3c79c0.
//
// Solidity: function accumulateInboxMessage(bytes32 inbox, bytes32 message) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) AccumulateInboxMessage(opts *bind.CallOpts, inbox [32]byte, message [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageTester.contract.Call(opts, &out, "accumulateInboxMessage", inbox, message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AccumulateInboxMessage is a free data retrieval call binding the contract method 0x8f3c79c0.
//
// Solidity: function accumulateInboxMessage(bytes32 inbox, bytes32 message) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) AccumulateInboxMessage(inbox [32]byte, message [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.AccumulateInboxMessage(&_MessageTester.CallOpts, inbox, message)
}

// AccumulateInboxMessage is a free data retrieval call binding the contract method 0x8f3c79c0.
//
// Solidity: function accumulateInboxMessage(bytes32 inbox, bytes32 message) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) AccumulateInboxMessage(inbox [32]byte, message [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.AccumulateInboxMessage(&_MessageTester.CallOpts, inbox, message)
}

// MessageHash is a free data retrieval call binding the contract method 0xbf009052.
//
// Solidity: function messageHash(uint8 messageType, address sender, uint64 blockNumber, uint64 timestamp, uint256 inboxSeqNum, uint256 gasPriceL1, bytes32 messageDataHash) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) MessageHash(opts *bind.CallOpts, messageType uint8, sender common.Address, blockNumber uint64, timestamp uint64, inboxSeqNum *big.Int, gasPriceL1 *big.Int, messageDataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageTester.contract.Call(opts, &out, "messageHash", messageType, sender, blockNumber, timestamp, inboxSeqNum, gasPriceL1, messageDataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageHash is a free data retrieval call binding the contract method 0xbf009052.
//
// Solidity: function messageHash(uint8 messageType, address sender, uint64 blockNumber, uint64 timestamp, uint256 inboxSeqNum, uint256 gasPriceL1, bytes32 messageDataHash) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) MessageHash(messageType uint8, sender common.Address, blockNumber uint64, timestamp uint64, inboxSeqNum *big.Int, gasPriceL1 *big.Int, messageDataHash [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.MessageHash(&_MessageTester.CallOpts, messageType, sender, blockNumber, timestamp, inboxSeqNum, gasPriceL1, messageDataHash)
}

// MessageHash is a free data retrieval call binding the contract method 0xbf009052.
//
// Solidity: function messageHash(uint8 messageType, address sender, uint64 blockNumber, uint64 timestamp, uint256 inboxSeqNum, uint256 gasPriceL1, bytes32 messageDataHash) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) MessageHash(messageType uint8, sender common.Address, blockNumber uint64, timestamp uint64, inboxSeqNum *big.Int, gasPriceL1 *big.Int, messageDataHash [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.MessageHash(&_MessageTester.CallOpts, messageType, sender, blockNumber, timestamp, inboxSeqNum, gasPriceL1, messageDataHash)
}

// OutboxWithoutOptTesterMetaData contains all meta data concerning the OutboxWithoutOptTester contract.
var OutboxWithoutOptTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"AlreadySpent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeCallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxProofLength\",\"type\":\"uint256\"}],\"name\":\"MerkleProofTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIndex\",\"type\":\"uint256\"}],\"name\":\"PathNotMinimal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proofLength\",\"type\":\"uint256\"}],\"name\":\"ProofTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"UnknownRoot\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"zero\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionIndex\",\"type\":\"uint256\"}],\"name\":\"OutBoxTransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2BlockHash\",\"type\":\"bytes32\"}],\"name\":\"SendRootUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OUTBOX_VERSION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"calculateItemHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"item\",\"type\":\"bytes32\"}],\"name\":\"calculateMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"executeTransactionSimulation\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1BatchNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1EthBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1OutputId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"roots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spent\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRollupAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2BlockHash\",\"type\":\"bytes32\"}],\"name\":\"updateSendRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5060805161130b61003060003960006106a2015261130b6000f3fe608060405234801561001057600080fd5b506004361061012b5760003560e01c806395fcea78116100ad578063c4d66de811610071578063c4d66de81461026d578063c75184df14610280578063cb23bcb5146102a0578063d5b5cc23146102b3578063e78cea92146102c657600080fd5b806395fcea78146101695780639f0c04bf14610216578063a04cee6014610229578063ae6dead71461023c578063b0f305371461025c57600080fd5b80635a129efe116100f45780635a129efe146101965780636ae71f12146101c957806372f2a8c7146101d157806380648b02146101d95780638515bc6a146101fe57600080fd5b80627436d31461013057806308635a9514610156578063119852711461016b578063288e5b10146101725780634654779014610185575b600080fd5b61014361013e366004610c79565b6102d9565b6040519081526020015b60405180910390f35b610169610164366004610d9e565b610316565b005b6000610143565b610169610180366004610e92565b610658565b6004546001600160801b0316610143565b6101b96101a4366004610f2d565b60026020526000908152604090205460ff1681565b604051901515815260200161014d565b610169610697565b600654610143565b6007546001600160a01b03165b6040516001600160a01b03909116815260200161014d565b600454600160801b90046001600160801b0316610143565b610143610224366004610f46565b61082c565b610169610237366004610fd4565b610871565b61014361024a366004610f2d565b60036020526000908152604090205481565b6005546001600160801b0316610143565b61016961027b366004610ff6565b6108b0565b610288600281565b6040516001600160801b03909116815260200161014d565b6000546101e6906001600160a01b031681565b6101436102c1366004610f2d565b610989565b6001546101e6906001600160a01b031681565b600061030e8484846040516020016102f391815260200190565b604051602081830303815290604052805190602001206109c5565b949350505050565b6000806103298a8a8a8a8a8a8a8a61082c565b905061036b8d8d808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508f9250859150610a689050565b915060008a6001600160a01b03168a6001600160a01b03167f20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab189648e6040516103b491815260200190565b60405180910390a450600060046040518060a00160405290816000820160009054906101000a90046001600160801b03166001600160801b03166001600160801b031681526020016000820160109054906101000a90046001600160801b03166001600160801b03166001600160801b031681526020016001820160009054906101000a90046001600160801b03166001600160801b03166001600160801b03168152602001600282015481526020016003820160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152505090506040518060a00160405280896001600160801b03168152602001886001600160801b03168152602001876001600160801b031681526020018381526020018b6001600160a01b0316815250600460008201518160000160006101000a8154816001600160801b0302191690836001600160801b0316021790555060208201518160000160106101000a8154816001600160801b0302191690836001600160801b0316021790555060408201518160010160006101000a8154816001600160801b0302191690836001600160801b031602179055506060820151816002015560808201518160030160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050506105e3898686868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610b6b92505050565b805160208201516001600160801b03918216600160801b91831691909102176004556040820151600580546001600160801b03191691909216179055606081015160065560800151600780546001600160a01b0319166001600160a01b03909216919091179055505050505050505050505050565b60405162461bcd60e51b815260206004820152600f60248201526e139bdd081a5b5c1b195b595b9d1959608a1b60448201526064015b60405180910390fd5b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156107255760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b606482015260840161068e565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61038054336001600160a01b0382161461078257604051631194af8760e11b81523360048201526001600160a01b038216602482015260440161068e565b600160009054906101000a90046001600160a01b03166001600160a01b031663cb23bcb56040518163ffffffff1660e01b815260040160206040518083038186803b1580156107d057600080fd5b505afa1580156107e4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610808919061101a565b600080546001600160a01b0319166001600160a01b03929092169190911790555050565b6000888888888888888860405160200161084d989796959493929190611037565b60405160208183030381529060405280519060200120905098975050505050505050565b60008281526003602052604080822083905551829184917fb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f67489190a35050565b6001546001600160a01b0316156108da57604051633bcd329760e21b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040805163cb23bcb560e01b8152905163cb23bcb591600480820192602092909190829003018186803b15801561092e57600080fd5b505afa158015610942573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610966919061101a565b600080546001600160a01b0319166001600160a01b039290921691909117905550565b60405162461bcd60e51b815260206004820152600e60248201526d1393d517d253541311535155115160921b604482015260009060640161068e565b82516000906101008111156109f857604051637ed6198f60e11b815260048101829052610100602482015260440161068e565b8260005b82811015610a5e576000878281518110610a1857610a18611090565b60200260200101519050816001901b871660001415610a4557826000528060205260406000209250610a55565b8060005282602052604060002092505b506001016109fc565b5095945050505050565b6000610100845110610a9257835160405163ab6a068360e01b815260040161068e91815260200190565b8351610a9f9060026111a2565b8310610ad6578284516002610ab491906111a2565b604051630b8a724b60e01b81526004810192909252602482015260440161068e565b6000610ae38585856102d9565b600081815260036020526040902054909150610b15576040516310e61af960e31b81526004810182905260240161068e565b60008481526002602052604090205460ff1615610b4857604051639715b8d360e01b81526004810185905260240161068e565b50506000828152600260205260409020805460ff19166001179055819392505050565b600154604051639e5d4c4960e01b815260009182916001600160a01b0390911690639e5d4c4990610ba4908890889088906004016111de565b600060405180830381600087803b158015610bbe57600080fd5b505af1158015610bd2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610bfa9190810190611227565b9150915081610c2c57805115610c135780518082602001fd5b604051631bb7daad60e11b815260040160405180910390fd5b5050505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715610c7157610c71610c33565b604052919050565b600080600060608486031215610c8e57600080fd5b83356001600160401b0380821115610ca557600080fd5b818601915086601f830112610cb957600080fd5b8135602082821115610ccd57610ccd610c33565b8160051b9250610cde818401610c49565b828152928401810192818101908a851115610cf857600080fd5b948201945b84861015610d1657853582529482019490820190610cfd565b9a918901359950506040909701359695505050505050565b6001600160a01b0381168114610d4357600080fd5b50565b8035610d5181610d2e565b919050565b60008083601f840112610d6857600080fd5b5081356001600160401b03811115610d7f57600080fd5b602083019150836020828501011115610d9757600080fd5b9250929050565b60008060008060008060008060008060006101208c8e031215610dc057600080fd5b8b356001600160401b0380821115610dd757600080fd5b818e0191508e601f830112610deb57600080fd5b813581811115610dfa57600080fd5b8f60208260051b8501011115610e0f57600080fd5b60208381019e50909c508e01359a50610e2a60408f01610d46565b9950610e3860608f01610d46565b985060808e0135975060a08e0135965060c08e0135955060e08e013594506101008e0135915080821115610e6b57600080fd5b50610e788e828f01610d56565b915080935050809150509295989b509295989b9093969950565b60008060008060008060008060006101008a8c031215610eb157600080fd5b8935985060208a0135610ec381610d2e565b975060408a0135610ed381610d2e565b965060608a0135955060808a0135945060a08a0135935060c08a0135925060e08a01356001600160401b03811115610f0a57600080fd5b610f168c828d01610d56565b915080935050809150509295985092959850929598565b600060208284031215610f3f57600080fd5b5035919050565b60008060008060008060008060e0898b031215610f6257600080fd5b8835610f6d81610d2e565b97506020890135610f7d81610d2e565b965060408901359550606089013594506080890135935060a0890135925060c08901356001600160401b03811115610fb457600080fd5b610fc08b828c01610d56565b999c989b5096995094979396929594505050565b60008060408385031215610fe757600080fd5b50508035926020909101359150565b60006020828403121561100857600080fd5b813561101381610d2e565b9392505050565b60006020828403121561102c57600080fd5b815161101381610d2e565b60006bffffffffffffffffffffffff19808b60601b168352808a60601b16601484015250876028830152866048830152856068830152846088830152828460a8840137506000910160a801908152979650505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600181815b808511156110f75781600019048211156110dd576110dd6110a6565b808516156110ea57918102915b93841c93908002906110c1565b509250929050565b60008261110e5750600161119c565b8161111b5750600061119c565b8160018114611131576002811461113b57611157565b600191505061119c565b60ff84111561114c5761114c6110a6565b50506001821b61119c565b5060208310610133831016604e8410600b841016171561117a575081810a61119c565b61118483836110bc565b8060001904821115611198576111986110a6565b0290505b92915050565b600061101383836110ff565b60005b838110156111c95781810151838201526020016111b1565b838111156111d8576000848401525b50505050565b60018060a01b038416815282602082015260606040820152600082518060608401526112118160808501602087016111ae565b601f01601f191691909101608001949350505050565b6000806040838503121561123a57600080fd5b8251801515811461124a57600080fd5b60208401519092506001600160401b038082111561126757600080fd5b818501915085601f83011261127b57600080fd5b81518181111561128d5761128d610c33565b6112a0601f8201601f1916602001610c49565b91508082528660208285010111156112b757600080fd5b6112c88160208401602086016111ae565b508092505050925092905056fea2646970667358221220173cbb3e5a80866102fe447f054268093048c75617ef53dc170453713061b76264736f6c63430008090033",
}

// OutboxWithoutOptTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use OutboxWithoutOptTesterMetaData.ABI instead.
var OutboxWithoutOptTesterABI = OutboxWithoutOptTesterMetaData.ABI

// OutboxWithoutOptTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OutboxWithoutOptTesterMetaData.Bin instead.
var OutboxWithoutOptTesterBin = OutboxWithoutOptTesterMetaData.Bin

// DeployOutboxWithoutOptTester deploys a new Ethereum contract, binding an instance of OutboxWithoutOptTester to it.
func DeployOutboxWithoutOptTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OutboxWithoutOptTester, error) {
	parsed, err := OutboxWithoutOptTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OutboxWithoutOptTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OutboxWithoutOptTester{OutboxWithoutOptTesterCaller: OutboxWithoutOptTesterCaller{contract: contract}, OutboxWithoutOptTesterTransactor: OutboxWithoutOptTesterTransactor{contract: contract}, OutboxWithoutOptTesterFilterer: OutboxWithoutOptTesterFilterer{contract: contract}}, nil
}

// OutboxWithoutOptTester is an auto generated Go binding around an Ethereum contract.
type OutboxWithoutOptTester struct {
	OutboxWithoutOptTesterCaller     // Read-only binding to the contract
	OutboxWithoutOptTesterTransactor // Write-only binding to the contract
	OutboxWithoutOptTesterFilterer   // Log filterer for contract events
}

// OutboxWithoutOptTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutboxWithoutOptTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxWithoutOptTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutboxWithoutOptTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxWithoutOptTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutboxWithoutOptTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxWithoutOptTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutboxWithoutOptTesterSession struct {
	Contract     *OutboxWithoutOptTester // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OutboxWithoutOptTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutboxWithoutOptTesterCallerSession struct {
	Contract *OutboxWithoutOptTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// OutboxWithoutOptTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutboxWithoutOptTesterTransactorSession struct {
	Contract     *OutboxWithoutOptTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// OutboxWithoutOptTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutboxWithoutOptTesterRaw struct {
	Contract *OutboxWithoutOptTester // Generic contract binding to access the raw methods on
}

// OutboxWithoutOptTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutboxWithoutOptTesterCallerRaw struct {
	Contract *OutboxWithoutOptTesterCaller // Generic read-only contract binding to access the raw methods on
}

// OutboxWithoutOptTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutboxWithoutOptTesterTransactorRaw struct {
	Contract *OutboxWithoutOptTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutboxWithoutOptTester creates a new instance of OutboxWithoutOptTester, bound to a specific deployed contract.
func NewOutboxWithoutOptTester(address common.Address, backend bind.ContractBackend) (*OutboxWithoutOptTester, error) {
	contract, err := bindOutboxWithoutOptTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTester{OutboxWithoutOptTesterCaller: OutboxWithoutOptTesterCaller{contract: contract}, OutboxWithoutOptTesterTransactor: OutboxWithoutOptTesterTransactor{contract: contract}, OutboxWithoutOptTesterFilterer: OutboxWithoutOptTesterFilterer{contract: contract}}, nil
}

// NewOutboxWithoutOptTesterCaller creates a new read-only instance of OutboxWithoutOptTester, bound to a specific deployed contract.
func NewOutboxWithoutOptTesterCaller(address common.Address, caller bind.ContractCaller) (*OutboxWithoutOptTesterCaller, error) {
	contract, err := bindOutboxWithoutOptTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTesterCaller{contract: contract}, nil
}

// NewOutboxWithoutOptTesterTransactor creates a new write-only instance of OutboxWithoutOptTester, bound to a specific deployed contract.
func NewOutboxWithoutOptTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*OutboxWithoutOptTesterTransactor, error) {
	contract, err := bindOutboxWithoutOptTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTesterTransactor{contract: contract}, nil
}

// NewOutboxWithoutOptTesterFilterer creates a new log filterer instance of OutboxWithoutOptTester, bound to a specific deployed contract.
func NewOutboxWithoutOptTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*OutboxWithoutOptTesterFilterer, error) {
	contract, err := bindOutboxWithoutOptTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTesterFilterer{contract: contract}, nil
}

// bindOutboxWithoutOptTester binds a generic wrapper to an already deployed contract.
func bindOutboxWithoutOptTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OutboxWithoutOptTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxWithoutOptTester.Contract.OutboxWithoutOptTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.OutboxWithoutOptTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.OutboxWithoutOptTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxWithoutOptTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.contract.Transact(opts, method, params...)
}

// OUTBOXVERSION is a free data retrieval call binding the contract method 0xc75184df.
//
// Solidity: function OUTBOX_VERSION() view returns(uint128)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) OUTBOXVERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "OUTBOX_VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OUTBOXVERSION is a free data retrieval call binding the contract method 0xc75184df.
//
// Solidity: function OUTBOX_VERSION() view returns(uint128)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) OUTBOXVERSION() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.OUTBOXVERSION(&_OutboxWithoutOptTester.CallOpts)
}

// OUTBOXVERSION is a free data retrieval call binding the contract method 0xc75184df.
//
// Solidity: function OUTBOX_VERSION() view returns(uint128)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) OUTBOXVERSION() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.OUTBOXVERSION(&_OutboxWithoutOptTester.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) Bridge() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.Bridge(&_OutboxWithoutOptTester.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) Bridge() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.Bridge(&_OutboxWithoutOptTester.CallOpts)
}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) CalculateItemHash(opts *bind.CallOpts, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "calculateItemHash", l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) CalculateItemHash(l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.CalculateItemHash(&_OutboxWithoutOptTester.CallOpts, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) CalculateItemHash(l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.CalculateItemHash(&_OutboxWithoutOptTester.CallOpts, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) CalculateMerkleRoot(opts *bind.CallOpts, proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "calculateMerkleRoot", proof, path, item)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) CalculateMerkleRoot(proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.CalculateMerkleRoot(&_OutboxWithoutOptTester.CallOpts, proof, path, item)
}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) CalculateMerkleRoot(proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.CalculateMerkleRoot(&_OutboxWithoutOptTester.CallOpts, proof, path, item)
}

// ExecuteTransactionSimulation is a free data retrieval call binding the contract method 0x288e5b10.
//
// Solidity: function executeTransactionSimulation(uint256 , address , address , uint256 , uint256 , uint256 , uint256 , bytes ) pure returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) ExecuteTransactionSimulation(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte) error {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "executeTransactionSimulation", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	if err != nil {
		return err
	}

	return err

}

// ExecuteTransactionSimulation is a free data retrieval call binding the contract method 0x288e5b10.
//
// Solidity: function executeTransactionSimulation(uint256 , address , address , uint256 , uint256 , uint256 , uint256 , bytes ) pure returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) ExecuteTransactionSimulation(arg0 *big.Int, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte) error {
	return _OutboxWithoutOptTester.Contract.ExecuteTransactionSimulation(&_OutboxWithoutOptTester.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// ExecuteTransactionSimulation is a free data retrieval call binding the contract method 0x288e5b10.
//
// Solidity: function executeTransactionSimulation(uint256 , address , address , uint256 , uint256 , uint256 , uint256 , bytes ) pure returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) ExecuteTransactionSimulation(arg0 *big.Int, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte) error {
	return _OutboxWithoutOptTester.Contract.ExecuteTransactionSimulation(&_OutboxWithoutOptTester.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// IsSpent is a free data retrieval call binding the contract method 0x5a129efe.
//
// Solidity: function isSpent(uint256 ) view returns(bool)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) IsSpent(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "isSpent", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSpent is a free data retrieval call binding the contract method 0x5a129efe.
//
// Solidity: function isSpent(uint256 ) view returns(bool)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) IsSpent(arg0 *big.Int) (bool, error) {
	return _OutboxWithoutOptTester.Contract.IsSpent(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// IsSpent is a free data retrieval call binding the contract method 0x5a129efe.
//
// Solidity: function isSpent(uint256 ) view returns(bool)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) IsSpent(arg0 *big.Int) (bool, error) {
	return _OutboxWithoutOptTester.Contract.IsSpent(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() pure returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1BatchNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1BatchNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() pure returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1BatchNum() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1BatchNum(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() pure returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1BatchNum() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1BatchNum(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1Block(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1Block")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1Block() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Block(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1Block() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Block(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1EthBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1EthBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1EthBlock() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1EthBlock(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1EthBlock() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1EthBlock(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1OutputId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1OutputId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1OutputId() ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1OutputId(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1OutputId() ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1OutputId(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1Sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1Sender() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Sender(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1Sender() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Sender(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1Timestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1Timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1Timestamp() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Timestamp(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1Timestamp() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Timestamp(&_OutboxWithoutOptTester.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) Rollup() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.Rollup(&_OutboxWithoutOptTester.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) Rollup() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.Rollup(&_OutboxWithoutOptTester.CallOpts)
}

// Roots is a free data retrieval call binding the contract method 0xae6dead7.
//
// Solidity: function roots(bytes32 ) view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) Roots(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "roots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Roots is a free data retrieval call binding the contract method 0xae6dead7.
//
// Solidity: function roots(bytes32 ) view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) Roots(arg0 [32]byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.Roots(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0xae6dead7.
//
// Solidity: function roots(bytes32 ) view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) Roots(arg0 [32]byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.Roots(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// Spent is a free data retrieval call binding the contract method 0xd5b5cc23.
//
// Solidity: function spent(uint256 ) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) Spent(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "spent", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Spent is a free data retrieval call binding the contract method 0xd5b5cc23.
//
// Solidity: function spent(uint256 ) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) Spent(arg0 *big.Int) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.Spent(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// Spent is a free data retrieval call binding the contract method 0xd5b5cc23.
//
// Solidity: function spent(uint256 ) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) Spent(arg0 *big.Int) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.Spent(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x08635a95.
//
// Solidity: function executeTransaction(bytes32[] proof, uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactor) ExecuteTransaction(opts *bind.TransactOpts, proof [][32]byte, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.contract.Transact(opts, "executeTransaction", proof, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x08635a95.
//
// Solidity: function executeTransaction(bytes32[] proof, uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) ExecuteTransaction(proof [][32]byte, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.ExecuteTransaction(&_OutboxWithoutOptTester.TransactOpts, proof, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x08635a95.
//
// Solidity: function executeTransaction(bytes32[] proof, uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorSession) ExecuteTransaction(proof [][32]byte, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.ExecuteTransaction(&_OutboxWithoutOptTester.TransactOpts, proof, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.contract.Transact(opts, "initialize", _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.Initialize(&_OutboxWithoutOptTester.TransactOpts, _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.Initialize(&_OutboxWithoutOptTester.TransactOpts, _bridge)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactor) PostUpgradeInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.contract.Transact(opts, "postUpgradeInit")
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) PostUpgradeInit() (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.PostUpgradeInit(&_OutboxWithoutOptTester.TransactOpts)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorSession) PostUpgradeInit() (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.PostUpgradeInit(&_OutboxWithoutOptTester.TransactOpts)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x6ae71f12.
//
// Solidity: function updateRollupAddress() returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactor) UpdateRollupAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.contract.Transact(opts, "updateRollupAddress")
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x6ae71f12.
//
// Solidity: function updateRollupAddress() returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) UpdateRollupAddress() (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.UpdateRollupAddress(&_OutboxWithoutOptTester.TransactOpts)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x6ae71f12.
//
// Solidity: function updateRollupAddress() returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorSession) UpdateRollupAddress() (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.UpdateRollupAddress(&_OutboxWithoutOptTester.TransactOpts)
}

// UpdateSendRoot is a paid mutator transaction binding the contract method 0xa04cee60.
//
// Solidity: function updateSendRoot(bytes32 root, bytes32 l2BlockHash) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactor) UpdateSendRoot(opts *bind.TransactOpts, root [32]byte, l2BlockHash [32]byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.contract.Transact(opts, "updateSendRoot", root, l2BlockHash)
}

// UpdateSendRoot is a paid mutator transaction binding the contract method 0xa04cee60.
//
// Solidity: function updateSendRoot(bytes32 root, bytes32 l2BlockHash) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) UpdateSendRoot(root [32]byte, l2BlockHash [32]byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.UpdateSendRoot(&_OutboxWithoutOptTester.TransactOpts, root, l2BlockHash)
}

// UpdateSendRoot is a paid mutator transaction binding the contract method 0xa04cee60.
//
// Solidity: function updateSendRoot(bytes32 root, bytes32 l2BlockHash) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorSession) UpdateSendRoot(root [32]byte, l2BlockHash [32]byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.UpdateSendRoot(&_OutboxWithoutOptTester.TransactOpts, root, l2BlockHash)
}

// OutboxWithoutOptTesterOutBoxTransactionExecutedIterator is returned from FilterOutBoxTransactionExecuted and is used to iterate over the raw logs and unpacked data for OutBoxTransactionExecuted events raised by the OutboxWithoutOptTester contract.
type OutboxWithoutOptTesterOutBoxTransactionExecutedIterator struct {
	Event *OutboxWithoutOptTesterOutBoxTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *OutboxWithoutOptTesterOutBoxTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutboxWithoutOptTesterOutBoxTransactionExecuted)
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
		it.Event = new(OutboxWithoutOptTesterOutBoxTransactionExecuted)
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
func (it *OutboxWithoutOptTesterOutBoxTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutboxWithoutOptTesterOutBoxTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutboxWithoutOptTesterOutBoxTransactionExecuted represents a OutBoxTransactionExecuted event raised by the OutboxWithoutOptTester contract.
type OutboxWithoutOptTesterOutBoxTransactionExecuted struct {
	To               common.Address
	L2Sender         common.Address
	Zero             *big.Int
	TransactionIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOutBoxTransactionExecuted is a free log retrieval operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed to, address indexed l2Sender, uint256 indexed zero, uint256 transactionIndex)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) FilterOutBoxTransactionExecuted(opts *bind.FilterOpts, to []common.Address, l2Sender []common.Address, zero []*big.Int) (*OutboxWithoutOptTesterOutBoxTransactionExecutedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var zeroRule []interface{}
	for _, zeroItem := range zero {
		zeroRule = append(zeroRule, zeroItem)
	}

	logs, sub, err := _OutboxWithoutOptTester.contract.FilterLogs(opts, "OutBoxTransactionExecuted", toRule, l2SenderRule, zeroRule)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTesterOutBoxTransactionExecutedIterator{contract: _OutboxWithoutOptTester.contract, event: "OutBoxTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchOutBoxTransactionExecuted is a free log subscription operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed to, address indexed l2Sender, uint256 indexed zero, uint256 transactionIndex)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) WatchOutBoxTransactionExecuted(opts *bind.WatchOpts, sink chan<- *OutboxWithoutOptTesterOutBoxTransactionExecuted, to []common.Address, l2Sender []common.Address, zero []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var zeroRule []interface{}
	for _, zeroItem := range zero {
		zeroRule = append(zeroRule, zeroItem)
	}

	logs, sub, err := _OutboxWithoutOptTester.contract.WatchLogs(opts, "OutBoxTransactionExecuted", toRule, l2SenderRule, zeroRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutboxWithoutOptTesterOutBoxTransactionExecuted)
				if err := _OutboxWithoutOptTester.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
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

// ParseOutBoxTransactionExecuted is a log parse operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed to, address indexed l2Sender, uint256 indexed zero, uint256 transactionIndex)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) ParseOutBoxTransactionExecuted(log types.Log) (*OutboxWithoutOptTesterOutBoxTransactionExecuted, error) {
	event := new(OutboxWithoutOptTesterOutBoxTransactionExecuted)
	if err := _OutboxWithoutOptTester.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OutboxWithoutOptTesterSendRootUpdatedIterator is returned from FilterSendRootUpdated and is used to iterate over the raw logs and unpacked data for SendRootUpdated events raised by the OutboxWithoutOptTester contract.
type OutboxWithoutOptTesterSendRootUpdatedIterator struct {
	Event *OutboxWithoutOptTesterSendRootUpdated // Event containing the contract specifics and raw log

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
func (it *OutboxWithoutOptTesterSendRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutboxWithoutOptTesterSendRootUpdated)
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
		it.Event = new(OutboxWithoutOptTesterSendRootUpdated)
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
func (it *OutboxWithoutOptTesterSendRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutboxWithoutOptTesterSendRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutboxWithoutOptTesterSendRootUpdated represents a SendRootUpdated event raised by the OutboxWithoutOptTester contract.
type OutboxWithoutOptTesterSendRootUpdated struct {
	OutputRoot  [32]byte
	L2BlockHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSendRootUpdated is a free log retrieval operation binding the contract event 0xb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f6748.
//
// Solidity: event SendRootUpdated(bytes32 indexed outputRoot, bytes32 indexed l2BlockHash)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) FilterSendRootUpdated(opts *bind.FilterOpts, outputRoot [][32]byte, l2BlockHash [][32]byte) (*OutboxWithoutOptTesterSendRootUpdatedIterator, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2BlockHashRule []interface{}
	for _, l2BlockHashItem := range l2BlockHash {
		l2BlockHashRule = append(l2BlockHashRule, l2BlockHashItem)
	}

	logs, sub, err := _OutboxWithoutOptTester.contract.FilterLogs(opts, "SendRootUpdated", outputRootRule, l2BlockHashRule)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTesterSendRootUpdatedIterator{contract: _OutboxWithoutOptTester.contract, event: "SendRootUpdated", logs: logs, sub: sub}, nil
}

// WatchSendRootUpdated is a free log subscription operation binding the contract event 0xb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f6748.
//
// Solidity: event SendRootUpdated(bytes32 indexed outputRoot, bytes32 indexed l2BlockHash)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) WatchSendRootUpdated(opts *bind.WatchOpts, sink chan<- *OutboxWithoutOptTesterSendRootUpdated, outputRoot [][32]byte, l2BlockHash [][32]byte) (event.Subscription, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2BlockHashRule []interface{}
	for _, l2BlockHashItem := range l2BlockHash {
		l2BlockHashRule = append(l2BlockHashRule, l2BlockHashItem)
	}

	logs, sub, err := _OutboxWithoutOptTester.contract.WatchLogs(opts, "SendRootUpdated", outputRootRule, l2BlockHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutboxWithoutOptTesterSendRootUpdated)
				if err := _OutboxWithoutOptTester.contract.UnpackLog(event, "SendRootUpdated", log); err != nil {
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

// ParseSendRootUpdated is a log parse operation binding the contract event 0xb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f6748.
//
// Solidity: event SendRootUpdated(bytes32 indexed outputRoot, bytes32 indexed l2BlockHash)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) ParseSendRootUpdated(log types.Log) (*OutboxWithoutOptTesterSendRootUpdated, error) {
	event := new(OutboxWithoutOptTesterSendRootUpdated)
	if err := _OutboxWithoutOptTester.contract.UnpackLog(event, "SendRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupMockMetaData contains all meta data concerning the RollupMock contract.
var RollupMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"WithdrawTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ZombieTriggered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"removeOldZombies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStakerFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516101e03803806101e083398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b61014d806100936000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806361373919146100465780638da5cb5b14610061578063edfd03ed1461008c575b600080fd5b61004e6100a1565b6040519081526020015b60405180910390f35b600054610074906001600160a01b031681565b6040516001600160a01b039091168152602001610058565b61009f61009a3660046100fe565b6100d2565b005b6040516000907f1c09fbbf7cfd024f5e4e5472dd87afd5d67ee5db6a0ca715bf508d96abce309f908290a150600090565b6040517fb774f793432a37585a7638b9afe49e91c478887a2c0fef32877508bf2f76429d90600090a150565b60006020828403121561011057600080fd5b503591905056fea26469706673582212204e821be1d241b8183557b117343b4ab5c37e83c4432dfbee2e117cf0bc35325d64736f6c63430008090033",
}

// RollupMockABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupMockMetaData.ABI instead.
var RollupMockABI = RollupMockMetaData.ABI

// RollupMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RollupMockMetaData.Bin instead.
var RollupMockBin = RollupMockMetaData.Bin

// DeployRollupMock deploys a new Ethereum contract, binding an instance of RollupMock to it.
func DeployRollupMock(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *RollupMock, error) {
	parsed, err := RollupMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RollupMockBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupMock{RollupMockCaller: RollupMockCaller{contract: contract}, RollupMockTransactor: RollupMockTransactor{contract: contract}, RollupMockFilterer: RollupMockFilterer{contract: contract}}, nil
}

// RollupMock is an auto generated Go binding around an Ethereum contract.
type RollupMock struct {
	RollupMockCaller     // Read-only binding to the contract
	RollupMockTransactor // Write-only binding to the contract
	RollupMockFilterer   // Log filterer for contract events
}

// RollupMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupMockSession struct {
	Contract     *RollupMock       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupMockCallerSession struct {
	Contract *RollupMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RollupMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupMockTransactorSession struct {
	Contract     *RollupMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RollupMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupMockRaw struct {
	Contract *RollupMock // Generic contract binding to access the raw methods on
}

// RollupMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupMockCallerRaw struct {
	Contract *RollupMockCaller // Generic read-only contract binding to access the raw methods on
}

// RollupMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupMockTransactorRaw struct {
	Contract *RollupMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupMock creates a new instance of RollupMock, bound to a specific deployed contract.
func NewRollupMock(address common.Address, backend bind.ContractBackend) (*RollupMock, error) {
	contract, err := bindRollupMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupMock{RollupMockCaller: RollupMockCaller{contract: contract}, RollupMockTransactor: RollupMockTransactor{contract: contract}, RollupMockFilterer: RollupMockFilterer{contract: contract}}, nil
}

// NewRollupMockCaller creates a new read-only instance of RollupMock, bound to a specific deployed contract.
func NewRollupMockCaller(address common.Address, caller bind.ContractCaller) (*RollupMockCaller, error) {
	contract, err := bindRollupMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupMockCaller{contract: contract}, nil
}

// NewRollupMockTransactor creates a new write-only instance of RollupMock, bound to a specific deployed contract.
func NewRollupMockTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupMockTransactor, error) {
	contract, err := bindRollupMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupMockTransactor{contract: contract}, nil
}

// NewRollupMockFilterer creates a new log filterer instance of RollupMock, bound to a specific deployed contract.
func NewRollupMockFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupMockFilterer, error) {
	contract, err := bindRollupMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupMockFilterer{contract: contract}, nil
}

// bindRollupMock binds a generic wrapper to an already deployed contract.
func bindRollupMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RollupMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupMock *RollupMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupMock.Contract.RollupMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupMock *RollupMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupMock.Contract.RollupMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupMock *RollupMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupMock.Contract.RollupMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupMock *RollupMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupMock *RollupMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupMock *RollupMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupMock.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupMock *RollupMockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupMock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupMock *RollupMockSession) Owner() (common.Address, error) {
	return _RollupMock.Contract.Owner(&_RollupMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupMock *RollupMockCallerSession) Owner() (common.Address, error) {
	return _RollupMock.Contract.Owner(&_RollupMock.CallOpts)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 ) returns()
func (_RollupMock *RollupMockTransactor) RemoveOldZombies(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _RollupMock.contract.Transact(opts, "removeOldZombies", arg0)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 ) returns()
func (_RollupMock *RollupMockSession) RemoveOldZombies(arg0 *big.Int) (*types.Transaction, error) {
	return _RollupMock.Contract.RemoveOldZombies(&_RollupMock.TransactOpts, arg0)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 ) returns()
func (_RollupMock *RollupMockTransactorSession) RemoveOldZombies(arg0 *big.Int) (*types.Transaction, error) {
	return _RollupMock.Contract.RemoveOldZombies(&_RollupMock.TransactOpts, arg0)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x61373919.
//
// Solidity: function withdrawStakerFunds() returns(uint256)
func (_RollupMock *RollupMockTransactor) WithdrawStakerFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupMock.contract.Transact(opts, "withdrawStakerFunds")
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x61373919.
//
// Solidity: function withdrawStakerFunds() returns(uint256)
func (_RollupMock *RollupMockSession) WithdrawStakerFunds() (*types.Transaction, error) {
	return _RollupMock.Contract.WithdrawStakerFunds(&_RollupMock.TransactOpts)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x61373919.
//
// Solidity: function withdrawStakerFunds() returns(uint256)
func (_RollupMock *RollupMockTransactorSession) WithdrawStakerFunds() (*types.Transaction, error) {
	return _RollupMock.Contract.WithdrawStakerFunds(&_RollupMock.TransactOpts)
}

// RollupMockWithdrawTriggeredIterator is returned from FilterWithdrawTriggered and is used to iterate over the raw logs and unpacked data for WithdrawTriggered events raised by the RollupMock contract.
type RollupMockWithdrawTriggeredIterator struct {
	Event *RollupMockWithdrawTriggered // Event containing the contract specifics and raw log

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
func (it *RollupMockWithdrawTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupMockWithdrawTriggered)
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
		it.Event = new(RollupMockWithdrawTriggered)
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
func (it *RollupMockWithdrawTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupMockWithdrawTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupMockWithdrawTriggered represents a WithdrawTriggered event raised by the RollupMock contract.
type RollupMockWithdrawTriggered struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTriggered is a free log retrieval operation binding the contract event 0x1c09fbbf7cfd024f5e4e5472dd87afd5d67ee5db6a0ca715bf508d96abce309f.
//
// Solidity: event WithdrawTriggered()
func (_RollupMock *RollupMockFilterer) FilterWithdrawTriggered(opts *bind.FilterOpts) (*RollupMockWithdrawTriggeredIterator, error) {

	logs, sub, err := _RollupMock.contract.FilterLogs(opts, "WithdrawTriggered")
	if err != nil {
		return nil, err
	}
	return &RollupMockWithdrawTriggeredIterator{contract: _RollupMock.contract, event: "WithdrawTriggered", logs: logs, sub: sub}, nil
}

// WatchWithdrawTriggered is a free log subscription operation binding the contract event 0x1c09fbbf7cfd024f5e4e5472dd87afd5d67ee5db6a0ca715bf508d96abce309f.
//
// Solidity: event WithdrawTriggered()
func (_RollupMock *RollupMockFilterer) WatchWithdrawTriggered(opts *bind.WatchOpts, sink chan<- *RollupMockWithdrawTriggered) (event.Subscription, error) {

	logs, sub, err := _RollupMock.contract.WatchLogs(opts, "WithdrawTriggered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupMockWithdrawTriggered)
				if err := _RollupMock.contract.UnpackLog(event, "WithdrawTriggered", log); err != nil {
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

// ParseWithdrawTriggered is a log parse operation binding the contract event 0x1c09fbbf7cfd024f5e4e5472dd87afd5d67ee5db6a0ca715bf508d96abce309f.
//
// Solidity: event WithdrawTriggered()
func (_RollupMock *RollupMockFilterer) ParseWithdrawTriggered(log types.Log) (*RollupMockWithdrawTriggered, error) {
	event := new(RollupMockWithdrawTriggered)
	if err := _RollupMock.contract.UnpackLog(event, "WithdrawTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupMockZombieTriggeredIterator is returned from FilterZombieTriggered and is used to iterate over the raw logs and unpacked data for ZombieTriggered events raised by the RollupMock contract.
type RollupMockZombieTriggeredIterator struct {
	Event *RollupMockZombieTriggered // Event containing the contract specifics and raw log

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
func (it *RollupMockZombieTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupMockZombieTriggered)
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
		it.Event = new(RollupMockZombieTriggered)
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
func (it *RollupMockZombieTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupMockZombieTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupMockZombieTriggered represents a ZombieTriggered event raised by the RollupMock contract.
type RollupMockZombieTriggered struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterZombieTriggered is a free log retrieval operation binding the contract event 0xb774f793432a37585a7638b9afe49e91c478887a2c0fef32877508bf2f76429d.
//
// Solidity: event ZombieTriggered()
func (_RollupMock *RollupMockFilterer) FilterZombieTriggered(opts *bind.FilterOpts) (*RollupMockZombieTriggeredIterator, error) {

	logs, sub, err := _RollupMock.contract.FilterLogs(opts, "ZombieTriggered")
	if err != nil {
		return nil, err
	}
	return &RollupMockZombieTriggeredIterator{contract: _RollupMock.contract, event: "ZombieTriggered", logs: logs, sub: sub}, nil
}

// WatchZombieTriggered is a free log subscription operation binding the contract event 0xb774f793432a37585a7638b9afe49e91c478887a2c0fef32877508bf2f76429d.
//
// Solidity: event ZombieTriggered()
func (_RollupMock *RollupMockFilterer) WatchZombieTriggered(opts *bind.WatchOpts, sink chan<- *RollupMockZombieTriggered) (event.Subscription, error) {

	logs, sub, err := _RollupMock.contract.WatchLogs(opts, "ZombieTriggered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupMockZombieTriggered)
				if err := _RollupMock.contract.UnpackLog(event, "ZombieTriggered", log); err != nil {
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

// ParseZombieTriggered is a log parse operation binding the contract event 0xb774f793432a37585a7638b9afe49e91c478887a2c0fef32877508bf2f76429d.
//
// Solidity: event ZombieTriggered()
func (_RollupMock *RollupMockFilterer) ParseZombieTriggered(log types.Log) (*RollupMockZombieTriggered, error) {
	event := new(RollupMockZombieTriggered)
	if err := _RollupMock.contract.UnpackLog(event, "ZombieTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestTokenMetaData contains all meta data concerning the TestToken contract.
var TestTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000b8838038062000b8883398101604081905262000034916200023f565b60408051808201825260098152682a32b9ba2a37b5b2b760b91b602080830191825283518085019094526002845261151560f21b9084015281519192916200007f9160039162000199565b5080516200009590600490602084019062000199565b505050620000aa3382620000b160201b60201c565b50620002bd565b6001600160a01b0382166200010c5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b806002600082825462000120919062000259565b90915550506001600160a01b038216600090815260208190526040812080548392906200014f90849062000259565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b828054620001a79062000280565b90600052602060002090601f016020900481019282620001cb576000855562000216565b82601f10620001e657805160ff191683800117855562000216565b8280016001018555821562000216579182015b8281111562000216578251825591602001919060010190620001f9565b506200022492915062000228565b5090565b5b8082111562000224576000815560010162000229565b6000602082840312156200025257600080fd5b5051919050565b600082198211156200027b57634e487b7160e01b600052601160045260246000fd5b500190565b600181811c908216806200029557607f821691505b60208210811415620002b757634e487b7160e01b600052602260045260246000fd5b50919050565b6108bb80620002cd6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461012357806370a082311461013657806395d89b411461015f578063a457c2d714610167578063a9059cbb1461017a578063dd62ed3e1461018d57600080fd5b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ef57806323b872dd14610101578063313ce56714610114575b600080fd5b6100b66101a0565b6040516100c391906106f8565b60405180910390f35b6100df6100da366004610769565b610232565b60405190151581526020016100c3565b6002545b6040519081526020016100c3565b6100df61010f366004610793565b61024a565b604051601281526020016100c3565b6100df610131366004610769565b61026e565b6100f36101443660046107cf565b6001600160a01b031660009081526020819052604090205490565b6100b66102ad565b6100df610175366004610769565b6102bc565b6100df610188366004610769565b610353565b6100f361019b3660046107f1565b610361565b6060600380546101af90610824565b80601f01602080910402602001604051908101604052809291908181526020018280546101db90610824565b80156102285780601f106101fd57610100808354040283529160200191610228565b820191906000526020600020905b81548152906001019060200180831161020b57829003601f168201915b5050505050905090565b60003361024081858561038c565b5060019392505050565b6000336102588582856104b0565b61026385858561052a565b506001949350505050565b3360008181526001602090815260408083206001600160a01b038716845290915281205490919061024090829086906102a890879061085f565b61038c565b6060600480546101af90610824565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909190838110156103465760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b610263828686840361038c565b60003361024081858561052a565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166103ee5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b606482015260840161033d565b6001600160a01b03821661044f5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b606482015260840161033d565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b60006104bc8484610361565b9050600019811461052457818110156105175760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161033d565b610524848484840361038c565b50505050565b6001600160a01b03831661058e5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161033d565b6001600160a01b0382166105f05760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b606482015260840161033d565b6001600160a01b038316600090815260208190526040902054818110156106685760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b606482015260840161033d565b6001600160a01b0380851660009081526020819052604080822085850390559185168152908120805484929061069f90849061085f565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516106eb91815260200190565b60405180910390a3610524565b600060208083528351808285015260005b8181101561072557858101830151858201604001528201610709565b81811115610737576000604083870101525b50601f01601f1916929092016040019392505050565b80356001600160a01b038116811461076457600080fd5b919050565b6000806040838503121561077c57600080fd5b6107858361074d565b946020939093013593505050565b6000806000606084860312156107a857600080fd5b6107b18461074d565b92506107bf6020850161074d565b9150604084013590509250925092565b6000602082840312156107e157600080fd5b6107ea8261074d565b9392505050565b6000806040838503121561080457600080fd5b61080d8361074d565b915061081b6020840161074d565b90509250929050565b600181811c9082168061083857607f821691505b6020821081141561085957634e487b7160e01b600052602260045260246000fd5b50919050565b6000821982111561088057634e487b7160e01b600052601160045260246000fd5b50019056fea26469706673582212205520cebf912a7b37ad2078443ea206cc6187d18972e4366e20d773c71e326a0464736f6c63430008090033",
}

// TestTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TestTokenMetaData.ABI instead.
var TestTokenABI = TestTokenMetaData.ABI

// TestTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestTokenMetaData.Bin instead.
var TestTokenBin = TestTokenMetaData.Bin

// DeployTestToken deploys a new Ethereum contract, binding an instance of TestToken to it.
func DeployTestToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *TestToken, error) {
	parsed, err := TestTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestTokenBin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestToken{TestTokenCaller: TestTokenCaller{contract: contract}, TestTokenTransactor: TestTokenTransactor{contract: contract}, TestTokenFilterer: TestTokenFilterer{contract: contract}}, nil
}

// TestToken is an auto generated Go binding around an Ethereum contract.
type TestToken struct {
	TestTokenCaller     // Read-only binding to the contract
	TestTokenTransactor // Write-only binding to the contract
	TestTokenFilterer   // Log filterer for contract events
}

// TestTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestTokenSession struct {
	Contract     *TestToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestTokenCallerSession struct {
	Contract *TestTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TestTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTokenTransactorSession struct {
	Contract     *TestTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TestTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestTokenRaw struct {
	Contract *TestToken // Generic contract binding to access the raw methods on
}

// TestTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestTokenCallerRaw struct {
	Contract *TestTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TestTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTokenTransactorRaw struct {
	Contract *TestTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestToken creates a new instance of TestToken, bound to a specific deployed contract.
func NewTestToken(address common.Address, backend bind.ContractBackend) (*TestToken, error) {
	contract, err := bindTestToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestToken{TestTokenCaller: TestTokenCaller{contract: contract}, TestTokenTransactor: TestTokenTransactor{contract: contract}, TestTokenFilterer: TestTokenFilterer{contract: contract}}, nil
}

// NewTestTokenCaller creates a new read-only instance of TestToken, bound to a specific deployed contract.
func NewTestTokenCaller(address common.Address, caller bind.ContractCaller) (*TestTokenCaller, error) {
	contract, err := bindTestToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestTokenCaller{contract: contract}, nil
}

// NewTestTokenTransactor creates a new write-only instance of TestToken, bound to a specific deployed contract.
func NewTestTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTokenTransactor, error) {
	contract, err := bindTestToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTokenTransactor{contract: contract}, nil
}

// NewTestTokenFilterer creates a new log filterer instance of TestToken, bound to a specific deployed contract.
func NewTestTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TestTokenFilterer, error) {
	contract, err := bindTestToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestTokenFilterer{contract: contract}, nil
}

// bindTestToken binds a generic wrapper to an already deployed contract.
func bindTestToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestToken *TestTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestToken.Contract.TestTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestToken *TestTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestToken.Contract.TestTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestToken *TestTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestToken.Contract.TestTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestToken *TestTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestToken *TestTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestToken *TestTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestToken *TestTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestToken *TestTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestToken.Contract.Allowance(&_TestToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestToken *TestTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestToken.Contract.Allowance(&_TestToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestToken *TestTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestToken *TestTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestToken.Contract.BalanceOf(&_TestToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestToken *TestTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestToken.Contract.BalanceOf(&_TestToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestToken *TestTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestToken *TestTokenSession) Decimals() (uint8, error) {
	return _TestToken.Contract.Decimals(&_TestToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestToken *TestTokenCallerSession) Decimals() (uint8, error) {
	return _TestToken.Contract.Decimals(&_TestToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestToken *TestTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestToken *TestTokenSession) Name() (string, error) {
	return _TestToken.Contract.Name(&_TestToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestToken *TestTokenCallerSession) Name() (string, error) {
	return _TestToken.Contract.Name(&_TestToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestToken *TestTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestToken *TestTokenSession) Symbol() (string, error) {
	return _TestToken.Contract.Symbol(&_TestToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestToken *TestTokenCallerSession) Symbol() (string, error) {
	return _TestToken.Contract.Symbol(&_TestToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestToken *TestTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestToken *TestTokenSession) TotalSupply() (*big.Int, error) {
	return _TestToken.Contract.TotalSupply(&_TestToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestToken *TestTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _TestToken.Contract.TotalSupply(&_TestToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestToken *TestTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestToken *TestTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.Approve(&_TestToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestToken *TestTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.Approve(&_TestToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestToken *TestTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestToken *TestTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.DecreaseAllowance(&_TestToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestToken *TestTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.DecreaseAllowance(&_TestToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestToken *TestTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestToken *TestTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.IncreaseAllowance(&_TestToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestToken *TestTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.IncreaseAllowance(&_TestToken.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TestToken *TestTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TestToken *TestTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.Transfer(&_TestToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TestToken *TestTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.Transfer(&_TestToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TestToken *TestTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TestToken *TestTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.TransferFrom(&_TestToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TestToken *TestTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.TransferFrom(&_TestToken.TransactOpts, from, to, amount)
}

// TestTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TestToken contract.
type TestTokenApprovalIterator struct {
	Event *TestTokenApproval // Event containing the contract specifics and raw log

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
func (it *TestTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestTokenApproval)
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
		it.Event = new(TestTokenApproval)
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
func (it *TestTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestTokenApproval represents a Approval event raised by the TestToken contract.
type TestTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestToken *TestTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TestTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TestTokenApprovalIterator{contract: _TestToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestToken *TestTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TestTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestTokenApproval)
				if err := _TestToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestToken *TestTokenFilterer) ParseApproval(log types.Log) (*TestTokenApproval, error) {
	event := new(TestTokenApproval)
	if err := _TestToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TestToken contract.
type TestTokenTransferIterator struct {
	Event *TestTokenTransfer // Event containing the contract specifics and raw log

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
func (it *TestTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestTokenTransfer)
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
		it.Event = new(TestTokenTransfer)
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
func (it *TestTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestTokenTransfer represents a Transfer event raised by the TestToken contract.
type TestTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestToken *TestTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestTokenTransferIterator{contract: _TestToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestToken *TestTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TestTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestTokenTransfer)
				if err := _TestToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestToken *TestTokenFilterer) ParseTransfer(log types.Log) (*TestTokenTransfer, error) {
	event := new(TestTokenTransfer)
	if err := _TestToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValueArrayTesterMetaData contains all meta data concerning the ValueArrayTester contract.
var ValueArrayTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"test\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610731806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063f8a8fd6d14610030575b600080fd5b61003861003a565b005b6040805160026020820181815260808301845260009383929083015b604080518082019091526000808252602082015281526020019060019003908161005657505090528051519091506002146100c45760405162461bcd60e51b815260206004820152600960248201526829aa20a92a2fa622a760b91b60448201526064015b60405180910390fd5b6100db60006100d360016103c2565b8391906103f5565b6100ea60016100d360026103c2565b6100fe6100f760036103c2565b829061041c565b80515160031461013b5760405162461bcd60e51b8152602060048201526008602482015267282aa9a42fa622a760c11b60448201526064016100bb565b60005b81515181101561020c5760006101548383610510565b905060008151600681111561016b5761016b61066f565b146101a85760405162461bcd60e51b815260206004820152600d60248201526c505553485f56414c5f5459504560981b60448201526064016100bb565b6101b382600161069b565b8160200151146101f95760405162461bcd60e51b8152602060048201526011602482015270505553485f56414c5f434f4e54454e545360781b60448201526064016100bb565b5080610204816106b3565b91505061013e565b50600061021882610548565b905060008151600681111561022f5761022f61066f565b1461026b5760405162461bcd60e51b815260206004820152600c60248201526b504f505f5245545f5459504560a01b60448201526064016100bb565b80602001516003146102b25760405162461bcd60e51b815260206004820152601060248201526f504f505f5245545f434f4e54454e545360801b60448201526064016100bb565b8151516002146102ee5760405162461bcd60e51b81526020600482015260076024820152662827a82fa622a760c91b60448201526064016100bb565b60005b8251518110156103bd5760006103078483610510565b905060008151600681111561031e5761031e61066f565b1461035a5760405162461bcd60e51b815260206004820152600c60248201526b504f505f56414c5f5459504560a01b60448201526064016100bb565b61036582600161069b565b8160200151146103aa5760405162461bcd60e51b815260206004820152601060248201526f504f505f56414c5f434f4e54454e545360801b60448201526064016100bb565b50806103b5816106b3565b9150506102f1565b505050565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b808360000151838151811061040c5761040c6106ce565b6020026020010181905250505050565b81515160009061042d90600161069b565b67ffffffffffffffff81111561044557610445610659565b60405190808252806020026020018201604052801561048a57816020015b60408051808201909152600080825260208201528152602001906001900390816104635790505b50905060005b8351518110156104e65783518051829081106104ae576104ae6106ce565b60200260200101518282815181106104c8576104c86106ce565b602002602001018190525080806104de906106b3565b915050610490565b508181846000015151815181106104ff576104ff6106ce565b602090810291909101015290915250565b60408051808201909152600080825260208201528251805183908110610538576105386106ce565b6020026020010151905092915050565b60408051808201909152600080825260208201528151805161056c906001906106e4565b8151811061057c5761057c6106ce565b602002602001015190506000600183600001515161059a91906106e4565b67ffffffffffffffff8111156105b2576105b2610659565b6040519080825280602002602001820160405280156105f757816020015b60408051808201909152600080825260208201528152602001906001900390816105d05790505b50905060005b815181101561065257835180518290811061061a5761061a6106ce565b6020026020010151828281518110610634576106346106ce565b6020026020010181905250808061064a906106b3565b9150506105fd565b5090915290565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082198211156106ae576106ae610685565b500190565b60006000198214156106c7576106c7610685565b5060010190565b634e487b7160e01b600052603260045260246000fd5b6000828210156106f6576106f6610685565b50039056fea26469706673582212202acf49f5bbd26486ef855e2fe94f59b6f316d9b57dcfa18e0aaf50ebcb90bc9a64736f6c63430008090033",
}

// ValueArrayTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use ValueArrayTesterMetaData.ABI instead.
var ValueArrayTesterABI = ValueArrayTesterMetaData.ABI

// ValueArrayTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValueArrayTesterMetaData.Bin instead.
var ValueArrayTesterBin = ValueArrayTesterMetaData.Bin

// DeployValueArrayTester deploys a new Ethereum contract, binding an instance of ValueArrayTester to it.
func DeployValueArrayTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValueArrayTester, error) {
	parsed, err := ValueArrayTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValueArrayTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValueArrayTester{ValueArrayTesterCaller: ValueArrayTesterCaller{contract: contract}, ValueArrayTesterTransactor: ValueArrayTesterTransactor{contract: contract}, ValueArrayTesterFilterer: ValueArrayTesterFilterer{contract: contract}}, nil
}

// ValueArrayTester is an auto generated Go binding around an Ethereum contract.
type ValueArrayTester struct {
	ValueArrayTesterCaller     // Read-only binding to the contract
	ValueArrayTesterTransactor // Write-only binding to the contract
	ValueArrayTesterFilterer   // Log filterer for contract events
}

// ValueArrayTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValueArrayTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueArrayTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValueArrayTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueArrayTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValueArrayTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueArrayTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValueArrayTesterSession struct {
	Contract     *ValueArrayTester // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueArrayTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValueArrayTesterCallerSession struct {
	Contract *ValueArrayTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ValueArrayTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValueArrayTesterTransactorSession struct {
	Contract     *ValueArrayTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ValueArrayTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValueArrayTesterRaw struct {
	Contract *ValueArrayTester // Generic contract binding to access the raw methods on
}

// ValueArrayTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValueArrayTesterCallerRaw struct {
	Contract *ValueArrayTesterCaller // Generic read-only contract binding to access the raw methods on
}

// ValueArrayTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValueArrayTesterTransactorRaw struct {
	Contract *ValueArrayTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValueArrayTester creates a new instance of ValueArrayTester, bound to a specific deployed contract.
func NewValueArrayTester(address common.Address, backend bind.ContractBackend) (*ValueArrayTester, error) {
	contract, err := bindValueArrayTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValueArrayTester{ValueArrayTesterCaller: ValueArrayTesterCaller{contract: contract}, ValueArrayTesterTransactor: ValueArrayTesterTransactor{contract: contract}, ValueArrayTesterFilterer: ValueArrayTesterFilterer{contract: contract}}, nil
}

// NewValueArrayTesterCaller creates a new read-only instance of ValueArrayTester, bound to a specific deployed contract.
func NewValueArrayTesterCaller(address common.Address, caller bind.ContractCaller) (*ValueArrayTesterCaller, error) {
	contract, err := bindValueArrayTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValueArrayTesterCaller{contract: contract}, nil
}

// NewValueArrayTesterTransactor creates a new write-only instance of ValueArrayTester, bound to a specific deployed contract.
func NewValueArrayTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*ValueArrayTesterTransactor, error) {
	contract, err := bindValueArrayTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValueArrayTesterTransactor{contract: contract}, nil
}

// NewValueArrayTesterFilterer creates a new log filterer instance of ValueArrayTester, bound to a specific deployed contract.
func NewValueArrayTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*ValueArrayTesterFilterer, error) {
	contract, err := bindValueArrayTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValueArrayTesterFilterer{contract: contract}, nil
}

// bindValueArrayTester binds a generic wrapper to an already deployed contract.
func bindValueArrayTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValueArrayTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueArrayTester *ValueArrayTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValueArrayTester.Contract.ValueArrayTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueArrayTester *ValueArrayTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueArrayTester.Contract.ValueArrayTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueArrayTester *ValueArrayTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueArrayTester.Contract.ValueArrayTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueArrayTester *ValueArrayTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValueArrayTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueArrayTester *ValueArrayTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueArrayTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueArrayTester *ValueArrayTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueArrayTester.Contract.contract.Transact(opts, method, params...)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() pure returns()
func (_ValueArrayTester *ValueArrayTesterCaller) Test(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ValueArrayTester.contract.Call(opts, &out, "test")

	if err != nil {
		return err
	}

	return err

}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() pure returns()
func (_ValueArrayTester *ValueArrayTesterSession) Test() error {
	return _ValueArrayTester.Contract.Test(&_ValueArrayTester.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() pure returns()
func (_ValueArrayTester *ValueArrayTesterCallerSession) Test() error {
	return _ValueArrayTester.Contract.Test(&_ValueArrayTester.CallOpts)
}
