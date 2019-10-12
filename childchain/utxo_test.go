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
	"math/big"
	"net/http"
	"testing"
)

func TestGetExitData(t *testing.T) {
	rs := `{ "version": "1.0", "success": true, "data": { "proof": "0xcedb8b31d1e4...", "sigs": "0xcc52d1d9ffe7...", "txbytes": "0x3eb6ae5b06f3...", "utxo_pos": 10000000010000000 } }`
	proofWanted := "0xcedb8b31d1e4..."
	sigsWanted := "0xcc52d1d9ffe7..."
	txbytesWanted := "0x3eb6ae5b06f3..."
	ms := createMockServer(t, rs, "/utxo.get_exit_data", StandardExitUTXOData{})
	defer ms.Close()
	chch, err := NewClient(ms.URL, &http.Client{})
	if err != nil {
		t.Errorf("unexpected error from creating new client: %v", err)
	}
	res, err := chch.GetUTXOExitData(10000000010000000)
	if err != nil {
		t.Errorf("unexpected error from getting exit data: %v", err)
	}
	if proofWanted != res.Data.Proof {
		t.Errorf("unexpected error, wanted %v, got %v", proofWanted, res.Data.Proof)
	}
	if sigsWanted != res.Data.Sigs {
		t.Errorf("unexpected error, wanted %v, got %v", sigsWanted, res.Data.Sigs)
	}
	if txbytesWanted != res.Data.Txbytes {
		t.Errorf("unexpected error, wanted %v, got %v", txbytesWanted, res.Data.Txbytes)
	}
}

func TestGetChallengeData(t *testing.T) {
	rs := `{ "version": "1.0", "success": true, "data": { "exit_id": 1718381723888, "input_index": 0, "sig": "0x6bfb9b2dbe32...", "txbytes": "0x3eb6ae5b06f3..." } }`
	exitWanted := big.NewInt(1718381723888)
	indexWanted := uint8(0)
	sigWanted := "0x6bfb9b2dbe32..."
	txbytesWanted := "0x3eb6ae5b06f3..."
	ms := createMockServer(t, rs, "/utxo.get_challenge_data", ChallengeUTXOData{})
	defer ms.Close()
	chch, err := NewClient(ms.URL, &http.Client{})
	if err != nil {
		t.Errorf("unexpected error from creating new client: %v", err)
	}
	res, err := chch.GetChallengeData(10000000010000000)
	if err != nil {
		t.Errorf("unexpected error from getting challenge data: %v", err)
	}
	if exitWanted.Cmp(res.Data.ExitId) != 0 {
		t.Errorf("unexpected error, wanted %v, got %v", exitWanted, res.Data.ExitId)
	}
	if indexWanted != res.Data.InputIndex {
		t.Errorf("unexpected error, wanted %v, got %v", indexWanted, res.Data.InputIndex)
	}
	if sigWanted != res.Data.Sig {
		t.Errorf("unexpected error, wanted %v, got %v", sigWanted, res.Data.Sig)
	}
	if txbytesWanted != res.Data.Txbytes {
		t.Errorf("unexpected error, wanted %v, got %v", txbytesWanted, res.Data.Txbytes)
	}
}
