import 'package:flutter/material.dart';

import '../core/theme.dart';
import '../models/farm.dart';

class FarmCard extends StatelessWidget {
  const FarmCard({super.key, required this.farm, required this.onTap});

  final Farm farm;
  final VoidCallback onTap;

  @override
  Widget build(BuildContext context) {
    final statusColor = farm.status.toLowerCase() == 'critical' || farm.status.toLowerCase() == 'alert'
        ? AppTheme.danger
        : farm.status.toLowerCase() == 'warning'
            ? AppTheme.warning
            : AppTheme.leaf;
    return Card(
      child: InkWell(
        borderRadius: BorderRadius.circular(8),
        onTap: onTap,
        child: Padding(
          padding: const EdgeInsets.all(14),
          child: Row(
            children: [
              CircleAvatar(
                backgroundColor: statusColor.withOpacity(0.12),
                foregroundColor: statusColor,
                child: const Icon(Icons.agriculture_outlined),
              ),
              const SizedBox(width: 12),
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(farm.name, style: Theme.of(context).textTheme.titleMedium?.copyWith(fontWeight: FontWeight.w700)),
                    const SizedBox(height: 2),
                    Text('${farm.crop} by ${farm.farmer}', maxLines: 1, overflow: TextOverflow.ellipsis),
                    const SizedBox(height: 2),
                    Text(farm.phone, style: Theme.of(context).textTheme.bodySmall?.copyWith(color: AppTheme.muted)),
                  ],
                ),
              ),
              const Icon(Icons.chevron_right),
            ],
          ),
        ),
      ),
    );
  }
}
