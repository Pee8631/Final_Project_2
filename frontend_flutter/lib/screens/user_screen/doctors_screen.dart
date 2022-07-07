import 'package:flutter/material.dart';
import 'package:frontend_flutter/screens/user_screen/doctor_profile_screen.dart';
import 'package:frontend_flutter/screens/user_screen/main_user_screen.dart';
import 'package:frontend_flutter/util/user_preferences.dart';
import 'package:http/http.dart' as http;

import '../../models/listDoctor.dart';
import '../../widget/appbar_doctors_screen.dart';

class DoctorsScreen extends StatefulWidget {
  final int DepartmentId;
  final String name;
  final int UserId;
  final String Profile;
  const DoctorsScreen(
      {Key? key,
      required this.name,
      required this.DepartmentId,
      required this.UserId,
      required this.Profile})
      : super(key: key);

  @override
  State<DoctorsScreen> createState() => _DoctorsScreenState();
}

class _DoctorsScreenState extends State<DoctorsScreen> {
  final profile = UserPreferences.myProfile;
  String Profile = 'assets/images/Profile_Default.png';
  late Future<List<ListDoctor>> getListDoctor;

  @override
  initState() {
    getListDoctor = _getListDoctor();
    super.initState();
  }

  Future<List<ListDoctor>> _getListDoctor() async {
    //FutureBuilder(future: getToken(), builder: (BuildContext context, AsyncSnapshot<void> snapshot) { },);
    var response = await http.get(Uri.parse(
        'http://10.0.2.2:8080/api/v1/users/getDepartment/' +
            widget.DepartmentId.toString()));
    if (response.statusCode == 200) {
      List<ListDoctor> results = listDoctorFromJson(response.body);
      return results;
    } else {
      throw ("User Not Found : " + response.reasonPhrase!);
    }
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: getListDoctor,
      builder:
          (BuildContext context, AsyncSnapshot<List<ListDoctor>> snapshot) {
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
          Profile = widget.Profile;
          return Scaffold(
            backgroundColor: Color.fromARGB(255, 208, 244, 255),
            appBar: buildAppBarBackToScreen(
                context,
                widget.name,
                MainUserScreen(
                  ScreenIndex: 0,
                ),
                Profile),
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
                var doctorData = snapshot.data![index];
                return ListTile(
                  selectedColor: Colors.grey,
                  onTap: () {
                    Navigator.pushReplacement(context,
                        MaterialPageRoute(builder: (context) {
                      return DoctorProfileScreen(
                          DoctorId: snapshot.data![index].id!,
                          Name: widget.name,
                          UserId: widget.UserId,
                          Profile: Profile);
                    }));
                  },
                  leading: ClipOval(
                    child: Image.asset(
                      doctorData.edges!.userHasPInfo![0].profile ?? '',
                      width: 60,
                      height: 60,
                    ),
                  ),
                  title: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      buildText(
                          "ชื่อ",
                          doctorData.edges!.userHasPInfo![0].prefix! +
                              " " +
                              doctorData.edges!.userHasPInfo![0].firstName! +
                              " " +
                              doctorData.edges!.userHasPInfo![0].lastName!),
                      buildText(
                          "โรงพยาบาล", doctorData.edges!.fromHospital!.name!),
                      buildText("แผนก", doctorData.edges!.hasDepartment!.name!),
                    ],
                  ),
                );
              },
              separatorBuilder: (BuildContext context, int index) =>
                  const Divider(color: Colors.black),
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

  Widget buildText(String title, String description) {
    return Row(
      children: [
        Expanded(
          flex: 1,
          child: Text(
            title,
            style: TextStyle(fontSize: 14, color: Colors.black87),
          ),
        ),
        Expanded(
          flex: 2,
          child: Text(
            description,
            style: TextStyle(fontSize: 14, color: Colors.black87),
          ),
        ),
      ],
    );
  }
}
