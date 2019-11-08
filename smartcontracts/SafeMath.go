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

// SafemathABI is the input ABI used to generate the binding from.
const SafemathABI = "[]"

// SafemathBin is the compiled bytecode used for deploying new contracts.
var SafemathBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a72315820b01f1bc24632d7564cc2c65a139b4c0dc3a5ad2fa49a498d7cdb224a6ca7d82164736f6c634300050c0032"

// DeploySafemath deploys a new Ethereum contract, binding an instance of Safemath to it.
func DeploySafemath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Safemath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafemathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafemathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Safemath{SafemathCaller: SafemathCaller{contract: contract}, SafemathTransactor: SafemathTransactor{contract: contract}, SafemathFilterer: SafemathFilterer{contract: contract}}, nil
}

// Safemath is an auto generated Go binding around an Ethereum contract.
type Safemath struct {
	SafemathCaller     // Read-only binding to the contract
	SafemathTransactor // Write-only binding to the contract
	SafemathFilterer   // Log filterer for contract events
}

// SafemathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafemathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafemathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafemathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafemathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafemathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafemathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafemathSession struct {
	Contract     *Safemath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafemathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafemathCallerSession struct {
	Contract *SafemathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafemathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafemathTransactorSession struct {
	Contract     *SafemathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafemathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafemathRaw struct {
	Contract *Safemath // Generic contract binding to access the raw methods on
}

// SafemathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafemathCallerRaw struct {
	Contract *SafemathCaller // Generic read-only contract binding to access the raw methods on
}

// SafemathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafemathTransactorRaw struct {
	Contract *SafemathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafemath creates a new instance of Safemath, bound to a specific deployed contract.
func NewSafemath(address common.Address, backend bind.ContractBackend) (*Safemath, error) {
	contract, err := bindSafemath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Safemath{SafemathCaller: SafemathCaller{contract: contract}, SafemathTransactor: SafemathTransactor{contract: contract}, SafemathFilterer: SafemathFilterer{contract: contract}}, nil
}

// NewSafemathCaller creates a new read-only instance of Safemath, bound to a specific deployed contract.
func NewSafemathCaller(address common.Address, caller bind.ContractCaller) (*SafemathCaller, error) {
	contract, err := bindSafemath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafemathCaller{contract: contract}, nil
}

// NewSafemathTransactor creates a new write-only instance of Safemath, bound to a specific deployed contract.
func NewSafemathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafemathTransactor, error) {
	contract, err := bindSafemath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafemathTransactor{contract: contract}, nil
}

// NewSafemathFilterer creates a new log filterer instance of Safemath, bound to a specific deployed contract.
func NewSafemathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafemathFilterer, error) {
	contract, err := bindSafemath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafemathFilterer{contract: contract}, nil
}

// bindSafemath binds a generic wrapper to an already deployed contract.
func bindSafemath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafemathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safemath *SafemathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Safemath.Contract.SafemathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safemath *SafemathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safemath.Contract.SafemathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safemath *SafemathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safemath.Contract.SafemathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safemath *SafemathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Safemath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safemath *SafemathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safemath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safemath *SafemathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safemath.Contract.contract.Transact(opts, method, params...)
}
