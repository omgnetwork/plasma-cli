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
	"net/http"

	"github.com/omisego/plasma-cli/childchain"
	"github.com/omisego/plasma-cli/plasma"
	"github.com/omisego/plasma-cli/util"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"testing"
)

func testDeposit(t *testing.T ) {
	d := plasma.PlasmaDeposit{PrivateKey: "0x598A2B06880B65CAA7C93A081D3A6235D7C667D18C81E8A14F918FC7D2C8270F", Client:"https://rinkeby.infura.io/v3/e2d2eee81e774cd3a4915d994dfec840" , Contract: "0x9631a230eaf33b51012fca494e4030d852bb9386", Amount: 1000, Owner: util.DeriveAddress("0x598A2B06880B65CAA7C93A081D3A6235D7C667D18C81E8A14F918FC7D2C8270F"), Currency: childchain.EthCurrency}
	d.DepositEthToPlasma()
	chch, err := childchain.NewClient(*watcherSubmitURL, &http.Client{})
	if err != nil {
		log.Errorf("unexpected error from creating new client: %v", err)
	}
	ptx := chch.NewPaymentTx()

	if err = ptx.AddOwner(util.DeriveAddress(privatekey)); err != nil {
		log.Errorf("unexpected error from Adding owner: %v", err)
	}

	if err = ptx.AddPayment(*amount, *to, *currency); err != nil {
		log.Errorf("unexpected error from Adding payment: %v", err)
	}
	if err = ptx.AddMetadata(*metadata); err != nil {
		log.Errorf("unexpected error from Adding metadata: %v", err)
	}

	if err = ptx.AddFee(*feeamount, *feetoken); err != nil {
		log.Errorf("unexpected error from Adding metadata: %v", err)
	}

	if err = childchain.BuildTransaction(ptx); err != nil {
		t.Errorf("unexpected error : %v", err)
	} 
	// _, err = childchain.SignTransaction(ptx, childchain.SignWithRawKeys(*privatekey))
	// if err != nil {
	// 	log.Errorf("unexpected error : %v", err)
	// }
	// res, err := childchain.SubmitTransaction(ptx)
	// if err != nil {
	// 	log.Errorf("unexpected error : %v", err)
	// }

}

