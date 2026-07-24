class SmsMessage {
  const SmsMessage({
    required this.id,
    required this.farmId,
    required this.phoneNumber,
    required this.message,
    required this.status,
    required this.provider,
    required this.sentAt,
  });

  final int id;
  final int farmId;
  final String phoneNumber;
  final String message;
  final String status;
  final String provider;
  final String sentAt;

  factory SmsMessage.fromJson(Map<String, dynamic> json) {
    return SmsMessage(
      id: _int(json['id']),
      farmId: _int(json['farm_id']),
      phoneNumber: _string(json['phone_number']),
      message: _string(json['message']),
      status: _string(json['status']),
      provider: _string(json['provider']),
      sentAt: _string(json['sent_at']),
    );
  }
}

int _int(dynamic value) => value is num ? value.toInt() : int.tryParse('$value') ?? 0;
String _string(dynamic value, [String fallback = '']) {
  final text = value?.toString().trim() ?? '';
  return text.isEmpty ? fallback : text;
}
