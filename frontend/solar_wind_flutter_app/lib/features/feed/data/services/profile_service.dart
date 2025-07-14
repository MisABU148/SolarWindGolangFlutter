// import 'package:dio/dio.dart';
// import '../models/user.dart';

// class ProfileService {
//   final Dio dio;

//   ProfileService({Dio? dio}) : dio = dio ?? Dio();

//   Future<User> fetchUser(int userId) async {
//     final response = await dio.get(
//       'https://solar-wind-gymbro.ru/profiles/api/users/$userId',
//       options: Options(headers: {
//         'Authorization-telegram-id': '637451540',
//       }),
//     );

//     return User.fromJson(response.data);
//   }
// }

import 'dart:async';
import '../models/user.dart';

class ProfileService {
  Future<User> fetchUser(int userId) async {
    print('Mock Profile Request for userId: $userId');

    await Future.delayed(const Duration(seconds: 1));

    return User(
      id: userId,
      username: 'mock_user_$userId',
      description: 'This is a mock user profile for testing purposes.',
      cityName: 'Mock City',
      preferredGymTime: [9, 18],
      sportName: ['Basketball', 'Yoga'],
    );
  }
}
