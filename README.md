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

Furthermore, you'll need a local Ethereum chain running. An easy way to do that is by using [ganache-cli](https://github.com/trufflesuite/ganache-cli).

Clone the repository and install the dependencies via [dep](https://golang.github.io/dep/):

```bash
dep ensure
```

Due to an [issue]() in dep, `libsecp256k1` needs to be copied from `go-ethereum` source to vendored version:

```bash
cp -r \
  "${GOPATH}/src/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1" \
  "vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/"
```

## Usage
Run a relayer, which starts listening to a pubsub topic:

```bash
go run cmd/relayer/main.go
```

Send a test message by running:

```bash
go run cmd/publisher/main.go
```
