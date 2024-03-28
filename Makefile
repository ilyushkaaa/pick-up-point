include .env
export


MIGRATION_FOLDER=$(CURDIR)/pkg/database/migrations

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) host=$(DB_HOST) port=$(DB_PORT) sslmode=disable" up

.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) host=$(DB_HOST) port=$(DB_PORT) sslmode=disable" down

.PHONY: build
build:
	go build cmd/pick_up_point_server/main.go
