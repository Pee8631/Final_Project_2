import 'package:flutter/material.dart';
import 'package:frontend_flutter/widget/small_profile_widget.dart';

AppBar buildAppBarMain(BuildContext context, String name) {
  return AppBar(
    // leading: context.widget.toString() == "DoctorsScreen"
    //     BackButton(
    //         color: Colors.white,
    //         onPressed: () {
    //           Navigator.pushReplacement(context,
    //               MaterialPageRoute(builder: (context) {
    //             return MainScreen();
    //           }));
    //         },
    //       )
    //     : null,
    backgroundColor: Color.fromARGB(232, 100, 180, 255),
    elevation: 0,
    actions: [
      Row(
        children: [
          Text(name + "", style: TextStyle(fontSize: 12, fontWeight: FontWeight.w900, color: Colors.black45)),
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
