import 'package:flutter/material.dart';

import '../../core/theme.dart';
import '../../models/farm.dart';
import '../../models/recommendation.dart';
import '../../services/api_service.dart';
import '../../services/farm_service.dart';
import '../../services/recommendation_service.dart';
import '../../services/sms_service.dart';
import '../../widgets/loading_widget.dart';
import '../../widgets/recommendation_card.dart';
import '../../widgets/stat_card.dart';

class FarmDetailsScreen extends StatefulWidget {
  const FarmDetailsScreen({super.key, required this.farmId});

  final int? farmId;

  @override
  State<FarmDetailsScreen> createState() => _FarmDetailsScreenState();
}

class _FarmDetailsScreenState extends State<FarmDetailsScreen> {
  final _api = ApiService();
  late final _farms = FarmService(_api);
  late final _recommendations = RecommendationService(_api);
  late final _sms = SmsService(_api);

  bool _loading = true;
  String _error = '';
  Farm? _farm;
  List<Recommendation> _alerts = const [];

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
    final id = widget.farmId;
    if (id == null) {
      setState(() {
        _loading = false;
        _error = 'Missing farm ID';
      });
      return;
    }
    setState(() {
      _loading = true;
      _error = '';
    });
    try {
      final results = await Future.wait([
        _farms.getFarm(id),
        _recommendations.listRecommendations(farmId: id),
      ]);
      setState(() {
        _farm = results[0] as Farm;
        _alerts = results[1] as List<Recommendation>;
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
    final farm = _farm;
    return Scaffold(
      appBar: AppBar(title: Text(farm?.name ?? 'Farm details')),
      body: _loading
          ? const LoadingWidget(message: 'Loading farm')
          : _error.isNotEmpty
              ? ErrorState(message: _error, onRetry: _load)
              : RefreshIndicator(
                  onRefresh: _load,
                  child: ListView(
                    padding: const EdgeInsets.all(16),
                    children: [
                      _FarmHeader(farm: farm!),
                      const SizedBox(height: 14),
                      GridView.count(
                        shrinkWrap: true,
                        physics: const NeverScrollableScrollPhysics(),
                        crossAxisCount: 2,
                        childAspectRatio: 1.25,
                        crossAxisSpacing: 10,
                        mainAxisSpacing: 10,
                        children: [
                          StatCard(title: 'Soil', value: '${farm.soilMoisture.toStringAsFixed(0)}%', icon: Icons.water_drop_outlined, color: AppTheme.sky),
                          StatCard(title: 'Temp', value: '${farm.temperature.toStringAsFixed(0)} C', icon: Icons.thermostat, color: AppTheme.danger),
                          StatCard(title: 'Rain', value: '${farm.rainProbability.toStringAsFixed(0)}%', icon: Icons.cloud_outlined, color: AppTheme.field),
                          StatCard(title: 'NDVI', value: farm.ndvi.toStringAsFixed(2), icon: Icons.grass, color: AppTheme.leaf),
                        ],
                      ),
                      const SizedBox(height: 16),
                      Row(
                        children: [
                          Expanded(
                            child: FilledButton.icon(
                              onPressed: _generate,
                              icon: const Icon(Icons.auto_awesome),
                              label: const Text('Generate alerts'),
                            ),
                          ),
                          const SizedBox(width: 10),
                          IconButton.filledTonal(
                            tooltip: 'Send top SMS',
                            onPressed: _sendTopSms,
                            icon: const Icon(Icons.sms_outlined),
                          ),
                        ],
                      ),
                      const SizedBox(height: 16),
                      Text('Recommendations', style: Theme.of(context).textTheme.titleLarge?.copyWith(fontWeight: FontWeight.w800)),
                      const SizedBox(height: 8),
                      if (_alerts.isEmpty)
                        const Card(child: Padding(padding: EdgeInsets.all(18), child: Text('No recommendations yet.')))
                      else
                        ..._alerts.map((rec) => Padding(
                              padding: const EdgeInsets.only(bottom: 10),
                              child: RecommendationCard(recommendation: rec, onSend: () => _sendSms(rec.message)),
                            )),
                    ],
                  ),
                ),
    );
  }

  Future<void> _generate() async {
    try {
      final generated = await _recommendations.generate(farmId: widget.farmId);
      setState(() => _alerts = generated);
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(const SnackBar(content: Text('Alerts generated')));
    } catch (error) {
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(error.toString())));
    }
  }

  Future<void> _sendTopSms() async {
    await _sendSms('');
  }

  Future<void> _sendSms(String message) async {
    try {
      await _sms.send(farmId: widget.farmId ?? 0, message: message);
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(const SnackBar(content: Text('SMS sent')));
    } catch (error) {
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(error.toString())));
    }
  }
}

class _FarmHeader extends StatelessWidget {
  const _FarmHeader({required this.farm});

  final Farm farm;

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(farm.crop, style: Theme.of(context).textTheme.headlineSmall?.copyWith(fontWeight: FontWeight.w800)),
            const SizedBox(height: 6),
            Text('${farm.farmer} | ${farm.phone}', style: const TextStyle(color: AppTheme.muted)),
            const SizedBox(height: 8),
            Wrap(
              spacing: 8,
              runSpacing: 8,
              children: [
                Chip(label: Text(farm.status)),
                Chip(label: Text(farm.preferredLanguage)),
                if (farm.county.isNotEmpty) Chip(label: Text(farm.county)),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
