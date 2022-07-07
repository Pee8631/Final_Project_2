import 'package:flutter/material.dart';
import 'package:frontend_flutter/page/event_editing_page.dart';
import 'package:frontend_flutter/widget/calendar_appoint_widget.dart';
import 'package:frontend_flutter/widget/calendar_doctor_widget.dart';
import 'package:frontend_flutter/widget/calendar_user_widget.dart';

class ScheduleDoctorScreen extends StatefulWidget {
  final int UserId;
  const ScheduleDoctorScreen({Key? key, required this.UserId})
      : super(key: key);

  @override
  State<ScheduleDoctorScreen> createState() => _ScheduleDoctorScreenState();
}

class _ScheduleDoctorScreenState extends State<ScheduleDoctorScreen> {
  var _selectedPage;

  @override
  Widget build(BuildContext context) {
    List<DropdownMenuItem<String>> _dropDownItem() {
      List<String> ddl = ["ตารางเวลาของคุณหมอ", "ตารางคิวที่จอง"];
      return ddl
          .map((value) => DropdownMenuItem(
                value: value,
                child: Text(value),
              ))
          .toList();
    }

    return Scaffold(
      backgroundColor: Color.fromARGB(255, 208, 244, 255),
      appBar: AppBar(
        backgroundColor: Color.fromARGB(255, 208, 244, 255),
        centerTitle: true,
        title: Row(
          children: [
            Padding(
              padding: const EdgeInsets.all(2.0),
              child: Text('หน้าต่างตารางเวลา:', textAlign: TextAlign.center, style: TextStyle( color: Colors.black, fontSize: 16,)),
            ),
            Padding(
              padding: const EdgeInsets.all(2.0),
              child: DropdownButton(
                value: _selectedPage,
                items: _dropDownItem(),
                onChanged: (value) {
                  setState(() {
                    _selectedPage = value;
                  });
                },
                hint: Text('ตารางเวลาของคุณหมอ'),
              ),
            ),
          ],
        ),
        shadowColor: Colors.transparent,
      ),
      body: builderCalendarWidget(),
      floatingActionButton: FloatingActionButton(
          child: Icon(
            Icons.add,
            color: Colors.white,
          ),
          backgroundColor: Color.fromARGB(230, 96, 239, 220),
          onPressed: () => Navigator.of(context).push(
                MaterialPageRoute(
                    builder: (context) =>
                        EventEditingPage(UserId: widget.UserId)),
              )),
    );
  }

  Widget builderCalendarWidget() {
    if (_selectedPage == "ตารางคิวที่จอง") {
      return CalendarAppointmentWidget();
    } else {
      return CalendarDoctorWidget();
    }
  }
}
