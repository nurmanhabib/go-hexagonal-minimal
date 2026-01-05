APP_NAME=api
DB_DSN?=root:pass@tcp(127.0.0.1:3306)/test

.PHONY: help run migrate-up migrate-down migrate-status

help:
	@echo "Available commands:"
	@echo "  make run              Run HTTP API"
	@echo "  make migrate-up       Run database migrations"
	@echo "  make migrate-down     Rollback last migration"
	@echo "  make migrate-status   Show migration status"

run:
	go run cmd/api/main.go

migrate-up:
	go run cmd/migrate/main.go

migrate-down:
	go run cmd/migrate/main.go down

migrate-status:
	go run cmd/migrate/main.go status
