.PHONY: help up down build tidy test lint

help:
	@echo "Available commands:"
	@echo "  make up      - Start all services"
	@echo "  make down    - Stop all services"
	@echo "  make build   - Build all services"
	@echo "  make tidy    - Clean go modules"

up:
	docker-compose up --build -d

down:
	docker-compose down -v

build:
	go build ./...

tidy:
	go mod tidy

test:
	go test ./... -v

lint:
	golangci-lint run