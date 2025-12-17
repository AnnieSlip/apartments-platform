# -----------------------
# Variables
# -----------------------
APP_NAME=apartments-api
PRECOMPUTE_NAME=precompute
WEEKLY_NAME=weekly-job

# Docker Compose
DC=docker-compose -f docker-compose.yml

# -----------------------
# Build
# -----------------------
build-api:
	go build -o bin/$(APP_NAME) cmd/api/main.go

build-precompute:
	go build -o bin/$(PRECOMPUTE_NAME) cmd/precompute/main.go

build-weekly:
	go build -o bin/$(WEEKLY_NAME) cmd/weekly-job/main.go

build-all: build-api build-precompute build-weekly

# -----------------------
# Run
# -----------------------
run-api:
	go run cmd/api/main.go

run-precompute:
	go run cmd/precompute/main.go

run-weekly:
	go run cmd/weekly-job/main.go

# -----------------------
# Docker
# -----------------------
docker-up:
	$(DC) up -d

docker-down:
	$(DC) down

docker-restart: docker-down docker-up

# -----------------------
# Migrations
# -----------------------
migrate-postgres:
	docker exec -i apartments-postgres psql -U postgres -d apartments -f migrations/postgres/001_init.sql

migrate-cassandra:
	docker exec -i apartments-cassandra cqlsh -f migrations/cassandra/001_init.cql

migrate-all: migrate-postgres migrate-cassandra

# -----------------------
# Clean
# -----------------------
clean:
	rm -rf bin/*

# -----------------------
# Help
# -----------------------
help:
	@echo "Makefile commands:"
	@echo "  build-api           Build API binary"
	@echo "  build-precompute    Build precompute binary"
	@echo "  build-weekly        Build weekly job binary"
	@echo "  build-all           Build all binaries"
	@echo "  run-api             Run API"
	@echo "  run-precompute      Run precompute job"
	@echo "  run-weekly          Run weekly job"
	@echo "  docker-up           Start all services via docker-compose"
	@echo "  docker-down         Stop all services via docker-compose"
	@echo "  migrate-postgres    Apply Postgres migrations"
	@echo "  migrate-cassandra   Apply Cassandra migrations"
	@echo "  clean               Remove binaries"
