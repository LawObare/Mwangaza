import '../models/sms.dart';
import 'api_service.dart';

class SmsService {
  final ApiService _api;

  SmsService(this._api);

  Future<List<SmsMessage>> getSmsHistory() async {
    final data = await _api.getList('/sms');
    return data
        .map((e) => SmsMessage.fromJson(e as Map<String, dynamic>))
        .toList();
  }

  Future<Map<String, dynamic>> sendSms(
      int farmId, String message, String phoneNumber) async {
    return _api.post('/sms/send', {
      'farm_id': farmId,
      'message': message,
      'phone_number': phoneNumber,
    });
  }
}