package satellite

import (
	"mwangaza/internal/models"
)

// FetchSatelliteData is the main service exposed to the rest of the application.
// It decides whether to use mock data or the live API based on configuration.
func FetchSatelliteData() (models.SatelliteData, error) {
	logRequestStarted()

	if UseMockData() {
		data := GetMockData()
		logFinished()
		return data, nil
	}

	raw, err := FetchLiveData()
	if err != nil {
		return models.SatelliteData{}, err
	}

	data, err := ParseSatelliteData(raw)
	if err != nil {
		return models.SatelliteData{}, err
	}

	logFinished()
	return data, nil
}
