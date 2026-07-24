// Package sms (mock.go) simulates sending an SMS without an API call.
//
// Function:
//   func SendMockSMS(phone, message string) string
//
// Simply returns "sent" after logging to console.
// Used when USE_MOCK_DATA=true in env.
// Called by service.go → SendSMS().
package sms