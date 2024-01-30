PROJECT_PATH := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
PROJECT_NAME = ch-app-store
MODULE_NAME := $(shell go list -m)
CMD_PATH= $(PROJECT_PATH)/cmd
BIN_PATH = $(PROJECT_PATH)/bin
GEN_PATH = $(PROJECT_PATH)/generated
TEST_ENV = "test"

# Application environment
STAGE ?= development
VERSION := $(shell git describe --exact-match --tags HEAD 2>/dev/null || git rev-parse --abbrev-ref HEAD)

#Flyway
FLYWAY_CONFIG_PATH = $(PROJECT_PATH)/config/flyway/$(STAGE).conf
FLYWAY_SQL_PATH = $(PROJECT_PATH)/resources/psql/migration

# Sql boiler
SQLBOILER_CONFIG_PATH = $(PROJECT_PATH)/config/sqlboiler/$(STAGE).toml

# Mockery
MOCKERY_TARGET_PATH = $(PROJECT_PATH)/internal
MOCKERY_OUTPUT_PATH = $(PROJECT_PATH)/test/mock

# Go environment
GOVERSION := $(shell go version | awk '{print $$3}')
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOPRIVATE ?= github.com/channel-io
LDFLAGS := -ldflags="-X ${MODULE_NAME}/tool.buildVersion=${VERSION}"

env:
	@echo "PROJECT_PATH:\t${PROJECT_PATH}"
	@echo "PROJECT_NAME:\t${PROJECT_NAME}"
	@echo "MODULE_NAME:\t${MODULE_NAME}"
	@echo "GOVERSION:\t${GOVERSION}"
	@echo "GOOS:\t\t${GOOS}"
	@echo "GOARCH:\t\t${GOARCH}"
	@echo "STAGE:\t\t${STAGE}"
	@echo "VERSION:\t${VERSION}"

installTools:
	# SqlBoiler
	go install github.com/volatiletech/sqlboiler/v4@v4.14.1
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.14.1
	# Mockery
	go install github.com/vektra/mockery/v2@v2.26.1
	# Swagger
	go install github.com/swaggo/swag/cmd/swag@v1.16.2

generate: flyway_clean flyway_migrate genBoiler genMock

genBoiler:
	sqlboiler --wipe --no-tests -o generated/models -c $(SQLBOILER_CONFIG_PATH) psql

genMock:
	#mockery --all --dir=$(MOCKERY_TARGET_PATH) --output=$(MOCKERY_OUTPUT_PATH) --keeptree --with-expecter --inpackage=false --packageprefix='mock'

init: installTools
	GOPRIVATE=${GOPRIVATE} go mod download

build: init generate
	GOOS=${GOOS} \
	GOARCH=${GOARCH} \
	go build ${LDFLAGS} \
	-o ${BIN_PATH}/${PROJECT_NAME}.${GOOS}.${GOARCH} \
	${CMD_PATH}

run:
	${BIN_PATH}/${PROJECT_NAME}.${GOOS}.${GOARCH}

dev: build run

test:
	GO_ENV=$(TEST_ENV) go test $(PROJECT_PATH)

clean: clean-bin clean-gen

clean-gen:
	rm -rf $(GEN_PATH)

clean-bin:
	rm -rf $(BIN_PATH)

docs: docs-gen docs-fmt

docs-gen:
	swag init -d api/http -g module.go -o api/http/swagger --pd

docs-fmt:
	swag fmt -d api/http

flyway_clean:
	flyway -configFiles=$(FLYWAY_CONFIG_PATH) -locations=filesystem:$(FLYWAY_SQL_PATH) clean

flyway_migrate:
	flyway -configFiles=$(FLYWAY_CONFIG_PATH) -locations=filesystem:$(FLYWAY_SQL_PATH) migrate

flyway_info:
	flyway -configFiles=$(FLYWAY_CONFIG_PATH) -locations=filesystem:$(FLYWAY_SQL_PATH) info

flyway_validate:
	flyway -configFiles=$(FLYWAY_CONFIG_PATH) -locations=filesystem:$(FLYWAY_SQL_PATH) validate

flyway_repair:
	flyway -configFiles=$(FLYWAY_CONFIG_PATH) -locations=filesystem:$(FLYWAY_SQL_PATH) repair

