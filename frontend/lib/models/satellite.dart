class SatelliteData {
  final double soilMoisture;
  final double temperature;
  final double rainProbability;
  final double ndvi;
  final double windSpeed;
  final String timestamp;

  SatelliteData({
    required this.soilMoisture,
    required this.temperature,
    required this.rainProbability,
    required this.ndvi,
    required this.windSpeed,
    required this.timestamp,
  });

  factory SatelliteData.fromJson(Map<String, dynamic> json) {
    return SatelliteData(
      soilMoisture: (json['soil_moisture'] as num).toDouble(),
      temperature: (json['temperature'] as num).toDouble(),
      rainProbability: (json['rain_probability'] as num).toDouble(),
      ndvi: (json['ndvi'] as num).toDouble(),
      windSpeed: (json['wind_speed'] as num).toDouble(),
      timestamp: json['timestamp'] as String,
    );
  }
}