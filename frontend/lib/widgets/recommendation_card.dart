import 'package:flutter/material.dart';
import '../models/recommendation.dart';

class RecommendationCard extends StatelessWidget {
  final Recommendation recommendation;

  const RecommendationCard({super.key, required this.recommendation});

  Color _severityColor(String severity) {
    switch (severity.toLowerCase()) {
      case 'low':
        return Colors.green;
      case 'medium':
        return Colors.orange;
      case 'high':
        return Colors.red;
      default:
        return Colors.grey;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Card(
      child: ListTile(
        leading: Icon(Icons.warning_amber_rounded,
            color: _severityColor(recommendation.severity)),
        title: Text(recommendation.type),
        subtitle: Text(recommendation.message),
        trailing: Text(recommendation.severity,
            style: TextStyle(color: _severityColor(recommendation.severity))),
      ),
    );
  }
}