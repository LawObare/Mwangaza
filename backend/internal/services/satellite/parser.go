package satellite

import (
	"encoding/json"
	"time"

	"mwangaza/backend/internal/models"
)

func ParseSatelliteData(data []byte) (models.SatelliteData, error) {
	var resp apiResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return models.SatelliteData{}, err
	}

	return models.SatelliteData{
		SoilMoisture: resp.Soil.Moisture,
		Temperature:  resp.Weather.Temperature,
		RainProb:     resp.Weather.RainProb,
		NDVI:         resp.Vegetation.NDVI,
		WindSpeed:    resp.Weather.WindSpeed,
		Timestamp:    time.Now(),
	}, nil
}