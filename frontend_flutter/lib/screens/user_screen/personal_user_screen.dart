import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:frontend_flutter/models/pInfo.dart';
import 'package:frontend_flutter/screens/user_screen/edit_personal_user_screen.dart';
import 'package:frontend_flutter/util/user_preferences.dart';
import 'package:frontend_flutter/widget/large_profile_widget.dart';
import 'package:frontend_flutter/widget/personal_button.dart';
import 'package:http/http.dart' as http;
import 'package:intl/intl.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../home_screen.dart';

class PersonalUserScreen extends StatefulWidget {
  const PersonalUserScreen({Key? key}) : super(key: key);

  @override
  State<PersonalUserScreen> createState() => _PersonalUserScreenState();
}

class _PersonalUserScreenState extends State<PersonalUserScreen> {
  final profile = UserPreferences.myProfile;
  final text = "แก้ไขโปรไฟล์ของคุณ";
  String _name = "FristName LastName";
  String _Profile = "assets/images/Profile_Default.png";
  var BD_format = DateFormat.yMMM();
  late Future<PInfo> pInfos;
  @override
  initState() {
    getName();
    pInfos = getPInfo();
    super.initState();
  }

  Future<void> getName() async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    final String? name = sharedPreferences.getString('username');
    setState(() => _name = name!);
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
      //setState(() => _UserId = results!);
      return await results;
    } else {
      throw ("User Not Found : " + response.reasonPhrase!);
    }
  }

  Future<PInfo> getPInfo() async {
    var UserId = await getUserId();
    var response = await http.get(
        Uri.parse('http://10.0.2.2:8080/api/v1/pinfos/' + UserId.toString()));
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
          _Profile = snapshot.data!.profile != null ? snapshot.data!.profile! :  _Profile;
          String IdCardNumber = snapshot.data == null
              ? 'ยังไม่มีข้อมูล'
              : snapshot.data!.idCardNumber;
          String BrithDate = snapshot.data!.brithDate == null
              ? 'ยังไม่มีข้อมูล'
              : BD_format.format(snapshot.data!.brithDate!);
          String Age = snapshot.data!.brithDate == null
              ? 'ยังไม่มีข้อมูล'
              : BD_format.format(DateTime(
                  DateTime.now().year - snapshot.data!.brithDate!.year));
          String Gender = 'ยังไม่มีข้อมูล';
          if (snapshot.data != null) {
            if (snapshot.data!.gender == 1) {
              Gender = 'ชาย';
            } else if (snapshot.data!.gender == 2) {
              Gender = 'หญิง';
            } else {
              Gender = 'ยังไม่มีข้อมูล';
            }
          }
          String BloodGroup = snapshot.data == null
              ? 'ยังไม่มีข้อมูล'
              : snapshot.data!.bloodGroup;
          String Address = snapshot.data!.address == null
              ? 'ยังไม่มีข้อมูล'
              : snapshot.data!.address!;
          String About = (snapshot.data!.about == '' || snapshot.data!.about == null)
              ? 'ยังไม่มีข้อมูล'
              : snapshot.data!.about!;
          return Scaffold(
            backgroundColor: Color.fromARGB(255, 208, 244, 255),
            body: ListView(
              physics: BouncingScrollPhysics(),
              children: [
                SizedBox(
                  height: 10,
                ),
                LargeProfileWidget(
                  imagePath: _Profile,
                  onClicked: () async {
                    Navigator.pushReplacement(context,
                        MaterialPageRoute(builder: (context) {
                      return EditPersonalUserScreen();
                    }));
                  },
                ),
                const SizedBox(
                  height: 24,
                ),
                buildName(snapshot.data!),
                buildData('เลขบัตรประชาชน', IdCardNumber),
                buildData('วันเดือนปีเกิด', BrithDate),
                buildData('อายุ', Age + ' ปี'),
                buildData('เพศ', Gender),
                buildData('กรุ๊ปเลือด', BloodGroup),
                buildDetails('ที่อยู่', Address),
                buildDetails('เกี่ยวกับ', About),
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
                            return EditPersonalUserScreen();
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
            _name,
            style: TextStyle(
                fontWeight: FontWeight.bold, fontSize: 15, color: Colors.grey),
          ),
        ],
      );

  buildData(String title, String data) => Container(
        padding: EdgeInsets.only(left: 48, top: 5, bottom: 5, right: 48),
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.center,
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Expanded(
              flex: 1,
              child: Text(
                title,
                style: TextStyle(
                  //fontWeight: FontWeight.bold,
                  fontSize: 16,
                ),
              ),
            ),
            Expanded(
              flex: 1,
              child: Text(
                data,
                style: TextStyle(
                  //fontWeight: FontWeight.bold,
                  color:
                      data != 'ยังไม่มีข้อมูล' ? Colors.black87 : Colors.grey,
                  fontSize: 16,
                ),
              ),
            ),
          ],
        ),
      );

  buildDetails(String title, String data) => Container(
        margin: const EdgeInsets.only(top: 5, bottom: 5,),
        padding: EdgeInsets.symmetric(horizontal: 48),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              title,
              style: TextStyle(
                //fontWeight: FontWeight.bold,
                fontSize: 18,
              ),
            ),
            const SizedBox(
              height: 16,
            ),
            Text(
              data,
              style: TextStyle(
                //fontWeight: FontWeight.bold,
                color: data != 'ยังไม่มีข้อมูล' ? Colors.black87 : Colors.grey,
                fontSize: 14,
              ),
            ),
          ],
        ),
      );
}
