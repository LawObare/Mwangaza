import 'package:flutter/material.dart';

class AppTheme {
  static const leaf = Color(0xFF236B45);
  static const field = Color(0xFF7A9B45);
  static const sky = Color(0xFF2F80A7);
  static const soil = Color(0xFF6A4E35);
  static const warning = Color(0xFFC26A2E);
  static const danger = Color(0xFFB33636);
  static const ink = Color(0xFF18221C);
  static const muted = Color(0xFF607066);
  static const surface = Color(0xFFF7FAF6);

  static ThemeData get lightTheme {
    final colorScheme = ColorScheme.fromSeed(
      seedColor: leaf,
      primary: leaf,
      secondary: sky,
      tertiary: field,
      error: danger,
      surface: surface,
      brightness: Brightness.light,
    );

    return ThemeData(
      useMaterial3: true,
      colorScheme: colorScheme,
      scaffoldBackgroundColor: surface,
      appBarTheme: const AppBarTheme(
        centerTitle: false,
        elevation: 0,
        backgroundColor: surface,
        foregroundColor: ink,
      ),
      cardTheme: CardTheme(
        elevation: 0,
        margin: EdgeInsets.zero,
        color: Colors.white,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(8),
          side: const BorderSide(color: Color(0xFFE2E8E2)),
        ),
      ),
      inputDecorationTheme: InputDecorationTheme(
        border: OutlineInputBorder(borderRadius: BorderRadius.circular(8)),
        filled: true,
        fillColor: Colors.white,
      ),
    );
  }
}
