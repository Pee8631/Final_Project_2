import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:frontend_flutter/models/appointment.dart';
import 'package:frontend_flutter/page/event_editing_page.dart';
import 'package:frontend_flutter/screens/user_screen/main_user_screen.dart';
import 'package:frontend_flutter/util/http_exception.dart';
import 'package:intl/intl.dart';
import 'package:syncfusion_flutter_calendar/calendar.dart';
import 'package:http/http.dart' as http;
import '../models/pInfo.dart';

class AppointmentViewingPage extends StatefulWidget {
  final Appointment appointment;
  final int? UserId;
  const AppointmentViewingPage(
      {Key? key, required this.appointment, this.UserId})
      : super(key: key);

  @override
  State<AppointmentViewingPage> createState() => _AppointmentViewingPageState();
}

class _AppointmentViewingPageState extends State<AppointmentViewingPage> {
  late AppointmentApi _appointment;
  late Future<PInfo> _futurePInfo;
  @override
  void initState() {
    _futurePInfo = futurePInfo();
    super.initState();
  }

  Future<AppointmentApi> getAppointments() async {
    final response = await http.get(Uri.parse(
        'http://10.0.2.2:8080/api/v1/appointments/' +
            widget.appointment.id.toString()));

    if (response.statusCode == 404 || response.reasonPhrase == "Not Found") {
      throw ('Appointment Not Found');
    }
    var results;
    if (response.statusCode == 200) {
      results = appointmentFromJson(response.body);
    } else {
      throw ('Appointment Not Found');
    }
    return await results;
  }

  Future<PInfo> futurePInfo() async {
    AppointmentApi appointments = await getAppointments();
    _appointment = appointments;
    var response = await http.get(Uri.parse(
        'http://10.0.2.2:8080/api/v1/pinfos/' +
            _appointment.doctorId.toString()));

    if (response.statusCode == 200) {
      final results = pInfoFromJson(response.body);
      return results;
    } else {
      throw ('PInfo Not Found');
    }
  }

  Future<PInfo> _getPInfo(int DoctorId) async {
    var response = await http.get(
        Uri.parse('http://10.0.2.2:8080/api/v1/pinfos/' + DoctorId.toString()));

    if (response.statusCode == 200) {
      var results = pInfoFromJson(response.body);
      return await results;
    } else {
      throw ('PInfo Not Found');
    }
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: _futurePInfo,
      builder: (BuildContext context, AsyncSnapshot<PInfo> snapshot) {
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
          PInfo pInfo = snapshot.data!;
          return Scaffold(
            backgroundColor: Color.fromARGB(255, 208, 244, 255),
            appBar: AppBar(
              backgroundColor: Color.fromARGB(232, 100, 180, 255),
              centerTitle: true,
              title: Text(
                'นัดหมาย',
                style: TextStyle(fontSize: 16),
              ),
              leading: CloseButton(),
            ),
            body: ListView(
              padding:
                  EdgeInsets.only(left: 16, right: 16, top: 32, bottom: 10),
              children: <Widget>[
                Container(
                  alignment: Alignment.center,
                  height: 40,
                  margin: EdgeInsets.only(top: 5, bottom: 5),
                  padding: EdgeInsets.only(
                    left: 10,
                    right: 10,
                  ),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(10.0),
                    color: Colors.white,
                  ),
                  child: buildDoctorData(pInfo),
                ),
                Container(
                  alignment: Alignment.center,
                  height: 40,
                  margin: EdgeInsets.only(top: 5, bottom: 5),
                  padding: EdgeInsets.only(
                    left: 10,
                    right: 10,
                  ),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(10.0),
                    color: Colors.white,
                  ),
                  child: buildStatus(_appointment.status),
                ),
                Container(
                  alignment: Alignment.center,
                  height: 110,
                  margin: EdgeInsets.only(top: 5, bottom: 5),
                  padding: EdgeInsets.only(
                    left: 10,
                    right: 10,
                  ),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(10.0),
                    color: Colors.white,
                  ),
                  child: buildDateTime(_appointment),
                ),
                Container(
                  alignment: Alignment.center,
                  height: 50,
                  margin: EdgeInsets.only(top: 5, bottom: 5),
                  padding: EdgeInsets.only(
                    left: 10,
                    right: 10,
                  ),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(10.0),
                    color: Colors.white,
                  ),
                  child: buildReason(_appointment.reasonForAppointment ?? ''),
                ),
                Container(
                  alignment: Alignment.topLeft,
                  height: 200,
                  margin: EdgeInsets.only(top: 5, bottom: 5),
                  padding: EdgeInsets.only(
                    left: 10,
                    right: 10,
                  ),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(10.0),
                    color: Colors.white,
                  ),
                  child: buildDetail(_appointment.detail ?? ''),
                ),
              ],
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

  buildDoctorData(PInfo pInfo) {
    return Row(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.end,
        children: [
          Expanded(
            flex: 1,
            child: Text(
              "คุณหมอ: ",
              textAlign: TextAlign.justify,
              style: TextStyle(fontSize: 16),
            ),
          ),
          Expanded(
            flex: 2,
            child: Text(
              pInfo.firstName + " " + pInfo.lastName,
              textAlign: TextAlign.left,
              style: TextStyle(fontSize: 16),
            ),
          ),
        ]);
  }

  buildStatus(String status) {
    String StatusInThai;
    Color StatusColor;
    if (status == "Confirm"){
      StatusInThai = "ยืนยันแล้ว";
      StatusColor = Colors.greenAccent;
    } else if (status == "Waiting") {
      StatusInThai = "กำลังรอการยืนยัน";
      StatusColor = Colors.yellowAccent;
    } else if (status == "Reject") {
      StatusInThai = "ปฏิเสธ";
      StatusColor = Colors.redAccent;
    } else {
      StatusInThai = "ยังไม่มีการจอง";
      StatusColor = Colors.grey;
    }
    return Row(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.end,
        children: [
          Expanded(
            flex: 1,
            child: Text(
              "สถานะนัดหมาย: ",
              textAlign: TextAlign.justify,
              style: TextStyle(fontSize: 16),
            ),
          ),
          Expanded(
            flex: 2,
            child: Text(
              StatusInThai,
              textAlign: TextAlign.left,
              style: TextStyle(fontSize: 16, color: StatusColor),
            ),
          ),
        ]);
  }

  Widget buildDateTime(AppointmentApi? appointment) {
    return Column(
      children: [
        Padding(
          padding: const EdgeInsets.only(top: 4.0, bottom: 4.0),
          child: buildDate('จนถึง:', appointment!.startTime!),
        ),
        Padding(
          padding: const EdgeInsets.only(top: 4.0, bottom: 4.0),
          child: buildDate('ตั้งแต่:', appointment.endTime!),
        )
      ],
    );
  }

  Widget buildDate(String Title, DateTime date) {
    return Row(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Expanded(
            flex: 1,
            child: Text(
              Title,
              textAlign: TextAlign.left,
              style: TextStyle(fontSize: 16),
            ),
          ),
          Expanded(
            flex: 2,
            child: Text(
              DateFormat.yMMMMEEEEd().add_Hm().format(date),
              textAlign: TextAlign.left,
              style: TextStyle(fontSize: 16),
            ),
          ),
        ]);
  }

  Widget buildReason(String reason) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Expanded(
          flex: 1,
          child: Padding(
            padding: const EdgeInsets.only(top: 4.0, bottom: 4.0),
            child: Text(
              'เรื่อง:',
              style: TextStyle(color: Colors.black87, fontSize: 16),
            ),
          ),
        ),
        Expanded(
          flex: 2,
          child: Padding(
            padding: const EdgeInsets.only(top: 4.0, bottom: 4.0),
            child: Text(
              reason,
              style: TextStyle(color: Colors.black87, fontSize: 16),
            ),
          ),
        ),
      ],
    );
  }

  Widget buildDetail(String Detail) {
    return Column(
      children: [
        Padding(
          padding: const EdgeInsets.only(top: 4.0, bottom: 4.0),
          child: Text(
            'รายละเอียด:',
            style: TextStyle(
              color: Colors.black87,
              fontSize: 16,
            ),
            textAlign: TextAlign.left,
          ),
        ),
        Padding(
          padding: const EdgeInsets.only(top: 10.0, bottom: 10.0),
          child: Text(
            Detail,
            style: TextStyle(color: Colors.black87, fontSize: 16),
            textAlign: TextAlign.left,
          ),
        ),
      ],
    );
  }
}
