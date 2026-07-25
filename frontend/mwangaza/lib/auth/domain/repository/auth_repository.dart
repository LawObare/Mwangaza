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
          print('🌐 Request: ${options.method} ${options.path}');
          print('📤 Headers: ${options.headers}');
          print('📤 Data: ${options.data}');
          
          final token = await spService.getToken();
          if (token != null && token.isNotEmpty) {
            options.headers['Authorization'] = 'Bearer $token';
            print('🔑 Token added to request');
          } else {
            print('⚠️ No token found');
          }
          return handler.next(options);
        },
        onResponse: (response, handler) {
          print('✅ Response: ${response.statusCode}');
          print('📥 Response data: ${response.data}');
          return handler.next(response);
        },
        onError: (DioException e, handler) async {
          print('❌ Dio Error:');
          print('  Status Code: ${e.response?.statusCode}');
          print('  Message: ${e.message}');
          print('  Response: ${e.response?.data}');
          print('  Request: ${e.requestOptions.method} ${e.requestOptions.path}');
          
          if (e.response?.statusCode == 401) {
            print('🔄 Token expired, clearing token');
            await spService.setToken('');
          }
          return handler.next(e);
        },
      ),
    );
  }

  Future<UserModel> signUp({
    required String email,
    required String password,
  }) async {
    print('📝 Starting signup process...');
    print('📧 Email: $email');
    
    try {
      final response = await _dio.post(
        '/auth/register',
        data: {
          'email': email,
          'password': password,
        },
      );

      print('✅ Signup response received');
      
      final responseData = Map<String, dynamic>.from(response.data);
      print('📥 Response data: $responseData');

      if (responseData['token'] != null) {
        await spService.setToken(responseData['token'].toString());
        print('🔑 Token saved successfully');
      } else {
        print('⚠️ No token in response');
      }

      final userMap = Map<String, dynamic>.from(responseData['user'] ?? {});
      
      final userData = {
        ...userMap,
        'token': responseData['token']?.toString() ?? '',
      };

      print('👤 User data: $userData');
      return UserModel.fromMap(userData);
    } on DioException catch (e) {
      print('❌ Signup DioException:');
      print('  Status: ${e.response?.statusCode}');
      print('  Data: ${e.response?.data}');
      print('  Message: ${e.message}');
      
      if (e.response != null) {
        if (e.response?.statusCode == 422) {
          final errors = e.response?.data['errors'];
          if (errors != null) {
            final errorMap = Map<String, dynamic>.from(errors);
            final errorMessages = errorMap.values
                .expand((x) => x as List)
                .join('\n');
            print('📝 Validation errors: $errorMessages');
            throw errorMessages;
          }
        }

        final message = e.response?.data['message'];
        throw message?.toString() ?? 'Registration failed';
      } else {
        throw e.message ?? 'Network error occurred';
      }
    } catch (e) {
      print('❌ Unexpected error in signup: $e');
      throw e.toString();
    }
  }

  Future<UserModel> login({
    required String email,
    required String password,
  }) async {
    print('🔐 Starting login process...');
    print('📧 Email: $email');
    print('🔑 Password: ${'*' * password.length}');
    
    try {
      final response = await _dio.post(
        '/auth/login',
        data: {'email': email, 'password': password},
      );

      print('✅ Login response received');
      print('📥 Status code: ${response.statusCode}');
      
      final responseData = Map<String, dynamic>.from(response.data);
      print('📥 Response data: $responseData');

      if (responseData['token'] != null) {
        await spService.setToken(responseData['token'].toString());
        print('🔑 Token saved successfully');
        print('🔑 Token: ${responseData['token'].toString().substring(0, min(20, responseData['token'].toString().length))}...');
      } else {
        print('⚠️ No token in response');
      }

      final userMap = Map<String, dynamic>.from(responseData['user'] ?? {});
      print('👤 User map: $userMap');
      
      final userData = {
        ...userMap,
        'token': responseData['token']?.toString() ?? '',
      };

      print('✅ Login successful for user: ${userData['email']}');
      return UserModel.fromMap(userData);
    } on DioException catch (e) {
      print('❌ Login DioException:');
      print('  Status: ${e.response?.statusCode}');
      print('  Data: ${e.response?.data}');
      print('  Message: ${e.message}');
      print('  Request: ${e.requestOptions.method} ${e.requestOptions.path}');
      
      if (e.response != null) {
        final message = e.response?.data['message'];
        print('📝 Server message: $message');
        throw message?.toString() ?? 'Login failed';
      } else {
        print('🌐 Network error: ${e.message}');
        throw e.message ?? 'Network error occurred';
      }
    } catch (e) {
      print('❌ Unexpected error in login: $e');
      throw e.toString();
    }
  }

  Future<UserModel?> getUserData() async {
    print('👤 Getting user data...');
    try {
      final token = await spService.getToken();
      if (token == null || token.isEmpty) {
        print('⚠️ No token found');
        return null;
      }
      print('🔑 Token found: ${token.substring(0, min(20, token.length))}...');

      final response = await _dio.get('/auth/user');
      print('✅ User data response received');

      final responseData = Map<String, dynamic>.from(response.data);
      print('👤 User data: $responseData');

      final userData = {...responseData, 'token': token};

      return UserModel.fromMap(userData);
    } on DioException catch (e) {
      print('❌ Error getting user data:');
      print('  Status: ${e.response?.statusCode}');
      print('  Data: ${e.response?.data}');
      
      if (e.response?.statusCode == 401) {
        print('🔄 Token expired, clearing token');
        await spService.setToken('');
        return null;
      }
      return null;
    } catch (e) {
      print('❌ Unexpected error getting user data: $e');
      return null;
    }
  }

  Future<void> logout() async {
    print('🚪 Logging out...');
    try {
      await _dio.post('/auth/logout');
      print('✅ Logout API call successful');
    } catch (e) {
      print('⚠️ Error during logout API call: $e');
    } finally {
      print('🧹 Clearing local token');
      await spService.setToken('');
      refreshDio();
      print('✅ Logout complete');
    }
  }

  void refreshDio() {
    print('🔄 Refreshing Dio headers');
    _dio.options.headers = {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
    };
  }
}

// Helper function for truncating strings
int min(int a, int b) => a < b ? a : b;
