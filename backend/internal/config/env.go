// Package config (env.go) provides helper functions for reading
// individual environment variables used across the backend.
//
// Functions:
//   IsMockEnabled() bool       — returns true if USE_MOCK_DATA == "true"
//   SpaceIoTBoxAPIKey() string — reads SPACEIOTBOX_API_KEY
//   SpaceIoTBoxBaseURL() string — reads SPACEIOTBOX_BASE_URL
//
// These are used by the satellite service (client.go, service.go)
// and the SMS service (service.go) to decide mock vs live mode.
package config