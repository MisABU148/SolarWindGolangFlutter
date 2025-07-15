import 'package:flutter/material.dart';
import 'enter_code.dart';
import 'package:url_launcher/url_launcher.dart';

class TelegramAuthScreen extends StatelessWidget {
  final void Function(BuildContext context) onNext;

  const TelegramAuthScreen({super.key, required this.onNext});

  void _openTelegramBot(BuildContext context) async {
    const botUrl = 'https://t.me/SolarWindAuthorization_bot';
    if (await canLaunchUrl(Uri.parse(botUrl))) {
      await launchUrl(Uri.parse(botUrl));
      onNext(context);
    } else {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text("Could not open Telegram.")),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Telegram Authorization')),
      body: Padding(
        padding: const EdgeInsets.all(24),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text(
              'To authorize, click the button below and write /start to the Telegram bot. You will receive a code.',
              style: TextStyle(fontSize: 16),
              textAlign: TextAlign.center,
            ),
            const SizedBox(height: 40),
            ElevatedButton(
              onPressed: () => _openTelegramBot(context),
              child: const Text('@SolarWindAuthorization_bot'),
            ),
          ],
        ),
      ),
    );
  }
}
