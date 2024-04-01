include .env
export


MIGRATION_FOLDER=$(CURDIR)/pkg/database/migrations

.PHONY: app_start
app_start:
	docker-compose up --build

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) host=$(DB_HOST) port=$(DB_PORT) sslmode=disable" up

.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) host=$(DB_HOST) port=$(DB_PORT) sslmode=disable" down

.PHONY: build
build:
	go build cmd/pick_up_point_server/main.go

.PHONY: test_env_up
test_env_up:
	docker-compose -f $(CURDIR)/tests/integration_tests/docker-compose.yml up

.PHONY: integration_tests_run
integration_tests_run:
	go test $(CURDIR)/tests/integration_tests/pick-up-points/...


.PHONY: unit_tests_run
unit_tests_run:
	go test $(CURDIR)/internal/pick-up_point/delivery/http/...
	go test $(CURDIR)/internal/pick-up_point/delivery/cli/...
	go test $(CURDIR)/internal/pick-up_point/service/...
	go test $(CURDIR)/internal/pick-up_point/storage/file/...
