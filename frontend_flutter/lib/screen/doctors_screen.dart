import 'package:flutter/material.dart';
import 'package:frontend_flutter/util/user_references.dart';

class DoctorsScreen extends StatelessWidget {
  final profile = UserReferences.myProfile;
  final List<String> textList = [
    "He'd have you",
    "Heed not the rabble",
    "Sound of screams",
    "Who scream",
    "Revolution is",
    "Revolution, they..."
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          backgroundColor: Colors.transparent,
          elevation: 0,
          title: Center(
              child: Text(
            "รายชื่อ คุณหมอ",
            style: TextStyle(fontSize: 20, color: Colors.black),
          )),
        ),
        body: ListView.separated(
          padding: const EdgeInsets.all(8.0),
          itemCount: textList.length,
          itemBuilder: (BuildContext context, int index) {
          return ListTile(
            selectedColor: Colors.grey,
            onTap: () {},
            leading: ClipOval(
            child: Image.asset(
              "assets/images/Profile_Default.png",
              width: 60,
              height: 60,
            ),
            ),
            title: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  "Firstname$index Lastname$index",
                  style: TextStyle(fontSize: 12, color: Colors.black), textAlign: TextAlign.left,
                ),
                Text(
                  "Hospital$index",
                  style: TextStyle(fontSize: 12, color: Colors.black),textAlign: TextAlign.left,
                ),
                Text(
                  "Clinic$index",
                  style: TextStyle(fontSize: 12, color: Colors.black),textAlign: TextAlign.left,
                ),
                Text(
                  "Price$index",
                  style: TextStyle(fontSize: 12, color: Colors.black),textAlign: TextAlign.left,
                ),
                Text(
                  "about : ${textList[index]}",
                  style: TextStyle(fontSize: 12, color: Colors.black),textAlign: TextAlign.left,
                ),
              ],
            ),
          );
        }, separatorBuilder: (BuildContext context, int index) => const Divider(color: Colors.black),));
  }
}
