import 'dart:convert';
import 'dart:async';
import 'package:flutter/material.dart';
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/model/user.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:http/http.dart' as http;

import 'home.dart';

class RegisterScreen extends StatefulWidget {
  const RegisterScreen({Key? key}) : super(key: key);

  @override
  _RegisterScreenState createState() => _RegisterScreenState();
}

class _RegisterScreenState extends State<RegisterScreen> {
  final formkey = GlobalKey<FormState>();
  User user = User(username: '', password: '');
  Future<User>? _futureUser;

  Future<void> createUser(String username, String password) async {
    final response = await http.post(
      Uri.parse('http://10.0.2.2:8080/api/v1/users'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        'username': username,
        'password': password,
      }),
    );

    if (response.statusCode == 200) {
      User userFromAPI = userFromJson(response.body);
      print(userFromAPI);
    } else {
      var error;
      if (response.body == "saving failed") {
        error = "ชื่อผู้ใช้นี้มีในระบบอยู่แล้ว";
      } else {
        error =
            "ไม่สามารถสมัครบัญชีใหม่ได้ \nกรุณาตรวจสอบชื่อผู้ใช้และรหัสผ่านใหม่อีกครั้ง";
      }
      Fluttertoast.showToast(msg: error, gravity: ToastGravity.CENTER);
    }
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
        future: _futureUser,
        builder: (context, snapshot) {
          /*if (snapshot.hasError) {
            return Scaffold(
              appBar: AppBar(
                title: Text("Error"),
              ),
              body: Center(
                child: Text("${snapshot.error}"),
              ),
            );
          }*/
          //if (snapshot.connectionState == ConnectionState.done) {
          return Scaffold(
              appBar: AppBar(
                title: Text("สร้างบัญชีผู้ใช้"),
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
                              child: Text("ลงทะเบียน",
                                  style: TextStyle(fontSize: 20)),
                              onPressed: () async {
                                if (formkey.currentState!.validate()) {
                                  formkey.currentState!.save();
                                  try {
                                    await createUser(
                                            user.username, user.password)
                                        .then((value) {
                                      Fluttertoast.showToast(
                                          msg: "สร้างบัญชีผู้ใช้เรียบร้อยแล้ว",
                                          gravity: ToastGravity.CENTER);
                                      print(
                                          "username = ${user.username} password = ${user.password} ");
                                      Navigator.pushReplacement(context,
                                          MaterialPageRoute(builder: (context) {
                                        return HomeScreen();
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
          //}
          /*return Scaffold(
            body: Center(
              child: CircularProgressIndicator(),
            ),
          );*/
        });
  }
}
