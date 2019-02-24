package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	//"crypto/rand"
	"encoding/json"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"os"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"strconv"

	"rootchain"

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
	depositOwner = deposit.Flag("owner", "Owner of the UTXOs").Required().String()
	depositAmount = deposit.Flag("amount", "Amount to deposit in wei").Required().String()
	transaction = kingpin.Command("transaction", "Create a transaction on the OmiseGO Plasma MoreVP network")
	blknum = transaction.Flag("blknum", "Block number").Required().Uint()
	txindex = transaction.Flag("txindex", "Transaction Index").Required().Uint()
	oindex = transaction.Flag("oindex", "Output Index").Required().Uint()
	cur12 = transaction.Flag("cur12", "Currency of the transaction").Required().String()
	toowner = transaction.Flag("toowner", "New owner of the UTXO").Required().String()
	fromowner = transaction.Flag("fromowner", "from an owner of the UTXO").Required().String()
	privatekey = transaction.Flag("privatekey", "privatekey to sign from owner of original UTXO").Required().String()
	toamount = transaction.Flag("toamount", "Amount to transact").Required().Uint()
	fromamount = transaction.Flag("fromamount", "original amount").Required().Uint()
	watcherSubmitURL = transaction.Flag("watcher", "FQDN of the Watcher in the format http://watcher.path.net").Required().String()
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

type Inner struct {
	Values []inputDeposit
	Values2 []interface{}
}

type InputTwo struct {
	OwnerAddress common.Address
    Currency common.Address
    Amount uint64
}

type deposits struct {
    OwnerAddress common.Address
    Currency common.Address
    Amount uint64
}

type second struct {
	Currency1 common.Address
	Currency2 common.Address
	Value uint
}

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
	toowner common.Address
	fromowner common.Address
	toamount uint
	fromamount uint
	privatekey string
}

type plasmaDeposit struct {
	privateKey string
	client string
	contract string
	amount string
	owner string
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

type inputDeposit struct {
	Txindex  uint    `json:"txindex"`
	Oindex   uint    `json:"oindex"`
	Blknum   uint    `json:"blknum"`
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

	amount, err := strconv.ParseInt(d.amount, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))

	auth.Value = big.NewInt(amount)   // in wei
	auth.GasLimit = uint64(210000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(d.contract)

	rlpInputs := buildRLPInput(removeLeadingZeroX(d.owner), d.amount)
	instance, err := rootchain.NewRootchain(address, client)
	if err != nil {
		log.Fatal(err)
	}
	t := &bind.TransactOpts{}
	t.From = fromAddress
	t.Signer = auth.Signer
	t.Value = big.NewInt(amount)
	tx, err := instance.Deposit(t, rlpInputs)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Deposit to Plasma MoreVP contract sent. Transaction: ", tx.Hash().Hex())
	}
}

// Build the RLP encoded input to the smart contract
// that deposit will accept.
func buildRLPInput(address string, value string) []byte {
	// [[[0,0,0],[0,0,0],[0,0,0],[0,0,0]],[[alice_raw, eth_raw, 10], [eth_raw, eth_raw, 0], [eth_raw, eth_raw, 0], [eth_raw, eth_raw, 0]]]
	v := make([]inputDeposit, 0)

	test := Inner{}

	NULL_INPUT := inputDeposit{Blknum: 0, Txindex: 0, Oindex: 0}
	v = append(v, NULL_INPUT)
	v = append(v, NULL_INPUT)
	v = append(v, NULL_INPUT)
	v = append(v, NULL_INPUT)

	test.Values = v

	cur := common.HexToAddress("0000000000000000000000000000000000000000")
	amount, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	t3 := InputTwo{}
	t3.OwnerAddress = common.HexToAddress(address)
	t3.Currency = cur
	t3.Amount = uint64(amount)

	s1 := second{Currency1: cur, Currency2: cur, Value: 0}
	s2 := second{Currency1: cur, Currency2: cur, Value: 0}
	s3 := second{Currency1: cur, Currency2: cur, Value: 0}

	arr := make([]interface{}, 0)
	arr = append(arr, t3)
	arr = append(arr, s1)
	arr = append(arr, s2)
	arr = append(arr, s3)

	test.Values2 = arr

	rlpEncoded, rerr := rlp.EncodeToBytes(test)
	if rerr != nil {
		log.Fatal(rerr)
	}

	return rlpEncoded
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

//sign a transaction with a key
func signTransaction(unsignedTx string, privateKey string) []byte {
	//hash the unsignedTx struct
	unsignedTxBytes, err := hex.DecodeString(filterZeroX(unsignedTx))
	if err != nil {
		log.Fatal(err)
	}
	hashed := crypto.Keccak256(unsignedTxBytes)
	priv, err := crypto.HexToECDSA(filterZeroX(privateKey))
	if err != nil {
		log.Fatal(err)
	}
	//sign the transaction
	signature, err := crypto.Sign(hashed, priv)
	if err != nil {
		log.Fatal(err)
	}
	//adding 27 to last byte, because ethereum
	copy(signature[64:], []uint8{signature[64] + 27})

	log.Info("transaction signed: ", signature)
	return signature
}

//create a basic transaction with 1 input splitted into 2 outputs
func (p *plasmaTransaction) createBasicTransaction() createdTx {
	//creates 1 input, 2 outputs tx
	NULL_ADDRESS := common.HexToAddress("0000000000000000000000000000000000000000")
	NULL_INPUT  := inputUTXO{Blknum: 0, Txindex: 0, Oindex: 0, Currency: NULL_ADDRESS}
	NULL_OUTPUT := outputUTXO{ OwnerAddress: NULL_ADDRESS, Amount: 0, Currency: NULL_ADDRESS }
	//1 single input
	singleInput := inputUTXO{Blknum: p.blknum, Txindex: p.txindex, Oindex: p.oindex, Currency: p.cur12}
	//output one is value you are sending
	outputOne := outputUTXO{ OwnerAddress: p.toowner, Amount: p.toamount, Currency: p.cur12 }
	//output two is the change
	if p.fromamount < p.toamount {
		log.Fatal("UTXO not large enough to be sent")
	}
	outputTwo := outputUTXO{ OwnerAddress: p.fromowner, Amount: p.fromamount - p.toamount, Currency: p.cur12 }

	var i []inputUTXO
	var o []outputUTXO
	i = append(i, singleInput, NULL_INPUT, NULL_INPUT, NULL_INPUT)
	o = append(o, outputOne, outputTwo, NULL_OUTPUT, NULL_OUTPUT)
	transaction := createdTx{Inputs: i, Outputs: o}

	log.Info("Raw Transaction: ", transaction)
	return transaction
}

//encode transaction with RLP
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
	log.Info("transaction UTXOs ", t)
	encodedBytes, err := rlp.EncodeToBytes(t)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Hex encoded transaction: ", encodedBytes)
	return hex.EncodeToString(encodedBytes)
}

//decode RLP, and rebuild transaction with signature, finally encode the whole thing
func buildSignedTransaction(signature []byte, unsignedTX string) []byte{
	var tx transactionToEncode
	//RLP decode unsignedTx
	decoded, err := hex.DecodeString(unsignedTX)
	if err != nil {
		log.Fatal(err)
	}
	rlp.DecodeBytes(decoded, &tx)
	//build Transaction
	txsig := Signature{Sig: signature}
	builtTx := transactionToBuild{Inputs: tx.Inputs, Outputs: tx.Outputs, Sig: txsig}
	//RLP encode built transaction
	encoded, err := rlp.EncodeToBytes(builtTx)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("transaction signed and built: ", hex.EncodeToString(encoded))
	return encoded
}

//submit transaction to endpoint, take tx byte and watcher url
func submitTransaction(tx []byte, w string) transactionSuccessResponse{
	txstring := "0x" + hex.EncodeToString(tx)

	// Build request
	var url strings.Builder
	url.WriteString(w)
	url.WriteString("/transaction.submit")
	postData := map[string]interface{}{"transaction": txstring}
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
	response := transactionSuccessResponse{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		log.Warning("Could not unmarshal successful response from the Watcher")
		errorInfo := transactionFailureResponse{}
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
	log.Info(response)
	return response
}

//minimal send transaction function, take utxo and send to an address
func (p *plasmaTransaction) sendBasicTransaction(w string) transactionSuccessResponse {
	k := p.createBasicTransaction()
	encoded := k.encodeTransaction()
	sig := signTransaction(encoded, p.privatekey)
	//log.Info(hex.EncodeToString(buildSignedTransaction(sig, encoded)))
	transaction := buildSignedTransaction(sig, encoded)
	return submitTransaction(transaction, w)
}

func main() {
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
		d := plasmaDeposit{privateKey: *privateKey, client: *client, contract: *contract, amount: *depositAmount, owner: *depositOwner}
		d.depositToPlasmaContract()
	case transaction.FullCommand():
		//plasma_cli transaction --blknum --txindex --oindex --cur12 --toowner --fromowner --privatekey --toamount --fromamount --watcher
		watcher := *watcherSubmitURL
		c := plasmaTransaction{
			blknum: *blknum,
			txindex: *txindex,
			oindex: *oindex,
			cur12: common.HexToAddress(*cur12),
			toowner: common.HexToAddress(*toowner),
			fromowner: common.HexToAddress(*fromowner),
			privatekey: *privatekey,
			toamount: *toamount,
			fromamount: *fromamount}
		c.sendBasicTransaction(watcher)
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
