# Getting Started

## Installation

### Building from source

Prerequistes: The [Go](https://golang.org/doc/install) compiler and GNU make

Run `make install` to build the code.`xrnd` is the blockchain daemon. `xrncli`
is the command line blockchain client. Both of them provide help messages when run.

## Configuring the command line client

By default the command line client will connect to a [local node](#running-a-local-node).
To connect to a testnet, you must know the testnet's chain ID and the address of the node.

`xrncli` can be configured to connect to a testnet node automatically by setting these
parameters in `$HOME/.xrncli/config/config.toml`. This file can be generated
automatically by running:

```sh
xrncli init --chain-id [chain ID] --node [node address]
```

Check current testnet status [here](../testnets).

Run `xrncli status` to see if you are able to connect to the block-chain node.

### Creating a key

To run any transactions you must have a key. A new one can be generated using:

```sh
xrncli keys add [key-name]
```

Make sure you save the seed mnemonic in a safe place!

## Running a local node

Initialize the node config

```sh
xrnd init
```

Start the node
```sh
xrnd init
```
