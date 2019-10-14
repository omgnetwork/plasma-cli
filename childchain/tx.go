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

// SignerFunc is any function that takes in raw byte
// and return signature of type [][]byte
type SignerFunc func([]byte) ([][]byte, error)

// Fee is a amount and currency of transaction fee
type Fee struct {
	Amount   uint64         `json:"amount"`
	Currency common.Address `json:"currency"`
}

// TransactionSubmitResponse is an expected
// response from making transaction to watcher
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

// PlasmaTransaction interface, can build, sign, submit
type PlasmaTransaction interface {
	Builder
	Signer
	Submitter
}

// Builder is a struct that can build
// plasma tx
type Builder interface {
	BuildTransaction() error
}

// Signer is a struct that can sign
// plasma tx
type Signer interface {
	SignTransaction(SignerFunc) ([][]byte, error)
}

// Submitter is a struct that can
// submit plasma tx
type Submitter interface {
	SubmitTransaction() (*TransactionSubmitResponse, error)
}

// BuildTransaction to prepare for signing
func BuildTransaction(b Builder) error {
	return b.BuildTransaction()
}

// SignTransaction with Signer interface and a signing function SignFunc
// agnostic to the actual implementation of signer function
func SignTransaction(s Signer, sf SignerFunc) ([][]byte, error) {
	return s.SignTransaction(sf)
}

// SubmitTransaction with submitter interface
func SubmitTransaction(s Submitter) (*TransactionSubmitResponse, error) {
	return s.SubmitTransaction()
}

// SignWithRawKeys takes private key as raw strings and return a function of type SignerFunc
// NOTE: this is a default convenience function for signing, user should implement different
// Signing function
func SignWithRawKeys(keys ...string) SignerFunc {
	return func(toSign []byte) ([][]byte, error) {
		if len(keys) > MaxOutputs {
			return nil, fmt.Errorf("error too many keys, maximum outputs are %v", MaxOutputs)
		}
		return util.SignHash(toSign, keys)
	}
}
