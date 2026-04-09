.PHONY: up down build test lint tidy

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

# Add per-service commands later, e.g., make user-up