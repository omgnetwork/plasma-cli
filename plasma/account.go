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
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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

// log the balance data in a human friendly format
func DisplayBalance(b WatcherBalanceFromAddress) {
	for _, v := range b.Data {
		log.Info("Currency: ", v.Currency)
		log.Info("Amount: ", v.Amount)
	}
}

// Retrieve the UTXOs associated with an address from the Watcher
func GetUTXOsFromAddress(address string, w string) (*WatcherUTXOsFromAddress, error) {
	// Build request
	var url strings.Builder
	url.WriteString(w)
	url.WriteString("/account.get_utxos")
	postData := map[string]interface{}{"address": address, "limit": "10000"}
	js, _ := json.Marshal(postData)
	r, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(js))
	if err != nil {
		return nil, err
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
	response := WatcherUTXOsFromAddress{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		return nil, err
	} 

	return &response, nil
}

// Get balance for a certain address
func GetBalance(address string, watcher string) (*WatcherBalanceFromAddress, error) {
	// Build request
	var url strings.Builder
	url.WriteString(watcher)
	url.WriteString("/account.get_balance")
	postData := map[string]interface{}{"address": address}
	js, _ := json.Marshal(postData)
	r, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(js))
	if err != nil {
		return nil, err
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
	response := WatcherBalanceFromAddress{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
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
