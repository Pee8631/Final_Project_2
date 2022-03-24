import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/model/errorMsg.dart';
import 'package:frontend_flutter/model/loginToken.dart';
import 'package:frontend_flutter/model/user.dart';
import 'package:frontend_flutter/screens/main_screen.dart';
import 'package:frontend_flutter/util/http_exception.dart';
import 'package:http/http.dart' as http;

//import 'main_screen.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({Key? key}) : super(key: key);

  @override
  _LoginScreenState createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  final formkey = GlobalKey<FormState>();
  User user = User(username: '', password: '', department: 0, hospital: 0);
  LoginToken loginToken = LoginToken(authToken: '', expiresAt: null, generatedAt: null, user: null);
  Future<User?> _futureUser() async {
    final url = Uri.parse('http://10.0.2.2:8080/api/v1/users/1');
    final response = await http.get(url);
    print(response.statusCode);
    final _results = userFromJson(response.body);
    return _results;
  }

  Future<void> authUser(User user) async {
    var response = await http.post(
      Uri.parse(
          'http://10.0.2.2:8080/api/v1/users/' + user.username.toString()),
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
        var token = loginTokenFromJson(response.body);
        print(token);
      } else {
        ErrorMsg? errorFromAPI = errorMsgFromJson(response.body);
        if (errorFromAPI.error == "Invalid Username") {
          throw "ไม่มีชื่อผู้ใช้นี้อยู่ในระบบ";
        } else if (errorFromAPI.error == "Invalid Password") {
          throw "รหัสผ่านไม่ถูกต้อง";
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
                                user.username = username ?? '';
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
                                user.password = password ?? '';
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
                                    await authUser(user).then((value) {
                                      formkey.currentState!.reset();
                                      Navigator.pushReplacement(context,
                                          MaterialPageRoute(builder: (context) {
                                        return MainScreen();
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
          return Scaffold(
            body: Center(
              child: CircularProgressIndicator(),
            ),
          );
        });
  }
}