import '../models/farm.dart';
import 'api_service.dart';

class FarmService {
  final ApiService _api;

  FarmService(this._api);

  Future<List<Farm>> getFarms() async {
    final data = await _api.getList('/farms');
    return data.map((e) => Farm.fromJson(e as Map<String, dynamic>)).toList();
  }

  Future<Farm> getFarm(int id) async {
    final data = await _api.get('/farms/$id');
    return Farm.fromJson(data);
  }
}