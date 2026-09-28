// Package middleware (auth.go) validates the JWTs issued by
// handlers.AuthHandler.
//
// Two wrappers are exposed:
//
//	Optional — attaches the caller to the request context when a valid Bearer
//	           token is present, and otherwise lets the request through
//	           anonymously so the offline demo keeps working.
//	Require  — rejects the request with 401 unless a valid token is present.
package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"mwangaza/internal/utils"
)

type contextKey string

const userIDKey contextKey = "user_id"

// AuthMiddleware signs and verifies tokens with a single shared secret.
type AuthMiddleware struct {
	secret []byte
}

func NewAuth(secret string) *AuthMiddleware {
	return &AuthMiddleware{secret: []byte(secret)}
}

// Secret exposes the signing secret so handlers can issue matching tokens.
func (m *AuthMiddleware) Secret() []byte {
	return m.secret
}

// Optional attaches the authenticated user to the request context when a valid
// Bearer token is supplied. Requests without a token continue anonymously.
func (m *AuthMiddleware) Optional(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, ok := m.authenticate(r)
		if ok {
			r = r.WithContext(context.WithValue(r.Context(), userIDKey, userID))
		}
		next.ServeHTTP(w, r)
	})
}

// Require rejects requests that do not carry a valid Bearer token.
func (m *AuthMiddleware) Require(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, ok := m.authenticate(r)
		if !ok {
			utils.Error(w, http.StatusUnauthorized, "authentication required")
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userIDKey, userID)))
	})
}

func (m *AuthMiddleware) authenticate(r *http.Request) (int, bool) {
	header := strings.TrimSpace(r.Header.Get("Authorization"))
	if header == "" {
		return 0, false
	}

	tokenString, found := bearerToken(header)
	if !found {
		return 0, false
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return m.secret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil || !token.Valid {
		return 0, false
	}

	return claimUserID(claims)
}

func bearerToken(header string) (string, bool) {
	parts := strings.Fields(header)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", false
	}
	if parts[1] == "" {
		return "", false
	}
	return parts[1], true
}

func claimUserID(claims jwt.MapClaims) (int, bool) {
	switch value := claims["user_id"].(type) {
	case float64:
		return int(value), true
	case int:
		return value, true
	case int64:
		return int(value), true
	default:
		return 0, false
	}
}

// UserIDFromContext returns the authenticated user id attached by Optional or
// Require.
func UserIDFromContext(ctx context.Context) (int, bool) {
	userID, ok := ctx.Value(userIDKey).(int)
	return userID, ok
}
