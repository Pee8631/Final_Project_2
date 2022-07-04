// To parse this JSON data, do
//
//     final listUser = listUserFromJson(jsonString);

import 'package:meta/meta.dart';
import 'dart:convert';

List<ListUser> listUserFromJson(String str) => List<ListUser>.from(json.decode(str).map((x) => ListUser.fromJson(x)));

String listUserToJson(List<ListUser> data) => json.encode(List<dynamic>.from(data.map((x) => x.toJson())));

class ListUser {
    ListUser({
        required this.id,
        required this.username,
        required this.password,
        required this.edges,
    });

    int id;
    String username;
    String password;
    ListUserEdges? edges;

    factory ListUser.fromJson(Map<String, dynamic> json) => ListUser(
        id: json["id"] == null ? null : json["id"],
        username: json["username"] == null ? null : json["username"],
        password: json["password"] == null ? null : json["password"],
        edges: json["edges"] == null ? null : ListUserEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "username": username,
        "password": password,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class ListUserEdges {
    ListUserEdges({
        required this.userHasPInfo,
        required this.hasDepartment,
        required this.fromHospital,
    });

    List<UserHasPInfo>? userHasPInfo;
    FromHospital? hasDepartment;
    FromHospital? fromHospital;

    factory ListUserEdges.fromJson(Map<String, dynamic> json) => ListUserEdges(
        userHasPInfo: json["user_has_PInfo"] == null ? null : List<UserHasPInfo>.from(json["user_has_PInfo"].map((x) => UserHasPInfo.fromJson(x))),
        hasDepartment: json["has_department"] == null ? null : FromHospital.fromJson(json["has_department"]),
        fromHospital: json["from_hospital"] == null ? null : FromHospital.fromJson(json["from_hospital"]),
    );

    Map<String, dynamic> toJson() => {
        "user_has_PInfo": userHasPInfo == null ? null : List<dynamic>.from(userHasPInfo!.map((x) => x.toJson())),
        "has_department": hasDepartment == null ? null : hasDepartment!.toJson(),
        "from_hospital": fromHospital == null ? null : fromHospital!.toJson(),
    };
}

class FromHospital {
    FromHospital({
        required this.id,
        required this.name,
        required this.edges,
    });

    int id;
    String name;
    FromHospitalEdges? edges;

    factory FromHospital.fromJson(Map<String, dynamic> json) => FromHospital(
        id: json["id"] == null ? null : json["id"],
        name: json["name"] == null ? null : json["name"],
        edges: json["edges"] == null ? null : FromHospitalEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "name": name,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class FromHospitalEdges {
    FromHospitalEdges();

    factory FromHospitalEdges.fromJson(Map<String, dynamic> json) => FromHospitalEdges(
    );

    Map<String, dynamic> toJson() => {
    };
}

class UserHasPInfo {
    UserHasPInfo({
        required this.id,
        required this.idCardNumber,
        required this.firstName,
        required this.lastName,
        required this.gender,
        required this.brithDate,
        required this.bloodGroup,
        required this.address,
        required this.edges,
    });

    int id;
    String idCardNumber;
    String firstName;
    String lastName;
    int gender;
    DateTime? brithDate;
    String bloodGroup;
    String address;
    FromHospitalEdges? edges;

    factory UserHasPInfo.fromJson(Map<String, dynamic> json) => UserHasPInfo(
        id: json["id"] == null ? null : json["id"],
        idCardNumber: json["idCardNumber"] == null ? null : json["idCardNumber"],
        firstName: json["firstName"] == null ? null : json["firstName"],
        lastName: json["lastName"] == null ? null : json["lastName"],
        gender: json["gender"] == null ? null : json["gender"],
        brithDate: json["brithDate"] == null ? null : DateTime.parse(json["brithDate"]),
        bloodGroup: json["bloodGroup"] == null ? null : json["bloodGroup"],
        address: json["address"] == null ? null : json["address"],
        edges: json["edges"] == null ? null : FromHospitalEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "idCardNumber": idCardNumber,
        "firstName": firstName,
        "lastName": lastName,
        "gender": gender,
        "brithDate": brithDate == null ? null : brithDate!.toIso8601String(),
        "bloodGroup": bloodGroup,
        "address": address,
        "edges": edges == null ? null : edges!.toJson(),
    };
}
