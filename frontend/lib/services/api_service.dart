import 'dart:convert';

import 'package:http/http.dart' as http;

import '../core/constants.dart';

class ApiService {
  ApiService({http.Client? client, String? baseUrl})
      : _client = client ?? http.Client(),
        baseUrl = baseUrl ?? AppConstants.apiBaseUrl;

  final http.Client _client;
  final String baseUrl;

  Future<dynamic> get(String path) async {
    final response = await _client.get(_uri(path));
    return _decode(response);
  }

  Future<dynamic> post(String path, [Map<String, dynamic>? body]) async {
    final response = await _client.post(
      _uri(path),
      headers: const {'Content-Type': 'application/json'},
      body: body == null ? null : jsonEncode(body),
    );
    return _decode(response);
  }

  Uri _uri(String path) {
    final normalized = path.startsWith('/') ? path : '/$path';
    return Uri.parse('$baseUrl$normalized');
  }

  dynamic _decode(http.Response response) {
    final decoded = response.body.isEmpty ? null : jsonDecode(response.body);
    if (response.statusCode >= 400) {
      final message = decoded is Map ? decoded['error']?.toString() : null;
      throw ApiException(message ?? 'Request failed (${response.statusCode})');
    }
    if (decoded is Map && decoded.containsKey('success')) {
      if (decoded['success'] == false) {
        throw ApiException(decoded['error']?.toString() ?? 'Request failed');
      }
      return decoded['data'];
    }
    return decoded;
  }

  void dispose() => _client.close();
}

class ApiException implements Exception {
  const ApiException(this.message);

  final String message;

  @override
  String toString() => message;
}
