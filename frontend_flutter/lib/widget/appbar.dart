import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:frontend_flutter/screen/home.dart';
import 'package:frontend_flutter/util/user_references.dart';
import 'package:frontend_flutter/widget/profile_widget.dart';

AppBar buildAppBar(BuildContext context) {
  final myprofile = UserReferences.myProfile;
  return AppBar(
    leading: BackButton(),
    backgroundColor: Colors.blue,
    elevation: 0,
    actions: [
      Row(
        children: [
          Positioned(bottom: 0, child : Text(myprofile.name,style: TextStyle(fontSize: 10)),),
          ProfileWidget(
            imagePath: myprofile.imagePath,
            onClicked: () async {
              Navigator.pushReplacement(context,
                  MaterialPageRoute(builder: (context) {
                return HomeScreen();
              }));
            },
          ),
        ],
      ),
      SizedBox(width: 10),
    ],
  );
}
