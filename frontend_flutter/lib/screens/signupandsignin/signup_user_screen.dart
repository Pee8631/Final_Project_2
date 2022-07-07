import 'dart:convert';
import 'dart:async';
import 'package:flutter/material.dart';
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/models/errorMsg.dart';
import 'package:frontend_flutter/models/user.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:frontend_flutter/screens/home_screen.dart';
import 'package:frontend_flutter/screens/signupandsignin/signin_user_screen.dart';
import 'package:frontend_flutter/util/http_exception.dart';
import 'package:frontend_flutter/widget/appbar_loginscreen.dart';
import 'package:http/http.dart' as http;
//import 'home.dart';

class SignUpUserScreen extends StatefulWidget {
  const SignUpUserScreen({Key? key}) : super(key: key);

  @override
  _SignUpUserScreenState createState() => _SignUpUserScreenState();
}

class _SignUpUserScreenState extends State<SignUpUserScreen> {
  final formkey = GlobalKey<FormState>();
  User user = User(
      username: '',
      password: '',
      department: 0,
      hospital: 0,
      roleId: 0,
      certification: null, pInfo: null);

  Future<void> createUser(String username, String password) async {
    final response = await http.post(
      Uri.parse('http://10.0.2.2:8080/api/v1/users'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        'username': username,
        'password': password,
        'department': 0,
        'hospital': 0,
        'roleId': 3,
        'certification': null,
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
  Widget build(BuildContext context) {
    return Scaffold(
        backgroundColor: Color.fromARGB(255, 208, 244, 255),
        appBar: buildAppBarLogin(context),
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
                        child: Text("Online Doctor Application",
                            style: TextStyle(
                              fontWeight: FontWeight.bold,
                              fontSize: 24,
                            ))),
                    Container(
                        margin: const EdgeInsets.all(15),
                        child: Text("สมัครสมาชิก",
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
                          user.username = username ?? '';
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
                          user.password = password ?? '';
                        },
                      ),
                    ),
                    SizedBox(
                      height: 15,
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
                        child:
                            Text("ลงทะเบียน", style: TextStyle(fontSize: 20)),
                        onPressed: () async {
                          if (formkey.currentState!.validate()) {
                            formkey.currentState!.save();
                            try {
                              await createUser(user.username.toString(),
                                      user.password.toString())
                                  .then((value) {
                                Fluttertoast.showToast(
                                    msg: "สร้างบัญชีผู้ใช้เรียบร้อยแล้ว",
                                    gravity: ToastGravity.CENTER);
                                Navigator.pushReplacement(context,
                                    MaterialPageRoute(builder: (context) {
                                  return SignInUserScreen();
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
                    buildButtonUserSignIn(),
                  ],
                ),
              )),
        ));
  }

  Widget buildButtonUserSignIn() {
    return Container(
      margin: const EdgeInsets.all(10),
      child: TextButton(
          child: Text(
            'กลับไปหน้าเข้าสู่ระบบ',
            style: TextStyle(fontFamily: ''),
          ),
          onPressed: () => Navigator.pushReplacement(context,
                  MaterialPageRoute(builder: (context) {
                return SignInUserScreen();
              }))),
    );
  }
}
