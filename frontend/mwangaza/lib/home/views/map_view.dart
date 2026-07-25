import 'dart:math';
import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:latlong2/latlong.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:geolocator/geolocator.dart';

class MapView extends StatefulWidget {
  const MapView({super.key});

  @override
  State<MapView> createState() => _MapViewState();
}

class _MapViewState extends State<MapView> {
  final MapController _mapController = MapController();

  // Default view: Nyanza region, Kenya (near Kisumu).
  LatLng _currentPosition = const LatLng(-0.0917, 34.7680);

  double _zoom = 9;

  // --- Field drawing state ---
  bool _isDrawing = false;
  final List<LatLng> _fieldPoints = [];
  double _fieldAreaSqMeters = 0;
  LatLng? _fieldCenter;

  bool _isLocating = false;
  String? _locationError;

  @override
  void initState() {
    super.initState();
    _getCurrentLocation();
  }

  Future<void> _getCurrentLocation() async {
    setState(() {
      _isLocating = true;
      _locationError = null;
    });

    try {
      bool serviceEnabled = await Geolocator.isLocationServiceEnabled();
      if (!serviceEnabled) {
        throw 'Location services are turned off on this device/browser.';
      }

      LocationPermission permission = await Geolocator.checkPermission();
      if (permission == LocationPermission.denied) {
        permission = await Geolocator.requestPermission();
        if (permission == LocationPermission.denied) {
          throw 'Location permission was denied.';
        }
      }

      if (permission == LocationPermission.deniedForever) {
        throw 'Location permission is permanently denied. '
            'Reset it in your browser/site settings.';
      }

      // Timeout so a stalled request (common with localhost/dev servers)
      // doesn't hang forever with no feedback.
      final position =
          await Geolocator.getCurrentPosition(
            desiredAccuracy: LocationAccuracy.high,
          ).timeout(
            const Duration(seconds: 10),
            onTimeout: () => throw 'Timed out waiting for your location.',
          );

      final newPosition = LatLng(position.latitude, position.longitude);

      if (!mounted) return;
      setState(() {
        _currentPosition = newPosition;
        _isLocating = false;
      });

      _mapController.move(newPosition, _zoom);
    } catch (e) {
      if (!mounted) return;
      setState(() {
        _isLocating = false;
        _locationError = e.toString();
      });
    }
  }

  // Spherical (WGS84-approx) polygon area in square meters.
  double _calculatePolygonArea(List<LatLng> points) {
    if (points.length < 3) return 0;
    const double earthRadius = 6378137; // meters
    double total = 0;
    for (int i = 0; i < points.length; i++) {
      final p1 = points[i];
      final p2 = points[(i + 1) % points.length];
      total +=
          _toRadians(p2.longitude - p1.longitude) *
          (2 + sin(_toRadians(p1.latitude)) + sin(_toRadians(p2.latitude)));
    }
    return (total * earthRadius * earthRadius / 2).abs();
  }

  double _toRadians(double degree) => degree * pi / 180;

  LatLng _calculateCenter(List<LatLng> points) {
    double lat = 0, lng = 0;
    for (final p in points) {
      lat += p.latitude;
      lng += p.longitude;
    }
    return LatLng(lat / points.length, lng / points.length);
  }

  void _recalculateField() {
    if (_fieldPoints.length >= 3) {
      _fieldAreaSqMeters = _calculatePolygonArea(_fieldPoints);
      _fieldCenter = _calculateCenter(_fieldPoints);
    } else {
      _fieldAreaSqMeters = 0;
      _fieldCenter = _fieldPoints.isNotEmpty
          ? _calculateCenter(_fieldPoints)
          : null;
    }
  }

  void _handleMapTap(LatLng point) {
    if (!_isDrawing) return;
    setState(() {
      _fieldPoints.add(point);
      _recalculateField();
    });
  }

  void _toggleDrawing() {
    setState(() {
      _isDrawing = !_isDrawing;
    });
  }

  void _undoLastPoint() {
    if (_fieldPoints.isEmpty) return;
    setState(() {
      _fieldPoints.removeLast();
      _recalculateField();
    });
  }

  void _clearField() {
    setState(() {
      _fieldPoints.clear();
      _fieldAreaSqMeters = 0;
      _fieldCenter = null;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Map View"),
        backgroundColor: Colors.white,
        foregroundColor: Colors.black,
        elevation: 1,
        actions: [
          IconButton(
            icon: const Icon(Icons.my_location),
            onPressed: _centerMap,
          ),
          IconButton(icon: const Icon(Icons.zoom_in), onPressed: _zoomIn),
          IconButton(icon: const Icon(Icons.zoom_out), onPressed: _zoomOut),
          IconButton(
            icon: Icon(
              Icons.edit_location_alt,
              color: _isDrawing ? Colors.green : null,
            ),
            tooltip: _isDrawing ? "Stop drawing field" : "Draw field",
            onPressed: _toggleDrawing,
          ),
        ],
      ),
      body: Stack(
        children: [
          FlutterMap(
            mapController: _mapController,
            options: MapOptions(
              initialCenter: _currentPosition,
              initialZoom: _zoom,
              onTap: (tapPosition, point) => _handleMapTap(point),
            ),
            children: [
              TileLayer(
                urlTemplate: 'https://tile.openstreetmap.org/{z}/{x}/{y}.png',
                userAgentPackageName: 'com.example.app',
              ),

              if (_fieldPoints.length >= 3)
                PolygonLayer(
                  polygons: [
                    Polygon(
                      points: _fieldPoints,
                      color: Colors.green.withOpacity(0.3),
                      borderColor: Colors.green,
                      borderStrokeWidth: 3,
                    ),
                  ],
                ),

              if (_fieldPoints.length >= 2)
                PolylineLayer(
                  polylines: [
                    Polyline(
                      points: _fieldPoints,
                      color: Colors.green,
                      strokeWidth: 3,
                    ),
                  ],
                ),

              MarkerLayer(
                markers: [
                  Marker(
                    point: _currentPosition,
                    width: 40,
                    height: 40,
                    child: const Icon(
                      Icons.location_pin,
                      color: Colors.red,
                      size: 40,
                    ),
                  ),
                  ..._fieldPoints.map(
                    (point) => Marker(
                      point: point,
                      width: 16,
                      height: 16,
                      child: Container(
                        decoration: BoxDecoration(
                          color: Colors.green,
                          shape: BoxShape.circle,
                          border: Border.all(color: Colors.white, width: 2),
                        ),
                      ),
                    ),
                  ),
                  if (_fieldCenter != null)
                    Marker(
                      point: _fieldCenter!,
                      width: 30,
                      height: 30,
                      child: const Icon(
                        Icons.center_focus_strong,
                        color: Colors.blue,
                        size: 26,
                      ),
                    ),
                ],
              ),

              RichAttributionWidget(
                attributions: [
                  TextSourceAttribution(
                    'OpenStreetMap contributors',
                    onTap: () {
                      launchUrl(
                        Uri.parse('https://openstreetmap.org/copyright'),
                      );
                    },
                  ),
                ],
              ),
            ],
          ),

          if (_locationError != null && !_isDrawing)
            Positioned(
              top: 12,
              left: 12,
              right: 12,
              child: Card(
                color: Colors.orange.shade50,
                elevation: 2,
                child: Padding(
                  padding: const EdgeInsets.symmetric(
                    horizontal: 12,
                    vertical: 8,
                  ),
                  child: Row(
                    children: [
                      const Icon(
                        Icons.warning_amber_rounded,
                        color: Colors.orange,
                        size: 20,
                      ),
                      const SizedBox(width: 8),
                      Expanded(
                        child: Text(
                          _locationError!,
                          style: const TextStyle(fontSize: 13),
                        ),
                      ),
                      TextButton(
                        onPressed: _isLocating ? null : _getCurrentLocation,
                        child: _isLocating
                            ? const SizedBox(
                                width: 16,
                                height: 16,
                                child: CircularProgressIndicator(
                                  strokeWidth: 2,
                                ),
                              )
                            : const Text("Retry"),
                      ),
                    ],
                  ),
                ),
              ),
            ),

          if (_isDrawing)
            Positioned(
              top: 12,
              left: 12,
              right: 12,
              child: Card(
                color: Colors.green.shade50,
                elevation: 2,
                child: Padding(
                  padding: const EdgeInsets.symmetric(
                    horizontal: 12,
                    vertical: 8,
                  ),
                  child: Row(
                    children: [
                      const Expanded(
                        child: Text(
                          "Tap the map to add field corners",
                          style: TextStyle(fontWeight: FontWeight.w500),
                        ),
                      ),
                      TextButton(
                        onPressed: _undoLastPoint,
                        child: const Text("Undo"),
                      ),
                      TextButton(
                        onPressed: _clearField,
                        child: const Text("Clear"),
                      ),
                    ],
                  ),
                ),
              ),
            ),

          Positioned(
            bottom: 20,
            left: 20,
            right: 20,
            child: Card(
              elevation: 4,
              child: Padding(
                padding: const EdgeInsets.all(16),
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    const Text(
                      'Location Information',
                      style: TextStyle(fontWeight: FontWeight.bold),
                    ),
                    const SizedBox(height: 8),
                    Text('Latitude: ${_currentPosition.latitude}'),
                    Text('Longitude: ${_currentPosition.longitude}'),
                    if (_fieldPoints.length >= 3) ...[
                      const Divider(height: 20),
                      const Text(
                        'Field Information',
                        style: TextStyle(fontWeight: FontWeight.bold),
                      ),
                      const SizedBox(height: 8),
                      Text(
                        'Area: ${_fieldAreaSqMeters.toStringAsFixed(1)} m² '
                        '(${(_fieldAreaSqMeters / 10000).toStringAsFixed(3)} ha, '
                        '${(_fieldAreaSqMeters * 0.000247105).toStringAsFixed(3)} acres)',
                      ),
                      if (_fieldCenter != null) ...[
                        Text(
                          'Center Latitude: ${_fieldCenter!.latitude.toStringAsFixed(6)}',
                        ),
                        Text(
                          'Center Longitude: ${_fieldCenter!.longitude.toStringAsFixed(6)}',
                        ),
                      ],
                    ],
                  ],
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }

  void _centerMap() {
    _mapController.move(_currentPosition, _zoom);
  }

  void _zoomIn() {
    setState(() {
      _zoom++;
    });
    _mapController.move(_currentPosition, _zoom);
  }

  void _zoomOut() {
    setState(() {
      _zoom--;
    });
    _mapController.move(_currentPosition, _zoom);
  }
}
