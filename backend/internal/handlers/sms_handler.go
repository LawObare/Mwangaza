// Package handlers (sms_handler.go) serves SMS-related endpoints.
//
// Endpoints:
//   GET  /api/sms       — returns sms_logs table (history)
//   POST /api/sms/send  — accepts { farm_id, message, phone_number },
//                          validates, sends via SMS service, logs to sms_logs
//
// The POST handler should:
//   1. Validate phone number (+254 format, 12 digits)
//   2. Validate farm_id exists in farms table
//   3. Call sms service (mock or Africa's Talking)
//   4. Insert result into sms_logs table
//   5. Return success/failure
package handlers