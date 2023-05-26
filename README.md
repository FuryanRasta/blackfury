<!--
parent:
  order: false
-->

<div align="center">
  <img src="https://user-images.githubusercontent.com/102999403/161656793-7a826432-de47-46ea-b212-72907f462b7d.gif" />
  <h1> Blackfury </h1>
</div>

<!-- TODO: add banner -->
<!-- ![banner](docs/ethermint.jpg) -->

<div align="center">
  <a href="https://github.com/furyanrasta/blackfury/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/furyanrasta/blackfury.svg" />
  </a>
  <a href="https://github.com/furyanrasta/blackfury/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/furyanrasta/blackfury.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/furyanrasta/blackfury">
    <img alt="GoDoc" src="https://godoc.org/github.com/furyanrasta/blackfury?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/furyanrasta/blackfury">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/furyanrasta/blackfury"/>
  </a>
</div>
<div align="center">
  <a href="https://discord.gg/ArXNfK99ae">
    <img alt="Discord" src="https://img.shields.io/discord/962917488180490250.svg" />
  </a>
  <a href="https://github.com/furyanrasta/blackfury/actions?query=branch%3Amain+workflow%3ALint">
    <img alt="Lint Status" src="https://github.com/furyanrasta/blackfury/actions/workflows/lint.yml/badge.svg?branch=main" />
  </a>
  <a href="https://codecov.io/gh/furyanrasta/blackfury">
    <img alt="Code Coverage" src="https://codecov.io/gh/furyanrasta/blackfury/branch/main/graph/badge.svg" />
  </a>
  <a href="https://twitter.com/BlackfuryFDN">
    <img alt="Twitter Follow Blackfury" src="https://img.shields.io/twitter/follow/BlackfuryFDN"/>
  </a>
</div>

Blackfury (ECH) is a scalable, high-throughput Proof-of-Stake blockchain that is fully compatible and
interoperable with Ethereum and Cosmos. It's built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/) which runs on top of [Tendermint Core](https://github.com/tendermint/tendermint) consensus engine with [EVMOS Ethermint](https://github.com/evmos/ethermint).

**Note**: Requires [Go 1.18+](https://golang.org/dl/)

## Installation

For prerequisites and detailed build instructions please read the [Installation](https://docs.ech.network) instructions. Once the dependencies are installed, run:

```bash
make install
```

Or check out the latest [release](https://github.com/furyanrasta/blackfury/releases).

## Genesis
To get onto our mainnet (highbury_710-1) genesis download the genesis.json here

`wget https://gist.githubusercontent.com/furyanrasta/ee862f58850fc1b5ee6a6fdccc3130d2/raw/55c2c4ea2fee8a9391d0dc55b2c272adb804054a/genesis.json`

and then move it into the blackfuryd config (after you have initilized your node)

`blackfuryd init <nodename> --chain-id highbury_710-1`

`mv genesis.json ~/.blackfuryd/config/`

## Quick Start

To learn how the Blackfury works from a high-level perspective, go to the [Introduction](https://docs.ech.network) section from the documentation. You can also check the instructions to [Run a Node](https://docs.ech.network). You can also read the Cosmos SDK and familiarize yourself with Cosmos SDK, Tendermint, and Ethermint.

## Community

The following chat channels and forums are a great spot to ask questions about Blackfury:

- [Blackfury Twitter](https://twitter.com/BlackfuryFDN)
- [Blackfury Discord](https://discord.gg/ArXNfK99ae)
- [Blackfury Telegram](https://t.me/blackfuryANN)
- [Official Website](https://ech.network)
- [Official Dapp](https://app.ech.network)

## Contributing

Looking for a good place to start contributing? Check out some [`good first issues`](https://github.com/furyanrasta/blackfury/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22).

Original codebase forked from [EVMOS](https://github.com/evmos/evmos)

For additional instructions, standards and style guides, please refer to the [Contributing](./CONTRIBUTING.md) document.
