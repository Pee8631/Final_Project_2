import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:frontend_flutter/models/department.dart';
import 'package:frontend_flutter/models/hospital.dart';
import 'package:frontend_flutter/models/user.dart';
import 'package:frontend_flutter/screens/signupandsignin/certification_doctor_screen.dart';
import 'package:frontend_flutter/screens/signupandsignin/pinfo_doctors_screen.dart';
import 'package:frontend_flutter/widget/appbor_logindoctorscreen.dart';
import 'package:http/http.dart' as http;

class DepHosDoctorScreen extends StatefulWidget {
  final User user;
  const DepHosDoctorScreen({Key? key, required this.user}) : super(key: key);

  @override
  State<DepHosDoctorScreen> createState() => _DepHosDoctorScreenState();
}

class _DepHosDoctorScreenState extends State<DepHosDoctorScreen> {
  final formkey = GlobalKey<FormState>();
  User _user = User(
      username: '',
      password: '',
      department: 0,
      hospital: 0,
      roleId: 0,
      certification: null,
      pInfo: null);
  late Future<List<Department>> _department;
  //['Ears', 'Eyes', 'Nose', 'Mouth', 'Kid']
  late Future<List<Hospital>> _hospital;
  //['example hospital', 'example 2 Hospital']

  @override
  initState() {
    _department = _getDepartment();
    _hospital = _getHospital();
    setState(() {
      _user = widget.user;
    });
    super.initState();
    // _getHospital().then((value) => setState(() => _hospital = value));
    // _getDepartment().then((value) => setState(() => _department = value));
  }

  Future<List<Department>> _getDepartment() async {
    final url = Uri.parse('http://10.0.2.2:8080/api/v1/departments');
    final response = await http.get(url);
    print(response.statusCode);
    print(response.body);
    final results = departmentsFromJson(response.body);
    return results;
    //setState(() => _department = results);
    //return await results;
  }

  Future<List<Hospital>> _getHospital() async {
    final url = Uri.parse('http://10.0.2.2:8080/api/v1/hospitals');
    final response = await http.get(url);
    print(response.statusCode);
    print(response.body);
    final results = hospitalFromJson(response.body);
    return results;
    //setState(() => _hospital = results);
  }

  late String valueDepartment;
  late String valueHospital;
  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
        future: Future.wait([_department, _hospital]),
        builder: (BuildContext context, AsyncSnapshot<List<dynamic>> snapshot) {
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
            final List<Department> department = snapshot.data?[0];
            final List<Hospital> hospital = snapshot.data?[1];
            valueDepartment = department[0].id.toString();
            valueHospital = hospital[0].id.toString();

            return Scaffold(
              backgroundColor: Color.fromARGB(255, 208, 244, 255),
              appBar: buildAppBarDoctorSignUp(
                  context,
                  PInfoDoctorScreen(
                    user: widget.user,
                  )),
              body: Form(
                key: formkey,
                child: Container(
                  padding: const EdgeInsets.all(20),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Container(
                        margin: const EdgeInsets.only(top: 10, bottom: 10),
                        child: Text(
                          'แผนก',
                          style: TextStyle(fontSize: 18),
                        ),
                      ),
                      Container(
                        padding:
                            EdgeInsets.symmetric(horizontal: 12, vertical: 4),
                        decoration: BoxDecoration(
                          borderRadius: BorderRadius.circular(12),
                          border: Border.all(color: Colors.black, width: 4),
                        ),
                        child: DropdownButtonHideUnderline(
                          child: new DropdownButton<String>(
                            value: valueDepartment,
                            isExpanded: true,
                            items: department
                                .map(
                                  (value) => new DropdownMenuItem<String>(
                                    value: value.id.toString(),
                                    child: new Text(
                                      value.name.toString(),
                                      style: TextStyle(
                                          fontWeight: FontWeight.bold,
                                          fontSize: 16),
                                    ),
                                  ),
                                )
                                .toList(),
                            onChanged: (valueDepartment) => setState(() {
                              print(valueDepartment);
                              this.valueDepartment = valueDepartment!;
                              _user.department = int.parse(valueDepartment);
                            }),
                          ),
                        ),
                      ),
                      Container(
                        margin: const EdgeInsets.only(top: 10, bottom: 10),
                        child: Text(
                          'โรงพยาบาล',
                          style: TextStyle(fontSize: 18),
                        ),
                      ),
                      Container(
                        padding:
                            EdgeInsets.symmetric(horizontal: 12, vertical: 4),
                        decoration: BoxDecoration(
                          borderRadius: BorderRadius.circular(12),
                          border: Border.all(color: Colors.black, width: 4),
                        ),
                        child: DropdownButtonHideUnderline(
                          child: new DropdownButton<String>(
                            value: valueHospital,
                            isExpanded: true,
                            items: hospital
                                .map(
                                  (value) => new DropdownMenuItem<String>(
                                    value: value.id.toString(),
                                    child: new Text(
                                      value.name.toString(),
                                      style: TextStyle(
                                          fontWeight: FontWeight.bold,
                                          fontSize: 16),
                                    ),
                                  ),
                                )
                                .toList(),
                            onChanged: (valueHospital) => setState(() {
                              print(valueHospital);
                              this.valueHospital = valueHospital!;
                              _user.hospital = int.parse(valueHospital);
                            }),
                          ),
                        ),
                      ),
                      Container(
                        margin: const EdgeInsets.only(top: 20.0, bottom: 10.0),
                        width: double.infinity,
                        height: 40,
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
                          child: Text("ต่อไป", style: TextStyle(fontSize: 16)),
                          onPressed: () async {
                            if (formkey.currentState!.validate()) {
                              formkey.currentState!.save();
                              try {
                                Navigator.pushReplacement(context,
                                    MaterialPageRoute(builder: (context) {
                                  return CreateCertificationScreen(
                                    user: _user,
                                  );
                                }));
                              } catch (error) {
                                Fluttertoast.showToast(
                                    msg: error.toString(),
                                    gravity: ToastGravity.CENTER);
                              }
                              formkey.currentState!.reset();
                            }
                          },
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            );
          } else {
            return Scaffold(
              body: Center(
                child: CircularProgressIndicator(),
              ),
            );
          }
        });
  }
}
