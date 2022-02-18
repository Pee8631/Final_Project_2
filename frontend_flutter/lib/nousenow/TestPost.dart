import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:frontend_flutter/screens/main_screen.dart';
import 'package:http/http.dart' as http;
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/model/user.dart';
import 'package:jwt_decoder/jwt_decoder.dart';

class TestPostScreen extends StatefulWidget {
  const TestPostScreen({Key? key}) : super(key: key);

  @override
  _TestPostScreenState createState() => _TestPostScreenState();
}

Future<dynamic> authUser(User user) async {
  var response = await http.post(
    Uri.parse('http://10.0.2.2:8080/api/v1/users/' + user.username),
    headers: <String, String>{
      'Content-Type': 'application/json; charset=UTF-8',
    },
    body: jsonEncode(<String, dynamic>{
      'username': user.username,
      'password': user.password,
    }),
  );
  if (response.statusCode == 200) {
    var token = JwtDecoder.decode(response.body);
    //print(token);
    return token;
  } else {
    return null;
  }
}

class _TestPostScreenState extends State<TestPostScreen> {
  final formkey = GlobalKey<FormState>();
  User user = User(username: '', password: '', department: 0 , hospital: 0);
  Future<User>? _futureUser;
  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
        future: _futureUser,
        builder: (context, snapshot) {
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
                            "Username",
                            style: TextStyle(fontSize: 20),
                          ),
                          TextFormField(
                            validator: MultiValidator([
                              RequiredValidator(
                                  errorText: "กรุณาป้อน Username"),
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
                                  var accessToken = await authUser(user);
                                  if (accessToken != null) {
                                    print(accessToken);
                                    print("บันทึกสำเร็จ");
                                    formkey.currentState!.reset();
                                    Navigator.pushReplacement(context,
                                        MaterialPageRoute(builder: (context) {
                                      return MainScreen(accessToken);
                                    }));
                                  }
                                } catch (error) {}
                              },
                            ),
                          ),
                        ],
                      ),
                    )),
              ));
        });
  }
}
