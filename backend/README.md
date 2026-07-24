# Mwangaza Backend

Hackathon-friendly Go REST API for the farm advisory demo.

## What It Does

- Serves farm, satellite, recommendation, and SMS endpoints
- Uses a lightweight file-backed demo store
- Falls back to mock satellite and SMS data when `USE_MOCK_DATA=true`

## Run It

1. `cd backend`
2. Copy `.env.example` to `.env` if needed
3. Run `go run ./cmd`

## Environment Variables

- `PORT` - HTTP server port, default `8080`
- `DATABASE_PATH` - path to the local demo store file, default `./data/lakenet.db`
- `USE_MOCK_DATA` - `true` for offline demo mode, default `true`
- `SPACEIOTBOX_API_KEY` - optional live satellite API key
- `SPACEIOTBOX_BASE_URL` - optional live satellite API base URL
- `SMS_API_KEY` - optional Africa's Talking API key
- `SMS_USERNAME` - optional Africa's Talking username
- `SMS_SENDER_ID` - optional sender ID

## Key Endpoints

- `GET /api/health`
- `GET /api/farms`
- `POST /api/farms`
- `GET /api/farms/{id}`
- `GET /api/satellite`
- `GET /api/recommendation`
- `POST /api/recommendation`
- `GET /api/sms`
- `POST /api/sms/send`
