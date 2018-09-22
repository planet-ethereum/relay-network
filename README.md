# Relay Network
[![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/planet-ethereum/Lobby)

## Motivation
Various projects aim to improve user experience by enabling recurring fee subscriptions, delegated conditional execution or sponsored gas. These second layer execution markets all share the need for a relayer that submits signed user transactions to the network.

The aim of this project is to implement a decentralized relay network, with enough modularity that could support the various use-cases mentioned above. A [prototype](https://github.com/hermes-network/the-executor) of this idea, with a focus on supporting [Gnosis Safe](https://gnosis-safe.readthedocs.io/en/latest/) as identity proxy and whisper as pubsub protocol, was recently implemented by the Hermes Network team at ETHBerlin.

## Requirements
In order to run a relayer, you'll need to have an IPFS node with pubsub enabled running. Install and setup IPFS, then run daemon:

```bash
ipfs daemon --enable-pubsub-experiment
```

Relayer also needs access to an Ethereum chain. By default, it connects to Rinkeby, through Infura.

Clone the repository and install the dependencies via [dep](https://golang.github.io/dep/):

```bash
dep ensure
```

## Usage
Relay-network currently only works with [ethbase](https://github.com/planet-ethereum/ethbase). For using with ethbase, start a relayer, which listens on a pubsub topic, and submits the logs to the ethbase contract:

```bash
MNEMONIC="____" go run cmd/relayer/main.go
```

Then, run `ethbasepub`, which subscribes to the logs emitted by event emitters registered in `Ethbase`, generates a proof for the log, and sends it to the `relayer` via IPFS pubsub.

```bash
go run cmd/ethbasepub/main.go
```
