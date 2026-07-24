// Package utils (response.go) provides standardized JSON response helpers.
//
// Every API handler uses these to return consistent response shapes:
//
// Success response:
//   { "success": true, "data": { ... } }
//
// Error response:
//   { "success": false, "error": "message" }
//
// Functions:
//   Success(c, data)       — 200 OK
//   Created(c, data)       — 201 Created
//   BadRequest(c, msg)     — 400 Bad Request
//   NotFound(c, msg)       — 404 Not Found
//   InternalError(c, msg)  — 500 Internal Server Error
package utils