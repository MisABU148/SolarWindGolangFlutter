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
  City? _selectedCity;

  void _search(String query) async {
    if (query.trim().isEmpty) {
      setState(() {
        cities = [];
      });
      return;
    }
    setState(() => isLoading = true);
    try {
      final results = await cityService.searchCities(query);
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
                  final isSelected = _selectedCity?.id == city.id;
                  return ListTile(
                    title: Text(city.name),
                    selected: isSelected,
                    onTap: () {
                      setState(() {
                        _selectedCity = city;
                      });
                      widget.onCitySelected(city);
                    },
                  );
                },
              ),
            ),
          if (_selectedCity != null)
            Padding(
              padding: const EdgeInsets.all(16.0),
              child: ElevatedButton(
                onPressed: () {
                  Navigator.pop(context, _selectedCity);
                },
                child: const Text('Далее'),
              ),
            ),
        ],
      ),
    );
  }
}

