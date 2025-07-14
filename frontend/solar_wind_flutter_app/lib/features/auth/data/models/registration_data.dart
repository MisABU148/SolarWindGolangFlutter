// lib/features/auth/data/models/registration_data.dart
import '../models/city.dart';
import '../models/sport.dart';

class RegistrationData {
  String? username;
  String? description;
  int? cityId;
  List<int> sportId = [];
  List<int> days = [1];
  String? gender; 
  String? preferredGender;
  DateTime? birthDate;

  Map<String, dynamic> toJson() => {
    "username": username,
    "firstName": "firstName",
    "lastName": "lastName",
    "description": description,
    "age": birthDate != null
        ? '${birthDate!.year.toString().padLeft(4, '0')}-${birthDate!.month.toString().padLeft(2, '0')}-${birthDate!.day.toString().padLeft(2, '0')}'
        : null,
    "gender": gender,
    "preferredGender": preferredGender,
    "cityId": cityId,
    "preferredGymTime": days,
    "sportId": sportId
  };
}
