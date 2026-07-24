package satellite

import (
	"encoding/json"
	"mwangaza/internal/models"
	"time"
)

// ParseSatelliteData converts SpaceIoTBox JSON into models.SatelliteData.
func ParseSatelliteData(data []byte) (models.SatelliteData, error) {
	logParsing()
	
	var apiResp APIResponse
	err := json.Unmarshal(data, &apiResp)
	if err != nil {
		return models.SatelliteData{}, ErrParsing
	}

	// Validate required fields (simple check)
	// In a real scenario, we might check if values are within reasonable ranges.

	timestamp, err := time.Parse(time.RFC3339, apiResp.Timestamp)
	if err != nil {
		// Fallback to current time if parsing fails, or return error
		timestamp = time.Now()
	}

	result := models.SatelliteData{
		SoilMoisture: apiResp.Soil.Moisture,
		Temperature:  apiResp.Weather.Temperature,
		RainProb:     apiResp.Weather.RainProbability,
		NDVI:         apiResp.Vegetation.NDVI,
		WindSpeed:    apiResp.Weather.WindSpeed,
		Timestamp:    timestamp,
	}

	return result, nil
}
