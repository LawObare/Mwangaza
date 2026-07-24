// Package satellite (mock.go) returns hardcoded satellite data for demos.
//
// Function:
//   func GetMockData() models.SatelliteData
//
// Returns consistent values:
//   SoilMoisture:  18.5
//   Temperature:   31.0
//   RainProb:      20.0
//   NDVI:          0.64
//   WindSpeed:     8.0
//   Timestamp:     time.Now()
//
// No randomness. Used when USE_MOCK_DATA=true in env.
// Called by service.go → FetchSatelliteData().
package satellite