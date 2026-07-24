import 'package:flutter/material.dart';

import '../../core/theme.dart';
import '../../models/farm.dart';
import '../../models/sms.dart';
import '../../routes/app_routes.dart';
import '../../services/api_service.dart';
import '../../services/farm_service.dart';
import '../../services/sms_service.dart';
import '../../widgets/loading_widget.dart';

class SMSScreen extends StatefulWidget {
  const SMSScreen({super.key});

  @override
  State<SMSScreen> createState() => _SMSScreenState();
}

class _SMSScreenState extends State<SMSScreen> {
  final _api = ApiService();
  late final _sms = SmsService(_api);
  late final _farms = FarmService(_api);

  bool _loading = true;
  String _error = '';
  List<SmsMessage> _logs = const [];
  List<Farm> _farmList = const [];

  @override
  void initState() {
    super.initState();
    _load();
  }

  @override
  void dispose() {
    _api.dispose();
    super.dispose();
  }

  Future<void> _load() async {
    setState(() {
      _loading = true;
      _error = '';
    });
    try {
      final results = await Future.wait([_sms.listSms(), _farms.listFarms()]);
      setState(() {
        _logs = results[0] as List<SmsMessage>;
        _farmList = results[1] as List<Farm>;
        _loading = false;
      });
    } catch (error) {
      setState(() {
        _error = error.toString();
        _loading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('SMS History'),
        actions: [
          IconButton(
            tooltip: 'Dashboard',
            onPressed: () => Navigator.pushReplacementNamed(context, AppRoutes.dashboard),
            icon: const Icon(Icons.dashboard_outlined),
          ),
        ],
      ),
      floatingActionButton: FloatingActionButton.extended(
        onPressed: _farmList.isEmpty ? null : _showSendSheet,
        icon: const Icon(Icons.send_outlined),
        label: const Text('Send'),
      ),
      body: _loading
          ? const LoadingWidget(message: 'Loading SMS logs')
          : _error.isNotEmpty
              ? ErrorState(message: _error, onRetry: _load)
              : RefreshIndicator(
                  onRefresh: _load,
                  child: ListView.separated(
                    padding: const EdgeInsets.all(16),
                    itemCount: _logs.isEmpty ? 1 : _logs.length,
                    separatorBuilder: (_, __) => const SizedBox(height: 10),
                    itemBuilder: (context, index) {
                      if (_logs.isEmpty) {
                        return const Card(child: Padding(padding: EdgeInsets.all(18), child: Text('No SMS messages yet.')));
                      }
                      final log = _logs[index];
                      return Card(
                        child: ListTile(
                          leading: const Icon(Icons.sms_outlined, color: AppTheme.sky),
                          title: Text(log.message, maxLines: 2, overflow: TextOverflow.ellipsis),
                          subtitle: Text('${log.phoneNumber} | ${log.provider} | ${log.sentAt}'),
                          trailing: Text(log.status),
                        ),
                      );
                    },
                  ),
                ),
    );
  }

  Future<void> _showSendSheet() async {
    final sent = await showModalBottomSheet<bool>(
      context: context,
      isScrollControlled: true,
      builder: (_) => _SendSmsSheet(farms: _farmList, service: _sms),
    );
    if (sent == true) _load();
  }
}

class _SendSmsSheet extends StatefulWidget {
  const _SendSmsSheet({required this.farms, required this.service});

  final List<Farm> farms;
  final SmsService service;

  @override
  State<_SendSmsSheet> createState() => _SendSmsSheetState();
}

class _SendSmsSheetState extends State<_SendSmsSheet> {
  late int _farmId;
  final _message = TextEditingController();
  bool _sending = false;

  @override
  void initState() {
    super.initState();
    _farmId = widget.farms.first.id;
  }

  @override
  void dispose() {
    _message.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.only(
        left: 16,
        right: 16,
        top: 16,
        bottom: MediaQuery.viewInsetsOf(context).bottom + 16,
      ),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text('Send SMS', style: Theme.of(context).textTheme.titleLarge?.copyWith(fontWeight: FontWeight.w800)),
          const SizedBox(height: 12),
          DropdownButtonFormField<int>(
            value: _farmId,
            decoration: const InputDecoration(labelText: 'Farm'),
            items: widget.farms.map((farm) => DropdownMenuItem(value: farm.id, child: Text('${farm.name} (${farm.crop})'))).toList(),
            onChanged: (value) => setState(() => _farmId = value ?? _farmId),
          ),
          const SizedBox(height: 10),
          TextField(
            controller: _message,
            minLines: 3,
            maxLines: 5,
            decoration: const InputDecoration(
              labelText: 'Message',
              hintText: 'Leave blank to send the latest highest-priority recommendation',
            ),
          ),
          const SizedBox(height: 14),
          SizedBox(
            width: double.infinity,
            child: FilledButton.icon(
              onPressed: _sending ? null : _send,
              icon: const Icon(Icons.send_outlined),
              label: const Text('Send SMS'),
            ),
          ),
        ],
      ),
    );
  }

  Future<void> _send() async {
    setState(() => _sending = true);
    try {
      await widget.service.send(farmId: _farmId, message: _message.text.trim());
      if (mounted) Navigator.pop(context, true);
    } catch (error) {
      setState(() => _sending = false);
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(error.toString())));
    }
  }
}
