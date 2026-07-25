// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:convert';

class Location {
  final double latitude;
  final double longitude;
  final int elevation;
  final String timezone;

  Location({
    required this.latitude,
    required this.longitude,
    required this.elevation,
    required this.timezone,
  });

  Location copyWith({
    double? latitude,
    double? longitude,
    int? elevation,
    String? timezone,
  }) {
    return Location(
      latitude: latitude ?? this.latitude,
      longitude: longitude ?? this.longitude,
      elevation: elevation ?? this.elevation,
      timezone: timezone ?? this.timezone,
    );
  }

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'latitude': latitude,
      'longitude': longitude,
      'elevation': elevation,
      'timezone': timezone,
    };
  }

  factory Location.fromMap(Map<String, dynamic> map) {
    return Location(
      latitude: map['latitude']?.toDouble() ?? 0.0,
      longitude: map['longitude']?.toDouble() ?? 0.0,
      elevation: map['elevation']?.toInt() ?? 0,
      timezone: map['timezone'] ?? '',
    );
  }

  String toJson() => json.encode(toMap());

  factory Location.fromJson(String source) =>
      Location.fromMap(json.decode(source) as Map<String, dynamic>);

  @override
  String toString() {
    return 'Location(latitude: $latitude, longitude: $longitude, elevation: $elevation, timezone: $timezone)';
  }

  @override
  bool operator ==(covariant Location other) {
    if (identical(this, other)) return true;

    return other.latitude == latitude &&
        other.longitude == longitude &&
        other.elevation == elevation &&
        other.timezone == timezone;
  }

  @override
  int get hashCode {
    return latitude.hashCode ^
        longitude.hashCode ^
        elevation.hashCode ^
        timezone.hashCode;
  }
}
