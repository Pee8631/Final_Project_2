import 'package:flutter/material.dart';
import 'package:frontend_flutter/screens/signupandsignin/signin_user_screen.dart';

AppBar buildAppBarLogin(BuildContext context) {
    return AppBar(
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
