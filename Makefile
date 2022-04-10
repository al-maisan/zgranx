.PHONY: all build

BIN_DIR := ./bin
version := $(shell git rev-parse --short=12 HEAD)
timestamp := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

all: build

build: proto
	rm -f $(BIN_DIR)/exaftx*
	go build -o $(BIN_DIR)/exaftx_server -v -ldflags \
  "-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/exaftx/server/main.go
	go build -o $(BIN_DIR)/exaftx_client -v -ldflags \
  "-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/exaftx/client/main.go


proto:
	protoc --go_out=. --go-grpc_out=. api/monitor.proto
