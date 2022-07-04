import 'dart:convert';
import 'dart:async';
import 'package:flutter/material.dart';
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/models/errorMsg.dart';
import 'package:frontend_flutter/models/user.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:frontend_flutter/screens/home_screen.dart';
import 'package:frontend_flutter/screens/signupandsignin/signin_userscreen.dart';
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
  User user = User(username: '', password: '', department: 0, hospital: 0, roleId: 0, certification: null);

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
        'roleId' : 3,
        'certification' : null,
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
        appBar: buildAppBarLogin(context),
        body: Container(
          padding: const EdgeInsets.all(10.0),
          child: Form(
              key: formkey,
              child: SingleChildScrollView(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      "ชื่อผู้ใช้",
                      style: TextStyle(fontSize: 20),
                    ),
                    TextFormField(
                      validator: MultiValidator([
                        RequiredValidator(errorText: "กรุณาป้อนชื่อผู้ใช้"),
                      ]),
                      onSaved: (username) {
                        user.username = username!;
                      },
                    ),
                    SizedBox(
                      height: 15,
                    ),
                    Text(
                      "รหัสผ่าน",
                      style: TextStyle(fontSize: 20),
                    ),
                    TextFormField(
                      validator: MultiValidator([
                        RequiredValidator(errorText: "กรุณาป้อนรหัสผ่าน"),
                        MinLengthValidator(6,
                            errorText: 'รหัสผ่านจะต้องมี 6 ตัวขึ้นไป')
                      ]),
                      obscureText: true,
                      onSaved: (password) {
                        user.password = password!;
                      },
                    ),
                    SizedBox(
                      height: 15,
                    ),
                    SizedBox(
                      width: double.infinity,
                      child: ElevatedButton(
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
                                  return SignInScreen();
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
                  ],
                ),
              )),
        ));
  }
}
