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
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

//payment tx struct
type PaymentTx struct {
	Client                    *Client
	CreateTransaction         CreateTransaction
	CreateTransactionResponse *CreateTransactionResponse
	Signatures                [][]byte
	TypedTransaction          TypedTransaction
}

type CreateTransaction struct {
	Owner    common.Address `json:"owner"`
	Payments []Payment      `json:"payments"`
	Fee      Fee            `json:"fee"`
	Metadata string         `json:"metadata"`
}

//payment struct
type Payment struct {
	Amount   uint64         `json:"amount"`
	Currency common.Address `json:"currency"`
	Owner    common.Address `json:"owner"`
}

// Typed transaction to be submitted
type TypedTransaction struct {
	Domain     Domain   `json:"domain"`
	Message    Message  `json:"message"`
	Signatures []string `json:"signatures"`
}

// Initialize a simple payment transaction with Fee and Metadata as a default value
func (c *Client) NewPaymentTx() (p *PaymentTx) {
	p = &PaymentTx{Client: c}
	p.AddFee(0, EthCurrency)
	p.AddMetadata(DefaultMetadata)
	return p
}

//add an owner to payment tx
func (p *PaymentTx) AddOwner(o string) error {
	if !common.IsHexAddress(o) {
		return fmt.Errorf("Owner is not a valid address")
	}
	p.CreateTransaction.Owner = common.HexToAddress(o)
	return nil
}

//add fee token and amount to payment tx
func (p *PaymentTx) AddFee(amount uint64, curr string) error {
	if !common.IsHexAddress(curr) {
		return fmt.Errorf("Fee currency is not a valid address")
	}
	p.CreateTransaction.Fee = Fee{Amount: amount, Currency: common.HexToAddress(curr)}
	return nil
}

//add metadata to payment tx
func (p *PaymentTx) AddMetadata(m string) error {
	p.CreateTransaction.Metadata = m
	return nil
}

//add Payment to payment slice in payment tx
func (p *PaymentTx) AddPayment(amount uint64, addr string, curr string) error {
	if !common.IsHexAddress(addr) {
		return fmt.Errorf("Recipient is not a valid address")
	}

	if !common.IsHexAddress(curr) {
		return fmt.Errorf("Payment currency is not a valid address")
	}
	payment := Payment{Amount: amount, Currency: common.HexToAddress(curr), Owner: common.HexToAddress(addr)}
	if len(p.CreateTransaction.Payments) == MaxOutputs {
		return fmt.Errorf("error too many payments, max outputs are %v", MaxOutputs)
	}
	p.CreateTransaction.Payments = append(p.CreateTransaction.Payments, payment)
	return nil
}

// Build transaction with simple payment transaction via transaction.create endpoint
func (p *PaymentTx) BuildTransaction() error {
	if len(p.CreateTransaction.Payments) == 0 {
		return errors.New("not enough arguement to build transaction")
	}
	rstring, err := p.Client.do(
		"/transaction.create",
		p.CreateTransaction,
	)
	response := CreateTransactionResponse{}
	if err != nil {
		return err
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		return jsonErr
	}

	// //payment cannot be completed in one transaction, this tx will perform a merge
	// if response.Data.Result == "intermediate" {
	// 	log.Info("enough balance to cover payment, but UTXOs are too small, this transaction will merge UTXOs..\n retry after confirmation")
	// }

	if response.Success == false {
		return fmt.Errorf(
			"Error creating transaction. \n Code: %v \n \n Description: %v",
			response.Data.Code,
			response.Data.Description,
		)
	}

	p.CreateTransactionResponse = &response
	return nil
}

// Sign transaction takes in any Signing function and runs in on the Signhash
func (p *PaymentTx) SignTransaction(signer SignerFunc) ([][]byte, error) {
	sigs, err := signer([]byte(p.CreateTransactionResponse.Data.Transactions[0].SignHash))
	if err != nil {
		return nil, fmt.Errorf("error signing transaction: %v", err)
	}
	p.Signatures = sigs
	return sigs, nil
}

// build typed transaction to be submitted on /transaction.submit_typed endpoint
func (p *PaymentTx) buildTypedTransaction() error {
	for _, s := range p.Signatures {
		p.TypedTransaction.Signatures = append(p.TypedTransaction.Signatures, fmt.Sprintf("0x%v", hex.EncodeToString(s)))
	}
	tx := p.CreateTransactionResponse.Data.Transactions[0].TypedData
	p.TypedTransaction.Domain = tx.Domain
	p.TypedTransaction.Message = tx.Message
	return nil
}

//submit payment transaction through /transaction.submit_typed endpoint
func (p *PaymentTx) SubmitTransaction() (*TransactionSubmitResponse, error) {
	if err := p.buildTypedTransaction(); err != nil {
		return nil, errors.New("error building typed transaction before submitting")
	}
	rstring, err := p.Client.do(
		"/transaction.submit_typed",
		p.TypedTransaction,
	)
	response := TransactionSubmitResponse{}
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if response.Success == false {
		return nil, fmt.Errorf(
			"Error submitting transaction. \n Code: %v \n \n Description: %v",
			response.Data.Code,
			response.Data.Description,
		)
	}

	return &response, nil
}
