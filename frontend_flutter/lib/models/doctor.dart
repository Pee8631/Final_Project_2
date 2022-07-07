// To parse this JSON data, do
//
//     final doctor = doctorFromJson(jsonString);

import 'package:meta/meta.dart';
import 'dart:convert';

Doctor doctorFromJson(String str) => Doctor.fromJson(json.decode(str));

String doctorToJson(Doctor data) => json.encode(data.toJson());

class Doctor {
    Doctor({
        required this.id,
        required this.username,
        required this.password,
        required this.edges,
    });

    int id;
    String username;
    String password;
    DoctorEdges? edges;

    factory Doctor.fromJson(Map<String, dynamic> json) => Doctor(
        id: json["id"] == null ? null : json["id"],
        username: json["username"] == null ? null : json["username"],
        password: json["password"] == null ? null : json["password"],
        edges: json["edges"] == null ? null : DoctorEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "username": username,
        "password": password,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class DoctorEdges {
    DoctorEdges({
        required this.doctorHasCertification,
        required this.userHasPInfo,
        required this.hasDepartment,
        required this.fromHospital,
    });

    List<DoctorHasCertification>? doctorHasCertification;
    List<UserHasPInfo>? userHasPInfo;
    HasDepartment? hasDepartment;
    FromHospital? fromHospital;

    factory DoctorEdges.fromJson(Map<String, dynamic> json) => DoctorEdges(
        doctorHasCertification: json["doctor_has_certification"] == null ? null : List<DoctorHasCertification>.from(json["doctor_has_certification"].map((x) => DoctorHasCertification.fromJson(x))),
        userHasPInfo: json["user_has_PInfo"] == null ? null : List<UserHasPInfo>.from(json["user_has_PInfo"].map((x) => UserHasPInfo.fromJson(x))),
        hasDepartment: json["has_department"] == null ? null : HasDepartment.fromJson(json["has_department"]),
        fromHospital: json["from_hospital"] == null ? null : FromHospital.fromJson(json["from_hospital"]),
    );

    Map<String, dynamic> toJson() => {
        "doctor_has_certification": doctorHasCertification == null ? null : List<dynamic>.from(doctorHasCertification!.map((x) => x.toJson())),
        "user_has_PInfo": userHasPInfo == null ? null : List<dynamic>.from(userHasPInfo!.map((x) => x.toJson())),
        "has_department": hasDepartment == null ? null : hasDepartment!.toJson(),
        "from_hospital": fromHospital == null ? null : fromHospital!.toJson(),
    };
}

class DoctorHasCertification {
    DoctorHasCertification({
        required this.id,
        required this.code,
        required this.diloma,
        required this.dateOfIssuing,
        required this.dateOfExp,
        required this.issuer,
        required this.edges,
    });

    int id;
    String code;
    String diloma;
    DateTime? dateOfIssuing;
    DateTime? dateOfExp;
    String issuer;
    DoctorHasCertificationEdges? edges;

    factory DoctorHasCertification.fromJson(Map<String, dynamic> json) => DoctorHasCertification(
        id: json["id"] == null ? null : json["id"],
        code: json["code"] == null ? null : json["code"],
        diloma: json["diloma"] == null ? null : json["diloma"],
        dateOfIssuing: json["dateOfIssuing"] == null ? null : DateTime.parse(json["dateOfIssuing"]),
        dateOfExp: json["dateOfExp"] == null ? null : DateTime.parse(json["dateOfExp"]),
        issuer: json["Issuer"] == null ? null : json["Issuer"],
        edges: json["edges"] == null ? null : DoctorHasCertificationEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "code": code,
        "diloma": diloma,
        "dateOfIssuing": dateOfIssuing == null ? null : dateOfIssuing!.toIso8601String(),
        "dateOfExp": dateOfExp == null ? null : dateOfExp!.toIso8601String(),
        "Issuer": issuer,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class DoctorHasCertificationEdges {
    DoctorHasCertificationEdges();

    factory DoctorHasCertificationEdges.fromJson(Map<String, dynamic> json) => DoctorHasCertificationEdges(
    );

    Map<String, dynamic> toJson() => {
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
    DoctorHasCertificationEdges? edges;

    factory FromHospital.fromJson(Map<String, dynamic> json) => FromHospital(
        id: json["id"] == null ? null : json["id"],
        name: json["name"] == null ? null : json["name"],
        edges: json["edges"] == null ? null : DoctorHasCertificationEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "name": name,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class HasDepartment {
    HasDepartment({
        required this.id,
        required this.name,
        required this.image,
        required this.edges,
    });

    int id;
    String name;
    String image;
    DoctorHasCertificationEdges? edges;

    factory HasDepartment.fromJson(Map<String, dynamic> json) => HasDepartment(
        id: json["id"] == null ? null : json["id"],
        name: json["name"] == null ? null : json["name"],
        image: json["image"] == null ? null : json["image"],
        edges: json["edges"] == null ? null : DoctorHasCertificationEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "name": name,
        "image": image,
        "edges": edges == null ? null : edges!.toJson(),
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

    int id;
    String? profile;
    String idCardNumber;
    String? prefix;
    String firstName;
    String lastName;
    int gender;
    DateTime? brithDate;
    String bloodGroup;
    String? address;
    String? about;
    DoctorHasCertificationEdges? edges;

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
        edges: json["edges"] == null ? null : DoctorHasCertificationEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "profile": profile == null ? null : profile,
        "idCardNumber": idCardNumber,
        "prefix": prefix == null ? null : prefix,
        "firstName": firstName,
        "lastName": lastName,
        "gender": gender,
        "brithDate": brithDate == null ? null : brithDate!.toIso8601String(),
        "bloodGroup": bloodGroup,
        "address": address == null ? null : address,
        "about": about == null ? null : about,
        "edges": edges == null ? null : edges!.toJson(),
    };
}
