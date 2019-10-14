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
	"net/http"
	"testing"
)

func TestGetTransaction(t *testing.T) {
	rs := `{ "version": "1.0", "success": true, "data": { "txindex": 5113, "txhash": "0x5df13a6bf96dbcf6e66d8babd6b55bd40d64d4320c3b115364c6588fc18c2a21", "metadata": "0x00000000000000000000000000000000000000000000000000000048656c6c6f", "txbytes": "0x5df13a6bee20000...", "block": { "timestamp": 1540365586, "hash": "0x0017372421f9a92bedb7163310918e623557ab5310befc14e67212b660c33bec", "eth_height": 97424, "blknum": 68290000 }, "inputs": [ { "blknum": 1000, "txindex": 111, "oindex": 0, "utxo_pos": 1000001110000, "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea", "currency": "0x0000000000000000000000000000000000000000", "amount": 10 } ], "outputs": [ { "blknum": 68290000, "txindex": 5113, "oindex": 0, "utxo_pos": 68290000051130000, "owner": "0xae8ae48796090ba693af60b5ea6be3686206523b", "currency": "0x0000000000000000000000000000000000000000", "amount": 2 }, { "blknum": 68290000, "txindex": 5113, "oindex": 1, "utxo_pos": 68290000051130000, "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea", "currency": "0x0000000000000000000000000000000000000000", "amount": 7 } ] } }`

	ms := createMockServer(t, rs, "/transaction.get", TransactionGetResponse{})
	defer ms.Close()
	chch, err := NewClient(ms.URL, &http.Client{})
	if err != nil {
		t.Errorf("unexpected error from creating new client: %v", err)
	}
	balance, err := chch.GetTransaction("0x5df13a6bf96dbcf6e66d8babd6b55bd40d64d4320c3b115364c6588fc18c2a21")
	if err != nil {
		t.Errorf("unexpected error from transaction.get: %v", err)
	}
	if balance.Data.Txindex != 5113 {
		t.Errorf("unexpected txindex returned, expecting %v, got %v", 5113, balance.Data.Txindex)
	}

	if balance.Data.Block.Blknum != 68290000 {
		t.Errorf("unexpected blknum returned, expecting %v, got %v", 68290000, balance.Data.Block.Blknum)
	}
}
