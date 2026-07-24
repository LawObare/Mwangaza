// Package auth (jwt.go) handles JWT token generation and validation.
//
// Functions:
//   GenerateToken(userID int, email string) (string, error)
//     — Creates a signed JWT with user claims and expiration (e.g., 24h)
//     — Uses a secret key from JWT_SECRET env var
//
//   ValidateToken(tokenString string) (claims, error)
//     — Parses and validates the JWT
//     — Returns user ID and email from claims
//     — Used by middleware/auth.go to protect routes
//
// The JWT secret should be at least 32 characters and set in .env.
// If not set, a hardcoded fallback can be used for development only.
package auth