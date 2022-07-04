import 'package:flutter/material.dart';
import 'package:frontend_flutter/screens/signupandsignin/signin_userscreen.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../screens/user_screen/edit_personal_user_screen.dart';

class SmallProfileWidget extends StatelessWidget {
  final String imagePath;

  const SmallProfileWidget({Key? key, required this.imagePath})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Stack(
        children: [
          buildImage(context),
        ],
      ),
    );
  }

  Widget buildImage(BuildContext context) {
    final image = NetworkImage(imagePath);

    return PopupMenuButton<int>(
        onSelected: (item) => onSelected(context, item),
        child: ClipOval(
          child: Material(
            color: Colors.transparent,
            child: Ink.image(
              image: image,
              fit: BoxFit.cover,
              width: 50,
              height: 50,
            ),
          ),
        ),
        itemBuilder: (BuildContext context) => [
              PopupMenuItem<int>(
                value: 0,
                child: Text('Edit Profile'),
              ),
              PopupMenuItem<int>(
                value: 1,
                child: Text('Settings'),
              ),
              PopupMenuItem<int>(
                value: 2,
                child: Text('Log Out'),
              ),
            ]);
  }

  onSelected(BuildContext context, int item) async {
    switch (item) {
      case 0:
        Navigator.of(context).push(MaterialPageRoute(
            builder: (BuildContext context) => EditPersonalUserScreen()));
        break;
    }
    switch (item) {
      case 1:
        print('Clicked Settings');
        break;
    }
    switch (item) {
      case 2:
        SharedPreferences sharedPreferences =
            await SharedPreferences.getInstance();
        sharedPreferences.clear();
        sharedPreferences.remove('username');
        sharedPreferences.remove('expireAt');
        sharedPreferences.remove('authToken');
        sharedPreferences.remove('role');
        Navigator.of(context).pushAndRemoveUntil(
            MaterialPageRoute(
                builder: (BuildContext context) => SignInScreen()),
            (route) => false);
        break;
    }
  }
}
