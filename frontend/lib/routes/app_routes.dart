// lib/routes/app_routes.dart — Route definitions and generator.
//
// Route names:
//   /               → DashboardScreen
//   /farms          → FarmsScreen
//   /farms/details  → FarmDetailsScreen (arguments: int farmId)
//   /map            → MapScreen
//   /sms            → SmsScreen
//
// Responsibilities:
//   - Define static route name constants.
//   - Implement generateRoute() with MaterialPageRoute for each screen.
//   - Pass arguments (e.g., farmId) from Navigator.pushNamed.
//
// Used by:
//   app.dart — onGenerateRoute references AppRoutes.generateRoute.
//   Every screen that calls Navigator.pushNamed() uses these constants.
class AppRoutes {}