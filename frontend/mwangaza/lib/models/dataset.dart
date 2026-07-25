// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:convert';

class Dataset {
  final int landuseCode;
  final String landuse;

  Dataset({required this.landuseCode, required this.landuse});

  Map<String, dynamic> toMap() {
    return <String, dynamic>{'landuse_code': landuseCode, 'landuse': landuse};
  }

  factory Dataset.fromMap(Map<String, dynamic> map) {
    return Dataset(
      landuseCode: map['landuse_code']?.toInt() ?? 0,
      landuse: map['landuse'] ?? '',
    );
  }

  String toJson() => json.encode(toMap());

  factory Dataset.fromJson(String source) =>
      Dataset.fromMap(json.decode(source) as Map<String, dynamic>);
}
