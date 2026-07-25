import 'package:dio/dio.dart';
import 'package:mwangaza/constants/backend_uri.dart';
import 'package:mwangaza/models/mainweather.dart';
import 'package:mwangaza/services/sp_service.dart';

class WeatherRepository {
  final SpService spService = SpService();
  late final Dio _dio;

  WeatherRepository() {
    _dio = Dio(
      BaseOptions(
        baseUrl: BackendUri.kijanibackendUri,
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json',
        },
      ),
    );

    _dio.interceptors.add(
      InterceptorsWrapper(
        onRequest: (options, handler) async {
          // Use Basic Authentication for Kijani API
          // Username: admin, Password: secret
          final String basicAuth = 'Basic YWRtaW46c2VjcmV0';
          options.headers['Authorization'] = basicAuth;

          // Also add Bearer token if needed for other endpoints
          // If your app uses both Bearer and Basic auth for different endpoints,
          // you can check the URL and apply the appropriate auth
          // Uncomment the following if you also need Bearer token for some endpoints
          /*
          if (options.path.startsWith('/auth/')) {
            final token = await spService.getToken();
            if (token != null && token.isNotEmpty) {
              options.headers['Authorization'] = 'Bearer $token';
            }
          }
          */

          return handler.next(options);
        },
        onError: (DioException e, handler) async {
          if (e.response?.statusCode == 401) {
            // Don't clear token for Basic Auth
            // Only clear if it's a Bearer token error
            if (e.requestOptions.headers['Authorization']?.startsWith(
                  'Bearer ',
                ) ??
                false) {
              await spService.clearAll();
            }
          }
          return handler.next(e);
        },
      ),
    );
  }

  /// Get weather forecast for a specific location
  Future<WeatherResponse> getWeatherForecast({
    required double latitude,
    required double longitude,
    int? forecastDays,
  }) async {
    try {
      final queryParams = <String, dynamic>{
        'lat': latitude, // Note: using 'lat' not 'latitude'
        'lon': longitude, // Note: using 'lon' not 'longitude'
      };

      if (forecastDays != null) {
        queryParams['forecast_days'] = forecastDays;
      }

      final response = await _dio.get(
        '/agro_climate/land', // Updated endpoint path
        queryParameters: queryParams,
      );

      final responseData = Map<String, dynamic>.from(response.data);
      return WeatherResponse.fromMap(responseData);
    } on DioException catch (e) {
      if (e.response != null) {
        final message =
            e.response?.data['message'] ?? 'Failed to fetch weather data';
        throw message.toString();
      } else {
        throw e.message ?? 'Network error occurred';
      }
    } catch (e) {
      throw e.toString();
    }
  }

  /// Get current weather for a location
  Future<WeatherResponse> getCurrentWeather({
    required double latitude,
    required double longitude,
  }) async {
    try {
      final response = await _dio.get(
        '/agro_climate/current', // Updated endpoint path
        queryParameters: {
          'lat': latitude, // Note: using 'lat' not 'latitude'
          'lon': longitude, // Note: using 'lon' not 'longitude'
        },
      );

      final responseData = Map<String, dynamic>.from(response.data);
      return WeatherResponse.fromMap(responseData);
    } on DioException catch (e) {
      if (e.response != null) {
        final message =
            e.response?.data['message'] ?? 'Failed to fetch current weather';
        throw message.toString();
      } else {
        throw e.message ?? 'Network error occurred';
      }
    } catch (e) {
      throw e.toString();
    }
  }

  /// Get weather forecast by location name
  Future<WeatherResponse> getWeatherByLocationName({
    required String locationName,
    int? forecastDays,
  }) async {
    try {
      final queryParams = <String, dynamic>{'location': locationName};

      if (forecastDays != null) {
        queryParams['forecast_days'] = forecastDays;
      }

      final response = await _dio.get(
        '/agro_climate/location', // Updated endpoint path
        queryParameters: queryParams,
      );

      final responseData = Map<String, dynamic>.from(response.data);
      return WeatherResponse.fromMap(responseData);
    } on DioException catch (e) {
      if (e.response != null) {
        final message =
            e.response?.data['message'] ?? 'Failed to fetch weather data';
        throw message.toString();
      } else {
        throw e.message ?? 'Network error occurred';
      }
    } catch (e) {
      throw e.toString();
    }
  }

  /// Get land data for a specific location
  Future<Map<String, dynamic>> getLandData({
    required double latitude,
    required double longitude,
  }) async {
    try {
      final response = await _dio.get(
        '/agro_climate/land',
        queryParameters: {'lat': latitude, 'lon': longitude},
      );

      return Map<String, dynamic>.from(response.data);
    } on DioException catch (e) {
      if (e.response != null) {
        final message =
            e.response?.data['message'] ?? 'Failed to fetch land data';
        throw message.toString();
      } else {
        throw e.message ?? 'Network error occurred';
      }
    } catch (e) {
      throw e.toString();
    }
  }
}
