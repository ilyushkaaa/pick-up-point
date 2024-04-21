include .env
export


MIGRATION_FOLDER=$(CURDIR)/pkg/infrastructure/database/migrations

.PHONY: app_start
app_start:
	docker-compose up -d zookeeper kafka1 kafka2 kafka3 postgres redis
	go run $(CURDIR)/cmd/pick_up_point_server/main.go


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

.PHONY: integration_tests_run
integration_tests_run:
	docker-compose -f $(CURDIR)/tests/integration_tests/docker-compose.yml up -d
	make migration-up DB_HOST=localhost DB_USER=postgres DB_PASS=test DB_NAME=test_postgres DB_PORT=5433
	DB_HOST=localhost DB_USER=postgres DB_PASS=test DB_NAME=test_postgres DB_PORT=5433 REDIS_HOST=localhost REDIS_PORT=6380 go test -tags=integration $(CURDIR)/tests/integration_tests/...
	docker-compose -f $(CURDIR)/tests/integration_tests/docker-compose.yml down

.PHONY: unit_tests_run
unit_tests_run:
	go test ./internal...

