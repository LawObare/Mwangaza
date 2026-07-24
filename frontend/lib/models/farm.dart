// lib/models/farm.dart — Farm data model.
//
// Fields:
//   int id, String name, String farmer, double latitude, double longitude,
//   String status, double? soilMoisture, double? temperature, String? lastSmsSent
//
// Responsibilities:
//   - Define a Farm class with a fromJson() factory constructor.
//   - JSON keys come from the Go backend GET /api/farms response.
//
// Used by:
//   services/farm_service.dart — deserializes API responses.
//   screens/farms/, screens/farm_details/, screens/map/ — display farm data.
//   widgets/farm_card.dart — renders farm info in cards.
class Farm {}