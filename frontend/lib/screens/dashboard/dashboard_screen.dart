import 'package:flutter/material.dart';
import '../../widgets/stat_card.dart';
import '../../widgets/satellite_card.dart';
import '../../widgets/recommendation_card.dart';
import '../../widgets/loading_widget.dart';
import '../../services/api_service.dart';
import '../../services/satellite_service.dart';
import '../../services/recommendation_service.dart';
import '../../models/satellite.dart';
import '../../models/recommendation.dart';

class DashboardScreen extends StatefulWidget {
  const DashboardScreen({super.key});

  @override
  State<DashboardScreen> createState() => _DashboardScreenState();
}

class _DashboardScreenState extends State<DashboardScreen> {
  final ApiService _api = ApiService();
  SatelliteData? _satelliteData;
  List<Recommendation>? _recommendations;
  bool _loading = true;

  @override
  void initState() {
    super.initState();
    _loadData();
  }

  Future<void> _loadData() async {
    final satelliteService = SatelliteService(_api);
    final recommendationService = RecommendationService(_api);

    try {
      final results = await Future.wait([
        satelliteService.getSatelliteData(),
        recommendationService.getRecommendations(),
      ]);
      setState(() {
        _satelliteData = results[0] as SatelliteData;
        _recommendations = results[1] as List<Recommendation>;
        _loading = false;
      });
    } catch (e) {
      setState(() => _loading = false);
    }
  }

  @override
  void dispose() {
    _api.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Mwangaza Dashboard')),
      drawer: _buildDrawer(context),
      body: _loading
          ? const LoadingWidget(message: 'Loading dashboard...')
          : RefreshIndicator(
              onRefresh: _loadData,
              child: SingleChildScrollView(
                physics: const AlwaysScrollableScrollPhysics(),
                padding: const EdgeInsets.all(16),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    if (_satelliteData != null) ...[
                      GridView.count(
                        crossAxisCount: 2,
                        shrinkWrap: true,
                        physics: const NeverScrollableScrollPhysics(),
                        childAspectRatio: 1.5,
                        children: [
                          StatCard(
                              title: 'Soil Moisture',
                              value: '${_satelliteData!.soilMoisture}%',
                              icon: Icons.water_drop),
                          StatCard(
                              title: 'Temperature',
                              value: '${_satelliteData!.temperature}°C',
                              icon: Icons.thermostat),
                          StatCard(
                              title: 'Rain Prob',
                              value: '${_satelliteData!.rainProbability}%',
                              icon: Icons.umbrella),
                          StatCard(
                              title: 'NDVI',
                              value: '${_satelliteData!.ndvi}',
                              icon: Icons.eco),
                        ],
                      ),
                      const SizedBox(height: 16),
                      SatelliteCard(data: _satelliteData!),
                    ],
                    if (_recommendations != null) ...[
                      const SizedBox(height: 16),
                      const Text('Recent Recommendations',
                          style: TextStyle(
                              fontSize: 16, fontWeight: FontWeight.bold)),
                      const SizedBox(height: 8),
                      ..._recommendations!
                          .take(3)
                          .map((r) => RecommendationCard(recommendation: r)),
                    ],
                  ],
                ),
              ),
            ),
    );
  }

  Widget _buildDrawer(BuildContext context) {
    return Drawer(
      child: ListView(
        padding: EdgeInsets.zero,
        children: [
          const DrawerHeader(
            decoration: BoxDecoration(color: Colors.teal),
            child: Text('Mwangaza', style: TextStyle(fontSize: 24, color: Colors.white)),
          ),
          ListTile(
            leading: const Icon(Icons.dashboard),
            title: const Text('Dashboard'),
            onTap: () => Navigator.pushReplacementNamed(context, '/'),
          ),
          ListTile(
            leading: const Icon(Icons.agriculture),
            title: const Text('Farms'),
            onTap: () => Navigator.pushReplacementNamed(context, '/farms'),
          ),
          ListTile(
            leading: const Icon(Icons.map),
            title: const Text('Map'),
            onTap: () => Navigator.pushReplacementNamed(context, '/map'),
          ),
          ListTile(
            leading: const Icon(Icons.sms),
            title: const Text('SMS History'),
            onTap: () => Navigator.pushReplacementNamed(context, '/sms'),
          ),
        ],
      ),
    );
  }
}