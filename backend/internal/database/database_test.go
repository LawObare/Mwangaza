package database

import (
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
