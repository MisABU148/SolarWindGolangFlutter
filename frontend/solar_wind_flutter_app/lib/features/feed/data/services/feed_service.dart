// import 'package:dio/dio.dart';
// import '../models/user.dart';

// class FeedService {
//   final Dio dio;

//   FeedService({Dio? dio}) : dio = dio ?? Dio();

//   Future<List<User>> fetchFeed() async {
//     final response = await dio.get(
//       'https://solar-wind-gymbro.ru/deckShuffle/api/create-deck',
//       options: Options(
//         headers: {
//           'Authorization-telegram-id': '637451540',
//         },
//       ),
//     );

//     if (response.statusCode == 200) {
//       final data = response.data as List;
//       return data.map((json) => User.fromJson(json)).toList();
//     } else {
//       throw Exception('Failed to load feed');
//     }
//   }
// }

import 'package:dio/dio.dart';
import '../models/user.dart';

class FeedService {
  final Dio dio;

  FeedService({Dio? dio}) : dio = dio ?? Dio();

  Future<List<User>> fetchFeed() async {
    await Future.delayed(const Duration(milliseconds: 500));

    return [
      User(
        id: 1,
        username: 'alex_fitness',
        description: 'I love morning workouts 💪',
        cityName: 'Berlin',
        preferredGymTime: [9, 10],
        sportName: ['Running', 'CrossFit'],
      ),
      User(
        id: 2,
        username: 'yoga_girl',
        description: 'Namaste 🌿',
        cityName: 'Munich',
        preferredGymTime: [18],
        sportName: ['Yoga', 'Pilates'],
      ),
      User(
        id: 3,
        username: 'night_lifter',
        description: 'Deadlifts at midnight 🏋️‍♂️',
        cityName: 'Hamburg',
        preferredGymTime: [23],
        sportName: ['Powerlifting'],
      ),
    ];
  }
}
