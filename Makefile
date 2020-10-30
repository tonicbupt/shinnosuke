.PHONY: shinnosuke

REVISION := $(shell git rev-parse HEAD || unknown)
BUILTAT := $(shell date +%Y-%m-%dT%H:%M:%S)
VERSION := $(shell git describe --tags $(shell git rev-list --tags --max-count=1) || git rev-parse --short HEAD)
GO_LDFLAGS ?= -s -X github.com/tonicbupt/shinnosuke/pkg/version.REVISION=$(REVISION) \
                 -X github.com/tonicbupt/shinnosuke/pkg/version.BUILTAT=$(BUILTAT) \
                 -X github.com/tonicbupt/shinnosuke/pkg/version.VERSION=$(VERSION)

all: shinnosuke

shinnosuke:
	mkdir -p bin
	go build -ldflags "$(GO_LDFLAGS)" -o bin/shinnosuke ./cmd/shinnosuke
test:
	go test -v -cover -count=1 -p=1 ./...
