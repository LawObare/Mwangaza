import 'package:dio/dio.dart';
import 'package:mwangaza/constants/backend_uri.dart';
import 'package:mwangaza/services/sp_service.dart';

/// Client for Mwangaza's API. Kijani remains the authentication provider; its
/// token is attached so the backend can retrieve the user's farm data.
class MwangazaApiClient {
  MwangazaApiClient({Dio? dio, SpService? spService})
      : _spService = spService ?? SpService(),
        _dio = dio ?? Dio(BaseOptions(baseUrl: BackendUri.mwangazaApiUri));

  final Dio _dio;
  final SpService _spService;

  Future<Response<T>> get<T>(String path, {Map<String, dynamic>? queryParameters}) {
    return _authorizedOptions().then(
      (options) => _dio.get<T>(path, queryParameters: queryParameters, options: options),
    );
  }

  Future<Response<T>> post<T>(String path, {Object? data}) {
    return _authorizedOptions().then((options) => _dio.post<T>(path, data: data, options: options));
  }

  Future<Options> _authorizedOptions() async {
    final token = await _spService.getToken();
    return Options(
      headers: {
        'Accept': 'application/json',
        if (token != null && token.isNotEmpty) 'Authorization': 'Bearer $token',
      },
    );
  }
}
