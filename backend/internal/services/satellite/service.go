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

func FetchForFarm(ctx context.Context, cfg config.Config, farm models.Farm) (models.SatelliteData, error) {
	if cfg.UseMockData || strings.TrimSpace(cfg.SpaceIoTBoxBaseURL) == "" {
		return MockForFarm(farm), nil
	}

	endpoint, err := buildEndpoint(cfg.SpaceIoTBoxBaseURL, farm)
	if err != nil {
		return MockForFarm(farm), nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return MockForFarm(farm), nil
	}

	if cfg.SpaceIoTBoxAPIKey != "" {
		req.Header.Set("Authorization", "Bearer "+cfg.SpaceIoTBoxAPIKey)
		req.Header.Set("X-API-Key", cfg.SpaceIoTBoxAPIKey)
	}

	resp, err := http.DefaultClient.Do(req)
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
		data.Source = "spaceiotbox"
	}

	return data, nil
}

func FetchAndStore(store *database.Store, cfg config.Config, farm models.Farm) (models.SatelliteData, error) {
	data, err := FetchForFarm(context.Background(), cfg, farm)
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
	if farm.ID > 0 {
		query.Set("farm_id", strconv.Itoa(farm.ID))
	}
	if farm.Latitude != 0 {
		query.Set("latitude", fmt.Sprintf("%.6f", farm.Latitude))
	}
	if farm.Longitude != 0 {
		query.Set("longitude", fmt.Sprintf("%.6f", farm.Longitude))
	}
	if farm.Crop != "" {
		query.Set("crop", farm.Crop)
	}
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
			switch typed := value.(type) {
			case float64:
				return typed
			case int:
				return float64(typed)
			case int32:
				return float64(typed)
			case int64:
				return float64(typed)
			case json.Number:
				if parsed, err := typed.Float64(); err == nil {
					return parsed
				}
			case string:
				if parsed, err := strconv.ParseFloat(strings.TrimSpace(typed), 64); err == nil {
					return parsed
				}
			}
		}
	}
	return 0
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
