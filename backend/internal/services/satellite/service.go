package satellite

import (
	"os"

	"github.com/laobare/mwangaza/backend/internal/models"
)

func FetchSatelliteData() (models.SatelliteData, error) {
	if os.Getenv("USE_MOCK_DATA") == "true" {
		return GetMockData(), nil
	}

	raw, err := FetchLiveData()
	if err != nil {
		return models.SatelliteData{}, err
	}

	return ParseSatelliteData(raw)
}