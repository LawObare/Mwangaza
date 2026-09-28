# API Contract

Defines the request/response contract between the Flutter frontend and Go backend.

## Endpoints

All endpoints return JSON with shape: `{ "success": bool, "data": ..., "error": "..." }`

| Method | Path              | Description              |
|--------|-------------------|--------------------------|
| POST   | /api/auth/register | Create an account; returns `{token, user}` |
| POST   | /api/auth/login    | Exchange credentials for a JWT |
| GET    | /api/auth/me       | Current account (requires Bearer token) |
| GET    | /api/health       | Health check             |
| GET    | /api/farms        | List all farms           |
| POST   | /api/farms        | Register a farm and fetch its first satellite snapshot |
| GET    | /api/farms/{id}   | Get single farm          |
| GET    | /api/satellite    | Get satellite data       |
| GET    | /api/recommendation | Get recommendations    |
| POST   | /api/recommendation | Generate recommendations for a farm |
| GET    | /api/sms          | Get SMS history          |
| POST   | /api/sms/send     | Send an SMS              |

## Error Codes

- 400 — Bad request (invalid input)
- 404 — Resource not found
- 500 — Internal server error