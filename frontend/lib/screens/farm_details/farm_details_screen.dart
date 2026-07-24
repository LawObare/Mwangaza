// lib/screens/farm_details/farm_details_screen.dart — Single farm detail screen.
//
// Layout:
//   AppBar with farm name.
//   Card with farm info (name, farmer, status, coordinates).
//   Grid of StatCards for soil moisture and temperature.
//   List of RecommendationCards for this farm.
//
// Responsibilities:
//   - Accept farmId as a constructor parameter.
//   - On init: fetch farm details and farm-specific recommendations in parallel.
//   - Show LoadingWidget while loading.
//   - Pull-to-refresh.
//
// Connects to:
//   services/farm_service.dart             — fetches single farm.
//   services/recommendation_service.dart   — fetches recommendations filtered by farm_id.
//   widgets/stat_card.dart                 — displays stats.
//   widgets/recommendation_card.dart       — displays recommendations.
//   widgets/loading_widget.dart            — loading state.
class FarmDetailsScreen {}