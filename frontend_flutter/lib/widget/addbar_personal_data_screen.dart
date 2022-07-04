import 'package:flutter/material.dart';
import 'package:frontend_flutter/widget/small_profile_widget.dart';

import '../screens/user_screen/main_user_screen.dart';

AppBar buildAppBarPersonalScreen(BuildContext context, String name) {
  return AppBar(
    leading: 
        BackButton(
            color: Colors.white,
            onPressed: () {
              Navigator.pushReplacement(context,
                  MaterialPageRoute(builder: (context) {
                return MainUserScreen(ScreenIndex: 4,);
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
