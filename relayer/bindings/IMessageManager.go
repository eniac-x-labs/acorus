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

// IMessageManagerMetaData contains all meta data concerning the IMessageManager contract.
var IMessageManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageAlreadySent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IMessageManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageManagerMetaData.ABI instead.
var IMessageManagerABI = IMessageManagerMetaData.ABI

// IMessageManager is an auto generated Go binding around an Ethereum contract.
type IMessageManager struct {
	IMessageManagerCaller     // Read-only binding to the contract
	IMessageManagerTransactor // Write-only binding to the contract
	IMessageManagerFilterer   // Log filterer for contract events
}

// IMessageManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageManagerSession struct {
	Contract     *IMessageManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMessageManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageManagerCallerSession struct {
	Contract *IMessageManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IMessageManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageManagerTransactorSession struct {
	Contract     *IMessageManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IMessageManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageManagerRaw struct {
	Contract *IMessageManager // Generic contract binding to access the raw methods on
}

// IMessageManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageManagerCallerRaw struct {
	Contract *IMessageManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageManagerTransactorRaw struct {
	Contract *IMessageManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageManager creates a new instance of IMessageManager, bound to a specific deployed contract.
func NewIMessageManager(address common.Address, backend bind.ContractBackend) (*IMessageManager, error) {
	contract, err := bindIMessageManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageManager{IMessageManagerCaller: IMessageManagerCaller{contract: contract}, IMessageManagerTransactor: IMessageManagerTransactor{contract: contract}, IMessageManagerFilterer: IMessageManagerFilterer{contract: contract}}, nil
}

// NewIMessageManagerCaller creates a new read-only instance of IMessageManager, bound to a specific deployed contract.
func NewIMessageManagerCaller(address common.Address, caller bind.ContractCaller) (*IMessageManagerCaller, error) {
	contract, err := bindIMessageManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageManagerCaller{contract: contract}, nil
}

// NewIMessageManagerTransactor creates a new write-only instance of IMessageManager, bound to a specific deployed contract.
func NewIMessageManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageManagerTransactor, error) {
	contract, err := bindIMessageManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageManagerTransactor{contract: contract}, nil
}

// NewIMessageManagerFilterer creates a new log filterer instance of IMessageManager, bound to a specific deployed contract.
func NewIMessageManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageManagerFilterer, error) {
	contract, err := bindIMessageManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageManagerFilterer{contract: contract}, nil
}

// bindIMessageManager binds a generic wrapper to an already deployed contract.
func bindIMessageManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMessageManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageManager *IMessageManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageManager.Contract.IMessageManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageManager *IMessageManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageManager.Contract.IMessageManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageManager *IMessageManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageManager.Contract.IMessageManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageManager *IMessageManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageManager *IMessageManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageManager *IMessageManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageManager.Contract.contract.Transact(opts, method, params...)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x1eb65ffb.
//
// Solidity: function claimMessage(uint256 sourceChainId, uint256 destChainId, address _to, uint256 _fee, uint256 _value, uint256 _nonce) returns()
func (_IMessageManager *IMessageManagerTransactor) ClaimMessage(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, _to common.Address, _fee *big.Int, _value *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _IMessageManager.contract.Transact(opts, "claimMessage", sourceChainId, destChainId, _to, _fee, _value, _nonce)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x1eb65ffb.
//
// Solidity: function claimMessage(uint256 sourceChainId, uint256 destChainId, address _to, uint256 _fee, uint256 _value, uint256 _nonce) returns()
func (_IMessageManager *IMessageManagerSession) ClaimMessage(sourceChainId *big.Int, destChainId *big.Int, _to common.Address, _fee *big.Int, _value *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _IMessageManager.Contract.ClaimMessage(&_IMessageManager.TransactOpts, sourceChainId, destChainId, _to, _fee, _value, _nonce)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x1eb65ffb.
//
// Solidity: function claimMessage(uint256 sourceChainId, uint256 destChainId, address _to, uint256 _fee, uint256 _value, uint256 _nonce) returns()
func (_IMessageManager *IMessageManagerTransactorSession) ClaimMessage(sourceChainId *big.Int, destChainId *big.Int, _to common.Address, _fee *big.Int, _value *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _IMessageManager.Contract.ClaimMessage(&_IMessageManager.TransactOpts, sourceChainId, destChainId, _to, _fee, _value, _nonce)
}

// SendMessage is a paid mutator transaction binding the contract method 0xc17d1e14.
//
// Solidity: function sendMessage(uint256 sourceChainId, uint256 destChainId, address _to, uint256 _value, uint256 _fee) returns()
func (_IMessageManager *IMessageManagerTransactor) SendMessage(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, _to common.Address, _value *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _IMessageManager.contract.Transact(opts, "sendMessage", sourceChainId, destChainId, _to, _value, _fee)
}

// SendMessage is a paid mutator transaction binding the contract method 0xc17d1e14.
//
// Solidity: function sendMessage(uint256 sourceChainId, uint256 destChainId, address _to, uint256 _value, uint256 _fee) returns()
func (_IMessageManager *IMessageManagerSession) SendMessage(sourceChainId *big.Int, destChainId *big.Int, _to common.Address, _value *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _IMessageManager.Contract.SendMessage(&_IMessageManager.TransactOpts, sourceChainId, destChainId, _to, _value, _fee)
}

// SendMessage is a paid mutator transaction binding the contract method 0xc17d1e14.
//
// Solidity: function sendMessage(uint256 sourceChainId, uint256 destChainId, address _to, uint256 _value, uint256 _fee) returns()
func (_IMessageManager *IMessageManagerTransactorSession) SendMessage(sourceChainId *big.Int, destChainId *big.Int, _to common.Address, _value *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _IMessageManager.Contract.SendMessage(&_IMessageManager.TransactOpts, sourceChainId, destChainId, _to, _value, _fee)
}

// IMessageManagerMessageClaimedIterator is returned from FilterMessageClaimed and is used to iterate over the raw logs and unpacked data for MessageClaimed events raised by the IMessageManager contract.
type IMessageManagerMessageClaimedIterator struct {
	Event *IMessageManagerMessageClaimed // Event containing the contract specifics and raw log

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
func (it *IMessageManagerMessageClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMessageManagerMessageClaimed)
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
		it.Event = new(IMessageManagerMessageClaimed)
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
func (it *IMessageManagerMessageClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMessageManagerMessageClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMessageManagerMessageClaimed represents a MessageClaimed event raised by the IMessageManager contract.
type IMessageManagerMessageClaimed struct {
	SourceChainId *big.Int
	DestChainId   *big.Int
	MessageHash   [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageClaimed is a free log retrieval operation binding the contract event 0x9ea158d96246b092af2f857f58de20289ede65682a6d7e52ae067f0e244242a0.
//
// Solidity: event MessageClaimed(uint256 sourceChainId, uint256 destChainId, bytes32 indexed _messageHash)
func (_IMessageManager *IMessageManagerFilterer) FilterMessageClaimed(opts *bind.FilterOpts, _messageHash [][32]byte) (*IMessageManagerMessageClaimedIterator, error) {

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _IMessageManager.contract.FilterLogs(opts, "MessageClaimed", _messageHashRule)
	if err != nil {
		return nil, err
	}
	return &IMessageManagerMessageClaimedIterator{contract: _IMessageManager.contract, event: "MessageClaimed", logs: logs, sub: sub}, nil
}

// WatchMessageClaimed is a free log subscription operation binding the contract event 0x9ea158d96246b092af2f857f58de20289ede65682a6d7e52ae067f0e244242a0.
//
// Solidity: event MessageClaimed(uint256 sourceChainId, uint256 destChainId, bytes32 indexed _messageHash)
func (_IMessageManager *IMessageManagerFilterer) WatchMessageClaimed(opts *bind.WatchOpts, sink chan<- *IMessageManagerMessageClaimed, _messageHash [][32]byte) (event.Subscription, error) {

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _IMessageManager.contract.WatchLogs(opts, "MessageClaimed", _messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMessageManagerMessageClaimed)
				if err := _IMessageManager.contract.UnpackLog(event, "MessageClaimed", log); err != nil {
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

// ParseMessageClaimed is a log parse operation binding the contract event 0x9ea158d96246b092af2f857f58de20289ede65682a6d7e52ae067f0e244242a0.
//
// Solidity: event MessageClaimed(uint256 sourceChainId, uint256 destChainId, bytes32 indexed _messageHash)
func (_IMessageManager *IMessageManagerFilterer) ParseMessageClaimed(log types.Log) (*IMessageManagerMessageClaimed, error) {
	event := new(IMessageManagerMessageClaimed)
	if err := _IMessageManager.contract.UnpackLog(event, "MessageClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMessageManagerMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the IMessageManager contract.
type IMessageManagerMessageSentIterator struct {
	Event *IMessageManagerMessageSent // Event containing the contract specifics and raw log

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
func (it *IMessageManagerMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMessageManagerMessageSent)
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
		it.Event = new(IMessageManagerMessageSent)
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
func (it *IMessageManagerMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMessageManagerMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMessageManagerMessageSent represents a MessageSent event raised by the IMessageManager contract.
type IMessageManagerMessageSent struct {
	SourceChainId *big.Int
	DestChainId   *big.Int
	From          common.Address
	To            common.Address
	Fee           *big.Int
	Value         *big.Int
	Nonce         *big.Int
	MessageHash   [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x1e88b107f16455e0599576d3f09ea10212fea6adff673dd5c13bdfb470c843b6.
//
// Solidity: event MessageSent(uint256 sourceChainId, uint256 destChainId, address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes32 indexed _messageHash)
func (_IMessageManager *IMessageManagerFilterer) FilterMessageSent(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _messageHash [][32]byte) (*IMessageManagerMessageSentIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _IMessageManager.contract.FilterLogs(opts, "MessageSent", _fromRule, _toRule, _messageHashRule)
	if err != nil {
		return nil, err
	}
	return &IMessageManagerMessageSentIterator{contract: _IMessageManager.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x1e88b107f16455e0599576d3f09ea10212fea6adff673dd5c13bdfb470c843b6.
//
// Solidity: event MessageSent(uint256 sourceChainId, uint256 destChainId, address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes32 indexed _messageHash)
func (_IMessageManager *IMessageManagerFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *IMessageManagerMessageSent, _from []common.Address, _to []common.Address, _messageHash [][32]byte) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _IMessageManager.contract.WatchLogs(opts, "MessageSent", _fromRule, _toRule, _messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMessageManagerMessageSent)
				if err := _IMessageManager.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0x1e88b107f16455e0599576d3f09ea10212fea6adff673dd5c13bdfb470c843b6.
//
// Solidity: event MessageSent(uint256 sourceChainId, uint256 destChainId, address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes32 indexed _messageHash)
func (_IMessageManager *IMessageManagerFilterer) ParseMessageSent(log types.Log) (*IMessageManagerMessageSent, error) {
	event := new(IMessageManagerMessageSent)
	if err := _IMessageManager.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
