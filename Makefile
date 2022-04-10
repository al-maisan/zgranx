.PHONY: all build

BIN_DIR := ./bin
version := $(shell git rev-parse --short=12 HEAD)
timestamp := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

all: build

build:
	rm -f $(BIN_DIR)/exaftx
	go build -o $(BIN_DIR)/exaftx -v -ldflags \
  "-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/exaftx/main.go
