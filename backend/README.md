# Mwangaza Backend

Go-based REST API for farm advisory system.

## Tech Stack

- **Framework:** Gin
- **Database:** SQLite
- **Satellite Data:** SpaceIoTBox API
- **SMS:** Africa's Talking

## Getting Started

1. Copy `.env.example` to `.env` and configure variables
2. Install Go dependencies: `go mod tidy`
3. Run the server: `go run cmd/main.go`

## Environment Variables

See `.env.example` for all required configuration.

## Project Structure

```
cmd/main.go          — Entry point
internal/config/     — Environment configuration
internal/database/   — SQLite connection, migrations, seeds
internal/handlers/   — HTTP request handlers
internal/middleware/  — CORS middleware
internal/models/      — Data structures
internal/routes/      — Route registration
internal/services/    — Business logic (satellite, recommendation, SMS)
internal/utils/       — Response helpers
data/                 — SQLite database file
docs/                 — Project documentation
```