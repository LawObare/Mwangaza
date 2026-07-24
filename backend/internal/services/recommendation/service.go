// Package recommendation (service.go) is the decision engine entry point.
//
// It analyzes satellite data for each farm and generates actionable
// recommendations stored in the recommendations table.
//
// Public function:
//   func GenerateRecommendations(db *sql.DB, data models.SatelliteData) error
//
// Internally delegates to:
//   - rules.go      — top-level decision logic
//   - irrigation.go — soil moisture thresholds → irrigation advice
//   - rainfall.go   — rain probability → delay/advance advice
//   - temperature.go — heat stress alerts
//   - ndvi.go       — vegetation health analysis
//
// The handler at handlers/recommendation_handler.go reads from the
// recommendations table; the handler does NOT call this service directly
// in the current architecture (the engine runs on a schedule or trigger).
package recommendation