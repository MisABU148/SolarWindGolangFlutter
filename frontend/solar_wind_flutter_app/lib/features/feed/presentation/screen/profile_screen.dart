import 'package:flutter/material.dart';
import 'package:solar_wind_flutter_app/features/feed/data/services/profile_service.dart';
import 'package:solar_wind_flutter_app/features/feed/data/services/like_service.dart';
import 'package:solar_wind_flutter_app/features/feed/data/models/user.dart';

class UserProfileScreen extends StatefulWidget {
  final int userId;

  const UserProfileScreen({super.key, required this.userId});

  @override
  State<UserProfileScreen> createState() => _UserProfileScreenState();
}

class _UserProfileScreenState extends State<UserProfileScreen> {
  final _profileService = ProfileService();
  final _likeService = LikeService();

  bool _isLoading = true;
  User? _user;
  bool _liked = false;

  @override
  void initState() {
    super.initState();
    _loadUser();
  }

  Future<void> _loadUser() async {
    try {
      final user = await _profileService.fetchUser(widget.userId);
      setState(() => _user = user);
    } catch (e) {
      print('Error loading profile: $e');
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Failed to load profile: $e')),
      );
    } finally {
      setState(() => _isLoading = false);
    }
  }

  Future<void> _like() async {
    try {
      await _likeService.likeUser(widget.userId);
      setState(() => _liked = true);
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('User liked!')),
      );
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Like failed: $e')),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    if (_isLoading) {
      return const Scaffold(body: Center(child: CircularProgressIndicator()));
    }

    if (_user == null) {
      return const Scaffold(body: Center(child: Text('User not found')));
    }

    return Scaffold(
      appBar: AppBar(
        title: Text(_user!.username),
        actions: [
          Padding(
            padding: const EdgeInsets.only(right: 12),
            child: _liked
                ? const Icon(Icons.favorite, color: Colors.red, size: 36)
                : IconButton(
                    icon: const Icon(Icons.favorite_border, size: 36),
                    onPressed: _like,
                    tooltip: 'Like',
                    color: Colors.redAccent,
                  ),
          )
        ],
      ),
      body: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              _user!.description,
              style: const TextStyle(fontSize: 20, fontWeight: FontWeight.w500),
            ),
            const SizedBox(height: 20),
            Text(
              'City: ${_user!.cityName}',
              style: const TextStyle(fontSize: 18, fontWeight: FontWeight.w400),
            ),
            const SizedBox(height: 20),
            Wrap(
              spacing: 8,
              children: _user!.sportName
                  .map((sport) => Chip(
                        label: Text(
                          sport,
                          style: const TextStyle(fontSize: 16),
                        ),
                      ))
                  .toList(),
            ),
          ],
        ),
      ),
    );
  }
}
