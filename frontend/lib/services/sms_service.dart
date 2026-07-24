import '../models/sms.dart';
import 'api_service.dart';

class SmsService {
  SmsService(this._api);

  final ApiService _api;

  Future<List<SmsMessage>> listSms() async {
    final data = await _api.get('/sms');
    return (data as List? ?? const [])
        .whereType<Map>()
        .map((item) => SmsMessage.fromJson(Map<String, dynamic>.from(item)))
        .toList();
  }

  Future<SmsMessage> send({
    required int farmId,
    String message = '',
    String phoneNumber = '',
  }) async {
    final data = await _api.post('/sms/send', {
      'farm_id': farmId,
      'message': message,
      'phone_number': phoneNumber,
    });
    return SmsMessage.fromJson(Map<String, dynamic>.from(data as Map));
  }
}
