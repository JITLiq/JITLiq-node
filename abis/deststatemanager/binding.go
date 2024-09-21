// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package deststatemanager

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

// DeststatemanagerMetaData contains all meta data concerning the Deststatemanager contract.
var DeststatemanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddressRegistry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NotGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEntrypoint\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"addressRegistry\",\"outputs\":[{\"internalType\":\"contractAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseBridgeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountDeduct\",\"type\":\"uint256\"}],\"name\":\"deductStakedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newOperators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"newTotalStakedFunds\",\"type\":\"uint256\"}],\"name\":\"syncSourceData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStakedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DeststatemanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DeststatemanagerMetaData.ABI instead.
var DeststatemanagerABI = DeststatemanagerMetaData.ABI

// Deststatemanager is an auto generated Go binding around an Ethereum contract.
type Deststatemanager struct {
	DeststatemanagerCaller     // Read-only binding to the contract
	DeststatemanagerTransactor // Write-only binding to the contract
	DeststatemanagerFilterer   // Log filterer for contract events
}

// DeststatemanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeststatemanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeststatemanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeststatemanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeststatemanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeststatemanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeststatemanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeststatemanagerSession struct {
	Contract     *Deststatemanager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeststatemanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeststatemanagerCallerSession struct {
	Contract *DeststatemanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DeststatemanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeststatemanagerTransactorSession struct {
	Contract     *DeststatemanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DeststatemanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeststatemanagerRaw struct {
	Contract *Deststatemanager // Generic contract binding to access the raw methods on
}

// DeststatemanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeststatemanagerCallerRaw struct {
	Contract *DeststatemanagerCaller // Generic read-only contract binding to access the raw methods on
}

// DeststatemanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeststatemanagerTransactorRaw struct {
	Contract *DeststatemanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeststatemanager creates a new instance of Deststatemanager, bound to a specific deployed contract.
func NewDeststatemanager(address common.Address, backend bind.ContractBackend) (*Deststatemanager, error) {
	contract, err := bindDeststatemanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deststatemanager{DeststatemanagerCaller: DeststatemanagerCaller{contract: contract}, DeststatemanagerTransactor: DeststatemanagerTransactor{contract: contract}, DeststatemanagerFilterer: DeststatemanagerFilterer{contract: contract}}, nil
}

// NewDeststatemanagerCaller creates a new read-only instance of Deststatemanager, bound to a specific deployed contract.
func NewDeststatemanagerCaller(address common.Address, caller bind.ContractCaller) (*DeststatemanagerCaller, error) {
	contract, err := bindDeststatemanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeststatemanagerCaller{contract: contract}, nil
}

// NewDeststatemanagerTransactor creates a new write-only instance of Deststatemanager, bound to a specific deployed contract.
func NewDeststatemanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DeststatemanagerTransactor, error) {
	contract, err := bindDeststatemanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeststatemanagerTransactor{contract: contract}, nil
}

// NewDeststatemanagerFilterer creates a new log filterer instance of Deststatemanager, bound to a specific deployed contract.
func NewDeststatemanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DeststatemanagerFilterer, error) {
	contract, err := bindDeststatemanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeststatemanagerFilterer{contract: contract}, nil
}

// bindDeststatemanager binds a generic wrapper to an already deployed contract.
func bindDeststatemanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DeststatemanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deststatemanager *DeststatemanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deststatemanager.Contract.DeststatemanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deststatemanager *DeststatemanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deststatemanager.Contract.DeststatemanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deststatemanager *DeststatemanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deststatemanager.Contract.DeststatemanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deststatemanager *DeststatemanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deststatemanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deststatemanager *DeststatemanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deststatemanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deststatemanager *DeststatemanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deststatemanager.Contract.contract.Transact(opts, method, params...)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_Deststatemanager *DeststatemanagerCaller) AddressRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deststatemanager.contract.Call(opts, &out, "addressRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_Deststatemanager *DeststatemanagerSession) AddressRegistry() (common.Address, error) {
	return _Deststatemanager.Contract.AddressRegistry(&_Deststatemanager.CallOpts)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_Deststatemanager *DeststatemanagerCallerSession) AddressRegistry() (common.Address, error) {
	return _Deststatemanager.Contract.AddressRegistry(&_Deststatemanager.CallOpts)
}

// BaseBridgeToken is a free data retrieval call binding the contract method 0x15e888b5.
//
// Solidity: function baseBridgeToken() view returns(address)
func (_Deststatemanager *DeststatemanagerCaller) BaseBridgeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deststatemanager.contract.Call(opts, &out, "baseBridgeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseBridgeToken is a free data retrieval call binding the contract method 0x15e888b5.
//
// Solidity: function baseBridgeToken() view returns(address)
func (_Deststatemanager *DeststatemanagerSession) BaseBridgeToken() (common.Address, error) {
	return _Deststatemanager.Contract.BaseBridgeToken(&_Deststatemanager.CallOpts)
}

// BaseBridgeToken is a free data retrieval call binding the contract method 0x15e888b5.
//
// Solidity: function baseBridgeToken() view returns(address)
func (_Deststatemanager *DeststatemanagerCallerSession) BaseBridgeToken() (common.Address, error) {
	return _Deststatemanager.Contract.BaseBridgeToken(&_Deststatemanager.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_Deststatemanager *DeststatemanagerCaller) GetOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Deststatemanager.contract.Call(opts, &out, "getOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_Deststatemanager *DeststatemanagerSession) GetOperators() ([]common.Address, error) {
	return _Deststatemanager.Contract.GetOperators(&_Deststatemanager.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_Deststatemanager *DeststatemanagerCallerSession) GetOperators() ([]common.Address, error) {
	return _Deststatemanager.Contract.GetOperators(&_Deststatemanager.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0xe28d4906.
//
// Solidity: function operators(uint256 ) view returns(address)
func (_Deststatemanager *DeststatemanagerCaller) Operators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Deststatemanager.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0xe28d4906.
//
// Solidity: function operators(uint256 ) view returns(address)
func (_Deststatemanager *DeststatemanagerSession) Operators(arg0 *big.Int) (common.Address, error) {
	return _Deststatemanager.Contract.Operators(&_Deststatemanager.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0xe28d4906.
//
// Solidity: function operators(uint256 ) view returns(address)
func (_Deststatemanager *DeststatemanagerCallerSession) Operators(arg0 *big.Int) (common.Address, error) {
	return _Deststatemanager.Contract.Operators(&_Deststatemanager.CallOpts, arg0)
}

// TotalStakedFunds is a free data retrieval call binding the contract method 0x18ec2d5c.
//
// Solidity: function totalStakedFunds() view returns(uint256)
func (_Deststatemanager *DeststatemanagerCaller) TotalStakedFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deststatemanager.contract.Call(opts, &out, "totalStakedFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakedFunds is a free data retrieval call binding the contract method 0x18ec2d5c.
//
// Solidity: function totalStakedFunds() view returns(uint256)
func (_Deststatemanager *DeststatemanagerSession) TotalStakedFunds() (*big.Int, error) {
	return _Deststatemanager.Contract.TotalStakedFunds(&_Deststatemanager.CallOpts)
}

// TotalStakedFunds is a free data retrieval call binding the contract method 0x18ec2d5c.
//
// Solidity: function totalStakedFunds() view returns(uint256)
func (_Deststatemanager *DeststatemanagerCallerSession) TotalStakedFunds() (*big.Int, error) {
	return _Deststatemanager.Contract.TotalStakedFunds(&_Deststatemanager.CallOpts)
}

// DeductStakedFunds is a paid mutator transaction binding the contract method 0xde17abe4.
//
// Solidity: function deductStakedFunds(uint256 amountDeduct) returns()
func (_Deststatemanager *DeststatemanagerTransactor) DeductStakedFunds(opts *bind.TransactOpts, amountDeduct *big.Int) (*types.Transaction, error) {
	return _Deststatemanager.contract.Transact(opts, "deductStakedFunds", amountDeduct)
}

// DeductStakedFunds is a paid mutator transaction binding the contract method 0xde17abe4.
//
// Solidity: function deductStakedFunds(uint256 amountDeduct) returns()
func (_Deststatemanager *DeststatemanagerSession) DeductStakedFunds(amountDeduct *big.Int) (*types.Transaction, error) {
	return _Deststatemanager.Contract.DeductStakedFunds(&_Deststatemanager.TransactOpts, amountDeduct)
}

// DeductStakedFunds is a paid mutator transaction binding the contract method 0xde17abe4.
//
// Solidity: function deductStakedFunds(uint256 amountDeduct) returns()
func (_Deststatemanager *DeststatemanagerTransactorSession) DeductStakedFunds(amountDeduct *big.Int) (*types.Transaction, error) {
	return _Deststatemanager.Contract.DeductStakedFunds(&_Deststatemanager.TransactOpts, amountDeduct)
}

// SyncSourceData is a paid mutator transaction binding the contract method 0xfc06d9c2.
//
// Solidity: function syncSourceData(address[] newOperators, uint256 newTotalStakedFunds) returns()
func (_Deststatemanager *DeststatemanagerTransactor) SyncSourceData(opts *bind.TransactOpts, newOperators []common.Address, newTotalStakedFunds *big.Int) (*types.Transaction, error) {
	return _Deststatemanager.contract.Transact(opts, "syncSourceData", newOperators, newTotalStakedFunds)
}

// SyncSourceData is a paid mutator transaction binding the contract method 0xfc06d9c2.
//
// Solidity: function syncSourceData(address[] newOperators, uint256 newTotalStakedFunds) returns()
func (_Deststatemanager *DeststatemanagerSession) SyncSourceData(newOperators []common.Address, newTotalStakedFunds *big.Int) (*types.Transaction, error) {
	return _Deststatemanager.Contract.SyncSourceData(&_Deststatemanager.TransactOpts, newOperators, newTotalStakedFunds)
}

// SyncSourceData is a paid mutator transaction binding the contract method 0xfc06d9c2.
//
// Solidity: function syncSourceData(address[] newOperators, uint256 newTotalStakedFunds) returns()
func (_Deststatemanager *DeststatemanagerTransactorSession) SyncSourceData(newOperators []common.Address, newTotalStakedFunds *big.Int) (*types.Transaction, error) {
	return _Deststatemanager.Contract.SyncSourceData(&_Deststatemanager.TransactOpts, newOperators, newTotalStakedFunds)
}
