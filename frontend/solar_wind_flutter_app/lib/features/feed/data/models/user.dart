class User {
  final int id;
  final String username;
  final String description;
  final String cityName;
  final List<int> preferredGymTime;
  final List<String> sportName;

  User({
    required this.id,
    required this.username,
    required this.description,
    required this.cityName,
    required this.preferredGymTime,
    required this.sportName,
  });

  factory User.fromJson(Map<String, dynamic> json) => User(
        id: json['id'],
        username: json['username'] ?? '',
        description: json['description'] ?? '',
        cityName: json['cityName'] ?? '',
        preferredGymTime: List<int>.from(json['preferredGymTime'] ?? []),
        sportName: List<String>.from(json['sportName'] ?? []),
      );
}
