include .env
export

APP_NAME := urlshortener
BIN_DIR := ./bin
CMD_DIR := ./cmd/urlshortener

.PHONY: build run fmt test clean

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)

run:
	go run $(CMD_DIR)

fmt:
	go fmt ./...

lint:
	go vet ./...

test:
	go test ./...

clean:
	rm -rf $(BIN_DIR)/*

# Database migration commands
.PHONY: migrate-up migrate-down migrate-status migrate-new

MIGRATE := migrate
MIGRATE_DIR := ./migrations

migrate-up:
	$(MIGRATE) -path $(MIGRATE_DIR) -database "$$DATABASE_URL" up

migrate-down:
	$(MIGRATE) -path $(MIGRATE_DIR) -database "$$DATABASE_URL" down 1

migrate-status:
	$(MIGRATE) -path $(MIGRATE_DIR) -database "$$DATABASE_URL" version

migrate-new:
ifndef NAME
	$(error NAME is required. Usage: make migrate-new NAME=create_table)
endif
	$(MIGRATE) create -ext sql -dir $(MIGRATE_DIR) -seq $(NAME)
