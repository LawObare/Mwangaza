// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:convert';

class VegetationIndices {
  final String date;
  final double ndvi;
  final double vegetationHealthIndex;
  final double vegetationConditionIndex;

  VegetationIndices({
    required this.date,
    required this.ndvi,
    required this.vegetationHealthIndex,
    required this.vegetationConditionIndex,
  });

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'date': date,
      'NDVI': ndvi,
      'vegetation_health_index': vegetationHealthIndex,
      'vegetation_condition_index': vegetationConditionIndex,
    };
  }

  factory VegetationIndices.fromMap(Map<String, dynamic> map) {
    return VegetationIndices(
      date: map['date'] ?? '',
      ndvi: map['NDVI']?.toDouble() ?? 0.0,
      vegetationHealthIndex: map['vegetation_health_index']?.toDouble() ?? 0.0,
      vegetationConditionIndex:
          map['vegetation_condition_index']?.toDouble() ?? 0.0,
    );
  }

  String toJson() => json.encode(toMap());

  factory VegetationIndices.fromJson(String source) =>
      VegetationIndices.fromMap(json.decode(source) as Map<String, dynamic>);
}
