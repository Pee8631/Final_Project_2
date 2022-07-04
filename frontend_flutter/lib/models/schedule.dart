// To parse this JSON data, do
//
//     final schedule = scheduleFromJson(jsonString);

import 'package:meta/meta.dart';
import 'dart:convert';

Schedule scheduleFromJson(String str) => Schedule.fromJson(json.decode(str));

String scheduleToJson(Schedule data) => json.encode(data.toJson());

List<Schedule> schedulesFromJson(String str) => List<Schedule>.from(json.decode(str).map((x) => Schedule.fromJson(x)));

String schedulesToJson(List<Schedule> data) => json.encode(List<dynamic>.from(data.map((x) => x.toJson())));

class Schedule {
    Schedule({
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
    DateTime startTime;
    DateTime endTime;
    Edges edges;

    factory Schedule.fromJson(Map<String, dynamic> json) => Schedule(
        id: json["id"],
        activity: json["activity"],
        detail: json["detail"],
        status: json["status"],
        startTime: DateTime.parse(json["startTime"]),
        endTime: DateTime.parse(json["endTime"]),
        edges: Edges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "activity": activity,
        "detail": detail,
        "status":status,
        "startTime": startTime.toIso8601String(),
        "endTime": endTime.toIso8601String(),
        "edges": edges.toJson(),
    };
}

class Edges {
    Edges();

    factory Edges.fromJson(Map<String, dynamic> json) => Edges(
    );

    Map<String, dynamic> toJson() => {
    };
}
