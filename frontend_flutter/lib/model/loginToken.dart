// To parse this JSON data, do
//
//     final loginToken = loginTokenFromJson(jsonString);

import 'dart:convert';

LoginToken loginTokenFromJson(String str) => LoginToken.fromJson(json.decode(str));

String loginTokenToJson(LoginToken data) => json.encode(data.toJson());

class LoginToken {
    LoginToken({
        required this.authToken,
        required this.expiresAt,
        required this.generatedAt,
        required this.user,
    });

    String? authToken;
    DateTime? expiresAt;
    DateTime? generatedAt;
    int? user;

    factory LoginToken.fromJson(Map<String, dynamic> json) => LoginToken(
        authToken: json["AuthToken"] == null ? null : json["AuthToken"],
        expiresAt: json["ExpiresAt"] == null ? null : DateTime.parse(json["ExpiresAt"]),
        generatedAt: json["GeneratedAt"] == null ? null : DateTime.parse(json["GeneratedAt"]),
        user: json["User"] == null ? null : json["User"],
    );

    Map<String, dynamic> toJson() => {
        "AuthToken": authToken == null ? null : authToken,
        "ExpiresAt": expiresAt == null ? null : expiresAt!.toIso8601String(),
        "GeneratedAt": generatedAt == null ? null : generatedAt!.toIso8601String(),
        "User": user == null ? null : user,
    };
}
