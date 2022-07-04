import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:frontend_flutter/models/pInfo.dart';
import 'package:frontend_flutter/provider/event_provider.dart';
import 'package:frontend_flutter/screens/share_screen/chats/chats_screen.dart';
import 'package:frontend_flutter/screens/user_screen/menu_user_screen.dart';
import 'package:frontend_flutter/screens/share_screen/notification_screen.dart';
import 'package:frontend_flutter/screens/user_screen/personal_user_screen.dart';
import 'package:frontend_flutter/screens/user_screen/schedule_user_srceen.dart';
import 'package:frontend_flutter/widget/appbar_mainscreen.dart';
import 'package:http/http.dart' as http;
import 'package:provider/provider.dart';
import 'package:shared_preferences/shared_preferences.dart';

class MainUserScreen extends StatefulWidget {
  final int ScreenIndex;
  const MainUserScreen({Key? key, this.ScreenIndex = 0}) : super(key: key);

  @override
  _MainUserScreenState createState() => _MainUserScreenState();
}

class _MainUserScreenState extends State<MainUserScreen> {
  PageController pageController = PageController();
  int _selectedIndex = 0;
  String _name = 'Fristname Lastname';
  Future<String>? _futureName;
  int _UserId = 0;
  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });

    pageController.jumpToPage(index);
  }

  @override
  initState() {
    getUserId();
    _getPInfo();

    pageController = PageController(initialPage: widget.ScreenIndex);
    setState(() {
      _selectedIndex = widget.ScreenIndex;
      _futureName = _getPInfo();
    });
    super.initState();
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

  Future<String> _getPInfo() async {
    await getUserId();
    //FutureBuilder(future: getToken(), builder: (BuildContext context, AsyncSnapshot<void> snapshot) { },);
    var response = await http.get(
        Uri.parse('http://10.0.2.2:8080/api/v1/pinfos/' + _UserId.toString()));
    var name = _name;
    if (response.statusCode == 200) {
      PInfo results = pInfoFromJson(response.body);
      if (results != null) {
        name = results.firstName + " " + results.lastName;
      }
    }
    return name;
  }

  Widget build(BuildContext context) {
    return FutureBuilder(
      future: _futureName,
      builder: (BuildContext context, AsyncSnapshot<String> snapshot) {
        _name = snapshot.data ?? _name;
        return Scaffold(
          backgroundColor: Color.fromARGB(255, 208, 244, 255),
          appBar: buildAppBarMain(context, _name),
          body: PageView(
            children: <Widget>[
              MenuScreen(name: _name, UserId: _UserId),
              ScheduleUserScreen(UserId: _UserId),
              ChatsScreen(),
              NotificationScreen(),
              PersonalUserScreen(),
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
                icon: Icon(Icons.schedule),
                label: 'ตาราง(ชั่วคราว)',
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
      },
    );
  }
}
