// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:convert';

import 'package:mwangaza/models/data.dart';
import 'package:mwangaza/models/forecast.dart';
import 'package:mwangaza/models/forecast_info.dart';
import 'package:mwangaza/models/location.dart';
import 'package:mwangaza/models/unit.dart';

class WeatherResponse {
  final Location location;
  final ForecastInfo forecastInfo;
  final Units units;
  final WeatherData data;
  final ForecastData forecastData;
  final String source;

  WeatherResponse({
    required this.location,
    required this.forecastInfo,
    required this.units,
    required this.data,
    required this.forecastData,
    required this.source,
  });

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'location': location.toMap(),
      'forecast_info': forecastInfo.toMap(),
      'units': units.toMap(),
      'data': data.toMap(),
      'forecast_data': forecastData.toMap(),
      'source': source,
    };
  }

  factory WeatherResponse.fromMap(Map<String, dynamic> map) {
    return WeatherResponse(
      location: Location.fromMap(map['location'] ?? {}),
      forecastInfo: ForecastInfo.fromMap(map['forecast_info'] ?? {}),
      units: Units.fromMap(map['units'] ?? {}),
      data: WeatherData.fromMap(map['data'] ?? {}),
      forecastData: ForecastData.fromMap(map['forecast_data'] ?? {}),
      source: map['source'] ?? '',
    );
  }

  String toJson() => json.encode(toMap());

  factory WeatherResponse.fromJson(String source) =>
      WeatherResponse.fromMap(json.decode(source) as Map<String, dynamic>);

  @override
  String toString() {
    return 'WeatherResponse(location: $location, forecastInfo: $forecastInfo, units: $units, data: $data, forecastData: $forecastData, source: $source)';
  }
}
