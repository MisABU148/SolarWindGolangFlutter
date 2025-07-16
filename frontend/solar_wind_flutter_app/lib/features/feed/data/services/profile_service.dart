import 'package:dio/dio.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../models/user.dart';
import 'package:solar_wind_flutter_app/mock/mock_users.dart';

class ProfileService {
  final Dio dio;

  ProfileService({Dio? dio}) : dio = dio ?? Dio();

  Future<User> fetchUser(int userId) async {
  final mockUser = _findMockUser(userId);
  if (mockUser != null) {
    return mockUser;
  }

  final prefs = await SharedPreferences.getInstance();
  final telegramId = prefs.getString('telegram_id');
  final token = prefs.getString('token');

  if (telegramId == null || token == null) {
    throw Exception('Missing token or telegram_id in SharedPreferences');
  }

  final response = await dio.get(
    'https://solar-wind-gymbro.ru/profiles/api/me',
    options: Options(headers: {
      'Authorization-telegram-id': userId,
      'Authorize': token,
    }),
  );

  return User.fromJson(response.data);
}

User? _findMockUser(int userId) {
  try {
    return mockUsers.firstWhere((u) => u.id == userId);
  } catch (_) {
    return null;
  }
}

}
