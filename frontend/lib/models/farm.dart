class Farm {
  final int id;
  final String name;
  final String farmer;
  final double latitude;
  final double longitude;
  final String status;
  final double? soilMoisture;
  final double? temperature;
  final String? lastSmsSent;

  Farm({
    required this.id,
    required this.name,
    required this.farmer,
    required this.latitude,
    required this.longitude,
    required this.status,
    this.soilMoisture,
    this.temperature,
    this.lastSmsSent,
  });

  factory Farm.fromJson(Map<String, dynamic> json) {
    return Farm(
      id: json['id'] as int,
      name: json['name'] as String,
      farmer: json['farmer'] as String,
      latitude: (json['latitude'] as num).toDouble(),
      longitude: (json['longitude'] as num).toDouble(),
      status: json['status'] as String,
      soilMoisture: (json['soil_moisture'] as num?)?.toDouble(),
      temperature: (json['temperature'] as num?)?.toDouble(),
      lastSmsSent: json['last_sms_sent'] as String?,
    );
  }
}