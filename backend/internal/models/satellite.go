package models

import "time"

// SatelliteData represents environmental data from SpaceIoTBox.
type SatelliteData struct {
	SoilMoisture float64   `json:"soil_moisture"`
	Temperature  float64   `json:"temperature"`
	RainProb     float64   `json:"rain_probability"`
	NDVI         float64   `json:"ndvi"`
	WindSpeed    float64   `json:"wind_speed"`
	Timestamp    time.Time `json:"timestamp"`
}