import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:latlong2/latlong.dart';
import 'package:url_launcher/url_launcher.dart';

class MapView extends StatefulWidget {
  const MapView({super.key});

  @override
  State<MapView> createState() => _MapViewState();
}

class _MapViewState extends State<MapView> {
  final MapController _mapController = MapController();

  final LatLng _currentPosition = const LatLng(-1.286389, 36.817223);

  final List<LatLng> _markers = [
    const LatLng(-1.286389, 36.817223),
    const LatLng(-1.2921, 36.8219),
    const LatLng(-1.2803, 36.8134),
  ];

  double _zoom = 13;

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
            icon: const Icon(Icons.add_location),
            onPressed: _showAddMarkerDialog,
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
            ),
            children: [
              TileLayer(
                urlTemplate: 'https://tile.openstreetmap.org/{z}/{x}/{y}.png',
                userAgentPackageName: 'com.example.app',
              ),

              MarkerLayer(
                markers: _markers
                    .map(
                      (point) => Marker(
                        point: point,
                        width: 40,
                        height: 40,
                        child: const Icon(
                          Icons.location_pin,
                          color: Colors.red,
                          size: 40,
                        ),
                      ),
                    )
                    .toList(),
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
                    const SizedBox(height: 8),
                    Text('Markers: ${_markers.length}'),
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

  void _showAddMarkerDialog() {
    showDialog(
      context: context,
      builder: (_) => AlertDialog(
        title: const Text("Add Marker"),
        content: const Text("Add a marker at the current location?"),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context),
            child: const Text("Cancel"),
          ),
          ElevatedButton(
            onPressed: () {
              setState(() {
                _markers.add(_currentPosition);
              });
              Navigator.pop(context);
            },
            child: const Text("Add Marker"),
          ),
        ],
      ),
    );
  }
}
