# elysium-app

[![Go Reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/elysiumorg/elysium-app)
[![GitHub Release](https://img.shields.io/github/v/release/elysiumorg/elysium-app)](https://github.com/elysiumorg/elysium-app/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/elysiumorg/elysium-app)](https://goreportcard.com/report/github.com/elysiumorg/elysium-app)
[![Lint](https://github.com/elysiumorg/elysium-app/actions/workflows/lint.yml/badge.svg)](https://github.com/elysiumorg/elysium-app/actions/workflows/lint.yml)
[![Tests / Code Coverage](https://github.com/elysiumorg/elysium-app/actions/workflows/test.yml/badge.svg)](https://github.com/elysiumorg/elysium-app/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/elysiumorg/elysium-app/branch/main/graph/badge.svg?token=CWGA4RLDS9)](https://codecov.io/gh/elysiumorg/elysium-app)
[![GitPOAP Badge](https://public-api.gitpoap.io/v1/repo/elysiumorg/elysium-app/badge)](https://www.gitpoap.io/gh/elysiumorg/elysium-app)

**elysium-app** is a blockchain application built using Cosmos SDK and [elysium-core](https://github.com/elysiumorg/elysium-core) in place of Tendermint.

## Diagram

```ascii
                ^  +-------------------------------+  ^
                |  |                               |  |
                |  |  State-machine = Application  |  |
                |  |                               |  |   elysium-app (built with Cosmos SDK)
                |  |            ^      +           |  |
                |  +----------- | ABCI | ----------+  v
Elysium        |  |            +      v           |  ^
validator or    |  |                               |  |
full consensus  |  |           Consensus           |  |
node            |  |                               |  |
                |  +-------------------------------+  |   elysium-core (fork of Tendermint Core)
                |  |                               |  |
                |  |           Networking          |  |
                |  |                               |  |
                v  +-------------------------------+  v
```

## Install

1. [Install Go](https://go.dev/doc/install) 1.18+
1. Clone this repo
1. Install the elysium-app CLI

    ```shell
    make install
    ```

## Usage

```sh
# Print help
elysium-appd --help
```

### Create your own single node devnet

```sh
# WARNING: this deletes config, data, and keyrings from previous local devnets
rm -r ~/.elysium-app

# Start a single node devnet
./scripts/single-node.sh

# Post data to the local devnet
elysium-appd tx blob PayForBlobs [hexNamespace] [hexBlob] [flags]
```

<!-- markdown-link-check-disable -->
<!-- markdown-link encounters an HTTP 503 on this link even though it works. -->
<!-- See https://github.com/elysiumorg/elysium-app/actions/runs/3296219513/jobs/5439416229#step:4:185 -->
See <https://docs.elysium.org/category/elysium-app> for more information
<!-- markdown-link-check-enable -->

## Contributing

### Tools

1. Install [golangci-lint](https://golangci-lint.run/usage/install/)
1. Install [markdownlint](https://github.com/DavidAnson/markdownlint)

### Helpful Commands

```sh
# Build a new elysium-app binary and output to build/elysium-appd
make build

# Run tests
make test

# Format code with linters (this assumes golangci-lint and markdownlint are installed)
make fmt

# Regenerate Protobuf files (this assumes Docker is running)
make proto-gen
```

### Package-specific documentation

- [Shares](https://pkg.go.dev/github.com/elysiumorg/elysium-app/pkg/shares)

## Careers

We are hiring Go engineers! Join us in building the future of blockchain scaling and interoperability. [Apply here](https://jobs.lever.co/elysium).
