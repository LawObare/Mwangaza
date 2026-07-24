import 'package:flutter/material.dart';

import '../../core/theme.dart';
import '../../models/farm.dart';
import '../../routes/app_routes.dart';
import '../../services/api_service.dart';
import '../../services/farm_service.dart';
import '../../widgets/loading_widget.dart';

class MapScreen extends StatefulWidget {
  const MapScreen({super.key});

  @override
  State<MapScreen> createState() => _MapScreenState();
}

class _MapScreenState extends State<MapScreen> {
  final _api = ApiService();
  late final _farms = FarmService(_api);

  bool _loading = true;
  String _error = '';
  List<Farm> _farmList = const [];

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
      final farms = await _farms.listFarms();
      setState(() {
        _farmList = farms;
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
      appBar: AppBar(
        title: const Text('Farm Map'),
        actions: [
          IconButton(
            tooltip: 'Dashboard',
            onPressed: () => Navigator.pushReplacementNamed(context, AppRoutes.dashboard),
            icon: const Icon(Icons.dashboard_outlined),
          ),
        ],
      ),
      body: _loading
          ? const LoadingWidget(message: 'Loading map')
          : _error.isNotEmpty
              ? ErrorState(message: _error, onRetry: _load)
              : RefreshIndicator(
                  onRefresh: _load,
                  child: ListView(
                    padding: const EdgeInsets.all(16),
                    children: [
                      AspectRatio(
                        aspectRatio: 1.45,
                        child: Card(
                          child: CustomPaint(
                            painter: _FarmMapPainter(_farmList),
                            child: const SizedBox.expand(),
                          ),
                        ),
                      ),
                      const SizedBox(height: 16),
                      Text('Farm Coordinates', style: Theme.of(context).textTheme.titleLarge?.copyWith(fontWeight: FontWeight.w800)),
                      const SizedBox(height: 8),
                      ..._farmList.map((farm) => Card(
                            child: ListTile(
                              leading: const Icon(Icons.place_outlined, color: AppTheme.leaf),
                              title: Text(farm.name),
                              subtitle: Text('${farm.crop} | ${farm.latitude.toStringAsFixed(4)}, ${farm.longitude.toStringAsFixed(4)}'),
                              onTap: () => Navigator.pushNamed(context, AppRoutes.farmDetails, arguments: farm.id),
                            ),
                          )),
                    ],
                  ),
                ),
    );
  }
}

class _FarmMapPainter extends CustomPainter {
  _FarmMapPainter(this.farms);

  final List<Farm> farms;

  @override
  void paint(Canvas canvas, Size size) {
    final background = Paint()..color = const Color(0xFFE9F2E5);
    final grid = Paint()
      ..color = const Color(0xFFD2DEC9)
      ..strokeWidth = 1;
    final marker = Paint()..color = AppTheme.leaf;
    final warning = Paint()..color = AppTheme.warning;

    canvas.drawRect(Offset.zero & size, background);
    for (var i = 1; i < 5; i++) {
      final x = size.width * i / 5;
      final y = size.height * i / 5;
      canvas.drawLine(Offset(x, 0), Offset(x, size.height), grid);
      canvas.drawLine(Offset(0, y), Offset(size.width, y), grid);
    }

    final usable = farms.where((farm) => farm.latitude != 0 || farm.longitude != 0).toList();
    if (usable.isEmpty) {
      final textPainter = TextPainter(
        text: const TextSpan(text: 'Farm locations appear here when coordinates are available', style: TextStyle(color: AppTheme.muted)),
        textDirection: TextDirection.ltr,
        textAlign: TextAlign.center,
      )..layout(maxWidth: size.width - 40);
      textPainter.paint(canvas, Offset((size.width - textPainter.width) / 2, (size.height - textPainter.height) / 2));
      return;
    }

    final minLat = usable.map((farm) => farm.latitude).reduce((a, b) => a < b ? a : b);
    final maxLat = usable.map((farm) => farm.latitude).reduce((a, b) => a > b ? a : b);
    final minLng = usable.map((farm) => farm.longitude).reduce((a, b) => a < b ? a : b);
    final maxLng = usable.map((farm) => farm.longitude).reduce((a, b) => a > b ? a : b);
    for (final farm in usable) {
      final x = _scale(farm.longitude, minLng, maxLng, 24, size.width - 24);
      final y = _scale(farm.latitude, minLat, maxLat, size.height - 24, 24);
      canvas.drawCircle(Offset(x, y), 8, farm.status.toLowerCase() == 'safe' ? marker : warning);
    }
  }

  double _scale(double value, double min, double max, double outMin, double outMax) {
    if ((max - min).abs() < 0.00001) return (outMin + outMax) / 2;
    return outMin + ((value - min) / (max - min)) * (outMax - outMin);
  }

  @override
  bool shouldRepaint(covariant _FarmMapPainter oldDelegate) => oldDelegate.farms != farms;
}
