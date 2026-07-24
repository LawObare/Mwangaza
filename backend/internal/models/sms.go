package models

type SmsMessage struct {
	ID          int    `json:"id"`
	FarmID      int    `json:"farm_id"`
	PhoneNumber string `json:"phone_number"`
	Message     string `json:"message"`
	Status      string `json:"status"`
	Provider    string `json:"provider"`
	SentAt      string `json:"sent_at"`
}
