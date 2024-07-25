# include
include .env


# Swagger
swag:
	@echo "Generating Swagger... "
	swag init --parseDependency --parseInternal -g internal/app/router/api.go --output docs/app

# local migrations
local-migration-up:
	@echo "Starting local migration... UP"
	migrate -path db/migration -database "postgresql://${LOCAL_POSTGRES_USER}:${LOCAL_POSTGRES_SECRET}@${LOCAL_POSTGRES_HOST}:${LOCAL_POSTGRES_PORT}/${LOCAL_POSTGRES_DBNAME}?sslmode=disable" -verbose up

local-migration-down:
	@echo "Starting local migration... DOWN"
	migrate -path db/migration -database "postgresql://${LOCAL_POSTGRES_USER}:${LOCAL_POSTGRES_SECRET}@${LOCAL_POSTGRES_HOST}:${LOCAL_POSTGRES_PORT}/${LOCAL_POSTGRES_DBNAME}?sslmode=disable" -verbose down

run:
	@echo "Running Golang App... LOCAL"
	go run cmd/main.go local

# internal commands
sqlc:
	@echo "Generating SQLC files..."
	sqlc generate

.PHONY: run sqlc local-migration-up local-migration-down swag

