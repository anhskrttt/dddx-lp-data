// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory

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

// FactoryMetaData contains all meta data concerning the Factory contract.
var FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"AcceptPauser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingPauser\",\"type\":\"address\"}],\"name\":\"SetPauser\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitializable\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isPair\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairCodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingPauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"setPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use FactoryMetaData.ABI instead.
var FactoryABI = FactoryMetaData.ABI

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Factory *FactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Factory *FactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Factory.Contract.AllPairs(&_Factory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Factory *FactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Factory.Contract.AllPairs(&_Factory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Factory *FactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Factory *FactorySession) AllPairsLength() (*big.Int, error) {
	return _Factory.Contract.AllPairsLength(&_Factory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Factory *FactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _Factory.Contract.AllPairsLength(&_Factory.CallOpts)
}

// GetInitializable is a free data retrieval call binding the contract method 0xeb13c4cf.
//
// Solidity: function getInitializable() view returns(address, address, bool)
func (_Factory *FactoryCaller) GetInitializable(opts *bind.CallOpts) (common.Address, common.Address, bool, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "getInitializable")

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetInitializable is a free data retrieval call binding the contract method 0xeb13c4cf.
//
// Solidity: function getInitializable() view returns(address, address, bool)
func (_Factory *FactorySession) GetInitializable() (common.Address, common.Address, bool, error) {
	return _Factory.Contract.GetInitializable(&_Factory.CallOpts)
}

// GetInitializable is a free data retrieval call binding the contract method 0xeb13c4cf.
//
// Solidity: function getInitializable() view returns(address, address, bool)
func (_Factory *FactoryCallerSession) GetInitializable() (common.Address, common.Address, bool, error) {
	return _Factory.Contract.GetInitializable(&_Factory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0x6801cc30.
//
// Solidity: function getPair(address , address , bool ) view returns(address)
func (_Factory *FactoryCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 bool) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "getPair", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0x6801cc30.
//
// Solidity: function getPair(address , address , bool ) view returns(address)
func (_Factory *FactorySession) GetPair(arg0 common.Address, arg1 common.Address, arg2 bool) (common.Address, error) {
	return _Factory.Contract.GetPair(&_Factory.CallOpts, arg0, arg1, arg2)
}

// GetPair is a free data retrieval call binding the contract method 0x6801cc30.
//
// Solidity: function getPair(address , address , bool ) view returns(address)
func (_Factory *FactoryCallerSession) GetPair(arg0 common.Address, arg1 common.Address, arg2 bool) (common.Address, error) {
	return _Factory.Contract.GetPair(&_Factory.CallOpts, arg0, arg1, arg2)
}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address ) view returns(bool)
func (_Factory *FactoryCaller) IsPair(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "isPair", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address ) view returns(bool)
func (_Factory *FactorySession) IsPair(arg0 common.Address) (bool, error) {
	return _Factory.Contract.IsPair(&_Factory.CallOpts, arg0)
}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address ) view returns(bool)
func (_Factory *FactoryCallerSession) IsPair(arg0 common.Address) (bool, error) {
	return _Factory.Contract.IsPair(&_Factory.CallOpts, arg0)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Factory *FactoryCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Factory *FactorySession) IsPaused() (bool, error) {
	return _Factory.Contract.IsPaused(&_Factory.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Factory *FactoryCallerSession) IsPaused() (bool, error) {
	return _Factory.Contract.IsPaused(&_Factory.CallOpts)
}

// PairCodeHash is a free data retrieval call binding the contract method 0x9aab9248.
//
// Solidity: function pairCodeHash() pure returns(bytes32)
func (_Factory *FactoryCaller) PairCodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "pairCodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PairCodeHash is a free data retrieval call binding the contract method 0x9aab9248.
//
// Solidity: function pairCodeHash() pure returns(bytes32)
func (_Factory *FactorySession) PairCodeHash() ([32]byte, error) {
	return _Factory.Contract.PairCodeHash(&_Factory.CallOpts)
}

// PairCodeHash is a free data retrieval call binding the contract method 0x9aab9248.
//
// Solidity: function pairCodeHash() pure returns(bytes32)
func (_Factory *FactoryCallerSession) PairCodeHash() ([32]byte, error) {
	return _Factory.Contract.PairCodeHash(&_Factory.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Factory *FactoryCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Factory *FactorySession) Pauser() (common.Address, error) {
	return _Factory.Contract.Pauser(&_Factory.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Factory *FactoryCallerSession) Pauser() (common.Address, error) {
	return _Factory.Contract.Pauser(&_Factory.CallOpts)
}

// PendingPauser is a free data retrieval call binding the contract method 0x9a7165e4.
//
// Solidity: function pendingPauser() view returns(address)
func (_Factory *FactoryCaller) PendingPauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "pendingPauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingPauser is a free data retrieval call binding the contract method 0x9a7165e4.
//
// Solidity: function pendingPauser() view returns(address)
func (_Factory *FactorySession) PendingPauser() (common.Address, error) {
	return _Factory.Contract.PendingPauser(&_Factory.CallOpts)
}

// PendingPauser is a free data retrieval call binding the contract method 0x9a7165e4.
//
// Solidity: function pendingPauser() view returns(address)
func (_Factory *FactoryCallerSession) PendingPauser() (common.Address, error) {
	return _Factory.Contract.PendingPauser(&_Factory.CallOpts)
}

// AcceptPauser is a paid mutator transaction binding the contract method 0x167a6f90.
//
// Solidity: function acceptPauser() returns()
func (_Factory *FactoryTransactor) AcceptPauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "acceptPauser")
}

// AcceptPauser is a paid mutator transaction binding the contract method 0x167a6f90.
//
// Solidity: function acceptPauser() returns()
func (_Factory *FactorySession) AcceptPauser() (*types.Transaction, error) {
	return _Factory.Contract.AcceptPauser(&_Factory.TransactOpts)
}

// AcceptPauser is a paid mutator transaction binding the contract method 0x167a6f90.
//
// Solidity: function acceptPauser() returns()
func (_Factory *FactoryTransactorSession) AcceptPauser() (*types.Transaction, error) {
	return _Factory.Contract.AcceptPauser(&_Factory.TransactOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0x82dfdce4.
//
// Solidity: function createPair(address tokenA, address tokenB, bool stable) returns(address pair)
func (_Factory *FactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createPair", tokenA, tokenB, stable)
}

// CreatePair is a paid mutator transaction binding the contract method 0x82dfdce4.
//
// Solidity: function createPair(address tokenA, address tokenB, bool stable) returns(address pair)
func (_Factory *FactorySession) CreatePair(tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _Factory.Contract.CreatePair(&_Factory.TransactOpts, tokenA, tokenB, stable)
}

// CreatePair is a paid mutator transaction binding the contract method 0x82dfdce4.
//
// Solidity: function createPair(address tokenA, address tokenB, bool stable) returns(address pair)
func (_Factory *FactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _Factory.Contract.CreatePair(&_Factory.TransactOpts, tokenA, tokenB, stable)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _state) returns()
func (_Factory *FactoryTransactor) SetPause(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setPause", _state)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _state) returns()
func (_Factory *FactorySession) SetPause(_state bool) (*types.Transaction, error) {
	return _Factory.Contract.SetPause(&_Factory.TransactOpts, _state)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _state) returns()
func (_Factory *FactoryTransactorSession) SetPause(_state bool) (*types.Transaction, error) {
	return _Factory.Contract.SetPause(&_Factory.TransactOpts, _state)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_Factory *FactoryTransactor) SetPauser(opts *bind.TransactOpts, _pauser common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setPauser", _pauser)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_Factory *FactorySession) SetPauser(_pauser common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetPauser(&_Factory.TransactOpts, _pauser)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_Factory *FactoryTransactorSession) SetPauser(_pauser common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetPauser(&_Factory.TransactOpts, _pauser)
}

// FactoryAcceptPauserIterator is returned from FilterAcceptPauser and is used to iterate over the raw logs and unpacked data for AcceptPauser events raised by the Factory contract.
type FactoryAcceptPauserIterator struct {
	Event *FactoryAcceptPauser // Event containing the contract specifics and raw log

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
func (it *FactoryAcceptPauserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryAcceptPauser)
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
		it.Event = new(FactoryAcceptPauser)
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
func (it *FactoryAcceptPauserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryAcceptPauserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryAcceptPauser represents a AcceptPauser event raised by the Factory contract.
type FactoryAcceptPauser struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAcceptPauser is a free log retrieval operation binding the contract event 0xa93c64beb9f312362ab29765df164e2454814a3ea1a2f2962fd38a96a7792e4f.
//
// Solidity: event AcceptPauser(address indexed pauser)
func (_Factory *FactoryFilterer) FilterAcceptPauser(opts *bind.FilterOpts, pauser []common.Address) (*FactoryAcceptPauserIterator, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "AcceptPauser", pauserRule)
	if err != nil {
		return nil, err
	}
	return &FactoryAcceptPauserIterator{contract: _Factory.contract, event: "AcceptPauser", logs: logs, sub: sub}, nil
}

// WatchAcceptPauser is a free log subscription operation binding the contract event 0xa93c64beb9f312362ab29765df164e2454814a3ea1a2f2962fd38a96a7792e4f.
//
// Solidity: event AcceptPauser(address indexed pauser)
func (_Factory *FactoryFilterer) WatchAcceptPauser(opts *bind.WatchOpts, sink chan<- *FactoryAcceptPauser, pauser []common.Address) (event.Subscription, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "AcceptPauser", pauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryAcceptPauser)
				if err := _Factory.contract.UnpackLog(event, "AcceptPauser", log); err != nil {
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

// ParseAcceptPauser is a log parse operation binding the contract event 0xa93c64beb9f312362ab29765df164e2454814a3ea1a2f2962fd38a96a7792e4f.
//
// Solidity: event AcceptPauser(address indexed pauser)
func (_Factory *FactoryFilterer) ParseAcceptPauser(log types.Log) (*FactoryAcceptPauser, error) {
	event := new(FactoryAcceptPauser)
	if err := _Factory.contract.UnpackLog(event, "AcceptPauser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the Factory contract.
type FactoryPairCreatedIterator struct {
	Event *FactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *FactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryPairCreated)
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
		it.Event = new(FactoryPairCreated)
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
func (it *FactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryPairCreated represents a PairCreated event raised by the Factory contract.
type FactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Stable bool
	Pair   common.Address
	Arg4   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0xc4805696c66d7cf352fc1d6bb633ad5ee82f6cb577c453024b6e0eb8306c6fc9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, bool stable, address pair, uint256 arg4)
func (_Factory *FactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*FactoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &FactoryPairCreatedIterator{contract: _Factory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0xc4805696c66d7cf352fc1d6bb633ad5ee82f6cb577c453024b6e0eb8306c6fc9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, bool stable, address pair, uint256 arg4)
func (_Factory *FactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *FactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryPairCreated)
				if err := _Factory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0xc4805696c66d7cf352fc1d6bb633ad5ee82f6cb577c453024b6e0eb8306c6fc9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, bool stable, address pair, uint256 arg4)
func (_Factory *FactoryFilterer) ParsePairCreated(log types.Log) (*FactoryPairCreated, error) {
	event := new(FactoryPairCreated)
	if err := _Factory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactorySetPauserIterator is returned from FilterSetPauser and is used to iterate over the raw logs and unpacked data for SetPauser events raised by the Factory contract.
type FactorySetPauserIterator struct {
	Event *FactorySetPauser // Event containing the contract specifics and raw log

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
func (it *FactorySetPauserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactorySetPauser)
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
		it.Event = new(FactorySetPauser)
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
func (it *FactorySetPauserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactorySetPauserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactorySetPauser represents a SetPauser event raised by the Factory contract.
type FactorySetPauser struct {
	PendingPauser common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetPauser is a free log retrieval operation binding the contract event 0xe02efb9e8f0fc21546730ab32d594f62d586e1bbb15bb5045edd0b1878a77b35.
//
// Solidity: event SetPauser(address indexed pendingPauser)
func (_Factory *FactoryFilterer) FilterSetPauser(opts *bind.FilterOpts, pendingPauser []common.Address) (*FactorySetPauserIterator, error) {

	var pendingPauserRule []interface{}
	for _, pendingPauserItem := range pendingPauser {
		pendingPauserRule = append(pendingPauserRule, pendingPauserItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "SetPauser", pendingPauserRule)
	if err != nil {
		return nil, err
	}
	return &FactorySetPauserIterator{contract: _Factory.contract, event: "SetPauser", logs: logs, sub: sub}, nil
}

// WatchSetPauser is a free log subscription operation binding the contract event 0xe02efb9e8f0fc21546730ab32d594f62d586e1bbb15bb5045edd0b1878a77b35.
//
// Solidity: event SetPauser(address indexed pendingPauser)
func (_Factory *FactoryFilterer) WatchSetPauser(opts *bind.WatchOpts, sink chan<- *FactorySetPauser, pendingPauser []common.Address) (event.Subscription, error) {

	var pendingPauserRule []interface{}
	for _, pendingPauserItem := range pendingPauser {
		pendingPauserRule = append(pendingPauserRule, pendingPauserItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "SetPauser", pendingPauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactorySetPauser)
				if err := _Factory.contract.UnpackLog(event, "SetPauser", log); err != nil {
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

// ParseSetPauser is a log parse operation binding the contract event 0xe02efb9e8f0fc21546730ab32d594f62d586e1bbb15bb5045edd0b1878a77b35.
//
// Solidity: event SetPauser(address indexed pendingPauser)
func (_Factory *FactoryFilterer) ParseSetPauser(log types.Log) (*FactorySetPauser, error) {
	event := new(FactorySetPauser)
	if err := _Factory.contract.UnpackLog(event, "SetPauser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
