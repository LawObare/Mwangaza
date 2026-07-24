// Package recommendation (irrigation.go) handles irrigation-related decisions.
//
// Logic:
//   - If soil_moisture < 15% → "Irrigate for 20 minutes" (priority: high)
//   - If soil_moisture 15-25% → "Monitor soil moisture" (priority: low)
//   - Otherwise → no irrigation recommendation
//
// Consumes soil moisture from SatelliteData.SoilMoisture.
// Called by the rule engine in rules.go.
package recommendation

import (
	"fmt"

	"mwangaza/internal/models"
)

func irrigationRule(farm models.Farm, data models.SatelliteData, createdAt string) *models.Recommendation {
	if data.SoilMoisture >= 20 {
		return nil
	}
	recommendation := recommendationFor(farm, createdAt, "irrigation",
		localized(farm, "Soil moisture is critically low. Irrigate within the next 12 hours to reduce crop stress.", "Unyevu wa udongo ni mdogo sana. Mwagilia ndani ya saa 12 zijazo ili kupunguza msongo wa zao."),
		fmt.Sprintf("soil moisture %.0f%% is below the 20%% threshold", data.SoilMoisture), "high", "MEDIUM")
	return &recommendation
}
