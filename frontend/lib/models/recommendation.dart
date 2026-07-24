// lib/models/recommendation.dart — Recommendation data model.
//
// Fields:
//   int farmId, String type, String message, String severity, String createdAt
//
// Responsibilities:
//   - Define a Recommendation class with a fromJson() factory constructor.
//   - JSON keys match the Go backend GET /api/recommendation response.
//
// Used by:
//   services/recommendation_service.dart — deserializes API response.
//   screens/dashboard/, screens/farm_details/ — shows recommendations.
//   widgets/recommendation_card.dart — renders each recommendation.
class Recommendation {}