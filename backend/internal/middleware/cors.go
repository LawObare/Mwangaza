// Package middleware (cors.go) provides CORS configuration for the Flutter frontend.
//
// The Flutter dashboard may run on a different origin (e.g., web build on localhost:3000
// or an emulator on 10.0.2.2:8080). This middleware allows all origins with standard
// methods (GET, POST, PUT, DELETE, OPTIONS) and headers (Content-Type, Authorization).
//
// Registered globally in routes.Setup() via r.Use().
// Preflight OPTIONS requests are handled with a 204 No Content response.
package middleware