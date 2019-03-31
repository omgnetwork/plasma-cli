# OmiseGO Plasma CLI

## What is this?

The Plasma CLI is used to interact with the OMG network from the command line. Technical details about Ethereum and Plasma are abstracted from the user.

Plasma CLI enables the user to:
* deposit ETH to the Plasma contract
* create transactions
* exit ETH from the OMG Network back to Ethereum

Connectivity to Ethereum via a local RPC node or Infura is supported.

ERC20 tokens are not currently supported.

## Installation

`plasma_cli` is available as binaries at https://github.com/omisego/plasma-cli/releases. Download the binary that matches your operating system.

After downloading the binary file (on macOS):
1. rename the file `$ mv "plasma-cli-v0.0.2-darwin-amd64" plasma_cli`
2. set execute permissions `$ chmod u+x ./plasma_cli`
3. run it `$ ./plasma_cli --help`

## Create an Account (Keypair)

```
plasma_cli create account
```

## Deposit ETH into the OMG Network

```
plasma_cli deposit --privatekey=<private key> --client=<local rpc server or Infura URL> --contract=<contract address> --amount=<amount in wei> --owner=<public key of the owner>
```


## Retrieve a List of UTXOs

```
plasma_cli get utxos --watcher=<watcher URL> --address=<public_address>
```

## Check the Balance of an Account

```
plasma_cli get balance --watcher=<watcher URL> --address=<public_address>
```

## Send ETH in the OMG Network

This function will either
1. Send the entire UTXO to an address, or
2. Split the UTXO to destination address and send the change to the owner

```
plasma_cli send --fromutxo=<UTXO position> --fromowner=<from address> --privatekey=<from privatekey> --toowner=<to address> --toamount=<to amount>  --watcher=<watcher url>
```

## Merging UTXOs

This function allows for merging 4 or less UTXOs into 1
```
plasma_cli merge --fromutxo=UTXO_1 --fromutxo=UTXO_2 --fromutxo=UTXO_3 --privatekey="private_key" --fromowner="from-owner"
```

## Exit to Ethereum

```
plasma_cli exit --utxo=<utxo ID> --privatekey=<private key> --contract=<contract address> --watcher=<watcher url> --client=<local rpc server or Infura URL>
```
