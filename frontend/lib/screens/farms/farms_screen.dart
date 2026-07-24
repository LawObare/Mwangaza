import 'package:flutter/material.dart';
import '../../widgets/farm_card.dart';
import '../../widgets/loading_widget.dart';
import '../../services/api_service.dart';
import '../../services/farm_service.dart';
import '../../models/farm.dart';

class FarmsScreen extends StatefulWidget {
  const FarmsScreen({super.key});

  @override
  State<FarmsScreen> createState() => _FarmsScreenState();
}

class _FarmsScreenState extends State<FarmsScreen> {
  final ApiService _api = ApiService();
  List<Farm>? _farms;
  bool _loading = true;

  @override
  void initState() {
    super.initState();
    _loadFarms();
  }

  Future<void> _loadFarms() async {
    final service = FarmService(_api);
    try {
      final farms = await service.getFarms();
      setState(() {
        _farms = farms;
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
      appBar: AppBar(title: const Text('Farms')),
      body: _loading
          ? const LoadingWidget(message: 'Loading farms...')
          : RefreshIndicator(
              onRefresh: _loadFarms,
              child: ListView.builder(
                padding: const EdgeInsets.all(8),
                itemCount: _farms?.length ?? 0,
                itemBuilder: (context, index) {
                  final farm = _farms![index];
                  return FarmCard(
                    farm: farm,
                    onTap: () =>
                        Navigator.pushNamed(context, '/farms/details',
                            arguments: farm.id),
                  );
                },
              ),
            ),
    );
  }
}