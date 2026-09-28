# Mwangaza — Farm Advisory Dashboard

Satellite-driven farm advisory system. Fetches weather data from OpenWeatherMap, runs a decision engine, and sends SMS alerts to farmers via Africa's Talking.

## Structure

```
backend/     — Go API + file-backed demo store
frontend/    — Flutter dashboard (OpenStreetMap)
database/    — SQL schema and seed scripts
docs/        — Project documentation
```

## Quick Start

**Backend:** `cd backend && go run cmd/main.go` (listens on `http://localhost:8080`)

**Frontend:** `cd frontend/mwangaza && flutter run`

Set `USE_MOCK_DATA=true` in `.env` for offline demo.

The Flutter app defaults to `http://10.0.2.2:8080/api`, which is the host alias
used by the Android emulator. For desktop, web, or a physical device, point it
at the backend that device can reach:

```
flutter run --dart-define=MWANGAZA_API_URL=http://localhost:8080/api
```
