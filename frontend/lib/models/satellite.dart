// lib/models/satellite.dart — Satellite data model.
//
// Fields:
//   double soilMoisture, double temperature, double rainProbability,
//   double ndvi, double windSpeed, String timestamp
//
// Responsibilities:
//   - Define a SatelliteData class with a fromJson() factory constructor.
//   - JSON keys match the Go backend GET /api/satellite response.
//
// Used by:
//   services/satellite_service.dart — deserializes API response.
//   screens/dashboard/ — displays current environmental data.
//   widgets/satellite_card.dart — renders the satellite data card.
class SatelliteData {}