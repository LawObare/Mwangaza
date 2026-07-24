// Package recommendation (ndvi.go) analyzes vegetation health from NDVI.
//
// Logic:
//   - If NDVI < 0.3 → "Poor vegetation health — check for pests/disease" (priority: high)
//   - If NDVI 0.3-0.5 → "Moderate vegetation health" (priority: low)
//   - If NDVI > 0.5 → no NDVI recommendation (healthy)
//
// Consumes SatelliteData.NDVI. Called by rules.go.
package recommendation