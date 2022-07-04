// To parse this JSON data, do
//
//     final hospital = hospitalFromJson(jsonString);

import 'dart:convert';

List<Hospital> hospitalFromJson(String str) => List<Hospital>.from(json.decode(str).map((x) => Hospital.fromJson(x)));

String hospitalToJson(List<Hospital> data) => json.encode(List<dynamic>.from(data.map((x) => x.toJson())));

class Hospital {
    Hospital({
        required this.id,
        required this.name,
        required this.edges,
    });

    int? id;
    String? name;
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

class Edges {
    Edges();

    factory Edges.fromJson(Map<String, dynamic> json) => Edges(
    );

    Map<String, dynamic> toJson() => {
    };
}
