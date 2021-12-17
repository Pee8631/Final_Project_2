import 'package:flutter/material.dart';
import 'package:frontend_flutter/widget/home_button.dart';
import 'login.dart';
import 'register.dart';

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
                RegisterScreen()),
            SizedBox(height: 5),
            HomeButton(Icon(Icons.login), "เข้าสู่ระบบ", 20, Colors.blue,
                LoginScreen()),
          ],
        ),
      ),
    );
  }
}