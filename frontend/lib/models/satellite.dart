class SatelliteData {
  const SatelliteData({
    required this.id,
    required this.farmId,
    required this.soilMoisture,
    required this.temperature,
    required this.rainProbability,
    required this.ndvi,
    required this.windSpeed,
    this.source = '',
    this.timestamp = '',
  });

  final int id;
  final int farmId;
  final double soilMoisture;
  final double temperature;
  final double rainProbability;
  final double ndvi;
  final double windSpeed;
  final String source;
  final String timestamp;

  factory SatelliteData.fromJson(Map<String, dynamic> json) {
    return SatelliteData(
      id: _int(json['id']),
      farmId: _int(json['farm_id']),
      soilMoisture: _double(json['soil_moisture']),
      temperature: _double(json['temperature']),
      rainProbability: _double(json['rain_probability']),
      ndvi: _double(json['ndvi']),
      windSpeed: _double(json['wind_speed']),
      source: _string(json['source']),
      timestamp: _string(json['timestamp']),
    );
  }
}

int _int(dynamic value) => value is num ? value.toInt() : int.tryParse('$value') ?? 0;
double _double(dynamic value) => value is num ? value.toDouble() : double.tryParse('$value') ?? 0;
String _string(dynamic value, [String fallback = '']) {
  final text = value?.toString().trim() ?? '';
  return text.isEmpty ? fallback : text;
}
