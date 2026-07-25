package auth

import (
	"errors"

	"mwangaza/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailExists  = errors.New("email already registered")
	ErrUserNotFound = errors.New("user not found")
)

type Repository interface {
	CreateUser(name, email, phone, passwordHash string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(name, email, phone, password string) (*models.User, error) {
	existing, _ := s.repo.GetUserByEmail(email)
	if existing != nil {
		return nil, ErrEmailExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return s.repo.CreateUser(name, email, phone, string(hash))
}

func (s *Service) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetUserByEmail(email)
}

func (s *Service) Repository() Repository {
	return s.repo
}