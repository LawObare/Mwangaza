// Package recommendation (temperature.go) handles temperature-based alerts.
//
// Logic:
//   - If temperature > 35°C → "Heat stress detected — increase watering" (priority: high)
//   - If temperature 30-35°C → "Monitor temperature" (priority: low)
//   - Otherwise → no temperature alert
//
// Consumes SatelliteData.Temperature. Called by rules.go.
package recommendation

import (
	"fmt"

	"mwangaza/internal/models"
)

func temperatureRule(farm models.Farm, data models.SatelliteData, createdAt string) *models.Recommendation {
	if data.Temperature < 33 {
		return nil
	}
	recommendation := recommendationFor(farm, createdAt, "temperature",
		localized(farm, "High temperature detected. Monitor for heat stress and irrigate in the evening where possible.", "Joto ni kubwa. Fuatilia dalili za msongo wa joto na mwagilia jioni inapowezekana."),
		fmt.Sprintf("temperature %.1fC is at or above the 33C heat threshold", data.Temperature), "medium", "MEDIUM")
	return &recommendation
}
