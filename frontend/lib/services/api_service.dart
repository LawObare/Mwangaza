// lib/services/api_service.dart — Base HTTP client for all API calls.
//
// Responsibilities:
//   - Wrap http.Client with methods: get(), getList(), post().
//   - Prepend AppConstants.baseUrl to all endpoints.
//   - Apply AppConstants.timeout to every request.
//   - Handle JSON encode/decode.
//   - Provide a dispose() method for cleanup.
//
// Used by:
//   farm_service.dart, satellite_service.dart, recommendation_service.dart, sms_service.dart
//   — each service receives an ApiService instance and calls its methods.
class ApiService {}