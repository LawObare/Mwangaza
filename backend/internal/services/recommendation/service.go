package recommendation

import (
	"fmt"
	"strings"
	"time"

	"mwangaza/internal/config"
	"mwangaza/internal/database"
	"mwangaza/internal/models"
	"mwangaza/internal/services/satellite"
)

func GenerateForFarm(store *database.Store, cfg config.Config, farm models.Farm) (models.Recommendation, models.SatelliteData, error) {
	data, err := satellite.FetchAndStore(store, cfg, farm)
	if err != nil {
		return models.Recommendation{}, models.SatelliteData{}, err
	}

	rec := GenerateRecommendation(farm, data)
	stored, err := store.AddRecommendation(rec)
	if err != nil {
		return models.Recommendation{}, data, err
	}

	return stored, data, nil
}

func GenerateRecommendation(farm models.Farm, data models.SatelliteData) models.Recommendation {
	crop := strings.ToLower(strings.TrimSpace(farm.Crop))
	now := time.Now().UTC().Format(time.RFC3339)

	recommendation := models.Recommendation{
		FarmID:    farm.ID,
		Crop:      farm.Crop,
		Type:      "monitoring",
		Message:   "Conditions are stable. Continue monitoring the field.",
		Severity:  "low",
		Priority:  "LOW",
		Reason:    "No critical threshold was triggered.",
		CreatedAt: now,
	}

	switch {
	case data.RainProbability >= 80:
		recommendation.Type = "rainfall"
		recommendation.Severity = "high"
		recommendation.Priority = "HIGH"
		switch crop {
		case "maize":
			recommendation.Message = "Heavy rainfall expected soon. Delay fertilizer application until after the rain."
		case "beans":
			recommendation.Message = "Heavy rainfall expected soon. Delay planting to avoid seed loss."
		case "rice":
			recommendation.Message = "Heavy rainfall expected soon. Check drainage and keep water levels under control."
		case "sugarcane":
			recommendation.Message = "Heavy rainfall expected soon. Suspend irrigation and protect field access."
		default:
			recommendation.Message = "Heavy rainfall expected soon. Delay field operations until conditions improve."
		}
		recommendation.Reason = fmt.Sprintf("rain probability %.0f%% is above the heavy-rain threshold", data.RainProbability)
	case data.SoilMoisture < 20 && data.RainProbability < 30:
		recommendation.Type = "irrigation"
		recommendation.Severity = "high"
		recommendation.Priority = "HIGH"
		switch crop {
		case "rice":
			recommendation.Message = "Water levels are low. Refill paddies or check irrigation canals."
		case "tea":
			recommendation.Message = "Soil is dry. Irrigate or increase moisture monitoring today."
		case "cassava":
			recommendation.Message = "Dry soil detected. Water young plants soon to reduce stress."
		default:
			recommendation.Message = "Dry soil detected. Irrigate today if water is available."
		}
		recommendation.Reason = fmt.Sprintf("soil moisture %.0f%% is below the dry threshold and rain chance is low", data.SoilMoisture)
	case data.Temperature >= 33:
		recommendation.Type = "temperature"
		recommendation.Severity = "medium"
		recommendation.Priority = "MEDIUM"
		switch crop {
		case "tea":
			recommendation.Message = "High temperature detected. Increase moisture monitoring and shade sensitive plots."
		case "maize":
			recommendation.Message = "High temperature detected. Monitor for heat stress and irrigate in the evening."
		default:
			recommendation.Message = "High temperature detected. Check water availability and monitor crop stress."
		}
		recommendation.Reason = fmt.Sprintf("temperature %.1fC is above the heat threshold", data.Temperature)
	case data.NDVI < 0.35:
		recommendation.Type = "ndvi"
		recommendation.Severity = "medium"
		recommendation.Priority = "MEDIUM"
		recommendation.Message = "Vegetation health looks weak. Inspect the field for pests, disease, or nutrient stress."
		recommendation.Reason = fmt.Sprintf("NDVI %.2f is below the vegetation-health threshold", data.NDVI)
	case data.RainProbability >= 55:
		recommendation.Type = "rainfall"
		recommendation.Severity = "medium"
		recommendation.Priority = "MEDIUM"
		recommendation.Message = "Rain is likely soon. Prepare field drains and avoid spraying before the rain."
		recommendation.Reason = fmt.Sprintf("rain probability %.0f%% suggests rain is coming", data.RainProbability)
	}

	return recommendation
}
