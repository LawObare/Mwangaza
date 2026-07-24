// Package routes (auth_routes.go) registers authentication endpoints.
//
// These routes are public (no JWT middleware):
//   POST /api/auth/register — auth_handler.Register
//   POST /api/auth/login    — auth_handler.Login
//
// Called from routes.Setup() to add auth routes to the /api group.
// Register before protected routes so the middleware chain is clean.
package routes