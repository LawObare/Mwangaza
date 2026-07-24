import 'package:flutter/material.dart';
import '../models/satellite.dart';

class SatelliteCard extends StatelessWidget {
  final SatelliteData data;

  const SatelliteCard({super.key, required this.data});

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            const Text('Satellite Data',
                style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold)),
            const Divider(),
            _row('Soil Moisture', '${data.soilMoisture}%'),
            _row('Temperature', '${data.temperature}°C'),
            _row('Rain Probability', '${data.rainProbability}%'),
            _row('NDVI', '${data.ndvi}'),
            _row('Wind Speed', '${data.windSpeed} m/s'),
          ],
        ),
      ),
    );
  }

  Widget _row(String label, String value) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 4),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [Text(label), Text(value, style: const TextStyle(fontWeight: FontWeight.w600))],
      ),
    );
  }
}