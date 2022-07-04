import 'package:flutter/material.dart';
import 'package:frontend_flutter/screens/doctor_screen/main_doctor_screen.dart';
import 'package:intl/intl.dart';
import 'package:syncfusion_flutter_calendar/calendar.dart';
import 'package:http/http.dart' as http;

class EventViewingPage extends StatelessWidget {
  final Appointment appointment;
  final int? UserId;
  const EventViewingPage({Key? key, required this.appointment, this.UserId})
      : super(key: key);

  Future<void> deleteSchedules() async {
    await http.delete(
      Uri.parse(
          'http://10.0.2.2:8080/api/v1/schedules/' + appointment.id.toString()),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Color.fromARGB(255, 208, 244, 255),
      appBar: AppBar(
        centerTitle: true,
        title: Text(
          'เวลาให้ที่ให้คำปรึกษา',
          style: TextStyle(fontSize: 16),
        ),
        leading: CloseButton(),
        actions: [
          Padding(
              padding: const EdgeInsets.all(8.0),
              child: IconButton(
                  icon: Icon(Icons.delete),
                  onPressed: () async {
                    await deleteSchedules().then((value) =>
                        Navigator.pushReplacement(context, MaterialPageRoute(
                          builder: (context) {
                            return MainDoctorScreen(ScreenIndex: 0);
                          },
                        )));
                  })),
        ],
      ),
      body: ListView(
        padding: EdgeInsets.only(left: 16, right: 16, top: 32, bottom: 10),
        children: <Widget>[
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
            child: buildDateTime(appointment),
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
            child: buildTopic(appointment.subject ?? ''),
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
            child: buildDetail(appointment.notes ?? ''),
          ),
        ],
      ),
    );
  }

  Widget buildDateTime(Appointment appointment) {
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

  Widget buildTopic(String topic) {
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
              topic,
              style: TextStyle(color: Colors.black87, fontSize: 16),
            ),
          ),
        ),
      ],
    );
  }

  Widget buildDetail(String detail) {
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
            detail,
            style: TextStyle(color: Colors.black87, fontSize: 16),
            textAlign: TextAlign.left,
          ),
        ),
      ],
    );
  }
}
