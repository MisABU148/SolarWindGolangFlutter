import 'package:dio/dio.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:hive/hive.dart';
import '../models/user.dart';

class FeedService {
  final Dio dio;

  FeedService({Dio? dio}) : dio = dio ?? Dio();

  Future<List<User>> fetchFeed() async {
    final prefs = await SharedPreferences.getInstance();
    final telegramId = prefs.getString('telegram_id');
    final token = prefs.getString('token');

    if (telegramId == null || token == null) {
      throw Exception('Missing token or telegram_id in SharedPreferences');
    }

    try {
      final response = await dio.get(
        'https://solar-wind-gymbro.ru/deckShuffle/api/create-deck',
        options: Options(headers: {
          'Authorization-telegram-id': telegramId,
          'Authorize': token,
        }),
      );

      if (response.statusCode == 200) {
        final data = response.data as List;
        final users = data.map((json) => User.fromJson(json)).toList();

        final enriched = [...users, ..._mockUsers];

        // 💾 Save to Hive
        final box = Hive.box('feedBox');
        await box.put('users', enriched);

        return enriched;
      } else {
        throw Exception('Failed to fetch feed. Status: ${response.statusCode}');
      }
    } catch (e) {
      // 🌐 Fallback to cached
      print('Failed to fetch online. Using cached. Error: $e');
      final box = Hive.box('feedBox');
      final cachedUsers = box.get('users') as List<User>?;

      if (cachedUsers != null) return cachedUsers;
      throw Exception('No internet and no cached data available');
    }
  }

  List<User> get _mockUsers => [
        User(
          id: 1001,
          username: 'fake_user_1',
          description: 'Just testing things out',
          cityName: 'Fake City',
          preferredGymTime: [10],
          sportName: ['TestSport'],
        ),
        User(
          id: 1002,
          username: 'test_account_2',
          description: 'Here for fun 🧪',
          cityName: 'SimCity',
          preferredGymTime: [17],
          sportName: ['Debugging'],
        ),
        User(
          id: 1003,
          username: 'gym_bot_3000',
          description: '💥 Automated fitness bot',
          cityName: 'Botland',
          preferredGymTime: [6],
          sportName: ['Powerlifting', 'Running'],
        ),
        User(
          id: 1004,
          username: 'qa_guru',
          description: 'Finding bugs and gains 🐛🏋️',
          cityName: 'QApolis',
          preferredGymTime: [8],
          sportName: ['CrossFit'],
        ),
        User(
          id: 1005,
          username: 'frontend_beast',
          description: 'CSS by day, squats by night',
          cityName: 'Codeville',
          preferredGymTime: [20],
          sportName: ['Yoga', 'Climbing'],
        ),
      ];
}



// import 'package:dio/dio.dart';
// import '../models/user.dart';

// class FeedService {
//   final Dio dio;

//   FeedService({Dio? dio}) : dio = dio ?? Dio();

//   Future<List<User>> fetchFeed() async {
//     await Future.delayed(const Duration(milliseconds: 500));

//     return [
//       User(
//         id: 1,
//         username: 'alex_fitness',
//         description: 'I love morning workouts 💪',
//         cityName: 'Berlin',
//         preferredGymTime: [9, 10],
//         sportName: ['Running', 'CrossFit'],
//       ),
//       User(
//         id: 2,
//         username: 'yoga_girl',
//         description: 'Namaste 🌿',
//         cityName: 'Munich',
//         preferredGymTime: [18],
//         sportName: ['Yoga', 'Pilates'],
//       ),
//       User(
//         id: 3,
//         username: 'night_lifter',
//         description: 'Deadlifts at midnight 🏋️‍♂️',
//         cityName: 'Hamburg',
//         preferredGymTime: [23],
//         sportName: ['Powerlifting'],
//       ),
//     ];
//   }
// }
