import 'package:flutter/material.dart';
import 'package:form_field_validator/form_field_validator.dart';
import 'package:frontend_flutter/screen/Menu_screen.dart';
import 'package:frontend_flutter/screen/main_screen.dart';
import 'package:frontend_flutter/screen/notification_screen.dart';
import 'package:frontend_flutter/screen/personal_screen.dart';
import 'package:frontend_flutter/widget/appbar_mainscreen.dart';

import 'doctors_screen.dart';

class EditPersonalScreen extends StatefulWidget {
  const EditPersonalScreen({Key? key}) : super(key: key);

  @override
  _EditPersonalScreenState createState() => _EditPersonalScreenState();
}

class _EditPersonalScreenState extends State<EditPersonalScreen> {
  PageController pageController = PageController();
  int _selectedIndex = 0;
  _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
    pageController.jumpToPage(index);
  }

  /*Map<String,dynamic> accessToken = {};
  turnBackPage(nextPage) {
    Navigator.pushReplacement(context, MaterialPageRoute(builder: (context) {
      return nextPage;
    }));
  }*/

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: buildAppBarMain(context),
      body: Container(
        padding: const EdgeInsets.all(10.0),
        child: Form(
          child: SingleChildScrollView(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  "อีเมล",
                  style: TextStyle(fontSize: 20),
                ),
                TextFormField(
                  validator: MultiValidator([
                    RequiredValidator(errorText: "กรุณาป้อนอีเมล"),
                    EmailValidator(errorText: "รูปแบบอีเมลไม่ถูกต้อง"),
                  ]),
                  onSaved: (Email) {},
                ),
                SizedBox(
                  height: 15,
                ),
                Text(
                  "เลขบัตรประชาชน",
                  style: TextStyle(fontSize: 20),
                ),
                TextFormField(
                  validator: MultiValidator([
                    RequiredValidator(errorText: "กรุณาป้อนเลขบัตรประชาชน"),
                  ]),
                  onSaved: (IDCardNumber) {},
                ),
                SizedBox(
                  height: 15,
                ),
                Text(
                  "ชื่อ",
                  style: TextStyle(fontSize: 20),
                ),
                TextFormField(
                  validator: MultiValidator([
                    RequiredValidator(errorText: "กรุณาป้อนชื่อ"),
                  ]),
                  onSaved: (Firstname) {},
                ),
                SizedBox(
                  height: 15,
                ),
                Text(
                  "นามสกุล",
                  style: TextStyle(fontSize: 20),
                ),
                TextFormField(
                  validator: MultiValidator([
                    RequiredValidator(errorText: "กรุณาป้อนนามสกุล"),
                  ]),
                  onSaved: (Lastname) {},
                ),
                SizedBox(
                  height: 15,
                ),
                Text(
                  "วันเดือนปีเกิด",
                  style: TextStyle(fontSize: 20),
                ),
                SizedBox(
                  height: 15,
                ),
                Text(
                  "เพศ",
                  style: TextStyle(fontSize: 20),
                ),
                
                SizedBox(
                  height: 15,
                ),
                Text(
                  "กรุ๊ปเลือด",
                  style: TextStyle(fontSize: 20),
                ),
              ],
            ),
          ),
        ),
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
