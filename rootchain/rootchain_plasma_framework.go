// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rootchain

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

// PlasmaFrameworkABI is the input ABI used to generate the binding from.
const PlasmaFrameworkABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isChildChainActivated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getNextExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"exitsQueues\",\"outputs\":[{\"internalType\":\"contractPriorityQueue\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"childBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_outputIds\",\"type\":\"bytes32[]\"}],\"name\":\"batchFlagOutputsSpent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addExitQueue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextChildBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deactivateNonReentrant\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"isExitGameSafeToUse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_vaultAddress\",\"type\":\"address\"}],\"name\":\"registerVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextDepositBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txType\",\"type\":\"uint256\"}],\"name\":\"protocols\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_outputIds\",\"type\":\"bytes32[]\"}],\"name\":\"isAnyOutputsSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"exitableAt\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structTxPosLib.TxPos\",\"name\":\"txPos\",\"type\":\"tuple\"},{\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"},{\"internalType\":\"contractIExitProcessor\",\"name\":\"exitProcessor\",\"type\":\"address\"}],\"name\":\"enqueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activateChildChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exitGame\",\"type\":\"address\"}],\"name\":\"exitGameToTxType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHILD_BLOCK_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isOutputSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txType\",\"type\":\"uint256\"}],\"name\":\"exitGames\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_blockRoot\",\"type\":\"bytes32\"}],\"name\":\"submitBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_blockRoot\",\"type\":\"bytes32\"}],\"name\":\"submitDepositBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delegations\",\"outputs\":[{\"internalType\":\"contractIExitProcessor\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"hasExitQueue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_outputId\",\"type\":\"bytes32\"}],\"name\":\"flagOutputSpent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minExitPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"topExitId\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"maxExitsToProcess\",\"type\":\"uint256\"}],\"name\":\"processExits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultAddress\",\"type\":\"address\"}],\"name\":\"vaultToId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_protocol\",\"type\":\"uint8\"}],\"name\":\"registerExitGame\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activateNonReentrant\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minExitPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initialImmuneVaults\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initialImmuneExitGames\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_authority\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_maintainer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"ChildChainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ExitQueueAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"processedNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ProcessedExitsNum\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priority\",\"type\":\"uint256\"}],\"name\":\"ExitQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"exitGameAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"protocol\",\"type\":\"uint8\"}],\"name\":\"ExitGameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"VaultRegistered\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMaintainer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PlasmaFramework is an auto generated Go binding around an Ethereum contract.
type PlasmaFramework struct {
	PlasmaFrameworkCaller     // Read-only binding to the contract
	PlasmaFrameworkTransactor // Write-only binding to the contract
	PlasmaFrameworkFilterer   // Log filterer for contract events
}

// PlasmaFrameworkCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlasmaFrameworkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaFrameworkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlasmaFrameworkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaFrameworkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlasmaFrameworkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaFrameworkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlasmaFrameworkSession struct {
	Contract     *PlasmaFramework  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlasmaFrameworkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlasmaFrameworkCallerSession struct {
	Contract *PlasmaFrameworkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PlasmaFrameworkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlasmaFrameworkTransactorSession struct {
	Contract     *PlasmaFrameworkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PlasmaFrameworkRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlasmaFrameworkRaw struct {
	Contract *PlasmaFramework // Generic contract binding to access the raw methods on
}

// PlasmaFrameworkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlasmaFrameworkCallerRaw struct {
	Contract *PlasmaFrameworkCaller // Generic read-only contract binding to access the raw methods on
}

// PlasmaFrameworkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlasmaFrameworkTransactorRaw struct {
	Contract *PlasmaFrameworkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlasmaFramework creates a new instance of PlasmaFramework, bound to a specific deployed contract.
func NewPlasmaFramework(address common.Address, backend bind.ContractBackend) (*PlasmaFramework, error) {
	contract, err := bindPlasmaFramework(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlasmaFramework{PlasmaFrameworkCaller: PlasmaFrameworkCaller{contract: contract}, PlasmaFrameworkTransactor: PlasmaFrameworkTransactor{contract: contract}, PlasmaFrameworkFilterer: PlasmaFrameworkFilterer{contract: contract}}, nil
}

// NewPlasmaFrameworkCaller creates a new read-only instance of PlasmaFramework, bound to a specific deployed contract.
func NewPlasmaFrameworkCaller(address common.Address, caller bind.ContractCaller) (*PlasmaFrameworkCaller, error) {
	contract, err := bindPlasmaFramework(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaFrameworkCaller{contract: contract}, nil
}

// NewPlasmaFrameworkTransactor creates a new write-only instance of PlasmaFramework, bound to a specific deployed contract.
func NewPlasmaFrameworkTransactor(address common.Address, transactor bind.ContractTransactor) (*PlasmaFrameworkTransactor, error) {
	contract, err := bindPlasmaFramework(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaFrameworkTransactor{contract: contract}, nil
}

// NewPlasmaFrameworkFilterer creates a new log filterer instance of PlasmaFramework, bound to a specific deployed contract.
func NewPlasmaFrameworkFilterer(address common.Address, filterer bind.ContractFilterer) (*PlasmaFrameworkFilterer, error) {
	contract, err := bindPlasmaFramework(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlasmaFrameworkFilterer{contract: contract}, nil
}

// bindPlasmaFramework binds a generic wrapper to an already deployed contract.
func bindPlasmaFramework(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlasmaFrameworkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlasmaFramework *PlasmaFrameworkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlasmaFramework.Contract.PlasmaFrameworkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlasmaFramework *PlasmaFrameworkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.PlasmaFrameworkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlasmaFramework *PlasmaFrameworkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.PlasmaFrameworkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlasmaFramework *PlasmaFrameworkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlasmaFramework.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlasmaFramework *PlasmaFrameworkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlasmaFramework *PlasmaFrameworkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.contract.Transact(opts, method, params...)
}

// TxPosLibTxPos is an auto generated low-level Go binding around an user-defined struct.
type TxPosLibTxPos struct {
	Value *big.Int
}

// CHILDBLOCKINTERVAL is a free data retrieval call binding the contract method 0xa831fa07.
//
// Solidity: function CHILD_BLOCK_INTERVAL() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCaller) CHILDBLOCKINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "CHILD_BLOCK_INTERVAL")
	return *ret0, err
}

// CHILDBLOCKINTERVAL is a free data retrieval call binding the contract method 0xa831fa07.
//
// Solidity: function CHILD_BLOCK_INTERVAL() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkSession) CHILDBLOCKINTERVAL() (*big.Int, error) {
	return _PlasmaFramework.Contract.CHILDBLOCKINTERVAL(&_PlasmaFramework.CallOpts)
}

// CHILDBLOCKINTERVAL is a free data retrieval call binding the contract method 0xa831fa07.
//
// Solidity: function CHILD_BLOCK_INTERVAL() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) CHILDBLOCKINTERVAL() (*big.Int, error) {
	return _PlasmaFramework.Contract.CHILDBLOCKINTERVAL(&_PlasmaFramework.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkSession) Authority() (common.Address, error) {
	return _PlasmaFramework.Contract.Authority(&_PlasmaFramework.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) Authority() (common.Address, error) {
	return _PlasmaFramework.Contract.Authority(&_PlasmaFramework.CallOpts)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 ) constant returns(bytes32 root, uint256 timestamp)
func (_PlasmaFramework *PlasmaFrameworkCaller) Blocks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Root      [32]byte
	Timestamp *big.Int
}, error) {
	ret := new(struct {
		Root      [32]byte
		Timestamp *big.Int
	})
	out := ret
	err := _PlasmaFramework.contract.Call(opts, out, "blocks", arg0)
	return *ret, err
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 ) constant returns(bytes32 root, uint256 timestamp)
func (_PlasmaFramework *PlasmaFrameworkSession) Blocks(arg0 *big.Int) (struct {
	Root      [32]byte
	Timestamp *big.Int
}, error) {
	return _PlasmaFramework.Contract.Blocks(&_PlasmaFramework.CallOpts, arg0)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 ) constant returns(bytes32 root, uint256 timestamp)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) Blocks(arg0 *big.Int) (struct {
	Root      [32]byte
	Timestamp *big.Int
}, error) {
	return _PlasmaFramework.Contract.Blocks(&_PlasmaFramework.CallOpts, arg0)
}

// ChildBlockInterval is a free data retrieval call binding the contract method 0x38a9e0bc.
//
// Solidity: function childBlockInterval() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCaller) ChildBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "childBlockInterval")
	return *ret0, err
}

// ChildBlockInterval is a free data retrieval call binding the contract method 0x38a9e0bc.
//
// Solidity: function childBlockInterval() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkSession) ChildBlockInterval() (*big.Int, error) {
	return _PlasmaFramework.Contract.ChildBlockInterval(&_PlasmaFramework.CallOpts)
}

// ChildBlockInterval is a free data retrieval call binding the contract method 0x38a9e0bc.
//
// Solidity: function childBlockInterval() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) ChildBlockInterval() (*big.Int, error) {
	return _PlasmaFramework.Contract.ChildBlockInterval(&_PlasmaFramework.CallOpts)
}

// Delegations is a free data retrieval call binding the contract method 0xc4336c1c.
//
// Solidity: function delegations(uint256 ) constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkCaller) Delegations(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "delegations", arg0)
	return *ret0, err
}

// Delegations is a free data retrieval call binding the contract method 0xc4336c1c.
//
// Solidity: function delegations(uint256 ) constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkSession) Delegations(arg0 *big.Int) (common.Address, error) {
	return _PlasmaFramework.Contract.Delegations(&_PlasmaFramework.CallOpts, arg0)
}

// Delegations is a free data retrieval call binding the contract method 0xc4336c1c.
//
// Solidity: function delegations(uint256 ) constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) Delegations(arg0 *big.Int) (common.Address, error) {
	return _PlasmaFramework.Contract.Delegations(&_PlasmaFramework.CallOpts, arg0)
}

// ExitGameToTxType is a free data retrieval call binding the contract method 0xa239e5a2.
//
// Solidity: function exitGameToTxType(address _exitGame) constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCaller) ExitGameToTxType(opts *bind.CallOpts, _exitGame common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "exitGameToTxType", _exitGame)
	return *ret0, err
}

// ExitGameToTxType is a free data retrieval call binding the contract method 0xa239e5a2.
//
// Solidity: function exitGameToTxType(address _exitGame) constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkSession) ExitGameToTxType(_exitGame common.Address) (*big.Int, error) {
	return _PlasmaFramework.Contract.ExitGameToTxType(&_PlasmaFramework.CallOpts, _exitGame)
}

// ExitGameToTxType is a free data retrieval call binding the contract method 0xa239e5a2.
//
// Solidity: function exitGameToTxType(address _exitGame) constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) ExitGameToTxType(_exitGame common.Address) (*big.Int, error) {
	return _PlasmaFramework.Contract.ExitGameToTxType(&_PlasmaFramework.CallOpts, _exitGame)
}

// ExitGames is a free data retrieval call binding the contract method 0xaf079764.
//
// Solidity: function exitGames(uint256 _txType) constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkCaller) ExitGames(opts *bind.CallOpts, _txType *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "exitGames", _txType)
	return *ret0, err
}

// ExitGames is a free data retrieval call binding the contract method 0xaf079764.
//
// Solidity: function exitGames(uint256 _txType) constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkSession) ExitGames(_txType *big.Int) (common.Address, error) {
	return _PlasmaFramework.Contract.ExitGames(&_PlasmaFramework.CallOpts, _txType)
}

// ExitGames is a free data retrieval call binding the contract method 0xaf079764.
//
// Solidity: function exitGames(uint256 _txType) constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) ExitGames(_txType *big.Int) (common.Address, error) {
	return _PlasmaFramework.Contract.ExitGames(&_PlasmaFramework.CallOpts, _txType)
}

// ExitsQueues is a free data retrieval call binding the contract method 0x3701b5ca.
//
// Solidity: function exitsQueues(bytes32 ) constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkCaller) ExitsQueues(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "exitsQueues", arg0)
	return *ret0, err
}

// ExitsQueues is a free data retrieval call binding the contract method 0x3701b5ca.
//
// Solidity: function exitsQueues(bytes32 ) constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkSession) ExitsQueues(arg0 [32]byte) (common.Address, error) {
	return _PlasmaFramework.Contract.ExitsQueues(&_PlasmaFramework.CallOpts, arg0)
}

// ExitsQueues is a free data retrieval call binding the contract method 0x3701b5ca.
//
// Solidity: function exitsQueues(bytes32 ) constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) ExitsQueues(arg0 [32]byte) (common.Address, error) {
	return _PlasmaFramework.Contract.ExitsQueues(&_PlasmaFramework.CallOpts, arg0)
}

// GetMaintainer is a free data retrieval call binding the contract method 0x4b0a72bc.
//
// Solidity: function getMaintainer() constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkCaller) GetMaintainer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "getMaintainer")
	return *ret0, err
}

// GetMaintainer is a free data retrieval call binding the contract method 0x4b0a72bc.
//
// Solidity: function getMaintainer() constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkSession) GetMaintainer() (common.Address, error) {
	return _PlasmaFramework.Contract.GetMaintainer(&_PlasmaFramework.CallOpts)
}

// GetMaintainer is a free data retrieval call binding the contract method 0x4b0a72bc.
//
// Solidity: function getMaintainer() constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) GetMaintainer() (common.Address, error) {
	return _PlasmaFramework.Contract.GetMaintainer(&_PlasmaFramework.CallOpts)
}

// GetNextExit is a free data retrieval call binding the contract method 0x296d5427.
//
// Solidity: function getNextExit(uint256 vaultId, address token) constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCaller) GetNextExit(opts *bind.CallOpts, vaultId *big.Int, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "getNextExit", vaultId, token)
	return *ret0, err
}

// GetNextExit is a free data retrieval call binding the contract method 0x296d5427.
//
// Solidity: function getNextExit(uint256 vaultId, address token) constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkSession) GetNextExit(vaultId *big.Int, token common.Address) (*big.Int, error) {
	return _PlasmaFramework.Contract.GetNextExit(&_PlasmaFramework.CallOpts, vaultId, token)
}

// GetNextExit is a free data retrieval call binding the contract method 0x296d5427.
//
// Solidity: function getNextExit(uint256 vaultId, address token) constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) GetNextExit(vaultId *big.Int, token common.Address) (*big.Int, error) {
	return _PlasmaFramework.Contract.GetNextExit(&_PlasmaFramework.CallOpts, vaultId, token)
}

// HasExitQueue is a free data retrieval call binding the contract method 0xc77b1ed6.
//
// Solidity: function hasExitQueue(uint256 vaultId, address token) constant returns(bool)
func (_PlasmaFramework *PlasmaFrameworkCaller) HasExitQueue(opts *bind.CallOpts, vaultId *big.Int, token common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "hasExitQueue", vaultId, token)
	return *ret0, err
}

// HasExitQueue is a free data retrieval call binding the contract method 0xc77b1ed6.
//
// Solidity: function hasExitQueue(uint256 vaultId, address token) constant returns(bool)
func (_PlasmaFramework *PlasmaFrameworkSession) HasExitQueue(vaultId *big.Int, token common.Address) (bool, error) {
	return _PlasmaFramework.Contract.HasExitQueue(&_PlasmaFramework.CallOpts, vaultId, token)
}

// HasExitQueue is a free data retrieval call binding the contract method 0xc77b1ed6.
//
// Solidity: function hasExitQueue(uint256 vaultId, address token) constant returns(bool)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) HasExitQueue(vaultId *big.Int, token common.Address) (bool, error) {
	return _PlasmaFramework.Contract.HasExitQueue(&_PlasmaFramework.CallOpts, vaultId, token)
}

// IsAnyOutputsSpent is a free data retrieval call binding the contract method 0x8e7e3890.
//
// Solidity: function isAnyOutputsSpent(bytes32[] _outputIds) constant returns(bool)
func (_PlasmaFramework *PlasmaFrameworkCaller) IsAnyOutputsSpent(opts *bind.CallOpts, _outputIds [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "isAnyOutputsSpent", _outputIds)
	return *ret0, err
}

// IsAnyOutputsSpent is a free data retrieval call binding the contract method 0x8e7e3890.
//
// Solidity: function isAnyOutputsSpent(bytes32[] _outputIds) constant returns(bool)
func (_PlasmaFramework *PlasmaFrameworkSession) IsAnyOutputsSpent(_outputIds [][32]byte) (bool, error) {
	return _PlasmaFramework.Contract.IsAnyOutputsSpent(&_PlasmaFramework.CallOpts, _outputIds)
}

// IsAnyOutputsSpent is a free data retrieval call binding the contract method 0x8e7e3890.
//
// Solidity: function isAnyOutputsSpent(bytes32[] _outputIds) constant returns(bool)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) IsAnyOutputsSpent(_outputIds [][32]byte) (bool, error) {
	return _PlasmaFramework.Contract.IsAnyOutputsSpent(&_PlasmaFramework.CallOpts, _outputIds)
}

// IsChildChainActivated is a free data retrieval call binding the contract method 0x0e71ee02.
//
// Solidity: function isChildChainActivated() constant returns(bool)
func (_PlasmaFramework *PlasmaFrameworkCaller) IsChildChainActivated(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "isChildChainActivated")
	return *ret0, err
}

// IsChildChainActivated is a free data retrieval call binding the contract method 0x0e71ee02.
//
// Solidity: function isChildChainActivated() constant returns(bool)
func (_PlasmaFramework *PlasmaFrameworkSession) IsChildChainActivated() (bool, error) {
	return _PlasmaFramework.Contract.IsChildChainActivated(&_PlasmaFramework.CallOpts)
}

// IsChildChainActivated is a free data retrieval call binding the contract method 0x0e71ee02.
//
// Solidity: function isChildChainActivated() constant returns(bool)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) IsChildChainActivated() (bool, error) {
	return _PlasmaFramework.Contract.IsChildChainActivated(&_PlasmaFramework.CallOpts)
}

// IsExitGameSafeToUse is a free data retrieval call binding the contract method 0x5a102dae.
//
// Solidity: function isExitGameSafeToUse(address _contract) constant returns(bool)
func (_PlasmaFramework *PlasmaFrameworkCaller) IsExitGameSafeToUse(opts *bind.CallOpts, _contract common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "isExitGameSafeToUse", _contract)
	return *ret0, err
}

// IsExitGameSafeToUse is a free data retrieval call binding the contract method 0x5a102dae.
//
// Solidity: function isExitGameSafeToUse(address _contract) constant returns(bool)
func (_PlasmaFramework *PlasmaFrameworkSession) IsExitGameSafeToUse(_contract common.Address) (bool, error) {
	return _PlasmaFramework.Contract.IsExitGameSafeToUse(&_PlasmaFramework.CallOpts, _contract)
}

// IsExitGameSafeToUse is a free data retrieval call binding the contract method 0x5a102dae.
//
// Solidity: function isExitGameSafeToUse(address _contract) constant returns(bool)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) IsExitGameSafeToUse(_contract common.Address) (bool, error) {
	return _PlasmaFramework.Contract.IsExitGameSafeToUse(&_PlasmaFramework.CallOpts, _contract)
}

// IsOutputSpent is a free data retrieval call binding the contract method 0xab84a94b.
//
// Solidity: function isOutputSpent(bytes32 ) constant returns(bool)
func (_PlasmaFramework *PlasmaFrameworkCaller) IsOutputSpent(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "isOutputSpent", arg0)
	return *ret0, err
}

// IsOutputSpent is a free data retrieval call binding the contract method 0xab84a94b.
//
// Solidity: function isOutputSpent(bytes32 ) constant returns(bool)
func (_PlasmaFramework *PlasmaFrameworkSession) IsOutputSpent(arg0 [32]byte) (bool, error) {
	return _PlasmaFramework.Contract.IsOutputSpent(&_PlasmaFramework.CallOpts, arg0)
}

// IsOutputSpent is a free data retrieval call binding the contract method 0xab84a94b.
//
// Solidity: function isOutputSpent(bytes32 ) constant returns(bool)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) IsOutputSpent(arg0 [32]byte) (bool, error) {
	return _PlasmaFramework.Contract.IsOutputSpent(&_PlasmaFramework.CallOpts, arg0)
}

// MinExitPeriod is a free data retrieval call binding the contract method 0xd4a2b4ef.
//
// Solidity: function minExitPeriod() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCaller) MinExitPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "minExitPeriod")
	return *ret0, err
}

// MinExitPeriod is a free data retrieval call binding the contract method 0xd4a2b4ef.
//
// Solidity: function minExitPeriod() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkSession) MinExitPeriod() (*big.Int, error) {
	return _PlasmaFramework.Contract.MinExitPeriod(&_PlasmaFramework.CallOpts)
}

// MinExitPeriod is a free data retrieval call binding the contract method 0xd4a2b4ef.
//
// Solidity: function minExitPeriod() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) MinExitPeriod() (*big.Int, error) {
	return _PlasmaFramework.Contract.MinExitPeriod(&_PlasmaFramework.CallOpts)
}

// NextChildBlock is a free data retrieval call binding the contract method 0x4ca8714f.
//
// Solidity: function nextChildBlock() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCaller) NextChildBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "nextChildBlock")
	return *ret0, err
}

// NextChildBlock is a free data retrieval call binding the contract method 0x4ca8714f.
//
// Solidity: function nextChildBlock() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkSession) NextChildBlock() (*big.Int, error) {
	return _PlasmaFramework.Contract.NextChildBlock(&_PlasmaFramework.CallOpts)
}

// NextChildBlock is a free data retrieval call binding the contract method 0x4ca8714f.
//
// Solidity: function nextChildBlock() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) NextChildBlock() (*big.Int, error) {
	return _PlasmaFramework.Contract.NextChildBlock(&_PlasmaFramework.CallOpts)
}

// NextDeposit is a free data retrieval call binding the contract method 0xa8cabcd5.
//
// Solidity: function nextDeposit() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCaller) NextDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "nextDeposit")
	return *ret0, err
}

// NextDeposit is a free data retrieval call binding the contract method 0xa8cabcd5.
//
// Solidity: function nextDeposit() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkSession) NextDeposit() (*big.Int, error) {
	return _PlasmaFramework.Contract.NextDeposit(&_PlasmaFramework.CallOpts)
}

// NextDeposit is a free data retrieval call binding the contract method 0xa8cabcd5.
//
// Solidity: function nextDeposit() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) NextDeposit() (*big.Int, error) {
	return _PlasmaFramework.Contract.NextDeposit(&_PlasmaFramework.CallOpts)
}

// NextDepositBlock is a free data retrieval call binding the contract method 0x8701fc5d.
//
// Solidity: function nextDepositBlock() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCaller) NextDepositBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "nextDepositBlock")
	return *ret0, err
}

// NextDepositBlock is a free data retrieval call binding the contract method 0x8701fc5d.
//
// Solidity: function nextDepositBlock() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkSession) NextDepositBlock() (*big.Int, error) {
	return _PlasmaFramework.Contract.NextDepositBlock(&_PlasmaFramework.CallOpts)
}

// NextDepositBlock is a free data retrieval call binding the contract method 0x8701fc5d.
//
// Solidity: function nextDepositBlock() constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) NextDepositBlock() (*big.Int, error) {
	return _PlasmaFramework.Contract.NextDepositBlock(&_PlasmaFramework.CallOpts)
}

// Protocols is a free data retrieval call binding the contract method 0x8c396220.
//
// Solidity: function protocols(uint256 _txType) constant returns(uint8)
func (_PlasmaFramework *PlasmaFrameworkCaller) Protocols(opts *bind.CallOpts, _txType *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "protocols", _txType)
	return *ret0, err
}

// Protocols is a free data retrieval call binding the contract method 0x8c396220.
//
// Solidity: function protocols(uint256 _txType) constant returns(uint8)
func (_PlasmaFramework *PlasmaFrameworkSession) Protocols(_txType *big.Int) (uint8, error) {
	return _PlasmaFramework.Contract.Protocols(&_PlasmaFramework.CallOpts, _txType)
}

// Protocols is a free data retrieval call binding the contract method 0x8c396220.
//
// Solidity: function protocols(uint256 _txType) constant returns(uint8)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) Protocols(_txType *big.Int) (uint8, error) {
	return _PlasmaFramework.Contract.Protocols(&_PlasmaFramework.CallOpts, _txType)
}

// VaultToId is a free data retrieval call binding the contract method 0xdfb494f0.
//
// Solidity: function vaultToId(address _vaultAddress) constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCaller) VaultToId(opts *bind.CallOpts, _vaultAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "vaultToId", _vaultAddress)
	return *ret0, err
}

// VaultToId is a free data retrieval call binding the contract method 0xdfb494f0.
//
// Solidity: function vaultToId(address _vaultAddress) constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkSession) VaultToId(_vaultAddress common.Address) (*big.Int, error) {
	return _PlasmaFramework.Contract.VaultToId(&_PlasmaFramework.CallOpts, _vaultAddress)
}

// VaultToId is a free data retrieval call binding the contract method 0xdfb494f0.
//
// Solidity: function vaultToId(address _vaultAddress) constant returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) VaultToId(_vaultAddress common.Address) (*big.Int, error) {
	return _PlasmaFramework.Contract.VaultToId(&_PlasmaFramework.CallOpts, _vaultAddress)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 _vaultId) constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkCaller) Vaults(opts *bind.CallOpts, _vaultId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlasmaFramework.contract.Call(opts, out, "vaults", _vaultId)
	return *ret0, err
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 _vaultId) constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkSession) Vaults(_vaultId *big.Int) (common.Address, error) {
	return _PlasmaFramework.Contract.Vaults(&_PlasmaFramework.CallOpts, _vaultId)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 _vaultId) constant returns(address)
func (_PlasmaFramework *PlasmaFrameworkCallerSession) Vaults(_vaultId *big.Int) (common.Address, error) {
	return _PlasmaFramework.Contract.Vaults(&_PlasmaFramework.CallOpts, _vaultId)
}

// ActivateChildChain is a paid mutator transaction binding the contract method 0xa11dcc34.
//
// Solidity: function activateChildChain() returns()
func (_PlasmaFramework *PlasmaFrameworkTransactor) ActivateChildChain(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaFramework.contract.Transact(opts, "activateChildChain")
}

// ActivateChildChain is a paid mutator transaction binding the contract method 0xa11dcc34.
//
// Solidity: function activateChildChain() returns()
func (_PlasmaFramework *PlasmaFrameworkSession) ActivateChildChain() (*types.Transaction, error) {
	return _PlasmaFramework.Contract.ActivateChildChain(&_PlasmaFramework.TransactOpts)
}

// ActivateChildChain is a paid mutator transaction binding the contract method 0xa11dcc34.
//
// Solidity: function activateChildChain() returns()
func (_PlasmaFramework *PlasmaFrameworkTransactorSession) ActivateChildChain() (*types.Transaction, error) {
	return _PlasmaFramework.Contract.ActivateChildChain(&_PlasmaFramework.TransactOpts)
}

// ActivateNonReentrant is a paid mutator transaction binding the contract method 0xff27b189.
//
// Solidity: function activateNonReentrant() returns()
func (_PlasmaFramework *PlasmaFrameworkTransactor) ActivateNonReentrant(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaFramework.contract.Transact(opts, "activateNonReentrant")
}

// ActivateNonReentrant is a paid mutator transaction binding the contract method 0xff27b189.
//
// Solidity: function activateNonReentrant() returns()
func (_PlasmaFramework *PlasmaFrameworkSession) ActivateNonReentrant() (*types.Transaction, error) {
	return _PlasmaFramework.Contract.ActivateNonReentrant(&_PlasmaFramework.TransactOpts)
}

// ActivateNonReentrant is a paid mutator transaction binding the contract method 0xff27b189.
//
// Solidity: function activateNonReentrant() returns()
func (_PlasmaFramework *PlasmaFrameworkTransactorSession) ActivateNonReentrant() (*types.Transaction, error) {
	return _PlasmaFramework.Contract.ActivateNonReentrant(&_PlasmaFramework.TransactOpts)
}

// AddExitQueue is a paid mutator transaction binding the contract method 0x45f20935.
//
// Solidity: function addExitQueue(uint256 vaultId, address token) returns()
func (_PlasmaFramework *PlasmaFrameworkTransactor) AddExitQueue(opts *bind.TransactOpts, vaultId *big.Int, token common.Address) (*types.Transaction, error) {
	return _PlasmaFramework.contract.Transact(opts, "addExitQueue", vaultId, token)
}

// AddExitQueue is a paid mutator transaction binding the contract method 0x45f20935.
//
// Solidity: function addExitQueue(uint256 vaultId, address token) returns()
func (_PlasmaFramework *PlasmaFrameworkSession) AddExitQueue(vaultId *big.Int, token common.Address) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.AddExitQueue(&_PlasmaFramework.TransactOpts, vaultId, token)
}

// AddExitQueue is a paid mutator transaction binding the contract method 0x45f20935.
//
// Solidity: function addExitQueue(uint256 vaultId, address token) returns()
func (_PlasmaFramework *PlasmaFrameworkTransactorSession) AddExitQueue(vaultId *big.Int, token common.Address) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.AddExitQueue(&_PlasmaFramework.TransactOpts, vaultId, token)
}

// BatchFlagOutputsSpent is a paid mutator transaction binding the contract method 0x42383d4a.
//
// Solidity: function batchFlagOutputsSpent(bytes32[] _outputIds) returns()
func (_PlasmaFramework *PlasmaFrameworkTransactor) BatchFlagOutputsSpent(opts *bind.TransactOpts, _outputIds [][32]byte) (*types.Transaction, error) {
	return _PlasmaFramework.contract.Transact(opts, "batchFlagOutputsSpent", _outputIds)
}

// BatchFlagOutputsSpent is a paid mutator transaction binding the contract method 0x42383d4a.
//
// Solidity: function batchFlagOutputsSpent(bytes32[] _outputIds) returns()
func (_PlasmaFramework *PlasmaFrameworkSession) BatchFlagOutputsSpent(_outputIds [][32]byte) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.BatchFlagOutputsSpent(&_PlasmaFramework.TransactOpts, _outputIds)
}

// BatchFlagOutputsSpent is a paid mutator transaction binding the contract method 0x42383d4a.
//
// Solidity: function batchFlagOutputsSpent(bytes32[] _outputIds) returns()
func (_PlasmaFramework *PlasmaFrameworkTransactorSession) BatchFlagOutputsSpent(_outputIds [][32]byte) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.BatchFlagOutputsSpent(&_PlasmaFramework.TransactOpts, _outputIds)
}

// DeactivateNonReentrant is a paid mutator transaction binding the contract method 0x4e0f7436.
//
// Solidity: function deactivateNonReentrant() returns()
func (_PlasmaFramework *PlasmaFrameworkTransactor) DeactivateNonReentrant(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaFramework.contract.Transact(opts, "deactivateNonReentrant")
}

// DeactivateNonReentrant is a paid mutator transaction binding the contract method 0x4e0f7436.
//
// Solidity: function deactivateNonReentrant() returns()
func (_PlasmaFramework *PlasmaFrameworkSession) DeactivateNonReentrant() (*types.Transaction, error) {
	return _PlasmaFramework.Contract.DeactivateNonReentrant(&_PlasmaFramework.TransactOpts)
}

// DeactivateNonReentrant is a paid mutator transaction binding the contract method 0x4e0f7436.
//
// Solidity: function deactivateNonReentrant() returns()
func (_PlasmaFramework *PlasmaFrameworkTransactorSession) DeactivateNonReentrant() (*types.Transaction, error) {
	return _PlasmaFramework.Contract.DeactivateNonReentrant(&_PlasmaFramework.TransactOpts)
}

// Enqueue is a paid mutator transaction binding the contract method 0x9524a123.
//
// Solidity: function enqueue(uint256 vaultId, address token, uint64 exitableAt, TxPosLibTxPos txPos, uint160 exitId, address exitProcessor) returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkTransactor) Enqueue(opts *bind.TransactOpts, vaultId *big.Int, token common.Address, exitableAt uint64, txPos TxPosLibTxPos, exitId *big.Int, exitProcessor common.Address) (*types.Transaction, error) {
	return _PlasmaFramework.contract.Transact(opts, "enqueue", vaultId, token, exitableAt, txPos, exitId, exitProcessor)
}

// Enqueue is a paid mutator transaction binding the contract method 0x9524a123.
//
// Solidity: function enqueue(uint256 vaultId, address token, uint64 exitableAt, TxPosLibTxPos txPos, uint160 exitId, address exitProcessor) returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkSession) Enqueue(vaultId *big.Int, token common.Address, exitableAt uint64, txPos TxPosLibTxPos, exitId *big.Int, exitProcessor common.Address) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.Enqueue(&_PlasmaFramework.TransactOpts, vaultId, token, exitableAt, txPos, exitId, exitProcessor)
}

// Enqueue is a paid mutator transaction binding the contract method 0x9524a123.
//
// Solidity: function enqueue(uint256 vaultId, address token, uint64 exitableAt, TxPosLibTxPos txPos, uint160 exitId, address exitProcessor) returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkTransactorSession) Enqueue(vaultId *big.Int, token common.Address, exitableAt uint64, txPos TxPosLibTxPos, exitId *big.Int, exitProcessor common.Address) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.Enqueue(&_PlasmaFramework.TransactOpts, vaultId, token, exitableAt, txPos, exitId, exitProcessor)
}

// FlagOutputSpent is a paid mutator transaction binding the contract method 0xd3329542.
//
// Solidity: function flagOutputSpent(bytes32 _outputId) returns()
func (_PlasmaFramework *PlasmaFrameworkTransactor) FlagOutputSpent(opts *bind.TransactOpts, _outputId [32]byte) (*types.Transaction, error) {
	return _PlasmaFramework.contract.Transact(opts, "flagOutputSpent", _outputId)
}

// FlagOutputSpent is a paid mutator transaction binding the contract method 0xd3329542.
//
// Solidity: function flagOutputSpent(bytes32 _outputId) returns()
func (_PlasmaFramework *PlasmaFrameworkSession) FlagOutputSpent(_outputId [32]byte) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.FlagOutputSpent(&_PlasmaFramework.TransactOpts, _outputId)
}

// FlagOutputSpent is a paid mutator transaction binding the contract method 0xd3329542.
//
// Solidity: function flagOutputSpent(bytes32 _outputId) returns()
func (_PlasmaFramework *PlasmaFrameworkTransactorSession) FlagOutputSpent(_outputId [32]byte) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.FlagOutputSpent(&_PlasmaFramework.TransactOpts, _outputId)
}

// ProcessExits is a paid mutator transaction binding the contract method 0xd9aa4b24.
//
// Solidity: function processExits(uint256 vaultId, address token, uint160 topExitId, uint256 maxExitsToProcess) returns()
func (_PlasmaFramework *PlasmaFrameworkTransactor) ProcessExits(opts *bind.TransactOpts, vaultId *big.Int, token common.Address, topExitId *big.Int, maxExitsToProcess *big.Int) (*types.Transaction, error) {
	return _PlasmaFramework.contract.Transact(opts, "processExits", vaultId, token, topExitId, maxExitsToProcess)
}

// ProcessExits is a paid mutator transaction binding the contract method 0xd9aa4b24.
//
// Solidity: function processExits(uint256 vaultId, address token, uint160 topExitId, uint256 maxExitsToProcess) returns()
func (_PlasmaFramework *PlasmaFrameworkSession) ProcessExits(vaultId *big.Int, token common.Address, topExitId *big.Int, maxExitsToProcess *big.Int) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.ProcessExits(&_PlasmaFramework.TransactOpts, vaultId, token, topExitId, maxExitsToProcess)
}

// ProcessExits is a paid mutator transaction binding the contract method 0xd9aa4b24.
//
// Solidity: function processExits(uint256 vaultId, address token, uint160 topExitId, uint256 maxExitsToProcess) returns()
func (_PlasmaFramework *PlasmaFrameworkTransactorSession) ProcessExits(vaultId *big.Int, token common.Address, topExitId *big.Int, maxExitsToProcess *big.Int) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.ProcessExits(&_PlasmaFramework.TransactOpts, vaultId, token, topExitId, maxExitsToProcess)
}

// RegisterExitGame is a paid mutator transaction binding the contract method 0xf2e208e7.
//
// Solidity: function registerExitGame(uint256 _txType, address _contract, uint8 _protocol) returns()
func (_PlasmaFramework *PlasmaFrameworkTransactor) RegisterExitGame(opts *bind.TransactOpts, _txType *big.Int, _contract common.Address, _protocol uint8) (*types.Transaction, error) {
	return _PlasmaFramework.contract.Transact(opts, "registerExitGame", _txType, _contract, _protocol)
}

// RegisterExitGame is a paid mutator transaction binding the contract method 0xf2e208e7.
//
// Solidity: function registerExitGame(uint256 _txType, address _contract, uint8 _protocol) returns()
func (_PlasmaFramework *PlasmaFrameworkSession) RegisterExitGame(_txType *big.Int, _contract common.Address, _protocol uint8) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.RegisterExitGame(&_PlasmaFramework.TransactOpts, _txType, _contract, _protocol)
}

// RegisterExitGame is a paid mutator transaction binding the contract method 0xf2e208e7.
//
// Solidity: function registerExitGame(uint256 _txType, address _contract, uint8 _protocol) returns()
func (_PlasmaFramework *PlasmaFrameworkTransactorSession) RegisterExitGame(_txType *big.Int, _contract common.Address, _protocol uint8) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.RegisterExitGame(&_PlasmaFramework.TransactOpts, _txType, _contract, _protocol)
}

// RegisterVault is a paid mutator transaction binding the contract method 0x6a51fd63.
//
// Solidity: function registerVault(uint256 _vaultId, address _vaultAddress) returns()
func (_PlasmaFramework *PlasmaFrameworkTransactor) RegisterVault(opts *bind.TransactOpts, _vaultId *big.Int, _vaultAddress common.Address) (*types.Transaction, error) {
	return _PlasmaFramework.contract.Transact(opts, "registerVault", _vaultId, _vaultAddress)
}

// RegisterVault is a paid mutator transaction binding the contract method 0x6a51fd63.
//
// Solidity: function registerVault(uint256 _vaultId, address _vaultAddress) returns()
func (_PlasmaFramework *PlasmaFrameworkSession) RegisterVault(_vaultId *big.Int, _vaultAddress common.Address) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.RegisterVault(&_PlasmaFramework.TransactOpts, _vaultId, _vaultAddress)
}

// RegisterVault is a paid mutator transaction binding the contract method 0x6a51fd63.
//
// Solidity: function registerVault(uint256 _vaultId, address _vaultAddress) returns()
func (_PlasmaFramework *PlasmaFrameworkTransactorSession) RegisterVault(_vaultId *big.Int, _vaultAddress common.Address) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.RegisterVault(&_PlasmaFramework.TransactOpts, _vaultId, _vaultAddress)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(bytes32 _blockRoot) returns()
func (_PlasmaFramework *PlasmaFrameworkTransactor) SubmitBlock(opts *bind.TransactOpts, _blockRoot [32]byte) (*types.Transaction, error) {
	return _PlasmaFramework.contract.Transact(opts, "submitBlock", _blockRoot)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(bytes32 _blockRoot) returns()
func (_PlasmaFramework *PlasmaFrameworkSession) SubmitBlock(_blockRoot [32]byte) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.SubmitBlock(&_PlasmaFramework.TransactOpts, _blockRoot)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(bytes32 _blockRoot) returns()
func (_PlasmaFramework *PlasmaFrameworkTransactorSession) SubmitBlock(_blockRoot [32]byte) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.SubmitBlock(&_PlasmaFramework.TransactOpts, _blockRoot)
}

// SubmitDepositBlock is a paid mutator transaction binding the contract method 0xbe5ac698.
//
// Solidity: function submitDepositBlock(bytes32 _blockRoot) returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkTransactor) SubmitDepositBlock(opts *bind.TransactOpts, _blockRoot [32]byte) (*types.Transaction, error) {
	return _PlasmaFramework.contract.Transact(opts, "submitDepositBlock", _blockRoot)
}

// SubmitDepositBlock is a paid mutator transaction binding the contract method 0xbe5ac698.
//
// Solidity: function submitDepositBlock(bytes32 _blockRoot) returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkSession) SubmitDepositBlock(_blockRoot [32]byte) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.SubmitDepositBlock(&_PlasmaFramework.TransactOpts, _blockRoot)
}

// SubmitDepositBlock is a paid mutator transaction binding the contract method 0xbe5ac698.
//
// Solidity: function submitDepositBlock(bytes32 _blockRoot) returns(uint256)
func (_PlasmaFramework *PlasmaFrameworkTransactorSession) SubmitDepositBlock(_blockRoot [32]byte) (*types.Transaction, error) {
	return _PlasmaFramework.Contract.SubmitDepositBlock(&_PlasmaFramework.TransactOpts, _blockRoot)
}

// PlasmaFrameworkBlockSubmittedIterator is returned from FilterBlockSubmitted and is used to iterate over the raw logs and unpacked data for BlockSubmitted events raised by the PlasmaFramework contract.
type PlasmaFrameworkBlockSubmittedIterator struct {
	Event *PlasmaFrameworkBlockSubmitted // Event containing the contract specifics and raw log

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
func (it *PlasmaFrameworkBlockSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaFrameworkBlockSubmitted)
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
		it.Event = new(PlasmaFrameworkBlockSubmitted)
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
func (it *PlasmaFrameworkBlockSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaFrameworkBlockSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaFrameworkBlockSubmitted represents a BlockSubmitted event raised by the PlasmaFramework contract.
type PlasmaFrameworkBlockSubmitted struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockSubmitted is a free log retrieval operation binding the contract event 0x5a978f4723b249ccf79cd7a658a8601ce1ff8b89fc770251a6be35216351ce32.
//
// Solidity: event BlockSubmitted(uint256 blockNumber)
func (_PlasmaFramework *PlasmaFrameworkFilterer) FilterBlockSubmitted(opts *bind.FilterOpts) (*PlasmaFrameworkBlockSubmittedIterator, error) {

	logs, sub, err := _PlasmaFramework.contract.FilterLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return &PlasmaFrameworkBlockSubmittedIterator{contract: _PlasmaFramework.contract, event: "BlockSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSubmitted is a free log subscription operation binding the contract event 0x5a978f4723b249ccf79cd7a658a8601ce1ff8b89fc770251a6be35216351ce32.
//
// Solidity: event BlockSubmitted(uint256 blockNumber)
func (_PlasmaFramework *PlasmaFrameworkFilterer) WatchBlockSubmitted(opts *bind.WatchOpts, sink chan<- *PlasmaFrameworkBlockSubmitted) (event.Subscription, error) {

	logs, sub, err := _PlasmaFramework.contract.WatchLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaFrameworkBlockSubmitted)
				if err := _PlasmaFramework.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
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

// ParseBlockSubmitted is a log parse operation binding the contract event 0x5a978f4723b249ccf79cd7a658a8601ce1ff8b89fc770251a6be35216351ce32.
//
// Solidity: event BlockSubmitted(uint256 blockNumber)
func (_PlasmaFramework *PlasmaFrameworkFilterer) ParseBlockSubmitted(log types.Log) (*PlasmaFrameworkBlockSubmitted, error) {
	event := new(PlasmaFrameworkBlockSubmitted)
	if err := _PlasmaFramework.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PlasmaFrameworkChildChainActivatedIterator is returned from FilterChildChainActivated and is used to iterate over the raw logs and unpacked data for ChildChainActivated events raised by the PlasmaFramework contract.
type PlasmaFrameworkChildChainActivatedIterator struct {
	Event *PlasmaFrameworkChildChainActivated // Event containing the contract specifics and raw log

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
func (it *PlasmaFrameworkChildChainActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaFrameworkChildChainActivated)
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
		it.Event = new(PlasmaFrameworkChildChainActivated)
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
func (it *PlasmaFrameworkChildChainActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaFrameworkChildChainActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaFrameworkChildChainActivated represents a ChildChainActivated event raised by the PlasmaFramework contract.
type PlasmaFrameworkChildChainActivated struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChildChainActivated is a free log retrieval operation binding the contract event 0xb8421a1acb5f1e701a4f11ecaad76fa438b3947e8dfd6960b6086130e68e0aed.
//
// Solidity: event ChildChainActivated(address authority)
func (_PlasmaFramework *PlasmaFrameworkFilterer) FilterChildChainActivated(opts *bind.FilterOpts) (*PlasmaFrameworkChildChainActivatedIterator, error) {

	logs, sub, err := _PlasmaFramework.contract.FilterLogs(opts, "ChildChainActivated")
	if err != nil {
		return nil, err
	}
	return &PlasmaFrameworkChildChainActivatedIterator{contract: _PlasmaFramework.contract, event: "ChildChainActivated", logs: logs, sub: sub}, nil
}

// WatchChildChainActivated is a free log subscription operation binding the contract event 0xb8421a1acb5f1e701a4f11ecaad76fa438b3947e8dfd6960b6086130e68e0aed.
//
// Solidity: event ChildChainActivated(address authority)
func (_PlasmaFramework *PlasmaFrameworkFilterer) WatchChildChainActivated(opts *bind.WatchOpts, sink chan<- *PlasmaFrameworkChildChainActivated) (event.Subscription, error) {

	logs, sub, err := _PlasmaFramework.contract.WatchLogs(opts, "ChildChainActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaFrameworkChildChainActivated)
				if err := _PlasmaFramework.contract.UnpackLog(event, "ChildChainActivated", log); err != nil {
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

// ParseChildChainActivated is a log parse operation binding the contract event 0xb8421a1acb5f1e701a4f11ecaad76fa438b3947e8dfd6960b6086130e68e0aed.
//
// Solidity: event ChildChainActivated(address authority)
func (_PlasmaFramework *PlasmaFrameworkFilterer) ParseChildChainActivated(log types.Log) (*PlasmaFrameworkChildChainActivated, error) {
	event := new(PlasmaFrameworkChildChainActivated)
	if err := _PlasmaFramework.contract.UnpackLog(event, "ChildChainActivated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PlasmaFrameworkExitGameRegisteredIterator is returned from FilterExitGameRegistered and is used to iterate over the raw logs and unpacked data for ExitGameRegistered events raised by the PlasmaFramework contract.
type PlasmaFrameworkExitGameRegisteredIterator struct {
	Event *PlasmaFrameworkExitGameRegistered // Event containing the contract specifics and raw log

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
func (it *PlasmaFrameworkExitGameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaFrameworkExitGameRegistered)
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
		it.Event = new(PlasmaFrameworkExitGameRegistered)
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
func (it *PlasmaFrameworkExitGameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaFrameworkExitGameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaFrameworkExitGameRegistered represents a ExitGameRegistered event raised by the PlasmaFramework contract.
type PlasmaFrameworkExitGameRegistered struct {
	TxType          *big.Int
	ExitGameAddress common.Address
	Protocol        uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExitGameRegistered is a free log retrieval operation binding the contract event 0x0ceb02f7474c01f5959171cead2f5fb2a44d9c16fe480725a2110b601b19db82.
//
// Solidity: event ExitGameRegistered(uint256 txType, address exitGameAddress, uint8 protocol)
func (_PlasmaFramework *PlasmaFrameworkFilterer) FilterExitGameRegistered(opts *bind.FilterOpts) (*PlasmaFrameworkExitGameRegisteredIterator, error) {

	logs, sub, err := _PlasmaFramework.contract.FilterLogs(opts, "ExitGameRegistered")
	if err != nil {
		return nil, err
	}
	return &PlasmaFrameworkExitGameRegisteredIterator{contract: _PlasmaFramework.contract, event: "ExitGameRegistered", logs: logs, sub: sub}, nil
}

// WatchExitGameRegistered is a free log subscription operation binding the contract event 0x0ceb02f7474c01f5959171cead2f5fb2a44d9c16fe480725a2110b601b19db82.
//
// Solidity: event ExitGameRegistered(uint256 txType, address exitGameAddress, uint8 protocol)
func (_PlasmaFramework *PlasmaFrameworkFilterer) WatchExitGameRegistered(opts *bind.WatchOpts, sink chan<- *PlasmaFrameworkExitGameRegistered) (event.Subscription, error) {

	logs, sub, err := _PlasmaFramework.contract.WatchLogs(opts, "ExitGameRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaFrameworkExitGameRegistered)
				if err := _PlasmaFramework.contract.UnpackLog(event, "ExitGameRegistered", log); err != nil {
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

// ParseExitGameRegistered is a log parse operation binding the contract event 0x0ceb02f7474c01f5959171cead2f5fb2a44d9c16fe480725a2110b601b19db82.
//
// Solidity: event ExitGameRegistered(uint256 txType, address exitGameAddress, uint8 protocol)
func (_PlasmaFramework *PlasmaFrameworkFilterer) ParseExitGameRegistered(log types.Log) (*PlasmaFrameworkExitGameRegistered, error) {
	event := new(PlasmaFrameworkExitGameRegistered)
	if err := _PlasmaFramework.contract.UnpackLog(event, "ExitGameRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PlasmaFrameworkExitQueueAddedIterator is returned from FilterExitQueueAdded and is used to iterate over the raw logs and unpacked data for ExitQueueAdded events raised by the PlasmaFramework contract.
type PlasmaFrameworkExitQueueAddedIterator struct {
	Event *PlasmaFrameworkExitQueueAdded // Event containing the contract specifics and raw log

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
func (it *PlasmaFrameworkExitQueueAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaFrameworkExitQueueAdded)
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
		it.Event = new(PlasmaFrameworkExitQueueAdded)
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
func (it *PlasmaFrameworkExitQueueAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaFrameworkExitQueueAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaFrameworkExitQueueAdded represents a ExitQueueAdded event raised by the PlasmaFramework contract.
type PlasmaFrameworkExitQueueAdded struct {
	VaultId *big.Int
	Token   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExitQueueAdded is a free log retrieval operation binding the contract event 0x0a925273fd754a41a580f8214815d5ae8a37367f561b3bd9d91a96730a418a17.
//
// Solidity: event ExitQueueAdded(uint256 vaultId, address token)
func (_PlasmaFramework *PlasmaFrameworkFilterer) FilterExitQueueAdded(opts *bind.FilterOpts) (*PlasmaFrameworkExitQueueAddedIterator, error) {

	logs, sub, err := _PlasmaFramework.contract.FilterLogs(opts, "ExitQueueAdded")
	if err != nil {
		return nil, err
	}
	return &PlasmaFrameworkExitQueueAddedIterator{contract: _PlasmaFramework.contract, event: "ExitQueueAdded", logs: logs, sub: sub}, nil
}

// WatchExitQueueAdded is a free log subscription operation binding the contract event 0x0a925273fd754a41a580f8214815d5ae8a37367f561b3bd9d91a96730a418a17.
//
// Solidity: event ExitQueueAdded(uint256 vaultId, address token)
func (_PlasmaFramework *PlasmaFrameworkFilterer) WatchExitQueueAdded(opts *bind.WatchOpts, sink chan<- *PlasmaFrameworkExitQueueAdded) (event.Subscription, error) {

	logs, sub, err := _PlasmaFramework.contract.WatchLogs(opts, "ExitQueueAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaFrameworkExitQueueAdded)
				if err := _PlasmaFramework.contract.UnpackLog(event, "ExitQueueAdded", log); err != nil {
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

// ParseExitQueueAdded is a log parse operation binding the contract event 0x0a925273fd754a41a580f8214815d5ae8a37367f561b3bd9d91a96730a418a17.
//
// Solidity: event ExitQueueAdded(uint256 vaultId, address token)
func (_PlasmaFramework *PlasmaFrameworkFilterer) ParseExitQueueAdded(log types.Log) (*PlasmaFrameworkExitQueueAdded, error) {
	event := new(PlasmaFrameworkExitQueueAdded)
	if err := _PlasmaFramework.contract.UnpackLog(event, "ExitQueueAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PlasmaFrameworkExitQueuedIterator is returned from FilterExitQueued and is used to iterate over the raw logs and unpacked data for ExitQueued events raised by the PlasmaFramework contract.
type PlasmaFrameworkExitQueuedIterator struct {
	Event *PlasmaFrameworkExitQueued // Event containing the contract specifics and raw log

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
func (it *PlasmaFrameworkExitQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaFrameworkExitQueued)
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
		it.Event = new(PlasmaFrameworkExitQueued)
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
func (it *PlasmaFrameworkExitQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaFrameworkExitQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaFrameworkExitQueued represents a ExitQueued event raised by the PlasmaFramework contract.
type PlasmaFrameworkExitQueued struct {
	ExitId   *big.Int
	Priority *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExitQueued is a free log retrieval operation binding the contract event 0xe15a4f223f922b625f5fdd941101a23fa0c097e522233d47a7cbea2167e92701.
//
// Solidity: event ExitQueued(uint160 indexed exitId, uint256 priority)
func (_PlasmaFramework *PlasmaFrameworkFilterer) FilterExitQueued(opts *bind.FilterOpts, exitId []*big.Int) (*PlasmaFrameworkExitQueuedIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PlasmaFramework.contract.FilterLogs(opts, "ExitQueued", exitIdRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaFrameworkExitQueuedIterator{contract: _PlasmaFramework.contract, event: "ExitQueued", logs: logs, sub: sub}, nil
}

// WatchExitQueued is a free log subscription operation binding the contract event 0xe15a4f223f922b625f5fdd941101a23fa0c097e522233d47a7cbea2167e92701.
//
// Solidity: event ExitQueued(uint160 indexed exitId, uint256 priority)
func (_PlasmaFramework *PlasmaFrameworkFilterer) WatchExitQueued(opts *bind.WatchOpts, sink chan<- *PlasmaFrameworkExitQueued, exitId []*big.Int) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PlasmaFramework.contract.WatchLogs(opts, "ExitQueued", exitIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaFrameworkExitQueued)
				if err := _PlasmaFramework.contract.UnpackLog(event, "ExitQueued", log); err != nil {
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

// ParseExitQueued is a log parse operation binding the contract event 0xe15a4f223f922b625f5fdd941101a23fa0c097e522233d47a7cbea2167e92701.
//
// Solidity: event ExitQueued(uint160 indexed exitId, uint256 priority)
func (_PlasmaFramework *PlasmaFrameworkFilterer) ParseExitQueued(log types.Log) (*PlasmaFrameworkExitQueued, error) {
	event := new(PlasmaFrameworkExitQueued)
	if err := _PlasmaFramework.contract.UnpackLog(event, "ExitQueued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PlasmaFrameworkProcessedExitsNumIterator is returned from FilterProcessedExitsNum and is used to iterate over the raw logs and unpacked data for ProcessedExitsNum events raised by the PlasmaFramework contract.
type PlasmaFrameworkProcessedExitsNumIterator struct {
	Event *PlasmaFrameworkProcessedExitsNum // Event containing the contract specifics and raw log

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
func (it *PlasmaFrameworkProcessedExitsNumIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaFrameworkProcessedExitsNum)
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
		it.Event = new(PlasmaFrameworkProcessedExitsNum)
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
func (it *PlasmaFrameworkProcessedExitsNumIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaFrameworkProcessedExitsNumIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaFrameworkProcessedExitsNum represents a ProcessedExitsNum event raised by the PlasmaFramework contract.
type PlasmaFrameworkProcessedExitsNum struct {
	ProcessedNum *big.Int
	VaultId      *big.Int
	Token        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProcessedExitsNum is a free log retrieval operation binding the contract event 0x7e35599fe450e626673ce6c37179ef9820aed16f76c1b350618ef7a612d13005.
//
// Solidity: event ProcessedExitsNum(uint256 processedNum, uint256 vaultId, address token)
func (_PlasmaFramework *PlasmaFrameworkFilterer) FilterProcessedExitsNum(opts *bind.FilterOpts) (*PlasmaFrameworkProcessedExitsNumIterator, error) {

	logs, sub, err := _PlasmaFramework.contract.FilterLogs(opts, "ProcessedExitsNum")
	if err != nil {
		return nil, err
	}
	return &PlasmaFrameworkProcessedExitsNumIterator{contract: _PlasmaFramework.contract, event: "ProcessedExitsNum", logs: logs, sub: sub}, nil
}

// WatchProcessedExitsNum is a free log subscription operation binding the contract event 0x7e35599fe450e626673ce6c37179ef9820aed16f76c1b350618ef7a612d13005.
//
// Solidity: event ProcessedExitsNum(uint256 processedNum, uint256 vaultId, address token)
func (_PlasmaFramework *PlasmaFrameworkFilterer) WatchProcessedExitsNum(opts *bind.WatchOpts, sink chan<- *PlasmaFrameworkProcessedExitsNum) (event.Subscription, error) {

	logs, sub, err := _PlasmaFramework.contract.WatchLogs(opts, "ProcessedExitsNum")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaFrameworkProcessedExitsNum)
				if err := _PlasmaFramework.contract.UnpackLog(event, "ProcessedExitsNum", log); err != nil {
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

// ParseProcessedExitsNum is a log parse operation binding the contract event 0x7e35599fe450e626673ce6c37179ef9820aed16f76c1b350618ef7a612d13005.
//
// Solidity: event ProcessedExitsNum(uint256 processedNum, uint256 vaultId, address token)
func (_PlasmaFramework *PlasmaFrameworkFilterer) ParseProcessedExitsNum(log types.Log) (*PlasmaFrameworkProcessedExitsNum, error) {
	event := new(PlasmaFrameworkProcessedExitsNum)
	if err := _PlasmaFramework.contract.UnpackLog(event, "ProcessedExitsNum", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PlasmaFrameworkVaultRegisteredIterator is returned from FilterVaultRegistered and is used to iterate over the raw logs and unpacked data for VaultRegistered events raised by the PlasmaFramework contract.
type PlasmaFrameworkVaultRegisteredIterator struct {
	Event *PlasmaFrameworkVaultRegistered // Event containing the contract specifics and raw log

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
func (it *PlasmaFrameworkVaultRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaFrameworkVaultRegistered)
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
		it.Event = new(PlasmaFrameworkVaultRegistered)
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
func (it *PlasmaFrameworkVaultRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaFrameworkVaultRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaFrameworkVaultRegistered represents a VaultRegistered event raised by the PlasmaFramework contract.
type PlasmaFrameworkVaultRegistered struct {
	VaultId      *big.Int
	VaultAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVaultRegistered is a free log retrieval operation binding the contract event 0x7051aac27f9b76ec8f37bd5796f2dcf402bd840a1c45c952a6eeb0e11bde0996.
//
// Solidity: event VaultRegistered(uint256 vaultId, address vaultAddress)
func (_PlasmaFramework *PlasmaFrameworkFilterer) FilterVaultRegistered(opts *bind.FilterOpts) (*PlasmaFrameworkVaultRegisteredIterator, error) {

	logs, sub, err := _PlasmaFramework.contract.FilterLogs(opts, "VaultRegistered")
	if err != nil {
		return nil, err
	}
	return &PlasmaFrameworkVaultRegisteredIterator{contract: _PlasmaFramework.contract, event: "VaultRegistered", logs: logs, sub: sub}, nil
}

// WatchVaultRegistered is a free log subscription operation binding the contract event 0x7051aac27f9b76ec8f37bd5796f2dcf402bd840a1c45c952a6eeb0e11bde0996.
//
// Solidity: event VaultRegistered(uint256 vaultId, address vaultAddress)
func (_PlasmaFramework *PlasmaFrameworkFilterer) WatchVaultRegistered(opts *bind.WatchOpts, sink chan<- *PlasmaFrameworkVaultRegistered) (event.Subscription, error) {

	logs, sub, err := _PlasmaFramework.contract.WatchLogs(opts, "VaultRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaFrameworkVaultRegistered)
				if err := _PlasmaFramework.contract.UnpackLog(event, "VaultRegistered", log); err != nil {
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

// ParseVaultRegistered is a log parse operation binding the contract event 0x7051aac27f9b76ec8f37bd5796f2dcf402bd840a1c45c952a6eeb0e11bde0996.
//
// Solidity: event VaultRegistered(uint256 vaultId, address vaultAddress)
func (_PlasmaFramework *PlasmaFrameworkFilterer) ParseVaultRegistered(log types.Log) (*PlasmaFrameworkVaultRegistered, error) {
	event := new(PlasmaFrameworkVaultRegistered)
	if err := _PlasmaFramework.contract.UnpackLog(event, "VaultRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

