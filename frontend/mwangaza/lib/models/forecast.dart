// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:convert';

class ForecastData {
  final List<String> time;
  final List<double> temperature;
  final List<double> felttemperature;
  final List<double> precipitation;
  final List<int> precipitationProbability;
  final List<double> convectivePrecipitation;
  final List<int> relativehumidity;
  final List<double> windspeed;
  final List<int> winddirection;
  final List<int> uvindex;
  final List<double> soiltemperature0to10cm;
  final List<int> soilmoisture0to10cm;
  final List<double> potentialevapotranspiration;
  final List<double> evapotranspiration;
  final List<int> sensibleheatflux;
  final List<double> skintemperature;
  final List<int> leafwetnessindex;
  final List<double> dewpointtemperature;
  final List<double> wetbulbtemperature;

  ForecastData({
    required this.time,
    required this.temperature,
    required this.felttemperature,
    required this.precipitation,
    required this.precipitationProbability,
    required this.convectivePrecipitation,
    required this.relativehumidity,
    required this.windspeed,
    required this.winddirection,
    required this.uvindex,
    required this.soiltemperature0to10cm,
    required this.soilmoisture0to10cm,
    required this.potentialevapotranspiration,
    required this.evapotranspiration,
    required this.sensibleheatflux,
    required this.skintemperature,
    required this.leafwetnessindex,
    required this.dewpointtemperature,
    required this.wetbulbtemperature,
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
      'uvindex': uvindex,
      'soiltemperature_0to10cm': soiltemperature0to10cm,
      'soilmoisture_0to10cm': soilmoisture0to10cm,
      'potentialevapotranspiration': potentialevapotranspiration,
      'evapotranspiration': evapotranspiration,
      'sensibleheatflux': sensibleheatflux,
      'skintemperature': skintemperature,
      'leafwetnessindex': leafwetnessindex,
      'dewpointtemperature': dewpointtemperature,
      'wetbulbtemperature': wetbulbtemperature,
    };
  }

  factory ForecastData.fromMap(Map<String, dynamic> map) {
    return ForecastData(
      time: List<String>.from(map['time'] ?? []),
      temperature: List<double>.from(
        map['temperature']?.map((x) => x?.toDouble() ?? 0.0) ?? [],
      ),
      felttemperature: List<double>.from(
        map['felttemperature']?.map((x) => x?.toDouble() ?? 0.0) ?? [],
      ),
      precipitation: List<double>.from(
        map['precipitation']?.map((x) => x?.toDouble() ?? 0.0) ?? [],
      ),
      precipitationProbability: List<int>.from(
        map['precipitation_probability']?.map((x) => x?.toInt() ?? 0) ?? [],
      ),
      convectivePrecipitation: List<double>.from(
        map['convective_precipitation']?.map((x) => x?.toDouble() ?? 0.0) ?? [],
      ),
      relativehumidity: List<int>.from(
        map['relativehumidity']?.map((x) => x?.toInt() ?? 0) ?? [],
      ),
      windspeed: List<double>.from(
        map['windspeed']?.map((x) => x?.toDouble() ?? 0.0) ?? [],
      ),
      winddirection: List<int>.from(
        map['winddirection']?.map((x) => x?.toInt() ?? 0) ?? [],
      ),
      uvindex: List<int>.from(
        map['uvindex']?.map((x) => x?.toInt() ?? 0) ?? [],
      ),
      soiltemperature0to10cm: List<double>.from(
        map['soiltemperature_0to10cm']?.map((x) => x?.toDouble() ?? 0.0) ?? [],
      ),
      soilmoisture0to10cm: List<int>.from(
        map['soilmoisture_0to10cm']?.map((x) => x?.toInt() ?? 0) ?? [],
      ),
      potentialevapotranspiration: List<double>.from(
        map['potentialevapotranspiration']?.map((x) => x?.toDouble() ?? 0.0) ??
            [],
      ),
      evapotranspiration: List<double>.from(
        map['evapotranspiration']?.map((x) => x?.toDouble() ?? 0.0) ?? [],
      ),
      sensibleheatflux: List<int>.from(
        map['sensibleheatflux']?.map((x) => x?.toInt() ?? 0) ?? [],
      ),
      skintemperature: List<double>.from(
        map['skintemperature']?.map((x) => x?.toDouble() ?? 0.0) ?? [],
      ),
      leafwetnessindex: List<int>.from(
        map['leafwetnessindex']?.map((x) => x?.toInt() ?? 0) ?? [],
      ),
      dewpointtemperature: List<double>.from(
        map['dewpointtemperature']?.map((x) => x?.toDouble() ?? 0.0) ?? [],
      ),
      wetbulbtemperature: List<double>.from(
        map['wetbulbtemperature']?.map((x) => x?.toDouble() ?? 0.0) ?? [],
      ),
    );
  }

  String toJson() => json.encode(toMap());

  factory ForecastData.fromJson(String source) =>
      ForecastData.fromMap(json.decode(source) as Map<String, dynamic>);
}
