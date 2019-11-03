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
	"math/big"
	"net/http"
	"strconv"
	"testing"

	"github.com/omisego/plasma-cli/util"
)

//test build payment transaction via transaction.create works properly
func TestBuildPaymentTx(t *testing.T) {
	rs := `{ "version": "1.0", "success": true, "data": { "result": "complete", "transactions": [ { "inputs": [ { "blknum": 123000, "txindex": 1, "oindex": 0, "utxo_pos": 123000001110000, "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea", "currency": "0x0000000000000000000000000000000000000000", "amount": 50 }, { "blknum": 277000, "txindex": 2340, "oindex": 3, "utxo_pos": 277000023400003, "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea", "currency": "0x0000000000000000000000000000000000000000", "amount": 75 } ], "outputs": [ { "amount": 100, "currency": "0x0000000000000000000000000000000000000000", "owner": "0xae8ae48796090ba693af60b5ea6be3686206523b" }, { "amount": 20, "currency": "0x0000000000000000000000000000000000000000", "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea" } ], "fee": { "amount": 5, "currency": "0x0000000000000000000000000000000000000000" }, "metadata": "0x5df13a6bf96dbcf6e66d8babd6b55bd40d64d4320c3b115364c6588fc18c2a21", "txbytes": "0x5df13a6bee20000...", "sign_hash": "0x7851b951edb0b9e88f0fc80e83461f71d0f4b1d4e44fae7d25a5d4ab6adc5d3d", "typed_data": { "types": { "EIP712Domain": [ { "name": "name", "type": "string" }, { "name": "version", "type": "string" }, { "name": "verifyingContract", "type": "address" }, { "name": "salt", "type": "bytes32" } ], "Transaction": [ { "name": "input0", "type": "Input" }, { "name": "input1", "type": "Input" }, { "name": "input2", "type": "Input" }, { "name": "input3", "type": "Input" }, { "name": "output0", "type": "Output" }, { "name": "output1", "type": "Output" }, { "name": "output2", "type": "Output" }, { "name": "output3", "type": "Output" }, { "name": "metadata", "type": "bytes32" } ], "Input": [ { "name": "blknum", "type": "uint256" }, { "name": "txindex", "type": "uint256" }, { "name": "oindex", "type": "uint256" } ], "Output": [ { "name": "owner", "type": "address" }, { "name": "currency", "type": "address" }, { "name": "amount", "type": "uint256" } ] }, "primaryType": "Transaction", "domain": { "name": "OMG Network", "salt": "0xfad5c7f626d80f9256ef01929f3beb96e058b8b4b0e3fe52d84f054c0e2a7a83", "verifyingContract": "0x44de0ec539b8c4a4b530c78620fe8320167f2f74", "version": "1" }, "message": { "input0": { "blknum": 123000, "txindex": 111, "oindex": 0 }, "input1": { "blknum": 277000, "txindex": 2340, "oindex": 3 }, "input2": { "blknum": 0, "txindex": 0, "oindex": 0 }, "input3": { "blknum": 0, "txindex": 0, "oindex": 0 }, "output0": { "owner": "0xae8ae48796090ba693af60b5ea6be3686206523b", "currency": "0x0000000000000000000000000000000000000000", "amount": 100 }, "output1": { "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea", "currency": "0x0000000000000000000000000000000000000000", "amount": 20 }, "output2": { "owner": "0x0000000000000000000000000000000000000000", "currency": "0x0000000000000000000000000000000000000000", "amount": 0 }, "output3": { "owner": "0x0000000000000000000000000000000000000000", "currency": "0x0000000000000000000000000000000000000000", "amount": 0 }, "metadata": "0x5df13a6bf96dbcf6e66d8babd6b55bd40d64d4320c3b115364c6588fc18c2a21" } } } ] } }`
	amountWanted := json.Number(strconv.FormatInt(50, 10))
	blknumWanted := int(123000)
	utxoWanted := big.NewInt(123000001110000)
	ownerWanted := "0xb3256026863eb6ae5b06fa396ab09069784ea8ea"
	currencyWanted := util.EthCurrency
	ms := createMockServer(t, rs, "/transaction.create", WatcherBalanceFromAddress{})
	defer ms.Close()
	chch, err := NewClient(ms.URL, &http.Client{})
	if err != nil {
		t.Errorf("unexpected error from creating new client: %v", err)
	}
	paymenttx := chch.NewPaymentTx()
	paymenttx.AddOwner("0xb3256026863eb6ae5b06fa396ab09069784ea8ea")
	paymenttx.AddFee(5, EthCurrency)
	paymenttx.AddMetadata(DefaultMetadata)
	paymenttx.AddPayment(100, "0xb3256026863eb6ae5b06fa396ab09069784ea8ea", EthCurrency)
	err = paymenttx.BuildTransaction()
	response := paymenttx.CreateTransactionResponse.Data
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if response.Transactions[0].Inputs[0].Amount != amountWanted {
		t.Errorf("Expected %v got %v", amountWanted, response.Transactions[0].Inputs[0].Amount)
	}
	if response.Transactions[0].Inputs[0].Blknum != blknumWanted {
		t.Errorf("Expected %v got %v", blknumWanted, response.Transactions[0].Inputs[0].Blknum)
	}
	if response.Transactions[0].Inputs[0].UtxoPos.Cmp(utxoWanted) != 0 {
		t.Errorf("Expected %v got %v", utxoWanted, response.Transactions[0].Inputs[0].UtxoPos)
	}
	if response.Transactions[0].Inputs[0].Owner != ownerWanted {
		t.Errorf("Expected %v got %v", ownerWanted, response.Transactions[0].Inputs[0].Owner)
	}
	if response.Transactions[0].Inputs[0].Currency != currencyWanted {
		t.Errorf("Expected %v got %v", currencyWanted, response.Transactions[0].Inputs[0].Currency)
	}

}

//TODO test signing payment

//TODO test submit payment transaction
