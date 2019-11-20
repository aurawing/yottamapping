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
const EthABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAllowMapping\",\"type\":\"bool\"}],\"name\":\"changeAllowMapping\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tos\",\"type\":\"address[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"multiTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"yottaCoin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"initPercent\",\"type\":\"uint8\"},{\"name\":\"periods\",\"type\":\"uint256[]\"},{\"name\":\"percents\",\"type\":\"uint8[]\"}],\"name\":\"addRule\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"timeT\",\"type\":\"uint256\"}],\"name\":\"addTimeT\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allowMapping\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_freeze\",\"type\":\"bool\"}],\"name\":\"freeze\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_targets\",\"type\":\"address[]\"},{\"name\":\"_freezes\",\"type\":\"bool[]\"}],\"name\":\"multiFreeze\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"freezeWithTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_targets\",\"type\":\"address[]\"},{\"name\":\"_timestamps\",\"type\":\"uint256[]\"}],\"name\":\"multiFreezeWithTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeRule\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"yottacoinAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Freezed\",\"type\":\"event\"}]"

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

// AllowMapping is a free data retrieval call binding the contract method 0xabfa90df.
//
// Solidity: function allowMapping() constant returns(bool)
func (_Eth *EthCaller) AllowMapping(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Eth.contract.Call(opts, out, "allowMapping")
	return *ret0, err
}

// AllowMapping is a free data retrieval call binding the contract method 0xabfa90df.
//
// Solidity: function allowMapping() constant returns(bool)
func (_Eth *EthSession) AllowMapping() (bool, error) {
	return _Eth.Contract.AllowMapping(&_Eth.CallOpts)
}

// AllowMapping is a free data retrieval call binding the contract method 0xabfa90df.
//
// Solidity: function allowMapping() constant returns(bool)
func (_Eth *EthCallerSession) AllowMapping() (bool, error) {
	return _Eth.Contract.AllowMapping(&_Eth.CallOpts)
}

// Register is a free data retrieval call binding the contract method 0x4420e486.
//
// Solidity: function register(address ) constant returns(bytes)
func (_Eth *EthCaller) Register(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Eth.contract.Call(opts, out, "register", arg0)
	return *ret0, err
}

// Register is a free data retrieval call binding the contract method 0x4420e486.
//
// Solidity: function register(address ) constant returns(bytes)
func (_Eth *EthSession) Register(arg0 common.Address) ([]byte, error) {
	return _Eth.Contract.Register(&_Eth.CallOpts, arg0)
}

// Register is a free data retrieval call binding the contract method 0x4420e486.
//
// Solidity: function register(address ) constant returns(bytes)
func (_Eth *EthCallerSession) Register(arg0 common.Address) ([]byte, error) {
	return _Eth.Contract.Register(&_Eth.CallOpts, arg0)
}

// YottaCoin is a free data retrieval call binding the contract method 0x7c5727ca.
//
// Solidity: function yottaCoin() constant returns(address)
func (_Eth *EthCaller) YottaCoin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Eth.contract.Call(opts, out, "yottaCoin")
	return *ret0, err
}

// YottaCoin is a free data retrieval call binding the contract method 0x7c5727ca.
//
// Solidity: function yottaCoin() constant returns(address)
func (_Eth *EthSession) YottaCoin() (common.Address, error) {
	return _Eth.Contract.YottaCoin(&_Eth.CallOpts)
}

// YottaCoin is a free data retrieval call binding the contract method 0x7c5727ca.
//
// Solidity: function yottaCoin() constant returns(address)
func (_Eth *EthCallerSession) YottaCoin() (common.Address, error) {
	return _Eth.Contract.YottaCoin(&_Eth.CallOpts)
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

// ChangeAllowMapping is a paid mutator transaction binding the contract method 0x09e2d914.
//
// Solidity: function changeAllowMapping(bool newAllowMapping) returns(bool)
func (_Eth *EthTransactor) ChangeAllowMapping(opts *bind.TransactOpts, newAllowMapping bool) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "changeAllowMapping", newAllowMapping)
}

// ChangeAllowMapping is a paid mutator transaction binding the contract method 0x09e2d914.
//
// Solidity: function changeAllowMapping(bool newAllowMapping) returns(bool)
func (_Eth *EthSession) ChangeAllowMapping(newAllowMapping bool) (*types.Transaction, error) {
	return _Eth.Contract.ChangeAllowMapping(&_Eth.TransactOpts, newAllowMapping)
}

// ChangeAllowMapping is a paid mutator transaction binding the contract method 0x09e2d914.
//
// Solidity: function changeAllowMapping(bool newAllowMapping) returns(bool)
func (_Eth *EthTransactorSession) ChangeAllowMapping(newAllowMapping bool) (*types.Transaction, error) {
	return _Eth.Contract.ChangeAllowMapping(&_Eth.TransactOpts, newAllowMapping)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 newThreshold) returns(bool)
func (_Eth *EthTransactor) ChangeThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "changeThreshold", newThreshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 newThreshold) returns(bool)
func (_Eth *EthSession) ChangeThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.ChangeThreshold(&_Eth.TransactOpts, newThreshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 newThreshold) returns(bool)
func (_Eth *EthTransactorSession) ChangeThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.ChangeThreshold(&_Eth.TransactOpts, newThreshold)
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

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_Eth *EthTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_Eth *EthSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.Withdraw(&_Eth.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_Eth *EthTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Eth.Contract.Withdraw(&_Eth.TransactOpts, _amount)
}

// EthFreezedIterator is returned from FilterFreezed and is used to iterate over the raw logs and unpacked data for Freezed events raised by the Eth contract.
type EthFreezedIterator struct {
	Event *EthFreezed // Event containing the contract specifics and raw log

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
func (it *EthFreezedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthFreezed)
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
		it.Event = new(EthFreezed)
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
func (it *EthFreezedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthFreezedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthFreezed represents a Freezed event raised by the Eth contract.
type EthFreezed struct {
	Sender  common.Address
	Balance *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFreezed is a free log retrieval operation binding the contract event 0x0774af36e8d3fbb76be8b6b764c3e3372138e0b332167b99da8bb7236cd1b3f3.
//
// Solidity: event Freezed(address sender, uint256 balance, bytes data)
func (_Eth *EthFilterer) FilterFreezed(opts *bind.FilterOpts) (*EthFreezedIterator, error) {

	logs, sub, err := _Eth.contract.FilterLogs(opts, "Freezed")
	if err != nil {
		return nil, err
	}
	return &EthFreezedIterator{contract: _Eth.contract, event: "Freezed", logs: logs, sub: sub}, nil
}

// WatchFreezed is a free log subscription operation binding the contract event 0x0774af36e8d3fbb76be8b6b764c3e3372138e0b332167b99da8bb7236cd1b3f3.
//
// Solidity: event Freezed(address sender, uint256 balance, bytes data)
func (_Eth *EthFilterer) WatchFreezed(opts *bind.WatchOpts, sink chan<- *EthFreezed) (event.Subscription, error) {

	logs, sub, err := _Eth.contract.WatchLogs(opts, "Freezed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthFreezed)
				if err := _Eth.contract.UnpackLog(event, "Freezed", log); err != nil {
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

// ParseFreezed is a log parse operation binding the contract event 0x0774af36e8d3fbb76be8b6b764c3e3372138e0b332167b99da8bb7236cd1b3f3.
//
// Solidity: event Freezed(address sender, uint256 balance, bytes data)
func (_Eth *EthFilterer) ParseFreezed(log types.Log) (*EthFreezed, error) {
	event := new(EthFreezed)
	if err := _Eth.contract.UnpackLog(event, "Freezed", log); err != nil {
		return nil, err
	}
	return event, nil
}
