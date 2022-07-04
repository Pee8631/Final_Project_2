import 'package:flutter/material.dart';
import 'package:frontend_flutter/widget/small_profile_widget.dart';

AppBar buildAppBarBackToScreen(BuildContext context, String name, Widget screen) {
  return AppBar(
    leading: 
        BackButton(
            color: Colors.white,
            onPressed: () {
              Navigator.pushReplacement(context,
                  MaterialPageRoute(builder: (context) {
                return screen;
              }));
            },
          ),
    backgroundColor: Colors.blue,
    elevation: 0,
    actions: [
      Row(
        children: [
          Text(name + "", style: TextStyle(fontSize: 10)),
          SizedBox(
            width: 10,
          ),
          SmallProfileWidget(
            imagePath:
                'https://www.jumpstarttech.com/files/2018/08/Network-Profile.png',
          ),
          // PopupMenuButton<int>(

          // ),
          //   itemBuilder: (BuildContext context) => [
          //     PopupMenuItem<int>(
          //       value: 0,
          //       child: Text('Settings'),
          //     ),
          //   ],
          // ),
          // SmallProfileWidget(
          //   imagePath:
          //       'https://www.jumpstarttech.com/files/2018/08/Network-Profile.png',
          //   onClicked: () async {

          //   },
          // ),
        ],
      ),
      SizedBox(width: 10),
    ],
  );
}
