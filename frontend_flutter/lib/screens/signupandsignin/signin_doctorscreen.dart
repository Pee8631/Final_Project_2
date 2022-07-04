import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/models/errorMsg.dart';
import 'package:frontend_flutter/models/loginToken.dart';
import 'package:frontend_flutter/models/user.dart';
import 'package:frontend_flutter/screens/doctor_screen/main_doctor_screen.dart';
import 'package:frontend_flutter/screens/signupandsignin/signup_doctorscreen.dart';
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
  User user =
      User(username: '', password: '', department: 0, hospital: 0, roleId: 0, certification: null);

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
        SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
        sharedPreferences.setString('username', loginToken.username);
        sharedPreferences.setString('authToken',loginToken.token!.authToken);
        sharedPreferences.setString('expireAt',
            loginToken.token!.expiresAt.toString());
        sharedPreferences.setInt('role',loginToken.role!.first.id);
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
                      style: TextStyle(fontSize: 18),
                    ),
                    TextFormField(
                      validator: MultiValidator([
                        RequiredValidator(errorText: "กรุณาป้อนชื่อผู้ใช้"),
                      ]),
                      onSaved: (username) {
                        user.username = username ?? '';
                      },
                    ),
                    SizedBox(
                      height: 15,
                    ),
                    Text(
                      "รหัสผ่าน",
                      style: TextStyle(fontSize: 18),
                    ),
                    TextFormField(
                      validator:
                          RequiredValidator(errorText: "กรุณาป้อนรหัสผ่าน"),
                      obscureText: true,
                      onSaved: (password) {
                        user.password = password ?? '';
                      },
                    ),
                    SizedBox(
                      height: 15,
                    ),
                    HomeButton(Icon(Icons.login), "สร้างบัญชีแพทย์", 18,
                        Colors.blue, SignUpDoctorScreen()),
                    Container(
                      decoration: BoxDecoration(border: Border.all(color: Colors.blueGrey)),
                      margin: const EdgeInsets.only(top: 3.0, bottom: 3.0),
                      width: double.infinity,
                      child: ElevatedButton(
                        child: Text("ลงชื่อเข้าใช้",
                            style: TextStyle(fontSize: 18)),
                        onPressed: () async {
                          if (formkey.currentState!.validate()) {
                            formkey.currentState!.save();
                          }
                          try {
                            await authUser(user).then((value) {
                              formkey.currentState!.reset();
                              Navigator.pushReplacement(context,
                                  MaterialPageRoute(builder: (context) {
                                return MainDoctorScreen(ScreenIndex: 0,);
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
                  ],
                ),
              )),
        ));
  }
}
