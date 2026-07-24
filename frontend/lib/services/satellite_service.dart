// lib/services/satellite_service.dart — Satellite data API service.
//
// Endpoint:
//   GET /api/satellite → getSatelliteData() returns SatelliteData
//
// Responsibilities:
//   - Use ApiService to call the backend satellite endpoint.
//   - Deserialize JSON into models/satellite.dart.
//
// Connects to:
//   api_service.dart       — makes HTTP requests.
//   models/satellite.dart  — target model.
//   screens/dashboard/     — displays satellite data in stat cards and satellite card.
class SatelliteService {}