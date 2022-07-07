import 'package:flutter/material.dart';
import 'package:frontend_flutter/screens/signupandsignin/signup_user_screen.dart';
import 'package:frontend_flutter/widget/home_button.dart';
import 'signupandsignin/signin_user_screen.dart';
import 'signupandsignin/signup_user_screen.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Register&Login"),
      ),
      body: Padding(
        padding: const EdgeInsets.fromLTRB(10, 50, 10, 0),
        child: Column(
          children: [
            Image.asset("assets/images/Default.png"),
            SizedBox(height: 5),
            HomeButton(Icon(Icons.add), "สร้างบัญชีผู้ใช้", 20, Colors.blue,
                SignUpUserScreen()),
            SizedBox(height: 5),
            HomeButton(Icon(Icons.login), "เข้าสู่ระบบ", 20, Colors.blue,
                SignInUserScreen()),
          ],
        ),
      ),
    );
  }
}
