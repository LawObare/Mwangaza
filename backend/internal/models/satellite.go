package models

type SatelliteData struct {
	ID              int     `json:"id"`
	FarmID          int     `json:"farm_id"`
	SoilMoisture    float64 `json:"soil_moisture"`
	Temperature     float64 `json:"temperature"`
	RainProbability float64 `json:"rain_probability"`
	NDVI            float64 `json:"ndvi"`
	WindSpeed       float64 `json:"wind_speed"`
	Source          string  `json:"source"`
	Timestamp       string  `json:"timestamp"`
}
