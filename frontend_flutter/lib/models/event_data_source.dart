import 'package:flutter/material.dart';
import 'package:frontend_flutter/models/event.dart';
import 'package:syncfusion_flutter_calendar/calendar.dart';

class EventDataSource extends CalendarDataSource{
  EventDataSource(List<Appointment> appointments){
    this.appointments = appointments;
  }

  Event getEvent(int index) => appointments![index] as Event;

  DateTime getStartTime(int index) => getEvent(index).from;

  DateTime getEndTime(int index) => getEvent(index).to;

  String getSubject(int index) => getEvent(index).title;

  String getDecsription(int index) => getEvent(index).decscription;

  Color getColor(int index) => getEvent(index).backgroundColor;

  bool isAllDay(int index) => getEvent(index).isAllDay;
}