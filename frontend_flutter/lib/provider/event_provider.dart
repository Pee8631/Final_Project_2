import 'package:flutter/foundation.dart';
import 'package:frontend_flutter/models/event.dart';

class EventProvider extends ChangeNotifier {
  List<Event> _events = [];

  List<Event> get events => _events.toList();

  DateTime _selectedDate = DateTime.now();

  DateTime get selectedDate => _selectedDate;

  void setDate(DateTime date) => _selectedDate = date;

  List<Event> get eventsOfSelectedDate => _events;

  void addEvent(Event event) {
    _events.add(event);

    notifyListeners();
  }

    void addListEvent(List<Event> events) {
    _events.addAll(events);

    notifyListeners();
  }

  void editEvent(Event newEvent, Event oldEvent) {
    final index = _events.indexOf(oldEvent);
    _events[index] = newEvent;

    notifyListeners();
  }

  void deleteEvent(Event event){
    final index = _events.indexOf(event);
    _events.removeAt(index);
  }
}