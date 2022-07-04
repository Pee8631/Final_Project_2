// To parse this JSON data, do
//
//     final notificationApi = notificationApiFromJson(jsonString);

import 'dart:convert';

List<NotificationApi> notificationApiFromJson(String str) => List<NotificationApi>.from(json.decode(str).map((x) => NotificationApi.fromJson(x)));

String notificationApiToJson(List<NotificationApi> data) => json.encode(List<dynamic>.from(data.map((x) => x.toJson())));

class NotificationApi {
    NotificationApi({
        required this.id,
        required this.message,
        required this.createdDate,
        required this.recipientId,
        required this.senderId,
        required this.appointmentId,
        required this.edges,
    });

    int id;
    String message;
    DateTime? createdDate;
    int recipientId;
    int senderId;
    int appointmentId;
    NotificationApiEdges? edges;

    factory NotificationApi.fromJson(Map<String, dynamic> json) => NotificationApi(
        id: json["id"] == null ? null : json["id"],
        message: json["Message"] == null ? null : json["Message"],
        createdDate: json["CreatedDate"] == null ? null : DateTime.parse(json["CreatedDate"]),
        recipientId: json["RecipientId"] == null ? null : json["RecipientId"],
        senderId: json["SenderId"] == null ? null : json["SenderId"],
        appointmentId: json["AppointmentId"] == null ? null : json["AppointmentId"],
        edges: json["edges"] == null ? null : NotificationApiEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "Message": message,
        "CreatedDate": createdDate == null ? null : createdDate!.toIso8601String(),
        "RecipientId": recipientId,
        "SenderId": senderId,
        "AppointmentId": appointmentId,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class NotificationApiEdges {
    NotificationApiEdges({
        required this.userNotification,
    });

    List<UserNotification>? userNotification;

    factory NotificationApiEdges.fromJson(Map<String, dynamic> json) => NotificationApiEdges(
        userNotification: json["user_notification"] == null ? null : List<UserNotification>.from(json["user_notification"].map((x) => UserNotification.fromJson(x))),
    );

    Map<String, dynamic> toJson() => {
        "user_notification": userNotification == null ? null : List<dynamic>.from(userNotification!.map((x) => x.toJson())),
    };
}

class UserNotification {
    UserNotification({
        required this.id,
        required this.username,
        required this.password,
        required this.edges,
    });

    int id;
    String username;
    String password;
    UserNotificationEdges? edges;

    factory UserNotification.fromJson(Map<String, dynamic> json) => UserNotification(
        id: json["id"] == null ? null : json["id"],
        username: json["username"] == null ? null : json["username"],
        password: json["password"] == null ? null : json["password"],
        edges: json["edges"] == null ? null : UserNotificationEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "username": username,
        "password": password,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class UserNotificationEdges {
    UserNotificationEdges();

    factory UserNotificationEdges.fromJson(Map<String, dynamic> json) => UserNotificationEdges(
    );

    Map<String, dynamic> toJson() => {
    };
}
