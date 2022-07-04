// To parse this JSON data, do
//
//     final scheduleTime = scheduleTimeFromJson(jsonString);

import 'dart:convert';

ScheduleTime scheduleTimeFromJson(String str) => ScheduleTime.fromJson(json.decode(str));

String scheduleTimeToJson(ScheduleTime data) => json.encode(data.toJson());

class ScheduleTime {
    ScheduleTime({
        required this.startTime,
        required this.stopTime,
        required this.schedule,
    });

    DateTime? startTime;
    DateTime? stopTime;
    int schedule;

    factory ScheduleTime.fromJson(Map<String, dynamic> json) => ScheduleTime(
        startTime: json["StartTime"] == null ? null : DateTime.parse(json["StartTime"]),
        stopTime: json["StopTime"] == null ? null : DateTime.parse(json["StopTime"]),
        schedule: json["Schedule"] == null ? null : json["Schedule"],
    );

    Map<String, dynamic> toJson() => {
        "StartTime": startTime == null ? null : startTime!.toIso8601String(),
        "StopTime": stopTime == null ? null : stopTime!.toIso8601String(),
        "Schedule": schedule,
    };
}
