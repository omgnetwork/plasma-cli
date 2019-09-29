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
	"testing"
)

// Test getting status from watcher
func TestGetStatus(t *testing.T) {
	rs := `{ "version": "1.0", "success": true, "data": { "last_validated_child_block_timestamp": 1558535130, "last_validated_child_block_number": 10000, "last_mined_child_block_timestamp": 1558535190, "last_mined_child_block_number": 11000, "last_seen_eth_block_timestamp": 1558535190, "last_seen_eth_block_number": 4427041, "contract_addr": "0x44de0ec539b8c4a4b530c78620fe8320167f2f74", "eth_syncing": true, "byzantine_events": [ { "event": "invalid_exit", "details": { "eth_height": 615440, "utxo_pos": 10000000010000000, "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea", "currency": "0x0000000000000000000000000000000000000000", "amount": 100 } }, { "event": "unchallenged_exit", "details": { "eth_height": 615440, "utxo_pos": 10000000010000000, "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea", "currency": "0x0000000000000000000000000000000000000000", "amount": 100 } } ], "services_synced_heights": [ { "service": "block_getter", "height": 4427041 }, { "service": "challenges_responds_processor", "height": 4427029 }, { "service": "competitor_processor", "height": 4427029 }, { "service": "convenience_deposit_processor", "height": 4427031 }, { "service": "convenience_exit_processor", "height": 4427029 }, { "service": "depositor", "height": 4427031 }, { "service": "exit_challenger", "height": 4427029 }, { "service": "exit_finalizer", "height": 4427029 }, { "service": "exit_processor", "height": 4427029 }, { "service": "ife_exit_finalizer", "height": 4427029 }, { "service": "in_flight_exit_processor", "height": 4427029 }, { "service": "piggyback_challenges_processor", "height": 4427029 }, { "service": "piggyback_processor", "height": 4427029 }, { "service": "root_chain_height", "height": 4427041 } ] } }`
	ms := createMockServer(t, rs, "/status.get", WatcherStatus{})
	defer ms.Close()
	status, err := GetWatcherStatus(ms.URL)
	if err != nil {
		t.Errorf("unexpected error from getting status: %v", err)
	}
	dbe := DisplayByzantineEvents(status)
	if dbe.InvalidExit != 1 {
		t.Errorf("unexpected number of invalid exits, expecting %v, got %v", 1, dbe.InvalidExit)
	}
	if dbe.UnchallengedExit != 1 {
		t.Errorf("unexpected number of unchallenged exits, expecting %v, got %v", 1, dbe.NonCanonical)
	}
}
