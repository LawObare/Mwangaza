import '../models/farm.dart';
import 'api_service.dart';

class FarmService {
  FarmService(this._api);

  final ApiService _api;

  Future<List<Farm>> listFarms() async {
    final data = await _api.get('/farms');
    return (data as List? ?? const [])
        .whereType<Map>()
        .map((item) => Farm.fromJson(Map<String, dynamic>.from(item)))
        .toList();
  }

  Future<Farm> getFarm(int id) async {
    final data = await _api.get('/farms/$id');
    return Farm.fromJson(Map<String, dynamic>.from(data as Map));
  }

  Future<Farm> createFarm(Farm farm) async {
    final data = await _api.post('/farms', farm.toCreateJson());
    return Farm.fromJson(Map<String, dynamic>.from(data as Map));
  }
}
