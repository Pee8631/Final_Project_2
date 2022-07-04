// To parse this JSON data, do
//
//     final department = departmentFromJson(jsonString);

import 'dart:convert';

List<Department> departmentsFromJson(String str) => List<Department>.from(json.decode(str).map((x) => Department.fromJson(x)));

String departmentsToJson(List<Department> data) => json.encode(List<dynamic>.from(data.map((x) => x.toJson())));

Department departmentFromJson(String str) => Department.fromJson(json.decode(str));

String departmentToJson(Department data) => json.encode(data.toJson());

class Department {
    Department({
        required this.id,
        required this.name,
        required this.image,
        required this.edges,
    });

    int id;
    String name;
    String image;
    DepartmentEdges? edges;

    factory Department.fromJson(Map<String, dynamic> json) => Department(
        id: json["id"] == null ? null : json["id"],
        name: json["name"] == null ? null : json["name"],
        image: json["image"] == null ? null : json["image"],
        edges: json["edges"] == null ? null : DepartmentEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "name": name,
        "image": image,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class DepartmentEdges {
    DepartmentEdges({
        required this.departmentHasDoctor,
    });

    List<DepartmentHasDoctor>? departmentHasDoctor;

    factory DepartmentEdges.fromJson(Map<String, dynamic> json) => DepartmentEdges(
        departmentHasDoctor: json["department_has_doctor"] == null ? null : List<DepartmentHasDoctor>.from(json["department_has_doctor"].map((x) => DepartmentHasDoctor.fromJson(x))),
    );

    Map<String, dynamic> toJson() => {
        "department_has_doctor": departmentHasDoctor == null ? null : List<dynamic>.from(departmentHasDoctor!.map((x) => x.toJson())),
    };
}

class DepartmentHasDoctor {
    DepartmentHasDoctor({
        required this.id,
        required this.username,
        required this.password,
        required this.edges,
    });

    int id;
    String username;
    String password;
    DepartmentHasDoctorEdges? edges;

    factory DepartmentHasDoctor.fromJson(Map<String, dynamic> json) => DepartmentHasDoctor(
        id: json["id"] == null ? null : json["id"],
        username: json["username"] == null ? null : json["username"],
        password: json["password"] == null ? null : json["password"],
        edges: json["edges"] == null ? null : DepartmentHasDoctorEdges.fromJson(json["edges"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "username": username,
        "password": password,
        "edges": edges == null ? null : edges!.toJson(),
    };
}

class DepartmentHasDoctorEdges {
    DepartmentHasDoctorEdges();

    factory DepartmentHasDoctorEdges.fromJson(Map<String, dynamic> json) => DepartmentHasDoctorEdges(
    );

    Map<String, dynamic> toJson() => {
    };
}

