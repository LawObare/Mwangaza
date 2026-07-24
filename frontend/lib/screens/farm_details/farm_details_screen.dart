import 'package:flutter/material.dart';
import '../../widgets/stat_card.dart';
import '../../widgets/recommendation_card.dart';
import '../../widgets/loading_widget.dart';
import '../../services/api_service.dart';
import '../../services/farm_service.dart';
import '../../services/recommendation_service.dart';
import '../../models/farm.dart';
import '../../models/recommendation.dart';

class FarmDetailsScreen extends StatefulWidget {
  final int farmId;

  const FarmDetailsScreen({super.key, required this.farmId});

  @override
  State<FarmDetailsScreen> createState() => _FarmDetailsScreenState();
}

class _FarmDetailsScreenState extends State<FarmDetailsScreen> {
  final ApiService _api = ApiService();
  Farm? _farm;
  List<Recommendation>? _recommendations;
  bool _loading = true;

  @override
  void initState() {
    super.initState();
    _loadDetails();
  }

  Future<void> _loadDetails() async {
    final farmService = FarmService(_api);
    final recService = RecommendationService(_api);

    try {
      final results = await Future.wait([
        farmService.getFarm(widget.farmId),
        recService.getRecommendationsForFarm(widget.farmId),
      ]);
      setState(() {
        _farm = results[0] as Farm;
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
      appBar: AppBar(title: Text(_farm?.name ?? 'Farm Details')),
      body: _loading
          ? const LoadingWidget()
          : RefreshIndicator(
              onRefresh: _loadDetails,
              child: SingleChildScrollView(
                physics: const AlwaysScrollableScrollPhysics(),
                padding: const EdgeInsets.all(16),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    if (_farm != null) ...[
                      Card(
                        child: Padding(
                          padding: const EdgeInsets.all(16),
                          child: Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              Text(_farm!.name,
                                  style: const TextStyle(
                                      fontSize: 20,
                                      fontWeight: FontWeight.bold)),
                              const SizedBox(height: 8),
                              Text('Farmer: ${_farm!.farmer}'),
                              Text('Status: ${_farm!.status}'),
                              Text(
                                  'Location: ${_farm!.latitude}, ${_farm!.longitude}'),
                            ],
                          ),
                        ),
                      ),
                      if (_farm!.soilMoisture != null) ...[
                        const SizedBox(height: 16),
                        GridView.count(
                          crossAxisCount: 2,
                          shrinkWrap: true,
                          physics: const NeverScrollableScrollPhysics(),
                          childAspectRatio: 1.5,
                          children: [
                            StatCard(
                                title: 'Soil Moisture',
                                value: '${_farm!.soilMoisture}%',
                                icon: Icons.water_drop),
                            StatCard(
                                title: 'Temperature',
                                value: '${_farm!.temperature}°C',
                                icon: Icons.thermostat),
                          ],
                        ),
                      ],
                    ],
                    if (_recommendations != null) ...[
                      const SizedBox(height: 16),
                      const Text('Recommendations',
                          style: TextStyle(
                              fontSize: 16, fontWeight: FontWeight.bold)),
                      const SizedBox(height: 8),
                      ..._recommendations!
                          .map((r) => RecommendationCard(recommendation: r)),
                    ],
                  ],
                ),
              ),
            ),
    );
  }
}