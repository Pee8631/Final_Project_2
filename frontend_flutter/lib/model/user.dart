// To parse this JSON data, do
//
//     final user = userFromJson(jsonString);

import 'dart:convert';

User userFromJson(String str) => User.fromJson(json.decode(str));

String userToJson(User data) => json.encode(data.toJson());

class User {
    User({
        required this.username,
        required this.password,
    });

    String username;
    String password;

    factory User.fromJson(Map<dynamic, dynamic> json) => User(
        username: json["Username"] == null ? null : json["Username"],
        password: json["Password"] == null ? null : json["Password"],
    );

    Map<String, dynamic> toJson() => {
        // ignore: unnecessary_null_comparison
        "Username": username == null ? null : username,
        // ignore: unnecessary_null_comparison
        "Password": password == null ? null : password,
    };
}
