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

package plasma

import (
	"encoding/json"
	"fmt"
	"net/http"
	"math/big"

	"github.com/omisego/plasma-cli/util"	
	log "github.com/sirupsen/logrus"
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

type InvalidExitDetails struct {
	EthHeight int     `json:"eth_height"`
	UtxoPos   *big.Int `json:"utxo_pos"`
	Owner     string  `json:"owner"`
	Currency  string  `json:"currency"`
	Amount    float64 `json:"amount"`
}

type ByzantineEvents struct {
	PiggyBack        int
	NonCanonical     int
	UnchallengedExit int
	InvalidExit      int
}

// Get the Watcher's status
func GetWatcherStatus(w string) (*WatcherStatus, error) {
	client := &http.Client{}
	rstring, err := util.SendChChReq(
		client,
		w,
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

//TODO: Display all events in the specs
//parse byzantine events data based on event names
func DisplayByzantineEvents(b *WatcherStatus) ByzantineEvents {
	var n int
	var p int
	var ue int
	var ie int

	for _, v := range b.Data.ByzantineEvents {
		switch v.Event {
		case "noncanonical_ife": //TODO: parse noncanonical_ife events
			n++
		case "piggyback_available": //TODO: parse piggyback_available events
			p++
		case "unchallenged_exit":
			ue++
			log.Infof("unchallenged exit: \n eth_height: %v, \n utxo_pos: %v, \n owner: %v, \n currency: %v, \n amount: %v",
				v.Details["eth_height"],
				fmt.Sprintf("%.0f", v.Details["utxo_pos"]),
				v.Details["owner"],
				v.Details["currency"],
				v.Details["amount"],
			)
		case "invalid_exit":
			ie++
			log.Infof("invalid exit: \n eth_height: %v, \n utxo_pos: %v, \n owner: %v, \n currency: %v, \n amount: %v",
				v.Details["eth_height"],
				fmt.Sprintf("%.0f", v.Details["utxo_pos"]),
				v.Details["owner"],
				v.Details["currency"],
				v.Details["amount"],
			)
		}
	}
	log.Infof(
		"Total Byzantine events:\n Invalid exits: %v, \n unchallenged exits: %v, \n Piggyback available: %v \n non-canonical ife: %v",
		ie,
		ue,
		p,
		n,
	)
	log.Info("Last validated Childchain block number: ", b.Data.LastValidatedChildBlockNumber)
	log.Info("Last mined Childchain block number: ", b.Data.LastMinedChildBlockNumber)

	return ByzantineEvents{NonCanonical: n, PiggyBack: p, UnchallengedExit: ue, InvalidExit: ie}
}
