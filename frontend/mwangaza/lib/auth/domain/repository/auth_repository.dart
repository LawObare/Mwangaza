import 'package:dio/dio.dart';
import 'package:mwangaza/constants/backend_uri.dart';
import 'package:mwangaza/models/user_model.dart';
import 'package:mwangaza/services/sp_service.dart';

class AuthRemoteRepository {
  final spService = SpService();

  late final Dio _dio;

  AuthRemoteRepository() {
    _dio = Dio(
      BaseOptions(
        baseUrl: BackendUri.kijaniApiUri,
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json',
        },
      ),
    );

    _dio.interceptors.add(
      InterceptorsWrapper(
        onRequest: (options, handler) async {
          final token = await spService.getToken();
          if (token != null && token.isNotEmpty) {
            options.headers['Authorization'] = 'Bearer $token';
          }
          return handler.next(options);
        },
        onError: (DioException e, handler) async {
          if (e.response?.statusCode == 401) {
            await spService.setToken('');
          }
          return handler.next(e);
        },
      ),
    );
  }

  Future<UserModel> signUp({
    //required String name,
    required String email,
    required String password,
  }) async {
    try {
      final response = await _dio.post(
        'auth/register',
        data: {
          //'name': name,
          'email': email,
          'password': password,
        },
      );

      final responseData = Map<String, dynamic>.from(response.data);

      return UserModel.fromMap(responseData);
    } on DioException catch (e) {
      if (e.response != null) {
        if (e.response?.statusCode == 422) {
          final errors = e.response?.data['errors'];
          if (errors != null) {
            final errorMap = Map<String, dynamic>.from(errors);

            final errorMessages = errorMap.values
                .expand((x) => x as List)
                .join('\n');
            throw errorMessages;
          }
        }

        final message = e.response?.data['detail'] ?? e.response?.data['message'];
        throw message?.toString() ?? 'Registration failed';
      } else {
        throw e.message ?? 'Network error occurred';
      }
    } catch (e) {
      throw e.toString();
    }
  }

  Future<UserModel> login({
    required String email,
    required String password,
  }) async {
    try {
      final response = await _dio.post(
        'auth/login',
        data: {'email': email, 'password': password},
      );

      final responseData = Map<String, dynamic>.from(response.data);

      final accessToken = responseData['access_token']?.toString();
      final refreshToken = responseData['refresh_token']?.toString();
      if (accessToken == null || refreshToken == null) {
        throw 'Kijani returned an invalid login response';
      }
      await spService.setToken(accessToken);
      await spService.setRefreshToken(refreshToken);

      final userResponse = await _dio.get('auth/me');
      final userData = Map<String, dynamic>.from(userResponse.data)..['token'] = accessToken;
      return UserModel.fromMap(userData);
    } on DioException catch (e) {
      if (e.response != null) {
        final message = e.response?.data['detail'] ?? e.response?.data['message'];
        throw message?.toString() ?? 'Login failed';
      } else {
        throw e.message ?? 'Network error occurred';
      }
    } catch (e) {
      throw e.toString();
    }
  }

  Future<UserModel?> getUserData() async {
    try {
      final token = await spService.getToken();
      if (token == null || token.isEmpty) {
        return null;
      }

      final response = await _dio.get('auth/me');

      final responseData = Map<String, dynamic>.from(response.data);

      final userData = {...responseData, 'token': token};

      return UserModel.fromMap(userData);
    } on DioException catch (e) {
      if (e.response?.statusCode == 401) {
        await spService.setToken('');
        return null;
      }
      return null;
    } catch (e) {
      return null;
    }
  }

  Future<void> logout() async {
    try {
      final refreshToken = await spService.getRefreshToken();
      if (refreshToken != null && refreshToken.isNotEmpty) {
        await _dio.post('auth/logout', data: {'refresh_token': refreshToken});
      }
    } finally {
      // Always clear local token
      await spService.clearAll();
      refreshDio();
    }
  }

  void refreshDio() {
    _dio.options.headers = {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
    };
  }
}
