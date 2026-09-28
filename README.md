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

## Run with Docker

Build the Flutter web bundle first — this bakes the backend URL into the app:

```bash
cd frontend/mwangaza
flutter build web --dart-define=MWANGAZA_API_URL=http://localhost:8080/api
cd ../..
```

Then start the whole stack:

```bash
docker compose up --build
```

- Backend: `http://localhost:8080` (health: `/api/health`)
- Frontend: `http://localhost:8000`

Configuration is read from `backend/.env` (copy `backend/.env.example` to
`backend/.env` first if needed). The demo store persists in the `backend-data`
named volume; the API key is never baked into the image — `.env` is excluded
from the build context. Stop everything with `docker compose down` (add `-v`
to also reset the stored data).
