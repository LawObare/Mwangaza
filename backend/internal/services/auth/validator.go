// Package auth (validator.go) validates registration and login input.
//
// Functions:
//   ValidateRegistration(name, email, phone, password string) error
//     — Checks: name not empty, email format, phone format (+254...), password min length (8)
//     — Returns descriptive error messages for each invalid field
//
//   ValidateLogin(email, password string) error
//     — Checks: email format, password not empty
//
// Called by service.go before any database operations.
// Keeps input validation isolated from business logic.
package auth