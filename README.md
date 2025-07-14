# URL Shortener API

A simple URL shortener web API written in Go.

## Structure

- `cmd/urlshortener/` - Main application entrypoint
- `internal/application/` - Application logic, use cases, etc
- `internal/presentation/` - HTTP handlers (controllers)
- `internal/domain/` - Domain models and logic
- `internal/infrastructure/` - Data persistence, etc

## Getting Started

1. Install Go (https://golang.org/doc/install)
2. Run `go mod tidy` to install dependencies
3. Run the app:
   ```bash
   go run ./cmd/urlshortener
   ```
