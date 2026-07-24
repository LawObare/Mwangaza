class Recommendation {
  const Recommendation({
    required this.id,
    required this.farmId,
    required this.type,
    required this.message,
    required this.severity,
    required this.priority,
    this.crop = '',
    this.confidence = '',
    this.reason = '',
    this.createdAt = '',
  });

  final int id;
  final int farmId;
  final String crop;
  final String type;
  final String message;
  final String severity;
  final String priority;
  final String confidence;
  final String reason;
  final String createdAt;

  factory Recommendation.fromJson(Map<String, dynamic> json) {
    return Recommendation(
      id: _int(json['id']),
      farmId: _int(json['farm_id']),
      crop: _string(json['crop']),
      type: _string(json['type'], 'monitoring'),
      message: _string(json['message']),
      severity: _string(json['severity'], 'low'),
      priority: _string(json['priority'], 'LOW').toUpperCase(),
      confidence: _string(json['confidence']),
      reason: _string(json['reason']),
      createdAt: _string(json['created_at']),
    );
  }

  static List<Recommendation> listFromFlexibleData(dynamic data) {
    if (data is List) {
      return data
          .whereType<Map>()
          .map((item) => Recommendation.fromJson(Map<String, dynamic>.from(item)))
          .toList();
    }
    if (data is Map) {
      final out = <Recommendation>[];
      for (final value in data.values) {
        out.addAll(listFromFlexibleData(value));
      }
      return out;
    }
    return const [];
  }
}

int _int(dynamic value) => value is num ? value.toInt() : int.tryParse('$value') ?? 0;
String _string(dynamic value, [String fallback = '']) {
  final text = value?.toString().trim() ?? '';
  return text.isEmpty ? fallback : text;
}
