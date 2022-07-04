// To parse this JSON data, do
//
//     final appointment = appointmentFromJson(jsonString);

import 'dart:convert';

AppointmentApi appointmentFromJson(String str) => AppointmentApi.fromJson(json.decode(str));

String appointmentToJson(AppointmentApi data) => json.encode(data.toJson());

List<AppointmentApi> appointmentsFromJson(String str) => List<AppointmentApi>.from(json.decode(str).map((x) => AppointmentApi.fromJson(x)));

String appointmentsToJson(List<AppointmentApi> data) => json.encode(List<dynamic>.from(data.map((x) => x.toJson())));

class AppointmentApi {
    AppointmentApi({
        required this.id,
        required this.reasonForAppointment,
        required this.detail,
        required this.startTime,
        required this.endTime,
        required this.status,
        required this.doctorId,
        required this.userId,
        required this.edges,
    });

    int id;
    String? reasonForAppointment;
    String? detail;
    DateTime? startTime;
    DateTime? endTime;
    String status;
    int doctorId;
    int? userId;
    AppointmentApiEdges? edges;

    factory AppointmentApi.fromJson(Map<String, dynamic> json) => AppointmentApi(
        id: json["id"] == null ? null : json["id"],
        reasonForAppointment: json["reasonForAppointment"] == null ? null : json["reasonForAppointment"],
        detail: json["detail"] == null ? null : json["detail"],
        startTime: json["startTime"] == null ? null : DateTime.parse(json["startTime"]),
        endTime: json["endTime"] == null ? null : DateTime.parse(json["endTime"]),
        status: json["status"] == null ? null : json["status"],
        doctorId: json["DoctorId"] == null ? null : json["DoctorId"],
        userId: json["UserId"] == null ? null : json["UserId"],
        edges: json["edges"] == null ? null : AppointmentApiEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "reasonForAppointment": reasonForAppointment == null ? null : reasonForAppointment,
        "detail": detail == null ? null : detail,
        "startTime": startTime == null ? null : startTime!.toIso8601String(),
        "endTime": endTime == null ? null : endTime!.toIso8601String(),
        "status": status,
        "DoctorId": doctorId,
        "UserId": userId == null ? null : userId,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class AppointmentApiEdges {
    AppointmentApiEdges({
        required this.appointmentChat,
    });

    AppointmentChat? appointmentChat;

    factory AppointmentApiEdges.fromJson(Map<String, dynamic> json) => AppointmentApiEdges(
        appointmentChat: json["appointment_chat"] == null ? null : AppointmentChat.fromJson(json["appointment_chat"]),
    );

    Map<String, dynamic> toJson() => {
        "appointment_chat": appointmentChat == null ? null : appointmentChat!.toJson(),
    };
}

class AppointmentChat {
    AppointmentChat({
        required this.id,
        required this.chatRoomName,
        required this.edges,
    });

    int id;
    String chatRoomName;
    AppointmentChatEdges? edges;

    factory AppointmentChat.fromJson(Map<String, dynamic> json) => AppointmentChat(
        id: json["id"] == null ? null : json["id"],
        chatRoomName: json["Chat_room_name"] == null ? null : json["Chat_room_name"],
        edges: json["edges"] == null ? null : AppointmentChatEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "Chat_room_name": chatRoomName,
        "edges": edges,
    };
}

class AppointmentChatEdges {
    AppointmentChatEdges();

    factory AppointmentChatEdges.fromJson(Map<String, dynamic> json) => AppointmentChatEdges(
    );
}