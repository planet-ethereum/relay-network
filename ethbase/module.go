// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbase

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ModuleABI is the input ABI used to generate the binding from.
const ModuleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_masterCopy\",\"type\":\"address\"}],\"name\":\"changeMasterCopy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"logs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isExecuted\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receipt\",\"type\":\"bytes\"},{\"name\":\"_parentNodes\",\"type\":\"bytes\"},{\"name\":\"_key\",\"type\":\"bytes\"},{\"name\":\"_logIndex\",\"type\":\"uint256\"},{\"name\":\"_blockHeader\",\"type\":\"bytes\"},{\"name\":\"_emitter\",\"type\":\"address\"},{\"name\":\"_eventTopic\",\"type\":\"bytes32\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_operation\",\"type\":\"uint8\"},{\"name\":\"_nonce\",\"type\":\"uint256\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"execConditional\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"operation\",\"type\":\"uint8\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getTransactionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Module is an auto generated Go binding around an Ethereum contract.
type Module struct {
	ModuleCaller     // Read-only binding to the contract
	ModuleTransactor // Write-only binding to the contract
	ModuleFilterer   // Log filterer for contract events
}

// ModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModuleSession struct {
	Contract     *Module           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModuleCallerSession struct {
	Contract *ModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModuleTransactorSession struct {
	Contract     *ModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModuleRaw struct {
	Contract *Module // Generic contract binding to access the raw methods on
}

// ModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModuleCallerRaw struct {
	Contract *ModuleCaller // Generic read-only contract binding to access the raw methods on
}

// ModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModuleTransactorRaw struct {
	Contract *ModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModule creates a new instance of Module, bound to a specific deployed contract.
func NewModule(address common.Address, backend bind.ContractBackend) (*Module, error) {
	contract, err := bindModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Module{ModuleCaller: ModuleCaller{contract: contract}, ModuleTransactor: ModuleTransactor{contract: contract}, ModuleFilterer: ModuleFilterer{contract: contract}}, nil
}

// NewModuleCaller creates a new read-only instance of Module, bound to a specific deployed contract.
func NewModuleCaller(address common.Address, caller bind.ContractCaller) (*ModuleCaller, error) {
	contract, err := bindModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModuleCaller{contract: contract}, nil
}

// NewModuleTransactor creates a new write-only instance of Module, bound to a specific deployed contract.
func NewModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*ModuleTransactor, error) {
	contract, err := bindModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModuleTransactor{contract: contract}, nil
}

// NewModuleFilterer creates a new log filterer instance of Module, bound to a specific deployed contract.
func NewModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*ModuleFilterer, error) {
	contract, err := bindModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModuleFilterer{contract: contract}, nil
}

// bindModule binds a generic wrapper to an already deployed contract.
func bindModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ModuleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Module *ModuleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Module.Contract.ModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Module *ModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Module.Contract.ModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Module *ModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Module.Contract.ModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Module *ModuleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Module.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Module *ModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Module.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Module *ModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Module.Contract.contract.Transact(opts, method, params...)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0x2b500041.
//
// Solidity: function getTransactionHash(to address, value uint256, data bytes, operation uint8, _nonce uint256) constant returns(bytes32)
func (_Module *ModuleCaller) GetTransactionHash(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Module.contract.Call(opts, out, "getTransactionHash", to, value, data, operation, _nonce)
	return *ret0, err
}

// GetTransactionHash is a free data retrieval call binding the contract method 0x2b500041.
//
// Solidity: function getTransactionHash(to address, value uint256, data bytes, operation uint8, _nonce uint256) constant returns(bytes32)
func (_Module *ModuleSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, _nonce *big.Int) ([32]byte, error) {
	return _Module.Contract.GetTransactionHash(&_Module.CallOpts, to, value, data, operation, _nonce)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0x2b500041.
//
// Solidity: function getTransactionHash(to address, value uint256, data bytes, operation uint8, _nonce uint256) constant returns(bytes32)
func (_Module *ModuleCallerSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, _nonce *big.Int) ([32]byte, error) {
	return _Module.Contract.GetTransactionHash(&_Module.CallOpts, to, value, data, operation, _nonce)
}

// IsExecuted is a free data retrieval call binding the contract method 0xe52cb36a.
//
// Solidity: function isExecuted( bytes32) constant returns(uint256)
func (_Module *ModuleCaller) IsExecuted(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Module.contract.Call(opts, out, "isExecuted", arg0)
	return *ret0, err
}

// IsExecuted is a free data retrieval call binding the contract method 0xe52cb36a.
//
// Solidity: function isExecuted( bytes32) constant returns(uint256)
func (_Module *ModuleSession) IsExecuted(arg0 [32]byte) (*big.Int, error) {
	return _Module.Contract.IsExecuted(&_Module.CallOpts, arg0)
}

// IsExecuted is a free data retrieval call binding the contract method 0xe52cb36a.
//
// Solidity: function isExecuted( bytes32) constant returns(uint256)
func (_Module *ModuleCallerSession) IsExecuted(arg0 [32]byte) (*big.Int, error) {
	return _Module.Contract.IsExecuted(&_Module.CallOpts, arg0)
}

// Logs is a free data retrieval call binding the contract method 0x9a02b653.
//
// Solidity: function logs( bytes32) constant returns(uint256)
func (_Module *ModuleCaller) Logs(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Module.contract.Call(opts, out, "logs", arg0)
	return *ret0, err
}

// Logs is a free data retrieval call binding the contract method 0x9a02b653.
//
// Solidity: function logs( bytes32) constant returns(uint256)
func (_Module *ModuleSession) Logs(arg0 [32]byte) (*big.Int, error) {
	return _Module.Contract.Logs(&_Module.CallOpts, arg0)
}

// Logs is a free data retrieval call binding the contract method 0x9a02b653.
//
// Solidity: function logs( bytes32) constant returns(uint256)
func (_Module *ModuleCallerSession) Logs(arg0 [32]byte) (*big.Int, error) {
	return _Module.Contract.Logs(&_Module.CallOpts, arg0)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Module *ModuleCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Module.contract.Call(opts, out, "manager")
	return *ret0, err
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Module *ModuleSession) Manager() (common.Address, error) {
	return _Module.Contract.Manager(&_Module.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Module *ModuleCallerSession) Manager() (common.Address, error) {
	return _Module.Contract.Manager(&_Module.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_Module *ModuleCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Module.contract.Call(opts, out, "nonce")
	return *ret0, err
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_Module *ModuleSession) Nonce() (*big.Int, error) {
	return _Module.Contract.Nonce(&_Module.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_Module *ModuleCallerSession) Nonce() (*big.Int, error) {
	return _Module.Contract.Nonce(&_Module.CallOpts)
}

// ChangeMasterCopy is a paid mutator transaction binding the contract method 0x7de7edef.
//
// Solidity: function changeMasterCopy(_masterCopy address) returns()
func (_Module *ModuleTransactor) ChangeMasterCopy(opts *bind.TransactOpts, _masterCopy common.Address) (*types.Transaction, error) {
	return _Module.contract.Transact(opts, "changeMasterCopy", _masterCopy)
}

// ChangeMasterCopy is a paid mutator transaction binding the contract method 0x7de7edef.
//
// Solidity: function changeMasterCopy(_masterCopy address) returns()
func (_Module *ModuleSession) ChangeMasterCopy(_masterCopy common.Address) (*types.Transaction, error) {
	return _Module.Contract.ChangeMasterCopy(&_Module.TransactOpts, _masterCopy)
}

// ChangeMasterCopy is a paid mutator transaction binding the contract method 0x7de7edef.
//
// Solidity: function changeMasterCopy(_masterCopy address) returns()
func (_Module *ModuleTransactorSession) ChangeMasterCopy(_masterCopy common.Address) (*types.Transaction, error) {
	return _Module.Contract.ChangeMasterCopy(&_Module.TransactOpts, _masterCopy)
}

// ExecConditional is a paid mutator transaction binding the contract method 0x605f2a9c.
//
// Solidity: function execConditional(_receipt bytes, _parentNodes bytes, _key bytes, _logIndex uint256, _blockHeader bytes, _emitter address, _eventTopic bytes32, _to address, _value uint256, _data bytes, _operation uint8, _nonce uint256, _signatures bytes) returns()
func (_Module *ModuleTransactor) ExecConditional(opts *bind.TransactOpts, _receipt []byte, _parentNodes []byte, _key []byte, _logIndex *big.Int, _blockHeader []byte, _emitter common.Address, _eventTopic [32]byte, _to common.Address, _value *big.Int, _data []byte, _operation uint8, _nonce *big.Int, _signatures []byte) (*types.Transaction, error) {
	return _Module.contract.Transact(opts, "execConditional", _receipt, _parentNodes, _key, _logIndex, _blockHeader, _emitter, _eventTopic, _to, _value, _data, _operation, _nonce, _signatures)
}

// ExecConditional is a paid mutator transaction binding the contract method 0x605f2a9c.
//
// Solidity: function execConditional(_receipt bytes, _parentNodes bytes, _key bytes, _logIndex uint256, _blockHeader bytes, _emitter address, _eventTopic bytes32, _to address, _value uint256, _data bytes, _operation uint8, _nonce uint256, _signatures bytes) returns()
func (_Module *ModuleSession) ExecConditional(_receipt []byte, _parentNodes []byte, _key []byte, _logIndex *big.Int, _blockHeader []byte, _emitter common.Address, _eventTopic [32]byte, _to common.Address, _value *big.Int, _data []byte, _operation uint8, _nonce *big.Int, _signatures []byte) (*types.Transaction, error) {
	return _Module.Contract.ExecConditional(&_Module.TransactOpts, _receipt, _parentNodes, _key, _logIndex, _blockHeader, _emitter, _eventTopic, _to, _value, _data, _operation, _nonce, _signatures)
}

// ExecConditional is a paid mutator transaction binding the contract method 0x605f2a9c.
//
// Solidity: function execConditional(_receipt bytes, _parentNodes bytes, _key bytes, _logIndex uint256, _blockHeader bytes, _emitter address, _eventTopic bytes32, _to address, _value uint256, _data bytes, _operation uint8, _nonce uint256, _signatures bytes) returns()
func (_Module *ModuleTransactorSession) ExecConditional(_receipt []byte, _parentNodes []byte, _key []byte, _logIndex *big.Int, _blockHeader []byte, _emitter common.Address, _eventTopic [32]byte, _to common.Address, _value *big.Int, _data []byte, _operation uint8, _nonce *big.Int, _signatures []byte) (*types.Transaction, error) {
	return _Module.Contract.ExecConditional(&_Module.TransactOpts, _receipt, _parentNodes, _key, _logIndex, _blockHeader, _emitter, _eventTopic, _to, _value, _data, _operation, _nonce, _signatures)
}

// Setup is a paid mutator transaction binding the contract method 0xba0bba40.
//
// Solidity: function setup() returns()
func (_Module *ModuleTransactor) Setup(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Module.contract.Transact(opts, "setup")
}

// Setup is a paid mutator transaction binding the contract method 0xba0bba40.
//
// Solidity: function setup() returns()
func (_Module *ModuleSession) Setup() (*types.Transaction, error) {
	return _Module.Contract.Setup(&_Module.TransactOpts)
}

// Setup is a paid mutator transaction binding the contract method 0xba0bba40.
//
// Solidity: function setup() returns()
func (_Module *ModuleTransactorSession) Setup() (*types.Transaction, error) {
	return _Module.Contract.Setup(&_Module.TransactOpts)
}
