import 'package:flutter/material.dart';
import 'package:solar_wind_flutter_app/features/feed/data/models/user.dart';
import 'package:solar_wind_flutter_app/features/feed/data/services/feed_service.dart';
import 'profile_screen.dart';
import 'my_profile_screen.dart';

class UserFeedScreen extends StatefulWidget {
  const UserFeedScreen({super.key});

  @override
  State<UserFeedScreen> createState() => _UserFeedScreenState();
}

class _UserFeedScreenState extends State<UserFeedScreen> {
  final FeedService _feedService = FeedService();
  List<User> _users = [];
  bool _isLoading = true;

  @override
  void initState() {
    super.initState();
    _fetchFeed();
  }

  Future<void> _fetchFeed() async {
    try {
      final users = await _feedService.fetchFeed();
      setState(() {
        _users = users;
      });
    } catch (e) {
      print('Error fetching feed: $e');
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Failed to load feed: $e')),
      );
    } finally {
      setState(() => _isLoading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        children: [
          Container(
            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 32),
            width: double.infinity,
            decoration: const BoxDecoration(
              gradient: LinearGradient(
                colors: [Color(0xFFF34BD5), Color(0xFFFBDA55)],
                begin: Alignment.topLeft,
                end: Alignment.bottomRight,
              ),
            ),
            child: Row(
              children: [
                Align(
                  alignment: Alignment.centerLeft,
                  child: Padding(
                    padding: const EdgeInsets.only(top: 20),
                    child: IconButton(
                      iconSize: 32,
                      icon: const Icon(Icons.person, color: Colors.white),
                      onPressed: () {
                        final userId = 637451540;

                        Navigator.push(
                          context,
                          MaterialPageRoute(
                            builder: (_) => MyProfileScreen(userId: userId),
                          ),
                        );
                      },
                    ),
                  ),
                ),
              ],
            ),
          ),
          const SizedBox(height: 8),
          Expanded(
            child: _isLoading
                ? const Center(child: CircularProgressIndicator())
                : ListView.builder(
                    itemCount: _users.length,
                    itemBuilder: (context, index) {
                      final user = _users[index];
                      return Padding(
                        padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
                        child: GestureDetector(
                          onTap: () {
                            Navigator.push(
                              context,
                              MaterialPageRoute(
                                builder: (_) => UserProfileScreen(userId: user.id),
                              ),
                            );
                          },
                          child: Card(
                            elevation: 2,
                            shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
                            child: Padding(
                              padding: const EdgeInsets.all(16),
                              child: Column(
                                crossAxisAlignment: CrossAxisAlignment.start,
                                children: [
                                  Text(user.username, style: const TextStyle(fontSize: 18, fontWeight: FontWeight.bold)),
                                  const SizedBox(height: 8),
                                  Text(user.description),
                                  const SizedBox(height: 8),
                                  Wrap(
                                    spacing: 8,
                                    children: user.sportName.map((sport) => Chip(label: Text(sport))).toList(),
                                  ),
                                ],
                              ),
                            ),
                          ),
                        ),

                      );
                    },
                  ),
          ),
        ],
      ),
    );
  }
}
