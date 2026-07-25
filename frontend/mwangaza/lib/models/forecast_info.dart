// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:convert';

class ForecastInfo {
  final int forecastDays;
  final String modelRun;

  ForecastInfo({required this.forecastDays, required this.modelRun});

  ForecastInfo copyWith({int? forecastDays, String? modelRun}) {
    return ForecastInfo(
      forecastDays: forecastDays ?? this.forecastDays,
      modelRun: modelRun ?? this.modelRun,
    );
  }

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'forecast_days': forecastDays,
      'model_run': modelRun,
    };
  }

  factory ForecastInfo.fromMap(Map<String, dynamic> map) {
    return ForecastInfo(
      forecastDays: map['forecast_days']?.toInt() ?? 0,
      modelRun: map['model_run'] ?? '',
    );
  }

  String toJson() => json.encode(toMap());

  factory ForecastInfo.fromJson(String source) =>
      ForecastInfo.fromMap(json.decode(source) as Map<String, dynamic>);

  @override
  String toString() {
    return 'ForecastInfo(forecastDays: $forecastDays, modelRun: $modelRun)';
  }

  @override
  bool operator ==(covariant ForecastInfo other) {
    if (identical(this, other)) return true;

    return other.forecastDays == forecastDays && other.modelRun == modelRun;
  }

  @override
  int get hashCode {
    return forecastDays.hashCode ^ modelRun.hashCode;
  }
}
