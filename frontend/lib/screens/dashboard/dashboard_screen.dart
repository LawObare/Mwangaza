// lib/screens/dashboard/dashboard_screen.dart — Main dashboard screen.
//
// Layout:
//   AppBar with "Mwangaza Dashboard" title.
//   Drawer with navigation to Farms, Map, SMS History.
//   Body: grid of StatCard widgets (soil moisture, temperature, rain, NDVI),
//         SatelliteCard with full data, and top 3 RecommendationCards.
//
// Responsibilities:
//   - On init: load satellite data and recommendations in parallel.
//   - Show LoadingWidget while data loads.
//   - Pull-to-refresh support via RefreshIndicator.
//   - Dispose ApiService when the screen is disposed.
//
// Connects to:
//   services/satellite_service.dart       — fetches satellite data.
//   services/recommendation_service.dart  — fetches recommendations.
//   widgets/stat_card.dart                — displays individual stats.
//   widgets/satellite_card.dart           — displays full satellite data.
//   widgets/recommendation_card.dart      — displays each recommendation.
//   widgets/loading_widget.dart           — loading state.
//   routes/app_routes.dart                — drawer navigation.
class DashboardScreen {}