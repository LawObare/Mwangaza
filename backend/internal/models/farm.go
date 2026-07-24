package models

// Farm represents a row in the farms table.
type Farm struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Farmer       string  `json:"farmer"`
	Phone        string  `json:"phone"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	SoilMoisture float64 `json:"soil_moisture,omitempty"`
	Temperature  float64 `json:"temperature,omitempty"`
	LastSMSSent  string  `json:"last_sms_sent,omitempty"`
}