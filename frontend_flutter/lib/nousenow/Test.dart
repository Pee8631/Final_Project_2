import 'package:http/http.dart' as http;
import 'package:flutter/material.dart';
import 'package:frontend_flutter/model/user.dart';

class Test extends StatefulWidget {
  const Test({Key? key}) : super(key: key);

  @override
  _Test createState() => _Test();
}

class _Test extends State<Test> {
  final formkey = GlobalKey<FormState>();
  User user = User(username: '', password: '', department: 0 , hospital: 0);

  get storage => null;

  void initState() {
    super.initState();
    getUser();
  }

  Future<User?> getUser() async {
    var url = "http://10.0.2.2:8080/api/v1/users/4";
    var response = await http.get(Uri.parse(url));
    print(response.body);
    User? DataFromAPI = userFromJson(response.body);
    return DataFromAPI;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text("Test AuthUser"),
        ),
        body: FutureBuilder(
          future: getUser(),
          builder: (BuildContext context, AsyncSnapshot<dynamic> snapshot) {
            if (snapshot.connectionState == ConnectionState.done) {
              var result = snapshot.data;
              return ListView(
                children: [
                  ListTile(
                    title: Text(result.username),
                  ),
                  ListTile(
                    title: Text(result.password),
                  ),
                ],
              );
            }
            return LinearProgressIndicator();
          },
        ));
  }
}
