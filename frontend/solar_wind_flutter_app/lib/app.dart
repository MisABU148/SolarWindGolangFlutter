import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'features/auth/data/services/registration_service.dart';
import 'features/auth/presentation/screens/about_me.dart';
import 'features/auth/presentation/screens/city_screen.dart';
import 'features/auth/presentation/screens/enter_code.dart';
import 'features/auth/presentation/screens/gender_and_birth_screen.dart';
import 'features/auth/presentation/screens/sport_screen.dart';
import 'features/auth/presentation/screens/tgbot_auth.dart';
import 'features/auth/presentation/screens/welcome_screen.dart';
import 'features/auth/presentation/state/registration_provider.dart';
import 'features/feed/presentation/screen/user_feed_screen.dart';

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  Future<Widget> _decideStartScreen() async {
    final prefs = await SharedPreferences.getInstance();

    final isRegistered = prefs.getBool('isRegistered') ?? false;
    final hasTelegramToken =
        prefs.containsKey('token') && prefs.containsKey('telegram_id');

    if (isRegistered) {
      return const UserFeedScreen();
    } else if (hasTelegramToken) {
      return _buildWelcomeFlow();
    } else {
      return TelegramAuthScreen(
        onNext: (ctx) {
          Navigator.pushReplacement(
            ctx,
            MaterialPageRoute(
              builder: (_) => EnterCodeScreen(
                onSuccess: (ctx) {
                  Navigator.pushReplacement(ctx, MaterialPageRoute(builder: (_) => _buildWelcomeFlow()));
                },
              ),
            ),
          );
        },
      );
    }
  }

  Widget _buildWelcomeFlow() {
    return WelcomeScreen(
      onNext: (ctx) {
        Navigator.push(
          ctx,
          MaterialPageRoute(
            builder: (_) => ChooseCityScreen(
              onCitySelected: (city) {
                Provider.of<RegistrationProvider>(ctx, listen: false).setCity(city);
                Navigator.push(
                  ctx,
                  MaterialPageRoute(
                    builder: (_) => ChooseSportScreen(
                      selectedCity: city,
                      onDone: (selectedSports) {
                        Provider.of<RegistrationProvider>(ctx, listen: false)
                            .setSports(selectedSports);
                        Navigator.push(
                          ctx,
                          MaterialPageRoute(
                            builder: (_) => FillProfileScreen(
                              onDone: (name, about) {
                                final provider =
                                    Provider.of<RegistrationProvider>(ctx, listen: false);
                                provider.setName(name);
                                provider.setAbout(about);
                                Navigator.push(
                                  ctx,
                                  MaterialPageRoute(
                                    builder: (_) => GenderAndBirthScreen(
                                      onNext: (gender, preferredGender, birthDate) async {
                                        final provider =
                                            Provider.of<RegistrationProvider>(ctx, listen: false);
                                        provider.setGender(gender);
                                        provider.setPreferredGender(preferredGender);
                                        provider.setBirthDate(birthDate);

                                        try {
                                          final service = RegistrationService();
                                          await service.register(provider.data);

                                          final prefs = await SharedPreferences.getInstance();
                                          await prefs.setBool('isRegistered', true);

                                          Navigator.pushReplacement(
                                            ctx,
                                            MaterialPageRoute(
                                              builder: (_) => const UserFeedScreen(),
                                            ),
                                          );
                                        } catch (e) {
                                          ScaffoldMessenger.of(ctx).showSnackBar(
                                            SnackBar(
                                              content: Text('Registration failed: $e'),
                                            ),
                                          );
                                        }
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
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'SolarWind App',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.pink),
        useMaterial3: true,
      ),
      debugShowCheckedModeBanner: false,
      home: FutureBuilder<Widget>(
        future: _decideStartScreen(),
        builder: (context, snapshot) {
          if (!snapshot.hasData) {
            return const Scaffold(
              body: Center(child: CircularProgressIndicator()),
            );
          }
          return snapshot.data!;
        },
      ),
    );
  }
}
