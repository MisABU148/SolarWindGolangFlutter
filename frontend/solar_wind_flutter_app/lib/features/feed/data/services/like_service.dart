import 'package:dio/dio.dart';
import 'package:shared_preferences/shared_preferences.dart';

class LikeService {
  final Dio dio;

  LikeService({Dio? dio}) : dio = dio ?? Dio();

  Future<void> likeUser(int user2) async {
    final prefs = await SharedPreferences.getInstance();
    final telegramId = prefs.getString('telegram_id');
    final token = prefs.getString('token');

    if (telegramId == null || token == null) {
      throw Exception('Missing token or telegram_id in SharedPreferences');
    }

    await dio.post(
      'https://solar-wind-gymbro.ru/likes/api/',
      options: Options(
        headers: {
          'receiver': user2,
          'Authorization-telegram-id': telegramId,
          'Authorize': token,
        },
      ),
    );
  }
}
