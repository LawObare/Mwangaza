// Package handlers (farm_handler.go) serves farm-related endpoints.
//
// Endpoints:
//   GET /api/farms      — returns all farms from the farms table
//   GET /api/farms/:id  — returns a single farm by id
//
// Each handler receives *sql.DB via a Handler struct or closure.
// Responses use utils.Success() / utils.ErrorResponse() helpers.
// Data is queried from the farms table and mapped to models.Farm.
package handlers