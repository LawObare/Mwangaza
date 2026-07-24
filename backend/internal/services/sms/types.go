// Package sms (types.go) defines request/response types for SMS operations.
//
// These types are used by the handler for request validation and by
// the Africa's Talking integration for API response parsing.
//
// Types:
//   SendSMSRequest — { farm_id, message, phone_number } — JSON binding for POST /api/sms/send
//   ATResponse     — Africa's Talking API response envelope
//   ATMessageData  — per-message status from AT response
//
// Not exported unless needed by handlers.
package sms