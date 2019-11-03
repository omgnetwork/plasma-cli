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

package childchain

import (
	"encoding/json"
	"math/big"
)

type CreateTransactionResponse struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    Data   `json:"data"`
}

type Data struct {
	Result       string         `json:"result"`
	Transactions []Transactions `json:"transactions"`
	Object       string         `json:"object"`
	Code         string         `json:"code"`
	Description  string         `json:"description"`
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

type Inputs struct {
	Blknum   int         `json:"blknum"`
	Txindex  int         `json:"txindex"`
	Oindex   int         `json:"oindex"`
	UtxoPos  *big.Int    `json:"utxo_pos"`
	Owner    string      `json:"owner"`
	Currency string      `json:"currency"`
	Amount   json.Number `json:"amount"`
}

type Outputs struct {
	Amount   json.Number `json:"amount"`
	Currency string      `json:"currency"`
	Owner    string      `json:"owner"`
}

type TypedData struct {
	Types       Types   `json:"types"`
	PrimaryType string  `json:"primaryType"`
	Domain      Domain  `json:"domain"`
	Message     Message `json:"message"`
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
	Owner    string      `json:"owner"`
	Currency string      `json:"currency"`
	Amount   json.Number `json:"amount"`
}
type Output1 struct {
	Owner    string      `json:"owner"`
	Currency string      `json:"currency"`
	Amount   json.Number `json:"amount"`
}
type Output2 struct {
	Owner    string      `json:"owner"`
	Currency string      `json:"currency"`
	Amount   json.Number `json:"amount"`
}
type Output3 struct {
	Owner    string      `json:"owner"`
	Currency string      `json:"currency"`
	Amount   json.Number `json:"amount"`
}
