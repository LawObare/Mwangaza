package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"mwangaza/internal/config"
	"mwangaza/internal/database"
	"mwangaza/internal/models"
)

func TestRecommendationGenerateReturnsArrays(t *testing.T) {
	store := testStore(t)
	farmOne := createTestFarm(t, store, "Maize")
	farmTwo := createTestFarm(t, store, "Rice")
	handler := NewRecommendationHandler(store, config.Config{UseMockData: true})

	t.Run("one farm", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		handler.Generate(recorder, httptest.NewRequest(http.MethodPost, "/api/recommendation?farm_id="+itoa(farmOne.ID), nil))
		if recorder.Code != http.StatusCreated {
			t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
		}
		var response struct {
			Data []models.Recommendation `json:"data"`
		}
		if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
			t.Fatal(err)
		}
		if len(response.Data) == 0 || response.Data[0].Confidence == "" {
			t.Fatalf("expected recommendation array with confidence: %#v", response.Data)
		}
	})

	t.Run("all farms", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		handler.Generate(recorder, httptest.NewRequest(http.MethodPost, "/api/recommendation", nil))
		if recorder.Code != http.StatusCreated {
			t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
		}
		var response struct {
			Data map[string][]models.Recommendation `json:"data"`
		}
		if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
			t.Fatal(err)
		}
		if len(response.Data[itoa(farmOne.ID)]) == 0 || len(response.Data[itoa(farmTwo.ID)]) == 0 {
			t.Fatalf("expected per-farm recommendation arrays: %#v", response.Data)
		}
	})
}

func TestSMSHandlerUsesHighestPriorityFromLatestBatch(t *testing.T) {
	store := testStore(t)
	farm := createTestFarm(t, store, "Maize")
	for _, recommendation := range []models.Recommendation{
		{FarmID: farm.ID, Message: "Monitor conditions", Priority: "MEDIUM", Severity: "medium", CreatedAt: "2026-07-24T10:00:00Z"},
		{FarmID: farm.ID, Message: "Urgent irrigation", Priority: "HIGH", Severity: "high", CreatedAt: "2026-07-24T10:00:00Z"},
	} {
		if _, err := store.AddRecommendation(recommendation); err != nil {
			t.Fatal(err)
		}
	}

	recorder := httptest.NewRecorder()
	NewSMSHandler(store, config.Config{UseMockData: true}).Send(recorder, httptest.NewRequest(http.MethodPost, "/api/sms/send", stringsReader(`{"farm_id":1}`)))
	if recorder.Code != http.StatusCreated {
		t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
	}
	var response struct {
		Data models.SmsMessage `json:"data"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	if response.Data.Message != "Urgent irrigation" {
		t.Fatalf("SMS message = %q, want highest-priority recommendation", response.Data.Message)
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

func createTestFarm(t *testing.T, store *database.Store, crop string) models.Farm {
	t.Helper()
	farm, err := store.CreateFarm(models.Farm{Name: crop + " Farm", Farmer: "Amina", Phone: "+254700000001", Crop: crop, PreferredLanguage: "English"})
	if err != nil {
		t.Fatal(err)
	}
	return farm
}

func itoa(value int) string                      { return fmt.Sprintf("%d", value) }
func stringsReader(value string) *strings.Reader { return strings.NewReader(value) }
