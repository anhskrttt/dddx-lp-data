// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gauge_factory

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

// GaugeFactoryMetaData contains all meta data concerning the GaugeFactory contract.
var GaugeFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bribe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ve\",\"type\":\"address\"}],\"name\":\"createGauge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bribe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"createGaugeSingle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last_gauge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GaugeFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use GaugeFactoryMetaData.ABI instead.
var GaugeFactoryABI = GaugeFactoryMetaData.ABI

// GaugeFactory is an auto generated Go binding around an Ethereum contract.
type GaugeFactory struct {
	GaugeFactoryCaller     // Read-only binding to the contract
	GaugeFactoryTransactor // Write-only binding to the contract
	GaugeFactoryFilterer   // Log filterer for contract events
}

// GaugeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type GaugeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GaugeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GaugeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GaugeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GaugeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GaugeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GaugeFactorySession struct {
	Contract     *GaugeFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GaugeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GaugeFactoryCallerSession struct {
	Contract *GaugeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// GaugeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GaugeFactoryTransactorSession struct {
	Contract     *GaugeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GaugeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type GaugeFactoryRaw struct {
	Contract *GaugeFactory // Generic contract binding to access the raw methods on
}

// GaugeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GaugeFactoryCallerRaw struct {
	Contract *GaugeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// GaugeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GaugeFactoryTransactorRaw struct {
	Contract *GaugeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGaugeFactory creates a new instance of GaugeFactory, bound to a specific deployed contract.
func NewGaugeFactory(address common.Address, backend bind.ContractBackend) (*GaugeFactory, error) {
	contract, err := bindGaugeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GaugeFactory{GaugeFactoryCaller: GaugeFactoryCaller{contract: contract}, GaugeFactoryTransactor: GaugeFactoryTransactor{contract: contract}, GaugeFactoryFilterer: GaugeFactoryFilterer{contract: contract}}, nil
}

// NewGaugeFactoryCaller creates a new read-only instance of GaugeFactory, bound to a specific deployed contract.
func NewGaugeFactoryCaller(address common.Address, caller bind.ContractCaller) (*GaugeFactoryCaller, error) {
	contract, err := bindGaugeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GaugeFactoryCaller{contract: contract}, nil
}

// NewGaugeFactoryTransactor creates a new write-only instance of GaugeFactory, bound to a specific deployed contract.
func NewGaugeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*GaugeFactoryTransactor, error) {
	contract, err := bindGaugeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GaugeFactoryTransactor{contract: contract}, nil
}

// NewGaugeFactoryFilterer creates a new log filterer instance of GaugeFactory, bound to a specific deployed contract.
func NewGaugeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*GaugeFactoryFilterer, error) {
	contract, err := bindGaugeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GaugeFactoryFilterer{contract: contract}, nil
}

// bindGaugeFactory binds a generic wrapper to an already deployed contract.
func bindGaugeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GaugeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GaugeFactory *GaugeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GaugeFactory.Contract.GaugeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GaugeFactory *GaugeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GaugeFactory.Contract.GaugeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GaugeFactory *GaugeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GaugeFactory.Contract.GaugeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GaugeFactory *GaugeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GaugeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GaugeFactory *GaugeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GaugeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GaugeFactory *GaugeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GaugeFactory.Contract.contract.Transact(opts, method, params...)
}

// LastGauge is a free data retrieval call binding the contract method 0x730a8bdb.
//
// Solidity: function last_gauge() view returns(address)
func (_GaugeFactory *GaugeFactoryCaller) LastGauge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GaugeFactory.contract.Call(opts, &out, "last_gauge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastGauge is a free data retrieval call binding the contract method 0x730a8bdb.
//
// Solidity: function last_gauge() view returns(address)
func (_GaugeFactory *GaugeFactorySession) LastGauge() (common.Address, error) {
	return _GaugeFactory.Contract.LastGauge(&_GaugeFactory.CallOpts)
}

// LastGauge is a free data retrieval call binding the contract method 0x730a8bdb.
//
// Solidity: function last_gauge() view returns(address)
func (_GaugeFactory *GaugeFactoryCallerSession) LastGauge() (common.Address, error) {
	return _GaugeFactory.Contract.LastGauge(&_GaugeFactory.CallOpts)
}

// CreateGauge is a paid mutator transaction binding the contract method 0x1c48e0fa.
//
// Solidity: function createGauge(address _pool, address _bribe, address _ve) returns(address)
func (_GaugeFactory *GaugeFactoryTransactor) CreateGauge(opts *bind.TransactOpts, _pool common.Address, _bribe common.Address, _ve common.Address) (*types.Transaction, error) {
	return _GaugeFactory.contract.Transact(opts, "createGauge", _pool, _bribe, _ve)
}

// CreateGauge is a paid mutator transaction binding the contract method 0x1c48e0fa.
//
// Solidity: function createGauge(address _pool, address _bribe, address _ve) returns(address)
func (_GaugeFactory *GaugeFactorySession) CreateGauge(_pool common.Address, _bribe common.Address, _ve common.Address) (*types.Transaction, error) {
	return _GaugeFactory.Contract.CreateGauge(&_GaugeFactory.TransactOpts, _pool, _bribe, _ve)
}

// CreateGauge is a paid mutator transaction binding the contract method 0x1c48e0fa.
//
// Solidity: function createGauge(address _pool, address _bribe, address _ve) returns(address)
func (_GaugeFactory *GaugeFactoryTransactorSession) CreateGauge(_pool common.Address, _bribe common.Address, _ve common.Address) (*types.Transaction, error) {
	return _GaugeFactory.Contract.CreateGauge(&_GaugeFactory.TransactOpts, _pool, _bribe, _ve)
}

// CreateGaugeSingle is a paid mutator transaction binding the contract method 0xaf4db2ec.
//
// Solidity: function createGaugeSingle(address _pool, address _bribe, address _ve, address _voter) returns(address)
func (_GaugeFactory *GaugeFactoryTransactor) CreateGaugeSingle(opts *bind.TransactOpts, _pool common.Address, _bribe common.Address, _ve common.Address, _voter common.Address) (*types.Transaction, error) {
	return _GaugeFactory.contract.Transact(opts, "createGaugeSingle", _pool, _bribe, _ve, _voter)
}

// CreateGaugeSingle is a paid mutator transaction binding the contract method 0xaf4db2ec.
//
// Solidity: function createGaugeSingle(address _pool, address _bribe, address _ve, address _voter) returns(address)
func (_GaugeFactory *GaugeFactorySession) CreateGaugeSingle(_pool common.Address, _bribe common.Address, _ve common.Address, _voter common.Address) (*types.Transaction, error) {
	return _GaugeFactory.Contract.CreateGaugeSingle(&_GaugeFactory.TransactOpts, _pool, _bribe, _ve, _voter)
}

// CreateGaugeSingle is a paid mutator transaction binding the contract method 0xaf4db2ec.
//
// Solidity: function createGaugeSingle(address _pool, address _bribe, address _ve, address _voter) returns(address)
func (_GaugeFactory *GaugeFactoryTransactorSession) CreateGaugeSingle(_pool common.Address, _bribe common.Address, _ve common.Address, _voter common.Address) (*types.Transaction, error) {
	return _GaugeFactory.Contract.CreateGaugeSingle(&_GaugeFactory.TransactOpts, _pool, _bribe, _ve, _voter)
}
