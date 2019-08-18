# OmiseGO Plasma CLI

The Plasma CLI is used to interact with the OMG network from the command line. Technical details about Ethereum and Plasma are abstracted from the user.

Plasma CLI enables the user to:
* deposit ETH to the Plasma contract
* create send, split and merge transactions
* exit ETH from the OMG Network back to Ethereum

Connectivity to Ethereum via a local RPC node or Infura is supported.

ERC20 tokens are not currently supported.

## Compatibility

This tool is compatible with version 0.1 of the OMG Network

## Installation

`plasma_cli` is available as binaries at https://github.com/omisego/plasma-cli/releases. Download the binary that matches your operating system.

After downloading the latest binary file (on macOS):
1. rename the file `$ mv "name_of_your_binary_file" plasma_cli`
2. set execute permissions `$ chmod u+x ./plasma_cli`
3. run it `$ ./plasma_cli --help`

## Create an Account (Keypair)

```
plasma_cli create account
```

## Deposit ETH into the OMG Network

```
plasma_cli deposit --privatekey="private_key" --client="local_rpc_server_or_Infura_URL" --contract="contract_address" --amount=amount_in_wei --owner="owner_address"
```


## Retrieve a List of UTXOs

```
plasma_cli get utxos --watcher="watcher_URL" --address="public_address"
```

## Check the Balance of an Account

```
plasma_cli get balance --watcher="watcher_URL" --address="public_address"
```
## Send transaction

creates a send transaction from the balance of the privatekey to a recipient wallet 
(currently only support 1 recipient). Makes an ETH transaction by default

```
plasma_cli send --to="address" --privatekey="privatekey" --amount=amount --watcher="watcher_url"
```

optional arguements:
```
--currency="token_address" 
--feetoken="token_address"
--feeamount=fee_amount
--metadata="hex_string"
```

## Exit to Ethereum

```
plasma_cli exit --utxo=UTXO_position --privatekey="private_key" --contract="contract_address" --watcher="watcher_url" --client="local_rpc_server_or_Infura_URL"
```
