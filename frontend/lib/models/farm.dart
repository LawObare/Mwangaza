class Farm {
  const Farm({
    required this.id,
    required this.name,
    required this.farmer,
    required this.phone,
    required this.crop,
    required this.status,
    this.county = '',
    this.subCounty = '',
    this.growthStage = '',
    this.preferredLanguage = 'English',
    this.latitude = 0,
    this.longitude = 0,
    this.soilMoisture = 0,
    this.temperature = 0,
    this.rainProbability = 0,
    this.ndvi = 0,
    this.windSpeed = 0,
    this.lastSmsSent = '',
  });

  final int id;
  final String name;
  final String farmer;
  final String phone;
  final String county;
  final String subCounty;
  final String crop;
  final String growthStage;
  final String preferredLanguage;
  final String status;
  final double latitude;
  final double longitude;
  final double soilMoisture;
  final double temperature;
  final double rainProbability;
  final double ndvi;
  final double windSpeed;
  final String lastSmsSent;

  factory Farm.fromJson(Map<String, dynamic> json) {
    return Farm(
      id: _int(json['id']),
      name: _string(json['name']),
      farmer: _string(json['farmer']),
      phone: _string(json['phone']),
      county: _string(json['county']),
      subCounty: _string(json['sub_county']),
      crop: _string(json['crop']),
      growthStage: _string(json['growth_stage']),
      preferredLanguage: _string(json['preferred_language'], 'English'),
      status: _string(json['status'], 'new'),
      latitude: _double(json['latitude']),
      longitude: _double(json['longitude']),
      soilMoisture: _double(json['soil_moisture']),
      temperature: _double(json['temperature']),
      rainProbability: _double(json['rain_probability']),
      ndvi: _double(json['ndvi']),
      windSpeed: _double(json['wind_speed']),
      lastSmsSent: _string(json['last_sms_sent']),
    );
  }

  Map<String, dynamic> toCreateJson() {
    return {
      'name': name,
      'farmer': farmer,
      'phone': phone,
      'county': county,
      'sub_county': subCounty,
      'crop': crop,
      'growth_stage': growthStage,
      'latitude': latitude,
      'longitude': longitude,
      'preferred_language': preferredLanguage,
    };
  }
}

int _int(dynamic value) => value is num ? value.toInt() : int.tryParse('$value') ?? 0;
double _double(dynamic value) => value is num ? value.toDouble() : double.tryParse('$value') ?? 0;
String _string(dynamic value, [String fallback = '']) {
  final text = value?.toString().trim() ?? '';
  return text.isEmpty ? fallback : text;
}
