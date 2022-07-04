// To parse this JSON data, do
//
//     final user = userFromJson(jsonString);

import 'dart:convert';

import 'package:frontend_flutter/models/certification.dart';

User userFromJson(String str) => User.fromJson(json.decode(str));

String userToJson(User data) => json.encode(data.toJson());

class User {
  User(
      {required this.username,
      required this.password,
      required this.department,
      required this.hospital,
      required this.roleId,
      required this.certification});

  String? username;
  String? password;
  int? department;
  int? hospital;
  int? roleId;
  Certification? certification;

  factory User.fromJson(Map<String, dynamic> json) => User(
        username: json["Username"] == null ? null : json["Username"],
        password: json["Password"] == null ? null : json["Password"],
        department: json["Department"] == null ? null : json["Department"],
        hospital: json["Hospital"] == null ? null : json["Hospital"],
        roleId: json["RoleId"] == null ? null : json["RoleId"],
        certification: json["Certification"] == null ? null : certificationFromJson(json["Certification"]),
      );

  Map<String, dynamic> toJson() => {
        "Username": username == null ? null : username,
        "Password": password == null ? null : password,
        "Department": department == null ? null : department,
        "Hospital": hospital == null ? null : hospital,
        "RoleId": roleId == null ? null : roleId,
        "Certification": certification == certificationToJson(certification!),
      };
}
