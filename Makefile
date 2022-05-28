# DONT EDIT. This file is synced from https://github.com/cloudquery/.github/misc/Makefile

.DEFAULT_GOAL := help

# This help function will automatically generate help/usage text for any make target that is commented with "##".
# Targets with a singe "#" description do not show up in the help text
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}'


.PHONY: install-cq 
install-cq:  ## install the latest version of CQ
	@if [[ "$(OS)" != "Darwin" && "$(OS)" != "Linux" && "$(OS)" != "Windows" ]]; then echo "\n Invalid OS set. Valid Options are Darwin, Linux and Windows. Example invocation is:\n make OS=Linux ARCH=arm64 install-cq \n For more information go to  https://docs.cloudquery.io/docs/getting-started \n"; exit 1; fi
	@if [[ "$(ARCH)" != "x86_64" && "$(ARCH)" != "arm64" ]]; then echo "\n Invalid ARCH set. Valid options are x86_64 and arm64. Example invocation is:\n make OS=Linux ARCH=arm64 install-cq \n For more information go to  https://docs.cloudquery.io/docs/getting-started \n"; exit 1; fi
	curl -L https://github.com/cloudquery/cloudquery/releases/latest/download/cloudquery_${OS}_${ARCH} -o cloudquery
	chmod a+x cloudquery

# start a timescale db running in a local container
.PHONY: ts-start
ts-start:
	docker run -p 5433:5432 -e POSTGRES_PASSWORD=pass -d timescale/timescaledb:latest-pg14

# stop the timescale db running in a local container
.PHONY: ts-stop
ts-stop:
	docker stop $$(docker ps -q --filter ancestor=timescale/timescaledb:latest-pg14)

# start a running docker container
.PHONY: pg-start
pg-start:
	docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d postgres

# stop a running docker container
.PHONY: pg-stop
pg-stop:
	docker stop $$(docker ps -q --filter ancestor=postgres:latest)

# connect to pg via cli
.PHONY: pg-connect
pg-connect:
	psql -h localhost -p 5432 -U postgres -d postgres

.PHONY: build
build:  ## build the cq provider
	go build -o cq-provider

.PHONY: run
run: build  ## build and run the cq provider
	CQ_PROVIDER_DEBUG=1 CQ_REATTACH_PROVIDERS=.cq_reattach ./cq-provider


.PHONY: fetch
fetch:  ## Run a fetch command
	CQ_PROVIDER_DEBUG=1 CQ_REATTACH_PROVIDERS=.cq_reattach ./cloudquery fetch --dsn "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable" -v --fail-on-error

.PHONY: generate-mocks
generate-mocks: ## Generate mocks for mock/unit testing 
	go generate ./client/services/...

.PHONY: test-unit
test-unit: ## Run unit tests
	go test -timeout 3m ./...

.PHONY: test-integration
test-integration: ## Run integration tests
	@if [[ "$(tableName)" == "" ]]; then go test -run=TestIntegration -timeout 3m -tags=integration ./...; else go test -run="TestIntegration/$(tableName)" -timeout 3m -tags=integration ./...; fi
