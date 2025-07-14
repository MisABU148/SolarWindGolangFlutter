import 'package:flutter/material.dart';
import 'features/auth/presentation/screens/welcome_screen.dart';
import 'features/auth/presentation/screens/city_screen.dart';
import 'features/auth/presentation/screens/sport_screen.dart';
import 'features/auth/presentation/screens/about_me.dart';
import 'features/auth/data/services/registration_service.dart';
import 'features/auth/presentation/state/registration_provider.dart';
import 'package:provider/provider.dart';
import 'features/auth/presentation/screens/gender_and_birth_screen.dart';


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
                    Provider.of<RegistrationProvider>(context, listen: false).setCity(city);
                    Navigator.push(
                      context,
                      MaterialPageRoute(
                        builder: (_) => ChooseSportScreen(
                          selectedCity: city,
                          onDone: (selectedSports) {
                            Provider.of<RegistrationProvider>(context, listen: false).setSports(selectedSports);
                            Navigator.push(
                              context,
                              MaterialPageRoute(
                                builder: (_) => FillProfileScreen(
                                  onDone: (name, about) {
                                    final provider = Provider.of<RegistrationProvider>(context, listen: false);
                                    provider.setName(name);
                                    provider.setAbout(about);

                                    Navigator.push(
                                      context,
                                      MaterialPageRoute(
                                        builder: (_) => GenderAndBirthScreen(
                                          onNext: (gender, preferredGender, birthDate) async {
                                            final provider = Provider.of<RegistrationProvider>(context, listen: false);
                                            provider.setGender(gender);
                                            provider.setPreferredGender(preferredGender);
                                            provider.setBirthDate(birthDate);

                                            try {
                                              final service = RegistrationService();
                                              await service.register(provider.data);

                                              ScaffoldMessenger.of(context).showSnackBar(
                                                const SnackBar(content: Text('Registration successful!')),
                                              );
                                            } catch (e) {
                                              ScaffoldMessenger.of(context).showSnackBar(
                                                SnackBar(content: Text('Registration failed: $e')),
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
        ),
      ),
    );
  }
}
