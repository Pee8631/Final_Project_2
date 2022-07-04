import 'dart:convert';

import 'package:date_picker_timeline/date_picker_timeline.dart';
import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/models/appointment.dart';
import 'package:frontend_flutter/screens/user_screen/doctor_profile_screen.dart';
import 'package:frontend_flutter/screens/user_screen/main_user_screen.dart';
import 'package:frontend_flutter/util/http_exception.dart';
import 'package:frontend_flutter/widget/appbar_doctors_screen.dart';
import 'package:group_button/group_button.dart';
import 'package:http/http.dart' as http;
import 'package:intl/intl.dart';

class AppointmentScreen extends StatefulWidget {
  final String Name;
  final int DoctorId;
  final int UserId;
  const AppointmentScreen(
      {Key? key,
      required this.Name,
      required this.DoctorId,
      required this.UserId})
      : super(key: key);

  @override
  State<AppointmentScreen> createState() => _AppointmentScreenState();
}

class _AppointmentScreenState extends State<AppointmentScreen> {
  DateTime _selectedValue = new DateTime.now();
  List<AppointmentApi> _appointments = <AppointmentApi>[];
  bool _IsSelected = false;
  AppointmentApi selectedAppointments = new AppointmentApi(
      detail: '',
      edges: null,
      endTime: null,
      id: 0,
      reasonForAppointment: '',
      startTime: null,
      status: '', doctorId: 0, userId: null);
  late Future<List<AppointmentApi>> _getAppointments;
  final formkey = GlobalKey<FormState>();

  @override
  initState() {
    _getAppointments = getAppointment();
    super.initState();
  }

  Future<List<AppointmentApi>> getAppointment() async {
    var response = await http.get(Uri.parse(
        'http://10.0.2.2:8080/api/v1/appointments/list/' +
            widget.DoctorId.toString()));
    if (response.statusCode == 200) {
      final results = appointmentsFromJson(response.body);
      return await results;
    } else {
      throw ("PInfo Not Found: " + response.reasonPhrase!);
    }
  }

  Future<void> updateAppointment(AppointmentApi appointment, int UserId) async {
    await http.put(
      Uri.parse('http://10.0.2.2:8080/api/v1/appointments/' +
          appointment.id.toString()),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        "ReasonForAppointment": appointment.reasonForAppointment,
        "Detail": appointment.detail,
        "Status": 'Waiting',
        "StartTime": appointment.startTime!.toIso8601String(),
        "EndTime": appointment.endTime!.toIso8601String(),
        "UserId": UserId,
      }),
    );

  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: _getAppointments,
      builder:
          (BuildContext context, AsyncSnapshot<List<AppointmentApi>> snapshot) {
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
        } else if (snapshot.connectionState == ConnectionState.done) {
          _appointments.clear();
          for (int i = 0; i < snapshot.data!.length; i++) {
            var startTimematter = new DateFormat('yyyy-MM-dd');
            String startDate =
                startTimematter.format(snapshot.data![i].startTime!);
            var selectedValuematter = new DateFormat('yyyy-MM-dd');
            String selectDate = selectedValuematter.format(_selectedValue);
            if (startDate == selectDate && snapshot.data![i].status == 'None') {
              _appointments.add(snapshot.data![i]);
            }
          }
          return Scaffold(
            appBar: buildAppBarBackToScreen(
                context,
                widget.Name,
                DoctorProfileScreen(
                  DoctorId: widget.DoctorId,
                  Name: widget.Name,
                  UserId: widget.UserId,
                )),
            body: Form(
              key: formkey,
              child: SingleChildScrollView(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: <Widget>[
                    buildListDate(snapshot.data!),
                    buildListAppointment(_appointments),
                    buildTextFieldTitle(),
                    buildTextFieldDetail(),
                    buildSaveButton(),
                  ],
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
      },
    );
  }

  Widget buildListDate(List<AppointmentApi> appointments) => Padding(
        padding: const EdgeInsets.all(8.0),
        child: DatePicker(
          DateTime.now(),
          initialSelectedDate: DateTime.now(),
          selectionColor: Colors.black,
          selectedTextColor: Colors.white,
          onDateChange: (date) {
            // New date selected
            setState(() {
              _appointments.clear();
              _selectedValue = date;
              for (int i = 0; i < appointments.length; i++) {
                var startTimematter = new DateFormat('yyyy-MM-dd');
                String startDate =
                    startTimematter.format(appointments[i].startTime!);
                var selectedValuematter = new DateFormat('yyyy-MM-dd');
                String selectDate = selectedValuematter.format(_selectedValue);
                if (startDate == selectDate &&
                    appointments[i].status == 'None') {
                  _appointments.add(appointments[i]);
                }
              }
            });
          },
        ),
      );

  buildListAppointment(List<AppointmentApi> appointments) {
    List<String> DateValue = <String>[];
    var DateValuematter = new DateFormat('kk:mm');
    appointments.forEach((item) {
      var results = DateValuematter.format(item.startTime!);
      DateValue.add(results);
    });

    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: GroupButton(
        isRadio: true,
        onSelected: (text, index, isSelected) {
          setState(() {
            selectedAppointments.id = appointments[index].id;
            selectedAppointments.startTime = appointments[index].startTime;
            selectedAppointments.endTime = appointments[index].endTime;
            _IsSelected = isSelected;
          });
        },
        buttons: DateValue,
      ),
    );
  }

  buildTextFieldTitle() => Padding(
        padding: const EdgeInsets.all(8.0),
        child: TextFormField(
          decoration: InputDecoration(
            border: OutlineInputBorder(),
            labelText: 'หัวข้อ',
          ),
          validator: MultiValidator([
            RequiredValidator(
                errorText:
                    "กรุณาป้อนเรื่องที่จะเข้ารับการตรวจสุขภาพหรือรับการรักษา"),
          ]),
          onChanged: (reason) {
            selectedAppointments.reasonForAppointment = reason;
          },
        ),
      );

  buildTextFieldDetail() => Padding(
        padding: const EdgeInsets.all(8.0),
        child: TextFormField(
          decoration: InputDecoration(
            border: OutlineInputBorder(),
            labelText: 'รายละเอียด',
          ),
          onChanged: (detail) {
            selectedAppointments.detail = detail;
          },
        ),
      );

  buildSaveButton() => Center(
        child: Padding(
          padding: const EdgeInsets.all(8.0),
          child: ElevatedButton(
              style: ElevatedButton.styleFrom(
                  fixedSize: Size(200, 40), onPrimary: Colors.white),
              child: Text('ยืนยันการนัดหมาย'),
              onPressed: () async {
                if (formkey.currentState!.validate()) {
                  try {
                    if (_IsSelected) {
                      await updateAppointment(
                              selectedAppointments, widget.UserId)
                          .then((value) {
                        formkey.currentState!.save();
                        Navigator.pushReplacement(context,
                            MaterialPageRoute(builder: (context) {
                          return MainUserScreen(
                            ScreenIndex: 0,
                          );
                        }));
                      });
                    } else {
                      Fluttertoast.showToast(
                          msg: 'กรุณาวัน เวลาที่ต้องการนัดหมาย',
                          gravity: ToastGravity.CENTER);
                    }
                  } on HttpException catch (error) {
                    Fluttertoast.showToast(
                        msg: error.toString(), gravity: ToastGravity.CENTER);
                  }
                }
              }),
        ),
      );
}
//  var b = appointments[index];
//             
//             return Text(
//               DateValuematter.format(b.startTime!),
//               style: TextStyle(fontSize: 16, color: Colors.black),
//             );