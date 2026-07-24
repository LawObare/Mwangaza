class SmsMessage {
  final int id;
  final int farmId;
  final String phoneNumber;
  final String message;
  final String status;
  final String sentAt;

  SmsMessage({
    required this.id,
    required this.farmId,
    required this.phoneNumber,
    required this.message,
    required this.status,
    required this.sentAt,
  });

  factory SmsMessage.fromJson(Map<String, dynamic> json) {
    return SmsMessage(
      id: json['id'] as int,
      farmId: json['farm_id'] as int,
      phoneNumber: json['phone_number'] as String,
      message: json['message'] as String,
      status: json['status'] as String,
      sentAt: json['sent_at'] as String,
    );
  }
}