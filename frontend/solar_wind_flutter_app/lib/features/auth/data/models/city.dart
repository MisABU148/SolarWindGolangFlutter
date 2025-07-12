class City {
  final String id;
  final String name;

  City({required this.id, required this.name});

  factory City.fromJson(Map<String, dynamic> json) => City(
        id: json['id'].toString(),
        name: json['cityName'],
      );
}
