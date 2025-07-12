import 'package:flutter/material.dart';
import 'features/auth/presentation/screens/welcome_screen.dart';

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
      home: WelcomeScreen(
        onNext: () {
          // TODO: Навигация на следующий экран
          // например: Navigator.push(...)
        },
      ),
    );
  }
}
