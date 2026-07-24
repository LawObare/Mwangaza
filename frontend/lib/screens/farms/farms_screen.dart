import 'package:flutter/material.dart';

import '../../models/farm.dart';
import '../../routes/app_routes.dart';
import '../../services/api_service.dart';
import '../../services/farm_service.dart';
import '../../widgets/farm_card.dart';
import '../../widgets/loading_widget.dart';

class FarmsScreen extends StatefulWidget {
  const FarmsScreen({super.key});

  @override
  State<FarmsScreen> createState() => _FarmsScreenState();
}

class _FarmsScreenState extends State<FarmsScreen> {
  final _api = ApiService();
  late final _service = FarmService(_api);

  bool _loading = true;
  String _error = '';
  List<Farm> _farms = const [];

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
      final farms = await _service.listFarms();
      setState(() {
        _farms = farms;
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
        title: const Text('Farms'),
        actions: [
          IconButton(
            tooltip: 'Dashboard',
            onPressed: () => Navigator.pushReplacementNamed(context, AppRoutes.dashboard),
            icon: const Icon(Icons.dashboard_outlined),
          ),
          IconButton(
            tooltip: 'SMS',
            onPressed: () => Navigator.pushReplacementNamed(context, AppRoutes.sms),
            icon: const Icon(Icons.sms_outlined),
          ),
        ],
      ),
      floatingActionButton: FloatingActionButton.extended(
        onPressed: _showCreateSheet,
        icon: const Icon(Icons.add),
        label: const Text('Farm'),
      ),
      body: _loading
          ? const LoadingWidget(message: 'Loading farms')
          : _error.isNotEmpty
              ? ErrorState(message: _error, onRetry: _load)
              : RefreshIndicator(
                  onRefresh: _load,
                  child: ListView.separated(
                    padding: const EdgeInsets.all(16),
                    itemBuilder: (context, index) {
                      final farm = _farms[index];
                      return FarmCard(
                        farm: farm,
                        onTap: () => Navigator.pushNamed(context, AppRoutes.farmDetails, arguments: farm.id),
                      );
                    },
                    separatorBuilder: (_, __) => const SizedBox(height: 10),
                    itemCount: _farms.length,
                  ),
                ),
    );
  }

  Future<void> _showCreateSheet() async {
    final created = await showModalBottomSheet<bool>(
      context: context,
      isScrollControlled: true,
      builder: (_) => _CreateFarmSheet(service: _service),
    );
    if (created == true) _load();
  }
}

class _CreateFarmSheet extends StatefulWidget {
  const _CreateFarmSheet({required this.service});

  final FarmService service;

  @override
  State<_CreateFarmSheet> createState() => _CreateFarmSheetState();
}

class _CreateFarmSheetState extends State<_CreateFarmSheet> {
  final _formKey = GlobalKey<FormState>();
  final _name = TextEditingController();
  final _farmer = TextEditingController();
  final _phone = TextEditingController();
  final _county = TextEditingController();
  final _crop = TextEditingController(text: 'Maize');
  final _lat = TextEditingController();
  final _lng = TextEditingController();
  String _language = 'English';
  bool _saving = false;

  @override
  void dispose() {
    _name.dispose();
    _farmer.dispose();
    _phone.dispose();
    _county.dispose();
    _crop.dispose();
    _lat.dispose();
    _lng.dispose();
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
      child: Form(
        key: _formKey,
        child: SingleChildScrollView(
          child: Column(
            mainAxisSize: MainAxisSize.min,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text('Register Farm', style: Theme.of(context).textTheme.titleLarge?.copyWith(fontWeight: FontWeight.w800)),
              const SizedBox(height: 12),
              _field(_name, 'Farm name'),
              _field(_farmer, 'Farmer name'),
              _field(_phone, 'Phone'),
              _field(_county, 'County', required: false),
              _field(_crop, 'Crop'),
              Row(
                children: [
                  Expanded(child: _field(_lat, 'Latitude', required: false, number: true)),
                  const SizedBox(width: 10),
                  Expanded(child: _field(_lng, 'Longitude', required: false, number: true)),
                ],
              ),
              SegmentedButton<String>(
                segments: const [
                  ButtonSegment(value: 'English', label: Text('English')),
                  ButtonSegment(value: 'Swahili', label: Text('Swahili')),
                ],
                selected: {_language},
                onSelectionChanged: (value) => setState(() => _language = value.first),
              ),
              const SizedBox(height: 14),
              SizedBox(
                width: double.infinity,
                child: FilledButton.icon(
                  onPressed: _saving ? null : _save,
                  icon: _saving ? const SizedBox.square(dimension: 18, child: CircularProgressIndicator(strokeWidth: 2)) : const Icon(Icons.save_outlined),
                  label: const Text('Create farm'),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget _field(TextEditingController controller, String label, {bool required = true, bool number = false}) {
    return Padding(
      padding: const EdgeInsets.only(bottom: 10),
      child: TextFormField(
        controller: controller,
        keyboardType: number ? TextInputType.number : TextInputType.text,
        decoration: InputDecoration(labelText: label),
        validator: required ? (value) => (value == null || value.trim().isEmpty) ? 'Required' : null : null,
      ),
    );
  }

  Future<void> _save() async {
    if (!_formKey.currentState!.validate()) return;
    setState(() => _saving = true);
    try {
      await widget.service.createFarm(Farm(
        id: 0,
        name: _name.text.trim(),
        farmer: _farmer.text.trim(),
        phone: _phone.text.trim(),
        county: _county.text.trim(),
        crop: _crop.text.trim(),
        preferredLanguage: _language,
        status: 'new',
        latitude: double.tryParse(_lat.text) ?? 0,
        longitude: double.tryParse(_lng.text) ?? 0,
      ));
      if (mounted) Navigator.pop(context, true);
    } catch (error) {
      setState(() => _saving = false);
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(error.toString())));
    }
  }
}
