import 'package:flutter/material.dart';

class Event {
  final String title;
  final String decscription;
  final DateTime from;
  final DateTime to;
  final Color backgroundColor;
  final bool isAllDay;

  const Event({
    required this.title,
    required this.decscription,
    required this.from,
    required this.to,
    this.backgroundColor = Colors.blue,
    this.isAllDay = false,
  });
}
