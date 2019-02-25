package util

import (
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"strings"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type Inner struct {
	Values []inputDeposit
	Values2 []interface{}
}

type inputDeposit struct {
	Txindex  uint    `json:"txindex"`
	Oindex   uint    `json:"oindex"`
	Blknum   uint    `json:"blknum"`
}

type InputTwo struct {
	OwnerAddress common.Address
    Currency common.Address
    Amount uint64
}

type second struct {
	Currency1 common.Address
	Currency2 common.Address
	Value uint
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

//sign a transaction with a key
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

// generate Account - Public and Privatekey
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

// filter our key string with 0x
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

// Build the RLP encoded input to the smart contract
// that deposit will accept.
func BuildRLPInput(address string, value string) []byte {
	// [[[0,0,0],[0,0,0],[0,0,0],[0,0,0]],[[alice_raw, eth_raw, 10], [eth_raw, eth_raw, 0], [eth_raw, eth_raw, 0], [eth_raw, eth_raw, 0]]]
	// TODO(jbunce): This works but it's a bloody mess. Need to make it pretty.
	v := make([]inputDeposit, 0)

	test := Inner{}

	NULL_INPUT := inputDeposit{Blknum: 0, Txindex: 0, Oindex: 0}
	v = append(v, NULL_INPUT)
	v = append(v, NULL_INPUT)
	v = append(v, NULL_INPUT)
	v = append(v, NULL_INPUT)

	test.Values = v

	cur := common.HexToAddress("0000000000000000000000000000000000000000")
	amount, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	t3 := InputTwo{}
	t3.OwnerAddress = common.HexToAddress(address)
	t3.Currency = cur
	t3.Amount = uint64(amount)

	s1 := second{Currency1: cur, Currency2: cur, Value: 0}
	s2 := second{Currency1: cur, Currency2: cur, Value: 0}
	s3 := second{Currency1: cur, Currency2: cur, Value: 0}

	arr := make([]interface{}, 0)
	arr = append(arr, t3)
	arr = append(arr, s1)
	arr = append(arr, s2)
	arr = append(arr, s3)

	test.Values2 = arr

	rlpEncoded, rerr := rlp.EncodeToBytes(test)
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
