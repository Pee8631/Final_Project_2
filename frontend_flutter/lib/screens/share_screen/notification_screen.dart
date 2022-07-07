import 'dart:convert';

import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:frontend_flutter/models/appointment.dart';
import 'package:frontend_flutter/models/notifications.dart';
import 'package:http/http.dart' as http;
import 'package:jiffy/jiffy.dart';
import 'package:shared_preferences/shared_preferences.dart';

class NotificationScreen extends StatefulWidget {
  const NotificationScreen({Key? key}) : super(key: key);

  @override
  State<NotificationScreen> createState() => _NotificationScreenState();
}

class _NotificationScreenState extends State<NotificationScreen> {
  late Future<List<NotificationApi>> _notifications;
  ScrollController _scrollController = ScrollController();
  // List<Doctor> Sender = <Doctor>[];
  List<AppointmentApi> Appointment = <AppointmentApi>[];
  late int UserId;
  late int RoleId;

  @override
  void initState() {
    setState(() {
      _notifications = _getNotifications();
    });
    super.initState();
  }

  Future<void> CheckRoleId() async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    final roleId = sharedPreferences.getInt('role');
    setState(() {
      RoleId = roleId!;
    });
  }

  Future<String> getToken() async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    final String? authToken = sharedPreferences.getString('authToken');
    //authToken?.substring(0, authToken.length - 1);
    return await authToken!;
  }

  Future<void> getUserId() async {
    //FutureBuilder(future: getToken(), builder: (BuildContext context, AsyncSnapshot<void> snapshot) { },);
    var response = await http.get(
        Uri.parse('http://10.0.2.2:8080/api/v1/tokens/' + await getToken()));
    if (response.statusCode == 200) {
      final int results = jsonDecode(response.body);
      setState(() {
        UserId = results;
      });
    } else {
      throw ("Token Not Found : " + response.reasonPhrase!);
    }
  }

  Future<List<NotificationApi>> _getNotifications() async {
    await getUserId();
    var response = await http.get(Uri.parse(
        'http://10.0.2.2:8080/api/v1/notifications/' + UserId.toString()));
    if (response.statusCode == 200) {
      List<NotificationApi> results =
          await notificationApiFromJson(response.body);
      if (results != null) {
        // for (int i = 0; i < results.length; i++) {
        //   var UserData = await _getUserData(results[i].senderId);
        //   Sender.add(UserData);
        // }
        for (int i = 0; i < results.length; i++) {
          var appointment = await _getAppointments(results[i].appointmentId);
          setState(() {
            Appointment.add(appointment);
          });
        }
      }

      return await results;
    } else {
      throw ("Notification Not Found : " + response.reasonPhrase!);
    }
  }

  Future<AppointmentApi> _getAppointments(int appointmentId) async {
    var response = await http.get(Uri.parse(
        'http://10.0.2.2:8080/api/v1/appointments/' +
            appointmentId.toString()));
    if (response.statusCode == 200) {
      final results = appointmentFromJson(response.body);
      return await results;
    } else {
      throw ("Appointment Not Found: " + response.reasonPhrase!);
    }
  }

  // Future<Doctor> _getUserData(int id) async {
  //   var response = await http
  //       .get(Uri.parse('http://10.0.2.2:8080/api/v1/users/' + id.toString()));
  //   if (response.statusCode == 200) {
  //     Doctor results = doctorFromJson(response.body);

  //     return await results;
  //   } else {
  //     throw ("User Not Found : " + response.reasonPhrase!);
  //   }
  // }

  Future<void> updateAppointment(int id, String status) async {
    await http.put(
      Uri.parse('http://10.0.2.2:8080/api/v1/appointments/' + id.toString()),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        "Status": status,
      }),
    );
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: _notifications,
      builder: (BuildContext context,
          AsyncSnapshot<List<NotificationApi>> snapshot) {
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
          List<NotificationApi> noti = snapshot.data!;
          List<AppointmentApi> appoint = Appointment;
          noti = List.from(noti.reversed);
          appoint = List.from(appoint.reversed);
          if (noti.isNotEmpty) {
            return Scaffold(
              backgroundColor: Color.fromARGB(255, 208, 244, 255),
              appBar: AppBar(
                backgroundColor: Colors.transparent,
                elevation: 0,
                title: Center(
                    child: Text(
                  "กล่องข้อความ",
                  style: TextStyle(
                      fontSize: 20,
                      fontWeight: FontWeight.bold,
                      color: Colors.black),
                )),
              ),
              body: ListView.separated(
                padding: const EdgeInsets.all(8.0),
                itemCount: noti.length,
                itemBuilder: (BuildContext context, int index) {
                  //var sender = Sender[index].edges!.userHasPInfo![0];
                  AppointmentApi appointment = new AppointmentApi(
                      id: 0,
                      reasonForAppointment: "",
                      detail: null,
                      startTime: null,
                      endTime: null,
                      status: "",
                      edges: null,
                      doctorId: 0,
                      userId: null);
                  var message = noti[index];
                  if (message.appointmentId == appoint[index]) {
                    appointment = appoint[index];
                  }
                  var time = Jiffy(message.createdDate)
                      .startOf(Units.MILLISECOND)
                      .fromNow();
                  return ListTile(
                    selectedColor: Colors.grey,
                    onTap: () {},
                    subtitle: Column(
                      crossAxisAlignment: CrossAxisAlignment.end,
                      children: [
                        Padding(
                          padding: const EdgeInsets.all(4.0),
                          child: Text(
                            message.message,
                            style: TextStyle(fontSize: 12, color: Colors.black),
                            textAlign: TextAlign.left,
                          ),
                        ),
                        Padding(
                          padding: const EdgeInsets.all(4.0),
                          child: Text(
                            time,
                            style: TextStyle(fontSize: 12, color: Colors.black),
                            textAlign: TextAlign.left,
                          ),
                        ),
                        if (appointment.doctorId == UserId) ...{
                          if (appointment.status == "Waiting") ...{
                            buildbuttonWidget(index),
                            if (appointment.status == "Confirm") ...{
                              Text(
                                "ยืนยันแล้ว",
                                style: TextStyle(fontSize: 16),
                              ),
                              if (appointment.status == "Rejected") ...{
                                Text(
                                  "ปฏเสธแล้ว",
                                  style: TextStyle(fontSize: 16),
                                ),
                              } else
                                ...{},
                            },
                          },
                        },
                      ],
                    ),
                  );
                },
                separatorBuilder: (BuildContext context, int index) =>
                    const Divider(
                  color: Colors.black,
                ),
              ),
            );
          } else {
            return Scaffold(
              backgroundColor: Color.fromARGB(255, 208, 244, 255),
              appBar: AppBar(
                title: Text(
                  "กล่องข้อความ",
                  style: TextStyle(
                      fontSize: 20,
                      fontWeight: FontWeight.bold,
                      color: Colors.black87),
                ),
                centerTitle: true,
                backgroundColor: Colors.transparent,
                shadowColor: Colors.transparent,
              ),
              body: Center(
                child: Text(
                  "คุณยังไม่มีข้อความเข้ามา",
                  style: TextStyle(
                      fontSize: 24,
                      fontWeight: FontWeight.bold,
                      color: Colors.black54),
                ),
              ),
            );
          }
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

  Widget buildbuttonWidget(int index) {
    return Container(
      margin: const EdgeInsets.all(8.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.end,
        children: [
          Padding(
            padding: const EdgeInsets.all(8.0),
            child: ElevatedButton(
                style: ElevatedButton.styleFrom(
                    shape: StadiumBorder(), onPrimary: Colors.white),
                child: Text("ยอมรับ", style: TextStyle(fontSize: 16)),
                onPressed: () {
                  updateAppointment(Appointment[index].id, 'Confirm').then(
                    (value) => Navigator.pop(context),
                  );
                }),
          ),
          Padding(
            padding: const EdgeInsets.all(8.0),
            child: ElevatedButton(
                style: ElevatedButton.styleFrom(
                    shape: StadiumBorder(), onPrimary: Colors.red),
                child: Text("ปฏิเสธ", style: TextStyle(fontSize: 16)),
                onPressed: () async {
                  updateAppointment(Appointment[index].id, 'Reject').then(
                    (value) => Navigator.pop(context),
                  );
                }),
          ),
        ],
      ),
    );
  }
}
