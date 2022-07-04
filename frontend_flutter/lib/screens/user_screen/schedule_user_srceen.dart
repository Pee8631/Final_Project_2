import 'package:flutter/material.dart';
import 'package:frontend_flutter/models/Meeting.dart';
import 'package:frontend_flutter/page/event_editing_page.dart';
import 'package:frontend_flutter/widget/calendar_doctor_widget.dart';
import 'package:frontend_flutter/widget/calendar_user_widget.dart';

class ScheduleUserScreen extends StatefulWidget {
  final int UserId;
  const ScheduleUserScreen({Key? key, required this.UserId}) : super(key: key);

  @override
  State<ScheduleUserScreen> createState() => _ScheduleUserScreenState();
}

class _ScheduleUserScreenState extends State<ScheduleUserScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: CalendarUserWidget(),
      // floatingActionButton: FloatingActionButton(
      //     child: Icon(
      //       Icons.add,
      //       color: Colors.white,
      //     ),
      //     backgroundColor: Colors.grey,
      //     onPressed: () => Navigator.of(context).push(
      //           MaterialPageRoute(
      //               builder: (context) =>
      //                   EventEditingPage(UserId: widget.UserId)),
      //         )),
    );
  }
}
