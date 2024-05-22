# Project directory structure
MODULE_NAME := $(shell go list -m)
PROJECT_PATH := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
PROJECT_NAME = ch-app-store
APPS := admin general
DOC_TARGETS := admin desk front general account

export PATH := ${PATH}:${GOPATH}/bin

# Artifacts
TARGET_DIR ?= ${PROJECT_PATH}/target
TARGET_BIN_DIR ?= ${TARGET_DIR}/bin
GENERATED_SRC_DIR ?= ${PROJECT_PATH}/generated

# Application environment
STAGE ?= development
VERSION := $(shell git describe --exact-match --tags HEAD 2>/dev/null || git rev-parse --abbrev-ref HEAD)

# Go environment
GOVERSION := $(shell go version | awk '{print $$3}')
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOPRIVATE ?= github.com/channel-io
LDFLAGS := -ldflags="-X ${MODULE_NAME}/tool.buildVersion=${VERSION}"

# Default setting for flyway migration (이 부분은 이야기 해봐야함)
PSQL_HOST ?= localhost
PSQL_DBNAME ?= app_store
PSQL_SCHEMA ?= public
PSQL_USER ?= app_store
PSQL_PASSWORD ?= ""

TEST_ENV = "test"

#----------------------------------- Delete 예정
#Flyway
FLYWAY_CONFIG_PATH = $(PROJECT_PATH)/config/flyway/$(STAGE).conf
FLYWAY_SQL_PATH = $(PROJECT_PATH)/resources/psql/migration

# Sql boiler
#SQLBOILER_CONFIG_PATH = $(PROJECT_PATH)/config/sqlboiler/$(STAGE).toml

# Mockery
MOCKERY_TARGET_PATH = $(PROJECT_PATH)/internal
MOCKERY_OUTPUT_PATH = $(PROJECT_PATH)/generated/mock
#----------------------------------- Delete 예정

env:
	@echo "PROJECT_PATH:\t${PROJECT_PATH}"
	@echo "PROJECT_NAME:\t${PROJECT_NAME}"
	@echo "MODULE_NAME:\t${MODULE_NAME}"
	@echo "GOVERSION:\t${GOVERSION}"
	@echo "GOOS:\t\t${GOOS}"
	@echo "GOARCH:\t\t${GOARCH}"
	@echo "STAGE:\t\t${STAGE}"
	@echo "VERSION:\t${VERSION}"

submodule-init:
	git submodule update --init --recursive

install-tools:
	# SqlBoiler
	go install github.com/volatiletech/sqlboiler/v4@v4.14.1
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.14.1
	# Mockery
	go install github.com/vektra/mockery/v2@v2.26.1
	# Swagger
	go install github.com/swaggo/swag/cmd/swag@v1.16.2

init: submodule-init install-tools
	GOPRIVATE=${GOPRIVATE} go mod download

generate: flyway-migrate gen-boiler gen-mock

gen-boiler:
	PSQL_DBNAME=${PSQL_DBNAME} \
	PSQL_SCHEMA=${PSQL_SCHEMA} \
	PSQL_HOST=${PSQL_HOST} \
	PSQL_PORT=5432 \
	PSQL_USER=${PSQL_USER} \
	PSQL_SSLMODE=disable \
	PSQL_PASSWORD=${PSQL_PASSWORD} \
	PSQL_BLACKLIST=flyway_schema_history \
	sqlboiler --wipe --no-tests -o ${GENERATED_SRC_DIR}/models psql


gen-mock:
	mockery --all --dir=$(MOCKERY_TARGET_PATH) --output=$(MOCKERY_OUTPUT_PATH) --keeptree --with-expecter --inpackage=false --packageprefix='mock'


build: init generate docs
	@for app in ${APPS} ; do \
		echo "Building ${PROJECT_NAME}-$$app..."; \
		GOOS=${GOOS} \
		GOARCH=${GOARCH} \
		go build ${LDFLAGS} \
		-o ${TARGET_BIN_DIR}/${PROJECT_NAME}-$$app.${GOOS}.${GOARCH} \
		${PROJECT_PATH}/cmd/$$app ; \
	done

run:
	# general 기본 run으로 설정되어 있음
	${TARGET_BIN_DIR}/${PROJECT_NAME}-general.${GOOS}.${GOARCH}

dev: build run

test: build
	GO_ENV=$(TEST_ENV) go test `go list ./... | grep -v ./generated` -p 1

clean: clean-target clean-gen

clean-gen:
	rm -rf ${GENERATED_SRC_DIR}

clean-target:
	rm -rf ${TARGET_DIR}

docs: docs-clean docs-gen docs-fmt

docs-gen:
	@for doc in ${DOC_TARGETS} ; do \
		swag init -d api/http/$$doc -g swagger.go -o api/http/swagger --pd --instanceName swagger_$$doc ; \
	done

docs-clean:
	rm -rf api/http/swagger

docs-fmt:
	swag fmt -d api/http


database-init:
	@# error 는 무시
	@# (TODO) DATABASE_HOST가 localhost가 아니면 경고 나오고 사용자 인풋 이후 진행되게 수정하기
	-@createuser -s -l -h ${PSQL_HOST} -p 5432 -U postgres ${PSQL_USER}
	-@createdb -E UTF8 -T template0 --lc-collate=C --lc-ctype=en_US.UTF-8 -h ${PSQL_HOST} -p 5432 -U ${PSQL_USER} ${PSQL_DBNAME}
	-@createdb -E UTF8 -T template0 --lc-collate=C --lc-ctype=en_US.UTF-8 -h ${PSQL_HOST} -p 5432 -U ${PSQL_USER} ${PSQL_DBNAME}_test

## ------------- FLYWAY -------------
ifndef DATABASE_PASSWORD
ifeq ($(strip $(STAGE)),exp)
AWS_PROFILE := ch-dev
AWS_PARAMETER := /channel/exp/rds/ch-dev/app_store/password
else ifeq ($(strip $(STAGE)),production)
AWS_PROFILE := ch-prod
AWS_PARAMETER := /channel/production/rds/ch3-psql14/app_store/password
endif
AWS_REGION ?= ap-northeast-2
DATABASE_PASSWORD := $(shell \
  aws ssm get-parameter \
    --profile $(AWS_PROFILE) \
    --region $(AWS_REGION) \
    --name $(AWS_PARAMETER) \
    --with-decryption 2>/dev/null \
  | jq -r '.Parameter.Value' \
)
endif

FLYWAY_CONFIG := ${PROJECT_PATH}/config/flyway/flyway.conf
FLYWAY_MIGRATION := ${PROJECT_PATH}/resource/psql/migration

FLYWAY_CMD=@PSQL_HOST=${PSQL_HOST} \
	PSQL_DBNAME=${PSQL_DBNAME} \
	PSQL_SCHEMA=${PSQL_SCHEMA} \
	PSQL_USER=${PSQL_USER} \
	PSQL_PASSWORD=${PSQL_PASSWORD} flyway \
	-configFiles=${FLYWAY_CONFIG} -locations=filesystem:${FLYWAY_MIGRATION}


flyway-migrate:
	${FLYWAY_CMD} migrate

flyway-info:
	${FLYWAY_CMD} info

flyway-validate:
	${FLYWAY_CMD} validate

flyway-repair:
	${FLYWAY_CMD} repair

# 아래 flyway-clean은 위험 ...
flyway-clean:
	${FLYWAY_CMD} clean

migrate: flyway-migrate
	@true
