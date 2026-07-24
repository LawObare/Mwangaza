import 'package:flutter/material.dart';

import '../core/theme.dart';
import '../models/satellite.dart';

class SatelliteCard extends StatelessWidget {
  const SatelliteCard({super.key, required this.data});

  final SatelliteData data;

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              children: [
                const Icon(Icons.satellite_alt_outlined, color: AppTheme.sky),
                const SizedBox(width: 8),
                Text('Latest SpaceIoTBox Snapshot', style: Theme.of(context).textTheme.titleMedium?.copyWith(fontWeight: FontWeight.w700)),
              ],
            ),
            const SizedBox(height: 14),
            _Line(label: 'Soil moisture', value: '${data.soilMoisture.toStringAsFixed(1)}%'),
            _Line(label: 'Temperature', value: '${data.temperature.toStringAsFixed(1)} C'),
            _Line(label: 'Rain probability', value: '${data.rainProbability.toStringAsFixed(1)}%'),
            _Line(label: 'NDVI', value: data.ndvi.toStringAsFixed(2)),
            _Line(label: 'Wind speed', value: '${data.windSpeed.toStringAsFixed(1)} km/h'),
            if (data.source.isNotEmpty) _Line(label: 'Source', value: data.source),
          ],
        ),
      ),
    );
  }
}

class _Line extends StatelessWidget {
  const _Line({required this.label, required this.value});

  final String label;
  final String value;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 5),
      child: Row(
        children: [
          Expanded(child: Text(label, style: const TextStyle(color: AppTheme.muted))),
          Text(value, style: const TextStyle(fontWeight: FontWeight.w700)),
        ],
      ),
    );
  }
}
