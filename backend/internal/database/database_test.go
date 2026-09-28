package database

import (
	"errors"
	"path/filepath"
	"testing"

	"mwangaza/internal/models"
)

func TestLatestRecommendationBatchForFarmOrdersByPriority(t *testing.T) {
	store, err := Open(filepath.Join(t.TempDir(), "store.json"))
	if err != nil {
		t.Fatal(err)
	}
	for _, recommendation := range []models.Recommendation{
		{FarmID: 1, Type: "medium", Priority: "MEDIUM", CreatedAt: "2026-07-24T10:00:00Z"},
		{FarmID: 1, Type: "high", Priority: "HIGH", CreatedAt: "2026-07-24T10:00:00Z"},
		{FarmID: 1, Type: "old", Priority: "HIGH", CreatedAt: "2026-07-24T09:00:00Z"},
	} {
		if _, err := store.AddRecommendation(recommendation); err != nil {
			t.Fatal(err)
		}
	}
	batch := store.LatestRecommendationBatchForFarm(1)
	if len(batch) != 2 || batch[0].Type != "high" || batch[1].Type != "medium" {
		t.Fatalf("latest ordered batch = %#v", batch)
	}
}

func TestStorePersistsUsersAcrossReopen(t *testing.T) {
	path := filepath.Join(t.TempDir(), "store.json")

	store, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}

	created, err := store.CreateUser("Amina Njeri", "Amina@Example.com", "+254700000001", "hash-one")
	if err != nil {
		t.Fatal(err)
	}
	if created.ID != 1 {
		t.Fatalf("first user id = %d, want 1", created.ID)
	}

	if _, err := store.CreateUser("John Ouma", "john@example.com", "", "hash-two"); err != nil {
		t.Fatal(err)
	}

	reopened, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}

	stored, found := reopened.GetUserByEmail("amina@example.com")
	if !found {
		t.Fatal("user was not restored from disk")
	}
	if stored.Name != "Amina Njeri" || stored.Password != "hash-one" {
		t.Fatalf("restored user = %#v", stored)
	}

	// IDs must continue after the persisted maximum, not restart at 1.
	next, err := reopened.CreateUser("Grace Wanjiku", "grace@example.com", "", "hash-three")
	if err != nil {
		t.Fatal(err)
	}
	if next.ID != 3 {
		t.Fatalf("user id after reopen = %d, want 3", next.ID)
	}
}

func TestStoreRejectsDuplicateEmailCaseInsensitively(t *testing.T) {
	store, err := Open(filepath.Join(t.TempDir(), "store.json"))
	if err != nil {
		t.Fatal(err)
	}

	if _, err := store.CreateUser("Amina", "amina@example.com", "", "hash"); err != nil {
		t.Fatal(err)
	}

	if _, err := store.CreateUser("Someone Else", "AMINA@EXAMPLE.COM", "", "hash"); !errors.Is(err, ErrUserExists) {
		t.Fatalf("duplicate email error = %v, want ErrUserExists", err)
	}
}
