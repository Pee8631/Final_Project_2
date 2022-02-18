// To parse this JSON data, do
//
//     final user = userFromJson(jsonString);

import 'package:meta/meta.dart';
import 'dart:convert';

User userFromJson(String str) => User.fromJson(json.decode(str));

String userToJson(User data) => json.encode(data.toJson());

class User {
    User({
        required this.username,
        required this.password,
        required this.department,
        required this.hospital,
    });

    String username;
    String password;
    int department;
    int hospital;

    factory User.fromJson(Map<String, dynamic> json) => User(
        username: json["Username"] == null ? null : json["Username"],
        password: json["Password"] == null ? null : json["Password"],
        department: json["Department"] == null ? null : json["Department"],
        hospital: json["Hospital"] == null ? null : json["Hospital"],
    );

    Map<String, dynamic> toJson() => {
        "Username": username == null ? null : username,
        "Password": password == null ? null : password,
        "Department": department == null ? null : department,
        "Hospital": hospital == null ? null : hospital,
    };
}

