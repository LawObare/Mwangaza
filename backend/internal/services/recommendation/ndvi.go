// Package recommendation (ndvi.go) analyzes vegetation health from NDVI.
//
// Logic:
//   - If NDVI < 0.3 → "Poor vegetation health — check for pests/disease" (priority: high)
//   - If NDVI 0.3-0.5 → "Moderate vegetation health" (priority: low)
//   - If NDVI > 0.5 → no NDVI recommendation (healthy)
//
// Consumes SatelliteData.NDVI. Called by rules.go.
package recommendation

import (
	"fmt"

	"mwangaza/internal/models"
)

func ndviRule(farm models.Farm, data models.SatelliteData, createdAt string) *models.Recommendation {
	if data.NDVI >= 0.35 {
		return nil
	}
	recommendation := recommendationFor(farm, createdAt, "ndvi",
		localized(farm, "Vegetation health is weak. Inspect the field for pests, disease, or nutrient stress.", "Afya ya mimea ni dhaifu. Kagua shamba kwa wadudu, magonjwa, au upungufu wa virutubisho."),
		fmt.Sprintf("NDVI %.2f is below the 0.35 vegetation-health threshold", data.NDVI), "medium", "MEDIUM")
	return &recommendation
}
