import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:mwangaza/services/mwangaza_api_client.dart';

class AddFarmView extends StatefulWidget {
  const AddFarmView({super.key});

  @override
  State<AddFarmView> createState() => _AddFarmViewState();
}

class _AddFarmViewState extends State<AddFarmView> {
  final _formKey = GlobalKey<FormState>();
  final _farmName = TextEditingController();
  final _farmerName = TextEditingController();
  final _phone = TextEditingController();
  final _county = TextEditingController();
  final _subCounty = TextEditingController();
  final _latitude = TextEditingController();
  final _longitude = TextEditingController();
  final _api = MwangazaApiClient();

  String _crop = 'Maize';
  String _language = 'English';
  bool _submitting = false;

  @override
  void dispose() {
    _farmName.dispose();
    _farmerName.dispose();
    _phone.dispose();
    _county.dispose();
    _subCounty.dispose();
    _latitude.dispose();
    _longitude.dispose();
    super.dispose();
  }

  Future<void> _submit() async {
    if (!_formKey.currentState!.validate()) return;

    setState(() => _submitting = true);
    try {
      await _api.post<Map<String, dynamic>>(
        'farms',
        data: {
          'name': _farmName.text.trim(),
          'farmer': _farmerName.text.trim(),
          'phone': _phone.text.trim(),
          'county': _county.text.trim(),
          'sub_county': _subCounty.text.trim(),
          'crop': _crop,
          'latitude': double.parse(_latitude.text.trim()),
          'longitude': double.parse(_longitude.text.trim()),
          'preferred_language': _language,
        },
      );
      if (!mounted) return;
      _formKey.currentState!.reset();
      _farmName.clear();
      _farmerName.clear();
      _phone.clear();
      _county.clear();
      _subCounty.clear();
      _latitude.clear();
      _longitude.clear();
      setState(() {
        _crop = 'Maize';
        _language = 'English';
      });
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Farm saved. Farm advice is being prepared.'), backgroundColor: Colors.green),
      );
    } on DioException catch (error) {
      final body = error.response?.data;
      final message = body is Map ? body['error']?.toString() : null;
      _showError(message ?? 'Could not save the farm. Check the backend connection and try again.');
    } catch (_) {
      _showError('Could not save the farm. Please try again.');
    } finally {
      if (mounted) setState(() => _submitting = false);
    }
  }

  void _showError(String message) {
    if (!mounted) return;
    ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(message), backgroundColor: Colors.red));
  }

  String? _required(String? value, String label) {
    if (value == null || value.trim().isEmpty) return '$label is required';
    return null;
  }

  String? _phoneValidator(String? value) {
    final missing = _required(value, 'Phone number');
    if (missing != null) return missing;
    if (!RegExp(r'^\+[1-9]\d{7,14}$').hasMatch(value!.trim())) {
      return 'Use international format, e.g. +254712345678';
    }
    return null;
  }

  String? _coordinateValidator(String? value, String label, double min, double max) {
    final parsed = double.tryParse(value?.trim() ?? '');
    if (parsed == null || parsed < min || parsed > max) {
      return '$label must be between $min and $max';
    }
    return null;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Register a Farm')),
      body: Center(
        child: ConstrainedBox(
          constraints: const BoxConstraints(maxWidth: 760),
          child: SingleChildScrollView(
            padding: const EdgeInsets.all(24),
            child: Form(
              key: _formKey,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text('Add a farm for advice and SMS alerts', style: Theme.of(context).textTheme.headlineSmall),
                  const SizedBox(height: 8),
                  const Text('The phone number is used only when an approved alert is sent to this farmer.'),
                  const SizedBox(height: 24),
                  _field(_farmName, 'Farm name', validator: (value) => _required(value, 'Farm name')),
                  _field(_farmerName, 'Farmer name', validator: (value) => _required(value, 'Farmer name')),
                  _field(_phone, 'Farmer phone number', keyboardType: TextInputType.phone, hint: '+254712345678', validator: _phoneValidator),
                  _field(_county, 'County', validator: (value) => _required(value, 'County')),
                  _field(_subCounty, 'Sub-county (optional)'),
                  DropdownButtonFormField<String>(
                    value: _crop,
                    decoration: const InputDecoration(labelText: 'Crop'),
                    items: const ['Maize', 'Rice', 'Beans', 'Tea', 'Sugarcane', 'Cassava']
                        .map((crop) => DropdownMenuItem(value: crop, child: Text(crop)))
                        .toList(),
                    onChanged: (value) => setState(() => _crop = value!),
                  ),
                  const SizedBox(height: 16),
                  Row(
                    children: [
                      Expanded(child: _field(_latitude, 'Latitude', keyboardType: const TextInputType.numberWithOptions(decimal: true, signed: true), hint: '-0.0917', validator: (value) => _coordinateValidator(value, 'Latitude', -5.1, 2.5))),
                      const SizedBox(width: 16),
                      Expanded(child: _field(_longitude, 'Longitude', keyboardType: const TextInputType.numberWithOptions(decimal: true, signed: true), hint: '34.7680', validator: (value) => _coordinateValidator(value, 'Longitude', 28.95, 36.7))),
                    ],
                  ),
                  const SizedBox(height: 6),
                  Text('Kijani live agro-climate coverage: latitude -5.1 to 2.5, longitude 28.95 to 36.7.', style: Theme.of(context).textTheme.bodySmall),
                  const SizedBox(height: 16),
                  DropdownButtonFormField<String>(
                    value: _language,
                    decoration: const InputDecoration(labelText: 'Preferred SMS language'),
                    items: const ['English', 'Swahili']
                        .map((language) => DropdownMenuItem(value: language, child: Text(language)))
                        .toList(),
                    onChanged: (value) => setState(() => _language = value!),
                  ),
                  const SizedBox(height: 28),
                  SizedBox(
                    width: double.infinity,
                    child: FilledButton.icon(
                      onPressed: _submitting ? null : _submit,
                      icon: _submitting
                          ? const SizedBox(height: 18, width: 18, child: CircularProgressIndicator(strokeWidth: 2))
                          : const Icon(Icons.add_location_alt_outlined),
                      label: Text(_submitting ? 'Saving farm...' : 'Save farm and get advice'),
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }

  Widget _field(
    TextEditingController controller,
    String label, {
    TextInputType? keyboardType,
    String? hint,
    String? Function(String?)? validator,
  }) {
    return Padding(
      padding: const EdgeInsets.only(bottom: 16),
      child: TextFormField(
        controller: controller,
        keyboardType: keyboardType,
        decoration: InputDecoration(labelText: label, hintText: hint),
        validator: validator,
      ),
    );
  }
}
