// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbase

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// EthbaseABI is the input ABI used to generate the binding from.
const EthbaseABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"eventId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"emitter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"eventName_\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"method\",\"type\":\"bytes4\"}],\"name\":\"Subscribed\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"emitter_\",\"type\":\"address\"},{\"name\":\"eventName_\",\"type\":\"bytes32\"},{\"name\":\"account_\",\"type\":\"address\"},{\"name\":\"method_\",\"type\":\"bytes4\"}],\"name\":\"subscribe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"eventId_\",\"type\":\"bytes32\"},{\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"unsubscribe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"eventId_\",\"type\":\"bytes32\"},{\"name\":\"account_\",\"type\":\"address\"},{\"name\":\"args_\",\"type\":\"bytes\"}],\"name\":\"invoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Ethbase is an auto generated Go binding around an Ethereum contract.
type Ethbase struct {
	EthbaseCaller     // Read-only binding to the contract
	EthbaseTransactor // Write-only binding to the contract
	EthbaseFilterer   // Log filterer for contract events
}

// EthbaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthbaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthbaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthbaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthbaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthbaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthbaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthbaseSession struct {
	Contract     *Ethbase          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthbaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthbaseCallerSession struct {
	Contract *EthbaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EthbaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthbaseTransactorSession struct {
	Contract     *EthbaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EthbaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthbaseRaw struct {
	Contract *Ethbase // Generic contract binding to access the raw methods on
}

// EthbaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthbaseCallerRaw struct {
	Contract *EthbaseCaller // Generic read-only contract binding to access the raw methods on
}

// EthbaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthbaseTransactorRaw struct {
	Contract *EthbaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthbase creates a new instance of Ethbase, bound to a specific deployed contract.
func NewEthbase(address common.Address, backend bind.ContractBackend) (*Ethbase, error) {
	contract, err := bindEthbase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethbase{EthbaseCaller: EthbaseCaller{contract: contract}, EthbaseTransactor: EthbaseTransactor{contract: contract}, EthbaseFilterer: EthbaseFilterer{contract: contract}}, nil
}

// NewEthbaseCaller creates a new read-only instance of Ethbase, bound to a specific deployed contract.
func NewEthbaseCaller(address common.Address, caller bind.ContractCaller) (*EthbaseCaller, error) {
	contract, err := bindEthbase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthbaseCaller{contract: contract}, nil
}

// NewEthbaseTransactor creates a new write-only instance of Ethbase, bound to a specific deployed contract.
func NewEthbaseTransactor(address common.Address, transactor bind.ContractTransactor) (*EthbaseTransactor, error) {
	contract, err := bindEthbase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthbaseTransactor{contract: contract}, nil
}

// NewEthbaseFilterer creates a new log filterer instance of Ethbase, bound to a specific deployed contract.
func NewEthbaseFilterer(address common.Address, filterer bind.ContractFilterer) (*EthbaseFilterer, error) {
	contract, err := bindEthbase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthbaseFilterer{contract: contract}, nil
}

// bindEthbase binds a generic wrapper to an already deployed contract.
func bindEthbase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthbaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethbase *EthbaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethbase.Contract.EthbaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethbase *EthbaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethbase.Contract.EthbaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethbase *EthbaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethbase.Contract.EthbaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethbase *EthbaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethbase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethbase *EthbaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethbase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethbase *EthbaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethbase.Contract.contract.Transact(opts, method, params...)
}

// Invoke is a paid mutator transaction binding the contract method 0x2e17e4c8.
//
// Solidity: function invoke(eventId_ bytes32, account_ address, args_ bytes) returns()
func (_Ethbase *EthbaseTransactor) Invoke(opts *bind.TransactOpts, eventId_ [32]byte, account_ common.Address, args_ []byte) (*types.Transaction, error) {
	return _Ethbase.contract.Transact(opts, "invoke", eventId_, account_, args_)
}

// Invoke is a paid mutator transaction binding the contract method 0x2e17e4c8.
//
// Solidity: function invoke(eventId_ bytes32, account_ address, args_ bytes) returns()
func (_Ethbase *EthbaseSession) Invoke(eventId_ [32]byte, account_ common.Address, args_ []byte) (*types.Transaction, error) {
	return _Ethbase.Contract.Invoke(&_Ethbase.TransactOpts, eventId_, account_, args_)
}

// Invoke is a paid mutator transaction binding the contract method 0x2e17e4c8.
//
// Solidity: function invoke(eventId_ bytes32, account_ address, args_ bytes) returns()
func (_Ethbase *EthbaseTransactorSession) Invoke(eventId_ [32]byte, account_ common.Address, args_ []byte) (*types.Transaction, error) {
	return _Ethbase.Contract.Invoke(&_Ethbase.TransactOpts, eventId_, account_, args_)
}

// Subscribe is a paid mutator transaction binding the contract method 0x4ff57300.
//
// Solidity: function subscribe(emitter_ address, eventName_ bytes32, account_ address, method_ bytes4) returns()
func (_Ethbase *EthbaseTransactor) Subscribe(opts *bind.TransactOpts, emitter_ common.Address, eventName_ [32]byte, account_ common.Address, method_ [4]byte) (*types.Transaction, error) {
	return _Ethbase.contract.Transact(opts, "subscribe", emitter_, eventName_, account_, method_)
}

// Subscribe is a paid mutator transaction binding the contract method 0x4ff57300.
//
// Solidity: function subscribe(emitter_ address, eventName_ bytes32, account_ address, method_ bytes4) returns()
func (_Ethbase *EthbaseSession) Subscribe(emitter_ common.Address, eventName_ [32]byte, account_ common.Address, method_ [4]byte) (*types.Transaction, error) {
	return _Ethbase.Contract.Subscribe(&_Ethbase.TransactOpts, emitter_, eventName_, account_, method_)
}

// Subscribe is a paid mutator transaction binding the contract method 0x4ff57300.
//
// Solidity: function subscribe(emitter_ address, eventName_ bytes32, account_ address, method_ bytes4) returns()
func (_Ethbase *EthbaseTransactorSession) Subscribe(emitter_ common.Address, eventName_ [32]byte, account_ common.Address, method_ [4]byte) (*types.Transaction, error) {
	return _Ethbase.Contract.Subscribe(&_Ethbase.TransactOpts, emitter_, eventName_, account_, method_)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xd59fd22f.
//
// Solidity: function unsubscribe(eventId_ bytes32, account_ address) returns()
func (_Ethbase *EthbaseTransactor) Unsubscribe(opts *bind.TransactOpts, eventId_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Ethbase.contract.Transact(opts, "unsubscribe", eventId_, account_)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xd59fd22f.
//
// Solidity: function unsubscribe(eventId_ bytes32, account_ address) returns()
func (_Ethbase *EthbaseSession) Unsubscribe(eventId_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Ethbase.Contract.Unsubscribe(&_Ethbase.TransactOpts, eventId_, account_)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xd59fd22f.
//
// Solidity: function unsubscribe(eventId_ bytes32, account_ address) returns()
func (_Ethbase *EthbaseTransactorSession) Unsubscribe(eventId_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Ethbase.Contract.Unsubscribe(&_Ethbase.TransactOpts, eventId_, account_)
}

// EthbaseSubscribedIterator is returned from FilterSubscribed and is used to iterate over the raw logs and unpacked data for Subscribed events raised by the Ethbase contract.
type EthbaseSubscribedIterator struct {
	Event *EthbaseSubscribed // Event containing the contract specifics and raw log

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
func (it *EthbaseSubscribedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthbaseSubscribed)
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
		it.Event = new(EthbaseSubscribed)
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
func (it *EthbaseSubscribedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthbaseSubscribedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthbaseSubscribed represents a Subscribed event raised by the Ethbase contract.
type EthbaseSubscribed struct {
	EventId   [32]byte
	Emitter   common.Address
	EventName [32]byte
	Account   common.Address
	Method    [4]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubscribed is a free log retrieval operation binding the contract event 0x04c9972fe596c2afdd8205c5d5cff160fdbbfc5254280890ce839b51fcdcad2f.
//
// Solidity: e Subscribed(eventId bytes32, emitter address, eventName_ bytes32, account address, method bytes4)
func (_Ethbase *EthbaseFilterer) FilterSubscribed(opts *bind.FilterOpts) (*EthbaseSubscribedIterator, error) {

	logs, sub, err := _Ethbase.contract.FilterLogs(opts, "Subscribed")
	if err != nil {
		return nil, err
	}
	return &EthbaseSubscribedIterator{contract: _Ethbase.contract, event: "Subscribed", logs: logs, sub: sub}, nil
}

// WatchSubscribed is a free log subscription operation binding the contract event 0x04c9972fe596c2afdd8205c5d5cff160fdbbfc5254280890ce839b51fcdcad2f.
//
// Solidity: e Subscribed(eventId bytes32, emitter address, eventName_ bytes32, account address, method bytes4)
func (_Ethbase *EthbaseFilterer) WatchSubscribed(opts *bind.WatchOpts, sink chan<- *EthbaseSubscribed) (event.Subscription, error) {

	logs, sub, err := _Ethbase.contract.WatchLogs(opts, "Subscribed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthbaseSubscribed)
				if err := _Ethbase.contract.UnpackLog(event, "Subscribed", log); err != nil {
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
