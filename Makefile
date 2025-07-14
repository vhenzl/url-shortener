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
