import 'package:flutter/material.dart';
import 'package:frontend_flutter/models/department.dart';
import 'package:frontend_flutter/screens/user_screen/doctors_screen.dart';
import 'package:frontend_flutter/util/user_preferences.dart';
import 'package:frontend_flutter/widget/menu_button.dart';
import 'package:http/http.dart' as http;

class MenuScreen extends StatefulWidget {
  final String name;
  final int UserId;
  const MenuScreen({Key? key, required this.name, required this.UserId}) : super(key: key);

  @override
  State<MenuScreen> createState() => _MenuScreenState();
}

class _MenuScreenState extends State<MenuScreen> {
  final profile = UserPreferences.myProfile;
  late Future<List<Department>> _futureGetDepartments;
  //late List<Department> _getDepartments;

  @override
  initState() {
    _futureGetDepartments = _getDepartments();
    super.initState();
  }

  Future<List<Department>> _getDepartments() async {
    //FutureBuilder(future: getToken(), builder: (BuildContext context, AsyncSnapshot<void> snapshot) { },);
    var response =
        await http.get(Uri.parse('http://10.0.2.2:8080/api/v1/departments'));
    if (response.statusCode == 200) {
      final results = departmentsFromJson(response.body);
      // setState(() {
      //   _getDepartments = results.toList();
      // });
      return results;
    } else {
      throw ("Not Found");
    }
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
        future: _futureGetDepartments,
        builder:
            (BuildContext context, AsyncSnapshot<List<Department>> snapshot) {
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
              backgroundColor: Color(0xBAD9EB),
              body: Padding(
                padding: const EdgeInsets.all(10.0),
                child: CustomScrollView(
                  slivers: <Widget>[
                    const SliverAppBar(
                      pinned: true,
                      elevation: 0,
                      expandedHeight: 50.0,
                      backgroundColor: Colors.transparent,
                      title: Center(
                        child: Text(
                          'แผนก',
                          style: TextStyle(
                            fontSize: 20,
                            color: Colors.black54,
                          ),
                        ),
                      ),
                    ),
                    SliverGrid(
                      gridDelegate:
                          const SliverGridDelegateWithMaxCrossAxisExtent(
                        maxCrossAxisExtent: 200.0,
                        mainAxisSpacing: 20.0,
                        mainAxisExtent: 150.0,
                        crossAxisSpacing: 15.0,
                        childAspectRatio: 2.0,
                      ),
                      delegate: SliverChildBuilderDelegate(
                        (BuildContext context, int index) {
                          var b = snapshot.data![index].name;
                          return Container(
                            alignment: Alignment.center,
                            child: Column(
                              children: [
                                MenuButtonWidget(
                                  imagePath:
                                      snapshot.data![index].image,
                                  onClicked: () async {
                                    Navigator.pushReplacement(context,
                                      MaterialPageRoute(builder: (context) {
                                    return DoctorsScreen(name: widget.name, DepartmentId: snapshot.data![index].id, UserId: widget.UserId);
                                  }));
                                  },
                                ),
                                Text(
                                  b,
                                  style: TextStyle(
                                      fontSize: 20, color: Colors.grey[700]),
                                ),
                              ],
                            ),
                          );
                        },
                        childCount: snapshot.data?.length,
                      ),
                    ),
                  ],
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
