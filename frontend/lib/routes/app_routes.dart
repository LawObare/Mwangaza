import 'package:flutter/material.dart';
import '../screens/dashboard/dashboard_screen.dart';
import '../screens/farms/farms_screen.dart';
import '../screens/farm_details/farm_details_screen.dart';
import '../screens/map/map_screen.dart';
import '../screens/sms/sms_screen.dart';

class AppRoutes {
  static const String dashboard = '/';
  static const String farms = '/farms';
  static const String farmDetails = '/farms/details';
  static const String map = '/map';
  static const String sms = '/sms';

  static Route<dynamic> generateRoute(RouteSettings settings) {
    switch (settings.name) {
      case dashboard:
        return MaterialPageRoute(builder: (_) => const DashboardScreen());
      case farms:
        return MaterialPageRoute(builder: (_) => const FarmsScreen());
      case farmDetails:
        final farmId = settings.arguments as int;
        return MaterialPageRoute(
            builder: (_) => FarmDetailsScreen(farmId: farmId));
      case map:
        return MaterialPageRoute(builder: (_) => const MapScreen());
      case sms:
        return MaterialPageRoute(builder: (_) => const SmsScreen());
      default:
        return MaterialPageRoute(
            builder: (_) => const Scaffold(
                  body: Center(child: Text('Route not found')),
                ));
    }
  }
}