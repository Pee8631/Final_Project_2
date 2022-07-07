import 'dart:convert';
import 'dart:async';
import 'package:flutter/material.dart';
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/models/certification.dart';
import 'package:frontend_flutter/models/department.dart';
import 'package:frontend_flutter/models/errorMsg.dart';
import 'package:frontend_flutter/models/hospital.dart';
import 'package:frontend_flutter/models/user.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:frontend_flutter/screens/signupandsignin/pinfo_doctors_screen.dart';
import 'package:frontend_flutter/screens/signupandsignin/signin_doctor_screen.dart';
import 'package:frontend_flutter/screens/signupandsignin/signin_user_screen.dart';
import 'package:frontend_flutter/util/http_exception.dart';
import 'package:frontend_flutter/widget/appbor_logindoctorscreen.dart';
import 'package:intl/intl.dart';
import 'package:http/http.dart' as http;
//import 'home.dart';

class SignUpDoctorScreen extends StatefulWidget {
  const SignUpDoctorScreen({Key? key}) : super(key: key);

  @override
  _SignUpDoctorScreenState createState() => _SignUpDoctorScreenState();
}

class _SignUpDoctorScreenState extends State<SignUpDoctorScreen> {
  final formkey = GlobalKey<FormState>();
  User _user = User(
      username: '',
      password: '',
      department: 0,
      hospital: 0,
      roleId: 0,
      certification: null,
      pInfo: null);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        backgroundColor: Color.fromARGB(255, 208, 244, 255),
        appBar: buildAppBarDoctorSignUp(context, SignInDoctorScreen()),
        body: Container(
          padding: const EdgeInsets.all(10.0),
          child: Form(
              key: formkey,
              child: SingleChildScrollView(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.center,
                  children: [
                    Container(
                      margin: const EdgeInsets.all(15),
                      child: Text("สมัครสมาชิกสำหรับหมอ",
                          style: TextStyle(
                              fontWeight: FontWeight.bold, fontSize: 18))),
                  Container(
                    margin: const EdgeInsets.only(top: 7.5, bottom: 7.5),
                    child: TextFormField(
                      validator: MultiValidator([
                        RequiredValidator(errorText: "กรุณาป้อนชื่อผู้ใช้"),
                      ]),
                      decoration: InputDecoration(
                        labelText: "ชื่อบัญชี",
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
                      onSaved: (username) {
                        _user.username = username!;
                      },
                    ),
                  ),
                  Container(
                    margin: const EdgeInsets.only(top: 7.5, bottom: 7.5),
                    child: TextFormField(
                      validator:
                          RequiredValidator(errorText: "กรุณาป้อนรหัสผ่าน"),
                      decoration: InputDecoration(
                        labelText: "รหัสผ่าน",
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
                      obscureText: true,
                      onSaved: (password) {
                        _user.password = password!;
                      },
                    ),
                  ),
                    SizedBox(
                      height: 15,
                    ),
                    Container(
                      margin: const EdgeInsets.only(top: 3.0, bottom: 3.0),
                      width: double.infinity,
                      height: 40,
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
                        child: Text("ต่อไป", style: TextStyle(fontSize: 16)),
                        onPressed: () async {
                          if (formkey.currentState!.validate()) {
                            formkey.currentState!.save();
                            try {
                              Navigator.pushReplacement(context,
                                  MaterialPageRoute(builder: (context) {
                                return PInfoDoctorScreen(
                                  user: _user,
                                );
                              }));
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
                  ],
                ),
              )),
        ));
  }
}
