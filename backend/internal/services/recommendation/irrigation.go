// Package recommendation (irrigation.go) handles irrigation-related decisions.
//
// Logic:
//   - If soil_moisture < 15% → "Irrigate for 20 minutes" (priority: high)
//   - If soil_moisture 15-25% → "Monitor soil moisture" (priority: low)
//   - Otherwise → no irrigation recommendation
//
// Consumes soil moisture from SatelliteData.SoilMoisture.
// Called by the rule engine in rules.go.
package recommendation