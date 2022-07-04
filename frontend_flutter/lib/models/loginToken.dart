// To parse this JSON data, do
//
//     final loginToken = loginTokenFromJson(jsonString);

import 'dart:convert';

LoginToken loginTokenFromJson(String str) => LoginToken.fromJson(json.decode(str));

String loginTokenToJson(LoginToken data) => json.encode(data.toJson());

class LoginToken {
    LoginToken({
        required this.department,
        required this.hospital,
        required this.password,
        required this.role,
        required this.token,
        required this.username,
    });

    Department? department;
    Hospital? hospital;
    String password;
    List<Role>? role;
    Token? token;
    String username;

    factory LoginToken.fromJson(Map<String, dynamic> json) => LoginToken(
        department: json["Department"] == null ? null : Department.fromJson(json["Department"]),
        hospital: json["Hospital"] == null ? null : Hospital.fromJson(json["Hospital"]),
        password: json["Password"] == null ? null : json["Password"],
        role: json["Role"] == null ? null : List<Role>.from(json["Role"].map((x) => Role.fromJson(x))),
        token: json["Token"] == null ? null : Token.fromJson(json["Token"]),
        username: json["Username"] == null ? null : json["Username"],
    );

    Map<String, dynamic> toJson() => {
        "Department": department == null ? null : department!.toJson(),
        "Hospital": hospital == null ? null : hospital!.toJson(),
        "Password": password == password,
        "Role": role == null ? null : List<dynamic>.from(role!.map((x) => x.toJson())),
        "Token": token == null ? null : token!.toJson(),
        "Username": username == username,
    };
}

class Department {
    Department({
        required this.id,
        required this.name,
        required this.edges,
    });

    int id;
    String name;
    Edges? edges;

    factory Department.fromJson(Map<String, dynamic> json) => Department(
        id: json["id"] == null ? null : json["id"],
        name: json["name"] == null ? null : json["name"],
        edges: json["edges"] == null ? null : Edges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id == null ? null : id,
        "name": name == null ? null : name,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class Hospital {
    Hospital({
        required this.id,
        required this.name,
        required this.edges,
    });

    int id;
    String name;
    Edges? edges;

    factory Hospital.fromJson(Map<String, dynamic> json) => Hospital(
        id: json["id"] == null ? null : json["id"],
        name: json["name"] == null ? null : json["name"],
        edges: json["edges"] == null ? null : Edges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id == null ? null : id,
        "name": name == null ? null : name,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class Role {
    Role({
        required this.id,
        required this.name,
        required this.edges,
    });

    int id;
    String name;
    Edges? edges;

    factory Role.fromJson(Map<String, dynamic> json) => Role(
        id: json["id"] == null ? null : json["id"],
        name: json["name"] == null ? null : json["name"],
        edges: json["edges"] == null ? null : Edges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id == null ? null : id,
        "name": name == null ? null : name,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class Edges {
    Edges();

    factory Edges.fromJson(Map<String, dynamic> json) => Edges(
    );

    Map<String, dynamic> toJson() => {
    };
}

class Token {
    Token({
        required this.authToken,
        required this.expiresAt,
        required this.generatedAt,
        required this.user,
    });

    String authToken;
    DateTime? expiresAt;
    DateTime? generatedAt;
    int user;

    factory Token.fromJson(Map<String, dynamic> json) => Token(
        authToken: json["AuthToken"] == null ? null : json["AuthToken"],
        expiresAt: json["ExpiresAt"] == null ? null : DateTime.parse(json["ExpiresAt"]),
        generatedAt: json["GeneratedAt"] == null ? null : DateTime.parse(json["GeneratedAt"]),
        user: json["User"] == null ? null : json["User"],
    );

    Map<String, dynamic> toJson() => {
        "AuthToken": authToken == authToken,
        "ExpiresAt": expiresAt == null ? null : expiresAt!.toIso8601String(),
        "GeneratedAt": generatedAt == null ? null : generatedAt!.toIso8601String(),
        "User": user == user,
    };
}
