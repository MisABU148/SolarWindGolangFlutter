// import 'package:dio/dio.dart';
// import '../models/user.dart';

// class LikeNotificationService {
//   final Dio dio;

//   LikeNotificationService({Dio? dio}) : dio = dio ?? Dio();

//   Future<List<User>> getLikedMeUsers() async {
//     final response = await dio.get(
//       'https://solar-wind-gymbro.ru/profiles/api/liked-me',
//       options: Options(headers: {
//         'Authorization-telegram-id': '637451540', 
//       }),
//     );

//     final List<dynamic> data = response.data;
//     return data.map((json) => User.fromJson(json)).toList();
//   }
// }

import 'dart:async';
import '../models/user.dart';

class LikeNotificationService {
  Future<List<User>> getLikedMeUsers() async {
    await Future.delayed(const Duration(milliseconds: 300));

    return [
      User(
        id: 1,
        username: 'alice',
        description: 'Loves yoga and morning runs',
        cityName: 'Berlin',
        sportName: ['Yoga', 'Running'],
        preferredGymTime: [1],
      ),
      User(
        id: 2,
        username: 'bob',
        description: 'Crossfit enthusiast',
        cityName: 'Munich',
        sportName: ['Crossfit'],
        preferredGymTime: [1],
      ),
    ];
  }
}
