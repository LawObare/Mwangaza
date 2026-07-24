import 'package:flutter/material.dart';

import 'core/theme.dart';
import 'routes/app_routes.dart';

class MwangazaApp extends StatelessWidget {
  const MwangazaApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Mwangaza',
      debugShowCheckedModeBanner: false,
      theme: AppTheme.lightTheme,
      initialRoute: AppRoutes.dashboard,
      onGenerateRoute: AppRoutes.generateRoute,
    );
  }
}
