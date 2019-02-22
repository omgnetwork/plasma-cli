package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	//"crypto/rand"
	"encoding/json"
	"encoding/hex"
	//"encoding/binary"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"strconv"

	"./rootchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	get = kingpin.Command("get", "Get a resource.").Default()
	getUTXO = get.Command("utxos", "Retrieve UTXO data from the Watcher service")
	watcherURL = get.Flag("watcher", "FQDN of the Watcher in the format http://watcher.path.net").Required().String()
	ownerUTXOAddress = get.Flag("address", "Owner address to search UTXOs").String()
	status = get.Command("status", "Get status from the Watcher")
	deposit = kingpin.Command("deposit", "Deposit ETH or ERC20 into the Plamsa MoreVP smart contract.")
	privateKey = deposit.Flag("privatekey", "Private key of the account used to send funds into Plasma MoreVP").Required().String()
	client = deposit.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()
	contract = deposit.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	transaction = kingpin.Command("transaction", "Create a transaction on the OmiseGO Plasma MoreVP network")
	blknum = transaction.Flag("block", "Block number").Required().String()
	txindex = transaction.Flag("txindex", "Transaction Index").Required().String()
	oindex = transaction.Flag("oindex", "Output Index").Required().String()
	cur12 = transaction.Flag("currency", "Currency of the transaction").Required().String()
	newowner = transaction.Flag("newowner", "New owner of the UTXO").Required().String()
	amount = transaction.Flag("amount", "Amount to transact").Required().String()
	signing = transaction.Flag("key", "Private key for signing the transaction").Required().String()
	exit = kingpin.Command("exit", "Standard exit a UTXO back to the root chain.")
	watcherExitURL = exit.Flag("watcher", "FQDN of the Watcher in the format http://watcher.path.net").Required().String()
	contractExit = exit.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	utxoPosition = exit.Flag("utxo", "UTXO Position that will be exited").Required().String()
	exitPrivateKey = exit.Flag("privatekey", "Private key of the UTXO owner").Required().String()
	clientExit = exit.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()
	process = kingpin.Command("process", "Process all exits that have completed the challenge period")
	processContract = process.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	processToken = process.Flag("token", "Token address to process for standard exits").Required().String()
	processPrivateKey = process.Flag("privatekey", "Private key used to fund the gas for the smart contract call").Required().String()
	processExitClient = process.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()
	create = kingpin.Command("create", "Create a resource.")
	createAccount = create.Command("account", "Create an account consisting of Public and Private key")
)

type processExit struct {
	contract string
	privateKey string
	token string
	client string
}

type standardExit struct {
	utxoPosition int
	privateKey string
	contract string
	client string
}

type standardExitUTXOError struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Object      string `json:"object"`
		Code        string `json:"code"`
		Description string `json:"description"`
		Messages    struct {
			ErrorKey string `json:"error_key"`
		} `json:"messages"`
	} `json:"data"`
}

type standardExitUTXOData struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		UtxoPos *big.Int  `json:"utxo_pos"`
		Txbytes string `json:"txbytes"`
		Proof   string `json:"proof"`
	} `json:"data"`
}

type plasmaTransaction struct {
	blknum uint
	txindex uint
	oindex uint
	cur12 common.Address 
	newowner common.Address 
	oldowner common.Address 
	toamount uint
	fromamount uint
}

type plasmaDeposit struct {
	privateKey string
	client string
	contract string
}

type watcherStatus struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		LastValidatedChildBlockNumber int           `json:"last_validated_child_block_number"`
		LastMinedChildBlockTimestamp  int           `json:"last_mined_child_block_timestamp"`
		LastMinedChildBlockNumber     int           `json:"last_mined_child_block_number"`
		EthSyncing                    bool          `json:"eth_syncing"`
		ByzantineEvents               []interface{} `json:"byzantine_events"`
	} `json:"data"`
}

type watcherUTXOsFromAddress struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    []struct {
		UtxoPos  int    `json:"utxo_pos"`
		Txindex  int    `json:"txindex"`
		Owner    string `json:"owner"`
		Oindex   int    `json:"oindex"`
		Currency string `json:"currency"`
		Blknum   int    `json:"blknum"`
		Amount   int    `json:"amount"`
	} `json:"data"`
}

type watcherError struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Object      string `json:"object"`
		Description string `json:"description"`
		Code        string `json:"code"`
	} `json:"data"`
}

type ownerUTXOs struct {
	Txindex  int    `json:"txindex"`
	Txbytes  string `json:"txbytes"`
	Oindex   int    `json:"oindex"`
	Currency common.Address `json:"currency"`
	Blknum   int    `json:"blknum"`
	Amount   int    `json:"amount"`
}

type inputUTXO struct {
	Txindex  uint    `json:"txindex"`
	Oindex   uint    `json:"oindex"`
	Currency common.Address `json:"currency"`
	Blknum   uint    `json:"blknum"`
	Amount   uint    `json:"amount"`
}

type outputUTXO struct {
	OwnerAddress common.Address  `json:"owner"`
	Amount uint `json:"amount"`
	Currency common.Address `json:"currency"`
}

type createdTx struct {
	Inputs []inputUTXO `json:"inputs"`
	Outputs []outputUTXO `json:"outputs"`
}

type input struct {
	Blknum uint
	Txindex uint
	Oindex uint
}

type output struct {
	OwnerAddress common.Address
	Currency common.Address
	Amount uint
}

type transactionToEncode struct {
	Inputs []input
	Outputs []output
}

type transactionToBuild struct {
	Sig Signature
	Inputs []input
	Outputs []output
}

const signatureLength = 65

type Signature struct {
	Sig []byte
}

//try big int, int and bytes32
type ecdsaSignature struct {
	R []byte
	S []byte
	V uint8
}

type transactionSuccessResponse struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Blknum  int    `json:"blknum"`
		Txindex int    `json:"txindex"`
		Txhash  string `json:"txhash"`
	} `json:"data"`
}

type transactionFailureResponse struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Object      string `json:"object"`
		Code        string `json:"code"`
		Description string `json:"description"`
		Messages    struct {
			ErrorKey string `json:"error_key"`
		} `json:"messages"`
	} `json:"data"`
}

type blockNumber struct {
	Hash string `json:"hash"`
}

type blockNumberError struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Object   string `json:"object"`
		Messages struct {
			ValidationError struct {
				Validator string `json:"validator"`
				Parameter string `json:"parameter"`
			} `json:"validation_error"`
		} `json:"messages"`
		Description string `json:"description"`
		Code        string `json:"code"`
	} `json:"data"`
}


// Add the full time include timezone into log messages
// INFO[2019-01-31T16:38:57+07:00]
func logFormatter() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)
}

//Retrieve the UTXOs associated with an address from the Watcher
func getUTXOsFromAddress(address string, w string) watcherUTXOsFromAddress {
	// Build request
	var url strings.Builder
	url.WriteString(w)
	url.WriteString("/account.get_utxos")
	postData := map[string]interface{}{"address": address, "limit": "10000"}
	js, _ := json.Marshal(postData)
	r, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(js))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Unmarshall the response
	response := watcherUTXOsFromAddress{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		log.Warning("Could not unmarshal successful response from the Watcher")
		errorInfo := watcherError{}
		processError := json.Unmarshal([]byte(rstring), &errorInfo)
		if processError != nil { // Response from the Watcher does not match a struct
			log.Fatal("Unknown response from Watcher API")
			panic("uh oh")
		}
		log.Warning("Unmarshalled JSON error response from the Watcher API")
		log.Error(errorInfo)
	} else {
		log.Info(resp.Status)
	}

	return response
}

//Deposit ETH into the already deployed Plasma MoreVP contract on Ethereum
func (d *plasmaDeposit) depositToPlasmaContract() {
	client, err := ethclient.Dial(d.client)
	if err != nil {
		log.Fatal(err)
	}
	
	privateKey, err := crypto.HexToECDSA(filterZeroX(d.privateKey))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(100)   // in wei
	auth.GasLimit = uint64(210000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(d.contract)
	// TODO(jbunce) Add support for the bytes representation of the RLP encoding from
	// a struct of the values. This implementation supports a fixed address, value,
	// and currency.
	//owner := "\x94J\x81\xbe\xec\xac\x91\x80'\x87\xfb\xcf\xb9v\x7f\xcb\xf8\x1d\xb1\xf5"
	rlpInputs := []byte("\xf8\xc3\xd0\xc3\x80\x80\x80\xc3\x80\x80\x80\xc3\x80\x80\x80\xc3\x80\x80\x80\xf8\xb0\xeb\x94\x94J\x81\xbe\xec\xac\x91\x80'\x87\xfb\xcf\xb9v\x7f\xcb\xf8\x1d\xb1\xf5\x94\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00d\xeb\x94\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x94\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x80\xeb\x94\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x94\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x80\xeb\x94\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x94\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x80")
	instance, err := rootchain.NewRootchain(address, client)
	if err != nil {
		log.Fatal(err)
	}
	t := &bind.TransactOpts{}
	t.From = fromAddress
	t.Signer = auth.Signer
	t.Value = big.NewInt(100)
	tx, err := instance.Deposit(t, rlpInputs)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Deposit to Plasma MoreVP contract sent. Transaction: ", tx.Hash().Hex())
	}
}

func getWatcherStatus(w string) {
	var url strings.Builder
	url.WriteString(w)
	url.WriteString("/status.get")
	r, err := http.NewRequest("POST", url.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Unmarshall the response
	response := watcherStatus{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		log.Error(jsonErr)
	}
	log.Info("Last validated Childchain block number: ", response.Data.LastValidatedChildBlockNumber)
	log.Info("Last mined Childchain block number: ", response.Data.LastMinedChildBlockNumber)
}

// Displays the UTXO ownership data in a human friendly format
func (u *watcherUTXOsFromAddress) displayUTXOS() {
	for _, value := range u.Data {
		log.Info("Owner: ", value.Owner)
		log.Info("Amount: ", value.Amount)
		log.Info("Block Number: ", value.Blknum)
		log.Info("Transaction Index: ", value.Txindex)
		log.Info("Output Index: ", value.Oindex)
		log.Info("Currency: ", value.Currency)
		log.Info("UTXO Position: ", value.UtxoPos)
	}
}

// Create a transaction from the user input from UTXO data.
// This needs RLP encoding and then signing with the private
// key of the sending party.
/*
func (t *plasmaTransaction) createPlasmaTransaction(privateKey string) {
	// RLP encode the transaction
	//transaction = createBasicTransaction(*t)
	transaction := *t
	rlpInputs, err := rlp.EncodeToBytes(transaction)
	if err != nil {
		log.Fatal("Error encoding RLP")
	} else {
		log.Info("Transaction data structure RLP encoded")
	}
	fmt.Printf("%v", rlpInputs)
  //rlpDecoded := rlp.DecodeBytes(rlpInputs, os.Stdout)
	//log.Printf("%v", rlpDecoded)
	//r, s, serr := ecdsa.Sign(rand.Reader, privateKey, hash)
}
*/
func convertStringToInt(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal("Could not convert value")
	}
	return i
}

func convertStringToFloat64(value string) float64 {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatal("Could not convert value")
	}
	return f
}

//Start a standard exit from user provided UTXO & private key
func (s *standardExit) startStandardExit(watcher string) {
	log.Info("Getting data needed to exit the UTXO from the Watcher")
	exit := getUTXOExitData(watcher, s.utxoPosition)
	exit.plasmaStartStandardExit(s.client, s.contract, s.privateKey)
}

//Start standard exit by calling the method in the smart contract
func (s *standardExitUTXOData) plasmaStartStandardExit(ethereumClient string, contract string, private string) {
	client, err := ethclient.Dial(ethereumClient)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(filterZeroX(private))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(31415926535)    // in wei
	auth.GasLimit = uint64(210000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(contract)
	instance, err := rootchain.NewRootchain(address, client)
	if err != nil {
		log.Fatal(err)
	}
	t := &bind.TransactOpts{}
	t.From = fromAddress
	t.Signer = auth.Signer
	t.Value = big.NewInt(31415926535) //STANDARD_EXIT_BOND in the smart contract
	t.GasLimit = 2000000

	txBytesHex, txErr := hex.DecodeString(removeLeadingZeroX(s.Data.Txbytes))
	if txErr != nil{
		log.Fatal(txErr)
	}

	proofBytesHex, proofErr := hex.DecodeString(removeLeadingZeroX(s.Data.Proof))
	if proofErr != nil {
		log.Fatal(proofErr)
	}
	tx, err := instance.StartStandardExit(t, s.Data.UtxoPos, []byte(txBytesHex), []byte(proofBytesHex))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Standard exit to Plasma MoreVP sent. Transaction: ", tx.Hash().Hex())
	}
}

//Make strings suitable for hex encoding
func removeLeadingZeroX(item string) string {
	cleanedString := strings.Replace(item, "0x", "", -1)
	return cleanedString
}

//filter our key string with 0x
func filterZeroX(item string) string {
	if strings.Contains(item, "0x") {
		return removeLeadingZeroX(item)
	} else {
		return item
	}
}

//Retrieve the UTXO exit data from the UTXO position
func getUTXOExitData(watcher string, utxoPosition int) standardExitUTXOData {
	// Build request
	var url strings.Builder
	url.WriteString(watcher)
	url.WriteString("/utxo.get_exit_data")
	postData := map[string]interface{}{"utxo_pos": utxoPosition}
	js, _ := json.Marshal(postData)
	r, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(js))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Unmarshall the response
	response := standardExitUTXOData{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		log.Warning("Could not unmarshal successful response from the Watcher")
		errorInfo := standardExitUTXOError{}
		processError := json.Unmarshal([]byte(rstring), &errorInfo)
		if processError != nil { // Response from the Watcher does not match a struct
			log.Fatal("Unknown response from Watcher API")
			panic("uh oh")
		}
		log.Warning("Unmarshalled JSON error response from the Watcher API")
		log.Error(errorInfo)
	} else {
		log.Info(resp.Status)
	}

	return response
}

//Calls the processExits in the Plasma smart contract to start processing exits that
//have completed the challenge period.
func (p *processExit) plasmaProcessExits(numberExitsToProcess int64) {
	client, err := ethclient.Dial(p.client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(filterZeroX(p.privateKey))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)    // in wei
	auth.GasLimit = uint64(210000) // in units
	auth.GasPrice = gasPrice

	contractAddress := p.contract

	address := common.HexToAddress(contractAddress)
	instance, err := rootchain.NewRootchain(address, client)
	if err != nil {
		log.Fatal(err)
	}
	t := &bind.TransactOpts{}
	t.From = fromAddress
	t.Signer = auth.Signer
	t.Value = big.NewInt(0)
	t.GasLimit = 2000000

	token := common.HexToAddress(p.token)


	tx, err := instance.ProcessExits(t, token, big.NewInt(0), big.NewInt(numberExitsToProcess))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Process exits request to Plasma MoreVP sent. Transaction: ", tx.Hash().Hex())
	}
}

//generate Account - Public and Privatekey
func generateAccount() {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	privateKey := hex.EncodeToString(key.D.Bytes())
	log.Info("Address is ", address)
	log.Info("Privatekey is ", privateKey)
}

//signature Signature
func signTransaction(unsignedTx string, privateKey string) []byte {
	//hash the unsignedTx struct
	unsignedTxBytes, _ := hex.DecodeString(filterZeroX(unsignedTx))
	hashed := crypto.Keccak256(unsignedTxBytes)
	priv, _ := crypto.HexToECDSA(filterZeroX(privateKey))
	//sign the transaction
	signaturebytes, _ := crypto.Sign(hashed, priv)
	fmt.Println(signaturebytes)
	fmt.Println(signaturebytes[0:32])
	fmt.Println(signaturebytes[32:64])
	fmt.Println(signaturebytes[64])
	//fmt.Println(hex.EncodeToString(signaturebytes))
	
	//copy(signature[:], signaturebytes)
	return signaturebytes //ecdsaSignature{R: signaturebytes[0:32], S: signaturebytes[32:64], V: 28}
}

//simple transaction implementation, taking a single UTXO and split into two outputs
/*
type plasmaTransaction struct {
	blknum int
	txindex int
	oindex int
	cur12 string
	newowner string
	amount float64
}

steps:
1. pickBasicUtxo(), loop through all UTXOs and return the biggest one
2. createBasicTransaction(), single input, two outputs
3. 
*/

func (w *watcherUTXOsFromAddress) pickBasicUtxo() {

}

/*
{"inputs":[{"utxo_pos":55000000000001,"txindex":0,"owner":"0x7dafb4442c112c76f2437af717bb8a1ea0f146bf","oindex":1,"currency":"0x0000000000000000000000000000000000000000","blknum":55000,"amount":63}],"outputs":[{"owner":"0xc53aee876d24b2e0634362f732118ebb67a20e14","currency":"0x0000000000000000000000000000000000000000","amount":1},{"owner":"0x7dafb4442c112c76f2437af717bb8a1ea0f146bf","currency":"0x0000000000000000000000000000000000000000","amount":62}]}
*/

func (p *plasmaTransaction) createBasicTransaction() createdTx {
		//creates 1 input, 2 outputs tx
		NULL_ADDRESS := common.HexToAddress("0000000000000000000000000000000000000000")
		NULL_INPUT  := inputUTXO{Blknum: 0, Txindex: 0, Oindex: 0, Currency: NULL_ADDRESS}
		NULL_OUTPUT := outputUTXO{ OwnerAddress: NULL_ADDRESS, Amount: 0, Currency: NULL_ADDRESS }
		//1 single input
		singleInput := inputUTXO{Blknum: p.blknum, Txindex: p.txindex, Oindex: p.oindex, Currency: p.cur12}
		//output one is value you are sending
		outputOne := outputUTXO{ OwnerAddress: p.newowner, Amount: p.toamount, Currency: p.cur12 }
		//output two is the change
		outputTwo := outputUTXO{ OwnerAddress: p.oldowner, Amount: p.fromamount - p.toamount, Currency: p.cur12 }
	
		var i []inputUTXO
		var o []outputUTXO
		i = append(i, singleInput, NULL_INPUT, NULL_INPUT, NULL_INPUT)
		o = append(o, outputOne, outputTwo, NULL_OUTPUT, NULL_OUTPUT)
		utxo := createdTx{Inputs: i, Outputs: o}
	
		//js, _ := json.Marshal(utxo)
		//fmt.Println(string(js))

		return utxo
}

func (c *createdTx) encodeTransaction() string {
	var t *transactionToEncode
	var i []input
	var o []output

	for _, val := range c.Inputs {
		t := input{Txindex: val.Txindex, Oindex: val.Oindex, Blknum: val.Blknum}
		i = append(i, t)
	}

	for _, val := range c.Outputs {
		t := output{OwnerAddress: val.OwnerAddress, Amount: val.Amount, Currency: val.Currency}
		o = append(o, t)
	}
	t = &transactionToEncode{Outputs: o, Inputs: i}

	encodedBytes, err := rlp.EncodeToBytes(t)
	if err != nil {
		fmt.Println(err)
	} 
	//fmt.Println(hex.EncodeToString(hey))
	return hex.EncodeToString(encodedBytes)
}


func buildSignedTransaction(signature []byte, unsignedTX string) {

	var tx transactionToEncode
	//var m type interface{}
	//RLP decode unsignedTx
	decoded, _ := hex.DecodeString(unsignedTX)
	rlp.DecodeBytes(decoded, &tx)
	//build Transaction
	realsig := Signature{Sig: signature}
	builtTx := transactionToBuild{Inputs: tx.Inputs, Outputs: tx.Outputs, Sig: realsig}
	//builtTx := m{tx, signature}
	//fmt.Println("struct is", binary.LittleEndian.Uint64(builtTx.Sig))

	//fmt.Println("struct is", builtTx)

	//RLP encode built transaction
	encoded, _ := rlp.EncodeToBytes(builtTx)
	fmt.Println(hex.EncodeToString(encoded))
	//var txx *transactionToBuild
	expected, _ := hex.DecodeString("f9010bf843b8413b02b7388ef50c8267287fa9a0cfde2798427d93aa10e1af77db0273b98622e42f08899d2973d6bfef8d4b2d904f66e3e3bcbf87ea388ae6ca7b09d0277221421bd3c6830134988001c3808080c3808080c3808080f8b0eb94c53aee876d24b2e0634362f732118ebb67a20e1494000000000000000000000000000000000000000001eb947dafb4442c112c76f2437af717bb8a1ea0f146bf94000000000000000000000000000000000000000038eb94000000000000000000000000000000000000000094000000000000000000000000000000000000000080eb94000000000000000000000000000000000000000094000000000000000000000000000000000000000080")
	//fmt.Println("answer should be", expected)
	//rlp.DecodeBytes(expected, txx)
	fmt.Println("answer should be", expected)
	//fmt.Println(txx.Transaction)
	/*
	//Append the signature with decoded unsignedTx list
	signaturesBytes, _ := hex.DecodeString(filterZeroX(signatures))
	s := append(tx, signaturesBytes...)
	//fmt.Println(s)
	encoded, _ := rlp.EncodeToBytes(s)
	fmt.Println(encoded)
	fmt.Println(hex.EncodeToString(encoded))
	//rlp encode both transaction + signature
	*/
}

func main() {
	//fmt.Println(signTransaction("0xF8C5D2C582CF088001C3808080C3808080C3808080F8B0EB94C53AEE876D24B2E0634362F732118EBB67A20E1494000000000000000000000000000000000000000002EB947DAFB4442C112C76F2437AF717BB8A1EA0F146BF94000000000000000000000000000000000000000040EB94000000000000000000000000000000000000000094000000000000000000000000000000000000000080EB94000000000000000000000000000000000000000094000000000000000000000000000000000000000080", "0x17b5c8b364f2687f416f2c414f9482abba984b9cf833cdfa20157377975d156e"))
	
	c := plasmaTransaction{blknum: 43000, txindex: 0, oindex: 1, cur12: common.HexToAddress("0000000000000000000000000000000000000000"), newowner: common.HexToAddress("0x7dafb4442c112c76f2437af717bb8a1ea0f146bf"), oldowner: common.HexToAddress("0xc53aee876d24b2e0634362f732118ebb67a20e14"), toamount: 1000, fromamount: 9801}
	k := c.createBasicTransaction()																			  
	encoded := k.encodeTransaction()
	sig := signTransaction(encoded, "3a0adcf75ed744225fbe3b07ba949d1a26685c2ee56d38025c2762dbdd4a73c6")
	buildSignedTransaction(sig, encoded)

	logFormatter()
	log.Info("Starting OmiseGO Plasma MoreVP CLI")
	switch kingpin.Parse() {
	case getUTXO.FullCommand():
		//plasma_cli get utxos --address=0x944A81BeECac91802787fBCFB9767FCBf81db1f5 --watcher=http://watcher.path.net
		ownerUTXOAddress := *ownerUTXOAddress
		if len(ownerUTXOAddress) == 0 {
			log.Error("Address is required to get UTXO data")
			os.Exit(1)
		}
		utxos := getUTXOsFromAddress(ownerUTXOAddress, *watcherURL)
		utxos.displayUTXOS()
	case status.FullCommand():
		//plamsa_cli get status --watcher=http://watcher.path.net
		getWatcherStatus(*watcherURL)
	case deposit.FullCommand():
		//plasma_cli deposit --privatekey=0x944A81BeECac91802787fBCFB9767FCBf81db1f5 --client=https://rinkeby.infura.io/v3/api_key --contract=0x457e2ec4ad356d3cb449e3bd4ba640d720c30377
		d := plasmaDeposit{privateKey: *privateKey, client: *client, contract: *contract}
		d.depositToPlasmaContract()
	case transaction.FullCommand():
		//plasma_cli transaction --block --txindex --oindex --currency --newowner --amount
		//t := plasmaTransaction{blknum: convertStringToInt(*blknum), txindex: convertStringToInt(*txindex), oindex: convertStringToInt(*oindex), cur12: *cur12, newowner: *newowner, amount: convertStringToFloat64(*amount)}
		//t.createPlasmaTransaction(*signing)
	case exit.FullCommand():
		//plasma_cli exit --utxo=1000000000 --privatekey=foo --contract=0x5bb7f2492487556e380e0bf960510277cdafd680 --watcher=watcher-staging.omg.network
		s := standardExit{utxoPosition: convertStringToInt(*utxoPosition), contract: *contractExit, privateKey: *exitPrivateKey, client: *clientExit}
		log.Info("Attempting to exit UTXO ", *utxoPosition)
		s.startStandardExit(*watcherExitURL)
	case process.FullCommand():
		//plasma_cli process --contract=0x5bb7f2492487556e380e0bf960510277cdafd680 --token 0x0 --privatekey=foo --client=https://rinkeby.infura.io/v3/api_key
		p := processExit{contract: *processContract, privateKey: *processPrivateKey, token: *processToken, client: *processExitClient}
		log.Info("Calling process exits in the Plasma contract")
		p.plasmaProcessExits(100)
	case createAccount.FullCommand():
		//plasma_cli create account 
		log.Info("Generating Keypair")
		generateAccount()
	}
}
