// lib/screens/sms/sms_screen.dart — SMS history screen.
//
// Layout:
//   AppBar with "SMS History" title.
//   ListView of SMS message cards.
//   Each card shows: icon (sent=green check, pending=orange pending),
//   message preview, phone number, timestamp, status badge.
//
// Responsibilities:
//   - On init: fetch SMS history from backend.
//   - Show LoadingWidget while loading.
//   - Pull-to-refresh.
//
// Connects to:
//   services/sms_service.dart   — fetches SMS history.
//   models/sms.dart             — data model for each message.
//   widgets/loading_widget.dart — loading state.
class SmsScreen {}