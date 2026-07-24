import 'package:flutter/material.dart';
import '../models/farm.dart';

class FarmCard extends StatelessWidget {
  final Farm farm;
  final VoidCallback onTap;

  const FarmCard({super.key, required this.farm, required this.onTap});

  Color _statusColor(String status) {
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
    return Card(
      child: ListTile(
        leading: CircleAvatar(
          backgroundColor: _statusColor(farm.status),
          child: Text('${farm.id}', style: const TextStyle(color: Colors.white)),
        ),
        title: Text(farm.name, style: const TextStyle(fontWeight: FontWeight.bold)),
        subtitle: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text('Farmer: ${farm.farmer}'),
            Text('Status: ${farm.status}'),
          ],
        ),
        trailing: const Icon(Icons.chevron_right),
        onTap: onTap,
      ),
    );
  }
}