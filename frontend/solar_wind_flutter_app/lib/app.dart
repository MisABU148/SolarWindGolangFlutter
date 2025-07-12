import 'package:flutter/material.dart';
import 'features/auth/presentation/screens/welcome_screen.dart';
import 'features/auth/presentation/screens/city_screen.dart';
import 'features/auth/presentation/screens/sport_screen.dart';
import 'features/auth/presentation/screens/about_me.dart';

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Your App',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.pink),
        useMaterial3: true,
      ),
      debugShowCheckedModeBanner: false,
      home: Builder(
        builder: (context) => WelcomeScreen(
          onNext: (ctx) {
            Navigator.push(
              ctx,
              MaterialPageRoute(
                builder: (_) => ChooseCityScreen(
                  onCitySelected: (city) {
                    Navigator.push(
                      context,
                      MaterialPageRoute(
                        builder: (_) => ChooseSportScreen(
                          selectedCity: city,
                          onDone: (selectedSports) {
                            print('City: ${city.name}');
                            print('Selected sports: ${selectedSports.map((s) => s.name).join(", ")}');
                            Navigator.push(
                              context,
                              MaterialPageRoute(
                                builder: (_) => FillProfileScreen(
                                  onDone: (name, about) {
                                    print('Имя: $name');
                                    print('О себе: $about');
                                    // здесь можно идти дальше — например, на экран загрузки фото или завершения профиля
                                  },
                                ),
                              ),
                            );
                          },
                        ),
                      ),
                    );
                  },
                ),
              ),
            );
          },
        ),
      ),
    );
  }
}
