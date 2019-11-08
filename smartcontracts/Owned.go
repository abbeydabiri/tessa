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
const SmarttokenABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SmarttokenBin is the compiled bytecode used for deploying new contracts.
var SmarttokenBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506103eb806100606000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806379ba50971461003b578063f2fde38b14610045575b600080fd5b610043610089565b005b6100876004803603602081101561005b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610272565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461012f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b81526020018061038c602b913960400191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610317576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603081526020018061035c6030913960400191505060405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fe4f6e6c7920636f6e7472616374206f776e65722063616e20706572666f726d2074686973207472616e73616374696f6e4f6e6c79206e6577206f776e65722063616e20706572666f726d2074686973207472616e73616374696f6ea265627a7a72315820316429997ad02139d97c276e9f22394138be7c536f810cd79a75275d6884c0fe64736f6c634300050c0032"

// DeploySmarttoken deploys a new Ethereum contract, binding an instance of Smarttoken to it.
func DeploySmarttoken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Smarttoken, error) {
	parsed, err := abi.JSON(strings.NewReader(SmarttokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SmarttokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Smarttoken{SmarttokenCaller: SmarttokenCaller{contract: contract}, SmarttokenTransactor: SmarttokenTransactor{contract: contract}, SmarttokenFilterer: SmarttokenFilterer{contract: contract}}, nil
}

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

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Smarttoken *SmarttokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smarttoken.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Smarttoken *SmarttokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _Smarttoken.Contract.AcceptOwnership(&_Smarttoken.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Smarttoken *SmarttokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Smarttoken.Contract.AcceptOwnership(&_Smarttoken.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Smarttoken *SmarttokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Smarttoken.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Smarttoken *SmarttokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Smarttoken.Contract.TransferOwnership(&_Smarttoken.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Smarttoken *SmarttokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Smarttoken.Contract.TransferOwnership(&_Smarttoken.TransactOpts, _newOwner)
}

// SmarttokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Smarttoken contract.
type SmarttokenOwnershipTransferredIterator struct {
	Event *SmarttokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SmarttokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmarttokenOwnershipTransferred)
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
		it.Event = new(SmarttokenOwnershipTransferred)
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
func (it *SmarttokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmarttokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmarttokenOwnershipTransferred represents a OwnershipTransferred event raised by the Smarttoken contract.
type SmarttokenOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _from, address indexed _to)
func (_Smarttoken *SmarttokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*SmarttokenOwnershipTransferredIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Smarttoken.contract.FilterLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &SmarttokenOwnershipTransferredIterator{contract: _Smarttoken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _from, address indexed _to)
func (_Smarttoken *SmarttokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SmarttokenOwnershipTransferred, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Smarttoken.contract.WatchLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmarttokenOwnershipTransferred)
				if err := _Smarttoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _from, address indexed _to)
func (_Smarttoken *SmarttokenFilterer) ParseOwnershipTransferred(log types.Log) (*SmarttokenOwnershipTransferred, error) {
	event := new(SmarttokenOwnershipTransferred)
	if err := _Smarttoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
