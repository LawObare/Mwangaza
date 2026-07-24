import 'package:flutter/material.dart';

import '../../core/constants.dart';
import '../../core/theme.dart';
import '../../models/farm.dart';
import '../../models/recommendation.dart';
import '../../models/satellite.dart';
import '../../routes/app_routes.dart';
import '../../services/api_service.dart';
import '../../services/farm_service.dart';
import '../../services/recommendation_service.dart';
import '../../services/satellite_service.dart';
import '../../widgets/loading_widget.dart';
import '../../widgets/recommendation_card.dart';
import '../../widgets/satellite_card.dart';
import '../../widgets/stat_card.dart';

class DashboardScreen extends StatefulWidget {
  const DashboardScreen({super.key});

  @override
  State<DashboardScreen> createState() => _DashboardScreenState();
}

class _DashboardScreenState extends State<DashboardScreen> {
  final _api = ApiService();
  late final _farms = FarmService(_api);
  late final _satellite = SatelliteService(_api);
  late final _recommendations = RecommendationService(_api);

  bool _loading = true;
  String _error = '';
  List<Farm> _farmList = const [];
  List<Recommendation> _alerts = const [];
  SatelliteData? _snapshot;

  @override
  void initState() {
    super.initState();
    _load();
  }

  @override
  void dispose() {
    _api.dispose();
    super.dispose();
  }

  Future<void> _load() async {
    setState(() {
      _loading = true;
      _error = '';
    });
    try {
      final results = await Future.wait([
        _farms.listFarms(),
        _recommendations.listRecommendations(),
        _satellite.latest(),
      ]);
      setState(() {
        _farmList = results[0] as List<Farm>;
        _alerts = results[1] as List<Recommendation>;
        _snapshot = results[2] as SatelliteData?;
        _loading = false;
      });
    } catch (error) {
      setState(() {
        _error = error.toString();
        _loading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text(AppConstants.appName)),
      drawer: const _AppDrawer(currentRoute: AppRoutes.dashboard),
      body: _loading
          ? const LoadingWidget(message: 'Loading farm intelligence')
          : _error.isNotEmpty
              ? ErrorState(message: _error, onRetry: _load)
              : RefreshIndicator(
                  onRefresh: _load,
                  child: ListView(
                    padding: const EdgeInsets.all(16),
                    children: [
                      Text('Operations Dashboard', style: Theme.of(context).textTheme.headlineSmall?.copyWith(fontWeight: FontWeight.w800)),
                      const SizedBox(height: 4),
                      const Text(AppConstants.supportText, style: TextStyle(color: AppTheme.muted)),
                      const SizedBox(height: 16),
                      _StatsGrid(farms: _farmList, snapshot: _snapshot, alerts: _alerts),
                      const SizedBox(height: 16),
                      if (_snapshot != null) SatelliteCard(data: _snapshot!),
                      const SizedBox(height: 16),
                      _SectionHeader(
                        title: 'Latest Alerts',
                        actionLabel: 'Generate',
                        onAction: _generateAlerts,
                      ),
                      const SizedBox(height: 8),
                      if (_alerts.isEmpty)
                        const _EmptyPanel(message: 'No recommendations yet. Generate alerts after the backend is running.')
                      else
                        ..._alerts.take(4).map((rec) => Padding(
                              padding: const EdgeInsets.only(bottom: 10),
                              child: RecommendationCard(recommendation: rec),
                            )),
                    ],
                  ),
                ),
    );
  }

  Future<void> _generateAlerts() async {
    try {
      final generated = await _recommendations.generate();
      setState(() => _alerts = generated);
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(const SnackBar(content: Text('Recommendations generated')));
      }
    } catch (error) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(error.toString())));
      }
    }
  }
}

class _StatsGrid extends StatelessWidget {
  const _StatsGrid({required this.farms, required this.snapshot, required this.alerts});

  final List<Farm> farms;
  final SatelliteData? snapshot;
  final List<Recommendation> alerts;

  @override
  Widget build(BuildContext context) {
    return GridView.count(
      shrinkWrap: true,
      physics: const NeverScrollableScrollPhysics(),
      crossAxisCount: MediaQuery.sizeOf(context).width > 760 ? 4 : 2,
      childAspectRatio: 1.2,
      crossAxisSpacing: 10,
      mainAxisSpacing: 10,
      children: [
        StatCard(title: 'Farms', value: farms.length.toString(), icon: Icons.agriculture_outlined, color: AppTheme.leaf),
        StatCard(title: 'Open alerts', value: alerts.length.toString(), icon: Icons.notifications_active_outlined, color: AppTheme.warning),
        StatCard(title: 'Soil moisture', value: '${(snapshot?.soilMoisture ?? 0).toStringAsFixed(0)}%', icon: Icons.water_drop_outlined, color: AppTheme.sky),
        StatCard(title: 'Temperature', value: '${(snapshot?.temperature ?? 0).toStringAsFixed(0)} C', icon: Icons.thermostat, color: AppTheme.danger),
      ],
    );
  }
}

class _SectionHeader extends StatelessWidget {
  const _SectionHeader({required this.title, required this.actionLabel, required this.onAction});

  final String title;
  final String actionLabel;
  final VoidCallback onAction;

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Expanded(child: Text(title, style: Theme.of(context).textTheme.titleLarge?.copyWith(fontWeight: FontWeight.w800))),
        FilledButton.icon(onPressed: onAction, icon: const Icon(Icons.auto_awesome), label: Text(actionLabel)),
      ],
    );
  }
}

class _EmptyPanel extends StatelessWidget {
  const _EmptyPanel({required this.message});

  final String message;

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(18),
        child: Text(message, textAlign: TextAlign.center),
      ),
    );
  }
}

class _AppDrawer extends StatelessWidget {
  const _AppDrawer({required this.currentRoute});

  final String currentRoute;

  @override
  Widget build(BuildContext context) {
    return NavigationDrawer(
      selectedIndex: [AppRoutes.dashboard, AppRoutes.farms, AppRoutes.map, AppRoutes.sms].indexOf(currentRoute),
      onDestinationSelected: (index) {
        final route = [AppRoutes.dashboard, AppRoutes.farms, AppRoutes.map, AppRoutes.sms][index];
        Navigator.pop(context);
        if (route != currentRoute) Navigator.pushReplacementNamed(context, route);
      },
      children: const [
        Padding(
          padding: EdgeInsets.fromLTRB(28, 28, 16, 12),
          child: Text('Mwangaza', style: TextStyle(fontSize: 24, fontWeight: FontWeight.w800)),
        ),
        NavigationDrawerDestination(icon: Icon(Icons.dashboard_outlined), label: Text('Dashboard')),
        NavigationDrawerDestination(icon: Icon(Icons.agriculture_outlined), label: Text('Farms')),
        NavigationDrawerDestination(icon: Icon(Icons.map_outlined), label: Text('Map')),
        NavigationDrawerDestination(icon: Icon(Icons.sms_outlined), label: Text('SMS')),
      ],
    );
  }
}
