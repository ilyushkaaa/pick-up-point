include .env
export


MIGRATION_FOLDER=$(CURDIR)/pkg/database/migrations
TEST_ENV_FILE=$(CURDIR)/.env.testing

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
	docker-compose --env-file $(CURDIR)/.env.testing -f $(CURDIR)/tests/integration_tests/docker-compose.yml up -d --build
	make migration-up DB_HOST=$(DB_HOST) DB_USER=$(DB_USER) DB_PASS=$(DB_PASS) DB_NAME=$(DB_NAME) DB_PORT=$(DB_PORT)

.PHONY: integration_tests_run
integration_tests_run:
	go test $(CURDIR)/tests/integration_tests/pick-up-points/... DB_HOST=$(DB_HOST) DB_USER=$(DB_USER) DB_PASS=$(DB_PASS) DB_NAME=$(DB_NAME) DB_PORT=$(DB_PORT)


.PHONY: unit_tests_run
unit_tests_run:
	go test $(CURDIR)/internal/pick-up_point/delivery/http/...
	go test $(CURDIR)/internal/pick-up_point/delivery/cli/...
	go test $(CURDIR)/internal/pick-up_point/service/...
	go test $(CURDIR)/internal/pick-up_point/storage/file/...
