// To parse this JSON data, do
//
//     final errorMsg = errorMsgFromJson(jsonString);

import 'dart:convert';

ErrorMsg errorMsgFromJson(String str) => ErrorMsg.fromJson(json.decode(str));

String errorMsgToJson(ErrorMsg data) => json.encode(data.toJson());

class ErrorMsg {
    ErrorMsg({
        required this.error,
        required this.status,
    });

    String? error;
    bool? status;

    factory ErrorMsg.fromJson(Map<String, dynamic> json) => ErrorMsg(
        error: json["error"] == null ? null : json["error"],
        status: json["status"] == null ? null : json["status"],
    );

    Map<String, dynamic> toJson() => {
        "error": error == null ? null : error,
        "status": status == null ? null : status,
    };
}
