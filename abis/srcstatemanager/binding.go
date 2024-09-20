// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package srcstatemanager

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

// ISourceOpStateManagerFeesData is an auto generated low-level Go binding around an user-defined struct.
type ISourceOpStateManagerFeesData struct {
	OperationFee *big.Int
	BridgeFee    *big.Int
}

// ISourceOpStateManagerFulfillerData is an auto generated low-level Go binding around an user-defined struct.
type ISourceOpStateManagerFulfillerData struct {
	FulfillAmount *big.Int
	Fulfiller     common.Address
}

// ISourceOpStateManagerOperatorData is an auto generated low-level Go binding around an user-defined struct.
type ISourceOpStateManagerOperatorData struct {
	CurrentStake   *big.Int
	CurrentHolding *big.Int
	Registered     bool
}

// SrcstatemanagerMetaData contains all meta data concerning the Srcstatemanager contract.
var SrcstatemanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_baseBridgeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddressRegistry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NotGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEntrypoint\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"OperatorStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LP_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressRegistry\",\"outputs\":[{\"internalType\":\"contractAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseBridgeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"}],\"name\":\"completeOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"expiry\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"orderAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"increaseStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentHolding\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"expiry\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"orderAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"operationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"}],\"internalType\":\"structISourceOpStateManager.FeesData\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentHolding\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"internalType\":\"structISourceOpStateManager.OperatorData\",\"name\":\"newOperatorData\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"deleteOperator\",\"type\":\"bool\"}],\"name\":\"syncOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"holdingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"init\",\"type\":\"bool\"}],\"name\":\"updateOperatorAllocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fulfillAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fulfiller\",\"type\":\"address\"}],\"internalType\":\"structISourceOpStateManager.FulfillerData[]\",\"name\":\"fulfillerData\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"lpFees\",\"type\":\"uint256\"}],\"name\":\"updatePendingRefunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SrcstatemanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use SrcstatemanagerMetaData.ABI instead.
var SrcstatemanagerABI = SrcstatemanagerMetaData.ABI

// Srcstatemanager is an auto generated Go binding around an Ethereum contract.
type Srcstatemanager struct {
	SrcstatemanagerCaller     // Read-only binding to the contract
	SrcstatemanagerTransactor // Write-only binding to the contract
	SrcstatemanagerFilterer   // Log filterer for contract events
}

// SrcstatemanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SrcstatemanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SrcstatemanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SrcstatemanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SrcstatemanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SrcstatemanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SrcstatemanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SrcstatemanagerSession struct {
	Contract     *Srcstatemanager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SrcstatemanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SrcstatemanagerCallerSession struct {
	Contract *SrcstatemanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SrcstatemanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SrcstatemanagerTransactorSession struct {
	Contract     *SrcstatemanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SrcstatemanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SrcstatemanagerRaw struct {
	Contract *Srcstatemanager // Generic contract binding to access the raw methods on
}

// SrcstatemanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SrcstatemanagerCallerRaw struct {
	Contract *SrcstatemanagerCaller // Generic read-only contract binding to access the raw methods on
}

// SrcstatemanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SrcstatemanagerTransactorRaw struct {
	Contract *SrcstatemanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSrcstatemanager creates a new instance of Srcstatemanager, bound to a specific deployed contract.
func NewSrcstatemanager(address common.Address, backend bind.ContractBackend) (*Srcstatemanager, error) {
	contract, err := bindSrcstatemanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Srcstatemanager{SrcstatemanagerCaller: SrcstatemanagerCaller{contract: contract}, SrcstatemanagerTransactor: SrcstatemanagerTransactor{contract: contract}, SrcstatemanagerFilterer: SrcstatemanagerFilterer{contract: contract}}, nil
}

// NewSrcstatemanagerCaller creates a new read-only instance of Srcstatemanager, bound to a specific deployed contract.
func NewSrcstatemanagerCaller(address common.Address, caller bind.ContractCaller) (*SrcstatemanagerCaller, error) {
	contract, err := bindSrcstatemanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SrcstatemanagerCaller{contract: contract}, nil
}

// NewSrcstatemanagerTransactor creates a new write-only instance of Srcstatemanager, bound to a specific deployed contract.
func NewSrcstatemanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*SrcstatemanagerTransactor, error) {
	contract, err := bindSrcstatemanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SrcstatemanagerTransactor{contract: contract}, nil
}

// NewSrcstatemanagerFilterer creates a new log filterer instance of Srcstatemanager, bound to a specific deployed contract.
func NewSrcstatemanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*SrcstatemanagerFilterer, error) {
	contract, err := bindSrcstatemanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SrcstatemanagerFilterer{contract: contract}, nil
}

// bindSrcstatemanager binds a generic wrapper to an already deployed contract.
func bindSrcstatemanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SrcstatemanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Srcstatemanager *SrcstatemanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Srcstatemanager.Contract.SrcstatemanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Srcstatemanager *SrcstatemanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.SrcstatemanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Srcstatemanager *SrcstatemanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.SrcstatemanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Srcstatemanager *SrcstatemanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Srcstatemanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Srcstatemanager *SrcstatemanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Srcstatemanager *SrcstatemanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.contract.Transact(opts, method, params...)
}

// LPFEE is a free data retrieval call binding the contract method 0x5e3f2727.
//
// Solidity: function LP_FEE() view returns(uint256)
func (_Srcstatemanager *SrcstatemanagerCaller) LPFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Srcstatemanager.contract.Call(opts, &out, "LP_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LPFEE is a free data retrieval call binding the contract method 0x5e3f2727.
//
// Solidity: function LP_FEE() view returns(uint256)
func (_Srcstatemanager *SrcstatemanagerSession) LPFEE() (*big.Int, error) {
	return _Srcstatemanager.Contract.LPFEE(&_Srcstatemanager.CallOpts)
}

// LPFEE is a free data retrieval call binding the contract method 0x5e3f2727.
//
// Solidity: function LP_FEE() view returns(uint256)
func (_Srcstatemanager *SrcstatemanagerCallerSession) LPFEE() (*big.Int, error) {
	return _Srcstatemanager.Contract.LPFEE(&_Srcstatemanager.CallOpts)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint256)
func (_Srcstatemanager *SrcstatemanagerCaller) MAXBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Srcstatemanager.contract.Call(opts, &out, "MAX_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint256)
func (_Srcstatemanager *SrcstatemanagerSession) MAXBPS() (*big.Int, error) {
	return _Srcstatemanager.Contract.MAXBPS(&_Srcstatemanager.CallOpts)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint256)
func (_Srcstatemanager *SrcstatemanagerCallerSession) MAXBPS() (*big.Int, error) {
	return _Srcstatemanager.Contract.MAXBPS(&_Srcstatemanager.CallOpts)
}

// OPERATORFEE is a free data retrieval call binding the contract method 0x7076c657.
//
// Solidity: function OPERATOR_FEE() view returns(uint256)
func (_Srcstatemanager *SrcstatemanagerCaller) OPERATORFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Srcstatemanager.contract.Call(opts, &out, "OPERATOR_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPERATORFEE is a free data retrieval call binding the contract method 0x7076c657.
//
// Solidity: function OPERATOR_FEE() view returns(uint256)
func (_Srcstatemanager *SrcstatemanagerSession) OPERATORFEE() (*big.Int, error) {
	return _Srcstatemanager.Contract.OPERATORFEE(&_Srcstatemanager.CallOpts)
}

// OPERATORFEE is a free data retrieval call binding the contract method 0x7076c657.
//
// Solidity: function OPERATOR_FEE() view returns(uint256)
func (_Srcstatemanager *SrcstatemanagerCallerSession) OPERATORFEE() (*big.Int, error) {
	return _Srcstatemanager.Contract.OPERATORFEE(&_Srcstatemanager.CallOpts)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_Srcstatemanager *SrcstatemanagerCaller) AddressRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Srcstatemanager.contract.Call(opts, &out, "addressRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_Srcstatemanager *SrcstatemanagerSession) AddressRegistry() (common.Address, error) {
	return _Srcstatemanager.Contract.AddressRegistry(&_Srcstatemanager.CallOpts)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_Srcstatemanager *SrcstatemanagerCallerSession) AddressRegistry() (common.Address, error) {
	return _Srcstatemanager.Contract.AddressRegistry(&_Srcstatemanager.CallOpts)
}

// BaseBridgeToken is a free data retrieval call binding the contract method 0x15e888b5.
//
// Solidity: function baseBridgeToken() view returns(address)
func (_Srcstatemanager *SrcstatemanagerCaller) BaseBridgeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Srcstatemanager.contract.Call(opts, &out, "baseBridgeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseBridgeToken is a free data retrieval call binding the contract method 0x15e888b5.
//
// Solidity: function baseBridgeToken() view returns(address)
func (_Srcstatemanager *SrcstatemanagerSession) BaseBridgeToken() (common.Address, error) {
	return _Srcstatemanager.Contract.BaseBridgeToken(&_Srcstatemanager.CallOpts)
}

// BaseBridgeToken is a free data retrieval call binding the contract method 0x15e888b5.
//
// Solidity: function baseBridgeToken() view returns(address)
func (_Srcstatemanager *SrcstatemanagerCallerSession) BaseBridgeToken() (common.Address, error) {
	return _Srcstatemanager.Contract.BaseBridgeToken(&_Srcstatemanager.CallOpts)
}

// OperatorData is a free data retrieval call binding the contract method 0x333d0ef0.
//
// Solidity: function operatorData(address ) view returns(uint256 currentStake, uint256 currentHolding, bool registered)
func (_Srcstatemanager *SrcstatemanagerCaller) OperatorData(opts *bind.CallOpts, arg0 common.Address) (struct {
	CurrentStake   *big.Int
	CurrentHolding *big.Int
	Registered     bool
}, error) {
	var out []interface{}
	err := _Srcstatemanager.contract.Call(opts, &out, "operatorData", arg0)

	outstruct := new(struct {
		CurrentStake   *big.Int
		CurrentHolding *big.Int
		Registered     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentStake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentHolding = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Registered = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// OperatorData is a free data retrieval call binding the contract method 0x333d0ef0.
//
// Solidity: function operatorData(address ) view returns(uint256 currentStake, uint256 currentHolding, bool registered)
func (_Srcstatemanager *SrcstatemanagerSession) OperatorData(arg0 common.Address) (struct {
	CurrentStake   *big.Int
	CurrentHolding *big.Int
	Registered     bool
}, error) {
	return _Srcstatemanager.Contract.OperatorData(&_Srcstatemanager.CallOpts, arg0)
}

// OperatorData is a free data retrieval call binding the contract method 0x333d0ef0.
//
// Solidity: function operatorData(address ) view returns(uint256 currentStake, uint256 currentHolding, bool registered)
func (_Srcstatemanager *SrcstatemanagerCallerSession) OperatorData(arg0 common.Address) (struct {
	CurrentStake   *big.Int
	CurrentHolding *big.Int
	Registered     bool
}, error) {
	return _Srcstatemanager.Contract.OperatorData(&_Srcstatemanager.CallOpts, arg0)
}

// OrderData is a free data retrieval call binding the contract method 0xc4cd9ed3.
//
// Solidity: function orderData(bytes32 ) view returns(bool fulfilled, uint32 expiry, uint256 orderAmount, address destAddress, address operator, (uint256,uint256) fees)
func (_Srcstatemanager *SrcstatemanagerCaller) OrderData(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Fulfilled   bool
	Expiry      uint32
	OrderAmount *big.Int
	DestAddress common.Address
	Operator    common.Address
	Fees        ISourceOpStateManagerFeesData
}, error) {
	var out []interface{}
	err := _Srcstatemanager.contract.Call(opts, &out, "orderData", arg0)

	outstruct := new(struct {
		Fulfilled   bool
		Expiry      uint32
		OrderAmount *big.Int
		DestAddress common.Address
		Operator    common.Address
		Fees        ISourceOpStateManagerFeesData
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fulfilled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Expiry = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.OrderAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DestAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Operator = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Fees = *abi.ConvertType(out[5], new(ISourceOpStateManagerFeesData)).(*ISourceOpStateManagerFeesData)

	return *outstruct, err

}

// OrderData is a free data retrieval call binding the contract method 0xc4cd9ed3.
//
// Solidity: function orderData(bytes32 ) view returns(bool fulfilled, uint32 expiry, uint256 orderAmount, address destAddress, address operator, (uint256,uint256) fees)
func (_Srcstatemanager *SrcstatemanagerSession) OrderData(arg0 [32]byte) (struct {
	Fulfilled   bool
	Expiry      uint32
	OrderAmount *big.Int
	DestAddress common.Address
	Operator    common.Address
	Fees        ISourceOpStateManagerFeesData
}, error) {
	return _Srcstatemanager.Contract.OrderData(&_Srcstatemanager.CallOpts, arg0)
}

// OrderData is a free data retrieval call binding the contract method 0xc4cd9ed3.
//
// Solidity: function orderData(bytes32 ) view returns(bool fulfilled, uint32 expiry, uint256 orderAmount, address destAddress, address operator, (uint256,uint256) fees)
func (_Srcstatemanager *SrcstatemanagerCallerSession) OrderData(arg0 [32]byte) (struct {
	Fulfilled   bool
	Expiry      uint32
	OrderAmount *big.Int
	DestAddress common.Address
	Operator    common.Address
	Fees        ISourceOpStateManagerFeesData
}, error) {
	return _Srcstatemanager.Contract.OrderData(&_Srcstatemanager.CallOpts, arg0)
}

// CompleteOrder is a paid mutator transaction binding the contract method 0xddf5df45.
//
// Solidity: function completeOrder(bytes32 orderId) returns()
func (_Srcstatemanager *SrcstatemanagerTransactor) CompleteOrder(opts *bind.TransactOpts, orderId [32]byte) (*types.Transaction, error) {
	return _Srcstatemanager.contract.Transact(opts, "completeOrder", orderId)
}

// CompleteOrder is a paid mutator transaction binding the contract method 0xddf5df45.
//
// Solidity: function completeOrder(bytes32 orderId) returns()
func (_Srcstatemanager *SrcstatemanagerSession) CompleteOrder(orderId [32]byte) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.CompleteOrder(&_Srcstatemanager.TransactOpts, orderId)
}

// CompleteOrder is a paid mutator transaction binding the contract method 0xddf5df45.
//
// Solidity: function completeOrder(bytes32 orderId) returns()
func (_Srcstatemanager *SrcstatemanagerTransactorSession) CompleteOrder(orderId [32]byte) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.CompleteOrder(&_Srcstatemanager.TransactOpts, orderId)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xf2738b5b.
//
// Solidity: function createOrder(bytes32 orderId, uint32 expiry, uint256 orderAmount, address destAddress, address operator, uint256 operationFee, uint256 bridgeFee) returns()
func (_Srcstatemanager *SrcstatemanagerTransactor) CreateOrder(opts *bind.TransactOpts, orderId [32]byte, expiry uint32, orderAmount *big.Int, destAddress common.Address, operator common.Address, operationFee *big.Int, bridgeFee *big.Int) (*types.Transaction, error) {
	return _Srcstatemanager.contract.Transact(opts, "createOrder", orderId, expiry, orderAmount, destAddress, operator, operationFee, bridgeFee)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xf2738b5b.
//
// Solidity: function createOrder(bytes32 orderId, uint32 expiry, uint256 orderAmount, address destAddress, address operator, uint256 operationFee, uint256 bridgeFee) returns()
func (_Srcstatemanager *SrcstatemanagerSession) CreateOrder(orderId [32]byte, expiry uint32, orderAmount *big.Int, destAddress common.Address, operator common.Address, operationFee *big.Int, bridgeFee *big.Int) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.CreateOrder(&_Srcstatemanager.TransactOpts, orderId, expiry, orderAmount, destAddress, operator, operationFee, bridgeFee)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xf2738b5b.
//
// Solidity: function createOrder(bytes32 orderId, uint32 expiry, uint256 orderAmount, address destAddress, address operator, uint256 operationFee, uint256 bridgeFee) returns()
func (_Srcstatemanager *SrcstatemanagerTransactorSession) CreateOrder(orderId [32]byte, expiry uint32, orderAmount *big.Int, destAddress common.Address, operator common.Address, operationFee *big.Int, bridgeFee *big.Int) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.CreateOrder(&_Srcstatemanager.TransactOpts, orderId, expiry, orderAmount, destAddress, operator, operationFee, bridgeFee)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0xeedad66b.
//
// Solidity: function increaseStake(uint256 stakeAmount) returns()
func (_Srcstatemanager *SrcstatemanagerTransactor) IncreaseStake(opts *bind.TransactOpts, stakeAmount *big.Int) (*types.Transaction, error) {
	return _Srcstatemanager.contract.Transact(opts, "increaseStake", stakeAmount)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0xeedad66b.
//
// Solidity: function increaseStake(uint256 stakeAmount) returns()
func (_Srcstatemanager *SrcstatemanagerSession) IncreaseStake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.IncreaseStake(&_Srcstatemanager.TransactOpts, stakeAmount)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0xeedad66b.
//
// Solidity: function increaseStake(uint256 stakeAmount) returns()
func (_Srcstatemanager *SrcstatemanagerTransactorSession) IncreaseStake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.IncreaseStake(&_Srcstatemanager.TransactOpts, stakeAmount)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x4a122a3d.
//
// Solidity: function registerOperator(uint256 stakeAmount) returns()
func (_Srcstatemanager *SrcstatemanagerTransactor) RegisterOperator(opts *bind.TransactOpts, stakeAmount *big.Int) (*types.Transaction, error) {
	return _Srcstatemanager.contract.Transact(opts, "registerOperator", stakeAmount)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x4a122a3d.
//
// Solidity: function registerOperator(uint256 stakeAmount) returns()
func (_Srcstatemanager *SrcstatemanagerSession) RegisterOperator(stakeAmount *big.Int) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.RegisterOperator(&_Srcstatemanager.TransactOpts, stakeAmount)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x4a122a3d.
//
// Solidity: function registerOperator(uint256 stakeAmount) returns()
func (_Srcstatemanager *SrcstatemanagerTransactorSession) RegisterOperator(stakeAmount *big.Int) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.RegisterOperator(&_Srcstatemanager.TransactOpts, stakeAmount)
}

// SyncOperator is a paid mutator transaction binding the contract method 0x4a2d72ff.
//
// Solidity: function syncOperator(address operator, (uint256,uint256,bool) newOperatorData, bool deleteOperator) returns()
func (_Srcstatemanager *SrcstatemanagerTransactor) SyncOperator(opts *bind.TransactOpts, operator common.Address, newOperatorData ISourceOpStateManagerOperatorData, deleteOperator bool) (*types.Transaction, error) {
	return _Srcstatemanager.contract.Transact(opts, "syncOperator", operator, newOperatorData, deleteOperator)
}

// SyncOperator is a paid mutator transaction binding the contract method 0x4a2d72ff.
//
// Solidity: function syncOperator(address operator, (uint256,uint256,bool) newOperatorData, bool deleteOperator) returns()
func (_Srcstatemanager *SrcstatemanagerSession) SyncOperator(operator common.Address, newOperatorData ISourceOpStateManagerOperatorData, deleteOperator bool) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.SyncOperator(&_Srcstatemanager.TransactOpts, operator, newOperatorData, deleteOperator)
}

// SyncOperator is a paid mutator transaction binding the contract method 0x4a2d72ff.
//
// Solidity: function syncOperator(address operator, (uint256,uint256,bool) newOperatorData, bool deleteOperator) returns()
func (_Srcstatemanager *SrcstatemanagerTransactorSession) SyncOperator(operator common.Address, newOperatorData ISourceOpStateManagerOperatorData, deleteOperator bool) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.SyncOperator(&_Srcstatemanager.TransactOpts, operator, newOperatorData, deleteOperator)
}

// UpdateOperatorAllocation is a paid mutator transaction binding the contract method 0x48522af9.
//
// Solidity: function updateOperatorAllocation(address operator, uint256 holdingAmount, uint256 stakeAmount, bool init) returns()
func (_Srcstatemanager *SrcstatemanagerTransactor) UpdateOperatorAllocation(opts *bind.TransactOpts, operator common.Address, holdingAmount *big.Int, stakeAmount *big.Int, init bool) (*types.Transaction, error) {
	return _Srcstatemanager.contract.Transact(opts, "updateOperatorAllocation", operator, holdingAmount, stakeAmount, init)
}

// UpdateOperatorAllocation is a paid mutator transaction binding the contract method 0x48522af9.
//
// Solidity: function updateOperatorAllocation(address operator, uint256 holdingAmount, uint256 stakeAmount, bool init) returns()
func (_Srcstatemanager *SrcstatemanagerSession) UpdateOperatorAllocation(operator common.Address, holdingAmount *big.Int, stakeAmount *big.Int, init bool) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.UpdateOperatorAllocation(&_Srcstatemanager.TransactOpts, operator, holdingAmount, stakeAmount, init)
}

// UpdateOperatorAllocation is a paid mutator transaction binding the contract method 0x48522af9.
//
// Solidity: function updateOperatorAllocation(address operator, uint256 holdingAmount, uint256 stakeAmount, bool init) returns()
func (_Srcstatemanager *SrcstatemanagerTransactorSession) UpdateOperatorAllocation(operator common.Address, holdingAmount *big.Int, stakeAmount *big.Int, init bool) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.UpdateOperatorAllocation(&_Srcstatemanager.TransactOpts, operator, holdingAmount, stakeAmount, init)
}

// UpdatePendingRefunds is a paid mutator transaction binding the contract method 0x0e1aa5f5.
//
// Solidity: function updatePendingRefunds((uint256,address)[] fulfillerData, uint256 lpFees) returns()
func (_Srcstatemanager *SrcstatemanagerTransactor) UpdatePendingRefunds(opts *bind.TransactOpts, fulfillerData []ISourceOpStateManagerFulfillerData, lpFees *big.Int) (*types.Transaction, error) {
	return _Srcstatemanager.contract.Transact(opts, "updatePendingRefunds", fulfillerData, lpFees)
}

// UpdatePendingRefunds is a paid mutator transaction binding the contract method 0x0e1aa5f5.
//
// Solidity: function updatePendingRefunds((uint256,address)[] fulfillerData, uint256 lpFees) returns()
func (_Srcstatemanager *SrcstatemanagerSession) UpdatePendingRefunds(fulfillerData []ISourceOpStateManagerFulfillerData, lpFees *big.Int) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.UpdatePendingRefunds(&_Srcstatemanager.TransactOpts, fulfillerData, lpFees)
}

// UpdatePendingRefunds is a paid mutator transaction binding the contract method 0x0e1aa5f5.
//
// Solidity: function updatePendingRefunds((uint256,address)[] fulfillerData, uint256 lpFees) returns()
func (_Srcstatemanager *SrcstatemanagerTransactorSession) UpdatePendingRefunds(fulfillerData []ISourceOpStateManagerFulfillerData, lpFees *big.Int) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.UpdatePendingRefunds(&_Srcstatemanager.TransactOpts, fulfillerData, lpFees)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 withdrawAmount) returns()
func (_Srcstatemanager *SrcstatemanagerTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Srcstatemanager.contract.Transact(opts, "withdrawStake", withdrawAmount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 withdrawAmount) returns()
func (_Srcstatemanager *SrcstatemanagerSession) WithdrawStake(withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.WithdrawStake(&_Srcstatemanager.TransactOpts, withdrawAmount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 withdrawAmount) returns()
func (_Srcstatemanager *SrcstatemanagerTransactorSession) WithdrawStake(withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Srcstatemanager.Contract.WithdrawStake(&_Srcstatemanager.TransactOpts, withdrawAmount)
}

// SrcstatemanagerOperatorStakedIterator is returned from FilterOperatorStaked and is used to iterate over the raw logs and unpacked data for OperatorStaked events raised by the Srcstatemanager contract.
type SrcstatemanagerOperatorStakedIterator struct {
	Event *SrcstatemanagerOperatorStaked // Event containing the contract specifics and raw log

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
func (it *SrcstatemanagerOperatorStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SrcstatemanagerOperatorStaked)
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
		it.Event = new(SrcstatemanagerOperatorStaked)
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
func (it *SrcstatemanagerOperatorStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SrcstatemanagerOperatorStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SrcstatemanagerOperatorStaked represents a OperatorStaked event raised by the Srcstatemanager contract.
type SrcstatemanagerOperatorStaked struct {
	Operator    common.Address
	StakeAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorStaked is a free log retrieval operation binding the contract event 0x0c7daee0b531bc77f6a329375c8a6564320bf8a3b67ae6d0fe509096b450d3d1.
//
// Solidity: event OperatorStaked(address indexed operator, uint256 stakeAmount)
func (_Srcstatemanager *SrcstatemanagerFilterer) FilterOperatorStaked(opts *bind.FilterOpts, operator []common.Address) (*SrcstatemanagerOperatorStakedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Srcstatemanager.contract.FilterLogs(opts, "OperatorStaked", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SrcstatemanagerOperatorStakedIterator{contract: _Srcstatemanager.contract, event: "OperatorStaked", logs: logs, sub: sub}, nil
}

// WatchOperatorStaked is a free log subscription operation binding the contract event 0x0c7daee0b531bc77f6a329375c8a6564320bf8a3b67ae6d0fe509096b450d3d1.
//
// Solidity: event OperatorStaked(address indexed operator, uint256 stakeAmount)
func (_Srcstatemanager *SrcstatemanagerFilterer) WatchOperatorStaked(opts *bind.WatchOpts, sink chan<- *SrcstatemanagerOperatorStaked, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Srcstatemanager.contract.WatchLogs(opts, "OperatorStaked", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SrcstatemanagerOperatorStaked)
				if err := _Srcstatemanager.contract.UnpackLog(event, "OperatorStaked", log); err != nil {
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

// ParseOperatorStaked is a log parse operation binding the contract event 0x0c7daee0b531bc77f6a329375c8a6564320bf8a3b67ae6d0fe509096b450d3d1.
//
// Solidity: event OperatorStaked(address indexed operator, uint256 stakeAmount)
func (_Srcstatemanager *SrcstatemanagerFilterer) ParseOperatorStaked(log types.Log) (*SrcstatemanagerOperatorStaked, error) {
	event := new(SrcstatemanagerOperatorStaked)
	if err := _Srcstatemanager.contract.UnpackLog(event, "OperatorStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SrcstatemanagerOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the Srcstatemanager contract.
type SrcstatemanagerOrderCreatedIterator struct {
	Event *SrcstatemanagerOrderCreated // Event containing the contract specifics and raw log

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
func (it *SrcstatemanagerOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SrcstatemanagerOrderCreated)
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
		it.Event = new(SrcstatemanagerOrderCreated)
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
func (it *SrcstatemanagerOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SrcstatemanagerOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SrcstatemanagerOrderCreated represents a OrderCreated event raised by the Srcstatemanager contract.
type SrcstatemanagerOrderCreated struct {
	OrderId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x918554b6bd6e2895ce6553de5de0e1a69db5289aa0e4fe193a0dcd1f14347477.
//
// Solidity: event OrderCreated(bytes32 indexed orderId)
func (_Srcstatemanager *SrcstatemanagerFilterer) FilterOrderCreated(opts *bind.FilterOpts, orderId [][32]byte) (*SrcstatemanagerOrderCreatedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Srcstatemanager.contract.FilterLogs(opts, "OrderCreated", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &SrcstatemanagerOrderCreatedIterator{contract: _Srcstatemanager.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x918554b6bd6e2895ce6553de5de0e1a69db5289aa0e4fe193a0dcd1f14347477.
//
// Solidity: event OrderCreated(bytes32 indexed orderId)
func (_Srcstatemanager *SrcstatemanagerFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *SrcstatemanagerOrderCreated, orderId [][32]byte) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Srcstatemanager.contract.WatchLogs(opts, "OrderCreated", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SrcstatemanagerOrderCreated)
				if err := _Srcstatemanager.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0x918554b6bd6e2895ce6553de5de0e1a69db5289aa0e4fe193a0dcd1f14347477.
//
// Solidity: event OrderCreated(bytes32 indexed orderId)
func (_Srcstatemanager *SrcstatemanagerFilterer) ParseOrderCreated(log types.Log) (*SrcstatemanagerOrderCreated, error) {
	event := new(SrcstatemanagerOrderCreated)
	if err := _Srcstatemanager.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SrcstatemanagerOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the Srcstatemanager contract.
type SrcstatemanagerOrderFulfilledIterator struct {
	Event *SrcstatemanagerOrderFulfilled // Event containing the contract specifics and raw log

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
func (it *SrcstatemanagerOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SrcstatemanagerOrderFulfilled)
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
		it.Event = new(SrcstatemanagerOrderFulfilled)
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
func (it *SrcstatemanagerOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SrcstatemanagerOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SrcstatemanagerOrderFulfilled represents a OrderFulfilled event raised by the Srcstatemanager contract.
type SrcstatemanagerOrderFulfilled struct {
	OrderId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0x48224e98144457f56f46e15959589d8155a9e29ca79439ab9ba37f60561b5c56.
//
// Solidity: event OrderFulfilled(bytes32 indexed orderId)
func (_Srcstatemanager *SrcstatemanagerFilterer) FilterOrderFulfilled(opts *bind.FilterOpts, orderId [][32]byte) (*SrcstatemanagerOrderFulfilledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Srcstatemanager.contract.FilterLogs(opts, "OrderFulfilled", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &SrcstatemanagerOrderFulfilledIterator{contract: _Srcstatemanager.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0x48224e98144457f56f46e15959589d8155a9e29ca79439ab9ba37f60561b5c56.
//
// Solidity: event OrderFulfilled(bytes32 indexed orderId)
func (_Srcstatemanager *SrcstatemanagerFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *SrcstatemanagerOrderFulfilled, orderId [][32]byte) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Srcstatemanager.contract.WatchLogs(opts, "OrderFulfilled", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SrcstatemanagerOrderFulfilled)
				if err := _Srcstatemanager.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
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

// ParseOrderFulfilled is a log parse operation binding the contract event 0x48224e98144457f56f46e15959589d8155a9e29ca79439ab9ba37f60561b5c56.
//
// Solidity: event OrderFulfilled(bytes32 indexed orderId)
func (_Srcstatemanager *SrcstatemanagerFilterer) ParseOrderFulfilled(log types.Log) (*SrcstatemanagerOrderFulfilled, error) {
	event := new(SrcstatemanagerOrderFulfilled)
	if err := _Srcstatemanager.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SrcstatemanagerWithdrawalInitiatedIterator is returned from FilterWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for WithdrawalInitiated events raised by the Srcstatemanager contract.
type SrcstatemanagerWithdrawalInitiatedIterator struct {
	Event *SrcstatemanagerWithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *SrcstatemanagerWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SrcstatemanagerWithdrawalInitiated)
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
		it.Event = new(SrcstatemanagerWithdrawalInitiated)
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
func (it *SrcstatemanagerWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SrcstatemanagerWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SrcstatemanagerWithdrawalInitiated represents a WithdrawalInitiated event raised by the Srcstatemanager contract.
type SrcstatemanagerWithdrawalInitiated struct {
	Operator       common.Address
	WithdrawAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalInitiated is a free log retrieval operation binding the contract event 0x6d92f7d3303f995bf21956bb0c51b388bae348eaf45c23debd2cfa3fcd9ec646.
//
// Solidity: event WithdrawalInitiated(address indexed operator, uint256 withdrawAmount)
func (_Srcstatemanager *SrcstatemanagerFilterer) FilterWithdrawalInitiated(opts *bind.FilterOpts, operator []common.Address) (*SrcstatemanagerWithdrawalInitiatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Srcstatemanager.contract.FilterLogs(opts, "WithdrawalInitiated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SrcstatemanagerWithdrawalInitiatedIterator{contract: _Srcstatemanager.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalInitiated is a free log subscription operation binding the contract event 0x6d92f7d3303f995bf21956bb0c51b388bae348eaf45c23debd2cfa3fcd9ec646.
//
// Solidity: event WithdrawalInitiated(address indexed operator, uint256 withdrawAmount)
func (_Srcstatemanager *SrcstatemanagerFilterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *SrcstatemanagerWithdrawalInitiated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Srcstatemanager.contract.WatchLogs(opts, "WithdrawalInitiated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SrcstatemanagerWithdrawalInitiated)
				if err := _Srcstatemanager.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
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

// ParseWithdrawalInitiated is a log parse operation binding the contract event 0x6d92f7d3303f995bf21956bb0c51b388bae348eaf45c23debd2cfa3fcd9ec646.
//
// Solidity: event WithdrawalInitiated(address indexed operator, uint256 withdrawAmount)
func (_Srcstatemanager *SrcstatemanagerFilterer) ParseWithdrawalInitiated(log types.Log) (*SrcstatemanagerWithdrawalInitiated, error) {
	event := new(SrcstatemanagerWithdrawalInitiated)
	if err := _Srcstatemanager.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
