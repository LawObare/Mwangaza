import '../models/recommendation.dart';
import 'api_service.dart';

class RecommendationService {
  final ApiService _api;

  RecommendationService(this._api);

  Future<List<Recommendation>> getRecommendations() async {
    final data = await _api.getList('/recommendation');
    return data
        .map((e) => Recommendation.fromJson(e as Map<String, dynamic>))
        .toList();
  }

  Future<List<Recommendation>> getRecommendationsForFarm(int farmId) async {
    final data = await _api.getList('/recommendation?farm_id=$farmId');
    return data
        .map((e) => Recommendation.fromJson(e as Map<String, dynamic>))
        .toList();
  }
}