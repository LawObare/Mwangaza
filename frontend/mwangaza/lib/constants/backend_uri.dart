class BackendUri {
  static const kijaniApiUri = 'https://api.kijanispace.eu/v1';
  static const mwangazaApiUri = String.fromEnvironment(
    'MWANGAZA_API_URL',
    defaultValue: 'http://10.0.2.2:8080/api',
  );
}
