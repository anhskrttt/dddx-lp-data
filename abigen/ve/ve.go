// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ve

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

// VeMetaData contains all meta data concerning the Ve contract.
var VeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_addr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Abstain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Attach\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"locktime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumve.DepositType\",\"name\":\"deposit_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ts\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Detach\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"SetVoter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"Supply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Voting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ts\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"abstain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_approved\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"attach\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"attachments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"balanceOfAtNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"balanceOfNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_t\",\"type\":\"uint256\"}],\"name\":\"balanceOfNFTAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"block_number\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lock_duration\",\"type\":\"uint256\"}],\"name\":\"create_lock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lock_duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"create_lock_for\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"deposit_for\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"detach\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"get_last_user_slope\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"increase_amount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lock_duration\",\"type\":\"uint256\"}],\"name\":\"increase_unlock_time\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"isApprovedOrOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"amount\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"locked__end\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"merge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownership_change\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"point_history\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"bias\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"slope\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"ts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blk\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"setVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slope_changes\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenIndex\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAtT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"user_point_epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"user_point_history\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"bias\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"slope\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"ts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blk\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"user_point_history__ts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"voting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VeABI is the input ABI used to generate the binding from.
// Deprecated: Use VeMetaData.ABI instead.
var VeABI = VeMetaData.ABI

// Ve is an auto generated Go binding around an Ethereum contract.
type Ve struct {
	VeCaller     // Read-only binding to the contract
	VeTransactor // Write-only binding to the contract
	VeFilterer   // Log filterer for contract events
}

// VeCaller is an auto generated read-only Go binding around an Ethereum contract.
type VeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VeSession struct {
	Contract     *Ve               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VeCallerSession struct {
	Contract *VeCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VeTransactorSession struct {
	Contract     *VeTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VeRaw is an auto generated low-level Go binding around an Ethereum contract.
type VeRaw struct {
	Contract *Ve // Generic contract binding to access the raw methods on
}

// VeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VeCallerRaw struct {
	Contract *VeCaller // Generic read-only contract binding to access the raw methods on
}

// VeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VeTransactorRaw struct {
	Contract *VeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVe creates a new instance of Ve, bound to a specific deployed contract.
func NewVe(address common.Address, backend bind.ContractBackend) (*Ve, error) {
	contract, err := bindVe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ve{VeCaller: VeCaller{contract: contract}, VeTransactor: VeTransactor{contract: contract}, VeFilterer: VeFilterer{contract: contract}}, nil
}

// NewVeCaller creates a new read-only instance of Ve, bound to a specific deployed contract.
func NewVeCaller(address common.Address, caller bind.ContractCaller) (*VeCaller, error) {
	contract, err := bindVe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VeCaller{contract: contract}, nil
}

// NewVeTransactor creates a new write-only instance of Ve, bound to a specific deployed contract.
func NewVeTransactor(address common.Address, transactor bind.ContractTransactor) (*VeTransactor, error) {
	contract, err := bindVe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VeTransactor{contract: contract}, nil
}

// NewVeFilterer creates a new log filterer instance of Ve, bound to a specific deployed contract.
func NewVeFilterer(address common.Address, filterer bind.ContractFilterer) (*VeFilterer, error) {
	contract, err := bindVe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VeFilterer{contract: contract}, nil
}

// bindVe binds a generic wrapper to an already deployed contract.
func bindVe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ve *VeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ve.Contract.VeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ve *VeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ve.Contract.VeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ve *VeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ve.Contract.VeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ve *VeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ve *VeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ve *VeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ve.Contract.contract.Transact(opts, method, params...)
}

// Attachments is a free data retrieval call binding the contract method 0x0d6a2033.
//
// Solidity: function attachments(uint256 ) view returns(uint256)
func (_Ve *VeCaller) Attachments(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "attachments", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Attachments is a free data retrieval call binding the contract method 0x0d6a2033.
//
// Solidity: function attachments(uint256 ) view returns(uint256)
func (_Ve *VeSession) Attachments(arg0 *big.Int) (*big.Int, error) {
	return _Ve.Contract.Attachments(&_Ve.CallOpts, arg0)
}

// Attachments is a free data retrieval call binding the contract method 0x0d6a2033.
//
// Solidity: function attachments(uint256 ) view returns(uint256)
func (_Ve *VeCallerSession) Attachments(arg0 *big.Int) (*big.Int, error) {
	return _Ve.Contract.Attachments(&_Ve.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Ve *VeCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Ve *VeSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Ve.Contract.BalanceOf(&_Ve.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Ve *VeCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Ve.Contract.BalanceOf(&_Ve.CallOpts, _owner)
}

// BalanceOfAtNFT is a free data retrieval call binding the contract method 0x8c2c9baf.
//
// Solidity: function balanceOfAtNFT(uint256 _tokenId, uint256 _block) view returns(uint256)
func (_Ve *VeCaller) BalanceOfAtNFT(opts *bind.CallOpts, _tokenId *big.Int, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "balanceOfAtNFT", _tokenId, _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAtNFT is a free data retrieval call binding the contract method 0x8c2c9baf.
//
// Solidity: function balanceOfAtNFT(uint256 _tokenId, uint256 _block) view returns(uint256)
func (_Ve *VeSession) BalanceOfAtNFT(_tokenId *big.Int, _block *big.Int) (*big.Int, error) {
	return _Ve.Contract.BalanceOfAtNFT(&_Ve.CallOpts, _tokenId, _block)
}

// BalanceOfAtNFT is a free data retrieval call binding the contract method 0x8c2c9baf.
//
// Solidity: function balanceOfAtNFT(uint256 _tokenId, uint256 _block) view returns(uint256)
func (_Ve *VeCallerSession) BalanceOfAtNFT(_tokenId *big.Int, _block *big.Int) (*big.Int, error) {
	return _Ve.Contract.BalanceOfAtNFT(&_Ve.CallOpts, _tokenId, _block)
}

// BalanceOfNFT is a free data retrieval call binding the contract method 0xe7e242d4.
//
// Solidity: function balanceOfNFT(uint256 _tokenId) view returns(uint256)
func (_Ve *VeCaller) BalanceOfNFT(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "balanceOfNFT", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfNFT is a free data retrieval call binding the contract method 0xe7e242d4.
//
// Solidity: function balanceOfNFT(uint256 _tokenId) view returns(uint256)
func (_Ve *VeSession) BalanceOfNFT(_tokenId *big.Int) (*big.Int, error) {
	return _Ve.Contract.BalanceOfNFT(&_Ve.CallOpts, _tokenId)
}

// BalanceOfNFT is a free data retrieval call binding the contract method 0xe7e242d4.
//
// Solidity: function balanceOfNFT(uint256 _tokenId) view returns(uint256)
func (_Ve *VeCallerSession) BalanceOfNFT(_tokenId *big.Int) (*big.Int, error) {
	return _Ve.Contract.BalanceOfNFT(&_Ve.CallOpts, _tokenId)
}

// BalanceOfNFTAt is a free data retrieval call binding the contract method 0xe0514aba.
//
// Solidity: function balanceOfNFTAt(uint256 _tokenId, uint256 _t) view returns(uint256)
func (_Ve *VeCaller) BalanceOfNFTAt(opts *bind.CallOpts, _tokenId *big.Int, _t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "balanceOfNFTAt", _tokenId, _t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfNFTAt is a free data retrieval call binding the contract method 0xe0514aba.
//
// Solidity: function balanceOfNFTAt(uint256 _tokenId, uint256 _t) view returns(uint256)
func (_Ve *VeSession) BalanceOfNFTAt(_tokenId *big.Int, _t *big.Int) (*big.Int, error) {
	return _Ve.Contract.BalanceOfNFTAt(&_Ve.CallOpts, _tokenId, _t)
}

// BalanceOfNFTAt is a free data retrieval call binding the contract method 0xe0514aba.
//
// Solidity: function balanceOfNFTAt(uint256 _tokenId, uint256 _t) view returns(uint256)
func (_Ve *VeCallerSession) BalanceOfNFTAt(_tokenId *big.Int, _t *big.Int) (*big.Int, error) {
	return _Ve.Contract.BalanceOfNFTAt(&_Ve.CallOpts, _tokenId, _t)
}

// BlockNumber is a free data retrieval call binding the contract method 0x25a58b56.
//
// Solidity: function block_number() view returns(uint256)
func (_Ve *VeCaller) BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "block_number")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockNumber is a free data retrieval call binding the contract method 0x25a58b56.
//
// Solidity: function block_number() view returns(uint256)
func (_Ve *VeSession) BlockNumber() (*big.Int, error) {
	return _Ve.Contract.BlockNumber(&_Ve.CallOpts)
}

// BlockNumber is a free data retrieval call binding the contract method 0x25a58b56.
//
// Solidity: function block_number() view returns(uint256)
func (_Ve *VeCallerSession) BlockNumber() (*big.Int, error) {
	return _Ve.Contract.BlockNumber(&_Ve.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ve *VeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ve *VeSession) Decimals() (uint8, error) {
	return _Ve.Contract.Decimals(&_Ve.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ve *VeCallerSession) Decimals() (uint8, error) {
	return _Ve.Contract.Decimals(&_Ve.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Ve *VeCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Ve *VeSession) Epoch() (*big.Int, error) {
	return _Ve.Contract.Epoch(&_Ve.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Ve *VeCallerSession) Epoch() (*big.Int, error) {
	return _Ve.Contract.Epoch(&_Ve.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address)
func (_Ve *VeCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "getApproved", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address)
func (_Ve *VeSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Ve.Contract.GetApproved(&_Ve.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address)
func (_Ve *VeCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Ve.Contract.GetApproved(&_Ve.CallOpts, _tokenId)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x461f711c.
//
// Solidity: function get_last_user_slope(uint256 _tokenId) view returns(int128)
func (_Ve *VeCaller) GetLastUserSlope(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "get_last_user_slope", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x461f711c.
//
// Solidity: function get_last_user_slope(uint256 _tokenId) view returns(int128)
func (_Ve *VeSession) GetLastUserSlope(_tokenId *big.Int) (*big.Int, error) {
	return _Ve.Contract.GetLastUserSlope(&_Ve.CallOpts, _tokenId)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x461f711c.
//
// Solidity: function get_last_user_slope(uint256 _tokenId) view returns(int128)
func (_Ve *VeCallerSession) GetLastUserSlope(_tokenId *big.Int) (*big.Int, error) {
	return _Ve.Contract.GetLastUserSlope(&_Ve.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_Ve *VeCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_Ve *VeSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Ve.Contract.IsApprovedForAll(&_Ve.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_Ve *VeCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Ve.Contract.IsApprovedForAll(&_Ve.CallOpts, _owner, _operator)
}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address _spender, uint256 _tokenId) view returns(bool)
func (_Ve *VeCaller) IsApprovedOrOwner(opts *bind.CallOpts, _spender common.Address, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "isApprovedOrOwner", _spender, _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address _spender, uint256 _tokenId) view returns(bool)
func (_Ve *VeSession) IsApprovedOrOwner(_spender common.Address, _tokenId *big.Int) (bool, error) {
	return _Ve.Contract.IsApprovedOrOwner(&_Ve.CallOpts, _spender, _tokenId)
}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address _spender, uint256 _tokenId) view returns(bool)
func (_Ve *VeCallerSession) IsApprovedOrOwner(_spender common.Address, _tokenId *big.Int) (bool, error) {
	return _Ve.Contract.IsApprovedOrOwner(&_Ve.CallOpts, _spender, _tokenId)
}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 ) view returns(int128 amount, uint256 end)
func (_Ve *VeCaller) Locked(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "locked", arg0)

	outstruct := new(struct {
		Amount *big.Int
		End    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 ) view returns(int128 amount, uint256 end)
func (_Ve *VeSession) Locked(arg0 *big.Int) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	return _Ve.Contract.Locked(&_Ve.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 ) view returns(int128 amount, uint256 end)
func (_Ve *VeCallerSession) Locked(arg0 *big.Int) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	return _Ve.Contract.Locked(&_Ve.CallOpts, arg0)
}

// LockedEnd is a free data retrieval call binding the contract method 0xf8a05763.
//
// Solidity: function locked__end(uint256 _tokenId) view returns(uint256)
func (_Ve *VeCaller) LockedEnd(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "locked__end", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedEnd is a free data retrieval call binding the contract method 0xf8a05763.
//
// Solidity: function locked__end(uint256 _tokenId) view returns(uint256)
func (_Ve *VeSession) LockedEnd(_tokenId *big.Int) (*big.Int, error) {
	return _Ve.Contract.LockedEnd(&_Ve.CallOpts, _tokenId)
}

// LockedEnd is a free data retrieval call binding the contract method 0xf8a05763.
//
// Solidity: function locked__end(uint256 _tokenId) view returns(uint256)
func (_Ve *VeCallerSession) LockedEnd(_tokenId *big.Int) (*big.Int, error) {
	return _Ve.Contract.LockedEnd(&_Ve.CallOpts, _tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ve *VeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ve *VeSession) Name() (string, error) {
	return _Ve.Contract.Name(&_Ve.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ve *VeCallerSession) Name() (string, error) {
	return _Ve.Contract.Name(&_Ve.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_Ve *VeCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_Ve *VeSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Ve.Contract.OwnerOf(&_Ve.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_Ve *VeCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Ve.Contract.OwnerOf(&_Ve.CallOpts, _tokenId)
}

// OwnershipChange is a free data retrieval call binding the contract method 0x6f548837.
//
// Solidity: function ownership_change(uint256 ) view returns(uint256)
func (_Ve *VeCaller) OwnershipChange(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "ownership_change", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipChange is a free data retrieval call binding the contract method 0x6f548837.
//
// Solidity: function ownership_change(uint256 ) view returns(uint256)
func (_Ve *VeSession) OwnershipChange(arg0 *big.Int) (*big.Int, error) {
	return _Ve.Contract.OwnershipChange(&_Ve.CallOpts, arg0)
}

// OwnershipChange is a free data retrieval call binding the contract method 0x6f548837.
//
// Solidity: function ownership_change(uint256 ) view returns(uint256)
func (_Ve *VeCallerSession) OwnershipChange(arg0 *big.Int) (*big.Int, error) {
	return _Ve.Contract.OwnershipChange(&_Ve.CallOpts, arg0)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 ) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Ve *VeCaller) PointHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "point_history", arg0)

	outstruct := new(struct {
		Bias  *big.Int
		Slope *big.Int
		Ts    *big.Int
		Blk   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bias = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ts = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Blk = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 ) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Ve *VeSession) PointHistory(arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Ve.Contract.PointHistory(&_Ve.CallOpts, arg0)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 ) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Ve *VeCallerSession) PointHistory(arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Ve.Contract.PointHistory(&_Ve.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 ) view returns(int128)
func (_Ve *VeCaller) SlopeChanges(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "slope_changes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 ) view returns(int128)
func (_Ve *VeSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _Ve.Contract.SlopeChanges(&_Ve.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 ) view returns(int128)
func (_Ve *VeCallerSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _Ve.Contract.SlopeChanges(&_Ve.CallOpts, arg0)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Ve *VeCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Ve *VeSession) Supply() (*big.Int, error) {
	return _Ve.Contract.Supply(&_Ve.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Ve *VeCallerSession) Supply() (*big.Int, error) {
	return _Ve.Contract.Supply(&_Ve.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_Ve *VeCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "supportsInterface", _interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_Ve *VeSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _Ve.Contract.SupportsInterface(&_Ve.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_Ve *VeCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _Ve.Contract.SupportsInterface(&_Ve.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ve *VeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ve *VeSession) Symbol() (string, error) {
	return _Ve.Contract.Symbol(&_Ve.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ve *VeCallerSession) Symbol() (string, error) {
	return _Ve.Contract.Symbol(&_Ve.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Ve *VeCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Ve *VeSession) Token() (common.Address, error) {
	return _Ve.Contract.Token(&_Ve.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Ve *VeCallerSession) Token() (common.Address, error) {
	return _Ve.Contract.Token(&_Ve.CallOpts)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _tokenIndex) view returns(uint256)
func (_Ve *VeCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _tokenIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "tokenOfOwnerByIndex", _owner, _tokenIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _tokenIndex) view returns(uint256)
func (_Ve *VeSession) TokenOfOwnerByIndex(_owner common.Address, _tokenIndex *big.Int) (*big.Int, error) {
	return _Ve.Contract.TokenOfOwnerByIndex(&_Ve.CallOpts, _owner, _tokenIndex)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _tokenIndex) view returns(uint256)
func (_Ve *VeCallerSession) TokenOfOwnerByIndex(_owner common.Address, _tokenIndex *big.Int) (*big.Int, error) {
	return _Ve.Contract.TokenOfOwnerByIndex(&_Ve.CallOpts, _owner, _tokenIndex)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Ve *VeCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Ve *VeSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Ve.Contract.TokenURI(&_Ve.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Ve *VeCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Ve.Contract.TokenURI(&_Ve.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ve *VeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ve *VeSession) TotalSupply() (*big.Int, error) {
	return _Ve.Contract.TotalSupply(&_Ve.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ve *VeCallerSession) TotalSupply() (*big.Int, error) {
	return _Ve.Contract.TotalSupply(&_Ve.CallOpts)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Ve *VeCaller) TotalSupplyAt(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "totalSupplyAt", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Ve *VeSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _Ve.Contract.TotalSupplyAt(&_Ve.CallOpts, _block)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Ve *VeCallerSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _Ve.Contract.TotalSupplyAt(&_Ve.CallOpts, _block)
}

// TotalSupplyAtT is a free data retrieval call binding the contract method 0x7116c60c.
//
// Solidity: function totalSupplyAtT(uint256 t) view returns(uint256)
func (_Ve *VeCaller) TotalSupplyAtT(opts *bind.CallOpts, t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "totalSupplyAtT", t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAtT is a free data retrieval call binding the contract method 0x7116c60c.
//
// Solidity: function totalSupplyAtT(uint256 t) view returns(uint256)
func (_Ve *VeSession) TotalSupplyAtT(t *big.Int) (*big.Int, error) {
	return _Ve.Contract.TotalSupplyAtT(&_Ve.CallOpts, t)
}

// TotalSupplyAtT is a free data retrieval call binding the contract method 0x7116c60c.
//
// Solidity: function totalSupplyAtT(uint256 t) view returns(uint256)
func (_Ve *VeCallerSession) TotalSupplyAtT(t *big.Int) (*big.Int, error) {
	return _Ve.Contract.TotalSupplyAtT(&_Ve.CallOpts, t)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0xe441135c.
//
// Solidity: function user_point_epoch(uint256 ) view returns(uint256)
func (_Ve *VeCaller) UserPointEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "user_point_epoch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPointEpoch is a free data retrieval call binding the contract method 0xe441135c.
//
// Solidity: function user_point_epoch(uint256 ) view returns(uint256)
func (_Ve *VeSession) UserPointEpoch(arg0 *big.Int) (*big.Int, error) {
	return _Ve.Contract.UserPointEpoch(&_Ve.CallOpts, arg0)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0xe441135c.
//
// Solidity: function user_point_epoch(uint256 ) view returns(uint256)
func (_Ve *VeCallerSession) UserPointEpoch(arg0 *big.Int) (*big.Int, error) {
	return _Ve.Contract.UserPointEpoch(&_Ve.CallOpts, arg0)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x1376f3da.
//
// Solidity: function user_point_history(uint256 , uint256 ) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Ve *VeCaller) UserPointHistory(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "user_point_history", arg0, arg1)

	outstruct := new(struct {
		Bias  *big.Int
		Slope *big.Int
		Ts    *big.Int
		Blk   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bias = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ts = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Blk = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserPointHistory is a free data retrieval call binding the contract method 0x1376f3da.
//
// Solidity: function user_point_history(uint256 , uint256 ) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Ve *VeSession) UserPointHistory(arg0 *big.Int, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Ve.Contract.UserPointHistory(&_Ve.CallOpts, arg0, arg1)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x1376f3da.
//
// Solidity: function user_point_history(uint256 , uint256 ) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Ve *VeCallerSession) UserPointHistory(arg0 *big.Int, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Ve.Contract.UserPointHistory(&_Ve.CallOpts, arg0, arg1)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0x1c984bc3.
//
// Solidity: function user_point_history__ts(uint256 _tokenId, uint256 _idx) view returns(uint256)
func (_Ve *VeCaller) UserPointHistoryTs(opts *bind.CallOpts, _tokenId *big.Int, _idx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "user_point_history__ts", _tokenId, _idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0x1c984bc3.
//
// Solidity: function user_point_history__ts(uint256 _tokenId, uint256 _idx) view returns(uint256)
func (_Ve *VeSession) UserPointHistoryTs(_tokenId *big.Int, _idx *big.Int) (*big.Int, error) {
	return _Ve.Contract.UserPointHistoryTs(&_Ve.CallOpts, _tokenId, _idx)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0x1c984bc3.
//
// Solidity: function user_point_history__ts(uint256 _tokenId, uint256 _idx) view returns(uint256)
func (_Ve *VeCallerSession) UserPointHistoryTs(_tokenId *big.Int, _idx *big.Int) (*big.Int, error) {
	return _Ve.Contract.UserPointHistoryTs(&_Ve.CallOpts, _tokenId, _idx)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Ve *VeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Ve *VeSession) Version() (string, error) {
	return _Ve.Contract.Version(&_Ve.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Ve *VeCallerSession) Version() (string, error) {
	return _Ve.Contract.Version(&_Ve.CallOpts)
}

// Voted is a free data retrieval call binding the contract method 0x8fbb38ff.
//
// Solidity: function voted(uint256 ) view returns(bool)
func (_Ve *VeCaller) Voted(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "voted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Voted is a free data retrieval call binding the contract method 0x8fbb38ff.
//
// Solidity: function voted(uint256 ) view returns(bool)
func (_Ve *VeSession) Voted(arg0 *big.Int) (bool, error) {
	return _Ve.Contract.Voted(&_Ve.CallOpts, arg0)
}

// Voted is a free data retrieval call binding the contract method 0x8fbb38ff.
//
// Solidity: function voted(uint256 ) view returns(bool)
func (_Ve *VeCallerSession) Voted(arg0 *big.Int) (bool, error) {
	return _Ve.Contract.Voted(&_Ve.CallOpts, arg0)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_Ve *VeCaller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ve.contract.Call(opts, &out, "voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_Ve *VeSession) Voter() (common.Address, error) {
	return _Ve.Contract.Voter(&_Ve.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_Ve *VeCallerSession) Voter() (common.Address, error) {
	return _Ve.Contract.Voter(&_Ve.CallOpts)
}

// Abstain is a paid mutator transaction binding the contract method 0xc1f0fb9f.
//
// Solidity: function abstain(uint256 _tokenId) returns()
func (_Ve *VeTransactor) Abstain(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "abstain", _tokenId)
}

// Abstain is a paid mutator transaction binding the contract method 0xc1f0fb9f.
//
// Solidity: function abstain(uint256 _tokenId) returns()
func (_Ve *VeSession) Abstain(_tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.Abstain(&_Ve.TransactOpts, _tokenId)
}

// Abstain is a paid mutator transaction binding the contract method 0xc1f0fb9f.
//
// Solidity: function abstain(uint256 _tokenId) returns()
func (_Ve *VeTransactorSession) Abstain(_tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.Abstain(&_Ve.TransactOpts, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) returns()
func (_Ve *VeTransactor) Approve(opts *bind.TransactOpts, _approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "approve", _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) returns()
func (_Ve *VeSession) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.Approve(&_Ve.TransactOpts, _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) returns()
func (_Ve *VeTransactorSession) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.Approve(&_Ve.TransactOpts, _approved, _tokenId)
}

// Attach is a paid mutator transaction binding the contract method 0xfbd3a29d.
//
// Solidity: function attach(uint256 _tokenId) returns()
func (_Ve *VeTransactor) Attach(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "attach", _tokenId)
}

// Attach is a paid mutator transaction binding the contract method 0xfbd3a29d.
//
// Solidity: function attach(uint256 _tokenId) returns()
func (_Ve *VeSession) Attach(_tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.Attach(&_Ve.TransactOpts, _tokenId)
}

// Attach is a paid mutator transaction binding the contract method 0xfbd3a29d.
//
// Solidity: function attach(uint256 _tokenId) returns()
func (_Ve *VeTransactorSession) Attach(_tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.Attach(&_Ve.TransactOpts, _tokenId)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Ve *VeTransactor) Checkpoint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "checkpoint")
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Ve *VeSession) Checkpoint() (*types.Transaction, error) {
	return _Ve.Contract.Checkpoint(&_Ve.TransactOpts)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Ve *VeTransactorSession) Checkpoint() (*types.Transaction, error) {
	return _Ve.Contract.Checkpoint(&_Ve.TransactOpts)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _lock_duration) returns(uint256)
func (_Ve *VeTransactor) CreateLock(opts *bind.TransactOpts, _value *big.Int, _lock_duration *big.Int) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "create_lock", _value, _lock_duration)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _lock_duration) returns(uint256)
func (_Ve *VeSession) CreateLock(_value *big.Int, _lock_duration *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.CreateLock(&_Ve.TransactOpts, _value, _lock_duration)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _lock_duration) returns(uint256)
func (_Ve *VeTransactorSession) CreateLock(_value *big.Int, _lock_duration *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.CreateLock(&_Ve.TransactOpts, _value, _lock_duration)
}

// CreateLockFor is a paid mutator transaction binding the contract method 0xd4e54c3b.
//
// Solidity: function create_lock_for(uint256 _value, uint256 _lock_duration, address _to) returns(uint256)
func (_Ve *VeTransactor) CreateLockFor(opts *bind.TransactOpts, _value *big.Int, _lock_duration *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "create_lock_for", _value, _lock_duration, _to)
}

// CreateLockFor is a paid mutator transaction binding the contract method 0xd4e54c3b.
//
// Solidity: function create_lock_for(uint256 _value, uint256 _lock_duration, address _to) returns(uint256)
func (_Ve *VeSession) CreateLockFor(_value *big.Int, _lock_duration *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Ve.Contract.CreateLockFor(&_Ve.TransactOpts, _value, _lock_duration, _to)
}

// CreateLockFor is a paid mutator transaction binding the contract method 0xd4e54c3b.
//
// Solidity: function create_lock_for(uint256 _value, uint256 _lock_duration, address _to) returns(uint256)
func (_Ve *VeTransactorSession) CreateLockFor(_value *big.Int, _lock_duration *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Ve.Contract.CreateLockFor(&_Ve.TransactOpts, _value, _lock_duration, _to)
}

// DepositFor is a paid mutator transaction binding the contract method 0xee99fe28.
//
// Solidity: function deposit_for(uint256 _tokenId, uint256 _value) returns()
func (_Ve *VeTransactor) DepositFor(opts *bind.TransactOpts, _tokenId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "deposit_for", _tokenId, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0xee99fe28.
//
// Solidity: function deposit_for(uint256 _tokenId, uint256 _value) returns()
func (_Ve *VeSession) DepositFor(_tokenId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.DepositFor(&_Ve.TransactOpts, _tokenId, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0xee99fe28.
//
// Solidity: function deposit_for(uint256 _tokenId, uint256 _value) returns()
func (_Ve *VeTransactorSession) DepositFor(_tokenId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.DepositFor(&_Ve.TransactOpts, _tokenId, _value)
}

// Detach is a paid mutator transaction binding the contract method 0x986b7d8a.
//
// Solidity: function detach(uint256 _tokenId) returns()
func (_Ve *VeTransactor) Detach(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "detach", _tokenId)
}

// Detach is a paid mutator transaction binding the contract method 0x986b7d8a.
//
// Solidity: function detach(uint256 _tokenId) returns()
func (_Ve *VeSession) Detach(_tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.Detach(&_Ve.TransactOpts, _tokenId)
}

// Detach is a paid mutator transaction binding the contract method 0x986b7d8a.
//
// Solidity: function detach(uint256 _tokenId) returns()
func (_Ve *VeTransactorSession) Detach(_tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.Detach(&_Ve.TransactOpts, _tokenId)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0xa183af52.
//
// Solidity: function increase_amount(uint256 _tokenId, uint256 _value) returns()
func (_Ve *VeTransactor) IncreaseAmount(opts *bind.TransactOpts, _tokenId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "increase_amount", _tokenId, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0xa183af52.
//
// Solidity: function increase_amount(uint256 _tokenId, uint256 _value) returns()
func (_Ve *VeSession) IncreaseAmount(_tokenId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.IncreaseAmount(&_Ve.TransactOpts, _tokenId, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0xa183af52.
//
// Solidity: function increase_amount(uint256 _tokenId, uint256 _value) returns()
func (_Ve *VeTransactorSession) IncreaseAmount(_tokenId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.IncreaseAmount(&_Ve.TransactOpts, _tokenId, _value)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xa4d855df.
//
// Solidity: function increase_unlock_time(uint256 _tokenId, uint256 _lock_duration) returns()
func (_Ve *VeTransactor) IncreaseUnlockTime(opts *bind.TransactOpts, _tokenId *big.Int, _lock_duration *big.Int) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "increase_unlock_time", _tokenId, _lock_duration)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xa4d855df.
//
// Solidity: function increase_unlock_time(uint256 _tokenId, uint256 _lock_duration) returns()
func (_Ve *VeSession) IncreaseUnlockTime(_tokenId *big.Int, _lock_duration *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.IncreaseUnlockTime(&_Ve.TransactOpts, _tokenId, _lock_duration)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xa4d855df.
//
// Solidity: function increase_unlock_time(uint256 _tokenId, uint256 _lock_duration) returns()
func (_Ve *VeTransactorSession) IncreaseUnlockTime(_tokenId *big.Int, _lock_duration *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.IncreaseUnlockTime(&_Ve.TransactOpts, _tokenId, _lock_duration)
}

// Merge is a paid mutator transaction binding the contract method 0xd1c2babb.
//
// Solidity: function merge(uint256 _from, uint256 _to) returns()
func (_Ve *VeTransactor) Merge(opts *bind.TransactOpts, _from *big.Int, _to *big.Int) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "merge", _from, _to)
}

// Merge is a paid mutator transaction binding the contract method 0xd1c2babb.
//
// Solidity: function merge(uint256 _from, uint256 _to) returns()
func (_Ve *VeSession) Merge(_from *big.Int, _to *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.Merge(&_Ve.TransactOpts, _from, _to)
}

// Merge is a paid mutator transaction binding the contract method 0xd1c2babb.
//
// Solidity: function merge(uint256 _from, uint256 _to) returns()
func (_Ve *VeTransactorSession) Merge(_from *big.Int, _to *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.Merge(&_Ve.TransactOpts, _from, _to)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Ve *VeTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Ve *VeSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.SafeTransferFrom(&_Ve.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Ve *VeTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.SafeTransferFrom(&_Ve.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Ve *VeTransactor) SafeTransferFrom0(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "safeTransferFrom0", _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Ve *VeSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ve.Contract.SafeTransferFrom0(&_Ve.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Ve *VeTransactorSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ve.Contract.SafeTransferFrom0(&_Ve.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Ve *VeTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Ve *VeSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Ve.Contract.SetApprovalForAll(&_Ve.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Ve *VeTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Ve.Contract.SetApprovalForAll(&_Ve.TransactOpts, _operator, _approved)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _voter) returns()
func (_Ve *VeTransactor) SetVoter(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "setVoter", _voter)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _voter) returns()
func (_Ve *VeSession) SetVoter(_voter common.Address) (*types.Transaction, error) {
	return _Ve.Contract.SetVoter(&_Ve.TransactOpts, _voter)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _voter) returns()
func (_Ve *VeTransactorSession) SetVoter(_voter common.Address) (*types.Transaction, error) {
	return _Ve.Contract.SetVoter(&_Ve.TransactOpts, _voter)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Ve *VeTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Ve *VeSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.TransferFrom(&_Ve.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Ve *VeTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.TransferFrom(&_Ve.TransactOpts, _from, _to, _tokenId)
}

// Voting is a paid mutator transaction binding the contract method 0xfd4a77f1.
//
// Solidity: function voting(uint256 _tokenId) returns()
func (_Ve *VeTransactor) Voting(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "voting", _tokenId)
}

// Voting is a paid mutator transaction binding the contract method 0xfd4a77f1.
//
// Solidity: function voting(uint256 _tokenId) returns()
func (_Ve *VeSession) Voting(_tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.Voting(&_Ve.TransactOpts, _tokenId)
}

// Voting is a paid mutator transaction binding the contract method 0xfd4a77f1.
//
// Solidity: function voting(uint256 _tokenId) returns()
func (_Ve *VeTransactorSession) Voting(_tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.Voting(&_Ve.TransactOpts, _tokenId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _tokenId) returns()
func (_Ve *VeTransactor) Withdraw(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.contract.Transact(opts, "withdraw", _tokenId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _tokenId) returns()
func (_Ve *VeSession) Withdraw(_tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.Withdraw(&_Ve.TransactOpts, _tokenId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _tokenId) returns()
func (_Ve *VeTransactorSession) Withdraw(_tokenId *big.Int) (*types.Transaction, error) {
	return _Ve.Contract.Withdraw(&_Ve.TransactOpts, _tokenId)
}

// VeAbstainIterator is returned from FilterAbstain and is used to iterate over the raw logs and unpacked data for Abstain events raised by the Ve contract.
type VeAbstainIterator struct {
	Event *VeAbstain // Event containing the contract specifics and raw log

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
func (it *VeAbstainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeAbstain)
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
		it.Event = new(VeAbstain)
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
func (it *VeAbstainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeAbstainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeAbstain represents a Abstain event raised by the Ve contract.
type VeAbstain struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAbstain is a free log retrieval operation binding the contract event 0x52f55c75573f8a773844e0c57474b4518e27fc4c3982e4ad60ed06af1f8bc06b.
//
// Solidity: event Abstain(uint256 tokenId)
func (_Ve *VeFilterer) FilterAbstain(opts *bind.FilterOpts) (*VeAbstainIterator, error) {

	logs, sub, err := _Ve.contract.FilterLogs(opts, "Abstain")
	if err != nil {
		return nil, err
	}
	return &VeAbstainIterator{contract: _Ve.contract, event: "Abstain", logs: logs, sub: sub}, nil
}

// WatchAbstain is a free log subscription operation binding the contract event 0x52f55c75573f8a773844e0c57474b4518e27fc4c3982e4ad60ed06af1f8bc06b.
//
// Solidity: event Abstain(uint256 tokenId)
func (_Ve *VeFilterer) WatchAbstain(opts *bind.WatchOpts, sink chan<- *VeAbstain) (event.Subscription, error) {

	logs, sub, err := _Ve.contract.WatchLogs(opts, "Abstain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeAbstain)
				if err := _Ve.contract.UnpackLog(event, "Abstain", log); err != nil {
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

// ParseAbstain is a log parse operation binding the contract event 0x52f55c75573f8a773844e0c57474b4518e27fc4c3982e4ad60ed06af1f8bc06b.
//
// Solidity: event Abstain(uint256 tokenId)
func (_Ve *VeFilterer) ParseAbstain(log types.Log) (*VeAbstain, error) {
	event := new(VeAbstain)
	if err := _Ve.contract.UnpackLog(event, "Abstain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ve contract.
type VeApprovalIterator struct {
	Event *VeApproval // Event containing the contract specifics and raw log

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
func (it *VeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeApproval)
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
		it.Event = new(VeApproval)
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
func (it *VeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeApproval represents a Approval event raised by the Ve contract.
type VeApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Ve *VeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*VeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Ve.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VeApprovalIterator{contract: _Ve.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Ve *VeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VeApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Ve.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeApproval)
				if err := _Ve.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Ve *VeFilterer) ParseApproval(log types.Log) (*VeApproval, error) {
	event := new(VeApproval)
	if err := _Ve.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Ve contract.
type VeApprovalForAllIterator struct {
	Event *VeApprovalForAll // Event containing the contract specifics and raw log

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
func (it *VeApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeApprovalForAll)
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
		it.Event = new(VeApprovalForAll)
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
func (it *VeApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeApprovalForAll represents a ApprovalForAll event raised by the Ve contract.
type VeApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Ve *VeFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*VeApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Ve.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &VeApprovalForAllIterator{contract: _Ve.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Ve *VeFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *VeApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Ve.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeApprovalForAll)
				if err := _Ve.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Ve *VeFilterer) ParseApprovalForAll(log types.Log) (*VeApprovalForAll, error) {
	event := new(VeApprovalForAll)
	if err := _Ve.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeAttachIterator is returned from FilterAttach and is used to iterate over the raw logs and unpacked data for Attach events raised by the Ve contract.
type VeAttachIterator struct {
	Event *VeAttach // Event containing the contract specifics and raw log

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
func (it *VeAttachIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeAttach)
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
		it.Event = new(VeAttach)
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
func (it *VeAttachIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeAttachIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeAttach represents a Attach event raised by the Ve contract.
type VeAttach struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAttach is a free log retrieval operation binding the contract event 0xe09f878577f252e784dfb3291480806d7a4ce506a7ebb7af0455aac9123a53b2.
//
// Solidity: event Attach(uint256 tokenId)
func (_Ve *VeFilterer) FilterAttach(opts *bind.FilterOpts) (*VeAttachIterator, error) {

	logs, sub, err := _Ve.contract.FilterLogs(opts, "Attach")
	if err != nil {
		return nil, err
	}
	return &VeAttachIterator{contract: _Ve.contract, event: "Attach", logs: logs, sub: sub}, nil
}

// WatchAttach is a free log subscription operation binding the contract event 0xe09f878577f252e784dfb3291480806d7a4ce506a7ebb7af0455aac9123a53b2.
//
// Solidity: event Attach(uint256 tokenId)
func (_Ve *VeFilterer) WatchAttach(opts *bind.WatchOpts, sink chan<- *VeAttach) (event.Subscription, error) {

	logs, sub, err := _Ve.contract.WatchLogs(opts, "Attach")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeAttach)
				if err := _Ve.contract.UnpackLog(event, "Attach", log); err != nil {
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

// ParseAttach is a log parse operation binding the contract event 0xe09f878577f252e784dfb3291480806d7a4ce506a7ebb7af0455aac9123a53b2.
//
// Solidity: event Attach(uint256 tokenId)
func (_Ve *VeFilterer) ParseAttach(log types.Log) (*VeAttach, error) {
	event := new(VeAttach)
	if err := _Ve.contract.UnpackLog(event, "Attach", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Ve contract.
type VeDepositIterator struct {
	Event *VeDeposit // Event containing the contract specifics and raw log

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
func (it *VeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeDeposit)
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
		it.Event = new(VeDeposit)
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
func (it *VeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeDeposit represents a Deposit event raised by the Ve contract.
type VeDeposit struct {
	Provider    common.Address
	TokenId     *big.Int
	Value       *big.Int
	Locktime    *big.Int
	DepositType uint8
	Ts          *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xff04ccafc360e16b67d682d17bd9503c4c6b9a131f6be6325762dc9ffc7de624.
//
// Solidity: event Deposit(address indexed provider, uint256 tokenId, uint256 value, uint256 indexed locktime, uint8 deposit_type, uint256 ts)
func (_Ve *VeFilterer) FilterDeposit(opts *bind.FilterOpts, provider []common.Address, locktime []*big.Int) (*VeDepositIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _Ve.contract.FilterLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return &VeDepositIterator{contract: _Ve.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xff04ccafc360e16b67d682d17bd9503c4c6b9a131f6be6325762dc9ffc7de624.
//
// Solidity: event Deposit(address indexed provider, uint256 tokenId, uint256 value, uint256 indexed locktime, uint8 deposit_type, uint256 ts)
func (_Ve *VeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *VeDeposit, provider []common.Address, locktime []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _Ve.contract.WatchLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeDeposit)
				if err := _Ve.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xff04ccafc360e16b67d682d17bd9503c4c6b9a131f6be6325762dc9ffc7de624.
//
// Solidity: event Deposit(address indexed provider, uint256 tokenId, uint256 value, uint256 indexed locktime, uint8 deposit_type, uint256 ts)
func (_Ve *VeFilterer) ParseDeposit(log types.Log) (*VeDeposit, error) {
	event := new(VeDeposit)
	if err := _Ve.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeDetachIterator is returned from FilterDetach and is used to iterate over the raw logs and unpacked data for Detach events raised by the Ve contract.
type VeDetachIterator struct {
	Event *VeDetach // Event containing the contract specifics and raw log

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
func (it *VeDetachIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeDetach)
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
		it.Event = new(VeDetach)
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
func (it *VeDetachIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeDetachIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeDetach represents a Detach event raised by the Ve contract.
type VeDetach struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDetach is a free log retrieval operation binding the contract event 0x78d521759fa7df961323979c63cb12f52e11c03f1832ff6bd20dc9de50196f75.
//
// Solidity: event Detach(uint256 tokenId)
func (_Ve *VeFilterer) FilterDetach(opts *bind.FilterOpts) (*VeDetachIterator, error) {

	logs, sub, err := _Ve.contract.FilterLogs(opts, "Detach")
	if err != nil {
		return nil, err
	}
	return &VeDetachIterator{contract: _Ve.contract, event: "Detach", logs: logs, sub: sub}, nil
}

// WatchDetach is a free log subscription operation binding the contract event 0x78d521759fa7df961323979c63cb12f52e11c03f1832ff6bd20dc9de50196f75.
//
// Solidity: event Detach(uint256 tokenId)
func (_Ve *VeFilterer) WatchDetach(opts *bind.WatchOpts, sink chan<- *VeDetach) (event.Subscription, error) {

	logs, sub, err := _Ve.contract.WatchLogs(opts, "Detach")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeDetach)
				if err := _Ve.contract.UnpackLog(event, "Detach", log); err != nil {
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

// ParseDetach is a log parse operation binding the contract event 0x78d521759fa7df961323979c63cb12f52e11c03f1832ff6bd20dc9de50196f75.
//
// Solidity: event Detach(uint256 tokenId)
func (_Ve *VeFilterer) ParseDetach(log types.Log) (*VeDetach, error) {
	event := new(VeDetach)
	if err := _Ve.contract.UnpackLog(event, "Detach", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeSetVoterIterator is returned from FilterSetVoter and is used to iterate over the raw logs and unpacked data for SetVoter events raised by the Ve contract.
type VeSetVoterIterator struct {
	Event *VeSetVoter // Event containing the contract specifics and raw log

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
func (it *VeSetVoterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeSetVoter)
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
		it.Event = new(VeSetVoter)
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
func (it *VeSetVoterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeSetVoterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeSetVoter represents a SetVoter event raised by the Ve contract.
type VeSetVoter struct {
	Voter common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetVoter is a free log retrieval operation binding the contract event 0xc6ff127433b785c51da9ae4088ee184c909b1a55b9afd82ae6c64224d3bc15d2.
//
// Solidity: event SetVoter(address indexed voter)
func (_Ve *VeFilterer) FilterSetVoter(opts *bind.FilterOpts, voter []common.Address) (*VeSetVoterIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Ve.contract.FilterLogs(opts, "SetVoter", voterRule)
	if err != nil {
		return nil, err
	}
	return &VeSetVoterIterator{contract: _Ve.contract, event: "SetVoter", logs: logs, sub: sub}, nil
}

// WatchSetVoter is a free log subscription operation binding the contract event 0xc6ff127433b785c51da9ae4088ee184c909b1a55b9afd82ae6c64224d3bc15d2.
//
// Solidity: event SetVoter(address indexed voter)
func (_Ve *VeFilterer) WatchSetVoter(opts *bind.WatchOpts, sink chan<- *VeSetVoter, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Ve.contract.WatchLogs(opts, "SetVoter", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeSetVoter)
				if err := _Ve.contract.UnpackLog(event, "SetVoter", log); err != nil {
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

// ParseSetVoter is a log parse operation binding the contract event 0xc6ff127433b785c51da9ae4088ee184c909b1a55b9afd82ae6c64224d3bc15d2.
//
// Solidity: event SetVoter(address indexed voter)
func (_Ve *VeFilterer) ParseSetVoter(log types.Log) (*VeSetVoter, error) {
	event := new(VeSetVoter)
	if err := _Ve.contract.UnpackLog(event, "SetVoter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the Ve contract.
type VeSupplyIterator struct {
	Event *VeSupply // Event containing the contract specifics and raw log

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
func (it *VeSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeSupply)
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
		it.Event = new(VeSupply)
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
func (it *VeSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeSupply represents a Supply event raised by the Ve contract.
type VeSupply struct {
	PrevSupply *big.Int
	Supply     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Ve *VeFilterer) FilterSupply(opts *bind.FilterOpts) (*VeSupplyIterator, error) {

	logs, sub, err := _Ve.contract.FilterLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return &VeSupplyIterator{contract: _Ve.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Ve *VeFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *VeSupply) (event.Subscription, error) {

	logs, sub, err := _Ve.contract.WatchLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeSupply)
				if err := _Ve.contract.UnpackLog(event, "Supply", log); err != nil {
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

// ParseSupply is a log parse operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Ve *VeFilterer) ParseSupply(log types.Log) (*VeSupply, error) {
	event := new(VeSupply)
	if err := _Ve.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ve contract.
type VeTransferIterator struct {
	Event *VeTransfer // Event containing the contract specifics and raw log

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
func (it *VeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeTransfer)
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
		it.Event = new(VeTransfer)
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
func (it *VeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeTransfer represents a Transfer event raised by the Ve contract.
type VeTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Ve *VeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*VeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Ve.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VeTransferIterator{contract: _Ve.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Ve *VeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VeTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Ve.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeTransfer)
				if err := _Ve.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Ve *VeFilterer) ParseTransfer(log types.Log) (*VeTransfer, error) {
	event := new(VeTransfer)
	if err := _Ve.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeVotingIterator is returned from FilterVoting and is used to iterate over the raw logs and unpacked data for Voting events raised by the Ve contract.
type VeVotingIterator struct {
	Event *VeVoting // Event containing the contract specifics and raw log

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
func (it *VeVotingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeVoting)
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
		it.Event = new(VeVoting)
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
func (it *VeVotingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeVotingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeVoting represents a Voting event raised by the Ve contract.
type VeVoting struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVoting is a free log retrieval operation binding the contract event 0x14ad08ba726cc02a573ab7d43e86fac0315865fcc1450a792ea6659f0c6d38b1.
//
// Solidity: event Voting(uint256 tokenId)
func (_Ve *VeFilterer) FilterVoting(opts *bind.FilterOpts) (*VeVotingIterator, error) {

	logs, sub, err := _Ve.contract.FilterLogs(opts, "Voting")
	if err != nil {
		return nil, err
	}
	return &VeVotingIterator{contract: _Ve.contract, event: "Voting", logs: logs, sub: sub}, nil
}

// WatchVoting is a free log subscription operation binding the contract event 0x14ad08ba726cc02a573ab7d43e86fac0315865fcc1450a792ea6659f0c6d38b1.
//
// Solidity: event Voting(uint256 tokenId)
func (_Ve *VeFilterer) WatchVoting(opts *bind.WatchOpts, sink chan<- *VeVoting) (event.Subscription, error) {

	logs, sub, err := _Ve.contract.WatchLogs(opts, "Voting")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeVoting)
				if err := _Ve.contract.UnpackLog(event, "Voting", log); err != nil {
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

// ParseVoting is a log parse operation binding the contract event 0x14ad08ba726cc02a573ab7d43e86fac0315865fcc1450a792ea6659f0c6d38b1.
//
// Solidity: event Voting(uint256 tokenId)
func (_Ve *VeFilterer) ParseVoting(log types.Log) (*VeVoting, error) {
	event := new(VeVoting)
	if err := _Ve.contract.UnpackLog(event, "Voting", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Ve contract.
type VeWithdrawIterator struct {
	Event *VeWithdraw // Event containing the contract specifics and raw log

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
func (it *VeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeWithdraw)
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
		it.Event = new(VeWithdraw)
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
func (it *VeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeWithdraw represents a Withdraw event raised by the Ve contract.
type VeWithdraw struct {
	Provider common.Address
	TokenId  *big.Int
	Value    *big.Int
	Ts       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed provider, uint256 tokenId, uint256 value, uint256 ts)
func (_Ve *VeFilterer) FilterWithdraw(opts *bind.FilterOpts, provider []common.Address) (*VeWithdrawIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Ve.contract.FilterLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return &VeWithdrawIterator{contract: _Ve.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed provider, uint256 tokenId, uint256 value, uint256 ts)
func (_Ve *VeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *VeWithdraw, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Ve.contract.WatchLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeWithdraw)
				if err := _Ve.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed provider, uint256 tokenId, uint256 value, uint256 ts)
func (_Ve *VeFilterer) ParseWithdraw(log types.Log) (*VeWithdraw, error) {
	event := new(VeWithdraw)
	if err := _Ve.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
