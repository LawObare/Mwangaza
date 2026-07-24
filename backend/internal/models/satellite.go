// Package models (satellite.go) defines the SatelliteData struct.
//
// This is the canonical representation of environmental data returned
// by the satellite service. It is what the Flutter dashboard displays.
//
// Fields:
//   SoilMoisture  float64 — percentage (0-100)
//   Temperature   float64 — Celsius
//   RainProb      float64 — percentage (0-100)
//   NDVI          float64 — Normalized Difference Vegetation Index (-1 to 1)
//   WindSpeed     float64 — meters per second
//   Timestamp     time.Time — when the data was recorded
//
// Populated by services/satellite/parser.go from SpaceIoTBox JSON
// or directly by services/satellite/mock.go.
package models