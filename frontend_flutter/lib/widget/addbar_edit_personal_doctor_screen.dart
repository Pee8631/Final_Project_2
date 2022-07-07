import 'package:flutter/material.dart';
import 'package:frontend_flutter/screens/doctor_screen/main_doctor_screen.dart';

AppBar buildAppBarEditPersonalDoctorScreen(BuildContext context) {
  return AppBar(
    leading: BackButton(
      color: Colors.white,
      onPressed: () {
        Navigator.pushReplacement(context,
            MaterialPageRoute(builder: (context) {
          return MainDoctorScreen(
            ScreenIndex: 3,
          );
        }));
      },
    ),
    backgroundColor: Color.fromARGB(232, 100, 180, 255),
    elevation: 0,
  );
}
