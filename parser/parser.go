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
	"os"

	"github.com/omisego/plasma-cli/plasma"
	"github.com/omisego/plasma-cli/util"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	get               = kingpin.Command("get", "Get a resource.").Default()
	getUTXO           = get.Command("utxos", "Retrieve UTXO data from the Watcher service")
	watcherURL        = get.Flag("watcher", "FQDN of the Watcher in the format http://watcher.path.net").Required().String()
	ownerUTXOAddress  = get.Flag("address", "Owner address to search UTXOs").String()
	status            = get.Command("status", "Get status from the Watcher")
	deposit           = kingpin.Command("deposit", "Deposit ETH or ERC20 into the Plamsa MoreVP smart contract.")
	privateKey        = deposit.Flag("privatekey", "Private key of the account used to send funds into Plasma MoreVP").Required().String()
	client            = deposit.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()
	contract          = deposit.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	depositOwner      = deposit.Flag("owner", "Owner of the UTXOs").Required().String()
	depositAmount     = deposit.Flag("amount", "Amount to deposit in wei").Required().Uint64()
	transaction       = kingpin.Command("transaction", "Create a transaction on the OmiseGO Plasma MoreVP network")
	toowner           = transaction.Flag("toowner", "New owner of the UTXO").Required().String()
	fromowner         = transaction.Flag("fromowner", "from an owner of the UTXO").Required().String()
	privatekey        = transaction.Flag("privatekey", "privatekey to sign from owner of original UTXO").Required().String()
	toamount          = transaction.Flag("toamount", "Amount to transact").Required().Uint()
	fromutxo          = transaction.Flag("fromutxo", "utxo position to send from").Required().Uint()
	watcherSubmitURL  = transaction.Flag("watcher", "FQDN of the Watcher in the format http://watcher.path.net").Required().String()
	exit              = kingpin.Command("exit", "Standard exit a UTXO back to the root chain.")
	watcherExitURL    = exit.Flag("watcher", "FQDN of the Watcher in the format http://watcher.path.net").Required().String()
	contractExit      = exit.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	utxoPosition      = exit.Flag("utxo", "UTXO Position that will be exited").Required().String()
	exitPrivateKey    = exit.Flag("privatekey", "Private key of the UTXO owner").Required().String()
	clientExit        = exit.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()
	process           = kingpin.Command("process", "Process all exits that have completed the challenge period")
	processContract   = process.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	processToken      = process.Flag("token", "Token address to process for standard exits").Required().String()
	processPrivateKey = process.Flag("privatekey", "Private key used to fund the gas for the smart contract call").Required().String()
	processExitClient = process.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()
	create            = kingpin.Command("create", "Create a resource.")
	createAccount     = create.Command("account", "Create an account consisting of Public and Private key")
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
		utxos := plasma.GetUTXOsFromAddress(ownerUTXOAddress, *watcherURL)
		util.DisplayUTXOS(utxos)
	case status.FullCommand():
		//plamsa_cli get status --watcher=http://watcher.path.net
		plasma.GetWatcherStatus(*watcherURL)
	case deposit.FullCommand():
		//plasma_cli deposit --privatekey=0x944A81BeECac91802787fBCFB9767FCBf81db1f5 --client=https://rinkeby.infura.io/v3/api_key --contract=0x457e2ec4ad356d3cb449e3bd4ba640d720c30377
		d := plasma.PlasmaDeposit{PrivateKey: *privateKey, Client: *client, Contract: *contract, Amount: *depositAmount, Owner: *depositOwner}
		d.DepositToPlasmaContract()
	case transaction.FullCommand():
		//plasma_cli transaction --fromutxo --fromowner --privatekey --toowner --toamount --watcher
		p := plasma.GetUTXO(*fromowner, *fromutxo, *watcherSubmitURL)
		watcher := *watcherSubmitURL

		c := plasma.PlasmaTransaction{
			Blknum:     uint(p.Blknum),
			Txindex:    uint(p.Txindex),
			Oindex:     uint(p.Oindex),
			Cur12:      common.HexToAddress(p.Currency),
			Toowner:    common.HexToAddress(*toowner),
			Fromowner:  common.HexToAddress(*fromowner),
			Privatekey: *privatekey,
			Toamount:   *toamount,
			Fromamount: uint(p.Amount)}
		c.SendBasicTransaction(watcher)
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
	}
}
