# Development Environment Setup

## Requirements
 * [Go](https://go.dev/doc/install) 1.17+ (to build the provider)
 * [pre-commit](https://pre-commit.com/) (to run pre-commit hooks)

## Quick Start

### Building

Clone the repository:

```bash
git clone https://github.com/andrewthetechie/cq-provider-datadog
```

Build the provider:

```
make build
```

### Running the provider in debug mode

1. Download [CloudQuery](https://github.com/cloudquery/cloudquery) latest version.
1. Execute `make run` and note of the `CQ_REATTACH_PROVIDERS` value.
1. Open another terminal and run `CQ_REATTACH_PROVIDERS=[VALUE_FROM_PREV] ./cloudquery fetch` 

> Make sure the authentication variables are exported in the provider process and not in cloudquery process.

See [docs](https://docs.cloudquery.io/docs/developers/debugging) for more details.

### Pre-commit

This repo uses pre-commit to run some tests and utilities before every commit to help you pass CI easier! Make sure you have [pre-commit](https://pre-commit.com/) installed and then run

```bash
pre-commit install
```

to install pre-commit's hooks.

You can run all pre-commit checks at any time with

```bash
pre-commit run -a
```

### Testing

The provider currently has very little testing. Please, add tests if you can!

#### Unit Tests

Unit Tests don't require any credentials or internet access

```bash
make test-unit # This runs go test ./...
```

Unit tests include:
- Specific resources tests. You can find those next to each resource, in the [`resources/services`](../../resources/services) folder.
- DB migration tests. You can find the code for these tests [here](../../resources/provider/provider_test.go).
- Client tests. You can find those in the [`client`](../../client) folder.

