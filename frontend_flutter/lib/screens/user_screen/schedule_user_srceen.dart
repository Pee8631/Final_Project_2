import 'package:flutter/material.dart';
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
    );
  }
}
