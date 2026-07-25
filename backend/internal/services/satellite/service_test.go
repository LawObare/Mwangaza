package satellite

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"mwangaza/internal/config"
	"mwangaza/internal/models"
)

func TestFetchForFarmUsesKijaniTokenAndMapsResponse(t *testing.T) {
	previousClient := httpClient
	httpClient = &http.Client{Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		if got := r.Header.Get("Authorization"); got != "Bearer kijani-user-token" {
			t.Fatalf("Authorization = %q, want Kijani bearer token", got)
		}
		if got := r.URL.Query().Get("lat"); got != "-0.091700" {
			t.Fatalf("lat = %q, want -0.091700", got)
		}
		if got := r.URL.Query().Get("lon"); got != "34.768000" {
			t.Fatalf("lon = %q, want 34.768000", got)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString(`{"data":{"soil_moisture":27.5,"temperature":29,"rain_probability":61,"ndvi":0.71,"wind_speed":3.2}}`)),
		}, nil
	})}
	t.Cleanup(func() { httpClient = previousClient })

	data, err := FetchForFarm(context.Background(), config.Config{
		KijaniDataURL: "https://kijani.example.test/data",
	}, models.Farm{ID: 7, Crop: "Maize", Latitude: -0.0917, Longitude: 34.768}, "kijani-user-token")
	if err != nil {
		t.Fatalf("FetchForFarm() error = %v", err)
	}
	if data.FarmID != 7 || data.Source != "kijani" || data.NDVI != 0.71 {
		t.Fatalf("unexpected mapped data: %#v", data)
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return fn(request)
}

func TestFetchForFarmUsesMockDataWithoutKijaniConfiguration(t *testing.T) {
	data, err := FetchForFarm(context.Background(), config.Config{}, models.Farm{ID: 1, Crop: "Maize"}, "")
	if err != nil {
		t.Fatalf("FetchForFarm() error = %v", err)
	}
	if data.Source != "mock" || data.FarmID != 1 {
		t.Fatalf("unexpected mock data: %#v", data)
	}
}

func TestParseResponseMapsKijaniAgroClimateForecast(t *testing.T) {
	raw := map[string]any{
		"source": "meteoblue",
		"data": map[string]any{
			"vegetation_indices": map[string]any{"NDVI": 0.32},
		},
		"forecast_data": map[string]any{
			"time":                      []any{"2026-07-25 00:00"},
			"temperature":               []any{20.4, 28.9, 25.0},
			"precipitation_probability": []any{15.0, 72.0, 38.0},
			"windspeed":                 []any{2.1, 5.4, 4.2},
			"soilmoisture_0to10cm":      []any{24.0, 18.5, 21.0},
		},
	}

	data := ParseResponse(raw)
	if data.Source != "meteoblue" || data.Timestamp != "2026-07-25 00:00" {
		t.Fatalf("metadata mapping failed: %#v", data)
	}
	if data.Temperature != 28.9 || data.RainProbability != 72 || data.WindSpeed != 5.4 || data.SoilMoisture != 18.5 || data.NDVI != 0.32 {
		t.Fatalf("forecast mapping failed: %#v", data)
	}
}
