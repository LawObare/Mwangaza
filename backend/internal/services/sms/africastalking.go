// Package sms (africastalking.go) integrates with Africa's Talking SMS API.
//
// It sends SMS messages via the Africa's Talking REST API.
//
// Required env:
//   SMS_API_KEY — Africa's Talking API key
//
// Steps:
//   1. POST to https://api.africastalking.com/version1/messaging
//   2. Headers: ApiKey, Accept: application/json
//   3. Body: { "username": "...", "to": phone, "message": text }
//   4. Parse response for status ("Sent", "Failed", etc.)
//
// Returns the status string and any error encountered.
// Called by service.go when USE_MOCK_DATA != "true".
package sms