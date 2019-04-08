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

package plasma

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/omisego/plasma-cli/rootchain"
	"github.com/omisego/plasma-cli/util"
	log "github.com/sirupsen/logrus"
)

type PlasmaDeposit struct {
	PrivateKey string
	Client     string
	Contract   string
	Amount     uint64
	Owner      string
	Currency   string
}

type transactionSuccessResponse struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Blknum  int    `json:"blknum"`
		Txindex int    `json:"txindex"`
		Txhash  string `json:"txhash"`
	} `json:"data"`
}

type transactionFailureResponse struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Object      string `json:"object"`
		Code        string `json:"code"`
		Description string `json:"description"`
		Messages    struct {
			ErrorKey string `json:"error_key"`
		} `json:"messages"`
	} `json:"data"`
}

type blockNumber struct {
	Hash string `json:"hash"`
}

type blockNumberError struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Object   string `json:"object"`
		Messages struct {
			ValidationError struct {
				Validator string `json:"validator"`
				Parameter string `json:"parameter"`
			} `json:"validation_error"`
		} `json:"messages"`
		Description string `json:"description"`
		Code        string `json:"code"`
	} `json:"data"`
}

type ProcessExit struct {
	Contract   string
	PrivateKey string
	Token      string
	Client     string
}

type StandardExit struct {
	UtxoPosition int
	PrivateKey   string
	Contract     string
	Client       string
}

type standardExitUTXOError struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Object      string `json:"object"`
		Code        string `json:"code"`
		Description string `json:"description"`
		Messages    struct {
			ErrorKey string `json:"error_key"`
		} `json:"messages"`
	} `json:"data"`
}

type StandardExitUTXOData struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		UtxoPos *big.Int `json:"utxo_pos"`
		Txbytes string   `json:"txbytes"`
		Proof   string   `json:"proof"`
	} `json:"data"`
}

type PlasmaTransaction struct {
	Blknum     uint
	Txindex    uint
	Oindex     uint
	Cur12      common.Address
	Toowner    common.Address
	Fromowner  common.Address
	Toamount   uint
	Fromamount uint
	Privatekey string
	Outputs    int
}

type MergeTransaction struct {
	Utxos      []SingleUTXO
	Fromowner  common.Address
	Privatekey string
}

type watcherStatus struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		LastValidatedChildBlockNumber int           `json:"last_validated_child_block_number"`
		LastMinedChildBlockTimestamp  int           `json:"last_mined_child_block_timestamp"`
		LastMinedChildBlockNumber     int           `json:"last_mined_child_block_number"`
		EthSyncing                    bool          `json:"eth_syncing"`
		ByzantineEvents               []interface{} `json:"byzantine_events"`
	} `json:"data"`
}

type watcherError struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Object      string `json:"object"`
		Description string `json:"description"`
		Code        string `json:"code"`
	} `json:"data"`
}

type inputUTXO struct {
	Txindex uint `json:"txindex"`
	Oindex  uint `json:"oindex"`
	Blknum  uint `json:"blknum"`
	Amount  uint `json:"amount"`
}

type outputUTXO struct {
	OwnerAddress common.Address `json:"owner"`
	Amount       uint           `json:"amount"`
	Currency     common.Address `json:"currency"`
}

type inputDeposit struct {
	Txindex uint `json:"txindex"`
	Oindex  uint `json:"oindex"`
	Blknum  uint `json:"blknum"`
}

type createdTx struct {
	Inputs  []inputUTXO  `json:"inputs"`
	Outputs []outputUTXO `json:"outputs"`
}

type input struct {
	Blknum  uint
	Txindex uint
	Oindex  uint
}

type output struct {
	OwnerAddress common.Address
	Currency     common.Address
	Amount       uint
}

type transactionToEncode struct {
	Inputs  []input
	Outputs []output
}

type transactionToBuild struct {
	Sig     [][]byte
	Inputs  []input
	Outputs []output
}

type SingleUTXO struct {
	UtxoPos  int    `json:"utxo_pos"`
	Txindex  int    `json:"txindex"`
	Owner    string `json:"owner"`
	Oindex   int    `json:"oindex"`
	Currency string `json:"currency"`
	Blknum   int    `json:"blknum"`
	Amount   int    `json:"amount"`
}

type exitDataUTXO struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		UtxoPos int64  `json:"utxo_pos"`
		Txbytes string `json:"txbytes"`
		Proof   string `json:"proof"`
	} `json:"data"`
}

type getUTXOError struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Object      string `json:"object"`
		Code        string `json:"code"`
		Description string `json:"description"`
		Messages    struct {
			ErrorKey string `json:"error_key"`
		} `json:"messages"`
	} `json:"data"`
}

// Create a basic transaction with 1 input and 1 output (2 if contains change)
func (p *PlasmaTransaction) createBasicTransaction() createdTx {
	NULL_ADDRESS := common.HexToAddress("0000000000000000000000000000000000000000")
	NULL_INPUT := inputUTXO{Blknum: 0, Txindex: 0, Oindex: 0}
	NULL_OUTPUT := outputUTXO{OwnerAddress: NULL_ADDRESS, Amount: 0, Currency: NULL_ADDRESS}

	singleInput := inputUTXO{Blknum: p.Blknum, Txindex: p.Txindex, Oindex: p.Oindex}

	var outputOne outputUTXO
	var outputTwo outputUTXO
	if p.Fromamount == p.Toamount {
		//send everything
		outputOne = outputUTXO{OwnerAddress: p.Toowner, Amount: p.Toamount, Currency: p.Cur12}
		outputTwo = outputUTXO{OwnerAddress: NULL_ADDRESS, Amount: 0, Currency: NULL_ADDRESS}
	} else {
		//send change to self
		outputOne = outputUTXO{OwnerAddress: p.Toowner, Amount: p.Toamount, Currency: p.Cur12}
		outputTwo = outputUTXO{OwnerAddress: p.Fromowner, Amount: p.Fromamount - p.Toamount, Currency: p.Cur12}
	}

	if p.Fromamount < p.Toamount {
		log.Fatal("UTXO not large enough to be sent")
	}

	var i []inputUTXO
	var o []outputUTXO
	i = append(i, singleInput, NULL_INPUT, NULL_INPUT, NULL_INPUT)
	o = append(o, outputOne, outputTwo, NULL_OUTPUT, NULL_OUTPUT)
	transaction := createdTx{Inputs: i, Outputs: o}

	return transaction
}

// Create a split transaction with 1 input and N output
func (p *PlasmaTransaction) createSplitTransaction() createdTx {
	totalAmount := p.Toamount * uint(p.Outputs)
	change := int(p.Fromamount) - int(totalAmount)
	if p.Outputs > 4 {
		log.Fatal("Exceeded maximum of 4 Outputs")
	}

	if change < 0 {
		log.Fatal("total output amount exceeded input amount")
	}

	if p.Outputs == 4 && change > 0 {
		log.Fatal("transaction has change, but outputs already reached maximum of 4")
	}

	NULL_ADDRESS := common.HexToAddress("0000000000000000000000000000000000000000")
	NULL_INPUT := inputUTXO{Blknum: 0, Txindex: 0, Oindex: 0}
	NULL_OUTPUT := outputUTXO{OwnerAddress: NULL_ADDRESS, Amount: 0, Currency: NULL_ADDRESS}
	//from a single input
	singleInput := inputUTXO{Blknum: p.Blknum, Txindex: p.Txindex, Oindex: p.Oindex}

	var i []inputUTXO
	var o []outputUTXO
	//no changes to be added
	if p.Fromamount == totalAmount {
		for no := 0; no < p.Outputs; no++ {
			o = append(o, outputUTXO{OwnerAddress: p.Toowner, Amount: p.Toamount, Currency: p.Cur12})
		}
	} else {
		for no := 0; no < p.Outputs; no++ {
			o = append(o, outputUTXO{OwnerAddress: p.Toowner, Amount: p.Toamount, Currency: p.Cur12})
		}
		//append change
		o = append(o, outputUTXO{OwnerAddress: p.Fromowner, Amount: uint(change), Currency: p.Cur12})
	}

	//fill the rest of outputs with null
	if len(o) < 4 {
		for ni := 0; ni < 4-len(o); ni++ {
			o = append(o, NULL_OUTPUT)
		}
	}

	i = append(i, singleInput, NULL_INPUT, NULL_INPUT, NULL_INPUT)
	transaction := createdTx{Inputs: i, Outputs: o}

	return transaction
}

// Create a merge transaction with 1 input and N output
func (m *MergeTransaction) createMergeTransaction() createdTx {
	NULL_ADDRESS := common.HexToAddress("0000000000000000000000000000000000000000")
	NULL_INPUT := inputUTXO{Blknum: 0, Txindex: 0, Oindex: 0}
	NULL_OUTPUT := outputUTXO{OwnerAddress: NULL_ADDRESS, Amount: 0, Currency: NULL_ADDRESS}
	var v uint
	for _, su := range m.Utxos {
		v = v + uint(su.Amount)
	}
	outputOne := outputUTXO{OwnerAddress: m.Fromowner, Amount: v, Currency: NULL_ADDRESS}

	var i []inputUTXO
	var o []outputUTXO

	for _, iu := range m.Utxos {
		in := inputUTXO{Blknum: uint(iu.Blknum), Txindex: uint(iu.Txindex), Oindex: uint(iu.Oindex)}
		i = append(i, in)
	}

	//fill the rest of inputs with null
	if len(i) < 4 {
		for ni := 0; ni <= 4-len(i); ni++ {
			i = append(i, NULL_INPUT)
		}
	}

	o = append(o, outputOne, NULL_OUTPUT, NULL_OUTPUT, NULL_OUTPUT)
	return createdTx{Inputs: i, Outputs: o}

}

// Encode transaction with RLP
func (c *createdTx) encodeTransaction() string {
	var t *transactionToEncode
	var i []input
	var o []output

	for _, val := range c.Inputs {
		t := input{Txindex: val.Txindex, Oindex: val.Oindex, Blknum: val.Blknum}
		i = append(i, t)
	}

	for _, val := range c.Outputs {
		t := output{OwnerAddress: val.OwnerAddress, Amount: val.Amount, Currency: val.Currency}
		o = append(o, t)
	}
	t = &transactionToEncode{Outputs: o, Inputs: i}

	encodedBytes, err := rlp.EncodeToBytes(t)
	if err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(encodedBytes)
}

// Submit transaction to endpoint, take tx byte and watcher URL
func submitTransaction(tx []byte, w string) transactionSuccessResponse {
	txstring := "0x" + hex.EncodeToString(tx)

	// Build request
	var url strings.Builder
	url.WriteString(w)
	url.WriteString("/transaction.submit")
	postData := map[string]interface{}{"transaction": txstring}
	js, _ := json.Marshal(postData)
	r, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(js))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Unmarshall the response
	response := transactionSuccessResponse{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		log.Warning("Could not unmarshal successful response from the Watcher")
		errorInfo := transactionFailureResponse{}
		processError := json.Unmarshal([]byte(rstring), &errorInfo)
		if processError != nil { // Response from the Watcher does not match a struct
			log.Fatal("Unknown response from Watcher API")
			panic("uh oh")
		}
		log.Warning("Unmarshalled JSON error response from the Watcher API")
		log.Error(errorInfo)
	} else {
		log.Info(resp.Status)
		log.Infof(
			"\n Response:\n Success: %v \n Blknum: %v \n txindex: %v\n Txhash: %v",
			response.Success,
			response.Data.Blknum,
			response.Data.Txindex,
			response.Data.Txhash,
		)
	}
	return response
}

// Wraps getUTXOsfromAddress, return a utxo from given position
func GetUTXO(address string, position uint, watcher string) SingleUTXO {
	tt := GetUTXOsFromAddress(address, watcher)
	var single SingleUTXO
	for _, t := range tt.Data {
		if uint(t.UtxoPos) == position {
			single = t
		}
	}
	return single
}

// Minimal send transaction function, take UTXO and send to an address
func (p *PlasmaTransaction) SendBasicTransaction(w string) transactionSuccessResponse {
	k := p.createBasicTransaction()
	encoded := k.encodeTransaction()
	var keys []string
	keys = append(keys, p.Privatekey)
	sig := util.SignTransaction(encoded, keys)
	//log.Info(hex.EncodeToString(buildSignedTransaction(sig, encoded)))
	transaction := buildSignedTransaction(sig, encoded)

	return submitTransaction(transaction, w)
}

// send split transaction to N outputs
func (p *PlasmaTransaction) SendSplitTransaction(w string) transactionSuccessResponse {
	k := p.createSplitTransaction()
	encoded := k.encodeTransaction()
	var keys []string
	keys = append(keys, p.Privatekey)
	sig := util.SignTransaction(encoded, keys)
	//log.Info(hex.EncodeToString(buildSignedTransaction(sig, encoded)))
	transaction := buildSignedTransaction(sig, encoded)

	return submitTransaction(transaction, w)
}

// form 1 input n output transaction
func (m *MergeTransaction) MergeBasicTransaction(w string) transactionSuccessResponse {
	if len(m.Utxos) > 4 {
		log.Fatal("Exceeded maximum of 4 UTXOs")
	}
	k := m.createMergeTransaction()
	encoded := k.encodeTransaction()
	var keys []string
	for range m.Utxos {
		keys = append(keys, m.Privatekey)
	}
	sig := util.SignTransaction(encoded, keys)
	transaction := buildSignedTransaction(sig, encoded)
	return submitTransaction(transaction, w)
}

// Deecode RLP, and rebuild transaction with signatures, finally encode the whole thing
func buildSignedTransaction(signatures [][]byte, unsignedTX string) []byte {
	var tx transactionToEncode
	//RLP decode unsignedTx
	decoded, err := hex.DecodeString(unsignedTX)
	if err != nil {
		log.Fatal(err)
	}
	rlp.DecodeBytes(decoded, &tx)
	//build Transaction
	var txsig [][]byte
	for _, s := range signatures {
		txsig = append(txsig, s)
	}
	builtTx := transactionToBuild{Inputs: tx.Inputs, Outputs: tx.Outputs, Sig: txsig}
	//RLP encode built transaction
	encoded, err := rlp.EncodeToBytes(builtTx)
	if err != nil {
		log.Fatal(err)
	}

	return encoded
}

// Start a standard exit from user provided UTXO & private key
func (s *StandardExit) StartStandardExit(watcher string) {
	log.Info("Getting data needed to exit the UTXO from the Watcher")
	exit, err := GetUTXOExitData(watcher, s.UtxoPosition)
	if err != nil{
		log.Fatal(err)
	}
	exit.StartStandardExit(s.Client, s.Contract, s.PrivateKey)
}

// Get the Watcher's status
func GetWatcherStatus(w string) {
	var url strings.Builder
	url.WriteString(w)
	url.WriteString("/status.get")
	r, err := http.NewRequest("POST", url.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Unmarshall the response
	response := watcherStatus{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		log.Error(jsonErr)
	}
	log.Info("Last validated Childchain block number: ", response.Data.LastValidatedChildBlockNumber)
	log.Info("Last mined Childchain block number: ", response.Data.LastMinedChildBlockNumber)
}

//Retrieve the UTXO exit data from the UTXO position
func GetUTXOExitData(watcher string, utxoPosition int) (StandardExitUTXOData, error) {
	// Build request
	var url strings.Builder
	url.WriteString(watcher)
	url.WriteString("/utxo.get_exit_data")
	postData := map[string]interface{}{"utxo_pos": utxoPosition}
	js, _ := json.Marshal(postData)
	r, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(js))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Unmarshall the response
	response := StandardExitUTXOData{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		return response, errors.New("No exit data found for UTXO provided")
	} else {
		log.Info(resp.Status)
	}

	return response, nil
}

// Retrieve the UTXOs associated with an address from the Watcher
func GetUTXOsFromAddress(address string, w string) util.WatcherUTXOsFromAddress {
	// Build request
	var url strings.Builder
	url.WriteString(w)
	url.WriteString("/account.get_utxos")
	postData := map[string]interface{}{"address": address, "limit": "10000"}
	js, _ := json.Marshal(postData)
	r, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(js))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Unmarshall the response
	response := util.WatcherUTXOsFromAddress{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		log.Warning("Could not unmarshal successful response from the Watcher")
		errorInfo := watcherError{}
		processError := json.Unmarshal([]byte(rstring), &errorInfo)
		if processError != nil { // Response from the Watcher does not match a struct
			log.Fatal("Unknown response from Watcher API")
			panic("uh oh")
		}
		log.Warning("Unmarshalled JSON error response from the Watcher API")
		log.Error(errorInfo)
	} else {
		log.Info(resp.Status)
	}

	return response
}

// Get balance for a certain address
func GetBalance(address string, watcher string) util.WatcherBalanceFromAddress {
	// Build request
	var url strings.Builder
	url.WriteString(watcher)
	url.WriteString("/account.get_balance")
	postData := map[string]interface{}{"address": address}
	js, _ := json.Marshal(postData)
	r, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(js))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Unmarshall the response
	response := util.WatcherBalanceFromAddress{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		log.Warning("Could not unmarshal successful response from the Watcher")
		errorInfo := watcherError{}
		processError := json.Unmarshal([]byte(rstring), &errorInfo)
		if processError != nil { // Response from the Watcher does not match a struct
			log.Fatal("Unknown response from Watcher API")
			panic("nani")
		}
		log.Warning("Unmarshalled JSON error response from the Watcher API")
		log.Error(errorInfo)
	} else {
		log.Info(resp.Status)
	}

	return response
}

// Start standard exit by calling the method in the smart contract
func (s *StandardExitUTXOData) StartStandardExit(ethereumClient string, contract string, private string) {
	client, err := ethclient.Dial(ethereumClient)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(private))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(31415926535) // in wei
	auth.GasLimit = uint64(210000)       // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(contract)
	instance, err := rootchain.NewRootchain(address, client)
	if err != nil {
		log.Fatal(err)
	}
	t := &bind.TransactOpts{}
	t.From = fromAddress
	t.Signer = auth.Signer
	t.Value = big.NewInt(31415926535) //STANDARD_EXIT_BOND in the smart contract
	t.GasLimit = 2000000

	txBytesHex, txErr := hex.DecodeString(util.RemoveLeadingZeroX(s.Data.Txbytes))
	if txErr != nil {
		log.Fatal(txErr)
	}

	proofBytesHex, proofErr := hex.DecodeString(util.RemoveLeadingZeroX(s.Data.Proof))
	if proofErr != nil {
		log.Fatal(proofErr)
	}
	tx, err := instance.StartStandardExit(t, s.Data.UtxoPos, []byte(txBytesHex), []byte(proofBytesHex))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Standard exit to Plasma MoreVP sent. Transaction: ", tx.Hash().Hex())
	}
}

// Deposit ETH into the already deployed Plasma MoreVP contract on Ethereum
func (d *PlasmaDeposit) DepositToPlasmaContract() {
	client, err := ethclient.Dial(d.Client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(d.PrivateKey))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))

	auth.Value = big.NewInt(int64(d.Amount)) // in wei
	auth.GasLimit = uint64(210000)           // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(d.Contract)

	rlpInputs := util.BuildRLPInput(util.RemoveLeadingZeroX(d.Owner), d.Currency, d.Amount)
	instance, err := rootchain.NewRootchain(address, client)
	if err != nil {
		log.Fatal(err)
	}
	t := &bind.TransactOpts{}
	t.From = fromAddress
	t.Signer = auth.Signer
	t.Value = big.NewInt(int64(d.Amount))
	tx, err := instance.Deposit(t, rlpInputs)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Deposit to Plasma MoreVP contract sent. Transaction: ", tx.Hash().Hex())
	}
}

// Calls the processExits in the Plasma smart contract to start processing exits that
// have completed the challenge period.
func ProcessExits(numberExitsToProcess int64, p ProcessExit) {
	client, err := ethclient.Dial(p.Client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(p.PrivateKey))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(210000) // in units
	auth.GasPrice = gasPrice

	contractAddress := p.Contract

	address := common.HexToAddress(contractAddress)
	instance, err := rootchain.NewRootchain(address, client)
	if err != nil {
		log.Fatal(err)
	}
	t := &bind.TransactOpts{}
	t.From = fromAddress
	t.Signer = auth.Signer
	t.Value = big.NewInt(0)
	t.GasLimit = 2000000

	token := common.HexToAddress(p.Token)

	tx, err := instance.ProcessExits(t, token, big.NewInt(0), big.NewInt(numberExitsToProcess))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Process exits request to Plasma MoreVP sent. Transaction: ", tx.Hash().Hex())
	}
}
