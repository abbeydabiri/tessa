// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smartcontracts

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ApproveandcallfallbackABI is the input ABI used to generate the binding from.
const ApproveandcallfallbackABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"receiveApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Approveandcallfallback is an auto generated Go binding around an Ethereum contract.
type Approveandcallfallback struct {
	ApproveandcallfallbackCaller     // Read-only binding to the contract
	ApproveandcallfallbackTransactor // Write-only binding to the contract
	ApproveandcallfallbackFilterer   // Log filterer for contract events
}

// ApproveandcallfallbackCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApproveandcallfallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApproveandcallfallbackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApproveandcallfallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApproveandcallfallbackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApproveandcallfallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApproveandcallfallbackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApproveandcallfallbackSession struct {
	Contract     *Approveandcallfallback // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ApproveandcallfallbackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApproveandcallfallbackCallerSession struct {
	Contract *ApproveandcallfallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ApproveandcallfallbackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApproveandcallfallbackTransactorSession struct {
	Contract     *ApproveandcallfallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ApproveandcallfallbackRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApproveandcallfallbackRaw struct {
	Contract *Approveandcallfallback // Generic contract binding to access the raw methods on
}

// ApproveandcallfallbackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApproveandcallfallbackCallerRaw struct {
	Contract *ApproveandcallfallbackCaller // Generic read-only contract binding to access the raw methods on
}

// ApproveandcallfallbackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApproveandcallfallbackTransactorRaw struct {
	Contract *ApproveandcallfallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApproveandcallfallback creates a new instance of Approveandcallfallback, bound to a specific deployed contract.
func NewApproveandcallfallback(address common.Address, backend bind.ContractBackend) (*Approveandcallfallback, error) {
	contract, err := bindApproveandcallfallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Approveandcallfallback{ApproveandcallfallbackCaller: ApproveandcallfallbackCaller{contract: contract}, ApproveandcallfallbackTransactor: ApproveandcallfallbackTransactor{contract: contract}, ApproveandcallfallbackFilterer: ApproveandcallfallbackFilterer{contract: contract}}, nil
}

// NewApproveandcallfallbackCaller creates a new read-only instance of Approveandcallfallback, bound to a specific deployed contract.
func NewApproveandcallfallbackCaller(address common.Address, caller bind.ContractCaller) (*ApproveandcallfallbackCaller, error) {
	contract, err := bindApproveandcallfallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApproveandcallfallbackCaller{contract: contract}, nil
}

// NewApproveandcallfallbackTransactor creates a new write-only instance of Approveandcallfallback, bound to a specific deployed contract.
func NewApproveandcallfallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*ApproveandcallfallbackTransactor, error) {
	contract, err := bindApproveandcallfallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApproveandcallfallbackTransactor{contract: contract}, nil
}

// NewApproveandcallfallbackFilterer creates a new log filterer instance of Approveandcallfallback, bound to a specific deployed contract.
func NewApproveandcallfallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*ApproveandcallfallbackFilterer, error) {
	contract, err := bindApproveandcallfallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApproveandcallfallbackFilterer{contract: contract}, nil
}

// bindApproveandcallfallback binds a generic wrapper to an already deployed contract.
func bindApproveandcallfallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApproveandcallfallbackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Approveandcallfallback *ApproveandcallfallbackRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Approveandcallfallback.Contract.ApproveandcallfallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Approveandcallfallback *ApproveandcallfallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Approveandcallfallback.Contract.ApproveandcallfallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Approveandcallfallback *ApproveandcallfallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Approveandcallfallback.Contract.ApproveandcallfallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Approveandcallfallback *ApproveandcallfallbackCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Approveandcallfallback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Approveandcallfallback *ApproveandcallfallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Approveandcallfallback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Approveandcallfallback *ApproveandcallfallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Approveandcallfallback.Contract.contract.Transact(opts, method, params...)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(address from, uint256 tokens, address token, bytes data) returns()
func (_Approveandcallfallback *ApproveandcallfallbackTransactor) ReceiveApproval(opts *bind.TransactOpts, from common.Address, tokens *big.Int, token common.Address, data []byte) (*types.Transaction, error) {
	return _Approveandcallfallback.contract.Transact(opts, "receiveApproval", from, tokens, token, data)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(address from, uint256 tokens, address token, bytes data) returns()
func (_Approveandcallfallback *ApproveandcallfallbackSession) ReceiveApproval(from common.Address, tokens *big.Int, token common.Address, data []byte) (*types.Transaction, error) {
	return _Approveandcallfallback.Contract.ReceiveApproval(&_Approveandcallfallback.TransactOpts, from, tokens, token, data)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(address from, uint256 tokens, address token, bytes data) returns()
func (_Approveandcallfallback *ApproveandcallfallbackTransactorSession) ReceiveApproval(from common.Address, tokens *big.Int, token common.Address, data []byte) (*types.Transaction, error) {
	return _Approveandcallfallback.Contract.ReceiveApproval(&_Approveandcallfallback.TransactOpts, from, tokens, token, data)
}
