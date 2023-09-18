SHELL:=/usr/bin/env bash

BIN_NAME:=dirshard
BIN_VERSION:=$(shell ./.version.sh)

default: help
# via https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## Print help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all
all: clean build build-linux-amd64 build-linux-arm64 build-linux-armv7 build-linux-armv6 build-darwin-amd64 build-darwin-arm64 ## Build for macOS (amd64, arm64) and Linux (amd64, arm64, armv7, armv6)

.PHONY: clean
clean: ## Remove build products (./out)
	rm -rf ./out

.PHONY: build
build: ## Build for the current platform & architecture to ./out
	mkdir -p out
	go build -ldflags="-X main.version=${BIN_VERSION}" -o ./out/${BIN_NAME} .

.PHONY: build-linux-amd64
build-linux-amd64: ## Build for Linux/amd64 to ./out
	env GOOS=linux GOARCH=amd64 go build -ldflags="-X main.version=${BIN_VERSION}" -o ./out/${BIN_NAME}-linux-amd64 .

.PHONY: build-linux-arm64
build-linux-arm64: ## Build for Linux/arm64 to ./out
	env GOOS=linux GOARCH=arm64 go build -ldflags="-X main.version=${BIN_VERSION}" -o ./out/${BIN_NAME}-linux-arm64 .

.PHONY: build-linux-armv7
build-linux-armv7: ## Build for Linux/armv7 to ./out
	env GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-X main.version=${BIN_VERSION}" -o ./out/${BIN_NAME}-linux-armv7 .

.PHONY: build-linux-armv6
build-linux-armv6: ## Build for Linux/armv6 to ./out
	env GOOS=linux GOARCH=arm GOARM=6 go build -ldflags="-X main.version=${BIN_VERSION}" -o ./out/${BIN_NAME}-linux-armv6 .

.PHONY: build-darwin-amd64
build-darwin-amd64: ## Build for macOS/amd64 to ./out
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-X main.version=${BIN_VERSION}" -o ./out/${BIN_NAME}-darwin-amd64 .

.PHONY: build-darwin-arm64
build-darwin-arm64: ## Build for macOS/arm64 to ./out
	env GOOS=darwin GOARCH=arm64 go build -ldflags="-X main.version=${BIN_VERSION}" -o ./out/${BIN_NAME}-darwin-arm64 .

.PHONY: lint
lint: ## Lint all source files in this repository (requires nektos/act: https://nektosact.com)
	act --artifact-server-path /tmp/artifacts -j lint

.PHONY: update-lint
update-lint: ## Pull updated images supporting the lint target (may fetch >10 GB!)
	docker pull catthehacker/ubuntu:full-latest

GOLINT_FILES:=$(shell find . -name '*.go' | grep -v /vendor/)
.PHONY: golint
golint: ## Lint all .go files with golint
	@for file in ${GOLINT_FILES} ;  do \
		echo "$$file" ; \
		golint $$file ; \
	done
