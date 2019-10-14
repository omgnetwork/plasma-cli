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
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omisego/plasma-cli/util"
)

const (
	EthCurrency     = "0x0000000000000000000000000000000000000000"
	DefaultMetadata = "0x0000000000000000000000000000000000000000000000000000000000000000"
	MaxOutputs      = 4
)

//a type signer func can sign any transaction bytes
type SignerFunc func([]byte) ([][]byte, error)

type Fee struct {
	Amount   uint64         `json:"amount"`
	Currency common.Address `json:"currency"`
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
		Messages    struct {
			ErrorKey string `json:"error_key"`
		} `json:"messages"`
	} `json:"data"`
}

// plasma transaction interface, can build, sign, submit
type PlasmaTransaction interface {
	Builder
	Signer
	Submitter
}

// plasma transaction that can be built
type Builder interface {
	BuildTransaction() error
}

// plasma transaction that can be signed
type Signer interface {
	SignTransaction(SignerFunc) ([][]byte, error)
}

// plasma transaction that can be submitted
type Submitter interface {
	SubmitTransaction() (TransactionSubmitResponse, error)
}

// build arbitrary transaction to prepare for signing
func BuildTransaction(b Builder) error {
	return b.BuildTransaction()
}

//sign transaction with Signer interface and a signing function SignFunc
func SignTransaction(s Signer, sf SignerFunc) ([][]byte, error) {
	return s.SignTransaction(sf)
}

//submit transaction with submitter interface
func SubmitTransaction(s Submitter) (TransactionSubmitResponse, error) {
	return s.SubmitTransaction()
}

// Sign with raw keys takes private key as raw strings and return a function of type SignerFunc
func SignWithRawKeys(keys ...string) SignerFunc {
	return func(toSign []byte) ([][]byte, error) {
		if len(keys) > MaxOutputs {
			return nil, fmt.Errorf("error too many keys, maximum outputs are %v", MaxOutputs)
		}
		return util.SignHash(toSign, keys)
	}
}
