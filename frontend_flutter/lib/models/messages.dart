// To parse this JSON data, do
//
//     final messages = messagesFromJson(jsonString);

import 'dart:convert';

List<Messages> messagesFromJson(String str) => List<Messages>.from(json.decode(str).map((x) => Messages.fromJson(x)));

String messagesToJson(List<Messages> data) => json.encode(List<dynamic>.from(data.map((x) => x.toJson())));

class Messages {
    Messages({
        required this.id,
        required this.messageText,
        required this.sentDateTime,
        required this.edges,
    });

    int id;
    String messageText;
    DateTime sentDateTime;
    MessageEdges edges;

    factory Messages.fromJson(Map<String, dynamic> json) => Messages(
        id: json["id"],
        messageText: json["message_text"],
        sentDateTime: DateTime.parse(json["sent_dateTime"]),
        edges: MessageEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "message_text": messageText,
        "sent_dateTime": sentDateTime.toIso8601String(),
        "edges": edges.toJson(),
    };
}

class MessageEdges {
    MessageEdges({
        required this.whatMessagesAreInThisChat,
        required this.whoSendMessages,
    });

    WhatMessagesAreInThisChat whatMessagesAreInThisChat;
    WhoSendMessages? whoSendMessages;

    factory MessageEdges.fromJson(Map<String, dynamic> json) => MessageEdges(
        whatMessagesAreInThisChat: WhatMessagesAreInThisChat.fromJson(json["What_messages_are_in_this_chat"]),
        whoSendMessages: json["Who_send_messages"] == null ? null : WhoSendMessages.fromJson(json["Who_send_messages"]),
    );

    Map<String, dynamic> toJson() => {
        "What_messages_are_in_this_chat": whatMessagesAreInThisChat.toJson(),
        "Who_send_messages": whoSendMessages == null ? null : whoSendMessages!.toJson(),
    };
}

class WhatMessagesAreInThisChat {
    WhatMessagesAreInThisChat({
        required this.id,
        required this.chatRoomName,
        required this.isLockChat,
        required this.edges,
    });

    int id;
    String chatRoomName;
    bool isLockChat;
    WhatMessagesAreInThisChatEdges edges;

    factory WhatMessagesAreInThisChat.fromJson(Map<String, dynamic> json) => WhatMessagesAreInThisChat(
        id: json["id"] == null ? null : json["id"],
        chatRoomName: json["Chat_room_name"] == null ? null : json["Chat_room_name"],
        isLockChat: json["IsLockChat"] == null ? false : json["IsLockChat"],
        edges: WhatMessagesAreInThisChatEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "Chat_room_name": chatRoomName,
        "IsLockChat": isLockChat,
        "edges": edges.toJson(),
    };
}

class WhatMessagesAreInThisChatEdges {
    WhatMessagesAreInThisChatEdges();

    factory WhatMessagesAreInThisChatEdges.fromJson(Map<String, dynamic> json) => WhatMessagesAreInThisChatEdges(
    );

    Map<String, dynamic> toJson() => {
    };
}

class WhoSendMessages {
    WhoSendMessages({
        required this.id,
        required this.username,
        required this.password,
        required this.edges,
    });

    int id;
    String username;
    String password;
    WhatMessagesAreInThisChatEdges edges;

    factory WhoSendMessages.fromJson(Map<String, dynamic> json) => WhoSendMessages(
        id: json["id"],
        username: json["username"],
        password: json["password"],
        edges: WhatMessagesAreInThisChatEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "username": username,
        "password": password,
        "edges": edges.toJson(),
    };
}
