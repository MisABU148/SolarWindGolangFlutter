import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'features/auth/presentation/state/registration_provider.dart';
import 'app.dart';
import 'package:hive_flutter/hive_flutter.dart';
import 'features/feed/data/models/user.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  
  await Hive.initFlutter();
  Hive.registerAdapter(UserAdapter()); 
  await Hive.openBox('feedBox');

  runApp(
    ChangeNotifierProvider(
      create: (_) => RegistrationProvider(),
      child: const MyApp(),
    ),
  );
}
