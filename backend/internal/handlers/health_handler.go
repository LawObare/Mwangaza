// Package handlers (health_handler.go) provides a simple health check.
//
// Endpoint:
//   GET /api/health — returns { "status": "ok" }
//
// Used by the Flutter app to verify backend reachability.
// Also useful for load balancers and deployment health probes.
package handlers