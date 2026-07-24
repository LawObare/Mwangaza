// lib/screens/farms/farms_screen.dart — Farm list screen.
//
// Layout:
//   AppBar with "Farms" title.
//   ListView of FarmCard widgets.
//
// Responsibilities:
//   - On init: fetch list of farms from backend.
//   - Show LoadingWidget while loading.
//   - Pull-to-refresh.
//   - Tap a farm card → navigate to FarmDetailsScreen with farm.id.
//
// Connects to:
//   services/farm_service.dart  — fetches farm list.
//   widgets/farm_card.dart       — renders each farm.
//   widgets/loading_widget.dart  — loading state.
//   screens/farm_details/        — navigates here on tap.
//   routes/app_routes.dart       — navigation constant.
class FarmsScreen {}