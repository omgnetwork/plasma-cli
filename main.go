package main

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

import (
	// "github.com/omisego/plasma-cli/parser"
	"encoding/hex"

	"github.com/omisego/plasma-cli/plasma"
	"github.com/omisego/plasma-cli/util"
)

func main() {
	// util.LogFormatter()
	// log.Info("Starting OmiseGO Plasma MoreVP CLI")
	p := plasma.NewCreateTransaction()
	p.Owner = "0xBddeAeE01f00e02c081D36c100D5DEe723cB9E17"
	p.Fee = plasma.Fee{
		Amount:   0,
		Currency: "0x0000000000000000000000000000000000000000",
	}
	p.Metadata = "0x0000000000000000000000000000000000000000000000000000000000000000"
	payment := plasma.Payments{
		Amount:   50,
		Owner:    "0x0527a37aa7081efcf405bd7c8fe36b01e91df27d",
		Currency: "0x0000000000000000000000000000000000000000",
	}
	p.Payments = []plasma.Payments{payment}
	p.WatcherEndpoint = "https://watcher-development-parity.omg.network/"
	resp, _ := p.CreateTransaction()

	pk := ""
	hash, _ := hex.DecodeString(util.FilterZeroX(resp.Data.Transactions[0].GetToSignHash()))
	//signatures := util.SignHash(hash, []string{pk})
	signatures, _ := plasma.Sign(plasma.SingleSigner{ToSign: hash, PrivateKey: pk})
	tx, _ := plasma.CreateTypedTransaction(
		resp.Data.Transactions[0].GetTypedData().Domain,
		resp.Data.Transactions[0].GetTypedData().Message,
		signatures,
		"https://watcher-development-parity.omg.network/",
	)
	plasma.Submit(tx)

	// parser.ParseArgs()
}
