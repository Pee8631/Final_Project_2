import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:frontend_flutter/models/pInfo.dart';
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
  const MainDoctorScreen({Key? key, required this.ScreenIndex})
      : super(key: key);

  @override
  State<MainDoctorScreen> createState() => _MainDoctorScreenState();
}

class _MainDoctorScreenState extends State<MainDoctorScreen> {
  PageController pageController = PageController();
  int _selectedIndex = 0;
  String _name = "Fristname Lastname";
  String _Profile = "assets/images/Profile_Default.png";
  Future<PInfo>? _futurePInfo;
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
      _futurePInfo = _getPInfo();
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

  Future<PInfo>? _getPInfo() async {
    await getUserId();
    //FutureBuilder(future: getToken(), builder: (BuildContext context, AsyncSnapshot<void> snapshot) { },);
    var response = await http.get(
        Uri.parse('http://10.0.2.2:8080/api/v1/pinfos/' + _UserId.toString()));
    //var name = _name;
    PInfo results = PInfo(
        about: '',
        address: '',
        bloodGroup: '',
        brithDate: null,
        firstName: '',
        gender: 0,
        id: 0,
        idCardNumber: '',
        lastName: '',
        prefix: '',
        profile: '',
        user: 0);
    if (response.statusCode == 200) {
      results = pInfoFromJson(response.body);
      // if (results != null) {
      //   name = results.firstName + " " + results.lastName;
      // }
    }
    return results;
  }

  Widget build(BuildContext context) {
    return FutureBuilder(
        future: _futurePInfo,
        builder: (BuildContext context, AsyncSnapshot<PInfo> snapshot) {
          _name = snapshot.data == null
              ? _name
              : snapshot.data!.firstName + " " + snapshot.data!.lastName;
          if (snapshot.data != null) {
            _Profile = snapshot.data!.profile != null
                ? snapshot.data!.profile!
                : _Profile;
          }

          return Scaffold(
            backgroundColor: Color.fromARGB(255, 208, 244, 255),
            appBar: buildAppBarMain(context, _name, _Profile),
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
                  backgroundColor: Color.fromARGB(232, 100, 180, 255),
                ),
                BottomNavigationBarItem(
                  icon: Icon(Icons.chat),
                  label: 'ห้องแชท',
                  backgroundColor: Color.fromARGB(232, 100, 180, 255),
                ),
                BottomNavigationBarItem(
                  icon: Icon(Icons.notifications_none),
                  label: 'แจ้งเตือน',
                  backgroundColor: Color.fromARGB(232, 100, 180, 255),
                ),
                BottomNavigationBarItem(
                  icon: Icon(Icons.person),
                  label: 'ส่วนตัว',
                  backgroundColor: Color.fromARGB(232, 100, 180, 255),
                ),
              ],
              currentIndex: _selectedIndex,
              selectedItemColor: Colors.white,
              unselectedItemColor: Colors.white,
              onTap: _onItemTapped,
            ),
          );
        });
  }
}
