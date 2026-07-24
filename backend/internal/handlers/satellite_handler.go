// Package handlers (satellite_handler.go) serves the satellite data endpoint.
//
// Endpoint:
//   GET /api/satellite — fetches latest satellite data via satellite.FetchSatelliteData()
//
// This handler calls the satellite service package (not the database directly).
// If USE_MOCK_DATA=true, the service returns mock.go data.
// Otherwise it hits the SpaceIoTBox API via client.go and parses via parser.go.
//
// Response shape:
//   { "success": true, "data": { "soil_moisture": ..., "temperature": ..., ... } }
package handlers