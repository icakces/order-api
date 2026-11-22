-include .env
export

DB_URL = mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

.PHONY: db-up db-down migrate-up migrate-down migrate-version run build test vet check

db-up:
	docker compose up -d
	docker compose ps

db-down:
	docker compose stop mysql

migrate-up:
	@migrate -path ./migrations -database "$(DB_URL)" up

migrate-down:
	@migrate -path ./migrations -database "$(DB_URL)" down 1

migrate-version:
	@migrate -path ./migrations -database "$(DB_URL)" version

run:
	go run ./cmd/api

build:
	go build -o bin/api ./cmd/api

test:
	go test ./... -count=1

vet:
	go vet ./...

check: vet test