import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'features/auth/presentation/state/registration_provider.dart';
import 'app.dart';

void main() {
  runApp(
    ChangeNotifierProvider(
      create: (_) => RegistrationProvider(),
      child: const MyApp(),
    ),
  );
}

