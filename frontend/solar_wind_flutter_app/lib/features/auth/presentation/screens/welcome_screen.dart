import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:solar_wind_flutter_app/features/auth/presentation/screens/city_screen.dart';
import 'package:solar_wind_flutter_app/features/auth/presentation/state/registration_provider.dart';

class WelcomeScreen extends StatelessWidget {
  const WelcomeScreen({super.key});

  static const mainPinkColor = Color.fromRGBO(243, 75, 213, 1); // #F34BD5
  static const mainYellowColor = Color.fromRGBO(251, 218, 85, 1); // #FBDA55

  void _goToCitySelection(BuildContext context) {
    Navigator.push(
      context,
      MaterialPageRoute(
        builder: (_) => ChooseCityScreen(),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        width: double.infinity,
        height: double.infinity,
        decoration: const BoxDecoration(
          gradient: LinearGradient(
            colors: [mainPinkColor, mainYellowColor],
            begin: Alignment.topCenter,
            end: Alignment.bottomCenter,
          ),
        ),
        child: SafeArea(
          child: Column(
            children: [
              const SizedBox(height: 40),
              Image.asset(
                'assets/images/logo.png',
                width: 260,
                height: 140,
                fit: BoxFit.contain,
              ),
              const SizedBox(height: 10),
              const Text(
                'Solar Wind',
                style: TextStyle(
                  color: mainPinkColor,
                  fontSize: 50,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 20),
              const Padding(
                padding: EdgeInsets.symmetric(horizontal: 40.0),
                child: Text(
                  'Move. Meet. Repeat.',
                  textAlign: TextAlign.center,
                  style: TextStyle(
                    fontSize: 24,
                    fontWeight: FontWeight.w500,
                    color: Colors.white,
                    shadows: [
                      Shadow(
                        color: mainPinkColor,
                        offset: Offset(1, 1),
                        blurRadius: 0,
                      )
                    ],
                  ),
                ),
              ),
              const Spacer(),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 20.0),
                child: SizedBox(
                  height: 52,
                  width: double.infinity,
                  child: ElevatedButton(
                    style: ElevatedButton.styleFrom(
                      backgroundColor: Colors.white,
                      foregroundColor: Colors.black,
                      textStyle: const TextStyle(fontSize: 16),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(8),
                      ),
                    ),
                    onPressed: () => _goToCitySelection(context),
                    child: const Text('Find your gym bro'),
                  ),
                ),
              ),
              const SizedBox(height: 40),
            ],
          ),
        ),
      ),
    );
  }
}
