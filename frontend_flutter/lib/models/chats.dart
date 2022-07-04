// To parse this JSON data, do
//
//     final chatss = chatsFromJson(jsonString);


import 'dart:convert';

Chats chatFromJson(String str) => Chats.fromJson(json.decode(str));

String chatToJson(Chats data) => json.encode(data.toJson());

List<Chats> chatsFromJson(String str) => List<Chats>.from(json.decode(str).map((x) => Chats.fromJson(x)));

String chatsToJson(List<Chats> data) => json.encode(List<dynamic>.from(data.map((x) => x.toJson())));


class Chats {
    Chats({
        required this.id,
        required this.chatRoomName,
        required this.isLockChat,
        required this.edges,
    });

    int id;
    String chatRoomName;
    bool isLockChat;
    ChatsEdges? edges;

    factory Chats.fromJson(Map<String, dynamic> json) => Chats(
        id: json["id"] == null ? null : json["id"],
        chatRoomName: json["Chat_room_name"] == null ? null : json["Chat_room_name"],
        isLockChat: json["IsLockChat"] == null ? false : json["IsLockChat"],
        edges: json["edges"] == null ? null : ChatsEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "Chat_room_name": chatRoomName,
        "IsLockChat": isLockChat,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class ChatsEdges {
    ChatsEdges({
        required this.chatUser,
        required this.chatMessage,
    });

    List<ChatUser>? chatUser;
    List<ChatMessage>? chatMessage;

    factory ChatsEdges.fromJson(Map<String, dynamic> json) => ChatsEdges(
        chatUser: json["chat_user"] == null ? null : List<ChatUser>.from(json["chat_user"].map((x) => ChatUser.fromJson(x))),
        chatMessage: json["chat_message"] == null ? null : List<ChatMessage>.from(json["chat_message"].map((x) => ChatMessage.fromJson(x))),
    );

    Map<String, dynamic> toJson() => {
        "chat_user": chatUser == null ? null : List<dynamic>.from(chatUser!.map((x) => x.toJson())),
        "chat_message": chatMessage == null ? null : List<dynamic>.from(chatMessage!.map((x) => x.toJson())),
    };
}

class ChatMessage {
    ChatMessage({
        required this.id,
        required this.messageText,
        required this.sentDateTime,
        required this.edges,
    });

    int id;
    String messageText;
    DateTime? sentDateTime;
    ChatMessageEdges? edges;

    factory ChatMessage.fromJson(Map<String, dynamic> json) => ChatMessage(
        id: json["id"] == null ? null : json["id"],
        messageText: json["message_text"] == null ? null : json["message_text"],
        sentDateTime: json["sent_dateTime"] == null ? null : DateTime.parse(json["sent_dateTime"]),
        edges: json["edges"] == null ? null : ChatMessageEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "message_text": messageText,
        "sent_dateTime": sentDateTime,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class ChatMessageEdges {
    ChatMessageEdges();

    factory ChatMessageEdges.fromJson(Map<String, dynamic> json) => ChatMessageEdges(
    );

    Map<String, dynamic> toJson() => {
    };
}

class ChatUser {
    ChatUser({
        required this.id,
        required this.username,
        required this.password,
        required this.edges,
    });

    int id;
    String username;
    String password;
    ChatMessageEdges? edges;

    factory ChatUser.fromJson(Map<String, dynamic> json) => ChatUser(
        id: json["id"] == null ? null : json["id"],
        username: json["username"] == null ? null : json["username"],
        password: json["password"] == null ? null : json["password"],
        edges: json["edges"] == null ? null : ChatMessageEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "username": username,
        "password": password,
        "edges": edges == null ? null : edges!.toJson(),
    };
}



