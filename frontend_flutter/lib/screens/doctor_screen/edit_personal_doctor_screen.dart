import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/models/pInfo.dart';
import 'package:frontend_flutter/screens/doctor_screen/main_doctor_screen.dart';
import 'package:frontend_flutter/screens/user_screen/main_user_screen.dart';
import 'package:frontend_flutter/util/http_exception.dart';
import 'package:frontend_flutter/widget/addbar_edit_personal_doctor_screen.dart';
import 'package:intl/intl.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

class EditPersonalDoctorScreen extends StatefulWidget {
  const EditPersonalDoctorScreen({Key? key}) : super(key: key);

  @override
  _EditPersonalDoctorScreenState createState() =>
      _EditPersonalDoctorScreenState();
}

class _EditPersonalDoctorScreenState extends State<EditPersonalDoctorScreen> {
  PageController pageController = PageController();
  final formkey = GlobalKey<FormState>();
  late Future<bool> _futurePInfos;
  bool _hasPInfo = false;
  String _name = 'FristName LastName';
  String _authTokens = 'No Token';
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

  Future<void> getToken() async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    final String? authToken = sharedPreferences.getString('authToken');
    //authToken?.substring(0, authToken.length - 1);
    setState(() => _authTokens = authToken!);
  }

  Future<String?> getName() async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    final String? name = sharedPreferences.getString('username');
    setState(() => _name = name!);
  }

  @override
  initState() {
    getName();

    _futurePInfos = _futurePInfo();
    super.initState();
  }

  Future<void> getUserId() async {
    //FutureBuilder(future: getToken(), builder: (BuildContext context, AsyncSnapshot<void> snapshot) { },);
    await getToken();
    var response = await http
        .get(Uri.parse('http://10.0.2.2:8080/api/v1/tokens/' + _authTokens));
    if (response.statusCode == 200) {
      final results = jsonDecode(response.body);
      setState(() => _pinfo.user = results);
    } else {
      setState(() => _pinfo.user = 0);
    }
  }

  Future<bool> _futurePInfo() async {
    await getUserId();
    var response;
    if (_pinfo.user == 0) {
      return false;
    } else {
      response = await http.get(Uri.parse(
          'http://10.0.2.2:8080/api/v1/pinfos/' + _pinfo.user.toString()));
    }

    if (response.statusCode == 404 || response.reasonPhrase == "Not Found") {
      return false;
    }
    if (response.statusCode == 200) {
      final results = pInfoFromJson(response.body);
      _hasPInfo = true;
      _pinfo = PInfo(
          address: results.address,
          bloodGroup: results.bloodGroup,
          brithDate: results.brithDate,
          firstName: results.firstName,
          gender: results.gender,
          idCardNumber: results.idCardNumber,
          lastName: results.lastName,
          user: results.user,
          id: results.id,
          about: results.about,
          prefix: results.prefix,
          profile: results.profile);
    } else {
      return false;
    }
    return true;
  }

  Future<void> createPInfo(PInfo pinfo) async {
    /*String _formatdate =
        new DateFormat('yyyy-MM-ddTHH:mm:ss.mmmuuuZ').format(_pinfo.brithDate ?? _selectdate);*/
    await http.post(
      Uri.parse('http://10.0.2.2:8080/api/v1/pinfos'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        "Id": 0,
        "IdCardNumber": pinfo.idCardNumber,
        "FirstName": pinfo.firstName,
        "LastName": pinfo.lastName,
        "Gender": pinfo.gender,
        "BrithDate": _pinfo.brithDate!.toIso8601String(),
        "BloodGroup": pinfo.bloodGroup,
        "Address": pinfo.address,
        "User": pinfo.user
      }),
    );
  }

  Future<void> deletePInfo() async {
    await http.delete(
      Uri.parse('http://10.0.2.2:8080/api/v1/pinfos/' + _pinfo.id.toString()),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
    );
  }

  Future<void> updatePInfo(PInfo pinfo) async {
    await http.put(
      Uri.parse('http://10.0.2.2:8080/api/v1/pinfos/' + _pinfo.id.toString()),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        "Id": _pinfo.id,
        "IdCardNumber": pinfo.idCardNumber,
        "FirstName": pinfo.firstName,
        "LastName": pinfo.lastName,
        "Gender": pinfo.gender,
        "BrithDate": pinfo.brithDate!.toIso8601String(),
        "BloodGroup": pinfo.bloodGroup,
        "Address": pinfo.address,
        "User": _pinfo.user
      }),
    );
  }

  DateTime _selectdate = new DateTime.now();

  @override
  Widget build(BuildContext context) {
    String _formatdate =
        new DateFormat.yMMMd().format(_pinfo.brithDate ?? _selectdate);
    return FutureBuilder(
        future: _futurePInfos,
        builder: (BuildContext context, AsyncSnapshot<bool> snapshot) {
          if (snapshot.hasError) {
            return Scaffold(
              appBar: AppBar(
                title: Text("Error"),
              ),
              body: Center(
                child: Text(
                    "ขออภัยไม่สามารถเชื่อมต่อกับ Server ได้ในขณะนี้\n\nError : ${snapshot.error}"),
              ),
            );
          }
          if (snapshot.connectionState == ConnectionState.done ||
              snapshot.hasData == true) {
            return Scaffold(
              backgroundColor: Color.fromARGB(255, 208, 244, 255),
              appBar: buildAppBarEditPersonalDoctorScreen(context),
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
                            initialValue: _pinfo.idCardNumber,
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
                                    initialValue: _pinfo.prefix,
                                    validator: MultiValidator([
                                      RequiredValidator(
                                          errorText: "กรุณาป้อนคำนำหน้านาม"),
                                    ]),
                                    decoration: InputDecoration(
                                      labelText: "คำนำหน้านาม",
                                      fillColor: Colors.white,
                                      focusedBorder: OutlineInputBorder(
                                        borderRadius:
                                            BorderRadius.circular(25.0),
                                        borderSide: BorderSide(
                                          color:
                                              Color.fromARGB(228, 96, 239, 220),
                                          width: 3.0,
                                        ),
                                      ),
                                      enabledBorder: OutlineInputBorder(
                                        borderRadius:
                                            BorderRadius.circular(25.0),
                                        borderSide: BorderSide(
                                          color: Color.fromARGB(
                                              235, 111, 137, 162),
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
                                    initialValue: _pinfo.firstName,
                                    validator: MultiValidator([
                                      RequiredValidator(
                                          errorText: "กรุณาป้อนชื่อ"),
                                    ]),
                                    decoration: InputDecoration(
                                      labelText: "ชื่อ",
                                      fillColor: Colors.white,
                                      focusedBorder: OutlineInputBorder(
                                        borderRadius:
                                            BorderRadius.circular(25.0),
                                        borderSide: BorderSide(
                                          color:
                                              Color.fromARGB(228, 96, 239, 220),
                                          width: 3.0,
                                        ),
                                      ),
                                      enabledBorder: OutlineInputBorder(
                                        borderRadius:
                                            BorderRadius.circular(25.0),
                                        borderSide: BorderSide(
                                          color: Color.fromARGB(
                                              235, 111, 137, 162),
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
                                  initialValue: _pinfo.lastName,
                                  validator: MultiValidator([
                                    RequiredValidator(
                                        errorText: "กรุณาป้อนนามสกุล"),
                                  ]),
                                  decoration: InputDecoration(
                                    labelText: "นามสกุล",
                                    fillColor: Colors.white,
                                    focusedBorder: OutlineInputBorder(
                                      borderRadius: BorderRadius.circular(25.0),
                                      borderSide: BorderSide(
                                        color:
                                            Color.fromARGB(228, 96, 239, 220),
                                        width: 3.0,
                                      ),
                                    ),
                                    enabledBorder: OutlineInputBorder(
                                      borderRadius: BorderRadius.circular(25.0),
                                      borderSide: BorderSide(
                                        color:
                                            Color.fromARGB(235, 111, 137, 162),
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
                                      borderRadius:
                                          BorderRadius.circular(32.0)),
                                  maximumSize: Size(100, 40), //////// HERE
                                ),
                                child: Text(_formatdate.toString()),
                                onPressed: () {
                                  showDatePicker(
                                          context: context,
                                          initialDate:
                                              _pinfo.brithDate ?? _selectdate,
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
                            initialValue: _pinfo.about,
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
                                Text("บันทึก", style: TextStyle(fontSize: 20)),
                            onPressed: () async {
                              if (formkey.currentState!.validate()) {
                                try {
                                  if (_hasPInfo == true) {
                                    await updatePInfo(_pinfo).then((value) {
                                      formkey.currentState!.save();
                                      Navigator.pushReplacement(context,
                                          MaterialPageRoute(builder: (context) {
                                        return MainDoctorScreen(
                                          ScreenIndex: 3,
                                        );
                                      }));
                                    });
                                  } else {
                                    await createPInfo(_pinfo).then((value) {
                                      formkey.currentState!.save();
                                      Navigator.pushReplacement(context,
                                          MaterialPageRoute(builder: (context) {
                                        return MainDoctorScreen(
                                          ScreenIndex: 3,
                                        );
                                      }));
                                    });
                                  }
                                } on HttpException catch (error) {
                                  Fluttertoast.showToast(
                                      msg: error.toString(),
                                      gravity: ToastGravity.CENTER);
                                }
                              }
                              formkey.currentState!.reset();
                              setState(() {
                                _hasPInfo = false;
                              });
                            },
                          ),
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
                            child: Text("ลบข้อมูล",
                                style: TextStyle(fontSize: 20)),
                            onPressed: () async {
                              try {
                                await deletePInfo().then((value) {
                                  Navigator.pushReplacement(context,
                                      MaterialPageRoute(builder: (context) {
                                    return MainUserScreen(
                                      ScreenIndex: 4,
                                    );
                                  }));
                                });
                              } on HttpException catch (error) {
                                Fluttertoast.showToast(
                                    msg: error.toString(),
                                    gravity: ToastGravity.CENTER);
                              }
                              setState(() {
                                _hasPInfo = false;
                              });
                            },
                          ),
                        ),
                        SizedBox(
                          height: 15,
                        ),
                      ],
                    ),
                  ),
                ),
              ),
            );
          } else {
            return Scaffold(
              body: Center(
                child: CircularProgressIndicator(),
              ),
            );
          }
        });
  }
}
