import 'package:flutter/material.dart';
import 'package:frontend_flutter/model/profile.dart';
import 'package:frontend_flutter/screens/edit_personal_screen.dart';
import 'package:frontend_flutter/util/user_references.dart';
import 'package:frontend_flutter/widget/button_widget.dart';
import 'package:frontend_flutter/widget/large_profile_widget.dart';
import 'package:frontend_flutter/widget/personal_button.dart';
import 'home_screen.dart';

class PersonalScreen extends StatelessWidget {
  const PersonalScreen({Key? key}) : super(key: key);
  final profile = UserReferences.myProfile;
  final text = 'แก้ไขโปรไฟล์ของคุณ';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: ListView(
        physics: BouncingScrollPhysics(),
        children: [
          SizedBox(
            height: 10,
          ),
          LargeProfileWidget(
            imagePath: profile.imagePath,
            onClicked: () async {
              Navigator.pushReplacement(context,
                  MaterialPageRoute(builder: (context) {
                return HomeScreen();
              }));
            },
          ),
          const SizedBox(
            height: 24,
          ),
          buildName(profile),
          const SizedBox(
            height: 24,
          ),
          Center(child: buildUpgradeButton()),
          const SizedBox(
            height: 24,
          ),
          buildAbout(profile),
          const SizedBox(
            height: 24,
          ),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 10),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Container(
                  height: 1,
                  color: Colors.grey,
                ),
                SizedBox(
                  height: 3,
                ),
                PersonalButtonWidget(
                  text: 'แก้ไขข้อมูลส่วนตัว',
                  onClicked: () {
                    Navigator.pushReplacement(context,
                        MaterialPageRoute(builder: (context) {
                      return EditPersonalScreen();
                    }));
                  },
                ),
                SizedBox(
                  height: 3,
                ),
                Container(
                  height: 1,
                  color: Colors.grey,
                ),
                SizedBox(
                  height: 3,
                ),
                PersonalButtonWidget(
                  text: 'รายการย้อนหลัง',
                  onClicked: () {},
                ),
                SizedBox(
                  height: 3,
                ),
                Container(
                  height: 1,
                  color: Colors.grey,
                ),
                SizedBox(
                  height: 3,
                ),
                PersonalButtonWidget(
                  text: 'คุณหมอคนโปรด',
                  onClicked: () {},
                ),
                SizedBox(
                  height: 3,
                ),
                Container(
                  height: 1,
                  color: Colors.grey,
                ),
                SizedBox(
                  height: 3,
                ),
                PersonalButtonWidget(
                  text: 'ตั้งค่า',
                  onClicked: () {},
                ),
                SizedBox(
                  height: 3,
                ),
                Container(
                  height: 1,
                  color: Colors.grey,
                ),
                SizedBox(
                  height: 3,
                ),
                PersonalButtonWidget(
                  text: 'ติดต่อเรา',
                  onClicked: () {},
                ),
                SizedBox(
                  height: 3,
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  Widget buildName(Profile profile) => Column(
        children: [
          Text(
            profile.name,
            style: TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 24,
            ),
          ),
          const SizedBox(
            height: 4,
          ),
          Text(
            profile.email,
            style: TextStyle(
                fontWeight: FontWeight.bold, fontSize: 15, color: Colors.grey),
          ),
        ],
      );

  Widget buildUpgradeButton() => ButtonWidget(
        onClicked: () {},
        text: text,
      );

  buildAbout(Profile profile) => Container(
        padding: EdgeInsets.symmetric(horizontal: 48),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              'เกี่ยวกับ',
              style: TextStyle(
                fontWeight: FontWeight.bold,
                fontSize: 24,
              ),
            ),
            const SizedBox(
              height: 16,
            ),
            Text(
              profile.about,
              style: TextStyle(
                fontWeight: FontWeight.bold,
                fontSize: 15,
              ),
            ),
          ],
        ),
      );
}
