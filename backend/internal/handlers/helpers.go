package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"mwangaza/internal/database"
	"mwangaza/internal/models"
)

func parseQueryFarmID(value string) (int, bool, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return 0, false, nil
	}

	id, err := strconv.Atoi(trimmed)
	if err != nil {
		return 0, false, fmt.Errorf("invalid farm_id")
	}

	return id, true, nil
}

func enrichFarm(store *database.Store, farm models.Farm) models.Farm {
	if sat, ok := store.LatestSatelliteForFarm(farm.ID); ok {
		farm.SoilMoisture = sat.SoilMoisture
		farm.Temperature = sat.Temperature
		farm.RainProbability = sat.RainProbability
		farm.NDVI = sat.NDVI
		farm.WindSpeed = sat.WindSpeed
		if farm.Status == "" || farm.Status == "new" {
			farm.Status = statusFromSatellite(sat)
		}
	}

	if rec, ok := store.LatestRecommendationForFarm(farm.ID); ok {
		farm.Status = statusFromSeverity(rec.Severity)
	} else if farm.Status == "" || farm.Status == "new" {
		if farm.SoilMoisture > 0 || farm.Temperature > 0 || farm.RainProbability > 0 || farm.NDVI > 0 {
			farm.Status = statusFromSatellite(models.SatelliteData{
				SoilMoisture:    farm.SoilMoisture,
				Temperature:     farm.Temperature,
				RainProbability: farm.RainProbability,
				NDVI:            farm.NDVI,
				WindSpeed:       farm.WindSpeed,
			})
		} else {
			farm.Status = "unknown"
		}
	}

	if smsLog, ok := store.LatestSMSForFarm(farm.ID); ok {
		farm.LastSmsSent = smsLog.SentAt
	}

	return farm
}

func statusFromSeverity(severity string) string {
	switch strings.ToLower(strings.TrimSpace(severity)) {
	case "high":
		return "urgent"
	case "medium":
		return "needs_attention"
	case "low":
		return "healthy"
	default:
		return "unknown"
	}
}

func statusFromSatellite(data models.SatelliteData) string {
	switch {
	case data.SoilMoisture < 15 || data.RainProbability >= 85 || data.Temperature >= 35 || data.NDVI < 0.25:
		return "urgent"
	case data.SoilMoisture < 25 || data.RainProbability >= 70 || data.Temperature >= 32 || data.NDVI < 0.35:
		return "needs_attention"
	default:
		return "healthy"
	}
}
