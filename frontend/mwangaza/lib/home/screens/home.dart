import 'package:flutter/material.dart';
import '../widget/animated_sidebar.dart';
import '../views/dashboard_view.dart';
import '../views/map_view.dart';
import '../views/notification_view.dart';
import '../views/insight_view.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  int _selectedIndex = 0;

  final List<Map<String, dynamic>> _menuItems = [
    {'title': 'Dashboard', 'icon': Icons.dashboard},
    {'title': 'Map', 'icon': Icons.map},
    {'title': 'Notification', 'icon': Icons.notifications},
    {'title': 'Insight', 'icon': Icons.insights},
    {'title': 'Logout', 'icon': Icons.logout},
  ];

  final List<Widget> _views = [
    const DashboardView(),
    const MapView(),
    //const NotificationView(),
    const InsightView(),
  ];

  void _onMenuItemSelected(int index) {
    if (_menuItems[index]['title'] == 'Logout') {
      _showLogoutDialog();
    } else {
      setState(() {
        _selectedIndex = index;
      });
    }
  }

  void _showLogoutDialog() {
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: const Text('Logout'),
        content: const Text('Are you sure you want to logout?'),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('Cancel'),
          ),
          TextButton(
            onPressed: () {
              Navigator.pop(context);
              ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(
                  content: Text('Logged out successfully'),
                  backgroundColor: Colors.green,
                ),
              );
            },
            style: TextButton.styleFrom(foregroundColor: Colors.red),
            child: const Text('Logout'),
          ),
        ],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Row(
        children: [
          AnimatedSidebar(
            menuItems: _menuItems,
            selectedIndex: _selectedIndex,
            onItemSelected: _onMenuItemSelected,
          ),
          Expanded(child: _views[_selectedIndex]),
        ],
      ),
    );
  }
}
