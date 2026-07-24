// Package recommendation (temperature.go) handles temperature-based alerts.
//
// Logic:
//   - If temperature > 35°C → "Heat stress detected — increase watering" (priority: high)
//   - If temperature 30-35°C → "Monitor temperature" (priority: low)
//   - Otherwise → no temperature alert
//
// Consumes SatelliteData.Temperature. Called by rules.go.
package recommendation