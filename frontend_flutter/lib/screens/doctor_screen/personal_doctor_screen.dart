import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:frontend_flutter/models/pInfo.dart';
import 'package:frontend_flutter/screens/doctor_screen/edit_personal_doctor_screen.dart';
import 'package:frontend_flutter/util/user_preferences.dart';
import 'package:frontend_flutter/widget/button_widget.dart';
import 'package:frontend_flutter/widget/large_profile_widget.dart';
import 'package:frontend_flutter/widget/personal_button.dart';
import 'package:http/http.dart' as http;
import 'package:intl/intl.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../home_screen.dart';

class PersonalDoctorScreen extends StatefulWidget {
  const PersonalDoctorScreen({Key? key}) : super(key: key);

  @override
  State<PersonalDoctorScreen> createState() => _PersonalDoctorScreenState();
}

class _PersonalDoctorScreenState extends State<PersonalDoctorScreen> {
  final profile = UserPreferences.myProfile;
  final text = 'แก้ไขโปรไฟล์ของคุณ';
  int _UserId = 0;
  late Future<PInfo> pInfos;
  @override
  initState() {
    pInfos = getPInfo();
    super.initState();
  }

  Future<String> getToken() async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    final String? authToken = sharedPreferences.getString('authToken');
    //authToken?.substring(0, authToken.length - 1);
    return await authToken!;
  }

  Future<int> getUserId() async {
    //FutureBuilder(future: getToken(), builder: (BuildContext context, AsyncSnapshot<void> snapshot) { },);
    var response = await http.get(
        Uri.parse('http://10.0.2.2:8080/api/v1/tokens/' + await getToken()));
    if (response.statusCode == 200) {
      final results = jsonDecode(response.body);
      // setState(() => _UserId = results!);
      return await results;
    } else {
      throw ("User Not Found : " + response.reasonPhrase!);
    }
  }


  Future<PInfo> getPInfo() async {
    var UserId = await getUserId();
    var response = await http.get(Uri.parse(
        'http://10.0.2.2:8080/api/v1/pinfos/' + UserId.toString()));
    if (response.statusCode == 200) {
      final results = pInfoFromJson(response.body);
      return results;
    } else {
      throw ("PInfo Not Found: " + response.reasonPhrase!);
    }
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: pInfos,
      builder: (BuildContext context, AsyncSnapshot<PInfo> snapshot) {
        if (snapshot.hasError) {
          return Scaffold(
            appBar: AppBar(
              title: Text("Error"),
            ),
            body: Center(
              child: Text(
                  "ขออภัยไม่สามารถเชื่อมต่อกับ Server ได้ในขณะนี้\n\nError : ${snapshot.error}"),
            ),
          );
        } else if (snapshot.connectionState == ConnectionState.done) {
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
                buildName(snapshot.data!),
                const SizedBox(
                  height: 24,
                ),
                Center(child: buildUpgradeButton()),
                const SizedBox(
                  height: 24,
                ),
                buildAbout(snapshot.data!),
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
                            return EditPersonalDoctorScreen();
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
                        text: 'เปลี่ยนรหัสผ่าน',
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
                    ],
                  ),
                ),
              ],
            ),
          );
        } else {
          return Scaffold(
            body: Center(
              child: CircularProgressIndicator(),
            ),
          );
        }
      },
    );
  }

  Widget buildName(PInfo profile) => Column(
        children: [
          Text(
            profile.firstName + ' ' + profile.lastName,
            style: TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 24,
            ),
          ),
          const SizedBox(
            height: 4,
          ),
          Text(
            DateFormat('yyyy-MM-dd').format(profile.brithDate!).toString(),
            style: TextStyle(
                fontWeight: FontWeight.bold, fontSize: 15, color: Colors.grey),
          ),
        ],
      );

  Widget buildUpgradeButton() => ButtonWidget(
        onClicked: () {},
        text: text,
      );

  buildAbout(PInfo profile) => Container(
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
              profile.address,
              style: TextStyle(
                fontWeight: FontWeight.bold,
                fontSize: 15,
              ),
            ),
          ],
        ),
      );
}
