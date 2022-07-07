import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/models/certification.dart';
import 'package:frontend_flutter/models/department.dart';
import 'package:frontend_flutter/models/errorMsg.dart';
import 'package:frontend_flutter/models/hospital.dart';
import 'package:frontend_flutter/models/user.dart';
import 'package:frontend_flutter/screens/signupandsignin/dep_hos_doctor_screen.dart';
import 'package:frontend_flutter/screens/signupandsignin/pinfo_doctors_screen.dart';
import 'package:frontend_flutter/screens/signupandsignin/signin_doctor_screen.dart';
import 'package:frontend_flutter/util/http_exception.dart';
import 'package:frontend_flutter/widget/appbor_logindoctorscreen.dart';
import 'package:http/http.dart' as http;
import 'package:intl/intl.dart';

class CreateCertificationScreen extends StatefulWidget {
  final User user;
  const CreateCertificationScreen({Key? key, required this.user})
      : super(key: key);

  @override
  State<CreateCertificationScreen> createState() =>
      _CreateCertificationScreenState();
}

class _CreateCertificationScreenState extends State<CreateCertificationScreen> {
  final formkey = GlobalKey<FormState>();
  User user = User(
      username: '',
      password: '',
      department: 0,
      hospital: 0,
      roleId: 0,
      certification: null,
      pInfo: null);
  Certification certification = Certification(
      code: '',
      dateOfExp: null,
      dateOfIssuing: null,
      diloma: '',
      issuer: '',
      user: null);

  Future<void> createUser(User user, Certification certification) async {
    final response = await http.post(
      Uri.parse('http://10.0.2.2:8080/api/v1/users'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        'username': user.username,
        'password': user.password,
        'department': user.department,
        'hospital': user.hospital,
        'roleId': 2,
        'Certification': certification,
        'PInfo': user.pInfo
      }),
    );
    try {
      if (response.statusCode == 200) {
        User? userFromAPI = userFromJson(response.body);
        print(userFromAPI);
      } else {
        ErrorMsg? errorFromAPI = errorMsgFromJson(response.body);
        if (errorFromAPI.error == "saving failed") {
          throw "ชื่อผู้ใช้นี้มีในระบบอยู่แล้ว";
        } else {
          throw "ไม่สามารถสมัครบัญชีใหม่ได้ \nกรุณาตรวจสอบชื่อผู้ใช้และรหัสผ่านใหม่อีกครั้ง";
        }
      }
    } on HttpException {
      rethrow;
    }
  }
  @override
  void initState() {
    setState(() {
      user = widget.user;
    });
    super.initState();
  }

  DateTime _selectdateOfIssuing = new DateTime.now();
  DateTime _selectdateOfExp = new DateTime.now();

  @override
  Widget build(BuildContext context) {
    String _formatdateOfIssuing = new DateFormat.yMMMd()
        .format(certification.dateOfIssuing ?? _selectdateOfIssuing);
    String _formatdateOfExp = new DateFormat.yMMMd()
        .format(certification.dateOfExp ?? _selectdateOfExp);

    return Scaffold(
      backgroundColor: Color.fromARGB(255, 208, 244, 255),
      appBar: buildAppBarDoctorSignUp(
          context,
          DepHosDoctorScreen(
            user: widget.user,
          )),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(20),
        child: Form(
          key: formkey,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Container(
                margin: const EdgeInsets.only(top: 10, bottom: 10),
                child: Center(
                  child: Text("ใบอนุญาต Certification",
                      style:
                          TextStyle(fontWeight: FontWeight.bold, fontSize: 20)),
                ),
              ),
              Container(
                margin: const EdgeInsets.only(top: 7.5, bottom: 7.5),
                child: TextFormField(
                  validator:
                      RequiredValidator(errorText: "กรุณาป้อนเลขวิชาชีพแพทย์"),
                  decoration: InputDecoration(
                    labelText: "เลขวิชาชีพแพทย์",
                    fillColor: Colors.white,
                    focusedBorder: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(25.0),
                      borderSide: BorderSide(
                        color: Color.fromARGB(228, 96, 239, 220),
                        width: 3.0,
                      ),
                    ),
                    enabledBorder: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(25.0),
                      borderSide: BorderSide(
                        color: Color.fromARGB(235, 111, 137, 162),
                        width: 3.0,
                      ),
                    ),
                  ),
                  onSaved: (Code) {
                    certification.code = Code!;
                  },
                ),
              ),
              Container(
                margin: const EdgeInsets.only(top: 7.5, bottom: 7.5),
                child: TextFormField(
                  validator:
                      RequiredValidator(errorText: "กรุณาป้อนเลขวิชาชีพแพทย์"),
                  decoration: InputDecoration(
                    labelText: "เลขใบประกาศนียบัตรวิชาชีพแพทย์",
                    fillColor: Colors.white,
                    focusedBorder: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(25.0),
                      borderSide: BorderSide(
                        color: Color.fromARGB(228, 96, 239, 220),
                        width: 3.0,
                      ),
                    ),
                    enabledBorder: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(25.0),
                      borderSide: BorderSide(
                        color: Color.fromARGB(235, 111, 137, 162),
                        width: 3.0,
                      ),
                    ),
                  ),
                  onSaved: (diloma) {
                    certification.diloma = diloma!;
                  },
                ),
              ),
              Row(
                children: [
                  Expanded(
                    flex: 2,
                    child: Text("วันที่ออกใบประกาศณียบัตร",
                        style: TextStyle(fontSize: 16)),
                  ),
                  Expanded(
                    flex: 1,
                    child: ElevatedButton(
                      style: ElevatedButton.styleFrom(
                        primary: Color.fromARGB(220, 96, 239, 220),
                        onPrimary: Colors.white,
                        shadowColor: Colors.greenAccent,
                        elevation: 3,
                        shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(32.0)),
                        maximumSize: Size(100, 40), //////// HERE
                      ),
                      child: Text(_formatdateOfIssuing.toString()),
                      onPressed: () {
                        showDatePicker(
                                context: context,
                                initialDate: certification.dateOfIssuing ??
                                    _selectdateOfIssuing,
                                firstDate: DateTime(1990),
                                lastDate: DateTime(2025))
                            .then((dateOfIssuing) {
                          setState(() {
                            certification.dateOfIssuing = dateOfIssuing!;
                          });
                        });
                      },
                    ),
                  ),
                ],
              ),
              Row(
                children: [
                  Expanded(
                    flex: 2,
                    child: Text("วันหมดอายุใบประกาศณียบัตร",
                        style: TextStyle(fontSize: 16)),
                  ),
                  Expanded(
                    flex: 1,
                    child: ElevatedButton(
                      style: ElevatedButton.styleFrom(
                        primary: Color.fromARGB(220, 96, 239, 220),
                        onPrimary: Colors.white,
                        shadowColor: Colors.greenAccent,
                        elevation: 3,
                        shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(32.0)),
                        maximumSize: Size(100, 40), //////// HERE
                      ),
                      child: Text(_formatdateOfExp.toString()),
                      onPressed: () {
                        showDatePicker(
                                context: context,
                                initialDate:
                                    certification.dateOfExp ?? _selectdateOfExp,
                                firstDate: DateTime(1990),
                                lastDate: DateTime(2025))
                            .then((dateOfExp) {
                          setState(() {
                            certification.dateOfExp = dateOfExp!;
                          });
                        });
                      },
                    ),
                  ),
                ],
              ),
              Container(
                margin: const EdgeInsets.only(top: 7.5, bottom: 7.5),
                child: TextFormField(
                  validator: RequiredValidator(
                      errorText: "กรุณาป้อนผู้ออกใบประกาศณียบัตร"),
                  decoration: InputDecoration(
                    labelText:
                        "ใคร หน่วยงานหรือองค์กรไหนเป็นผู้ออกใบประกาศณียบัตร",
                    fillColor: Colors.white,
                    focusedBorder: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(25.0),
                      borderSide: BorderSide(
                        color: Color.fromARGB(228, 96, 239, 220),
                        width: 3.0,
                      ),
                    ),
                    enabledBorder: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(25.0),
                      borderSide: BorderSide(
                        color: Color.fromARGB(235, 111, 137, 162),
                        width: 3.0,
                      ),
                    ),
                  ),
                  onSaved: (issuer) {
                    certification.issuer = issuer!;
                  },
                ),
              ),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton(
                  style: ElevatedButton.styleFrom(
                    primary: Color.fromARGB(220, 96, 239, 220),
                    onPrimary: Colors.white,
                    shadowColor: Colors.greenAccent,
                    elevation: 3,
                    shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(32.0)),
                    maximumSize: Size(100, 40), //////// HERE
                  ),
                  child: Text("ลงทะเบียน", style: TextStyle(fontSize: 20)),
                  onPressed: () async {
                    if (formkey.currentState!.validate()) {
                      formkey.currentState!.save();
                      try {
                        await createUser(user, certification).then((value) {
                          Fluttertoast.showToast(
                              msg: "สร้างบัญชีผู้ใช้เรียบร้อยแล้ว",
                              gravity: ToastGravity.CENTER);
                          Navigator.pushReplacement(context,
                              MaterialPageRoute(builder: (context) {
                            return SignInDoctorScreen();
                          }));
                        });
                      } catch (error) {
                        Fluttertoast.showToast(
                            msg: error.toString(),
                            gravity: ToastGravity.CENTER);
                      }
                      formkey.currentState!.reset();
                    }
                  },
                ),
              ),
              Text(user.username!),
              Text(user.pInfo!.firstName),
              Text(user.hospital.toString()),
              Text(user.department.toString())
            ],
          ),
        ),
      ),
    );
  }
}
