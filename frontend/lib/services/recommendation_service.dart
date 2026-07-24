// lib/services/recommendation_service.dart — Recommendation API service.
//
// Endpoints:
//   GET /api/recommendation           → getRecommendations() returns List<Recommendation>
//   GET /api/recommendation?farm_id=  → getRecommendationsForFarm(id) returns List<Recommendation>
//
// Responsibilities:
//   - Use ApiService to call the backend recommendation endpoint.
//   - Support optional farm_id filtering.
//   - Deserialize JSON into models/recommendation.dart.
//
// Connects to:
//   api_service.dart             — makes HTTP requests.
//   models/recommendation.dart   — target model.
//   screens/dashboard/, screens/farm_details/ — consumers.
class RecommendationService {}