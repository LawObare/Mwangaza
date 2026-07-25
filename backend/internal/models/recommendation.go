package models

type Recommendation struct {
	ID         int    `json:"id"`
	FarmID     int    `json:"farm_id"`
	Crop       string `json:"crop"`
	Type       string `json:"type"`
	Message    string `json:"message"`
	Severity   string `json:"severity"`
	Priority   string `json:"priority"`
	Confidence string `json:"confidence"`
	Reason     string `json:"reason"`
	CreatedAt  string `json:"created_at"`
}
