import 'package:flutter/material.dart';
import 'package:frontend_flutter/models/pInfo.dart';
import 'package:frontend_flutter/screens/user_screen/main_user_screen.dart';
import 'package:http/http.dart' as http;
import '../../../constants.dart';
import '../../../models/messages.dart';
import 'components/body.dart';

class MessagesScreen extends StatefulWidget {
  const MessagesScreen(
      {Key? key, required this.chatId, required this.PInfoUser})
      : super(key: key);
  final int chatId;
  final PInfo PInfoUser;
  @override
  State<MessagesScreen> createState() => _MessagesScreenState();
}

class _MessagesScreenState extends State<MessagesScreen> {
  late Future<List<Messages>> _futureMessages;

  @override
  initState() {
    _futureMessages = _futureMessage();
    super.initState();
  }

  Future<List<Messages>> _futureMessage() async {
    var response = await http.get(Uri.parse(
        'http://10.0.2.2:8080/api/v1/messages/' + widget.chatId.toString()));
    // if (response.statusCode == 404 || response.reasonPhrase == "Not Found") {
    //   return false;
    // }
    final results = messagesFromJson(response.body);
    return results;
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
        future: _futureMessages,
        builder:
            (BuildContext context, AsyncSnapshot<List<Messages>> snapshot) {
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
              backgroundColor: Color.fromARGB(255, 208, 244, 255),
              appBar: buildAppBar(widget.PInfoUser),
              body: Body(widget.chatId, snapshot.data, widget.PInfoUser),
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

  AppBar buildAppBar(PInfo pInfoUser) {
    String Profile = pInfoUser.profile == null
        ? 'assets/images/Profile_Default.png'
        : pInfoUser.profile!;
    String ChatName = pInfoUser.firstName == null && pInfoUser.lastName == null
        ? ''
        : pInfoUser.firstName + ' ' + pInfoUser.lastName;
    return AppBar(
      backgroundColor: Color.fromARGB(232, 100, 180, 255),
      automaticallyImplyLeading: false,
      title: Row(
        children: [
          BackButton(
              // onPressed: () => Navigator.pushReplacement(context,
              //       MaterialPageRoute(builder: (context) {
              //     return MainUserScreen(ScreenIndex: 2,);
              //   })),
              ),
          CircleAvatar(
            backgroundImage: AssetImage(Profile),
          ),
          SizedBox(width: kDefaultPadding * 0.75),
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                ChatName,
                style: TextStyle(fontSize: 16),
              ),
            ],
          )
        ],
      ),
      // actions: [
      //   IconButton(
      //     icon: Icon(Icons.local_phone),
      //     onPressed: () {},
      //   ),
      //   IconButton(
      //     icon: Icon(Icons.videocam),
      //     onPressed: () {},
      //   ),
      //   SizedBox(width: kDefaultPadding / 2),
      // ],
    );
  }
}
