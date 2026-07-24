// Package main is the application entry point.
// It loads configuration from env, initializes the SQLite database,
// runs migrations, optionally seeds data, and starts the Gin HTTP server.
//
// Flow:
//   1. config.Load() → reads PORT, DATABASE_PATH, SPACEIOTBOX_* env vars
//   2. database.Connect(cfg.DatabasePath) → opens SQLite via mattn/go-sqlite3
//   3. database.Migrate() → creates farms, satellite_data, recommendations, sms_logs tables
//   4. routes.Setup(db) → registers all API handlers on a Gin engine
//   5. r.Run(":" + PORT) → starts listening
//
// The db *sql.DB instance is passed through to handlers/services via dependency injection.
package main