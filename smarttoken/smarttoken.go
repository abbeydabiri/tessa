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

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", tokenOwner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_ERC20 *ERC20Session) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_ERC20 *ERC20CallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, tokenOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", tokenOwner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_ERC20 *ERC20Session) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, tokenOwner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, tokens)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(tokenOwner indexed address, spender indexed address, tokens uint256)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(tokenOwner indexed address, spender indexed address, tokens uint256)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, tokens uint256)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, tokens uint256)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x608060405234801561001057600080fd5b50600080546001600160a01b0319163317905561023a806100326000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806379ba50971461003b578063f2fde38b14610045575b600080fd5b61004361006b565b005b6100436004803603602081101561005b57600080fd5b50356001600160a01b0316610118565b6001546001600160a01b031633146100b45760405162461bcd60e51b815260040180806020018281038252602b8152602001806101b4602b913960400191505060405180910390fd5b600154600080546040516001600160a01b0393841693909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a360018054600080546001600160a01b03199081166001600160a01b03841617909155169055565b6000546001600160a01b031633146101615760405162461bcd60e51b81526004018080602001828103825260308152602001806101846030913960400191505060405180910390fd5b600180546001600160a01b0319166001600160a01b039290921691909117905556fe4f6e6c7920636f6e7472616374206f776e65722063616e20706572666f726d2074686973207472616e73616374696f6e4f6e6c79206e6577206f776e65722063616e20706572666f726d2074686973207472616e73616374696f6ea265627a7a72315820edc9d5223e4f9a89cf122b173f255bf6985b77234fb80801a3e9e7c228b1779064736f6c637829302e352e31342d646576656c6f702e323031392e31312e31352b636f6d6d69742e33653034666436650059`

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
	OwnedFilterer   // Log filterer for contract events
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// NewOwnedFilterer creates a new log filterer instance of Owned, bound to a specific deployed contract.
func NewOwnedFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnedFilterer, error) {
	contract, err := bindOwned(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnedFilterer{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedSession) AcceptOwnership() (*types.Transaction, error) {
	return _Owned.Contract.AcceptOwnership(&_Owned.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Owned.Contract.AcceptOwnership(&_Owned.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Owned *OwnedTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Owned *OwnedSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.TransferOwnership(&_Owned.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Owned *OwnedTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.TransferOwnership(&_Owned.TransactOpts, _newOwner)
}

// OwnedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Owned contract.
type OwnedOwnershipTransferredIterator struct {
	Event *OwnedOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnedOwnershipTransferred)
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
		it.Event = new(OwnedOwnershipTransferred)
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
func (it *OwnedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnedOwnershipTransferred represents a OwnershipTransferred event raised by the Owned contract.
type OwnedOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_from indexed address, _to indexed address)
func (_Owned *OwnedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*OwnedOwnershipTransferredIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Owned.contract.FilterLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &OwnedOwnershipTransferredIterator{contract: _Owned.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_from indexed address, _to indexed address)
func (_Owned *OwnedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnedOwnershipTransferred, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Owned.contract.WatchLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnedOwnershipTransferred)
				if err := _Owned.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x6080604052348015600f57600080fd5b50606580601d6000396000f3fe6080604052600080fdfea265627a7a72315820b15582b54f899844cd9ce14469fa9135acee5be1e2b1b22294f62eadeab5507e64736f6c637829302e352e31342d646576656c6f702e323031392e31312e31352b636f6d6d69742e33653034666436650059`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SmartTokenABI is the input ABI used to generate the binding from.
const SmartTokenABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_maxTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_seed\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addToBlacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"blacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"curTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"mintTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"removeFromBlacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SmartTokenBin is the compiled bytecode used for deploying new contracts.
const SmartTokenBin = `0x60806040523480156200001157600080fd5b506040516200143c3803806200143c833981810160405260808110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b5060409081526020820151910151600080546001600160a01b031916331790556004805460ff1916601217815585519294509092501062000225576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180620013f16022913960400191505060405180910390fd5b600284511162000281576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180620013cd6024913960400191505060405180910390fd5b600a8211620002dc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180620014136029913960400191505060405180910390fd5b8251620002f190600290602086019062000372565b5083516200030790600390602087019062000372565b5060045460ff16600a0a8181026005819055908302600655336000818152600760209081526040808320859055805194855251929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050505062000417565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003b557805160ff1916838001178555620003e5565b82800160010185558215620003e5579182015b82811115620003e5578251825591602001919060010190620003c8565b50620003f3929150620003f7565b5090565b6200041491905b80821115620003f35760008155600101620003fe565b90565b610fa680620004276000396000f3fe6080604052600436106101095760003560e01c806370a0823111610095578063cae9ca5111610064578063cae9ca51146103cc578063dd62ed3e14610494578063f0dda65c146104cf578063f2fde38b14610508578063f9f92be41461053b57610109565b806370a082311461033457806379ba50971461036757806395d89b411461037e578063a9059cbb1461039357610109565b806323b872dd116100dc57806323b872dd1461024b5780632ab4d0521461028e578063313ce567146102a357806344337ea1146102ce578063537df3b61461030157610109565b806306fdde0314610138578063095ea7b3146101c257806309cc98541461020f57806318160ddd14610236575b60405133903480156108fc02916000818181858888f19350505050158015610135573d6000803e3d6000fd5b50005b34801561014457600080fd5b5061014d61056e565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561018757818101518382015260200161016f565b50505050905090810190601f1680156101b45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101ce57600080fd5b506101fb600480360360408110156101e557600080fd5b506001600160a01b0381351690602001356105f9565b604080519115158252519081900360200190f35b34801561021b57600080fd5b506102246106b0565b60408051918252519081900360200190f35b34801561024257600080fd5b506102246106b6565b34801561025757600080fd5b506101fb6004803603606081101561026e57600080fd5b506001600160a01b038135811691602081013590911690604001356106e8565b34801561029a57600080fd5b50610224610808565b3480156102af57600080fd5b506102b861080e565b6040805160ff9092168252519081900360200190f35b3480156102da57600080fd5b506101fb600480360360208110156102f157600080fd5b50356001600160a01b0316610817565b34801561030d57600080fd5b506101fb6004803603602081101561032457600080fd5b50356001600160a01b031661088a565b34801561034057600080fd5b506102246004803603602081101561035757600080fd5b50356001600160a01b03166108f9565b34801561037357600080fd5b5061037c610914565b005b34801561038a57600080fd5b5061014d6109c1565b34801561039f57600080fd5b506101fb600480360360408110156103b657600080fd5b506001600160a01b038135169060200135610a1c565b3480156103d857600080fd5b506101fb600480360360608110156103ef57600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561041f57600080fd5b82018360208201111561043157600080fd5b8035906020019184600183028401116401000000008311171561045357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610b0f945050505050565b3480156104a057600080fd5b50610224600480360360408110156104b757600080fd5b506001600160a01b0381358116916020013516610ca7565b3480156104db57600080fd5b506101fb600480360360408110156104f257600080fd5b506001600160a01b038135169060200135610cd2565b34801561051457600080fd5b5061037c6004803603602081101561052b57600080fd5b50356001600160a01b0316610e29565b34801561054757600080fd5b506101fb6004803603602081101561055e57600080fd5b50356001600160a01b0316610e94565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156105f15780601f106105c6576101008083540402835291602001916105f1565b820191906000526020600020905b8154815290600101906020018083116105d457829003601f168201915b505050505081565b3360009081526008602052604081205460ff16156106485760405162461bcd60e51b8152600401808060200182810382526021815260200180610eff6021913960400191505060405180910390fd5b3360008181526009602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b60055481565b6000805260076020527f6d5257204ebe7d88fd91ae87941cb2dd9d8062b64ae5a2bd2d28ec40b9fbf6df546005540390565b6001600160a01b03831660009081526007602052604081205461070b9083610ea9565b6001600160a01b038581166000908152600760205260408120929092559054163314610784576001600160a01b038416600090815260096020908152604080832033845290915290205461075f9083610ea9565b6001600160a01b03851660009081526009602090815260408083203384529091529020555b6001600160a01b0383166000908152600760205260409020546107a79083610ebe565b6001600160a01b0380851660008181526007602090815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a35060019392505050565b60065481565b60045460ff1681565b600080546001600160a01b031633146108615760405162461bcd60e51b8152600401808060200182810382526030815260200180610ecf6030913960400191505060405180910390fd5b506001600160a01b03166000908152600860205260409020805460ff1916600190811790915590565b600080546001600160a01b031633146108d45760405162461bcd60e51b8152600401808060200182810382526030815260200180610ecf6030913960400191505060405180910390fd5b506001600160a01b03166000908152600860205260409020805460ff19169055600190565b6001600160a01b031660009081526007602052604090205490565b6001546001600160a01b0316331461095d5760405162461bcd60e51b815260040180806020018281038252602b815260200180610f20602b913960400191505060405180910390fd5b600154600080546040516001600160a01b0393841693909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a360018054600080546001600160a01b03199081166001600160a01b03841617909155169055565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105f15780601f106105c6576101008083540402835291602001916105f1565b3360009081526008602052604081205460ff1615610a6b5760405162461bcd60e51b8152600401808060200182810382526021815260200180610eff6021913960400191505060405180910390fd5b33600090815260076020526040902054610a859083610ea9565b33600090815260076020526040808220929092556001600160a01b03851681522054610ab19083610ebe565b6001600160a01b0384166000818152600760209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b3360009081526008602052604081205460ff1615610b5e5760405162461bcd60e51b8152600401808060200182810382526021815260200180610eff6021913960400191505060405180910390fd5b3360008181526009602090815260408083206001600160a01b03891680855290835292819020879055805187815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a3604051638f4ffcb160e01b815233600482018181526024830186905230604484018190526080606485019081528651608486015286516001600160a01b038a1695638f4ffcb195948a94938a939192909160a490910190602085019080838360005b83811015610c36578181015183820152602001610c1e565b50505050905090810190601f168015610c635780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b158015610c8557600080fd5b505af1158015610c99573d6000803e3d6000fd5b506001979650505050505050565b6001600160a01b03918216600090815260096020908152604080832093909416825291909152205490565b600080546001600160a01b03163314610d1c5760405162461bcd60e51b8152600401808060200182810382526030815260200180610ecf6030913960400191505060405180910390fd5b6001600160a01b03831660009081526008602052604090205460ff1615610d745760405162461bcd60e51b8152600401808060200182810382526021815260200180610eff6021913960400191505060405180910390fd5b60045460055460ff909116600a0a9290920291610d919083610ebe565b60058190556006541015610da457600080fd5b6001600160a01b038316600090815260076020526040902054610dc79083610ebe565b6001600160a01b0380851660008181526007602090815260408083209590955590548451878152945192949316927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a350600192915050565b6000546001600160a01b03163314610e725760405162461bcd60e51b8152600401808060200182810382526030815260200180610ecf6030913960400191505060405180910390fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b60086020526000908152604090205460ff1681565b600082821115610eb857600080fd5b50900390565b818101828110156106aa57600080fdfe4f6e6c7920636f6e7472616374206f776e65722063616e20706572666f726d2074686973207472616e73616374696f6e54686973206163636f756e7420686173206265656e20626c61636b6c69737465644f6e6c79206e6577206f776e65722063616e20706572666f726d2074686973207472616e73616374696f6ea265627a7a72315820eb913982702d1aec6add31eff9d716a631076e23f7f74b100006a952ad1a532964736f6c637829302e352e31342d646576656c6f702e323031392e31312e31352b636f6d6d69742e336530346664366500595f73796d626f6c206d75737420626520332063686172616374657273206d696e696d756d5f6e616d65206d75737420626520342063686172616374657273206d696e696d756d5f6d6178546f74616c537570706c79206d7573742062652031303020756e697473206d696e696d756d`

// DeploySmartToken deploys a new Ethereum contract, binding an instance of SmartToken to it.
func DeploySmartToken(auth *bind.TransactOpts, backend bind.ContractBackend, _symbol string, _name string, _maxTotalSupply *big.Int, _seed *big.Int) (common.Address, *types.Transaction, *SmartToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SmartTokenBin), backend, _symbol, _name, _maxTotalSupply, _seed)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SmartToken{SmartTokenCaller: SmartTokenCaller{contract: contract}, SmartTokenTransactor: SmartTokenTransactor{contract: contract}, SmartTokenFilterer: SmartTokenFilterer{contract: contract}}, nil
}

// SmartToken is an auto generated Go binding around an Ethereum contract.
type SmartToken struct {
	SmartTokenCaller     // Read-only binding to the contract
	SmartTokenTransactor // Write-only binding to the contract
	SmartTokenFilterer   // Log filterer for contract events
}

// SmartTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartTokenSession struct {
	Contract     *SmartToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartTokenCallerSession struct {
	Contract *SmartTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SmartTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartTokenTransactorSession struct {
	Contract     *SmartTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SmartTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartTokenRaw struct {
	Contract *SmartToken // Generic contract binding to access the raw methods on
}

// SmartTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartTokenCallerRaw struct {
	Contract *SmartTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SmartTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartTokenTransactorRaw struct {
	Contract *SmartTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartToken creates a new instance of SmartToken, bound to a specific deployed contract.
func NewSmartToken(address common.Address, backend bind.ContractBackend) (*SmartToken, error) {
	contract, err := bindSmartToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SmartToken{SmartTokenCaller: SmartTokenCaller{contract: contract}, SmartTokenTransactor: SmartTokenTransactor{contract: contract}, SmartTokenFilterer: SmartTokenFilterer{contract: contract}}, nil
}

// NewSmartTokenCaller creates a new read-only instance of SmartToken, bound to a specific deployed contract.
func NewSmartTokenCaller(address common.Address, caller bind.ContractCaller) (*SmartTokenCaller, error) {
	contract, err := bindSmartToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartTokenCaller{contract: contract}, nil
}

// NewSmartTokenTransactor creates a new write-only instance of SmartToken, bound to a specific deployed contract.
func NewSmartTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartTokenTransactor, error) {
	contract, err := bindSmartToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartTokenTransactor{contract: contract}, nil
}

// NewSmartTokenFilterer creates a new log filterer instance of SmartToken, bound to a specific deployed contract.
func NewSmartTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartTokenFilterer, error) {
	contract, err := bindSmartToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartTokenFilterer{contract: contract}, nil
}

// bindSmartToken binds a generic wrapper to an already deployed contract.
func bindSmartToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartToken *SmartTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartToken.Contract.SmartTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartToken *SmartTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartToken.Contract.SmartTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartToken *SmartTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartToken.Contract.SmartTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartToken *SmartTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartToken *SmartTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartToken *SmartTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_SmartToken *SmartTokenCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartToken.contract.Call(opts, out, "allowance", tokenOwner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_SmartToken *SmartTokenSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _SmartToken.Contract.Allowance(&_SmartToken.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_SmartToken *SmartTokenCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _SmartToken.Contract.Allowance(&_SmartToken.CallOpts, tokenOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_SmartToken *SmartTokenCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartToken.contract.Call(opts, out, "balanceOf", tokenOwner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_SmartToken *SmartTokenSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _SmartToken.Contract.BalanceOf(&_SmartToken.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_SmartToken *SmartTokenCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _SmartToken.Contract.BalanceOf(&_SmartToken.CallOpts, tokenOwner)
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist( address) constant returns(bool)
func (_SmartToken *SmartTokenCaller) Blacklist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartToken.contract.Call(opts, out, "blacklist", arg0)
	return *ret0, err
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist( address) constant returns(bool)
func (_SmartToken *SmartTokenSession) Blacklist(arg0 common.Address) (bool, error) {
	return _SmartToken.Contract.Blacklist(&_SmartToken.CallOpts, arg0)
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist( address) constant returns(bool)
func (_SmartToken *SmartTokenCallerSession) Blacklist(arg0 common.Address) (bool, error) {
	return _SmartToken.Contract.Blacklist(&_SmartToken.CallOpts, arg0)
}

// CurTotalSupply is a free data retrieval call binding the contract method 0x09cc9854.
//
// Solidity: function curTotalSupply() constant returns(uint256)
func (_SmartToken *SmartTokenCaller) CurTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartToken.contract.Call(opts, out, "curTotalSupply")
	return *ret0, err
}

// CurTotalSupply is a free data retrieval call binding the contract method 0x09cc9854.
//
// Solidity: function curTotalSupply() constant returns(uint256)
func (_SmartToken *SmartTokenSession) CurTotalSupply() (*big.Int, error) {
	return _SmartToken.Contract.CurTotalSupply(&_SmartToken.CallOpts)
}

// CurTotalSupply is a free data retrieval call binding the contract method 0x09cc9854.
//
// Solidity: function curTotalSupply() constant returns(uint256)
func (_SmartToken *SmartTokenCallerSession) CurTotalSupply() (*big.Int, error) {
	return _SmartToken.Contract.CurTotalSupply(&_SmartToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SmartToken *SmartTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SmartToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SmartToken *SmartTokenSession) Decimals() (uint8, error) {
	return _SmartToken.Contract.Decimals(&_SmartToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SmartToken *SmartTokenCallerSession) Decimals() (uint8, error) {
	return _SmartToken.Contract.Decimals(&_SmartToken.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() constant returns(uint256)
func (_SmartToken *SmartTokenCaller) MaxTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartToken.contract.Call(opts, out, "maxTotalSupply")
	return *ret0, err
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() constant returns(uint256)
func (_SmartToken *SmartTokenSession) MaxTotalSupply() (*big.Int, error) {
	return _SmartToken.Contract.MaxTotalSupply(&_SmartToken.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() constant returns(uint256)
func (_SmartToken *SmartTokenCallerSession) MaxTotalSupply() (*big.Int, error) {
	return _SmartToken.Contract.MaxTotalSupply(&_SmartToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SmartToken *SmartTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SmartToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SmartToken *SmartTokenSession) Name() (string, error) {
	return _SmartToken.Contract.Name(&_SmartToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SmartToken *SmartTokenCallerSession) Name() (string, error) {
	return _SmartToken.Contract.Name(&_SmartToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SmartToken *SmartTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SmartToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SmartToken *SmartTokenSession) Symbol() (string, error) {
	return _SmartToken.Contract.Symbol(&_SmartToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SmartToken *SmartTokenCallerSession) Symbol() (string, error) {
	return _SmartToken.Contract.Symbol(&_SmartToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SmartToken *SmartTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SmartToken *SmartTokenSession) TotalSupply() (*big.Int, error) {
	return _SmartToken.Contract.TotalSupply(&_SmartToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SmartToken *SmartTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SmartToken.Contract.TotalSupply(&_SmartToken.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SmartToken *SmartTokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartToken.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SmartToken *SmartTokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _SmartToken.Contract.AcceptOwnership(&_SmartToken.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SmartToken *SmartTokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SmartToken.Contract.AcceptOwnership(&_SmartToken.TransactOpts)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x44337ea1.
//
// Solidity: function addToBlacklist(user address) returns(success bool)
func (_SmartToken *SmartTokenTransactor) AddToBlacklist(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _SmartToken.contract.Transact(opts, "addToBlacklist", user)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x44337ea1.
//
// Solidity: function addToBlacklist(user address) returns(success bool)
func (_SmartToken *SmartTokenSession) AddToBlacklist(user common.Address) (*types.Transaction, error) {
	return _SmartToken.Contract.AddToBlacklist(&_SmartToken.TransactOpts, user)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x44337ea1.
//
// Solidity: function addToBlacklist(user address) returns(success bool)
func (_SmartToken *SmartTokenTransactorSession) AddToBlacklist(user common.Address) (*types.Transaction, error) {
	return _SmartToken.Contract.AddToBlacklist(&_SmartToken.TransactOpts, user)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_SmartToken *SmartTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _SmartToken.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_SmartToken *SmartTokenSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _SmartToken.Contract.Approve(&_SmartToken.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_SmartToken *SmartTokenTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _SmartToken.Contract.Approve(&_SmartToken.TransactOpts, spender, tokens)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(spender address, tokens uint256, data bytes) returns(success bool)
func (_SmartToken *SmartTokenTransactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _SmartToken.contract.Transact(opts, "approveAndCall", spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(spender address, tokens uint256, data bytes) returns(success bool)
func (_SmartToken *SmartTokenSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _SmartToken.Contract.ApproveAndCall(&_SmartToken.TransactOpts, spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(spender address, tokens uint256, data bytes) returns(success bool)
func (_SmartToken *SmartTokenTransactorSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _SmartToken.Contract.ApproveAndCall(&_SmartToken.TransactOpts, spender, tokens, data)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf0dda65c.
//
// Solidity: function mintTokens(to address, tokens uint256) returns(success bool)
func (_SmartToken *SmartTokenTransactor) MintTokens(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _SmartToken.contract.Transact(opts, "mintTokens", to, tokens)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf0dda65c.
//
// Solidity: function mintTokens(to address, tokens uint256) returns(success bool)
func (_SmartToken *SmartTokenSession) MintTokens(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _SmartToken.Contract.MintTokens(&_SmartToken.TransactOpts, to, tokens)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf0dda65c.
//
// Solidity: function mintTokens(to address, tokens uint256) returns(success bool)
func (_SmartToken *SmartTokenTransactorSession) MintTokens(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _SmartToken.Contract.MintTokens(&_SmartToken.TransactOpts, to, tokens)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x537df3b6.
//
// Solidity: function removeFromBlacklist(user address) returns(success bool)
func (_SmartToken *SmartTokenTransactor) RemoveFromBlacklist(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _SmartToken.contract.Transact(opts, "removeFromBlacklist", user)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x537df3b6.
//
// Solidity: function removeFromBlacklist(user address) returns(success bool)
func (_SmartToken *SmartTokenSession) RemoveFromBlacklist(user common.Address) (*types.Transaction, error) {
	return _SmartToken.Contract.RemoveFromBlacklist(&_SmartToken.TransactOpts, user)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x537df3b6.
//
// Solidity: function removeFromBlacklist(user address) returns(success bool)
func (_SmartToken *SmartTokenTransactorSession) RemoveFromBlacklist(user common.Address) (*types.Transaction, error) {
	return _SmartToken.Contract.RemoveFromBlacklist(&_SmartToken.TransactOpts, user)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_SmartToken *SmartTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _SmartToken.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_SmartToken *SmartTokenSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _SmartToken.Contract.Transfer(&_SmartToken.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_SmartToken *SmartTokenTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _SmartToken.Contract.Transfer(&_SmartToken.TransactOpts, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_SmartToken *SmartTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _SmartToken.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_SmartToken *SmartTokenSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _SmartToken.Contract.TransferFrom(&_SmartToken.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_SmartToken *SmartTokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _SmartToken.Contract.TransferFrom(&_SmartToken.TransactOpts, from, to, tokens)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_SmartToken *SmartTokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _SmartToken.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_SmartToken *SmartTokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _SmartToken.Contract.TransferOwnership(&_SmartToken.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_SmartToken *SmartTokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _SmartToken.Contract.TransferOwnership(&_SmartToken.TransactOpts, _newOwner)
}

// SmartTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SmartToken contract.
type SmartTokenApprovalIterator struct {
	Event *SmartTokenApproval // Event containing the contract specifics and raw log

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
func (it *SmartTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartTokenApproval)
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
		it.Event = new(SmartTokenApproval)
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
func (it *SmartTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartTokenApproval represents a Approval event raised by the SmartToken contract.
type SmartTokenApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(tokenOwner indexed address, spender indexed address, tokens uint256)
func (_SmartToken *SmartTokenFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*SmartTokenApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SmartToken.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SmartTokenApprovalIterator{contract: _SmartToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(tokenOwner indexed address, spender indexed address, tokens uint256)
func (_SmartToken *SmartTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SmartTokenApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SmartToken.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartTokenApproval)
				if err := _SmartToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// SmartTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SmartToken contract.
type SmartTokenOwnershipTransferredIterator struct {
	Event *SmartTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SmartTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartTokenOwnershipTransferred)
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
		it.Event = new(SmartTokenOwnershipTransferred)
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
func (it *SmartTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartTokenOwnershipTransferred represents a OwnershipTransferred event raised by the SmartToken contract.
type SmartTokenOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_from indexed address, _to indexed address)
func (_SmartToken *SmartTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*SmartTokenOwnershipTransferredIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _SmartToken.contract.FilterLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &SmartTokenOwnershipTransferredIterator{contract: _SmartToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_from indexed address, _to indexed address)
func (_SmartToken *SmartTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SmartTokenOwnershipTransferred, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _SmartToken.contract.WatchLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartTokenOwnershipTransferred)
				if err := _SmartToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SmartTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SmartToken contract.
type SmartTokenTransferIterator struct {
	Event *SmartTokenTransfer // Event containing the contract specifics and raw log

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
func (it *SmartTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartTokenTransfer)
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
		it.Event = new(SmartTokenTransfer)
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
func (it *SmartTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartTokenTransfer represents a Transfer event raised by the SmartToken contract.
type SmartTokenTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, tokens uint256)
func (_SmartToken *SmartTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SmartTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SmartToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SmartTokenTransferIterator{contract: _SmartToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, tokens uint256)
func (_SmartToken *SmartTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SmartTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SmartToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartTokenTransfer)
				if err := _SmartToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ApproveAndCallFallBackABI is the input ABI used to generate the binding from.
const ApproveAndCallFallBackABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"receiveApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ApproveAndCallFallBackBin is the compiled bytecode used for deploying new contracts.
const ApproveAndCallFallBackBin = `0x`

// DeployApproveAndCallFallBack deploys a new Ethereum contract, binding an instance of ApproveAndCallFallBack to it.
func DeployApproveAndCallFallBack(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ApproveAndCallFallBack, error) {
	parsed, err := abi.JSON(strings.NewReader(ApproveAndCallFallBackABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ApproveAndCallFallBackBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ApproveAndCallFallBack{ApproveAndCallFallBackCaller: ApproveAndCallFallBackCaller{contract: contract}, ApproveAndCallFallBackTransactor: ApproveAndCallFallBackTransactor{contract: contract}, ApproveAndCallFallBackFilterer: ApproveAndCallFallBackFilterer{contract: contract}}, nil
}

// ApproveAndCallFallBack is an auto generated Go binding around an Ethereum contract.
type ApproveAndCallFallBack struct {
	ApproveAndCallFallBackCaller     // Read-only binding to the contract
	ApproveAndCallFallBackTransactor // Write-only binding to the contract
	ApproveAndCallFallBackFilterer   // Log filterer for contract events
}

// ApproveAndCallFallBackCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApproveAndCallFallBackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApproveAndCallFallBackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApproveAndCallFallBackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApproveAndCallFallBackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApproveAndCallFallBackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApproveAndCallFallBackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApproveAndCallFallBackSession struct {
	Contract     *ApproveAndCallFallBack // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ApproveAndCallFallBackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApproveAndCallFallBackCallerSession struct {
	Contract *ApproveAndCallFallBackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ApproveAndCallFallBackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApproveAndCallFallBackTransactorSession struct {
	Contract     *ApproveAndCallFallBackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ApproveAndCallFallBackRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApproveAndCallFallBackRaw struct {
	Contract *ApproveAndCallFallBack // Generic contract binding to access the raw methods on
}

// ApproveAndCallFallBackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApproveAndCallFallBackCallerRaw struct {
	Contract *ApproveAndCallFallBackCaller // Generic read-only contract binding to access the raw methods on
}

// ApproveAndCallFallBackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApproveAndCallFallBackTransactorRaw struct {
	Contract *ApproveAndCallFallBackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApproveAndCallFallBack creates a new instance of ApproveAndCallFallBack, bound to a specific deployed contract.
func NewApproveAndCallFallBack(address common.Address, backend bind.ContractBackend) (*ApproveAndCallFallBack, error) {
	contract, err := bindApproveAndCallFallBack(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ApproveAndCallFallBack{ApproveAndCallFallBackCaller: ApproveAndCallFallBackCaller{contract: contract}, ApproveAndCallFallBackTransactor: ApproveAndCallFallBackTransactor{contract: contract}, ApproveAndCallFallBackFilterer: ApproveAndCallFallBackFilterer{contract: contract}}, nil
}

// NewApproveAndCallFallBackCaller creates a new read-only instance of ApproveAndCallFallBack, bound to a specific deployed contract.
func NewApproveAndCallFallBackCaller(address common.Address, caller bind.ContractCaller) (*ApproveAndCallFallBackCaller, error) {
	contract, err := bindApproveAndCallFallBack(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApproveAndCallFallBackCaller{contract: contract}, nil
}

// NewApproveAndCallFallBackTransactor creates a new write-only instance of ApproveAndCallFallBack, bound to a specific deployed contract.
func NewApproveAndCallFallBackTransactor(address common.Address, transactor bind.ContractTransactor) (*ApproveAndCallFallBackTransactor, error) {
	contract, err := bindApproveAndCallFallBack(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApproveAndCallFallBackTransactor{contract: contract}, nil
}

// NewApproveAndCallFallBackFilterer creates a new log filterer instance of ApproveAndCallFallBack, bound to a specific deployed contract.
func NewApproveAndCallFallBackFilterer(address common.Address, filterer bind.ContractFilterer) (*ApproveAndCallFallBackFilterer, error) {
	contract, err := bindApproveAndCallFallBack(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApproveAndCallFallBackFilterer{contract: contract}, nil
}

// bindApproveAndCallFallBack binds a generic wrapper to an already deployed contract.
func bindApproveAndCallFallBack(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApproveAndCallFallBackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApproveAndCallFallBack *ApproveAndCallFallBackRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ApproveAndCallFallBack.Contract.ApproveAndCallFallBackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApproveAndCallFallBack *ApproveAndCallFallBackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApproveAndCallFallBack.Contract.ApproveAndCallFallBackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApproveAndCallFallBack *ApproveAndCallFallBackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApproveAndCallFallBack.Contract.ApproveAndCallFallBackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApproveAndCallFallBack *ApproveAndCallFallBackCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ApproveAndCallFallBack.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApproveAndCallFallBack *ApproveAndCallFallBackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApproveAndCallFallBack.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApproveAndCallFallBack *ApproveAndCallFallBackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApproveAndCallFallBack.Contract.contract.Transact(opts, method, params...)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(from address, tokens uint256, token address, data bytes) returns()
func (_ApproveAndCallFallBack *ApproveAndCallFallBackTransactor) ReceiveApproval(opts *bind.TransactOpts, from common.Address, tokens *big.Int, token common.Address, data []byte) (*types.Transaction, error) {
	return _ApproveAndCallFallBack.contract.Transact(opts, "receiveApproval", from, tokens, token, data)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(from address, tokens uint256, token address, data bytes) returns()
func (_ApproveAndCallFallBack *ApproveAndCallFallBackSession) ReceiveApproval(from common.Address, tokens *big.Int, token common.Address, data []byte) (*types.Transaction, error) {
	return _ApproveAndCallFallBack.Contract.ReceiveApproval(&_ApproveAndCallFallBack.TransactOpts, from, tokens, token, data)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(from address, tokens uint256, token address, data bytes) returns()
func (_ApproveAndCallFallBack *ApproveAndCallFallBackTransactorSession) ReceiveApproval(from common.Address, tokens *big.Int, token common.Address, data []byte) (*types.Transaction, error) {
	return _ApproveAndCallFallBack.Contract.ReceiveApproval(&_ApproveAndCallFallBack.TransactOpts, from, tokens, token, data)
}
