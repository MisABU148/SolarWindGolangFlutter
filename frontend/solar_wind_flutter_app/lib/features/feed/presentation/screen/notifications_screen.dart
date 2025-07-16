import 'package:flutter/material.dart';
import 'package:solar_wind_flutter_app/features/feed/data/models/user.dart';
import 'package:solar_wind_flutter_app/features/feed/data/services/notifications_service.dart';
import 'package:solar_wind_flutter_app/features/feed/presentation/screen/profile_screen.dart';


class NotificationsScreen extends StatefulWidget {
  const NotificationsScreen({super.key});

  @override
  State<NotificationsScreen> createState() => _NotificationsScreenState();
}

class _NotificationsScreenState extends State<NotificationsScreen> {
  final _service = LikeNotificationService();
  bool _isLoading = true;
  List<User> _likedMeUsers = [];

  @override
  void initState() {
    super.initState();
    _fetchLikedUsers();
  }

  Future<void> _fetchLikedUsers() async {
    try {
      final users = await _service.getLikedMeUsers();
      setState(() => _likedMeUsers = users);
    } catch (e) {
      print('Error: $e');
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Failed to load notifications')),
      );
    } finally {
      setState(() => _isLoading = false);
    }
  }

  Widget _buildUserCard(User user) {
    return Card(
      margin: const EdgeInsets.symmetric(vertical: 8),
      child: ListTile(
        title: Text(user.username),
        subtitle: Text(user.description),
        trailing: const Icon(Icons.favorite, color: Colors.redAccent),
          onTap: () {
            Navigator.push(
                context,
                MaterialPageRoute(
                builder: (_) => UserProfileScreen(userId: user.id),
                ),
            );
        },
      ),
    );
  }

  @override
Widget build(BuildContext context) {
  final theme = Theme.of(context);

  return Scaffold(
    appBar: AppBar(
      title: const Text('Notifications'),
      backgroundColor: theme.colorScheme.primary,
      foregroundColor: theme.colorScheme.onPrimary,
    ),
    body: _isLoading
        ? Center(
            child: CircularProgressIndicator(color: theme.colorScheme.primary),
          )
        : _likedMeUsers.isEmpty
            ? Center(
                child: Text(
                  "No one liked you yet 🥲",
                  style: theme.textTheme.bodyMedium?.copyWith(color: theme.colorScheme.onBackground),
                ),
              )
            : ListView.builder(
                padding: const EdgeInsets.all(16),
                itemCount: _likedMeUsers.length,
                itemBuilder: (_, index) {
                  final user = _likedMeUsers[index];
                  return Card(
                    margin: const EdgeInsets.symmetric(vertical: 8),
                    color: theme.colorScheme.surface,
                    child: ListTile(
                      title: Text(
                        user.username,
                        style: theme.textTheme.titleMedium?.copyWith(color: theme.colorScheme.onSurface),
                      ),
                      subtitle: Text(
                        user.description,
                        style: theme.textTheme.bodyMedium?.copyWith(color: theme.colorScheme.onSurfaceVariant),
                      ),
                      trailing: Icon(Icons.favorite, color: theme.colorScheme.error),
                      onTap: () {
                        Navigator.push(
                          context,
                          MaterialPageRoute(
                            builder: (_) => UserProfileScreen(userId: user.id),
                          ),
                        );
                      },
                    ),
                  );
                },
              ),
  );
}

}
