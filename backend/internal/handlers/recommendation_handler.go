// Package handlers (recommendation_handler.go) serves recommendation endpoints.
//
// Endpoints:
//   GET /api/recommendation          — returns all recommendations
//   GET /api/recommendation?farm_id= — filters by farm_id
//
// Queries the recommendations table. Each row includes a message,
// priority level (low/medium/high), reason, and timestamp.
// The recommendation engine (services/recommendation) populates this table.
package handlers