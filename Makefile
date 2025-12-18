# -----------------------
# Variables
# -----------------------
APP_NAME=apartments-api
DC=docker-compose -f docker-compose.yml

# -----------------------
# Build
# -----------------------
build-api:
	go build -o bin/$(APP_NAME) cmd/api/main.go

# -----------------------
# Run locally (no Docker)
# -----------------------
run-api:
	go run cmd/api/main.go

# -----------------------
# Docker
# -----------------------
docker-up:
	$(DC) up -d

docker-down:
	$(DC) down

docker-restart: docker-down docker-up

docker-logs:
	$(DC) logs -f

# -----------------------
# Migrations
# -----------------------
migrate-postgres:
	docker exec -i apartments-postgres psql \
		-U postgres \
		-d apartments \
		-f migrations/postgres/001_init.sql

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
	@echo "  build-api        Build API binary"
	@echo "  run-api          Run API locally"
	@echo "  docker-up        Start services (Postgres + Elasticsearch + API)"
	@echo "  docker-down      Stop services"
	@echo "  docker-restart   Restart services"
	@echo "  docker-logs      View docker logs"
	@echo "  migrate-postgres Apply Postgres migrations"
	@echo "  clean            Remove binaries"
