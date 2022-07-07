import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/models/errorMsg.dart';
import 'package:frontend_flutter/models/loginToken.dart';
import 'package:frontend_flutter/models/user.dart';
import 'package:frontend_flutter/screens/doctor_screen/main_doctor_screen.dart';
import 'package:frontend_flutter/screens/signupandsignin/signin_user_screen.dart';
import 'package:frontend_flutter/screens/signupandsignin/signup_doctor_screen.dart';
import 'package:frontend_flutter/util/http_exception.dart';
import 'package:frontend_flutter/widget/appbar_loginscreen.dart';
import 'package:frontend_flutter/widget/home_button.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

//import 'main_screen.dart';

class SignInDoctorScreen extends StatefulWidget {
  const SignInDoctorScreen({Key? key}) : super(key: key);

  @override
  _SignInDoctorScreenState createState() => _SignInDoctorScreenState();
}

class _SignInDoctorScreenState extends State<SignInDoctorScreen> {
  final formkey = GlobalKey<FormState>();
  User user = User(
      username: '',
      password: '',
      department: 0,
      hospital: 0,
      roleId: 0,
      certification: null,
      pInfo: null);

  Future<void> authUser(User user) async {
    var response = await http.post(
      Uri.parse(
          'http://10.0.2.2:8080/api/v1/users/' + user.username.toString()),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        'Username': user.username,
        'Password': user.password,
        'RoleId': 2,
      }),
    );
    try {
      if (response.statusCode == 200) {
        var loginToken = loginTokenFromJson(response.body);
        SharedPreferences sharedPreferences =
            await SharedPreferences.getInstance();
        sharedPreferences.setString('username', loginToken.username);
        sharedPreferences.setString('authToken', loginToken.token!.authToken);
        sharedPreferences.setString(
            'expireAt', loginToken.token!.expiresAt.toString());
        sharedPreferences.setInt('role', loginToken.role!.first.id);
      } else {
        ErrorMsg? errorFromAPI = errorMsgFromJson(response.body);
        if (errorFromAPI.error == "Invalid Username") {
          throw "ไม่มีชื่อผู้ใช้นี้อยู่ในระบบ";
        } else if (errorFromAPI.error == "Invalid Password") {
          throw "รหัสผ่านไม่ถูกต้อง";
        } else if (errorFromAPI.error == "User has not permission to sign in") {
          throw "คุณไม่สามารถเข้าสู่ระบบของแพทย์ ด้วยบัญชีนี้ได้";
        } else {
          throw "ไม่สามารถเข้าสู่ระบบได้ในขณะนี้";
        }
      }
    } on HttpException {
      rethrow;
    }
  }

  /*String removeSomeWord(String msgError) {
    msgError = msgError.replaceAll(new RegExp(r'[^\w\s]+'), '');
    msgError = msgError.replaceAll("error", '');
    msgError = msgError.replaceAll("ent", '');
    return msgError;
  }*/
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
              child: Container(
                padding: const EdgeInsets.all(15),
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
                        child: Text("เข้าสู่ระบบ",
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
                    Container(
                      margin: const EdgeInsets.only(top: 3.0, bottom: 3.0),
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
                            Text("เข้าสู่ระบบ", style: TextStyle(fontSize: 18)),
                        onPressed: () async {
                          if (formkey.currentState!.validate()) {
                            formkey.currentState!.save();
                          }
                          try {
                            await authUser(user).then((value) {
                              formkey.currentState!.reset();
                              Navigator.pushReplacement(context,
                                  MaterialPageRoute(builder: (context) {
                                return MainDoctorScreen(
                                  ScreenIndex: 0,
                                );
                              }));
                            });
                          } catch (error) {
                            Fluttertoast.showToast(
                                msg: error.toString(),
                                gravity: ToastGravity.CENTER);
                          }
                        },
                      ),
                    ),
                    buildButtonDoctorSignUp(),
                  ],
                ),
              ),
            )),
      ),
      bottomNavigationBar: buildButtonDoctorSignIn(),
    );
  }

  Widget buildButtonDoctorSignUp() {
    return Row(
      crossAxisAlignment: CrossAxisAlignment.center,
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Text('คุณยังไม่มีบัญชีใช่หรือไม่'),
        TextButton(
            child: Text(
              'สมัครสมาชิก',
              style: TextStyle(fontFamily: ''),
            ),
            onPressed: () => Navigator.pushReplacement(context,
                    MaterialPageRoute(builder: (context) {
                  return SignUpDoctorScreen();
                }))),
      ],
    );
  }

  Widget buildButtonDoctorSignIn() {
    return Container(
      margin: const EdgeInsets.all(10),
      child: TextButton(
          child: Text(
            'กลับไปหน้าเข้าสู่ระบบของผู้ใช้ทั่วไป',
            style: TextStyle(fontFamily: ''),
          ),
          onPressed: () => Navigator.pushReplacement(context,
                  MaterialPageRoute(builder: (context) {
                return SignInUserScreen();
              }))),
    );
  }
}
