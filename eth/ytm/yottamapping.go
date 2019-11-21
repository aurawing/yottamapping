// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ytm

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

// YtmABI is the input ABI used to generate the binding from.
const YtmABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAllowMapping\",\"type\":\"bool\"}],\"name\":\"changeAllowMapping\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tos\",\"type\":\"address[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"multiTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"yottaCoin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"initPercent\",\"type\":\"uint8\"},{\"name\":\"periods\",\"type\":\"uint256[]\"},{\"name\":\"percents\",\"type\":\"uint8[]\"}],\"name\":\"addRule\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"timeT\",\"type\":\"uint256\"}],\"name\":\"addTimeT\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allowMapping\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_freeze\",\"type\":\"bool\"}],\"name\":\"freeze\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_targets\",\"type\":\"address[]\"},{\"name\":\"_freezes\",\"type\":\"bool[]\"}],\"name\":\"multiFreeze\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"freezeWithTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_targets\",\"type\":\"address[]\"},{\"name\":\"_timestamps\",\"type\":\"uint256[]\"}],\"name\":\"multiFreezeWithTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeRule\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"yottacoinAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Freezed\",\"type\":\"event\"}]"

// Ytm is an auto generated Go binding around an Ethereum contract.
type Ytm struct {
	YtmCaller     // Read-only binding to the contract
	YtmTransactor // Write-only binding to the contract
	YtmFilterer   // Log filterer for contract events
}

// YtmCaller is an auto generated read-only Go binding around an Ethereum contract.
type YtmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YtmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YtmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YtmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YtmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YtmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YtmSession struct {
	Contract     *Ytm              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YtmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YtmCallerSession struct {
	Contract *YtmCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// YtmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YtmTransactorSession struct {
	Contract     *YtmTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YtmRaw is an auto generated low-level Go binding around an Ethereum contract.
type YtmRaw struct {
	Contract *Ytm // Generic contract binding to access the raw methods on
}

// YtmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YtmCallerRaw struct {
	Contract *YtmCaller // Generic read-only contract binding to access the raw methods on
}

// YtmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YtmTransactorRaw struct {
	Contract *YtmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYtm creates a new instance of Ytm, bound to a specific deployed contract.
func NewYtm(address common.Address, backend bind.ContractBackend) (*Ytm, error) {
	contract, err := bindYtm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ytm{YtmCaller: YtmCaller{contract: contract}, YtmTransactor: YtmTransactor{contract: contract}, YtmFilterer: YtmFilterer{contract: contract}}, nil
}

// NewYtmCaller creates a new read-only instance of Ytm, bound to a specific deployed contract.
func NewYtmCaller(address common.Address, caller bind.ContractCaller) (*YtmCaller, error) {
	contract, err := bindYtm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YtmCaller{contract: contract}, nil
}

// NewYtmTransactor creates a new write-only instance of Ytm, bound to a specific deployed contract.
func NewYtmTransactor(address common.Address, transactor bind.ContractTransactor) (*YtmTransactor, error) {
	contract, err := bindYtm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YtmTransactor{contract: contract}, nil
}

// NewYtmFilterer creates a new log filterer instance of Ytm, bound to a specific deployed contract.
func NewYtmFilterer(address common.Address, filterer bind.ContractFilterer) (*YtmFilterer, error) {
	contract, err := bindYtm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YtmFilterer{contract: contract}, nil
}

// bindYtm binds a generic wrapper to an already deployed contract.
func bindYtm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YtmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ytm *YtmRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ytm.Contract.YtmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ytm *YtmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ytm.Contract.YtmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ytm *YtmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ytm.Contract.YtmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ytm *YtmCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ytm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ytm *YtmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ytm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ytm *YtmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ytm.Contract.contract.Transact(opts, method, params...)
}

// AllowMapping is a free data retrieval call binding the contract method 0xabfa90df.
//
// Solidity: function allowMapping() constant returns(bool)
func (_Ytm *YtmCaller) AllowMapping(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ytm.contract.Call(opts, out, "allowMapping")
	return *ret0, err
}

// AllowMapping is a free data retrieval call binding the contract method 0xabfa90df.
//
// Solidity: function allowMapping() constant returns(bool)
func (_Ytm *YtmSession) AllowMapping() (bool, error) {
	return _Ytm.Contract.AllowMapping(&_Ytm.CallOpts)
}

// AllowMapping is a free data retrieval call binding the contract method 0xabfa90df.
//
// Solidity: function allowMapping() constant returns(bool)
func (_Ytm *YtmCallerSession) AllowMapping() (bool, error) {
	return _Ytm.Contract.AllowMapping(&_Ytm.CallOpts)
}

// Register is a free data retrieval call binding the contract method 0x4420e486.
//
// Solidity: function register(address ) constant returns(bytes)
func (_Ytm *YtmCaller) Register(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Ytm.contract.Call(opts, out, "register", arg0)
	return *ret0, err
}

// Register is a free data retrieval call binding the contract method 0x4420e486.
//
// Solidity: function register(address ) constant returns(bytes)
func (_Ytm *YtmSession) Register(arg0 common.Address) ([]byte, error) {
	return _Ytm.Contract.Register(&_Ytm.CallOpts, arg0)
}

// Register is a free data retrieval call binding the contract method 0x4420e486.
//
// Solidity: function register(address ) constant returns(bytes)
func (_Ytm *YtmCallerSession) Register(arg0 common.Address) ([]byte, error) {
	return _Ytm.Contract.Register(&_Ytm.CallOpts, arg0)
}

// YottaCoin is a free data retrieval call binding the contract method 0x7c5727ca.
//
// Solidity: function yottaCoin() constant returns(address)
func (_Ytm *YtmCaller) YottaCoin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ytm.contract.Call(opts, out, "yottaCoin")
	return *ret0, err
}

// YottaCoin is a free data retrieval call binding the contract method 0x7c5727ca.
//
// Solidity: function yottaCoin() constant returns(address)
func (_Ytm *YtmSession) YottaCoin() (common.Address, error) {
	return _Ytm.Contract.YottaCoin(&_Ytm.CallOpts)
}

// YottaCoin is a free data retrieval call binding the contract method 0x7c5727ca.
//
// Solidity: function yottaCoin() constant returns(address)
func (_Ytm *YtmCallerSession) YottaCoin() (common.Address, error) {
	return _Ytm.Contract.YottaCoin(&_Ytm.CallOpts)
}

// AddRule is a paid mutator transaction binding the contract method 0x99f9b55e.
//
// Solidity: function addRule(address addr, uint8 initPercent, uint256[] periods, uint8[] percents) returns(bool)
func (_Ytm *YtmTransactor) AddRule(opts *bind.TransactOpts, addr common.Address, initPercent uint8, periods []*big.Int, percents []uint8) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "addRule", addr, initPercent, periods, percents)
}

// AddRule is a paid mutator transaction binding the contract method 0x99f9b55e.
//
// Solidity: function addRule(address addr, uint8 initPercent, uint256[] periods, uint8[] percents) returns(bool)
func (_Ytm *YtmSession) AddRule(addr common.Address, initPercent uint8, periods []*big.Int, percents []uint8) (*types.Transaction, error) {
	return _Ytm.Contract.AddRule(&_Ytm.TransactOpts, addr, initPercent, periods, percents)
}

// AddRule is a paid mutator transaction binding the contract method 0x99f9b55e.
//
// Solidity: function addRule(address addr, uint8 initPercent, uint256[] periods, uint8[] percents) returns(bool)
func (_Ytm *YtmTransactorSession) AddRule(addr common.Address, initPercent uint8, periods []*big.Int, percents []uint8) (*types.Transaction, error) {
	return _Ytm.Contract.AddRule(&_Ytm.TransactOpts, addr, initPercent, periods, percents)
}

// AddTimeT is a paid mutator transaction binding the contract method 0xa2c8a927.
//
// Solidity: function addTimeT(address addr, uint256 timeT) returns(bool)
func (_Ytm *YtmTransactor) AddTimeT(opts *bind.TransactOpts, addr common.Address, timeT *big.Int) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "addTimeT", addr, timeT)
}

// AddTimeT is a paid mutator transaction binding the contract method 0xa2c8a927.
//
// Solidity: function addTimeT(address addr, uint256 timeT) returns(bool)
func (_Ytm *YtmSession) AddTimeT(addr common.Address, timeT *big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.AddTimeT(&_Ytm.TransactOpts, addr, timeT)
}

// AddTimeT is a paid mutator transaction binding the contract method 0xa2c8a927.
//
// Solidity: function addTimeT(address addr, uint256 timeT) returns(bool)
func (_Ytm *YtmTransactorSession) AddTimeT(addr common.Address, timeT *big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.AddTimeT(&_Ytm.TransactOpts, addr, timeT)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Ytm *YtmTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Ytm *YtmSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.Approve(&_Ytm.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Ytm *YtmTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.Approve(&_Ytm.TransactOpts, _spender, _value)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns(bool)
func (_Ytm *YtmTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns(bool)
func (_Ytm *YtmSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Ytm.Contract.ChangeAdmin(&_Ytm.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns(bool)
func (_Ytm *YtmTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Ytm.Contract.ChangeAdmin(&_Ytm.TransactOpts, newAdmin)
}

// ChangeAllowMapping is a paid mutator transaction binding the contract method 0x09e2d914.
//
// Solidity: function changeAllowMapping(bool newAllowMapping) returns(bool)
func (_Ytm *YtmTransactor) ChangeAllowMapping(opts *bind.TransactOpts, newAllowMapping bool) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "changeAllowMapping", newAllowMapping)
}

// ChangeAllowMapping is a paid mutator transaction binding the contract method 0x09e2d914.
//
// Solidity: function changeAllowMapping(bool newAllowMapping) returns(bool)
func (_Ytm *YtmSession) ChangeAllowMapping(newAllowMapping bool) (*types.Transaction, error) {
	return _Ytm.Contract.ChangeAllowMapping(&_Ytm.TransactOpts, newAllowMapping)
}

// ChangeAllowMapping is a paid mutator transaction binding the contract method 0x09e2d914.
//
// Solidity: function changeAllowMapping(bool newAllowMapping) returns(bool)
func (_Ytm *YtmTransactorSession) ChangeAllowMapping(newAllowMapping bool) (*types.Transaction, error) {
	return _Ytm.Contract.ChangeAllowMapping(&_Ytm.TransactOpts, newAllowMapping)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 newThreshold) returns(bool)
func (_Ytm *YtmTransactor) ChangeThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "changeThreshold", newThreshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 newThreshold) returns(bool)
func (_Ytm *YtmSession) ChangeThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.ChangeThreshold(&_Ytm.TransactOpts, newThreshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 newThreshold) returns(bool)
func (_Ytm *YtmTransactorSession) ChangeThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.ChangeThreshold(&_Ytm.TransactOpts, newThreshold)
}

// Freeze is a paid mutator transaction binding the contract method 0xbf120ae5.
//
// Solidity: function freeze(address _target, bool _freeze) returns(bool)
func (_Ytm *YtmTransactor) Freeze(opts *bind.TransactOpts, _target common.Address, _freeze bool) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "freeze", _target, _freeze)
}

// Freeze is a paid mutator transaction binding the contract method 0xbf120ae5.
//
// Solidity: function freeze(address _target, bool _freeze) returns(bool)
func (_Ytm *YtmSession) Freeze(_target common.Address, _freeze bool) (*types.Transaction, error) {
	return _Ytm.Contract.Freeze(&_Ytm.TransactOpts, _target, _freeze)
}

// Freeze is a paid mutator transaction binding the contract method 0xbf120ae5.
//
// Solidity: function freeze(address _target, bool _freeze) returns(bool)
func (_Ytm *YtmTransactorSession) Freeze(_target common.Address, _freeze bool) (*types.Transaction, error) {
	return _Ytm.Contract.Freeze(&_Ytm.TransactOpts, _target, _freeze)
}

// FreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd70907b0.
//
// Solidity: function freezeWithTimestamp(address _target, uint256 _timestamp) returns(bool)
func (_Ytm *YtmTransactor) FreezeWithTimestamp(opts *bind.TransactOpts, _target common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "freezeWithTimestamp", _target, _timestamp)
}

// FreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd70907b0.
//
// Solidity: function freezeWithTimestamp(address _target, uint256 _timestamp) returns(bool)
func (_Ytm *YtmSession) FreezeWithTimestamp(_target common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.FreezeWithTimestamp(&_Ytm.TransactOpts, _target, _timestamp)
}

// FreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd70907b0.
//
// Solidity: function freezeWithTimestamp(address _target, uint256 _timestamp) returns(bool)
func (_Ytm *YtmTransactorSession) FreezeWithTimestamp(_target common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.FreezeWithTimestamp(&_Ytm.TransactOpts, _target, _timestamp)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ytm *YtmTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ytm *YtmSession) Kill() (*types.Transaction, error) {
	return _Ytm.Contract.Kill(&_Ytm.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ytm *YtmTransactorSession) Kill() (*types.Transaction, error) {
	return _Ytm.Contract.Kill(&_Ytm.TransactOpts)
}

// MultiFreeze is a paid mutator transaction binding the contract method 0xc878dad9.
//
// Solidity: function multiFreeze(address[] _targets, bool[] _freezes) returns(bool)
func (_Ytm *YtmTransactor) MultiFreeze(opts *bind.TransactOpts, _targets []common.Address, _freezes []bool) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "multiFreeze", _targets, _freezes)
}

// MultiFreeze is a paid mutator transaction binding the contract method 0xc878dad9.
//
// Solidity: function multiFreeze(address[] _targets, bool[] _freezes) returns(bool)
func (_Ytm *YtmSession) MultiFreeze(_targets []common.Address, _freezes []bool) (*types.Transaction, error) {
	return _Ytm.Contract.MultiFreeze(&_Ytm.TransactOpts, _targets, _freezes)
}

// MultiFreeze is a paid mutator transaction binding the contract method 0xc878dad9.
//
// Solidity: function multiFreeze(address[] _targets, bool[] _freezes) returns(bool)
func (_Ytm *YtmTransactorSession) MultiFreeze(_targets []common.Address, _freezes []bool) (*types.Transaction, error) {
	return _Ytm.Contract.MultiFreeze(&_Ytm.TransactOpts, _targets, _freezes)
}

// MultiFreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd950c432.
//
// Solidity: function multiFreezeWithTimestamp(address[] _targets, uint256[] _timestamps) returns(bool)
func (_Ytm *YtmTransactor) MultiFreezeWithTimestamp(opts *bind.TransactOpts, _targets []common.Address, _timestamps []*big.Int) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "multiFreezeWithTimestamp", _targets, _timestamps)
}

// MultiFreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd950c432.
//
// Solidity: function multiFreezeWithTimestamp(address[] _targets, uint256[] _timestamps) returns(bool)
func (_Ytm *YtmSession) MultiFreezeWithTimestamp(_targets []common.Address, _timestamps []*big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.MultiFreezeWithTimestamp(&_Ytm.TransactOpts, _targets, _timestamps)
}

// MultiFreezeWithTimestamp is a paid mutator transaction binding the contract method 0xd950c432.
//
// Solidity: function multiFreezeWithTimestamp(address[] _targets, uint256[] _timestamps) returns(bool)
func (_Ytm *YtmTransactorSession) MultiFreezeWithTimestamp(_targets []common.Address, _timestamps []*big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.MultiFreezeWithTimestamp(&_Ytm.TransactOpts, _targets, _timestamps)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0x1e89d545.
//
// Solidity: function multiTransfer(address[] _tos, uint256[] _values) returns(bool)
func (_Ytm *YtmTransactor) MultiTransfer(opts *bind.TransactOpts, _tos []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "multiTransfer", _tos, _values)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0x1e89d545.
//
// Solidity: function multiTransfer(address[] _tos, uint256[] _values) returns(bool)
func (_Ytm *YtmSession) MultiTransfer(_tos []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.MultiTransfer(&_Ytm.TransactOpts, _tos, _values)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0x1e89d545.
//
// Solidity: function multiTransfer(address[] _tos, uint256[] _values) returns(bool)
func (_Ytm *YtmTransactorSession) MultiTransfer(_tos []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.MultiTransfer(&_Ytm.TransactOpts, _tos, _values)
}

// RemoveRule is a paid mutator transaction binding the contract method 0xdf21950f.
//
// Solidity: function removeRule(address addr) returns(bool)
func (_Ytm *YtmTransactor) RemoveRule(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "removeRule", addr)
}

// RemoveRule is a paid mutator transaction binding the contract method 0xdf21950f.
//
// Solidity: function removeRule(address addr) returns(bool)
func (_Ytm *YtmSession) RemoveRule(addr common.Address) (*types.Transaction, error) {
	return _Ytm.Contract.RemoveRule(&_Ytm.TransactOpts, addr)
}

// RemoveRule is a paid mutator transaction binding the contract method 0xdf21950f.
//
// Solidity: function removeRule(address addr) returns(bool)
func (_Ytm *YtmTransactorSession) RemoveRule(addr common.Address) (*types.Transaction, error) {
	return _Ytm.Contract.RemoveRule(&_Ytm.TransactOpts, addr)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Ytm *YtmTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Ytm *YtmSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.Transfer(&_Ytm.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Ytm *YtmTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.Transfer(&_Ytm.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Ytm *YtmTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Ytm *YtmSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.TransferFrom(&_Ytm.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Ytm *YtmTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.TransferFrom(&_Ytm.TransactOpts, _from, _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_Ytm *YtmTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Ytm.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_Ytm *YtmSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.Withdraw(&_Ytm.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_Ytm *YtmTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Ytm.Contract.Withdraw(&_Ytm.TransactOpts, _amount)
}

// YtmFreezedIterator is returned from FilterFreezed and is used to iterate over the raw logs and unpacked data for Freezed events raised by the Ytm contract.
type YtmFreezedIterator struct {
	Event *YtmFreezed // Event containing the contract specifics and raw log

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
func (it *YtmFreezedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YtmFreezed)
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
		it.Event = new(YtmFreezed)
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
func (it *YtmFreezedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YtmFreezedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YtmFreezed represents a Freezed event raised by the Ytm contract.
type YtmFreezed struct {
	Sender  common.Address
	Balance *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFreezed is a free log retrieval operation binding the contract event 0x0774af36e8d3fbb76be8b6b764c3e3372138e0b332167b99da8bb7236cd1b3f3.
//
// Solidity: event Freezed(address sender, uint256 balance, bytes data)
func (_Ytm *YtmFilterer) FilterFreezed(opts *bind.FilterOpts) (*YtmFreezedIterator, error) {

	logs, sub, err := _Ytm.contract.FilterLogs(opts, "Freezed")
	if err != nil {
		return nil, err
	}
	return &YtmFreezedIterator{contract: _Ytm.contract, event: "Freezed", logs: logs, sub: sub}, nil
}

// WatchFreezed is a free log subscription operation binding the contract event 0x0774af36e8d3fbb76be8b6b764c3e3372138e0b332167b99da8bb7236cd1b3f3.
//
// Solidity: event Freezed(address sender, uint256 balance, bytes data)
func (_Ytm *YtmFilterer) WatchFreezed(opts *bind.WatchOpts, sink chan<- *YtmFreezed) (event.Subscription, error) {

	logs, sub, err := _Ytm.contract.WatchLogs(opts, "Freezed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YtmFreezed)
				if err := _Ytm.contract.UnpackLog(event, "Freezed", log); err != nil {
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
func (_Ytm *YtmFilterer) ParseFreezed(log types.Log) (*YtmFreezed, error) {
	event := new(YtmFreezed)
	if err := _Ytm.contract.UnpackLog(event, "Freezed", log); err != nil {
		return nil, err
	}
	return event, nil
}
