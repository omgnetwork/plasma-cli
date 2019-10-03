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
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpySigner struct {
	Called        bool
	ErrorResponse error
	Result        [][]byte
}

type SpySubmiter struct {
	Called        bool
	ErrorResponse error
	Result        TransactionSubmitResponse
}

func (s *SpySigner) Sign() ([][]byte, error) {
	s.Called = true
	return s.Result, s.ErrorResponse
}

func (s *SpySubmiter) Submit() (*TransactionSubmitResponse, error) {
	s.Called = true
	return &s.Result, s.ErrorResponse
}

// Test the transaction.create/ calls via mocked server
func TestTransactionCreate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "POST" {
			t.Errorf("unexpected Method call %s, expecting POST request", r.Method)
		}
		if r.URL.EscapedPath() != "/transaction.create" {
			t.Errorf("unexpected URL path %s, path expected to have /transaction.create", r.URL.EscapedPath())
		}
		mr := `{ "version": "1.0", "success": true, "data": { "result": "complete", "transactions": [ { "inputs": [ { "blknum": 123000, "txindex": 111, "oindex": 0, "utxo_pos": 123000001110000, "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea", "currency": "0x0000000000000000000000000000000000000000", "amount": 50 }, { "blknum": 277000, "txindex": 2340, "oindex": 3, "utxo_pos": 277000023400003, "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea", "currency": "0x0000000000000000000000000000000000000000", "amount": 75 } ], "outputs": [ { "amount": 100, "currency": "0x0000000000000000000000000000000000000000", "owner": "0xae8ae48796090ba693af60b5ea6be3686206523b" }, { "amount": 20, "currency": "0x0000000000000000000000000000000000000000", "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea" } ], "fee": { "amount": 5, "currency": "0x0000000000000000000000000000000000000000" }, "metadata": "0x5df13a6bf96dbcf6e66d8babd6b55bd40d64d4320c3b115364c6588fc18c2a21", "txbytes": "0x5df13a6bee20000...", "sign_hash": "0x7851b951edb0b9e88f0fc80e83461f71d0f4b1d4e44fae7d25a5d4ab6adc5d3d", "typed_data": { "types": { "EIP712Domain": [ { "name": "name", "type": "string" }, { "name": "version", "type": "string" }, { "name": "verifyingContract", "type": "address" }, { "name": "salt", "type": "bytes32" } ], "Transaction": [ { "name": "input0", "type": "Input" }, { "name": "input1", "type": "Input" }, { "name": "input2", "type": "Input" }, { "name": "input3", "type": "Input" }, { "name": "output0", "type": "Output" }, { "name": "output1", "type": "Output" }, { "name": "output2", "type": "Output" }, { "name": "output3", "type": "Output" }, { "name": "metadata", "type": "bytes32" } ], "Input": [ { "name": "blknum", "type": "uint256" }, { "name": "txindex", "type": "uint256" }, { "name": "oindex", "type": "uint256" } ], "Output": [ { "name": "owner", "type": "address" }, { "name": "currency", "type": "address" }, { "name": "amount", "type": "uint256" } ] }, "primaryType": "Transaction", "domain": { "name": "OMG Network", "salt": "0xfad5c7f626d80f9256ef01929f3beb96e058b8b4b0e3fe52d84f054c0e2a7a83", "verifyingContract": "0x44de0ec539b8c4a4b530c78620fe8320167f2f74", "version": "1" }, "message": { "input0": { "blknum": 123000, "txindex": 111, "oindex": 0 }, "input1": { "blknum": 277000, "txindex": 2340, "oindex": 3 }, "input2": { "blknum": 0, "txindex": 0, "oindex": 0 }, "input3": { "blknum": 0, "txindex": 0, "oindex": 0 }, "output0": { "owner": "0xae8ae48796090ba693af60b5ea6be3686206523b", "currency": "0x0000000000000000000000000000000000000000", "amount": 100 }, "output1": { "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea", "currency": "0x0000000000000000000000000000000000000000", "amount": 20 }, "output2": { "owner": "0x0000000000000000000000000000000000000000", "currency": "0x0000000000000000000000000000000000000000", "amount": 0 }, "output3": { "owner": "0x0000000000000000000000000000000000000000", "currency": "0x0000000000000000000000000000000000000000", "amount": 0 }, "metadata": "0x5df13a6bf96dbcf6e66d8babd6b55bd40d64d4320c3b115364c6588fc18c2a21" } } } ] } }`
		//expecting correct JSON post request
		cr := CreateTransaction{}
		rstring, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("unexpected error reading from req body: %v", err)
		}
		jsonErr := json.Unmarshal([]byte(rstring), &cr)

		if jsonErr != nil {
			t.Errorf("unexpected error unmarshalling JSON req %v", err)
		}
		fmt.Fprintln(w, mr)

	}))
	defer ts.Close()
	mw := ts.URL
	p := NewCreateTransaction()
	p.Owner = "0xb3256026863eb6ae5b06fa396ab09069784ea8ea"
	p.Fee = Fee{Amount: 5, Currency: EthCurrency}
	p.Metadata = DefaultMetadata
	payment := Payments{
		Amount:   100,
		Owner:    "0xb3256026863eb6ae5b06fa396ab09069784ea8ea",
		Currency: EthCurrency,
	}
	p.Payments = []Payments{payment}
	p.WatcherEndpoint = mw
	_, err := p.CreateTransaction()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

// Test the transaction.submit_typed/ calls via mocked server
func TestTransactionSubmitTyped(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "POST" {
			t.Errorf("unexpected Method call %s, expecting POST request", r.Method)
		}
		if r.URL.EscapedPath() != "/transaction.submit_typed" {
			t.Errorf("unexpected URL path %s, path expected to have /transaction.submit_typed", r.URL.EscapedPath())
		}
		rr := `{ "version": "1.0", "success": true, "data": { "blknum": 123000, "txindex": 111, "txhash": "0xbdf562c24ace032176e27621073df58ce1c6f65de3b5932343b70ba03c72132d" } }`
		sr := TypedTransaction{}
		//expecting correct JSON post request
		rstring, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("unexpected error reading from req body: %v", err)
		}
		jsonErr := json.Unmarshal([]byte(rstring), &sr)

		if jsonErr != nil {
			t.Errorf("unexpected error unmarshalling JSON req %v", err)
		}

		// check that values in the request is correct
		fmt.Fprintln(w, rr)
	}))
	mw := ts.URL
	defer ts.Close()
	ctt, err := CreateTypedTransaction(Domain{}, Message{}, [][]byte{}, mw)
	if err != nil {
		t.Errorf("unexpected error creating typed transaction %v", err)
	}
	_, err = ctt.Submit()
	if err != nil {
		t.Errorf("unexpected error submitting typed transaction %v", err)
	}
}

func TestGetTransaction(t *testing.T) {
	rs := `{ "version": "1.0", "success": true, "data": { "txindex": 5113, "txhash": "0x5df13a6bf96dbcf6e66d8babd6b55bd40d64d4320c3b115364c6588fc18c2a21", "metadata": "0x00000000000000000000000000000000000000000000000000000048656c6c6f", "txbytes": "0x5df13a6bee20000...", "block": { "timestamp": 1540365586, "hash": "0x0017372421f9a92bedb7163310918e623557ab5310befc14e67212b660c33bec", "eth_height": 97424, "blknum": 68290000 }, "inputs": [ { "blknum": 1000, "txindex": 111, "oindex": 0, "utxo_pos": 1000001110000, "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea", "currency": "0x0000000000000000000000000000000000000000", "amount": 10 } ], "outputs": [ { "blknum": 68290000, "txindex": 5113, "oindex": 0, "utxo_pos": 68290000051130000, "owner": "0xae8ae48796090ba693af60b5ea6be3686206523b", "currency": "0x0000000000000000000000000000000000000000", "amount": 2 }, { "blknum": 68290000, "txindex": 5113, "oindex": 1, "utxo_pos": 68290000051130000, "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea", "currency": "0x0000000000000000000000000000000000000000", "amount": 7 } ] } }`

	ms := createMockServer(t, rs, "/transaction.get", TransactionGetResponse{})
	defer ms.Close()
	balance, err := GetTransaction("0x5df13a6bf96dbcf6e66d8babd6b55bd40d64d4320c3b115364c6588fc18c2a21", ms.URL)
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

// Test the transaction signing returns proper result and error
func TestSigning(t *testing.T) {
	ss := SpySigner{}
	_, err := Sign(&ss)
	if err != nil {
		t.Errorf("unexpected error from signing: %v", err)
	}
	if !ss.Called {
		t.Errorf("expected signer method to get called, got false")
	}
	bs := SpySigner{
		ErrorResponse: errors.New("mock error from signer"),
	}
	_, err = Sign(&bs)
	if err == nil {
		t.Errorf("expected error to return from signer, got nil")
	}
}

// Test the transaction submit returns proper result and error
func TestSubmit(t *testing.T) {
	ss := SpySubmiter{}
	_, err := Submit(&ss)
	if err != nil {
		t.Errorf("unexpected error from submitting: %v", err)
	}
	if !ss.Called {
		t.Errorf("expected submiting method to get called, got false")
	}
	bs := SpySubmiter{
		ErrorResponse: errors.New("mock error from submit method"),
	}
	_, err = Submit(&bs)
	if err == nil {
		t.Errorf("expected error to return from signer, got nil")
	}

}
