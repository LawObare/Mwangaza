# Database

## Technology

SQLite (via `github.com/mattn/go-sqlite3`).

## Location

The database file is stored at `./data/lakenet.db` (configurable via `DATABASE_PATH` env var).

## Tables

### farms
Stores registered farms with location and contact info.

### satellite_data
Stores environmental readings per farm (soil moisture, temperature, rain, NDVI, wind speed).

### recommendations
Stores generated advice per farm with priority levels.

### sms_logs
Records all SMS messages sent through the system.

## Migration

Tables are created automatically on startup via `CREATE TABLE IF NOT EXISTS`.

## Seeding

Sample data is inserted when tables are empty (development/demo mode).