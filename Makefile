.PHONY: all build proto

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
	$(shell find internal/ | grep pb.go$ | xargs rm -f)
	protoc --go_out=. --go-grpc_out=. api/monitor.proto
	protoc --experimental_allow_proto3_optional --go_out=. --go-grpc_out=. api/exa.proto
