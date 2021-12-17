import 'package:flutter/material.dart';
import 'package:frontend_flutter/util/user_references.dart';

class DoctorsScreen extends StatelessWidget {
  const DoctorsScreen({Key? key})
      : super(key: key);
  final profile = UserReferences.myProfile;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column( children: [
        Text("หน้าจอ คุณหมอ"),
      ],),
      );
  }
}
