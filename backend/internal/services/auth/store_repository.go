package auth

import (
	"errors"

	"mwangaza/internal/database"
	"mwangaza/internal/models"
)

// StoreRepository adapts the file-backed store to the Repository interface so
// accounts can be registered and verified without an external identity
// provider.
type StoreRepository struct {
	store *database.Store
}

func NewStoreRepository(store *database.Store) *StoreRepository {
	return &StoreRepository{store: store}
}

func (r *StoreRepository) CreateUser(name, email, phone, passwordHash string) (*models.User, error) {
	user, err := r.store.CreateUser(name, email, phone, passwordHash)
	if err != nil {
		if errors.Is(err, database.ErrUserExists) {
			return nil, ErrEmailExists
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail returns (nil, nil) when no account matches, which is the
// contract Service.Register and the login handler rely on.
func (r *StoreRepository) GetUserByEmail(email string) (*models.User, error) {
	user, found := r.store.GetUserByEmail(email)
	if !found {
		return nil, nil
	}
	return &user, nil
}

func (r *StoreRepository) GetUserByID(id int) (*models.User, error) {
	user, found := r.store.GetUserByID(id)
	if !found {
		return nil, nil
	}
	return &user, nil
}
