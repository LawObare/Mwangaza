// lib/models/sms.dart — SMS message data model.
//
// Fields:
//   int id, int farmId, String phoneNumber, String message, String status, String sentAt
//
// Responsibilities:
//   - Define an SmsMessage class with a fromJson() factory constructor.
//   - JSON keys match the Go backend GET /api/sms response.
//
// Used by:
//   services/sms_service.dart — deserializes API response.
//   screens/sms/ — displays SMS history list.
class SmsMessage {}