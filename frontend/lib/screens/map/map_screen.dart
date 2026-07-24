// lib/screens/map/map_screen.dart — OpenStreetMap screen with farm markers.
//
// Layout:
//   Full-screen FlutterMap widget using OpenStreetMap tiles.
//   Markers at each farm's latitude/longitude.
//   Color-coded markers: green (healthy), orange (needs attention), red (urgent).
//   Tap marker → bottom sheet with farm summary + "View Details" button.
//
// Responsibilities:
//   - On init: fetch all farms from backend.
//   - Place markers for each farm.
//   - Show farm details in a bottom sheet on marker tap.
//   - Navigate to FarmDetailsScreen from the bottom sheet.
//
// Connects to:
//   services/farm_service.dart    — fetches farm locations.
//   screens/farm_details/         — navigates on "View Details" tap.
//   routes/app_routes.dart        — navigation constant.
//   pubspec.yaml                  — requires flutter_map and latlong2 packages.
class MapScreen {}