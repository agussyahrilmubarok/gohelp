import 'package:shelf_dart_api/users/dtos/user_dto.dart';

class UserAndPassword {
  final UserDto user;
  final String password;

  UserAndPassword({required this.user, required this.password});
}
