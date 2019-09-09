// Copyright 2019 OmiseGO Pte Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rootchain

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// RootchainABI is the input ABI used to generate the binding from.
const RootchainABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"nextFeeExit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint192\"}],\"name\":\"exits\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint192\"}],\"name\":\"inFlightExits\",\"outputs\":[{\"name\":\"exitStartTimestamp\",\"type\":\"uint256\"},{\"name\":\"exitMap\",\"type\":\"uint256\"},{\"name\":\"bondOwner\",\"type\":\"address\"},{\"name\":\"oldestCompetitor\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_initOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tx\",\"type\":\"bytes\"},{\"name\":\"_outputIndex\",\"type\":\"uint256\"}],\"name\":\"getInFlightExitOutput\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_inFlightTx\",\"type\":\"bytes\"},{\"name\":\"_inFlightTxPos\",\"type\":\"uint256\"},{\"name\":\"_inFlightTxInclusionProof\",\"type\":\"bytes\"}],\"name\":\"respondToNonCanonicalChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_outputId\",\"type\":\"uint256\"}],\"name\":\"getExitableTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_inFlightTx\",\"type\":\"bytes\"},{\"name\":\"_inFlightTxInputIndex\",\"type\":\"uint8\"},{\"name\":\"_spendingTx\",\"type\":\"bytes\"},{\"name\":\"_spendingTxInputIndex\",\"type\":\"uint8\"},{\"name\":\"_spendingTxSig\",\"type\":\"bytes\"}],\"name\":\"challengeInFlightExitInputSpent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"startFeeExit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"piggybackBond\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_feeExitId\",\"type\":\"uint256\"}],\"name\":\"getFeeExitPriority\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextChildBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_exitId\",\"type\":\"uint192\"},{\"name\":\"_outputId\",\"type\":\"uint256\"}],\"name\":\"getStandardExitPriority\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_inFlightTx\",\"type\":\"bytes\"},{\"name\":\"_inFlightTxInputIndex\",\"type\":\"uint8\"},{\"name\":\"_competingTx\",\"type\":\"bytes\"},{\"name\":\"_competingTxInputIndex\",\"type\":\"uint8\"},{\"name\":\"_competingTxId\",\"type\":\"uint256\"},{\"name\":\"_competingTxInclusionProof\",\"type\":\"bytes\"},{\"name\":\"_competingTxSig\",\"type\":\"bytes\"}],\"name\":\"challengeInFlightExitNotCanonical\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"exitableTimestamp\",\"type\":\"uint64\"}],\"name\":\"isMature\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_inFlightTx\",\"type\":\"bytes\"},{\"name\":\"_inFlightTxOutputId\",\"type\":\"uint256\"},{\"name\":\"_inFlightTxInclusionProof\",\"type\":\"bytes\"},{\"name\":\"_spendingTx\",\"type\":\"bytes\"},{\"name\":\"_spendingTxInputIndex\",\"type\":\"uint256\"},{\"name\":\"_spendingTxSig\",\"type\":\"bytes\"}],\"name\":\"challengeInFlightExitOutputSpent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"clearFlag\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_inFlightTx\",\"type\":\"bytes\"},{\"name\":\"_outputIndex\",\"type\":\"uint8\"}],\"name\":\"piggybackInFlightExit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tx\",\"type\":\"bytes\"}],\"name\":\"getInFlightExitId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint192\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_inFlightTx\",\"type\":\"bytes\"},{\"name\":\"_inputTxs\",\"type\":\"bytes\"},{\"name\":\"_inputTxsInclusionProofs\",\"type\":\"bytes\"},{\"name\":\"_inFlightTxSigs\",\"type\":\"bytes\"}],\"name\":\"startInFlightExit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextDepositBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getNextExit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint192\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_topUtxoPos\",\"type\":\"uint256\"},{\"name\":\"_exitsToProcess\",\"type\":\"uint256\"}],\"name\":\"processExits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"standardExitBond\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_depositTx\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"hasToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"flagged\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHILD_BLOCK_INTERVAL\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_outputId\",\"type\":\"uint192\"},{\"name\":\"_outputTx\",\"type\":\"bytes\"},{\"name\":\"_outputTxInclusionProof\",\"type\":\"bytes\"}],\"name\":\"startStandardExit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inFlightExitBond\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blockRoot\",\"type\":\"bytes32\"}],\"name\":\"submitBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"priority\",\"type\":\"uint256\"}],\"name\":\"markStandard\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_outputId\",\"type\":\"uint192\"},{\"name\":\"_challengeTx\",\"type\":\"bytes\"},{\"name\":\"_inputIndex\",\"type\":\"uint8\"},{\"name\":\"_challengeTxSig\",\"type\":\"bytes\"}],\"name\":\"challengeStandardExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"exitsQueues\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setFlag\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"priority\",\"type\":\"uint256\"}],\"name\":\"markInFlight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_outputId\",\"type\":\"uint256\"}],\"name\":\"getStandardExitId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint192\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_EXIT_PERIOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_outputId\",\"type\":\"uint256\"},{\"name\":\"_tx\",\"type\":\"bytes\"}],\"name\":\"getInFlightExitPriority\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_depositTx\",\"type\":\"bytes\"}],\"name\":\"depositFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"blknum\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"outputId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ExitStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"utxoPos\",\"type\":\"uint256\"}],\"name\":\"ExitFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"utxoPos\",\"type\":\"uint256\"}],\"name\":\"ExitChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"InFlightExitStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"outputIndex\",\"type\":\"uint256\"}],\"name\":\"InFlightExitPiggybacked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"challengeTxPosition\",\"type\":\"uint256\"}],\"name\":\"InFlightExitChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"challengeTxPosition\",\"type\":\"uint256\"}],\"name\":\"InFlightExitChallengeResponded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"outputId\",\"type\":\"uint256\"}],\"name\":\"InFlightExitOutputBlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"inFlightExitId\",\"type\":\"uint192\"},{\"indexed\":false,\"name\":\"outputId\",\"type\":\"uint256\"}],\"name\":\"InFlightExitFinalized\",\"type\":\"event\"}]"

// Rootchain is an auto generated Go binding around an Ethereum contract.
type Rootchain struct {
	RootchainCaller     // Read-only binding to the contract
	RootchainTransactor // Write-only binding to the contract
	RootchainFilterer   // Log filterer for contract events
}

// RootchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootchainSession struct {
	Contract     *Rootchain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootchainCallerSession struct {
	Contract *RootchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RootchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootchainTransactorSession struct {
	Contract     *RootchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RootchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootchainRaw struct {
	Contract *Rootchain // Generic contract binding to access the raw methods on
}

// RootchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootchainCallerRaw struct {
	Contract *RootchainCaller // Generic read-only contract binding to access the raw methods on
}

// RootchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootchainTransactorRaw struct {
	Contract *RootchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootchain creates a new instance of Rootchain, bound to a specific deployed contract.
func NewRootchain(address common.Address, backend bind.ContractBackend) (*Rootchain, error) {
	contract, err := bindRootchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rootchain{RootchainCaller: RootchainCaller{contract: contract}, RootchainTransactor: RootchainTransactor{contract: contract}, RootchainFilterer: RootchainFilterer{contract: contract}}, nil
}

// NewRootchainCaller creates a new read-only instance of Rootchain, bound to a specific deployed contract.
func NewRootchainCaller(address common.Address, caller bind.ContractCaller) (*RootchainCaller, error) {
	contract, err := bindRootchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootchainCaller{contract: contract}, nil
}

// NewRootchainTransactor creates a new write-only instance of Rootchain, bound to a specific deployed contract.
func NewRootchainTransactor(address common.Address, transactor bind.ContractTransactor) (*RootchainTransactor, error) {
	contract, err := bindRootchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootchainTransactor{contract: contract}, nil
}

// NewRootchainFilterer creates a new log filterer instance of Rootchain, bound to a specific deployed contract.
func NewRootchainFilterer(address common.Address, filterer bind.ContractFilterer) (*RootchainFilterer, error) {
	contract, err := bindRootchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootchainFilterer{contract: contract}, nil
}

// bindRootchain binds a generic wrapper to an already deployed contract.
func bindRootchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootchainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rootchain *RootchainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rootchain.Contract.RootchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rootchain *RootchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rootchain.Contract.RootchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rootchain *RootchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rootchain.Contract.RootchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rootchain *RootchainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rootchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rootchain *RootchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rootchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rootchain *RootchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rootchain.Contract.contract.Transact(opts, method, params...)
}

// CHILDBLOCKINTERVAL is a free data retrieval call binding the contract method 0xa831fa07.
//
// Solidity: function CHILD_BLOCK_INTERVAL() constant returns(uint256)
func (_Rootchain *RootchainCaller) CHILDBLOCKINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "CHILD_BLOCK_INTERVAL")
	return *ret0, err
}

// CHILDBLOCKINTERVAL is a free data retrieval call binding the contract method 0xa831fa07.
//
// Solidity: function CHILD_BLOCK_INTERVAL() constant returns(uint256)
func (_Rootchain *RootchainSession) CHILDBLOCKINTERVAL() (*big.Int, error) {
	return _Rootchain.Contract.CHILDBLOCKINTERVAL(&_Rootchain.CallOpts)
}

// CHILDBLOCKINTERVAL is a free data retrieval call binding the contract method 0xa831fa07.
//
// Solidity: function CHILD_BLOCK_INTERVAL() constant returns(uint256)
func (_Rootchain *RootchainCallerSession) CHILDBLOCKINTERVAL() (*big.Int, error) {
	return _Rootchain.Contract.CHILDBLOCKINTERVAL(&_Rootchain.CallOpts)
}

// MINEXITPERIOD is a free data retrieval call binding the contract method 0xdf49c648.
//
// Solidity: function MIN_EXIT_PERIOD() constant returns(uint256)
func (_Rootchain *RootchainCaller) MINEXITPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "MIN_EXIT_PERIOD")
	return *ret0, err
}

// MINEXITPERIOD is a free data retrieval call binding the contract method 0xdf49c648.
//
// Solidity: function MIN_EXIT_PERIOD() constant returns(uint256)
func (_Rootchain *RootchainSession) MINEXITPERIOD() (*big.Int, error) {
	return _Rootchain.Contract.MINEXITPERIOD(&_Rootchain.CallOpts)
}

// MINEXITPERIOD is a free data retrieval call binding the contract method 0xdf49c648.
//
// Solidity: function MIN_EXIT_PERIOD() constant returns(uint256)
func (_Rootchain *RootchainCallerSession) MINEXITPERIOD() (*big.Int, error) {
	return _Rootchain.Contract.MINEXITPERIOD(&_Rootchain.CallOpts)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks( uint256) constant returns(root bytes32, timestamp uint256)
func (_Rootchain *RootchainCaller) Blocks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Root      [32]byte
	Timestamp *big.Int
}, error) {
	ret := new(struct {
		Root      [32]byte
		Timestamp *big.Int
	})
	out := ret
	err := _Rootchain.contract.Call(opts, out, "blocks", arg0)
	return *ret, err
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks( uint256) constant returns(root bytes32, timestamp uint256)
func (_Rootchain *RootchainSession) Blocks(arg0 *big.Int) (struct {
	Root      [32]byte
	Timestamp *big.Int
}, error) {
	return _Rootchain.Contract.Blocks(&_Rootchain.CallOpts, arg0)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks( uint256) constant returns(root bytes32, timestamp uint256)
func (_Rootchain *RootchainCallerSession) Blocks(arg0 *big.Int) (struct {
	Root      [32]byte
	Timestamp *big.Int
}, error) {
	return _Rootchain.Contract.Blocks(&_Rootchain.CallOpts, arg0)
}

// ClearFlag is a free data retrieval call binding the contract method 0x739471ca.
//
// Solidity: function clearFlag(_value uint256) constant returns(uint256)
func (_Rootchain *RootchainCaller) ClearFlag(opts *bind.CallOpts, _value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "clearFlag", _value)
	return *ret0, err
}

// ClearFlag is a free data retrieval call binding the contract method 0x739471ca.
//
// Solidity: function clearFlag(_value uint256) constant returns(uint256)
func (_Rootchain *RootchainSession) ClearFlag(_value *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.ClearFlag(&_Rootchain.CallOpts, _value)
}

// ClearFlag is a free data retrieval call binding the contract method 0x739471ca.
//
// Solidity: function clearFlag(_value uint256) constant returns(uint256)
func (_Rootchain *RootchainCallerSession) ClearFlag(_value *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.ClearFlag(&_Rootchain.CallOpts, _value)
}

// Exits is a free data retrieval call binding the contract method 0x178914ca.
//
// Solidity: function exits( uint192) constant returns(owner address, token address, amount uint256)
func (_Rootchain *RootchainCaller) Exits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner  common.Address
	Token  common.Address
	Amount *big.Int
}, error) {
	ret := new(struct {
		Owner  common.Address
		Token  common.Address
		Amount *big.Int
	})
	out := ret
	err := _Rootchain.contract.Call(opts, out, "exits", arg0)
	return *ret, err
}

// Exits is a free data retrieval call binding the contract method 0x178914ca.
//
// Solidity: function exits( uint192) constant returns(owner address, token address, amount uint256)
func (_Rootchain *RootchainSession) Exits(arg0 *big.Int) (struct {
	Owner  common.Address
	Token  common.Address
	Amount *big.Int
}, error) {
	return _Rootchain.Contract.Exits(&_Rootchain.CallOpts, arg0)
}

// Exits is a free data retrieval call binding the contract method 0x178914ca.
//
// Solidity: function exits( uint192) constant returns(owner address, token address, amount uint256)
func (_Rootchain *RootchainCallerSession) Exits(arg0 *big.Int) (struct {
	Owner  common.Address
	Token  common.Address
	Amount *big.Int
}, error) {
	return _Rootchain.Contract.Exits(&_Rootchain.CallOpts, arg0)
}

// ExitsQueues is a free data retrieval call binding the contract method 0xd11f045c.
//
// Solidity: function exitsQueues( address) constant returns(address)
func (_Rootchain *RootchainCaller) ExitsQueues(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "exitsQueues", arg0)
	return *ret0, err
}

// ExitsQueues is a free data retrieval call binding the contract method 0xd11f045c.
//
// Solidity: function exitsQueues( address) constant returns(address)
func (_Rootchain *RootchainSession) ExitsQueues(arg0 common.Address) (common.Address, error) {
	return _Rootchain.Contract.ExitsQueues(&_Rootchain.CallOpts, arg0)
}

// ExitsQueues is a free data retrieval call binding the contract method 0xd11f045c.
//
// Solidity: function exitsQueues( address) constant returns(address)
func (_Rootchain *RootchainCallerSession) ExitsQueues(arg0 common.Address) (common.Address, error) {
	return _Rootchain.Contract.ExitsQueues(&_Rootchain.CallOpts, arg0)
}

// Flagged is a free data retrieval call binding the contract method 0xa370c668.
//
// Solidity: function flagged(_value uint256) constant returns(bool)
func (_Rootchain *RootchainCaller) Flagged(opts *bind.CallOpts, _value *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "flagged", _value)
	return *ret0, err
}

// Flagged is a free data retrieval call binding the contract method 0xa370c668.
//
// Solidity: function flagged(_value uint256) constant returns(bool)
func (_Rootchain *RootchainSession) Flagged(_value *big.Int) (bool, error) {
	return _Rootchain.Contract.Flagged(&_Rootchain.CallOpts, _value)
}

// Flagged is a free data retrieval call binding the contract method 0xa370c668.
//
// Solidity: function flagged(_value uint256) constant returns(bool)
func (_Rootchain *RootchainCallerSession) Flagged(_value *big.Int) (bool, error) {
	return _Rootchain.Contract.Flagged(&_Rootchain.CallOpts, _value)
}

// GetDepositBlockNumber is a free data retrieval call binding the contract method 0xe24c1f36.
//
// Solidity: function getDepositBlockNumber() constant returns(uint256)
func (_Rootchain *RootchainCaller) GetDepositBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "getDepositBlockNumber")
	return *ret0, err
}

// GetDepositBlockNumber is a free data retrieval call binding the contract method 0xe24c1f36.
//
// Solidity: function getDepositBlockNumber() constant returns(uint256)
func (_Rootchain *RootchainSession) GetDepositBlockNumber() (*big.Int, error) {
	return _Rootchain.Contract.GetDepositBlockNumber(&_Rootchain.CallOpts)
}

// GetDepositBlockNumber is a free data retrieval call binding the contract method 0xe24c1f36.
//
// Solidity: function getDepositBlockNumber() constant returns(uint256)
func (_Rootchain *RootchainCallerSession) GetDepositBlockNumber() (*big.Int, error) {
	return _Rootchain.Contract.GetDepositBlockNumber(&_Rootchain.CallOpts)
}

// GetExitableTimestamp is a free data retrieval call binding the contract method 0x3ed39580.
//
// Solidity: function getExitableTimestamp(_outputId uint256) constant returns(uint256)
func (_Rootchain *RootchainCaller) GetExitableTimestamp(opts *bind.CallOpts, _outputId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "getExitableTimestamp", _outputId)
	return *ret0, err
}

// GetExitableTimestamp is a free data retrieval call binding the contract method 0x3ed39580.
//
// Solidity: function getExitableTimestamp(_outputId uint256) constant returns(uint256)
func (_Rootchain *RootchainSession) GetExitableTimestamp(_outputId *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.GetExitableTimestamp(&_Rootchain.CallOpts, _outputId)
}

// GetExitableTimestamp is a free data retrieval call binding the contract method 0x3ed39580.
//
// Solidity: function getExitableTimestamp(_outputId uint256) constant returns(uint256)
func (_Rootchain *RootchainCallerSession) GetExitableTimestamp(_outputId *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.GetExitableTimestamp(&_Rootchain.CallOpts, _outputId)
}

// GetFeeExitPriority is a free data retrieval call binding the contract method 0x46fdfa7b.
//
// Solidity: function getFeeExitPriority(_feeExitId uint256) constant returns(uint256)
func (_Rootchain *RootchainCaller) GetFeeExitPriority(opts *bind.CallOpts, _feeExitId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "getFeeExitPriority", _feeExitId)
	return *ret0, err
}

// GetFeeExitPriority is a free data retrieval call binding the contract method 0x46fdfa7b.
//
// Solidity: function getFeeExitPriority(_feeExitId uint256) constant returns(uint256)
func (_Rootchain *RootchainSession) GetFeeExitPriority(_feeExitId *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.GetFeeExitPriority(&_Rootchain.CallOpts, _feeExitId)
}

// GetFeeExitPriority is a free data retrieval call binding the contract method 0x46fdfa7b.
//
// Solidity: function getFeeExitPriority(_feeExitId uint256) constant returns(uint256)
func (_Rootchain *RootchainCallerSession) GetFeeExitPriority(_feeExitId *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.GetFeeExitPriority(&_Rootchain.CallOpts, _feeExitId)
}

// GetInFlightExitId is a free data retrieval call binding the contract method 0x839c78f9.
//
// Solidity: function getInFlightExitId(_tx bytes) constant returns(uint192)
func (_Rootchain *RootchainCaller) GetInFlightExitId(opts *bind.CallOpts, _tx []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "getInFlightExitId", _tx)
	return *ret0, err
}

// GetInFlightExitId is a free data retrieval call binding the contract method 0x839c78f9.
//
// Solidity: function getInFlightExitId(_tx bytes) constant returns(uint192)
func (_Rootchain *RootchainSession) GetInFlightExitId(_tx []byte) (*big.Int, error) {
	return _Rootchain.Contract.GetInFlightExitId(&_Rootchain.CallOpts, _tx)
}

// GetInFlightExitId is a free data retrieval call binding the contract method 0x839c78f9.
//
// Solidity: function getInFlightExitId(_tx bytes) constant returns(uint192)
func (_Rootchain *RootchainCallerSession) GetInFlightExitId(_tx []byte) (*big.Int, error) {
	return _Rootchain.Contract.GetInFlightExitId(&_Rootchain.CallOpts, _tx)
}

// GetInFlightExitOutput is a free data retrieval call binding the contract method 0x31294253.
//
// Solidity: function getInFlightExitOutput(_tx bytes, _outputIndex uint256) constant returns(address, address, uint256)
func (_Rootchain *RootchainCaller) GetInFlightExitOutput(opts *bind.CallOpts, _tx []byte, _outputIndex *big.Int) (common.Address, common.Address, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Rootchain.contract.Call(opts, out, "getInFlightExitOutput", _tx, _outputIndex)
	return *ret0, *ret1, *ret2, err
}

// GetInFlightExitOutput is a free data retrieval call binding the contract method 0x31294253.
//
// Solidity: function getInFlightExitOutput(_tx bytes, _outputIndex uint256) constant returns(address, address, uint256)
func (_Rootchain *RootchainSession) GetInFlightExitOutput(_tx []byte, _outputIndex *big.Int) (common.Address, common.Address, *big.Int, error) {
	return _Rootchain.Contract.GetInFlightExitOutput(&_Rootchain.CallOpts, _tx, _outputIndex)
}

// GetInFlightExitOutput is a free data retrieval call binding the contract method 0x31294253.
//
// Solidity: function getInFlightExitOutput(_tx bytes, _outputIndex uint256) constant returns(address, address, uint256)
func (_Rootchain *RootchainCallerSession) GetInFlightExitOutput(_tx []byte, _outputIndex *big.Int) (common.Address, common.Address, *big.Int, error) {
	return _Rootchain.Contract.GetInFlightExitOutput(&_Rootchain.CallOpts, _tx, _outputIndex)
}

// GetInFlightExitPriority is a free data retrieval call binding the contract method 0xf53d04be.
//
// Solidity: function getInFlightExitPriority(_outputId uint256, _tx bytes) constant returns(uint256)
func (_Rootchain *RootchainCaller) GetInFlightExitPriority(opts *bind.CallOpts, _outputId *big.Int, _tx []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "getInFlightExitPriority", _outputId, _tx)
	return *ret0, err
}

// GetInFlightExitPriority is a free data retrieval call binding the contract method 0xf53d04be.
//
// Solidity: function getInFlightExitPriority(_outputId uint256, _tx bytes) constant returns(uint256)
func (_Rootchain *RootchainSession) GetInFlightExitPriority(_outputId *big.Int, _tx []byte) (*big.Int, error) {
	return _Rootchain.Contract.GetInFlightExitPriority(&_Rootchain.CallOpts, _outputId, _tx)
}

// GetInFlightExitPriority is a free data retrieval call binding the contract method 0xf53d04be.
//
// Solidity: function getInFlightExitPriority(_outputId uint256, _tx bytes) constant returns(uint256)
func (_Rootchain *RootchainCallerSession) GetInFlightExitPriority(_outputId *big.Int, _tx []byte) (*big.Int, error) {
	return _Rootchain.Contract.GetInFlightExitPriority(&_Rootchain.CallOpts, _outputId, _tx)
}

// GetNextExit is a free data retrieval call binding the contract method 0x8bfe0aec.
//
// Solidity: function getNextExit(_token address) constant returns(uint64, uint192, bool)
func (_Rootchain *RootchainCaller) GetNextExit(opts *bind.CallOpts, _token common.Address) (uint64, *big.Int, bool, error) {
	var (
		ret0 = new(uint64)
		ret1 = new(*big.Int)
		ret2 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Rootchain.contract.Call(opts, out, "getNextExit", _token)
	return *ret0, *ret1, *ret2, err
}

// GetNextExit is a free data retrieval call binding the contract method 0x8bfe0aec.
//
// Solidity: function getNextExit(_token address) constant returns(uint64, uint192, bool)
func (_Rootchain *RootchainSession) GetNextExit(_token common.Address) (uint64, *big.Int, bool, error) {
	return _Rootchain.Contract.GetNextExit(&_Rootchain.CallOpts, _token)
}

// GetNextExit is a free data retrieval call binding the contract method 0x8bfe0aec.
//
// Solidity: function getNextExit(_token address) constant returns(uint64, uint192, bool)
func (_Rootchain *RootchainCallerSession) GetNextExit(_token common.Address) (uint64, *big.Int, bool, error) {
	return _Rootchain.Contract.GetNextExit(&_Rootchain.CallOpts, _token)
}

// GetStandardExitId is a free data retrieval call binding the contract method 0xdef3c897.
//
// Solidity: function getStandardExitId(_outputId uint256) constant returns(uint192)
func (_Rootchain *RootchainCaller) GetStandardExitId(opts *bind.CallOpts, _outputId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "getStandardExitId", _outputId)
	return *ret0, err
}

// GetStandardExitId is a free data retrieval call binding the contract method 0xdef3c897.
//
// Solidity: function getStandardExitId(_outputId uint256) constant returns(uint192)
func (_Rootchain *RootchainSession) GetStandardExitId(_outputId *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.GetStandardExitId(&_Rootchain.CallOpts, _outputId)
}

// GetStandardExitId is a free data retrieval call binding the contract method 0xdef3c897.
//
// Solidity: function getStandardExitId(_outputId uint256) constant returns(uint192)
func (_Rootchain *RootchainCallerSession) GetStandardExitId(_outputId *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.GetStandardExitId(&_Rootchain.CallOpts, _outputId)
}

// GetStandardExitPriority is a free data retrieval call binding the contract method 0x4cf307db.
//
// Solidity: function getStandardExitPriority(_exitId uint192, _outputId uint256) constant returns(uint256)
func (_Rootchain *RootchainCaller) GetStandardExitPriority(opts *bind.CallOpts, _exitId *big.Int, _outputId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "getStandardExitPriority", _exitId, _outputId)
	return *ret0, err
}

// GetStandardExitPriority is a free data retrieval call binding the contract method 0x4cf307db.
//
// Solidity: function getStandardExitPriority(_exitId uint192, _outputId uint256) constant returns(uint256)
func (_Rootchain *RootchainSession) GetStandardExitPriority(_exitId *big.Int, _outputId *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.GetStandardExitPriority(&_Rootchain.CallOpts, _exitId, _outputId)
}

// GetStandardExitPriority is a free data retrieval call binding the contract method 0x4cf307db.
//
// Solidity: function getStandardExitPriority(_exitId uint192, _outputId uint256) constant returns(uint256)
func (_Rootchain *RootchainCallerSession) GetStandardExitPriority(_exitId *big.Int, _outputId *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.GetStandardExitPriority(&_Rootchain.CallOpts, _exitId, _outputId)
}

// HasToken is a free data retrieval call binding the contract method 0x9bb0f599.
//
// Solidity: function hasToken(_token address) constant returns(bool)
func (_Rootchain *RootchainCaller) HasToken(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "hasToken", _token)
	return *ret0, err
}

// HasToken is a free data retrieval call binding the contract method 0x9bb0f599.
//
// Solidity: function hasToken(_token address) constant returns(bool)
func (_Rootchain *RootchainSession) HasToken(_token common.Address) (bool, error) {
	return _Rootchain.Contract.HasToken(&_Rootchain.CallOpts, _token)
}

// HasToken is a free data retrieval call binding the contract method 0x9bb0f599.
//
// Solidity: function hasToken(_token address) constant returns(bool)
func (_Rootchain *RootchainCallerSession) HasToken(_token common.Address) (bool, error) {
	return _Rootchain.Contract.HasToken(&_Rootchain.CallOpts, _token)
}

// InFlightExitBond is a free data retrieval call binding the contract method 0xb55dfc08.
//
// Solidity: function inFlightExitBond() constant returns(uint256)
func (_Rootchain *RootchainCaller) InFlightExitBond(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "inFlightExitBond")
	return *ret0, err
}

// InFlightExitBond is a free data retrieval call binding the contract method 0xb55dfc08.
//
// Solidity: function inFlightExitBond() constant returns(uint256)
func (_Rootchain *RootchainSession) InFlightExitBond() (*big.Int, error) {
	return _Rootchain.Contract.InFlightExitBond(&_Rootchain.CallOpts)
}

// InFlightExitBond is a free data retrieval call binding the contract method 0xb55dfc08.
//
// Solidity: function inFlightExitBond() constant returns(uint256)
func (_Rootchain *RootchainCallerSession) InFlightExitBond() (*big.Int, error) {
	return _Rootchain.Contract.InFlightExitBond(&_Rootchain.CallOpts)
}

// InFlightExits is a free data retrieval call binding the contract method 0x17a10af3.
//
// Solidity: function inFlightExits( uint192) constant returns(exitStartTimestamp uint256, exitMap uint256, bondOwner address, oldestCompetitor uint256)
func (_Rootchain *RootchainCaller) InFlightExits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ExitStartTimestamp *big.Int
	ExitMap            *big.Int
	BondOwner          common.Address
	OldestCompetitor   *big.Int
}, error) {
	ret := new(struct {
		ExitStartTimestamp *big.Int
		ExitMap            *big.Int
		BondOwner          common.Address
		OldestCompetitor   *big.Int
	})
	out := ret
	err := _Rootchain.contract.Call(opts, out, "inFlightExits", arg0)
	return *ret, err
}

// InFlightExits is a free data retrieval call binding the contract method 0x17a10af3.
//
// Solidity: function inFlightExits( uint192) constant returns(exitStartTimestamp uint256, exitMap uint256, bondOwner address, oldestCompetitor uint256)
func (_Rootchain *RootchainSession) InFlightExits(arg0 *big.Int) (struct {
	ExitStartTimestamp *big.Int
	ExitMap            *big.Int
	BondOwner          common.Address
	OldestCompetitor   *big.Int
}, error) {
	return _Rootchain.Contract.InFlightExits(&_Rootchain.CallOpts, arg0)
}

// InFlightExits is a free data retrieval call binding the contract method 0x17a10af3.
//
// Solidity: function inFlightExits( uint192) constant returns(exitStartTimestamp uint256, exitMap uint256, bondOwner address, oldestCompetitor uint256)
func (_Rootchain *RootchainCallerSession) InFlightExits(arg0 *big.Int) (struct {
	ExitStartTimestamp *big.Int
	ExitMap            *big.Int
	BondOwner          common.Address
	OldestCompetitor   *big.Int
}, error) {
	return _Rootchain.Contract.InFlightExits(&_Rootchain.CallOpts, arg0)
}

// IsMature is a free data retrieval call binding the contract method 0x52817086.
//
// Solidity: function isMature(exitableTimestamp uint64) constant returns(bool)
func (_Rootchain *RootchainCaller) IsMature(opts *bind.CallOpts, exitableTimestamp uint64) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "isMature", exitableTimestamp)
	return *ret0, err
}

// IsMature is a free data retrieval call binding the contract method 0x52817086.
//
// Solidity: function isMature(exitableTimestamp uint64) constant returns(bool)
func (_Rootchain *RootchainSession) IsMature(exitableTimestamp uint64) (bool, error) {
	return _Rootchain.Contract.IsMature(&_Rootchain.CallOpts, exitableTimestamp)
}

// IsMature is a free data retrieval call binding the contract method 0x52817086.
//
// Solidity: function isMature(exitableTimestamp uint64) constant returns(bool)
func (_Rootchain *RootchainCallerSession) IsMature(exitableTimestamp uint64) (bool, error) {
	return _Rootchain.Contract.IsMature(&_Rootchain.CallOpts, exitableTimestamp)
}

// MarkInFlight is a free data retrieval call binding the contract method 0xdb48110f.
//
// Solidity: function markInFlight(priority uint256) constant returns(uint256)
func (_Rootchain *RootchainCaller) MarkInFlight(opts *bind.CallOpts, priority *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "markInFlight", priority)
	return *ret0, err
}

// MarkInFlight is a free data retrieval call binding the contract method 0xdb48110f.
//
// Solidity: function markInFlight(priority uint256) constant returns(uint256)
func (_Rootchain *RootchainSession) MarkInFlight(priority *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.MarkInFlight(&_Rootchain.CallOpts, priority)
}

// MarkInFlight is a free data retrieval call binding the contract method 0xdb48110f.
//
// Solidity: function markInFlight(priority uint256) constant returns(uint256)
func (_Rootchain *RootchainCallerSession) MarkInFlight(priority *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.MarkInFlight(&_Rootchain.CallOpts, priority)
}

// MarkStandard is a free data retrieval call binding the contract method 0xc0158b86.
//
// Solidity: function markStandard(priority uint256) constant returns(uint256)
func (_Rootchain *RootchainCaller) MarkStandard(opts *bind.CallOpts, priority *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "markStandard", priority)
	return *ret0, err
}

// MarkStandard is a free data retrieval call binding the contract method 0xc0158b86.
//
// Solidity: function markStandard(priority uint256) constant returns(uint256)
func (_Rootchain *RootchainSession) MarkStandard(priority *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.MarkStandard(&_Rootchain.CallOpts, priority)
}

// MarkStandard is a free data retrieval call binding the contract method 0xc0158b86.
//
// Solidity: function markStandard(priority uint256) constant returns(uint256)
func (_Rootchain *RootchainCallerSession) MarkStandard(priority *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.MarkStandard(&_Rootchain.CallOpts, priority)
}

// NextChildBlock is a free data retrieval call binding the contract method 0x4ca8714f.
//
// Solidity: function nextChildBlock() constant returns(uint256)
func (_Rootchain *RootchainCaller) NextChildBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "nextChildBlock")
	return *ret0, err
}

// NextChildBlock is a free data retrieval call binding the contract method 0x4ca8714f.
//
// Solidity: function nextChildBlock() constant returns(uint256)
func (_Rootchain *RootchainSession) NextChildBlock() (*big.Int, error) {
	return _Rootchain.Contract.NextChildBlock(&_Rootchain.CallOpts)
}

// NextChildBlock is a free data retrieval call binding the contract method 0x4ca8714f.
//
// Solidity: function nextChildBlock() constant returns(uint256)
func (_Rootchain *RootchainCallerSession) NextChildBlock() (*big.Int, error) {
	return _Rootchain.Contract.NextChildBlock(&_Rootchain.CallOpts)
}

// NextDepositBlock is a free data retrieval call binding the contract method 0x8701fc5d.
//
// Solidity: function nextDepositBlock() constant returns(uint256)
func (_Rootchain *RootchainCaller) NextDepositBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "nextDepositBlock")
	return *ret0, err
}

// NextDepositBlock is a free data retrieval call binding the contract method 0x8701fc5d.
//
// Solidity: function nextDepositBlock() constant returns(uint256)
func (_Rootchain *RootchainSession) NextDepositBlock() (*big.Int, error) {
	return _Rootchain.Contract.NextDepositBlock(&_Rootchain.CallOpts)
}

// NextDepositBlock is a free data retrieval call binding the contract method 0x8701fc5d.
//
// Solidity: function nextDepositBlock() constant returns(uint256)
func (_Rootchain *RootchainCallerSession) NextDepositBlock() (*big.Int, error) {
	return _Rootchain.Contract.NextDepositBlock(&_Rootchain.CallOpts)
}

// NextFeeExit is a free data retrieval call binding the contract method 0x03f3536e.
//
// Solidity: function nextFeeExit() constant returns(uint256)
func (_Rootchain *RootchainCaller) NextFeeExit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "nextFeeExit")
	return *ret0, err
}

// NextFeeExit is a free data retrieval call binding the contract method 0x03f3536e.
//
// Solidity: function nextFeeExit() constant returns(uint256)
func (_Rootchain *RootchainSession) NextFeeExit() (*big.Int, error) {
	return _Rootchain.Contract.NextFeeExit(&_Rootchain.CallOpts)
}

// NextFeeExit is a free data retrieval call binding the contract method 0x03f3536e.
//
// Solidity: function nextFeeExit() constant returns(uint256)
func (_Rootchain *RootchainCallerSession) NextFeeExit() (*big.Int, error) {
	return _Rootchain.Contract.NextFeeExit(&_Rootchain.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_Rootchain *RootchainCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_Rootchain *RootchainSession) Operator() (common.Address, error) {
	return _Rootchain.Contract.Operator(&_Rootchain.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_Rootchain *RootchainCallerSession) Operator() (common.Address, error) {
	return _Rootchain.Contract.Operator(&_Rootchain.CallOpts)
}

// PiggybackBond is a free data retrieval call binding the contract method 0x45ab91e2.
//
// Solidity: function piggybackBond() constant returns(uint256)
func (_Rootchain *RootchainCaller) PiggybackBond(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "piggybackBond")
	return *ret0, err
}

// PiggybackBond is a free data retrieval call binding the contract method 0x45ab91e2.
//
// Solidity: function piggybackBond() constant returns(uint256)
func (_Rootchain *RootchainSession) PiggybackBond() (*big.Int, error) {
	return _Rootchain.Contract.PiggybackBond(&_Rootchain.CallOpts)
}

// PiggybackBond is a free data retrieval call binding the contract method 0x45ab91e2.
//
// Solidity: function piggybackBond() constant returns(uint256)
func (_Rootchain *RootchainCallerSession) PiggybackBond() (*big.Int, error) {
	return _Rootchain.Contract.PiggybackBond(&_Rootchain.CallOpts)
}

// SetFlag is a free data retrieval call binding the contract method 0xd40c79f0.
//
// Solidity: function setFlag(_value uint256) constant returns(uint256)
func (_Rootchain *RootchainCaller) SetFlag(opts *bind.CallOpts, _value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "setFlag", _value)
	return *ret0, err
}

// SetFlag is a free data retrieval call binding the contract method 0xd40c79f0.
//
// Solidity: function setFlag(_value uint256) constant returns(uint256)
func (_Rootchain *RootchainSession) SetFlag(_value *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.SetFlag(&_Rootchain.CallOpts, _value)
}

// SetFlag is a free data retrieval call binding the contract method 0xd40c79f0.
//
// Solidity: function setFlag(_value uint256) constant returns(uint256)
func (_Rootchain *RootchainCallerSession) SetFlag(_value *big.Int) (*big.Int, error) {
	return _Rootchain.Contract.SetFlag(&_Rootchain.CallOpts, _value)
}

// StandardExitBond is a free data retrieval call binding the contract method 0x97ecc2eb.
//
// Solidity: function standardExitBond() constant returns(uint256)
func (_Rootchain *RootchainCaller) StandardExitBond(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rootchain.contract.Call(opts, out, "standardExitBond")
	return *ret0, err
}

// StandardExitBond is a free data retrieval call binding the contract method 0x97ecc2eb.
//
// Solidity: function standardExitBond() constant returns(uint256)
func (_Rootchain *RootchainSession) StandardExitBond() (*big.Int, error) {
	return _Rootchain.Contract.StandardExitBond(&_Rootchain.CallOpts)
}

// StandardExitBond is a free data retrieval call binding the contract method 0x97ecc2eb.
//
// Solidity: function standardExitBond() constant returns(uint256)
func (_Rootchain *RootchainCallerSession) StandardExitBond() (*big.Int, error) {
	return _Rootchain.Contract.StandardExitBond(&_Rootchain.CallOpts)
}

// InitOperator is a paid mutator transaction binding the contract method 0x2cc5034f.
//
// Solidity: function _initOperator() returns()
func (_Rootchain *RootchainTransactor) InitOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "_initOperator")
}

// InitOperator is a paid mutator transaction binding the contract method 0x2cc5034f.
//
// Solidity: function _initOperator() returns()
func (_Rootchain *RootchainSession) InitOperator() (*types.Transaction, error) {
	return _Rootchain.Contract.InitOperator(&_Rootchain.TransactOpts)
}

// InitOperator is a paid mutator transaction binding the contract method 0x2cc5034f.
//
// Solidity: function _initOperator() returns()
func (_Rootchain *RootchainTransactorSession) InitOperator() (*types.Transaction, error) {
	return _Rootchain.Contract.InitOperator(&_Rootchain.TransactOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(_token address) returns()
func (_Rootchain *RootchainTransactor) AddToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "addToken", _token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(_token address) returns()
func (_Rootchain *RootchainSession) AddToken(_token common.Address) (*types.Transaction, error) {
	return _Rootchain.Contract.AddToken(&_Rootchain.TransactOpts, _token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(_token address) returns()
func (_Rootchain *RootchainTransactorSession) AddToken(_token common.Address) (*types.Transaction, error) {
	return _Rootchain.Contract.AddToken(&_Rootchain.TransactOpts, _token)
}

// ChallengeInFlightExitInputSpent is a paid mutator transaction binding the contract method 0x40823bb7.
//
// Solidity: function challengeInFlightExitInputSpent(_inFlightTx bytes, _inFlightTxInputIndex uint8, _spendingTx bytes, _spendingTxInputIndex uint8, _spendingTxSig bytes) returns()
func (_Rootchain *RootchainTransactor) ChallengeInFlightExitInputSpent(opts *bind.TransactOpts, _inFlightTx []byte, _inFlightTxInputIndex uint8, _spendingTx []byte, _spendingTxInputIndex uint8, _spendingTxSig []byte) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "challengeInFlightExitInputSpent", _inFlightTx, _inFlightTxInputIndex, _spendingTx, _spendingTxInputIndex, _spendingTxSig)
}

// ChallengeInFlightExitInputSpent is a paid mutator transaction binding the contract method 0x40823bb7.
//
// Solidity: function challengeInFlightExitInputSpent(_inFlightTx bytes, _inFlightTxInputIndex uint8, _spendingTx bytes, _spendingTxInputIndex uint8, _spendingTxSig bytes) returns()
func (_Rootchain *RootchainSession) ChallengeInFlightExitInputSpent(_inFlightTx []byte, _inFlightTxInputIndex uint8, _spendingTx []byte, _spendingTxInputIndex uint8, _spendingTxSig []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.ChallengeInFlightExitInputSpent(&_Rootchain.TransactOpts, _inFlightTx, _inFlightTxInputIndex, _spendingTx, _spendingTxInputIndex, _spendingTxSig)
}

// ChallengeInFlightExitInputSpent is a paid mutator transaction binding the contract method 0x40823bb7.
//
// Solidity: function challengeInFlightExitInputSpent(_inFlightTx bytes, _inFlightTxInputIndex uint8, _spendingTx bytes, _spendingTxInputIndex uint8, _spendingTxSig bytes) returns()
func (_Rootchain *RootchainTransactorSession) ChallengeInFlightExitInputSpent(_inFlightTx []byte, _inFlightTxInputIndex uint8, _spendingTx []byte, _spendingTxInputIndex uint8, _spendingTxSig []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.ChallengeInFlightExitInputSpent(&_Rootchain.TransactOpts, _inFlightTx, _inFlightTxInputIndex, _spendingTx, _spendingTxInputIndex, _spendingTxSig)
}

// ChallengeInFlightExitNotCanonical is a paid mutator transaction binding the contract method 0x4ffa32cf.
//
// Solidity: function challengeInFlightExitNotCanonical(_inFlightTx bytes, _inFlightTxInputIndex uint8, _competingTx bytes, _competingTxInputIndex uint8, _competingTxId uint256, _competingTxInclusionProof bytes, _competingTxSig bytes) returns()
func (_Rootchain *RootchainTransactor) ChallengeInFlightExitNotCanonical(opts *bind.TransactOpts, _inFlightTx []byte, _inFlightTxInputIndex uint8, _competingTx []byte, _competingTxInputIndex uint8, _competingTxId *big.Int, _competingTxInclusionProof []byte, _competingTxSig []byte) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "challengeInFlightExitNotCanonical", _inFlightTx, _inFlightTxInputIndex, _competingTx, _competingTxInputIndex, _competingTxId, _competingTxInclusionProof, _competingTxSig)
}

// ChallengeInFlightExitNotCanonical is a paid mutator transaction binding the contract method 0x4ffa32cf.
//
// Solidity: function challengeInFlightExitNotCanonical(_inFlightTx bytes, _inFlightTxInputIndex uint8, _competingTx bytes, _competingTxInputIndex uint8, _competingTxId uint256, _competingTxInclusionProof bytes, _competingTxSig bytes) returns()
func (_Rootchain *RootchainSession) ChallengeInFlightExitNotCanonical(_inFlightTx []byte, _inFlightTxInputIndex uint8, _competingTx []byte, _competingTxInputIndex uint8, _competingTxId *big.Int, _competingTxInclusionProof []byte, _competingTxSig []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.ChallengeInFlightExitNotCanonical(&_Rootchain.TransactOpts, _inFlightTx, _inFlightTxInputIndex, _competingTx, _competingTxInputIndex, _competingTxId, _competingTxInclusionProof, _competingTxSig)
}

// ChallengeInFlightExitNotCanonical is a paid mutator transaction binding the contract method 0x4ffa32cf.
//
// Solidity: function challengeInFlightExitNotCanonical(_inFlightTx bytes, _inFlightTxInputIndex uint8, _competingTx bytes, _competingTxInputIndex uint8, _competingTxId uint256, _competingTxInclusionProof bytes, _competingTxSig bytes) returns()
func (_Rootchain *RootchainTransactorSession) ChallengeInFlightExitNotCanonical(_inFlightTx []byte, _inFlightTxInputIndex uint8, _competingTx []byte, _competingTxInputIndex uint8, _competingTxId *big.Int, _competingTxInclusionProof []byte, _competingTxSig []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.ChallengeInFlightExitNotCanonical(&_Rootchain.TransactOpts, _inFlightTx, _inFlightTxInputIndex, _competingTx, _competingTxInputIndex, _competingTxId, _competingTxInclusionProof, _competingTxSig)
}

// ChallengeInFlightExitOutputSpent is a paid mutator transaction binding the contract method 0x6c5bada0.
//
// Solidity: function challengeInFlightExitOutputSpent(_inFlightTx bytes, _inFlightTxOutputId uint256, _inFlightTxInclusionProof bytes, _spendingTx bytes, _spendingTxInputIndex uint256, _spendingTxSig bytes) returns()
func (_Rootchain *RootchainTransactor) ChallengeInFlightExitOutputSpent(opts *bind.TransactOpts, _inFlightTx []byte, _inFlightTxOutputId *big.Int, _inFlightTxInclusionProof []byte, _spendingTx []byte, _spendingTxInputIndex *big.Int, _spendingTxSig []byte) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "challengeInFlightExitOutputSpent", _inFlightTx, _inFlightTxOutputId, _inFlightTxInclusionProof, _spendingTx, _spendingTxInputIndex, _spendingTxSig)
}

// ChallengeInFlightExitOutputSpent is a paid mutator transaction binding the contract method 0x6c5bada0.
//
// Solidity: function challengeInFlightExitOutputSpent(_inFlightTx bytes, _inFlightTxOutputId uint256, _inFlightTxInclusionProof bytes, _spendingTx bytes, _spendingTxInputIndex uint256, _spendingTxSig bytes) returns()
func (_Rootchain *RootchainSession) ChallengeInFlightExitOutputSpent(_inFlightTx []byte, _inFlightTxOutputId *big.Int, _inFlightTxInclusionProof []byte, _spendingTx []byte, _spendingTxInputIndex *big.Int, _spendingTxSig []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.ChallengeInFlightExitOutputSpent(&_Rootchain.TransactOpts, _inFlightTx, _inFlightTxOutputId, _inFlightTxInclusionProof, _spendingTx, _spendingTxInputIndex, _spendingTxSig)
}

// ChallengeInFlightExitOutputSpent is a paid mutator transaction binding the contract method 0x6c5bada0.
//
// Solidity: function challengeInFlightExitOutputSpent(_inFlightTx bytes, _inFlightTxOutputId uint256, _inFlightTxInclusionProof bytes, _spendingTx bytes, _spendingTxInputIndex uint256, _spendingTxSig bytes) returns()
func (_Rootchain *RootchainTransactorSession) ChallengeInFlightExitOutputSpent(_inFlightTx []byte, _inFlightTxOutputId *big.Int, _inFlightTxInclusionProof []byte, _spendingTx []byte, _spendingTxInputIndex *big.Int, _spendingTxSig []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.ChallengeInFlightExitOutputSpent(&_Rootchain.TransactOpts, _inFlightTx, _inFlightTxOutputId, _inFlightTxInclusionProof, _spendingTx, _spendingTxInputIndex, _spendingTxSig)
}

// ChallengeStandardExit is a paid mutator transaction binding the contract method 0xe5d99629.
//
// Solidity: function challengeStandardExit(_outputId uint192, _challengeTx bytes, _inputIndex uint8, _challengeTxSig bytes) returns()
func (_Rootchain *RootchainTransactor) ChallengeStandardExit(opts *bind.TransactOpts, _outputId *big.Int, _challengeTx []byte, _inputIndex uint8, _challengeTxSig []byte) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "challengeStandardExit", _outputId, _challengeTx, _inputIndex, _challengeTxSig)
}

// ChallengeStandardExit is a paid mutator transaction binding the contract method 0xe5d99629.
//
// Solidity: function challengeStandardExit(_outputId uint192, _challengeTx bytes, _inputIndex uint8, _challengeTxSig bytes) returns()
func (_Rootchain *RootchainSession) ChallengeStandardExit(_outputId *big.Int, _challengeTx []byte, _inputIndex uint8, _challengeTxSig []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.ChallengeStandardExit(&_Rootchain.TransactOpts, _outputId, _challengeTx, _inputIndex, _challengeTxSig)
}

// ChallengeStandardExit is a paid mutator transaction binding the contract method 0xe5d99629.
//
// Solidity: function challengeStandardExit(_outputId uint192, _challengeTx bytes, _inputIndex uint8, _challengeTxSig bytes) returns()
func (_Rootchain *RootchainTransactorSession) ChallengeStandardExit(_outputId *big.Int, _challengeTx []byte, _inputIndex uint8, _challengeTxSig []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.ChallengeStandardExit(&_Rootchain.TransactOpts, _outputId, _challengeTx, _inputIndex, _challengeTxSig)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(_depositTx bytes) returns()
func (_Rootchain *RootchainTransactor) Deposit(opts *bind.TransactOpts, _depositTx []byte) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "deposit", _depositTx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(_depositTx bytes) returns()
func (_Rootchain *RootchainSession) Deposit(_depositTx []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.Deposit(&_Rootchain.TransactOpts, _depositTx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(_depositTx bytes) returns()
func (_Rootchain *RootchainTransactorSession) Deposit(_depositTx []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.Deposit(&_Rootchain.TransactOpts, _depositTx)
}

// DepositFrom is a paid mutator transaction binding the contract method 0xf8cbd6d2.
//
// Solidity: function depositFrom(_depositTx bytes) returns()
func (_Rootchain *RootchainTransactor) DepositFrom(opts *bind.TransactOpts, _depositTx []byte) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "depositFrom", _depositTx)
}

// DepositFrom is a paid mutator transaction binding the contract method 0xf8cbd6d2.
//
// Solidity: function depositFrom(_depositTx bytes) returns()
func (_Rootchain *RootchainSession) DepositFrom(_depositTx []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.DepositFrom(&_Rootchain.TransactOpts, _depositTx)
}

// DepositFrom is a paid mutator transaction binding the contract method 0xf8cbd6d2.
//
// Solidity: function depositFrom(_depositTx bytes) returns()
func (_Rootchain *RootchainTransactorSession) DepositFrom(_depositTx []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.DepositFrom(&_Rootchain.TransactOpts, _depositTx)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Rootchain *RootchainTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Rootchain *RootchainSession) Init() (*types.Transaction, error) {
	return _Rootchain.Contract.Init(&_Rootchain.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Rootchain *RootchainTransactorSession) Init() (*types.Transaction, error) {
	return _Rootchain.Contract.Init(&_Rootchain.TransactOpts)
}

// PiggybackInFlightExit is a paid mutator transaction binding the contract method 0x750b05ac.
//
// Solidity: function piggybackInFlightExit(_inFlightTx bytes, _outputIndex uint8) returns()
func (_Rootchain *RootchainTransactor) PiggybackInFlightExit(opts *bind.TransactOpts, _inFlightTx []byte, _outputIndex uint8) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "piggybackInFlightExit", _inFlightTx, _outputIndex)
}

// PiggybackInFlightExit is a paid mutator transaction binding the contract method 0x750b05ac.
//
// Solidity: function piggybackInFlightExit(_inFlightTx bytes, _outputIndex uint8) returns()
func (_Rootchain *RootchainSession) PiggybackInFlightExit(_inFlightTx []byte, _outputIndex uint8) (*types.Transaction, error) {
	return _Rootchain.Contract.PiggybackInFlightExit(&_Rootchain.TransactOpts, _inFlightTx, _outputIndex)
}

// PiggybackInFlightExit is a paid mutator transaction binding the contract method 0x750b05ac.
//
// Solidity: function piggybackInFlightExit(_inFlightTx bytes, _outputIndex uint8) returns()
func (_Rootchain *RootchainTransactorSession) PiggybackInFlightExit(_inFlightTx []byte, _outputIndex uint8) (*types.Transaction, error) {
	return _Rootchain.Contract.PiggybackInFlightExit(&_Rootchain.TransactOpts, _inFlightTx, _outputIndex)
}

// ProcessExits is a paid mutator transaction binding the contract method 0x90d45ef5.
//
// Solidity: function processExits(_token address, _topUtxoPos uint256, _exitsToProcess uint256) returns()
func (_Rootchain *RootchainTransactor) ProcessExits(opts *bind.TransactOpts, _token common.Address, _topUtxoPos *big.Int, _exitsToProcess *big.Int) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "processExits", _token, _topUtxoPos, _exitsToProcess)
}

// ProcessExits is a paid mutator transaction binding the contract method 0x90d45ef5.
//
// Solidity: function processExits(_token address, _topUtxoPos uint256, _exitsToProcess uint256) returns()
func (_Rootchain *RootchainSession) ProcessExits(_token common.Address, _topUtxoPos *big.Int, _exitsToProcess *big.Int) (*types.Transaction, error) {
	return _Rootchain.Contract.ProcessExits(&_Rootchain.TransactOpts, _token, _topUtxoPos, _exitsToProcess)
}

// ProcessExits is a paid mutator transaction binding the contract method 0x90d45ef5.
//
// Solidity: function processExits(_token address, _topUtxoPos uint256, _exitsToProcess uint256) returns()
func (_Rootchain *RootchainTransactorSession) ProcessExits(_token common.Address, _topUtxoPos *big.Int, _exitsToProcess *big.Int) (*types.Transaction, error) {
	return _Rootchain.Contract.ProcessExits(&_Rootchain.TransactOpts, _token, _topUtxoPos, _exitsToProcess)
}

// RespondToNonCanonicalChallenge is a paid mutator transaction binding the contract method 0x36660de4.
//
// Solidity: function respondToNonCanonicalChallenge(_inFlightTx bytes, _inFlightTxPos uint256, _inFlightTxInclusionProof bytes) returns()
func (_Rootchain *RootchainTransactor) RespondToNonCanonicalChallenge(opts *bind.TransactOpts, _inFlightTx []byte, _inFlightTxPos *big.Int, _inFlightTxInclusionProof []byte) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "respondToNonCanonicalChallenge", _inFlightTx, _inFlightTxPos, _inFlightTxInclusionProof)
}

// RespondToNonCanonicalChallenge is a paid mutator transaction binding the contract method 0x36660de4.
//
// Solidity: function respondToNonCanonicalChallenge(_inFlightTx bytes, _inFlightTxPos uint256, _inFlightTxInclusionProof bytes) returns()
func (_Rootchain *RootchainSession) RespondToNonCanonicalChallenge(_inFlightTx []byte, _inFlightTxPos *big.Int, _inFlightTxInclusionProof []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.RespondToNonCanonicalChallenge(&_Rootchain.TransactOpts, _inFlightTx, _inFlightTxPos, _inFlightTxInclusionProof)
}

// RespondToNonCanonicalChallenge is a paid mutator transaction binding the contract method 0x36660de4.
//
// Solidity: function respondToNonCanonicalChallenge(_inFlightTx bytes, _inFlightTxPos uint256, _inFlightTxInclusionProof bytes) returns()
func (_Rootchain *RootchainTransactorSession) RespondToNonCanonicalChallenge(_inFlightTx []byte, _inFlightTxPos *big.Int, _inFlightTxInclusionProof []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.RespondToNonCanonicalChallenge(&_Rootchain.TransactOpts, _inFlightTx, _inFlightTxPos, _inFlightTxInclusionProof)
}

// StartFeeExit is a paid mutator transaction binding the contract method 0x40b8d53a.
//
// Solidity: function startFeeExit(_token address, _amount uint256) returns()
func (_Rootchain *RootchainTransactor) StartFeeExit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "startFeeExit", _token, _amount)
}

// StartFeeExit is a paid mutator transaction binding the contract method 0x40b8d53a.
//
// Solidity: function startFeeExit(_token address, _amount uint256) returns()
func (_Rootchain *RootchainSession) StartFeeExit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Rootchain.Contract.StartFeeExit(&_Rootchain.TransactOpts, _token, _amount)
}

// StartFeeExit is a paid mutator transaction binding the contract method 0x40b8d53a.
//
// Solidity: function startFeeExit(_token address, _amount uint256) returns()
func (_Rootchain *RootchainTransactorSession) StartFeeExit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Rootchain.Contract.StartFeeExit(&_Rootchain.TransactOpts, _token, _amount)
}

// StartInFlightExit is a paid mutator transaction binding the contract method 0x84612195.
//
// Solidity: function startInFlightExit(_inFlightTx bytes, _inputTxs bytes, _inputTxsInclusionProofs bytes, _inFlightTxSigs bytes) returns()
func (_Rootchain *RootchainTransactor) StartInFlightExit(opts *bind.TransactOpts, _inFlightTx []byte, _inputTxs []byte, _inputTxsInclusionProofs []byte, _inFlightTxSigs []byte) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "startInFlightExit", _inFlightTx, _inputTxs, _inputTxsInclusionProofs, _inFlightTxSigs)
}

// StartInFlightExit is a paid mutator transaction binding the contract method 0x84612195.
//
// Solidity: function startInFlightExit(_inFlightTx bytes, _inputTxs bytes, _inputTxsInclusionProofs bytes, _inFlightTxSigs bytes) returns()
func (_Rootchain *RootchainSession) StartInFlightExit(_inFlightTx []byte, _inputTxs []byte, _inputTxsInclusionProofs []byte, _inFlightTxSigs []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.StartInFlightExit(&_Rootchain.TransactOpts, _inFlightTx, _inputTxs, _inputTxsInclusionProofs, _inFlightTxSigs)
}

// StartInFlightExit is a paid mutator transaction binding the contract method 0x84612195.
//
// Solidity: function startInFlightExit(_inFlightTx bytes, _inputTxs bytes, _inputTxsInclusionProofs bytes, _inFlightTxSigs bytes) returns()
func (_Rootchain *RootchainTransactorSession) StartInFlightExit(_inFlightTx []byte, _inputTxs []byte, _inputTxsInclusionProofs []byte, _inFlightTxSigs []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.StartInFlightExit(&_Rootchain.TransactOpts, _inFlightTx, _inputTxs, _inputTxsInclusionProofs, _inFlightTxSigs)
}

// StartStandardExit is a paid mutator transaction binding the contract method 0xa86e0efb.
//
// Solidity: function startStandardExit(_outputId uint192, _outputTx bytes, _outputTxInclusionProof bytes) returns()
func (_Rootchain *RootchainTransactor) StartStandardExit(opts *bind.TransactOpts, _outputId *big.Int, _outputTx []byte, _outputTxInclusionProof []byte) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "startStandardExit", _outputId, _outputTx, _outputTxInclusionProof)
}

// StartStandardExit is a paid mutator transaction binding the contract method 0xa86e0efb.
//
// Solidity: function startStandardExit(_outputId uint192, _outputTx bytes, _outputTxInclusionProof bytes) returns()
func (_Rootchain *RootchainSession) StartStandardExit(_outputId *big.Int, _outputTx []byte, _outputTxInclusionProof []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.StartStandardExit(&_Rootchain.TransactOpts, _outputId, _outputTx, _outputTxInclusionProof)
}

// StartStandardExit is a paid mutator transaction binding the contract method 0xa86e0efb.
//
// Solidity: function startStandardExit(_outputId uint192, _outputTx bytes, _outputTxInclusionProof bytes) returns()
func (_Rootchain *RootchainTransactorSession) StartStandardExit(_outputId *big.Int, _outputTx []byte, _outputTxInclusionProof []byte) (*types.Transaction, error) {
	return _Rootchain.Contract.StartStandardExit(&_Rootchain.TransactOpts, _outputId, _outputTx, _outputTxInclusionProof)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(_blockRoot bytes32) returns()
func (_Rootchain *RootchainTransactor) SubmitBlock(opts *bind.TransactOpts, _blockRoot [32]byte) (*types.Transaction, error) {
	return _Rootchain.contract.Transact(opts, "submitBlock", _blockRoot)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(_blockRoot bytes32) returns()
func (_Rootchain *RootchainSession) SubmitBlock(_blockRoot [32]byte) (*types.Transaction, error) {
	return _Rootchain.Contract.SubmitBlock(&_Rootchain.TransactOpts, _blockRoot)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(_blockRoot bytes32) returns()
func (_Rootchain *RootchainTransactorSession) SubmitBlock(_blockRoot [32]byte) (*types.Transaction, error) {
	return _Rootchain.Contract.SubmitBlock(&_Rootchain.TransactOpts, _blockRoot)
}

// RootchainBlockSubmittedIterator is returned from FilterBlockSubmitted and is used to iterate over the raw logs and unpacked data for BlockSubmitted events raised by the Rootchain contract.
type RootchainBlockSubmittedIterator struct {
	Event *RootchainBlockSubmitted // Event containing the contract specifics and raw log

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
func (it *RootchainBlockSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootchainBlockSubmitted)
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
		it.Event = new(RootchainBlockSubmitted)
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
func (it *RootchainBlockSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootchainBlockSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootchainBlockSubmitted represents a BlockSubmitted event raised by the Rootchain contract.
type RootchainBlockSubmitted struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockSubmitted is a free log retrieval operation binding the contract event 0x5a978f4723b249ccf79cd7a658a8601ce1ff8b89fc770251a6be35216351ce32.
//
// Solidity: e BlockSubmitted(blockNumber uint256)
func (_Rootchain *RootchainFilterer) FilterBlockSubmitted(opts *bind.FilterOpts) (*RootchainBlockSubmittedIterator, error) {

	logs, sub, err := _Rootchain.contract.FilterLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return &RootchainBlockSubmittedIterator{contract: _Rootchain.contract, event: "BlockSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSubmitted is a free log subscription operation binding the contract event 0x5a978f4723b249ccf79cd7a658a8601ce1ff8b89fc770251a6be35216351ce32.
//
// Solidity: e BlockSubmitted(blockNumber uint256)
func (_Rootchain *RootchainFilterer) WatchBlockSubmitted(opts *bind.WatchOpts, sink chan<- *RootchainBlockSubmitted) (event.Subscription, error) {

	logs, sub, err := _Rootchain.contract.WatchLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootchainBlockSubmitted)
				if err := _Rootchain.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
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

// RootchainDepositCreatedIterator is returned from FilterDepositCreated and is used to iterate over the raw logs and unpacked data for DepositCreated events raised by the Rootchain contract.
type RootchainDepositCreatedIterator struct {
	Event *RootchainDepositCreated // Event containing the contract specifics and raw log

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
func (it *RootchainDepositCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootchainDepositCreated)
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
		it.Event = new(RootchainDepositCreated)
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
func (it *RootchainDepositCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootchainDepositCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootchainDepositCreated represents a DepositCreated event raised by the Rootchain contract.
type RootchainDepositCreated struct {
	Depositor common.Address
	Blknum    *big.Int
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositCreated is a free log retrieval operation binding the contract event 0x18569122d84f30025bb8dffb33563f1bdbfb9637f21552b11b8305686e9cb307.
//
// Solidity: e DepositCreated(depositor indexed address, blknum indexed uint256, token indexed address, amount uint256)
func (_Rootchain *RootchainFilterer) FilterDepositCreated(opts *bind.FilterOpts, depositor []common.Address, blknum []*big.Int, token []common.Address) (*RootchainDepositCreatedIterator, error) {

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

	logs, sub, err := _Rootchain.contract.FilterLogs(opts, "DepositCreated", depositorRule, blknumRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &RootchainDepositCreatedIterator{contract: _Rootchain.contract, event: "DepositCreated", logs: logs, sub: sub}, nil
}

// WatchDepositCreated is a free log subscription operation binding the contract event 0x18569122d84f30025bb8dffb33563f1bdbfb9637f21552b11b8305686e9cb307.
//
// Solidity: e DepositCreated(depositor indexed address, blknum indexed uint256, token indexed address, amount uint256)
func (_Rootchain *RootchainFilterer) WatchDepositCreated(opts *bind.WatchOpts, sink chan<- *RootchainDepositCreated, depositor []common.Address, blknum []*big.Int, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Rootchain.contract.WatchLogs(opts, "DepositCreated", depositorRule, blknumRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootchainDepositCreated)
				if err := _Rootchain.contract.UnpackLog(event, "DepositCreated", log); err != nil {
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

// RootchainExitChallengedIterator is returned from FilterExitChallenged and is used to iterate over the raw logs and unpacked data for ExitChallenged events raised by the Rootchain contract.
type RootchainExitChallengedIterator struct {
	Event *RootchainExitChallenged // Event containing the contract specifics and raw log

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
func (it *RootchainExitChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootchainExitChallenged)
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
		it.Event = new(RootchainExitChallenged)
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
func (it *RootchainExitChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootchainExitChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootchainExitChallenged represents a ExitChallenged event raised by the Rootchain contract.
type RootchainExitChallenged struct {
	UtxoPos *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExitChallenged is a free log retrieval operation binding the contract event 0x5dfba526c59b25f899f935c5b0d5b8739e97e4d89c38c158eca3192ea34b87d8.
//
// Solidity: e ExitChallenged(utxoPos indexed uint256)
func (_Rootchain *RootchainFilterer) FilterExitChallenged(opts *bind.FilterOpts, utxoPos []*big.Int) (*RootchainExitChallengedIterator, error) {

	var utxoPosRule []interface{}
	for _, utxoPosItem := range utxoPos {
		utxoPosRule = append(utxoPosRule, utxoPosItem)
	}

	logs, sub, err := _Rootchain.contract.FilterLogs(opts, "ExitChallenged", utxoPosRule)
	if err != nil {
		return nil, err
	}
	return &RootchainExitChallengedIterator{contract: _Rootchain.contract, event: "ExitChallenged", logs: logs, sub: sub}, nil
}

// WatchExitChallenged is a free log subscription operation binding the contract event 0x5dfba526c59b25f899f935c5b0d5b8739e97e4d89c38c158eca3192ea34b87d8.
//
// Solidity: e ExitChallenged(utxoPos indexed uint256)
func (_Rootchain *RootchainFilterer) WatchExitChallenged(opts *bind.WatchOpts, sink chan<- *RootchainExitChallenged, utxoPos []*big.Int) (event.Subscription, error) {

	var utxoPosRule []interface{}
	for _, utxoPosItem := range utxoPos {
		utxoPosRule = append(utxoPosRule, utxoPosItem)
	}

	logs, sub, err := _Rootchain.contract.WatchLogs(opts, "ExitChallenged", utxoPosRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootchainExitChallenged)
				if err := _Rootchain.contract.UnpackLog(event, "ExitChallenged", log); err != nil {
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

// RootchainExitFinalizedIterator is returned from FilterExitFinalized and is used to iterate over the raw logs and unpacked data for ExitFinalized events raised by the Rootchain contract.
type RootchainExitFinalizedIterator struct {
	Event *RootchainExitFinalized // Event containing the contract specifics and raw log

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
func (it *RootchainExitFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootchainExitFinalized)
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
		it.Event = new(RootchainExitFinalized)
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
func (it *RootchainExitFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootchainExitFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootchainExitFinalized represents a ExitFinalized event raised by the Rootchain contract.
type RootchainExitFinalized struct {
	UtxoPos *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExitFinalized is a free log retrieval operation binding the contract event 0x36805ae734d39c3675a2dfe79c18d15a03680ff0589e7119958435a3bdd61b36.
//
// Solidity: e ExitFinalized(utxoPos indexed uint256)
func (_Rootchain *RootchainFilterer) FilterExitFinalized(opts *bind.FilterOpts, utxoPos []*big.Int) (*RootchainExitFinalizedIterator, error) {

	var utxoPosRule []interface{}
	for _, utxoPosItem := range utxoPos {
		utxoPosRule = append(utxoPosRule, utxoPosItem)
	}

	logs, sub, err := _Rootchain.contract.FilterLogs(opts, "ExitFinalized", utxoPosRule)
	if err != nil {
		return nil, err
	}
	return &RootchainExitFinalizedIterator{contract: _Rootchain.contract, event: "ExitFinalized", logs: logs, sub: sub}, nil
}

// WatchExitFinalized is a free log subscription operation binding the contract event 0x36805ae734d39c3675a2dfe79c18d15a03680ff0589e7119958435a3bdd61b36.
//
// Solidity: e ExitFinalized(utxoPos indexed uint256)
func (_Rootchain *RootchainFilterer) WatchExitFinalized(opts *bind.WatchOpts, sink chan<- *RootchainExitFinalized, utxoPos []*big.Int) (event.Subscription, error) {

	var utxoPosRule []interface{}
	for _, utxoPosItem := range utxoPos {
		utxoPosRule = append(utxoPosRule, utxoPosItem)
	}

	logs, sub, err := _Rootchain.contract.WatchLogs(opts, "ExitFinalized", utxoPosRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootchainExitFinalized)
				if err := _Rootchain.contract.UnpackLog(event, "ExitFinalized", log); err != nil {
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

// RootchainExitStartedIterator is returned from FilterExitStarted and is used to iterate over the raw logs and unpacked data for ExitStarted events raised by the Rootchain contract.
type RootchainExitStartedIterator struct {
	Event *RootchainExitStarted // Event containing the contract specifics and raw log

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
func (it *RootchainExitStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootchainExitStarted)
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
		it.Event = new(RootchainExitStarted)
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
func (it *RootchainExitStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootchainExitStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootchainExitStarted represents a ExitStarted event raised by the Rootchain contract.
type RootchainExitStarted struct {
	Owner    common.Address
	OutputId *big.Int
	Amount   *big.Int
	Token    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExitStarted is a free log retrieval operation binding the contract event 0x774f4dfa540020b406b2c42f8c27303b81aa155dc7e197270ec27ba79e33d250.
//
// Solidity: e ExitStarted(owner indexed address, outputId uint256, amount uint256, token address)
func (_Rootchain *RootchainFilterer) FilterExitStarted(opts *bind.FilterOpts, owner []common.Address) (*RootchainExitStartedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Rootchain.contract.FilterLogs(opts, "ExitStarted", ownerRule)
	if err != nil {
		return nil, err
	}
	return &RootchainExitStartedIterator{contract: _Rootchain.contract, event: "ExitStarted", logs: logs, sub: sub}, nil
}

// WatchExitStarted is a free log subscription operation binding the contract event 0x774f4dfa540020b406b2c42f8c27303b81aa155dc7e197270ec27ba79e33d250.
//
// Solidity: e ExitStarted(owner indexed address, outputId uint256, amount uint256, token address)
func (_Rootchain *RootchainFilterer) WatchExitStarted(opts *bind.WatchOpts, sink chan<- *RootchainExitStarted, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Rootchain.contract.WatchLogs(opts, "ExitStarted", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootchainExitStarted)
				if err := _Rootchain.contract.UnpackLog(event, "ExitStarted", log); err != nil {
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

// RootchainInFlightExitChallengeRespondedIterator is returned from FilterInFlightExitChallengeResponded and is used to iterate over the raw logs and unpacked data for InFlightExitChallengeResponded events raised by the Rootchain contract.
type RootchainInFlightExitChallengeRespondedIterator struct {
	Event *RootchainInFlightExitChallengeResponded // Event containing the contract specifics and raw log

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
func (it *RootchainInFlightExitChallengeRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootchainInFlightExitChallengeResponded)
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
		it.Event = new(RootchainInFlightExitChallengeResponded)
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
func (it *RootchainInFlightExitChallengeRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootchainInFlightExitChallengeRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootchainInFlightExitChallengeResponded represents a InFlightExitChallengeResponded event raised by the Rootchain contract.
type RootchainInFlightExitChallengeResponded struct {
	Challenger          common.Address
	TxHash              [32]byte
	ChallengeTxPosition *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitChallengeResponded is a free log retrieval operation binding the contract event 0x637cc4a7148767df19331a5c7dfb6d31f0a7e159a3dbb28a716be18c8c74f768.
//
// Solidity: e InFlightExitChallengeResponded(challenger address, txHash bytes32, challengeTxPosition uint256)
func (_Rootchain *RootchainFilterer) FilterInFlightExitChallengeResponded(opts *bind.FilterOpts) (*RootchainInFlightExitChallengeRespondedIterator, error) {

	logs, sub, err := _Rootchain.contract.FilterLogs(opts, "InFlightExitChallengeResponded")
	if err != nil {
		return nil, err
	}
	return &RootchainInFlightExitChallengeRespondedIterator{contract: _Rootchain.contract, event: "InFlightExitChallengeResponded", logs: logs, sub: sub}, nil
}

// WatchInFlightExitChallengeResponded is a free log subscription operation binding the contract event 0x637cc4a7148767df19331a5c7dfb6d31f0a7e159a3dbb28a716be18c8c74f768.
//
// Solidity: e InFlightExitChallengeResponded(challenger address, txHash bytes32, challengeTxPosition uint256)
func (_Rootchain *RootchainFilterer) WatchInFlightExitChallengeResponded(opts *bind.WatchOpts, sink chan<- *RootchainInFlightExitChallengeResponded) (event.Subscription, error) {

	logs, sub, err := _Rootchain.contract.WatchLogs(opts, "InFlightExitChallengeResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootchainInFlightExitChallengeResponded)
				if err := _Rootchain.contract.UnpackLog(event, "InFlightExitChallengeResponded", log); err != nil {
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

// RootchainInFlightExitChallengedIterator is returned from FilterInFlightExitChallenged and is used to iterate over the raw logs and unpacked data for InFlightExitChallenged events raised by the Rootchain contract.
type RootchainInFlightExitChallengedIterator struct {
	Event *RootchainInFlightExitChallenged // Event containing the contract specifics and raw log

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
func (it *RootchainInFlightExitChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootchainInFlightExitChallenged)
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
		it.Event = new(RootchainInFlightExitChallenged)
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
func (it *RootchainInFlightExitChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootchainInFlightExitChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootchainInFlightExitChallenged represents a InFlightExitChallenged event raised by the Rootchain contract.
type RootchainInFlightExitChallenged struct {
	Challenger          common.Address
	TxHash              [32]byte
	ChallengeTxPosition *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitChallenged is a free log retrieval operation binding the contract event 0x687401968e501bda2d2d6f880dd1a0a56ff50b1787185ee0b6f4c3fb9fc417ab.
//
// Solidity: e InFlightExitChallenged(challenger indexed address, txHash bytes32, challengeTxPosition uint256)
func (_Rootchain *RootchainFilterer) FilterInFlightExitChallenged(opts *bind.FilterOpts, challenger []common.Address) (*RootchainInFlightExitChallengedIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Rootchain.contract.FilterLogs(opts, "InFlightExitChallenged", challengerRule)
	if err != nil {
		return nil, err
	}
	return &RootchainInFlightExitChallengedIterator{contract: _Rootchain.contract, event: "InFlightExitChallenged", logs: logs, sub: sub}, nil
}

// WatchInFlightExitChallenged is a free log subscription operation binding the contract event 0x687401968e501bda2d2d6f880dd1a0a56ff50b1787185ee0b6f4c3fb9fc417ab.
//
// Solidity: e InFlightExitChallenged(challenger indexed address, txHash bytes32, challengeTxPosition uint256)
func (_Rootchain *RootchainFilterer) WatchInFlightExitChallenged(opts *bind.WatchOpts, sink chan<- *RootchainInFlightExitChallenged, challenger []common.Address) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Rootchain.contract.WatchLogs(opts, "InFlightExitChallenged", challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootchainInFlightExitChallenged)
				if err := _Rootchain.contract.UnpackLog(event, "InFlightExitChallenged", log); err != nil {
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

// RootchainInFlightExitFinalizedIterator is returned from FilterInFlightExitFinalized and is used to iterate over the raw logs and unpacked data for InFlightExitFinalized events raised by the Rootchain contract.
type RootchainInFlightExitFinalizedIterator struct {
	Event *RootchainInFlightExitFinalized // Event containing the contract specifics and raw log

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
func (it *RootchainInFlightExitFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootchainInFlightExitFinalized)
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
		it.Event = new(RootchainInFlightExitFinalized)
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
func (it *RootchainInFlightExitFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootchainInFlightExitFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootchainInFlightExitFinalized represents a InFlightExitFinalized event raised by the Rootchain contract.
type RootchainInFlightExitFinalized struct {
	InFlightExitId *big.Int
	OutputId       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitFinalized is a free log retrieval operation binding the contract event 0xe79fbb7bdf83c454ec9bd621a4e216dda59abe8837976fc9c848cab9f4a9c181.
//
// Solidity: e InFlightExitFinalized(inFlightExitId uint192, outputId uint256)
func (_Rootchain *RootchainFilterer) FilterInFlightExitFinalized(opts *bind.FilterOpts) (*RootchainInFlightExitFinalizedIterator, error) {

	logs, sub, err := _Rootchain.contract.FilterLogs(opts, "InFlightExitFinalized")
	if err != nil {
		return nil, err
	}
	return &RootchainInFlightExitFinalizedIterator{contract: _Rootchain.contract, event: "InFlightExitFinalized", logs: logs, sub: sub}, nil
}

// WatchInFlightExitFinalized is a free log subscription operation binding the contract event 0xe79fbb7bdf83c454ec9bd621a4e216dda59abe8837976fc9c848cab9f4a9c181.
//
// Solidity: e InFlightExitFinalized(inFlightExitId uint192, outputId uint256)
func (_Rootchain *RootchainFilterer) WatchInFlightExitFinalized(opts *bind.WatchOpts, sink chan<- *RootchainInFlightExitFinalized) (event.Subscription, error) {

	logs, sub, err := _Rootchain.contract.WatchLogs(opts, "InFlightExitFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootchainInFlightExitFinalized)
				if err := _Rootchain.contract.UnpackLog(event, "InFlightExitFinalized", log); err != nil {
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

// RootchainInFlightExitOutputBlockedIterator is returned from FilterInFlightExitOutputBlocked and is used to iterate over the raw logs and unpacked data for InFlightExitOutputBlocked events raised by the Rootchain contract.
type RootchainInFlightExitOutputBlockedIterator struct {
	Event *RootchainInFlightExitOutputBlocked // Event containing the contract specifics and raw log

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
func (it *RootchainInFlightExitOutputBlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootchainInFlightExitOutputBlocked)
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
		it.Event = new(RootchainInFlightExitOutputBlocked)
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
func (it *RootchainInFlightExitOutputBlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootchainInFlightExitOutputBlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootchainInFlightExitOutputBlocked represents a InFlightExitOutputBlocked event raised by the Rootchain contract.
type RootchainInFlightExitOutputBlocked struct {
	Challenger common.Address
	TxHash     [32]byte
	OutputId   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitOutputBlocked is a free log retrieval operation binding the contract event 0x23540ff515aa4f71173bfad1bdd9d0af58ead4a8bf5e1f3ad0ae9f3696802784.
//
// Solidity: e InFlightExitOutputBlocked(challenger indexed address, txHash bytes32, outputId uint256)
func (_Rootchain *RootchainFilterer) FilterInFlightExitOutputBlocked(opts *bind.FilterOpts, challenger []common.Address) (*RootchainInFlightExitOutputBlockedIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Rootchain.contract.FilterLogs(opts, "InFlightExitOutputBlocked", challengerRule)
	if err != nil {
		return nil, err
	}
	return &RootchainInFlightExitOutputBlockedIterator{contract: _Rootchain.contract, event: "InFlightExitOutputBlocked", logs: logs, sub: sub}, nil
}

// WatchInFlightExitOutputBlocked is a free log subscription operation binding the contract event 0x23540ff515aa4f71173bfad1bdd9d0af58ead4a8bf5e1f3ad0ae9f3696802784.
//
// Solidity: e InFlightExitOutputBlocked(challenger indexed address, txHash bytes32, outputId uint256)
func (_Rootchain *RootchainFilterer) WatchInFlightExitOutputBlocked(opts *bind.WatchOpts, sink chan<- *RootchainInFlightExitOutputBlocked, challenger []common.Address) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Rootchain.contract.WatchLogs(opts, "InFlightExitOutputBlocked", challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootchainInFlightExitOutputBlocked)
				if err := _Rootchain.contract.UnpackLog(event, "InFlightExitOutputBlocked", log); err != nil {
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

// RootchainInFlightExitPiggybackedIterator is returned from FilterInFlightExitPiggybacked and is used to iterate over the raw logs and unpacked data for InFlightExitPiggybacked events raised by the Rootchain contract.
type RootchainInFlightExitPiggybackedIterator struct {
	Event *RootchainInFlightExitPiggybacked // Event containing the contract specifics and raw log

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
func (it *RootchainInFlightExitPiggybackedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootchainInFlightExitPiggybacked)
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
		it.Event = new(RootchainInFlightExitPiggybacked)
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
func (it *RootchainInFlightExitPiggybackedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootchainInFlightExitPiggybackedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootchainInFlightExitPiggybacked represents a InFlightExitPiggybacked event raised by the Rootchain contract.
type RootchainInFlightExitPiggybacked struct {
	Owner       common.Address
	TxHash      [32]byte
	OutputIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitPiggybacked is a free log retrieval operation binding the contract event 0x07e1a98d5a6eb9c83cdc6ab7d8862424f9ea3f5f414f8e19c9d6fdda0c073533.
//
// Solidity: e InFlightExitPiggybacked(owner indexed address, txHash bytes32, outputIndex uint256)
func (_Rootchain *RootchainFilterer) FilterInFlightExitPiggybacked(opts *bind.FilterOpts, owner []common.Address) (*RootchainInFlightExitPiggybackedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Rootchain.contract.FilterLogs(opts, "InFlightExitPiggybacked", ownerRule)
	if err != nil {
		return nil, err
	}
	return &RootchainInFlightExitPiggybackedIterator{contract: _Rootchain.contract, event: "InFlightExitPiggybacked", logs: logs, sub: sub}, nil
}

// WatchInFlightExitPiggybacked is a free log subscription operation binding the contract event 0x07e1a98d5a6eb9c83cdc6ab7d8862424f9ea3f5f414f8e19c9d6fdda0c073533.
//
// Solidity: e InFlightExitPiggybacked(owner indexed address, txHash bytes32, outputIndex uint256)
func (_Rootchain *RootchainFilterer) WatchInFlightExitPiggybacked(opts *bind.WatchOpts, sink chan<- *RootchainInFlightExitPiggybacked, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Rootchain.contract.WatchLogs(opts, "InFlightExitPiggybacked", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootchainInFlightExitPiggybacked)
				if err := _Rootchain.contract.UnpackLog(event, "InFlightExitPiggybacked", log); err != nil {
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

// RootchainInFlightExitStartedIterator is returned from FilterInFlightExitStarted and is used to iterate over the raw logs and unpacked data for InFlightExitStarted events raised by the Rootchain contract.
type RootchainInFlightExitStartedIterator struct {
	Event *RootchainInFlightExitStarted // Event containing the contract specifics and raw log

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
func (it *RootchainInFlightExitStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootchainInFlightExitStarted)
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
		it.Event = new(RootchainInFlightExitStarted)
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
func (it *RootchainInFlightExitStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootchainInFlightExitStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootchainInFlightExitStarted represents a InFlightExitStarted event raised by the Rootchain contract.
type RootchainInFlightExitStarted struct {
	Initiator common.Address
	TxHash    [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitStarted is a free log retrieval operation binding the contract event 0xd5f1fe9d48880b57daa227004b16d320c0eb885d6c49d472d54c16a05fa3179e.
//
// Solidity: e InFlightExitStarted(initiator indexed address, txHash bytes32)
func (_Rootchain *RootchainFilterer) FilterInFlightExitStarted(opts *bind.FilterOpts, initiator []common.Address) (*RootchainInFlightExitStartedIterator, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _Rootchain.contract.FilterLogs(opts, "InFlightExitStarted", initiatorRule)
	if err != nil {
		return nil, err
	}
	return &RootchainInFlightExitStartedIterator{contract: _Rootchain.contract, event: "InFlightExitStarted", logs: logs, sub: sub}, nil
}

// WatchInFlightExitStarted is a free log subscription operation binding the contract event 0xd5f1fe9d48880b57daa227004b16d320c0eb885d6c49d472d54c16a05fa3179e.
//
// Solidity: e InFlightExitStarted(initiator indexed address, txHash bytes32)
func (_Rootchain *RootchainFilterer) WatchInFlightExitStarted(opts *bind.WatchOpts, sink chan<- *RootchainInFlightExitStarted, initiator []common.Address) (event.Subscription, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _Rootchain.contract.WatchLogs(opts, "InFlightExitStarted", initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootchainInFlightExitStarted)
				if err := _Rootchain.contract.UnpackLog(event, "InFlightExitStarted", log); err != nil {
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

// RootchainTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the Rootchain contract.
type RootchainTokenAddedIterator struct {
	Event *RootchainTokenAdded // Event containing the contract specifics and raw log

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
func (it *RootchainTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootchainTokenAdded)
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
		it.Event = new(RootchainTokenAdded)
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
func (it *RootchainTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootchainTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootchainTokenAdded represents a TokenAdded event raised by the Rootchain contract.
type RootchainTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: e TokenAdded(token address)
func (_Rootchain *RootchainFilterer) FilterTokenAdded(opts *bind.FilterOpts) (*RootchainTokenAddedIterator, error) {

	logs, sub, err := _Rootchain.contract.FilterLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return &RootchainTokenAddedIterator{contract: _Rootchain.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: e TokenAdded(token address)
func (_Rootchain *RootchainFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *RootchainTokenAdded) (event.Subscription, error) {

	logs, sub, err := _Rootchain.contract.WatchLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootchainTokenAdded)
				if err := _Rootchain.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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
