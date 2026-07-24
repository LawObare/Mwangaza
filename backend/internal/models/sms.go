package models

// SmsLog represents a row in the sms_logs table.
type SmsLog struct {
	ID      int    `json:"id"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
	Status  string `json:"status"`
	SentAt  string `json:"sent_at"`
}