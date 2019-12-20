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
	"testing"
	"github.com/omisego/plasma-cli/util"
	"fmt"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type TestEnv struct {
	Watcher, Privatekey, Publickey, EthClient, EthVault, ExitGame, PlasmaFramework string
	UtxoPos, DepositAmount, ExitToProcess int

}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := os.Getenv(name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func loadTestEnv() (*TestEnv, error){
	err := godotenv.Load("./integration_test.env")
	if err != nil {
		return nil, fmt.Errorf("error loading env in test: %v", err)
	}
	env := TestEnv{
		Watcher: os.Getenv("WATCHER"),
		Privatekey: os.Getenv("PRIVKEY"),
		Publickey: os.Getenv("PUBKEY"),
		EthClient: os.Getenv("ETH_CLIENT"),
		EthVault: os.Getenv("ETH_VAULT_CONTRACT"),
		ExitGame: os.Getenv("EXIT_GAME_CONTRACT"),
		PlasmaFramework: os.Getenv("PLASMA_FRAMEWORK_CONTRACT"),
		UtxoPos: getEnvAsInt("UTXO_POS_FOR_EXIT", 1),
		DepositAmount: getEnvAsInt("DEPOSIT_AMOUNT",1),
		ExitToProcess: getEnvAsInt("EXIT_TO_PROCESS",1),
	}
	return &env, nil
}
func TestGetStandardExitBond(t *testing.T) {
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error loading test env in get exit bond test: %v", err)
	}
	res, err := GetStandardExitBond(env.EthClient, env.ExitGame, env.Privatekey)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	fmt.Printf("exit bond value: %v", res)
}

func TestStartStandardExit(t *testing.T){
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error loading test env in standard exit test: %v", err)
	}
	// get standard exit bond
	bond, err := GetStandardExitBond(env.EthClient, env.ExitGame, env.Privatekey)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	// get exit UTXO info from watcher
	exit, err := GetUTXOExitData(env.Watcher, env.UtxoPos)
	if err != nil {
t.Errorf("unexpected error from getting UTXO exit data %v", err)
	}
	// call start standard exit 
	res, err := exit.StartStandardExit(env.EthClient, env.ExitGame, env.Privatekey, bond)
	if err != nil {
t.Errorf("unexpected error from, starting exit: %v", err)
	}
	fmt.Printf("standard exit tx hash: %s \n", res)
}

func TestDepositEth(t *testing.T) {
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error loading test env in standard exit test: %v", err)
	}

	d := PlasmaDeposit{PrivateKey: env.Privatekey, Client: env.EthClient, Contract: env.EthVault, Amount: uint64(env.DepositAmount), Owner: util.DeriveAddress(env.Privatekey), Currency: util.EthCurrency}
	res,  err := d.DepositEthToPlasma()
	if err != nil {
		t.Error(err)
	} 
	fmt.Printf("deposit tx hash: %s \n", res)
}

func TestProcessExit(t *testing.T) {
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error loading test env in standard exit test: %v", err)
	}

	p := ProcessExit{Contract: env.PlasmaFramework, PrivateKey: env.Privatekey, Token: util.EthCurrency, Client: env.EthClient}
	res, err := ProcessExits(1, int64(env.ExitToProcess), p)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("process exit tx hash: %s \n", res)

}
