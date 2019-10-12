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
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		LastValidatedChildBlockNumber int  `json:"last_validated_child_block_number"`
		LastMinedChildBlockTimestamp  int  `json:"last_mined_child_block_timestamp"`
		LastMinedChildBlockNumber     int  `json:"last_mined_child_block_number"`
		EthSyncing                    bool `json:"eth_syncing"`
		ByzantineEvents               []struct {
			Event   string                 `json:"event"`
			Details map[string]interface{} `json:"details"`
		} `json:"byzantine_events"`
	} `json:"data"`
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
