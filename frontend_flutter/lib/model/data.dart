// To parse this JSON data, do
//
//     final data = dataFromJson(jsonString);
import 'dart:convert';

Data dataFromJson(String str) => Data.fromJson(json.decode(str));

String dataToJson(Data data) => json.encode(data.toJson());

class Data {
    Data({
        required this.idCardNumber,
        required this.firstName,
        required this.lastName,
        required this.gender,
        required this.brithDate,
        required this.bloodGroup,
        required this.address,
        required this.user,
    });

    String? idCardNumber;
    String? firstName;
    String? lastName;
    int? gender;
    DateTime? brithDate;
    String? bloodGroup;
    String? address;
    int? user;

    factory Data.fromJson(Map<String, dynamic> json) => Data(
        idCardNumber: json["IdCardNumber"] == null ? null : json["IdCardNumber"],
        firstName: json["FirstName"] == null ? null : json["FirstName"],
        lastName: json["LastName"] == null ? null : json["LastName"],
        gender: json["Gender"] == null ? null : json["Gender"],
        brithDate: json["BrithDate"] == null ? null : DateTime.parse(json["BrithDate"]),
        bloodGroup: json["BloodGroup"] == null ? null : json["BloodGroup"],
        address: json["Address"] == null ? null : json["Address"],
        user: json["User"] == null ? null : json["User"],
    );

    Map<String, dynamic> toJson() => {
        "IdCardNumber": idCardNumber == null ? null : idCardNumber,
        "FirstName": firstName == null ? null : firstName,
        "LastName": lastName == null ? null : lastName,
        "Gender": gender == null ? null : gender,
        "BrithDate": brithDate == null ? null : brithDate!.toIso8601String(),
        "BloodGroup": bloodGroup == null ? null : bloodGroup,
        "Address": address == null ? null : address,
        "User": user == null ? null : user,
    };
}
