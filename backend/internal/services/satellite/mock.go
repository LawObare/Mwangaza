package satellite

import (
	"time"

	"mwangaza/backend/internal/models"
)

func GetMockData() models.SatelliteData {
	return models.SatelliteData{
		SoilMoisture: 18.5,
		Temperature:  31,
		RainProb:     20,
		NDVI:         0.64,
		WindSpeed:    8,
		Timestamp:    time.Now(),
	}
}