import 'package:flutter/material.dart';
import 'package:frontend_flutter/models/doctor.dart';
import 'package:frontend_flutter/screens/user_screen/appointment_screen.dart';
import 'package:frontend_flutter/screens/user_screen/doctors_screen.dart';
import 'package:frontend_flutter/widget/appbar_doctors_screen.dart';
import 'package:http/http.dart' as http;

class DoctorProfileScreen extends StatefulWidget {
  final int DoctorId;
  final String Name;
  final int UserId;
  final String Profile;
  const DoctorProfileScreen(
      {Key? key,
      required this.DoctorId,
      required this.Name,
      required this.UserId,
      required this.Profile})
      : super(key: key);

  @override
  State<DoctorProfileScreen> createState() => _DoctorProfileScreenState();
}

class _DoctorProfileScreenState extends State<DoctorProfileScreen> {
  late Future<Doctor> _user;
  String Profile = 'assets/images/Profile_Default.png';
  @override
  initState() {
    _user = getDoctor();
    super.initState();
  }

  Future<Doctor> getDoctor() async {
    var response = await http.get(Uri.parse(
        'http://10.0.2.2:8080/api/v1/users/' + widget.DoctorId.toString()));
    if (response.statusCode == 200) {
      final results = doctorFromJson(response.body);
      return results;
    } else {
      throw ("PInfo Not Found: " + response.reasonPhrase!);
    }
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: _user,
      builder: (BuildContext context, AsyncSnapshot<Doctor> snapshot) {
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
          Profile = widget.Profile;
          return Scaffold(
            backgroundColor: Color.fromARGB(255, 208, 244, 255),
            appBar: buildAppBarBackToScreen(
                context,
                widget.Name,
                DoctorsScreen(
                  DepartmentId: snapshot.data!.edges!.hasDepartment!.id,
                  name: widget.Name,
                  UserId: widget.UserId,
                  Profile: Profile,
                ),
                Profile),
            body: Container(
                margin: const EdgeInsets.all(20),
                child: buildDoctorData(snapshot.data!)),
            bottomNavigationBar: BottomAppBar(
              color: Colors.transparent,
              child: buildButton(widget.Name, widget.DoctorId),
              elevation: 0,
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

  Widget buildDoctorData(Doctor doctor) {
    //_Profile = snapshot.data!.profile != null ? _Profile : snapshot.data!.profile!;
    var Profile;
    if (doctor.edges!.userHasPInfo != null) {
      Profile = doctor.edges!.userHasPInfo![0].profile != null
          ? doctor.edges!.userHasPInfo![0].profile
          : 'assets/images/Profile_Default.png';
    }

    return Container(
      margin: const EdgeInsets.all(10.0),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.start,
        children: [
          ClipRRect(
            borderRadius: BorderRadius.circular(8.0),
            child: Material(
              color: Colors.transparent,
              child: Ink.image(
                image:
                    AssetImage(Profile ?? 'assets/images/Profile_Default.png'),
                fit: BoxFit.cover,
                width: 250,
                height: 250,
              ),
            ),
          ),
          Container(
              width: double.infinity,
              margin: const EdgeInsets.only(top: 20, bottom: 20),
              child: buildDoctorInfo(doctor)),
        ],
      ),
    );
  }

  Widget buildDoctorInfo(Doctor doctor) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Expanded(
          flex: 1,
          child: buildTitle(),
        ),
        Expanded(
          flex: 2,
          child: buildData(doctor),
        ),
      ],
    );
  }

  Widget buildTitle() {
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        buildText("ชื่อ"),
        buildText("โรงพยาบาล"),
        buildText("แผนก"),
        buildText("เกี่ยวกับ "),
      ],
    );
  }

  Widget buildData(Doctor doctor) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        buildText(
          doctor.edges!.userHasPInfo![0].prefix! +
              " " +
              doctor.edges!.userHasPInfo![0].firstName +
              " " +
              doctor.edges!.userHasPInfo![0].lastName,
        ),
        buildText(
          doctor.edges!.fromHospital!.name,
        ),
        buildText(
          doctor.edges!.hasDepartment!.name,
        ),
        buildText(
          doctor.edges!.userHasPInfo![0].about!,
        ),
      ],
    );
  }

  Widget buildText(String text) {
    return Container(
      padding: const EdgeInsets.all(4.0),
      width: double.maxFinite,
      child: Text(
        text,
        style: TextStyle(fontSize: 16, color: Colors.black),
      ),
    );
  }

  Widget buildButton(String name, int doctorId) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: SizedBox(
        width: double.maxFinite,
        child: ElevatedButton(
          style: ElevatedButton.styleFrom(
            primary: Color.fromARGB(220, 96, 239, 220),
            onPrimary: Colors.white,
            shadowColor: Colors.greenAccent,
            elevation: 3,
            shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(32.0)),
            maximumSize: Size(100, 40), //////// HERE
          ),
          child: Text('นัดหมาย'),
          onPressed: () {
            Navigator.pushReplacement(context,
                MaterialPageRoute(builder: (context) {
              return AppointmentScreen(
                  Name: name,
                  DoctorId: doctorId,
                  UserId: widget.UserId,
                  Profile: Profile);
            }));
          },
        ),
      ),
    );
  }
}
