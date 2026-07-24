import 'package:flutter/material.dart';
import '../../widgets/loading_widget.dart';
import '../../services/api_service.dart';
import '../../services/sms_service.dart';
import '../../models/sms.dart';

class SmsScreen extends StatefulWidget {
  const SmsScreen({super.key});

  @override
  State<SmsScreen> createState() => _SmsScreenState();
}

class _SmsScreenState extends State<SmsScreen> {
  final ApiService _api = ApiService();
  List<SmsMessage>? _messages;
  bool _loading = true;

  @override
  void initState() {
    super.initState();
    _loadMessages();
  }

  Future<void> _loadMessages() async {
    final service = SmsService(_api);
    try {
      final messages = await service.getSmsHistory();
      setState(() {
        _messages = messages;
        _loading = false;
      });
    } catch (e) {
      setState(() => _loading = false);
    }
  }

  @override
  void dispose() {
    _api.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('SMS History')),
      body: _loading
          ? const LoadingWidget(message: 'Loading SMS history...')
          : RefreshIndicator(
              onRefresh: _loadMessages,
              child: ListView.builder(
                padding: const EdgeInsets.all(8),
                itemCount: _messages?.length ?? 0,
                itemBuilder: (context, index) {
                  final msg = _messages![index];
                  return Card(
                    child: ListTile(
                      leading: Icon(
                        msg.status == 'sent'
                            ? Icons.check_circle
                            : Icons.pending,
                        color: msg.status == 'sent' ? Colors.green : Colors.orange,
                      ),
                      title: Text(msg.message, maxLines: 2, overflow: TextOverflow.ellipsis),
                      subtitle: Text('${msg.phoneNumber} • ${msg.sentAt}'),
                      trailing: Text(msg.status),
                    ),
                  );
                },
              ),
            ),
    );
  }
}