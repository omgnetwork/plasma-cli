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

// plasma transaction that can be built
type Builder interface {
	BuildTransaction() (interface{}, error)
}

// plasma transaction that can be signed
type Signer interface {
	SignTransaction() ([]byte, error)
}

// plasma transaction that can be submitted
type Submitter interface {
	SubmitTransaction() (TransactionSubmitResponse, error)
}

// build arbitrary transaction to prepare for signing
func (c *Client) BuildTransaction(b Builder) (interface{}, error) {
	return b.BuildTransaction()
}

//sign transaction with signer interface
func SignTransaction(s Signer) ([]byte, error) {
	return s.SignTransaction()
}

//submit transaction with submitter interface
func (c *Client) SubmitTransaction(s Submitter) (TransactionSubmitResponse, error) {
	return s.SubmitTransaction()
}
