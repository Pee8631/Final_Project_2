import 'package:flutter/material.dart';
import 'package:frontend_flutter/screens/signupandsignin/signin_userscreen.dart';

AppBar buildAppBarLogin(BuildContext context) {
  String text;
  if (context.widget.toString() == "SignInScreen") {
    text = "หน้าเข้าสู่ระบบสำหรับผู้ใช้ทั่วไป";
  } else if (context.widget.toString() == "SignUpDoctorScreen") {
    text = "หน้าสมัครสมาชิกสำหรับแพทย์";
  } else if (context.widget.toString() == "SignInDoctorScreen") {
    text = "หน้าเข้าสู่ระบบสำหรับแพทย์";
  } else {
    text = "หน้าสมัครสมาชิกสำหรับผู้ใช้ทั่วไป";
  }
  if (context.widget.toString() == "SignInScreen") {
    return AppBar(
    backgroundColor: Colors.blue,
    elevation: 0,
    actions: [
      Row(
        children: [Text(text)],
      ),
      SizedBox(width: 10),
    ],
  );
  } else {
    return AppBar(
    leading: BackButton(
      color: Colors.white,
      onPressed: () {
        Navigator.pushReplacement(context,
            MaterialPageRoute(builder: (context) {
          return SignInScreen();
        }));
      },
    ),
    backgroundColor: Colors.blue,
    elevation: 0,
    actions: [
      Row(
        children: [Text(text)],
      ),
      SizedBox(width: 10),
    ],
  );
  }
  
}
