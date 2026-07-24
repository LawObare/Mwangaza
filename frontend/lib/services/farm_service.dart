// lib/services/farm_service.dart — Farm API service.
//
// Endpoints:
//   GET /api/farms      → getFarms() returns List<Farm>
//   GET /api/farms/:id  → getFarm(id) returns Farm
//
// Responsibilities:
//   - Use ApiService to call backend endpoints.
//   - Deserialize JSON responses into model objects.
//
// Connects to:
//   api_service.dart  — makes HTTP requests.
//   models/farm.dart   — target model for deserialization.
//   screens/farms/, screens/map/ — consumers of farm data.
class FarmService {}