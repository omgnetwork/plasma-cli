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

// PaymentExitGameABI is the input ABI used to generate the binding from.
const PaymentExitGameABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newBondSize\",\"type\":\"uint128\"}],\"name\":\"updateStartStandardExitBondSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inFlightTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"inFlightTxInputIndex\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"challengingTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"challengingTxInputIndex\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"challengingTxWitness\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"inputTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"inputUtxoPos\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"spendingConditionOptionalArgs\",\"type\":\"bytes\"}],\"internalType\":\"structPaymentInFlightExitRouterArgs.ChallengeInputSpentArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"challengeInFlightExitInputSpent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newBondSize\",\"type\":\"uint128\"}],\"name\":\"updatePiggybackBondSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inFlightTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"inFlightTxPos\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"inFlightTxInclusionProof\",\"type\":\"bytes\"}],\"name\":\"respondToNonCanonicalChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"piggybackBondSize\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"}],\"name\":\"standardExits\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"exitable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"utxoPos\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"outputId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"exitTarget\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bondSize\",\"type\":\"uint256\"}],\"internalType\":\"structPaymentExitDataModel.StandardExit\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startIFEBondSize\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_PB_BOND_SIZE\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"exitingTx\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"inputIndex\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"witness\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"spendingConditionOptionalArgs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"outputGuardPreimage\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"challengeTxPos\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"challengeTxInclusionProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"challengeTxConfirmSig\",\"type\":\"bytes\"}],\"internalType\":\"structPaymentStandardExitRouterArgs.ChallengeStandardExitArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"challengeStandardExit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inputTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"inputUtxoPos\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"inFlightTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"inFlightTxInputIndex\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"competingTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"competingTxInputIndex\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"outputGuardPreimage\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"competingTxPos\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"competingTxInclusionProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"competingTxWitness\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"competingTxConfirmSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"competingTxSpendingConditionOptionalArgs\",\"type\":\"bytes\"}],\"internalType\":\"structPaymentInFlightExitRouterArgs.ChallengeCanonicityArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"challengeInFlightExitNotCanonical\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"utxoPos\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rlpOutputTx\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"outputGuardPreimage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"outputTxInclusionProof\",\"type\":\"bytes\"}],\"internalType\":\"structPaymentStandardExitRouterArgs.StartStandardExitArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"startStandardExit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inFlightTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"outputIndex\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"outputGuardPreimage\",\"type\":\"bytes\"}],\"internalType\":\"structPaymentInFlightExitRouterArgs.PiggybackInFlightExitOnOutputArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"piggybackInFlightExitOnOutput\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BOND_LOWER_BOUND_DIVISOR\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BOND_UPPER_BOUND_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newBondSize\",\"type\":\"uint128\"}],\"name\":\"updateStartIFEBondSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inFlightTx\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"inFlightTxInclusionProof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"outputUtxoPos\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"challengingTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"challengingTxInputIndex\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"challengingTxWitness\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"spendingConditionOptionalArgs\",\"type\":\"bytes\"}],\"internalType\":\"structPaymentInFlightExitRouterArgs.ChallengeOutputSpent\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"challengeInFlightExitOutputSpent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inFlightTx\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputTxs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"inputUtxosPos\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"outputGuardPreimagesForInputs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"inputTxsInclusionProofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"inputTxsConfirmSigs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"inFlightTxWitnesses\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"inputSpendingConditionOptionalArgs\",\"type\":\"bytes[]\"}],\"internalType\":\"structPaymentInFlightExitRouterArgs.StartExitArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"startInFlightExit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_BOND_SIZE\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_IFE_BOND_SIZE\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"}],\"name\":\"inFlightExits\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isCanonical\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"exitStartTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"exitMap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"exitTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"piggybackBondSize\",\"type\":\"uint256\"}],\"internalType\":\"structPaymentExitDataModel.WithdrawData[4]\",\"name\":\"inputs\",\"type\":\"tuple[4]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"exitTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"piggybackBondSize\",\"type\":\"uint256\"}],\"internalType\":\"structPaymentExitDataModel.WithdrawData[4]\",\"name\":\"outputs\",\"type\":\"tuple[4]\"},{\"internalType\":\"addresspayable\",\"name\":\"bondOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bondSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldestCompetitorPosition\",\"type\":\"uint256\"}],\"internalType\":\"structPaymentExitDataModel.InFlightExit\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inFlightTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"inputIndex\",\"type\":\"uint16\"}],\"internalType\":\"structPaymentInFlightExitRouterArgs.PiggybackInFlightExitOnInputArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"piggybackInFlightExitOnInput\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startStandardExitBondSize\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractPlasmaFramework\",\"name\":\"framework\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ethVaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"erc20VaultId\",\"type\":\"uint256\"},{\"internalType\":\"contractOutputGuardHandlerRegistry\",\"name\":\"outputGuardHandlerRegistry\",\"type\":\"address\"},{\"internalType\":\"contractSpendingConditionRegistry\",\"name\":\"spendingConditionRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIStateTransitionVerifier\",\"name\":\"stateTransitionVerifier\",\"type\":\"address\"},{\"internalType\":\"contractITxFinalizationVerifier\",\"name\":\"txFinalizationVerifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"supportTxType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"safeGasStipend\",\"type\":\"uint256\"}],\"internalType\":\"structPaymentExitGame.PaymentExitGameArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"bondSize\",\"type\":\"uint128\"}],\"name\":\"IFEBondUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"bondSize\",\"type\":\"uint128\"}],\"name\":\"PiggybackBondUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"InFlightExitStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"exitTarget\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"inputIndex\",\"type\":\"uint16\"}],\"name\":\"InFlightExitInputPiggybacked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"InFlightExitOmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"outputIndex\",\"type\":\"uint16\"}],\"name\":\"InFlightExitOutputWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"inputIndex\",\"type\":\"uint16\"}],\"name\":\"InFlightExitInputWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"exitTarget\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"outputIndex\",\"type\":\"uint16\"}],\"name\":\"InFlightExitOutputPiggybacked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengeTxPosition\",\"type\":\"uint256\"}],\"name\":\"InFlightExitChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengeTxPosition\",\"type\":\"uint256\"}],\"name\":\"InFlightExitChallengeResponded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"inputIndex\",\"type\":\"uint16\"}],\"name\":\"InFlightExitInputBlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ifeTxHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"outputIndex\",\"type\":\"uint16\"}],\"name\":\"InFlightExitOutputBlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"bondSize\",\"type\":\"uint128\"}],\"name\":\"StandardExitBondUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"}],\"name\":\"ExitStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"utxoPos\",\"type\":\"uint256\"}],\"name\":\"ExitChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"}],\"name\":\"ExitOmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"}],\"name\":\"ExitFinalized\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"processExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isDeposit\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_txBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_utxoPos\",\"type\":\"uint256\"}],\"name\":\"getStandardExitId\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"\",\"type\":\"uint160\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_txBytes\",\"type\":\"bytes\"}],\"name\":\"getInFlightExitId\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"\",\"type\":\"uint160\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// PaymentExitGame is an auto generated Go binding around an Ethereum contract.
type PaymentExitGame struct {
	PaymentExitGameCaller     // Read-only binding to the contract
	PaymentExitGameTransactor // Write-only binding to the contract
	PaymentExitGameFilterer   // Log filterer for contract events
}

// PaymentExitGameCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentExitGameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentExitGameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentExitGameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentExitGameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentExitGameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentExitGameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentExitGameSession struct {
	Contract     *PaymentExitGame  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentExitGameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentExitGameCallerSession struct {
	Contract *PaymentExitGameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PaymentExitGameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentExitGameTransactorSession struct {
	Contract     *PaymentExitGameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PaymentExitGameRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentExitGameRaw struct {
	Contract *PaymentExitGame // Generic contract binding to access the raw methods on
}

// PaymentExitGameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentExitGameCallerRaw struct {
	Contract *PaymentExitGameCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentExitGameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentExitGameTransactorRaw struct {
	Contract *PaymentExitGameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentExitGame creates a new instance of PaymentExitGame, bound to a specific deployed contract.
func NewPaymentExitGame(address common.Address, backend bind.ContractBackend) (*PaymentExitGame, error) {
	contract, err := bindPaymentExitGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGame{PaymentExitGameCaller: PaymentExitGameCaller{contract: contract}, PaymentExitGameTransactor: PaymentExitGameTransactor{contract: contract}, PaymentExitGameFilterer: PaymentExitGameFilterer{contract: contract}}, nil
}

// NewPaymentExitGameCaller creates a new read-only instance of PaymentExitGame, bound to a specific deployed contract.
func NewPaymentExitGameCaller(address common.Address, caller bind.ContractCaller) (*PaymentExitGameCaller, error) {
	contract, err := bindPaymentExitGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameCaller{contract: contract}, nil
}

// NewPaymentExitGameTransactor creates a new write-only instance of PaymentExitGame, bound to a specific deployed contract.
func NewPaymentExitGameTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentExitGameTransactor, error) {
	contract, err := bindPaymentExitGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameTransactor{contract: contract}, nil
}

// NewPaymentExitGameFilterer creates a new log filterer instance of PaymentExitGame, bound to a specific deployed contract.
func NewPaymentExitGameFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentExitGameFilterer, error) {
	contract, err := bindPaymentExitGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameFilterer{contract: contract}, nil
}

// bindPaymentExitGame binds a generic wrapper to an already deployed contract.
func bindPaymentExitGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentExitGameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentExitGame *PaymentExitGameRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PaymentExitGame.Contract.PaymentExitGameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentExitGame *PaymentExitGameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.PaymentExitGameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentExitGame *PaymentExitGameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.PaymentExitGameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentExitGame *PaymentExitGameCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PaymentExitGame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentExitGame *PaymentExitGameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentExitGame *PaymentExitGameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.contract.Transact(opts, method, params...)
}

// PaymentExitDataModelInFlightExit is an auto generated low-level Go binding around an user-defined struct.
type PaymentExitDataModelInFlightExit struct {
	IsCanonical              bool
	ExitStartTimestamp       uint64
	ExitMap                  *big.Int
	Position                 *big.Int
	Inputs                   [4]PaymentExitDataModelWithdrawData
	Outputs                  [4]PaymentExitDataModelWithdrawData
	BondOwner                common.Address
	BondSize                 *big.Int
	OldestCompetitorPosition *big.Int
}

// PaymentExitDataModelStandardExit is an auto generated low-level Go binding around an user-defined struct.
type PaymentExitDataModelStandardExit struct {
	Exitable   bool
	UtxoPos    *big.Int
	OutputId   [32]byte
	ExitTarget common.Address
	Amount     *big.Int
	BondSize   *big.Int
}

// PaymentExitDataModelWithdrawData is an auto generated low-level Go binding around an user-defined struct.
type PaymentExitDataModelWithdrawData struct {
	OutputId          [32]byte
	ExitTarget        common.Address
	Token             common.Address
	Amount            *big.Int
	PiggybackBondSize *big.Int
}

// PaymentInFlightExitRouterArgsChallengeCanonicityArgs is an auto generated low-level Go binding around an user-defined struct.
type PaymentInFlightExitRouterArgsChallengeCanonicityArgs struct {
	InputTx                                  []byte
	InputUtxoPos                             *big.Int
	InFlightTx                               []byte
	InFlightTxInputIndex                     uint16
	CompetingTx                              []byte
	CompetingTxInputIndex                    uint16
	OutputGuardPreimage                      []byte
	CompetingTxPos                           *big.Int
	CompetingTxInclusionProof                []byte
	CompetingTxWitness                       []byte
	CompetingTxConfirmSig                    []byte
	CompetingTxSpendingConditionOptionalArgs []byte
}

// PaymentInFlightExitRouterArgsChallengeInputSpentArgs is an auto generated low-level Go binding around an user-defined struct.
type PaymentInFlightExitRouterArgsChallengeInputSpentArgs struct {
	InFlightTx                    []byte
	InFlightTxInputIndex          uint16
	ChallengingTx                 []byte
	ChallengingTxInputIndex       uint16
	ChallengingTxWitness          []byte
	InputTx                       []byte
	InputUtxoPos                  *big.Int
	SpendingConditionOptionalArgs []byte
}

// PaymentInFlightExitRouterArgsChallengeOutputSpent is an auto generated low-level Go binding around an user-defined struct.
type PaymentInFlightExitRouterArgsChallengeOutputSpent struct {
	InFlightTx                    []byte
	InFlightTxInclusionProof      []byte
	OutputUtxoPos                 *big.Int
	ChallengingTx                 []byte
	ChallengingTxInputIndex       uint16
	ChallengingTxWitness          []byte
	SpendingConditionOptionalArgs []byte
}

// PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs is an auto generated low-level Go binding around an user-defined struct.
type PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs struct {
	InFlightTx []byte
	InputIndex uint16
}

// PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs is an auto generated low-level Go binding around an user-defined struct.
type PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs struct {
	InFlightTx          []byte
	OutputIndex         uint16
	OutputGuardPreimage []byte
}

// PaymentInFlightExitRouterArgsStartExitArgs is an auto generated low-level Go binding around an user-defined struct.
type PaymentInFlightExitRouterArgsStartExitArgs struct {
	InFlightTx                         []byte
	InputTxs                           [][]byte
	InputUtxosPos                      []*big.Int
	OutputGuardPreimagesForInputs      [][]byte
	InputTxsInclusionProofs            [][]byte
	InputTxsConfirmSigs                [][]byte
	InFlightTxWitnesses                [][]byte
	InputSpendingConditionOptionalArgs [][]byte
}

// PaymentStandardExitRouterArgsChallengeStandardExitArgs is an auto generated low-level Go binding around an user-defined struct.
type PaymentStandardExitRouterArgsChallengeStandardExitArgs struct {
	ExitId                        *big.Int
	ExitingTx                     []byte
	ChallengeTx                   []byte
	InputIndex                    uint16
	Witness                       []byte
	SpendingConditionOptionalArgs []byte
	OutputGuardPreimage           []byte
	ChallengeTxPos                *big.Int
	ChallengeTxInclusionProof     []byte
	ChallengeTxConfirmSig         []byte
}

// PaymentStandardExitRouterArgsStartStandardExitArgs is an auto generated low-level Go binding around an user-defined struct.
type PaymentStandardExitRouterArgsStartStandardExitArgs struct {
	UtxoPos                *big.Int
	RlpOutputTx            []byte
	OutputGuardPreimage    []byte
	OutputTxInclusionProof []byte
}

// BONDLOWERBOUNDDIVISOR is a free data retrieval call binding the contract method 0x905d6a99.
//
// Solidity: function BOND_LOWER_BOUND_DIVISOR() constant returns(uint16)
func (_PaymentExitGame *PaymentExitGameCaller) BONDLOWERBOUNDDIVISOR(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _PaymentExitGame.contract.Call(opts, out, "BOND_LOWER_BOUND_DIVISOR")
	return *ret0, err
}

// BONDLOWERBOUNDDIVISOR is a free data retrieval call binding the contract method 0x905d6a99.
//
// Solidity: function BOND_LOWER_BOUND_DIVISOR() constant returns(uint16)
func (_PaymentExitGame *PaymentExitGameSession) BONDLOWERBOUNDDIVISOR() (uint16, error) {
	return _PaymentExitGame.Contract.BONDLOWERBOUNDDIVISOR(&_PaymentExitGame.CallOpts)
}

// BONDLOWERBOUNDDIVISOR is a free data retrieval call binding the contract method 0x905d6a99.
//
// Solidity: function BOND_LOWER_BOUND_DIVISOR() constant returns(uint16)
func (_PaymentExitGame *PaymentExitGameCallerSession) BONDLOWERBOUNDDIVISOR() (uint16, error) {
	return _PaymentExitGame.Contract.BONDLOWERBOUNDDIVISOR(&_PaymentExitGame.CallOpts)
}

// BONDUPPERBOUNDMULTIPLIER is a free data retrieval call binding the contract method 0xa0e403b1.
//
// Solidity: function BOND_UPPER_BOUND_MULTIPLIER() constant returns(uint16)
func (_PaymentExitGame *PaymentExitGameCaller) BONDUPPERBOUNDMULTIPLIER(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _PaymentExitGame.contract.Call(opts, out, "BOND_UPPER_BOUND_MULTIPLIER")
	return *ret0, err
}

// BONDUPPERBOUNDMULTIPLIER is a free data retrieval call binding the contract method 0xa0e403b1.
//
// Solidity: function BOND_UPPER_BOUND_MULTIPLIER() constant returns(uint16)
func (_PaymentExitGame *PaymentExitGameSession) BONDUPPERBOUNDMULTIPLIER() (uint16, error) {
	return _PaymentExitGame.Contract.BONDUPPERBOUNDMULTIPLIER(&_PaymentExitGame.CallOpts)
}

// BONDUPPERBOUNDMULTIPLIER is a free data retrieval call binding the contract method 0xa0e403b1.
//
// Solidity: function BOND_UPPER_BOUND_MULTIPLIER() constant returns(uint16)
func (_PaymentExitGame *PaymentExitGameCallerSession) BONDUPPERBOUNDMULTIPLIER() (uint16, error) {
	return _PaymentExitGame.Contract.BONDUPPERBOUNDMULTIPLIER(&_PaymentExitGame.CallOpts)
}

// INITIALBONDSIZE is a free data retrieval call binding the contract method 0xc170ecf5.
//
// Solidity: function INITIAL_BOND_SIZE() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameCaller) INITIALBONDSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitGame.contract.Call(opts, out, "INITIAL_BOND_SIZE")
	return *ret0, err
}

// INITIALBONDSIZE is a free data retrieval call binding the contract method 0xc170ecf5.
//
// Solidity: function INITIAL_BOND_SIZE() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameSession) INITIALBONDSIZE() (*big.Int, error) {
	return _PaymentExitGame.Contract.INITIALBONDSIZE(&_PaymentExitGame.CallOpts)
}

// INITIALBONDSIZE is a free data retrieval call binding the contract method 0xc170ecf5.
//
// Solidity: function INITIAL_BOND_SIZE() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameCallerSession) INITIALBONDSIZE() (*big.Int, error) {
	return _PaymentExitGame.Contract.INITIALBONDSIZE(&_PaymentExitGame.CallOpts)
}

// INITIALIFEBONDSIZE is a free data retrieval call binding the contract method 0xc96e7c05.
//
// Solidity: function INITIAL_IFE_BOND_SIZE() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameCaller) INITIALIFEBONDSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitGame.contract.Call(opts, out, "INITIAL_IFE_BOND_SIZE")
	return *ret0, err
}

// INITIALIFEBONDSIZE is a free data retrieval call binding the contract method 0xc96e7c05.
//
// Solidity: function INITIAL_IFE_BOND_SIZE() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameSession) INITIALIFEBONDSIZE() (*big.Int, error) {
	return _PaymentExitGame.Contract.INITIALIFEBONDSIZE(&_PaymentExitGame.CallOpts)
}

// INITIALIFEBONDSIZE is a free data retrieval call binding the contract method 0xc96e7c05.
//
// Solidity: function INITIAL_IFE_BOND_SIZE() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameCallerSession) INITIALIFEBONDSIZE() (*big.Int, error) {
	return _PaymentExitGame.Contract.INITIALIFEBONDSIZE(&_PaymentExitGame.CallOpts)
}

// INITIALPBBONDSIZE is a free data retrieval call binding the contract method 0x7e2b2efe.
//
// Solidity: function INITIAL_PB_BOND_SIZE() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameCaller) INITIALPBBONDSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitGame.contract.Call(opts, out, "INITIAL_PB_BOND_SIZE")
	return *ret0, err
}

// INITIALPBBONDSIZE is a free data retrieval call binding the contract method 0x7e2b2efe.
//
// Solidity: function INITIAL_PB_BOND_SIZE() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameSession) INITIALPBBONDSIZE() (*big.Int, error) {
	return _PaymentExitGame.Contract.INITIALPBBONDSIZE(&_PaymentExitGame.CallOpts)
}

// INITIALPBBONDSIZE is a free data retrieval call binding the contract method 0x7e2b2efe.
//
// Solidity: function INITIAL_PB_BOND_SIZE() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameCallerSession) INITIALPBBONDSIZE() (*big.Int, error) {
	return _PaymentExitGame.Contract.INITIALPBBONDSIZE(&_PaymentExitGame.CallOpts)
}

// GetInFlightExitId is a free data retrieval call binding the contract method 0x839c78f9.
//
// Solidity: function getInFlightExitId(bytes _txBytes) constant returns(uint160)
func (_PaymentExitGame *PaymentExitGameCaller) GetInFlightExitId(opts *bind.CallOpts, _txBytes []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitGame.contract.Call(opts, out, "getInFlightExitId", _txBytes)
	return *ret0, err
}

// GetInFlightExitId is a free data retrieval call binding the contract method 0x839c78f9.
//
// Solidity: function getInFlightExitId(bytes _txBytes) constant returns(uint160)
func (_PaymentExitGame *PaymentExitGameSession) GetInFlightExitId(_txBytes []byte) (*big.Int, error) {
	return _PaymentExitGame.Contract.GetInFlightExitId(&_PaymentExitGame.CallOpts, _txBytes)
}

// GetInFlightExitId is a free data retrieval call binding the contract method 0x839c78f9.
//
// Solidity: function getInFlightExitId(bytes _txBytes) constant returns(uint160)
func (_PaymentExitGame *PaymentExitGameCallerSession) GetInFlightExitId(_txBytes []byte) (*big.Int, error) {
	return _PaymentExitGame.Contract.GetInFlightExitId(&_PaymentExitGame.CallOpts, _txBytes)
}

// GetStandardExitId is a free data retrieval call binding the contract method 0xb22d3e14.
//
// Solidity: function getStandardExitId(bool _isDeposit, bytes _txBytes, uint256 _utxoPos) constant returns(uint160)
func (_PaymentExitGame *PaymentExitGameCaller) GetStandardExitId(opts *bind.CallOpts, _isDeposit bool, _txBytes []byte, _utxoPos *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitGame.contract.Call(opts, out, "getStandardExitId", _isDeposit, _txBytes, _utxoPos)
	return *ret0, err
}

// GetStandardExitId is a free data retrieval call binding the contract method 0xb22d3e14.
//
// Solidity: function getStandardExitId(bool _isDeposit, bytes _txBytes, uint256 _utxoPos) constant returns(uint160)
func (_PaymentExitGame *PaymentExitGameSession) GetStandardExitId(_isDeposit bool, _txBytes []byte, _utxoPos *big.Int) (*big.Int, error) {
	return _PaymentExitGame.Contract.GetStandardExitId(&_PaymentExitGame.CallOpts, _isDeposit, _txBytes, _utxoPos)
}

// GetStandardExitId is a free data retrieval call binding the contract method 0xb22d3e14.
//
// Solidity: function getStandardExitId(bool _isDeposit, bytes _txBytes, uint256 _utxoPos) constant returns(uint160)
func (_PaymentExitGame *PaymentExitGameCallerSession) GetStandardExitId(_isDeposit bool, _txBytes []byte, _utxoPos *big.Int) (*big.Int, error) {
	return _PaymentExitGame.Contract.GetStandardExitId(&_PaymentExitGame.CallOpts, _isDeposit, _txBytes, _utxoPos)
}

// InFlightExits is a free data retrieval call binding the contract method 0xe47a4631.
//
// Solidity: function inFlightExits(uint160 exitId) constant returns(PaymentExitDataModelInFlightExit)
func (_PaymentExitGame *PaymentExitGameCaller) InFlightExits(opts *bind.CallOpts, exitId *big.Int) (PaymentExitDataModelInFlightExit, error) {
	var (
		ret0 = new(PaymentExitDataModelInFlightExit)
	)
	out := ret0
	err := _PaymentExitGame.contract.Call(opts, out, "inFlightExits", exitId)
	return *ret0, err
}

// InFlightExits is a free data retrieval call binding the contract method 0xe47a4631.
//
// Solidity: function inFlightExits(uint160 exitId) constant returns(PaymentExitDataModelInFlightExit)
func (_PaymentExitGame *PaymentExitGameSession) InFlightExits(exitId *big.Int) (PaymentExitDataModelInFlightExit, error) {
	return _PaymentExitGame.Contract.InFlightExits(&_PaymentExitGame.CallOpts, exitId)
}

// InFlightExits is a free data retrieval call binding the contract method 0xe47a4631.
//
// Solidity: function inFlightExits(uint160 exitId) constant returns(PaymentExitDataModelInFlightExit)
func (_PaymentExitGame *PaymentExitGameCallerSession) InFlightExits(exitId *big.Int) (PaymentExitDataModelInFlightExit, error) {
	return _PaymentExitGame.Contract.InFlightExits(&_PaymentExitGame.CallOpts, exitId)
}

// PiggybackBondSize is a free data retrieval call binding the contract method 0x3f1214cf.
//
// Solidity: function piggybackBondSize() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameCaller) PiggybackBondSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitGame.contract.Call(opts, out, "piggybackBondSize")
	return *ret0, err
}

// PiggybackBondSize is a free data retrieval call binding the contract method 0x3f1214cf.
//
// Solidity: function piggybackBondSize() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameSession) PiggybackBondSize() (*big.Int, error) {
	return _PaymentExitGame.Contract.PiggybackBondSize(&_PaymentExitGame.CallOpts)
}

// PiggybackBondSize is a free data retrieval call binding the contract method 0x3f1214cf.
//
// Solidity: function piggybackBondSize() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameCallerSession) PiggybackBondSize() (*big.Int, error) {
	return _PaymentExitGame.Contract.PiggybackBondSize(&_PaymentExitGame.CallOpts)
}

// StandardExits is a free data retrieval call binding the contract method 0x5c0e683a.
//
// Solidity: function standardExits(uint160 exitId) constant returns(PaymentExitDataModelStandardExit)
func (_PaymentExitGame *PaymentExitGameCaller) StandardExits(opts *bind.CallOpts, exitId *big.Int) (PaymentExitDataModelStandardExit, error) {
	var (
		ret0 = new(PaymentExitDataModelStandardExit)
	)
	out := ret0
	err := _PaymentExitGame.contract.Call(opts, out, "standardExits", exitId)
	return *ret0, err
}

// StandardExits is a free data retrieval call binding the contract method 0x5c0e683a.
//
// Solidity: function standardExits(uint160 exitId) constant returns(PaymentExitDataModelStandardExit)
func (_PaymentExitGame *PaymentExitGameSession) StandardExits(exitId *big.Int) (PaymentExitDataModelStandardExit, error) {
	return _PaymentExitGame.Contract.StandardExits(&_PaymentExitGame.CallOpts, exitId)
}

// StandardExits is a free data retrieval call binding the contract method 0x5c0e683a.
//
// Solidity: function standardExits(uint160 exitId) constant returns(PaymentExitDataModelStandardExit)
func (_PaymentExitGame *PaymentExitGameCallerSession) StandardExits(exitId *big.Int) (PaymentExitDataModelStandardExit, error) {
	return _PaymentExitGame.Contract.StandardExits(&_PaymentExitGame.CallOpts, exitId)
}

// StartIFEBondSize is a free data retrieval call binding the contract method 0x7702fa17.
//
// Solidity: function startIFEBondSize() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameCaller) StartIFEBondSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitGame.contract.Call(opts, out, "startIFEBondSize")
	return *ret0, err
}

// StartIFEBondSize is a free data retrieval call binding the contract method 0x7702fa17.
//
// Solidity: function startIFEBondSize() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameSession) StartIFEBondSize() (*big.Int, error) {
	return _PaymentExitGame.Contract.StartIFEBondSize(&_PaymentExitGame.CallOpts)
}

// StartIFEBondSize is a free data retrieval call binding the contract method 0x7702fa17.
//
// Solidity: function startIFEBondSize() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameCallerSession) StartIFEBondSize() (*big.Int, error) {
	return _PaymentExitGame.Contract.StartIFEBondSize(&_PaymentExitGame.CallOpts)
}

// StartStandardExitBondSize is a free data retrieval call binding the contract method 0xfe32a124.
//
// Solidity: function startStandardExitBondSize() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameCaller) StartStandardExitBondSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitGame.contract.Call(opts, out, "startStandardExitBondSize")
	return *ret0, err
}

// StartStandardExitBondSize is a free data retrieval call binding the contract method 0xfe32a124.
//
// Solidity: function startStandardExitBondSize() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameSession) StartStandardExitBondSize() (*big.Int, error) {
	return _PaymentExitGame.Contract.StartStandardExitBondSize(&_PaymentExitGame.CallOpts)
}

// StartStandardExitBondSize is a free data retrieval call binding the contract method 0xfe32a124.
//
// Solidity: function startStandardExitBondSize() constant returns(uint128)
func (_PaymentExitGame *PaymentExitGameCallerSession) StartStandardExitBondSize() (*big.Int, error) {
	return _PaymentExitGame.Contract.StartStandardExitBondSize(&_PaymentExitGame.CallOpts)
}

// ChallengeInFlightExitInputSpent is a paid mutator transaction binding the contract method 0x2d3cd2e9.
//
// Solidity: function challengeInFlightExitInputSpent(PaymentInFlightExitRouterArgsChallengeInputSpentArgs args) returns()
func (_PaymentExitGame *PaymentExitGameTransactor) ChallengeInFlightExitInputSpent(opts *bind.TransactOpts, args PaymentInFlightExitRouterArgsChallengeInputSpentArgs) (*types.Transaction, error) {
	return _PaymentExitGame.contract.Transact(opts, "challengeInFlightExitInputSpent", args)
}

// ChallengeInFlightExitInputSpent is a paid mutator transaction binding the contract method 0x2d3cd2e9.
//
// Solidity: function challengeInFlightExitInputSpent(PaymentInFlightExitRouterArgsChallengeInputSpentArgs args) returns()
func (_PaymentExitGame *PaymentExitGameSession) ChallengeInFlightExitInputSpent(args PaymentInFlightExitRouterArgsChallengeInputSpentArgs) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.ChallengeInFlightExitInputSpent(&_PaymentExitGame.TransactOpts, args)
}

// ChallengeInFlightExitInputSpent is a paid mutator transaction binding the contract method 0x2d3cd2e9.
//
// Solidity: function challengeInFlightExitInputSpent(PaymentInFlightExitRouterArgsChallengeInputSpentArgs args) returns()
func (_PaymentExitGame *PaymentExitGameTransactorSession) ChallengeInFlightExitInputSpent(args PaymentInFlightExitRouterArgsChallengeInputSpentArgs) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.ChallengeInFlightExitInputSpent(&_PaymentExitGame.TransactOpts, args)
}

// ChallengeInFlightExitNotCanonical is a paid mutator transaction binding the contract method 0x8c9963f0.
//
// Solidity: function challengeInFlightExitNotCanonical(PaymentInFlightExitRouterArgsChallengeCanonicityArgs args) returns()
func (_PaymentExitGame *PaymentExitGameTransactor) ChallengeInFlightExitNotCanonical(opts *bind.TransactOpts, args PaymentInFlightExitRouterArgsChallengeCanonicityArgs) (*types.Transaction, error) {
	return _PaymentExitGame.contract.Transact(opts, "challengeInFlightExitNotCanonical", args)
}

// ChallengeInFlightExitNotCanonical is a paid mutator transaction binding the contract method 0x8c9963f0.
//
// Solidity: function challengeInFlightExitNotCanonical(PaymentInFlightExitRouterArgsChallengeCanonicityArgs args) returns()
func (_PaymentExitGame *PaymentExitGameSession) ChallengeInFlightExitNotCanonical(args PaymentInFlightExitRouterArgsChallengeCanonicityArgs) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.ChallengeInFlightExitNotCanonical(&_PaymentExitGame.TransactOpts, args)
}

// ChallengeInFlightExitNotCanonical is a paid mutator transaction binding the contract method 0x8c9963f0.
//
// Solidity: function challengeInFlightExitNotCanonical(PaymentInFlightExitRouterArgsChallengeCanonicityArgs args) returns()
func (_PaymentExitGame *PaymentExitGameTransactorSession) ChallengeInFlightExitNotCanonical(args PaymentInFlightExitRouterArgsChallengeCanonicityArgs) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.ChallengeInFlightExitNotCanonical(&_PaymentExitGame.TransactOpts, args)
}

// ChallengeInFlightExitOutputSpent is a paid mutator transaction binding the contract method 0xb8fd17fa.
//
// Solidity: function challengeInFlightExitOutputSpent(PaymentInFlightExitRouterArgsChallengeOutputSpent args) returns()
func (_PaymentExitGame *PaymentExitGameTransactor) ChallengeInFlightExitOutputSpent(opts *bind.TransactOpts, args PaymentInFlightExitRouterArgsChallengeOutputSpent) (*types.Transaction, error) {
	return _PaymentExitGame.contract.Transact(opts, "challengeInFlightExitOutputSpent", args)
}

// ChallengeInFlightExitOutputSpent is a paid mutator transaction binding the contract method 0xb8fd17fa.
//
// Solidity: function challengeInFlightExitOutputSpent(PaymentInFlightExitRouterArgsChallengeOutputSpent args) returns()
func (_PaymentExitGame *PaymentExitGameSession) ChallengeInFlightExitOutputSpent(args PaymentInFlightExitRouterArgsChallengeOutputSpent) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.ChallengeInFlightExitOutputSpent(&_PaymentExitGame.TransactOpts, args)
}

// ChallengeInFlightExitOutputSpent is a paid mutator transaction binding the contract method 0xb8fd17fa.
//
// Solidity: function challengeInFlightExitOutputSpent(PaymentInFlightExitRouterArgsChallengeOutputSpent args) returns()
func (_PaymentExitGame *PaymentExitGameTransactorSession) ChallengeInFlightExitOutputSpent(args PaymentInFlightExitRouterArgsChallengeOutputSpent) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.ChallengeInFlightExitOutputSpent(&_PaymentExitGame.TransactOpts, args)
}

// ChallengeStandardExit is a paid mutator transaction binding the contract method 0x8301d179.
//
// Solidity: function challengeStandardExit(PaymentStandardExitRouterArgsChallengeStandardExitArgs args) returns()
func (_PaymentExitGame *PaymentExitGameTransactor) ChallengeStandardExit(opts *bind.TransactOpts, args PaymentStandardExitRouterArgsChallengeStandardExitArgs) (*types.Transaction, error) {
	return _PaymentExitGame.contract.Transact(opts, "challengeStandardExit", args)
}

// ChallengeStandardExit is a paid mutator transaction binding the contract method 0x8301d179.
//
// Solidity: function challengeStandardExit(PaymentStandardExitRouterArgsChallengeStandardExitArgs args) returns()
func (_PaymentExitGame *PaymentExitGameSession) ChallengeStandardExit(args PaymentStandardExitRouterArgsChallengeStandardExitArgs) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.ChallengeStandardExit(&_PaymentExitGame.TransactOpts, args)
}

// ChallengeStandardExit is a paid mutator transaction binding the contract method 0x8301d179.
//
// Solidity: function challengeStandardExit(PaymentStandardExitRouterArgsChallengeStandardExitArgs args) returns()
func (_PaymentExitGame *PaymentExitGameTransactorSession) ChallengeStandardExit(args PaymentStandardExitRouterArgsChallengeStandardExitArgs) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.ChallengeStandardExit(&_PaymentExitGame.TransactOpts, args)
}

// PiggybackInFlightExitOnInput is a paid mutator transaction binding the contract method 0xe5bc60c7.
//
// Solidity: function piggybackInFlightExitOnInput(PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs args) returns()
func (_PaymentExitGame *PaymentExitGameTransactor) PiggybackInFlightExitOnInput(opts *bind.TransactOpts, args PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs) (*types.Transaction, error) {
	return _PaymentExitGame.contract.Transact(opts, "piggybackInFlightExitOnInput", args)
}

// PiggybackInFlightExitOnInput is a paid mutator transaction binding the contract method 0xe5bc60c7.
//
// Solidity: function piggybackInFlightExitOnInput(PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs args) returns()
func (_PaymentExitGame *PaymentExitGameSession) PiggybackInFlightExitOnInput(args PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.PiggybackInFlightExitOnInput(&_PaymentExitGame.TransactOpts, args)
}

// PiggybackInFlightExitOnInput is a paid mutator transaction binding the contract method 0xe5bc60c7.
//
// Solidity: function piggybackInFlightExitOnInput(PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs args) returns()
func (_PaymentExitGame *PaymentExitGameTransactorSession) PiggybackInFlightExitOnInput(args PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.PiggybackInFlightExitOnInput(&_PaymentExitGame.TransactOpts, args)
}

// PiggybackInFlightExitOnOutput is a paid mutator transaction binding the contract method 0x8ec32e8f.
//
// Solidity: function piggybackInFlightExitOnOutput(PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs args) returns()
func (_PaymentExitGame *PaymentExitGameTransactor) PiggybackInFlightExitOnOutput(opts *bind.TransactOpts, args PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs) (*types.Transaction, error) {
	return _PaymentExitGame.contract.Transact(opts, "piggybackInFlightExitOnOutput", args)
}

// PiggybackInFlightExitOnOutput is a paid mutator transaction binding the contract method 0x8ec32e8f.
//
// Solidity: function piggybackInFlightExitOnOutput(PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs args) returns()
func (_PaymentExitGame *PaymentExitGameSession) PiggybackInFlightExitOnOutput(args PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.PiggybackInFlightExitOnOutput(&_PaymentExitGame.TransactOpts, args)
}

// PiggybackInFlightExitOnOutput is a paid mutator transaction binding the contract method 0x8ec32e8f.
//
// Solidity: function piggybackInFlightExitOnOutput(PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs args) returns()
func (_PaymentExitGame *PaymentExitGameTransactorSession) PiggybackInFlightExitOnOutput(args PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.PiggybackInFlightExitOnOutput(&_PaymentExitGame.TransactOpts, args)
}

// ProcessExit is a paid mutator transaction binding the contract method 0x68496351.
//
// Solidity: function processExit(uint160 exitId, uint256 , address token) returns()
func (_PaymentExitGame *PaymentExitGameTransactor) ProcessExit(opts *bind.TransactOpts, exitId *big.Int, arg1 *big.Int, token common.Address) (*types.Transaction, error) {
	return _PaymentExitGame.contract.Transact(opts, "processExit", exitId, arg1, token)
}

// ProcessExit is a paid mutator transaction binding the contract method 0x68496351.
//
// Solidity: function processExit(uint160 exitId, uint256 , address token) returns()
func (_PaymentExitGame *PaymentExitGameSession) ProcessExit(exitId *big.Int, arg1 *big.Int, token common.Address) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.ProcessExit(&_PaymentExitGame.TransactOpts, exitId, arg1, token)
}

// ProcessExit is a paid mutator transaction binding the contract method 0x68496351.
//
// Solidity: function processExit(uint160 exitId, uint256 , address token) returns()
func (_PaymentExitGame *PaymentExitGameTransactorSession) ProcessExit(exitId *big.Int, arg1 *big.Int, token common.Address) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.ProcessExit(&_PaymentExitGame.TransactOpts, exitId, arg1, token)
}

// RespondToNonCanonicalChallenge is a paid mutator transaction binding the contract method 0x36660de4.
//
// Solidity: function respondToNonCanonicalChallenge(bytes inFlightTx, uint256 inFlightTxPos, bytes inFlightTxInclusionProof) returns()
func (_PaymentExitGame *PaymentExitGameTransactor) RespondToNonCanonicalChallenge(opts *bind.TransactOpts, inFlightTx []byte, inFlightTxPos *big.Int, inFlightTxInclusionProof []byte) (*types.Transaction, error) {
	return _PaymentExitGame.contract.Transact(opts, "respondToNonCanonicalChallenge", inFlightTx, inFlightTxPos, inFlightTxInclusionProof)
}

// RespondToNonCanonicalChallenge is a paid mutator transaction binding the contract method 0x36660de4.
//
// Solidity: function respondToNonCanonicalChallenge(bytes inFlightTx, uint256 inFlightTxPos, bytes inFlightTxInclusionProof) returns()
func (_PaymentExitGame *PaymentExitGameSession) RespondToNonCanonicalChallenge(inFlightTx []byte, inFlightTxPos *big.Int, inFlightTxInclusionProof []byte) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.RespondToNonCanonicalChallenge(&_PaymentExitGame.TransactOpts, inFlightTx, inFlightTxPos, inFlightTxInclusionProof)
}

// RespondToNonCanonicalChallenge is a paid mutator transaction binding the contract method 0x36660de4.
//
// Solidity: function respondToNonCanonicalChallenge(bytes inFlightTx, uint256 inFlightTxPos, bytes inFlightTxInclusionProof) returns()
func (_PaymentExitGame *PaymentExitGameTransactorSession) RespondToNonCanonicalChallenge(inFlightTx []byte, inFlightTxPos *big.Int, inFlightTxInclusionProof []byte) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.RespondToNonCanonicalChallenge(&_PaymentExitGame.TransactOpts, inFlightTx, inFlightTxPos, inFlightTxInclusionProof)
}

// StartInFlightExit is a paid mutator transaction binding the contract method 0xb90424b5.
//
// Solidity: function startInFlightExit(PaymentInFlightExitRouterArgsStartExitArgs args) returns()
func (_PaymentExitGame *PaymentExitGameTransactor) StartInFlightExit(opts *bind.TransactOpts, args PaymentInFlightExitRouterArgsStartExitArgs) (*types.Transaction, error) {
	return _PaymentExitGame.contract.Transact(opts, "startInFlightExit", args)
}

// StartInFlightExit is a paid mutator transaction binding the contract method 0xb90424b5.
//
// Solidity: function startInFlightExit(PaymentInFlightExitRouterArgsStartExitArgs args) returns()
func (_PaymentExitGame *PaymentExitGameSession) StartInFlightExit(args PaymentInFlightExitRouterArgsStartExitArgs) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.StartInFlightExit(&_PaymentExitGame.TransactOpts, args)
}

// StartInFlightExit is a paid mutator transaction binding the contract method 0xb90424b5.
//
// Solidity: function startInFlightExit(PaymentInFlightExitRouterArgsStartExitArgs args) returns()
func (_PaymentExitGame *PaymentExitGameTransactorSession) StartInFlightExit(args PaymentInFlightExitRouterArgsStartExitArgs) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.StartInFlightExit(&_PaymentExitGame.TransactOpts, args)
}

// StartStandardExit is a paid mutator transaction binding the contract method 0x8d6f74cf.
//
// Solidity: function startStandardExit(PaymentStandardExitRouterArgsStartStandardExitArgs args) returns()
func (_PaymentExitGame *PaymentExitGameTransactor) StartStandardExit(opts *bind.TransactOpts, args PaymentStandardExitRouterArgsStartStandardExitArgs) (*types.Transaction, error) {
	return _PaymentExitGame.contract.Transact(opts, "startStandardExit", args)
}

// StartStandardExit is a paid mutator transaction binding the contract method 0x8d6f74cf.
//
// Solidity: function startStandardExit(PaymentStandardExitRouterArgsStartStandardExitArgs args) returns()
func (_PaymentExitGame *PaymentExitGameSession) StartStandardExit(args PaymentStandardExitRouterArgsStartStandardExitArgs) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.StartStandardExit(&_PaymentExitGame.TransactOpts, args)
}

// StartStandardExit is a paid mutator transaction binding the contract method 0x8d6f74cf.
//
// Solidity: function startStandardExit(PaymentStandardExitRouterArgsStartStandardExitArgs args) returns()
func (_PaymentExitGame *PaymentExitGameTransactorSession) StartStandardExit(args PaymentStandardExitRouterArgsStartStandardExitArgs) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.StartStandardExit(&_PaymentExitGame.TransactOpts, args)
}

// UpdatePiggybackBondSize is a paid mutator transaction binding the contract method 0x315fb7da.
//
// Solidity: function updatePiggybackBondSize(uint128 newBondSize) returns()
func (_PaymentExitGame *PaymentExitGameTransactor) UpdatePiggybackBondSize(opts *bind.TransactOpts, newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitGame.contract.Transact(opts, "updatePiggybackBondSize", newBondSize)
}

// UpdatePiggybackBondSize is a paid mutator transaction binding the contract method 0x315fb7da.
//
// Solidity: function updatePiggybackBondSize(uint128 newBondSize) returns()
func (_PaymentExitGame *PaymentExitGameSession) UpdatePiggybackBondSize(newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.UpdatePiggybackBondSize(&_PaymentExitGame.TransactOpts, newBondSize)
}

// UpdatePiggybackBondSize is a paid mutator transaction binding the contract method 0x315fb7da.
//
// Solidity: function updatePiggybackBondSize(uint128 newBondSize) returns()
func (_PaymentExitGame *PaymentExitGameTransactorSession) UpdatePiggybackBondSize(newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.UpdatePiggybackBondSize(&_PaymentExitGame.TransactOpts, newBondSize)
}

// UpdateStartIFEBondSize is a paid mutator transaction binding the contract method 0xb177ba23.
//
// Solidity: function updateStartIFEBondSize(uint128 newBondSize) returns()
func (_PaymentExitGame *PaymentExitGameTransactor) UpdateStartIFEBondSize(opts *bind.TransactOpts, newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitGame.contract.Transact(opts, "updateStartIFEBondSize", newBondSize)
}

// UpdateStartIFEBondSize is a paid mutator transaction binding the contract method 0xb177ba23.
//
// Solidity: function updateStartIFEBondSize(uint128 newBondSize) returns()
func (_PaymentExitGame *PaymentExitGameSession) UpdateStartIFEBondSize(newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.UpdateStartIFEBondSize(&_PaymentExitGame.TransactOpts, newBondSize)
}

// UpdateStartIFEBondSize is a paid mutator transaction binding the contract method 0xb177ba23.
//
// Solidity: function updateStartIFEBondSize(uint128 newBondSize) returns()
func (_PaymentExitGame *PaymentExitGameTransactorSession) UpdateStartIFEBondSize(newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.UpdateStartIFEBondSize(&_PaymentExitGame.TransactOpts, newBondSize)
}

// UpdateStartStandardExitBondSize is a paid mutator transaction binding the contract method 0x163d7bc7.
//
// Solidity: function updateStartStandardExitBondSize(uint128 newBondSize) returns()
func (_PaymentExitGame *PaymentExitGameTransactor) UpdateStartStandardExitBondSize(opts *bind.TransactOpts, newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitGame.contract.Transact(opts, "updateStartStandardExitBondSize", newBondSize)
}

// UpdateStartStandardExitBondSize is a paid mutator transaction binding the contract method 0x163d7bc7.
//
// Solidity: function updateStartStandardExitBondSize(uint128 newBondSize) returns()
func (_PaymentExitGame *PaymentExitGameSession) UpdateStartStandardExitBondSize(newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.UpdateStartStandardExitBondSize(&_PaymentExitGame.TransactOpts, newBondSize)
}

// UpdateStartStandardExitBondSize is a paid mutator transaction binding the contract method 0x163d7bc7.
//
// Solidity: function updateStartStandardExitBondSize(uint128 newBondSize) returns()
func (_PaymentExitGame *PaymentExitGameTransactorSession) UpdateStartStandardExitBondSize(newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitGame.Contract.UpdateStartStandardExitBondSize(&_PaymentExitGame.TransactOpts, newBondSize)
}

// PaymentExitGameExitChallengedIterator is returned from FilterExitChallenged and is used to iterate over the raw logs and unpacked data for ExitChallenged events raised by the PaymentExitGame contract.
type PaymentExitGameExitChallengedIterator struct {
	Event *PaymentExitGameExitChallenged // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameExitChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameExitChallenged)
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
		it.Event = new(PaymentExitGameExitChallenged)
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
func (it *PaymentExitGameExitChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameExitChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameExitChallenged represents a ExitChallenged event raised by the PaymentExitGame contract.
type PaymentExitGameExitChallenged struct {
	UtxoPos *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExitChallenged is a free log retrieval operation binding the contract event 0x5dfba526c59b25f899f935c5b0d5b8739e97e4d89c38c158eca3192ea34b87d8.
//
// Solidity: event ExitChallenged(uint256 indexed utxoPos)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterExitChallenged(opts *bind.FilterOpts, utxoPos []*big.Int) (*PaymentExitGameExitChallengedIterator, error) {

	var utxoPosRule []interface{}
	for _, utxoPosItem := range utxoPos {
		utxoPosRule = append(utxoPosRule, utxoPosItem)
	}

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "ExitChallenged", utxoPosRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameExitChallengedIterator{contract: _PaymentExitGame.contract, event: "ExitChallenged", logs: logs, sub: sub}, nil
}

// WatchExitChallenged is a free log subscription operation binding the contract event 0x5dfba526c59b25f899f935c5b0d5b8739e97e4d89c38c158eca3192ea34b87d8.
//
// Solidity: event ExitChallenged(uint256 indexed utxoPos)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchExitChallenged(opts *bind.WatchOpts, sink chan<- *PaymentExitGameExitChallenged, utxoPos []*big.Int) (event.Subscription, error) {

	var utxoPosRule []interface{}
	for _, utxoPosItem := range utxoPos {
		utxoPosRule = append(utxoPosRule, utxoPosItem)
	}

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "ExitChallenged", utxoPosRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameExitChallenged)
				if err := _PaymentExitGame.contract.UnpackLog(event, "ExitChallenged", log); err != nil {
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

// ParseExitChallenged is a log parse operation binding the contract event 0x5dfba526c59b25f899f935c5b0d5b8739e97e4d89c38c158eca3192ea34b87d8.
//
// Solidity: event ExitChallenged(uint256 indexed utxoPos)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseExitChallenged(log types.Log) (*PaymentExitGameExitChallenged, error) {
	event := new(PaymentExitGameExitChallenged)
	if err := _PaymentExitGame.contract.UnpackLog(event, "ExitChallenged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGameExitFinalizedIterator is returned from FilterExitFinalized and is used to iterate over the raw logs and unpacked data for ExitFinalized events raised by the PaymentExitGame contract.
type PaymentExitGameExitFinalizedIterator struct {
	Event *PaymentExitGameExitFinalized // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameExitFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameExitFinalized)
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
		it.Event = new(PaymentExitGameExitFinalized)
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
func (it *PaymentExitGameExitFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameExitFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameExitFinalized represents a ExitFinalized event raised by the PaymentExitGame contract.
type PaymentExitGameExitFinalized struct {
	ExitId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitFinalized is a free log retrieval operation binding the contract event 0x0adb29b0831e081044cefe31155c1f2b2b85ad3613a480a5f901ee287addef55.
//
// Solidity: event ExitFinalized(uint160 indexed exitId)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterExitFinalized(opts *bind.FilterOpts, exitId []*big.Int) (*PaymentExitGameExitFinalizedIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "ExitFinalized", exitIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameExitFinalizedIterator{contract: _PaymentExitGame.contract, event: "ExitFinalized", logs: logs, sub: sub}, nil
}

// WatchExitFinalized is a free log subscription operation binding the contract event 0x0adb29b0831e081044cefe31155c1f2b2b85ad3613a480a5f901ee287addef55.
//
// Solidity: event ExitFinalized(uint160 indexed exitId)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchExitFinalized(opts *bind.WatchOpts, sink chan<- *PaymentExitGameExitFinalized, exitId []*big.Int) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "ExitFinalized", exitIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameExitFinalized)
				if err := _PaymentExitGame.contract.UnpackLog(event, "ExitFinalized", log); err != nil {
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

// ParseExitFinalized is a log parse operation binding the contract event 0x0adb29b0831e081044cefe31155c1f2b2b85ad3613a480a5f901ee287addef55.
//
// Solidity: event ExitFinalized(uint160 indexed exitId)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseExitFinalized(log types.Log) (*PaymentExitGameExitFinalized, error) {
	event := new(PaymentExitGameExitFinalized)
	if err := _PaymentExitGame.contract.UnpackLog(event, "ExitFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGameExitOmittedIterator is returned from FilterExitOmitted and is used to iterate over the raw logs and unpacked data for ExitOmitted events raised by the PaymentExitGame contract.
type PaymentExitGameExitOmittedIterator struct {
	Event *PaymentExitGameExitOmitted // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameExitOmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameExitOmitted)
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
		it.Event = new(PaymentExitGameExitOmitted)
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
func (it *PaymentExitGameExitOmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameExitOmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameExitOmitted represents a ExitOmitted event raised by the PaymentExitGame contract.
type PaymentExitGameExitOmitted struct {
	ExitId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitOmitted is a free log retrieval operation binding the contract event 0x91b69d9a0a91ef4f81d0232886cb1880ada4ccdadc8981a3fede280c284bc46a.
//
// Solidity: event ExitOmitted(uint160 indexed exitId)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterExitOmitted(opts *bind.FilterOpts, exitId []*big.Int) (*PaymentExitGameExitOmittedIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "ExitOmitted", exitIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameExitOmittedIterator{contract: _PaymentExitGame.contract, event: "ExitOmitted", logs: logs, sub: sub}, nil
}

// WatchExitOmitted is a free log subscription operation binding the contract event 0x91b69d9a0a91ef4f81d0232886cb1880ada4ccdadc8981a3fede280c284bc46a.
//
// Solidity: event ExitOmitted(uint160 indexed exitId)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchExitOmitted(opts *bind.WatchOpts, sink chan<- *PaymentExitGameExitOmitted, exitId []*big.Int) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "ExitOmitted", exitIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameExitOmitted)
				if err := _PaymentExitGame.contract.UnpackLog(event, "ExitOmitted", log); err != nil {
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

// ParseExitOmitted is a log parse operation binding the contract event 0x91b69d9a0a91ef4f81d0232886cb1880ada4ccdadc8981a3fede280c284bc46a.
//
// Solidity: event ExitOmitted(uint160 indexed exitId)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseExitOmitted(log types.Log) (*PaymentExitGameExitOmitted, error) {
	event := new(PaymentExitGameExitOmitted)
	if err := _PaymentExitGame.contract.UnpackLog(event, "ExitOmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGameExitStartedIterator is returned from FilterExitStarted and is used to iterate over the raw logs and unpacked data for ExitStarted events raised by the PaymentExitGame contract.
type PaymentExitGameExitStartedIterator struct {
	Event *PaymentExitGameExitStarted // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameExitStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameExitStarted)
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
		it.Event = new(PaymentExitGameExitStarted)
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
func (it *PaymentExitGameExitStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameExitStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameExitStarted represents a ExitStarted event raised by the PaymentExitGame contract.
type PaymentExitGameExitStarted struct {
	Owner  common.Address
	ExitId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitStarted is a free log retrieval operation binding the contract event 0xdd6f755cba05d0a420007aef6afc05e4889ab424505e2e440ecd1c434ba7082e.
//
// Solidity: event ExitStarted(address indexed owner, uint160 exitId)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterExitStarted(opts *bind.FilterOpts, owner []common.Address) (*PaymentExitGameExitStartedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "ExitStarted", ownerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameExitStartedIterator{contract: _PaymentExitGame.contract, event: "ExitStarted", logs: logs, sub: sub}, nil
}

// WatchExitStarted is a free log subscription operation binding the contract event 0xdd6f755cba05d0a420007aef6afc05e4889ab424505e2e440ecd1c434ba7082e.
//
// Solidity: event ExitStarted(address indexed owner, uint160 exitId)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchExitStarted(opts *bind.WatchOpts, sink chan<- *PaymentExitGameExitStarted, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "ExitStarted", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameExitStarted)
				if err := _PaymentExitGame.contract.UnpackLog(event, "ExitStarted", log); err != nil {
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

// ParseExitStarted is a log parse operation binding the contract event 0xdd6f755cba05d0a420007aef6afc05e4889ab424505e2e440ecd1c434ba7082e.
//
// Solidity: event ExitStarted(address indexed owner, uint160 exitId)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseExitStarted(log types.Log) (*PaymentExitGameExitStarted, error) {
	event := new(PaymentExitGameExitStarted)
	if err := _PaymentExitGame.contract.UnpackLog(event, "ExitStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGameIFEBondUpdatedIterator is returned from FilterIFEBondUpdated and is used to iterate over the raw logs and unpacked data for IFEBondUpdated events raised by the PaymentExitGame contract.
type PaymentExitGameIFEBondUpdatedIterator struct {
	Event *PaymentExitGameIFEBondUpdated // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameIFEBondUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameIFEBondUpdated)
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
		it.Event = new(PaymentExitGameIFEBondUpdated)
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
func (it *PaymentExitGameIFEBondUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameIFEBondUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameIFEBondUpdated represents a IFEBondUpdated event raised by the PaymentExitGame contract.
type PaymentExitGameIFEBondUpdated struct {
	BondSize *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIFEBondUpdated is a free log retrieval operation binding the contract event 0x63ccac61f03626a0e717e96da239ff8996745b6c4ddcaffdedc88d7ef938f5ec.
//
// Solidity: event IFEBondUpdated(uint128 bondSize)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterIFEBondUpdated(opts *bind.FilterOpts) (*PaymentExitGameIFEBondUpdatedIterator, error) {

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "IFEBondUpdated")
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameIFEBondUpdatedIterator{contract: _PaymentExitGame.contract, event: "IFEBondUpdated", logs: logs, sub: sub}, nil
}

// WatchIFEBondUpdated is a free log subscription operation binding the contract event 0x63ccac61f03626a0e717e96da239ff8996745b6c4ddcaffdedc88d7ef938f5ec.
//
// Solidity: event IFEBondUpdated(uint128 bondSize)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchIFEBondUpdated(opts *bind.WatchOpts, sink chan<- *PaymentExitGameIFEBondUpdated) (event.Subscription, error) {

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "IFEBondUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameIFEBondUpdated)
				if err := _PaymentExitGame.contract.UnpackLog(event, "IFEBondUpdated", log); err != nil {
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

// ParseIFEBondUpdated is a log parse operation binding the contract event 0x63ccac61f03626a0e717e96da239ff8996745b6c4ddcaffdedc88d7ef938f5ec.
//
// Solidity: event IFEBondUpdated(uint128 bondSize)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseIFEBondUpdated(log types.Log) (*PaymentExitGameIFEBondUpdated, error) {
	event := new(PaymentExitGameIFEBondUpdated)
	if err := _PaymentExitGame.contract.UnpackLog(event, "IFEBondUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGameInFlightExitChallengeRespondedIterator is returned from FilterInFlightExitChallengeResponded and is used to iterate over the raw logs and unpacked data for InFlightExitChallengeResponded events raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitChallengeRespondedIterator struct {
	Event *PaymentExitGameInFlightExitChallengeResponded // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameInFlightExitChallengeRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameInFlightExitChallengeResponded)
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
		it.Event = new(PaymentExitGameInFlightExitChallengeResponded)
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
func (it *PaymentExitGameInFlightExitChallengeRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameInFlightExitChallengeRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameInFlightExitChallengeResponded represents a InFlightExitChallengeResponded event raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitChallengeResponded struct {
	Challenger          common.Address
	TxHash              [32]byte
	ChallengeTxPosition *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitChallengeResponded is a free log retrieval operation binding the contract event 0x637cc4a7148767df19331a5c7dfb6d31f0a7e159a3dbb28a716be18c8c74f768.
//
// Solidity: event InFlightExitChallengeResponded(address challenger, bytes32 txHash, uint256 challengeTxPosition)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterInFlightExitChallengeResponded(opts *bind.FilterOpts) (*PaymentExitGameInFlightExitChallengeRespondedIterator, error) {

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "InFlightExitChallengeResponded")
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameInFlightExitChallengeRespondedIterator{contract: _PaymentExitGame.contract, event: "InFlightExitChallengeResponded", logs: logs, sub: sub}, nil
}

// WatchInFlightExitChallengeResponded is a free log subscription operation binding the contract event 0x637cc4a7148767df19331a5c7dfb6d31f0a7e159a3dbb28a716be18c8c74f768.
//
// Solidity: event InFlightExitChallengeResponded(address challenger, bytes32 txHash, uint256 challengeTxPosition)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchInFlightExitChallengeResponded(opts *bind.WatchOpts, sink chan<- *PaymentExitGameInFlightExitChallengeResponded) (event.Subscription, error) {

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "InFlightExitChallengeResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameInFlightExitChallengeResponded)
				if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitChallengeResponded", log); err != nil {
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

// ParseInFlightExitChallengeResponded is a log parse operation binding the contract event 0x637cc4a7148767df19331a5c7dfb6d31f0a7e159a3dbb28a716be18c8c74f768.
//
// Solidity: event InFlightExitChallengeResponded(address challenger, bytes32 txHash, uint256 challengeTxPosition)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseInFlightExitChallengeResponded(log types.Log) (*PaymentExitGameInFlightExitChallengeResponded, error) {
	event := new(PaymentExitGameInFlightExitChallengeResponded)
	if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitChallengeResponded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGameInFlightExitChallengedIterator is returned from FilterInFlightExitChallenged and is used to iterate over the raw logs and unpacked data for InFlightExitChallenged events raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitChallengedIterator struct {
	Event *PaymentExitGameInFlightExitChallenged // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameInFlightExitChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameInFlightExitChallenged)
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
		it.Event = new(PaymentExitGameInFlightExitChallenged)
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
func (it *PaymentExitGameInFlightExitChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameInFlightExitChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameInFlightExitChallenged represents a InFlightExitChallenged event raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitChallenged struct {
	Challenger          common.Address
	TxHash              [32]byte
	ChallengeTxPosition *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitChallenged is a free log retrieval operation binding the contract event 0x687401968e501bda2d2d6f880dd1a0a56ff50b1787185ee0b6f4c3fb9fc417ab.
//
// Solidity: event InFlightExitChallenged(address indexed challenger, bytes32 txHash, uint256 challengeTxPosition)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterInFlightExitChallenged(opts *bind.FilterOpts, challenger []common.Address) (*PaymentExitGameInFlightExitChallengedIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "InFlightExitChallenged", challengerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameInFlightExitChallengedIterator{contract: _PaymentExitGame.contract, event: "InFlightExitChallenged", logs: logs, sub: sub}, nil
}

// WatchInFlightExitChallenged is a free log subscription operation binding the contract event 0x687401968e501bda2d2d6f880dd1a0a56ff50b1787185ee0b6f4c3fb9fc417ab.
//
// Solidity: event InFlightExitChallenged(address indexed challenger, bytes32 txHash, uint256 challengeTxPosition)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchInFlightExitChallenged(opts *bind.WatchOpts, sink chan<- *PaymentExitGameInFlightExitChallenged, challenger []common.Address) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "InFlightExitChallenged", challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameInFlightExitChallenged)
				if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitChallenged", log); err != nil {
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

// ParseInFlightExitChallenged is a log parse operation binding the contract event 0x687401968e501bda2d2d6f880dd1a0a56ff50b1787185ee0b6f4c3fb9fc417ab.
//
// Solidity: event InFlightExitChallenged(address indexed challenger, bytes32 txHash, uint256 challengeTxPosition)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseInFlightExitChallenged(log types.Log) (*PaymentExitGameInFlightExitChallenged, error) {
	event := new(PaymentExitGameInFlightExitChallenged)
	if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitChallenged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGameInFlightExitInputBlockedIterator is returned from FilterInFlightExitInputBlocked and is used to iterate over the raw logs and unpacked data for InFlightExitInputBlocked events raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitInputBlockedIterator struct {
	Event *PaymentExitGameInFlightExitInputBlocked // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameInFlightExitInputBlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameInFlightExitInputBlocked)
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
		it.Event = new(PaymentExitGameInFlightExitInputBlocked)
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
func (it *PaymentExitGameInFlightExitInputBlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameInFlightExitInputBlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameInFlightExitInputBlocked represents a InFlightExitInputBlocked event raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitInputBlocked struct {
	Challenger common.Address
	TxHash     [32]byte
	InputIndex uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitInputBlocked is a free log retrieval operation binding the contract event 0x4794045885b2e249895bc730deb7caedc2e6b4b49d56ccd696b69d1f4944b8ec.
//
// Solidity: event InFlightExitInputBlocked(address indexed challenger, bytes32 txHash, uint16 inputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterInFlightExitInputBlocked(opts *bind.FilterOpts, challenger []common.Address) (*PaymentExitGameInFlightExitInputBlockedIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "InFlightExitInputBlocked", challengerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameInFlightExitInputBlockedIterator{contract: _PaymentExitGame.contract, event: "InFlightExitInputBlocked", logs: logs, sub: sub}, nil
}

// WatchInFlightExitInputBlocked is a free log subscription operation binding the contract event 0x4794045885b2e249895bc730deb7caedc2e6b4b49d56ccd696b69d1f4944b8ec.
//
// Solidity: event InFlightExitInputBlocked(address indexed challenger, bytes32 txHash, uint16 inputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchInFlightExitInputBlocked(opts *bind.WatchOpts, sink chan<- *PaymentExitGameInFlightExitInputBlocked, challenger []common.Address) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "InFlightExitInputBlocked", challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameInFlightExitInputBlocked)
				if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitInputBlocked", log); err != nil {
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

// ParseInFlightExitInputBlocked is a log parse operation binding the contract event 0x4794045885b2e249895bc730deb7caedc2e6b4b49d56ccd696b69d1f4944b8ec.
//
// Solidity: event InFlightExitInputBlocked(address indexed challenger, bytes32 txHash, uint16 inputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseInFlightExitInputBlocked(log types.Log) (*PaymentExitGameInFlightExitInputBlocked, error) {
	event := new(PaymentExitGameInFlightExitInputBlocked)
	if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitInputBlocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGameInFlightExitInputPiggybackedIterator is returned from FilterInFlightExitInputPiggybacked and is used to iterate over the raw logs and unpacked data for InFlightExitInputPiggybacked events raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitInputPiggybackedIterator struct {
	Event *PaymentExitGameInFlightExitInputPiggybacked // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameInFlightExitInputPiggybackedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameInFlightExitInputPiggybacked)
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
		it.Event = new(PaymentExitGameInFlightExitInputPiggybacked)
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
func (it *PaymentExitGameInFlightExitInputPiggybackedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameInFlightExitInputPiggybackedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameInFlightExitInputPiggybacked represents a InFlightExitInputPiggybacked event raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitInputPiggybacked struct {
	ExitTarget common.Address
	TxHash     [32]byte
	InputIndex uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitInputPiggybacked is a free log retrieval operation binding the contract event 0xa93c0e9b202feaf554acf6ef1185b898c9f214da16e51740b06b5f7487b018e5.
//
// Solidity: event InFlightExitInputPiggybacked(address indexed exitTarget, bytes32 txHash, uint16 inputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterInFlightExitInputPiggybacked(opts *bind.FilterOpts, exitTarget []common.Address) (*PaymentExitGameInFlightExitInputPiggybackedIterator, error) {

	var exitTargetRule []interface{}
	for _, exitTargetItem := range exitTarget {
		exitTargetRule = append(exitTargetRule, exitTargetItem)
	}

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "InFlightExitInputPiggybacked", exitTargetRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameInFlightExitInputPiggybackedIterator{contract: _PaymentExitGame.contract, event: "InFlightExitInputPiggybacked", logs: logs, sub: sub}, nil
}

// WatchInFlightExitInputPiggybacked is a free log subscription operation binding the contract event 0xa93c0e9b202feaf554acf6ef1185b898c9f214da16e51740b06b5f7487b018e5.
//
// Solidity: event InFlightExitInputPiggybacked(address indexed exitTarget, bytes32 txHash, uint16 inputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchInFlightExitInputPiggybacked(opts *bind.WatchOpts, sink chan<- *PaymentExitGameInFlightExitInputPiggybacked, exitTarget []common.Address) (event.Subscription, error) {

	var exitTargetRule []interface{}
	for _, exitTargetItem := range exitTarget {
		exitTargetRule = append(exitTargetRule, exitTargetItem)
	}

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "InFlightExitInputPiggybacked", exitTargetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameInFlightExitInputPiggybacked)
				if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitInputPiggybacked", log); err != nil {
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

// ParseInFlightExitInputPiggybacked is a log parse operation binding the contract event 0xa93c0e9b202feaf554acf6ef1185b898c9f214da16e51740b06b5f7487b018e5.
//
// Solidity: event InFlightExitInputPiggybacked(address indexed exitTarget, bytes32 txHash, uint16 inputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseInFlightExitInputPiggybacked(log types.Log) (*PaymentExitGameInFlightExitInputPiggybacked, error) {
	event := new(PaymentExitGameInFlightExitInputPiggybacked)
	if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitInputPiggybacked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGameInFlightExitInputWithdrawnIterator is returned from FilterInFlightExitInputWithdrawn and is used to iterate over the raw logs and unpacked data for InFlightExitInputWithdrawn events raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitInputWithdrawnIterator struct {
	Event *PaymentExitGameInFlightExitInputWithdrawn // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameInFlightExitInputWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameInFlightExitInputWithdrawn)
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
		it.Event = new(PaymentExitGameInFlightExitInputWithdrawn)
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
func (it *PaymentExitGameInFlightExitInputWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameInFlightExitInputWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameInFlightExitInputWithdrawn represents a InFlightExitInputWithdrawn event raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitInputWithdrawn struct {
	ExitId     *big.Int
	InputIndex uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitInputWithdrawn is a free log retrieval operation binding the contract event 0x4446ec118df0062851669a051944a56f276ef134203ac88cbb6032bd10b6aeec.
//
// Solidity: event InFlightExitInputWithdrawn(uint160 indexed exitId, uint16 inputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterInFlightExitInputWithdrawn(opts *bind.FilterOpts, exitId []*big.Int) (*PaymentExitGameInFlightExitInputWithdrawnIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "InFlightExitInputWithdrawn", exitIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameInFlightExitInputWithdrawnIterator{contract: _PaymentExitGame.contract, event: "InFlightExitInputWithdrawn", logs: logs, sub: sub}, nil
}

// WatchInFlightExitInputWithdrawn is a free log subscription operation binding the contract event 0x4446ec118df0062851669a051944a56f276ef134203ac88cbb6032bd10b6aeec.
//
// Solidity: event InFlightExitInputWithdrawn(uint160 indexed exitId, uint16 inputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchInFlightExitInputWithdrawn(opts *bind.WatchOpts, sink chan<- *PaymentExitGameInFlightExitInputWithdrawn, exitId []*big.Int) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "InFlightExitInputWithdrawn", exitIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameInFlightExitInputWithdrawn)
				if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitInputWithdrawn", log); err != nil {
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

// ParseInFlightExitInputWithdrawn is a log parse operation binding the contract event 0x4446ec118df0062851669a051944a56f276ef134203ac88cbb6032bd10b6aeec.
//
// Solidity: event InFlightExitInputWithdrawn(uint160 indexed exitId, uint16 inputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseInFlightExitInputWithdrawn(log types.Log) (*PaymentExitGameInFlightExitInputWithdrawn, error) {
	event := new(PaymentExitGameInFlightExitInputWithdrawn)
	if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitInputWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGameInFlightExitOmittedIterator is returned from FilterInFlightExitOmitted and is used to iterate over the raw logs and unpacked data for InFlightExitOmitted events raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitOmittedIterator struct {
	Event *PaymentExitGameInFlightExitOmitted // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameInFlightExitOmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameInFlightExitOmitted)
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
		it.Event = new(PaymentExitGameInFlightExitOmitted)
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
func (it *PaymentExitGameInFlightExitOmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameInFlightExitOmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameInFlightExitOmitted represents a InFlightExitOmitted event raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitOmitted struct {
	ExitId *big.Int
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitOmitted is a free log retrieval operation binding the contract event 0x438fb52bbaecea88ef59bcef337f0f8d5ba28dc74f727c0c176f81ff2c7ab81f.
//
// Solidity: event InFlightExitOmitted(uint160 indexed exitId, address token)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterInFlightExitOmitted(opts *bind.FilterOpts, exitId []*big.Int) (*PaymentExitGameInFlightExitOmittedIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "InFlightExitOmitted", exitIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameInFlightExitOmittedIterator{contract: _PaymentExitGame.contract, event: "InFlightExitOmitted", logs: logs, sub: sub}, nil
}

// WatchInFlightExitOmitted is a free log subscription operation binding the contract event 0x438fb52bbaecea88ef59bcef337f0f8d5ba28dc74f727c0c176f81ff2c7ab81f.
//
// Solidity: event InFlightExitOmitted(uint160 indexed exitId, address token)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchInFlightExitOmitted(opts *bind.WatchOpts, sink chan<- *PaymentExitGameInFlightExitOmitted, exitId []*big.Int) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "InFlightExitOmitted", exitIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameInFlightExitOmitted)
				if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitOmitted", log); err != nil {
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

// ParseInFlightExitOmitted is a log parse operation binding the contract event 0x438fb52bbaecea88ef59bcef337f0f8d5ba28dc74f727c0c176f81ff2c7ab81f.
//
// Solidity: event InFlightExitOmitted(uint160 indexed exitId, address token)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseInFlightExitOmitted(log types.Log) (*PaymentExitGameInFlightExitOmitted, error) {
	event := new(PaymentExitGameInFlightExitOmitted)
	if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitOmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGameInFlightExitOutputBlockedIterator is returned from FilterInFlightExitOutputBlocked and is used to iterate over the raw logs and unpacked data for InFlightExitOutputBlocked events raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitOutputBlockedIterator struct {
	Event *PaymentExitGameInFlightExitOutputBlocked // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameInFlightExitOutputBlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameInFlightExitOutputBlocked)
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
		it.Event = new(PaymentExitGameInFlightExitOutputBlocked)
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
func (it *PaymentExitGameInFlightExitOutputBlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameInFlightExitOutputBlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameInFlightExitOutputBlocked represents a InFlightExitOutputBlocked event raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitOutputBlocked struct {
	Challenger  common.Address
	IfeTxHash   [32]byte
	OutputIndex uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitOutputBlocked is a free log retrieval operation binding the contract event 0xcbe8dad2e7fcbfe0dcba2f9b2e44f122c66cd26dc0808a0f7e9ec41e4fe285bf.
//
// Solidity: event InFlightExitOutputBlocked(address indexed challenger, bytes32 ifeTxHash, uint16 outputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterInFlightExitOutputBlocked(opts *bind.FilterOpts, challenger []common.Address) (*PaymentExitGameInFlightExitOutputBlockedIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "InFlightExitOutputBlocked", challengerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameInFlightExitOutputBlockedIterator{contract: _PaymentExitGame.contract, event: "InFlightExitOutputBlocked", logs: logs, sub: sub}, nil
}

// WatchInFlightExitOutputBlocked is a free log subscription operation binding the contract event 0xcbe8dad2e7fcbfe0dcba2f9b2e44f122c66cd26dc0808a0f7e9ec41e4fe285bf.
//
// Solidity: event InFlightExitOutputBlocked(address indexed challenger, bytes32 ifeTxHash, uint16 outputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchInFlightExitOutputBlocked(opts *bind.WatchOpts, sink chan<- *PaymentExitGameInFlightExitOutputBlocked, challenger []common.Address) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "InFlightExitOutputBlocked", challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameInFlightExitOutputBlocked)
				if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitOutputBlocked", log); err != nil {
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

// ParseInFlightExitOutputBlocked is a log parse operation binding the contract event 0xcbe8dad2e7fcbfe0dcba2f9b2e44f122c66cd26dc0808a0f7e9ec41e4fe285bf.
//
// Solidity: event InFlightExitOutputBlocked(address indexed challenger, bytes32 ifeTxHash, uint16 outputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseInFlightExitOutputBlocked(log types.Log) (*PaymentExitGameInFlightExitOutputBlocked, error) {
	event := new(PaymentExitGameInFlightExitOutputBlocked)
	if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitOutputBlocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGameInFlightExitOutputPiggybackedIterator is returned from FilterInFlightExitOutputPiggybacked and is used to iterate over the raw logs and unpacked data for InFlightExitOutputPiggybacked events raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitOutputPiggybackedIterator struct {
	Event *PaymentExitGameInFlightExitOutputPiggybacked // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameInFlightExitOutputPiggybackedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameInFlightExitOutputPiggybacked)
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
		it.Event = new(PaymentExitGameInFlightExitOutputPiggybacked)
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
func (it *PaymentExitGameInFlightExitOutputPiggybackedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameInFlightExitOutputPiggybackedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameInFlightExitOutputPiggybacked represents a InFlightExitOutputPiggybacked event raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitOutputPiggybacked struct {
	ExitTarget  common.Address
	TxHash      [32]byte
	OutputIndex uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitOutputPiggybacked is a free log retrieval operation binding the contract event 0x6ecd8e79a5f67f6c12b54371ada2ffb41bc128c61d9ac1e969f0aa2aca46cd78.
//
// Solidity: event InFlightExitOutputPiggybacked(address indexed exitTarget, bytes32 txHash, uint16 outputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterInFlightExitOutputPiggybacked(opts *bind.FilterOpts, exitTarget []common.Address) (*PaymentExitGameInFlightExitOutputPiggybackedIterator, error) {

	var exitTargetRule []interface{}
	for _, exitTargetItem := range exitTarget {
		exitTargetRule = append(exitTargetRule, exitTargetItem)
	}

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "InFlightExitOutputPiggybacked", exitTargetRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameInFlightExitOutputPiggybackedIterator{contract: _PaymentExitGame.contract, event: "InFlightExitOutputPiggybacked", logs: logs, sub: sub}, nil
}

// WatchInFlightExitOutputPiggybacked is a free log subscription operation binding the contract event 0x6ecd8e79a5f67f6c12b54371ada2ffb41bc128c61d9ac1e969f0aa2aca46cd78.
//
// Solidity: event InFlightExitOutputPiggybacked(address indexed exitTarget, bytes32 txHash, uint16 outputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchInFlightExitOutputPiggybacked(opts *bind.WatchOpts, sink chan<- *PaymentExitGameInFlightExitOutputPiggybacked, exitTarget []common.Address) (event.Subscription, error) {

	var exitTargetRule []interface{}
	for _, exitTargetItem := range exitTarget {
		exitTargetRule = append(exitTargetRule, exitTargetItem)
	}

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "InFlightExitOutputPiggybacked", exitTargetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameInFlightExitOutputPiggybacked)
				if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitOutputPiggybacked", log); err != nil {
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

// ParseInFlightExitOutputPiggybacked is a log parse operation binding the contract event 0x6ecd8e79a5f67f6c12b54371ada2ffb41bc128c61d9ac1e969f0aa2aca46cd78.
//
// Solidity: event InFlightExitOutputPiggybacked(address indexed exitTarget, bytes32 txHash, uint16 outputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseInFlightExitOutputPiggybacked(log types.Log) (*PaymentExitGameInFlightExitOutputPiggybacked, error) {
	event := new(PaymentExitGameInFlightExitOutputPiggybacked)
	if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitOutputPiggybacked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGameInFlightExitOutputWithdrawnIterator is returned from FilterInFlightExitOutputWithdrawn and is used to iterate over the raw logs and unpacked data for InFlightExitOutputWithdrawn events raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitOutputWithdrawnIterator struct {
	Event *PaymentExitGameInFlightExitOutputWithdrawn // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameInFlightExitOutputWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameInFlightExitOutputWithdrawn)
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
		it.Event = new(PaymentExitGameInFlightExitOutputWithdrawn)
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
func (it *PaymentExitGameInFlightExitOutputWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameInFlightExitOutputWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameInFlightExitOutputWithdrawn represents a InFlightExitOutputWithdrawn event raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitOutputWithdrawn struct {
	ExitId      *big.Int
	OutputIndex uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitOutputWithdrawn is a free log retrieval operation binding the contract event 0xa241c6deaf193e53a1b002d779e4f247bf5d57ba0be5a753e628dfcee645a4f7.
//
// Solidity: event InFlightExitOutputWithdrawn(uint160 indexed exitId, uint16 outputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterInFlightExitOutputWithdrawn(opts *bind.FilterOpts, exitId []*big.Int) (*PaymentExitGameInFlightExitOutputWithdrawnIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "InFlightExitOutputWithdrawn", exitIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameInFlightExitOutputWithdrawnIterator{contract: _PaymentExitGame.contract, event: "InFlightExitOutputWithdrawn", logs: logs, sub: sub}, nil
}

// WatchInFlightExitOutputWithdrawn is a free log subscription operation binding the contract event 0xa241c6deaf193e53a1b002d779e4f247bf5d57ba0be5a753e628dfcee645a4f7.
//
// Solidity: event InFlightExitOutputWithdrawn(uint160 indexed exitId, uint16 outputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchInFlightExitOutputWithdrawn(opts *bind.WatchOpts, sink chan<- *PaymentExitGameInFlightExitOutputWithdrawn, exitId []*big.Int) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "InFlightExitOutputWithdrawn", exitIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameInFlightExitOutputWithdrawn)
				if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitOutputWithdrawn", log); err != nil {
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

// ParseInFlightExitOutputWithdrawn is a log parse operation binding the contract event 0xa241c6deaf193e53a1b002d779e4f247bf5d57ba0be5a753e628dfcee645a4f7.
//
// Solidity: event InFlightExitOutputWithdrawn(uint160 indexed exitId, uint16 outputIndex)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseInFlightExitOutputWithdrawn(log types.Log) (*PaymentExitGameInFlightExitOutputWithdrawn, error) {
	event := new(PaymentExitGameInFlightExitOutputWithdrawn)
	if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitOutputWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGameInFlightExitStartedIterator is returned from FilterInFlightExitStarted and is used to iterate over the raw logs and unpacked data for InFlightExitStarted events raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitStartedIterator struct {
	Event *PaymentExitGameInFlightExitStarted // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameInFlightExitStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameInFlightExitStarted)
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
		it.Event = new(PaymentExitGameInFlightExitStarted)
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
func (it *PaymentExitGameInFlightExitStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameInFlightExitStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameInFlightExitStarted represents a InFlightExitStarted event raised by the PaymentExitGame contract.
type PaymentExitGameInFlightExitStarted struct {
	Initiator common.Address
	TxHash    [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitStarted is a free log retrieval operation binding the contract event 0xd5f1fe9d48880b57daa227004b16d320c0eb885d6c49d472d54c16a05fa3179e.
//
// Solidity: event InFlightExitStarted(address indexed initiator, bytes32 txHash)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterInFlightExitStarted(opts *bind.FilterOpts, initiator []common.Address) (*PaymentExitGameInFlightExitStartedIterator, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "InFlightExitStarted", initiatorRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameInFlightExitStartedIterator{contract: _PaymentExitGame.contract, event: "InFlightExitStarted", logs: logs, sub: sub}, nil
}

// WatchInFlightExitStarted is a free log subscription operation binding the contract event 0xd5f1fe9d48880b57daa227004b16d320c0eb885d6c49d472d54c16a05fa3179e.
//
// Solidity: event InFlightExitStarted(address indexed initiator, bytes32 txHash)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchInFlightExitStarted(opts *bind.WatchOpts, sink chan<- *PaymentExitGameInFlightExitStarted, initiator []common.Address) (event.Subscription, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "InFlightExitStarted", initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameInFlightExitStarted)
				if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitStarted", log); err != nil {
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

// ParseInFlightExitStarted is a log parse operation binding the contract event 0xd5f1fe9d48880b57daa227004b16d320c0eb885d6c49d472d54c16a05fa3179e.
//
// Solidity: event InFlightExitStarted(address indexed initiator, bytes32 txHash)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseInFlightExitStarted(log types.Log) (*PaymentExitGameInFlightExitStarted, error) {
	event := new(PaymentExitGameInFlightExitStarted)
	if err := _PaymentExitGame.contract.UnpackLog(event, "InFlightExitStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGamePiggybackBondUpdatedIterator is returned from FilterPiggybackBondUpdated and is used to iterate over the raw logs and unpacked data for PiggybackBondUpdated events raised by the PaymentExitGame contract.
type PaymentExitGamePiggybackBondUpdatedIterator struct {
	Event *PaymentExitGamePiggybackBondUpdated // Event containing the contract specifics and raw log

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
func (it *PaymentExitGamePiggybackBondUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGamePiggybackBondUpdated)
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
		it.Event = new(PaymentExitGamePiggybackBondUpdated)
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
func (it *PaymentExitGamePiggybackBondUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGamePiggybackBondUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGamePiggybackBondUpdated represents a PiggybackBondUpdated event raised by the PaymentExitGame contract.
type PaymentExitGamePiggybackBondUpdated struct {
	BondSize *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPiggybackBondUpdated is a free log retrieval operation binding the contract event 0x30ce291291dca42b8168ab211fafc36d3d4868292ed3259874ed00864432e20e.
//
// Solidity: event PiggybackBondUpdated(uint128 bondSize)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterPiggybackBondUpdated(opts *bind.FilterOpts) (*PaymentExitGamePiggybackBondUpdatedIterator, error) {

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "PiggybackBondUpdated")
	if err != nil {
		return nil, err
	}
	return &PaymentExitGamePiggybackBondUpdatedIterator{contract: _PaymentExitGame.contract, event: "PiggybackBondUpdated", logs: logs, sub: sub}, nil
}

// WatchPiggybackBondUpdated is a free log subscription operation binding the contract event 0x30ce291291dca42b8168ab211fafc36d3d4868292ed3259874ed00864432e20e.
//
// Solidity: event PiggybackBondUpdated(uint128 bondSize)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchPiggybackBondUpdated(opts *bind.WatchOpts, sink chan<- *PaymentExitGamePiggybackBondUpdated) (event.Subscription, error) {

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "PiggybackBondUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGamePiggybackBondUpdated)
				if err := _PaymentExitGame.contract.UnpackLog(event, "PiggybackBondUpdated", log); err != nil {
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

// ParsePiggybackBondUpdated is a log parse operation binding the contract event 0x30ce291291dca42b8168ab211fafc36d3d4868292ed3259874ed00864432e20e.
//
// Solidity: event PiggybackBondUpdated(uint128 bondSize)
func (_PaymentExitGame *PaymentExitGameFilterer) ParsePiggybackBondUpdated(log types.Log) (*PaymentExitGamePiggybackBondUpdated, error) {
	event := new(PaymentExitGamePiggybackBondUpdated)
	if err := _PaymentExitGame.contract.UnpackLog(event, "PiggybackBondUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitGameStandardExitBondUpdatedIterator is returned from FilterStandardExitBondUpdated and is used to iterate over the raw logs and unpacked data for StandardExitBondUpdated events raised by the PaymentExitGame contract.
type PaymentExitGameStandardExitBondUpdatedIterator struct {
	Event *PaymentExitGameStandardExitBondUpdated // Event containing the contract specifics and raw log

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
func (it *PaymentExitGameStandardExitBondUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitGameStandardExitBondUpdated)
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
		it.Event = new(PaymentExitGameStandardExitBondUpdated)
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
func (it *PaymentExitGameStandardExitBondUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitGameStandardExitBondUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitGameStandardExitBondUpdated represents a StandardExitBondUpdated event raised by the PaymentExitGame contract.
type PaymentExitGameStandardExitBondUpdated struct {
	BondSize *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStandardExitBondUpdated is a free log retrieval operation binding the contract event 0xd21ccbbae3d29826d53310efdc993c64e27496dd31694b292f1a120344ec37ab.
//
// Solidity: event StandardExitBondUpdated(uint128 bondSize)
func (_PaymentExitGame *PaymentExitGameFilterer) FilterStandardExitBondUpdated(opts *bind.FilterOpts) (*PaymentExitGameStandardExitBondUpdatedIterator, error) {

	logs, sub, err := _PaymentExitGame.contract.FilterLogs(opts, "StandardExitBondUpdated")
	if err != nil {
		return nil, err
	}
	return &PaymentExitGameStandardExitBondUpdatedIterator{contract: _PaymentExitGame.contract, event: "StandardExitBondUpdated", logs: logs, sub: sub}, nil
}

// WatchStandardExitBondUpdated is a free log subscription operation binding the contract event 0xd21ccbbae3d29826d53310efdc993c64e27496dd31694b292f1a120344ec37ab.
//
// Solidity: event StandardExitBondUpdated(uint128 bondSize)
func (_PaymentExitGame *PaymentExitGameFilterer) WatchStandardExitBondUpdated(opts *bind.WatchOpts, sink chan<- *PaymentExitGameStandardExitBondUpdated) (event.Subscription, error) {

	logs, sub, err := _PaymentExitGame.contract.WatchLogs(opts, "StandardExitBondUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitGameStandardExitBondUpdated)
				if err := _PaymentExitGame.contract.UnpackLog(event, "StandardExitBondUpdated", log); err != nil {
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

// ParseStandardExitBondUpdated is a log parse operation binding the contract event 0xd21ccbbae3d29826d53310efdc993c64e27496dd31694b292f1a120344ec37ab.
//
// Solidity: event StandardExitBondUpdated(uint128 bondSize)
func (_PaymentExitGame *PaymentExitGameFilterer) ParseStandardExitBondUpdated(log types.Log) (*PaymentExitGameStandardExitBondUpdated, error) {
	event := new(PaymentExitGameStandardExitBondUpdated)
	if err := _PaymentExitGame.contract.UnpackLog(event, "StandardExitBondUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

