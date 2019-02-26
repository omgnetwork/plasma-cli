# OmiseGO Plasma CLI

## What is this?

The Plasma CLI is used for interaction with OmiseGO services from the command line. It allows the user to deposit to the Plasma contract, create transactions, and exit back to the root chain with the Ethereum & MoreVP Plasma details abstracted from the user. Both local RPC node for connecting to Ethereum and Infura are supported.

## How do I deposit into the contract?

```
plasma_cli deposit --privatekey=<private key> --client=<local rpc server or Infura URL> --contract=<contract address> --amount=<amount in wei> --owner=<public key of the owner>
```


## How do get the UTXOs that I own?

```
plasma_cli.go get utxos --watcher=<watcher URL> --address=<public key of owner>
```

## How do I create a transaction?

```
plasma_cli transaction --fromutxo=<UTXO position> --fromowner=<from address> --privatekey=<from privatekey> --toowner=<to address> --toamount=<to amount>  --watcher=<watcher url>
```

## How do I exit back to Ethereum?

```
plasma_cli exit --utxo=<utxo ID> --privatekey=<private key> --contract=<contract address> --watcher=<watcher url> --client=<local rpc server or Infura URL>
```