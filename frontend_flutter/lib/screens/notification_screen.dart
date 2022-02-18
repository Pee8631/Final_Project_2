import 'package:flutter/material.dart';
import 'package:frontend_flutter/util/user_references.dart';

class NotificationScreen extends StatelessWidget {
  NotificationScreen({Key? key}) : super(key: key);
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
            "แจ้งเตือน",
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
            title: Text(
              "title$index",
              style: TextStyle(fontSize: 12, color: Colors.black),
              textAlign: TextAlign.left,
            ),
            subtitle: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  "subtitle$index : ${textList[index]}",
                  style: TextStyle(fontSize: 12, color: Colors.black),
                  textAlign: TextAlign.left,
                ),
              ],
            ),
          );
        }, separatorBuilder: (BuildContext context, int index) => const Divider(color: Colors.black,),));
  }
}
