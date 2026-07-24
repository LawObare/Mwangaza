// Package sms (service.go) is the SMS service entry point.
//
// It decides whether to send via Africa's Talking API or mock,
// based on USE_MOCK_DATA env var.
//
// Public function:
//   func SendSMS(phone, message string) (status string, err error)
//
// If mock → returns "sent" immediately.
// If live → calls africastalking.go to send via API.
//
// The handler at handlers/sms_handler.go calls this and logs to sms_logs.
package sms