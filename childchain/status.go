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
)

//only displaying invalid exit details for now
type WatcherStatus struct {
	Data struct {
		ByzantineEvents []struct {
			Event   string                 `json:"event"`
			Details map[string]interface{} `json:"details"`
		} `json:"byzantine_events"`
		ContractAddr  string `json:"contract_addr"`
		EthSyncing    bool   `json:"eth_syncing"`
		InFlightExits []struct {
			EthHeight          int           `json:"eth_height"`
			PiggybackedInputs  []interface{} `json:"piggybacked_inputs"`
			PiggybackedOutputs []interface{} `json:"piggybacked_outputs"`
			Txbytes            string        `json:"txbytes"`
			Txhash             string        `json:"txhash"`
		} `json:"in_flight_exits"`
		LastMinedChildBlockNumber        int `json:"last_mined_child_block_number"`
		LastMinedChildBlockTimestamp     int `json:"last_mined_child_block_timestamp"`
		LastSeenEthBlockNumber           int `json:"last_seen_eth_block_number"`
		LastSeenEthBlockTimestamp        int `json:"last_seen_eth_block_timestamp"`
		LastValidatedChildBlockNumber    int `json:"last_validated_child_block_number"`
		LastValidatedChildBlockTimestamp int `json:"last_validated_child_block_timestamp"`
		ServicesSyncedHeights            []struct {
			Height  int    `json:"height"`
			Service string `json:"service"`
		} `json:"services_synced_heights"`
	} `json:"data"`
	Success bool   `json:"success"`
	Version string `json:"version"`
}

// Get the Watcher's status
func (c *Client) GetWatcherStatus() (*WatcherStatus, error) {
	rstring, err := c.do(
		"/status.get",
		nil,
	)
	response := WatcherStatus{}
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &response, nil
}
