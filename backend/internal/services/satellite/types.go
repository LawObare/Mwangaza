package satellite

type apiResponse struct {
	Weather    weather    `json:"weather"`
	Soil       soil       `json:"soil"`
	Vegetation vegetation `json:"vegetation"`
}

type soil struct {
	Moisture float64 `json:"moisture"`
}

type weather struct {
	Temperature   float64 `json:"temperature"`
	RainProb      float64 `json:"rain_probability"`
	WindSpeed     float64 `json:"wind_speed"`
}

type vegetation struct {
	NDVI float64 `json:"ndvi"`
}