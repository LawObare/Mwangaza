package recommendation

import (
	"time"

	"mwangaza/internal/config"
	"mwangaza/internal/database"
	"mwangaza/internal/models"
	"mwangaza/internal/services/satellite"
)

func GenerateForFarm(store *database.Store, cfg config.Config, farm models.Farm) ([]models.Recommendation, models.SatelliteData, error) {
	return GenerateForFarmWithToken(store, cfg, farm, "")
}

// GenerateForFarmWithToken uses the logged-in Kijani user's token when it
// refreshes environmental data. Background jobs call GenerateForFarm instead
// and rely on the configured server API key.
func GenerateForFarmWithToken(store *database.Store, cfg config.Config, farm models.Farm, accessToken string) ([]models.Recommendation, models.SatelliteData, error) {
	data, err := satellite.FetchAndStore(store, cfg, farm, accessToken)
	if err != nil {
		return nil, models.SatelliteData{}, err
	}

	recommendations := GenerateRecommendations(farm, data)
	stored := make([]models.Recommendation, 0, len(recommendations))
	for _, recommendation := range recommendations {
		recommendation, err = store.AddRecommendation(recommendation)
		if err != nil {
			return nil, data, err
		}
		stored = append(stored, recommendation)
	}

	return stored, data, nil
}

// GenerateRecommendations evaluates every applicable rule and returns a single
// monitoring result only when the farm does not need an alert.
func GenerateRecommendations(farm models.Farm, data models.SatelliteData) []models.Recommendation {
	now := time.Now().UTC().Format(time.RFC3339)
	recommendations := make([]models.Recommendation, 0, 7)
	for _, rule := range []rule{
		droughtStressRule, heatStressRule, heavyRainRule, rainfallRule,
		irrigationRule, temperatureRule, ndviRule, windRule,
	} {
		if recommendation := rule(farm, data, now); recommendation != nil {
			recommendations = append(recommendations, *recommendation)
		}
	}
	if len(recommendations) == 0 {
		recommendations = append(recommendations, monitoringRecommendation(farm, now))
	}
	return recommendations
}

// GenerateRecommendation remains as a compatibility helper for callers that
// need the most important current result.
func GenerateRecommendation(farm models.Farm, data models.SatelliteData) models.Recommendation {
	return HighestPriority(GenerateRecommendations(farm, data))
}

// HighestPriority returns the most urgent recommendation while preserving rule
// evaluation order when two recommendations have the same priority.
func HighestPriority(recommendations []models.Recommendation) models.Recommendation {
	if len(recommendations) == 0 {
		return models.Recommendation{}
	}
	best := recommendations[0]
	for _, recommendation := range recommendations[1:] {
		if priorityRank(recommendation.Priority) > priorityRank(best.Priority) {
			best = recommendation
		}
	}
	return best
}

func priorityRank(priority string) int {
	switch priority {
	case "HIGH":
		return 3
	case "MEDIUM":
		return 2
	default:
		return 1
	}
}
