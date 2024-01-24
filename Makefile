STAGE ?= dev
PROJECT_PATH := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
CMD_PATH= $(PROJECT_PATH)/cmd/appstore
BIN_PATH = $(PROJECT_PATH)/bin/main
TEST_ENV = "test"

#Flyway
FLYWAY_CONFIG_PATH = $(PROJECT_PATH)/config/flyway/$(STAGE).conf
FLYWAY_SQL_PATH = $(PROJECT_PATH)/resources/psql/migration

# Sql boiler
SQLBOILER_CONFIG_PATH = $(PROJECT_PATH)/config/sqlboiler/$(STAGE).toml

# Mockery
MOCKERY_TARGET_PATH = $(PROJECT_PATH)/internal
MOCKERY_OUTPUT_PATH = $(PROJECT_PATH)/test/mock

build: generate
	go build -o $(BIN_PATH) $(CMD_PATH)

test:
	GO_ENV=$(TEST_ENV) go test $(PROJECT_PATH)

clean:
	rm $(BIN_PATH)

installTools:
	# SqlBoiler
	go install github.com/volatiletech/sqlboiler/v4@v4.14.1
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.14.1
	# Mockery
	go install github.com/vektra/mockery/v2@v2.26.1

generate: flyway_clean flyway_migrate genBoiler genMock

genBoiler:
	sqlboiler --wipe --no-tests -o generated/models -c $(SQLBOILER_CONFIG_PATH) psql

genMock:
	mockery --all --dir=$(MOCKERY_TARGET_PATH) --output=$(MOCKERY_OUTPUT_PATH) --keeptree --with-expecter --inpackage=false --packageprefix='mock'

genSwaggo:
	# api/http/{front|desk|admin}/handler/util/handler.go 파일의 swaggerInstanceName 상수와 instanceName 옵션의 값이 동일해야 합니다.
#	swag init -d api/http/front -g module.go -o api/http/front/swagger --instanceName swagger_front --pd
#	swag init -d api/http/desk -g module.go -o api/http/desk/swagger --instanceName swagger_desk --pd
#	swag init -d api/http/admin  -g module.go -o api/http/admin/swagger  --instanceName swagger_admin  --pd
	swag init -d api/http/general  -g module.go -o api/http/general/swagger  --instanceName swagger_general  --pd

swagFmt:
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

