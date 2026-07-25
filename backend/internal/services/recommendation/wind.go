package recommendation

import (
	"fmt"

	"mwangaza/internal/models"
)

func windRule(farm models.Farm, data models.SatelliteData, createdAt string) *models.Recommendation {
	if data.WindSpeed <= 30 {
		return nil
	}
	recommendation := recommendationFor(farm, createdAt, "wind",
		localized(farm, "Strong wind detected. Avoid spraying pesticides today to prevent drift.", "Upepo mkali umetambuliwa. Epuka kunyunyizia dawa leo ili kuzuia dawa kupeperushwa."),
		fmt.Sprintf("wind speed %.0f km/h is above the 30 km/h spraying-safety threshold", data.WindSpeed), "high", "HIGH")
	return &recommendation
}
