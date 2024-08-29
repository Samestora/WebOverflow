include .env
DOCKER_COMPOSE_FILE=docker-compose.yml
MIGRATE_PATH=$(shell which migrate)
.PHONY: all build run test clean

build:
	@echo "Go is building..."
	go build -o ./tmp/main ./cmd

docker-build:
	@echo "Docker is building for prod..."
	docker build .

docker-up:
	@echo "Starting docker contaiers..."
	docker compose -f ${DOCKER_COMPOSE_FILE} up -d

docker-down:
	@echo "Terminating docker containers..."
	docker compose -f ${DOCKER_COMPOSE_FILE} down

migrate-up:
	@echo "Migrate make..."
	$(MIGRATE_PATH) -path ./platform/migrations/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@localhost/${DB_NAME}?sslmode=disable" up

all: build docker-build docker-up migrate-up
