import 'package:dio/dio.dart';
import '../models/registration_data.dart';

class RegistrationService {
  final Dio dio;

  RegistrationService({Dio? dio}) : dio = dio ?? Dio();

  Future<void> register(RegistrationData data) async {
    print(data.toJson());
    final response = await dio.post(
      'https://solar-wind-gymbro.ru/profiles/api/me',
      data: data.toJson(),
      options: Options(
        headers: {
          'Authorization':
              'bYpCKtI2Mgs0zHbpEq7m1k5Nq3KX7QLrL4LTFjl1X5swvQjCHuvCkBAFBctAPK0VDmBJTK5DzUa8f1oNf5qWgzII2Z9xzmsfXJxvsXuRLqHTYiyQweggdfKylGcHgKT5LhknMAYKkf5xRokYxCTBJVfr4mkMOO2R8JGUFNsyNOBs4VpITSwpOqfNoGaHGBV3n5ciOsI014gNWYfjuxQXepOtAkG8W9hwk2c9aLhDd1ULj1VUT6Hwcl1ky0gDooKV',
          'Authorization-telegram-id': '637451540',
        },
      ),
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to register. Status: ${response.statusCode}');
    }
  }
}
