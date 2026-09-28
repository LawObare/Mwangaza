import 'package:dio/dio.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mwangaza/constants/backend_uri.dart';

void main() {
  group('BackendUri', () {
    test('the Mwangaza base URL always ends with a slash', () {
      // Dio joins baseUrl and path by concatenation, so a missing trailing
      // slash turns `farms` into `/apifarms`.
      expect(BackendUri.mwangazaApiUri, endsWith('/'));
    });

    test('Mwangaza paths resolve under /api', () {
      final request = RequestOptions(
        baseUrl: BackendUri.mwangazaApiUri,
        path: 'farms',
      );

      expect(request.uri.toString(), 'http://10.0.2.2:8080/api/farms');
    });
  });
}
