import 'package:dio/dio.dart';
import '../../../auth/data/models/registration_data.dart';

class ProfileUpdateService {
  final Dio dio;

  ProfileUpdateService({Dio? dio}) : dio = dio ?? Dio();

  Future<void> updateProfile({
    required int userId,
    required RegistrationData data,
  }) async {
    final payload = data.toJson();

    final response = await dio.put(
      'https://solar-wind-gymbro.ru/profiles/api/users/$userId',
      data: payload,
      options: Options(
        headers: {
          'Authorization-telegram-id': '637451540',
        },
      ),
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to update profile. Status: ${response.statusCode}');
    }
  }
}
