# OmiseGO Plasma CLI

## What is this?

The Plasma CLI is used for interacting with OMG network from the command line. It allows the user to deposit to the Plasma contract, create transactions, and exit back to the root chain with the Ethereum & MoreVP Plasma details abstracted from the user. Both local RPC node for connecting to Ethereum and Infura are supported.

The CLI currently does not support ERC20 tokens

## Installation

`plasma_cli` is available as binaries at https://github.com/omisego/plasma-cli/releases, download one of the assets, depending on your OS.

After downloading the binary file (on macOS):
1. rename the file `$ mv "plasma-cli-v0.0.2-darwin-amd64" plasma_cli`
2. set permission `$ chmod +x ./plasma_cli`
3. try running it `$ ./plasma_cli --help`

## Creating Account (Keypair)

```
plasma_cli create account
```

## Making Deposits

```
plasma_cli deposit --privatekey=<private key> --client=<local rpc server or Infura URL> --contract=<contract address> --amount=<amount in wei> --owner=<public key of the owner>
```


## Getting UTXOs

```
plasma_cli get utxos --watcher=<watcher URL> --address=<public key of owner>
```

## Checking Balance

```
plasma_cli get balance --watcher=<watcher URL> --address=<public key of owner>
```

## Creating simple transaction

NOTE: This function will either 
1. Send the entire UTXO to an address, or
2. Split the UTXO to destination address and send the change to the owner

Support for more complex transactions will come later

```
plasma_cli transaction --fromutxo=<UTXO position> --fromowner=<from address> --privatekey=<from privatekey> --toowner=<to address> --toamount=<to amount>  --watcher=<watcher url>
```

## Exit to Rootchain

```
plasma_cli exit --utxo=<utxo ID> --privatekey=<private key> --contract=<contract address> --watcher=<watcher url> --client=<local rpc server or Infura URL>
```