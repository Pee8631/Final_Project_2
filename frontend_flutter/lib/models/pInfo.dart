// To parse this JSON data, do
//
//     final pInfo = pInfoFromJson(jsonString);

import 'package:meta/meta.dart';
import 'dart:convert';

PInfo pInfoFromJson(String str) => PInfo.fromJson(json.decode(str));

String pInfoToJson(PInfo data) => json.encode(data.toJson());

class PInfo {
    PInfo({
        required this.about,
        required this.address,
        required this.bloodGroup,
        required this.brithDate,
        required this.firstName,
        required this.gender,
        required this.id,
        required this.idCardNumber,
        required this.lastName,
        required this.prefix,
        required this.profile,
        required this.user,
    });

    String? about;
    String? address;
    String bloodGroup;
    DateTime? brithDate;
    String firstName;
    int gender;
    int id;
    String idCardNumber;
    String lastName;
    String? prefix;
    String? profile;
    int user;

    factory PInfo.fromJson(Map<String, dynamic> json) => PInfo(
        about: json["About"] == null ? null : json["About"],
        address: json["Address"] == null ? null : json["Address"],
        bloodGroup: json["BloodGroup"] == null ? null : json["BloodGroup"],
        brithDate: json["BrithDate"] == null ? null : DateTime.parse(json["BrithDate"]),
        firstName: json["FirstName"] == null ? null : json["FirstName"],
        gender: json["Gender"] == null ? null : json["Gender"],
        id: json["Id"] == null ? null : json["Id"],
        idCardNumber: json["IdCardNumber"] == null ? null : json["IdCardNumber"],
        lastName: json["LastName"] == null ? null : json["LastName"],
        prefix: json["Prefix"] == null ? null : json["Prefix"],
        profile: json["Profile"] == null ? null : json["Profile"],
        user: json["User"] == null ? null : json["User"],
    );

    Map<String, dynamic> toJson() => {
        "About": about == null ? null : about,
        "Address": address == null ? null : address,
        "BloodGroup": bloodGroup,
        "BrithDate": brithDate == null ? null : brithDate!.toIso8601String(),
        "FirstName": firstName,
        "Gender": gender,
        "Id": id,
        "IdCardNumber": idCardNumber,
        "LastName": lastName,
        "Prefix": prefix == null ? null : prefix,
        "Profile": profile == null ? null : profile,
        "User": user,
    };
}
