package database

import (
	"os"
	"path/filepath"
)

func Migrate(store *Store) error {
	if store == nil {
		return nil
	}

	store.mu.Lock()
	defer store.mu.Unlock()

	if store.path == "" {
		store.path = "./data/lakenet.db"
	}

	return os.MkdirAll(filepath.Dir(store.path), 0o755)
}
