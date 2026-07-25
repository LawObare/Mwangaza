package recommendation

import (
	"fmt"

	"mwangaza/internal/models"
)

func droughtStressRule(farm models.Farm, data models.SatelliteData, createdAt string) *models.Recommendation {
	if data.NDVI >= 0.35 || data.SoilMoisture >= 20 {
		return nil
	}
	recommendation := recommendationFor(farm, createdAt, "drought_stress",
		localized(farm, "Possible drought stress detected. Prioritize irrigation and inspect the crop today.", "Dalili za ukame zimetambuliwa. Tanguliza umwagiliaji na kagua zao leo."),
		fmt.Sprintf("NDVI %.2f and soil moisture %.0f%% indicate drought stress", data.NDVI, data.SoilMoisture), "high", "HIGH")
	return &recommendation
}

func heatStressRule(farm models.Farm, data models.SatelliteData, createdAt string) *models.Recommendation {
	if data.Temperature <= 34 || data.SoilMoisture >= 20 {
		return nil
	}
	recommendation := recommendationFor(farm, createdAt, "heat_stress",
		localized(farm, "Heat stress risk is high. Irrigate in the evening and reduce crop exposure where possible.", "Hatari ya msongo wa joto ni kubwa. Mwagilia jioni na punguza athari ya jua inapowezekana."),
		fmt.Sprintf("temperature %.1fC and soil moisture %.0f%% indicate heat stress", data.Temperature, data.SoilMoisture), "high", "HIGH")
	return &recommendation
}
