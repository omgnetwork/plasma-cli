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
	"encoding/hex"
	"os"

	"github.com/omisego/plasma-cli/plasma"
	"github.com/omisego/plasma-cli/util"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	get              = kingpin.Command("get", "Get a resource.").Default()
	getUTXO          = get.Command("utxos", "Retrieve UTXO data from the Watcher service")
	watcherURL       = get.Flag("watcher", "FQDN of the Watcher in the format https://watcher.path.net").Required().String()
	ownerUTXOAddress = get.Flag("address", "Owner address to search UTXOs").String()
	getBalance       = get.Command("balance", "Retrieve balance of an address from the Watcher service")
	status           = get.Command("status", "Get status from the Watcher")
	transaction      = get.Command("transaction", "get transaction details")
	txHash           = get.Flag("txhash", "transaction hash of the transaction you would like to get the information for").String()

	getExit             = get.Command("exit", "Get UTXO exit information")
	getExitUTXOPosition = getExit.Flag("utxo", "Get UTXO exit information").Required().Int()

	deposit         = kingpin.Command("deposit", "Deposit ETH or ERC20 into the Plamsa MoreVP smart contract.")
	privateKey      = deposit.Flag("privatekey", "Private key of the account used to send funds into Plasma MoreVP").Required().String()
	client          = deposit.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()
	contract        = deposit.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	depositOwner    = deposit.Flag("owner", "Owner of the UTXOs").Required().String()
	depositAmount   = deposit.Flag("amount", "Amount to deposit in wei").Required().Uint64()
	depositCurrency = deposit.Flag("currency", "Currency of the deposit. Example: ETH").Required().String()

	send             = kingpin.Command("send", "Create a transaction on the OmiseGO Plasma MoreVP network")
	to               = send.Flag("to", "Wallet address of the recipient").Required().String()
	privatekey       = send.Flag("privatekey", "Privatekey from a wallet to send from").Required().String()
	amount           = send.Flag("amount", "Amount to transact").Required().Uint64()
	currency         = send.Flag("currency", "currency of the amount to send, default to ETH").Default(plasma.EthCurrency).String()
	watcherSubmitURL = send.Flag("watcher", "FQDN of the Watcher in the format http://watcher.path.net").Required().String()
	feetoken         = send.Flag("feetoken", "set the token to be used as transaction fee, default to ETH").Default(plasma.EthCurrency).String()
	feeamount        = send.Flag("feeamount", "set the amount to be used as transaction fee, default to 0").Default("0").Uint64()
	metadata         = send.Flag("metadata", "additional metadata to send with the transaction, up to 32 bytes").Default(plasma.DefaultMetadata).String()

	exit           = kingpin.Command("exit", "Standard exit a UTXO back to the root chain.")
	watcherExitURL = exit.Flag("watcher", "FQDN of the Watcher in the format http://watcher.path.net").Required().String()
	contractExit   = exit.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	utxoPosition   = exit.Flag("utxo", "UTXO Position that will be exited").Required().String()
	exitPrivateKey = exit.Flag("privatekey", "Private key of the UTXO owner").Required().String()
	clientExit     = exit.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()

	process           = kingpin.Command("process", "Process all exits that have completed the challenge period")
	processContract   = process.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	processToken      = process.Flag("token", "Token address to process for standard exits").Required().String()
	processPrivateKey = process.Flag("privatekey", "Private key used to fund the gas for the smart contract call").Required().String()
	processExitClient = process.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()

	create        = kingpin.Command("create", "Create a resource.")
	createAccount = create.Command("account", "Create an account consisting of Public and Private key")

	challengeexit  = kingpin.Command("challenge exit", "Challenge an invalid exit on the root chain")
	ceContract     = challengeexit.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	ceClient       = challengeexit.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()
	cePrivateKey   = challengeexit.Flag("privatekey", "Private key used to fund the gas for the smart contract call").Required().String()
	ceUTXOPosition = challengeexit.Flag("utxo", "the UTXO to call challenge exit").Required().String()
	ceWatcherURL   = challengeexit.Flag("watcher", "FQDN of the Watcher in the format http://watcher.path.net").Required().String()
)

func ParseArgs() {
	switch kingpin.Parse() {
	case getUTXO.FullCommand():
		//plasma_cli get utxos --address=0x944A81BeECac91802787fBCFB9767FCBf81db1f5 --watcher=http://watcher.path.net
		ownerUTXOAddress := *ownerUTXOAddress
		if len(ownerUTXOAddress) == 0 {
			log.Error("Address is required to get UTXO data")
			os.Exit(1)
		}
		utxos, err := plasma.GetUTXOsFromAddress(ownerUTXOAddress, *watcherURL)
		if err != nil {
			log.Error(err)
		}
		plasma.DisplayUTXOS(utxos)
	case getBalance.FullCommand():
		//plamsa_cli get balance --address=0x944A81BeECac91802787fBCFB9767FCBf81db1f5 --watcher=http://watcher.path.net
		ownerBalanceAddress := *ownerUTXOAddress
		if len(ownerBalanceAddress) == 0 {
			log.Error("Address is required to get balance")
			os.Exit(1)
		}
		balance, err := plasma.GetBalance(ownerBalanceAddress, *watcherURL)
		if err != nil {
			log.Error(err)
		}
		plasma.DisplayBalance(balance)
	case status.FullCommand():
		//plamsa_cli get status --watcher=http://watcher.path.net
		ws, err := plasma.GetWatcherStatus(*watcherURL)
		if err != nil {
			log.Error(err)
		}
		plasma.DisplayByzantineEvents(ws)
	case deposit.FullCommand():
		//plasma_cli deposit --privatekey=0x944A81BeECac91802787fBCFB9767FCBf81db1f5 --client=https://rinkeby.infura.io/v3/api_key --contract=0x457e2ec4ad356d3cb449e3bd4ba640d720c30377 --currency=ETH
		d := plasma.PlasmaDeposit{PrivateKey: *privateKey, Client: *client, Contract: *contract, Amount: *depositAmount, Owner: *depositOwner, Currency: *depositCurrency}
		d.DepositToPlasmaContract()
	case send.FullCommand():
		//plasma_cli send --to --privatekey --amount --currency --watcher
		p := plasma.NewCreateTransaction()
		p.Owner = util.DeriveAddress(*privatekey)
		p.Fee = plasma.Fee{Amount: *feeamount, Currency: *feetoken}
		p.Metadata = *metadata
		payment := plasma.Payments{
			Amount:   *amount,
			Owner:    *to,
			Currency: *currency,
		}
		p.Payments = []plasma.Payments{payment}
		p.WatcherEndpoint = *watcherSubmitURL
		resp, _ := p.CreateTransaction()

		pk := *privatekey
		hash, _ := hex.DecodeString(util.FilterZeroX(resp.Data.Transactions[0].GetToSignHash()))
		signatures, _ := plasma.Sign(plasma.SingleSigner{ToSign: hash, PrivateKey: pk})
		tx, _ := plasma.CreateTypedTransaction(
			resp.Data.Transactions[0].GetTypedData().Domain,
			resp.Data.Transactions[0].GetTypedData().Message,
			signatures,
			*watcherSubmitURL,
		)
		r, _ := plasma.Submit(tx)
		plasma.DisplaySubmitResponse(r)
	case transaction.FullCommand():
		gtx, err := plasma.GetTransaction(*txHash, *watcherURL)
		if err != nil {
			log.Errorf("got error: %v", err)
		}
		plasma.DisplayGetResponse(gtx)
	case exit.FullCommand():
		//plasma_cli exit --utxo=1000000000 --privatekey=foo --contract=0x5bb7f2492487556e380e0bf960510277cdafd680 --watcher=ari.omg.network
		s := plasma.StandardExit{UtxoPosition: util.ConvertStringToInt(*utxoPosition), Contract: *contractExit, PrivateKey: *exitPrivateKey, Client: *clientExit}
		log.Info("Attempting to exit UTXO ", *utxoPosition)
		s.StartStandardExit(*watcherExitURL)
	case process.FullCommand():
		//plasma_cli process --contract=0x5bb7f2492487556e380e0bf960510277cdafd680 --token 0x0 --privatekey=foo --client=https://rinkeby.infura.io/v3/api_key
		p := plasma.ProcessExit{Contract: *processContract, PrivateKey: *processPrivateKey, Token: *processToken, Client: *processExitClient}
		log.Info("Calling process exits in the Plasma contract")
		plasma.ProcessExits(100, p)
	case createAccount.FullCommand():
		//plasma_cli create account
		log.Info("Generating Keypair")
		util.GenerateAccount()
	case getExit.FullCommand():
		//plasma_cli get exit --watcher=https://watcher.ari.omg.network --utxo=1000000000000
		log.Info("Getting UTXO exit data")
		exitData, err := plasma.GetUTXOExitData(*watcherURL, *getExitUTXOPosition)
		if err != nil {
			log.Warn(err)
			os.Exit(0)
		}
		log.Info("UTXO Position: ", exitData.Data.UtxoPos, " Proof: ", exitData.Data.Proof)
	case challengeexit.FullCommand():
		//plasma_cli challengeexit --contract="" --client="" --privatekey="" --utxo="" --watcher=""
		challengeData, err := plasma.GetChallengeData(*ceWatcherURL, util.ConvertStringToInt(*ceUTXOPosition))
		log.Info(challengeData)
		if err != nil {
			log.Fatalf("got error retrieving challenge data %v", err)
		}
		challengeData.ChallengeInvalidExit(*ceClient, *ceContract, *cePrivateKey)
	}
}
