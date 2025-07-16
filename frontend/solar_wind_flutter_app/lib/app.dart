import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'features/auth/presentation/screens/welcome_screen.dart';
import 'features/auth/presentation/screens/tgbot_auth.dart';
import 'features/auth/presentation/screens/enter_code.dart';
import 'features/feed/presentation/screen/user_feed_screen.dart';
import 'features/auth/presentation/state/registration_provider.dart';

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  Future<Widget> _decideStartScreen() async {
    final prefs = await SharedPreferences.getInstance();

    final isRegistered = prefs.getBool('isRegistered') ?? false;

    if (isRegistered) {
      return const UserFeedScreen();
    } else {
      return WelcomeScreen();
    }
  }

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => RegistrationProvider(),
      child: MaterialApp(
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
      ),
    );
  }
}
