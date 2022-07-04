// To parse this JSON data, do
//
//     final certification = certificationFromJson(jsonString);

import 'dart:convert';

Certification certificationFromJson(String str) => Certification.fromJson(json.decode(str));

String certificationToJson(Certification data) => json.encode(data.toJson());

class Certification {
    Certification({
        required this.code,
        required this.diloma,
        required this.dateOfIssuing,
        required this.dateOfExp,
        required this.issuer,
        required this.user,
    });

    String? code;
    String? diloma;
    DateTime? dateOfIssuing;
    DateTime? dateOfExp;
    String? issuer;
    int? user;

    factory Certification.fromJson(Map<String, dynamic> json) => Certification(
        code: json["Code"] == null ? null : json["Code"],
        diloma: json["Diloma"] == null ? null : json["Diloma"],
        dateOfIssuing: json["DateOfIssuing"] == null ? null : DateTime.parse(json["DateOfIssuing"]),
        dateOfExp: json["DateOfExp"] == null ? null : DateTime.parse(json["DateOfExp"]),
        issuer: json["Issuer"] == null ? null : json["Issuer"],
        user: json["User"] == null ? null : json["User"],
    );

    Map<String, dynamic> toJson() => {
        "Code": code == null ? null : code,
        "Diloma": diloma == null ? null : diloma,
        "DateOfIssuing": dateOfIssuing == null ? null : dateOfIssuing!.toIso8601String(),
        "DateOfExp": dateOfExp == null ? null : dateOfExp!.toIso8601String(),
        "Issuer": issuer == null ? null : issuer,
        "User": user == null ? null : user,
    };
}
