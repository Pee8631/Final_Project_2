import 'package:flutter/material.dart';
import 'package:frontend_flutter/screens/user_screen/doctor_profile_screen.dart';
import 'package:frontend_flutter/screens/user_screen/main_user_screen.dart';
import 'package:frontend_flutter/util/user_preferences.dart';
import 'package:http/http.dart' as http;

import '../../models/listUser.dart';
import '../../widget/appbar_doctors_screen.dart';


class DoctorsScreen extends StatefulWidget {
  final int DepartmentId;
  final String name;
  final int UserId;
  const DoctorsScreen(
      {Key? key, required this.name, required this.DepartmentId, required this.UserId})
      : super(key: key);

  @override
  State<DoctorsScreen> createState() => _DoctorsScreenState();
}

class _DoctorsScreenState extends State<DoctorsScreen> {
  final profile = UserPreferences.myProfile;
  PageController pageController = PageController();
  //int _selectedIndex = 0;
  final List<String> textList = [
    "He'd have you",
    "Heed not the rabble",
    "Sound of screams",
    "Who scream",
    "Revolution is",
    "Revolution, they..."
  ];

  late Future<List<ListUser>> getListDoctor;
  //late List<Department> _getDepartments;

  @override
  initState() {
    getListDoctor = _getListDoctor();
    super.initState();
  }
  //   void _onItemTapped(int index) {
  //   setState(() {
  //     _selectedIndex = index;
  //   });
  //   pageController.jumpToPage(index);
  // }

  Future<List<ListUser>> _getListDoctor() async {
    //FutureBuilder(future: getToken(), builder: (BuildContext context, AsyncSnapshot<void> snapshot) { },);
    var response = await http.get(Uri.parse(
        'http://10.0.2.2:8080/api/v1/users/getDepartment/' +
            widget.DepartmentId.toString()));
    if (response.statusCode == 200) {
      List<ListUser> results = listUserFromJson(response.body);
      return results;
    } else {
      throw ("User Not Found : " + response.reasonPhrase!);
    }
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: getListDoctor,
      builder: (BuildContext context, AsyncSnapshot<List<ListUser>> snapshot) {
        if (snapshot.hasError) {
          return Scaffold(
            appBar: AppBar(
              title: Text("Error"),
            ),
            body: Center(
              child: Text("Error : ${snapshot.error}"),
            ),
          );
        } else if (snapshot.connectionState == ConnectionState.done) {
          return Scaffold(
            appBar: buildAppBarBackToScreen(context, widget.name, MainUserScreen(ScreenIndex: 0,)),
            // appBar: AppBar(
            //   backgroundColor: Colors.transparent,
            //   elevation: 0,
            //   title: Center(
            //       child:
            //     Text(
            //   "รายชื่อ คุณหมอ",
            //   style: TextStyle(fontSize: 20, color: Colors.black),
            // ),
            //   ),
            // ),
            body: ListView.separated(
              padding: const EdgeInsets.all(8.0),
              itemCount: snapshot.data!.length,
              itemBuilder: (BuildContext context, int index) {
                return ListTile(
                  selectedColor: Colors.grey,
                  onTap: () {
                    Navigator.pushReplacement(context,
                        MaterialPageRoute(builder: (context) {
                      return DoctorProfileScreen(DoctorId: snapshot.data![index].id, Name: widget.name, UserId: widget.UserId);
                    }));
                  },
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
                        snapshot.data![index].edges!.userHasPInfo![0]
                                .firstName +
                            " " +
                            snapshot
                                .data![index].edges!.userHasPInfo![0].lastName,
                        style: TextStyle(fontSize: 12, color: Colors.black),
                        textAlign: TextAlign.left,
                      ),
                      Text(
                        snapshot.data![index].edges!.fromHospital!.name,
                        style: TextStyle(fontSize: 12, color: Colors.black),
                        textAlign: TextAlign.left,
                      ),
                      Text(
                        snapshot.data![index].edges!.hasDepartment!.name,
                        style: TextStyle(fontSize: 12, color: Colors.black),
                        textAlign: TextAlign.left,
                      ),
                      Text(
                        "Price$index",
                        style: TextStyle(fontSize: 12, color: Colors.black),
                        textAlign: TextAlign.left,
                      ),
                      Text(
                        "about : ${textList[index]}",
                        style: TextStyle(fontSize: 12, color: Colors.black),
                        textAlign: TextAlign.left,
                      ),
                    ],
                  ),
                );
              },
              separatorBuilder: (BuildContext context, int index) =>
                  const Divider(color: Colors.black),
            ),
            // bottomNavigationBar: BottomNavigationBar(
            //   items: const <BottomNavigationBarItem>[
            //     BottomNavigationBarItem(
            //       icon: Icon(Icons.home),
            //       label: 'หน้าหลัก',
            //       backgroundColor: Colors.blue,
            //     ),
            //     BottomNavigationBarItem(
            //       icon: Icon(Icons.schedule),
            //       label: 'ตาราง(ชั่วคราว)',
            //       backgroundColor: Colors.blue,
            //     ),
            //     BottomNavigationBarItem(
            //       icon: Icon(Icons.chat),
            //       label: 'ห้องแชท',
            //       backgroundColor: Colors.blue,
            //     ),
            //     BottomNavigationBarItem(
            //       icon: Icon(Icons.notifications_none),
            //       label: 'แจ้งเตือน',
            //       backgroundColor: Colors.blue,
            //     ),
            //     BottomNavigationBarItem(
            //       icon: Icon(Icons.person),
            //       label: 'ส่วนตัว',
            //       backgroundColor: Colors.blue,
            //     ),
            //   ],
            //   currentIndex: _selectedIndex,
            //   selectedItemColor: Colors.white,
            //   unselectedItemColor: Colors.white,
            //   onTap: _onItemTapped,
            // ),
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
}
