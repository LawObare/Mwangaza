import 'package:flutter/material.dart';

class InsightView extends StatelessWidget {
  const InsightView({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.white,
        elevation: 1,
        title: const Text(
          'Insights',
          style: TextStyle(color: Colors.black87, fontWeight: FontWeight.w600),
        ),
        actions: [
          IconButton(
            icon: const Icon(Icons.download),
            onPressed: () {
              ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(
                  content: Text('Downloading insights...'),
                  duration: Duration(seconds: 1),
                ),
              );
            },
            color: Colors.grey[600],
          ),
        ],
      ),
      body: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Overview Cards
            Row(
              children: [
                _buildInsightCard(
                  'Total Sessions',
                  '12,847',
                  '+12.5%',
                  Colors.blue,
                  Icons.people,
                ),
                const SizedBox(width: 12),
                _buildInsightCard(
                  'Conversion Rate',
                  '3.2%',
                  '+0.8%',
                  Colors.green,
                  Icons.trending_up,
                ),
              ],
            ),
            const SizedBox(height: 12),
            Row(
              children: [
                _buildInsightCard(
                  'Avg. Duration',
                  '4m 32s',
                  '+15s',
                  Colors.orange,
                  Icons.timer,
                ),
                const SizedBox(width: 12),
                _buildInsightCard(
                  'Bounce Rate',
                  '24.8%',
                  '-2.1%',
                  Colors.purple,
                  Icons.exit_to_app,
                ),
              ],
            ),
            const SizedBox(height: 24),

            // Recent Insights
            const Text(
              'Recent Insights',
              style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 16),
            Expanded(
              child: ListView.separated(
                itemCount: 8,
                separatorBuilder: (context, index) => const Divider(),
                itemBuilder: (context, index) {
                  final insights = [
                    'User engagement increased by 15% this week',
                    'New feature adoption rate is 67%',
                    'Average session duration is up by 12%',
                    'Mobile app usage grew by 23%',
                    'Customer satisfaction score: 4.8/5',
                    'Daily active users reached 8,500',
                    'Revenue grew by 18% month-over-month',
                    'Support tickets decreased by 22%',
                  ];
                  return ListTile(
                    leading: CircleAvatar(
                      backgroundColor: Colors
                          .primaries[index % Colors.primaries.length]
                          .withOpacity(0.1),
                      child: Icon(
                        Icons.lightbulb,
                        color:
                            Colors.primaries[index % Colors.primaries.length],
                        size: 20,
                      ),
                    ),
                    title: Text(
                      insights[index],
                      style: const TextStyle(fontSize: 14),
                    ),
                    trailing: const Icon(Icons.arrow_forward_ios, size: 16),
                    onTap: () {
                      ScaffoldMessenger.of(context).showSnackBar(
                        SnackBar(
                          content: Text('Viewing insight ${index + 1}'),
                          duration: const Duration(seconds: 1),
                        ),
                      );
                    },
                  );
                },
              ),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildInsightCard(
    String title,
    String value,
    String change,
    Color color,
    IconData icon,
  ) {
    return Expanded(
      child: Card(
        child: Padding(
          padding: const EdgeInsets.all(12),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Row(
                children: [
                  Container(
                    padding: const EdgeInsets.all(6),
                    decoration: BoxDecoration(
                      color: color.withOpacity(0.1),
                      borderRadius: BorderRadius.circular(6),
                    ),
                    child: Icon(icon, color: color, size: 16),
                  ),
                  const Spacer(),
                  Container(
                    padding: const EdgeInsets.symmetric(
                      horizontal: 6,
                      vertical: 2,
                    ),
                    decoration: BoxDecoration(
                      color: change.startsWith('+')
                          ? Colors.green.withOpacity(0.1)
                          : Colors.red.withOpacity(0.1),
                      borderRadius: BorderRadius.circular(4),
                    ),
                    child: Text(
                      change,
                      style: TextStyle(
                        fontSize: 10,
                        color: change.startsWith('+')
                            ? Colors.green
                            : Colors.red,
                        fontWeight: FontWeight.w600,
                      ),
                    ),
                  ),
                ],
              ),
              const SizedBox(height: 8),
              Text(
                value,
                style: const TextStyle(
                  fontSize: 18,
                  fontWeight: FontWeight.bold,
                ),
              ),
              Text(
                title,
                style: TextStyle(fontSize: 12, color: Colors.grey[600]),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
