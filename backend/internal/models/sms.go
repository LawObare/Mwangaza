// Package models (sms.go) defines the SmsLog struct.
//
// Represents a row in the sms_logs table. Records every SMS sent via the system.
//
// Fields:
//   ID      int    — primary key
//   Phone   string — recipient phone number
//   Message string — SMS body text
//   Status  string — "sent", "failed", "pending"
//   SentAt  string — ISO 8601 timestamp
//
// Used by handlers to return SMS history to the Flutter dashboard.
package models