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

// EthVaultABI is the input ABI used to generate the binding from.
const EthVaultABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositVerifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"name\":\"setDepositVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEffectiveDepositVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newDepositVerifierMaturityTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"safeGasStipend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPlasmaFramework\",\"name\":\"_framework\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_safeGasStipend\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blknum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nextDepositVerifier\",\"type\":\"address\"}],\"name\":\"SetDepositVerifierCalled\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_depositTx\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthVault is an auto generated Go binding around an Ethereum contract.
type EthVault struct {
	EthVaultCaller     // Read-only binding to the contract
	EthVaultTransactor // Write-only binding to the contract
}

// EthVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthVaultSession struct {
	Contract     *EthVault         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthVaultCallerSession struct {
	Contract *EthVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EthVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthVaultTransactorSession struct {
	Contract     *EthVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthVaultRaw struct {
	Contract *EthVault // Generic contract binding to access the raw methods on
}

// EthVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthVaultCallerRaw struct {
	Contract *EthVaultCaller // Generic read-only contract binding to access the raw methods on
}

// EthVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthVaultTransactorRaw struct {
	Contract *EthVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthVault creates a new instance of EthVault, bound to a specific deployed contract.
func NewEthVault(address common.Address, backend bind.ContractBackend) (*EthVault, error) {
	contract, err := bindEthVault(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthVault{EthVaultCaller: EthVaultCaller{contract: contract}, EthVaultTransactor: EthVaultTransactor{contract: contract}}, nil
}

// NewEthVaultCaller creates a new read-only instance of EthVault, bound to a specific deployed contract.
func NewEthVaultCaller(address common.Address, caller bind.ContractCaller) (*EthVaultCaller, error) {
	contract, err := bindEthVault(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &EthVaultCaller{contract: contract}, nil
}

// NewEthVaultTransactor creates a new write-only instance of EthVault, bound to a specific deployed contract.
func NewEthVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*EthVaultTransactor, error) {
	contract, err := bindEthVault(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &EthVaultTransactor{contract: contract}, nil
}

// bindEthVault binds a generic wrapper to an already deployed contract.
func bindEthVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, nil), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthVault *EthVaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthVault.Contract.EthVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthVault *EthVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthVault.Contract.EthVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthVault *EthVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthVault.Contract.EthVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthVault *EthVaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthVault *EthVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthVault *EthVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthVault.Contract.contract.Transact(opts, method, params...)
}

// DepositVerifiers is a free data retrieval call binding the contract method 0x02e9b07b.
//
// Solidity: function depositVerifiers( uint256) constant returns(address)
func (_EthVault *EthVaultCaller) DepositVerifiers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthVault.contract.Call(opts, out, "depositVerifiers", arg0)
	return *ret0, err
}

// DepositVerifiers is a free data retrieval call binding the contract method 0x02e9b07b.
//
// Solidity: function depositVerifiers( uint256) constant returns(address)
func (_EthVault *EthVaultSession) DepositVerifiers(arg0 *big.Int) (common.Address, error) {
	return _EthVault.Contract.DepositVerifiers(&_EthVault.CallOpts, arg0)
}

// DepositVerifiers is a free data retrieval call binding the contract method 0x02e9b07b.
//
// Solidity: function depositVerifiers( uint256) constant returns(address)
func (_EthVault *EthVaultCallerSession) DepositVerifiers(arg0 *big.Int) (common.Address, error) {
	return _EthVault.Contract.DepositVerifiers(&_EthVault.CallOpts, arg0)
}

// GetEffectiveDepositVerifier is a free data retrieval call binding the contract method 0x52518b7f.
//
// Solidity: function getEffectiveDepositVerifier() constant returns(address)
func (_EthVault *EthVaultCaller) GetEffectiveDepositVerifier(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthVault.contract.Call(opts, out, "getEffectiveDepositVerifier")
	return *ret0, err
}

// GetEffectiveDepositVerifier is a free data retrieval call binding the contract method 0x52518b7f.
//
// Solidity: function getEffectiveDepositVerifier() constant returns(address)
func (_EthVault *EthVaultSession) GetEffectiveDepositVerifier() (common.Address, error) {
	return _EthVault.Contract.GetEffectiveDepositVerifier(&_EthVault.CallOpts)
}

// GetEffectiveDepositVerifier is a free data retrieval call binding the contract method 0x52518b7f.
//
// Solidity: function getEffectiveDepositVerifier() constant returns(address)
func (_EthVault *EthVaultCallerSession) GetEffectiveDepositVerifier() (common.Address, error) {
	return _EthVault.Contract.GetEffectiveDepositVerifier(&_EthVault.CallOpts)
}

// NewDepositVerifierMaturityTimestamp is a free data retrieval call binding the contract method 0xa7e5994f.
//
// Solidity: function newDepositVerifierMaturityTimestamp() constant returns(uint256)
func (_EthVault *EthVaultCaller) NewDepositVerifierMaturityTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthVault.contract.Call(opts, out, "newDepositVerifierMaturityTimestamp")
	return *ret0, err
}

// NewDepositVerifierMaturityTimestamp is a free data retrieval call binding the contract method 0xa7e5994f.
//
// Solidity: function newDepositVerifierMaturityTimestamp() constant returns(uint256)
func (_EthVault *EthVaultSession) NewDepositVerifierMaturityTimestamp() (*big.Int, error) {
	return _EthVault.Contract.NewDepositVerifierMaturityTimestamp(&_EthVault.CallOpts)
}

// NewDepositVerifierMaturityTimestamp is a free data retrieval call binding the contract method 0xa7e5994f.
//
// Solidity: function newDepositVerifierMaturityTimestamp() constant returns(uint256)
func (_EthVault *EthVaultCallerSession) NewDepositVerifierMaturityTimestamp() (*big.Int, error) {
	return _EthVault.Contract.NewDepositVerifierMaturityTimestamp(&_EthVault.CallOpts)
}

// SafeGasStipend is a free data retrieval call binding the contract method 0xccbd2176.
//
// Solidity: function safeGasStipend() constant returns(uint256)
func (_EthVault *EthVaultCaller) SafeGasStipend(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthVault.contract.Call(opts, out, "safeGasStipend")
	return *ret0, err
}

// SafeGasStipend is a free data retrieval call binding the contract method 0xccbd2176.
//
// Solidity: function safeGasStipend() constant returns(uint256)
func (_EthVault *EthVaultSession) SafeGasStipend() (*big.Int, error) {
	return _EthVault.Contract.SafeGasStipend(&_EthVault.CallOpts)
}

// SafeGasStipend is a free data retrieval call binding the contract method 0xccbd2176.
//
// Solidity: function safeGasStipend() constant returns(uint256)
func (_EthVault *EthVaultCallerSession) SafeGasStipend() (*big.Int, error) {
	return _EthVault.Contract.SafeGasStipend(&_EthVault.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(_depositTx bytes) returns()
func (_EthVault *EthVaultTransactor) Deposit(opts *bind.TransactOpts, _depositTx []byte) (*types.Transaction, error) {
	return _EthVault.contract.Transact(opts, "deposit", _depositTx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(_depositTx bytes) returns()
func (_EthVault *EthVaultSession) Deposit(_depositTx []byte) (*types.Transaction, error) {
	return _EthVault.Contract.Deposit(&_EthVault.TransactOpts, _depositTx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(_depositTx bytes) returns()
func (_EthVault *EthVaultTransactorSession) Deposit(_depositTx []byte) (*types.Transaction, error) {
	return _EthVault.Contract.Deposit(&_EthVault.TransactOpts, _depositTx)
}

// SetDepositVerifier is a paid mutator transaction binding the contract method 0x2c6ce78b.
//
// Solidity: function setDepositVerifier(_verifier address) returns()
func (_EthVault *EthVaultTransactor) SetDepositVerifier(opts *bind.TransactOpts, _verifier common.Address) (*types.Transaction, error) {
	return _EthVault.contract.Transact(opts, "setDepositVerifier", _verifier)
}

// SetDepositVerifier is a paid mutator transaction binding the contract method 0x2c6ce78b.
//
// Solidity: function setDepositVerifier(_verifier address) returns()
func (_EthVault *EthVaultSession) SetDepositVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _EthVault.Contract.SetDepositVerifier(&_EthVault.TransactOpts, _verifier)
}

// SetDepositVerifier is a paid mutator transaction binding the contract method 0x2c6ce78b.
//
// Solidity: function setDepositVerifier(_verifier address) returns()
func (_EthVault *EthVaultTransactorSession) SetDepositVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _EthVault.Contract.SetDepositVerifier(&_EthVault.TransactOpts, _verifier)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(receiver address, amount uint256) returns()
func (_EthVault *EthVaultTransactor) Withdraw(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthVault.contract.Transact(opts, "withdraw", receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(receiver address, amount uint256) returns()
func (_EthVault *EthVaultSession) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthVault.Contract.Withdraw(&_EthVault.TransactOpts, receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(receiver address, amount uint256) returns()
func (_EthVault *EthVaultTransactorSession) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthVault.Contract.Withdraw(&_EthVault.TransactOpts, receiver, amount)
}

