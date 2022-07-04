import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:frontend_flutter/screens/doctor_screen/personal_doctor_screen.dart';
import 'package:frontend_flutter/screens/doctor_screen/schedule_doctor_screen.dart';
import 'package:frontend_flutter/screens/share_screen/chats/chats_screen.dart';
import 'package:frontend_flutter/screens/share_screen/notification_screen.dart';
import 'package:http/http.dart' as http;
import 'package:frontend_flutter/provider/event_provider.dart';
import 'package:frontend_flutter/widget/appbar_mainscreen.dart';
import 'package:provider/provider.dart';
import 'package:shared_preferences/shared_preferences.dart';

class MainDoctorScreen extends StatefulWidget {
  final int ScreenIndex;
  const MainDoctorScreen({Key? key, required this.ScreenIndex}) : super(key: key);

  @override
  State<MainDoctorScreen> createState() => _MainDoctorScreenState();
}

class _MainDoctorScreenState extends State<MainDoctorScreen> {
   PageController pageController = PageController();
  int _selectedIndex = 0;
  String _name = 'Fristname Lastname';
  int _UserId = 0;
  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });

    pageController.jumpToPage(index);
  }

  @override
  initState() {
    getName();
    getUserId();
    super.initState();
    pageController = PageController(initialPage: widget.ScreenIndex);
    setState(() {
      _selectedIndex = widget.ScreenIndex;
    });
  }

  Future<void> getName() async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    final String? name = sharedPreferences.getString('username');
    setState(() => _name = name!);
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

  Widget build(BuildContext context) {
    return ChangeNotifierProvider<EventProvider>(
      create: (context) => EventProvider(),
      child: Consumer<EventProvider>(
        builder: (context, provider, child) => Scaffold(
          appBar: buildAppBarMain(context, _name),
          body: PageView(
            children: <Widget>[
              ScheduleDoctorScreen(UserId: _UserId),
              ChatsScreen(),
              NotificationScreen(),
              PersonalDoctorScreen(),
            ],
            controller: pageController,
          ),
          bottomNavigationBar: BottomNavigationBar(
            items: const <BottomNavigationBarItem>[
              BottomNavigationBarItem(
                icon: Icon(Icons.home),
                label: 'หน้าหลัก',
                backgroundColor: Colors.blue,
              ),
              BottomNavigationBarItem(
                icon: Icon(Icons.chat),
                label: 'ห้องแชท',
                backgroundColor: Colors.blue,
              ),
              BottomNavigationBarItem(
                icon: Icon(Icons.notifications_none),
                label: 'แจ้งเตือน',
                backgroundColor: Colors.blue,
              ),
              BottomNavigationBarItem(
                icon: Icon(Icons.person),
                label: 'ส่วนตัว',
                backgroundColor: Colors.blue,
              ),
            ],
            currentIndex: _selectedIndex,
            selectedItemColor: Colors.white,
            unselectedItemColor: Colors.white,
            onTap: _onItemTapped,
          ),
        ),
      ),
    );
  }
}