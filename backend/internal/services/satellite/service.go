package satellite

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"mwangaza/internal/config"
	"mwangaza/internal/database"
	"mwangaza/internal/models"
)

var httpClient = http.DefaultClient

func FetchForFarm(ctx context.Context, cfg config.Config, farm models.Farm, accessToken string) (models.SatelliteData, error) {
	if cfg.UseMockData || strings.TrimSpace(cfg.KijaniDataURL) == "" {
		return MockForFarm(farm), nil
	}

	endpoint, err := buildEndpoint(cfg.KijaniDataURL, farm)
	if err != nil {
		return MockForFarm(farm), nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return MockForFarm(farm), nil
	}

	if accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+accessToken)
	} else if cfg.KijaniAPIKey != "" {
		req.Header.Set("X-API-Key", cfg.KijaniAPIKey)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return MockForFarm(farm), nil
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return MockForFarm(farm), nil
	}

	var raw any
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return MockForFarm(farm), nil
	}

	data := ParseResponse(raw)
	data.FarmID = farm.ID
	if data.Timestamp == "" {
		data.Timestamp = time.Now().UTC().Format(time.RFC3339)
	}
	if data.Source == "" {
		data.Source = "kijani"
	}

	return data, nil
}

func FetchAndStore(store *database.Store, cfg config.Config, farm models.Farm, accessToken string) (models.SatelliteData, error) {
	data, err := FetchForFarm(context.Background(), cfg, farm, accessToken)
	if err != nil {
		return models.SatelliteData{}, err
	}

	stored, err := store.AddSatelliteData(data)
	if err != nil {
		return models.SatelliteData{}, err
	}

	return stored, nil
}

func ParseResponse(raw any) models.SatelliteData {
	if root, ok := raw.(map[string]any); ok {
		if _, ok := root["forecast_data"].(map[string]any); ok {
			return parseKijaniAgroClimate(root)
		}
	}

	payload := unwrapMap(raw)

	return models.SatelliteData{
		SoilMoisture:    floatValue(payload, "soil_moisture", "soilMoisture", "soil"),
		Temperature:     floatValue(payload, "temperature", "temp"),
		RainProbability: floatValue(payload, "rain_probability", "rainProbability", "rain_prob"),
		NDVI:            floatValue(payload, "ndvi"),
		WindSpeed:       floatValue(payload, "wind_speed", "windSpeed"),
		Source:          stringValue(payload, "source", "provider"),
		Timestamp:       stringValue(payload, "timestamp", "recorded_at", "recordedAt"),
	}
}

// parseKijaniAgroClimate converts Kijani's hourly five-day forecast into the
// single risk snapshot used by Mwangaza's recommendation rules. We assess the
// first 24 forecast hours: maximum heat/rain/wind, minimum soil moisture, and
// the latest satellite-derived NDVI.
func parseKijaniAgroClimate(payload map[string]any) models.SatelliteData {
	forecast, _ := payload["forecast_data"].(map[string]any)
	data, _ := payload["data"].(map[string]any)
	vegetation, _ := data["vegetation_indices"].(map[string]any)

	return models.SatelliteData{
		SoilMoisture:    minForecastValue(forecast, "soilmoisture_0to10cm"),
		Temperature:     maxForecastValue(forecast, "temperature"),
		RainProbability: maxForecastValue(forecast, "precipitation_probability"),
		NDVI:            floatValue(vegetation, "NDVI", "ndvi"),
		WindSpeed:       maxForecastValue(forecast, "windspeed", "wind_speed"),
		Source:          stringValue(payload, "source"),
		Timestamp:       firstForecastTime(forecast),
	}
}

func forecastWindow(payload map[string]any, key string) []float64 {
	values, ok := payload[key].([]any)
	if !ok {
		return nil
	}
	limit := len(values)
	if limit > 24 {
		limit = 24
	}
	result := make([]float64, 0, limit)
	for _, value := range values[:limit] {
		if number, ok := numericValue(value); ok {
			result = append(result, number)
		}
	}
	return result
}

func maxForecastValue(payload map[string]any, keys ...string) float64 {
	var max float64
	found := false
	for _, key := range keys {
		for _, value := range forecastWindow(payload, key) {
			if !found || value > max {
				max, found = value, true
			}
		}
	}
	return max
}

func minForecastValue(payload map[string]any, keys ...string) float64 {
	var min float64
	found := false
	for _, key := range keys {
		for _, value := range forecastWindow(payload, key) {
			if !found || value < min {
				min, found = value, true
			}
		}
	}
	return min
}

func firstForecastTime(payload map[string]any) string {
	values, ok := payload["time"].([]any)
	if !ok || len(values) == 0 {
		return ""
	}
	value, _ := values[0].(string)
	return value
}

func MockForFarm(farm models.Farm) models.SatelliteData {
	now := time.Now().UTC().Format(time.RFC3339)
	crop := strings.ToLower(strings.TrimSpace(farm.Crop))

	data := models.SatelliteData{
		FarmID:    farm.ID,
		Source:    "mock",
		Timestamp: now,
	}

	switch crop {
	case "maize":
		data.SoilMoisture = 18
		data.Temperature = 33
		data.RainProbability = 12
		data.NDVI = 0.42
		data.WindSpeed = 4
	case "rice":
		data.SoilMoisture = 68
		data.Temperature = 27
		data.RainProbability = 84
		data.NDVI = 0.64
		data.WindSpeed = 5
	case "beans":
		data.SoilMoisture = 24
		data.Temperature = 29
		data.RainProbability = 76
		data.NDVI = 0.48
		data.WindSpeed = 3
	case "tea":
		data.SoilMoisture = 22
		data.Temperature = 34
		data.RainProbability = 35
		data.NDVI = 0.29
		data.WindSpeed = 6
	case "sugarcane":
		data.SoilMoisture = 19
		data.Temperature = 31
		data.RainProbability = 18
		data.NDVI = 0.44
		data.WindSpeed = 4
	case "cassava":
		data.SoilMoisture = 21
		data.Temperature = 32
		data.RainProbability = 28
		data.NDVI = 0.33
		data.WindSpeed = 5
	default:
		profiles := []models.SatelliteData{
			{SoilMoisture: 18, Temperature: 33, RainProbability: 12, NDVI: 0.42, WindSpeed: 4},
			{SoilMoisture: 68, Temperature: 27, RainProbability: 84, NDVI: 0.64, WindSpeed: 5},
			{SoilMoisture: 24, Temperature: 29, RainProbability: 76, NDVI: 0.48, WindSpeed: 3},
		}
		index := 0
		if farm.ID > 0 {
			index = (farm.ID - 1) % len(profiles)
		}
		data.SoilMoisture = profiles[index].SoilMoisture
		data.Temperature = profiles[index].Temperature
		data.RainProbability = profiles[index].RainProbability
		data.NDVI = profiles[index].NDVI
		data.WindSpeed = profiles[index].WindSpeed
	}

	return data
}

func buildEndpoint(base string, farm models.Farm) (string, error) {
	parsed, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	query := parsed.Query()
	// Kijani's published agro-climate endpoint requires WGS84 `lat` and
	// `lon`. Farm ID and crop remain local Mwangaza metadata; they are not
	// query parameters in Kijani's API contract.
	query.Set("lat", fmt.Sprintf("%.6f", farm.Latitude))
	query.Set("lon", fmt.Sprintf("%.6f", farm.Longitude))
	parsed.RawQuery = query.Encode()
	return parsed.String(), nil
}

func unwrapMap(raw any) map[string]any {
	if payload, ok := raw.(map[string]any); ok {
		if nested, ok := payload["data"].(map[string]any); ok {
			return nested
		}
		return payload
	}
	return map[string]any{}
}

func floatValue(payload map[string]any, keys ...string) float64 {
	for _, key := range keys {
		if value, ok := payload[key]; ok {
			if parsed, ok := numericValue(value); ok {
				return parsed
			}
		}
	}
	return 0
}

func numericValue(value any) (float64, bool) {
	switch typed := value.(type) {
	case float64:
		return typed, true
	case int:
		return float64(typed), true
	case int32:
		return float64(typed), true
	case int64:
		return float64(typed), true
	case json.Number:
		parsed, err := typed.Float64()
		return parsed, err == nil
	case string:
		parsed, err := strconv.ParseFloat(strings.TrimSpace(typed), 64)
		return parsed, err == nil
	default:
		return 0, false
	}
}

func stringValue(payload map[string]any, keys ...string) string {
	for _, key := range keys {
		if value, ok := payload[key]; ok {
			switch typed := value.(type) {
			case string:
				return typed
			case json.Number:
				return typed.String()
			case float64:
				return strconv.FormatFloat(typed, 'f', -1, 64)
			case int:
				return strconv.Itoa(typed)
			case int64:
				return strconv.FormatInt(typed, 10)
			}
		}
	}
	return ""
}
