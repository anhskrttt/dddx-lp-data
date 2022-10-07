// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voter

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

// VoterMetaData contains all meta data concerning the Voter contract.
var VoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__ve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gauges\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bribes\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"weight\",\"type\":\"int256\"}],\"name\":\"Abstained\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Attach\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Detach\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributeReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bribe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"GaugeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NotifyReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"weight\",\"type\":\"int256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"whitelister\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"Whitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_ve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"attachTokenToGauge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bribefactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bribes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_bribes\",\"type\":\"address[]\"},{\"internalType\":\"address[][]\",\"name\":\"_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"claimBribes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_fees\",\"type\":\"address[]\"},{\"internalType\":\"address[][]\",\"name\":\"_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"},{\"internalType\":\"address[][]\",\"name\":\"_tokens\",\"type\":\"address[][]\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"createGauge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"detachTokenFromGauge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finish\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"}],\"name\":\"distributeFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distro\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emitDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emitWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gaugefactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gauges\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isGauge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listing_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"poke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolForGauge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolVote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"}],\"name\":\"updateFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"updateForRange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"updateGauge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_poolVote\",\"type\":\"address[]\"},{\"internalType\":\"int256[]\",\"name\":\"_weights\",\"type\":\"int256[]\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"weights\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"whitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VoterABI is the input ABI used to generate the binding from.
// Deprecated: Use VoterMetaData.ABI instead.
var VoterABI = VoterMetaData.ABI

// Voter is an auto generated Go binding around an Ethereum contract.
type Voter struct {
	VoterCaller     // Read-only binding to the contract
	VoterTransactor // Write-only binding to the contract
	VoterFilterer   // Log filterer for contract events
}

// VoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoterSession struct {
	Contract     *Voter            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoterCallerSession struct {
	Contract *VoterCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoterTransactorSession struct {
	Contract     *VoterTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoterRaw struct {
	Contract *Voter // Generic contract binding to access the raw methods on
}

// VoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoterCallerRaw struct {
	Contract *VoterCaller // Generic read-only contract binding to access the raw methods on
}

// VoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoterTransactorRaw struct {
	Contract *VoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoter creates a new instance of Voter, bound to a specific deployed contract.
func NewVoter(address common.Address, backend bind.ContractBackend) (*Voter, error) {
	contract, err := bindVoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voter{VoterCaller: VoterCaller{contract: contract}, VoterTransactor: VoterTransactor{contract: contract}, VoterFilterer: VoterFilterer{contract: contract}}, nil
}

// NewVoterCaller creates a new read-only instance of Voter, bound to a specific deployed contract.
func NewVoterCaller(address common.Address, caller bind.ContractCaller) (*VoterCaller, error) {
	contract, err := bindVoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoterCaller{contract: contract}, nil
}

// NewVoterTransactor creates a new write-only instance of Voter, bound to a specific deployed contract.
func NewVoterTransactor(address common.Address, transactor bind.ContractTransactor) (*VoterTransactor, error) {
	contract, err := bindVoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoterTransactor{contract: contract}, nil
}

// NewVoterFilterer creates a new log filterer instance of Voter, bound to a specific deployed contract.
func NewVoterFilterer(address common.Address, filterer bind.ContractFilterer) (*VoterFilterer, error) {
	contract, err := bindVoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoterFilterer{contract: contract}, nil
}

// bindVoter binds a generic wrapper to an already deployed contract.
func bindVoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voter *VoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voter.Contract.VoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voter *VoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voter.Contract.VoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voter *VoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voter.Contract.VoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voter *VoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voter *VoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voter *VoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voter.Contract.contract.Transact(opts, method, params...)
}

// Ve is a free data retrieval call binding the contract method 0x8dd598fb.
//
// Solidity: function _ve() view returns(address)
func (_Voter *VoterCaller) Ve(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "_ve")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ve is a free data retrieval call binding the contract method 0x8dd598fb.
//
// Solidity: function _ve() view returns(address)
func (_Voter *VoterSession) Ve() (common.Address, error) {
	return _Voter.Contract.Ve(&_Voter.CallOpts)
}

// Ve is a free data retrieval call binding the contract method 0x8dd598fb.
//
// Solidity: function _ve() view returns(address)
func (_Voter *VoterCallerSession) Ve() (common.Address, error) {
	return _Voter.Contract.Ve(&_Voter.CallOpts)
}

// Bribefactory is a free data retrieval call binding the contract method 0x38752a9d.
//
// Solidity: function bribefactory() view returns(address)
func (_Voter *VoterCaller) Bribefactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "bribefactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bribefactory is a free data retrieval call binding the contract method 0x38752a9d.
//
// Solidity: function bribefactory() view returns(address)
func (_Voter *VoterSession) Bribefactory() (common.Address, error) {
	return _Voter.Contract.Bribefactory(&_Voter.CallOpts)
}

// Bribefactory is a free data retrieval call binding the contract method 0x38752a9d.
//
// Solidity: function bribefactory() view returns(address)
func (_Voter *VoterCallerSession) Bribefactory() (common.Address, error) {
	return _Voter.Contract.Bribefactory(&_Voter.CallOpts)
}

// Bribes is a free data retrieval call binding the contract method 0xa8c5d95a.
//
// Solidity: function bribes(address ) view returns(address)
func (_Voter *VoterCaller) Bribes(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "bribes", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bribes is a free data retrieval call binding the contract method 0xa8c5d95a.
//
// Solidity: function bribes(address ) view returns(address)
func (_Voter *VoterSession) Bribes(arg0 common.Address) (common.Address, error) {
	return _Voter.Contract.Bribes(&_Voter.CallOpts, arg0)
}

// Bribes is a free data retrieval call binding the contract method 0xa8c5d95a.
//
// Solidity: function bribes(address ) view returns(address)
func (_Voter *VoterCallerSession) Bribes(arg0 common.Address) (common.Address, error) {
	return _Voter.Contract.Bribes(&_Voter.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_Voter *VoterCaller) Claimable(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "claimable", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_Voter *VoterSession) Claimable(arg0 common.Address) (*big.Int, error) {
	return _Voter.Contract.Claimable(&_Voter.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_Voter *VoterCallerSession) Claimable(arg0 common.Address) (*big.Int, error) {
	return _Voter.Contract.Claimable(&_Voter.CallOpts, arg0)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Voter *VoterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Voter *VoterSession) Factory() (common.Address, error) {
	return _Voter.Contract.Factory(&_Voter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Voter *VoterCallerSession) Factory() (common.Address, error) {
	return _Voter.Contract.Factory(&_Voter.CallOpts)
}

// Gaugefactory is a free data retrieval call binding the contract method 0x68c3acb3.
//
// Solidity: function gaugefactory() view returns(address)
func (_Voter *VoterCaller) Gaugefactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "gaugefactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gaugefactory is a free data retrieval call binding the contract method 0x68c3acb3.
//
// Solidity: function gaugefactory() view returns(address)
func (_Voter *VoterSession) Gaugefactory() (common.Address, error) {
	return _Voter.Contract.Gaugefactory(&_Voter.CallOpts)
}

// Gaugefactory is a free data retrieval call binding the contract method 0x68c3acb3.
//
// Solidity: function gaugefactory() view returns(address)
func (_Voter *VoterCallerSession) Gaugefactory() (common.Address, error) {
	return _Voter.Contract.Gaugefactory(&_Voter.CallOpts)
}

// Gauges is a free data retrieval call binding the contract method 0xb9a09fd5.
//
// Solidity: function gauges(address ) view returns(address)
func (_Voter *VoterCaller) Gauges(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "gauges", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gauges is a free data retrieval call binding the contract method 0xb9a09fd5.
//
// Solidity: function gauges(address ) view returns(address)
func (_Voter *VoterSession) Gauges(arg0 common.Address) (common.Address, error) {
	return _Voter.Contract.Gauges(&_Voter.CallOpts, arg0)
}

// Gauges is a free data retrieval call binding the contract method 0xb9a09fd5.
//
// Solidity: function gauges(address ) view returns(address)
func (_Voter *VoterCallerSession) Gauges(arg0 common.Address) (common.Address, error) {
	return _Voter.Contract.Gauges(&_Voter.CallOpts, arg0)
}

// IsGauge is a free data retrieval call binding the contract method 0xaa79979b.
//
// Solidity: function isGauge(address ) view returns(bool)
func (_Voter *VoterCaller) IsGauge(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "isGauge", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGauge is a free data retrieval call binding the contract method 0xaa79979b.
//
// Solidity: function isGauge(address ) view returns(bool)
func (_Voter *VoterSession) IsGauge(arg0 common.Address) (bool, error) {
	return _Voter.Contract.IsGauge(&_Voter.CallOpts, arg0)
}

// IsGauge is a free data retrieval call binding the contract method 0xaa79979b.
//
// Solidity: function isGauge(address ) view returns(bool)
func (_Voter *VoterCallerSession) IsGauge(arg0 common.Address) (bool, error) {
	return _Voter.Contract.IsGauge(&_Voter.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_Voter *VoterCaller) IsWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "isWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_Voter *VoterSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _Voter.Contract.IsWhitelisted(&_Voter.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_Voter *VoterCallerSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _Voter.Contract.IsWhitelisted(&_Voter.CallOpts, arg0)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_Voter *VoterCaller) Length(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "length")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_Voter *VoterSession) Length() (*big.Int, error) {
	return _Voter.Contract.Length(&_Voter.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_Voter *VoterCallerSession) Length() (*big.Int, error) {
	return _Voter.Contract.Length(&_Voter.CallOpts)
}

// ListingFee is a free data retrieval call binding the contract method 0xb980777a.
//
// Solidity: function listing_fee() view returns(uint256)
func (_Voter *VoterCaller) ListingFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "listing_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ListingFee is a free data retrieval call binding the contract method 0xb980777a.
//
// Solidity: function listing_fee() view returns(uint256)
func (_Voter *VoterSession) ListingFee() (*big.Int, error) {
	return _Voter.Contract.ListingFee(&_Voter.CallOpts)
}

// ListingFee is a free data retrieval call binding the contract method 0xb980777a.
//
// Solidity: function listing_fee() view returns(uint256)
func (_Voter *VoterCallerSession) ListingFee() (*big.Int, error) {
	return _Voter.Contract.ListingFee(&_Voter.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Voter *VoterCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Voter *VoterSession) Minter() (common.Address, error) {
	return _Voter.Contract.Minter(&_Voter.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Voter *VoterCallerSession) Minter() (common.Address, error) {
	return _Voter.Contract.Minter(&_Voter.CallOpts)
}

// PoolForGauge is a free data retrieval call binding the contract method 0x06d6a1b2.
//
// Solidity: function poolForGauge(address ) view returns(address)
func (_Voter *VoterCaller) PoolForGauge(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "poolForGauge", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolForGauge is a free data retrieval call binding the contract method 0x06d6a1b2.
//
// Solidity: function poolForGauge(address ) view returns(address)
func (_Voter *VoterSession) PoolForGauge(arg0 common.Address) (common.Address, error) {
	return _Voter.Contract.PoolForGauge(&_Voter.CallOpts, arg0)
}

// PoolForGauge is a free data retrieval call binding the contract method 0x06d6a1b2.
//
// Solidity: function poolForGauge(address ) view returns(address)
func (_Voter *VoterCallerSession) PoolForGauge(arg0 common.Address) (common.Address, error) {
	return _Voter.Contract.PoolForGauge(&_Voter.CallOpts, arg0)
}

// PoolVote is a free data retrieval call binding the contract method 0xa86a366d.
//
// Solidity: function poolVote(uint256 , uint256 ) view returns(address)
func (_Voter *VoterCaller) PoolVote(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "poolVote", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolVote is a free data retrieval call binding the contract method 0xa86a366d.
//
// Solidity: function poolVote(uint256 , uint256 ) view returns(address)
func (_Voter *VoterSession) PoolVote(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Voter.Contract.PoolVote(&_Voter.CallOpts, arg0, arg1)
}

// PoolVote is a free data retrieval call binding the contract method 0xa86a366d.
//
// Solidity: function poolVote(uint256 , uint256 ) view returns(address)
func (_Voter *VoterCallerSession) PoolVote(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Voter.Contract.PoolVote(&_Voter.CallOpts, arg0, arg1)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_Voter *VoterCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "pools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_Voter *VoterSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _Voter.Contract.Pools(&_Voter.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_Voter *VoterCallerSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _Voter.Contract.Pools(&_Voter.CallOpts, arg0)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_Voter *VoterCaller) TotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "totalWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_Voter *VoterSession) TotalWeight() (*big.Int, error) {
	return _Voter.Contract.TotalWeight(&_Voter.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_Voter *VoterCallerSession) TotalWeight() (*big.Int, error) {
	return _Voter.Contract.TotalWeight(&_Voter.CallOpts)
}

// UsedWeights is a free data retrieval call binding the contract method 0x79e93824.
//
// Solidity: function usedWeights(uint256 ) view returns(uint256)
func (_Voter *VoterCaller) UsedWeights(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "usedWeights", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsedWeights is a free data retrieval call binding the contract method 0x79e93824.
//
// Solidity: function usedWeights(uint256 ) view returns(uint256)
func (_Voter *VoterSession) UsedWeights(arg0 *big.Int) (*big.Int, error) {
	return _Voter.Contract.UsedWeights(&_Voter.CallOpts, arg0)
}

// UsedWeights is a free data retrieval call binding the contract method 0x79e93824.
//
// Solidity: function usedWeights(uint256 ) view returns(uint256)
func (_Voter *VoterCallerSession) UsedWeights(arg0 *big.Int) (*big.Int, error) {
	return _Voter.Contract.UsedWeights(&_Voter.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd23254b4.
//
// Solidity: function votes(uint256 , address ) view returns(int256)
func (_Voter *VoterCaller) Votes(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "votes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xd23254b4.
//
// Solidity: function votes(uint256 , address ) view returns(int256)
func (_Voter *VoterSession) Votes(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Voter.Contract.Votes(&_Voter.CallOpts, arg0, arg1)
}

// Votes is a free data retrieval call binding the contract method 0xd23254b4.
//
// Solidity: function votes(uint256 , address ) view returns(int256)
func (_Voter *VoterCallerSession) Votes(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Voter.Contract.Votes(&_Voter.CallOpts, arg0, arg1)
}

// Weights is a free data retrieval call binding the contract method 0xa7cac846.
//
// Solidity: function weights(address ) view returns(int256)
func (_Voter *VoterCaller) Weights(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "weights", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Weights is a free data retrieval call binding the contract method 0xa7cac846.
//
// Solidity: function weights(address ) view returns(int256)
func (_Voter *VoterSession) Weights(arg0 common.Address) (*big.Int, error) {
	return _Voter.Contract.Weights(&_Voter.CallOpts, arg0)
}

// Weights is a free data retrieval call binding the contract method 0xa7cac846.
//
// Solidity: function weights(address ) view returns(int256)
func (_Voter *VoterCallerSession) Weights(arg0 common.Address) (*big.Int, error) {
	return _Voter.Contract.Weights(&_Voter.CallOpts, arg0)
}

// AttachTokenToGauge is a paid mutator transaction binding the contract method 0x698473e3.
//
// Solidity: function attachTokenToGauge(uint256 tokenId, address account) returns()
func (_Voter *VoterTransactor) AttachTokenToGauge(opts *bind.TransactOpts, tokenId *big.Int, account common.Address) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "attachTokenToGauge", tokenId, account)
}

// AttachTokenToGauge is a paid mutator transaction binding the contract method 0x698473e3.
//
// Solidity: function attachTokenToGauge(uint256 tokenId, address account) returns()
func (_Voter *VoterSession) AttachTokenToGauge(tokenId *big.Int, account common.Address) (*types.Transaction, error) {
	return _Voter.Contract.AttachTokenToGauge(&_Voter.TransactOpts, tokenId, account)
}

// AttachTokenToGauge is a paid mutator transaction binding the contract method 0x698473e3.
//
// Solidity: function attachTokenToGauge(uint256 tokenId, address account) returns()
func (_Voter *VoterTransactorSession) AttachTokenToGauge(tokenId *big.Int, account common.Address) (*types.Transaction, error) {
	return _Voter.Contract.AttachTokenToGauge(&_Voter.TransactOpts, tokenId, account)
}

// ClaimBribes is a paid mutator transaction binding the contract method 0x7715ee75.
//
// Solidity: function claimBribes(address[] _bribes, address[][] _tokens, uint256 _tokenId) returns()
func (_Voter *VoterTransactor) ClaimBribes(opts *bind.TransactOpts, _bribes []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "claimBribes", _bribes, _tokens, _tokenId)
}

// ClaimBribes is a paid mutator transaction binding the contract method 0x7715ee75.
//
// Solidity: function claimBribes(address[] _bribes, address[][] _tokens, uint256 _tokenId) returns()
func (_Voter *VoterSession) ClaimBribes(_bribes []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.ClaimBribes(&_Voter.TransactOpts, _bribes, _tokens, _tokenId)
}

// ClaimBribes is a paid mutator transaction binding the contract method 0x7715ee75.
//
// Solidity: function claimBribes(address[] _bribes, address[][] _tokens, uint256 _tokenId) returns()
func (_Voter *VoterTransactorSession) ClaimBribes(_bribes []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.ClaimBribes(&_Voter.TransactOpts, _bribes, _tokens, _tokenId)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x666256aa.
//
// Solidity: function claimFees(address[] _fees, address[][] _tokens, uint256 _tokenId) returns()
func (_Voter *VoterTransactor) ClaimFees(opts *bind.TransactOpts, _fees []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "claimFees", _fees, _tokens, _tokenId)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x666256aa.
//
// Solidity: function claimFees(address[] _fees, address[][] _tokens, uint256 _tokenId) returns()
func (_Voter *VoterSession) ClaimFees(_fees []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.ClaimFees(&_Voter.TransactOpts, _fees, _tokens, _tokenId)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x666256aa.
//
// Solidity: function claimFees(address[] _fees, address[][] _tokens, uint256 _tokenId) returns()
func (_Voter *VoterTransactorSession) ClaimFees(_fees []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.ClaimFees(&_Voter.TransactOpts, _fees, _tokens, _tokenId)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x20b1cb6f.
//
// Solidity: function claimRewards(address[] _gauges, address[][] _tokens) returns()
func (_Voter *VoterTransactor) ClaimRewards(opts *bind.TransactOpts, _gauges []common.Address, _tokens [][]common.Address) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "claimRewards", _gauges, _tokens)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x20b1cb6f.
//
// Solidity: function claimRewards(address[] _gauges, address[][] _tokens) returns()
func (_Voter *VoterSession) ClaimRewards(_gauges []common.Address, _tokens [][]common.Address) (*types.Transaction, error) {
	return _Voter.Contract.ClaimRewards(&_Voter.TransactOpts, _gauges, _tokens)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x20b1cb6f.
//
// Solidity: function claimRewards(address[] _gauges, address[][] _tokens) returns()
func (_Voter *VoterTransactorSession) ClaimRewards(_gauges []common.Address, _tokens [][]common.Address) (*types.Transaction, error) {
	return _Voter.Contract.ClaimRewards(&_Voter.TransactOpts, _gauges, _tokens)
}

// CreateGauge is a paid mutator transaction binding the contract method 0xa5f4301e.
//
// Solidity: function createGauge(address _pool) returns(address)
func (_Voter *VoterTransactor) CreateGauge(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "createGauge", _pool)
}

// CreateGauge is a paid mutator transaction binding the contract method 0xa5f4301e.
//
// Solidity: function createGauge(address _pool) returns(address)
func (_Voter *VoterSession) CreateGauge(_pool common.Address) (*types.Transaction, error) {
	return _Voter.Contract.CreateGauge(&_Voter.TransactOpts, _pool)
}

// CreateGauge is a paid mutator transaction binding the contract method 0xa5f4301e.
//
// Solidity: function createGauge(address _pool) returns(address)
func (_Voter *VoterTransactorSession) CreateGauge(_pool common.Address) (*types.Transaction, error) {
	return _Voter.Contract.CreateGauge(&_Voter.TransactOpts, _pool)
}

// DetachTokenFromGauge is a paid mutator transaction binding the contract method 0x411b1f77.
//
// Solidity: function detachTokenFromGauge(uint256 tokenId, address account) returns()
func (_Voter *VoterTransactor) DetachTokenFromGauge(opts *bind.TransactOpts, tokenId *big.Int, account common.Address) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "detachTokenFromGauge", tokenId, account)
}

// DetachTokenFromGauge is a paid mutator transaction binding the contract method 0x411b1f77.
//
// Solidity: function detachTokenFromGauge(uint256 tokenId, address account) returns()
func (_Voter *VoterSession) DetachTokenFromGauge(tokenId *big.Int, account common.Address) (*types.Transaction, error) {
	return _Voter.Contract.DetachTokenFromGauge(&_Voter.TransactOpts, tokenId, account)
}

// DetachTokenFromGauge is a paid mutator transaction binding the contract method 0x411b1f77.
//
// Solidity: function detachTokenFromGauge(uint256 tokenId, address account) returns()
func (_Voter *VoterTransactorSession) DetachTokenFromGauge(tokenId *big.Int, account common.Address) (*types.Transaction, error) {
	return _Voter.Contract.DetachTokenFromGauge(&_Voter.TransactOpts, tokenId, account)
}

// Distribute is a paid mutator transaction binding the contract method 0x6138889b.
//
// Solidity: function distribute(address[] _gauges) returns()
func (_Voter *VoterTransactor) Distribute(opts *bind.TransactOpts, _gauges []common.Address) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "distribute", _gauges)
}

// Distribute is a paid mutator transaction binding the contract method 0x6138889b.
//
// Solidity: function distribute(address[] _gauges) returns()
func (_Voter *VoterSession) Distribute(_gauges []common.Address) (*types.Transaction, error) {
	return _Voter.Contract.Distribute(&_Voter.TransactOpts, _gauges)
}

// Distribute is a paid mutator transaction binding the contract method 0x6138889b.
//
// Solidity: function distribute(address[] _gauges) returns()
func (_Voter *VoterTransactorSession) Distribute(_gauges []common.Address) (*types.Transaction, error) {
	return _Voter.Contract.Distribute(&_Voter.TransactOpts, _gauges)
}

// Distribute0 is a paid mutator transaction binding the contract method 0x63453ae1.
//
// Solidity: function distribute(address _gauge) returns()
func (_Voter *VoterTransactor) Distribute0(opts *bind.TransactOpts, _gauge common.Address) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "distribute0", _gauge)
}

// Distribute0 is a paid mutator transaction binding the contract method 0x63453ae1.
//
// Solidity: function distribute(address _gauge) returns()
func (_Voter *VoterSession) Distribute0(_gauge common.Address) (*types.Transaction, error) {
	return _Voter.Contract.Distribute0(&_Voter.TransactOpts, _gauge)
}

// Distribute0 is a paid mutator transaction binding the contract method 0x63453ae1.
//
// Solidity: function distribute(address _gauge) returns()
func (_Voter *VoterTransactorSession) Distribute0(_gauge common.Address) (*types.Transaction, error) {
	return _Voter.Contract.Distribute0(&_Voter.TransactOpts, _gauge)
}

// Distribute1 is a paid mutator transaction binding the contract method 0x7625391a.
//
// Solidity: function distribute(uint256 start, uint256 finish) returns()
func (_Voter *VoterTransactor) Distribute1(opts *bind.TransactOpts, start *big.Int, finish *big.Int) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "distribute1", start, finish)
}

// Distribute1 is a paid mutator transaction binding the contract method 0x7625391a.
//
// Solidity: function distribute(uint256 start, uint256 finish) returns()
func (_Voter *VoterSession) Distribute1(start *big.Int, finish *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.Distribute1(&_Voter.TransactOpts, start, finish)
}

// Distribute1 is a paid mutator transaction binding the contract method 0x7625391a.
//
// Solidity: function distribute(uint256 start, uint256 finish) returns()
func (_Voter *VoterTransactorSession) Distribute1(start *big.Int, finish *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.Distribute1(&_Voter.TransactOpts, start, finish)
}

// Distribute2 is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_Voter *VoterTransactor) Distribute2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "distribute2")
}

// Distribute2 is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_Voter *VoterSession) Distribute2() (*types.Transaction, error) {
	return _Voter.Contract.Distribute2(&_Voter.TransactOpts)
}

// Distribute2 is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_Voter *VoterTransactorSession) Distribute2() (*types.Transaction, error) {
	return _Voter.Contract.Distribute2(&_Voter.TransactOpts)
}

// DistributeFees is a paid mutator transaction binding the contract method 0xc527ee1f.
//
// Solidity: function distributeFees(address[] _gauges) returns()
func (_Voter *VoterTransactor) DistributeFees(opts *bind.TransactOpts, _gauges []common.Address) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "distributeFees", _gauges)
}

// DistributeFees is a paid mutator transaction binding the contract method 0xc527ee1f.
//
// Solidity: function distributeFees(address[] _gauges) returns()
func (_Voter *VoterSession) DistributeFees(_gauges []common.Address) (*types.Transaction, error) {
	return _Voter.Contract.DistributeFees(&_Voter.TransactOpts, _gauges)
}

// DistributeFees is a paid mutator transaction binding the contract method 0xc527ee1f.
//
// Solidity: function distributeFees(address[] _gauges) returns()
func (_Voter *VoterTransactorSession) DistributeFees(_gauges []common.Address) (*types.Transaction, error) {
	return _Voter.Contract.DistributeFees(&_Voter.TransactOpts, _gauges)
}

// Distro is a paid mutator transaction binding the contract method 0x47b3c6ba.
//
// Solidity: function distro() returns()
func (_Voter *VoterTransactor) Distro(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "distro")
}

// Distro is a paid mutator transaction binding the contract method 0x47b3c6ba.
//
// Solidity: function distro() returns()
func (_Voter *VoterSession) Distro() (*types.Transaction, error) {
	return _Voter.Contract.Distro(&_Voter.TransactOpts)
}

// Distro is a paid mutator transaction binding the contract method 0x47b3c6ba.
//
// Solidity: function distro() returns()
func (_Voter *VoterTransactorSession) Distro() (*types.Transaction, error) {
	return _Voter.Contract.Distro(&_Voter.TransactOpts)
}

// EmitDeposit is a paid mutator transaction binding the contract method 0xa61c713a.
//
// Solidity: function emitDeposit(uint256 tokenId, address account, uint256 amount) returns()
func (_Voter *VoterTransactor) EmitDeposit(opts *bind.TransactOpts, tokenId *big.Int, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "emitDeposit", tokenId, account, amount)
}

// EmitDeposit is a paid mutator transaction binding the contract method 0xa61c713a.
//
// Solidity: function emitDeposit(uint256 tokenId, address account, uint256 amount) returns()
func (_Voter *VoterSession) EmitDeposit(tokenId *big.Int, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.EmitDeposit(&_Voter.TransactOpts, tokenId, account, amount)
}

// EmitDeposit is a paid mutator transaction binding the contract method 0xa61c713a.
//
// Solidity: function emitDeposit(uint256 tokenId, address account, uint256 amount) returns()
func (_Voter *VoterTransactorSession) EmitDeposit(tokenId *big.Int, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.EmitDeposit(&_Voter.TransactOpts, tokenId, account, amount)
}

// EmitWithdraw is a paid mutator transaction binding the contract method 0xea94ee44.
//
// Solidity: function emitWithdraw(uint256 tokenId, address account, uint256 amount) returns()
func (_Voter *VoterTransactor) EmitWithdraw(opts *bind.TransactOpts, tokenId *big.Int, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "emitWithdraw", tokenId, account, amount)
}

// EmitWithdraw is a paid mutator transaction binding the contract method 0xea94ee44.
//
// Solidity: function emitWithdraw(uint256 tokenId, address account, uint256 amount) returns()
func (_Voter *VoterSession) EmitWithdraw(tokenId *big.Int, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.EmitWithdraw(&_Voter.TransactOpts, tokenId, account, amount)
}

// EmitWithdraw is a paid mutator transaction binding the contract method 0xea94ee44.
//
// Solidity: function emitWithdraw(uint256 tokenId, address account, uint256 amount) returns()
func (_Voter *VoterTransactorSession) EmitWithdraw(tokenId *big.Int, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.EmitWithdraw(&_Voter.TransactOpts, tokenId, account, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _tokens, address _minter) returns()
func (_Voter *VoterTransactor) Initialize(opts *bind.TransactOpts, _tokens []common.Address, _minter common.Address) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "initialize", _tokens, _minter)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _tokens, address _minter) returns()
func (_Voter *VoterSession) Initialize(_tokens []common.Address, _minter common.Address) (*types.Transaction, error) {
	return _Voter.Contract.Initialize(&_Voter.TransactOpts, _tokens, _minter)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _tokens, address _minter) returns()
func (_Voter *VoterTransactorSession) Initialize(_tokens []common.Address, _minter common.Address) (*types.Transaction, error) {
	return _Voter.Contract.Initialize(&_Voter.TransactOpts, _tokens, _minter)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 amount) returns()
func (_Voter *VoterTransactor) NotifyRewardAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "notifyRewardAmount", amount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 amount) returns()
func (_Voter *VoterSession) NotifyRewardAmount(amount *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.NotifyRewardAmount(&_Voter.TransactOpts, amount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 amount) returns()
func (_Voter *VoterTransactorSession) NotifyRewardAmount(amount *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.NotifyRewardAmount(&_Voter.TransactOpts, amount)
}

// Poke is a paid mutator transaction binding the contract method 0x32145f90.
//
// Solidity: function poke(uint256 _tokenId) returns()
func (_Voter *VoterTransactor) Poke(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "poke", _tokenId)
}

// Poke is a paid mutator transaction binding the contract method 0x32145f90.
//
// Solidity: function poke(uint256 _tokenId) returns()
func (_Voter *VoterSession) Poke(_tokenId *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.Poke(&_Voter.TransactOpts, _tokenId)
}

// Poke is a paid mutator transaction binding the contract method 0x32145f90.
//
// Solidity: function poke(uint256 _tokenId) returns()
func (_Voter *VoterTransactorSession) Poke(_tokenId *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.Poke(&_Voter.TransactOpts, _tokenId)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 _tokenId) returns()
func (_Voter *VoterTransactor) Reset(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "reset", _tokenId)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 _tokenId) returns()
func (_Voter *VoterSession) Reset(_tokenId *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.Reset(&_Voter.TransactOpts, _tokenId)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 _tokenId) returns()
func (_Voter *VoterTransactorSession) Reset(_tokenId *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.Reset(&_Voter.TransactOpts, _tokenId)
}

// UpdateAll is a paid mutator transaction binding the contract method 0x53d78693.
//
// Solidity: function updateAll() returns()
func (_Voter *VoterTransactor) UpdateAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "updateAll")
}

// UpdateAll is a paid mutator transaction binding the contract method 0x53d78693.
//
// Solidity: function updateAll() returns()
func (_Voter *VoterSession) UpdateAll() (*types.Transaction, error) {
	return _Voter.Contract.UpdateAll(&_Voter.TransactOpts)
}

// UpdateAll is a paid mutator transaction binding the contract method 0x53d78693.
//
// Solidity: function updateAll() returns()
func (_Voter *VoterTransactorSession) UpdateAll() (*types.Transaction, error) {
	return _Voter.Contract.UpdateAll(&_Voter.TransactOpts)
}

// UpdateFor is a paid mutator transaction binding the contract method 0xd560b0d7.
//
// Solidity: function updateFor(address[] _gauges) returns()
func (_Voter *VoterTransactor) UpdateFor(opts *bind.TransactOpts, _gauges []common.Address) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "updateFor", _gauges)
}

// UpdateFor is a paid mutator transaction binding the contract method 0xd560b0d7.
//
// Solidity: function updateFor(address[] _gauges) returns()
func (_Voter *VoterSession) UpdateFor(_gauges []common.Address) (*types.Transaction, error) {
	return _Voter.Contract.UpdateFor(&_Voter.TransactOpts, _gauges)
}

// UpdateFor is a paid mutator transaction binding the contract method 0xd560b0d7.
//
// Solidity: function updateFor(address[] _gauges) returns()
func (_Voter *VoterTransactorSession) UpdateFor(_gauges []common.Address) (*types.Transaction, error) {
	return _Voter.Contract.UpdateFor(&_Voter.TransactOpts, _gauges)
}

// UpdateForRange is a paid mutator transaction binding the contract method 0x9b6a9d72.
//
// Solidity: function updateForRange(uint256 start, uint256 end) returns()
func (_Voter *VoterTransactor) UpdateForRange(opts *bind.TransactOpts, start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "updateForRange", start, end)
}

// UpdateForRange is a paid mutator transaction binding the contract method 0x9b6a9d72.
//
// Solidity: function updateForRange(uint256 start, uint256 end) returns()
func (_Voter *VoterSession) UpdateForRange(start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.UpdateForRange(&_Voter.TransactOpts, start, end)
}

// UpdateForRange is a paid mutator transaction binding the contract method 0x9b6a9d72.
//
// Solidity: function updateForRange(uint256 start, uint256 end) returns()
func (_Voter *VoterTransactorSession) UpdateForRange(start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.UpdateForRange(&_Voter.TransactOpts, start, end)
}

// UpdateGauge is a paid mutator transaction binding the contract method 0x6ecbe38a.
//
// Solidity: function updateGauge(address _gauge) returns()
func (_Voter *VoterTransactor) UpdateGauge(opts *bind.TransactOpts, _gauge common.Address) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "updateGauge", _gauge)
}

// UpdateGauge is a paid mutator transaction binding the contract method 0x6ecbe38a.
//
// Solidity: function updateGauge(address _gauge) returns()
func (_Voter *VoterSession) UpdateGauge(_gauge common.Address) (*types.Transaction, error) {
	return _Voter.Contract.UpdateGauge(&_Voter.TransactOpts, _gauge)
}

// UpdateGauge is a paid mutator transaction binding the contract method 0x6ecbe38a.
//
// Solidity: function updateGauge(address _gauge) returns()
func (_Voter *VoterTransactorSession) UpdateGauge(_gauge common.Address) (*types.Transaction, error) {
	return _Voter.Contract.UpdateGauge(&_Voter.TransactOpts, _gauge)
}

// Vote is a paid mutator transaction binding the contract method 0xfecdad60.
//
// Solidity: function vote(uint256 tokenId, address[] _poolVote, int256[] _weights) returns()
func (_Voter *VoterTransactor) Vote(opts *bind.TransactOpts, tokenId *big.Int, _poolVote []common.Address, _weights []*big.Int) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "vote", tokenId, _poolVote, _weights)
}

// Vote is a paid mutator transaction binding the contract method 0xfecdad60.
//
// Solidity: function vote(uint256 tokenId, address[] _poolVote, int256[] _weights) returns()
func (_Voter *VoterSession) Vote(tokenId *big.Int, _poolVote []common.Address, _weights []*big.Int) (*types.Transaction, error) {
	return _Voter.Contract.Vote(&_Voter.TransactOpts, tokenId, _poolVote, _weights)
}

// Vote is a paid mutator transaction binding the contract method 0xfecdad60.
//
// Solidity: function vote(uint256 tokenId, address[] _poolVote, int256[] _weights) returns()
func (_Voter *VoterTransactorSession) Vote(tokenId *big.Int, _poolVote []common.Address, _weights []*big.Int) (*types.Transaction, error) {
	return _Voter.Contract.Vote(&_Voter.TransactOpts, tokenId, _poolVote, _weights)
}

// Whitelist is a paid mutator transaction binding the contract method 0x98fc55d8.
//
// Solidity: function whitelist(address _token, uint256 _tokenId) returns()
func (_Voter *VoterTransactor) Whitelist(opts *bind.TransactOpts, _token common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "whitelist", _token, _tokenId)
}

// Whitelist is a paid mutator transaction binding the contract method 0x98fc55d8.
//
// Solidity: function whitelist(address _token, uint256 _tokenId) returns()
func (_Voter *VoterSession) Whitelist(_token common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.Whitelist(&_Voter.TransactOpts, _token, _tokenId)
}

// Whitelist is a paid mutator transaction binding the contract method 0x98fc55d8.
//
// Solidity: function whitelist(address _token, uint256 _tokenId) returns()
func (_Voter *VoterTransactorSession) Whitelist(_token common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.Whitelist(&_Voter.TransactOpts, _token, _tokenId)
}

// VoterAbstainedIterator is returned from FilterAbstained and is used to iterate over the raw logs and unpacked data for Abstained events raised by the Voter contract.
type VoterAbstainedIterator struct {
	Event *VoterAbstained // Event containing the contract specifics and raw log

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
func (it *VoterAbstainedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterAbstained)
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
		it.Event = new(VoterAbstained)
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
func (it *VoterAbstainedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterAbstainedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterAbstained represents a Abstained event raised by the Voter contract.
type VoterAbstained struct {
	TokenId *big.Int
	Weight  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAbstained is a free log retrieval operation binding the contract event 0x6b3894ce60b9bbe9d93f1a4e6fc25b6b93cd8222e73ab6348d79c596f5b51de9.
//
// Solidity: event Abstained(uint256 tokenId, int256 weight)
func (_Voter *VoterFilterer) FilterAbstained(opts *bind.FilterOpts) (*VoterAbstainedIterator, error) {

	logs, sub, err := _Voter.contract.FilterLogs(opts, "Abstained")
	if err != nil {
		return nil, err
	}
	return &VoterAbstainedIterator{contract: _Voter.contract, event: "Abstained", logs: logs, sub: sub}, nil
}

// WatchAbstained is a free log subscription operation binding the contract event 0x6b3894ce60b9bbe9d93f1a4e6fc25b6b93cd8222e73ab6348d79c596f5b51de9.
//
// Solidity: event Abstained(uint256 tokenId, int256 weight)
func (_Voter *VoterFilterer) WatchAbstained(opts *bind.WatchOpts, sink chan<- *VoterAbstained) (event.Subscription, error) {

	logs, sub, err := _Voter.contract.WatchLogs(opts, "Abstained")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterAbstained)
				if err := _Voter.contract.UnpackLog(event, "Abstained", log); err != nil {
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

// ParseAbstained is a log parse operation binding the contract event 0x6b3894ce60b9bbe9d93f1a4e6fc25b6b93cd8222e73ab6348d79c596f5b51de9.
//
// Solidity: event Abstained(uint256 tokenId, int256 weight)
func (_Voter *VoterFilterer) ParseAbstained(log types.Log) (*VoterAbstained, error) {
	event := new(VoterAbstained)
	if err := _Voter.contract.UnpackLog(event, "Abstained", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterAttachIterator is returned from FilterAttach and is used to iterate over the raw logs and unpacked data for Attach events raised by the Voter contract.
type VoterAttachIterator struct {
	Event *VoterAttach // Event containing the contract specifics and raw log

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
func (it *VoterAttachIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterAttach)
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
		it.Event = new(VoterAttach)
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
func (it *VoterAttachIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterAttachIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterAttach represents a Attach event raised by the Voter contract.
type VoterAttach struct {
	Owner   common.Address
	Gauge   common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAttach is a free log retrieval operation binding the contract event 0x60940192810a6fb3bce3fd3e2e3a13fd6ccc7605e963fb87ee971aba829989bd.
//
// Solidity: event Attach(address indexed owner, address indexed gauge, uint256 tokenId)
func (_Voter *VoterFilterer) FilterAttach(opts *bind.FilterOpts, owner []common.Address, gauge []common.Address) (*VoterAttachIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _Voter.contract.FilterLogs(opts, "Attach", ownerRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return &VoterAttachIterator{contract: _Voter.contract, event: "Attach", logs: logs, sub: sub}, nil
}

// WatchAttach is a free log subscription operation binding the contract event 0x60940192810a6fb3bce3fd3e2e3a13fd6ccc7605e963fb87ee971aba829989bd.
//
// Solidity: event Attach(address indexed owner, address indexed gauge, uint256 tokenId)
func (_Voter *VoterFilterer) WatchAttach(opts *bind.WatchOpts, sink chan<- *VoterAttach, owner []common.Address, gauge []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _Voter.contract.WatchLogs(opts, "Attach", ownerRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterAttach)
				if err := _Voter.contract.UnpackLog(event, "Attach", log); err != nil {
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

// ParseAttach is a log parse operation binding the contract event 0x60940192810a6fb3bce3fd3e2e3a13fd6ccc7605e963fb87ee971aba829989bd.
//
// Solidity: event Attach(address indexed owner, address indexed gauge, uint256 tokenId)
func (_Voter *VoterFilterer) ParseAttach(log types.Log) (*VoterAttach, error) {
	event := new(VoterAttach)
	if err := _Voter.contract.UnpackLog(event, "Attach", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Voter contract.
type VoterDepositIterator struct {
	Event *VoterDeposit // Event containing the contract specifics and raw log

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
func (it *VoterDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterDeposit)
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
		it.Event = new(VoterDeposit)
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
func (it *VoterDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterDeposit represents a Deposit event raised by the Voter contract.
type VoterDeposit struct {
	Lp      common.Address
	Gauge   common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed lp, address indexed gauge, uint256 tokenId, uint256 amount)
func (_Voter *VoterFilterer) FilterDeposit(opts *bind.FilterOpts, lp []common.Address, gauge []common.Address) (*VoterDepositIterator, error) {

	var lpRule []interface{}
	for _, lpItem := range lp {
		lpRule = append(lpRule, lpItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _Voter.contract.FilterLogs(opts, "Deposit", lpRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return &VoterDepositIterator{contract: _Voter.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed lp, address indexed gauge, uint256 tokenId, uint256 amount)
func (_Voter *VoterFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *VoterDeposit, lp []common.Address, gauge []common.Address) (event.Subscription, error) {

	var lpRule []interface{}
	for _, lpItem := range lp {
		lpRule = append(lpRule, lpItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _Voter.contract.WatchLogs(opts, "Deposit", lpRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterDeposit)
				if err := _Voter.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed lp, address indexed gauge, uint256 tokenId, uint256 amount)
func (_Voter *VoterFilterer) ParseDeposit(log types.Log) (*VoterDeposit, error) {
	event := new(VoterDeposit)
	if err := _Voter.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterDetachIterator is returned from FilterDetach and is used to iterate over the raw logs and unpacked data for Detach events raised by the Voter contract.
type VoterDetachIterator struct {
	Event *VoterDetach // Event containing the contract specifics and raw log

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
func (it *VoterDetachIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterDetach)
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
		it.Event = new(VoterDetach)
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
func (it *VoterDetachIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterDetachIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterDetach represents a Detach event raised by the Voter contract.
type VoterDetach struct {
	Owner   common.Address
	Gauge   common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDetach is a free log retrieval operation binding the contract event 0xae268d9aab12f3605f58efd74fd3801fa812b03fdb44317eb70f46dff0e19e22.
//
// Solidity: event Detach(address indexed owner, address indexed gauge, uint256 tokenId)
func (_Voter *VoterFilterer) FilterDetach(opts *bind.FilterOpts, owner []common.Address, gauge []common.Address) (*VoterDetachIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _Voter.contract.FilterLogs(opts, "Detach", ownerRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return &VoterDetachIterator{contract: _Voter.contract, event: "Detach", logs: logs, sub: sub}, nil
}

// WatchDetach is a free log subscription operation binding the contract event 0xae268d9aab12f3605f58efd74fd3801fa812b03fdb44317eb70f46dff0e19e22.
//
// Solidity: event Detach(address indexed owner, address indexed gauge, uint256 tokenId)
func (_Voter *VoterFilterer) WatchDetach(opts *bind.WatchOpts, sink chan<- *VoterDetach, owner []common.Address, gauge []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _Voter.contract.WatchLogs(opts, "Detach", ownerRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterDetach)
				if err := _Voter.contract.UnpackLog(event, "Detach", log); err != nil {
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

// ParseDetach is a log parse operation binding the contract event 0xae268d9aab12f3605f58efd74fd3801fa812b03fdb44317eb70f46dff0e19e22.
//
// Solidity: event Detach(address indexed owner, address indexed gauge, uint256 tokenId)
func (_Voter *VoterFilterer) ParseDetach(log types.Log) (*VoterDetach, error) {
	event := new(VoterDetach)
	if err := _Voter.contract.UnpackLog(event, "Detach", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterDistributeRewardIterator is returned from FilterDistributeReward and is used to iterate over the raw logs and unpacked data for DistributeReward events raised by the Voter contract.
type VoterDistributeRewardIterator struct {
	Event *VoterDistributeReward // Event containing the contract specifics and raw log

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
func (it *VoterDistributeRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterDistributeReward)
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
		it.Event = new(VoterDistributeReward)
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
func (it *VoterDistributeRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterDistributeRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterDistributeReward represents a DistributeReward event raised by the Voter contract.
type VoterDistributeReward struct {
	Sender common.Address
	Gauge  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributeReward is a free log retrieval operation binding the contract event 0x4fa9693cae526341d334e2862ca2413b2e503f1266255f9e0869fb36e6d89b17.
//
// Solidity: event DistributeReward(address indexed sender, address indexed gauge, uint256 amount)
func (_Voter *VoterFilterer) FilterDistributeReward(opts *bind.FilterOpts, sender []common.Address, gauge []common.Address) (*VoterDistributeRewardIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _Voter.contract.FilterLogs(opts, "DistributeReward", senderRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return &VoterDistributeRewardIterator{contract: _Voter.contract, event: "DistributeReward", logs: logs, sub: sub}, nil
}

// WatchDistributeReward is a free log subscription operation binding the contract event 0x4fa9693cae526341d334e2862ca2413b2e503f1266255f9e0869fb36e6d89b17.
//
// Solidity: event DistributeReward(address indexed sender, address indexed gauge, uint256 amount)
func (_Voter *VoterFilterer) WatchDistributeReward(opts *bind.WatchOpts, sink chan<- *VoterDistributeReward, sender []common.Address, gauge []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _Voter.contract.WatchLogs(opts, "DistributeReward", senderRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterDistributeReward)
				if err := _Voter.contract.UnpackLog(event, "DistributeReward", log); err != nil {
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

// ParseDistributeReward is a log parse operation binding the contract event 0x4fa9693cae526341d334e2862ca2413b2e503f1266255f9e0869fb36e6d89b17.
//
// Solidity: event DistributeReward(address indexed sender, address indexed gauge, uint256 amount)
func (_Voter *VoterFilterer) ParseDistributeReward(log types.Log) (*VoterDistributeReward, error) {
	event := new(VoterDistributeReward)
	if err := _Voter.contract.UnpackLog(event, "DistributeReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterGaugeCreatedIterator is returned from FilterGaugeCreated and is used to iterate over the raw logs and unpacked data for GaugeCreated events raised by the Voter contract.
type VoterGaugeCreatedIterator struct {
	Event *VoterGaugeCreated // Event containing the contract specifics and raw log

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
func (it *VoterGaugeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterGaugeCreated)
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
		it.Event = new(VoterGaugeCreated)
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
func (it *VoterGaugeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterGaugeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterGaugeCreated represents a GaugeCreated event raised by the Voter contract.
type VoterGaugeCreated struct {
	Gauge   common.Address
	Creator common.Address
	Bribe   common.Address
	Pool    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGaugeCreated is a free log retrieval operation binding the contract event 0x48d3c521fd0d5541640f58c6d6381eed7cb2e8c9df421ae165a4f4c2d221ee0d.
//
// Solidity: event GaugeCreated(address indexed gauge, address creator, address indexed bribe, address indexed pool)
func (_Voter *VoterFilterer) FilterGaugeCreated(opts *bind.FilterOpts, gauge []common.Address, bribe []common.Address, pool []common.Address) (*VoterGaugeCreatedIterator, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	var bribeRule []interface{}
	for _, bribeItem := range bribe {
		bribeRule = append(bribeRule, bribeItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Voter.contract.FilterLogs(opts, "GaugeCreated", gaugeRule, bribeRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &VoterGaugeCreatedIterator{contract: _Voter.contract, event: "GaugeCreated", logs: logs, sub: sub}, nil
}

// WatchGaugeCreated is a free log subscription operation binding the contract event 0x48d3c521fd0d5541640f58c6d6381eed7cb2e8c9df421ae165a4f4c2d221ee0d.
//
// Solidity: event GaugeCreated(address indexed gauge, address creator, address indexed bribe, address indexed pool)
func (_Voter *VoterFilterer) WatchGaugeCreated(opts *bind.WatchOpts, sink chan<- *VoterGaugeCreated, gauge []common.Address, bribe []common.Address, pool []common.Address) (event.Subscription, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	var bribeRule []interface{}
	for _, bribeItem := range bribe {
		bribeRule = append(bribeRule, bribeItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Voter.contract.WatchLogs(opts, "GaugeCreated", gaugeRule, bribeRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterGaugeCreated)
				if err := _Voter.contract.UnpackLog(event, "GaugeCreated", log); err != nil {
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

// ParseGaugeCreated is a log parse operation binding the contract event 0x48d3c521fd0d5541640f58c6d6381eed7cb2e8c9df421ae165a4f4c2d221ee0d.
//
// Solidity: event GaugeCreated(address indexed gauge, address creator, address indexed bribe, address indexed pool)
func (_Voter *VoterFilterer) ParseGaugeCreated(log types.Log) (*VoterGaugeCreated, error) {
	event := new(VoterGaugeCreated)
	if err := _Voter.contract.UnpackLog(event, "GaugeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the Voter contract.
type VoterInitializeIterator struct {
	Event *VoterInitialize // Event containing the contract specifics and raw log

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
func (it *VoterInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterInitialize)
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
		it.Event = new(VoterInitialize)
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
func (it *VoterInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterInitialize represents a Initialize event raised by the Voter contract.
type VoterInitialize struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x36b1453565f45af7b509b59d5e2eac8f1948ea9e3e193db6663b4101afb6382c.
//
// Solidity: event Initialize(address indexed minter)
func (_Voter *VoterFilterer) FilterInitialize(opts *bind.FilterOpts, minter []common.Address) (*VoterInitializeIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Voter.contract.FilterLogs(opts, "Initialize", minterRule)
	if err != nil {
		return nil, err
	}
	return &VoterInitializeIterator{contract: _Voter.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x36b1453565f45af7b509b59d5e2eac8f1948ea9e3e193db6663b4101afb6382c.
//
// Solidity: event Initialize(address indexed minter)
func (_Voter *VoterFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *VoterInitialize, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Voter.contract.WatchLogs(opts, "Initialize", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterInitialize)
				if err := _Voter.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0x36b1453565f45af7b509b59d5e2eac8f1948ea9e3e193db6663b4101afb6382c.
//
// Solidity: event Initialize(address indexed minter)
func (_Voter *VoterFilterer) ParseInitialize(log types.Log) (*VoterInitialize, error) {
	event := new(VoterInitialize)
	if err := _Voter.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterNotifyRewardIterator is returned from FilterNotifyReward and is used to iterate over the raw logs and unpacked data for NotifyReward events raised by the Voter contract.
type VoterNotifyRewardIterator struct {
	Event *VoterNotifyReward // Event containing the contract specifics and raw log

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
func (it *VoterNotifyRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterNotifyReward)
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
		it.Event = new(VoterNotifyReward)
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
func (it *VoterNotifyRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterNotifyRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterNotifyReward represents a NotifyReward event raised by the Voter contract.
type VoterNotifyReward struct {
	Sender common.Address
	Reward common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotifyReward is a free log retrieval operation binding the contract event 0xf70d5c697de7ea828df48e5c4573cb2194c659f1901f70110c52b066dcf50826.
//
// Solidity: event NotifyReward(address indexed sender, address indexed reward, uint256 amount)
func (_Voter *VoterFilterer) FilterNotifyReward(opts *bind.FilterOpts, sender []common.Address, reward []common.Address) (*VoterNotifyRewardIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _Voter.contract.FilterLogs(opts, "NotifyReward", senderRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &VoterNotifyRewardIterator{contract: _Voter.contract, event: "NotifyReward", logs: logs, sub: sub}, nil
}

// WatchNotifyReward is a free log subscription operation binding the contract event 0xf70d5c697de7ea828df48e5c4573cb2194c659f1901f70110c52b066dcf50826.
//
// Solidity: event NotifyReward(address indexed sender, address indexed reward, uint256 amount)
func (_Voter *VoterFilterer) WatchNotifyReward(opts *bind.WatchOpts, sink chan<- *VoterNotifyReward, sender []common.Address, reward []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _Voter.contract.WatchLogs(opts, "NotifyReward", senderRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterNotifyReward)
				if err := _Voter.contract.UnpackLog(event, "NotifyReward", log); err != nil {
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
// Solidity: event NotifyReward(address indexed sender, address indexed reward, uint256 amount)
func (_Voter *VoterFilterer) ParseNotifyReward(log types.Log) (*VoterNotifyReward, error) {
	event := new(VoterNotifyReward)
	if err := _Voter.contract.UnpackLog(event, "NotifyReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Voter contract.
type VoterVotedIterator struct {
	Event *VoterVoted // Event containing the contract specifics and raw log

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
func (it *VoterVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterVoted)
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
		it.Event = new(VoterVoted)
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
func (it *VoterVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterVoted represents a Voted event raised by the Voter contract.
type VoterVoted struct {
	Voter   common.Address
	TokenId *big.Int
	Weight  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x1263a2295e53acd6ef8f655b8afc11fa0f2cf11925be7aa1757d741ef32a926c.
//
// Solidity: event Voted(address indexed voter, uint256 tokenId, int256 weight)
func (_Voter *VoterFilterer) FilterVoted(opts *bind.FilterOpts, voter []common.Address) (*VoterVotedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Voter.contract.FilterLogs(opts, "Voted", voterRule)
	if err != nil {
		return nil, err
	}
	return &VoterVotedIterator{contract: _Voter.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x1263a2295e53acd6ef8f655b8afc11fa0f2cf11925be7aa1757d741ef32a926c.
//
// Solidity: event Voted(address indexed voter, uint256 tokenId, int256 weight)
func (_Voter *VoterFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *VoterVoted, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Voter.contract.WatchLogs(opts, "Voted", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterVoted)
				if err := _Voter.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x1263a2295e53acd6ef8f655b8afc11fa0f2cf11925be7aa1757d741ef32a926c.
//
// Solidity: event Voted(address indexed voter, uint256 tokenId, int256 weight)
func (_Voter *VoterFilterer) ParseVoted(log types.Log) (*VoterVoted, error) {
	event := new(VoterVoted)
	if err := _Voter.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the Voter contract.
type VoterWhitelistedIterator struct {
	Event *VoterWhitelisted // Event containing the contract specifics and raw log

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
func (it *VoterWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterWhitelisted)
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
		it.Event = new(VoterWhitelisted)
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
func (it *VoterWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterWhitelisted represents a Whitelisted event raised by the Voter contract.
type VoterWhitelisted struct {
	Whitelister common.Address
	Token       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0x6661a7108aecd07864384529117d96c319c1163e3010c01390f6b704726e07de.
//
// Solidity: event Whitelisted(address indexed whitelister, address indexed token)
func (_Voter *VoterFilterer) FilterWhitelisted(opts *bind.FilterOpts, whitelister []common.Address, token []common.Address) (*VoterWhitelistedIterator, error) {

	var whitelisterRule []interface{}
	for _, whitelisterItem := range whitelister {
		whitelisterRule = append(whitelisterRule, whitelisterItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Voter.contract.FilterLogs(opts, "Whitelisted", whitelisterRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &VoterWhitelistedIterator{contract: _Voter.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0x6661a7108aecd07864384529117d96c319c1163e3010c01390f6b704726e07de.
//
// Solidity: event Whitelisted(address indexed whitelister, address indexed token)
func (_Voter *VoterFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *VoterWhitelisted, whitelister []common.Address, token []common.Address) (event.Subscription, error) {

	var whitelisterRule []interface{}
	for _, whitelisterItem := range whitelister {
		whitelisterRule = append(whitelisterRule, whitelisterItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Voter.contract.WatchLogs(opts, "Whitelisted", whitelisterRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterWhitelisted)
				if err := _Voter.contract.UnpackLog(event, "Whitelisted", log); err != nil {
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

// ParseWhitelisted is a log parse operation binding the contract event 0x6661a7108aecd07864384529117d96c319c1163e3010c01390f6b704726e07de.
//
// Solidity: event Whitelisted(address indexed whitelister, address indexed token)
func (_Voter *VoterFilterer) ParseWhitelisted(log types.Log) (*VoterWhitelisted, error) {
	event := new(VoterWhitelisted)
	if err := _Voter.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Voter contract.
type VoterWithdrawIterator struct {
	Event *VoterWithdraw // Event containing the contract specifics and raw log

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
func (it *VoterWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterWithdraw)
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
		it.Event = new(VoterWithdraw)
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
func (it *VoterWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterWithdraw represents a Withdraw event raised by the Voter contract.
type VoterWithdraw struct {
	Lp      common.Address
	Gauge   common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: event Withdraw(address indexed lp, address indexed gauge, uint256 tokenId, uint256 amount)
func (_Voter *VoterFilterer) FilterWithdraw(opts *bind.FilterOpts, lp []common.Address, gauge []common.Address) (*VoterWithdrawIterator, error) {

	var lpRule []interface{}
	for _, lpItem := range lp {
		lpRule = append(lpRule, lpItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _Voter.contract.FilterLogs(opts, "Withdraw", lpRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return &VoterWithdrawIterator{contract: _Voter.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: event Withdraw(address indexed lp, address indexed gauge, uint256 tokenId, uint256 amount)
func (_Voter *VoterFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *VoterWithdraw, lp []common.Address, gauge []common.Address) (event.Subscription, error) {

	var lpRule []interface{}
	for _, lpItem := range lp {
		lpRule = append(lpRule, lpItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _Voter.contract.WatchLogs(opts, "Withdraw", lpRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterWithdraw)
				if err := _Voter.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: event Withdraw(address indexed lp, address indexed gauge, uint256 tokenId, uint256 amount)
func (_Voter *VoterFilterer) ParseWithdraw(log types.Log) (*VoterWithdraw, error) {
	event := new(VoterWithdraw)
	if err := _Voter.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
