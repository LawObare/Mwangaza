package models

type Farm struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Farmer            string  `json:"farmer"`
	Phone             string  `json:"phone"`
	County            string  `json:"county"`
	SubCounty         string  `json:"sub_county,omitempty"`
	Crop              string  `json:"crop"`
	GrowthStage       string  `json:"growth_stage,omitempty"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	PreferredLanguage string  `json:"preferred_language,omitempty"`
	Status            string  `json:"status"`
	SoilMoisture      float64 `json:"soil_moisture"`
	Temperature       float64 `json:"temperature"`
	RainProbability   float64 `json:"rain_probability"`
	NDVI              float64 `json:"ndvi"`
	WindSpeed         float64 `json:"wind_speed"`
	LastSmsSent       string  `json:"last_sms_sent"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}
