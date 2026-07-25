import 'package:shared_preferences/shared_preferences.dart';

class SpService {
  static const String _tokenKey = 'auth_token';
  static const String _refreshTokenKey = 'refresh_token';

  final SharedPreferencesAsync _prefs = SharedPreferencesAsync();

  // Token methods
  Future<void> setToken(String token) async {
    await _prefs.setString(_tokenKey, token);
  }

  Future<String?> getToken() async {
    return await _prefs.getString(_tokenKey);
  }

  Future<void> setRefreshToken(String token) async {
    await _prefs.setString(_refreshTokenKey, token);
  }

  Future<String?> getRefreshToken() async {
    return await _prefs.getString(_refreshTokenKey);
  }

  // Clear all data on logout
  Future<void> clearAll() async {
    await _prefs.remove(_tokenKey);
    await _prefs.remove(_refreshTokenKey);
  }
}
