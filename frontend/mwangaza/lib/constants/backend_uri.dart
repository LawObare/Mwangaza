class BackendUri {
  static const _mwangazaApiUri = String.fromEnvironment(
    'MWANGAZA_API_URL',
    defaultValue: 'http://10.0.2.2:8080/api',
  );

  /// Base URL of the Mwangaza backend.
  ///
  /// Dio joins [BaseOptions.baseUrl] with a relative path by plain string
  /// concatenation, so the base has to end with `/`. Without it, `farms`
  /// would resolve to `/apifarms` instead of `/api/farms`.
  static String get mwangazaApiUri =>
      _mwangazaApiUri.endsWith('/') ? _mwangazaApiUri : '$_mwangazaApiUri/';
}
