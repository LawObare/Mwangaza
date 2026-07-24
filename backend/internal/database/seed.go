// Package database (seed.go) inserts sample rows for development/demo.
//
// Seeds:
//   - 3 farms (Green Valley Farm, Sunrise Acres, Hilltop Farm)
//   - 3 satellite_data rows (one per farm, with mock sensor readings)
//   - 3 recommendations (irrigation, rainfall, temperature alerts)
//   - 3 sms_logs (sent/failed examples)
//
// Only runs when the tables are empty (checked via SELECT COUNT(*)).
// Called from main.go after database.Migrate().
package database