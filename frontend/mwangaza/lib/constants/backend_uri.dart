class BackendUri {
  /// Kijani v1 REST API — used for auth (historical, now replaced by Mwangaza backend)
  static const kijaniApiUri = 'https://api.kijanispace.eu/v1';

  /// Kijani agro-climate / satellite data API — used by WeatherRepository
  static const kijanibackendUri = 'https://api.kijanispace.eu';

  /// Mwangaza own backend — used for auth (register / login)
  static const mwangazaApiUri = String.fromEnvironment(
    'MWANGAZA_API_URL',
    defaultValue: 'http://10.0.2.2:8080/api',
  );
}
