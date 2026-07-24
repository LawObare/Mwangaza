package models

// User represents an admin account.
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}