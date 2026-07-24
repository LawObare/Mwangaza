# System Architecture

## High-Level Design

```
┌──────────────────┐
│  Flutter Dashboard│
│  (Admin UI)       │
└────────┬─────────┘
         │ HTTP (REST)
         ▼
┌──────────────────┐
│  Go Backend (Gin) │
└────────┬─────────┘
    ┌────┼────┬────┐
    ▼    ▼    ▼    ▼
 SpaceIoT  SQLite  SMS
 API              Service
                   │
                   ▼
             Africa's Talking
                   │
                   ▼
            Farmer Phone (SMS)
```

## Components

### Flutter Dashboard
- Screens: Dashboard, Farms, Farm Details, Map, SMS History
- Map uses OpenStreetMap via `flutter_map` package
- Consumes Go backend REST API

### Go Backend
- Gin HTTP router with CORS middleware
- SQLite for persistence
- Satellite service (mock or live SpaceIoTBox)
- Recommendation engine (threshold-based rules)
- SMS service (mock or Africa's Talking)

### SQLite Database
- Tables: farms, satellite_data, recommendations, sms_logs
- Single file at `data/lakenet.db`
- Auto-migrated on startup

### External APIs
- SpaceIoTBox — satellite environmental data
- Africa's Talking — SMS gateway