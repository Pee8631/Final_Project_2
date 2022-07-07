import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:frontend_flutter/models/doctorSchedule.dart';
import 'package:frontend_flutter/page/event_viewing_page.dart';
import 'package:frontend_flutter/widget/calendar_user_widget.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:syncfusion_flutter_calendar/calendar.dart';
import 'package:http/http.dart' as http;

class CalendarDoctorWidget extends StatefulWidget {
  const CalendarDoctorWidget({Key? key}) : super(key: key);

  @override
  State<CalendarDoctorWidget> createState() => _CalendarDoctorWidgetState();
}

late Future<Map<DateTime, List<Appointment>>> _dataCollection;

class _CalendarDoctorWidgetState extends State<CalendarDoctorWidget> {
  final CalendarController _controller = CalendarController();
  late DoctorSchedule _schedule;
  late EventDataSource _calendarDataSource;

  var _selectedPage;
  @override
  void initState() {
    _calendarDataSource = EventDataSource(<Appointment>[]);
    _dataCollection = getAppointments();
    super.initState();
  }

  Future<String> getToken() async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    final String? authToken = sharedPreferences.getString('authToken');
    //authToken?.substring(0, authToken.length - 1);
    return await authToken!;
  }

  Future<int> getUserId() async {
    //FutureBuilder(future: getToken(), builder: (BuildContext context, AsyncSnapshot<void> snapshot) { },);
    var response = await http.get(
        Uri.parse('http://10.0.2.2:8080/api/v1/tokens/' + await getToken()));
    if (response.statusCode == 200) {
      final results = jsonDecode(response.body);
      return await results!;
    } else {
      throw ('User Not Found: ' + response.reasonPhrase!);
    }
  }

  Future<DoctorSchedule> _getSchedule() async {
    int userId = await getUserId();
    final response = await http.get(Uri.parse(
        'http://10.0.2.2:8080/api/v1/schedules/' + userId.toString()));

    if (response.statusCode == 404 || response.reasonPhrase == "Not Found") {
      throw ('Schedule Not Found');
    }
    var results;
    if (response.statusCode == 200) {
      results = doctorScheduleFromJson(response.body);
      setState(() {
        _schedule = results;
      });
    } else {
      throw ('Schedule Not Found');
    }
    return results;
  }

  Future<Map<DateTime, List<Appointment>>> getAppointments() async {
    await _getSchedule();
    var _dataCollection = <DateTime, List<Appointment>>{};
    if (_schedule.edges!.doctorSchedule != null) {
      for (int i = 0; i < _schedule.edges!.doctorSchedule!.length; i++) {
        final Appointment meeting = Appointment(
            id: _schedule.edges!.doctorSchedule![i].id,
            subject: _schedule.edges!.doctorSchedule![i].activity,
            startTime: _schedule.edges!.doctorSchedule![i].startTime!,
            endTime: _schedule.edges!.doctorSchedule![i].endTime!,
            notes: _schedule.edges!.doctorSchedule![i].detail,
            color: Color(0xFF01A1EF),
            isAllDay: false);

        final DateTime rangeStartDate = DateTime(
            _schedule.edges!.doctorSchedule![i].startTime!.year,
            _schedule.edges!.doctorSchedule![i].startTime!.month,
            _schedule.edges!.doctorSchedule![i].startTime!.day);
        if (_dataCollection.containsKey(rangeStartDate)) {
          final List<Appointment> meetings = _dataCollection[rangeStartDate]!;
          meetings.add(meeting);
          _dataCollection[rangeStartDate] = meetings;
        } else {
          _dataCollection[rangeStartDate] = [meeting];
        }
      }
    }
    return _dataCollection;
  }

  @override
  Widget build(BuildContext context) {
    return SfCalendar(
      view: CalendarView.week,
      allowedViews: [
        CalendarView.day,
        CalendarView.week,
      ],
      //viewHeaderStyle: ViewHeaderStyle(backgroundColor: viewHeaderColor),
      dataSource: _calendarDataSource,
      controller: _controller,
      firstDayOfWeek: 7,
      selectionDecoration: BoxDecoration(
        color: Color.fromARGB(235, 190, 242, 229),
        border: Border.all(color: Colors.blueGrey, width: 2),
        borderRadius: const BorderRadius.all(Radius.circular(4)),
        shape: BoxShape.rectangle,
      ),
      monthViewSettings: MonthViewSettings(
          appointmentDisplayMode: MonthAppointmentDisplayMode.appointment),
      initialSelectedDate: DateTime.now(),
      cellBorderColor: Colors.blueGrey,
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
      onTap: (details) {
        if (details.appointments == null) return;
        Appointment event = details.appointments!.first;
        Navigator.of(context).push(MaterialPageRoute(
            builder: ((context) => EventViewingPage(
                appointment: event, UserId: _schedule.id))));
      },
      // onLongPress: (details) async {
      //   //final provider = Provider.of<EventProvider>(context, listen: false);
      //   //provider.setDate(details.date!);
      //   showModalBottomSheet(
      //     context: context,
      //     builder: (context) => TasksWidget(
      //         selectedDate: details.date!,
      //         calendarDataSource: _calendarDataSource,
      //         schedulesId: 1),
      //   );
      // },
    );
    //         } else {
    //           return Scaffold(
    //             body: Center(
    //               child: CircularProgressIndicator(),
    //             ),
    //           );
    //         }
    //       });
  }
}

class EventDataSource extends CalendarDataSource {
  EventDataSource(List<Appointment> appointments) {
    this.appointments = appointments;
  }

  @override
  int getId(int index) {
    return appointments![index].id;
  }

  @override
  DateTime getStartTime(int index) {
    return appointments![index].from;
  }

  @override
  DateTime getEndTime(int index) {
    return appointments![index].to;
  }

  @override
  bool isAllDay(int index) {
    return appointments![index].isAllDay;
  }

  @override
  String? getNotes(int index) {
    return appointments![index].notes;
  }

  @override
  String getSubject(int index) {
    return appointments![index].subject;
  }

  @override
  Color getColor(int index) {
    return appointments![index].color;
  }

  @override
  Future<void> handleLoadMore(DateTime startDate, DateTime endDate) async {
    Map<DateTime, List<Appointment>> dataCollection = await _dataCollection;
    final List<Appointment> meetings = <Appointment>[];
    DateTime appStartDate = startDate;
    DateTime appEndDate = endDate.add(Duration(days: 2));

    while (appStartDate.isBefore(appEndDate)) {
      final List<Appointment>? data = dataCollection[appStartDate];
      if (data == null) {
        appStartDate = appStartDate.add(Duration(days: 1));
        continue;
      }
      for (final Appointment meeting in data) {
        if (appointments!.contains(meeting)) {
          continue;
        }
        meetings.add(meeting);
      }
      appStartDate = appStartDate.add(Duration(days: 1));
    }
    appointments!.addAll(meetings);
    notifyListeners(CalendarDataSourceAction.add, meetings);
  }
}
