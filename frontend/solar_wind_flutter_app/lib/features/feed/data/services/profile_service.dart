import 'package:dio/dio.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../models/user.dart';

class ProfileService {
  final Dio dio;

  ProfileService({Dio? dio}) : dio = dio ?? Dio();

  Future<User> fetchUser(int userId) async {
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
}

// import 'dart:async';
// import '../models/user.dart';

// class ProfileService {
//   Future<User> fetchUser(int userId) async {
//     print('Mock Profile Request for userId: $userId');

//     await Future.delayed(const Duration(seconds: 1));

//     return User(
//       id: userId,
//       username: 'mock_user_$userId',
//       description: 'This is a mock user profile for testing purposes.',
//       cityName: 'Mock City',
//       preferredGymTime: [9, 18],
//       sportName: ['Basketball', 'Yoga'],
//     );
//   }
// }
