# Relay Network
## Motivation
Various projects aim to improve user experience by enabling recurring fee subscriptions, delegated conditional execution or sponsored gas. These second layer execution markets all share the need for a relayer that submits signed user transactions to the network.

The aim of this project is to implement a decentralized relay network, with enough modularity that could support the various use-cases mentioned above. A [prototype](https://github.com/hermes-network/the-executor) of this idea, with a focus on supporting [Gnosis Safe](https://gnosis-safe.readthedocs.io/en/latest/) as identity proxy and whisper as pubsub protocol, was recently implemented by the Hermes Network team at ETHBerlin.
