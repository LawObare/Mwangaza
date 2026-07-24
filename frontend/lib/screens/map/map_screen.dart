import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:latlong2/latlong.dart';
import '../../services/api_service.dart';
import '../../services/farm_service.dart';
import '../../models/farm.dart';
import '../../widgets/loading_widget.dart';

class MapScreen extends StatefulWidget {
  const MapScreen({super.key});

  @override
  State<MapScreen> createState() => _MapScreenState();
}

class _MapScreenState extends State<MapScreen> {
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

  Color _markerColor(String status) {
    switch (status.toLowerCase()) {
      case 'healthy':
        return Colors.green;
      case 'needs attention':
        return Colors.orange;
      case 'urgent':
        return Colors.red;
      default:
        return Colors.grey;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Farm Map')),
      body: _loading
          ? const LoadingWidget(message: 'Loading map...')
          : FlutterMap(
              options: const MapOptions(
                initialCenter: LatLng(-1.286389, 36.817223),
                initialZoom: 7,
              ),
              children: [
                TileLayer(
                  urlTemplate: 'https://tile.openstreetmap.org/{z}/{x}/{y}.png',
                  userAgentPackageName: 'com.mwangaza.app',
                ),
                if (_farms != null)
                  MarkerLayer(
                    markers: _farms!.map((farm) {
                      return Marker(
                        point: LatLng(farm.latitude, farm.longitude),
                        width: 40,
                        height: 40,
                        child: GestureDetector(
                          onTap: () => _showFarmDetails(context, farm),
                          child: Icon(Icons.location_on,
                              color: _markerColor(farm.status), size: 40),
                        ),
                      );
                    }).toList(),
                  ),
              ],
            ),
    );
  }

  void _showFarmDetails(BuildContext context, Farm farm) {
    showModalBottomSheet(
      context: context,
      builder: (_) => Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(farm.name,
                style:
                    const TextStyle(fontSize: 18, fontWeight: FontWeight.bold)),
            const SizedBox(height: 8),
            Text('Farmer: ${farm.farmer}'),
            Text('Status: ${farm.status}'),
            if (farm.soilMoisture != null)
              Text('Soil Moisture: ${farm.soilMoisture}%'),
            const SizedBox(height: 16),
            ElevatedButton(
              onPressed: () {
                Navigator.pop(context);
                Navigator.pushNamed(context, '/farms/details',
                    arguments: farm.id);
              },
              child: const Text('View Details'),
            ),
          ],
        ),
      ),
    );
  }
}