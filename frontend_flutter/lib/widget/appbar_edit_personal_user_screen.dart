import 'package:flutter/material.dart';
import 'package:frontend_flutter/screens/doctor_screen/main_doctor_screen.dart';
import 'package:frontend_flutter/screens/user_screen/main_user_screen.dart';

AppBar buildAppBarEditPersonalUserScreen(BuildContext context) {
  return AppBar(
    leading: BackButton(
      color: Colors.white,
      onPressed: () {
        Navigator.pushReplacement(context,
            MaterialPageRoute(builder: (context) {
          return MainUserScreen(
            ScreenIndex: 4,
          );
        }));
      },
    ),
    backgroundColor: Color.fromARGB(232, 100, 180, 255),
    elevation: 0,
  );
}
