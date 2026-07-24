// Package database (migrate.go) creates the SQLite schema on startup.
//
// Tables to create:
//   farms            - id, name, farmer, phone, latitude, longitude
//   satellite_data   - id, farm_id (FK), soil_moisture, temperature,
//                      rain_probability, ndvi, wind_speed, recorded_at
//   recommendations  - id, farm_id (FK), message, priority, reason, created_at
//   sms_logs         - id, phone, message, status, sent_at
//
// Run on every server start; uses CREATE TABLE IF NOT EXISTS.
// Called from main.go after database.Connect().
package database