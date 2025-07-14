import 'package:dio/dio.dart';

class LikeService {
  final Dio dio;

  LikeService({Dio? dio}) : dio = dio ?? Dio();

  Future<void> likeUser(int userId) async {
    await dio.post(
      'https://solar-wind-gymbro.ru/profiles/api/like/$userId',
      options: Options(headers: {
        'Authorization-telegram-id': '637451540',
      }),
    );
  }
}
