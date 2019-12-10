// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plasma

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

// EthvaultABI is the input ABI used to generate the binding from.
const EthvaultABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositVerifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"name\":\"setDepositVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEffectiveDepositVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newDepositVerifierMaturityTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"safeGasStipend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPlasmaFramework\",\"name\":\"_framework\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_safeGasStipend\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blknum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nextDepositVerifier\",\"type\":\"address\"}],\"name\":\"SetDepositVerifierCalled\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_depositTx\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Ethvault is an auto generated Go binding around an Ethereum contract.
type Ethvault struct {
	EthvaultCaller     // Read-only binding to the contract
	EthvaultTransactor // Write-only binding to the contract
	EthvaultFilterer   // Log filterer for contract events
}

// EthvaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthvaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthvaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthvaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthvaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthvaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthvaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthvaultSession struct {
	Contract     *Ethvault         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthvaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthvaultCallerSession struct {
	Contract *EthvaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EthvaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthvaultTransactorSession struct {
	Contract     *EthvaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthvaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthvaultRaw struct {
	Contract *Ethvault // Generic contract binding to access the raw methods on
}

// EthvaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthvaultCallerRaw struct {
	Contract *EthvaultCaller // Generic read-only contract binding to access the raw methods on
}

// EthvaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthvaultTransactorRaw struct {
	Contract *EthvaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthvault creates a new instance of Ethvault, bound to a specific deployed contract.
func NewEthvault(address common.Address, backend bind.ContractBackend) (*Ethvault, error) {
	contract, err := bindEthvault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethvault{EthvaultCaller: EthvaultCaller{contract: contract}, EthvaultTransactor: EthvaultTransactor{contract: contract}, EthvaultFilterer: EthvaultFilterer{contract: contract}}, nil
}

// NewEthvaultCaller creates a new read-only instance of Ethvault, bound to a specific deployed contract.
func NewEthvaultCaller(address common.Address, caller bind.ContractCaller) (*EthvaultCaller, error) {
	contract, err := bindEthvault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthvaultCaller{contract: contract}, nil
}

// NewEthvaultTransactor creates a new write-only instance of Ethvault, bound to a specific deployed contract.
func NewEthvaultTransactor(address common.Address, transactor bind.ContractTransactor) (*EthvaultTransactor, error) {
	contract, err := bindEthvault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthvaultTransactor{contract: contract}, nil
}

// NewEthvaultFilterer creates a new log filterer instance of Ethvault, bound to a specific deployed contract.
func NewEthvaultFilterer(address common.Address, filterer bind.ContractFilterer) (*EthvaultFilterer, error) {
	contract, err := bindEthvault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthvaultFilterer{contract: contract}, nil
}

// bindEthvault binds a generic wrapper to an already deployed contract.
func bindEthvault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthvaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethvault *EthvaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethvault.Contract.EthvaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethvault *EthvaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethvault.Contract.EthvaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethvault *EthvaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethvault.Contract.EthvaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethvault *EthvaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethvault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethvault *EthvaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethvault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethvault *EthvaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethvault.Contract.contract.Transact(opts, method, params...)
}

// DepositVerifiers is a free data retrieval call binding the contract method 0x02e9b07b.
//
// Solidity: function depositVerifiers(uint256 ) constant returns(address)
func (_Ethvault *EthvaultCaller) DepositVerifiers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ethvault.contract.Call(opts, out, "depositVerifiers", arg0)
	return *ret0, err
}

// DepositVerifiers is a free data retrieval call binding the contract method 0x02e9b07b.
//
// Solidity: function depositVerifiers(uint256 ) constant returns(address)
func (_Ethvault *EthvaultSession) DepositVerifiers(arg0 *big.Int) (common.Address, error) {
	return _Ethvault.Contract.DepositVerifiers(&_Ethvault.CallOpts, arg0)
}

// DepositVerifiers is a free data retrieval call binding the contract method 0x02e9b07b.
//
// Solidity: function depositVerifiers(uint256 ) constant returns(address)
func (_Ethvault *EthvaultCallerSession) DepositVerifiers(arg0 *big.Int) (common.Address, error) {
	return _Ethvault.Contract.DepositVerifiers(&_Ethvault.CallOpts, arg0)
}

// GetEffectiveDepositVerifier is a free data retrieval call binding the contract method 0x52518b7f.
//
// Solidity: function getEffectiveDepositVerifier() constant returns(address)
func (_Ethvault *EthvaultCaller) GetEffectiveDepositVerifier(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ethvault.contract.Call(opts, out, "getEffectiveDepositVerifier")
	return *ret0, err
}

// GetEffectiveDepositVerifier is a free data retrieval call binding the contract method 0x52518b7f.
//
// Solidity: function getEffectiveDepositVerifier() constant returns(address)
func (_Ethvault *EthvaultSession) GetEffectiveDepositVerifier() (common.Address, error) {
	return _Ethvault.Contract.GetEffectiveDepositVerifier(&_Ethvault.CallOpts)
}

// GetEffectiveDepositVerifier is a free data retrieval call binding the contract method 0x52518b7f.
//
// Solidity: function getEffectiveDepositVerifier() constant returns(address)
func (_Ethvault *EthvaultCallerSession) GetEffectiveDepositVerifier() (common.Address, error) {
	return _Ethvault.Contract.GetEffectiveDepositVerifier(&_Ethvault.CallOpts)
}

// NewDepositVerifierMaturityTimestamp is a free data retrieval call binding the contract method 0xa7e5994f.
//
// Solidity: function newDepositVerifierMaturityTimestamp() constant returns(uint256)
func (_Ethvault *EthvaultCaller) NewDepositVerifierMaturityTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethvault.contract.Call(opts, out, "newDepositVerifierMaturityTimestamp")
	return *ret0, err
}

// NewDepositVerifierMaturityTimestamp is a free data retrieval call binding the contract method 0xa7e5994f.
//
// Solidity: function newDepositVerifierMaturityTimestamp() constant returns(uint256)
func (_Ethvault *EthvaultSession) NewDepositVerifierMaturityTimestamp() (*big.Int, error) {
	return _Ethvault.Contract.NewDepositVerifierMaturityTimestamp(&_Ethvault.CallOpts)
}

// NewDepositVerifierMaturityTimestamp is a free data retrieval call binding the contract method 0xa7e5994f.
//
// Solidity: function newDepositVerifierMaturityTimestamp() constant returns(uint256)
func (_Ethvault *EthvaultCallerSession) NewDepositVerifierMaturityTimestamp() (*big.Int, error) {
	return _Ethvault.Contract.NewDepositVerifierMaturityTimestamp(&_Ethvault.CallOpts)
}

// SafeGasStipend is a free data retrieval call binding the contract method 0xccbd2176.
//
// Solidity: function safeGasStipend() constant returns(uint256)
func (_Ethvault *EthvaultCaller) SafeGasStipend(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethvault.contract.Call(opts, out, "safeGasStipend")
	return *ret0, err
}

// SafeGasStipend is a free data retrieval call binding the contract method 0xccbd2176.
//
// Solidity: function safeGasStipend() constant returns(uint256)
func (_Ethvault *EthvaultSession) SafeGasStipend() (*big.Int, error) {
	return _Ethvault.Contract.SafeGasStipend(&_Ethvault.CallOpts)
}

// SafeGasStipend is a free data retrieval call binding the contract method 0xccbd2176.
//
// Solidity: function safeGasStipend() constant returns(uint256)
func (_Ethvault *EthvaultCallerSession) SafeGasStipend() (*big.Int, error) {
	return _Ethvault.Contract.SafeGasStipend(&_Ethvault.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(bytes _depositTx) returns()
func (_Ethvault *EthvaultTransactor) Deposit(opts *bind.TransactOpts, _depositTx []byte) (*types.Transaction, error) {
	return _Ethvault.contract.Transact(opts, "deposit", _depositTx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(bytes _depositTx) returns()
func (_Ethvault *EthvaultSession) Deposit(_depositTx []byte) (*types.Transaction, error) {
	return _Ethvault.Contract.Deposit(&_Ethvault.TransactOpts, _depositTx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(bytes _depositTx) returns()
func (_Ethvault *EthvaultTransactorSession) Deposit(_depositTx []byte) (*types.Transaction, error) {
	return _Ethvault.Contract.Deposit(&_Ethvault.TransactOpts, _depositTx)
}

// SetDepositVerifier is a paid mutator transaction binding the contract method 0x2c6ce78b.
//
// Solidity: function setDepositVerifier(address _verifier) returns()
func (_Ethvault *EthvaultTransactor) SetDepositVerifier(opts *bind.TransactOpts, _verifier common.Address) (*types.Transaction, error) {
	return _Ethvault.contract.Transact(opts, "setDepositVerifier", _verifier)
}

// SetDepositVerifier is a paid mutator transaction binding the contract method 0x2c6ce78b.
//
// Solidity: function setDepositVerifier(address _verifier) returns()
func (_Ethvault *EthvaultSession) SetDepositVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _Ethvault.Contract.SetDepositVerifier(&_Ethvault.TransactOpts, _verifier)
}

// SetDepositVerifier is a paid mutator transaction binding the contract method 0x2c6ce78b.
//
// Solidity: function setDepositVerifier(address _verifier) returns()
func (_Ethvault *EthvaultTransactorSession) SetDepositVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _Ethvault.Contract.SetDepositVerifier(&_Ethvault.TransactOpts, _verifier)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address receiver, uint256 amount) returns()
func (_Ethvault *EthvaultTransactor) Withdraw(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ethvault.contract.Transact(opts, "withdraw", receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address receiver, uint256 amount) returns()
func (_Ethvault *EthvaultSession) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ethvault.Contract.Withdraw(&_Ethvault.TransactOpts, receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address receiver, uint256 amount) returns()
func (_Ethvault *EthvaultTransactorSession) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ethvault.Contract.Withdraw(&_Ethvault.TransactOpts, receiver, amount)
}

// EthvaultDepositCreatedIterator is returned from FilterDepositCreated and is used to iterate over the raw logs and unpacked data for DepositCreated events raised by the Ethvault contract.
type EthvaultDepositCreatedIterator struct {
	Event *EthvaultDepositCreated // Event containing the contract specifics and raw log

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
func (it *EthvaultDepositCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthvaultDepositCreated)
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
		it.Event = new(EthvaultDepositCreated)
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
func (it *EthvaultDepositCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthvaultDepositCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthvaultDepositCreated represents a DepositCreated event raised by the Ethvault contract.
type EthvaultDepositCreated struct {
	Depositor common.Address
	Blknum    *big.Int
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositCreated is a free log retrieval operation binding the contract event 0x18569122d84f30025bb8dffb33563f1bdbfb9637f21552b11b8305686e9cb307.
//
// Solidity: event DepositCreated(address indexed depositor, uint256 indexed blknum, address indexed token, uint256 amount)
func (_Ethvault *EthvaultFilterer) FilterDepositCreated(opts *bind.FilterOpts, depositor []common.Address, blknum []*big.Int, token []common.Address) (*EthvaultDepositCreatedIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var blknumRule []interface{}
	for _, blknumItem := range blknum {
		blknumRule = append(blknumRule, blknumItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Ethvault.contract.FilterLogs(opts, "DepositCreated", depositorRule, blknumRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &EthvaultDepositCreatedIterator{contract: _Ethvault.contract, event: "DepositCreated", logs: logs, sub: sub}, nil
}

// WatchDepositCreated is a free log subscription operation binding the contract event 0x18569122d84f30025bb8dffb33563f1bdbfb9637f21552b11b8305686e9cb307.
//
// Solidity: event DepositCreated(address indexed depositor, uint256 indexed blknum, address indexed token, uint256 amount)
func (_Ethvault *EthvaultFilterer) WatchDepositCreated(opts *bind.WatchOpts, sink chan<- *EthvaultDepositCreated, depositor []common.Address, blknum []*big.Int, token []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var blknumRule []interface{}
	for _, blknumItem := range blknum {
		blknumRule = append(blknumRule, blknumItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Ethvault.contract.WatchLogs(opts, "DepositCreated", depositorRule, blknumRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthvaultDepositCreated)
				if err := _Ethvault.contract.UnpackLog(event, "DepositCreated", log); err != nil {
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

// ParseDepositCreated is a log parse operation binding the contract event 0x18569122d84f30025bb8dffb33563f1bdbfb9637f21552b11b8305686e9cb307.
//
// Solidity: event DepositCreated(address indexed depositor, uint256 indexed blknum, address indexed token, uint256 amount)
func (_Ethvault *EthvaultFilterer) ParseDepositCreated(log types.Log) (*EthvaultDepositCreated, error) {
	event := new(EthvaultDepositCreated)
	if err := _Ethvault.contract.UnpackLog(event, "DepositCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthvaultEthWithdrawnIterator is returned from FilterEthWithdrawn and is used to iterate over the raw logs and unpacked data for EthWithdrawn events raised by the Ethvault contract.
type EthvaultEthWithdrawnIterator struct {
	Event *EthvaultEthWithdrawn // Event containing the contract specifics and raw log

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
func (it *EthvaultEthWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthvaultEthWithdrawn)
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
		it.Event = new(EthvaultEthWithdrawn)
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
func (it *EthvaultEthWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthvaultEthWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthvaultEthWithdrawn represents a EthWithdrawn event raised by the Ethvault contract.
type EthvaultEthWithdrawn struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(address indexed receiver, uint256 amount)
func (_Ethvault *EthvaultFilterer) FilterEthWithdrawn(opts *bind.FilterOpts, receiver []common.Address) (*EthvaultEthWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Ethvault.contract.FilterLogs(opts, "EthWithdrawn", receiverRule)
	if err != nil {
		return nil, err
	}
	return &EthvaultEthWithdrawnIterator{contract: _Ethvault.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(address indexed receiver, uint256 amount)
func (_Ethvault *EthvaultFilterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *EthvaultEthWithdrawn, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Ethvault.contract.WatchLogs(opts, "EthWithdrawn", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthvaultEthWithdrawn)
				if err := _Ethvault.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
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

// ParseEthWithdrawn is a log parse operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(address indexed receiver, uint256 amount)
func (_Ethvault *EthvaultFilterer) ParseEthWithdrawn(log types.Log) (*EthvaultEthWithdrawn, error) {
	event := new(EthvaultEthWithdrawn)
	if err := _Ethvault.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthvaultSetDepositVerifierCalledIterator is returned from FilterSetDepositVerifierCalled and is used to iterate over the raw logs and unpacked data for SetDepositVerifierCalled events raised by the Ethvault contract.
type EthvaultSetDepositVerifierCalledIterator struct {
	Event *EthvaultSetDepositVerifierCalled // Event containing the contract specifics and raw log

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
func (it *EthvaultSetDepositVerifierCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthvaultSetDepositVerifierCalled)
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
		it.Event = new(EthvaultSetDepositVerifierCalled)
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
func (it *EthvaultSetDepositVerifierCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthvaultSetDepositVerifierCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthvaultSetDepositVerifierCalled represents a SetDepositVerifierCalled event raised by the Ethvault contract.
type EthvaultSetDepositVerifierCalled struct {
	NextDepositVerifier common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetDepositVerifierCalled is a free log retrieval operation binding the contract event 0x11e5a7570e0de4b31179f8c9cf5e8a1e6cfd23731ec7a5a06e3006d71dd321fd.
//
// Solidity: event SetDepositVerifierCalled(address nextDepositVerifier)
func (_Ethvault *EthvaultFilterer) FilterSetDepositVerifierCalled(opts *bind.FilterOpts) (*EthvaultSetDepositVerifierCalledIterator, error) {

	logs, sub, err := _Ethvault.contract.FilterLogs(opts, "SetDepositVerifierCalled")
	if err != nil {
		return nil, err
	}
	return &EthvaultSetDepositVerifierCalledIterator{contract: _Ethvault.contract, event: "SetDepositVerifierCalled", logs: logs, sub: sub}, nil
}

// WatchSetDepositVerifierCalled is a free log subscription operation binding the contract event 0x11e5a7570e0de4b31179f8c9cf5e8a1e6cfd23731ec7a5a06e3006d71dd321fd.
//
// Solidity: event SetDepositVerifierCalled(address nextDepositVerifier)
func (_Ethvault *EthvaultFilterer) WatchSetDepositVerifierCalled(opts *bind.WatchOpts, sink chan<- *EthvaultSetDepositVerifierCalled) (event.Subscription, error) {

	logs, sub, err := _Ethvault.contract.WatchLogs(opts, "SetDepositVerifierCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthvaultSetDepositVerifierCalled)
				if err := _Ethvault.contract.UnpackLog(event, "SetDepositVerifierCalled", log); err != nil {
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

// ParseSetDepositVerifierCalled is a log parse operation binding the contract event 0x11e5a7570e0de4b31179f8c9cf5e8a1e6cfd23731ec7a5a06e3006d71dd321fd.
//
// Solidity: event SetDepositVerifierCalled(address nextDepositVerifier)
func (_Ethvault *EthvaultFilterer) ParseSetDepositVerifierCalled(log types.Log) (*EthvaultSetDepositVerifierCalled, error) {
	event := new(EthvaultSetDepositVerifierCalled)
	if err := _Ethvault.contract.UnpackLog(event, "SetDepositVerifierCalled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthvaultWithdrawFailedIterator is returned from FilterWithdrawFailed and is used to iterate over the raw logs and unpacked data for WithdrawFailed events raised by the Ethvault contract.
type EthvaultWithdrawFailedIterator struct {
	Event *EthvaultWithdrawFailed // Event containing the contract specifics and raw log

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
func (it *EthvaultWithdrawFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthvaultWithdrawFailed)
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
		it.Event = new(EthvaultWithdrawFailed)
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
func (it *EthvaultWithdrawFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthvaultWithdrawFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthvaultWithdrawFailed represents a WithdrawFailed event raised by the Ethvault contract.
type EthvaultWithdrawFailed struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFailed is a free log retrieval operation binding the contract event 0xa2269912b47133fae1d7f448c9284ea248951ac29b8c7c41d301f8721a38d10d.
//
// Solidity: event WithdrawFailed(address indexed receiver, uint256 amount)
func (_Ethvault *EthvaultFilterer) FilterWithdrawFailed(opts *bind.FilterOpts, receiver []common.Address) (*EthvaultWithdrawFailedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Ethvault.contract.FilterLogs(opts, "WithdrawFailed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &EthvaultWithdrawFailedIterator{contract: _Ethvault.contract, event: "WithdrawFailed", logs: logs, sub: sub}, nil
}

// WatchWithdrawFailed is a free log subscription operation binding the contract event 0xa2269912b47133fae1d7f448c9284ea248951ac29b8c7c41d301f8721a38d10d.
//
// Solidity: event WithdrawFailed(address indexed receiver, uint256 amount)
func (_Ethvault *EthvaultFilterer) WatchWithdrawFailed(opts *bind.WatchOpts, sink chan<- *EthvaultWithdrawFailed, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Ethvault.contract.WatchLogs(opts, "WithdrawFailed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthvaultWithdrawFailed)
				if err := _Ethvault.contract.UnpackLog(event, "WithdrawFailed", log); err != nil {
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

// ParseWithdrawFailed is a log parse operation binding the contract event 0xa2269912b47133fae1d7f448c9284ea248951ac29b8c7c41d301f8721a38d10d.
//
// Solidity: event WithdrawFailed(address indexed receiver, uint256 amount)
func (_Ethvault *EthvaultFilterer) ParseWithdrawFailed(log types.Log) (*EthvaultWithdrawFailed, error) {
	event := new(EthvaultWithdrawFailed)
	if err := _Ethvault.contract.UnpackLog(event, "WithdrawFailed", log); err != nil {
		return nil, err
	}
	return event, nil
}

