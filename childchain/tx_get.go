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

	log "github.com/sirupsen/logrus"
)

type TransactionGetResponse struct {
	Version string    `json:"version"`
	Success bool      `json:"success"`
	Data    GetTxData `json:"data"`
}

type GetTxData struct {
	Txindex  int       `json:"txindex"`
	Txhash   string    `json:"txhash"`
	Metadata string    `json:"metadata"`
	Txbytes  string    `json:"txbytes"`
	Block    Block     `json:"block"`
	Inputs   []Inputs  `json:"inputs"`
	Outputs  []Outputs `json:"outputs"`
}

type Block struct {
	Timestamp int    `json:"timestamp"`
	Hash      string `json:"hash"`
	EthHeight int    `json:"eth_height"`
	Blknum    int    `json:"blknum"`
}

// GetTransaction fetches data about
// a certain transaction hash via
// transaction.get endpoint
func (c *Client) GetTransaction(txHash string) (*TransactionGetResponse, error) {
	// client := &http.Client{}
	postData := map[string]interface{}{"id": txHash}
	rstring, err := c.do(
		"/transaction.get",
		postData,
	)
	response := TransactionGetResponse{}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		log.Warning("Could not unmarshal successful response from the Watcher")
		var errorInfo ClientError
		processError := json.Unmarshal([]byte(rstring), &errorInfo)
		if processError != nil { // Response from the Watcher does not match a struct
			return nil, err
		}
	}

	return &response, nil
}
