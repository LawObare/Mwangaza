# Mwangaza — Farm Advisory Dashboard

Satellite-driven farm advisory system. Fetches weather data from OpenWeatherMap, runs a decision engine, and sends SMS alerts to farmers via Africa's Talking.

## Structure

```
backend/     — Go API + file-backed demo store
frontend/    — Flutter dashboard (OpenStreetMap)
database/    — SQL schema and seed scripts
docs/        — Project documentation
```

## Prerequisites

- **Go 1.25+** — `go version`
- **Flutter** (Dart SDK `^3.12.2`) — `flutter --version`
- **Docker + Docker Compose** — only needed for the containerized run
- **API keys** (optional): OpenWeatherMap (`WEATHER_API_KEY`) and Africa's
  Talking (`SMS_API_KEY`). Neither is required for the offline demo — the
  backend falls back to mock weather/SMS data when `USE_MOCK_DATA=true`.

## 1. Run the backend (Go API)

```bash
cd backend
cp .env.example .env        # skip if backend/.env already exists
go run ./cmd
```

The API listens on `http://localhost:8080`. Verify it in another terminal:

```bash
curl http://localhost:8080/api/health
# {"success":true,"data":{"status":"ok",...}}
```

Notes:

- The SQLite demo store is created, migrated, and seeded automatically on
  first start at `backend/data/lakenet.db`. The reference SQL lives in
  `database/schema.sql` and `database/seed.sql`.
- For a fully offline demo set `USE_MOCK_DATA=true` in `backend/.env`.
- Values exported in your shell override the `.env` file, e.g.
  `PORT=9090 go run ./cmd`.
- Full environment-variable reference: [`backend/README.md`](backend/README.md).

## 2. Run the frontend (Flutter dashboard)

```bash
cd frontend/mwangaza
flutter pub get
flutter run
```

The backend URL is compiled into the app at build time via the
`MWANGAZA_API_URL` define (see `lib/constants/backend_uri.dart`):

- Default: `http://10.0.2.2:8080/api` — the host alias used by the **Android
  emulator**, so plain `flutter run` works against a local backend there.
- For **desktop, web, or a physical device**, point it at the backend that
  device can reach (use your machine's LAN IP for a physical device):

```bash
flutter run --dart-define=MWANGAZA_API_URL=http://localhost:8080/api
```

## Run the whole stack with Docker

Build the Flutter web bundle first — this bakes the backend URL into the app:

```bash
cd frontend/mwangaza
flutter build web --dart-define=MWANGAZA_API_URL=http://localhost:8080/api
cd ../..
```

Then start everything:

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

## Tests

```bash
cd backend && go test ./...              # Go unit tests
cd frontend/mwangaza && flutter test     # Flutter tests
```
