// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package safe

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// SafeABI is the input ABI used to generate the binding from.
const SafeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"addOwnerWithThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR_TYPEHASH\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModule\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"module\",\"type\":\"address\"}],\"name\":\"enableModule\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_masterCopy\",\"type\":\"address\"}],\"name\":\"changeMasterCopy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SENTINEL_MODULES\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SENTINEL_OWNERS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getModules\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SAFE_MSG_TYPEHASH\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SAFE_TX_TYPEHASH\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"prevModule\",\"type\":\"address\"},{\"name\":\"module\",\"type\":\"address\"}],\"name\":\"disableModule\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"prevOwner\",\"type\":\"address\"},{\"name\":\"oldOwner\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"swapOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"prevOwner\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"ExecutionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"RemovedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ChangedThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"module\",\"type\":\"address\"}],\"name\":\"EnabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"module\",\"type\":\"address\"}],\"name\":\"DisabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractCreation\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_threshold\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"operation\",\"type\":\"uint8\"},{\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"name\":\"dataGas\",\"type\":\"uint256\"},{\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"name\":\"gasToken\",\"type\":\"address\"},{\"name\":\"refundReceiver\",\"type\":\"address\"},{\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"execTransaction\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"requiredTxGas\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hashToApprove\",\"type\":\"bytes32\"}],\"name\":\"approveHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"signMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"name\":\"isValid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"operation\",\"type\":\"uint8\"},{\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"name\":\"dataGas\",\"type\":\"uint256\"},{\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"name\":\"gasToken\",\"type\":\"address\"},{\"name\":\"refundReceiver\",\"type\":\"address\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"encodeTransactionData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"operation\",\"type\":\"uint8\"},{\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"name\":\"dataGas\",\"type\":\"uint256\"},{\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"name\":\"gasToken\",\"type\":\"address\"},{\"name\":\"refundReceiver\",\"type\":\"address\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getTransactionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Safe is an auto generated Go binding around an Ethereum contract.
type Safe struct {
	SafeCaller     // Read-only binding to the contract
	SafeTransactor // Write-only binding to the contract
	SafeFilterer   // Log filterer for contract events
}

// SafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeSession struct {
	Contract     *Safe             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCallerSession struct {
	Contract *SafeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeTransactorSession struct {
	Contract     *SafeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeRaw struct {
	Contract *Safe // Generic contract binding to access the raw methods on
}

// SafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCallerRaw struct {
	Contract *SafeCaller // Generic read-only contract binding to access the raw methods on
}

// SafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeTransactorRaw struct {
	Contract *SafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafe creates a new instance of Safe, bound to a specific deployed contract.
func NewSafe(address common.Address, backend bind.ContractBackend) (*Safe, error) {
	contract, err := bindSafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Safe{SafeCaller: SafeCaller{contract: contract}, SafeTransactor: SafeTransactor{contract: contract}, SafeFilterer: SafeFilterer{contract: contract}}, nil
}

// NewSafeCaller creates a new read-only instance of Safe, bound to a specific deployed contract.
func NewSafeCaller(address common.Address, caller bind.ContractCaller) (*SafeCaller, error) {
	contract, err := bindSafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCaller{contract: contract}, nil
}

// NewSafeTransactor creates a new write-only instance of Safe, bound to a specific deployed contract.
func NewSafeTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeTransactor, error) {
	contract, err := bindSafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeTransactor{contract: contract}, nil
}

// NewSafeFilterer creates a new log filterer instance of Safe, bound to a specific deployed contract.
func NewSafeFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeFilterer, error) {
	contract, err := bindSafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeFilterer{contract: contract}, nil
}

// bindSafe binds a generic wrapper to an already deployed contract.
func bindSafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safe *SafeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Safe.Contract.SafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safe *SafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safe.Contract.SafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safe *SafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safe.Contract.SafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safe *SafeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Safe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safe *SafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safe *SafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safe.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATORTYPEHASH is a free data retrieval call binding the contract method 0x1db61b54.
//
// Solidity: function DOMAIN_SEPARATOR_TYPEHASH() constant returns(bytes32)
func (_Safe *SafeCaller) DOMAINSEPARATORTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "DOMAIN_SEPARATOR_TYPEHASH")
	return *ret0, err
}

// DOMAINSEPARATORTYPEHASH is a free data retrieval call binding the contract method 0x1db61b54.
//
// Solidity: function DOMAIN_SEPARATOR_TYPEHASH() constant returns(bytes32)
func (_Safe *SafeSession) DOMAINSEPARATORTYPEHASH() ([32]byte, error) {
	return _Safe.Contract.DOMAINSEPARATORTYPEHASH(&_Safe.CallOpts)
}

// DOMAINSEPARATORTYPEHASH is a free data retrieval call binding the contract method 0x1db61b54.
//
// Solidity: function DOMAIN_SEPARATOR_TYPEHASH() constant returns(bytes32)
func (_Safe *SafeCallerSession) DOMAINSEPARATORTYPEHASH() ([32]byte, error) {
	return _Safe.Contract.DOMAINSEPARATORTYPEHASH(&_Safe.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() constant returns(string)
func (_Safe *SafeCaller) NAME(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "NAME")
	return *ret0, err
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() constant returns(string)
func (_Safe *SafeSession) NAME() (string, error) {
	return _Safe.Contract.NAME(&_Safe.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() constant returns(string)
func (_Safe *SafeCallerSession) NAME() (string, error) {
	return _Safe.Contract.NAME(&_Safe.CallOpts)
}

// SAFEMSGTYPEHASH is a free data retrieval call binding the contract method 0xc0856ffc.
//
// Solidity: function SAFE_MSG_TYPEHASH() constant returns(bytes32)
func (_Safe *SafeCaller) SAFEMSGTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "SAFE_MSG_TYPEHASH")
	return *ret0, err
}

// SAFEMSGTYPEHASH is a free data retrieval call binding the contract method 0xc0856ffc.
//
// Solidity: function SAFE_MSG_TYPEHASH() constant returns(bytes32)
func (_Safe *SafeSession) SAFEMSGTYPEHASH() ([32]byte, error) {
	return _Safe.Contract.SAFEMSGTYPEHASH(&_Safe.CallOpts)
}

// SAFEMSGTYPEHASH is a free data retrieval call binding the contract method 0xc0856ffc.
//
// Solidity: function SAFE_MSG_TYPEHASH() constant returns(bytes32)
func (_Safe *SafeCallerSession) SAFEMSGTYPEHASH() ([32]byte, error) {
	return _Safe.Contract.SAFEMSGTYPEHASH(&_Safe.CallOpts)
}

// SAFETXTYPEHASH is a free data retrieval call binding the contract method 0xccafc387.
//
// Solidity: function SAFE_TX_TYPEHASH() constant returns(bytes32)
func (_Safe *SafeCaller) SAFETXTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "SAFE_TX_TYPEHASH")
	return *ret0, err
}

// SAFETXTYPEHASH is a free data retrieval call binding the contract method 0xccafc387.
//
// Solidity: function SAFE_TX_TYPEHASH() constant returns(bytes32)
func (_Safe *SafeSession) SAFETXTYPEHASH() ([32]byte, error) {
	return _Safe.Contract.SAFETXTYPEHASH(&_Safe.CallOpts)
}

// SAFETXTYPEHASH is a free data retrieval call binding the contract method 0xccafc387.
//
// Solidity: function SAFE_TX_TYPEHASH() constant returns(bytes32)
func (_Safe *SafeCallerSession) SAFETXTYPEHASH() ([32]byte, error) {
	return _Safe.Contract.SAFETXTYPEHASH(&_Safe.CallOpts)
}

// SENTINELMODULES is a free data retrieval call binding the contract method 0x85e332cd.
//
// Solidity: function SENTINEL_MODULES() constant returns(address)
func (_Safe *SafeCaller) SENTINELMODULES(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "SENTINEL_MODULES")
	return *ret0, err
}

// SENTINELMODULES is a free data retrieval call binding the contract method 0x85e332cd.
//
// Solidity: function SENTINEL_MODULES() constant returns(address)
func (_Safe *SafeSession) SENTINELMODULES() (common.Address, error) {
	return _Safe.Contract.SENTINELMODULES(&_Safe.CallOpts)
}

// SENTINELMODULES is a free data retrieval call binding the contract method 0x85e332cd.
//
// Solidity: function SENTINEL_MODULES() constant returns(address)
func (_Safe *SafeCallerSession) SENTINELMODULES() (common.Address, error) {
	return _Safe.Contract.SENTINELMODULES(&_Safe.CallOpts)
}

// SENTINELOWNERS is a free data retrieval call binding the contract method 0x8cff6355.
//
// Solidity: function SENTINEL_OWNERS() constant returns(address)
func (_Safe *SafeCaller) SENTINELOWNERS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "SENTINEL_OWNERS")
	return *ret0, err
}

// SENTINELOWNERS is a free data retrieval call binding the contract method 0x8cff6355.
//
// Solidity: function SENTINEL_OWNERS() constant returns(address)
func (_Safe *SafeSession) SENTINELOWNERS() (common.Address, error) {
	return _Safe.Contract.SENTINELOWNERS(&_Safe.CallOpts)
}

// SENTINELOWNERS is a free data retrieval call binding the contract method 0x8cff6355.
//
// Solidity: function SENTINEL_OWNERS() constant returns(address)
func (_Safe *SafeCallerSession) SENTINELOWNERS() (common.Address, error) {
	return _Safe.Contract.SENTINELOWNERS(&_Safe.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Safe *SafeCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Safe *SafeSession) VERSION() (string, error) {
	return _Safe.Contract.VERSION(&_Safe.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Safe *SafeCallerSession) VERSION() (string, error) {
	return _Safe.Contract.VERSION(&_Safe.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() constant returns(bytes32)
func (_Safe *SafeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "domainSeparator")
	return *ret0, err
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() constant returns(bytes32)
func (_Safe *SafeSession) DomainSeparator() ([32]byte, error) {
	return _Safe.Contract.DomainSeparator(&_Safe.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() constant returns(bytes32)
func (_Safe *SafeCallerSession) DomainSeparator() ([32]byte, error) {
	return _Safe.Contract.DomainSeparator(&_Safe.CallOpts)
}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(to address, value uint256, data bytes, operation uint8, safeTxGas uint256, dataGas uint256, gasPrice uint256, gasToken address, refundReceiver address, _nonce uint256) constant returns(bytes)
func (_Safe *SafeCaller) EncodeTransactionData(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, dataGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "encodeTransactionData", to, value, data, operation, safeTxGas, dataGas, gasPrice, gasToken, refundReceiver, _nonce)
	return *ret0, err
}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(to address, value uint256, data bytes, operation uint8, safeTxGas uint256, dataGas uint256, gasPrice uint256, gasToken address, refundReceiver address, _nonce uint256) constant returns(bytes)
func (_Safe *SafeSession) EncodeTransactionData(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, dataGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	return _Safe.Contract.EncodeTransactionData(&_Safe.CallOpts, to, value, data, operation, safeTxGas, dataGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(to address, value uint256, data bytes, operation uint8, safeTxGas uint256, dataGas uint256, gasPrice uint256, gasToken address, refundReceiver address, _nonce uint256) constant returns(bytes)
func (_Safe *SafeCallerSession) EncodeTransactionData(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, dataGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	return _Safe.Contract.EncodeTransactionData(&_Safe.CallOpts, to, value, data, operation, safeTxGas, dataGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(message bytes) constant returns(bytes32)
func (_Safe *SafeCaller) GetMessageHash(opts *bind.CallOpts, message []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "getMessageHash", message)
	return *ret0, err
}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(message bytes) constant returns(bytes32)
func (_Safe *SafeSession) GetMessageHash(message []byte) ([32]byte, error) {
	return _Safe.Contract.GetMessageHash(&_Safe.CallOpts, message)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(message bytes) constant returns(bytes32)
func (_Safe *SafeCallerSession) GetMessageHash(message []byte) ([32]byte, error) {
	return _Safe.Contract.GetMessageHash(&_Safe.CallOpts, message)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() constant returns(address[])
func (_Safe *SafeCaller) GetModules(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "getModules")
	return *ret0, err
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() constant returns(address[])
func (_Safe *SafeSession) GetModules() ([]common.Address, error) {
	return _Safe.Contract.GetModules(&_Safe.CallOpts)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() constant returns(address[])
func (_Safe *SafeCallerSession) GetModules() ([]common.Address, error) {
	return _Safe.Contract.GetModules(&_Safe.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_Safe *SafeCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "getOwners")
	return *ret0, err
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_Safe *SafeSession) GetOwners() ([]common.Address, error) {
	return _Safe.Contract.GetOwners(&_Safe.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_Safe *SafeCallerSession) GetOwners() ([]common.Address, error) {
	return _Safe.Contract.GetOwners(&_Safe.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() constant returns(uint256)
func (_Safe *SafeCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "getThreshold")
	return *ret0, err
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() constant returns(uint256)
func (_Safe *SafeSession) GetThreshold() (*big.Int, error) {
	return _Safe.Contract.GetThreshold(&_Safe.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() constant returns(uint256)
func (_Safe *SafeCallerSession) GetThreshold() (*big.Int, error) {
	return _Safe.Contract.GetThreshold(&_Safe.CallOpts)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(to address, value uint256, data bytes, operation uint8, safeTxGas uint256, dataGas uint256, gasPrice uint256, gasToken address, refundReceiver address, _nonce uint256) constant returns(bytes32)
func (_Safe *SafeCaller) GetTransactionHash(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, dataGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "getTransactionHash", to, value, data, operation, safeTxGas, dataGas, gasPrice, gasToken, refundReceiver, _nonce)
	return *ret0, err
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(to address, value uint256, data bytes, operation uint8, safeTxGas uint256, dataGas uint256, gasPrice uint256, gasToken address, refundReceiver address, _nonce uint256) constant returns(bytes32)
func (_Safe *SafeSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, dataGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _Safe.Contract.GetTransactionHash(&_Safe.CallOpts, to, value, data, operation, safeTxGas, dataGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(to address, value uint256, data bytes, operation uint8, safeTxGas uint256, dataGas uint256, gasPrice uint256, gasToken address, refundReceiver address, _nonce uint256) constant returns(bytes32)
func (_Safe *SafeCallerSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, dataGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _Safe.Contract.GetTransactionHash(&_Safe.CallOpts, to, value, data, operation, safeTxGas, dataGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(owner address) constant returns(bool)
func (_Safe *SafeCaller) IsOwner(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "isOwner", owner)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(owner address) constant returns(bool)
func (_Safe *SafeSession) IsOwner(owner common.Address) (bool, error) {
	return _Safe.Contract.IsOwner(&_Safe.CallOpts, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(owner address) constant returns(bool)
func (_Safe *SafeCallerSession) IsOwner(owner common.Address) (bool, error) {
	return _Safe.Contract.IsOwner(&_Safe.CallOpts, owner)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_Safe *SafeCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "nonce")
	return *ret0, err
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_Safe *SafeSession) Nonce() (*big.Int, error) {
	return _Safe.Contract.Nonce(&_Safe.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_Safe *SafeCallerSession) Nonce() (*big.Int, error) {
	return _Safe.Contract.Nonce(&_Safe.CallOpts)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(owner address, _threshold uint256) returns()
func (_Safe *SafeTransactor) AddOwnerWithThreshold(opts *bind.TransactOpts, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "addOwnerWithThreshold", owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(owner address, _threshold uint256) returns()
func (_Safe *SafeSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Safe.Contract.AddOwnerWithThreshold(&_Safe.TransactOpts, owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(owner address, _threshold uint256) returns()
func (_Safe *SafeTransactorSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Safe.Contract.AddOwnerWithThreshold(&_Safe.TransactOpts, owner, _threshold)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(hashToApprove bytes32) returns()
func (_Safe *SafeTransactor) ApproveHash(opts *bind.TransactOpts, hashToApprove [32]byte) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "approveHash", hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(hashToApprove bytes32) returns()
func (_Safe *SafeSession) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _Safe.Contract.ApproveHash(&_Safe.TransactOpts, hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(hashToApprove bytes32) returns()
func (_Safe *SafeTransactorSession) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _Safe.Contract.ApproveHash(&_Safe.TransactOpts, hashToApprove)
}

// ChangeMasterCopy is a paid mutator transaction binding the contract method 0x7de7edef.
//
// Solidity: function changeMasterCopy(_masterCopy address) returns()
func (_Safe *SafeTransactor) ChangeMasterCopy(opts *bind.TransactOpts, _masterCopy common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "changeMasterCopy", _masterCopy)
}

// ChangeMasterCopy is a paid mutator transaction binding the contract method 0x7de7edef.
//
// Solidity: function changeMasterCopy(_masterCopy address) returns()
func (_Safe *SafeSession) ChangeMasterCopy(_masterCopy common.Address) (*types.Transaction, error) {
	return _Safe.Contract.ChangeMasterCopy(&_Safe.TransactOpts, _masterCopy)
}

// ChangeMasterCopy is a paid mutator transaction binding the contract method 0x7de7edef.
//
// Solidity: function changeMasterCopy(_masterCopy address) returns()
func (_Safe *SafeTransactorSession) ChangeMasterCopy(_masterCopy common.Address) (*types.Transaction, error) {
	return _Safe.Contract.ChangeMasterCopy(&_Safe.TransactOpts, _masterCopy)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(_threshold uint256) returns()
func (_Safe *SafeTransactor) ChangeThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "changeThreshold", _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(_threshold uint256) returns()
func (_Safe *SafeSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _Safe.Contract.ChangeThreshold(&_Safe.TransactOpts, _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(_threshold uint256) returns()
func (_Safe *SafeTransactorSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _Safe.Contract.ChangeThreshold(&_Safe.TransactOpts, _threshold)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(prevModule address, module address) returns()
func (_Safe *SafeTransactor) DisableModule(opts *bind.TransactOpts, prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "disableModule", prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(prevModule address, module address) returns()
func (_Safe *SafeSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _Safe.Contract.DisableModule(&_Safe.TransactOpts, prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(prevModule address, module address) returns()
func (_Safe *SafeTransactorSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _Safe.Contract.DisableModule(&_Safe.TransactOpts, prevModule, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(module address) returns()
func (_Safe *SafeTransactor) EnableModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "enableModule", module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(module address) returns()
func (_Safe *SafeSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _Safe.Contract.EnableModule(&_Safe.TransactOpts, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(module address) returns()
func (_Safe *SafeTransactorSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _Safe.Contract.EnableModule(&_Safe.TransactOpts, module)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(to address, value uint256, data bytes, operation uint8, safeTxGas uint256, dataGas uint256, gasPrice uint256, gasToken address, refundReceiver address, signatures bytes) returns(success bool)
func (_Safe *SafeTransactor) ExecTransaction(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, dataGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "execTransaction", to, value, data, operation, safeTxGas, dataGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(to address, value uint256, data bytes, operation uint8, safeTxGas uint256, dataGas uint256, gasPrice uint256, gasToken address, refundReceiver address, signatures bytes) returns(success bool)
func (_Safe *SafeSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, dataGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _Safe.Contract.ExecTransaction(&_Safe.TransactOpts, to, value, data, operation, safeTxGas, dataGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(to address, value uint256, data bytes, operation uint8, safeTxGas uint256, dataGas uint256, gasPrice uint256, gasToken address, refundReceiver address, signatures bytes) returns(success bool)
func (_Safe *SafeTransactorSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, dataGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _Safe.Contract.ExecTransaction(&_Safe.TransactOpts, to, value, data, operation, safeTxGas, dataGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(to address, value uint256, data bytes, operation uint8) returns(success bool)
func (_Safe *SafeTransactor) ExecTransactionFromModule(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "execTransactionFromModule", to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(to address, value uint256, data bytes, operation uint8) returns(success bool)
func (_Safe *SafeSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Safe.Contract.ExecTransactionFromModule(&_Safe.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(to address, value uint256, data bytes, operation uint8) returns(success bool)
func (_Safe *SafeTransactorSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Safe.Contract.ExecTransactionFromModule(&_Safe.TransactOpts, to, value, data, operation)
}

// IsValidSignature is a paid mutator transaction binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(_data bytes, _signature bytes) returns(isValid bool)
func (_Safe *SafeTransactor) IsValidSignature(opts *bind.TransactOpts, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "isValidSignature", _data, _signature)
}

// IsValidSignature is a paid mutator transaction binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(_data bytes, _signature bytes) returns(isValid bool)
func (_Safe *SafeSession) IsValidSignature(_data []byte, _signature []byte) (*types.Transaction, error) {
	return _Safe.Contract.IsValidSignature(&_Safe.TransactOpts, _data, _signature)
}

// IsValidSignature is a paid mutator transaction binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(_data bytes, _signature bytes) returns(isValid bool)
func (_Safe *SafeTransactorSession) IsValidSignature(_data []byte, _signature []byte) (*types.Transaction, error) {
	return _Safe.Contract.IsValidSignature(&_Safe.TransactOpts, _data, _signature)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(prevOwner address, owner address, _threshold uint256) returns()
func (_Safe *SafeTransactor) RemoveOwner(opts *bind.TransactOpts, prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "removeOwner", prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(prevOwner address, owner address, _threshold uint256) returns()
func (_Safe *SafeSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Safe.Contract.RemoveOwner(&_Safe.TransactOpts, prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(prevOwner address, owner address, _threshold uint256) returns()
func (_Safe *SafeTransactorSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Safe.Contract.RemoveOwner(&_Safe.TransactOpts, prevOwner, owner, _threshold)
}

// RequiredTxGas is a paid mutator transaction binding the contract method 0xc4ca3a9c.
//
// Solidity: function requiredTxGas(to address, value uint256, data bytes, operation uint8) returns(uint256)
func (_Safe *SafeTransactor) RequiredTxGas(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "requiredTxGas", to, value, data, operation)
}

// RequiredTxGas is a paid mutator transaction binding the contract method 0xc4ca3a9c.
//
// Solidity: function requiredTxGas(to address, value uint256, data bytes, operation uint8) returns(uint256)
func (_Safe *SafeSession) RequiredTxGas(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Safe.Contract.RequiredTxGas(&_Safe.TransactOpts, to, value, data, operation)
}

// RequiredTxGas is a paid mutator transaction binding the contract method 0xc4ca3a9c.
//
// Solidity: function requiredTxGas(to address, value uint256, data bytes, operation uint8) returns(uint256)
func (_Safe *SafeTransactorSession) RequiredTxGas(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Safe.Contract.RequiredTxGas(&_Safe.TransactOpts, to, value, data, operation)
}

// Setup is a paid mutator transaction binding the contract method 0x0ec78d9e.
//
// Solidity: function setup(_owners address[], _threshold uint256, to address, data bytes) returns()
func (_Safe *SafeTransactor) Setup(opts *bind.TransactOpts, _owners []common.Address, _threshold *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "setup", _owners, _threshold, to, data)
}

// Setup is a paid mutator transaction binding the contract method 0x0ec78d9e.
//
// Solidity: function setup(_owners address[], _threshold uint256, to address, data bytes) returns()
func (_Safe *SafeSession) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Safe.Contract.Setup(&_Safe.TransactOpts, _owners, _threshold, to, data)
}

// Setup is a paid mutator transaction binding the contract method 0x0ec78d9e.
//
// Solidity: function setup(_owners address[], _threshold uint256, to address, data bytes) returns()
func (_Safe *SafeTransactorSession) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Safe.Contract.Setup(&_Safe.TransactOpts, _owners, _threshold, to, data)
}

// SignMessage is a paid mutator transaction binding the contract method 0x85a5affe.
//
// Solidity: function signMessage(_data bytes) returns()
func (_Safe *SafeTransactor) SignMessage(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "signMessage", _data)
}

// SignMessage is a paid mutator transaction binding the contract method 0x85a5affe.
//
// Solidity: function signMessage(_data bytes) returns()
func (_Safe *SafeSession) SignMessage(_data []byte) (*types.Transaction, error) {
	return _Safe.Contract.SignMessage(&_Safe.TransactOpts, _data)
}

// SignMessage is a paid mutator transaction binding the contract method 0x85a5affe.
//
// Solidity: function signMessage(_data bytes) returns()
func (_Safe *SafeTransactorSession) SignMessage(_data []byte) (*types.Transaction, error) {
	return _Safe.Contract.SignMessage(&_Safe.TransactOpts, _data)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(prevOwner address, oldOwner address, newOwner address) returns()
func (_Safe *SafeTransactor) SwapOwner(opts *bind.TransactOpts, prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "swapOwner", prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(prevOwner address, oldOwner address, newOwner address) returns()
func (_Safe *SafeSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Safe.Contract.SwapOwner(&_Safe.TransactOpts, prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(prevOwner address, oldOwner address, newOwner address) returns()
func (_Safe *SafeTransactorSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Safe.Contract.SwapOwner(&_Safe.TransactOpts, prevOwner, oldOwner, newOwner)
}

// SafeAddedOwnerIterator is returned from FilterAddedOwner and is used to iterate over the raw logs and unpacked data for AddedOwner events raised by the Safe contract.
type SafeAddedOwnerIterator struct {
	Event *SafeAddedOwner // Event containing the contract specifics and raw log

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
func (it *SafeAddedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeAddedOwner)
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
		it.Event = new(SafeAddedOwner)
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
func (it *SafeAddedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeAddedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeAddedOwner represents a AddedOwner event raised by the Safe contract.
type SafeAddedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddedOwner is a free log retrieval operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: e AddedOwner(owner address)
func (_Safe *SafeFilterer) FilterAddedOwner(opts *bind.FilterOpts) (*SafeAddedOwnerIterator, error) {

	logs, sub, err := _Safe.contract.FilterLogs(opts, "AddedOwner")
	if err != nil {
		return nil, err
	}
	return &SafeAddedOwnerIterator{contract: _Safe.contract, event: "AddedOwner", logs: logs, sub: sub}, nil
}

// WatchAddedOwner is a free log subscription operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: e AddedOwner(owner address)
func (_Safe *SafeFilterer) WatchAddedOwner(opts *bind.WatchOpts, sink chan<- *SafeAddedOwner) (event.Subscription, error) {

	logs, sub, err := _Safe.contract.WatchLogs(opts, "AddedOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeAddedOwner)
				if err := _Safe.contract.UnpackLog(event, "AddedOwner", log); err != nil {
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

// SafeChangedThresholdIterator is returned from FilterChangedThreshold and is used to iterate over the raw logs and unpacked data for ChangedThreshold events raised by the Safe contract.
type SafeChangedThresholdIterator struct {
	Event *SafeChangedThreshold // Event containing the contract specifics and raw log

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
func (it *SafeChangedThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeChangedThreshold)
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
		it.Event = new(SafeChangedThreshold)
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
func (it *SafeChangedThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeChangedThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeChangedThreshold represents a ChangedThreshold event raised by the Safe contract.
type SafeChangedThreshold struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedThreshold is a free log retrieval operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: e ChangedThreshold(threshold uint256)
func (_Safe *SafeFilterer) FilterChangedThreshold(opts *bind.FilterOpts) (*SafeChangedThresholdIterator, error) {

	logs, sub, err := _Safe.contract.FilterLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return &SafeChangedThresholdIterator{contract: _Safe.contract, event: "ChangedThreshold", logs: logs, sub: sub}, nil
}

// WatchChangedThreshold is a free log subscription operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: e ChangedThreshold(threshold uint256)
func (_Safe *SafeFilterer) WatchChangedThreshold(opts *bind.WatchOpts, sink chan<- *SafeChangedThreshold) (event.Subscription, error) {

	logs, sub, err := _Safe.contract.WatchLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeChangedThreshold)
				if err := _Safe.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
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

// SafeContractCreationIterator is returned from FilterContractCreation and is used to iterate over the raw logs and unpacked data for ContractCreation events raised by the Safe contract.
type SafeContractCreationIterator struct {
	Event *SafeContractCreation // Event containing the contract specifics and raw log

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
func (it *SafeContractCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeContractCreation)
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
		it.Event = new(SafeContractCreation)
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
func (it *SafeContractCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeContractCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeContractCreation represents a ContractCreation event raised by the Safe contract.
type SafeContractCreation struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractCreation is a free log retrieval operation binding the contract event 0x4db17dd5e4732fb6da34a148104a592783ca119a1e7bb8829eba6cbadef0b511.
//
// Solidity: e ContractCreation(newContract address)
func (_Safe *SafeFilterer) FilterContractCreation(opts *bind.FilterOpts) (*SafeContractCreationIterator, error) {

	logs, sub, err := _Safe.contract.FilterLogs(opts, "ContractCreation")
	if err != nil {
		return nil, err
	}
	return &SafeContractCreationIterator{contract: _Safe.contract, event: "ContractCreation", logs: logs, sub: sub}, nil
}

// WatchContractCreation is a free log subscription operation binding the contract event 0x4db17dd5e4732fb6da34a148104a592783ca119a1e7bb8829eba6cbadef0b511.
//
// Solidity: e ContractCreation(newContract address)
func (_Safe *SafeFilterer) WatchContractCreation(opts *bind.WatchOpts, sink chan<- *SafeContractCreation) (event.Subscription, error) {

	logs, sub, err := _Safe.contract.WatchLogs(opts, "ContractCreation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeContractCreation)
				if err := _Safe.contract.UnpackLog(event, "ContractCreation", log); err != nil {
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

// SafeDisabledModuleIterator is returned from FilterDisabledModule and is used to iterate over the raw logs and unpacked data for DisabledModule events raised by the Safe contract.
type SafeDisabledModuleIterator struct {
	Event *SafeDisabledModule // Event containing the contract specifics and raw log

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
func (it *SafeDisabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeDisabledModule)
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
		it.Event = new(SafeDisabledModule)
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
func (it *SafeDisabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeDisabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeDisabledModule represents a DisabledModule event raised by the Safe contract.
type SafeDisabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisabledModule is a free log retrieval operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: e DisabledModule(module address)
func (_Safe *SafeFilterer) FilterDisabledModule(opts *bind.FilterOpts) (*SafeDisabledModuleIterator, error) {

	logs, sub, err := _Safe.contract.FilterLogs(opts, "DisabledModule")
	if err != nil {
		return nil, err
	}
	return &SafeDisabledModuleIterator{contract: _Safe.contract, event: "DisabledModule", logs: logs, sub: sub}, nil
}

// WatchDisabledModule is a free log subscription operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: e DisabledModule(module address)
func (_Safe *SafeFilterer) WatchDisabledModule(opts *bind.WatchOpts, sink chan<- *SafeDisabledModule) (event.Subscription, error) {

	logs, sub, err := _Safe.contract.WatchLogs(opts, "DisabledModule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeDisabledModule)
				if err := _Safe.contract.UnpackLog(event, "DisabledModule", log); err != nil {
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

// SafeEnabledModuleIterator is returned from FilterEnabledModule and is used to iterate over the raw logs and unpacked data for EnabledModule events raised by the Safe contract.
type SafeEnabledModuleIterator struct {
	Event *SafeEnabledModule // Event containing the contract specifics and raw log

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
func (it *SafeEnabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeEnabledModule)
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
		it.Event = new(SafeEnabledModule)
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
func (it *SafeEnabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeEnabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeEnabledModule represents a EnabledModule event raised by the Safe contract.
type SafeEnabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnabledModule is a free log retrieval operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: e EnabledModule(module address)
func (_Safe *SafeFilterer) FilterEnabledModule(opts *bind.FilterOpts) (*SafeEnabledModuleIterator, error) {

	logs, sub, err := _Safe.contract.FilterLogs(opts, "EnabledModule")
	if err != nil {
		return nil, err
	}
	return &SafeEnabledModuleIterator{contract: _Safe.contract, event: "EnabledModule", logs: logs, sub: sub}, nil
}

// WatchEnabledModule is a free log subscription operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: e EnabledModule(module address)
func (_Safe *SafeFilterer) WatchEnabledModule(opts *bind.WatchOpts, sink chan<- *SafeEnabledModule) (event.Subscription, error) {

	logs, sub, err := _Safe.contract.WatchLogs(opts, "EnabledModule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeEnabledModule)
				if err := _Safe.contract.UnpackLog(event, "EnabledModule", log); err != nil {
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

// SafeExecutionFailedIterator is returned from FilterExecutionFailed and is used to iterate over the raw logs and unpacked data for ExecutionFailed events raised by the Safe contract.
type SafeExecutionFailedIterator struct {
	Event *SafeExecutionFailed // Event containing the contract specifics and raw log

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
func (it *SafeExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeExecutionFailed)
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
		it.Event = new(SafeExecutionFailed)
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
func (it *SafeExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeExecutionFailed represents a ExecutionFailed event raised by the Safe contract.
type SafeExecutionFailed struct {
	TxHash [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailed is a free log retrieval operation binding the contract event 0xabfd711ecdd15ae3a6b3ad16ff2e9d81aec026a39d16725ee164be4fbf857a7c.
//
// Solidity: e ExecutionFailed(txHash bytes32)
func (_Safe *SafeFilterer) FilterExecutionFailed(opts *bind.FilterOpts) (*SafeExecutionFailedIterator, error) {

	logs, sub, err := _Safe.contract.FilterLogs(opts, "ExecutionFailed")
	if err != nil {
		return nil, err
	}
	return &SafeExecutionFailedIterator{contract: _Safe.contract, event: "ExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchExecutionFailed is a free log subscription operation binding the contract event 0xabfd711ecdd15ae3a6b3ad16ff2e9d81aec026a39d16725ee164be4fbf857a7c.
//
// Solidity: e ExecutionFailed(txHash bytes32)
func (_Safe *SafeFilterer) WatchExecutionFailed(opts *bind.WatchOpts, sink chan<- *SafeExecutionFailed) (event.Subscription, error) {

	logs, sub, err := _Safe.contract.WatchLogs(opts, "ExecutionFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeExecutionFailed)
				if err := _Safe.contract.UnpackLog(event, "ExecutionFailed", log); err != nil {
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

// SafeRemovedOwnerIterator is returned from FilterRemovedOwner and is used to iterate over the raw logs and unpacked data for RemovedOwner events raised by the Safe contract.
type SafeRemovedOwnerIterator struct {
	Event *SafeRemovedOwner // Event containing the contract specifics and raw log

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
func (it *SafeRemovedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeRemovedOwner)
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
		it.Event = new(SafeRemovedOwner)
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
func (it *SafeRemovedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeRemovedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeRemovedOwner represents a RemovedOwner event raised by the Safe contract.
type SafeRemovedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemovedOwner is a free log retrieval operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: e RemovedOwner(owner address)
func (_Safe *SafeFilterer) FilterRemovedOwner(opts *bind.FilterOpts) (*SafeRemovedOwnerIterator, error) {

	logs, sub, err := _Safe.contract.FilterLogs(opts, "RemovedOwner")
	if err != nil {
		return nil, err
	}
	return &SafeRemovedOwnerIterator{contract: _Safe.contract, event: "RemovedOwner", logs: logs, sub: sub}, nil
}

// WatchRemovedOwner is a free log subscription operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: e RemovedOwner(owner address)
func (_Safe *SafeFilterer) WatchRemovedOwner(opts *bind.WatchOpts, sink chan<- *SafeRemovedOwner) (event.Subscription, error) {

	logs, sub, err := _Safe.contract.WatchLogs(opts, "RemovedOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeRemovedOwner)
				if err := _Safe.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
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
