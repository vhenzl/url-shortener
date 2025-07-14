# URL Shortener API

A simple URL shortener web API written in Go.

## Structure

- `cmd/urlshortener/` - Main application entrypoint
- `internal/app/` - App-wide logic, wiring, use cases, service layer
- `internal/handler/` - HTTP handlers (controllers)
- `internal/domain/` - Domain: core logic, models, interfaces
- `internal/storage/` - Data persistence (in-memory, DBs, etc)

## Getting Started

1. Install Go (https://golang.org/doc/install)
2. Run `go mod tidy` to install dependencies
3. Run the app:
   ```bash
   go run ./cmd/urlshortener
   ```
