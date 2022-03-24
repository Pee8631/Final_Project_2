import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/model/data.dart';
import 'package:frontend_flutter/screens/main_screen.dart';
import 'package:frontend_flutter/util/http_exception.dart';
import 'package:frontend_flutter/widget/appbar_mainscreen.dart';
import 'package:intl/intl.dart';
import 'package:http/http.dart' as http;

class EditPersonalScreen extends StatefulWidget {
  const EditPersonalScreen({Key? key}) : super(key: key);

  @override
  _EditPersonalScreenState createState() => _EditPersonalScreenState();
}

class _EditPersonalScreenState extends State<EditPersonalScreen> {
  PageController pageController = PageController();
  final formkey = GlobalKey<FormState>();
  bool _hasData = false;
  Data _data = Data(
      address: '',
      bloodGroup: '',
      brithDate: null,
      firstName: '',
      gender: null,
      idCardNumber: '',
      lastName: '',
      user: null);

  Future<Data?> _futureData() async {
    final url = Uri.parse('http://10.0.2.2:8080/api/v1/datas/4');
    final response = await http.get(url);
    print(response.statusCode);
    print(response.body);
    final results = dataFromJson(response.body);
    print(results.brithDate);
    if (results.address != null ||
        results.idCardNumber != null ||
        results.firstName != null ||
        results.lastName != null ||
        results.gender != null ||
        results.brithDate != null ||
        results.bloodGroup != null) {
        _hasData = true;
        _data = Data(
            address: results.address,
            bloodGroup: results.bloodGroup,
            brithDate: results.brithDate,
            firstName: results.firstName,
            gender: results.gender,
            idCardNumber: results.idCardNumber,
            lastName: results.lastName,
            user: results.user);
    }
    return _data;
  }

  Future<void> createData(Data data) async {
    /*String _formatdate =
        new DateFormat('yyyy-MM-ddTHH:mm:ss.mmmuuuZ').format(_data.brithDate ?? _selectdate);*/
    await http.post(
      Uri.parse('http://10.0.2.2:8080/api/v1/datas'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        "IdCardNumber": data.idCardNumber,
        "FirstName": data.firstName,
        "LastName": data.lastName,
        "Gender": data.gender,
        "BrithDate": _data.brithDate!.toIso8601String(),
        "BloodGroup": data.bloodGroup,
        "Address": data.address,
        "User": 1
      }),
    );
  }

  Future<void> deleteData() async {
    await http.delete(
      Uri.parse('http://10.0.2.2:8080/api/v1/datas/4'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
    );
  }

  Future<void> updateData(Data data) async {
    await http.put(
      Uri.parse('http://10.0.2.2:8080/api/v1/datas/4'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        "IdCardNumber": data.idCardNumber,
        "FirstName": data.firstName,
        "LastName": data.lastName,
        "Gender": data.gender,
        "BrithDate": data.brithDate!.toIso8601String(),
        "BloodGroup": data.bloodGroup,
        "Address": data.address,
        "User": data.user
      }),
    );
  }

  DateTime _selectdate = new DateTime.now();

  @override
  Widget build(BuildContext context) {
    String _formatdate =
        new DateFormat.yMMMd().format(_data.brithDate ?? _selectdate);
    return FutureBuilder(
        future: _futureData(),
        builder: (BuildContext context, AsyncSnapshot<dynamic> snapshot) {
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
          if (snapshot.connectionState == ConnectionState.done) {
            return Scaffold(
              appBar: buildAppBarMain(context),
              body: Container(
                padding: const EdgeInsets.all(10.0),
                child: Form(
                  key: formkey,
                  child: SingleChildScrollView(
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: <Widget>[
                        SizedBox(
                          height: 15,
                        ),
                        Text(
                          "เลขบัตรประชาชน",
                          style: TextStyle(fontSize: 20),
                        ),
                        SizedBox(
                          height: 5,
                        ),
                        TextFormField(
                          initialValue: _data.idCardNumber ?? '',
                          decoration: InputDecoration(
                            border: OutlineInputBorder(),
                            labelText: 'ตัวอย่าง 1-2345-67890-12-3',
                          ),
                          validator: MultiValidator([
                            RequiredValidator(
                                errorText: "กรุณาป้อนเลขบัตรประชาชน"),
                          ]),
                          /*onSaved: (IDCardNumber) {
                            _data.idCardNumber = IDCardNumber;
                          },*/
                          onChanged: (IDCardNumber) {
                            _data.idCardNumber = IDCardNumber;
                          },
                        ),
                        SizedBox(
                          height: 15,
                        ),
                        Text(
                          "ชื่อ",
                          style: TextStyle(fontSize: 20),
                        ),
                        SizedBox(
                          height: 5,
                        ),
                        TextFormField(
                            initialValue: _data.firstName ?? '',
                            decoration: InputDecoration(
                              border: OutlineInputBorder(),
                              labelText: 'ชื่อ',
                            ),
                            validator: MultiValidator([
                              RequiredValidator(errorText: "กรุณาป้อนชื่อ"),
                            ]),
                            /*onSaved: (Firstname) {
                              _data.firstName = Firstname;
                            },*/
                            onChanged: (Firstname) {
                              _data.firstName = Firstname;
                            }),
                        SizedBox(
                          height: 15,
                        ),
                        Text(
                          "นามสกุล",
                          style: TextStyle(fontSize: 20),
                        ),
                        SizedBox(
                          height: 5,
                        ),
                        TextFormField(
                          initialValue: _data.lastName ?? '',
                          decoration: InputDecoration(
                            border: OutlineInputBorder(),
                            labelText: 'นามสกุล',
                          ),
                          validator: MultiValidator([
                            RequiredValidator(errorText: "กรุณาป้อนนามสกุล"),
                          ]),
                          /*onSaved: (Lastname) {
                            _data.lastName = Lastname;
                          },*/
                          onChanged: (Lastname) {
                            _data.lastName = Lastname;
                          },
                        ),
                        SizedBox(
                          height: 15,
                        ),
                        Text(
                          "วันเดือนปีเกิด",
                          style: TextStyle(fontSize: 20),
                        ),
                        SizedBox(
                          height: 5,
                        ),
                        ElevatedButton(
                          child: Text(_formatdate.toString()),
                          onPressed: () {
                            showDatePicker(
                                    context: context,
                                    initialDate: _data.brithDate ?? _selectdate,
                                    firstDate: DateTime(1990),
                                    lastDate: DateTime(2025))
                                .then((date) {
                              setState(() {
                              _data.brithDate = date!;
                              });
                            });
                          },
                        ),
                        SizedBox(
                          height: 15,
                        ),
                        Text(
                          "เพศ",
                          style: TextStyle(fontSize: 20),
                        ),
                        SizedBox(
                          height: 5,
                        ),
                        ListTile(
                          title: const Text('Male'),
                          leading: Radio<int>(
                            value: 1,
                            groupValue: _data.gender,
                            onChanged: (int? value) {
                              setState(() {
                                _data.gender = value;
                              });
                            },
                          ),
                        ),
                        ListTile(
                          title: const Text('Female'),
                          leading: Radio<int>(
                            value: 2,
                            groupValue: _data.gender,
                            onChanged: (int? value) {
                              setState(() {
                                _data.gender = value;
                              });
                            },
                          ),
                        ),
                        SizedBox(
                          height: 15,
                        ),
                        Text(
                          "กรุ๊ปเลือด",
                          style: TextStyle(fontSize: 20),
                        ),
                        SizedBox(
                          height: 5,
                        ),
                        ListTile(
                          title: const Text('A'),
                          leading: Radio<String>(
                            value: 'A',
                            groupValue: _data.bloodGroup,
                            onChanged: (String? value) {
                              setState(() {
                                _data.bloodGroup = value;
                              });
                            },
                          ),
                        ),
                        ListTile(
                          title: const Text('B'),
                          leading: Radio<String>(
                            value: 'B',
                            groupValue: _data.bloodGroup,
                            onChanged: (String? value) {
                              setState(() {
                                _data.bloodGroup = value;
                              });
                            },
                          ),
                        ),
                        ListTile(
                          title: const Text('AB'),
                          leading: Radio<String>(
                            value: 'AB',
                            groupValue: _data.bloodGroup,
                            onChanged: (String? value) {
                              setState(() {
                                _data.bloodGroup = value;
                                formkey.currentState!.save();
                              });
                            },
                          ),
                        ),
                        ListTile(
                          title: const Text('O'),
                          leading: Radio<String>(
                            value: 'O',
                            groupValue: _data.bloodGroup,
                            onChanged: (String? value) {
                              setState(() {
                                _data.bloodGroup = value;
                                formkey.currentState!.save();
                              });
                            },
                          ),
                        ),
                        SizedBox(
                          height: 15,
                        ),
                        Text(
                          "ที่อยู่",
                          style: TextStyle(fontSize: 20),
                        ),
                        SizedBox(
                          height: 5,
                        ),
                        TextFormField(
                          initialValue: _data.address ?? '',
                          decoration: InputDecoration(
                            border: OutlineInputBorder(),
                            labelText:
                                'บ้านเลขที่ หมู่ ตำบล อำเภอ จังหวัด รหัสไปรษณีย์',
                          ),
                          validator: MultiValidator([
                            RequiredValidator(errorText: "กรุณาป้อนที่อยู่"),
                          ]),
                          onSaved: (Address) {
                            _data.address = Address;
                            formkey.currentState!.save();
                          },
                        ),
                        SizedBox(
                          height: 15,
                        ),
                        SizedBox(
                          width: double.infinity,
                          child: ElevatedButton(
                            child:
                                Text("บันทึก", style: TextStyle(fontSize: 20)),
                            onPressed: () async {
                              if (formkey.currentState!.validate()) {
                                formkey.currentState!.save();
                                try {
                                  if (_hasData == true) {
                                    await updateData(_data).then((value) {
                                      Navigator.pushReplacement(context,
                                          MaterialPageRoute(builder: (context) {
                                        return MainScreen();
                                      }));
                                    });
                                  } else {
                                    await createData(_data).then((value) {
                                      Navigator.pushReplacement(context,
                                          MaterialPageRoute(builder: (context) {
                                        return MainScreen();
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
                                _hasData = false;
                              });
                            },
                          ),
                        ),
                        SizedBox(
                          width: double.infinity,
                          child: ElevatedButton(
                            child: Text("ลบข้อมูล",
                                style: TextStyle(fontSize: 20)),
                            onPressed: () async {                           
                                try {
                                  await deleteData().then((value) {
                                    Navigator.pushReplacement(context,
                                        MaterialPageRoute(builder: (context) {
                                      return MainScreen();
                                    }));
                                  });
                                } on HttpException catch (error) {
                                  Fluttertoast.showToast(
                                      msg: error.toString(),
                                      gravity: ToastGravity.CENTER);
                                }
                              setState(() {
                                _hasData = false;
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
          }
          return Scaffold(
            body: Center(
              child: CircularProgressIndicator(),
            ),
          );
        });
  }
}
