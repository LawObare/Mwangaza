// Package auth (service.go) is the public entry point for authentication logic.
//
// It coordinates registration and login workflows:
//
// Register(name, email, phone, password) → (User, error)
//   1. Validate input fields
//   2. Check email uniqueness via user_repository
//   3. Hash password via password.go
//   4. Store user via user_repository
//   5. Return created user (without password)
//
// Login(email, password) → (token string, User, error)
//   1. Lookup user by email via user_repository
//   2. Compare password hash via password.go
//   3. Generate JWT via jwt.go
//   4. Return token and user
//
// The handler at handlers/auth_handler.go calls this service.
package auth