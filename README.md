# URL Shortener API

A simple URL shortener web API written in Go.

## Structure

- `cmd/urlshortener/` - Main application entrypoint
- `db/migrations/` - Database schema migrations
- `internal/application/` - Application logic, use cases, etc
- `internal/presentation/` - HTTP handlers (controllers)
- `internal/domain/` - Domain models and logic
- `internal/infrastructure/` - Data persistence, etc

## Prerequisites

1. Install Go (version 1.24 or later):
   - Follow instructions at https://golang.org/doc/install
   - Add `$HOME/go/bin` to your PATH:
     ```bash
     echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
     source ~/.bashrc
     ```

2. Install golang-migrate:
   ```bash
   go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```

## Getting Started

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Create `.env` file from `.env.example`

3. Run migrations:
   ```bash
   make migrate-up
   ```

4. Run the app:
   ```bash
   go run ./cmd/urlshortener
   ```
