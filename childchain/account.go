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
	"fmt"
	"math/big"

	"github.com/omisego/plasma-cli/util"
	log "github.com/sirupsen/logrus"
)

// UTXOs of an address returned from watcher
type WatcherUTXOsFromAddress struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    []struct {
		UtxoPos  *big.Int `json:"utxo_pos"`
		Txindex  int      `json:"txindex"`
		Owner    string   `json:"owner"`
		Oindex   int      `json:"oindex"`
		Currency string   `json:"currency"`
		Blknum   int      `json:"blknum"`
		Amount   float64  `json:"amount"`
	} `json:"data"`
}

// balance of an address returned from watcher
type WatcherBalanceFromAddress struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    []struct {
		Currency string  `json:"currency"`
		Amount   float64 `json:"amount"`
	} `json:"data"`
}

// Retrieve the UTXOs associated with an address from the Watcher
func (c *Client) GetUTXOsFromAddress(address string) (*WatcherUTXOsFromAddress, error) {
	if util.ValidateHex(address) != nil {
		return nil, fmt.Errorf("error validating address in GetUTXOsFromAddress(): %v", util.ValidateHex(address))
	}
	postData := map[string]interface{}{"address": address, "limit": "10000"}
	rstring, err := c.do(
		"/account.get_utxos",
		postData,
	)
	response := WatcherUTXOsFromAddress{}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		return nil, err
	}

	return &response, nil
}

// Get balance for a certain address
func (c *Client) GetBalance(address string) (*WatcherBalanceFromAddress, error) {
	if util.ValidateHex(address) != nil {
		return nil, fmt.Errorf("error validating address in GetBalance(): %v", util.ValidateHex(address))
	}
	postData := map[string]interface{}{"address": address}
	rstring, err := c.do(
		"/account.get_balance",
		postData,
	)
	response := WatcherBalanceFromAddress{}
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
