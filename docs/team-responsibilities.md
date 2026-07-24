# Team Responsibilities

## Person 1 — Backend Core
- `cmd/main.go` — Server entry point
- `internal/config/` — Environment configuration
- `internal/database/` — SQLite connection, migrations, seeding
- `internal/routes/` — Route registration with Gin
- `internal/handlers/` — All HTTP handlers
- `internal/middleware/` — CORS middleware
- `internal/utils/` — Response helpers

## Person 2 — Satellite Service
- `internal/services/satellite/` — SpaceIoTBox integration
- `service.go` — Public entry point (FetchSatelliteData)
- `client.go` — HTTP client for SpaceIoTBox API
- `parser.go` — JSON parsing into models.SatelliteData
- `mock.go` — Mock data for offline demo
- `types.go` — API response structs

## Person 3 — Decision Engine
- `internal/services/recommendation/` — Recommendation logic
- `service.go` — Entry point
- `rules.go` — Top-level decision logic
- `irrigation.go` — Soil moisture thresholds
- `rainfall.go` — Rain probability analysis
- `temperature.go` — Heat stress detection
- `ndvi.go` — Vegetation health analysis

## Person 4 — SMS Service
- `internal/services/sms/` — SMS integration
- `service.go` — Entry point (mock vs live)
- `africastalking.go` — Africa's Talking API client
- `mock.go` — Mock SMS sender
- `types.go` — SMS request/response types

## Person 5 — Flutter Frontend
- `frontend/lib/` — Entire Flutter application
- Screens: Dashboard, Farms, Farm Details, Map, SMS History
- Widgets: StatCard, FarmCard, SatelliteCard, RecommendationCard, LoadingWidget
- Services: API client and per-feature services
- Models: Farm, SatelliteData, Recommendation, SmsMessage