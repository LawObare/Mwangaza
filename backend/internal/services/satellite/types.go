package satellite

// APIResponse represents the top-level response from SpaceIoTBox.
type APIResponse struct {
	Weather    Weather    `json:"weather"`
	Soil       Soil       `json:"soil"`
	Vegetation Vegetation `json:"vegetation"`
	Timestamp  string     `json:"timestamp"`
}

// Weather represents weather data from the API.
type Weather struct {
	Temperature     float64 `json:"temperature"`
	RainProbability float64 `json:"rain_probability"`
	WindSpeed       float64 `json:"wind_speed"`
}

// Soil represents soil data from the API.
type Soil struct {
	Moisture float64 `json:"moisture"`
}

// Vegetation represents vegetation data from the API.
type Vegetation struct {
	NDVI float64 `json:"ndvi"`
}
