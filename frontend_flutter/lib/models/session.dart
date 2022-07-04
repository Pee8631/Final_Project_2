import 'loginToken.dart';

class Session {
  final String username;
  final Role role;
  final String authToken;
  final DateTime expiresAt;

  const Session({
    required this.username,
    required this.role,
    required this.authToken,
    required this.expiresAt,
  });
}
