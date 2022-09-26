// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pair

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

// BaseV1PairObservation is an auto generated low-level Go binding around an user-defined struct.
type BaseV1PairObservation struct {
	Timestamp          *big.Int
	Reserve0Cumulative *big.Int
	Reserve1Cumulative *big.Int
}

// PairMetaData contains all meta data concerning the Pair contract.
var PairMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Fees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockTimestampLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"claimed0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimed1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimable0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimable1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"current\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentCumulativePrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserve0Cumulative\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve1Cumulative\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"claimed0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimed1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_reserve0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserve1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestampLast\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"index0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"index1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastObservation\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve0Cumulative\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve1Cumulative\",\"type\":\"uint256\"}],\"internalType\":\"structBaseV1Pair.Observation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadata\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dec0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dec1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r1\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"st\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"t0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t1\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"observationLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"observations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve0Cumulative\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve1Cumulative\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"points\",\"type\":\"uint256\"}],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"granularity\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"points\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"window\",\"type\":\"uint256\"}],\"name\":\"sample\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supplyIndex0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supplyIndex1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PairABI is the input ABI used to generate the binding from.
// Deprecated: Use PairMetaData.ABI instead.
var PairABI = PairMetaData.ABI

// Pair is an auto generated Go binding around an Ethereum contract.
type Pair struct {
	PairCaller     // Read-only binding to the contract
	PairTransactor // Write-only binding to the contract
	PairFilterer   // Log filterer for contract events
}

// PairCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairSession struct {
	Contract     *Pair             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairCallerSession struct {
	Contract *PairCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairTransactorSession struct {
	Contract     *PairTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairRaw struct {
	Contract *Pair // Generic contract binding to access the raw methods on
}

// PairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairCallerRaw struct {
	Contract *PairCaller // Generic read-only contract binding to access the raw methods on
}

// PairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairTransactorRaw struct {
	Contract *PairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPair creates a new instance of Pair, bound to a specific deployed contract.
func NewPair(address common.Address, backend bind.ContractBackend) (*Pair, error) {
	contract, err := bindPair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pair{PairCaller: PairCaller{contract: contract}, PairTransactor: PairTransactor{contract: contract}, PairFilterer: PairFilterer{contract: contract}}, nil
}

// NewPairCaller creates a new read-only instance of Pair, bound to a specific deployed contract.
func NewPairCaller(address common.Address, caller bind.ContractCaller) (*PairCaller, error) {
	contract, err := bindPair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairCaller{contract: contract}, nil
}

// NewPairTransactor creates a new write-only instance of Pair, bound to a specific deployed contract.
func NewPairTransactor(address common.Address, transactor bind.ContractTransactor) (*PairTransactor, error) {
	contract, err := bindPair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairTransactor{contract: contract}, nil
}

// NewPairFilterer creates a new log filterer instance of Pair, bound to a specific deployed contract.
func NewPairFilterer(address common.Address, filterer bind.ContractFilterer) (*PairFilterer, error) {
	contract, err := bindPair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairFilterer{contract: contract}, nil
}

// bindPair binds a generic wrapper to an already deployed contract.
func bindPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pair *PairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pair.Contract.PairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pair *PairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pair.Contract.PairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pair *PairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pair.Contract.PairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pair *PairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pair *PairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pair *PairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pair.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pair *PairCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pair *PairSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Pair.Contract.Allowance(&_Pair.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pair *PairCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Pair.Contract.Allowance(&_Pair.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Pair *PairCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Pair *PairSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Pair.Contract.BalanceOf(&_Pair.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Pair *PairCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Pair.Contract.BalanceOf(&_Pair.CallOpts, arg0)
}

// BlockTimestampLast is a free data retrieval call binding the contract method 0xc5700a02.
//
// Solidity: function blockTimestampLast() view returns(uint256)
func (_Pair *PairCaller) BlockTimestampLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "blockTimestampLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockTimestampLast is a free data retrieval call binding the contract method 0xc5700a02.
//
// Solidity: function blockTimestampLast() view returns(uint256)
func (_Pair *PairSession) BlockTimestampLast() (*big.Int, error) {
	return _Pair.Contract.BlockTimestampLast(&_Pair.CallOpts)
}

// BlockTimestampLast is a free data retrieval call binding the contract method 0xc5700a02.
//
// Solidity: function blockTimestampLast() view returns(uint256)
func (_Pair *PairCallerSession) BlockTimestampLast() (*big.Int, error) {
	return _Pair.Contract.BlockTimestampLast(&_Pair.CallOpts)
}

// Claimable0 is a free data retrieval call binding the contract method 0x4d5a9f8a.
//
// Solidity: function claimable0(address ) view returns(uint256)
func (_Pair *PairCaller) Claimable0(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "claimable0", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable0 is a free data retrieval call binding the contract method 0x4d5a9f8a.
//
// Solidity: function claimable0(address ) view returns(uint256)
func (_Pair *PairSession) Claimable0(arg0 common.Address) (*big.Int, error) {
	return _Pair.Contract.Claimable0(&_Pair.CallOpts, arg0)
}

// Claimable0 is a free data retrieval call binding the contract method 0x4d5a9f8a.
//
// Solidity: function claimable0(address ) view returns(uint256)
func (_Pair *PairCallerSession) Claimable0(arg0 common.Address) (*big.Int, error) {
	return _Pair.Contract.Claimable0(&_Pair.CallOpts, arg0)
}

// Claimable1 is a free data retrieval call binding the contract method 0xa1ac4d13.
//
// Solidity: function claimable1(address ) view returns(uint256)
func (_Pair *PairCaller) Claimable1(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "claimable1", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable1 is a free data retrieval call binding the contract method 0xa1ac4d13.
//
// Solidity: function claimable1(address ) view returns(uint256)
func (_Pair *PairSession) Claimable1(arg0 common.Address) (*big.Int, error) {
	return _Pair.Contract.Claimable1(&_Pair.CallOpts, arg0)
}

// Claimable1 is a free data retrieval call binding the contract method 0xa1ac4d13.
//
// Solidity: function claimable1(address ) view returns(uint256)
func (_Pair *PairCallerSession) Claimable1(arg0 common.Address) (*big.Int, error) {
	return _Pair.Contract.Claimable1(&_Pair.CallOpts, arg0)
}

// Current is a free data retrieval call binding the contract method 0x517b3f82.
//
// Solidity: function current(address tokenIn, uint256 amountIn) view returns(uint256 amountOut)
func (_Pair *PairCaller) Current(opts *bind.CallOpts, tokenIn common.Address, amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "current", tokenIn, amountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Current is a free data retrieval call binding the contract method 0x517b3f82.
//
// Solidity: function current(address tokenIn, uint256 amountIn) view returns(uint256 amountOut)
func (_Pair *PairSession) Current(tokenIn common.Address, amountIn *big.Int) (*big.Int, error) {
	return _Pair.Contract.Current(&_Pair.CallOpts, tokenIn, amountIn)
}

// Current is a free data retrieval call binding the contract method 0x517b3f82.
//
// Solidity: function current(address tokenIn, uint256 amountIn) view returns(uint256 amountOut)
func (_Pair *PairCallerSession) Current(tokenIn common.Address, amountIn *big.Int) (*big.Int, error) {
	return _Pair.Contract.Current(&_Pair.CallOpts, tokenIn, amountIn)
}

// CurrentCumulativePrices is a free data retrieval call binding the contract method 0x1df8c717.
//
// Solidity: function currentCumulativePrices() view returns(uint256 reserve0Cumulative, uint256 reserve1Cumulative, uint256 blockTimestamp)
func (_Pair *PairCaller) CurrentCumulativePrices(opts *bind.CallOpts) (struct {
	Reserve0Cumulative *big.Int
	Reserve1Cumulative *big.Int
	BlockTimestamp     *big.Int
}, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "currentCumulativePrices")

	outstruct := new(struct {
		Reserve0Cumulative *big.Int
		Reserve1Cumulative *big.Int
		BlockTimestamp     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0Cumulative = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1Cumulative = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentCumulativePrices is a free data retrieval call binding the contract method 0x1df8c717.
//
// Solidity: function currentCumulativePrices() view returns(uint256 reserve0Cumulative, uint256 reserve1Cumulative, uint256 blockTimestamp)
func (_Pair *PairSession) CurrentCumulativePrices() (struct {
	Reserve0Cumulative *big.Int
	Reserve1Cumulative *big.Int
	BlockTimestamp     *big.Int
}, error) {
	return _Pair.Contract.CurrentCumulativePrices(&_Pair.CallOpts)
}

// CurrentCumulativePrices is a free data retrieval call binding the contract method 0x1df8c717.
//
// Solidity: function currentCumulativePrices() view returns(uint256 reserve0Cumulative, uint256 reserve1Cumulative, uint256 blockTimestamp)
func (_Pair *PairCallerSession) CurrentCumulativePrices() (struct {
	Reserve0Cumulative *big.Int
	Reserve1Cumulative *big.Int
	BlockTimestamp     *big.Int
}, error) {
	return _Pair.Contract.CurrentCumulativePrices(&_Pair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pair *PairCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pair *PairSession) Decimals() (uint8, error) {
	return _Pair.Contract.Decimals(&_Pair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pair *PairCallerSession) Decimals() (uint8, error) {
	return _Pair.Contract.Decimals(&_Pair.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(address)
func (_Pair *PairCaller) Fees(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "fees")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(address)
func (_Pair *PairSession) Fees() (common.Address, error) {
	return _Pair.Contract.Fees(&_Pair.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(address)
func (_Pair *PairCallerSession) Fees() (common.Address, error) {
	return _Pair.Contract.Fees(&_Pair.CallOpts)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn) view returns(uint256)
func (_Pair *PairCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, tokenIn common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "getAmountOut", amountIn, tokenIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn) view returns(uint256)
func (_Pair *PairSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address) (*big.Int, error) {
	return _Pair.Contract.GetAmountOut(&_Pair.CallOpts, amountIn, tokenIn)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn) view returns(uint256)
func (_Pair *PairCallerSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address) (*big.Int, error) {
	return _Pair.Contract.GetAmountOut(&_Pair.CallOpts, amountIn, tokenIn)
}

// GetFees is a free data retrieval call binding the contract method 0x9af608c9.
//
// Solidity: function getFees(address account) view returns(uint256 claimed0, uint256 claimed1)
func (_Pair *PairCaller) GetFees(opts *bind.CallOpts, account common.Address) (struct {
	Claimed0 *big.Int
	Claimed1 *big.Int
}, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "getFees", account)

	outstruct := new(struct {
		Claimed0 *big.Int
		Claimed1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Claimed0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Claimed1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetFees is a free data retrieval call binding the contract method 0x9af608c9.
//
// Solidity: function getFees(address account) view returns(uint256 claimed0, uint256 claimed1)
func (_Pair *PairSession) GetFees(account common.Address) (struct {
	Claimed0 *big.Int
	Claimed1 *big.Int
}, error) {
	return _Pair.Contract.GetFees(&_Pair.CallOpts, account)
}

// GetFees is a free data retrieval call binding the contract method 0x9af608c9.
//
// Solidity: function getFees(address account) view returns(uint256 claimed0, uint256 claimed1)
func (_Pair *PairCallerSession) GetFees(account common.Address) (struct {
	Claimed0 *big.Int
	Claimed1 *big.Int
}, error) {
	return _Pair.Contract.GetFees(&_Pair.CallOpts, account)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1, uint256 _blockTimestampLast)
func (_Pair *PairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast *big.Int
}, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1, uint256 _blockTimestampLast)
func (_Pair *PairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast *big.Int
}, error) {
	return _Pair.Contract.GetReserves(&_Pair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1, uint256 _blockTimestampLast)
func (_Pair *PairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast *big.Int
}, error) {
	return _Pair.Contract.GetReserves(&_Pair.CallOpts)
}

// Index0 is a free data retrieval call binding the contract method 0x32c0defd.
//
// Solidity: function index0() view returns(uint256)
func (_Pair *PairCaller) Index0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "index0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Index0 is a free data retrieval call binding the contract method 0x32c0defd.
//
// Solidity: function index0() view returns(uint256)
func (_Pair *PairSession) Index0() (*big.Int, error) {
	return _Pair.Contract.Index0(&_Pair.CallOpts)
}

// Index0 is a free data retrieval call binding the contract method 0x32c0defd.
//
// Solidity: function index0() view returns(uint256)
func (_Pair *PairCallerSession) Index0() (*big.Int, error) {
	return _Pair.Contract.Index0(&_Pair.CallOpts)
}

// Index1 is a free data retrieval call binding the contract method 0xbda39cad.
//
// Solidity: function index1() view returns(uint256)
func (_Pair *PairCaller) Index1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "index1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Index1 is a free data retrieval call binding the contract method 0xbda39cad.
//
// Solidity: function index1() view returns(uint256)
func (_Pair *PairSession) Index1() (*big.Int, error) {
	return _Pair.Contract.Index1(&_Pair.CallOpts)
}

// Index1 is a free data retrieval call binding the contract method 0xbda39cad.
//
// Solidity: function index1() view returns(uint256)
func (_Pair *PairCallerSession) Index1() (*big.Int, error) {
	return _Pair.Contract.Index1(&_Pair.CallOpts)
}

// LastObservation is a free data retrieval call binding the contract method 0x8a7b8cf2.
//
// Solidity: function lastObservation() view returns((uint256,uint256,uint256))
func (_Pair *PairCaller) LastObservation(opts *bind.CallOpts) (BaseV1PairObservation, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "lastObservation")

	if err != nil {
		return *new(BaseV1PairObservation), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseV1PairObservation)).(*BaseV1PairObservation)

	return out0, err

}

// LastObservation is a free data retrieval call binding the contract method 0x8a7b8cf2.
//
// Solidity: function lastObservation() view returns((uint256,uint256,uint256))
func (_Pair *PairSession) LastObservation() (BaseV1PairObservation, error) {
	return _Pair.Contract.LastObservation(&_Pair.CallOpts)
}

// LastObservation is a free data retrieval call binding the contract method 0x8a7b8cf2.
//
// Solidity: function lastObservation() view returns((uint256,uint256,uint256))
func (_Pair *PairCallerSession) LastObservation() (BaseV1PairObservation, error) {
	return _Pair.Contract.LastObservation(&_Pair.CallOpts)
}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(uint256 dec0, uint256 dec1, uint256 r0, uint256 r1, bool st, address t0, address t1)
func (_Pair *PairCaller) Metadata(opts *bind.CallOpts) (struct {
	Dec0 *big.Int
	Dec1 *big.Int
	R0   *big.Int
	R1   *big.Int
	St   bool
	T0   common.Address
	T1   common.Address
}, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "metadata")

	outstruct := new(struct {
		Dec0 *big.Int
		Dec1 *big.Int
		R0   *big.Int
		R1   *big.Int
		St   bool
		T0   common.Address
		T1   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dec0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Dec1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.R0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.R1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.St = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.T0 = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.T1 = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(uint256 dec0, uint256 dec1, uint256 r0, uint256 r1, bool st, address t0, address t1)
func (_Pair *PairSession) Metadata() (struct {
	Dec0 *big.Int
	Dec1 *big.Int
	R0   *big.Int
	R1   *big.Int
	St   bool
	T0   common.Address
	T1   common.Address
}, error) {
	return _Pair.Contract.Metadata(&_Pair.CallOpts)
}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(uint256 dec0, uint256 dec1, uint256 r0, uint256 r1, bool st, address t0, address t1)
func (_Pair *PairCallerSession) Metadata() (struct {
	Dec0 *big.Int
	Dec1 *big.Int
	R0   *big.Int
	R1   *big.Int
	St   bool
	T0   common.Address
	T1   common.Address
}, error) {
	return _Pair.Contract.Metadata(&_Pair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pair *PairCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pair *PairSession) Name() (string, error) {
	return _Pair.Contract.Name(&_Pair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pair *PairCallerSession) Name() (string, error) {
	return _Pair.Contract.Name(&_Pair.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Pair *PairCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Pair *PairSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Pair.Contract.Nonces(&_Pair.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Pair *PairCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Pair.Contract.Nonces(&_Pair.CallOpts, arg0)
}

// ObservationLength is a free data retrieval call binding the contract method 0xebeb31db.
//
// Solidity: function observationLength() view returns(uint256)
func (_Pair *PairCaller) ObservationLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "observationLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ObservationLength is a free data retrieval call binding the contract method 0xebeb31db.
//
// Solidity: function observationLength() view returns(uint256)
func (_Pair *PairSession) ObservationLength() (*big.Int, error) {
	return _Pair.Contract.ObservationLength(&_Pair.CallOpts)
}

// ObservationLength is a free data retrieval call binding the contract method 0xebeb31db.
//
// Solidity: function observationLength() view returns(uint256)
func (_Pair *PairCallerSession) ObservationLength() (*big.Int, error) {
	return _Pair.Contract.ObservationLength(&_Pair.CallOpts)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint256 timestamp, uint256 reserve0Cumulative, uint256 reserve1Cumulative)
func (_Pair *PairCaller) Observations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp          *big.Int
	Reserve0Cumulative *big.Int
	Reserve1Cumulative *big.Int
}, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "observations", arg0)

	outstruct := new(struct {
		Timestamp          *big.Int
		Reserve0Cumulative *big.Int
		Reserve1Cumulative *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve0Cumulative = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Reserve1Cumulative = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint256 timestamp, uint256 reserve0Cumulative, uint256 reserve1Cumulative)
func (_Pair *PairSession) Observations(arg0 *big.Int) (struct {
	Timestamp          *big.Int
	Reserve0Cumulative *big.Int
	Reserve1Cumulative *big.Int
}, error) {
	return _Pair.Contract.Observations(&_Pair.CallOpts, arg0)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint256 timestamp, uint256 reserve0Cumulative, uint256 reserve1Cumulative)
func (_Pair *PairCallerSession) Observations(arg0 *big.Int) (struct {
	Timestamp          *big.Int
	Reserve0Cumulative *big.Int
	Reserve1Cumulative *big.Int
}, error) {
	return _Pair.Contract.Observations(&_Pair.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0x5881c475.
//
// Solidity: function prices(address tokenIn, uint256 amountIn, uint256 points) view returns(uint256[])
func (_Pair *PairCaller) Prices(opts *bind.CallOpts, tokenIn common.Address, amountIn *big.Int, points *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "prices", tokenIn, amountIn, points)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Prices is a free data retrieval call binding the contract method 0x5881c475.
//
// Solidity: function prices(address tokenIn, uint256 amountIn, uint256 points) view returns(uint256[])
func (_Pair *PairSession) Prices(tokenIn common.Address, amountIn *big.Int, points *big.Int) ([]*big.Int, error) {
	return _Pair.Contract.Prices(&_Pair.CallOpts, tokenIn, amountIn, points)
}

// Prices is a free data retrieval call binding the contract method 0x5881c475.
//
// Solidity: function prices(address tokenIn, uint256 amountIn, uint256 points) view returns(uint256[])
func (_Pair *PairCallerSession) Prices(tokenIn common.Address, amountIn *big.Int, points *big.Int) ([]*big.Int, error) {
	return _Pair.Contract.Prices(&_Pair.CallOpts, tokenIn, amountIn, points)
}

// Quote is a free data retrieval call binding the contract method 0x9e8cc04b.
//
// Solidity: function quote(address tokenIn, uint256 amountIn, uint256 granularity) view returns(uint256 amountOut)
func (_Pair *PairCaller) Quote(opts *bind.CallOpts, tokenIn common.Address, amountIn *big.Int, granularity *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "quote", tokenIn, amountIn, granularity)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0x9e8cc04b.
//
// Solidity: function quote(address tokenIn, uint256 amountIn, uint256 granularity) view returns(uint256 amountOut)
func (_Pair *PairSession) Quote(tokenIn common.Address, amountIn *big.Int, granularity *big.Int) (*big.Int, error) {
	return _Pair.Contract.Quote(&_Pair.CallOpts, tokenIn, amountIn, granularity)
}

// Quote is a free data retrieval call binding the contract method 0x9e8cc04b.
//
// Solidity: function quote(address tokenIn, uint256 amountIn, uint256 granularity) view returns(uint256 amountOut)
func (_Pair *PairCallerSession) Quote(tokenIn common.Address, amountIn *big.Int, granularity *big.Int) (*big.Int, error) {
	return _Pair.Contract.Quote(&_Pair.CallOpts, tokenIn, amountIn, granularity)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_Pair *PairCaller) Reserve0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "reserve0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_Pair *PairSession) Reserve0() (*big.Int, error) {
	return _Pair.Contract.Reserve0(&_Pair.CallOpts)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_Pair *PairCallerSession) Reserve0() (*big.Int, error) {
	return _Pair.Contract.Reserve0(&_Pair.CallOpts)
}

// Reserve0CumulativeLast is a free data retrieval call binding the contract method 0xbf944dbc.
//
// Solidity: function reserve0CumulativeLast() view returns(uint256)
func (_Pair *PairCaller) Reserve0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "reserve0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve0CumulativeLast is a free data retrieval call binding the contract method 0xbf944dbc.
//
// Solidity: function reserve0CumulativeLast() view returns(uint256)
func (_Pair *PairSession) Reserve0CumulativeLast() (*big.Int, error) {
	return _Pair.Contract.Reserve0CumulativeLast(&_Pair.CallOpts)
}

// Reserve0CumulativeLast is a free data retrieval call binding the contract method 0xbf944dbc.
//
// Solidity: function reserve0CumulativeLast() view returns(uint256)
func (_Pair *PairCallerSession) Reserve0CumulativeLast() (*big.Int, error) {
	return _Pair.Contract.Reserve0CumulativeLast(&_Pair.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_Pair *PairCaller) Reserve1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "reserve1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_Pair *PairSession) Reserve1() (*big.Int, error) {
	return _Pair.Contract.Reserve1(&_Pair.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_Pair *PairCallerSession) Reserve1() (*big.Int, error) {
	return _Pair.Contract.Reserve1(&_Pair.CallOpts)
}

// Reserve1CumulativeLast is a free data retrieval call binding the contract method 0xc245febc.
//
// Solidity: function reserve1CumulativeLast() view returns(uint256)
func (_Pair *PairCaller) Reserve1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "reserve1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve1CumulativeLast is a free data retrieval call binding the contract method 0xc245febc.
//
// Solidity: function reserve1CumulativeLast() view returns(uint256)
func (_Pair *PairSession) Reserve1CumulativeLast() (*big.Int, error) {
	return _Pair.Contract.Reserve1CumulativeLast(&_Pair.CallOpts)
}

// Reserve1CumulativeLast is a free data retrieval call binding the contract method 0xc245febc.
//
// Solidity: function reserve1CumulativeLast() view returns(uint256)
func (_Pair *PairCallerSession) Reserve1CumulativeLast() (*big.Int, error) {
	return _Pair.Contract.Reserve1CumulativeLast(&_Pair.CallOpts)
}

// Sample is a free data retrieval call binding the contract method 0x13345fe1.
//
// Solidity: function sample(address tokenIn, uint256 amountIn, uint256 points, uint256 window) view returns(uint256[])
func (_Pair *PairCaller) Sample(opts *bind.CallOpts, tokenIn common.Address, amountIn *big.Int, points *big.Int, window *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "sample", tokenIn, amountIn, points, window)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Sample is a free data retrieval call binding the contract method 0x13345fe1.
//
// Solidity: function sample(address tokenIn, uint256 amountIn, uint256 points, uint256 window) view returns(uint256[])
func (_Pair *PairSession) Sample(tokenIn common.Address, amountIn *big.Int, points *big.Int, window *big.Int) ([]*big.Int, error) {
	return _Pair.Contract.Sample(&_Pair.CallOpts, tokenIn, amountIn, points, window)
}

// Sample is a free data retrieval call binding the contract method 0x13345fe1.
//
// Solidity: function sample(address tokenIn, uint256 amountIn, uint256 points, uint256 window) view returns(uint256[])
func (_Pair *PairCallerSession) Sample(tokenIn common.Address, amountIn *big.Int, points *big.Int, window *big.Int) ([]*big.Int, error) {
	return _Pair.Contract.Sample(&_Pair.CallOpts, tokenIn, amountIn, points, window)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(bool)
func (_Pair *PairCaller) Stable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "stable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(bool)
func (_Pair *PairSession) Stable() (bool, error) {
	return _Pair.Contract.Stable(&_Pair.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(bool)
func (_Pair *PairCallerSession) Stable() (bool, error) {
	return _Pair.Contract.Stable(&_Pair.CallOpts)
}

// SupplyIndex0 is a free data retrieval call binding the contract method 0x9f767c88.
//
// Solidity: function supplyIndex0(address ) view returns(uint256)
func (_Pair *PairCaller) SupplyIndex0(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "supplyIndex0", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyIndex0 is a free data retrieval call binding the contract method 0x9f767c88.
//
// Solidity: function supplyIndex0(address ) view returns(uint256)
func (_Pair *PairSession) SupplyIndex0(arg0 common.Address) (*big.Int, error) {
	return _Pair.Contract.SupplyIndex0(&_Pair.CallOpts, arg0)
}

// SupplyIndex0 is a free data retrieval call binding the contract method 0x9f767c88.
//
// Solidity: function supplyIndex0(address ) view returns(uint256)
func (_Pair *PairCallerSession) SupplyIndex0(arg0 common.Address) (*big.Int, error) {
	return _Pair.Contract.SupplyIndex0(&_Pair.CallOpts, arg0)
}

// SupplyIndex1 is a free data retrieval call binding the contract method 0x205aabf1.
//
// Solidity: function supplyIndex1(address ) view returns(uint256)
func (_Pair *PairCaller) SupplyIndex1(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "supplyIndex1", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyIndex1 is a free data retrieval call binding the contract method 0x205aabf1.
//
// Solidity: function supplyIndex1(address ) view returns(uint256)
func (_Pair *PairSession) SupplyIndex1(arg0 common.Address) (*big.Int, error) {
	return _Pair.Contract.SupplyIndex1(&_Pair.CallOpts, arg0)
}

// SupplyIndex1 is a free data retrieval call binding the contract method 0x205aabf1.
//
// Solidity: function supplyIndex1(address ) view returns(uint256)
func (_Pair *PairCallerSession) SupplyIndex1(arg0 common.Address) (*big.Int, error) {
	return _Pair.Contract.SupplyIndex1(&_Pair.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pair *PairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pair *PairSession) Symbol() (string, error) {
	return _Pair.Contract.Symbol(&_Pair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pair *PairCallerSession) Symbol() (string, error) {
	return _Pair.Contract.Symbol(&_Pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pair *PairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pair *PairSession) Token0() (common.Address, error) {
	return _Pair.Contract.Token0(&_Pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pair *PairCallerSession) Token0() (common.Address, error) {
	return _Pair.Contract.Token0(&_Pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pair *PairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pair *PairSession) Token1() (common.Address, error) {
	return _Pair.Contract.Token1(&_Pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pair *PairCallerSession) Token1() (common.Address, error) {
	return _Pair.Contract.Token1(&_Pair.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() view returns(address, address)
func (_Pair *PairCaller) Tokens(opts *bind.CallOpts) (common.Address, common.Address, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "tokens")

	if err != nil {
		return *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() view returns(address, address)
func (_Pair *PairSession) Tokens() (common.Address, common.Address, error) {
	return _Pair.Contract.Tokens(&_Pair.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() view returns(address, address)
func (_Pair *PairCallerSession) Tokens() (common.Address, common.Address, error) {
	return _Pair.Contract.Tokens(&_Pair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pair *PairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pair *PairSession) TotalSupply() (*big.Int, error) {
	return _Pair.Contract.TotalSupply(&_Pair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pair *PairCallerSession) TotalSupply() (*big.Int, error) {
	return _Pair.Contract.TotalSupply(&_Pair.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Pair *PairTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Pair *PairSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pair.Contract.Approve(&_Pair.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Pair *PairTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pair.Contract.Approve(&_Pair.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pair *PairTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pair *PairSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Burn(&_Pair.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pair *PairTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Burn(&_Pair.TransactOpts, to)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns(uint256 claimed0, uint256 claimed1)
func (_Pair *PairTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns(uint256 claimed0, uint256 claimed1)
func (_Pair *PairSession) ClaimFees() (*types.Transaction, error) {
	return _Pair.Contract.ClaimFees(&_Pair.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns(uint256 claimed0, uint256 claimed1)
func (_Pair *PairTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _Pair.Contract.ClaimFees(&_Pair.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pair *PairTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pair *PairSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Mint(&_Pair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pair *PairTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Mint(&_Pair.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pair *PairTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pair *PairSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pair.Contract.Permit(&_Pair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pair *PairTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pair.Contract.Permit(&_Pair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pair *PairTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pair *PairSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Skim(&_Pair.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pair *PairTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Skim(&_Pair.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pair *PairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pair *PairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Pair.Contract.Swap(&_Pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pair *PairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Pair.Contract.Swap(&_Pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pair *PairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pair *PairSession) Sync() (*types.Transaction, error) {
	return _Pair.Contract.Sync(&_Pair.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pair *PairTransactorSession) Sync() (*types.Transaction, error) {
	return _Pair.Contract.Sync(&_Pair.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Pair *PairTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Pair *PairSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pair.Contract.Transfer(&_Pair.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Pair *PairTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pair.Contract.Transfer(&_Pair.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Pair *PairTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Pair *PairSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pair.Contract.TransferFrom(&_Pair.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Pair *PairTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pair.Contract.TransferFrom(&_Pair.TransactOpts, src, dst, amount)
}

// PairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pair contract.
type PairApprovalIterator struct {
	Event *PairApproval // Event containing the contract specifics and raw log

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
func (it *PairApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairApproval)
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
		it.Event = new(PairApproval)
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
func (it *PairApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairApproval represents a Approval event raised by the Pair contract.
type PairApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Pair *PairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PairApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PairApprovalIterator{contract: _Pair.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Pair *PairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairApproval)
				if err := _Pair.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Pair *PairFilterer) ParseApproval(log types.Log) (*PairApproval, error) {
	event := new(PairApproval)
	if err := _Pair.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Pair contract.
type PairBurnIterator struct {
	Event *PairBurn // Event containing the contract specifics and raw log

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
func (it *PairBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairBurn)
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
		it.Event = new(PairBurn)
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
func (it *PairBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairBurn represents a Burn event raised by the Pair contract.
type PairBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Pair *PairFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PairBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PairBurnIterator{contract: _Pair.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Pair *PairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PairBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairBurn)
				if err := _Pair.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Pair *PairFilterer) ParseBurn(log types.Log) (*PairBurn, error) {
	event := new(PairBurn)
	if err := _Pair.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Pair contract.
type PairClaimIterator struct {
	Event *PairClaim // Event containing the contract specifics and raw log

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
func (it *PairClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairClaim)
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
		it.Event = new(PairClaim)
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
func (it *PairClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairClaim represents a Claim event raised by the Pair contract.
type PairClaim struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x865ca08d59f5cb456e85cd2f7ef63664ea4f73327414e9d8152c4158b0e94645.
//
// Solidity: event Claim(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1)
func (_Pair *PairFilterer) FilterClaim(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*PairClaimIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Claim", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &PairClaimIterator{contract: _Pair.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x865ca08d59f5cb456e85cd2f7ef63664ea4f73327414e9d8152c4158b0e94645.
//
// Solidity: event Claim(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1)
func (_Pair *PairFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *PairClaim, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Claim", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairClaim)
				if err := _Pair.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x865ca08d59f5cb456e85cd2f7ef63664ea4f73327414e9d8152c4158b0e94645.
//
// Solidity: event Claim(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1)
func (_Pair *PairFilterer) ParseClaim(log types.Log) (*PairClaim, error) {
	event := new(PairClaim)
	if err := _Pair.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairFeesIterator is returned from FilterFees and is used to iterate over the raw logs and unpacked data for Fees events raised by the Pair contract.
type PairFeesIterator struct {
	Event *PairFees // Event containing the contract specifics and raw log

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
func (it *PairFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairFees)
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
		it.Event = new(PairFees)
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
func (it *PairFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairFees represents a Fees event raised by the Pair contract.
type PairFees struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFees is a free log retrieval operation binding the contract event 0x112c256902bf554b6ed882d2936687aaeb4225e8cd5b51303c90ca6cf43a8602.
//
// Solidity: event Fees(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pair *PairFilterer) FilterFees(opts *bind.FilterOpts, sender []common.Address) (*PairFeesIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Fees", senderRule)
	if err != nil {
		return nil, err
	}
	return &PairFeesIterator{contract: _Pair.contract, event: "Fees", logs: logs, sub: sub}, nil
}

// WatchFees is a free log subscription operation binding the contract event 0x112c256902bf554b6ed882d2936687aaeb4225e8cd5b51303c90ca6cf43a8602.
//
// Solidity: event Fees(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pair *PairFilterer) WatchFees(opts *bind.WatchOpts, sink chan<- *PairFees, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Fees", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairFees)
				if err := _Pair.contract.UnpackLog(event, "Fees", log); err != nil {
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

// ParseFees is a log parse operation binding the contract event 0x112c256902bf554b6ed882d2936687aaeb4225e8cd5b51303c90ca6cf43a8602.
//
// Solidity: event Fees(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pair *PairFilterer) ParseFees(log types.Log) (*PairFees, error) {
	event := new(PairFees)
	if err := _Pair.contract.UnpackLog(event, "Fees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Pair contract.
type PairMintIterator struct {
	Event *PairMint // Event containing the contract specifics and raw log

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
func (it *PairMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairMint)
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
		it.Event = new(PairMint)
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
func (it *PairMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairMint represents a Mint event raised by the Pair contract.
type PairMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pair *PairFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*PairMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &PairMintIterator{contract: _Pair.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pair *PairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PairMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairMint)
				if err := _Pair.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pair *PairFilterer) ParseMint(log types.Log) (*PairMint, error) {
	event := new(PairMint)
	if err := _Pair.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Pair contract.
type PairSwapIterator struct {
	Event *PairSwap // Event containing the contract specifics and raw log

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
func (it *PairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairSwap)
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
		it.Event = new(PairSwap)
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
func (it *PairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairSwap represents a Swap event raised by the Pair contract.
type PairSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Pair *PairFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PairSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PairSwapIterator{contract: _Pair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Pair *PairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PairSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairSwap)
				if err := _Pair.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Pair *PairFilterer) ParseSwap(log types.Log) (*PairSwap, error) {
	event := new(PairSwap)
	if err := _Pair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the Pair contract.
type PairSyncIterator struct {
	Event *PairSync // Event containing the contract specifics and raw log

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
func (it *PairSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairSync)
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
		it.Event = new(PairSync)
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
func (it *PairSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairSync represents a Sync event raised by the Pair contract.
type PairSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_Pair *PairFilterer) FilterSync(opts *bind.FilterOpts) (*PairSyncIterator, error) {

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &PairSyncIterator{contract: _Pair.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_Pair *PairFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *PairSync) (event.Subscription, error) {

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairSync)
				if err := _Pair.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_Pair *PairFilterer) ParseSync(log types.Log) (*PairSync, error) {
	event := new(PairSync)
	if err := _Pair.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pair contract.
type PairTransferIterator struct {
	Event *PairTransfer // Event containing the contract specifics and raw log

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
func (it *PairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairTransfer)
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
		it.Event = new(PairTransfer)
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
func (it *PairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairTransfer represents a Transfer event raised by the Pair contract.
type PairTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Pair *PairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PairTransferIterator{contract: _Pair.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Pair *PairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairTransfer)
				if err := _Pair.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Pair *PairFilterer) ParseTransfer(log types.Log) (*PairTransfer, error) {
	event := new(PairTransfer)
	if err := _Pair.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
