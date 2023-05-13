# elysium-app

[![Go Reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/furyaxyz/elysium-app)
[![mdBook Specs](https://img.shields.io/badge/mdBook-specs-blue)](https://furyaxyz.github.io/elysium-app/)
[![GitHub Release](https://img.shields.io/github/v/release/furyaxyz/elysium-app)](https://github.com/furyaxyz/elysium-app/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/furyaxyz/elysium-app)](https://goreportcard.com/report/github.com/furyaxyz/elysium-app)
[![Lint](https://github.com/furyaxyz/elysium-app/actions/workflows/lint.yml/badge.svg)](https://github.com/furyaxyz/elysium-app/actions/workflows/lint.yml)
[![Tests / Code Coverage](https://github.com/furyaxyz/elysium-app/actions/workflows/test.yml/badge.svg)](https://github.com/furyaxyz/elysium-app/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/furyaxyz/elysium-app/branch/main/graph/badge.svg?token=CWGA4RLDS9)](https://app.codecov.io/gh/furyaxyz/elysium-app/tree/main)
[![GitPOAP Badge](https://public-api.gitpoap.io/v1/repo/furyaxyz/elysium-app/badge)](https://www.gitpoap.io/gh/furyaxyz/elysium-app)

elysium-app is a blockchain application built using parts of the Cosmos stack. elysium-app uses

- [furyaxyz/cosmos-sdk](https://github.com/furyaxyz/cosmos-sdk) a fork of [cosmos/cosmos-sdk](https://github.com/cosmos/cosmos-sdk)
- [furyaxyz/elysium-core](https://github.com/furyaxyz/elysium-core) a fork of [cometbft/cometbft](https://github.com/cometbft/cometbft)

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
                |  +-------------------------------+  |   elysium-core (fork of CometBFT)
                |  |                               |  |
                |  |           Networking          |  |
                |  |                               |  |
                v  +-------------------------------+  v
```

## Install

1. [Install Go](https://go.dev/doc/install) 1.20
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

### Environment variables

| Variable        | Explanation                        | Default value                                            | Required |
| --------------- | ---------------------------------- | -------------------------------------------------------- | -------- |
| `CELESITA_HOME` | Home directory for the application | User home dir. [Ref](https://pkg.go.dev/os#UserHomeDir). | Optional |

### Create your own single node devnet

```sh

# Start a single node devnet using the pre-installed elysium app
./scripts/single-node.sh

# Build and start a single node devnet
./scripts/build-run-single-node.sh

# Post data to the local devnet
elysium-appd tx blob PayForBlobs [hexNamespace] [hexBlob] [flags]
```

**Note:** please note that the `./scripts/` commands above, created a random `tmp` directory and keeps all data/configs there.

<!-- markdown-link-check-disable -->
<!-- markdown-link encounters an HTTP 503 on this link even though it works. -->
<!-- See https://github.com/furyaxyz/elysium-app/actions/runs/3296219513/jobs/5439416229#step:4:185 -->
See <https://docs.elysium.org/category/elysium-app> for more information
<!-- markdown-link-check-enable -->

## Contributing

### Tools

1. Install [golangci-lint](https://golangci-lint.run/usage/install/)
1. Install [markdownlint](https://github.com/DavidAnson/markdownlint)
1. Install [hadolint](https://github.com/hadolint/hadolint)
1. Install [yamllint](https://yamllint.readthedocs.io/en/stable/quickstart.html)

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

### Docs

Package-specific docs aim to explain implementation details for developers that are contributing to these packages. The [specs](https://furyaxyz.github.io/elysium-app/) aim to explain the protocol as a whole for developers building on top of Elysium.

- [pkg/shares](https://pkg.go.dev/github.com/furyaxyz/elysium-app/pkg/shares)
- [pkg/wrapper](https://github.com/furyaxyz/elysium-app/blob/main/pkg/wrapper)
- [x/blob](https://github.com/furyaxyz/elysium-app/tree/main/x/blob)
- [x/qgb](https://github.com/furyaxyz/elysium-app/tree/main/x/qgb)

## Careers

We are hiring Go engineers! Join us in building the future of blockchain scaling and interoperability. [Apply here](https://jobs.lever.co/elysium).
