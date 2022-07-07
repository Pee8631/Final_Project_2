import 'dart:convert';
import 'package:cell_calendar/cell_calendar.dart';
import 'package:flutter/material.dart';
import 'package:frontend_flutter/models/appointment.dart';
import 'package:frontend_flutter/models/pInfo.dart';
import 'package:frontend_flutter/page/appointment_viewing_page.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:syncfusion_flutter_calendar/calendar.dart';
import 'package:http/http.dart' as http;

class CalendarAppointmentWidget extends StatefulWidget {
  const CalendarAppointmentWidget({Key? key}) : super(key: key);

  @override
  State<CalendarAppointmentWidget> createState() => _CalendarAppointmentWidgetState();
}

late Future<Map<DateTime, List<Appointment>>> _dataCollection;

class _CalendarAppointmentWidgetState extends State<CalendarAppointmentWidget> {
  final CalendarController _controller = CalendarController();
  late List<AppointmentApi> _appointments;
  late int _UserId;
  @override
  void initState() {
    setState(() {
      _calendarDataSource = EventDataSource(<Appointment>[]);
      _dataCollection = getAppointments();
    });
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

  Future<List<AppointmentApi>> _getAppointments() async {
    int userId = await getUserId();
    _UserId = userId;
    final response = await http.get(Uri.parse(
        'http://10.0.2.2:8080/api/v1/appointments/list/' + userId.toString()));

    if (response.statusCode == 404 || response.reasonPhrase == "Not Found") {
      throw ('Appointment Not Found');
    }
    var results;
    if (response.statusCode == 200) {
      results = appointmentsFromJson(response.body);
    } else {
      throw ('Appointment Not Found');
    }
    return await results;
  }

  Future<Map<DateTime, List<Appointment>>> getAppointments() async {
    List<AppointmentApi> appointments = await _getAppointments();

    _appointments = appointments;

    var _dataCollection = <DateTime, List<Appointment>>{};
    if (_appointments != null) {
      Color color =Colors.grey;
      for (int i = 0; i < _appointments.length; i++) {
        if( _appointments[i].status == "Confirm" ) {
          color = Colors.greenAccent;
        } else if ( _appointments[i].status == "Waiting") {
          color = Colors.yellowAccent;
        } else if ( _appointments[i].status == "Reject") {
          color = Colors.redAccent;
        } else {
          color = Colors.grey;
        }
        final Appointment meeting = Appointment(
            id: _appointments[i].id,
            subject: _appointments[i].reasonForAppointment ?? "",
            startTime: _appointments[i].startTime!,
            endTime: _appointments[i].endTime!,
            notes: _appointments[i].detail ?? "",
            color:  color,
            isAllDay: false);

        final DateTime rangeStartDate = DateTime(
            _appointments[i].startTime!.year,
            _appointments[i].startTime!.month,
            _appointments[i].startTime!.day);
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

  late EventDataSource _calendarDataSource;
  @override
  Widget build(BuildContext context) {
    return SfCalendar(
      view: CalendarView.week,
      allowedViews: [
        CalendarView.day,
        CalendarView.week,
      ],
      selectionDecoration: BoxDecoration(
        color: Color.fromARGB(235, 190, 242, 229),
        border: Border.all(color: Colors.blueGrey, width: 2),
        borderRadius: const BorderRadius.all(Radius.circular(4)),
        shape: BoxShape.rectangle,
      ),
      cellBorderColor: Colors.blueGrey,
      dataSource: _calendarDataSource,
      controller: _controller,
      firstDayOfWeek: 7,
      initialSelectedDate: DateTime.now(),
      loadMoreWidgetBuilder:
          (BuildContext context, LoadMoreCallback loadMoreAppointments) {
        return FutureBuilder(
          future: loadMoreAppointments(),
          builder: (context, snapShot) {
            return Container(
              alignment: Alignment.center,
              child: CircularProgressIndicator(
                valueColor: AlwaysStoppedAnimation(
                  Colors.green,
                ),
              ),
            );
          },
        );
      },
      onTap: (details) async {
        if (details.appointments != null) {
          var selected = details.appointments!.first;
          Navigator.of(context).push(MaterialPageRoute(
              builder: ((context) => AppointmentViewingPage(
                  appointment: selected, UserId: _UserId))));
        }
      },
    );
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
    DateTime appStartDate = startDate.add(Duration(days: -1));
    DateTime appEndDate = endDate.add(Duration(days: 1));

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
