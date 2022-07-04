import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:frontend_flutter/page/event_viewing_page.dart';
import 'package:frontend_flutter/widget/calendar_doctor_widget.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';
import 'package:syncfusion_flutter_calendar/calendar.dart';
import 'package:syncfusion_flutter_core/theme.dart';

class TasksWidget extends StatefulWidget {
  final DateTime selectedDate;
  final int schedulesId;
  final EventDataSource calendarDataSource;
  const TasksWidget(
      {Key? key,
      required this.schedulesId,
      required this.selectedDate,
      required this.calendarDataSource})
      : super(key: key);

  @override
  State<TasksWidget> createState() => _TasksWidgetState();
}

class _TasksWidgetState extends State<TasksWidget> {
  late int _UserId;
  @override
  void initState() {
    getUserId();
    super.initState();
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
      final results = jsonDecode(response.body);
      setState(() => _UserId = results!);
    } else {
      throw ("User Not Found : " + response.reasonPhrase!);
    }
  }

  @override
  Widget build(BuildContext context) {
    // int id = 0;
    // for (int i = 0; i < widget.doctorSchedules.edges!.doctorSchedule!.length; i++) {
    //   if (widget.events[i].title == widget.doctorSchedules.edges!.doctorSchedule![i].activity &&
    //       widget.events[i].decscription == widget.doctorSchedules.edges!.doctorSchedule![i].detail &&
    //       widget.events[i].from == widget.doctorSchedules.edges!.doctorSchedule![i].startTime &&
    //       widget.events[i].to == widget.doctorSchedules.edges!.doctorSchedule![i].endTime) {
    //         setState(() {
    //           id = widget.doctorSchedules.edges!.doctorSchedule![i].id;
    //         });
    //       }
    //   //provider.addEvent(event);
    // }
    ///final provider = Provider.of<EventProvider>(context);
    //final events = Provider.of<EventProvider>(context).events;
    //final selectedEvents = provider.eventsOfSelectedDate;
    // if (widget.events.isEmpty) {
    //   return Center(
    //     child: Text(
    //       'No Events found',
    //       style: TextStyle(color: Colors.black, fontSize: 24),
    //     ),
    //   );
    // }
    // for(int i = 0; i < widget.calendarDataSource.appointments!.length; i++) {
    //   widget.calendarDataSource.appointments![i]
    // }
    return SfCalendarTheme(
      data: SfCalendarThemeData(
        timeTextStyle: TextStyle(color: Colors.black, fontSize: 16),
      ),
      child: SfCalendar(
        view: CalendarView.timelineDay,
        dataSource: widget
            .calendarDataSource, //EventDataSource(new Appointment(startTime: startTime, endTime: endTime)),
        initialDisplayDate: widget.selectedDate,
        appointmentBuilder: appointmentBuilder,
        headerHeight: 0,
        todayHighlightColor: Colors.grey,
        loadMoreWidgetBuilder:
            (BuildContext context, LoadMoreCallback loadMoreAppointments) {
          return FutureBuilder(
            future: loadMoreAppointments(),
            builder: (context, snapShot) {
              return Container(
                alignment: Alignment.center,
                child: CircularProgressIndicator(
                  valueColor: AlwaysStoppedAnimation(Colors.blue),
                ),
              );
            },
          );
        },
        selectionDecoration: BoxDecoration(color: Colors.grey.withOpacity(0.7)),
        onTap: (details) {
          if (details.appointments == null) return;
          final event = details.appointments!.first;
          Navigator.of(context).push(MaterialPageRoute(
              builder: ((context) => EventViewingPage(
                appointment: event, UserId: _UserId))));
        },
      ),
    );
  }

  Widget appointmentBuilder(
    BuildContext context,
    CalendarAppointmentDetails details
  ) {
    var widgets;
    details.appointments.forEach((item) => {
      widgets[item] =  Container(
        width: details.bounds.width,
        height: details.bounds.height,
        decoration: BoxDecoration(
            color: item.color.withOpacity(0.5),
            borderRadius: BorderRadius.circular(12)),
        child: Center(
          child: Text(
            item.subject,
            maxLines: 2,
            overflow: TextOverflow.ellipsis,
            style: TextStyle(
                color: Colors.black, fontSize: 16, fontWeight: FontWeight.bold),
          ),
        ),
      ),
    });

    return widgets;
  }
}
