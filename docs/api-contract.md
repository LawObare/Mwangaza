# API Contract

Defines the request/response contract between the Flutter frontend and Go backend.

## Endpoints

All endpoints return JSON with shape: `{ "success": bool, "data": ..., "error": "..." }`

| Method | Path              | Description              |
|--------|-------------------|--------------------------|
| GET    | /api/health       | Health check             |
| GET    | /api/farms        | List all farms           |
| GET    | /api/farms/:id    | Get single farm          |
| GET    | /api/satellite    | Get satellite data       |
| GET    | /api/recommendation | Get recommendations    |
| GET    | /api/sms          | Get SMS history          |
| POST   | /api/sms/send     | Send an SMS              |

## Error Codes

- 400 — Bad request (invalid input)
- 404 — Resource not found
- 500 — Internal server error