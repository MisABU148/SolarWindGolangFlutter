import 'package:dio/dio.dart';
import '../models/user.dart';

class FeedService {
  final Dio dio;

  FeedService({Dio? dio}) : dio = dio ?? Dio();

  Future<List<User>> fetchFeed() async {
    final response = await dio.get(
      'https://solar-wind-gymbro.ru/deckShuffle/api/create-deck',
      options: Options(
        headers: {
          'Authorization-telegram-id': '637451540',
        },
      ),
    );

    if (response.statusCode == 200) {
      final data = response.data as List;
      return data.map((json) => User.fromJson(json)).toList();
    } else {
      throw Exception('Failed to load feed');
    }
  }
}
