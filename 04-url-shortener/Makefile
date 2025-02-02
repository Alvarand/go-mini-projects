include .env

up:
	docker-compose up --detach

down:
	docker-compose down

migrate:
	goose -dir migrations/  postgres "host=localhost port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) database=$(POSTGRES_DB)" up

run:
	go run cmd/main.go
