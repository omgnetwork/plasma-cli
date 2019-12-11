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
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

const (
	EthCurrency     = "0x0000000000000000000000000000000000000000"
	DefaultMetadata = "0x0000000000000000000000000000000000000000000000000000000000000000"
)


//deposit parent for ALD
type DepositTransaction struct {
	OutputType  uint
	UTXOInputs  []InputDeposit
	UTXOOutputs []interface{}
	MetaData common.Hash
}


type InputDeposit struct {
	Txindex uint `json:"txindex"`
	Oindex  uint `json:"oindex"`
	Blknum  uint `json:"blknum"`
}

type DepositOutput struct {
	OutputType uint64
	OutputGuard common.Address
	Token common.Address
	Amount uint64
}


//sign an already hashed tx bytes
func SignHash(hashed []byte, privateKeys []string) ([][]byte, error) {
	var sigs [][]byte
	for _, pk := range privateKeys {
		priv, err := crypto.HexToECDSA(FilterZeroX(pk))
		if err != nil {
			return nil, err
		}
		//sign the transaction
		signature, err := crypto.Sign(hashed, priv)
		if err != nil {
			return nil, err
		}
		//adding 27 to last byte, because ethereum
		copy(signature[64:], []uint8{signature[64] + 27})
		sigs = append(sigs, signature)
	}

	return sigs, nil
}

// Generate Account - Public and Privatekey
func GenerateAccount() (string, string) {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	privateKey := hex.EncodeToString(key.D.Bytes())
	log.Infof("\n Address: %s \n Privatekey: 0x%s ", address, strings.ToUpper(privateKey))
	return address, privateKey
}

// Derive Address from Private Key
func DeriveAddress(p string) string {
	key, err := crypto.HexToECDSA(FilterZeroX(p))
	if err != nil {
		log.Fatal(err)
	}
	addr := crypto.PubkeyToAddress(key.PublicKey).Hex()
	return addr
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

// Build a deposit transaction, consists of empty inputs
// and a single output UTXO of a specified address, currency, value and transaction type 
func BuildALDRLPInput(address, currency string, value, txtype uint64) []byte {
	deposit := DepositTransaction{OutputType: uint(txtype), MetaData: common.HexToHash(DefaultMetadata)}

	cur := common.HexToAddress(EthCurrency)
	// create a single ownership output 
	ownership := DepositOutput{}
	ownership.OutputGuard = common.HexToAddress(address)
	ownership.Token = cur
	ownership.Amount = value
	ownership.OutputType = txtype
	// we skip making inputs in ALD
	UTXOOutputs := make([]interface{}, 0)
	UTXOOutputs = append(UTXOOutputs, ownership)
	deposit.UTXOOutputs = UTXOOutputs
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
