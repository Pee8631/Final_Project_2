import 'package:flutter/material.dart';

AppBar buildAppBarDoctorSignUp(BuildContext context,Widget screen) {
  return AppBar(
    leading: BackButton(
      onPressed: () => Navigator.pushReplacement(context,
          MaterialPageRoute(builder: (context) {
        return screen;
      })),
    ),
    backgroundColor: Color.fromARGB(232, 100, 180, 255),
    elevation: 0,
    actions: [
      Container(
          alignment: Alignment.center,
          margin: const EdgeInsets.all(10),
          child: Text('Online Doctor Application')),
    ],
  );
}
