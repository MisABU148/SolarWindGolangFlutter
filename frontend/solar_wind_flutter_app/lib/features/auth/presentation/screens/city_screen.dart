import 'package:flutter/material.dart';
import '../../data/services/city_service.dart';
import '../../data/models/city.dart';

class ChooseCityScreen extends StatefulWidget {
  final void Function(City) onCitySelected;

  const ChooseCityScreen({super.key, required this.onCitySelected});

  @override
  State<ChooseCityScreen> createState() => _ChooseCityScreenState();
}

class _ChooseCityScreenState extends State<ChooseCityScreen> {
  final cityService = CityService();
  final controller = TextEditingController();
  List<City> cities = [];
  bool isLoading = false;

  void _search(String query) async {
    setState(() => isLoading = true);
    try {
      final results = await cityService.searchCities(query);
      print(query);
      setState(() => cities = results);
    } catch (e) {
      print(e);
    } finally {
      setState(() => isLoading = false);
    }
  }

  @override
  void dispose() {
    controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Choose City')),
      body: Column(
        children: [
          Padding(
            padding: const EdgeInsets.all(16.0),
            child: TextField(
              controller: controller,
              onChanged: _search,
              decoration: const InputDecoration(
                hintText: 'Search city...',
                border: OutlineInputBorder(),
              ),
            ),
          ),
          if (isLoading)
            const CircularProgressIndicator()
          else
            Expanded(
              child: ListView.builder(
                itemCount: cities.length,
                itemBuilder: (context, index) {
                  final city = cities[index];
                  return ListTile(
                    title: Text(city.name),
                    onTap: () => widget.onCitySelected(city),
                  );
                },
              ),
            ),
        ],
      ),
    );
  }
}
