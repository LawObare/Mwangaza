// Package repositories (user_repository.go) handles user database operations.
//
// Functions:
//   CreateUser(user *models.User) error
//     — INSERT into users table
//     — Returns error if email already exists (UNIQUE constraint)
//
//   GetUserByEmail(email string) (*models.User, error)
//     — SELECT from users WHERE email = ?
//     — Returns nil if not found
//
//   GetUserByID(id int) (*models.User, error)
//     — SELECT from users WHERE id = ?
//     — Used by middleware to load user context
//
// Each function takes *sql.DB as a parameter (injected from main.go).
// The user_repository is called by services/auth/service.go.
package repositories