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

