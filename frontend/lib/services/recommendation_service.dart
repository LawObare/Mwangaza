import '../models/recommendation.dart';
import 'api_service.dart';

class RecommendationService {
  RecommendationService(this._api);

  final ApiService _api;

  Future<List<Recommendation>> listRecommendations({int? farmId}) async {
    final query = farmId == null ? '' : '?farm_id=$farmId';
    final data = await _api.get('/recommendation$query');
    return Recommendation.listFromFlexibleData(data);
  }

  Future<List<Recommendation>> generate({int? farmId}) async {
    final query = farmId == null ? '' : '?farm_id=$farmId';
    final data = await _api.post('/recommendation$query');
    return Recommendation.listFromFlexibleData(data);
  }
}
