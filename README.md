# Mwangaza — Farm Advisory Dashboard

Satellite-driven farm advisory system. Fetches environmental data from SpaceIoTBox, runs a decision engine, and sends SMS alerts to farmers via Africa's Talking.

## Structure

```
backend/     — Go + Gin REST API (SQLite database)
frontend/    — Flutter dashboard (OpenStreetMap)
database/    — SQL schema and seed scripts
docs/        — Project documentation
```

## Quick Start

**Backend:** `cd backend && go run cmd/main.go`

**Frontend:** `cd frontend && flutter run`

Set `USE_MOCK_DATA=true` in `.env` for offline demo.