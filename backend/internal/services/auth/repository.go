package auth

import (
	"database/sql"
	"time"

	"mwangaza/internal/models"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) CreateUser(name, email, phone, passwordHash string) (*models.User, error) {
	result, err := r.db.Exec(
		"INSERT INTO users (name, email, phone, password_hash) VALUES (?, ?, ?, ?)",
		name, email, phone, passwordHash,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:        int(id),
		Name:      name,
		Email:     email,
		Phone:     phone,
		CreatedAt: time.Now().Format(time.RFC3339),
	}, nil
}

func (r *SQLiteRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	var passwordHash string
	err := r.db.QueryRow(
		"SELECT id, name, email, phone, password_hash, created_at FROM users WHERE email = ?",
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &passwordHash, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return &user, nil
}

func (r *SQLiteRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	var passwordHash string
	err := r.db.QueryRow(
		"SELECT id, name, email, phone, password_hash, created_at FROM users WHERE id = ?",
		id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &passwordHash, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return &user, nil
}