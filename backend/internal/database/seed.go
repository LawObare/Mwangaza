package database

import (
	"time"

	"mwangaza/internal/models"
)

func Seed(store *Store) error {
	if store == nil {
		return nil
	}

	if len(store.ListFarms()) > 0 {
		return nil
	}

	seeded := []models.Farm{
		{
			Name:              "Green Valley Farm",
			Farmer:            "Amina Njeri",
			Phone:             "+254700000001",
			County:            "Kakamega",
			SubCounty:         "Mumias",
			Crop:              "Maize",
			GrowthStage:       "Vegetative",
			Latitude:          -0.3071,
			Longitude:         34.7894,
			PreferredLanguage: "English",
		},
		{
			Name:              "Sunrise Acres",
			Farmer:            "John Ouma",
			Phone:             "+254700000002",
			County:            "Siaya",
			SubCounty:         "Bondo",
			Crop:              "Rice",
			GrowthStage:       "Flowering",
			Latitude:          -0.0931,
			Longitude:         34.2677,
			PreferredLanguage: "Swahili",
		},
		{
			Name:              "Hilltop Farm",
			Farmer:            "Grace Wanjiku",
			Phone:             "+254700000003",
			County:            "Nyeri",
			SubCounty:         "Kieni",
			Crop:              "Tea",
			GrowthStage:       "Harvesting",
			Latitude:          -0.4201,
			Longitude:         36.9476,
			PreferredLanguage: "English",
		},
	}

	for i, farm := range seeded {
		created, err := store.CreateFarm(farm)
		if err != nil {
			return err
		}

		satellite := models.SatelliteData{
			FarmID:          created.ID,
			SoilMoisture:    []float64{18, 68, 22}[i],
			Temperature:     []float64{33, 27, 34}[i],
			RainProbability: []float64{12, 84, 35}[i],
			NDVI:            []float64{0.42, 0.64, 0.29}[i],
			WindSpeed:       []float64{4, 5, 6}[i],
			Source:          "seed",
			Timestamp:       time.Now().UTC().Add(-time.Duration(3-i) * time.Hour).Format(time.RFC3339),
		}
		if _, err := store.AddSatelliteData(satellite); err != nil {
			return err
		}

		recommendation := models.Recommendation{
			FarmID:    created.ID,
			Crop:      created.Crop,
			Type:      "monitoring",
			Message:   "Conditions are stable. Continue monitoring the field.",
			Severity:  "low",
			Priority:  "LOW",
			Reason:    "Seed data",
			CreatedAt: time.Now().UTC().Format(time.RFC3339),
		}

		switch i {
		case 0:
			recommendation.Type = "irrigation"
			recommendation.Message = "Dry soil detected. Irrigate today before noon."
			recommendation.Severity = "high"
			recommendation.Priority = "HIGH"
			recommendation.Reason = "soil moisture below the dry threshold"
		case 1:
			recommendation.Type = "rainfall"
			recommendation.Message = "Heavy rainfall expected tomorrow. Keep drainage clear and delay fertilizer."
			recommendation.Severity = "high"
			recommendation.Priority = "HIGH"
			recommendation.Reason = "rain probability is above the heavy-rain threshold"
		case 2:
			recommendation.Type = "temperature"
			recommendation.Message = "High temperature and low NDVI detected. Increase moisture monitoring and inspect the crop."
			recommendation.Severity = "medium"
			recommendation.Priority = "MEDIUM"
			recommendation.Reason = "temperature and NDVI indicate stress"
		}

		if _, err := store.AddRecommendation(recommendation); err != nil {
			return err
		}

		sms := models.SmsMessage{
			FarmID:      created.ID,
			PhoneNumber: created.Phone,
			Message:     recommendation.Message,
			Status:      "sent",
			Provider:    "mock",
			SentAt:      time.Now().UTC().Format(time.RFC3339),
		}
		if _, err := store.AddSMSLog(sms); err != nil {
			return err
		}
	}

	return nil
}
