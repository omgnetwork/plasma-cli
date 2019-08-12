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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/omisego/plasma-cli/util"
	log "github.com/sirupsen/logrus"
)

type PlasmaTransaction interface {
	Submit() (TransactionSubmitResponse, error)
}

type TransactionSigner interface {
	Sign() ([][]byte, error)
}

// plasma transaction to call via the convenient transaction.create endpoint
type CreateTransaction struct {
	Owner           string     `json:"owner"`
	Payments        []Payments `json:"payments"`
	Fee             Fee        `json:"fee"`
	Metadata        string     `json:"metadata"`
	WatcherEndpoint string
}
type Payments struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Owner    string `json:"owner"`
}

type CreateTransactionResponse struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    Data   `json:"data"`
}
type Inputs struct {
	Blknum   int    `json:"blknum"`
	Txindex  int    `json:"txindex"`
	Oindex   int    `json:"oindex"`
	UtxoPos  int64  `json:"utxo_pos"`
	Owner    string `json:"owner"`
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}
type Outputs struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Owner    string `json:"owner"`
}
type Fee struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}
type EIP712Domain struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type Transaction struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type Input struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type Output struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type Types struct {
	EIP712Domain []EIP712Domain `json:"EIP712Domain"`
	Transaction  []Transaction  `json:"Transaction"`
	Input        []Input        `json:"Input"`
	Output       []Output       `json:"Output"`
}
type Domain struct {
	Name              string `json:"name"`
	Salt              string `json:"salt"`
	VerifyingContract string `json:"verifyingContract"`
	Version           string `json:"version"`
}
type Input0 struct {
	Blknum  int `json:"blknum"`
	Txindex int `json:"txindex"`
	Oindex  int `json:"oindex"`
}
type Input1 struct {
	Blknum  int `json:"blknum"`
	Txindex int `json:"txindex"`
	Oindex  int `json:"oindex"`
}
type Input2 struct {
	Blknum  int `json:"blknum"`
	Txindex int `json:"txindex"`
	Oindex  int `json:"oindex"`
}
type Input3 struct {
	Blknum  int `json:"blknum"`
	Txindex int `json:"txindex"`
	Oindex  int `json:"oindex"`
}
type Output0 struct {
	Owner    string `json:"owner"`
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}
type Output1 struct {
	Owner    string `json:"owner"`
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}
type Output2 struct {
	Owner    string `json:"owner"`
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}
type Output3 struct {
	Owner    string `json:"owner"`
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}
type Message struct {
	Input0   Input0  `json:"input0"`
	Input1   Input1  `json:"input1"`
	Input2   Input2  `json:"input2"`
	Input3   Input3  `json:"input3"`
	Output0  Output0 `json:"output0"`
	Output1  Output1 `json:"output1"`
	Output2  Output2 `json:"output2"`
	Output3  Output3 `json:"output3"`
	Metadata string  `json:"metadata"`
}
type TypedData struct {
	Types       Types   `json:"types"`
	PrimaryType string  `json:"primaryType"`
	Domain      Domain  `json:"domain"`
	Message     Message `json:"message"`
}
type Transactions struct {
	Inputs    []Inputs  `json:"inputs"`
	Outputs   []Outputs `json:"outputs"`
	Fee       Fee       `json:"fee"`
	Metadata  string    `json:"metadata"`
	Txbytes   string    `json:"txbytes"`
	SignHash  string    `json:"sign_hash"`
	TypedData TypedData `json:"typed_data"`
}
type Data struct {
	Result       string         `json:"result"`
	Transactions []Transactions `json:"transactions"`
	Object       string         `json:"object"`
	Code         string         `json:"code"`
	Description  string         `json:"description"`
}

type Messages struct {
	ErrorKey string `json:"error_key"`
}
type ErrorData struct {
	Object      string   `json:"object"`
	Code        string   `json:"code"`
	Description string   `json:"description"`
	Messages    Messages `json:"messages"`
}
type TypedTransaction struct {
	Domain          Domain   `json:"domain"`
	Message         Message  `json:"message"`
	Signatures      []string `json:"signatures"`
	WatcherEndpoint string
}

type TransactionSubmitResponse struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Blknum      int    `json:"blknum"`
		Txindex     int    `json:"txindex"`
		Txhash      string `json:"txhash"`
		Object      string `json:"object"`
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"data"`
}

type TransactionSubmitFailureResponse struct {
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

type SingleSigner struct {
	ToSign     []byte
	PrivateKey string
}

func NewCreateTransaction() CreateTransaction {
	return CreateTransaction{}
}

//call the transaction.create HTTP endpoint
func (c *CreateTransaction) CreateTransaction() (*CreateTransactionResponse, error) {
	// Build request
	var url strings.Builder
	url.WriteString(c.WatcherEndpoint)
	url.WriteString("/transaction.create")
	js, _ := json.Marshal(c)
	r, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(js))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshall the response
	response := CreateTransactionResponse{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)

	if jsonErr != nil {
		return nil, jsonErr
	}

	//payment cannot be completed in one atomic transaction, this tx will perform a merge
	if response.Data.Result == "intermediate" {
		log.Info("enough balance to cover payment, but UTXOs are too small, this transaction will merge UTXOs..\n retry after confirmation")
	}

	if response.Success == false {
		log.Fatalf(
			"Error creating transaction. \n Code: %v \n \n Description: %v",
			response.Data.Code,
			response.Data.Description,
		)
	}

	return &response, nil
}

//getting a typed data hash to be signed
func (t *Transactions) GetToSignHash() string {
	return t.SignHash
}

//getting typed data
func (t *Transactions) GetTypedData() TypedData {
	return t.TypedData
}

//create a typed transaction to submit
func CreateTypedTransaction(d Domain, m Message, sigs [][]byte, w string) (TypedTransaction, error) {
	var hexsigs []string
	var ttx TypedTransaction
	for _, s := range sigs {
		hexsigs = append(hexsigs, fmt.Sprintf("0x%v", hex.EncodeToString(s)))
	}
	ttx.Domain = d
	ttx.Message = m
	ttx.Signatures = hexsigs
	ttx.WatcherEndpoint = w
	return ttx, nil
}

//sign with a single private key
func (s SingleSigner) Sign() ([][]byte, error) {
	return util.SignHash(s.ToSign, []string{s.PrivateKey})
}

//submit transaction to submit_typed endpoint
func (t TypedTransaction) Submit() (TransactionSubmitResponse, error) {
	// Build request
	var url strings.Builder
	url.WriteString(t.WatcherEndpoint)
	url.WriteString("/transaction.submit_typed")
	js, _ := json.Marshal(t)
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
	response := TransactionSubmitResponse{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		log.Warning("Could not unmarshal successful response from the Watcher")
		errorInfo := TransactionSubmitFailureResponse{}
		processError := json.Unmarshal([]byte(rstring), &errorInfo)
		if processError != nil { // Response from the Watcher does not match a struct
			log.Fatal("Unknown response from Watcher API")
			panic("uh oh")
		}
		log.Warning("Unmarshalled JSON error response from the Watcher API")
		log.Error(errorInfo)
	}
	if response.Success == true {
		log.Info(resp.Status)
		log.Infof(
			"\n Response:\n Success: %v \n Blknum: %v \n txindex: %v\n Txhash: %v",
			response.Success,
			response.Data.Blknum,
			response.Data.Txindex,
			response.Data.Txhash,
		)
	} else {
		log.Info(resp.Status)
		log.Fatalf(
			"\n Error submitting transaction:\n Object: %v \n Code: %v \n Description: %v\n Error key: %v",
			response.Data.Object,
			response.Data.Code,
			response.Data.Description,
		)
	}

	return response, nil
}

func Submit(p PlasmaTransaction) (ts TransactionSubmitResponse, err error) {
	ts, err = p.Submit()
	if err != nil {
		return
	}
	return ts, nil
}

func Sign(t TransactionSigner) (sigs [][]byte, err error) {
	sigs, err = t.Sign()
	if err != nil {
		return
	}
	return sigs, nil
}
