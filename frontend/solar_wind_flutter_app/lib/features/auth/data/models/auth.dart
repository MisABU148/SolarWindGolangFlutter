class Auth {
  final String id;
  final String token;

  Auth({required this.id, required this.token});

  factory Auth.fromJson(Map<String, dynamic> json) {
    return Auth(
      id: json['id'].toString(),
      token: json['token'],
    );
  }
}
