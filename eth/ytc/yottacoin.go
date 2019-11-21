// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ytc

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

// YtcABI is the input ABI used to generate the binding from.
const YtcABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tos\",\"type\":\"address[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"multiTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"initPercent\",\"type\":\"uint8\"},{\"name\":\"periods\",\"type\":\"uint256[]\"},{\"name\":\"percents\",\"type\":\"uint8[]\"}],\"name\":\"addRule\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"timeT\",\"type\":\"uint256\"}],\"name\":\"addTimeT\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_freeze\",\"type\":\"bool\"}],\"name\":\"freeze\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"getFrozenAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_targets\",\"type\":\"address[]\"},{\"name\":\"_freezes\",\"type\":\"bool[]\"}],\"name\":\"multiFreeze\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferfix\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"freezeWithTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_targets\",\"type\":\"address[]\"},{\"name\":\"_timestamps\",\"type\":\"uint256[]\"}],\"name\":\"multiFreezeWithTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeRule\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"getFrozenTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\"},{\"name\":\"totalTokenSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// Ytc is an auto generated Go binding around an Ethereum contract.
type Ytc struct {
	YtcCaller     // Read-only binding to the contract
	YtcTransactor // Write-only binding to the contract
	YtcFilterer   // Log filterer for contract events
}

// YtcCaller is an auto generated read-only Go binding around an Ethereum contract.
type YtcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YtcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YtcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YtcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YtcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YtcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YtcSession struct {
	Contract     *Ytc              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YtcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YtcCallerSession struct {
	Contract *YtcCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// YtcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YtcTransactorSession struct {
	Contract     *YtcTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YtcRaw is an auto generated low-level Go binding around an Ethereum contract.
type YtcRaw struct {
	Contract *Ytc // Generic contract binding to access the raw methods on
}

// YtcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YtcCallerRaw struct {
	Contract *YtcCaller // Generic read-only contract binding to access the raw methods on
}

// YtcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YtcTransactorRaw struct {
	Contract *YtcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYtc creates a new instance of Ytc, bound to a specific deployed contract.
func NewYtc(address common.Address, backend bind.ContractBackend) (*Ytc, error) {
	contract, err := bindYtc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ytc{YtcCaller: YtcCaller{contract: contract}, YtcTransactor: YtcTransactor{contract: contract}, YtcFilterer: YtcFilterer{contract: contract}}, nil
}

// NewYtcCaller creates a new read-only instance of Ytc, bound to a specific deployed contract.
func NewYtcCaller(address common.Address, caller bind.ContractCaller) (*YtcCaller, error) {
	contract, err := bindYtc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YtcCaller{contract: contract}, nil
}

// NewYtcTransactor creates a new write-only instance of Ytc, bound to a specific deployed contract.
func NewYtcTransactor(address common.Address, transactor bind.ContractTransactor) (*YtcTransactor, error) {
	contract, err := bindYtc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YtcTransactor{contract: contract}, nil
}

// NewYtcFilterer creates a new log filterer instance of Ytc, bound to a specific deployed contract.
func NewYtcFilterer(address common.Address, filterer bind.ContractFilterer) (*YtcFilterer, error) {
	contract, err := bindYtc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YtcFilterer{contract: contract}, nil
}

// bindYtc binds a generic wrapper to an already deployed contract.
func bindYtc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YtcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ytc *YtcRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ytc.Contract.YtcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ytc *YtcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ytc.Contract.YtcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ytc *YtcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ytc.Contract.YtcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ytc *YtcCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ytc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ytc *YtcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ytc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ytc *YtcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ytc.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_Ytc *YtcCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ytc.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_Ytc *YtcSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Ytc.Contract.Allowance(&_Ytc.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_Ytc *YtcCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Ytc.Contract.Allowance(&_Ytc.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Ytc *YtcCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ytc.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Ytc *YtcSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Ytc.Contract.BalanceOf(&_Ytc.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Ytc *YtcCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Ytc.Contract.BalanceOf(&_Ytc.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Ytc *YtcCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Ytc.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Ytc *YtcSession) Decimals() (uint8, error) {
	return _Ytc.Contract.Decimals(&_Ytc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Ytc *YtcCallerSession) Decimals() (uint8, error) {
	return _Ytc.Contract.Decimals(&_Ytc.CallOpts)
}

// GetFrozenAccount is a free data retrieval call binding the contract method 0xc4977807.
//
// Solidity: function getFrozenAccount(address _target) constant returns(bool)
func (_Ytc *YtcCaller) GetFrozenAccount(opts *bind.CallOpts, _target common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ytc.contract.Call(opts, out, "getFrozenAccount", _target)
	return *ret0, err
}

// GetFrozenAccount is a free data retrieval call binding the contract method 0xc4977807.
//
// Solidity: function getFrozenAccount(address _target) constant returns(bool)
func (_Ytc *YtcSession) GetFrozenAccount(_target common.Address) (bool, error) {
	return _Ytc.Contract.GetFrozenAccount(&_Ytc.CallOpts, _target)
}

// GetFrozenAccount is a free data retrieval call binding the contract method 0xc4977807.
//
// Solidity: function getFrozenAccount(address _target) constant returns(bool)
func (_Ytc *YtcCallerSession) GetFrozenAccount(_target common.Address) (bool, error) {
	return _Ytc.Contract.GetFrozenAccount(&_Ytc.CallOpts, _target)
}

// GetFrozenTimestamp is a free data retrieval call binding the contract method 0xe6ad5bc7.
//
// Solidity: function getFrozenTimestamp(address _target) constant returns(uint256)
func (_Ytc *YtcCaller) GetFrozenTimestamp(opts *bind.CallOpts, _target common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ytc.contract.Call(opts, out, "getFrozenTimestamp", _target)
	return *ret0, err
}

// GetFrozenTimestamp is a free data retrieval call binding the contract method 0xe6ad5bc7.
//
// Solidity: function getFrozenTimestamp(address _target) constant returns(uint256)
func (_Ytc *YtcSession) GetFrozenTimestamp(_target common.Address) (*big.Int, error) {
	return _Ytc.Contract.GetFrozenTimestamp(&_Ytc.CallOpts, _target)
}

// GetFrozenTimestamp is a free data retrieval call binding the contract method 0xe6ad5bc7.
//
// Solidity: function getFrozenTimestamp(address _target) constant returns(uint256)
func (_Ytc *YtcCallerSession) GetFrozenTimestamp(_target common.Address) (*big.Int, error) {
	return _Ytc.Contract.GetFrozenTimestamp(&_Ytc.CallOpts, _target)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Ytc *YtcCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ytc.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Ytc *YtcSession) Name() (string, error) {
	return _Ytc.Contract.Name(&_Ytc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Ytc *YtcCallerSession) Name() (string, error) {
	return _Ytc.Contract.Name(&_Ytc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Ytc *YtcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ytc.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Ytc *YtcSession) Symbol() (string, error) {
	return _Ytc.Contract.Symbol(&_Ytc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Ytc *YtcCallerSession) Symbol() (string, error) {
	return _Ytc.Contract.Symbol(&_Ytc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Ytc *YtcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ytc.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Ytc *YtcSession) TotalSupply() (*big.Int, error) {
	return _Ytc.Contract.TotalSupply(&_Ytc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Ytc *YtcCallerSession) TotalSupply() (*big.Int, error) {
	return _Ytc.Contract.TotalSupply(&_Ytc.CallOpts)
}

// AddRule is a paid mutator transaction binding the contract method 0x99f9b55e.
//
// Solidity: function addRule(address addr, uint8 initPercent, uint256[] periods, uint8[] percents) returns(bool)
func (_Ytc *YtcTransactor) AddRule(opts *bind.TransactOpts, addr common.Address, initPercent uint8, periods []*big.Int, percents []uint8) (*types.Transaction, error) {
	return _Ytc.contract.Transact(opts, "addRule", addr, initPercent, periods, percents)
}

// AddRule is a paid mutator transaction binding the contract method 0x99f9b55e.
//
// Solidity: function addRule(address addr, uint8 initPercent, uint256[] periods, uint8[] percents) returns(bool)
func (_Ytc *YtcSession) AddRule(addr common.Address, initPercent uint8, periods []*big.Int, percents []uint8) (*types.Transaction, error) {
	return _Ytc.Contract.AddRule(&_Ytc.TransactOpts, addr, initPercent, periods, percents)
}

// AddRule is a paid mutator transaction binding the contract method 0x99f9b55e.
//
// Solidity: function addRule(address addr, uint8 initPercent, uint256[] periods, uint8[] percents) returns(bool)
func (_Ytc *YtcTransactorSession) AddRule(addr common.Address, initPercent uint8, periods []*big.Int, percents []uint8) (*types.Transaction, error) {
	return _Ytc.Contract.AddRule(&_Ytc.TransactOpts, addr, initPercent, periods, percents)
}

// AddTimeT is a paid mutator transaction binding the contract method 0xa2c8a927.
//
// Solidity: function addTimeT(address addr, uint256 timeT) returns(bool)
func (_Ytc *YtcTransactor) AddTimeT(opts *bind.TransactOpts, addr common.Address, timeT *big.Int) (*types.Transaction, error) {
	return _Ytc.contract.Transact(opts, "addTimeT", addr, timeT)
}

// AddTimeT is a paid mutator transaction binding the contract method 0xa2c8a927.
//
// Solidity: function addTimeT(address addr, uint256 timeT) returns(bool)
func (_Ytc *YtcSession) AddTimeT(addr common.Address, timeT *big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.AddTimeT(&_Ytc.TransactOpts, addr, timeT)
}

// AddTimeT is a paid mutator transaction binding the contract method 0xa2c8a927.
//
// Solidity: function addTimeT(address addr, uint256 timeT) returns(bool)
func (_Ytc *YtcTransactorSession) AddTimeT(addr common.Address, timeT *big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.AddTimeT(&_Ytc.TransactOpts, addr, timeT)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Ytc *YtcTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytc.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Ytc *YtcSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.Approve(&_Ytc.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Ytc *YtcTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.Approve(&_Ytc.TransactOpts, _spender, _value)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns(bool)
func (_Ytc *YtcTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Ytc.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns(bool)
func (_Ytc *YtcSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Ytc.Contract.ChangeAdmin(&_Ytc.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns(bool)
func (_Ytc *YtcTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Ytc.Contract.ChangeAdmin(&_Ytc.TransactOpts, newAdmin)
}

// Freeze is a paid mutator transaction binding the contract method 0xbf120ae5.
//
// Solidity: function freeze(address _target, bool _freeze) returns(bool)
func (_Ytc *YtcTransactor) Freeze(opts *bind.TransactOpts, _target common.Address, _freeze bool) (*types.Transaction, error) {
	return _Ytc.contract.Transact(opts, "freeze", _target, _freeze)
}

// Freeze is a paid mutator transaction binding the contract method 0xbf120ae5.
//
// Solidity: function freeze(address _target, bool _freeze) returns(bool)
func (_Ytc *YtcSession) Freeze(_target common.Address, _freeze bool) (*types.Transaction, error) {
	return _Ytc.Contract.Freeze(&_Ytc.TransactOpts, _target, _freeze)
}

// Freeze is a paid mutator transaction binding the contract method 0xbf120ae5.
//
// Solidity: function freeze(address _target, bool _freeze) returns(bool)
func (_Ytc *YtcTransactorSession) Freeze(_target common.Address, _freeze bool) (*types.Transaction, error) {
	return _Ytc.Contract.Freeze(&_Ytc.TransactOpts, _target, _freeze)
}

// FreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd70907b0.
//
// Solidity: function freezeWithTimestamp(address _target, uint256 _timestamp) returns(bool)
func (_Ytc *YtcTransactor) FreezeWithTimestamp(opts *bind.TransactOpts, _target common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _Ytc.contract.Transact(opts, "freezeWithTimestamp", _target, _timestamp)
}

// FreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd70907b0.
//
// Solidity: function freezeWithTimestamp(address _target, uint256 _timestamp) returns(bool)
func (_Ytc *YtcSession) FreezeWithTimestamp(_target common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.FreezeWithTimestamp(&_Ytc.TransactOpts, _target, _timestamp)
}

// FreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd70907b0.
//
// Solidity: function freezeWithTimestamp(address _target, uint256 _timestamp) returns(bool)
func (_Ytc *YtcTransactorSession) FreezeWithTimestamp(_target common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.FreezeWithTimestamp(&_Ytc.TransactOpts, _target, _timestamp)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ytc *YtcTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ytc.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ytc *YtcSession) Kill() (*types.Transaction, error) {
	return _Ytc.Contract.Kill(&_Ytc.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ytc *YtcTransactorSession) Kill() (*types.Transaction, error) {
	return _Ytc.Contract.Kill(&_Ytc.TransactOpts)
}

// MultiFreeze is a paid mutator transaction binding the contract method 0xc878dad9.
//
// Solidity: function multiFreeze(address[] _targets, bool[] _freezes) returns(bool)
func (_Ytc *YtcTransactor) MultiFreeze(opts *bind.TransactOpts, _targets []common.Address, _freezes []bool) (*types.Transaction, error) {
	return _Ytc.contract.Transact(opts, "multiFreeze", _targets, _freezes)
}

// MultiFreeze is a paid mutator transaction binding the contract method 0xc878dad9.
//
// Solidity: function multiFreeze(address[] _targets, bool[] _freezes) returns(bool)
func (_Ytc *YtcSession) MultiFreeze(_targets []common.Address, _freezes []bool) (*types.Transaction, error) {
	return _Ytc.Contract.MultiFreeze(&_Ytc.TransactOpts, _targets, _freezes)
}

// MultiFreeze is a paid mutator transaction binding the contract method 0xc878dad9.
//
// Solidity: function multiFreeze(address[] _targets, bool[] _freezes) returns(bool)
func (_Ytc *YtcTransactorSession) MultiFreeze(_targets []common.Address, _freezes []bool) (*types.Transaction, error) {
	return _Ytc.Contract.MultiFreeze(&_Ytc.TransactOpts, _targets, _freezes)
}

// MultiFreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd950c432.
//
// Solidity: function multiFreezeWithTimestamp(address[] _targets, uint256[] _timestamps) returns(bool)
func (_Ytc *YtcTransactor) MultiFreezeWithTimestamp(opts *bind.TransactOpts, _targets []common.Address, _timestamps []*big.Int) (*types.Transaction, error) {
	return _Ytc.contract.Transact(opts, "multiFreezeWithTimestamp", _targets, _timestamps)
}

// MultiFreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd950c432.
//
// Solidity: function multiFreezeWithTimestamp(address[] _targets, uint256[] _timestamps) returns(bool)
func (_Ytc *YtcSession) MultiFreezeWithTimestamp(_targets []common.Address, _timestamps []*big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.MultiFreezeWithTimestamp(&_Ytc.TransactOpts, _targets, _timestamps)
}

// MultiFreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd950c432.
//
// Solidity: function multiFreezeWithTimestamp(address[] _targets, uint256[] _timestamps) returns(bool)
func (_Ytc *YtcTransactorSession) MultiFreezeWithTimestamp(_targets []common.Address, _timestamps []*big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.MultiFreezeWithTimestamp(&_Ytc.TransactOpts, _targets, _timestamps)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0x1e89d545.
//
// Solidity: function multiTransfer(address[] _tos, uint256[] _values) returns(bool)
func (_Ytc *YtcTransactor) MultiTransfer(opts *bind.TransactOpts, _tos []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Ytc.contract.Transact(opts, "multiTransfer", _tos, _values)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0x1e89d545.
//
// Solidity: function multiTransfer(address[] _tos, uint256[] _values) returns(bool)
func (_Ytc *YtcSession) MultiTransfer(_tos []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.MultiTransfer(&_Ytc.TransactOpts, _tos, _values)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0x1e89d545.
//
// Solidity: function multiTransfer(address[] _tos, uint256[] _values) returns(bool)
func (_Ytc *YtcTransactorSession) MultiTransfer(_tos []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.MultiTransfer(&_Ytc.TransactOpts, _tos, _values)
}

// RemoveRule is a paid mutator transaction binding the contract method 0xdf21950f.
//
// Solidity: function removeRule(address addr) returns(bool)
func (_Ytc *YtcTransactor) RemoveRule(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Ytc.contract.Transact(opts, "removeRule", addr)
}

// RemoveRule is a paid mutator transaction binding the contract method 0xdf21950f.
//
// Solidity: function removeRule(address addr) returns(bool)
func (_Ytc *YtcSession) RemoveRule(addr common.Address) (*types.Transaction, error) {
	return _Ytc.Contract.RemoveRule(&_Ytc.TransactOpts, addr)
}

// RemoveRule is a paid mutator transaction binding the contract method 0xdf21950f.
//
// Solidity: function removeRule(address addr) returns(bool)
func (_Ytc *YtcTransactorSession) RemoveRule(addr common.Address) (*types.Transaction, error) {
	return _Ytc.Contract.RemoveRule(&_Ytc.TransactOpts, addr)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Ytc *YtcTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytc.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Ytc *YtcSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.Transfer(&_Ytc.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Ytc *YtcTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.Transfer(&_Ytc.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Ytc *YtcTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytc.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Ytc *YtcSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.TransferFrom(&_Ytc.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Ytc *YtcTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.TransferFrom(&_Ytc.TransactOpts, _from, _to, _value)
}

// Transferfix is a paid mutator transaction binding the contract method 0xd54c8a56.
//
// Solidity: function transferfix(address _to, uint256 _value) returns()
func (_Ytc *YtcTransactor) Transferfix(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytc.contract.Transact(opts, "transferfix", _to, _value)
}

// Transferfix is a paid mutator transaction binding the contract method 0xd54c8a56.
//
// Solidity: function transferfix(address _to, uint256 _value) returns()
func (_Ytc *YtcSession) Transferfix(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.Transferfix(&_Ytc.TransactOpts, _to, _value)
}

// Transferfix is a paid mutator transaction binding the contract method 0xd54c8a56.
//
// Solidity: function transferfix(address _to, uint256 _value) returns()
func (_Ytc *YtcTransactorSession) Transferfix(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytc.Contract.Transferfix(&_Ytc.TransactOpts, _to, _value)
}

// YtcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ytc contract.
type YtcApprovalIterator struct {
	Event *YtcApproval // Event containing the contract specifics and raw log

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
func (it *YtcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YtcApproval)
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
		it.Event = new(YtcApproval)
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
func (it *YtcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YtcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YtcApproval represents a Approval event raised by the Ytc contract.
type YtcApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ytc *YtcFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*YtcApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ytc.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &YtcApprovalIterator{contract: _Ytc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ytc *YtcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *YtcApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ytc.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YtcApproval)
				if err := _Ytc.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Ytc *YtcFilterer) ParseApproval(log types.Log) (*YtcApproval, error) {
	event := new(YtcApproval)
	if err := _Ytc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// YtcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ytc contract.
type YtcTransferIterator struct {
	Event *YtcTransfer // Event containing the contract specifics and raw log

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
func (it *YtcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YtcTransfer)
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
		it.Event = new(YtcTransfer)
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
func (it *YtcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YtcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YtcTransfer represents a Transfer event raised by the Ytc contract.
type YtcTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ytc *YtcFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*YtcTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ytc.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YtcTransferIterator{contract: _Ytc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ytc *YtcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *YtcTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ytc.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YtcTransfer)
				if err := _Ytc.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Ytc *YtcFilterer) ParseTransfer(log types.Log) (*YtcTransfer, error) {
	event := new(YtcTransfer)
	if err := _Ytc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
