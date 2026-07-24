// Package config loads environment configuration for the backend.
// It reads values from .env or system environment and exposes them
// as a typed Config struct.
//
// Required env vars:
//   PORT              - HTTP server port (default: "8080")
//   DATABASE_PATH     - path to SQLite database file (default: "./data/lakenet.db")
//   SPACEIOTBOX_API_KEY   - API key for SpaceIoTBox
//   SPACEIOTBOX_BASE_URL  - base URL for SpaceIoTBox API
//   USE_MOCK_DATA         - "true" to use mock satellite/SMS data (default: "true")
//   SMS_API_KEY           - API key for Africa's Talking
//
// Config struct is consumed by main.go to initialize the database and services.
package config