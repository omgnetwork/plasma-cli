// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rootchain

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ERC20VaultABI is the input ABI used to generate the binding from.
const ERC20VaultABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositVerifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"name\":\"setDepositVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEffectiveDepositVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newDepositVerifierMaturityTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPlasmaFramework\",\"name\":\"_framework\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Erc20Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blknum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nextDepositVerifier\",\"type\":\"address\"}],\"name\":\"SetDepositVerifierCalled\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"depositTx\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20Vault is an auto generated Go binding around an Ethereum contract.
type ERC20Vault struct {
	ERC20VaultCaller     // Read-only binding to the contract
	ERC20VaultTransactor // Write-only binding to the contract
}

// ERC20VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20VaultSession struct {
	Contract     *ERC20Vault       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20VaultCallerSession struct {
	Contract *ERC20VaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20VaultTransactorSession struct {
	Contract     *ERC20VaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20VaultRaw struct {
	Contract *ERC20Vault // Generic contract binding to access the raw methods on
}

// ERC20VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20VaultCallerRaw struct {
	Contract *ERC20VaultCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20VaultTransactorRaw struct {
	Contract *ERC20VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Vault creates a new instance of ERC20Vault, bound to a specific deployed contract.
func NewERC20Vault(address common.Address, backend bind.ContractBackend) (*ERC20Vault, error) {
	contract, err := bindERC20Vault(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Vault{ERC20VaultCaller: ERC20VaultCaller{contract: contract}, ERC20VaultTransactor: ERC20VaultTransactor{contract: contract}}, nil
}

// NewERC20VaultCaller creates a new read-only instance of ERC20Vault, bound to a specific deployed contract.
func NewERC20VaultCaller(address common.Address, caller bind.ContractCaller) (*ERC20VaultCaller, error) {
	contract, err := bindERC20Vault(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20VaultCaller{contract: contract}, nil
}

// NewERC20VaultTransactor creates a new write-only instance of ERC20Vault, bound to a specific deployed contract.
func NewERC20VaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20VaultTransactor, error) {
	contract, err := bindERC20Vault(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC20VaultTransactor{contract: contract}, nil
}

// bindERC20Vault binds a generic wrapper to an already deployed contract.
func bindERC20Vault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Vault *ERC20VaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Vault.Contract.ERC20VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Vault *ERC20VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Vault.Contract.ERC20VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Vault *ERC20VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Vault.Contract.ERC20VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Vault *ERC20VaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Vault *ERC20VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Vault *ERC20VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Vault.Contract.contract.Transact(opts, method, params...)
}

// DepositVerifiers is a free data retrieval call binding the contract method 0x02e9b07b.
//
// Solidity: function depositVerifiers( uint256) constant returns(address)
func (_ERC20Vault *ERC20VaultCaller) DepositVerifiers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Vault.contract.Call(opts, out, "depositVerifiers", arg0)
	return *ret0, err
}

// DepositVerifiers is a free data retrieval call binding the contract method 0x02e9b07b.
//
// Solidity: function depositVerifiers( uint256) constant returns(address)
func (_ERC20Vault *ERC20VaultSession) DepositVerifiers(arg0 *big.Int) (common.Address, error) {
	return _ERC20Vault.Contract.DepositVerifiers(&_ERC20Vault.CallOpts, arg0)
}

// DepositVerifiers is a free data retrieval call binding the contract method 0x02e9b07b.
//
// Solidity: function depositVerifiers( uint256) constant returns(address)
func (_ERC20Vault *ERC20VaultCallerSession) DepositVerifiers(arg0 *big.Int) (common.Address, error) {
	return _ERC20Vault.Contract.DepositVerifiers(&_ERC20Vault.CallOpts, arg0)
}

// GetEffectiveDepositVerifier is a free data retrieval call binding the contract method 0x52518b7f.
//
// Solidity: function getEffectiveDepositVerifier() constant returns(address)
func (_ERC20Vault *ERC20VaultCaller) GetEffectiveDepositVerifier(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Vault.contract.Call(opts, out, "getEffectiveDepositVerifier")
	return *ret0, err
}

// GetEffectiveDepositVerifier is a free data retrieval call binding the contract method 0x52518b7f.
//
// Solidity: function getEffectiveDepositVerifier() constant returns(address)
func (_ERC20Vault *ERC20VaultSession) GetEffectiveDepositVerifier() (common.Address, error) {
	return _ERC20Vault.Contract.GetEffectiveDepositVerifier(&_ERC20Vault.CallOpts)
}

// GetEffectiveDepositVerifier is a free data retrieval call binding the contract method 0x52518b7f.
//
// Solidity: function getEffectiveDepositVerifier() constant returns(address)
func (_ERC20Vault *ERC20VaultCallerSession) GetEffectiveDepositVerifier() (common.Address, error) {
	return _ERC20Vault.Contract.GetEffectiveDepositVerifier(&_ERC20Vault.CallOpts)
}

// NewDepositVerifierMaturityTimestamp is a free data retrieval call binding the contract method 0xa7e5994f.
//
// Solidity: function newDepositVerifierMaturityTimestamp() constant returns(uint256)
func (_ERC20Vault *ERC20VaultCaller) NewDepositVerifierMaturityTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Vault.contract.Call(opts, out, "newDepositVerifierMaturityTimestamp")
	return *ret0, err
}

// NewDepositVerifierMaturityTimestamp is a free data retrieval call binding the contract method 0xa7e5994f.
//
// Solidity: function newDepositVerifierMaturityTimestamp() constant returns(uint256)
func (_ERC20Vault *ERC20VaultSession) NewDepositVerifierMaturityTimestamp() (*big.Int, error) {
	return _ERC20Vault.Contract.NewDepositVerifierMaturityTimestamp(&_ERC20Vault.CallOpts)
}

// NewDepositVerifierMaturityTimestamp is a free data retrieval call binding the contract method 0xa7e5994f.
//
// Solidity: function newDepositVerifierMaturityTimestamp() constant returns(uint256)
func (_ERC20Vault *ERC20VaultCallerSession) NewDepositVerifierMaturityTimestamp() (*big.Int, error) {
	return _ERC20Vault.Contract.NewDepositVerifierMaturityTimestamp(&_ERC20Vault.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(depositTx bytes) returns()
func (_ERC20Vault *ERC20VaultTransactor) Deposit(opts *bind.TransactOpts, depositTx []byte) (*types.Transaction, error) {
	return _ERC20Vault.contract.Transact(opts, "deposit", depositTx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(depositTx bytes) returns()
func (_ERC20Vault *ERC20VaultSession) Deposit(depositTx []byte) (*types.Transaction, error) {
	return _ERC20Vault.Contract.Deposit(&_ERC20Vault.TransactOpts, depositTx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(depositTx bytes) returns()
func (_ERC20Vault *ERC20VaultTransactorSession) Deposit(depositTx []byte) (*types.Transaction, error) {
	return _ERC20Vault.Contract.Deposit(&_ERC20Vault.TransactOpts, depositTx)
}

// SetDepositVerifier is a paid mutator transaction binding the contract method 0x2c6ce78b.
//
// Solidity: function setDepositVerifier(_verifier address) returns()
func (_ERC20Vault *ERC20VaultTransactor) SetDepositVerifier(opts *bind.TransactOpts, _verifier common.Address) (*types.Transaction, error) {
	return _ERC20Vault.contract.Transact(opts, "setDepositVerifier", _verifier)
}

// SetDepositVerifier is a paid mutator transaction binding the contract method 0x2c6ce78b.
//
// Solidity: function setDepositVerifier(_verifier address) returns()
func (_ERC20Vault *ERC20VaultSession) SetDepositVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _ERC20Vault.Contract.SetDepositVerifier(&_ERC20Vault.TransactOpts, _verifier)
}

// SetDepositVerifier is a paid mutator transaction binding the contract method 0x2c6ce78b.
//
// Solidity: function setDepositVerifier(_verifier address) returns()
func (_ERC20Vault *ERC20VaultTransactorSession) SetDepositVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _ERC20Vault.Contract.SetDepositVerifier(&_ERC20Vault.TransactOpts, _verifier)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(receiver address, token address, amount uint256) returns()
func (_ERC20Vault *ERC20VaultTransactor) Withdraw(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Vault.contract.Transact(opts, "withdraw", receiver, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(receiver address, token address, amount uint256) returns()
func (_ERC20Vault *ERC20VaultSession) Withdraw(receiver common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Vault.Contract.Withdraw(&_ERC20Vault.TransactOpts, receiver, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(receiver address, token address, amount uint256) returns()
func (_ERC20Vault *ERC20VaultTransactorSession) Withdraw(receiver common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Vault.Contract.Withdraw(&_ERC20Vault.TransactOpts, receiver, token, amount)
}

