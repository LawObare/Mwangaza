// Package recommendation (rainfall.go) analyzes rain probability data.
//
// Logic:
//   - If rain_probability > 60% → "Delay irrigation — rain expected" (priority: medium)
//   - If rain_probability 30-60% → "Monitor weather" (priority: low)
//   - Otherwise → no rainfall recommendation
//
// Consumes SatelliteData.RainProb. Called by rules.go.
package recommendation

import (
	"fmt"

	"mwangaza/internal/models"
)

func heavyRainRule(farm models.Farm, data models.SatelliteData, createdAt string) *models.Recommendation {
	if data.RainProbability < 80 {
		return nil
	}
	recommendation := recommendationFor(farm, createdAt, "rainfall",
		localized(farm, "Heavy rainfall is expected soon. Delay fertilizer and field operations, and keep drainage clear.", "Mvua kubwa inatarajiwa hivi karibuni. Ahirisha kuweka mbolea na kazi za shambani, na hakikisha mifereji iko wazi."),
		fmt.Sprintf("rain probability %.0f%% is at or above the 80%% heavy-rain threshold", data.RainProbability), "high", "HIGH")
	return &recommendation
}

func rainfallRule(farm models.Farm, data models.SatelliteData, createdAt string) *models.Recommendation {
	if data.RainProbability < 55 || data.RainProbability >= 80 {
		return nil
	}
	recommendation := recommendationFor(farm, createdAt, "rainfall",
		localized(farm, "Rain is likely soon. Prepare field drains and avoid spraying before the rain.", "Mvua inawezekana hivi karibuni. Andaa mifereji na epuka kunyunyizia dawa kabla ya mvua."),
		fmt.Sprintf("rain probability %.0f%% is within the 55-79%% likely-rain range", data.RainProbability), "medium", "MEDIUM")
	return &recommendation
}
