// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gauge

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

// GaugeMetaData contains all meta data concerning the Gauge contract.
var GaugeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stake\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bribe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__ve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimed0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimed1\",\"type\":\"uint256\"}],\"name\":\"ClaimFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NotifyReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_ve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxRuns\",\"type\":\"uint256\"}],\"name\":\"batchRewardPerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bribe\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceOf\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"claimed0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimed1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"derivedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"derivedBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"derivedSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getPriorBalanceIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getPriorRewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getPriorSupplyIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastEarn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"left\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardPerTokenCheckpoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardPerTokenNumCheckpoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsListLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"supplyCheckpoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyNumCheckpoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GaugeABI is the input ABI used to generate the binding from.
// Deprecated: Use GaugeMetaData.ABI instead.
var GaugeABI = GaugeMetaData.ABI

// Gauge is an auto generated Go binding around an Ethereum contract.
type Gauge struct {
	GaugeCaller     // Read-only binding to the contract
	GaugeTransactor // Write-only binding to the contract
	GaugeFilterer   // Log filterer for contract events
}

// GaugeCaller is an auto generated read-only Go binding around an Ethereum contract.
type GaugeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GaugeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GaugeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GaugeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GaugeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GaugeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GaugeSession struct {
	Contract     *Gauge            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GaugeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GaugeCallerSession struct {
	Contract *GaugeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GaugeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GaugeTransactorSession struct {
	Contract     *GaugeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GaugeRaw is an auto generated low-level Go binding around an Ethereum contract.
type GaugeRaw struct {
	Contract *Gauge // Generic contract binding to access the raw methods on
}

// GaugeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GaugeCallerRaw struct {
	Contract *GaugeCaller // Generic read-only contract binding to access the raw methods on
}

// GaugeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GaugeTransactorRaw struct {
	Contract *GaugeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGauge creates a new instance of Gauge, bound to a specific deployed contract.
func NewGauge(address common.Address, backend bind.ContractBackend) (*Gauge, error) {
	contract, err := bindGauge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gauge{GaugeCaller: GaugeCaller{contract: contract}, GaugeTransactor: GaugeTransactor{contract: contract}, GaugeFilterer: GaugeFilterer{contract: contract}}, nil
}

// NewGaugeCaller creates a new read-only instance of Gauge, bound to a specific deployed contract.
func NewGaugeCaller(address common.Address, caller bind.ContractCaller) (*GaugeCaller, error) {
	contract, err := bindGauge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GaugeCaller{contract: contract}, nil
}

// NewGaugeTransactor creates a new write-only instance of Gauge, bound to a specific deployed contract.
func NewGaugeTransactor(address common.Address, transactor bind.ContractTransactor) (*GaugeTransactor, error) {
	contract, err := bindGauge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GaugeTransactor{contract: contract}, nil
}

// NewGaugeFilterer creates a new log filterer instance of Gauge, bound to a specific deployed contract.
func NewGaugeFilterer(address common.Address, filterer bind.ContractFilterer) (*GaugeFilterer, error) {
	contract, err := bindGauge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GaugeFilterer{contract: contract}, nil
}

// bindGauge binds a generic wrapper to an already deployed contract.
func bindGauge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GaugeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gauge *GaugeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gauge.Contract.GaugeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gauge *GaugeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gauge.Contract.GaugeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gauge *GaugeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gauge.Contract.GaugeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gauge *GaugeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gauge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gauge *GaugeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gauge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gauge *GaugeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gauge.Contract.contract.Transact(opts, method, params...)
}

// Ve is a free data retrieval call binding the contract method 0x8dd598fb.
//
// Solidity: function _ve() view returns(address)
func (_Gauge *GaugeCaller) Ve(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "_ve")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ve is a free data retrieval call binding the contract method 0x8dd598fb.
//
// Solidity: function _ve() view returns(address)
func (_Gauge *GaugeSession) Ve() (common.Address, error) {
	return _Gauge.Contract.Ve(&_Gauge.CallOpts)
}

// Ve is a free data retrieval call binding the contract method 0x8dd598fb.
//
// Solidity: function _ve() view returns(address)
func (_Gauge *GaugeCallerSession) Ve() (common.Address, error) {
	return _Gauge.Contract.Ve(&_Gauge.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Gauge *GaugeCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Gauge *GaugeSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.BalanceOf(&_Gauge.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Gauge *GaugeCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.BalanceOf(&_Gauge.CallOpts, arg0)
}

// Bribe is a free data retrieval call binding the contract method 0x37d0208c.
//
// Solidity: function bribe() view returns(address)
func (_Gauge *GaugeCaller) Bribe(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "bribe")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bribe is a free data retrieval call binding the contract method 0x37d0208c.
//
// Solidity: function bribe() view returns(address)
func (_Gauge *GaugeSession) Bribe() (common.Address, error) {
	return _Gauge.Contract.Bribe(&_Gauge.CallOpts)
}

// Bribe is a free data retrieval call binding the contract method 0x37d0208c.
//
// Solidity: function bribe() view returns(address)
func (_Gauge *GaugeCallerSession) Bribe() (common.Address, error) {
	return _Gauge.Contract.Bribe(&_Gauge.CallOpts)
}

// Checkpoints is a free data retrieval call binding the contract method 0x0cdfebfa.
//
// Solidity: function checkpoints(address , uint256 ) view returns(uint256 timestamp, uint256 balanceOf)
func (_Gauge *GaugeCaller) Checkpoints(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Timestamp *big.Int
	BalanceOf *big.Int
}, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "checkpoints", arg0, arg1)

	outstruct := new(struct {
		Timestamp *big.Int
		BalanceOf *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BalanceOf = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Checkpoints is a free data retrieval call binding the contract method 0x0cdfebfa.
//
// Solidity: function checkpoints(address , uint256 ) view returns(uint256 timestamp, uint256 balanceOf)
func (_Gauge *GaugeSession) Checkpoints(arg0 common.Address, arg1 *big.Int) (struct {
	Timestamp *big.Int
	BalanceOf *big.Int
}, error) {
	return _Gauge.Contract.Checkpoints(&_Gauge.CallOpts, arg0, arg1)
}

// Checkpoints is a free data retrieval call binding the contract method 0x0cdfebfa.
//
// Solidity: function checkpoints(address , uint256 ) view returns(uint256 timestamp, uint256 balanceOf)
func (_Gauge *GaugeCallerSession) Checkpoints(arg0 common.Address, arg1 *big.Int) (struct {
	Timestamp *big.Int
	BalanceOf *big.Int
}, error) {
	return _Gauge.Contract.Checkpoints(&_Gauge.CallOpts, arg0, arg1)
}

// DerivedBalance is a free data retrieval call binding the contract method 0xd35e2544.
//
// Solidity: function derivedBalance(address account) view returns(uint256)
func (_Gauge *GaugeCaller) DerivedBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "derivedBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DerivedBalance is a free data retrieval call binding the contract method 0xd35e2544.
//
// Solidity: function derivedBalance(address account) view returns(uint256)
func (_Gauge *GaugeSession) DerivedBalance(account common.Address) (*big.Int, error) {
	return _Gauge.Contract.DerivedBalance(&_Gauge.CallOpts, account)
}

// DerivedBalance is a free data retrieval call binding the contract method 0xd35e2544.
//
// Solidity: function derivedBalance(address account) view returns(uint256)
func (_Gauge *GaugeCallerSession) DerivedBalance(account common.Address) (*big.Int, error) {
	return _Gauge.Contract.DerivedBalance(&_Gauge.CallOpts, account)
}

// DerivedBalances is a free data retrieval call binding the contract method 0x63fb415b.
//
// Solidity: function derivedBalances(address ) view returns(uint256)
func (_Gauge *GaugeCaller) DerivedBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "derivedBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DerivedBalances is a free data retrieval call binding the contract method 0x63fb415b.
//
// Solidity: function derivedBalances(address ) view returns(uint256)
func (_Gauge *GaugeSession) DerivedBalances(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.DerivedBalances(&_Gauge.CallOpts, arg0)
}

// DerivedBalances is a free data retrieval call binding the contract method 0x63fb415b.
//
// Solidity: function derivedBalances(address ) view returns(uint256)
func (_Gauge *GaugeCallerSession) DerivedBalances(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.DerivedBalances(&_Gauge.CallOpts, arg0)
}

// DerivedSupply is a free data retrieval call binding the contract method 0xd7da4bb0.
//
// Solidity: function derivedSupply() view returns(uint256)
func (_Gauge *GaugeCaller) DerivedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "derivedSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DerivedSupply is a free data retrieval call binding the contract method 0xd7da4bb0.
//
// Solidity: function derivedSupply() view returns(uint256)
func (_Gauge *GaugeSession) DerivedSupply() (*big.Int, error) {
	return _Gauge.Contract.DerivedSupply(&_Gauge.CallOpts)
}

// DerivedSupply is a free data retrieval call binding the contract method 0xd7da4bb0.
//
// Solidity: function derivedSupply() view returns(uint256)
func (_Gauge *GaugeCallerSession) DerivedSupply() (*big.Int, error) {
	return _Gauge.Contract.DerivedSupply(&_Gauge.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address token, address account) view returns(uint256)
func (_Gauge *GaugeCaller) Earned(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "earned", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address token, address account) view returns(uint256)
func (_Gauge *GaugeSession) Earned(token common.Address, account common.Address) (*big.Int, error) {
	return _Gauge.Contract.Earned(&_Gauge.CallOpts, token, account)
}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address token, address account) view returns(uint256)
func (_Gauge *GaugeCallerSession) Earned(token common.Address, account common.Address) (*big.Int, error) {
	return _Gauge.Contract.Earned(&_Gauge.CallOpts, token, account)
}

// Fees0 is a free data retrieval call binding the contract method 0x93f1c442.
//
// Solidity: function fees0() view returns(uint256)
func (_Gauge *GaugeCaller) Fees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "fees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees0 is a free data retrieval call binding the contract method 0x93f1c442.
//
// Solidity: function fees0() view returns(uint256)
func (_Gauge *GaugeSession) Fees0() (*big.Int, error) {
	return _Gauge.Contract.Fees0(&_Gauge.CallOpts)
}

// Fees0 is a free data retrieval call binding the contract method 0x93f1c442.
//
// Solidity: function fees0() view returns(uint256)
func (_Gauge *GaugeCallerSession) Fees0() (*big.Int, error) {
	return _Gauge.Contract.Fees0(&_Gauge.CallOpts)
}

// Fees1 is a free data retrieval call binding the contract method 0x4c02a21c.
//
// Solidity: function fees1() view returns(uint256)
func (_Gauge *GaugeCaller) Fees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "fees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees1 is a free data retrieval call binding the contract method 0x4c02a21c.
//
// Solidity: function fees1() view returns(uint256)
func (_Gauge *GaugeSession) Fees1() (*big.Int, error) {
	return _Gauge.Contract.Fees1(&_Gauge.CallOpts)
}

// Fees1 is a free data retrieval call binding the contract method 0x4c02a21c.
//
// Solidity: function fees1() view returns(uint256)
func (_Gauge *GaugeCallerSession) Fees1() (*big.Int, error) {
	return _Gauge.Contract.Fees1(&_Gauge.CallOpts)
}

// GetPriorBalanceIndex is a free data retrieval call binding the contract method 0x115c6f39.
//
// Solidity: function getPriorBalanceIndex(address account, uint256 timestamp) view returns(uint256)
func (_Gauge *GaugeCaller) GetPriorBalanceIndex(opts *bind.CallOpts, account common.Address, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "getPriorBalanceIndex", account, timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorBalanceIndex is a free data retrieval call binding the contract method 0x115c6f39.
//
// Solidity: function getPriorBalanceIndex(address account, uint256 timestamp) view returns(uint256)
func (_Gauge *GaugeSession) GetPriorBalanceIndex(account common.Address, timestamp *big.Int) (*big.Int, error) {
	return _Gauge.Contract.GetPriorBalanceIndex(&_Gauge.CallOpts, account, timestamp)
}

// GetPriorBalanceIndex is a free data retrieval call binding the contract method 0x115c6f39.
//
// Solidity: function getPriorBalanceIndex(address account, uint256 timestamp) view returns(uint256)
func (_Gauge *GaugeCallerSession) GetPriorBalanceIndex(account common.Address, timestamp *big.Int) (*big.Int, error) {
	return _Gauge.Contract.GetPriorBalanceIndex(&_Gauge.CallOpts, account, timestamp)
}

// GetPriorRewardPerToken is a free data retrieval call binding the contract method 0xfd314098.
//
// Solidity: function getPriorRewardPerToken(address token, uint256 timestamp) view returns(uint256, uint256)
func (_Gauge *GaugeCaller) GetPriorRewardPerToken(opts *bind.CallOpts, token common.Address, timestamp *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "getPriorRewardPerToken", token, timestamp)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPriorRewardPerToken is a free data retrieval call binding the contract method 0xfd314098.
//
// Solidity: function getPriorRewardPerToken(address token, uint256 timestamp) view returns(uint256, uint256)
func (_Gauge *GaugeSession) GetPriorRewardPerToken(token common.Address, timestamp *big.Int) (*big.Int, *big.Int, error) {
	return _Gauge.Contract.GetPriorRewardPerToken(&_Gauge.CallOpts, token, timestamp)
}

// GetPriorRewardPerToken is a free data retrieval call binding the contract method 0xfd314098.
//
// Solidity: function getPriorRewardPerToken(address token, uint256 timestamp) view returns(uint256, uint256)
func (_Gauge *GaugeCallerSession) GetPriorRewardPerToken(token common.Address, timestamp *big.Int) (*big.Int, *big.Int, error) {
	return _Gauge.Contract.GetPriorRewardPerToken(&_Gauge.CallOpts, token, timestamp)
}

// GetPriorSupplyIndex is a free data retrieval call binding the contract method 0x76f4be36.
//
// Solidity: function getPriorSupplyIndex(uint256 timestamp) view returns(uint256)
func (_Gauge *GaugeCaller) GetPriorSupplyIndex(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "getPriorSupplyIndex", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorSupplyIndex is a free data retrieval call binding the contract method 0x76f4be36.
//
// Solidity: function getPriorSupplyIndex(uint256 timestamp) view returns(uint256)
func (_Gauge *GaugeSession) GetPriorSupplyIndex(timestamp *big.Int) (*big.Int, error) {
	return _Gauge.Contract.GetPriorSupplyIndex(&_Gauge.CallOpts, timestamp)
}

// GetPriorSupplyIndex is a free data retrieval call binding the contract method 0x76f4be36.
//
// Solidity: function getPriorSupplyIndex(uint256 timestamp) view returns(uint256)
func (_Gauge *GaugeCallerSession) GetPriorSupplyIndex(timestamp *big.Int) (*big.Int, error) {
	return _Gauge.Contract.GetPriorSupplyIndex(&_Gauge.CallOpts, timestamp)
}

// IsReward is a free data retrieval call binding the contract method 0x4d5ce038.
//
// Solidity: function isReward(address ) view returns(bool)
func (_Gauge *GaugeCaller) IsReward(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "isReward", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReward is a free data retrieval call binding the contract method 0x4d5ce038.
//
// Solidity: function isReward(address ) view returns(bool)
func (_Gauge *GaugeSession) IsReward(arg0 common.Address) (bool, error) {
	return _Gauge.Contract.IsReward(&_Gauge.CallOpts, arg0)
}

// IsReward is a free data retrieval call binding the contract method 0x4d5ce038.
//
// Solidity: function isReward(address ) view returns(bool)
func (_Gauge *GaugeCallerSession) IsReward(arg0 common.Address) (bool, error) {
	return _Gauge.Contract.IsReward(&_Gauge.CallOpts, arg0)
}

// LastEarn is a free data retrieval call binding the contract method 0xa495e5b5.
//
// Solidity: function lastEarn(address , address ) view returns(uint256)
func (_Gauge *GaugeCaller) LastEarn(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "lastEarn", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastEarn is a free data retrieval call binding the contract method 0xa495e5b5.
//
// Solidity: function lastEarn(address , address ) view returns(uint256)
func (_Gauge *GaugeSession) LastEarn(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Gauge.Contract.LastEarn(&_Gauge.CallOpts, arg0, arg1)
}

// LastEarn is a free data retrieval call binding the contract method 0xa495e5b5.
//
// Solidity: function lastEarn(address , address ) view returns(uint256)
func (_Gauge *GaugeCallerSession) LastEarn(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Gauge.Contract.LastEarn(&_Gauge.CallOpts, arg0, arg1)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address token) view returns(uint256)
func (_Gauge *GaugeCaller) LastTimeRewardApplicable(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "lastTimeRewardApplicable", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address token) view returns(uint256)
func (_Gauge *GaugeSession) LastTimeRewardApplicable(token common.Address) (*big.Int, error) {
	return _Gauge.Contract.LastTimeRewardApplicable(&_Gauge.CallOpts, token)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address token) view returns(uint256)
func (_Gauge *GaugeCallerSession) LastTimeRewardApplicable(token common.Address) (*big.Int, error) {
	return _Gauge.Contract.LastTimeRewardApplicable(&_Gauge.CallOpts, token)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0x2ce9aead.
//
// Solidity: function lastUpdateTime(address ) view returns(uint256)
func (_Gauge *GaugeCaller) LastUpdateTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "lastUpdateTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0x2ce9aead.
//
// Solidity: function lastUpdateTime(address ) view returns(uint256)
func (_Gauge *GaugeSession) LastUpdateTime(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.LastUpdateTime(&_Gauge.CallOpts, arg0)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0x2ce9aead.
//
// Solidity: function lastUpdateTime(address ) view returns(uint256)
func (_Gauge *GaugeCallerSession) LastUpdateTime(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.LastUpdateTime(&_Gauge.CallOpts, arg0)
}

// Left is a free data retrieval call binding the contract method 0x99bcc052.
//
// Solidity: function left(address token) view returns(uint256)
func (_Gauge *GaugeCaller) Left(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "left", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Left is a free data retrieval call binding the contract method 0x99bcc052.
//
// Solidity: function left(address token) view returns(uint256)
func (_Gauge *GaugeSession) Left(token common.Address) (*big.Int, error) {
	return _Gauge.Contract.Left(&_Gauge.CallOpts, token)
}

// Left is a free data retrieval call binding the contract method 0x99bcc052.
//
// Solidity: function left(address token) view returns(uint256)
func (_Gauge *GaugeCallerSession) Left(token common.Address) (*big.Int, error) {
	return _Gauge.Contract.Left(&_Gauge.CallOpts, token)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint256)
func (_Gauge *GaugeCaller) NumCheckpoints(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "numCheckpoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint256)
func (_Gauge *GaugeSession) NumCheckpoints(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.NumCheckpoints(&_Gauge.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint256)
func (_Gauge *GaugeCallerSession) NumCheckpoints(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.NumCheckpoints(&_Gauge.CallOpts, arg0)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xda09d19d.
//
// Solidity: function periodFinish(address ) view returns(uint256)
func (_Gauge *GaugeCaller) PeriodFinish(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "periodFinish", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xda09d19d.
//
// Solidity: function periodFinish(address ) view returns(uint256)
func (_Gauge *GaugeSession) PeriodFinish(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.PeriodFinish(&_Gauge.CallOpts, arg0)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xda09d19d.
//
// Solidity: function periodFinish(address ) view returns(uint256)
func (_Gauge *GaugeCallerSession) PeriodFinish(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.PeriodFinish(&_Gauge.CallOpts, arg0)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address token) view returns(uint256)
func (_Gauge *GaugeCaller) RewardPerToken(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "rewardPerToken", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address token) view returns(uint256)
func (_Gauge *GaugeSession) RewardPerToken(token common.Address) (*big.Int, error) {
	return _Gauge.Contract.RewardPerToken(&_Gauge.CallOpts, token)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address token) view returns(uint256)
func (_Gauge *GaugeCallerSession) RewardPerToken(token common.Address) (*big.Int, error) {
	return _Gauge.Contract.RewardPerToken(&_Gauge.CallOpts, token)
}

// RewardPerTokenCheckpoints is a free data retrieval call binding the contract method 0x01316ddf.
//
// Solidity: function rewardPerTokenCheckpoints(address , uint256 ) view returns(uint256 timestamp, uint256 rewardPerToken)
func (_Gauge *GaugeCaller) RewardPerTokenCheckpoints(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Timestamp      *big.Int
	RewardPerToken *big.Int
}, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "rewardPerTokenCheckpoints", arg0, arg1)

	outstruct := new(struct {
		Timestamp      *big.Int
		RewardPerToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardPerToken = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardPerTokenCheckpoints is a free data retrieval call binding the contract method 0x01316ddf.
//
// Solidity: function rewardPerTokenCheckpoints(address , uint256 ) view returns(uint256 timestamp, uint256 rewardPerToken)
func (_Gauge *GaugeSession) RewardPerTokenCheckpoints(arg0 common.Address, arg1 *big.Int) (struct {
	Timestamp      *big.Int
	RewardPerToken *big.Int
}, error) {
	return _Gauge.Contract.RewardPerTokenCheckpoints(&_Gauge.CallOpts, arg0, arg1)
}

// RewardPerTokenCheckpoints is a free data retrieval call binding the contract method 0x01316ddf.
//
// Solidity: function rewardPerTokenCheckpoints(address , uint256 ) view returns(uint256 timestamp, uint256 rewardPerToken)
func (_Gauge *GaugeCallerSession) RewardPerTokenCheckpoints(arg0 common.Address, arg1 *big.Int) (struct {
	Timestamp      *big.Int
	RewardPerToken *big.Int
}, error) {
	return _Gauge.Contract.RewardPerTokenCheckpoints(&_Gauge.CallOpts, arg0, arg1)
}

// RewardPerTokenNumCheckpoints is a free data retrieval call binding the contract method 0xaa479652.
//
// Solidity: function rewardPerTokenNumCheckpoints(address ) view returns(uint256)
func (_Gauge *GaugeCaller) RewardPerTokenNumCheckpoints(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "rewardPerTokenNumCheckpoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenNumCheckpoints is a free data retrieval call binding the contract method 0xaa479652.
//
// Solidity: function rewardPerTokenNumCheckpoints(address ) view returns(uint256)
func (_Gauge *GaugeSession) RewardPerTokenNumCheckpoints(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.RewardPerTokenNumCheckpoints(&_Gauge.CallOpts, arg0)
}

// RewardPerTokenNumCheckpoints is a free data retrieval call binding the contract method 0xaa479652.
//
// Solidity: function rewardPerTokenNumCheckpoints(address ) view returns(uint256)
func (_Gauge *GaugeCallerSession) RewardPerTokenNumCheckpoints(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.RewardPerTokenNumCheckpoints(&_Gauge.CallOpts, arg0)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0x9ce43f90.
//
// Solidity: function rewardPerTokenStored(address ) view returns(uint256)
func (_Gauge *GaugeCaller) RewardPerTokenStored(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "rewardPerTokenStored", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0x9ce43f90.
//
// Solidity: function rewardPerTokenStored(address ) view returns(uint256)
func (_Gauge *GaugeSession) RewardPerTokenStored(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.RewardPerTokenStored(&_Gauge.CallOpts, arg0)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0x9ce43f90.
//
// Solidity: function rewardPerTokenStored(address ) view returns(uint256)
func (_Gauge *GaugeCallerSession) RewardPerTokenStored(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.RewardPerTokenStored(&_Gauge.CallOpts, arg0)
}

// RewardRate is a free data retrieval call binding the contract method 0x221ca18c.
//
// Solidity: function rewardRate(address ) view returns(uint256)
func (_Gauge *GaugeCaller) RewardRate(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "rewardRate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x221ca18c.
//
// Solidity: function rewardRate(address ) view returns(uint256)
func (_Gauge *GaugeSession) RewardRate(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.RewardRate(&_Gauge.CallOpts, arg0)
}

// RewardRate is a free data retrieval call binding the contract method 0x221ca18c.
//
// Solidity: function rewardRate(address ) view returns(uint256)
func (_Gauge *GaugeCallerSession) RewardRate(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.RewardRate(&_Gauge.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(address)
func (_Gauge *GaugeCaller) Rewards(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(address)
func (_Gauge *GaugeSession) Rewards(arg0 *big.Int) (common.Address, error) {
	return _Gauge.Contract.Rewards(&_Gauge.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(address)
func (_Gauge *GaugeCallerSession) Rewards(arg0 *big.Int) (common.Address, error) {
	return _Gauge.Contract.Rewards(&_Gauge.CallOpts, arg0)
}

// RewardsListLength is a free data retrieval call binding the contract method 0xe6886396.
//
// Solidity: function rewardsListLength() view returns(uint256)
func (_Gauge *GaugeCaller) RewardsListLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "rewardsListLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsListLength is a free data retrieval call binding the contract method 0xe6886396.
//
// Solidity: function rewardsListLength() view returns(uint256)
func (_Gauge *GaugeSession) RewardsListLength() (*big.Int, error) {
	return _Gauge.Contract.RewardsListLength(&_Gauge.CallOpts)
}

// RewardsListLength is a free data retrieval call binding the contract method 0xe6886396.
//
// Solidity: function rewardsListLength() view returns(uint256)
func (_Gauge *GaugeCallerSession) RewardsListLength() (*big.Int, error) {
	return _Gauge.Contract.RewardsListLength(&_Gauge.CallOpts)
}

// Stake is a free data retrieval call binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() view returns(address)
func (_Gauge *GaugeCaller) Stake(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "stake")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stake is a free data retrieval call binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() view returns(address)
func (_Gauge *GaugeSession) Stake() (common.Address, error) {
	return _Gauge.Contract.Stake(&_Gauge.CallOpts)
}

// Stake is a free data retrieval call binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() view returns(address)
func (_Gauge *GaugeCallerSession) Stake() (common.Address, error) {
	return _Gauge.Contract.Stake(&_Gauge.CallOpts)
}

// SupplyCheckpoints is a free data retrieval call binding the contract method 0xf7412baf.
//
// Solidity: function supplyCheckpoints(uint256 ) view returns(uint256 timestamp, uint256 supply)
func (_Gauge *GaugeCaller) SupplyCheckpoints(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp *big.Int
	Supply    *big.Int
}, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "supplyCheckpoints", arg0)

	outstruct := new(struct {
		Timestamp *big.Int
		Supply    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Supply = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SupplyCheckpoints is a free data retrieval call binding the contract method 0xf7412baf.
//
// Solidity: function supplyCheckpoints(uint256 ) view returns(uint256 timestamp, uint256 supply)
func (_Gauge *GaugeSession) SupplyCheckpoints(arg0 *big.Int) (struct {
	Timestamp *big.Int
	Supply    *big.Int
}, error) {
	return _Gauge.Contract.SupplyCheckpoints(&_Gauge.CallOpts, arg0)
}

// SupplyCheckpoints is a free data retrieval call binding the contract method 0xf7412baf.
//
// Solidity: function supplyCheckpoints(uint256 ) view returns(uint256 timestamp, uint256 supply)
func (_Gauge *GaugeCallerSession) SupplyCheckpoints(arg0 *big.Int) (struct {
	Timestamp *big.Int
	Supply    *big.Int
}, error) {
	return _Gauge.Contract.SupplyCheckpoints(&_Gauge.CallOpts, arg0)
}

// SupplyNumCheckpoints is a free data retrieval call binding the contract method 0xe8111a12.
//
// Solidity: function supplyNumCheckpoints() view returns(uint256)
func (_Gauge *GaugeCaller) SupplyNumCheckpoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "supplyNumCheckpoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyNumCheckpoints is a free data retrieval call binding the contract method 0xe8111a12.
//
// Solidity: function supplyNumCheckpoints() view returns(uint256)
func (_Gauge *GaugeSession) SupplyNumCheckpoints() (*big.Int, error) {
	return _Gauge.Contract.SupplyNumCheckpoints(&_Gauge.CallOpts)
}

// SupplyNumCheckpoints is a free data retrieval call binding the contract method 0xe8111a12.
//
// Solidity: function supplyNumCheckpoints() view returns(uint256)
func (_Gauge *GaugeCallerSession) SupplyNumCheckpoints() (*big.Int, error) {
	return _Gauge.Contract.SupplyNumCheckpoints(&_Gauge.CallOpts)
}

// TokenIds is a free data retrieval call binding the contract method 0xfc97a303.
//
// Solidity: function tokenIds(address ) view returns(uint256)
func (_Gauge *GaugeCaller) TokenIds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "tokenIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIds is a free data retrieval call binding the contract method 0xfc97a303.
//
// Solidity: function tokenIds(address ) view returns(uint256)
func (_Gauge *GaugeSession) TokenIds(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.TokenIds(&_Gauge.CallOpts, arg0)
}

// TokenIds is a free data retrieval call binding the contract method 0xfc97a303.
//
// Solidity: function tokenIds(address ) view returns(uint256)
func (_Gauge *GaugeCallerSession) TokenIds(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.TokenIds(&_Gauge.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gauge *GaugeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gauge *GaugeSession) TotalSupply() (*big.Int, error) {
	return _Gauge.Contract.TotalSupply(&_Gauge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gauge *GaugeCallerSession) TotalSupply() (*big.Int, error) {
	return _Gauge.Contract.TotalSupply(&_Gauge.CallOpts)
}

// UserRewardPerTokenStored is a free data retrieval call binding the contract method 0x3ca068b6.
//
// Solidity: function userRewardPerTokenStored(address , address ) view returns(uint256)
func (_Gauge *GaugeCaller) UserRewardPerTokenStored(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "userRewardPerTokenStored", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenStored is a free data retrieval call binding the contract method 0x3ca068b6.
//
// Solidity: function userRewardPerTokenStored(address , address ) view returns(uint256)
func (_Gauge *GaugeSession) UserRewardPerTokenStored(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Gauge.Contract.UserRewardPerTokenStored(&_Gauge.CallOpts, arg0, arg1)
}

// UserRewardPerTokenStored is a free data retrieval call binding the contract method 0x3ca068b6.
//
// Solidity: function userRewardPerTokenStored(address , address ) view returns(uint256)
func (_Gauge *GaugeCallerSession) UserRewardPerTokenStored(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Gauge.Contract.UserRewardPerTokenStored(&_Gauge.CallOpts, arg0, arg1)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_Gauge *GaugeCaller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_Gauge *GaugeSession) Voter() (common.Address, error) {
	return _Gauge.Contract.Voter(&_Gauge.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_Gauge *GaugeCallerSession) Voter() (common.Address, error) {
	return _Gauge.Contract.Voter(&_Gauge.CallOpts)
}

// BatchRewardPerToken is a paid mutator transaction binding the contract method 0x5a45d052.
//
// Solidity: function batchRewardPerToken(address token, uint256 maxRuns) returns()
func (_Gauge *GaugeTransactor) BatchRewardPerToken(opts *bind.TransactOpts, token common.Address, maxRuns *big.Int) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "batchRewardPerToken", token, maxRuns)
}

// BatchRewardPerToken is a paid mutator transaction binding the contract method 0x5a45d052.
//
// Solidity: function batchRewardPerToken(address token, uint256 maxRuns) returns()
func (_Gauge *GaugeSession) BatchRewardPerToken(token common.Address, maxRuns *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.BatchRewardPerToken(&_Gauge.TransactOpts, token, maxRuns)
}

// BatchRewardPerToken is a paid mutator transaction binding the contract method 0x5a45d052.
//
// Solidity: function batchRewardPerToken(address token, uint256 maxRuns) returns()
func (_Gauge *GaugeTransactorSession) BatchRewardPerToken(token common.Address, maxRuns *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.BatchRewardPerToken(&_Gauge.TransactOpts, token, maxRuns)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns(uint256 claimed0, uint256 claimed1)
func (_Gauge *GaugeTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns(uint256 claimed0, uint256 claimed1)
func (_Gauge *GaugeSession) ClaimFees() (*types.Transaction, error) {
	return _Gauge.Contract.ClaimFees(&_Gauge.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns(uint256 claimed0, uint256 claimed1)
func (_Gauge *GaugeTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _Gauge.Contract.ClaimFees(&_Gauge.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 amount, uint256 tokenId) returns()
func (_Gauge *GaugeTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "deposit", amount, tokenId)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 amount, uint256 tokenId) returns()
func (_Gauge *GaugeSession) Deposit(amount *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.Deposit(&_Gauge.TransactOpts, amount, tokenId)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 amount, uint256 tokenId) returns()
func (_Gauge *GaugeTransactorSession) Deposit(amount *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.Deposit(&_Gauge.TransactOpts, amount, tokenId)
}

// DepositAll is a paid mutator transaction binding the contract method 0xc6f678bd.
//
// Solidity: function depositAll(uint256 tokenId) returns()
func (_Gauge *GaugeTransactor) DepositAll(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "depositAll", tokenId)
}

// DepositAll is a paid mutator transaction binding the contract method 0xc6f678bd.
//
// Solidity: function depositAll(uint256 tokenId) returns()
func (_Gauge *GaugeSession) DepositAll(tokenId *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.DepositAll(&_Gauge.TransactOpts, tokenId)
}

// DepositAll is a paid mutator transaction binding the contract method 0xc6f678bd.
//
// Solidity: function depositAll(uint256 tokenId) returns()
func (_Gauge *GaugeTransactorSession) DepositAll(tokenId *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.DepositAll(&_Gauge.TransactOpts, tokenId)
}

// GetReward is a paid mutator transaction binding the contract method 0x31279d3d.
//
// Solidity: function getReward(address account, address[] tokens) returns()
func (_Gauge *GaugeTransactor) GetReward(opts *bind.TransactOpts, account common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "getReward", account, tokens)
}

// GetReward is a paid mutator transaction binding the contract method 0x31279d3d.
//
// Solidity: function getReward(address account, address[] tokens) returns()
func (_Gauge *GaugeSession) GetReward(account common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _Gauge.Contract.GetReward(&_Gauge.TransactOpts, account, tokens)
}

// GetReward is a paid mutator transaction binding the contract method 0x31279d3d.
//
// Solidity: function getReward(address account, address[] tokens) returns()
func (_Gauge *GaugeTransactorSession) GetReward(account common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _Gauge.Contract.GetReward(&_Gauge.TransactOpts, account, tokens)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address token, uint256 amount) returns()
func (_Gauge *GaugeTransactor) NotifyRewardAmount(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "notifyRewardAmount", token, amount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address token, uint256 amount) returns()
func (_Gauge *GaugeSession) NotifyRewardAmount(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.NotifyRewardAmount(&_Gauge.TransactOpts, token, amount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address token, uint256 amount) returns()
func (_Gauge *GaugeTransactorSession) NotifyRewardAmount(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.NotifyRewardAmount(&_Gauge.TransactOpts, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Gauge *GaugeTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Gauge *GaugeSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.Withdraw(&_Gauge.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Gauge *GaugeTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.Withdraw(&_Gauge.TransactOpts, amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Gauge *GaugeTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Gauge *GaugeSession) WithdrawAll() (*types.Transaction, error) {
	return _Gauge.Contract.WithdrawAll(&_Gauge.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Gauge *GaugeTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _Gauge.Contract.WithdrawAll(&_Gauge.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0xfdb483c7.
//
// Solidity: function withdrawToken(uint256 amount, uint256 tokenId) returns()
func (_Gauge *GaugeTransactor) WithdrawToken(opts *bind.TransactOpts, amount *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "withdrawToken", amount, tokenId)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0xfdb483c7.
//
// Solidity: function withdrawToken(uint256 amount, uint256 tokenId) returns()
func (_Gauge *GaugeSession) WithdrawToken(amount *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.WithdrawToken(&_Gauge.TransactOpts, amount, tokenId)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0xfdb483c7.
//
// Solidity: function withdrawToken(uint256 amount, uint256 tokenId) returns()
func (_Gauge *GaugeTransactorSession) WithdrawToken(amount *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.WithdrawToken(&_Gauge.TransactOpts, amount, tokenId)
}

// GaugeClaimFeesIterator is returned from FilterClaimFees and is used to iterate over the raw logs and unpacked data for ClaimFees events raised by the Gauge contract.
type GaugeClaimFeesIterator struct {
	Event *GaugeClaimFees // Event containing the contract specifics and raw log

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
func (it *GaugeClaimFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GaugeClaimFees)
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
		it.Event = new(GaugeClaimFees)
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
func (it *GaugeClaimFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GaugeClaimFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GaugeClaimFees represents a ClaimFees event raised by the Gauge contract.
type GaugeClaimFees struct {
	From     common.Address
	Claimed0 *big.Int
	Claimed1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimFees is a free log retrieval operation binding the contract event 0xbc567d6cbad26368064baa0ab5a757be46aae4d70f707f9203d9d9b6c8ccbfa3.
//
// Solidity: event ClaimFees(address indexed from, uint256 claimed0, uint256 claimed1)
func (_Gauge *GaugeFilterer) FilterClaimFees(opts *bind.FilterOpts, from []common.Address) (*GaugeClaimFeesIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Gauge.contract.FilterLogs(opts, "ClaimFees", fromRule)
	if err != nil {
		return nil, err
	}
	return &GaugeClaimFeesIterator{contract: _Gauge.contract, event: "ClaimFees", logs: logs, sub: sub}, nil
}

// WatchClaimFees is a free log subscription operation binding the contract event 0xbc567d6cbad26368064baa0ab5a757be46aae4d70f707f9203d9d9b6c8ccbfa3.
//
// Solidity: event ClaimFees(address indexed from, uint256 claimed0, uint256 claimed1)
func (_Gauge *GaugeFilterer) WatchClaimFees(opts *bind.WatchOpts, sink chan<- *GaugeClaimFees, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Gauge.contract.WatchLogs(opts, "ClaimFees", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GaugeClaimFees)
				if err := _Gauge.contract.UnpackLog(event, "ClaimFees", log); err != nil {
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

// ParseClaimFees is a log parse operation binding the contract event 0xbc567d6cbad26368064baa0ab5a757be46aae4d70f707f9203d9d9b6c8ccbfa3.
//
// Solidity: event ClaimFees(address indexed from, uint256 claimed0, uint256 claimed1)
func (_Gauge *GaugeFilterer) ParseClaimFees(log types.Log) (*GaugeClaimFees, error) {
	event := new(GaugeClaimFees)
	if err := _Gauge.contract.UnpackLog(event, "ClaimFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GaugeClaimRewardsIterator is returned from FilterClaimRewards and is used to iterate over the raw logs and unpacked data for ClaimRewards events raised by the Gauge contract.
type GaugeClaimRewardsIterator struct {
	Event *GaugeClaimRewards // Event containing the contract specifics and raw log

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
func (it *GaugeClaimRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GaugeClaimRewards)
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
		it.Event = new(GaugeClaimRewards)
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
func (it *GaugeClaimRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GaugeClaimRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GaugeClaimRewards represents a ClaimRewards event raised by the Gauge contract.
type GaugeClaimRewards struct {
	From   common.Address
	Reward common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimRewards is a free log retrieval operation binding the contract event 0x9aa05b3d70a9e3e2f004f039648839560576334fb45c81f91b6db03ad9e2efc9.
//
// Solidity: event ClaimRewards(address indexed from, address indexed reward, uint256 amount)
func (_Gauge *GaugeFilterer) FilterClaimRewards(opts *bind.FilterOpts, from []common.Address, reward []common.Address) (*GaugeClaimRewardsIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _Gauge.contract.FilterLogs(opts, "ClaimRewards", fromRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &GaugeClaimRewardsIterator{contract: _Gauge.contract, event: "ClaimRewards", logs: logs, sub: sub}, nil
}

// WatchClaimRewards is a free log subscription operation binding the contract event 0x9aa05b3d70a9e3e2f004f039648839560576334fb45c81f91b6db03ad9e2efc9.
//
// Solidity: event ClaimRewards(address indexed from, address indexed reward, uint256 amount)
func (_Gauge *GaugeFilterer) WatchClaimRewards(opts *bind.WatchOpts, sink chan<- *GaugeClaimRewards, from []common.Address, reward []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _Gauge.contract.WatchLogs(opts, "ClaimRewards", fromRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GaugeClaimRewards)
				if err := _Gauge.contract.UnpackLog(event, "ClaimRewards", log); err != nil {
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

// ParseClaimRewards is a log parse operation binding the contract event 0x9aa05b3d70a9e3e2f004f039648839560576334fb45c81f91b6db03ad9e2efc9.
//
// Solidity: event ClaimRewards(address indexed from, address indexed reward, uint256 amount)
func (_Gauge *GaugeFilterer) ParseClaimRewards(log types.Log) (*GaugeClaimRewards, error) {
	event := new(GaugeClaimRewards)
	if err := _Gauge.contract.UnpackLog(event, "ClaimRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GaugeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Gauge contract.
type GaugeDepositIterator struct {
	Event *GaugeDeposit // Event containing the contract specifics and raw log

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
func (it *GaugeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GaugeDeposit)
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
		it.Event = new(GaugeDeposit)
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
func (it *GaugeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GaugeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GaugeDeposit represents a Deposit event raised by the Gauge contract.
type GaugeDeposit struct {
	From    common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed from, uint256 tokenId, uint256 amount)
func (_Gauge *GaugeFilterer) FilterDeposit(opts *bind.FilterOpts, from []common.Address) (*GaugeDepositIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Gauge.contract.FilterLogs(opts, "Deposit", fromRule)
	if err != nil {
		return nil, err
	}
	return &GaugeDepositIterator{contract: _Gauge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed from, uint256 tokenId, uint256 amount)
func (_Gauge *GaugeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *GaugeDeposit, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Gauge.contract.WatchLogs(opts, "Deposit", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GaugeDeposit)
				if err := _Gauge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed from, uint256 tokenId, uint256 amount)
func (_Gauge *GaugeFilterer) ParseDeposit(log types.Log) (*GaugeDeposit, error) {
	event := new(GaugeDeposit)
	if err := _Gauge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GaugeNotifyRewardIterator is returned from FilterNotifyReward and is used to iterate over the raw logs and unpacked data for NotifyReward events raised by the Gauge contract.
type GaugeNotifyRewardIterator struct {
	Event *GaugeNotifyReward // Event containing the contract specifics and raw log

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
func (it *GaugeNotifyRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GaugeNotifyReward)
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
		it.Event = new(GaugeNotifyReward)
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
func (it *GaugeNotifyRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GaugeNotifyRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GaugeNotifyReward represents a NotifyReward event raised by the Gauge contract.
type GaugeNotifyReward struct {
	From   common.Address
	Reward common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotifyReward is a free log retrieval operation binding the contract event 0xf70d5c697de7ea828df48e5c4573cb2194c659f1901f70110c52b066dcf50826.
//
// Solidity: event NotifyReward(address indexed from, address indexed reward, uint256 amount)
func (_Gauge *GaugeFilterer) FilterNotifyReward(opts *bind.FilterOpts, from []common.Address, reward []common.Address) (*GaugeNotifyRewardIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _Gauge.contract.FilterLogs(opts, "NotifyReward", fromRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &GaugeNotifyRewardIterator{contract: _Gauge.contract, event: "NotifyReward", logs: logs, sub: sub}, nil
}

// WatchNotifyReward is a free log subscription operation binding the contract event 0xf70d5c697de7ea828df48e5c4573cb2194c659f1901f70110c52b066dcf50826.
//
// Solidity: event NotifyReward(address indexed from, address indexed reward, uint256 amount)
func (_Gauge *GaugeFilterer) WatchNotifyReward(opts *bind.WatchOpts, sink chan<- *GaugeNotifyReward, from []common.Address, reward []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _Gauge.contract.WatchLogs(opts, "NotifyReward", fromRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GaugeNotifyReward)
				if err := _Gauge.contract.UnpackLog(event, "NotifyReward", log); err != nil {
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

// ParseNotifyReward is a log parse operation binding the contract event 0xf70d5c697de7ea828df48e5c4573cb2194c659f1901f70110c52b066dcf50826.
//
// Solidity: event NotifyReward(address indexed from, address indexed reward, uint256 amount)
func (_Gauge *GaugeFilterer) ParseNotifyReward(log types.Log) (*GaugeNotifyReward, error) {
	event := new(GaugeNotifyReward)
	if err := _Gauge.contract.UnpackLog(event, "NotifyReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GaugeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Gauge contract.
type GaugeWithdrawIterator struct {
	Event *GaugeWithdraw // Event containing the contract specifics and raw log

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
func (it *GaugeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GaugeWithdraw)
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
		it.Event = new(GaugeWithdraw)
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
func (it *GaugeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GaugeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GaugeWithdraw represents a Withdraw event raised by the Gauge contract.
type GaugeWithdraw struct {
	From    common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed from, uint256 tokenId, uint256 amount)
func (_Gauge *GaugeFilterer) FilterWithdraw(opts *bind.FilterOpts, from []common.Address) (*GaugeWithdrawIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Gauge.contract.FilterLogs(opts, "Withdraw", fromRule)
	if err != nil {
		return nil, err
	}
	return &GaugeWithdrawIterator{contract: _Gauge.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed from, uint256 tokenId, uint256 amount)
func (_Gauge *GaugeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *GaugeWithdraw, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Gauge.contract.WatchLogs(opts, "Withdraw", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GaugeWithdraw)
				if err := _Gauge.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed from, uint256 tokenId, uint256 amount)
func (_Gauge *GaugeFilterer) ParseWithdraw(log types.Log) (*GaugeWithdraw, error) {
	event := new(GaugeWithdraw)
	if err := _Gauge.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
