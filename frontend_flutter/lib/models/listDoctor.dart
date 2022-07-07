// To parse this JSON data, do
//
//     final listDoctor = listDoctorFromJson(jsonString);

import 'package:meta/meta.dart';
import 'dart:convert';

List<ListDoctor> listDoctorFromJson(String str) => List<ListDoctor>.from(json.decode(str).map((x) => ListDoctor.fromJson(x)));

String listDoctorToJson(List<ListDoctor> data) => json.encode(List<dynamic>.from(data.map((x) => x.toJson())));

class ListDoctor {
    ListDoctor({
        required this.id,
        required this.username,
        required this.password,
        required this.edges,
    });

    int? id;
    String? username;
    String? password;
    ListDoctorEdges? edges;

    factory ListDoctor.fromJson(Map<String, dynamic> json) => ListDoctor(
        id: json["id"] == null ? null : json["id"],
        username: json["username"] == null ? null : json["username"],
        password: json["password"] == null ? null : json["password"],
        edges: json["edges"] == null ? null : ListDoctorEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id == null ? null : id,
        "username": username == null ? null : username,
        "password": password == null ? null : password,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class ListDoctorEdges {
    ListDoctorEdges({
        required this.userHasPInfo,
        required this.hasDepartment,
        required this.fromHospital,
    });

    List<UserHasPInfo>? userHasPInfo;
    FromHospital? hasDepartment;
    FromHospital? fromHospital;

    factory ListDoctorEdges.fromJson(Map<String, dynamic> json) => ListDoctorEdges(
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
        required this.image,
    });

    int? id;
    String? name;
    FromHospitalEdges? edges;
    String? image;

    factory FromHospital.fromJson(Map<String, dynamic> json) => FromHospital(
        id: json["id"] == null ? null : json["id"],
        name: json["name"] == null ? null : json["name"],
        edges: json["edges"] == null ? null : FromHospitalEdges.fromJson(json["edges"]),
        image: json["image"] == null ? null : json["image"],
    );

    Map<String, dynamic> toJson() => {
        "id": id == null ? null : id,
        "name": name == null ? null : name,
        "edges": edges == null ? null : edges!.toJson(),
        "image": image == null ? null : image,
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
        required this.profile,
        required this.idCardNumber,
        required this.prefix,
        required this.firstName,
        required this.lastName,
        required this.gender,
        required this.brithDate,
        required this.bloodGroup,
        required this.address,
        required this.about,
        required this.edges,
    });

    int? id;
    String? profile;
    String? idCardNumber;
    String? prefix;
    String? firstName;
    String? lastName;
    int? gender;
    DateTime? brithDate;
    String? bloodGroup;
    String? address;
    String? about;
    FromHospitalEdges? edges;

    factory UserHasPInfo.fromJson(Map<String, dynamic> json) => UserHasPInfo(
        id: json["id"] == null ? null : json["id"],
        profile: json["profile"] == null ? null : json["profile"],
        idCardNumber: json["idCardNumber"] == null ? null : json["idCardNumber"],
        prefix: json["prefix"] == null ? null : json["prefix"],
        firstName: json["firstName"] == null ? null : json["firstName"],
        lastName: json["lastName"] == null ? null : json["lastName"],
        gender: json["gender"] == null ? null : json["gender"],
        brithDate: json["brithDate"] == null ? null : DateTime.parse(json["brithDate"]),
        bloodGroup: json["bloodGroup"] == null ? null : json["bloodGroup"],
        address: json["address"] == null ? null : json["address"],
        about: json["about"] == null ? null : json["about"],
        edges: json["edges"] == null ? null : FromHospitalEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id == null ? null : id,
        "profile": profile == null ? null : profile,
        "idCardNumber": idCardNumber == null ? null : idCardNumber,
        "prefix": prefix == null ? null : prefix,
        "firstName": firstName == null ? null : firstName,
        "lastName": lastName == null ? null : lastName,
        "gender": gender == null ? null : gender,
        "brithDate": brithDate == null ? null : brithDate!.toIso8601String(),
        "bloodGroup": bloodGroup == null ? null : bloodGroup,
        "address": address == null ? null : address,
        "about": about == null ? null : about,
        "edges": edges == null ? null : edges!.toJson(),
    };
}
