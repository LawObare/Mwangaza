import 'package:dio/dio.dart';
import 'package:mwangaza/constants/backend_uri.dart';
import 'package:mwangaza/models/user_model.dart';
import 'package:mwangaza/services/sp_service.dart';

/// Talks to the Mwangaza backend's local account API (`/api/auth/*`), which
/// keeps accounts in the demo store. Sign-up, sign-in and session restore
/// therefore work without an external identity provider.
///
/// The backend answers with the project-wide envelope
/// `{ "success": bool, "data": ..., "error": "..." }`, so every response is
/// unwrapped here before it reaches the UI.
class AuthRemoteRepository {
  AuthRemoteRepository({Dio? dio, SpService? spService})
      : _spService = spService ?? SpService(),
        _dio = dio ??
            Dio(
              BaseOptions(
                baseUrl: BackendUri.mwangazaApiUri,
                headers: {
                  'Content-Type': 'application/json',
                  'Accept': 'application/json',
                },
              ),
            ) {
    _dio.interceptors.add(
      InterceptorsWrapper(
        onRequest: (options, handler) async {
          final token = await _spService.getToken();
          if (token != null && token.isNotEmpty) {
            options.headers['Authorization'] = 'Bearer $token';
          }
          handler.next(options);
        },
        onError: (error, handler) async {
          if (error.response?.statusCode == 401) {
            await _spService.clearAll();
          }
          handler.next(error);
        },
      ),
    );
  }

  final Dio _dio;
  final SpService _spService;

  Future<UserModel> signUp({
    required String email,
    required String password,
  }) async {
    try {
      final response = await _dio.post<Map<String, dynamic>>(
        'auth/register',
        data: {'email': email, 'password': password},
      );
      return _saveSession(response.data);
    } on DioException catch (e) {
      throw _messageOf(e, fallback: 'Registration failed');
    } catch (e) {
      throw e.toString();
    }
  }

  Future<UserModel> login({
    required String email,
    required String password,
  }) async {
    try {
      final response = await _dio.post<Map<String, dynamic>>(
        'auth/login',
        data: {'email': email, 'password': password},
      );
      return _saveSession(response.data);
    } on DioException catch (e) {
      throw _messageOf(e, fallback: 'Login failed');
    } catch (e) {
      throw e.toString();
    }
  }

  Future<UserModel?> getUserData() async {
    final token = await _spService.getToken();
    if (token == null || token.isEmpty) return null;

    try {
      final response = await _dio.get<Map<String, dynamic>>('auth/me');
      return _saveSession(response.data, fallbackToken: token);
    } on DioException {
      // A 401 means the stored token expired; the interceptor cleared it.
      return null;
    } catch (e) {
      return null;
    }
  }

  /// Clears the stored session. The backend issues stateless JWTs, so there is
  /// no server-side session to revoke.
  Future<void> logout() async {
    await _spService.clearAll();
    refreshDio();
  }

  void refreshDio() {
    _dio.options.headers = {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
    };
  }

  /// Unwraps the `{success, data, error}` envelope, persists the token (when
  /// the endpoint returned one) and hands back the user.
  UserModel _saveSession(Map<String, dynamic>? body, {String? fallbackToken}) {
    final data = _unwrap(body);

    final sessionToken = data['token']?.toString() ?? fallbackToken ?? '';
    if (sessionToken.isNotEmpty) {
      _spService.setToken(sessionToken);
    }

    final user = Map<String, dynamic>.from(data['user'] as Map? ?? data);
    if (sessionToken.isNotEmpty) user['token'] = sessionToken;

    return UserModel.fromMap(user);
  }

  Map<String, dynamic> _unwrap(Map<String, dynamic>? body) {
    if (body == null) return <String, dynamic>{};
    final data = body['data'];
    if (data is Map) return Map<String, dynamic>.from(data);
    return Map<String, dynamic>.from(body);
  }

  String _messageOf(DioException error, {required String fallback}) {
    final body = error.response?.data;
    if (body is Map) {
      final message = body['error'] ?? body['message'];
      if (message != null && message.toString().isNotEmpty) {
        return message.toString();
      }
    }
    return error.message ?? fallback;
  }
}
