class AppConstants {
  static const apiBaseUrl = String.fromEnvironment(
    'API_BASE_URL',
    defaultValue: 'http://localhost:8080/api',
  );

  static const appName = 'Mwangaza';
  static const supportText = 'Farm alerts from satellite conditions';
}
