// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:convert';

import 'package:mwangaza/models/dataset.dart';
import 'package:mwangaza/models/vegetation_indices.dart';

class WeatherData {
  final Dataset dataset;
  final VegetationIndices vegetationIndices;

  WeatherData({required this.dataset, required this.vegetationIndices});

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'dataset': dataset.toMap(),
      'vegetation_indices': vegetationIndices.toMap(),
    };
  }

  factory WeatherData.fromMap(Map<String, dynamic> map) {
    return WeatherData(
      dataset: Dataset.fromMap(map['dataset'] ?? {}),
      vegetationIndices: VegetationIndices.fromMap(
        map['vegetation_indices'] ?? {},
      ),
    );
  }

  String toJson() => json.encode(toMap());

  factory WeatherData.fromJson(String source) =>
      WeatherData.fromMap(json.decode(source) as Map<String, dynamic>);
}
