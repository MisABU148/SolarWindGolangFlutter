import 'package:flutter/material.dart';
import 'package:solar_wind_flutter_app/features/feed/data/services/profile_service.dart';
import 'package:solar_wind_flutter_app/features/feed/data/services/profile_update_service.dart';
import 'package:solar_wind_flutter_app/features/auth/data/models/registration_data.dart';
import 'package:solar_wind_flutter_app/features/auth/data/models/city.dart';
import 'package:solar_wind_flutter_app/features/auth/data/models/sport.dart';
import 'package:solar_wind_flutter_app/features/auth/data/services/city_service.dart';
import 'package:solar_wind_flutter_app/features/auth/data/services/sport_service.dart';

class MyProfileScreen extends StatefulWidget {
  final int userId;

  const MyProfileScreen({super.key, required this.userId});

  @override
  State<MyProfileScreen> createState() => _MyProfileScreenState();
}

class _MyProfileScreenState extends State<MyProfileScreen> {
  final _formKey = GlobalKey<FormState>();
  final _profileService = ProfileService();
  final _updateService = ProfileUpdateService();
  final _cityService = CityService();
  final _sportService = SportService();

  bool _isLoading = true;
  RegistrationData _data = RegistrationData();

  // Поиск и выбор города
  List<City> _citySearchResults = [];
  final TextEditingController _citySearchController = TextEditingController();
  City? _selectedCity;

  // Поиск и выбор спортов
  List<Sport> _sportSearchResults = [];
  final TextEditingController _sportSearchController = TextEditingController();

  List<Sport> _selectedSports = [];

  @override
  void initState() {
    super.initState();
    _loadInitialData();

    _citySearchController.addListener(_onCitySearchChanged);
    _sportSearchController.addListener(_onSportSearchChanged);
  }

  @override
  void dispose() {
    _citySearchController.dispose();
    _sportSearchController.dispose();
    super.dispose();
  }

  Future<void> _loadInitialData() async {
    try {
      final user = await _profileService.fetchUser(widget.userId);
      final cities = await _cityService.searchCities('');
      final sports = await _sportService.searchSports('');

      final city = cities.firstWhere(
        (c) => c.name == user.cityName,
        orElse: () => cities.first, // Город всегда есть
      );

      final selectedSports = sports.where((sport) => user.sportName.contains(sport.name)).toList();

      setState(() {
        _data.username = user.username;
        _data.description = user.description;
        _selectedCity = city;
        _data.cityId = city.id;
        _selectedSports = selectedSports;
        _data.sportId = selectedSports.map((s) => s.id).toList();
        _data.days = user.preferredGymTime;

        _citySearchResults = cities;
        _sportSearchResults = sports;

        _isLoading = false;

        _citySearchController.text = city.name;
      });
    } catch (e) {
      print("Error loading user: $e");
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text("Failed to load profile")),
      );
      setState(() => _isLoading = false);
    }
  }

  // Поиск города
  Future<void> _onCitySearchChanged() async {
    final query = _citySearchController.text.trim();
    if (query.isEmpty) {
      setState(() => _citySearchResults = []);
      return;
    }
    final results = await _cityService.searchCities(query);
    setState(() => _citySearchResults = results);
  }

  // Поиск спорта
  Future<void> _onSportSearchChanged() async {
    final query = _sportSearchController.text.trim();
    if (query.isEmpty) {
      setState(() => _sportSearchResults = []);
      return;
    }
    final results = await _sportService.searchSports(query);
    setState(() => _sportSearchResults = results);
  }

  void _selectCity(City city) {
    setState(() {
      _selectedCity = city;
      _data.cityId = city.id;
      _citySearchController.text = city.name;
      _citySearchResults = [];
      FocusScope.of(context).unfocus(); // Скрыть клавиатуру
    });
  }

  void _addSport(Sport sport) {
    if (_selectedSports.any((s) => s.id == sport.id)) return;
    setState(() {
      _selectedSports.add(sport);
      _data.sportId.add(sport.id);
      _sportSearchController.clear();
      _sportSearchResults = [];
    });
  }

  void _removeSport(Sport sport) {
    setState(() {
      _selectedSports.removeWhere((s) => s.id == sport.id);
      _data.sportId.removeWhere((id) => id == sport.id);
    });
  }

  void _saveProfile() async {
    if (_formKey.currentState!.validate()) {
      try {
        await _updateService.updateProfile(userId: widget.userId, data: _data);
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text("Profile updated")),
        );
      } catch (e) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text("Failed to update profile: $e")),
        );
      }
    }
  }

  Widget _buildCitySearchField(ThemeData theme) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text("City", style: theme.textTheme.labelLarge?.copyWith(fontWeight: FontWeight.bold)),
        const SizedBox(height: 8),
        TextFormField(
          controller: _citySearchController,
          decoration: InputDecoration(
            hintText: "Search city",
            border: const OutlineInputBorder(),
            fillColor: theme.colorScheme.surface,
            filled: true,
          ),
          validator: (value) {
            if (_selectedCity == null) return 'Select a city';
            return null;
          },
          style: theme.textTheme.bodyMedium,
        ),
        if (_citySearchResults.isNotEmpty)
          Container(
            constraints: const BoxConstraints(maxHeight: 150),
            decoration: BoxDecoration(
              border: Border.all(color: theme.colorScheme.outline),
              borderRadius: BorderRadius.circular(4),
              color: theme.colorScheme.surface,
            ),
            child: ListView.builder(
              shrinkWrap: true,
              itemCount: _citySearchResults.length,
              itemBuilder: (context, index) {
                final city = _citySearchResults[index];
                return ListTile(
                  title: Text(city.name, style: theme.textTheme.bodyMedium),
                  onTap: () => _selectCity(city),
                );
              },
            ),
          ),
      ],
    );
  }

  Widget _buildSportSearchField(ThemeData theme) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text("Sports", style: theme.textTheme.labelLarge?.copyWith(fontWeight: FontWeight.bold)),
        const SizedBox(height: 8),
        Wrap(
          spacing: 8,
          runSpacing: 4,
          children: _selectedSports
              .map(
                (sport) => Chip(
                  label: Text(
                    sport.name,
                    style: theme.textTheme.bodyMedium?.copyWith(color: theme.colorScheme.onSurface),
                  ),
                  backgroundColor: theme.colorScheme.secondaryContainer,
                  onDeleted: () => _removeSport(sport),
                ),
              )
              .toList(),
        ),
        const SizedBox(height: 8),
        TextField(
          controller: _sportSearchController,
          decoration: InputDecoration(
            hintText: "Search sports",
            border: const OutlineInputBorder(),
            fillColor: theme.colorScheme.surface,
            filled: true,
          ),
          style: theme.textTheme.bodyMedium,
        ),
        if (_sportSearchResults.isNotEmpty)
          Container(
            constraints: const BoxConstraints(maxHeight: 150),
            decoration: BoxDecoration(
              border: Border.all(color: theme.colorScheme.outline),
              borderRadius: BorderRadius.circular(4),
              color: theme.colorScheme.surface,
            ),
            child: ListView.builder(
              shrinkWrap: true,
              itemCount: _sportSearchResults.length,
              itemBuilder: (context, index) {
                final sport = _sportSearchResults[index];
                return ListTile(
                  title: Text(sport.name, style: theme.textTheme.bodyMedium),
                  onTap: () => _addSport(sport),
                );
              },
            ),
          ),
      ],
    );
  }

  @override
Widget build(BuildContext context) {
  final theme = Theme.of(context);

  return Scaffold(
    appBar: AppBar(
      title: const Text("My Profile"),
      backgroundColor: theme.colorScheme.primaryContainer,
      foregroundColor: theme.colorScheme.onPrimaryContainer,
      elevation: 0,
    ),
    body: _isLoading
        ? Center(
            child: CircularProgressIndicator(color: theme.colorScheme.primary),
          )
        : Padding(
            padding: const EdgeInsets.all(16),
            child: Form(
              key: _formKey,
              child: ListView(
                children: [
                  Text(
                    "Edit your profile",
                    style: theme.textTheme.headlineSmall?.copyWith(fontWeight: FontWeight.bold),
                  ),
                  const SizedBox(height: 16),
                    TextFormField(
                      initialValue: _data.username,
                      decoration: InputDecoration(
                        labelText: 'Username',
                        border: const OutlineInputBorder(),
                        fillColor: theme.colorScheme.surface,
                        filled: true,
                      ),
                      style: theme.textTheme.bodyMedium,
                      onChanged: (value) => _data.username = value,
                      validator: (value) =>
                          value == null || value.isEmpty ? 'Enter username' : null,
                    ),
                    const SizedBox(height: 12),
                    TextFormField(
                      initialValue: _data.description,
                      decoration: InputDecoration(
                        labelText: 'Description',
                        border: const OutlineInputBorder(),
                        fillColor: theme.colorScheme.surface,
                        filled: true,
                      ),
                      style: theme.textTheme.bodyMedium,
                      maxLines: 3,
                      onChanged: (value) => _data.description = value,
                    ),
                    const SizedBox(height: 12),
                    _buildCitySearchField(theme),
                    const SizedBox(height: 12),
                    _buildSportSearchField(theme),
                    const SizedBox(height: 24),
                    ElevatedButton(
                      onPressed: _saveProfile,
                      child: const Text("Save Changes"),
                    ),
                  ],
                ),
              ),
            ),
    );
  }
}
