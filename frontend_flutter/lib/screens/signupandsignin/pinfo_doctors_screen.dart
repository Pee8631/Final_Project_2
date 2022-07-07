import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/models/pInfo.dart';
import 'package:frontend_flutter/models/user.dart';
import 'package:frontend_flutter/screens/signupandsignin/certification_doctor_screen.dart';
import 'package:frontend_flutter/screens/signupandsignin/dep_hos_doctor_screen.dart';
import 'package:frontend_flutter/screens/signupandsignin/signup_doctor_screen.dart';
import 'package:frontend_flutter/util/http_exception.dart';
import 'package:frontend_flutter/widget/appbor_logindoctorscreen.dart';
import 'package:http/http.dart' as http;
import 'package:intl/intl.dart';
import 'package:shared_preferences/shared_preferences.dart';

class PInfoDoctorScreen extends StatefulWidget {
  final User user;
  const PInfoDoctorScreen({Key? key, required this.user}) : super(key: key);

  @override
  State<PInfoDoctorScreen> createState() => _PInfoDoctorScreenState();
}

class _PInfoDoctorScreenState extends State<PInfoDoctorScreen> {
  PageController pageController = PageController();
  final formkey = GlobalKey<FormState>();
  late User _user;
  PInfo _pinfo = PInfo(
      address: '',
      bloodGroup: '',
      brithDate: null,
      firstName: '',
      gender: 0,
      idCardNumber: '',
      lastName: '',
      user: 0,
      id: 0,
      about: '',
      prefix: '',
      profile: '');
  @override
  void initState() {
    setState(() {
      _user = widget.user;
    });
    super.initState();
  }

  DateTime _selectdate = new DateTime.now();

  @override
  Widget build(BuildContext context) {
    String _formatdate =
        new DateFormat.yMMMd().format(_pinfo.brithDate ?? _selectdate);
    return Scaffold(
      backgroundColor: Color.fromARGB(255, 208, 244, 255),
      appBar: buildAppBarDoctorSignUp(context, SignUpDoctorScreen()),
      body: Container(
        padding: const EdgeInsets.all(10.0),
        child: Form(
          key: formkey,
          child: SingleChildScrollView(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: <Widget>[
                Container(
                  margin: const EdgeInsets.only(top: 7.5, bottom: 7.5),
                  child: TextFormField(
                    validator: MultiValidator([
                      RequiredValidator(
                          errorText: "กรุณาป้อนเลขบัตรประชาชน 13 หลัก"),
                      MaxLengthValidator(13,
                          errorText: 'เลขบัตรประชาชน 13 หลัก'),
                      MinLengthValidator(13,
                          errorText: 'เลขบัตรประชาชน 13 หลัก')
                    ]),
                    decoration: InputDecoration(
                      labelText: "เลขบัตรประชาชน 13 หลัก",
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
                    onSaved: (IDCardNumber) {
                      _pinfo.idCardNumber = IDCardNumber!;
                    },
                  ),
                ),
                Row(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Expanded(
                      flex: 1,
                      child: Container(
                        margin: const EdgeInsets.only(
                            top: 7.5, bottom: 7.5, right: 1),
                        child: TextFormField(
                            validator: MultiValidator([
                              RequiredValidator(
                                  errorText: "กรุณาป้อนคำนำหน้านาม"),
                            ]),
                            decoration: InputDecoration(
                              labelText: "คำนำหน้านาม",
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
                            onChanged: (Prefix) {
                              _pinfo.prefix = Prefix;
                            }),
                      ),
                    ),
                    Expanded(
                      flex: 2,
                      child: Container(
                        margin: const EdgeInsets.only(
                            left: 1, top: 7.5, bottom: 7.5, right: 1),
                        child: TextFormField(
                            validator: MultiValidator([
                              RequiredValidator(errorText: "กรุณาป้อนชื่อ"),
                            ]),
                            decoration: InputDecoration(
                              labelText: "ชื่อ",
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
                            onChanged: (Firstname) {
                              _pinfo.firstName = Firstname;
                            }),
                      ),
                    ),
                    Expanded(
                      flex: 2,
                      child: Container(
                        margin: const EdgeInsets.only(
                          left: 1,
                          top: 7.5,
                          bottom: 7.5,
                        ),
                        child: TextFormField(
                          validator: MultiValidator([
                            RequiredValidator(errorText: "กรุณาป้อนนามสกุล"),
                          ]),
                          decoration: InputDecoration(
                            labelText: "นามสกุล",
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
                          onChanged: (Lastname) {
                            _pinfo.lastName = Lastname;
                          },
                        ),
                      ),
                    ),
                  ],
                ),
                Row(
                  crossAxisAlignment: CrossAxisAlignment.center,
                  children: [
                    Expanded(
                      flex: 1,
                      child: Text(
                        "วันเดือนปีเกิด",
                        style: TextStyle(fontSize: 16),
                      ),
                    ),
                    Expanded(
                      flex: 1,
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
                        child: Text(_formatdate.toString()),
                        onPressed: () {
                          showDatePicker(
                                  context: context,
                                  initialDate: _pinfo.brithDate ?? _selectdate,
                                  firstDate: DateTime(1990),
                                  lastDate: DateTime(2025))
                              .then((date) {
                            setState(() {
                              _pinfo.brithDate = date!;
                            });
                          });
                        },
                      ),
                    ),
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                      flex: 1,
                      child: Text(
                        "เพศ",
                        style: TextStyle(fontSize: 16),
                      ),
                    ),
                    Expanded(
                      flex: 2,
                      child: ListTile(
                        title: const Text('ชาย'),
                        leading: Radio<int>(
                          value: 1,
                          groupValue: _pinfo.gender,
                          onChanged: (int? value) {
                            setState(() {
                              _pinfo.gender = value!;
                            });
                          },
                        ),
                      ),
                    ),
                    Expanded(
                      flex: 2,
                      child: ListTile(
                        title: const Text('ผู้หญิง'),
                        leading: Radio<int>(
                          value: 2,
                          groupValue: _pinfo.gender,
                          onChanged: (int? value) {
                            setState(() {
                              _pinfo.gender = value!;
                            });
                          },
                        ),
                      ),
                    ),
                  ],
                ),
                Text(
                  "กรุ๊ปเลือด",
                  style: TextStyle(fontSize: 16),
                ),
                Row(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  mainAxisAlignment: MainAxisAlignment.start,
                  children: [
                    Expanded(
                      flex: 1,
                      child: ListTile(
                        title: const Text('A'),
                        leading: Radio<String>(
                          value: 'A',
                          groupValue: _pinfo.bloodGroup,
                          onChanged: (String? value) {
                            setState(() {
                              _pinfo.bloodGroup = value!;
                            });
                          },
                        ),
                      ),
                    ),
                    Expanded(
                      flex: 1,
                      child: ListTile(
                        title: const Text('B'),
                        leading: Radio<String>(
                          value: 'B',
                          groupValue: _pinfo.bloodGroup,
                          onChanged: (String? value) {
                            setState(() {
                              _pinfo.bloodGroup = value!;
                            });
                          },
                        ),
                      ),
                    ),
                    Expanded(
                      flex: 1,
                      child: ListTile(
                        title: Text(
                          'AB',
                          style: TextStyle(fontSize: 16),
                        ),
                        leading: Radio<String>(
                          value: 'AB',
                          groupValue: _pinfo.bloodGroup,
                          onChanged: (String? value) {
                            setState(() {
                              _pinfo.bloodGroup = value!;
                            });
                          },
                        ),
                      ),
                    ),
                    Expanded(
                      flex: 1,
                      child: ListTile(
                        title: const Text('O'),
                        leading: Radio<String>(
                          value: 'O',
                          groupValue: _pinfo.bloodGroup,
                          onChanged: (String? value) {
                            setState(() {
                              _pinfo.bloodGroup = value!;
                            });
                          },
                        ),
                      ),
                    ),
                  ],
                ),
                Container(
                  margin: const EdgeInsets.only(top: 5, bottom: 5),
                  child: TextFormField(
                    initialValue: _pinfo.address,
                    keyboardType: TextInputType.multiline,
                    maxLines: null,
                    decoration: InputDecoration(
                      labelText:
                          'ที่อยู่ เช่น บ้านเลขที่ หมู่ ตำบล อำเภอ จังหวัด รหัสไปรษณีย์',
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
                    onChanged: (Address) {
                      _pinfo.address = Address;
                    },
                  ),
                ),
                Container(
                  margin: const EdgeInsets.only(top: 5, bottom: 5),
                  child: TextFormField(
                    keyboardType: TextInputType.multiline,
                    maxLines: null,
                    initialValue: _pinfo.address,
                    decoration: InputDecoration(
                      labelText: 'คุณอยากอธิบายอะไรที่เกี่ยวกับตัวเอง',
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
                    onChanged: (About) {
                      _pinfo.about = About;
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
                    child: Text("ต่อไป", style: TextStyle(fontSize: 20)),
                    onPressed: () async {
                      if (formkey.currentState!.validate()) {
                        formkey.currentState!.save();
                        setState(() {
                          _user.pInfo = _pinfo;
                        });
                        try {
                          Navigator.pushReplacement(context,
                              MaterialPageRoute(builder: (context) {
                            return DepHosDoctorScreen(
                              user: widget.user,
                            );
                          }));
                        } on HttpException catch (error) {
                          Fluttertoast.showToast(
                              msg: error.toString(),
                              gravity: ToastGravity.CENTER);
                        }
                      }
                      formkey.currentState!.reset();
                    },
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
