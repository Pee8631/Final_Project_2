import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:frontend_flutter/models/schedule.dart';
import 'package:frontend_flutter/screens/user_screen/main_user_screen.dart';
//import 'package:frontend_flutter/provider/event_provider.dart';
import 'package:http/http.dart' as http;
import 'package:interval_time_picker/interval_time_picker.dart';
import 'package:syncfusion_flutter_calendar/calendar.dart';
import 'package:time_picker_widget/time_picker_widget.dart';
import '../util/utils.dart';

class EventEditingPage extends StatefulWidget {
  final Appointment? appointment;
  final int UserId;
  const EventEditingPage({Key? key, this.appointment, required this.UserId})
      : super(key: key);

  @override
  State<EventEditingPage> createState() => _EventEditingPageState();
}

class _EventEditingPageState extends State<EventEditingPage> {
  final _formKey = GlobalKey<FormState>();
  final titleController = TextEditingController();
  final decsriptionController = TextEditingController();
  late int id;
  late DateTime fromDate;
  late DateTime toDate;
  late Schedule _schedule;
  bool isAllDay = false;

  @override
  void initState() {
    super.initState();

    if (widget.appointment == null) {
      fromDate = new DateTime.utc(DateTime.now().year, DateTime.now().month,
          DateTime.now().day, DateTime.now().hour, 0);
      toDate = new DateTime.utc(DateTime.now().year, DateTime.now().month,
              DateTime.now().day, DateTime.now().hour, 0)
          .add(Duration(hours: 1));
    } else {
      final event = widget.appointment!;
      id = event.id as int;
      titleController.text = event.subject;
      decsriptionController.text = event.notes!;
      fromDate = event.startTime;
      toDate = event.endTime;
    }
  }

  @override
  void dispose() {
    titleController.dispose();

    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          leading: CloseButton(),
          actions: buildEditingActions(),
        ),
        body: SingleChildScrollView(
            padding: EdgeInsets.all(12),
            child: Form(
              key: _formKey,
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: <Widget>[
                  buildTitle(),
                  SizedBox(height: 12),
                  if (isAllDay == false) ...[
                    buildDateTimePicker(),
                  ],
                  //buildAllDayEvents(),
                  buildDescriptions(),
                ],
              ),
            )));
  }

  List<Widget> buildEditingActions() => [
        ElevatedButton.icon(
            style: ElevatedButton.styleFrom(
              primary: Colors.transparent,
              shadowColor: Colors.transparent,
            ),
            onPressed: saveForm,
            icon: Icon(Icons.done),
            label: Text("Save"))
      ];

  Widget buildTitle() => TextFormField(
        style: TextStyle(fontSize: 24),
        decoration:
            InputDecoration(border: UnderlineInputBorder(), hintText: 'หัวข้อ'),
        onFieldSubmitted: (_) => saveForm(),
        validator: (title) =>
            title != null && title.isEmpty ? 'กรุณาระบุหัวข้อ' : null,
        controller: titleController,
      );

  Widget buildDateTimePicker() => Column(
        children: [buildDay(), buildTime()],
      );

  Widget buildDay() => buildHeader(
        header: 'วัน',
        child: Row(
          children: [
            Expanded(
              flex: 2,
              child: buildDropdownField(
                text: Utils.toDate(fromDate),
                onClicked: () => pickFromDateTime(pickDate: true),
              ),
            ),
            // Expanded(
            //   flex: 2,
            //   child: buildDropdownField(
            //     text: Utils.toDate(toDate),
            //     onClicked: () => pickToDateTime(pickDate: true),
            //   ),
            // ),
          ],
        ),
      );

  Widget buildTime() => buildHeader(
        header: 'เวลา',
        child: Row(
          children: [
            Expanded(
              child: buildDropdownField(
                text: 'ตั้งแต่ :         ' + Utils.toTime(fromDate),
                onClicked: () => pickFromDateTime(pickDate: false),
              ),
            ),
            Expanded(
              child: buildDropdownField(
                text: 'จนถึง :         ' + Utils.toTime(toDate),
                onClicked: () => pickToDateTime(pickDate: false),
              ),
            ),
          ],
        ),
      );

  Future pickToDateTime({required bool pickDate}) async {
    final date = await pickDateTime(toDate,
        pickDate: pickDate, firstDate: pickDate ? fromDate : null);
    if (date == null) return;

    if (date.isAfter(toDate)) {
      toDate = DateTime(
          fromDate.year, fromDate.month, fromDate.day, date.hour, date.minute);
    }

    setState(() => toDate = date);
  }

  Future pickFromDateTime({required bool pickDate}) async {
    final date = await pickDateTime(fromDate, pickDate: pickDate);
    if (date == null) return;

    if (date.isAfter(toDate)) {
      toDate =
          DateTime(date.year, date.month, date.day, date.hour, date.minute);
    }

    setState(() => fromDate = date);
  }

  Future<DateTime?> pickDateTime(DateTime initialDate,
      {required bool pickDate, DateTime? firstDate}) async {
    if (pickDate) {
      final date = await showDatePicker(
        context: context,
        initialDate: initialDate,
        firstDate: firstDate ?? DateTime(2015, 8),
        lastDate: DateTime(2101),
      );
      if (date == null) return null;

      final time =
          Duration(hours: initialDate.hour, minutes: initialDate.minute);

      return date.add(time);
    } else {
      final timeOfDay = await showIntervalTimePicker(
        context: context,
        interval: 5,
        visibleStep: VisibleStep.Fifteenths,
        initialTime: TimeOfDay.fromDateTime(initialDate),
        //     selectableTimePredicate: (time) => initialDate.hour.indexOf(time!.hour) != -1 &&   time.minute % 15 == 0).then(
        // (time) =>
        //     setState(() => selectedTime = time?.format(context)),
      );

      if (timeOfDay == null) return null;

      final date =
          DateTime(initialDate.year, initialDate.month, initialDate.day);

      final time = Duration(hours: timeOfDay.hour, minutes: timeOfDay.minute);

      return date.add(time);
    }
  }

  Widget buildDropdownField(
          {required String text, required VoidCallback onClicked}) =>
      ListTile(
        title: Text(text),
        trailing: Icon(Icons.arrow_drop_down),
        onTap: onClicked,
      );

  Widget buildHeader({
    required String header,
    required Widget child,
  }) =>
      Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(header, style: TextStyle(fontWeight: FontWeight.bold)),
          child,
        ],
      );

  //Save here for later
  Future<void> createSchedule(String Title, String decscription,
      DateTime startTime, DateTime endTime, int userId) async {
    /*String _formatdate =
        new DateFormat('yyyy-MM-ddTHH:mm:ss.mmmuuuZ').format(_pinfo.brithDate ?? _selectdate);*/
    var response = await http.post(
      Uri.parse('http://10.0.2.2:8080/api/v1/schedules'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        'Activity': Title,
        'Detail': decscription,
        'Status': 'Active',
        'StartTime': startTime.toIso8601String(),
        'EndTime': endTime.toIso8601String(),
        'UserId': userId
      }),
    );
    final results;
    if (response.statusCode == 200) {
      results = await scheduleFromJson(response.body);
      setState(() {
        _schedule = results;
      });
    } else {
      throw ("Can't Save Schedule To Server.");
    }
  }

  Future<void> editSchedule(int id, String Title, String decscription,
      DateTime startTime, DateTime endTime, int userId) async {
    /*String _formatdate =
        new DateFormat('yyyy-MM-ddTHH:mm:ss.mmmuuuZ').format(_pinfo.brithDate ?? _selectdate);*/
    var response = await http.put(
      Uri.parse('http://10.0.2.2:8080/api/v1/schedules/' + id.toString()),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        'Activity': Title,
        'Detail': decscription,
        'Status': 'Active',
        'StartTime': startTime.toIso8601String(),
        'EndTime': endTime.toIso8601String(),
        'UserId': userId
      }),
    );
    final results;
    if (response.statusCode == 200) {
      results = scheduleFromJson(response.body);
      setState(() {
        _schedule = results;
      });
    } else {
      throw ("Can't Save Schedule To Server.");
    }
  }

  Future saveForm() async {
    final isValid = _formKey.currentState!.validate();
    if (isValid) {
      if (toDate.isAfter(fromDate)) {
        if (widget.appointment == null) {
          await createSchedule(titleController.text, decsriptionController.text,
              fromDate, toDate, widget.UserId);
        } else {
          await editSchedule(
              widget.appointment!.id! as int,
              titleController.text,
              decsriptionController.text,
              fromDate,
              toDate,
              widget.UserId);
        }
        if (_schedule != null) {
          // final event = Event(
          //     title: _schedule.activity,
          //     from: _schedule.startTime,
          //     to: _schedule.endTime,
          //     decscription: _schedule.detail,
          //     isAllDay: isAllDay);

          // final isEditing = widget.event != null;

          // final provider = Provider.of<EventProvider>(context, listen: false);
          // if (isEditing) {
          //   provider.editEvent(event, widget.event!);

          //   Navigator.of(context).pop();
          // } else {
          //   provider.addEvent(event);
          // }
          Navigator.pop(context, MaterialPageRoute(builder: (context) {
            return MainUserScreen(
              ScreenIndex: 1,
            );
          }));
        } else {
          throw ('Schedule was null');
        }
      } else {
        Fluttertoast.showToast(
            msg:
                "คุณตั้งเวลาเริ่มต้นและสิ้นสุดไม่ถูกต้อง\nกรุณาตั้งเวลาใหม่อีกครั้ง",
            gravity: ToastGravity.CENTER);
      }
    }
  }

  Widget buildAllDayEvents() => Row(children: [
        Expanded(
          child: Text("All Day Events: ", style: TextStyle(fontSize: 16)),
        ),
        Expanded(
          child: Checkbox(
              value: isAllDay,
              onChanged: (isAllDay) => setState(() {
                    this.isAllDay = isAllDay!;
                  })),
        )
      ]);

  Widget buildDescriptions() => TextFormField(
        style: TextStyle(fontSize: 20),
        decoration: InputDecoration(
            border: UnderlineInputBorder(), hintText: 'รายละเอียด'),
        onFieldSubmitted: (_) => saveForm(),
        controller: decsriptionController,
      );
}
