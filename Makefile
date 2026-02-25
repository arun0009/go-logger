.DEFAULT_GOAL := help
SHELL := /bin/bash

# Project variables
BINARY_NAME=go-logger
PKG=./pkg/logger/...

.PHONY: help build test bench vet format tidy clean

help: ## Show this help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-z0-9A-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the library
	go build ./...

test: ## Run all unit tests
	go test -v $(PKG)

bench: ## Run benchmarks
	go test -bench=. -benchmem $(PKG)

vet: ## Run go vet
	go vet $(PKG)

tidy: ## Tidy up go.mod and go.sum
	go mod tidy

format: ## Format source code
	go fmt $(PKG)

clean: ## Remove build artifacts and test cache
	go clean -testcache
	rm -rf target/
	rm -f $(BINARY_NAME)
	rm -f *.out
	rm -f *.test