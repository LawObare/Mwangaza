import 'package:flutter/material.dart';

import '../core/theme.dart';
import '../models/recommendation.dart';

class RecommendationCard extends StatelessWidget {
  const RecommendationCard({super.key, required this.recommendation, this.onSend});

  final Recommendation recommendation;
  final VoidCallback? onSend;

  @override
  Widget build(BuildContext context) {
    final color = _priorityColor(recommendation.priority);
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(14),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              children: [
                Icon(_typeIcon(recommendation.type), color: color),
                const SizedBox(width: 8),
                Expanded(
                  child: Text(
                    _title(recommendation.type),
                    style: Theme.of(context).textTheme.titleMedium?.copyWith(fontWeight: FontWeight.w700),
                  ),
                ),
                _Chip(label: recommendation.priority, color: color),
              ],
            ),
            const SizedBox(height: 10),
            Text(recommendation.message),
            if (recommendation.reason.isNotEmpty) ...[
              const SizedBox(height: 8),
              Text(
                recommendation.reason,
                style: Theme.of(context).textTheme.bodySmall?.copyWith(color: AppTheme.muted),
              ),
            ],
            const SizedBox(height: 12),
            Row(
              children: [
                if (recommendation.confidence.isNotEmpty)
                  _Chip(label: 'Confidence ${recommendation.confidence}', color: AppTheme.sky),
                const Spacer(),
                if (onSend != null)
                  IconButton.filledTonal(
                    tooltip: 'Send SMS',
                    onPressed: onSend,
                    icon: const Icon(Icons.sms_outlined),
                  ),
              ],
            ),
          ],
        ),
      ),
    );
  }

  Color _priorityColor(String priority) {
    switch (priority.toUpperCase()) {
      case 'HIGH':
        return AppTheme.danger;
      case 'MEDIUM':
        return AppTheme.warning;
      default:
        return AppTheme.leaf;
    }
  }

  IconData _typeIcon(String type) {
    switch (type) {
      case 'rainfall':
        return Icons.water_drop_outlined;
      case 'irrigation':
        return Icons.water;
      case 'temperature':
      case 'heat_stress':
        return Icons.thermostat;
      case 'wind':
        return Icons.air;
      case 'ndvi':
      case 'drought_stress':
        return Icons.grass;
      default:
        return Icons.monitor_heart_outlined;
    }
  }

  String _title(String type) {
    return type.replaceAll('_', ' ').split(' ').map((part) {
      if (part.isEmpty) return part;
      return '${part[0].toUpperCase()}${part.substring(1)}';
    }).join(' ');
  }
}

class _Chip extends StatelessWidget {
  const _Chip({required this.label, required this.color});

  final String label;
  final Color color;

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 9, vertical: 5),
      decoration: BoxDecoration(
        color: color.withOpacity(0.12),
        borderRadius: BorderRadius.circular(99),
      ),
      child: Text(
        label,
        style: Theme.of(context).textTheme.labelSmall?.copyWith(color: color, fontWeight: FontWeight.w700),
      ),
    );
  }
}
