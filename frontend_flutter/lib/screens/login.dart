import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/model/loginToken.dart';
import 'package:frontend_flutter/model/user.dart';
import 'package:frontend_flutter/screens/main_screen.dart';
import 'package:http/http.dart' as http;
import 'package:jwt_decoder/jwt_decoder.dart';

//import 'main_screen.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({Key? key}) : super(key: key);

  @override
  _LoginScreenState createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  final formkey = GlobalKey<FormState>();
  User user = User(username: '', password: '', department: 0, hospital: 0);
  LoginToken loginToken = LoginToken(authorized: false, exp: 0, userId: '');
  Future<User> _futureUser() async {
    final url = Uri.parse('http://10.0.2.2:8080/api/v1/users/1');
    final response = await http.get(url);
    print(response.statusCode);
    final _results = userFromJson(response.body);
    return _results;
  }

  Future<Map<String, dynamic>?> authUser(User user) async {
    var response = await http.post(
      Uri.parse('http://10.0.2.2:8080/api/v1/users/' + user.username),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        'username': user.username,
        'password': user.password,
        'department': 0,
        'hospital': 0,
      }),
    );
    try {
      if (response.statusCode == 200) {
        Map<String, dynamic> token = JwtDecoder.decode(response.body);
        print(token);
        return token;
      } else {
        Map<String, dynamic> err = {'error': response.body};
        return err;
      }
    } catch (error) {
      Fluttertoast.showToast(
          msg: error.toString(), gravity: ToastGravity.CENTER);
    }
  }

  String removeSomeWord(String msgError) {
    msgError = msgError.replaceAll(new RegExp(r'[^\w\s]+'), '');
    msgError = msgError.replaceAll("error", '');
    msgError = msgError.replaceAll("ent", '');
    return msgError;
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
        future: _futureUser(),
        builder: (context, snapshot) {
          if (snapshot.hasError) {
            return Scaffold(
              appBar: AppBar(
                title: Text("Error"),
              ),
              body: Center(
                child: Text("${snapshot.error}"),
              ),
            );
          }
          if (snapshot.connectionState == ConnectionState.done) {
            return Scaffold(
                appBar: AppBar(
                  title: Text("เข้าสู่ระบบ"),
                ),
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
                                RequiredValidator(
                                    errorText: "กรุณาป้อนชื่อผู้ใช้"),
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
                              validator: RequiredValidator(
                                  errorText: "กรุณาป้อนรหัสผ่าน"),
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
                                child: Text("ลงชื่อเข้าใช้",
                                    style: TextStyle(fontSize: 20)),
                                onPressed: () async {
                                  if (formkey.currentState!.validate()) {
                                    formkey.currentState!.save();
                                  }
                                  try {
                                    Map<String, dynamic>? accessToken =
                                        await authUser(user);

                                    if (accessToken!['authorized'] == true) {
                                      print("บันทึกสำเร็จ");
                                      print(accessToken);
                                      formkey.currentState!.reset();
                                      Navigator.pushReplacement(context,
                                          MaterialPageRoute(builder: (context) {
                                        return MainScreen(accessToken);
                                      }));
                                    } else {
                                      String msgError =
                                          accessToken.values.toString();
                                      msgError = removeSomeWord(msgError);
                                      print(msgError);
                                      switch (msgError) {
                                        case " user not found":
                                          msgError =
                                              "ไม่พบชื่อผู้ใช้ในระบบ\nกรุณาตรวจสอบชื่อผู้ใช้ใหม่อีกครั่ง";
                                          break;
                                        case "invalid Password":
                                          msgError =
                                              "รหัสผ่านไม่ถูกต้อง\nกรุณากรอกรหัสผ่านใหม่อีกครั่ง";
                                          break;
                                        default:
                                          msgError =
                                              "ไม่สามารถล็อกอินได้ \nกรุณาตรวจสอบชื่อบัญชีและรหัสผ่านใหม่อีกครั้ง";
                                      }
                                      Fluttertoast.showToast(
                                          msg: msgError,
                                          gravity: ToastGravity.CENTER);
                                    }
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
          return Scaffold(
            body: Center(
              child: CircularProgressIndicator(),
            ),
          );
        });
  }
}
