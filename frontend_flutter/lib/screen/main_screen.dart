import 'package:flutter/material.dart';
import 'package:frontend_flutter/screen/Menu_screen.dart';
import 'package:frontend_flutter/screen/doctors_screen.dart';
import 'package:frontend_flutter/screen/notification_screen.dart';
import 'package:frontend_flutter/screen/personal_screen.dart';
import 'package:frontend_flutter/widget/appbar_mainscreen.dart';

class MainScreen extends StatefulWidget {
  const MainScreen(Map<String, dynamic> accessToken, {Key? key})
      : super(key: key);

  @override
  _MainScreenState createState() => _MainScreenState();
}

class _MainScreenState extends State<MainScreen> {

  PageController pageController = PageController();
  int _selectedIndex = 0;
  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
    pageController.jumpToPage(index);
  }

  Widget build(BuildContext context) {
    return Scaffold(
      appBar: buildAppBarMain(context),
      body: PageView(
        children: <Widget>[
          MenuScreen(),
          DoctorsScreen(),
          NotificationScreen(),
          PersonalScreen(),
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
            icon: Icon(Icons.health_and_safety),
            label: 'คุณหมอ',
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
    );
  }
}
