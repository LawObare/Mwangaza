package alerts

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"mwangaza/internal/config"
	"mwangaza/internal/database"
	"mwangaza/internal/models"
)

func TestRunOnceSendsHighestAlertAndRespectsCooldown(t *testing.T) {
	store := testStore(t)
	farm := createFarm(t, store, "Maize")
	runner := NewRunner(store, config.Config{
		UseMockData:           true,
		AutoAlertsSMSCooldown: time.Hour,
	}, nil)

	summary := runner.RunOnce(context.Background())
	if summary.FarmsChecked != 1 || summary.BatchesSaved != 1 || summary.SMSSent != 1 || summary.Errors != 0 {
		t.Fatalf("summary = %#v", summary)
	}
	logs := store.ListSMSLogs()
	if len(logs) != 1 {
		t.Fatalf("SMS logs = %d, want 1", len(logs))
	}
	if logs[0].FarmID != farm.ID || logs[0].Message == "" {
		t.Fatalf("SMS log = %#v", logs[0])
	}

	summary = runner.RunOnce(context.Background())
	if summary.SMSSent != 0 || summary.Skipped != 1 {
		t.Fatalf("cooldown summary = %#v", summary)
	}
	if logs := store.ListSMSLogs(); len(logs) != 1 {
		t.Fatalf("SMS logs after cooldown run = %d, want 1", len(logs))
	}
}

func TestRunOnceDoesNotSendMonitoringRecommendation(t *testing.T) {
	store := testStore(t)
	createFarm(t, store, "Coffee")
	runner := Runner{
		Store: store,
		Config: config.Config{
			UseMockData:           true,
			AutoAlertsSMSCooldown: time.Hour,
		},
		generate: func(store *database.Store, cfg config.Config, farm models.Farm) ([]models.Recommendation, models.SatelliteData, error) {
			stored, err := store.AddRecommendation(models.Recommendation{
				FarmID:     farm.ID,
				Crop:       farm.Crop,
				Type:       "monitoring",
				Message:    "Keep monitoring crop conditions.",
				Severity:   "low",
				Priority:   "LOW",
				Confidence: "LOW",
			})
			if err != nil {
				return nil, models.SatelliteData{}, err
			}
			return []models.Recommendation{stored}, models.SatelliteData{}, nil
		},
	}

	summary := runner.RunOnce(context.Background())
	if summary.FarmsChecked != 1 || summary.BatchesSaved != 1 || summary.SMSSent != 0 || summary.Skipped != 1 || summary.Errors != 0 {
		t.Fatalf("summary = %#v", summary)
	}
	if logs := store.ListSMSLogs(); len(logs) != 0 {
		t.Fatalf("SMS logs = %d, want 0", len(logs))
	}
	recommendations := store.ListRecommendations(nil)
	if len(recommendations) != 1 || recommendations[0].Type != "monitoring" {
		t.Fatalf("recommendations = %#v", recommendations)
	}
}

func testStore(t *testing.T) *database.Store {
	t.Helper()
	store, err := database.Open(filepath.Join(t.TempDir(), "store.json"))
	if err != nil {
		t.Fatal(err)
	}
	return store
}

func createFarm(t *testing.T, store *database.Store, crop string) models.Farm {
	t.Helper()
	farm, err := store.CreateFarm(models.Farm{
		Name:              crop + " Farm",
		Farmer:            "Amina",
		Phone:             "+254700000001",
		Crop:              crop,
		PreferredLanguage: "English",
	})
	if err != nil {
		t.Fatal(err)
	}
	return farm
}
