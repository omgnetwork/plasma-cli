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

package parser

import (
	"fmt"

	"github.com/omisego/plasma-cli/childchain"
	log "github.com/sirupsen/logrus"
)

type ByzantineEventCounts struct {
	PiggyBack        int
	NonCanonical     int
	UnchallengedExit int
	InvalidExit      int
}

// log the UTXO data in a human friendly format
func DisplayUTXOS(u *childchain.WatcherUTXOsFromAddress) {
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
func DisplayBalance(b *childchain.WatcherBalanceFromAddress) {
	for _, v := range b.Data {
		log.Info("Currency: ", v.Currency)
		log.Info("Amount: ", v.Amount)
	}
}

//TODO: Display all events in the specs
//parse byzantine events data based on event names
func DisplayByzantineEvents(b *childchain.WatcherStatus) ByzantineEventCounts {
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

	return ByzantineEventCounts{NonCanonical: n, PiggyBack: p, UnchallengedExit: ue, InvalidExit: ie}
}
