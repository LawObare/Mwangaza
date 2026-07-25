// Package recommendation (rules.go) contains the top-level decision logic.
//
// It evaluates all environmental factors (soil moisture, temperature,
// rain probability, NDVI, wind speed) against defined thresholds
// and delegates to specialized sub-files (irrigation, rainfall, etc.).
//
// Thresholds (example):
//
//	Soil moisture < 15%   → irrigation needed
//	Rain probability > 60% → delay irrigation
//	Temperature > 35°C    → heat stress warning
//	NDVI < 0.3            → poor vegetation health
//
// Each check returns a Recommendation with a priority level and message.
package recommendation

import (
	"strings"

	"mwangaza/internal/models"
)

type rule func(models.Farm, models.SatelliteData, string) *models.Recommendation

func recommendationFor(farm models.Farm, createdAt, kind, message, reason, severity, confidence string) models.Recommendation {
	return models.Recommendation{
		FarmID: farm.ID, Crop: farm.Crop, Type: kind, Message: message,
		Severity: severity, Priority: strings.ToUpper(severity), Confidence: confidence,
		Reason: reason, CreatedAt: createdAt,
	}
}

func monitoringRecommendation(farm models.Farm, createdAt string) models.Recommendation {
	return recommendationFor(farm, createdAt, "monitoring",
		localized(farm, "Conditions are stable. Continue monitoring the field.", "Hali ya shamba ni tulivu. Endelea kufuatilia shamba."),
		"No alert threshold was triggered.", "low", "LOW")
}

func localized(farm models.Farm, english, swahili string) string {
	if strings.EqualFold(strings.TrimSpace(farm.PreferredLanguage), "swahili") {
		return swahili + cropNote(farm, true)
	}
	return english + cropNote(farm, false)
}

func cropNote(farm models.Farm, swahili bool) string {
	crop := strings.ToLower(strings.TrimSpace(farm.Crop))
	english := map[string]string{
		"maize": " Focus on the maize crop.", "beans": " Check the bean plants closely.", "rice": " Check paddy water management.",
		"sugarcane": " Protect the sugarcane rows.", "tea": " Monitor the tea bushes.", "cassava": " Check young cassava plants.",
		"tomatoes": " Inspect tomato plants and fruit.", "onions": " Inspect onion beds.", "coffee": " Monitor the coffee trees.", "bananas": " Check banana mats and leaves.",
	}
	swahiliNotes := map[string]string{
		"maize": " Zingatia zao la mahindi.", "beans": " Kagua mimea ya maharagwe kwa karibu.", "rice": " Kagua usimamizi wa maji shambani.",
		"sugarcane": " Linda mistari ya miwa.", "tea": " Fuatilia vichaka vya chai.", "cassava": " Kagua mimea michanga ya muhogo.",
		"tomatoes": " Kagua mimea na matunda ya nyanya.", "onions": " Kagua matuta ya vitunguu.", "coffee": " Fuatilia miti ya kahawa.", "bananas": " Kagua migomba na majani yake.",
	}
	if swahili {
		return swahiliNotes[crop]
	}
	return english[crop]
}
