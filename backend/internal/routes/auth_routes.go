// Package routes (auth_routes.go) registers the local authentication endpoints
// on the /api group.
//
//	POST /api/auth/register - create an account, returns a JWT
//	POST /api/auth/login    - exchange credentials for a JWT
//	GET  /api/auth/me       - current account (requires a valid Bearer token)
//
// Called from routes.Setup(); kept in its own file so the whole auth surface
// can be reviewed in one place.
package routes

import (
	"net/http"

	"mwangaza/internal/handlers"
	"mwangaza/internal/middleware"
)

func RegisterAuthRoutes(mux *http.ServeMux, handler *handlers.AuthHandler, auth *middleware.AuthMiddleware) {
	mux.HandleFunc("POST /api/auth/register", handler.Register)
	mux.HandleFunc("POST /api/auth/login", handler.Login)
	mux.Handle("GET /api/auth/me", auth.Require(http.HandlerFunc(handler.Me)))
}
