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

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"eventId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"emitter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"eventName_\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"method\",\"type\":\"bytes4\"}],\"name\":\"Subscribed\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"emitter_\",\"type\":\"address\"},{\"name\":\"eventName_\",\"type\":\"bytes32\"},{\"name\":\"account_\",\"type\":\"address\"},{\"name\":\"method_\",\"type\":\"bytes4\"}],\"name\":\"subscribe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"eventId_\",\"type\":\"bytes32\"},{\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"unsubscribe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"eventId_\",\"type\":\"bytes32\"},{\"name\":\"account_\",\"type\":\"address\"},{\"name\":\"args_\",\"type\":\"bytes\"}],\"name\":\"invoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// Invoke is a paid mutator transaction binding the contract method 0x2e17e4c8.
//
// Solidity: function invoke(eventId_ bytes32, account_ address, args_ bytes) returns()
func (_Registry *RegistryTransactor) Invoke(opts *bind.TransactOpts, eventId_ [32]byte, account_ common.Address, args_ []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "invoke", eventId_, account_, args_)
}

// Invoke is a paid mutator transaction binding the contract method 0x2e17e4c8.
//
// Solidity: function invoke(eventId_ bytes32, account_ address, args_ bytes) returns()
func (_Registry *RegistrySession) Invoke(eventId_ [32]byte, account_ common.Address, args_ []byte) (*types.Transaction, error) {
	return _Registry.Contract.Invoke(&_Registry.TransactOpts, eventId_, account_, args_)
}

// Invoke is a paid mutator transaction binding the contract method 0x2e17e4c8.
//
// Solidity: function invoke(eventId_ bytes32, account_ address, args_ bytes) returns()
func (_Registry *RegistryTransactorSession) Invoke(eventId_ [32]byte, account_ common.Address, args_ []byte) (*types.Transaction, error) {
	return _Registry.Contract.Invoke(&_Registry.TransactOpts, eventId_, account_, args_)
}

// Subscribe is a paid mutator transaction binding the contract method 0x4ff57300.
//
// Solidity: function subscribe(emitter_ address, eventName_ bytes32, account_ address, method_ bytes4) returns()
func (_Registry *RegistryTransactor) Subscribe(opts *bind.TransactOpts, emitter_ common.Address, eventName_ [32]byte, account_ common.Address, method_ [4]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "subscribe", emitter_, eventName_, account_, method_)
}

// Subscribe is a paid mutator transaction binding the contract method 0x4ff57300.
//
// Solidity: function subscribe(emitter_ address, eventName_ bytes32, account_ address, method_ bytes4) returns()
func (_Registry *RegistrySession) Subscribe(emitter_ common.Address, eventName_ [32]byte, account_ common.Address, method_ [4]byte) (*types.Transaction, error) {
	return _Registry.Contract.Subscribe(&_Registry.TransactOpts, emitter_, eventName_, account_, method_)
}

// Subscribe is a paid mutator transaction binding the contract method 0x4ff57300.
//
// Solidity: function subscribe(emitter_ address, eventName_ bytes32, account_ address, method_ bytes4) returns()
func (_Registry *RegistryTransactorSession) Subscribe(emitter_ common.Address, eventName_ [32]byte, account_ common.Address, method_ [4]byte) (*types.Transaction, error) {
	return _Registry.Contract.Subscribe(&_Registry.TransactOpts, emitter_, eventName_, account_, method_)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xd59fd22f.
//
// Solidity: function unsubscribe(eventId_ bytes32, account_ address) returns()
func (_Registry *RegistryTransactor) Unsubscribe(opts *bind.TransactOpts, eventId_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unsubscribe", eventId_, account_)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xd59fd22f.
//
// Solidity: function unsubscribe(eventId_ bytes32, account_ address) returns()
func (_Registry *RegistrySession) Unsubscribe(eventId_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Unsubscribe(&_Registry.TransactOpts, eventId_, account_)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xd59fd22f.
//
// Solidity: function unsubscribe(eventId_ bytes32, account_ address) returns()
func (_Registry *RegistryTransactorSession) Unsubscribe(eventId_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Unsubscribe(&_Registry.TransactOpts, eventId_, account_)
}

// RegistrySubscribedIterator is returned from FilterSubscribed and is used to iterate over the raw logs and unpacked data for Subscribed events raised by the Registry contract.
type RegistrySubscribedIterator struct {
	Event *RegistrySubscribed // Event containing the contract specifics and raw log

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
func (it *RegistrySubscribedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrySubscribed)
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
		it.Event = new(RegistrySubscribed)
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
func (it *RegistrySubscribedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrySubscribedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrySubscribed represents a Subscribed event raised by the Registry contract.
type RegistrySubscribed struct {
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
func (_Registry *RegistryFilterer) FilterSubscribed(opts *bind.FilterOpts) (*RegistrySubscribedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Subscribed")
	if err != nil {
		return nil, err
	}
	return &RegistrySubscribedIterator{contract: _Registry.contract, event: "Subscribed", logs: logs, sub: sub}, nil
}

// WatchSubscribed is a free log subscription operation binding the contract event 0x04c9972fe596c2afdd8205c5d5cff160fdbbfc5254280890ce839b51fcdcad2f.
//
// Solidity: e Subscribed(eventId bytes32, emitter address, eventName_ bytes32, account address, method bytes4)
func (_Registry *RegistryFilterer) WatchSubscribed(opts *bind.WatchOpts, sink chan<- *RegistrySubscribed) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Subscribed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrySubscribed)
				if err := _Registry.contract.UnpackLog(event, "Subscribed", log); err != nil {
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
