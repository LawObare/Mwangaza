// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:convert';

class Units {
  final String time;
  final String temperature;
  final String felttemperature;
  final String precipitation;
  final String precipitationProbability;
  final String convectivePrecipitation;
  final String relativehumidity;
  final String windspeed;
  final String winddirection;
  final String soiltemperature0to10cm;
  final String soilmoisture0to10cm;
  final String potentialevapotranspiration;
  final String evapotranspiration;
  final String sensibleheatflux;
  final String skintemperature;
  final String leafwetness;
  final String dewpointtemperature;
  final String wetbulbtemperature;
  final String ndvi;
  final String vegetationHealthIndex;
  final String vegetationConditionIndex;

  Units({
    required this.time,
    required this.temperature,
    required this.felttemperature,
    required this.precipitation,
    required this.precipitationProbability,
    required this.convectivePrecipitation,
    required this.relativehumidity,
    required this.windspeed,
    required this.winddirection,
    required this.soiltemperature0to10cm,
    required this.soilmoisture0to10cm,
    required this.potentialevapotranspiration,
    required this.evapotranspiration,
    required this.sensibleheatflux,
    required this.skintemperature,
    required this.leafwetness,
    required this.dewpointtemperature,
    required this.wetbulbtemperature,
    required this.ndvi,
    required this.vegetationHealthIndex,
    required this.vegetationConditionIndex,
  });

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'time': time,
      'temperature': temperature,
      'felttemperature': felttemperature,
      'precipitation': precipitation,
      'precipitation_probability': precipitationProbability,
      'convective_precipitation': convectivePrecipitation,
      'relativehumidity': relativehumidity,
      'windspeed': windspeed,
      'winddirection': winddirection,
      'soiltemperature_0to10cm': soiltemperature0to10cm,
      'soilmoisture_0to10cm': soilmoisture0to10cm,
      'potentialevapotranspiration': potentialevapotranspiration,
      'evapotranspiration': evapotranspiration,
      'sensibleheatflux': sensibleheatflux,
      'skintemperature': skintemperature,
      'leafwetness': leafwetness,
      'dewpointtemperature': dewpointtemperature,
      'wetbulbtemperature': wetbulbtemperature,
      'NDVI': ndvi,
      'vegetation_health_index': vegetationHealthIndex,
      'vegetation_condition_index': vegetationConditionIndex,
    };
  }

  factory Units.fromMap(Map<String, dynamic> map) {
    return Units(
      time: map['time'] ?? '',
      temperature: map['temperature'] ?? '',
      felttemperature: map['felttemperature'] ?? '',
      precipitation: map['precipitation'] ?? '',
      precipitationProbability: map['precipitation_probability'] ?? '',
      convectivePrecipitation: map['convective_precipitation'] ?? '',
      relativehumidity: map['relativehumidity'] ?? '',
      windspeed: map['windspeed'] ?? '',
      winddirection: map['winddirection'] ?? '',
      soiltemperature0to10cm: map['soiltemperature_0to10cm'] ?? '',
      soilmoisture0to10cm: map['soilmoisture_0to10cm'] ?? '',
      potentialevapotranspiration: map['potentialevapotranspiration'] ?? '',
      evapotranspiration: map['evapotranspiration'] ?? '',
      sensibleheatflux: map['sensibleheatflux'] ?? '',
      skintemperature: map['skintemperature'] ?? '',
      leafwetness: map['leafwetness'] ?? '',
      dewpointtemperature: map['dewpointtemperature'] ?? '',
      wetbulbtemperature: map['wetbulbtemperature'] ?? '',
      ndvi: map['NDVI'] ?? '',
      vegetationHealthIndex: map['vegetation_health_index'] ?? '',
      vegetationConditionIndex: map['vegetation_condition_index'] ?? '',
    );
  }

  String toJson() => json.encode(toMap());

  factory Units.fromJson(String source) =>
      Units.fromMap(json.decode(source) as Map<String, dynamic>);
}
