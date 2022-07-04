// To parse this JSON data, do
//
//     final doctorSchedule = doctorScheduleFromJson(jsonString);

import 'package:meta/meta.dart';
import 'dart:convert';

DoctorSchedule doctorScheduleFromJson(String str) => DoctorSchedule.fromJson(json.decode(str));

String doctorScheduleToJson(DoctorSchedule data) => json.encode(data.toJson());

class DoctorSchedule {
    DoctorSchedule({
        required this.id,
        required this.username,
        required this.password,
        required this.edges,
    });

    int id;
    String username;
    String password;
    DoctorScheduleEdges? edges;

    factory DoctorSchedule.fromJson(Map<String, dynamic> json) => DoctorSchedule(
        id: json["id"] == null ? null : json["id"],
        username: json["username"] == null ? null : json["username"],
        password: json["password"] == null ? null : json["password"],
        edges: json["edges"] == null ? null : DoctorScheduleEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "username": username,
        "password": password,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class DoctorScheduleEdges {
    DoctorScheduleEdges({
        required this.doctorSchedule,
    });

    List<DoctorScheduleElement>? doctorSchedule;

    factory DoctorScheduleEdges.fromJson(Map<String, dynamic> json) => DoctorScheduleEdges(
        doctorSchedule: json["doctor_schedule"] == null ? null : List<DoctorScheduleElement>.from(json["doctor_schedule"].map((x) => DoctorScheduleElement.fromJson(x))),
    );

    Map<String, dynamic> toJson() => {
        "doctor_schedule": doctorSchedule == null ? null : List<dynamic>.from(doctorSchedule!.map((x) => x.toJson())),
    };
}

class DoctorScheduleElement {
    DoctorScheduleElement({
        required this.id,
        required this.activity,
        required this.detail,
        required this.status,
        required this.startTime,
        required this.endTime,
        required this.edges,
    });

    int id;
    String activity;
    String? detail;
    String status;
    DateTime? startTime;
    DateTime? endTime;
    DoctorScheduleEdgesClass? edges;

    factory DoctorScheduleElement.fromJson(Map<String, dynamic> json) => DoctorScheduleElement(
        id: json["id"] == null ? null : json["id"],
        activity: json["activity"] == null ? null : json["activity"],
        detail: json["detail"] == null ? null : json["detail"],
        status: json["status"] == null ? null : json["status"],
        startTime: json["startTime"] == null ? null : DateTime.parse(json["startTime"]),
        endTime: json["endTime"] == null ? null : DateTime.parse(json["endTime"]),
        edges: json["edges"] == null ? null : DoctorScheduleEdgesClass.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "activity": activity,
        "detail": detail,
        "status": status,
        "startTime": startTime == null ? null : startTime!.toIso8601String(),
        "endTime": endTime == null ? null : endTime!.toIso8601String(),
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class DoctorScheduleEdgesClass {
    DoctorScheduleEdgesClass();

    factory DoctorScheduleEdgesClass.fromJson(Map<String, dynamic> json) => DoctorScheduleEdgesClass(
    );

    Map<String, dynamic> toJson() => {
    };
}
