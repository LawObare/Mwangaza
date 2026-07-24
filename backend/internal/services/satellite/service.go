// Package satellite (service.go) is the public entry point for satellite data.
//
// This is the only file other packages should import. It exposes:
//   func FetchSatelliteData() (models.SatelliteData, error)
//
// Decision logic:
//   - If USE_MOCK_DATA == "true" → return GetMockData()
//   - Otherwise → FetchLiveData() → ParseSatelliteData(raw) → return
//
// The handler at handlers/satellite_handler.go calls this function.
package satellite