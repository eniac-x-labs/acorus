// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package challengegen

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

// ChallengeLibChallenge is an auto generated low-level Go binding around an user-defined struct.
type ChallengeLibChallenge struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}

// ChallengeLibParticipant is an auto generated low-level Go binding around an user-defined struct.
type ChallengeLibParticipant struct {
	Addr     common.Address
	TimeLeft *big.Int
}

// ChallengeLibSegmentSelection is an auto generated low-level Go binding around an user-defined struct.
type ChallengeLibSegmentSelection struct {
	OldSegmentsStart  *big.Int
	OldSegmentsLength *big.Int
	OldSegments       [][32]byte
	ChallengePosition *big.Int
}

// GlobalState is an auto generated low-level Go binding around an user-defined struct.
type GlobalState struct {
	Bytes32Vals [2][32]byte
	U64Vals     [2]uint64
}

// ChallengeLibMetaData contains all meta data concerning the ChallengeLib contract.
var ChallengeLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220be45a96ce6243b77a059183fe5e02a3439a1cffd58f2017cfb910ec79b578a3c64736f6c63430008090033",
}

// ChallengeLibABI is the input ABI used to generate the binding from.
// Deprecated: Use ChallengeLibMetaData.ABI instead.
var ChallengeLibABI = ChallengeLibMetaData.ABI

// ChallengeLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChallengeLibMetaData.Bin instead.
var ChallengeLibBin = ChallengeLibMetaData.Bin

// DeployChallengeLib deploys a new Ethereum contract, binding an instance of ChallengeLib to it.
func DeployChallengeLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeLib, error) {
	parsed, err := ChallengeLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChallengeLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeLib{ChallengeLibCaller: ChallengeLibCaller{contract: contract}, ChallengeLibTransactor: ChallengeLibTransactor{contract: contract}, ChallengeLibFilterer: ChallengeLibFilterer{contract: contract}}, nil
}

// ChallengeLib is an auto generated Go binding around an Ethereum contract.
type ChallengeLib struct {
	ChallengeLibCaller     // Read-only binding to the contract
	ChallengeLibTransactor // Write-only binding to the contract
	ChallengeLibFilterer   // Log filterer for contract events
}

// ChallengeLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeLibSession struct {
	Contract     *ChallengeLib     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeLibCallerSession struct {
	Contract *ChallengeLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ChallengeLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeLibTransactorSession struct {
	Contract     *ChallengeLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ChallengeLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeLibRaw struct {
	Contract *ChallengeLib // Generic contract binding to access the raw methods on
}

// ChallengeLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeLibCallerRaw struct {
	Contract *ChallengeLibCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeLibTransactorRaw struct {
	Contract *ChallengeLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeLib creates a new instance of ChallengeLib, bound to a specific deployed contract.
func NewChallengeLib(address common.Address, backend bind.ContractBackend) (*ChallengeLib, error) {
	contract, err := bindChallengeLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeLib{ChallengeLibCaller: ChallengeLibCaller{contract: contract}, ChallengeLibTransactor: ChallengeLibTransactor{contract: contract}, ChallengeLibFilterer: ChallengeLibFilterer{contract: contract}}, nil
}

// NewChallengeLibCaller creates a new read-only instance of ChallengeLib, bound to a specific deployed contract.
func NewChallengeLibCaller(address common.Address, caller bind.ContractCaller) (*ChallengeLibCaller, error) {
	contract, err := bindChallengeLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeLibCaller{contract: contract}, nil
}

// NewChallengeLibTransactor creates a new write-only instance of ChallengeLib, bound to a specific deployed contract.
func NewChallengeLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeLibTransactor, error) {
	contract, err := bindChallengeLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeLibTransactor{contract: contract}, nil
}

// NewChallengeLibFilterer creates a new log filterer instance of ChallengeLib, bound to a specific deployed contract.
func NewChallengeLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeLibFilterer, error) {
	contract, err := bindChallengeLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeLibFilterer{contract: contract}, nil
}

// bindChallengeLib binds a generic wrapper to an already deployed contract.
func bindChallengeLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChallengeLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeLib *ChallengeLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeLib.Contract.ChallengeLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeLib *ChallengeLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeLib.Contract.ChallengeLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeLib *ChallengeLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeLib.Contract.ChallengeLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeLib *ChallengeLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeLib *ChallengeLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeLib *ChallengeLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeLib.Contract.contract.Transact(opts, method, params...)
}

// ChallengeManagerMetaData contains all meta data concerning the ChallengeManager contract.
var ChallengeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"challengeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"chainHashes\",\"type\":\"bytes32[]\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIChallengeManager.ChallengeTerminationType\",\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"ChallengeEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockSteps\",\"type\":\"uint256\"}],\"name\":\"ExecutionChallengeBegun\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"startState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"endState\",\"type\":\"tuple\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"newSegments\",\"type\":\"bytes32[]\"}],\"name\":\"bisectExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"machineStatuses\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"globalStateHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256\",\"name\":\"numSteps\",\"type\":\"uint256\"}],\"name\":\"challengeExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"challengeInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"current\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastMoveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessages\",\"type\":\"uint64\"},{\"internalType\":\"enumChallengeLib.ChallengeMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"internalType\":\"structChallengeLib.Challenge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"challenges\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"current\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastMoveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessages\",\"type\":\"uint64\"},{\"internalType\":\"enumChallengeLib.ChallengeMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"clearChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"startAndEndMachineStatuses_\",\"type\":\"uint8[2]\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState[2]\",\"name\":\"startAndEndGlobalStates_\",\"type\":\"tuple[2]\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"asserter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"asserterTimeLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengerTimeLeft_\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"currentResponder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIChallengeResultReceiver\",\"name\":\"resultReceiver_\",\"type\":\"address\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"sequencerInbox_\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"osp_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"isTimedOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProveExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"osp\",\"outputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"osp_\",\"type\":\"address\"}],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resultReceiver\",\"outputs\":[{\"internalType\":\"contractIChallengeResultReceiver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalChallengesCreated\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5060805161312f61003760003960008181610e02015261130d015261312f6000f3fe608060405234801561001057600080fd5b50600436106100eb5760003560e01c80639ede42b9116100925780639ede42b91461025c578063a521b0321461027f578063c474d2c514610292578063d248d124146102a5578063e78cea92146102b8578063ee35f327146102cb578063f26a62c6146102de578063f8c8765e146102f1578063fb7be0a11461030457600080fd5b806314eab5e7146100f05780631b45c86a1461012057806323a9ef23146101355780633504f1d71461016057806356e9df97146101735780635ef489e6146101865780637fd07a9c146101995780638f1d3776146101b9575b600080fd5b6101036100fe36600461269b565b610317565b6040516001600160401b0390911681526020015b60405180910390f35b61013361012e36600461272e565b61061f565b005b61014861014336600461272e565b6106ef565b6040516001600160a01b039091168152602001610117565b600254610148906001600160a01b031681565b61013361018136600461272e565b610713565b600054610103906001600160401b031681565b6101ac6101a736600461272e565b610881565b604051610117919061278b565b6102496101c73660046127fd565b6001602081815260009283526040928390208351808501855281546001600160a01b0390811682529382015481840152845180860190955260028201549093168452600381015491840191909152600481015460058201546006830154600790930154939493919290916001600160401b03811690600160401b900460ff1687565b6040516101179796959493929190612816565b61026f61026a36600461272e565b61095a565b6040519015158152602001610117565b61013361028d366004612873565b610981565b6101336102a0366004612917565b610df7565b6101336102b3366004612934565b610ec1565b600454610148906001600160a01b031681565b600354610148906001600160a01b031681565b600554610148906001600160a01b031681565b6101336102ff3660046129c6565b611302565b610133610312366004612a22565b61142e565b6002546000906001600160a01b0316331461036c5760405162461bcd60e51b815260206004820152601060248201526f13d3931657d493d313155417d0d2105360821b60448201526064015b60405180910390fd5b6040805160028082526060820183526000926020830190803683370190505090506103c261039d60208b018b612ac6565b6103bd8a60005b608002018036038101906103b89190612b85565b611a9e565b611b1f565b816000815181106103d5576103d5612ab0565b60209081029190910101526104048960016020020160208101906103f99190612ac6565b6103bd8a60016103a4565b8160018151811061041757610417612ab0565b6020908102919091010152600080548190819061043c906001600160401b0316612c33565b91906101000a8154816001600160401b0302191690836001600160401b031602179055905060006001600160401b0316816001600160401b0316141561048457610484612c5a565b6001600160401b0381166000908152600160205260408120600581018d9055906104be6104b9368d90038d0160808e01612b85565b611c43565b905060026104d260408e0160208f01612ac6565b60038111156104e3576104e3612761565b148061051157506000610506610501368e90038e0160808f01612b85565b611c58565b6001600160401b0316115b15610524578061052081612c33565b9150505b6007820180546040805180820182526001600160a01b038d811680835260209283018d90526002880180546001600160a01b03199081169092179055600388018d905583518085018552918e16808352919092018b90528654909116178555600185018990554260048601556001600160401b0384811668ffffffffffffffffff1990931692909217600160401b179092559051908416907f76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a906105ee908e906080820190612cba565b60405180910390a261060c8360008c6001600160401b031687611c67565b5090925050505b98975050505050505050565b60006001600160401b038216600090815260016020526040902060070154600160401b900460ff16600281111561065857610658612761565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906106985760405162461bcd60e51b81526004016103639190612cd6565b506106a28161095a565b6106e15760405162461bcd60e51b815260206004820152601060248201526f54494d454f55545f444541444c494e4560801b6044820152606401610363565b6106ec816000611cfd565b50565b6001600160401b03166000908152600160205260409020546001600160a01b031690565b6002546001600160a01b031633146107605760405162461bcd60e51b815260206004820152601060248201526f2727aa2fa922a9afa922a1a2a4ab22a960811b6044820152606401610363565b60006001600160401b038216600090815260016020526040902060070154600160401b900460ff16600281111561079957610799612761565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906107d95760405162461bcd60e51b81526004016103639190612cd6565b506001600160401b038116600081815260016020819052604080832080546001600160a01b031990811682559281018490556002810180549093169092556003808301849055600483018490556005830184905560068301939093556007909101805468ffffffffffffffffff19169055517ffdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f409161087691612d2b565b60405180910390a250565b6108896125f6565b6001600160401b0382811660009081526001602081815260409283902083516101208101855281546001600160a01b0390811660e0830190815294830154610100830152938152845180860186526002808401549095168152600383015481850152928101929092526004810154938201939093526005830154606082015260068301546080820152600783015493841660a08201529260c0840191600160401b90910460ff169081111561094057610940612761565b600281111561095157610951612761565b90525092915050565b6001600160401b038116600090815260016020526040812061097b90611e2b565b92915050565b6001600160401b0384166000908152600160205260408120859185916109a6846106ef565b6001600160a01b0316336001600160a01b0316146109d65760405162461bcd60e51b815260040161036390612d45565b6109df8461095a565b156109fc5760405162461bcd60e51b815260040161036390612d6a565b6000826002811115610a1057610a10612761565b1415610a7e5760006007820154600160401b900460ff166002811115610a3857610a38612761565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b81525090610a785760405162461bcd60e51b81526004016103639190612cd6565b50610b3d565b6001826002811115610a9257610a92612761565b1415610adc5760016007820154600160401b900460ff166002811115610aba57610aba612761565b14610ad75760405162461bcd60e51b815260040161036390612d91565b610b3d565b6002826002811115610af057610af0612761565b1415610b355760026007820154600160401b900460ff166002811115610b1857610b18612761565b14610ad75760405162461bcd60e51b815260040161036390612db9565b610b3d612c5a565b610b8b83356020850135610b546040870187612de5565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250611e4392505050565b816006015414610bad5760405162461bcd60e51b815260040161036390612e35565b6002610bbc6040850185612de5565b90501080610be757506001610bd46040850185612de5565b610bdf929150612e58565b836060013510155b15610c045760405162461bcd60e51b815260040161036390612e6f565b600080610c1089611e7a565b9150915060018111610c505760405162461bcd60e51b81526020600482015260096024820152681513d3d7d4d213d49560ba1b6044820152606401610363565b806028811115610c5e575060285b610c69816001612e9a565b8814610ca65760405162461bcd60e51b815260206004820152600c60248201526b57524f4e475f44454752454560a01b6044820152606401610363565b50610cf08989896000818110610cbe57610cbe612ab0565b602002919091013590508a8a610cd5600182612e58565b818110610ce457610ce4612ab0565b90506020020135611f0b565b610d2f8a83838b8b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250611c6792505050565b50600090505b6007820154600160401b900460ff166002811115610d5557610d55612761565b1415610d615750610dee565b6040805180820190915281546001600160a01b03168152600182015460208201526004820154610d919042612e58565b81602001818151610da29190612e58565b90525060028201805483546001600160a01b038083166001600160a01b031992831617865560038601805460018801558551929093169116179091556020909101519055426004909101555b50505050505050565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161415610e405760405162461bcd60e51b815260040161036390612eb2565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61038054336001600160a01b03821614610e9d57604051631194af8760e11b81523360048201526001600160a01b0382166024820152604401610363565b5050600580546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160401b038416600090815260016020526040902084908490600290610ee9846106ef565b6001600160a01b0316336001600160a01b031614610f195760405162461bcd60e51b815260040161036390612d45565b610f228461095a565b15610f3f5760405162461bcd60e51b815260040161036390612d6a565b6000826002811115610f5357610f53612761565b1415610fc15760006007820154600160401b900460ff166002811115610f7b57610f7b612761565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b81525090610fbb5760405162461bcd60e51b81526004016103639190612cd6565b50611080565b6001826002811115610fd557610fd5612761565b141561101f5760016007820154600160401b900460ff166002811115610ffd57610ffd612761565b1461101a5760405162461bcd60e51b815260040161036390612d91565b611080565b600282600281111561103357611033612761565b14156110785760026007820154600160401b900460ff16600281111561105b5761105b612761565b1461101a5760405162461bcd60e51b815260040161036390612db9565b611080612c5a565b61109783356020850135610b546040870187612de5565b8160060154146110b95760405162461bcd60e51b815260040161036390612e35565b60026110c86040850185612de5565b905010806110f3575060016110e06040850185612de5565b6110eb929150612e58565b836060013510155b156111105760405162461bcd60e51b815260040161036390612e6f565b6001600160401b038816600090815260016020526040812090806111338a611e7a565b9092509050600181146111585760405162461bcd60e51b815260040161036390612efe565b5060055460408051808201825260078501546001600160401b031681526004546001600160a01b0390811660208301526000931691635d3adcfb919085906111a2908f018f612de5565b8f606001358181106111b6576111b6612ab0565b905060200201358d8d6040518663ffffffff1660e01b81526004016111df959493929190612f20565b60206040518083038186803b1580156111f757600080fd5b505afa15801561120b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061122f9190612f77565b905061123e60408b018b612de5565b61124d60608d01356001612e9a565b81811061125c5761125c612ab0565b905060200201358114156112a15760405162461bcd60e51b815260206004820152600c60248201526b14d0535157d3d4d417d1539160a21b6044820152606401610363565b6040516001600160401b038c16907fc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e90600090a26112f68b6001600160401b0316600090815260016020526040812060060155565b5060009150610d359050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016141561134b5760405162461bcd60e51b815260040161036390612eb2565b6002546001600160a01b0316156113935760405162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b6044820152606401610363565b6001600160a01b0384166113de5760405162461bcd60e51b81526020600482015260126024820152712727afa922a9aaa62a2fa922a1a2a4ab22a960711b6044820152606401610363565b600280546001600160a01b039586166001600160a01b0319918216179091556003805494861694821694909417909355600480549285169284169290921790915560058054919093169116179055565b6001600160401b038516600090815260016020819052604090912086918691611456846106ef565b6001600160a01b0316336001600160a01b0316146114865760405162461bcd60e51b815260040161036390612d45565b61148f8461095a565b156114ac5760405162461bcd60e51b815260040161036390612d6a565b60008260028111156114c0576114c0612761565b141561152e5760006007820154600160401b900460ff1660028111156114e8576114e8612761565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906115285760405162461bcd60e51b81526004016103639190612cd6565b506115ed565b600182600281111561154257611542612761565b141561158c5760016007820154600160401b900460ff16600281111561156a5761156a612761565b146115875760405162461bcd60e51b815260040161036390612d91565b6115ed565b60028260028111156115a0576115a0612761565b14156115e55760026007820154600160401b900460ff1660028111156115c8576115c8612761565b146115875760405162461bcd60e51b815260040161036390612db9565b6115ed612c5a565b61160483356020850135610b546040870187612de5565b8160060154146116265760405162461bcd60e51b815260040161036390612e35565b60026116356040850185612de5565b905010806116605750600161164d6040850185612de5565b611658929150612e58565b836060013510155b1561167d5760405162461bcd60e51b815260040161036390612e6f565b60018510156116c45760405162461bcd60e51b815260206004820152601360248201527210d2105313115391d157d513d3d7d4d213d495606a1b6044820152606401610363565b6508000000000085111561170f5760405162461bcd60e51b81526020600482015260126024820152714348414c4c454e47455f544f4f5f4c4f4e4760701b6044820152606401610363565b6117518861173161172360208b018b612ac6565b8960005b6020020135611b1f565b61174c61174460408c0160208d01612ac6565b8a6001611727565b611f0b565b6001600160401b038916600090815260016020526040812090806117748b611e7a565b91509150806001146117985760405162461bcd60e51b815260040161036390612efe565b60016117a760208c018c612ac6565b60038111156117b8576117b8612761565b14611872576117cd60408b0160208c01612ac6565b60038111156117de576117de612761565b6117eb60208c018c612ac6565b60038111156117fc576117fc612761565b14801561180d5750883560208a0135145b6118495760405162461bcd60e51b815260206004820152600d60248201526c48414c5445445f4348414e474560981b6044820152606401610363565b61186a8c6001600160401b0316600090815260016020526040812060060155565b5050506119d9565b600261188460408c0160208d01612ac6565b600381111561189557611895612761565b14156118de57883560208a0135146118de5760405162461bcd60e51b815260206004820152600c60248201526b4552524f525f4348414e474560a01b6044820152606401610363565b6040805160028082526060820183526000926020830190803683375050506005850154909150611910908b3590611fe0565b8160008151811061192357611923612ab0565b60209081029190910101526119518b60016020020160208101906119479190612ac6565b60208c013561216c565b8160018151811061196457611964612ab0565b602090810291909101015260078401805460ff60401b1916600160411b1790556119918d60008b84611c67565b8c6001600160401b03167f24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db846040516119cc91815260200190565b60405180910390a2505050505b60006007820154600160401b900460ff1660028111156119fb576119fb612761565b1415611a075750611a94565b6040805180820190915281546001600160a01b03168152600182015460208201526004820154611a379042612e58565b81602001818151611a489190612e58565b90525060028201805483546001600160a01b038083166001600160a01b031992831617865560038601805460018801558551929093169116179091556020909101519055426004909101555b5050505050505050565b80518051602091820151828401518051908401516040516c23b637b130b61039ba30ba329d60991b95810195909552602d850193909352604d8401919091526001600160c01b031960c091821b8116606d85015291901b166075820152600090607d015b604051602081830303815290604052805190602001209050919050565b60006001836003811115611b3557611b35612761565b1415611b7b576040516b213637b1b59039ba30ba329d60a11b6020820152602c8101839052604c015b60405160208183030381529060405280519060200120905061097b565b6002836003811115611b8f57611b8f612761565b1415611bc55760405174213637b1b59039ba30ba32961032b93937b932b21d60591b602082015260358101839052605501611b5e565b6003836003811115611bd957611bd9612761565b1415611c085760405174213637b1b59039ba30ba3296103a37b7903330b91d60591b6020820152603501611b5e565b60405162461bcd60e51b815260206004820152601060248201526f4241445f424c4f434b5f53544154555360801b6044820152606401610363565b6020810151600090815b602002015192915050565b60208101516000906001611c4d565b6001821015611c7857611c78612c5a565b600281511015611c8a57611c8a612c5a565b6000611c97848484611e43565b6001600160401b038616600081815260016020526040908190206006018390555191925082917f86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d6890611cee90889088908890612f90565b60405180910390a35050505050565b6001600160401b03821660008181526001602081905260408083206002808201805483546001600160a01b0319808216865596850188905595811690915560038301869055600480840187905560058401879055600684019690965560078301805468ffffffffffffffffff1916905590549251630357aa4960e01b8152948501959095526001600160a01b03948516602485018190529285166044850181905290949293909290911690630357aa4990606401600060405180830381600087803b158015611dcb57600080fd5b505af1158015611ddf573d6000803e3d6000fd5b50505050846001600160401b03167ffdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f4085604051611e1c9190612d2b565b60405180910390a25050505050565b6001810154600090611e3c83612215565b1192915050565b6000838383604051602001611e5a93929190612fe5565b6040516020818303038152906040528051906020012090505b9392505050565b600080806001611e8d6040860186612de5565b611e98929150612e58565b9050611ea881602086013561303d565b9150611eb8606085013583613051565b611ec3908535612e9a565b92506002611ed46040860186612de5565b611edf929150612e58565b84606001351415611f0557611ef8816020860135613070565b611f029083612e9a565b91505b50915091565b81611f196040850185612de5565b8560600135818110611f2d57611f2d612ab0565b9050602002013514611f6f5760405162461bcd60e51b815260206004820152600b60248201526a15d493d391d7d4d510549560aa1b6044820152606401610363565b80611f7d6040850185612de5565b611f8c60608701356001612e9a565b818110611f9b57611f9b612ab0565b905060200201351415611fdb5760405162461bcd60e51b815260206004820152600860248201526714d0535157d1539160c21b6044820152606401610363565b505050565b60408051600380825260808201909252600091829190816020015b6040805180820190915260008082526020820152815260200190600190039081611ffb57505060408051808201825260008082526020918201819052825180840190935260048352908201529091508160008151811061205d5761205d612ab0565b60200260200101819052506120726000612227565b8160018151811061208557612085612ab0565b602002602001018190525061209a6000612227565b816002815181106120ad576120ad612ab0565b60209081029190910181019190915260408051808301825283815281518083019092528082526000928201929092526120fd60408051606080820183529181019182529081526000602082015290565b604080518082018252606080825260006020808401829052845161012081018652828152908101879052938401859052908301829052608083018a905260a0830181905260c0830181905260e083015261010082018890529061215f8161225a565b9998505050505050505050565b6000600183600381111561218257612182612761565b14156121995781604051602001611b5e9190613084565b60028360038111156121ad576121ad612761565b14156121d7576040516f26b0b1b434b7329032b93937b932b21d60811b6020820152603001611b5e565b60038360038111156121eb576121eb612761565b1415611c08576040516f26b0b1b434b732903a37b7903330b91d60811b6020820152603001611b5e565b600081600401544261097b9190612e58565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b6000808251600381111561227057612270612761565b1415612326576122838260200151612413565b6122908360400151612413565b61229d8460600151612498565b608085015160a086015160c087015160e0808901516101008a01516040516f26b0b1b434b73290393ab73734b7339d60811b602082015260308101999099526050890197909752607088019590955260908701939093526001600160e01b031991831b821660b0870152821b811660b486015291901b1660b883015260bc82015260dc01611b02565b60018251600381111561233b5761233b612761565b1415612356578160800151604051602001611b029190613084565b60028251600381111561236b5761236b612761565b1415612395576040516f26b0b1b434b7329032b93937b932b21d60811b6020820152603001611b02565b6003825160038111156123aa576123aa612761565b14156123d4576040516f26b0b1b434b732903a37b7903330b91d60811b6020820152603001611b02565b60405162461bcd60e51b815260206004820152600f60248201526e4241445f4d4143485f53544154555360881b6044820152606401610363565b919050565b60208101518151515160005b8181101561249157835161243c906124379083612531565b612569565b6040516b2b30b63ab29039ba30b1b59d60a11b6020820152602c810191909152604c8101849052606c016040516020818303038152906040528051906020012092508080612489906130a9565b91505061241f565b5050919050565b602081015160005b82515181101561252b576124d0836000015182815181106124c3576124c3612ab0565b6020026020010151612586565b6040517129ba30b1b590333930b6b29039ba30b1b59d60711b60208201526032810191909152605281018390526072016040516020818303038152906040528051906020012091508080612523906130a9565b9150506124a0565b50919050565b6040805180820190915260008082526020820152825180518390811061255957612559612ab0565b6020026020010151905092915050565b600081600001518260200151604051602001611b029291906130c4565b60006125958260000151612569565b602080840151604080860151606087015191516b29ba30b1b590333930b6b29d60a11b94810194909452602c840194909452604c8301919091526001600160e01b031960e093841b8116606c840152921b9091166070820152607401611b02565b604080516101208101909152600060e0820181815261010083019190915281908152602001612635604080518082019091526000808252602082015290565b815260006020820181905260408201819052606082018190526080820181905260a09091015290565b806040810183101561097b57600080fd5b80356001600160401b038116811461240e57600080fd5b6001600160a01b03811681146106ec57600080fd5b600080600080600080600080610200898b0312156126b857600080fd5b883597506126c98a60208b0161265e565b965061016089018a8111156126dd57600080fd5b60608a0196506126ec8161266f565b9550506101808901356126fe81612686565b93506101a089013561270f81612686565b979a96995094979396929592945050506101c0820135916101e0013590565b60006020828403121561274057600080fd5b611e738261266f565b80516001600160a01b03168252602090810151910152565b634e487b7160e01b600052602160045260246000fd5b6003811061278757612787612761565b9052565b60006101208201905061279f828451612749565b60208301516127b16040840182612749565b5060408301516080830152606083015160a0830152608083015160c08301526001600160401b0360a08401511660e083015260c08301516127f6610100840182612777565b5092915050565b60006020828403121561280f57600080fd5b5035919050565b6101208101612825828a612749565b6128326040830189612749565b8660808301528560a08301528460c08301526001600160401b03841660e0830152610613610100830184612777565b60006080828403121561252b57600080fd5b6000806000806060858703121561288957600080fd5b6128928561266f565b935060208501356001600160401b03808211156128ae57600080fd5b6128ba88838901612861565b945060408701359150808211156128d057600080fd5b818701915087601f8301126128e457600080fd5b8135818111156128f357600080fd5b8860208260051b850101111561290857600080fd5b95989497505060200194505050565b60006020828403121561292957600080fd5b8135611e7381612686565b6000806000806060858703121561294a57600080fd5b6129538561266f565b935060208501356001600160401b038082111561296f57600080fd5b61297b88838901612861565b9450604087013591508082111561299157600080fd5b818701915087601f8301126129a557600080fd5b8135818111156129b457600080fd5b88602082850101111561290857600080fd5b600080600080608085870312156129dc57600080fd5b84356129e781612686565b935060208501356129f781612686565b92506040850135612a0781612686565b91506060850135612a1781612686565b939692955090935050565b600080600080600060e08688031215612a3a57600080fd5b612a438661266f565b945060208601356001600160401b03811115612a5e57600080fd5b612a6a88828901612861565b945050612a7a876040880161265e565b9250612a89876080880161265e565b9497939650919460c0013592915050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600060208284031215612ad857600080fd5b813560048110611e7357600080fd5b604080519081016001600160401b0381118282101715612b0957612b09612a9a565b60405290565b600082601f830112612b2057600080fd5b604051604081018181106001600160401b0382111715612b4257612b42612a9a565b8060405250806040840185811115612b5957600080fd5b845b81811015612b7a57612b6c8161266f565b835260209283019201612b5b565b509195945050505050565b600060808284031215612b9757600080fd5b604051604081018181106001600160401b0382111715612bb957612bb9612a9a565b604052601f83018413612bcb57600080fd5b612bd3612ae7565b806040850186811115612be557600080fd5b855b81811015612bff578035845260209384019301612be7565b50818452612c0d8782612b0f565b6020850152509195945050505050565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b0380831681811415612c5057612c50612c1d565b6001019392505050565b634e487b7160e01b600052600160045260246000fd5b604081833760006040838101828152908301915b6002811015612cb3576001600160401b03612c9e8461266f565b16825260209283019290910190600101612c84565b5050505050565b6101008101612cc98285612c70565b611e736080830184612c70565b600060208083528351808285015260005b81811015612d0357858101830151858201604001528201612ce7565b81811115612d15576000604083870101525b50601f01601f1916929092016040019392505050565b6020810160048310612d3f57612d3f612761565b91905290565b6020808252600b908201526a21a420a62fa9a2a72222a960a91b604082015260600190565b6020808252600d908201526c4348414c5f444541444c494e4560981b604082015260600190565b6020808252600e908201526d4348414c5f4e4f545f424c4f434b60901b604082015260600190565b60208082526012908201527121a420a62fa727aa2fa2ac22a1aaaa24a7a760711b604082015260600190565b6000808335601e19843603018112612dfc57600080fd5b8301803591506001600160401b03821115612e1657600080fd5b6020019150600581901b3603821315612e2e57600080fd5b9250929050565b6020808252600990820152684249535f535441544560b81b604082015260600190565b600082821015612e6a57612e6a612c1d565b500390565b6020808252601190820152704241445f4348414c4c454e47455f504f5360781b604082015260600190565b60008219821115612ead57612ead612c1d565b500190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b602080825260089082015267544f4f5f4c4f4e4760c01b604082015260600190565b8551815260018060a01b03602087015116602082015284604082015283606082015260a060808201528160a0820152818360c0830137600081830160c090810191909152601f909201601f19160101949350505050565b600060208284031215612f8957600080fd5b5051919050565b6000606082018583526020858185015260606040850152818551808452608086019150828701935060005b81811015612fd757845183529383019391830191600101612fbb565b509098975050505050505050565b83815260006020848184015260408301845182860160005b8281101561301957815184529284019290840190600101612ffd565b509198975050505050505050565b634e487b7160e01b600052601260045260246000fd5b60008261304c5761304c613027565b500490565b600081600019048311821515161561306b5761306b612c1d565b500290565b60008261307f5761307f613027565b500690565b7026b0b1b434b732903334b734b9b432b21d60791b8152601181019190915260310190565b60006000198214156130bd576130bd612c1d565b5060010190565b652b30b63ab29d60d11b81526000600784106130e2576130e2612761565b5060f89290921b600683015260078201526027019056fea26469706673582212204e9beb8667db81401ddf0d2848d4f1418bb69fc687e03f3e46f3c6878310471464736f6c63430008090033",
}

// ChallengeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ChallengeManagerMetaData.ABI instead.
var ChallengeManagerABI = ChallengeManagerMetaData.ABI

// ChallengeManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChallengeManagerMetaData.Bin instead.
var ChallengeManagerBin = ChallengeManagerMetaData.Bin

// DeployChallengeManager deploys a new Ethereum contract, binding an instance of ChallengeManager to it.
func DeployChallengeManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeManager, error) {
	parsed, err := ChallengeManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChallengeManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeManager{ChallengeManagerCaller: ChallengeManagerCaller{contract: contract}, ChallengeManagerTransactor: ChallengeManagerTransactor{contract: contract}, ChallengeManagerFilterer: ChallengeManagerFilterer{contract: contract}}, nil
}

// ChallengeManager is an auto generated Go binding around an Ethereum contract.
type ChallengeManager struct {
	ChallengeManagerCaller     // Read-only binding to the contract
	ChallengeManagerTransactor // Write-only binding to the contract
	ChallengeManagerFilterer   // Log filterer for contract events
}

// ChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeManagerSession struct {
	Contract     *ChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeManagerCallerSession struct {
	Contract *ChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeManagerTransactorSession struct {
	Contract     *ChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeManagerRaw struct {
	Contract *ChallengeManager // Generic contract binding to access the raw methods on
}

// ChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeManagerCallerRaw struct {
	Contract *ChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeManagerTransactorRaw struct {
	Contract *ChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeManager creates a new instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManager(address common.Address, backend bind.ContractBackend) (*ChallengeManager, error) {
	contract, err := bindChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeManager{ChallengeManagerCaller: ChallengeManagerCaller{contract: contract}, ChallengeManagerTransactor: ChallengeManagerTransactor{contract: contract}, ChallengeManagerFilterer: ChallengeManagerFilterer{contract: contract}}, nil
}

// NewChallengeManagerCaller creates a new read-only instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*ChallengeManagerCaller, error) {
	contract, err := bindChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerCaller{contract: contract}, nil
}

// NewChallengeManagerTransactor creates a new write-only instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeManagerTransactor, error) {
	contract, err := bindChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerTransactor{contract: contract}, nil
}

// NewChallengeManagerFilterer creates a new log filterer instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeManagerFilterer, error) {
	contract, err := bindChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerFilterer{contract: contract}, nil
}

// bindChallengeManager binds a generic wrapper to an already deployed contract.
func bindChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChallengeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManager *ChallengeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeManager.Contract.ChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManager *ChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManager *ChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManager *ChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManager *ChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManager *ChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_ChallengeManager *ChallengeManagerCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_ChallengeManager *ChallengeManagerSession) Bridge() (common.Address, error) {
	return _ChallengeManager.Contract.Bridge(&_ChallengeManager.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_ChallengeManager *ChallengeManagerCallerSession) Bridge() (common.Address, error) {
	return _ChallengeManager.Contract.Bridge(&_ChallengeManager.CallOpts)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_ChallengeManager *ChallengeManagerCaller) ChallengeInfo(opts *bind.CallOpts, challengeIndex uint64) (ChallengeLibChallenge, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "challengeInfo", challengeIndex)

	if err != nil {
		return *new(ChallengeLibChallenge), err
	}

	out0 := *abi.ConvertType(out[0], new(ChallengeLibChallenge)).(*ChallengeLibChallenge)

	return out0, err

}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_ChallengeManager *ChallengeManagerSession) ChallengeInfo(challengeIndex uint64) (ChallengeLibChallenge, error) {
	return _ChallengeManager.Contract.ChallengeInfo(&_ChallengeManager.CallOpts, challengeIndex)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_ChallengeManager *ChallengeManagerCallerSession) ChallengeInfo(challengeIndex uint64) (ChallengeLibChallenge, error) {
	return _ChallengeManager.Contract.ChallengeInfo(&_ChallengeManager.CallOpts, challengeIndex)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_ChallengeManager *ChallengeManagerCaller) Challenges(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "challenges", arg0)

	outstruct := new(struct {
		Current            ChallengeLibParticipant
		Next               ChallengeLibParticipant
		LastMoveTimestamp  *big.Int
		WasmModuleRoot     [32]byte
		ChallengeStateHash [32]byte
		MaxInboxMessages   uint64
		Mode               uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Current = *abi.ConvertType(out[0], new(ChallengeLibParticipant)).(*ChallengeLibParticipant)
	outstruct.Next = *abi.ConvertType(out[1], new(ChallengeLibParticipant)).(*ChallengeLibParticipant)
	outstruct.LastMoveTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.WasmModuleRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.ChallengeStateHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.MaxInboxMessages = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.Mode = *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return *outstruct, err

}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_ChallengeManager *ChallengeManagerSession) Challenges(arg0 *big.Int) (struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	return _ChallengeManager.Contract.Challenges(&_ChallengeManager.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_ChallengeManager *ChallengeManagerCallerSession) Challenges(arg0 *big.Int) (struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	return _ChallengeManager.Contract.Challenges(&_ChallengeManager.CallOpts, arg0)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_ChallengeManager *ChallengeManagerCaller) CurrentResponder(opts *bind.CallOpts, challengeIndex uint64) (common.Address, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "currentResponder", challengeIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_ChallengeManager *ChallengeManagerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _ChallengeManager.Contract.CurrentResponder(&_ChallengeManager.CallOpts, challengeIndex)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_ChallengeManager *ChallengeManagerCallerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _ChallengeManager.Contract.CurrentResponder(&_ChallengeManager.CallOpts, challengeIndex)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_ChallengeManager *ChallengeManagerCaller) IsTimedOut(opts *bind.CallOpts, challengeIndex uint64) (bool, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "isTimedOut", challengeIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_ChallengeManager *ChallengeManagerSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _ChallengeManager.Contract.IsTimedOut(&_ChallengeManager.CallOpts, challengeIndex)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_ChallengeManager *ChallengeManagerCallerSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _ChallengeManager.Contract.IsTimedOut(&_ChallengeManager.CallOpts, challengeIndex)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_ChallengeManager *ChallengeManagerCaller) Osp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "osp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_ChallengeManager *ChallengeManagerSession) Osp() (common.Address, error) {
	return _ChallengeManager.Contract.Osp(&_ChallengeManager.CallOpts)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_ChallengeManager *ChallengeManagerCallerSession) Osp() (common.Address, error) {
	return _ChallengeManager.Contract.Osp(&_ChallengeManager.CallOpts)
}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_ChallengeManager *ChallengeManagerCaller) ResultReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "resultReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_ChallengeManager *ChallengeManagerSession) ResultReceiver() (common.Address, error) {
	return _ChallengeManager.Contract.ResultReceiver(&_ChallengeManager.CallOpts)
}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_ChallengeManager *ChallengeManagerCallerSession) ResultReceiver() (common.Address, error) {
	return _ChallengeManager.Contract.ResultReceiver(&_ChallengeManager.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_ChallengeManager *ChallengeManagerCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_ChallengeManager *ChallengeManagerSession) SequencerInbox() (common.Address, error) {
	return _ChallengeManager.Contract.SequencerInbox(&_ChallengeManager.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_ChallengeManager *ChallengeManagerCallerSession) SequencerInbox() (common.Address, error) {
	return _ChallengeManager.Contract.SequencerInbox(&_ChallengeManager.CallOpts)
}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_ChallengeManager *ChallengeManagerCaller) TotalChallengesCreated(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "totalChallengesCreated")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_ChallengeManager *ChallengeManagerSession) TotalChallengesCreated() (uint64, error) {
	return _ChallengeManager.Contract.TotalChallengesCreated(&_ChallengeManager.CallOpts)
}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_ChallengeManager *ChallengeManagerCallerSession) TotalChallengesCreated() (uint64, error) {
	return _ChallengeManager.Contract.TotalChallengesCreated(&_ChallengeManager.CallOpts)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_ChallengeManager *ChallengeManagerTransactor) BisectExecution(opts *bind.TransactOpts, challengeIndex uint64, selection ChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "bisectExecution", challengeIndex, selection, newSegments)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_ChallengeManager *ChallengeManagerSession) BisectExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.BisectExecution(&_ChallengeManager.TransactOpts, challengeIndex, selection, newSegments)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) BisectExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.BisectExecution(&_ChallengeManager.TransactOpts, challengeIndex, selection, newSegments)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_ChallengeManager *ChallengeManagerTransactor) ChallengeExecution(opts *bind.TransactOpts, challengeIndex uint64, selection ChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "challengeExecution", challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_ChallengeManager *ChallengeManagerSession) ChallengeExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeExecution(&_ChallengeManager.TransactOpts, challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) ChallengeExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeExecution(&_ChallengeManager.TransactOpts, challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_ChallengeManager *ChallengeManagerTransactor) ClearChallenge(opts *bind.TransactOpts, challengeIndex uint64) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "clearChallenge", challengeIndex)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_ChallengeManager *ChallengeManagerSession) ClearChallenge(challengeIndex uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ClearChallenge(&_ChallengeManager.TransactOpts, challengeIndex)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) ClearChallenge(challengeIndex uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ClearChallenge(&_ChallengeManager.TransactOpts, challengeIndex)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_ChallengeManager *ChallengeManagerTransactor) CreateChallenge(opts *bind.TransactOpts, wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "createChallenge", wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_ChallengeManager *ChallengeManagerSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _ChallengeManager.Contract.CreateChallenge(&_ChallengeManager.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_ChallengeManager *ChallengeManagerTransactorSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _ChallengeManager.Contract.CreateChallenge(&_ChallengeManager.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_ChallengeManager *ChallengeManagerTransactor) Initialize(opts *bind.TransactOpts, resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "initialize", resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_ChallengeManager *ChallengeManagerSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.Initialize(&_ChallengeManager.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.Initialize(&_ChallengeManager.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_ChallengeManager *ChallengeManagerTransactor) OneStepProveExecution(opts *bind.TransactOpts, challengeIndex uint64, selection ChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "oneStepProveExecution", challengeIndex, selection, proof)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_ChallengeManager *ChallengeManagerSession) OneStepProveExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.OneStepProveExecution(&_ChallengeManager.TransactOpts, challengeIndex, selection, proof)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) OneStepProveExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.OneStepProveExecution(&_ChallengeManager.TransactOpts, challengeIndex, selection, proof)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address osp_) returns()
func (_ChallengeManager *ChallengeManagerTransactor) PostUpgradeInit(opts *bind.TransactOpts, osp_ common.Address) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "postUpgradeInit", osp_)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address osp_) returns()
func (_ChallengeManager *ChallengeManagerSession) PostUpgradeInit(osp_ common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.PostUpgradeInit(&_ChallengeManager.TransactOpts, osp_)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address osp_) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) PostUpgradeInit(osp_ common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.PostUpgradeInit(&_ChallengeManager.TransactOpts, osp_)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_ChallengeManager *ChallengeManagerTransactor) Timeout(opts *bind.TransactOpts, challengeIndex uint64) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "timeout", challengeIndex)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_ChallengeManager *ChallengeManagerSession) Timeout(challengeIndex uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.Timeout(&_ChallengeManager.TransactOpts, challengeIndex)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) Timeout(challengeIndex uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.Timeout(&_ChallengeManager.TransactOpts, challengeIndex)
}

// ChallengeManagerBisectedIterator is returned from FilterBisected and is used to iterate over the raw logs and unpacked data for Bisected events raised by the ChallengeManager contract.
type ChallengeManagerBisectedIterator struct {
	Event *ChallengeManagerBisected // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerBisected)
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
		it.Event = new(ChallengeManagerBisected)
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
func (it *ChallengeManagerBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerBisected represents a Bisected event raised by the ChallengeManager contract.
type ChallengeManagerBisected struct {
	ChallengeIndex          uint64
	ChallengeRoot           [32]byte
	ChallengedSegmentStart  *big.Int
	ChallengedSegmentLength *big.Int
	ChainHashes             [][32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBisected is a free log retrieval operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_ChallengeManager *ChallengeManagerFilterer) FilterBisected(opts *bind.FilterOpts, challengeIndex []uint64, challengeRoot [][32]byte) (*ChallengeManagerBisectedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerBisectedIterator{contract: _ChallengeManager.contract, event: "Bisected", logs: logs, sub: sub}, nil
}

// WatchBisected is a free log subscription operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_ChallengeManager *ChallengeManagerFilterer) WatchBisected(opts *bind.WatchOpts, sink chan<- *ChallengeManagerBisected, challengeIndex []uint64, challengeRoot [][32]byte) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerBisected)
				if err := _ChallengeManager.contract.UnpackLog(event, "Bisected", log); err != nil {
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

// ParseBisected is a log parse operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_ChallengeManager *ChallengeManagerFilterer) ParseBisected(log types.Log) (*ChallengeManagerBisected, error) {
	event := new(ChallengeManagerBisected)
	if err := _ChallengeManager.contract.UnpackLog(event, "Bisected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeManagerChallengeEndedIterator is returned from FilterChallengeEnded and is used to iterate over the raw logs and unpacked data for ChallengeEnded events raised by the ChallengeManager contract.
type ChallengeManagerChallengeEndedIterator struct {
	Event *ChallengeManagerChallengeEnded // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerChallengeEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerChallengeEnded)
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
		it.Event = new(ChallengeManagerChallengeEnded)
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
func (it *ChallengeManagerChallengeEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerChallengeEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerChallengeEnded represents a ChallengeEnded event raised by the ChallengeManager contract.
type ChallengeManagerChallengeEnded struct {
	ChallengeIndex uint64
	Kind           uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChallengeEnded is a free log retrieval operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_ChallengeManager *ChallengeManagerFilterer) FilterChallengeEnded(opts *bind.FilterOpts, challengeIndex []uint64) (*ChallengeManagerChallengeEndedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerChallengeEndedIterator{contract: _ChallengeManager.contract, event: "ChallengeEnded", logs: logs, sub: sub}, nil
}

// WatchChallengeEnded is a free log subscription operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_ChallengeManager *ChallengeManagerFilterer) WatchChallengeEnded(opts *bind.WatchOpts, sink chan<- *ChallengeManagerChallengeEnded, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerChallengeEnded)
				if err := _ChallengeManager.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
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

// ParseChallengeEnded is a log parse operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_ChallengeManager *ChallengeManagerFilterer) ParseChallengeEnded(log types.Log) (*ChallengeManagerChallengeEnded, error) {
	event := new(ChallengeManagerChallengeEnded)
	if err := _ChallengeManager.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeManagerExecutionChallengeBegunIterator is returned from FilterExecutionChallengeBegun and is used to iterate over the raw logs and unpacked data for ExecutionChallengeBegun events raised by the ChallengeManager contract.
type ChallengeManagerExecutionChallengeBegunIterator struct {
	Event *ChallengeManagerExecutionChallengeBegun // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerExecutionChallengeBegunIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerExecutionChallengeBegun)
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
		it.Event = new(ChallengeManagerExecutionChallengeBegun)
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
func (it *ChallengeManagerExecutionChallengeBegunIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerExecutionChallengeBegunIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerExecutionChallengeBegun represents a ExecutionChallengeBegun event raised by the ChallengeManager contract.
type ChallengeManagerExecutionChallengeBegun struct {
	ChallengeIndex uint64
	BlockSteps     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExecutionChallengeBegun is a free log retrieval operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_ChallengeManager *ChallengeManagerFilterer) FilterExecutionChallengeBegun(opts *bind.FilterOpts, challengeIndex []uint64) (*ChallengeManagerExecutionChallengeBegunIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerExecutionChallengeBegunIterator{contract: _ChallengeManager.contract, event: "ExecutionChallengeBegun", logs: logs, sub: sub}, nil
}

// WatchExecutionChallengeBegun is a free log subscription operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_ChallengeManager *ChallengeManagerFilterer) WatchExecutionChallengeBegun(opts *bind.WatchOpts, sink chan<- *ChallengeManagerExecutionChallengeBegun, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerExecutionChallengeBegun)
				if err := _ChallengeManager.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
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

// ParseExecutionChallengeBegun is a log parse operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_ChallengeManager *ChallengeManagerFilterer) ParseExecutionChallengeBegun(log types.Log) (*ChallengeManagerExecutionChallengeBegun, error) {
	event := new(ChallengeManagerExecutionChallengeBegun)
	if err := _ChallengeManager.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeManagerInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the ChallengeManager contract.
type ChallengeManagerInitiatedChallengeIterator struct {
	Event *ChallengeManagerInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerInitiatedChallenge)
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
		it.Event = new(ChallengeManagerInitiatedChallenge)
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
func (it *ChallengeManagerInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerInitiatedChallenge represents a InitiatedChallenge event raised by the ChallengeManager contract.
type ChallengeManagerInitiatedChallenge struct {
	ChallengeIndex uint64
	StartState     GlobalState
	EndState       GlobalState
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_ChallengeManager *ChallengeManagerFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts, challengeIndex []uint64) (*ChallengeManagerInitiatedChallengeIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerInitiatedChallengeIterator{contract: _ChallengeManager.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_ChallengeManager *ChallengeManagerFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeManagerInitiatedChallenge, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerInitiatedChallenge)
				if err := _ChallengeManager.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_ChallengeManager *ChallengeManagerFilterer) ParseInitiatedChallenge(log types.Log) (*ChallengeManagerInitiatedChallenge, error) {
	event := new(ChallengeManagerInitiatedChallenge)
	if err := _ChallengeManager.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeManagerOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the ChallengeManager contract.
type ChallengeManagerOneStepProofCompletedIterator struct {
	Event *ChallengeManagerOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerOneStepProofCompleted)
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
		it.Event = new(ChallengeManagerOneStepProofCompleted)
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
func (it *ChallengeManagerOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerOneStepProofCompleted represents a OneStepProofCompleted event raised by the ChallengeManager contract.
type ChallengeManagerOneStepProofCompleted struct {
	ChallengeIndex uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_ChallengeManager *ChallengeManagerFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts, challengeIndex []uint64) (*ChallengeManagerOneStepProofCompletedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerOneStepProofCompletedIterator{contract: _ChallengeManager.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_ChallengeManager *ChallengeManagerFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *ChallengeManagerOneStepProofCompleted, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerOneStepProofCompleted)
				if err := _ChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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

// ParseOneStepProofCompleted is a log parse operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_ChallengeManager *ChallengeManagerFilterer) ParseOneStepProofCompleted(log types.Log) (*ChallengeManagerOneStepProofCompleted, error) {
	event := new(ChallengeManagerOneStepProofCompleted)
	if err := _ChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IChallengeManagerMetaData contains all meta data concerning the IChallengeManager contract.
var IChallengeManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"challengeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"chainHashes\",\"type\":\"bytes32[]\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIChallengeManager.ChallengeTerminationType\",\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"ChallengeEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockSteps\",\"type\":\"uint256\"}],\"name\":\"ExecutionChallengeBegun\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"startState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"endState\",\"type\":\"tuple\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex_\",\"type\":\"uint64\"}],\"name\":\"challengeInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"current\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastMoveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessages\",\"type\":\"uint64\"},{\"internalType\":\"enumChallengeLib.ChallengeMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"internalType\":\"structChallengeLib.Challenge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex_\",\"type\":\"uint64\"}],\"name\":\"clearChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"startAndEndMachineStatuses_\",\"type\":\"uint8[2]\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState[2]\",\"name\":\"startAndEndGlobalStates_\",\"type\":\"tuple[2]\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"asserter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"asserterTimeLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengerTimeLeft_\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"currentResponder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIChallengeResultReceiver\",\"name\":\"resultReceiver_\",\"type\":\"address\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"sequencerInbox_\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"osp_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"isTimedOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex_\",\"type\":\"uint64\"}],\"name\":\"timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IChallengeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IChallengeManagerMetaData.ABI instead.
var IChallengeManagerABI = IChallengeManagerMetaData.ABI

// IChallengeManager is an auto generated Go binding around an Ethereum contract.
type IChallengeManager struct {
	IChallengeManagerCaller     // Read-only binding to the contract
	IChallengeManagerTransactor // Write-only binding to the contract
	IChallengeManagerFilterer   // Log filterer for contract events
}

// IChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeManagerSession struct {
	Contract     *IChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeManagerCallerSession struct {
	Contract *IChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeManagerTransactorSession struct {
	Contract     *IChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeManagerRaw struct {
	Contract *IChallengeManager // Generic contract binding to access the raw methods on
}

// IChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeManagerCallerRaw struct {
	Contract *IChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactorRaw struct {
	Contract *IChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeManager creates a new instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManager(address common.Address, backend bind.ContractBackend) (*IChallengeManager, error) {
	contract, err := bindIChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeManager{IChallengeManagerCaller: IChallengeManagerCaller{contract: contract}, IChallengeManagerTransactor: IChallengeManagerTransactor{contract: contract}, IChallengeManagerFilterer: IChallengeManagerFilterer{contract: contract}}, nil
}

// NewIChallengeManagerCaller creates a new read-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*IChallengeManagerCaller, error) {
	contract, err := bindIChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerCaller{contract: contract}, nil
}

// NewIChallengeManagerTransactor creates a new write-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeManagerTransactor, error) {
	contract, err := bindIChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerTransactor{contract: contract}, nil
}

// NewIChallengeManagerFilterer creates a new log filterer instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeManagerFilterer, error) {
	contract, err := bindIChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerFilterer{contract: contract}, nil
}

// bindIChallengeManager binds a generic wrapper to an already deployed contract.
func bindIChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IChallengeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.IChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex_) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_IChallengeManager *IChallengeManagerCaller) ChallengeInfo(opts *bind.CallOpts, challengeIndex_ uint64) (ChallengeLibChallenge, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "challengeInfo", challengeIndex_)

	if err != nil {
		return *new(ChallengeLibChallenge), err
	}

	out0 := *abi.ConvertType(out[0], new(ChallengeLibChallenge)).(*ChallengeLibChallenge)

	return out0, err

}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex_) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_IChallengeManager *IChallengeManagerSession) ChallengeInfo(challengeIndex_ uint64) (ChallengeLibChallenge, error) {
	return _IChallengeManager.Contract.ChallengeInfo(&_IChallengeManager.CallOpts, challengeIndex_)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex_) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_IChallengeManager *IChallengeManagerCallerSession) ChallengeInfo(challengeIndex_ uint64) (ChallengeLibChallenge, error) {
	return _IChallengeManager.Contract.ChallengeInfo(&_IChallengeManager.CallOpts, challengeIndex_)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_IChallengeManager *IChallengeManagerCaller) CurrentResponder(opts *bind.CallOpts, challengeIndex uint64) (common.Address, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "currentResponder", challengeIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_IChallengeManager *IChallengeManagerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _IChallengeManager.Contract.CurrentResponder(&_IChallengeManager.CallOpts, challengeIndex)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_IChallengeManager *IChallengeManagerCallerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _IChallengeManager.Contract.CurrentResponder(&_IChallengeManager.CallOpts, challengeIndex)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_IChallengeManager *IChallengeManagerCaller) IsTimedOut(opts *bind.CallOpts, challengeIndex uint64) (bool, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "isTimedOut", challengeIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_IChallengeManager *IChallengeManagerSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _IChallengeManager.Contract.IsTimedOut(&_IChallengeManager.CallOpts, challengeIndex)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_IChallengeManager *IChallengeManagerCallerSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _IChallengeManager.Contract.IsTimedOut(&_IChallengeManager.CallOpts, challengeIndex)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex_) returns()
func (_IChallengeManager *IChallengeManagerTransactor) ClearChallenge(opts *bind.TransactOpts, challengeIndex_ uint64) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "clearChallenge", challengeIndex_)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex_) returns()
func (_IChallengeManager *IChallengeManagerSession) ClearChallenge(challengeIndex_ uint64) (*types.Transaction, error) {
	return _IChallengeManager.Contract.ClearChallenge(&_IChallengeManager.TransactOpts, challengeIndex_)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex_) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) ClearChallenge(challengeIndex_ uint64) (*types.Transaction, error) {
	return _IChallengeManager.Contract.ClearChallenge(&_IChallengeManager.TransactOpts, challengeIndex_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_IChallengeManager *IChallengeManagerTransactor) CreateChallenge(opts *bind.TransactOpts, wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "createChallenge", wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_IChallengeManager *IChallengeManagerSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _IChallengeManager.Contract.CreateChallenge(&_IChallengeManager.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_IChallengeManager *IChallengeManagerTransactorSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _IChallengeManager.Contract.CreateChallenge(&_IChallengeManager.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_IChallengeManager *IChallengeManagerTransactor) Initialize(opts *bind.TransactOpts, resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "initialize", resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_IChallengeManager *IChallengeManagerSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Initialize(&_IChallengeManager.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Initialize(&_IChallengeManager.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex_) returns()
func (_IChallengeManager *IChallengeManagerTransactor) Timeout(opts *bind.TransactOpts, challengeIndex_ uint64) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "timeout", challengeIndex_)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex_) returns()
func (_IChallengeManager *IChallengeManagerSession) Timeout(challengeIndex_ uint64) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Timeout(&_IChallengeManager.TransactOpts, challengeIndex_)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex_) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) Timeout(challengeIndex_ uint64) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Timeout(&_IChallengeManager.TransactOpts, challengeIndex_)
}

// IChallengeManagerBisectedIterator is returned from FilterBisected and is used to iterate over the raw logs and unpacked data for Bisected events raised by the IChallengeManager contract.
type IChallengeManagerBisectedIterator struct {
	Event *IChallengeManagerBisected // Event containing the contract specifics and raw log

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
func (it *IChallengeManagerBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IChallengeManagerBisected)
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
		it.Event = new(IChallengeManagerBisected)
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
func (it *IChallengeManagerBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IChallengeManagerBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IChallengeManagerBisected represents a Bisected event raised by the IChallengeManager contract.
type IChallengeManagerBisected struct {
	ChallengeIndex          uint64
	ChallengeRoot           [32]byte
	ChallengedSegmentStart  *big.Int
	ChallengedSegmentLength *big.Int
	ChainHashes             [][32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBisected is a free log retrieval operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_IChallengeManager *IChallengeManagerFilterer) FilterBisected(opts *bind.FilterOpts, challengeIndex []uint64, challengeRoot [][32]byte) (*IChallengeManagerBisectedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _IChallengeManager.contract.FilterLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerBisectedIterator{contract: _IChallengeManager.contract, event: "Bisected", logs: logs, sub: sub}, nil
}

// WatchBisected is a free log subscription operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_IChallengeManager *IChallengeManagerFilterer) WatchBisected(opts *bind.WatchOpts, sink chan<- *IChallengeManagerBisected, challengeIndex []uint64, challengeRoot [][32]byte) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _IChallengeManager.contract.WatchLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IChallengeManagerBisected)
				if err := _IChallengeManager.contract.UnpackLog(event, "Bisected", log); err != nil {
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

// ParseBisected is a log parse operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_IChallengeManager *IChallengeManagerFilterer) ParseBisected(log types.Log) (*IChallengeManagerBisected, error) {
	event := new(IChallengeManagerBisected)
	if err := _IChallengeManager.contract.UnpackLog(event, "Bisected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IChallengeManagerChallengeEndedIterator is returned from FilterChallengeEnded and is used to iterate over the raw logs and unpacked data for ChallengeEnded events raised by the IChallengeManager contract.
type IChallengeManagerChallengeEndedIterator struct {
	Event *IChallengeManagerChallengeEnded // Event containing the contract specifics and raw log

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
func (it *IChallengeManagerChallengeEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IChallengeManagerChallengeEnded)
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
		it.Event = new(IChallengeManagerChallengeEnded)
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
func (it *IChallengeManagerChallengeEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IChallengeManagerChallengeEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IChallengeManagerChallengeEnded represents a ChallengeEnded event raised by the IChallengeManager contract.
type IChallengeManagerChallengeEnded struct {
	ChallengeIndex uint64
	Kind           uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChallengeEnded is a free log retrieval operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_IChallengeManager *IChallengeManagerFilterer) FilterChallengeEnded(opts *bind.FilterOpts, challengeIndex []uint64) (*IChallengeManagerChallengeEndedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.FilterLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerChallengeEndedIterator{contract: _IChallengeManager.contract, event: "ChallengeEnded", logs: logs, sub: sub}, nil
}

// WatchChallengeEnded is a free log subscription operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_IChallengeManager *IChallengeManagerFilterer) WatchChallengeEnded(opts *bind.WatchOpts, sink chan<- *IChallengeManagerChallengeEnded, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.WatchLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IChallengeManagerChallengeEnded)
				if err := _IChallengeManager.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
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

// ParseChallengeEnded is a log parse operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_IChallengeManager *IChallengeManagerFilterer) ParseChallengeEnded(log types.Log) (*IChallengeManagerChallengeEnded, error) {
	event := new(IChallengeManagerChallengeEnded)
	if err := _IChallengeManager.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IChallengeManagerExecutionChallengeBegunIterator is returned from FilterExecutionChallengeBegun and is used to iterate over the raw logs and unpacked data for ExecutionChallengeBegun events raised by the IChallengeManager contract.
type IChallengeManagerExecutionChallengeBegunIterator struct {
	Event *IChallengeManagerExecutionChallengeBegun // Event containing the contract specifics and raw log

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
func (it *IChallengeManagerExecutionChallengeBegunIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IChallengeManagerExecutionChallengeBegun)
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
		it.Event = new(IChallengeManagerExecutionChallengeBegun)
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
func (it *IChallengeManagerExecutionChallengeBegunIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IChallengeManagerExecutionChallengeBegunIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IChallengeManagerExecutionChallengeBegun represents a ExecutionChallengeBegun event raised by the IChallengeManager contract.
type IChallengeManagerExecutionChallengeBegun struct {
	ChallengeIndex uint64
	BlockSteps     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExecutionChallengeBegun is a free log retrieval operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_IChallengeManager *IChallengeManagerFilterer) FilterExecutionChallengeBegun(opts *bind.FilterOpts, challengeIndex []uint64) (*IChallengeManagerExecutionChallengeBegunIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.FilterLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerExecutionChallengeBegunIterator{contract: _IChallengeManager.contract, event: "ExecutionChallengeBegun", logs: logs, sub: sub}, nil
}

// WatchExecutionChallengeBegun is a free log subscription operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_IChallengeManager *IChallengeManagerFilterer) WatchExecutionChallengeBegun(opts *bind.WatchOpts, sink chan<- *IChallengeManagerExecutionChallengeBegun, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.WatchLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IChallengeManagerExecutionChallengeBegun)
				if err := _IChallengeManager.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
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

// ParseExecutionChallengeBegun is a log parse operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_IChallengeManager *IChallengeManagerFilterer) ParseExecutionChallengeBegun(log types.Log) (*IChallengeManagerExecutionChallengeBegun, error) {
	event := new(IChallengeManagerExecutionChallengeBegun)
	if err := _IChallengeManager.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IChallengeManagerInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the IChallengeManager contract.
type IChallengeManagerInitiatedChallengeIterator struct {
	Event *IChallengeManagerInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *IChallengeManagerInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IChallengeManagerInitiatedChallenge)
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
		it.Event = new(IChallengeManagerInitiatedChallenge)
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
func (it *IChallengeManagerInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IChallengeManagerInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IChallengeManagerInitiatedChallenge represents a InitiatedChallenge event raised by the IChallengeManager contract.
type IChallengeManagerInitiatedChallenge struct {
	ChallengeIndex uint64
	StartState     GlobalState
	EndState       GlobalState
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_IChallengeManager *IChallengeManagerFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts, challengeIndex []uint64) (*IChallengeManagerInitiatedChallengeIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.FilterLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerInitiatedChallengeIterator{contract: _IChallengeManager.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_IChallengeManager *IChallengeManagerFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *IChallengeManagerInitiatedChallenge, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.WatchLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IChallengeManagerInitiatedChallenge)
				if err := _IChallengeManager.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_IChallengeManager *IChallengeManagerFilterer) ParseInitiatedChallenge(log types.Log) (*IChallengeManagerInitiatedChallenge, error) {
	event := new(IChallengeManagerInitiatedChallenge)
	if err := _IChallengeManager.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IChallengeManagerOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the IChallengeManager contract.
type IChallengeManagerOneStepProofCompletedIterator struct {
	Event *IChallengeManagerOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *IChallengeManagerOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IChallengeManagerOneStepProofCompleted)
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
		it.Event = new(IChallengeManagerOneStepProofCompleted)
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
func (it *IChallengeManagerOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IChallengeManagerOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IChallengeManagerOneStepProofCompleted represents a OneStepProofCompleted event raised by the IChallengeManager contract.
type IChallengeManagerOneStepProofCompleted struct {
	ChallengeIndex uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_IChallengeManager *IChallengeManagerFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts, challengeIndex []uint64) (*IChallengeManagerOneStepProofCompletedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.FilterLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerOneStepProofCompletedIterator{contract: _IChallengeManager.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_IChallengeManager *IChallengeManagerFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *IChallengeManagerOneStepProofCompleted, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.WatchLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IChallengeManagerOneStepProofCompleted)
				if err := _IChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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

// ParseOneStepProofCompleted is a log parse operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_IChallengeManager *IChallengeManagerFilterer) ParseOneStepProofCompleted(log types.Log) (*IChallengeManagerOneStepProofCompleted, error) {
	event := new(IChallengeManagerOneStepProofCompleted)
	if err := _IChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IChallengeResultReceiverMetaData contains all meta data concerning the IChallengeResultReceiver contract.
var IChallengeResultReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengeIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IChallengeResultReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IChallengeResultReceiverMetaData.ABI instead.
var IChallengeResultReceiverABI = IChallengeResultReceiverMetaData.ABI

// IChallengeResultReceiver is an auto generated Go binding around an Ethereum contract.
type IChallengeResultReceiver struct {
	IChallengeResultReceiverCaller     // Read-only binding to the contract
	IChallengeResultReceiverTransactor // Write-only binding to the contract
	IChallengeResultReceiverFilterer   // Log filterer for contract events
}

// IChallengeResultReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeResultReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeResultReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeResultReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeResultReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeResultReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeResultReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeResultReceiverSession struct {
	Contract     *IChallengeResultReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IChallengeResultReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeResultReceiverCallerSession struct {
	Contract *IChallengeResultReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IChallengeResultReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeResultReceiverTransactorSession struct {
	Contract     *IChallengeResultReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IChallengeResultReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeResultReceiverRaw struct {
	Contract *IChallengeResultReceiver // Generic contract binding to access the raw methods on
}

// IChallengeResultReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeResultReceiverCallerRaw struct {
	Contract *IChallengeResultReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeResultReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeResultReceiverTransactorRaw struct {
	Contract *IChallengeResultReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeResultReceiver creates a new instance of IChallengeResultReceiver, bound to a specific deployed contract.
func NewIChallengeResultReceiver(address common.Address, backend bind.ContractBackend) (*IChallengeResultReceiver, error) {
	contract, err := bindIChallengeResultReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeResultReceiver{IChallengeResultReceiverCaller: IChallengeResultReceiverCaller{contract: contract}, IChallengeResultReceiverTransactor: IChallengeResultReceiverTransactor{contract: contract}, IChallengeResultReceiverFilterer: IChallengeResultReceiverFilterer{contract: contract}}, nil
}

// NewIChallengeResultReceiverCaller creates a new read-only instance of IChallengeResultReceiver, bound to a specific deployed contract.
func NewIChallengeResultReceiverCaller(address common.Address, caller bind.ContractCaller) (*IChallengeResultReceiverCaller, error) {
	contract, err := bindIChallengeResultReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeResultReceiverCaller{contract: contract}, nil
}

// NewIChallengeResultReceiverTransactor creates a new write-only instance of IChallengeResultReceiver, bound to a specific deployed contract.
func NewIChallengeResultReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeResultReceiverTransactor, error) {
	contract, err := bindIChallengeResultReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeResultReceiverTransactor{contract: contract}, nil
}

// NewIChallengeResultReceiverFilterer creates a new log filterer instance of IChallengeResultReceiver, bound to a specific deployed contract.
func NewIChallengeResultReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeResultReceiverFilterer, error) {
	contract, err := bindIChallengeResultReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeResultReceiverFilterer{contract: contract}, nil
}

// bindIChallengeResultReceiver binds a generic wrapper to an already deployed contract.
func bindIChallengeResultReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IChallengeResultReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeResultReceiver *IChallengeResultReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeResultReceiver.Contract.IChallengeResultReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeResultReceiver *IChallengeResultReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeResultReceiver.Contract.IChallengeResultReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeResultReceiver *IChallengeResultReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeResultReceiver.Contract.IChallengeResultReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeResultReceiver *IChallengeResultReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeResultReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeResultReceiver *IChallengeResultReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeResultReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeResultReceiver *IChallengeResultReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeResultReceiver.Contract.contract.Transact(opts, method, params...)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x0357aa49.
//
// Solidity: function completeChallenge(uint256 challengeIndex, address winner, address loser) returns()
func (_IChallengeResultReceiver *IChallengeResultReceiverTransactor) CompleteChallenge(opts *bind.TransactOpts, challengeIndex *big.Int, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _IChallengeResultReceiver.contract.Transact(opts, "completeChallenge", challengeIndex, winner, loser)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x0357aa49.
//
// Solidity: function completeChallenge(uint256 challengeIndex, address winner, address loser) returns()
func (_IChallengeResultReceiver *IChallengeResultReceiverSession) CompleteChallenge(challengeIndex *big.Int, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _IChallengeResultReceiver.Contract.CompleteChallenge(&_IChallengeResultReceiver.TransactOpts, challengeIndex, winner, loser)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x0357aa49.
//
// Solidity: function completeChallenge(uint256 challengeIndex, address winner, address loser) returns()
func (_IChallengeResultReceiver *IChallengeResultReceiverTransactorSession) CompleteChallenge(challengeIndex *big.Int, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _IChallengeResultReceiver.Contract.CompleteChallenge(&_IChallengeResultReceiver.TransactOpts, challengeIndex, winner, loser)
}
