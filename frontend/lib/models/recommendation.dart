class Recommendation {
  final int farmId;
  final String type;
  final String message;
  final String severity;
  final String createdAt;

  Recommendation({
    required this.farmId,
    required this.type,
    required this.message,
    required this.severity,
    required this.createdAt,
  });

  factory Recommendation.fromJson(Map<String, dynamic> json) {
    return Recommendation(
      farmId: json['farm_id'] as int,
      type: json['type'] as String,
      message: json['message'] as String,
      severity: json['severity'] as String,
      createdAt: json['created_at'] as String,
    );
  }
}