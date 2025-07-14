import 'package:flutter/material.dart';

class GenderAndBirthScreen extends StatefulWidget {
  final void Function(String gender, String preferredGender, DateTime birthDate) onNext;

  const GenderAndBirthScreen({super.key, required this.onNext});

  @override
  State<GenderAndBirthScreen> createState() => _GenderAndBirthScreenState();
}

class _GenderAndBirthScreenState extends State<GenderAndBirthScreen> {
  String? gender;
  String? preferredGender;
  DateTime? birthDate;

  void _selectDate() async {
    final now = DateTime.now();
    final picked = await showDatePicker(
      context: context,
      initialDate: DateTime(now.year - 18),
      firstDate: DateTime(now.year - 100),
      lastDate: now,
    );

    if (picked != null) {
      setState(() => birthDate = picked);
    }
  }

  bool get isValid =>
      gender != null && preferredGender != null && birthDate != null;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Tell us about you')),
      body: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          children: [
            const Text('Your Gender'),
            Row(
              children: [
                ChoiceChip(
                  label: const Text('Male'),
                  selected: gender == 'male',
                  onSelected: (_) => setState(() => gender = 'male'),
                ),
                const SizedBox(width: 8),
                ChoiceChip(
                  label: const Text('Female'),
                  selected: gender == 'female',
                  onSelected: (_) => setState(() => gender = 'female'),
                ),
              ],
            ),
            const SizedBox(height: 16),
            const Text('Preferred Gender'),
            Row(
              children: [
                ChoiceChip(
                  label: const Text('Male'),
                  selected: preferredGender == 'male',
                  onSelected: (_) => setState(() => preferredGender = 'male'),
                ),
                const SizedBox(width: 8),
                ChoiceChip(
                  label: const Text('Female'),
                  selected: preferredGender == 'female',
                  onSelected: (_) => setState(() => preferredGender = 'female'),
                ),
                const SizedBox(width: 8),
                ChoiceChip(
                  label: const Text('Doesn’t matter'),
                  selected: preferredGender == 'any',
                  onSelected: (_) => setState(() => preferredGender = 'any'),
                ),
              ],
            ),
            const SizedBox(height: 16),
            Row(
              children: [
                Text(birthDate != null
                    ? 'Birth Date: ${birthDate!.toLocal()}'.split(' ')[0]
                    : 'Choose your birth date'),
                const SizedBox(width: 8),
                ElevatedButton(
                  onPressed: _selectDate,
                  child: const Text('Select Date'),
                ),
              ],
            ),
            const Spacer(),
            ElevatedButton(
              onPressed: isValid
                  ? () => widget.onNext(gender!, preferredGender!, birthDate!)
                  : null,
              child: const Text('Next'),
            ),
          ],
        ),
      ),
    );
  }
}
