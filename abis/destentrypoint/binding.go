// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package destentrypoint

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

// IEntityFeesData is an auto generated low-level Go binding around an user-defined struct.
type IEntityFeesData struct {
	OperationFee *big.Int
	BridgeFee    *big.Int
}

// IEntityFulfillerData is an auto generated low-level Go binding around an user-defined struct.
type IEntityFulfillerData struct {
	FulfillAmount *big.Int
	Fulfiller     common.Address
}

// IEntityOrderData is an auto generated low-level Go binding around an user-defined struct.
type IEntityOrderData struct {
	Fulfilled   bool
	Expiry      uint32
	OrderAmount *big.Int
	DestAddress common.Address
	Operator    common.Address
	Fees        IEntityFeesData
}

// MessagingFee is an auto generated low-level Go binding around an user-defined struct.
type MessagingFee struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}

// DestentrypointMetaData contains all meta data concerning the Destentrypoint contract.
var DestentrypointMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lzEndpoint\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_sourceChainEid\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddressRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEndpointCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"optionType\",\"type\":\"uint16\"}],\"name\":\"InvalidOptionType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LzTokenUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"NoPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"NotEnoughNative\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NotGovernance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"OnlyPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"}],\"name\":\"FulfilledOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"name\":\"PeerSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_EID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressRegistry\",\"outputs\":[{\"internalType\":\"contractAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpointV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"expiry\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"orderAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"operationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"}],\"internalType\":\"structIEntity.FeesData\",\"name\":\"fees\",\"type\":\"tuple\"}],\"internalType\":\"structIEntity.OrderData\",\"name\":\"orderData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fulfillAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fulfiller\",\"type\":\"address\"}],\"internalType\":\"structIEntity.FulfillerData[]\",\"name\":\"fulfillerData\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"operatorSignatures\",\"type\":\"bytes[]\"}],\"name\":\"fulfillOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"payInLzToken\",\"type\":\"bool\"}],\"name\":\"lzQuote\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"fee\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oAppVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"senderVersion\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receiverVersion\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"peers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"setPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DestentrypointABI is the input ABI used to generate the binding from.
// Deprecated: Use DestentrypointMetaData.ABI instead.
var DestentrypointABI = DestentrypointMetaData.ABI

// Destentrypoint is an auto generated Go binding around an Ethereum contract.
type Destentrypoint struct {
	DestentrypointCaller     // Read-only binding to the contract
	DestentrypointTransactor // Write-only binding to the contract
	DestentrypointFilterer   // Log filterer for contract events
}

// DestentrypointCaller is an auto generated read-only Go binding around an Ethereum contract.
type DestentrypointCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestentrypointTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DestentrypointTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestentrypointFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DestentrypointFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestentrypointSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DestentrypointSession struct {
	Contract     *Destentrypoint   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DestentrypointCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DestentrypointCallerSession struct {
	Contract *DestentrypointCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DestentrypointTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DestentrypointTransactorSession struct {
	Contract     *DestentrypointTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DestentrypointRaw is an auto generated low-level Go binding around an Ethereum contract.
type DestentrypointRaw struct {
	Contract *Destentrypoint // Generic contract binding to access the raw methods on
}

// DestentrypointCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DestentrypointCallerRaw struct {
	Contract *DestentrypointCaller // Generic read-only contract binding to access the raw methods on
}

// DestentrypointTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DestentrypointTransactorRaw struct {
	Contract *DestentrypointTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDestentrypoint creates a new instance of Destentrypoint, bound to a specific deployed contract.
func NewDestentrypoint(address common.Address, backend bind.ContractBackend) (*Destentrypoint, error) {
	contract, err := bindDestentrypoint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Destentrypoint{DestentrypointCaller: DestentrypointCaller{contract: contract}, DestentrypointTransactor: DestentrypointTransactor{contract: contract}, DestentrypointFilterer: DestentrypointFilterer{contract: contract}}, nil
}

// NewDestentrypointCaller creates a new read-only instance of Destentrypoint, bound to a specific deployed contract.
func NewDestentrypointCaller(address common.Address, caller bind.ContractCaller) (*DestentrypointCaller, error) {
	contract, err := bindDestentrypoint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DestentrypointCaller{contract: contract}, nil
}

// NewDestentrypointTransactor creates a new write-only instance of Destentrypoint, bound to a specific deployed contract.
func NewDestentrypointTransactor(address common.Address, transactor bind.ContractTransactor) (*DestentrypointTransactor, error) {
	contract, err := bindDestentrypoint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DestentrypointTransactor{contract: contract}, nil
}

// NewDestentrypointFilterer creates a new log filterer instance of Destentrypoint, bound to a specific deployed contract.
func NewDestentrypointFilterer(address common.Address, filterer bind.ContractFilterer) (*DestentrypointFilterer, error) {
	contract, err := bindDestentrypoint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DestentrypointFilterer{contract: contract}, nil
}

// bindDestentrypoint binds a generic wrapper to an already deployed contract.
func bindDestentrypoint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DestentrypointMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destentrypoint *DestentrypointRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Destentrypoint.Contract.DestentrypointCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destentrypoint *DestentrypointRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destentrypoint.Contract.DestentrypointTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destentrypoint *DestentrypointRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destentrypoint.Contract.DestentrypointTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destentrypoint *DestentrypointCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Destentrypoint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destentrypoint *DestentrypointTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destentrypoint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destentrypoint *DestentrypointTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destentrypoint.Contract.contract.Transact(opts, method, params...)
}

// SOURCECHAINEID is a free data retrieval call binding the contract method 0x6099a385.
//
// Solidity: function SOURCE_CHAIN_EID() view returns(uint32)
func (_Destentrypoint *DestentrypointCaller) SOURCECHAINEID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Destentrypoint.contract.Call(opts, &out, "SOURCE_CHAIN_EID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SOURCECHAINEID is a free data retrieval call binding the contract method 0x6099a385.
//
// Solidity: function SOURCE_CHAIN_EID() view returns(uint32)
func (_Destentrypoint *DestentrypointSession) SOURCECHAINEID() (uint32, error) {
	return _Destentrypoint.Contract.SOURCECHAINEID(&_Destentrypoint.CallOpts)
}

// SOURCECHAINEID is a free data retrieval call binding the contract method 0x6099a385.
//
// Solidity: function SOURCE_CHAIN_EID() view returns(uint32)
func (_Destentrypoint *DestentrypointCallerSession) SOURCECHAINEID() (uint32, error) {
	return _Destentrypoint.Contract.SOURCECHAINEID(&_Destentrypoint.CallOpts)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_Destentrypoint *DestentrypointCaller) AddressRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Destentrypoint.contract.Call(opts, &out, "addressRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_Destentrypoint *DestentrypointSession) AddressRegistry() (common.Address, error) {
	return _Destentrypoint.Contract.AddressRegistry(&_Destentrypoint.CallOpts)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_Destentrypoint *DestentrypointCallerSession) AddressRegistry() (common.Address, error) {
	return _Destentrypoint.Contract.AddressRegistry(&_Destentrypoint.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Destentrypoint *DestentrypointCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Destentrypoint.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Destentrypoint *DestentrypointSession) Endpoint() (common.Address, error) {
	return _Destentrypoint.Contract.Endpoint(&_Destentrypoint.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Destentrypoint *DestentrypointCallerSession) Endpoint() (common.Address, error) {
	return _Destentrypoint.Contract.Endpoint(&_Destentrypoint.CallOpts)
}

// LzQuote is a free data retrieval call binding the contract method 0x686be10e.
//
// Solidity: function lzQuote(uint32 dstEid, bytes payload, bool payInLzToken) view returns((uint256,uint256) fee)
func (_Destentrypoint *DestentrypointCaller) LzQuote(opts *bind.CallOpts, dstEid uint32, payload []byte, payInLzToken bool) (MessagingFee, error) {
	var out []interface{}
	err := _Destentrypoint.contract.Call(opts, &out, "lzQuote", dstEid, payload, payInLzToken)

	if err != nil {
		return *new(MessagingFee), err
	}

	out0 := *abi.ConvertType(out[0], new(MessagingFee)).(*MessagingFee)

	return out0, err

}

// LzQuote is a free data retrieval call binding the contract method 0x686be10e.
//
// Solidity: function lzQuote(uint32 dstEid, bytes payload, bool payInLzToken) view returns((uint256,uint256) fee)
func (_Destentrypoint *DestentrypointSession) LzQuote(dstEid uint32, payload []byte, payInLzToken bool) (MessagingFee, error) {
	return _Destentrypoint.Contract.LzQuote(&_Destentrypoint.CallOpts, dstEid, payload, payInLzToken)
}

// LzQuote is a free data retrieval call binding the contract method 0x686be10e.
//
// Solidity: function lzQuote(uint32 dstEid, bytes payload, bool payInLzToken) view returns((uint256,uint256) fee)
func (_Destentrypoint *DestentrypointCallerSession) LzQuote(dstEid uint32, payload []byte, payInLzToken bool) (MessagingFee, error) {
	return _Destentrypoint.Contract.LzQuote(&_Destentrypoint.CallOpts, dstEid, payload, payInLzToken)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() view returns(uint64 senderVersion, uint64 receiverVersion)
func (_Destentrypoint *DestentrypointCaller) OAppVersion(opts *bind.CallOpts) (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	var out []interface{}
	err := _Destentrypoint.contract.Call(opts, &out, "oAppVersion")

	outstruct := new(struct {
		SenderVersion   uint64
		ReceiverVersion uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SenderVersion = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ReceiverVersion = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() view returns(uint64 senderVersion, uint64 receiverVersion)
func (_Destentrypoint *DestentrypointSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Destentrypoint.Contract.OAppVersion(&_Destentrypoint.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() view returns(uint64 senderVersion, uint64 receiverVersion)
func (_Destentrypoint *DestentrypointCallerSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Destentrypoint.Contract.OAppVersion(&_Destentrypoint.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Destentrypoint *DestentrypointCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Destentrypoint.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Destentrypoint *DestentrypointSession) Owner() (common.Address, error) {
	return _Destentrypoint.Contract.Owner(&_Destentrypoint.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Destentrypoint *DestentrypointCallerSession) Owner() (common.Address, error) {
	return _Destentrypoint.Contract.Owner(&_Destentrypoint.CallOpts)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Destentrypoint *DestentrypointCaller) Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error) {
	var out []interface{}
	err := _Destentrypoint.contract.Call(opts, &out, "peers", eid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Destentrypoint *DestentrypointSession) Peers(eid uint32) ([32]byte, error) {
	return _Destentrypoint.Contract.Peers(&_Destentrypoint.CallOpts, eid)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Destentrypoint *DestentrypointCallerSession) Peers(eid uint32) ([32]byte, error) {
	return _Destentrypoint.Contract.Peers(&_Destentrypoint.CallOpts, eid)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xc2208e9b.
//
// Solidity: function fulfillOrder(bytes32 orderId, (bool,uint32,uint256,address,address,(uint256,uint256)) orderData, (uint256,address)[] fulfillerData, bytes[] operatorSignatures) returns()
func (_Destentrypoint *DestentrypointTransactor) FulfillOrder(opts *bind.TransactOpts, orderId [32]byte, orderData IEntityOrderData, fulfillerData []IEntityFulfillerData, operatorSignatures [][]byte) (*types.Transaction, error) {
	return _Destentrypoint.contract.Transact(opts, "fulfillOrder", orderId, orderData, fulfillerData, operatorSignatures)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xc2208e9b.
//
// Solidity: function fulfillOrder(bytes32 orderId, (bool,uint32,uint256,address,address,(uint256,uint256)) orderData, (uint256,address)[] fulfillerData, bytes[] operatorSignatures) returns()
func (_Destentrypoint *DestentrypointSession) FulfillOrder(orderId [32]byte, orderData IEntityOrderData, fulfillerData []IEntityFulfillerData, operatorSignatures [][]byte) (*types.Transaction, error) {
	return _Destentrypoint.Contract.FulfillOrder(&_Destentrypoint.TransactOpts, orderId, orderData, fulfillerData, operatorSignatures)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xc2208e9b.
//
// Solidity: function fulfillOrder(bytes32 orderId, (bool,uint32,uint256,address,address,(uint256,uint256)) orderData, (uint256,address)[] fulfillerData, bytes[] operatorSignatures) returns()
func (_Destentrypoint *DestentrypointTransactorSession) FulfillOrder(orderId [32]byte, orderData IEntityOrderData, fulfillerData []IEntityFulfillerData, operatorSignatures [][]byte) (*types.Transaction, error) {
	return _Destentrypoint.Contract.FulfillOrder(&_Destentrypoint.TransactOpts, orderId, orderData, fulfillerData, operatorSignatures)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destentrypoint *DestentrypointTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destentrypoint.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destentrypoint *DestentrypointSession) RenounceOwnership() (*types.Transaction, error) {
	return _Destentrypoint.Contract.RenounceOwnership(&_Destentrypoint.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destentrypoint *DestentrypointTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Destentrypoint.Contract.RenounceOwnership(&_Destentrypoint.TransactOpts)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Destentrypoint *DestentrypointTransactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _Destentrypoint.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Destentrypoint *DestentrypointSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Destentrypoint.Contract.SetDelegate(&_Destentrypoint.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Destentrypoint *DestentrypointTransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Destentrypoint.Contract.SetDelegate(&_Destentrypoint.TransactOpts, _delegate)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Destentrypoint *DestentrypointTransactor) SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Destentrypoint.contract.Transact(opts, "setPeer", _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Destentrypoint *DestentrypointSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Destentrypoint.Contract.SetPeer(&_Destentrypoint.TransactOpts, _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Destentrypoint *DestentrypointTransactorSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Destentrypoint.Contract.SetPeer(&_Destentrypoint.TransactOpts, _eid, _peer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Destentrypoint *DestentrypointTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Destentrypoint.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Destentrypoint *DestentrypointSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Destentrypoint.Contract.TransferOwnership(&_Destentrypoint.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Destentrypoint *DestentrypointTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Destentrypoint.Contract.TransferOwnership(&_Destentrypoint.TransactOpts, newOwner)
}

// DestentrypointFulfilledOrderIterator is returned from FilterFulfilledOrder and is used to iterate over the raw logs and unpacked data for FulfilledOrder events raised by the Destentrypoint contract.
type DestentrypointFulfilledOrderIterator struct {
	Event *DestentrypointFulfilledOrder // Event containing the contract specifics and raw log

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
func (it *DestentrypointFulfilledOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestentrypointFulfilledOrder)
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
		it.Event = new(DestentrypointFulfilledOrder)
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
func (it *DestentrypointFulfilledOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestentrypointFulfilledOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestentrypointFulfilledOrder represents a FulfilledOrder event raised by the Destentrypoint contract.
type DestentrypointFulfilledOrder struct {
	OrderId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFulfilledOrder is a free log retrieval operation binding the contract event 0x8028dba6722378cccc0fa0893fdab0a32ec1f24adbf6d78a0694c283ae767949.
//
// Solidity: event FulfilledOrder(bytes32 indexed orderId)
func (_Destentrypoint *DestentrypointFilterer) FilterFulfilledOrder(opts *bind.FilterOpts, orderId [][32]byte) (*DestentrypointFulfilledOrderIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Destentrypoint.contract.FilterLogs(opts, "FulfilledOrder", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &DestentrypointFulfilledOrderIterator{contract: _Destentrypoint.contract, event: "FulfilledOrder", logs: logs, sub: sub}, nil
}

// WatchFulfilledOrder is a free log subscription operation binding the contract event 0x8028dba6722378cccc0fa0893fdab0a32ec1f24adbf6d78a0694c283ae767949.
//
// Solidity: event FulfilledOrder(bytes32 indexed orderId)
func (_Destentrypoint *DestentrypointFilterer) WatchFulfilledOrder(opts *bind.WatchOpts, sink chan<- *DestentrypointFulfilledOrder, orderId [][32]byte) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Destentrypoint.contract.WatchLogs(opts, "FulfilledOrder", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestentrypointFulfilledOrder)
				if err := _Destentrypoint.contract.UnpackLog(event, "FulfilledOrder", log); err != nil {
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

// ParseFulfilledOrder is a log parse operation binding the contract event 0x8028dba6722378cccc0fa0893fdab0a32ec1f24adbf6d78a0694c283ae767949.
//
// Solidity: event FulfilledOrder(bytes32 indexed orderId)
func (_Destentrypoint *DestentrypointFilterer) ParseFulfilledOrder(log types.Log) (*DestentrypointFulfilledOrder, error) {
	event := new(DestentrypointFulfilledOrder)
	if err := _Destentrypoint.contract.UnpackLog(event, "FulfilledOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestentrypointOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Destentrypoint contract.
type DestentrypointOwnershipTransferredIterator struct {
	Event *DestentrypointOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DestentrypointOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestentrypointOwnershipTransferred)
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
		it.Event = new(DestentrypointOwnershipTransferred)
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
func (it *DestentrypointOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestentrypointOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestentrypointOwnershipTransferred represents a OwnershipTransferred event raised by the Destentrypoint contract.
type DestentrypointOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Destentrypoint *DestentrypointFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DestentrypointOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Destentrypoint.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DestentrypointOwnershipTransferredIterator{contract: _Destentrypoint.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Destentrypoint *DestentrypointFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DestentrypointOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Destentrypoint.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestentrypointOwnershipTransferred)
				if err := _Destentrypoint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Destentrypoint *DestentrypointFilterer) ParseOwnershipTransferred(log types.Log) (*DestentrypointOwnershipTransferred, error) {
	event := new(DestentrypointOwnershipTransferred)
	if err := _Destentrypoint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestentrypointPeerSetIterator is returned from FilterPeerSet and is used to iterate over the raw logs and unpacked data for PeerSet events raised by the Destentrypoint contract.
type DestentrypointPeerSetIterator struct {
	Event *DestentrypointPeerSet // Event containing the contract specifics and raw log

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
func (it *DestentrypointPeerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestentrypointPeerSet)
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
		it.Event = new(DestentrypointPeerSet)
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
func (it *DestentrypointPeerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestentrypointPeerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestentrypointPeerSet represents a PeerSet event raised by the Destentrypoint contract.
type DestentrypointPeerSet struct {
	Eid  uint32
	Peer [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPeerSet is a free log retrieval operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Destentrypoint *DestentrypointFilterer) FilterPeerSet(opts *bind.FilterOpts) (*DestentrypointPeerSetIterator, error) {

	logs, sub, err := _Destentrypoint.contract.FilterLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return &DestentrypointPeerSetIterator{contract: _Destentrypoint.contract, event: "PeerSet", logs: logs, sub: sub}, nil
}

// WatchPeerSet is a free log subscription operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Destentrypoint *DestentrypointFilterer) WatchPeerSet(opts *bind.WatchOpts, sink chan<- *DestentrypointPeerSet) (event.Subscription, error) {

	logs, sub, err := _Destentrypoint.contract.WatchLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestentrypointPeerSet)
				if err := _Destentrypoint.contract.UnpackLog(event, "PeerSet", log); err != nil {
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

// ParsePeerSet is a log parse operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Destentrypoint *DestentrypointFilterer) ParsePeerSet(log types.Log) (*DestentrypointPeerSet, error) {
	event := new(DestentrypointPeerSet)
	if err := _Destentrypoint.contract.UnpackLog(event, "PeerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
