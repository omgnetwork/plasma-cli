// Copyright 2019 OmiseGO Pte Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may obtain a copy of the License at
// you may not use this file except in compliance with the License.
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
	"encoding/json"
	"net/http"

	"github.com/omisego/plasma-cli/util"

	log "github.com/sirupsen/logrus"
)

// UTXOs of an address returned from watcher
type WatcherUTXOsFromAddress struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    []struct {
		UtxoPos  int     `json:"utxo_pos"`
		Txindex  int     `json:"txindex"`
		Owner    string  `json:"owner"`
		Oindex   int     `json:"oindex"`
		Currency string  `json:"currency"`
		Blknum   int     `json:"blknum"`
		Amount   float64 `json:"amount"`
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

// log the UTXO data in a human friendly format
func DisplayUTXOS(u *WatcherUTXOsFromAddress) {
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

// log the balance data in a human friendly format
func DisplayBalance(b *WatcherBalanceFromAddress) {
	for _, v := range b.Data {
		log.Info("Currency: ", v.Currency)
		log.Info("Amount: ", v.Amount)
	}
}

// Retrieve the UTXOs associated with an address from the Watcher
func GetUTXOsFromAddress(address string, w string) (*WatcherUTXOsFromAddress, error) {
	client := &http.Client{}
	postData := map[string]interface{}{"address": address, "limit": "10000"}
	rstring, err := util.MakeChChReq(
		client,
		w,
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
func GetBalance(address string, watcher string) (*WatcherBalanceFromAddress, error) {
	client := &http.Client{}
	postData := map[string]interface{}{"address": address}
	rstring, err := util.MakeChChReq(
		client,
		watcher,
		"/account.get_balance",
		postData,
	)
	response := WatcherBalanceFromAddress{}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		log.Warning("Could not unmarshal successful response from the Watcher")
		errorInfo := watcherError{}
		processError := json.Unmarshal([]byte(rstring), &errorInfo)
		if processError != nil { // Response from the Watcher does not match a struct
			return nil, err
		}
	}

	return &response, nil
}
