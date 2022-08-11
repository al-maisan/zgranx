.PHONY: all build proto

BIN_DIR := ./bin
version := $(shell git rev-parse --short=12 HEAD)
timestamp := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

sdbdir=/tmp/mariadb/T-801/sdb
sdbinitdir="$(ROOT_DIR)/deployments/sdb"
pdbdir=/tmp/mariadb/T-801/pdb
pdbinitdir="$(ROOT_DIR)/deployments/pdb"


all: build


build: proto
	rm -f $(BIN_DIR)/dit
	go build -o $(BIN_DIR)/dit -v -ldflags \
  "-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/dit/main.go
	rm -f $(BIN_DIR)/etc
	go build -o $(BIN_DIR)/etc -v -ldflags \
  "-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/etc/main.go
	rm -f $(BIN_DIR)/huobi
	go build -o $(BIN_DIR)/huobi -v -ldflags \
  "-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/exa/huobi/main.go
	rm -f $(BIN_DIR)/ma
	go build -o $(BIN_DIR)/ma -v -ldflags \
  "-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/ma/main.go
	rm -f $(BIN_DIR)/tma
	go build -o $(BIN_DIR)/tma -v -ldflags \
  "-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/ma/main.go


proto:
	$(shell find internal/ | grep pb.go$ | xargs rm -f)
	protoc --go_out=. --go-grpc_out=. api/monitor.proto
	protoc --experimental_allow_proto3_optional --go_out=. --go-grpc_out=. api/exa.proto
	protoc --experimental_allow_proto3_optional --go_out=. --go-grpc_out=. api/ma.proto
	protoc --experimental_allow_proto3_optional --go_opt=module=github.com/alphabot-fi/T-801 --go_out=. --go-grpc_out=. api/base.proto


dockerinit:
	-docker container prune -f >/dev/null 2>&1
	-docker network create T-801net >/dev/null 2>&1


sdbinit: sdbhalt
	-docker container prune -f >/dev/null 2>&1
	-sudo rm -rf $(sdbdir)
	-docker run -p 3306:3306 --detach -v $(sdbdir):/var/lib/mysql:z  -v $(sdbinitdir):/docker-entrypoint-initdb.d:z --network T-801net --name sT-801db --env MARIADB_USER=$(T_801_SDB_USER) --env MARIADB_PASSWORD=$(T_801_SDB_PASSWORD) --env MARIADB_ROOT_PASSWORD=$(T_801_SDB_ROOT_PASSWORD) --env MARIADB_DATABASE=$(T_801_SDB_DATABASE) mariadb:latest


sdbstart:
	-docker container prune -f >/dev/null 2>&1
	-docker run -p 3306:3306 --detach -v $(sdbdir):/var/lib/mysql  --network T-801net --name sT-801db --env MARIADB_USER=$(T_801_SDB_USER) --env MARIADB_PASSWORD=$(T_801_SDB_PASSWORD) --env MARIADB_ROOT_PASSWORD=$(T_801_SDB_ROOT_PASSWORD) mariadb:latest


sdbhalt:
	-docker stop sT-801db
	-docker container prune -f >/dev/null 2>&1

sdbprompt:
	-docker container prune -f >/dev/null 2>&1
	-docker run --network T-801net -it --rm mariadb mysql -h sT-801db -u $(T_801_SDB_USER) -D $(T_801_SDB_DATABASE) -p$(T_801_SDB_PASSWORD)


pdbinit: pdbhalt
	-docker container prune -f >/dev/null 2>&1
	-sudo rm -rf $(pdbdir)
	-docker run -p 3307:3306 --detach -v $(pdbdir):/var/lib/mysql:z  -v $(pdbinitdir):/docker-entrypoint-initdb.d:z --network T-801net --name pT-801db --env MARIADB_USER=$(T_801_PDB_USER) --env MARIADB_PASSWORD=$(T_801_PDB_PASSWORD) --env MARIADB_ROOT_PASSWORD=$(T_801_PDB_ROOT_PASSWORD) --env MARIADB_DATABASE=$(T_801_PDB_DATABASE) mariadb:latest
	-sleep 4
	-docker exec -i pT-801db mysql -u root -p$(T_801_PDB_ROOT_PASSWORD) -D$(T_801_PDB_DATABASE) -e "SET GLOBAL max_allowed_packet=1072731894;"


pdbstart:
	-docker container prune -f >/dev/null 2>&1
	-docker run -p 3307:3306 --detach -v $(pdbdir):/var/lib/mysql  --network T-801net --name pT-801db --env MARIADB_USER=$(T_801_PDB_USER) --env MARIADB_PASSWORD=$(T_801_PDB_PASSWORD) --env MARIADB_ROOT_PASSWORD=$(T_801_PDB_ROOT_PASSWORD) mariadb:latest


pdbhalt:
	-docker stop pT-801db
	-docker container prune -f >/dev/null 2>&1

pdbprompt:
	-docker container prune -f >/dev/null 2>&1
	-docker run --network T-801net -it --rm mariadb mysql -h pT-801db -u $(T_801_PDB_USER) -D $(T_801_PDB_DATABASE) -p$(T_801_PDB_PASSWORD)
