// lib/services/sms_service.dart — SMS API service.
//
// Endpoints:
//   GET  /api/sms       → getSmsHistory() returns List<SmsMessage>
//   POST /api/sms/send  → sendSms(farmId, message, phone) sends an SMS
//
// Responsibilities:
//   - GET sms history from backend.
//   - POST new SMS to send via Africa's Talking (or mock).
//   - Deserialize JSON into models/sms.dart.
//
// Connects to:
//   api_service.dart  — makes HTTP requests.
//   models/sms.dart   — target model.
//   screens/sms/      — displays history.
//   widgets/ (send button in dashboard detail screens).
class SmsService {}