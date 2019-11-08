// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smarttoken

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

// SmarttokenABI is the input ABI used to generate the binding from.
const SmarttokenABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"receiveApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Smarttoken is an auto generated Go binding around an Ethereum contract.
type Smarttoken struct {
	SmarttokenCaller     // Read-only binding to the contract
	SmarttokenTransactor // Write-only binding to the contract
	SmarttokenFilterer   // Log filterer for contract events
}

// SmarttokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmarttokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmarttokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmarttokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmarttokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmarttokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmarttokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmarttokenSession struct {
	Contract     *Smarttoken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmarttokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmarttokenCallerSession struct {
	Contract *SmarttokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SmarttokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmarttokenTransactorSession struct {
	Contract     *SmarttokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SmarttokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmarttokenRaw struct {
	Contract *Smarttoken // Generic contract binding to access the raw methods on
}

// SmarttokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmarttokenCallerRaw struct {
	Contract *SmarttokenCaller // Generic read-only contract binding to access the raw methods on
}

// SmarttokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmarttokenTransactorRaw struct {
	Contract *SmarttokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmarttoken creates a new instance of Smarttoken, bound to a specific deployed contract.
func NewSmarttoken(address common.Address, backend bind.ContractBackend) (*Smarttoken, error) {
	contract, err := bindSmarttoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Smarttoken{SmarttokenCaller: SmarttokenCaller{contract: contract}, SmarttokenTransactor: SmarttokenTransactor{contract: contract}, SmarttokenFilterer: SmarttokenFilterer{contract: contract}}, nil
}

// NewSmarttokenCaller creates a new read-only instance of Smarttoken, bound to a specific deployed contract.
func NewSmarttokenCaller(address common.Address, caller bind.ContractCaller) (*SmarttokenCaller, error) {
	contract, err := bindSmarttoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmarttokenCaller{contract: contract}, nil
}

// NewSmarttokenTransactor creates a new write-only instance of Smarttoken, bound to a specific deployed contract.
func NewSmarttokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SmarttokenTransactor, error) {
	contract, err := bindSmarttoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmarttokenTransactor{contract: contract}, nil
}

// NewSmarttokenFilterer creates a new log filterer instance of Smarttoken, bound to a specific deployed contract.
func NewSmarttokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SmarttokenFilterer, error) {
	contract, err := bindSmarttoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmarttokenFilterer{contract: contract}, nil
}

// bindSmarttoken binds a generic wrapper to an already deployed contract.
func bindSmarttoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmarttokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smarttoken *SmarttokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Smarttoken.Contract.SmarttokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smarttoken *SmarttokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smarttoken.Contract.SmarttokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smarttoken *SmarttokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smarttoken.Contract.SmarttokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smarttoken *SmarttokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Smarttoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smarttoken *SmarttokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smarttoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smarttoken *SmarttokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smarttoken.Contract.contract.Transact(opts, method, params...)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(address from, uint256 tokens, address token, bytes data) returns()
func (_Smarttoken *SmarttokenTransactor) ReceiveApproval(opts *bind.TransactOpts, from common.Address, tokens *big.Int, token common.Address, data []byte) (*types.Transaction, error) {
	return _Smarttoken.contract.Transact(opts, "receiveApproval", from, tokens, token, data)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(address from, uint256 tokens, address token, bytes data) returns()
func (_Smarttoken *SmarttokenSession) ReceiveApproval(from common.Address, tokens *big.Int, token common.Address, data []byte) (*types.Transaction, error) {
	return _Smarttoken.Contract.ReceiveApproval(&_Smarttoken.TransactOpts, from, tokens, token, data)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(address from, uint256 tokens, address token, bytes data) returns()
func (_Smarttoken *SmarttokenTransactorSession) ReceiveApproval(from common.Address, tokens *big.Int, token common.Address, data []byte) (*types.Transaction, error) {
	return _Smarttoken.Contract.ReceiveApproval(&_Smarttoken.TransactOpts, from, tokens, token, data)
}
