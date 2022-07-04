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
import 'package:frontend_flutter/screens/signupandsignin/signin_userscreen.dart';
import 'package:frontend_flutter/util/http_exception.dart';
import 'package:frontend_flutter/widget/appbar_loginscreen.dart';
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
  User user = User(
      username: '',
      password: '',
      department: 0,
      hospital: 0,
      roleId: 0,
      certification: null);
  Certification certification = Certification(
      code: '',
      dateOfExp: null,
      dateOfIssuing: null,
      diloma: '',
      issuer: '',
      user: null);

  late Future<List<Department>> _department;
  //['Ears', 'Eyes', 'Nose', 'Mouth', 'Kid']
  late Future<List<Hospital>> _hospital;
  //['example hospital', 'example 2 Hospital']

  String? valueDepartment;
  String? valueHospital;

  @override
  initState() {
    _department = _getDepartment();
    _hospital = _getHospital();
    super.initState();
    // _getHospital().then((value) => setState(() => _hospital = value));
    // _getDepartment().then((value) => setState(() => _department = value));
  }

  Future<List<Department>> _getDepartment() async {
    final url = Uri.parse('http://10.0.2.2:8080/api/v1/departments');
    final response = await http.get(url);
    print(response.statusCode);
    print(response.body);
    final results = departmentsFromJson(response.body);
    return results;
    //setState(() => _department = results);
    //return await results;
  }

  Future<List<Hospital>> _getHospital() async {
    final url = Uri.parse('http://10.0.2.2:8080/api/v1/hospitals');
    final response = await http.get(url);
    print(response.statusCode);
    print(response.body);
    final results = hospitalFromJson(response.body);
    return results;
    //setState(() => _hospital = results);
  }

  Future<void> createUser(User user, Certification certification) async {
    final response = await http.post(
      Uri.parse('http://10.0.2.2:8080/api/v1/users'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        'username': user.username,
        'password': user.password,
        'department': user.department,
        'hospital': user.hospital,
        'roleId': 2,
        'Certification': certification
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

  DateTime _selectdateOfIssuing = new DateTime.now();
  DateTime _selectdateOfExp = new DateTime.now();

  @override
  Widget build(BuildContext context) {
    String _formatdateOfIssuing = new DateFormat.yMMMd()
        .format(certification.dateOfIssuing ?? _selectdateOfIssuing);
    String _formatdateOfExp = new DateFormat.yMMMd()
        .format(certification.dateOfExp ?? _selectdateOfExp);
    return FutureBuilder(
      future: Future.wait([_department, _hospital]),
      builder: (BuildContext context, AsyncSnapshot<List<dynamic>> snapshot) {
        if (snapshot.hasError) {
          return Scaffold(
            appBar: AppBar(
              title: Text("Error"),
            ),
            body: Center(
              child: Text("Error : ${snapshot.error}"),
            ),
          );
        } else if (snapshot.connectionState == ConnectionState.done ||
            snapshot.hasData == true) {
          final List<Department> department = snapshot.data?[0];
          final List<Hospital> hospital = snapshot.data?[1];
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
                          Container(
                            margin: EdgeInsets.all(16),
                            padding: EdgeInsets.symmetric(
                                horizontal: 12, vertical: 4),
                            decoration: BoxDecoration(
                              borderRadius: BorderRadius.circular(12),
                              border: Border.all(color: Colors.black, width: 4),
                            ),
                            child: DropdownButtonHideUnderline(
                              child: DropdownButton<String>(
                                  value: valueDepartment,
                                  isExpanded: true,
                                  items: department
                                      .map(
                                        (value) => DropdownMenuItem(
                                          value: value.id.toString(),
                                          child: Text(
                                            value.name.toString(),
                                            style: TextStyle(
                                                fontWeight: FontWeight.bold,
                                                fontSize: 20),
                                          ),
                                        ),
                                      )
                                      .toList(),
                                  onChanged: (valueDepartment) => setState(() {
                                        this.valueDepartment = valueDepartment;
                                        user.department =
                                            int.parse(valueDepartment!);
                                      })),
                            ),
                          ),
                          Container(
                            margin: EdgeInsets.all(16),
                            padding: EdgeInsets.symmetric(
                                horizontal: 12, vertical: 4),
                            decoration: BoxDecoration(
                              borderRadius: BorderRadius.circular(12),
                              border: Border.all(color: Colors.black, width: 4),
                            ),
                            child: DropdownButtonHideUnderline(
                              child: DropdownButton<String>(
                                  value: valueHospital,
                                  isExpanded: true,
                                  items: hospital
                                      .map(
                                        (value) => DropdownMenuItem(
                                          value: value.id.toString(),
                                          child: Text(
                                            value.name.toString(),
                                            style: TextStyle(
                                                fontWeight: FontWeight.bold,
                                                fontSize: 20),
                                          ),
                                        ),
                                      )
                                      .toList(),
                                  onChanged: (valueHospital) => setState(() {
                                        this.valueHospital = valueHospital;
                                        user.hospital =
                                            int.parse(valueHospital!);
                                      })),
                            ),
                          ),
                          Text("ใบอนุญาต Certification",
                              style: TextStyle(fontSize: 20)),
                          SizedBox(
                            height: 20,
                          ),
                          Text("เลขวิชาชีพแพทย์",
                              style: TextStyle(fontSize: 20)),
                          TextFormField(
                            validator: MultiValidator([
                              RequiredValidator(
                                  errorText: "กรุณาป้อนเลขวิชาชีพแพทย์"),
                            ]),
                            onSaved: (Code) {
                              certification.code = Code!;
                            },
                          ),
                          Text(
                              "ไฟล์ประกาศนียบัตรวิชาชีพแพทย์(กรุณาส่งมาเป็น Link)",
                              style: TextStyle(fontSize: 20)),
                          TextFormField(
                            validator: MultiValidator([
                              RequiredValidator(
                                  errorText: "กรุณาป้อนเลขวิชาชีพแพทย์"),
                            ]),
                            onSaved: (diloma) {
                              certification.diloma = diloma!;
                            },
                          ),
                          Text("วันที่ออกใบประกาศณียบัตร",
                              style: TextStyle(fontSize: 20)),
                          ElevatedButton(
                            child: Text(_formatdateOfIssuing.toString()),
                            onPressed: () {
                              showDatePicker(
                                      context: context,
                                      initialDate:
                                          certification.dateOfIssuing ??
                                              _selectdateOfIssuing,
                                      firstDate: DateTime(1990),
                                      lastDate: DateTime(2025))
                                  .then((dateOfIssuing) {
                                setState(() {
                                  certification.dateOfIssuing = dateOfIssuing!;
                                });
                              });
                            },
                          ),
                          Text("วันหมดอายุใบประกาศณียบัตร",
                              style: TextStyle(fontSize: 20)),
                          ElevatedButton(
                            child: Text(_formatdateOfExp.toString()),
                            onPressed: () {
                              showDatePicker(
                                      context: context,
                                      initialDate: certification.dateOfExp ??
                                          _selectdateOfExp,
                                      firstDate: DateTime(1990),
                                      lastDate: DateTime(2025))
                                  .then((dateOfExp) {
                                setState(() {
                                  certification.dateOfExp = dateOfExp!;
                                });
                              });
                            },
                          ),
                          Text(
                              "ใคร หน่วยงานหรือองค์กรไหนเป็นผู้ออกใบประกาศณียบัตร",
                              style: TextStyle(fontSize: 20)),
                          TextFormField(
                            validator: MultiValidator([
                              RequiredValidator(
                                  errorText: "กรุณาป้อนเลขวิชาชีพแพทย์"),
                            ]),
                            onSaved: (issuer) {
                              certification.issuer = issuer!;
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
                                    await createUser(user, certification)
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
        } else {
          return Scaffold(
            body: Center(
              child: CircularProgressIndicator(),
            ),
          );
        }
      },
    );
  }
}
