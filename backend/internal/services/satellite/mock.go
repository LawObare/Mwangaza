package satellite

import (
	"mwangaza/internal/models"
	"time"
)

// GetMockData returns hardcoded satellite data for demonstrations.
func GetMockData() models.SatelliteData {
	return models.SatelliteData{
		SoilMoisture: 18.5,
		Temperature:  31.2,
		RainProb:     20.0,
		NDVI:         0.63,
		WindSpeed:    7.4,
		Timestamp:    time.Now(),
	}
}
