// Package auth (password.go) handles password hashing and comparison.
//
// Functions:
//   HashPassword(password string) (string, error)
//     — Uses bcrypt with default cost to hash the password
//     — Returns the hashed string for storage
//
//   CheckPassword(password, hash string) bool
//     — Compares a plaintext password against a bcrypt hash
//     — Returns true if they match
//
// Called by:
//   service.go — HashPassword during registration
//   service.go — CheckPassword during login
package auth