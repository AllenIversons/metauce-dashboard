// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
)

// CodeIQLFInviteCode is an auto generated low-level Go binding around an user-defined struct.
type CodeIQLFInviteCode struct {
	AllowUser   bool
	Owner       common.Address
	UserAddress common.Address
	Power       *big.Int
	Attr1       string
	Attr2       string
}

// InvitecodeMetaData contains all meta data concerning the Invitecode contract.
var InvitecodeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\",\"payable\":true},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_code\",\"type\":\"string\"}],\"name\":\"isValidCode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"checkUserInviteCodes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_code\",\"type\":\"string\"}],\"name\":\"registerInviteCodes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_code\",\"type\":\"string\"}],\"name\":\"loginInviteCodes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_codes\",\"type\":\"string[]\"}],\"name\":\"registerInviteCodesOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allow\",\"type\":\"bool\"}],\"name\":\"addAllowAppore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getAllowAppore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getMyAllInviteCode\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowUser\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"attr1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"attr2\",\"type\":\"string\"}],\"internalType\":\"structCodeIQLF.InviteCode[]\",\"name\":\"_inviteCodes\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getTotalPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_power\",\"type\":\"uint256\"}],\"name\":\"inviteCodePower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"setMinResourceCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinResourceCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]",
}

// InvitecodeABI is the input ABI used to generate the binding from.
// Deprecated: Use InvitecodeMetaData.ABI instead.
var InvitecodeABI = InvitecodeMetaData.ABI

// Invitecode is an auto generated Go binding around an Ethereum contract.
type Invitecode struct {
	InvitecodeCaller     // Read-only binding to the contract
	InvitecodeTransactor // Write-only binding to the contract
	InvitecodeFilterer   // Log filterer for contract events
}

// InvitecodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type InvitecodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvitecodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InvitecodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvitecodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InvitecodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvitecodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InvitecodeSession struct {
	Contract     *Invitecode       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InvitecodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InvitecodeCallerSession struct {
	Contract *InvitecodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// InvitecodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InvitecodeTransactorSession struct {
	Contract     *InvitecodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// InvitecodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type InvitecodeRaw struct {
	Contract *Invitecode // Generic contract binding to access the raw methods on
}

// InvitecodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InvitecodeCallerRaw struct {
	Contract *InvitecodeCaller // Generic read-only contract binding to access the raw methods on
}

// InvitecodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InvitecodeTransactorRaw struct {
	Contract *InvitecodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInvitecode creates a new instance of Invitecode, bound to a specific deployed contract.
func NewInvitecode(address common.Address, backend bind.ContractBackend) (*Invitecode, error) {
	contract, err := bindInvitecode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Invitecode{InvitecodeCaller: InvitecodeCaller{contract: contract}, InvitecodeTransactor: InvitecodeTransactor{contract: contract}, InvitecodeFilterer: InvitecodeFilterer{contract: contract}}, nil
}

// NewInvitecodeCaller creates a new read-only instance of Invitecode, bound to a specific deployed contract.
func NewInvitecodeCaller(address common.Address, caller bind.ContractCaller) (*InvitecodeCaller, error) {
	contract, err := bindInvitecode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InvitecodeCaller{contract: contract}, nil
}

// NewInvitecodeTransactor creates a new write-only instance of Invitecode, bound to a specific deployed contract.
func NewInvitecodeTransactor(address common.Address, transactor bind.ContractTransactor) (*InvitecodeTransactor, error) {
	contract, err := bindInvitecode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InvitecodeTransactor{contract: contract}, nil
}

// NewInvitecodeFilterer creates a new log filterer instance of Invitecode, bound to a specific deployed contract.
func NewInvitecodeFilterer(address common.Address, filterer bind.ContractFilterer) (*InvitecodeFilterer, error) {
	contract, err := bindInvitecode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InvitecodeFilterer{contract: contract}, nil
}

// bindInvitecode binds a generic wrapper to an already deployed contract.
func bindInvitecode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InvitecodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Invitecode *InvitecodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Invitecode.Contract.InvitecodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Invitecode *InvitecodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Invitecode.Contract.InvitecodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Invitecode *InvitecodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Invitecode.Contract.InvitecodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Invitecode *InvitecodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Invitecode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Invitecode *InvitecodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Invitecode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Invitecode *InvitecodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Invitecode.Contract.contract.Transact(opts, method, params...)
}

// CheckUserInviteCodes is a free data retrieval call binding the contract method 0x058f0205.
//
// Solidity: function checkUserInviteCodes(address _address) view returns(bool)
func (_Invitecode *InvitecodeCaller) CheckUserInviteCodes(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Invitecode.contract.Call(opts, &out, "checkUserInviteCodes", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckUserInviteCodes is a free data retrieval call binding the contract method 0x058f0205.
//
// Solidity: function checkUserInviteCodes(address _address) view returns(bool)
func (_Invitecode *InvitecodeSession) CheckUserInviteCodes(_address common.Address) (bool, error) {
	return _Invitecode.Contract.CheckUserInviteCodes(&_Invitecode.CallOpts, _address)
}

// CheckUserInviteCodes is a free data retrieval call binding the contract method 0x058f0205.
//
// Solidity: function checkUserInviteCodes(address _address) view returns(bool)
func (_Invitecode *InvitecodeCallerSession) CheckUserInviteCodes(_address common.Address) (bool, error) {
	return _Invitecode.Contract.CheckUserInviteCodes(&_Invitecode.CallOpts, _address)
}

// GetAllowAppore is a free data retrieval call binding the contract method 0x5ac17c2e.
//
// Solidity: function getAllowAppore(address _owner) view returns(bool)
func (_Invitecode *InvitecodeCaller) GetAllowAppore(opts *bind.CallOpts, _owner common.Address) (bool, error) {
	var out []interface{}
	err := _Invitecode.contract.Call(opts, &out, "getAllowAppore", _owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetAllowAppore is a free data retrieval call binding the contract method 0x5ac17c2e.
//
// Solidity: function getAllowAppore(address _owner) view returns(bool)
func (_Invitecode *InvitecodeSession) GetAllowAppore(_owner common.Address) (bool, error) {
	return _Invitecode.Contract.GetAllowAppore(&_Invitecode.CallOpts, _owner)
}

// GetAllowAppore is a free data retrieval call binding the contract method 0x5ac17c2e.
//
// Solidity: function getAllowAppore(address _owner) view returns(bool)
func (_Invitecode *InvitecodeCallerSession) GetAllowAppore(_owner common.Address) (bool, error) {
	return _Invitecode.Contract.GetAllowAppore(&_Invitecode.CallOpts, _owner)
}

// GetMinResourceCount is a free data retrieval call binding the contract method 0x7a364836.
//
// Solidity: function getMinResourceCount() view returns(uint256)
func (_Invitecode *InvitecodeCaller) GetMinResourceCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Invitecode.contract.Call(opts, &out, "getMinResourceCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinResourceCount is a free data retrieval call binding the contract method 0x7a364836.
//
// Solidity: function getMinResourceCount() view returns(uint256)
func (_Invitecode *InvitecodeSession) GetMinResourceCount() (*big.Int, error) {
	return _Invitecode.Contract.GetMinResourceCount(&_Invitecode.CallOpts)
}

// GetMinResourceCount is a free data retrieval call binding the contract method 0x7a364836.
//
// Solidity: function getMinResourceCount() view returns(uint256)
func (_Invitecode *InvitecodeCallerSession) GetMinResourceCount() (*big.Int, error) {
	return _Invitecode.Contract.GetMinResourceCount(&_Invitecode.CallOpts)
}

// GetMyAllInviteCode is a free data retrieval call binding the contract method 0xadc0ae9b.
//
// Solidity: function getMyAllInviteCode(address _owner) view returns((bool,address,address,uint256,string,string)[] _inviteCodes)
func (_Invitecode *InvitecodeCaller) GetMyAllInviteCode(opts *bind.CallOpts, _owner common.Address) ([]CodeIQLFInviteCode, error) {
	var out []interface{}
	err := _Invitecode.contract.Call(opts, &out, "getMyAllInviteCode", _owner)

	if err != nil {
		return *new([]CodeIQLFInviteCode), err
	}

	out0 := *abi.ConvertType(out[0], new([]CodeIQLFInviteCode)).(*[]CodeIQLFInviteCode)

	return out0, err

}

// GetMyAllInviteCode is a free data retrieval call binding the contract method 0xadc0ae9b.
//
// Solidity: function getMyAllInviteCode(address _owner) view returns((bool,address,address,uint256,string,string)[] _inviteCodes)
func (_Invitecode *InvitecodeSession) GetMyAllInviteCode(_owner common.Address) ([]CodeIQLFInviteCode, error) {
	return _Invitecode.Contract.GetMyAllInviteCode(&_Invitecode.CallOpts, _owner)
}

// GetMyAllInviteCode is a free data retrieval call binding the contract method 0xadc0ae9b.
//
// Solidity: function getMyAllInviteCode(address _owner) view returns((bool,address,address,uint256,string,string)[] _inviteCodes)
func (_Invitecode *InvitecodeCallerSession) GetMyAllInviteCode(_owner common.Address) ([]CodeIQLFInviteCode, error) {
	return _Invitecode.Contract.GetMyAllInviteCode(&_Invitecode.CallOpts, _owner)
}

// GetTotalPower is a free data retrieval call binding the contract method 0xd7dc81cd.
//
// Solidity: function getTotalPower(address _owner) view returns(uint256)
func (_Invitecode *InvitecodeCaller) GetTotalPower(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Invitecode.contract.Call(opts, &out, "getTotalPower", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPower is a free data retrieval call binding the contract method 0xd7dc81cd.
//
// Solidity: function getTotalPower(address _owner) view returns(uint256)
func (_Invitecode *InvitecodeSession) GetTotalPower(_owner common.Address) (*big.Int, error) {
	return _Invitecode.Contract.GetTotalPower(&_Invitecode.CallOpts, _owner)
}

// GetTotalPower is a free data retrieval call binding the contract method 0xd7dc81cd.
//
// Solidity: function getTotalPower(address _owner) view returns(uint256)
func (_Invitecode *InvitecodeCallerSession) GetTotalPower(_owner common.Address) (*big.Int, error) {
	return _Invitecode.Contract.GetTotalPower(&_Invitecode.CallOpts, _owner)
}

// IsValidCode is a free data retrieval call binding the contract method 0x25ed64a0.
//
// Solidity: function isValidCode(string _code) view returns(bool)
func (_Invitecode *InvitecodeCaller) IsValidCode(opts *bind.CallOpts, _code string) (bool, error) {
	var out []interface{}
	err := _Invitecode.contract.Call(opts, &out, "isValidCode", _code)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidCode is a free data retrieval call binding the contract method 0x25ed64a0.
//
// Solidity: function isValidCode(string _code) view returns(bool)
func (_Invitecode *InvitecodeSession) IsValidCode(_code string) (bool, error) {
	return _Invitecode.Contract.IsValidCode(&_Invitecode.CallOpts, _code)
}

// IsValidCode is a free data retrieval call binding the contract method 0x25ed64a0.
//
// Solidity: function isValidCode(string _code) view returns(bool)
func (_Invitecode *InvitecodeCallerSession) IsValidCode(_code string) (bool, error) {
	return _Invitecode.Contract.IsValidCode(&_Invitecode.CallOpts, _code)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Invitecode *InvitecodeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Invitecode.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Invitecode *InvitecodeSession) Owner() (common.Address, error) {
	return _Invitecode.Contract.Owner(&_Invitecode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Invitecode *InvitecodeCallerSession) Owner() (common.Address, error) {
	return _Invitecode.Contract.Owner(&_Invitecode.CallOpts)
}

// AddAllowAppore is a paid mutator transaction binding the contract method 0xff6a78f5.
//
// Solidity: function addAllowAppore(address _owner, bool _allow) returns()
func (_Invitecode *InvitecodeTransactor) AddAllowAppore(opts *bind.TransactOpts, _owner common.Address, _allow bool) (*types.Transaction, error) {
	return _Invitecode.contract.Transact(opts, "addAllowAppore", _owner, _allow)
}

// AddAllowAppore is a paid mutator transaction binding the contract method 0xff6a78f5.
//
// Solidity: function addAllowAppore(address _owner, bool _allow) returns()
func (_Invitecode *InvitecodeSession) AddAllowAppore(_owner common.Address, _allow bool) (*types.Transaction, error) {
	return _Invitecode.Contract.AddAllowAppore(&_Invitecode.TransactOpts, _owner, _allow)
}

// AddAllowAppore is a paid mutator transaction binding the contract method 0xff6a78f5.
//
// Solidity: function addAllowAppore(address _owner, bool _allow) returns()
func (_Invitecode *InvitecodeTransactorSession) AddAllowAppore(_owner common.Address, _allow bool) (*types.Transaction, error) {
	return _Invitecode.Contract.AddAllowAppore(&_Invitecode.TransactOpts, _owner, _allow)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Invitecode *InvitecodeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Invitecode.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Invitecode *InvitecodeSession) Initialize() (*types.Transaction, error) {
	return _Invitecode.Contract.Initialize(&_Invitecode.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Invitecode *InvitecodeTransactorSession) Initialize() (*types.Transaction, error) {
	return _Invitecode.Contract.Initialize(&_Invitecode.TransactOpts)
}

// InviteCodePower is a paid mutator transaction binding the contract method 0x2f19ede0.
//
// Solidity: function inviteCodePower(address _address, uint256 _power) returns(uint256)
func (_Invitecode *InvitecodeTransactor) InviteCodePower(opts *bind.TransactOpts, _address common.Address, _power *big.Int) (*types.Transaction, error) {
	return _Invitecode.contract.Transact(opts, "inviteCodePower", _address, _power)
}

// InviteCodePower is a paid mutator transaction binding the contract method 0x2f19ede0.
//
// Solidity: function inviteCodePower(address _address, uint256 _power) returns(uint256)
func (_Invitecode *InvitecodeSession) InviteCodePower(_address common.Address, _power *big.Int) (*types.Transaction, error) {
	return _Invitecode.Contract.InviteCodePower(&_Invitecode.TransactOpts, _address, _power)
}

// InviteCodePower is a paid mutator transaction binding the contract method 0x2f19ede0.
//
// Solidity: function inviteCodePower(address _address, uint256 _power) returns(uint256)
func (_Invitecode *InvitecodeTransactorSession) InviteCodePower(_address common.Address, _power *big.Int) (*types.Transaction, error) {
	return _Invitecode.Contract.InviteCodePower(&_Invitecode.TransactOpts, _address, _power)
}

// LoginInviteCodes is a paid mutator transaction binding the contract method 0x553fc285.
//
// Solidity: function loginInviteCodes(address _address, string _code) returns()
func (_Invitecode *InvitecodeTransactor) LoginInviteCodes(opts *bind.TransactOpts, _address common.Address, _code string) (*types.Transaction, error) {
	return _Invitecode.contract.Transact(opts, "loginInviteCodes", _address, _code)
}

// LoginInviteCodes is a paid mutator transaction binding the contract method 0x553fc285.
//
// Solidity: function loginInviteCodes(address _address, string _code) returns()
func (_Invitecode *InvitecodeSession) LoginInviteCodes(_address common.Address, _code string) (*types.Transaction, error) {
	return _Invitecode.Contract.LoginInviteCodes(&_Invitecode.TransactOpts, _address, _code)
}

// LoginInviteCodes is a paid mutator transaction binding the contract method 0x553fc285.
//
// Solidity: function loginInviteCodes(address _address, string _code) returns()
func (_Invitecode *InvitecodeTransactorSession) LoginInviteCodes(_address common.Address, _code string) (*types.Transaction, error) {
	return _Invitecode.Contract.LoginInviteCodes(&_Invitecode.TransactOpts, _address, _code)
}

// RegisterInviteCodes is a paid mutator transaction binding the contract method 0xb24a0cc9.
//
// Solidity: function registerInviteCodes(address from, string _code) returns()
func (_Invitecode *InvitecodeTransactor) RegisterInviteCodes(opts *bind.TransactOpts, from common.Address, _code string) (*types.Transaction, error) {
	return _Invitecode.contract.Transact(opts, "registerInviteCodes", from, _code)
}

// RegisterInviteCodes is a paid mutator transaction binding the contract method 0xb24a0cc9.
//
// Solidity: function registerInviteCodes(address from, string _code) returns()
func (_Invitecode *InvitecodeSession) RegisterInviteCodes(from common.Address, _code string) (*types.Transaction, error) {
	return _Invitecode.Contract.RegisterInviteCodes(&_Invitecode.TransactOpts, from, _code)
}

// RegisterInviteCodes is a paid mutator transaction binding the contract method 0xb24a0cc9.
//
// Solidity: function registerInviteCodes(address from, string _code) returns()
func (_Invitecode *InvitecodeTransactorSession) RegisterInviteCodes(from common.Address, _code string) (*types.Transaction, error) {
	return _Invitecode.Contract.RegisterInviteCodes(&_Invitecode.TransactOpts, from, _code)
}

// RegisterInviteCodesOwner is a paid mutator transaction binding the contract method 0xbef22d25.
//
// Solidity: function registerInviteCodesOwner(address _owner, string[] _codes) returns()
func (_Invitecode *InvitecodeTransactor) RegisterInviteCodesOwner(opts *bind.TransactOpts, _owner common.Address, _codes []string) (*types.Transaction, error) {
	return _Invitecode.contract.Transact(opts, "registerInviteCodesOwner", _owner, _codes)
}

// RegisterInviteCodesOwner is a paid mutator transaction binding the contract method 0xbef22d25.
//
// Solidity: function registerInviteCodesOwner(address _owner, string[] _codes) returns()
func (_Invitecode *InvitecodeSession) RegisterInviteCodesOwner(_owner common.Address, _codes []string) (*types.Transaction, error) {
	return _Invitecode.Contract.RegisterInviteCodesOwner(&_Invitecode.TransactOpts, _owner, _codes)
}

// RegisterInviteCodesOwner is a paid mutator transaction binding the contract method 0xbef22d25.
//
// Solidity: function registerInviteCodesOwner(address _owner, string[] _codes) returns()
func (_Invitecode *InvitecodeTransactorSession) RegisterInviteCodesOwner(_owner common.Address, _codes []string) (*types.Transaction, error) {
	return _Invitecode.Contract.RegisterInviteCodesOwner(&_Invitecode.TransactOpts, _owner, _codes)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Invitecode *InvitecodeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Invitecode.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Invitecode *InvitecodeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Invitecode.Contract.RenounceOwnership(&_Invitecode.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Invitecode *InvitecodeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Invitecode.Contract.RenounceOwnership(&_Invitecode.TransactOpts)
}

// SetMinResourceCount is a paid mutator transaction binding the contract method 0x2726f796.
//
// Solidity: function setMinResourceCount(uint256 count) returns()
func (_Invitecode *InvitecodeTransactor) SetMinResourceCount(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _Invitecode.contract.Transact(opts, "setMinResourceCount", count)
}

// SetMinResourceCount is a paid mutator transaction binding the contract method 0x2726f796.
//
// Solidity: function setMinResourceCount(uint256 count) returns()
func (_Invitecode *InvitecodeSession) SetMinResourceCount(count *big.Int) (*types.Transaction, error) {
	return _Invitecode.Contract.SetMinResourceCount(&_Invitecode.TransactOpts, count)
}

// SetMinResourceCount is a paid mutator transaction binding the contract method 0x2726f796.
//
// Solidity: function setMinResourceCount(uint256 count) returns()
func (_Invitecode *InvitecodeTransactorSession) SetMinResourceCount(count *big.Int) (*types.Transaction, error) {
	return _Invitecode.Contract.SetMinResourceCount(&_Invitecode.TransactOpts, count)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Invitecode *InvitecodeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Invitecode.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Invitecode *InvitecodeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Invitecode.Contract.TransferOwnership(&_Invitecode.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Invitecode *InvitecodeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Invitecode.Contract.TransferOwnership(&_Invitecode.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Invitecode *InvitecodeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Invitecode.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Invitecode *InvitecodeSession) Receive() (*types.Transaction, error) {
	return _Invitecode.Contract.Receive(&_Invitecode.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Invitecode *InvitecodeTransactorSession) Receive() (*types.Transaction, error) {
	return _Invitecode.Contract.Receive(&_Invitecode.TransactOpts)
}

// InvitecodeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Invitecode contract.
type InvitecodeOwnershipTransferredIterator struct {
	Event *InvitecodeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InvitecodeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InvitecodeOwnershipTransferred)
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
		it.Event = new(InvitecodeOwnershipTransferred)
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
func (it *InvitecodeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InvitecodeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InvitecodeOwnershipTransferred represents a OwnershipTransferred event raised by the Invitecode contract.
type InvitecodeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Invitecode *InvitecodeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InvitecodeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Invitecode.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InvitecodeOwnershipTransferredIterator{contract: _Invitecode.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Invitecode *InvitecodeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InvitecodeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Invitecode.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InvitecodeOwnershipTransferred)
				if err := _Invitecode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Invitecode *InvitecodeFilterer) ParseOwnershipTransferred(log types.Log) (*InvitecodeOwnershipTransferred, error) {
	event := new(InvitecodeOwnershipTransferred)
	if err := _Invitecode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
