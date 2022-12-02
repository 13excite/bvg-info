SHELL := /bin/bash

# constant variables
GIT_COMMIT 		= $(shell git rev-parse HEAD)
BINARY_TAR_DIR 	= $(BINARY_NAME)-$(GIT_COMMIT)
BINARY_TAR_FILE	= $(BINARY_TAR_DIR).tar.gz
BUILD_VERSION 	= $(shell cat VERSION.txt)
BUILD_DATE 		= $(shell date -u '+%Y-%m-%d_%H:%M:%S')

# LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: lint fmt

fmt:
	@gofmt -l -w $(SRC)

lint:
	@echo 'running linter...'
	@golangci-lint run ./...

