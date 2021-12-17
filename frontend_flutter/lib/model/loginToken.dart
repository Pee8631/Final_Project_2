// To parse this JSON data, do
//
//     final loginToken = loginTokenFromJson(jsonString);

import 'dart:convert';

LoginToken loginTokenFromJson(String str) =>
    LoginToken.fromJson(json.decode(str));

String loginTokenToJson(LoginToken data) => json.encode(data.toJson());

class LoginToken {
  LoginToken({
    required this.authorized,
    required this.exp,
    required this.userId,
  });

  bool authorized;
  int exp;
  String userId;

  factory LoginToken.fromJson(Map<String, dynamic> json) => LoginToken(
        authorized: json["authorized"] == null ? null : json["authorized"],
        exp: json["exp"] == null ? null : json["exp"],
        userId: json["user_id"] == null ? null : json["user_id"],
      );

  Map<String, dynamic> toJson() => {
        // ignore: unnecessary_null_comparison
        "authorized": authorized == null ? null : authorized,
        // ignore: unnecessary_null_comparison
        "exp": exp == null ? null : exp,
        // ignore: unnecessary_null_comparison
        "user_id": userId == null ? null : userId,
      };
}
