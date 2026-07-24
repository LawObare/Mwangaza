import '../models/satellite.dart';
import 'api_service.dart';

class SatelliteService {
  final ApiService _api;

  SatelliteService(this._api);

  Future<SatelliteData> getSatelliteData() async {
    final data = await _api.get('/satellite');
    return SatelliteData.fromJson(data);
  }
}