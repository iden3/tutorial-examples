// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sate

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SateABI is the input ABI used to generate the binding from.
const SateABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifierContractAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockN\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"StateUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"blockN\",\"type\":\"uint64\"}],\"name\":\"getStateDataByBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getStateDataById\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getStateDataByTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"getTransitionInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"setVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"isOldStateGenesis\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"name\":\"transitState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Sate is an auto generated Go binding around an Ethereum contract.
type Sate struct {
	SateCaller     // Read-only binding to the contract
	SateTransactor // Write-only binding to the contract
	SateFilterer   // Log filterer for contract events
}

// SateCaller is an auto generated read-only Go binding around an Ethereum contract.
type SateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SateSession struct {
	Contract     *Sate             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SateCallerSession struct {
	Contract *SateCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SateTransactorSession struct {
	Contract     *SateTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SateRaw is an auto generated low-level Go binding around an Ethereum contract.
type SateRaw struct {
	Contract *Sate // Generic contract binding to access the raw methods on
}

// SateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SateCallerRaw struct {
	Contract *SateCaller // Generic read-only contract binding to access the raw methods on
}

// SateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SateTransactorRaw struct {
	Contract *SateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSate creates a new instance of Sate, bound to a specific deployed contract.
func NewSate(address common.Address, backend bind.ContractBackend) (*Sate, error) {
	contract, err := bindSate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sate{SateCaller: SateCaller{contract: contract}, SateTransactor: SateTransactor{contract: contract}, SateFilterer: SateFilterer{contract: contract}}, nil
}

// NewSateCaller creates a new read-only instance of Sate, bound to a specific deployed contract.
func NewSateCaller(address common.Address, caller bind.ContractCaller) (*SateCaller, error) {
	contract, err := bindSate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SateCaller{contract: contract}, nil
}

// NewSateTransactor creates a new write-only instance of Sate, bound to a specific deployed contract.
func NewSateTransactor(address common.Address, transactor bind.ContractTransactor) (*SateTransactor, error) {
	contract, err := bindSate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SateTransactor{contract: contract}, nil
}

// NewSateFilterer creates a new log filterer instance of Sate, bound to a specific deployed contract.
func NewSateFilterer(address common.Address, filterer bind.ContractFilterer) (*SateFilterer, error) {
	contract, err := bindSate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SateFilterer{contract: contract}, nil
}

// bindSate binds a generic wrapper to an already deployed contract.
func bindSate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sate *SateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sate.Contract.SateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sate *SateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sate.Contract.SateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sate *SateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sate.Contract.SateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sate *SateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sate *SateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sate *SateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sate.Contract.contract.Transact(opts, method, params...)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 id) view returns(uint256)
func (_Sate *SateCaller) GetState(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Sate.contract.Call(opts, &out, "getState", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 id) view returns(uint256)
func (_Sate *SateSession) GetState(id *big.Int) (*big.Int, error) {
	return _Sate.Contract.GetState(&_Sate.CallOpts, id)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 id) view returns(uint256)
func (_Sate *SateCallerSession) GetState(id *big.Int) (*big.Int, error) {
	return _Sate.Contract.GetState(&_Sate.CallOpts, id)
}

// GetStateDataByBlock is a free data retrieval call binding the contract method 0xd8dcd971.
//
// Solidity: function getStateDataByBlock(uint256 id, uint64 blockN) view returns(uint64, uint64, uint256)
func (_Sate *SateCaller) GetStateDataByBlock(opts *bind.CallOpts, id *big.Int, blockN uint64) (uint64, uint64, *big.Int, error) {
	var out []interface{}
	err := _Sate.contract.Call(opts, &out, "getStateDataByBlock", id, blockN)

	if err != nil {
		return *new(uint64), *new(uint64), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetStateDataByBlock is a free data retrieval call binding the contract method 0xd8dcd971.
//
// Solidity: function getStateDataByBlock(uint256 id, uint64 blockN) view returns(uint64, uint64, uint256)
func (_Sate *SateSession) GetStateDataByBlock(id *big.Int, blockN uint64) (uint64, uint64, *big.Int, error) {
	return _Sate.Contract.GetStateDataByBlock(&_Sate.CallOpts, id, blockN)
}

// GetStateDataByBlock is a free data retrieval call binding the contract method 0xd8dcd971.
//
// Solidity: function getStateDataByBlock(uint256 id, uint64 blockN) view returns(uint64, uint64, uint256)
func (_Sate *SateCallerSession) GetStateDataByBlock(id *big.Int, blockN uint64) (uint64, uint64, *big.Int, error) {
	return _Sate.Contract.GetStateDataByBlock(&_Sate.CallOpts, id, blockN)
}

// GetStateDataById is a free data retrieval call binding the contract method 0xc8d1e53e.
//
// Solidity: function getStateDataById(uint256 id) view returns(uint64, uint64, uint256)
func (_Sate *SateCaller) GetStateDataById(opts *bind.CallOpts, id *big.Int) (uint64, uint64, *big.Int, error) {
	var out []interface{}
	err := _Sate.contract.Call(opts, &out, "getStateDataById", id)

	if err != nil {
		return *new(uint64), *new(uint64), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetStateDataById is a free data retrieval call binding the contract method 0xc8d1e53e.
//
// Solidity: function getStateDataById(uint256 id) view returns(uint64, uint64, uint256)
func (_Sate *SateSession) GetStateDataById(id *big.Int) (uint64, uint64, *big.Int, error) {
	return _Sate.Contract.GetStateDataById(&_Sate.CallOpts, id)
}

// GetStateDataById is a free data retrieval call binding the contract method 0xc8d1e53e.
//
// Solidity: function getStateDataById(uint256 id) view returns(uint64, uint64, uint256)
func (_Sate *SateCallerSession) GetStateDataById(id *big.Int) (uint64, uint64, *big.Int, error) {
	return _Sate.Contract.GetStateDataById(&_Sate.CallOpts, id)
}

// GetStateDataByTime is a free data retrieval call binding the contract method 0x0281bec2.
//
// Solidity: function getStateDataByTime(uint256 id, uint64 timestamp) view returns(uint64, uint64, uint256)
func (_Sate *SateCaller) GetStateDataByTime(opts *bind.CallOpts, id *big.Int, timestamp uint64) (uint64, uint64, *big.Int, error) {
	var out []interface{}
	err := _Sate.contract.Call(opts, &out, "getStateDataByTime", id, timestamp)

	if err != nil {
		return *new(uint64), *new(uint64), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetStateDataByTime is a free data retrieval call binding the contract method 0x0281bec2.
//
// Solidity: function getStateDataByTime(uint256 id, uint64 timestamp) view returns(uint64, uint64, uint256)
func (_Sate *SateSession) GetStateDataByTime(id *big.Int, timestamp uint64) (uint64, uint64, *big.Int, error) {
	return _Sate.Contract.GetStateDataByTime(&_Sate.CallOpts, id, timestamp)
}

// GetStateDataByTime is a free data retrieval call binding the contract method 0x0281bec2.
//
// Solidity: function getStateDataByTime(uint256 id, uint64 timestamp) view returns(uint64, uint64, uint256)
func (_Sate *SateCallerSession) GetStateDataByTime(id *big.Int, timestamp uint64) (uint64, uint64, *big.Int, error) {
	return _Sate.Contract.GetStateDataByTime(&_Sate.CallOpts, id, timestamp)
}

// GetTransitionInfo is a free data retrieval call binding the contract method 0xbb795715.
//
// Solidity: function getTransitionInfo(uint256 state) view returns(uint256, uint256, uint64, uint64, uint256, uint256)
func (_Sate *SateCaller) GetTransitionInfo(opts *bind.CallOpts, state *big.Int) (*big.Int, *big.Int, uint64, uint64, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Sate.contract.Call(opts, &out, "getTransitionInfo", state)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(uint64), *new(uint64), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new(uint64)).(*uint64)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetTransitionInfo is a free data retrieval call binding the contract method 0xbb795715.
//
// Solidity: function getTransitionInfo(uint256 state) view returns(uint256, uint256, uint64, uint64, uint256, uint256)
func (_Sate *SateSession) GetTransitionInfo(state *big.Int) (*big.Int, *big.Int, uint64, uint64, *big.Int, *big.Int, error) {
	return _Sate.Contract.GetTransitionInfo(&_Sate.CallOpts, state)
}

// GetTransitionInfo is a free data retrieval call binding the contract method 0xbb795715.
//
// Solidity: function getTransitionInfo(uint256 state) view returns(uint256, uint256, uint64, uint64, uint256, uint256)
func (_Sate *SateCallerSession) GetTransitionInfo(state *big.Int) (*big.Int, *big.Int, uint64, uint64, *big.Int, *big.Int, error) {
	return _Sate.Contract.GetTransitionInfo(&_Sate.CallOpts, state)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sate *SateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sate.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sate *SateSession) Owner() (common.Address, error) {
	return _Sate.Contract.Owner(&_Sate.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sate *SateCallerSession) Owner() (common.Address, error) {
	return _Sate.Contract.Owner(&_Sate.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Sate *SateCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sate.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Sate *SateSession) Verifier() (common.Address, error) {
	return _Sate.Contract.Verifier(&_Sate.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Sate *SateCallerSession) Verifier() (common.Address, error) {
	return _Sate.Contract.Verifier(&_Sate.CallOpts)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address newVerifier) returns()
func (_Sate *SateTransactor) SetVerifier(opts *bind.TransactOpts, newVerifier common.Address) (*types.Transaction, error) {
	return _Sate.contract.Transact(opts, "setVerifier", newVerifier)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address newVerifier) returns()
func (_Sate *SateSession) SetVerifier(newVerifier common.Address) (*types.Transaction, error) {
	return _Sate.Contract.SetVerifier(&_Sate.TransactOpts, newVerifier)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address newVerifier) returns()
func (_Sate *SateTransactorSession) SetVerifier(newVerifier common.Address) (*types.Transaction, error) {
	return _Sate.Contract.SetVerifier(&_Sate.TransactOpts, newVerifier)
}

// TransitState is a paid mutator transaction binding the contract method 0x89fbb92b.
//
// Solidity: function transitState(uint256 id, uint256 oldState, uint256 newState, uint256 isOldStateGenesis, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_Sate *SateTransactor) TransitState(opts *bind.TransactOpts, id *big.Int, oldState *big.Int, newState *big.Int, isOldStateGenesis *big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _Sate.contract.Transact(opts, "transitState", id, oldState, newState, isOldStateGenesis, a, b, c)
}

// TransitState is a paid mutator transaction binding the contract method 0x89fbb92b.
//
// Solidity: function transitState(uint256 id, uint256 oldState, uint256 newState, uint256 isOldStateGenesis, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_Sate *SateSession) TransitState(id *big.Int, oldState *big.Int, newState *big.Int, isOldStateGenesis *big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _Sate.Contract.TransitState(&_Sate.TransactOpts, id, oldState, newState, isOldStateGenesis, a, b, c)
}

// TransitState is a paid mutator transaction binding the contract method 0x89fbb92b.
//
// Solidity: function transitState(uint256 id, uint256 oldState, uint256 newState, uint256 isOldStateGenesis, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_Sate *SateTransactorSession) TransitState(id *big.Int, oldState *big.Int, newState *big.Int, isOldStateGenesis *big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _Sate.Contract.TransitState(&_Sate.TransactOpts, id, oldState, newState, isOldStateGenesis, a, b, c)
}

// SateStateUpdatedIterator is returned from FilterStateUpdated and is used to iterate over the raw logs and unpacked data for StateUpdated events raised by the Sate contract.
type SateStateUpdatedIterator struct {
	Event *SateStateUpdated // Event containing the contract specifics and raw log

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
func (it *SateStateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SateStateUpdated)
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
		it.Event = new(SateStateUpdated)
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
func (it *SateStateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SateStateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SateStateUpdated represents a StateUpdated event raised by the Sate contract.
type SateStateUpdated struct {
	Id        *big.Int
	BlockN    uint64
	Timestamp uint64
	State     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStateUpdated is a free log retrieval operation binding the contract event 0x81c6f328b24014ef550c34a433275b52f3a8a0f32aa871adec069ab526a02390.
//
// Solidity: event StateUpdated(uint256 id, uint64 blockN, uint64 timestamp, uint256 state)
func (_Sate *SateFilterer) FilterStateUpdated(opts *bind.FilterOpts) (*SateStateUpdatedIterator, error) {

	logs, sub, err := _Sate.contract.FilterLogs(opts, "StateUpdated")
	if err != nil {
		return nil, err
	}
	return &SateStateUpdatedIterator{contract: _Sate.contract, event: "StateUpdated", logs: logs, sub: sub}, nil
}

// WatchStateUpdated is a free log subscription operation binding the contract event 0x81c6f328b24014ef550c34a433275b52f3a8a0f32aa871adec069ab526a02390.
//
// Solidity: event StateUpdated(uint256 id, uint64 blockN, uint64 timestamp, uint256 state)
func (_Sate *SateFilterer) WatchStateUpdated(opts *bind.WatchOpts, sink chan<- *SateStateUpdated) (event.Subscription, error) {

	logs, sub, err := _Sate.contract.WatchLogs(opts, "StateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SateStateUpdated)
				if err := _Sate.contract.UnpackLog(event, "StateUpdated", log); err != nil {
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

// ParseStateUpdated is a log parse operation binding the contract event 0x81c6f328b24014ef550c34a433275b52f3a8a0f32aa871adec069ab526a02390.
//
// Solidity: event StateUpdated(uint256 id, uint64 blockN, uint64 timestamp, uint256 state)
func (_Sate *SateFilterer) ParseStateUpdated(log types.Log) (*SateStateUpdated, error) {
	event := new(SateStateUpdated)
	if err := _Sate.contract.UnpackLog(event, "StateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
