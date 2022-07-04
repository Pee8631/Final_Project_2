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
  const DoctorProfileScreen(
      {Key? key, required this.DoctorId, required this.Name, required this.UserId})
      : super(key: key);

  @override
  State<DoctorProfileScreen> createState() => _DoctorProfileScreenState();
}

class _DoctorProfileScreenState extends State<DoctorProfileScreen> {
  late Future<Doctor> _user;

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
          return Scaffold(
            appBar: buildAppBarBackToScreen(
                context,
                widget.Name,
                DoctorsScreen(
                  DepartmentId: snapshot.data!.edges!.hasDepartment!.id,
                  name: widget.Name, UserId: widget.UserId,
                )),
            body: ListView(
              physics: BouncingScrollPhysics(),
              children: [
                buildDoctorData(snapshot.data!),
                buildButton(widget.Name, widget.DoctorId),
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

  Widget buildDoctorData(Doctor doctor) {
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
                image: NetworkImage(
                    'https://i.pinimg.com/474x/65/25/a0/6525a08f1df98a2e3a545fe2ace4be47.jpg'),
                fit: BoxFit.cover,
                width: 250,
                height: 250,
              ),
            ),
          ),
          buildDoctorInfo(doctor),
        ],
      ),
    );
  }

    
  Widget buildDoctorInfo(Doctor doctor) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        buildTitle(),
        buildData(doctor),
    ],);
  }

  Widget buildTitle() {
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        buildText("ชื่อ"),
        buildText("โรงพยาบาล"),
        buildText("แผนก"),
        buildText("วุฒิบัตร"),
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
          "น.พ." +
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
          doctor.edges!.doctorHasCertification![0].code,
        ),
        buildText(
          doctor.edges!.hasDepartment!.name,
        ),
      ],
    );
  }

  Widget buildText(String text) {
    return Container(
      padding: const EdgeInsets.all(8.0),
      alignment: Alignment.centerLeft,
      child: Text(
        text,
        style: TextStyle(fontSize: 16, color: Colors.black),
        textAlign: TextAlign.left,
      ),
    );
  }
  
  Widget buildButton(String name, int doctorId) {
    return SizedBox(
      width: double.infinity,
      child: ElevatedButton(
        child: Text(
          'นัดหมาย',
          style: TextStyle(fontSize: 16),
        ),
        onPressed: () {
          Navigator.pushReplacement(context,
              MaterialPageRoute(builder: (context) {
            return AppointmentScreen(Name: name, DoctorId: doctorId, UserId: widget.UserId);
          }));
        },
      ),
    );
  }
}
