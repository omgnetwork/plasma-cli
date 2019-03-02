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

package util

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type DepositParent struct {
	UTXOInputs  []InputDeposit
	UTXOOutputs []interface{}
}

type InputDeposit struct {
	Txindex uint `json:"txindex"`
	Oindex  uint `json:"oindex"`
	Blknum  uint `json:"blknum"`
}

type DepositOwnership struct {
	OwnerAddress common.Address
	Currency     common.Address
	Amount       uint64
}

type OutputUTXO struct {
	Currency1 common.Address
	Currency2 common.Address
	Value     uint
}

type WatcherUTXOsFromAddress struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    []struct {
		UtxoPos  int    `json:"utxo_pos"`
		Txindex  int    `json:"txindex"`
		Owner    string `json:"owner"`
		Oindex   int    `json:"oindex"`
		Currency string `json:"currency"`
		Blknum   int    `json:"blknum"`
		Amount   int    `json:"amount"`
	} `json:"data"`
}

// Sign a transaction with a key
func SignTransaction(unsignedTx string, privateKey string) []byte {
	//hash the unsignedTx struct
	unsignedTxBytes, err := hex.DecodeString(FilterZeroX(unsignedTx))
	if err != nil {
		log.Fatal(err)
	}
	hashed := crypto.Keccak256(unsignedTxBytes)
	priv, err := crypto.HexToECDSA(FilterZeroX(privateKey))
	if err != nil {
		log.Fatal(err)
	}
	//sign the transaction
	signature, err := crypto.Sign(hashed, priv)
	if err != nil {
		log.Fatal(err)
	}
	//adding 27 to last byte, because ethereum
	copy(signature[64:], []uint8{signature[64] + 27})

	log.Info("transaction signed: ", signature)
	return signature
}

// Generate Account - Public and Privatekey
func GenerateAccount() {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	privateKey := hex.EncodeToString(key.D.Bytes())
	log.Info("Address is ", address)
	log.Info("Privatekey is ", privateKey)
}

// Make strings suitable for hex encoding
func RemoveLeadingZeroX(item string) string {
	cleanedString := strings.Replace(item, "0x", "", -1)
	return cleanedString
}

// Filter our key string with 0x
func FilterZeroX(item string) string {
	if strings.Contains(item, "0x") {
		return RemoveLeadingZeroX(item)
	} else {
		return item
	}
}

func ConvertStringToInt(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal("Could not convert value")
	}
	return i
}

func convertStringToFloat64(value string) float64 {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatal("Could not convert value")
	}
	return f
}

// Displays the UTXO ownership data in a human friendly format
func DisplayUTXOS(u WatcherUTXOsFromAddress) {
	for _, value := range u.Data {
		log.Info("Owner: ", value.Owner)
		log.Info("Amount: ", value.Amount)
		log.Info("Block Number: ", value.Blknum)
		log.Info("Transaction Index: ", value.Txindex)
		log.Info("Output Index: ", value.Oindex)
		log.Info("Currency: ", value.Currency)
		log.Info("UTXO Position: ", value.UtxoPos)
	}
}

// Build the RLP encoded input to the smart contract that
// deposit() will accept. The format of the UTXO inputs, outputs,
// and ownership data is critical. If any value is incorrect the
// deposit will be rejected by the smart contract and transaction
// reversed. The correct format is:
// [[[0,0,0],[0,0,0],[0,0,0],[0,0,0]],[[owner_public, eth_raw, 10], [eth_raw, eth_raw, 0], [eth_raw, eth_raw, 0], [eth_raw, eth_raw, 0]]]
// where eth_raw is 20 bytes of 0.
func BuildRLPInput(address string, value uint64) []byte {
	depositUTXOPositions := make([]InputDeposit, 0)
	deposit := DepositParent{}

	NULL_INPUT := InputDeposit{Blknum: 0, Txindex: 0, Oindex: 0}
	for i := 0; i <= 4; i++ {
		depositUTXOPositions = append(depositUTXOPositions, NULL_INPUT)
	}
	deposit.UTXOInputs = depositUTXOPositions

	// TODO(jbunce): Allow deposits from ERC20 tokens
	cur := common.HexToAddress("0000000000000000000000000000000000000000")

	// Define the UTXO ownership and currency of the deposit
	ownership := DepositOwnership{}
	ownership.OwnerAddress = common.HexToAddress(address)
	ownership.Currency = cur
	ownership.Amount = value

	UTXOOutputs := make([]interface{}, 0)
	UTXOOutputs = append(UTXOOutputs, ownership)

	NULL_OUTPUT := OutputUTXO{Currency1: cur, Currency2: cur, Value: 0}
	for i := 0; i <= 3; i++ {
		UTXOOutputs = append(UTXOOutputs, NULL_OUTPUT)
	}
	deposit.UTXOOutputs = UTXOOutputs

	// The actual RLP encoding happens here
	rlpEncoded, rerr := rlp.EncodeToBytes(deposit)
	if rerr != nil {
		log.Fatal(rerr)
	}

	return rlpEncoded
}

// Add the full time include timezone into log messages
// INFO[2019-01-31T16:38:57+07:00]
func LogFormatter() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)
}
