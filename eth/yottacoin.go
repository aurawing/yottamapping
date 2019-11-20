// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EthABI is the input ABI used to generate the binding from.
const EthABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tos\",\"type\":\"address[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"multiTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"initPercent\",\"type\":\"uint8\"},{\"name\":\"periods\",\"type\":\"uint256[]\"},{\"name\":\"percents\",\"type\":\"uint8[]\"}],\"name\":\"addRule\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"timeT\",\"type\":\"uint256\"}],\"name\":\"addTimeT\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_freeze\",\"type\":\"bool\"}],\"name\":\"freeze\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"getFrozenAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_targets\",\"type\":\"address[]\"},{\"name\":\"_freezes\",\"type\":\"bool[]\"}],\"name\":\"multiFreeze\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferfix\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"freezeWithTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_targets\",\"type\":\"address[]\"},{\"name\":\"_timestamps\",\"type\":\"uint256[]\"}],\"name\":\"multiFreezeWithTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeRule\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"getFrozenTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\"},{\"name\":\"totalTokenSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// Eth is an auto generated Go binding around an Ethereum contract.
type Eth struct {
	EthCaller     // Read-only binding to the contract
	EthTransactor // Write-only binding to the contract
	EthFilterer   // Log filterer for contract events
}

// EthCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthSession struct {
	Contract     *Eth              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthCallerSession struct {
	Contract *EthCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthTransactorSession struct {
	Contract     *EthTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthRaw struct {
	Contract *Eth // Generic contract binding to access the raw methods on
}

// EthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthCallerRaw struct {
	Contract *EthCaller // Generic read-only contract binding to access the raw methods on
}

// EthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthTransactorRaw struct {
	Contract *EthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEth creates a new instance of Eth, bound to a specific deployed contract.
func NewEth(address common.Address, backend bind.ContractBackend) (*Eth, error) {
	contract, err := bindEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eth{EthCaller: EthCaller{contract: contract}, EthTransactor: EthTransactor{contract: contract}, EthFilterer: EthFilterer{contract: contract}}, nil
}

// NewEthCaller creates a new read-only instance of Eth, bound to a specific deployed contract.
func NewEthCaller(address common.Address, caller bind.ContractCaller) (*EthCaller, error) {
	contract, err := bindEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthCaller{contract: contract}, nil
}

// NewEthTransactor creates a new write-only instance of Eth, bound to a specific deployed contract.
func NewEthTransactor(address common.Address, transactor bind.ContractTransactor) (*EthTransactor, error) {
	contract, err := bindEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthTransactor{contract: contract}, nil
}

// NewEthFilterer creates a new log filterer instance of Eth, bound to a specific deployed contract.
func NewEthFilterer(address common.Address, filterer bind.ContractFilterer) (*EthFilterer, error) {
	contract, err := bindEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthFilterer{contract: contract}, nil
}

// bindEth binds a generic wrapper to an already deployed contract.
func bindEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eth *EthRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Eth.Contract.EthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eth *EthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth.Contract.EthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eth *EthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eth.Contract.EthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eth *EthCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Eth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eth *EthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eth *EthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eth.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_Eth *EthCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Eth.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_Eth *EthSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Eth.Contract.Allowance(&_Eth.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_Eth *EthCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Eth.Contract.Allowance(&_Eth.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Eth *EthCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Eth.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Eth *EthSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Eth.Contract.BalanceOf(&_Eth.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Eth *EthCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Eth.Contract.BalanceOf(&_Eth.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Eth *EthCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Eth.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Eth *EthSession) Decimals() (uint8, error) {
	return _Eth.Contract.Decimals(&_Eth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Eth *EthCallerSession) Decimals() (uint8, error) {
	return _Eth.Contract.Decimals(&_Eth.CallOpts)
}

// GetFrozenAccount is a free data retrieval call binding the contract method 0xc4977807.
//
// Solidity: function getFrozenAccount(address _target) constant returns(bool)
func (_Eth *EthCaller) GetFrozenAccount(opts *bind.CallOpts, _target common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Eth.contract.Call(opts, out, "getFrozenAccount", _target)
	return *ret0, err
}

// GetFrozenAccount is a free data retrieval call binding the contract method 0xc4977807.
//
// Solidity: function getFrozenAccount(address _target) constant returns(bool)
func (_Eth *EthSession) GetFrozenAccount(_target common.Address) (bool, error) {
	return _Eth.Contract.GetFrozenAccount(&_Eth.CallOpts, _target)
}

// GetFrozenAccount is a free data retrieval call binding the contract method 0xc4977807.
//
// Solidity: function getFrozenAccount(address _target) constant returns(bool)
func (_Eth *EthCallerSession) GetFrozenAccount(_target common.Address) (bool, error) {
	return _Eth.Contract.GetFrozenAccount(&_Eth.CallOpts, _target)
}

// GetFrozenTimestamp is a free data retrieval call binding the contract method 0xe6ad5bc7.
//
// Solidity: function getFrozenTimestamp(address _target) constant returns(uint256)
func (_Eth *EthCaller) GetFrozenTimestamp(opts *bind.CallOpts, _target common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Eth.contract.Call(opts, out, "getFrozenTimestamp", _target)
	return *ret0, err
}

// GetFrozenTimestamp is a free data retrieval call binding the contract method 0xe6ad5bc7.
//
// Solidity: function getFrozenTimestamp(address _target) constant returns(uint256)
func (_Eth *EthSession) GetFrozenTimestamp(_target common.Address) (*big.Int, error) {
	return _Eth.Contract.GetFrozenTimestamp(&_Eth.CallOpts, _target)
}

// GetFrozenTimestamp is a free data retrieval call binding the contract method 0xe6ad5bc7.
//
// Solidity: function getFrozenTimestamp(address _target) constant returns(uint256)
func (_Eth *EthCallerSession) GetFrozenTimestamp(_target common.Address) (*big.Int, error) {
	return _Eth.Contract.GetFrozenTimestamp(&_Eth.CallOpts, _target)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Eth *EthCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Eth.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Eth *EthSession) Name() (string, error) {
	return _Eth.Contract.Name(&_Eth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Eth *EthCallerSession) Name() (string, error) {
	return _Eth.Contract.Name(&_Eth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Eth *EthCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Eth.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Eth *EthSession) Symbol() (string, error) {
	return _Eth.Contract.Symbol(&_Eth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Eth *EthCallerSession) Symbol() (string, error) {
	return _Eth.Contract.Symbol(&_Eth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Eth *EthCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Eth.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Eth *EthSession) TotalSupply() (*big.Int, error) {
	return _Eth.Contract.TotalSupply(&_Eth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Eth *EthCallerSession) TotalSupply() (*big.Int, error) {
	return _Eth.Contract.TotalSupply(&_Eth.CallOpts)
}

// AddRule is a paid mutator transaction binding the contract method 0x99f9b55e.
//
// Solidity: function addRule(address addr, uint8 initPercent, uint256[] periods, uint8[] percents) returns(bool)
func (_Eth *EthTransactor) AddRule(opts *bind.TransactOpts, addr common.Address, initPercent uint8, periods []*big.Int, percents []uint8) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "addRule", addr, initPercent, periods, percents)
}

// AddRule is a paid mutator transaction binding the contract method 0x99f9b55e.
//
// Solidity: function addRule(address addr, uint8 initPercent, uint256[] periods, uint8[] percents) returns(bool)
func (_Eth *EthSession) AddRule(addr common.Address, initPercent uint8, periods []*big.Int, percents []uint8) (*types.Transaction, error) {
	return _Eth.Contract.AddRule(&_Eth.TransactOpts, addr, initPercent, periods, percents)
}

// AddRule is a paid mutator transaction binding the contract method 0x99f9b55e.
//
// Solidity: function addRule(address addr, uint8 initPercent, uint256[] periods, uint8[] percents) returns(bool)
func (_Eth *EthTransactorSession) AddRule(addr common.Address, initPercent uint8, periods []*big.Int, percents []uint8) (*types.Transaction, error) {
	return _Eth.Contract.AddRule(&_Eth.TransactOpts, addr, initPercent, periods, percents)
}

// AddTimeT is a paid mutator transaction binding the contract method 0xa2c8a927.
//
// Solidity: function addTimeT(address addr, uint256 timeT) returns(bool)
func (_Eth *EthTransactor) AddTimeT(opts *bind.TransactOpts, addr common.Address, timeT *big.Int) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "addTimeT", addr, timeT)
}

// AddTimeT is a paid mutator transaction binding the contract method 0xa2c8a927.
//
// Solidity: function addTimeT(address addr, uint256 timeT) returns(bool)
func (_Eth *EthSession) AddTimeT(addr common.Address, timeT *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.AddTimeT(&_Eth.TransactOpts, addr, timeT)
}

// AddTimeT is a paid mutator transaction binding the contract method 0xa2c8a927.
//
// Solidity: function addTimeT(address addr, uint256 timeT) returns(bool)
func (_Eth *EthTransactorSession) AddTimeT(addr common.Address, timeT *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.AddTimeT(&_Eth.TransactOpts, addr, timeT)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Eth *EthTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Eth *EthSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.Approve(&_Eth.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Eth *EthTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.Approve(&_Eth.TransactOpts, _spender, _value)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns(bool)
func (_Eth *EthTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns(bool)
func (_Eth *EthSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Eth.Contract.ChangeAdmin(&_Eth.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns(bool)
func (_Eth *EthTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Eth.Contract.ChangeAdmin(&_Eth.TransactOpts, newAdmin)
}

// Freeze is a paid mutator transaction binding the contract method 0xbf120ae5.
//
// Solidity: function freeze(address _target, bool _freeze) returns(bool)
func (_Eth *EthTransactor) Freeze(opts *bind.TransactOpts, _target common.Address, _freeze bool) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "freeze", _target, _freeze)
}

// Freeze is a paid mutator transaction binding the contract method 0xbf120ae5.
//
// Solidity: function freeze(address _target, bool _freeze) returns(bool)
func (_Eth *EthSession) Freeze(_target common.Address, _freeze bool) (*types.Transaction, error) {
	return _Eth.Contract.Freeze(&_Eth.TransactOpts, _target, _freeze)
}

// Freeze is a paid mutator transaction binding the contract method 0xbf120ae5.
//
// Solidity: function freeze(address _target, bool _freeze) returns(bool)
func (_Eth *EthTransactorSession) Freeze(_target common.Address, _freeze bool) (*types.Transaction, error) {
	return _Eth.Contract.Freeze(&_Eth.TransactOpts, _target, _freeze)
}

// FreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd70907b0.
//
// Solidity: function freezeWithTimestamp(address _target, uint256 _timestamp) returns(bool)
func (_Eth *EthTransactor) FreezeWithTimestamp(opts *bind.TransactOpts, _target common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "freezeWithTimestamp", _target, _timestamp)
}

// FreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd70907b0.
//
// Solidity: function freezeWithTimestamp(address _target, uint256 _timestamp) returns(bool)
func (_Eth *EthSession) FreezeWithTimestamp(_target common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.FreezeWithTimestamp(&_Eth.TransactOpts, _target, _timestamp)
}

// FreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd70907b0.
//
// Solidity: function freezeWithTimestamp(address _target, uint256 _timestamp) returns(bool)
func (_Eth *EthTransactorSession) FreezeWithTimestamp(_target common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.FreezeWithTimestamp(&_Eth.TransactOpts, _target, _timestamp)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Eth *EthTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Eth *EthSession) Kill() (*types.Transaction, error) {
	return _Eth.Contract.Kill(&_Eth.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Eth *EthTransactorSession) Kill() (*types.Transaction, error) {
	return _Eth.Contract.Kill(&_Eth.TransactOpts)
}

// MultiFreeze is a paid mutator transaction binding the contract method 0xc878dad9.
//
// Solidity: function multiFreeze(address[] _targets, bool[] _freezes) returns(bool)
func (_Eth *EthTransactor) MultiFreeze(opts *bind.TransactOpts, _targets []common.Address, _freezes []bool) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "multiFreeze", _targets, _freezes)
}

// MultiFreeze is a paid mutator transaction binding the contract method 0xc878dad9.
//
// Solidity: function multiFreeze(address[] _targets, bool[] _freezes) returns(bool)
func (_Eth *EthSession) MultiFreeze(_targets []common.Address, _freezes []bool) (*types.Transaction, error) {
	return _Eth.Contract.MultiFreeze(&_Eth.TransactOpts, _targets, _freezes)
}

// MultiFreeze is a paid mutator transaction binding the contract method 0xc878dad9.
//
// Solidity: function multiFreeze(address[] _targets, bool[] _freezes) returns(bool)
func (_Eth *EthTransactorSession) MultiFreeze(_targets []common.Address, _freezes []bool) (*types.Transaction, error) {
	return _Eth.Contract.MultiFreeze(&_Eth.TransactOpts, _targets, _freezes)
}

// MultiFreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd950c432.
//
// Solidity: function multiFreezeWithTimestamp(address[] _targets, uint256[] _timestamps) returns(bool)
func (_Eth *EthTransactor) MultiFreezeWithTimestamp(opts *bind.TransactOpts, _targets []common.Address, _timestamps []*big.Int) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "multiFreezeWithTimestamp", _targets, _timestamps)
}

// MultiFreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd950c432.
//
// Solidity: function multiFreezeWithTimestamp(address[] _targets, uint256[] _timestamps) returns(bool)
func (_Eth *EthSession) MultiFreezeWithTimestamp(_targets []common.Address, _timestamps []*big.Int) (*types.Transaction, error) {
	return _Eth.Contract.MultiFreezeWithTimestamp(&_Eth.TransactOpts, _targets, _timestamps)
}

// MultiFreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd950c432.
//
// Solidity: function multiFreezeWithTimestamp(address[] _targets, uint256[] _timestamps) returns(bool)
func (_Eth *EthTransactorSession) MultiFreezeWithTimestamp(_targets []common.Address, _timestamps []*big.Int) (*types.Transaction, error) {
	return _Eth.Contract.MultiFreezeWithTimestamp(&_Eth.TransactOpts, _targets, _timestamps)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0x1e89d545.
//
// Solidity: function multiTransfer(address[] _tos, uint256[] _values) returns(bool)
func (_Eth *EthTransactor) MultiTransfer(opts *bind.TransactOpts, _tos []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "multiTransfer", _tos, _values)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0x1e89d545.
//
// Solidity: function multiTransfer(address[] _tos, uint256[] _values) returns(bool)
func (_Eth *EthSession) MultiTransfer(_tos []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Eth.Contract.MultiTransfer(&_Eth.TransactOpts, _tos, _values)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0x1e89d545.
//
// Solidity: function multiTransfer(address[] _tos, uint256[] _values) returns(bool)
func (_Eth *EthTransactorSession) MultiTransfer(_tos []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Eth.Contract.MultiTransfer(&_Eth.TransactOpts, _tos, _values)
}

// RemoveRule is a paid mutator transaction binding the contract method 0xdf21950f.
//
// Solidity: function removeRule(address addr) returns(bool)
func (_Eth *EthTransactor) RemoveRule(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "removeRule", addr)
}

// RemoveRule is a paid mutator transaction binding the contract method 0xdf21950f.
//
// Solidity: function removeRule(address addr) returns(bool)
func (_Eth *EthSession) RemoveRule(addr common.Address) (*types.Transaction, error) {
	return _Eth.Contract.RemoveRule(&_Eth.TransactOpts, addr)
}

// RemoveRule is a paid mutator transaction binding the contract method 0xdf21950f.
//
// Solidity: function removeRule(address addr) returns(bool)
func (_Eth *EthTransactorSession) RemoveRule(addr common.Address) (*types.Transaction, error) {
	return _Eth.Contract.RemoveRule(&_Eth.TransactOpts, addr)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Eth *EthTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Eth *EthSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.Transfer(&_Eth.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Eth *EthTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.Transfer(&_Eth.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Eth *EthTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Eth *EthSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.TransferFrom(&_Eth.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Eth *EthTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.TransferFrom(&_Eth.TransactOpts, _from, _to, _value)
}

// Transferfix is a paid mutator transaction binding the contract method 0xd54c8a56.
//
// Solidity: function transferfix(address _to, uint256 _value) returns()
func (_Eth *EthTransactor) Transferfix(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "transferfix", _to, _value)
}

// Transferfix is a paid mutator transaction binding the contract method 0xd54c8a56.
//
// Solidity: function transferfix(address _to, uint256 _value) returns()
func (_Eth *EthSession) Transferfix(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.Transferfix(&_Eth.TransactOpts, _to, _value)
}

// Transferfix is a paid mutator transaction binding the contract method 0xd54c8a56.
//
// Solidity: function transferfix(address _to, uint256 _value) returns()
func (_Eth *EthTransactorSession) Transferfix(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.Transferfix(&_Eth.TransactOpts, _to, _value)
}

// EthApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Eth contract.
type EthApprovalIterator struct {
	Event *EthApproval // Event containing the contract specifics and raw log

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
func (it *EthApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthApproval)
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
		it.Event = new(EthApproval)
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
func (it *EthApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthApproval represents a Approval event raised by the Eth contract.
type EthApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Eth *EthFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EthApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Eth.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EthApprovalIterator{contract: _Eth.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Eth *EthFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EthApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Eth.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthApproval)
				if err := _Eth.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Eth *EthFilterer) ParseApproval(log types.Log) (*EthApproval, error) {
	event := new(EthApproval)
	if err := _Eth.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Eth contract.
type EthTransferIterator struct {
	Event *EthTransfer // Event containing the contract specifics and raw log

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
func (it *EthTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthTransfer)
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
		it.Event = new(EthTransfer)
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
func (it *EthTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthTransfer represents a Transfer event raised by the Eth contract.
type EthTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Eth *EthFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EthTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Eth.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EthTransferIterator{contract: _Eth.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Eth *EthFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EthTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Eth.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthTransfer)
				if err := _Eth.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Eth *EthFilterer) ParseTransfer(log types.Log) (*EthTransfer, error) {
	event := new(EthTransfer)
	if err := _Eth.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
