import 'package:flutter/material.dart';
import 'package:mwangaza/auth/domain/repository/auth_repository.dart';
import 'package:mwangaza/auth/screens/login.dart';
import '../widget/animated_sidebar.dart';
import '../views/dashboard_view.dart';
import '../views/map_view.dart';
import '../views/insight_view.dart';
import '../views/sms_logs_view.dart';
import '../../farms/views/add_farm_view.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  int _selectedIndex = 0;
  bool _isLoggingOut = false;

  final AuthRemoteRepository _authRepo = AuthRemoteRepository();

  final List<Map<String, dynamic>> _menuItems = [
    {'title': 'Dashboard', 'icon': Icons.dashboard},
    {'title': 'Add Farm', 'icon': Icons.add_location_alt_outlined},
    {'title': 'Map', 'icon': Icons.map},
    {'title': 'Insight', 'icon': Icons.insights},
    {'title': 'SMS Logs', 'icon': Icons.sms_outlined},
    {'title': 'Logout', 'icon': Icons.logout},
  ];

  final List<Widget> _views = [
    const DashboardView(),
    const AddFarmView(),
    const MapView(),
    const InsightView(),
    const SmsLogsView(),
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
            onPressed: _handleLogout,
            style: TextButton.styleFrom(foregroundColor: Colors.red),
            child: _isLoggingOut
                ? const SizedBox(
                    height: 20,
                    width: 20,
                    child: CircularProgressIndicator(
                      strokeWidth: 2,
                      color: Colors.red,
                    ),
                  )
                : const Text('Logout'),
          ),
        ],
      ),
    );
  }

  Future<void> _handleLogout() async {
    setState(() {
      _isLoggingOut = true;
    });

    try {
      await _authRepo.logout();

      if (mounted) {
        Navigator.pop(context); // Close the dialog
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text('Logged out successfully'),
            backgroundColor: Colors.green,
          ),
        );

        // Navigate to login page
        Navigator.pushReplacement(
          context,
          MaterialPageRoute(builder: (context) => const LoginPage()),
        );
      }
    } catch (e) {
      if (mounted) {
        Navigator.pop(context); // Close the dialog
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text('Logout failed: ${e.toString()}'),
            backgroundColor: Colors.red,
          ),
        );
      }
    } finally {
      if (mounted) {
        setState(() {
          _isLoggingOut = false;
        });
      }
    }
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
