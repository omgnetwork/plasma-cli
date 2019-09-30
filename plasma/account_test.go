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
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"math/big"
)

// Mock test server
func createMockServer(t *testing.T, rString string, url string, rStruct interface{}) *httptest.Server {
	ms := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "POST" {
			t.Errorf("unexpected Method call %s, expecting POST request", r.Method)
		}
		if r.URL.EscapedPath() != url {
			t.Errorf("unexpected URL path %s, path expected to have %s", r.URL.EscapedPath(), url)
		}
		//expecting correct JSON post request
		rstring, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("unexpected error reading from req body: %v", err)
		}
		jsonErr := json.Unmarshal([]byte(rstring), &rStruct)

		if jsonErr != nil {
			t.Errorf("unexpected error unmarshalling JSON req %v", err)
		}

		// return response string
		fmt.Fprintln(w, rString)
	}))
	return ms
}

// Test getting balance for an account
func TestGetBalance(t *testing.T) {
	rs := `{ "data": [ { "amount": 1000, "currency": "0x0000000000000000000000000000000000000000" } ], "success": true, "version": "0.2" }`
	amountWanted := float64(1000)
	currencyWanted := EthCurrency
	ms := createMockServer(t, rs, "/account.get_balance", WatcherBalanceFromAddress{})
	defer ms.Close()
	balance, err := GetBalance("0xBddeAeE01f00e02c081D36c100D5DEe723cB9E17", ms.URL)
	if err != nil {
		t.Errorf("unexpected error from getting balance: %v", err)
	}
	if balance.Data[0].Amount != amountWanted {
		t.Errorf("unexpected balance amount returned, expecting %v, got %v", amountWanted, balance.Data[0].Amount)
	}
	if balance.Data[0].Currency != currencyWanted {
		t.Errorf("unexpected balance currency returned, expecting %s, got %s", currencyWanted, balance.Data[0].Currency)
	}
}

// Test getting balance for an account with Exponent number
func TestGetBalanceWithExponent(t *testing.T) {
	rs := `{ "data": [ { "amount": 3.335e+21, "currency": "0x0000000000000000000000000000000000000000" } ], "success": true, "version": "0.2" }`
	amountWanted := 3.335e+21
	currencyWanted := EthCurrency
	ms := createMockServer(t, rs, "/account.get_balance", WatcherBalanceFromAddress{})
	defer ms.Close()
	balance, err := GetBalance("0xBddeAeE01f00e02c081D36c100D5DEe723cB9E17", ms.URL)
	if err != nil {
		t.Errorf("unexpected error from getting balance: %v", err)
	}
	if balance.Data[0].Amount != amountWanted {
		t.Errorf("unexpected balance amount returned, expecting %v, got %v", amountWanted, balance.Data[0].Amount)
	}
	if balance.Data[0].Currency != currencyWanted {
		t.Errorf("unexpected balance currency returned, expecting %s, got %s", currencyWanted, balance.Data[0].Currency)
	}
}

// Test getting UTXOs for an account
func TestGetUTXOs(t *testing.T) {
	rs := `{ "data": [ { "amount": 123, "blknum": 1741002, "currency": "0x0000000000000000000000000000000000000000", "oindex": 0, "owner": "0x4522fb44c2ab359e76ecc75c22c9409690f12241", "txindex": 0, "utxo_pos": 1741002000000000 } ], "success": true, "version": "0.2" }`
	expectedAmount := float64(123)
	expectedUtxo := big.NewInt(1741002000000000)
	ms := createMockServer(t, rs, "/account.get_utxos", WatcherUTXOsFromAddress{})
	defer ms.Close()
	balance, err := GetUTXOsFromAddress("0x4522fb44c2ab359e76ecc75c22c9409690f12241", ms.URL)
	if err != nil {
		t.Errorf("unexpected error from getting balance: %v", err)
	}
	if balance.Data[0].Amount != expectedAmount {
		t.Errorf("unexpected utxo amount returned, expecting %v, got %v", expectedAmount, balance.Data[0].Amount)
	}
	if balance.Data[0].UtxoPos.Cmp(expectedUtxo) != 0 {
		t.Errorf("unexpected utxo position returned, expecting %v, got %v", expectedUtxo, balance.Data[0].UtxoPos)
	}
}

// Test getting UTXOs for an account with Exponent number
func TestGetUTXOsWithExponent(t *testing.T) {
	rs := `{ "data": [ { "amount": 1.949e+21, "blknum": 1741002, "currency": "0x0000000000000000000000000000000000000000", "oindex": 0, "owner": "0x4522fb44c2ab359e76ecc75c22c9409690f12241", "txindex": 0, "utxo_pos": 1741002000000000 } ], "success": true, "version": "0.2" }`
	amountWanted := 1.949e+21
	ms := createMockServer(t, rs, "/account.get_utxos", WatcherUTXOsFromAddress{})
	defer ms.Close()
	res, err := GetUTXOsFromAddress("0x4522fb44c2ab359e76ecc75c22c9409690f12241", ms.URL)
	if err != nil {
		t.Errorf("unexpected error from getting balance: %v", err)
	}
	if amountWanted != res.Data[0].Amount {
		t.Errorf("unexpected error, wanted amount %v, got %v", amountWanted, res.Data[0].Amount)
	}
}

func TestGetUTXOsWithLargeUtxoPos(t *testing.T) {
	rs := `{ "data": [ { "amount": 1.949e+21, "blknum": 1741002, "currency": "0x0000000000000000000000000000000000000000", "oindex": 0, "owner": "0x4522fb44c2ab359e76ecc75c22c9409690f12241", "txindex": 0, "utxo_pos": 1741002000000000100012344443 } ], "success": true, "version": "0.2" }`
	amountWanted := 1.949e+21
	utxoWanted, _ := new(big.Int).SetString("1741002000000000100012344443", 10)
	ms := createMockServer(t, rs, "/account.get_utxos", WatcherUTXOsFromAddress{})
	defer ms.Close()
	res, err := GetUTXOsFromAddress("0x4522fb44c2ab359e76ecc75c22c9409690f12241", ms.URL)
	if err != nil {
		t.Errorf("unexpected error from getting balance: %v", err)
	}
	if amountWanted != res.Data[0].Amount {
		t.Errorf("unexpected error, wanted amount %v, got %v", amountWanted, res.Data[0].Amount)
	}

	if utxoWanted.Cmp(res.Data[0].UtxoPos) != 0 {
		t.Errorf("unexpected error, wanted amount %v, got %v", utxoWanted, res.Data[0].UtxoPos)
	}
}

