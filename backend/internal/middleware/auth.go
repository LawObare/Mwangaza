package middleware

import "net/http"

// Auth is a placeholder JWT middleware for protected routes.
// Currently passes all requests through. TODO: extract Bearer token,
// validate JWT, and set user context.
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: extract Bearer token, validate JWT, set user context
		next.ServeHTTP(w, r)
	})
}