import 'package:flutter/material.dart';

class FillProfileScreen extends StatefulWidget {
  final void Function(String name, String about) onDone;

  const FillProfileScreen({super.key, required this.onDone});

  @override
  State<FillProfileScreen> createState() => _FillProfileScreenState();
}

class _FillProfileScreenState extends State<FillProfileScreen> {
  final nameController = TextEditingController();
  final aboutController = TextEditingController();

  bool get isNameValid => nameController.text.trim().isNotEmpty;

  void _onChanged() => setState(() {}); // обновляет состояние для кнопки

  @override
  void initState() {
    super.initState();
    nameController.addListener(_onChanged);
  }

  @override
  void dispose() {
    nameController.removeListener(_onChanged);
    nameController.dispose();
    aboutController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Tell us about yourself')),
      body: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          children: [
            // 👤 Имя
            TextField(
              controller: nameController,
              decoration: const InputDecoration(
                labelText: 'Your Name *',
                border: OutlineInputBorder(),
              ),
            ),
            const SizedBox(height: 16),
            // 📄 Описание
            TextField(
              controller: aboutController,
              decoration: const InputDecoration(
                labelText: 'About You',
                border: OutlineInputBorder(),
              ),
              maxLines: 4,
            ),
            const Spacer(),
            // 🔘 Кнопка Далее
            SizedBox(
              width: double.infinity,
              height: 52,
              child: ElevatedButton(
                onPressed: isNameValid
                    ? () => widget.onDone(
                          nameController.text.trim(),
                          aboutController.text.trim(),
                        )
                    : null,
                child: const Text('Next'),
              ),
            )
          ],
        ),
      ),
    );
  }
}
