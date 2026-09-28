package satellite

import (
	"bytes"
	"context"
	"io"
	"math"
	"net/http"
	"testing"
	"time"

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
		WeatherDataURL: "https://kijani.example.test/data",
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

func TestFetchForFarmUsesMockDataWithoutWeatherConfiguration(t *testing.T) {
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

func TestFetchForFarmUsesOpenWeatherMapCredentialsAndMapsForecast(t *testing.T) {
	previousClient := httpClient
	httpClient = &http.Client{Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		query := r.URL.Query()
		if got := query.Get("appid"); got != "test-openweather-key" {
			t.Fatalf("appid = %q, want test-openweather-key", got)
		}
		if got := query.Get("units"); got != "metric" {
			t.Fatalf("units = %q, want metric", got)
		}
		if got := query.Get("lat"); got != "-0.091700" {
			t.Fatalf("lat = %q, want -0.091700", got)
		}
		if got := query.Get("lon"); got != "34.768000" {
			t.Fatalf("lon = %q, want 34.768000", got)
		}
		// OpenWeatherMap authenticates via `appid`; the caller's JWT and the
		// legacy API-key header must never be sent to a third party.
		if got := r.Header.Get("Authorization"); got != "" {
			t.Fatalf("Authorization = %q, want no auth header for OpenWeatherMap", got)
		}
		if got := r.Header.Get("X-API-Key"); got != "" {
			t.Fatalf("X-API-Key = %q, want no api-key header for OpenWeatherMap", got)
		}
		body := `{"cod":"200","list":[
			{"dt":1759058400,"main":{"temp":28.4},"wind":{"speed":4},"pop":0.25},
			{"dt":1759069200,"main":{"temp":31.5},"wind":{"speed":10},"pop":0.8},
			{"dt":1759155600,"main":{"temp":99.0},"wind":{"speed":50},"pop":1.0}
		]}`
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString(body)),
		}, nil
	})}
	t.Cleanup(func() { httpClient = previousClient })

	data, err := FetchForFarm(context.Background(), config.Config{
		WeatherDataURL: "https://api.openweathermap.org/data/2.5/forecast",
		WeatherAPIKey:  "test-openweather-key",
	}, models.Farm{ID: 7, Crop: "Maize", Latitude: -0.0917, Longitude: 34.768}, "mwangaza-user-jwt")
	if err != nil {
		t.Fatalf("FetchForFarm() error = %v", err)
	}

	wantTimestamp := time.Unix(1759058400, 0).UTC().Format(time.RFC3339)
	if data.FarmID != 7 || data.Source != "openweathermap" || data.Timestamp != wantTimestamp {
		t.Fatalf("unexpected metadata: %#v", data)
	}
	// Only the first 24 forecast hours count; the 27-hour entry (temp 99)
	// must be ignored.
	if data.Temperature != 31.5 {
		t.Fatalf("temperature = %v, want 31.5", data.Temperature)
	}
	if math.Abs(data.RainProbability-80) > 0.001 {
		t.Fatalf("rain probability = %v, want 80", data.RainProbability)
	}
	if math.Abs(data.WindSpeed-36) > 0.001 {
		t.Fatalf("wind speed = %v, want 36 km/h (10 m/s)", data.WindSpeed)
	}
	// OpenWeatherMap reports neither soil moisture nor NDVI; zeros mean
	// "not available" and must stay zero.
	if data.SoilMoisture != 0 || data.NDVI != 0 {
		t.Fatalf("unexpected fabricated soil/NDVI: %#v", data)
	}
}

func TestParseResponseMapsOpenWeatherCurrentWeather(t *testing.T) {
	raw := map[string]any{
		"dt":      float64(1759058400),
		"main":    map[string]any{"temp": 26.7},
		"wind":    map[string]any{"speed": 5.5},
		"weather": []any{map[string]any{"main": "Clear"}},
	}

	data := ParseResponse(raw)
	if data.Source != "openweathermap" || data.Temperature != 26.7 {
		t.Fatalf("metadata mapping failed: %#v", data)
	}
	if math.Abs(data.WindSpeed-19.8) > 0.001 {
		t.Fatalf("wind speed = %v, want 19.8 km/h (5.5 m/s)", data.WindSpeed)
	}
	wantTimestamp := time.Unix(1759058400, 0).UTC().Format(time.RFC3339)
	if data.Timestamp != wantTimestamp {
		t.Fatalf("timestamp = %q, want %q", data.Timestamp, wantTimestamp)
	}
	if data.SoilMoisture != 0 || data.NDVI != 0 || data.RainProbability != 0 {
		t.Fatalf("unexpected fabricated metrics: %#v", data)
	}
}
