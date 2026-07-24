// Package recommendation (rainfall.go) analyzes rain probability data.
//
// Logic:
//   - If rain_probability > 60% → "Delay irrigation — rain expected" (priority: medium)
//   - If rain_probability 30-60% → "Monitor weather" (priority: low)
//   - Otherwise → no rainfall recommendation
//
// Consumes SatelliteData.RainProb. Called by rules.go.
package recommendation