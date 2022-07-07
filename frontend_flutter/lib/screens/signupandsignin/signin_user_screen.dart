import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/models/errorMsg.dart';
import 'package:frontend_flutter/models/loginToken.dart';
import 'package:frontend_flutter/models/user.dart';
import 'package:frontend_flutter/screens/signupandsignin/signin_doctor_screen.dart';
import 'package:frontend_flutter/screens/signupandsignin/signup_doctor_screen.dart';
import 'package:frontend_flutter/screens/signupandsignin/signup_user_screen.dart';
import 'package:frontend_flutter/screens/user_screen/main_user_screen.dart';
import 'package:frontend_flutter/util/http_exception.dart';
import 'package:frontend_flutter/widget/appbar_loginscreen.dart';
import 'package:frontend_flutter/widget/home_button.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

//import 'main_screen.dart';

class SignInUserScreen extends StatefulWidget {
  const SignInUserScreen({Key? key}) : super(key: key);

  @override
  _SignInUserScreenState createState() => _SignInUserScreenState();
}

class _SignInUserScreenState extends State<SignInUserScreen> {
  final formkey = GlobalKey<FormState>();
  User user = User(
      username: '',
      password: '',
      department: 0,
      hospital: 0,
      roleId: 0,
      certification: null, pInfo: null);

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
        'RoleId': 3,
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
          throw "คุณไม่สามารถเข้าสู่ระบบของผู้ใช้ ด้วยบัญชีนี้ได้";
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
      body: Form(
          key: formkey,
          child: SingleChildScrollView(
            child: Container(
              padding: const EdgeInsets.all(15),
              margin: const EdgeInsets.only(top: 20, bottom: 20),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.center,
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Container(
                      margin: const EdgeInsets.all(15),
                      child: Text("Online Doctor Application",
                          style: TextStyle(
                              fontWeight: FontWeight.bold, fontSize: 24,))),
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
                      child:
                          Text("เข้าสู่ระบบ", style: TextStyle(fontSize: 16)),
                      onPressed: () async {
                        if (formkey.currentState!.validate()) {
                          formkey.currentState!.save();
                        }
                        try {
                          await authUser(user).then((value) {
                            formkey.currentState!.reset();
                            Navigator.pushReplacement(context,
                                MaterialPageRoute(builder: (context) {
                              return MainUserScreen(
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
                  buildButtonUserSignUp(),
                ],
              ),
            ),
          )),
      bottomNavigationBar: buildButtonDoctorSignIn(),
    );
  }

  Widget buildButtonUserSignUp() {
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
                  return SignUpUserScreen();
                }))),
      ],
    );
  }

  Widget buildButtonDoctorSignIn() {
    return Row(
      crossAxisAlignment: CrossAxisAlignment.center,
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Text('คุณเป็นหมอหรือไม่'),
        TextButton(
            child: Text(
              'ลงชื่อเข้าใช้สำหรับหมอ',
              style: TextStyle(fontFamily: ''),
            ),
            onPressed: () => Navigator.pushReplacement(context,
                    MaterialPageRoute(builder: (context) {
                  return SignInDoctorScreen();
                }))),
      ],
    );
  }
}
