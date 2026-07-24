import '../models/satellite.dart';
import 'api_service.dart';

class SatelliteService {
  SatelliteService(this._api);

  final ApiService _api;

  Future<SatelliteData?> latest() async {
    final data = await _api.get('/satellite');
    if (data is Map) {
      return SatelliteData.fromJson(Map<String, dynamic>.from(data));
    }
    return null;
  }
}
