import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:frontend_flutter/models/appointment.dart';
//import 'package:frontend_flutter/components/filled_outline_button.dart';
//import 'package:frontend_flutter/models/Chat.dart';
import 'package:frontend_flutter/models/chats.dart';
import 'package:frontend_flutter/models/pInfo.dart';
import 'package:frontend_flutter/screens/share_screen/messages/message_screen.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:http/http.dart' as http;

import '../../../../models/CreateChat.dart';
import 'chat_card.dart';

class Body extends StatefulWidget {
  const Body({Key? key}) : super(key: key);

  @override
  State<Body> createState() => _BodyState();
}

class _BodyState extends State<Body> {
  // String _name = 'FristName LastName';
  late Future<List<Chats>> _futureChats;
  CreateChat chats = new CreateChat();
  List<PInfo> pInfoUser = <PInfo>[];
  late int UserId;
  late int Role;

  @override
  initState() {
    setState(() {
      _futureChats = _futureChat();
      CheckRoleId();
    });
    super.initState();
  }

  Future<String> getToken() async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    final String? authToken = sharedPreferences.getString('authToken');
    //authToken?.substring(0, authToken.length - 1);
    return authToken!;
  }

  Future<int> getUserId() async {
    //FutureBuilder(future: getToken(), builder: (BuildContext context, AsyncSnapshot<void> snapshot) { },);
    String _authTokens = await getToken();
    var response = await http
        .get(Uri.parse('http://10.0.2.2:8080/api/v1/tokens/' + _authTokens));
    if (response.statusCode == 200) {
      final results = jsonDecode(response.body);
      setState(() => UserId = results);
      return results;
    } else {
      throw ("Not Found: User ID");
    }
  }

  Future<void> CheckRoleId() async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    final roleId = sharedPreferences.getInt('role');
    setState(() {
      Role = roleId!;
    });
  }

  Future<List<Chats>> _futureChat() async {
    int _getUserId = await getUserId();
    var response = await http.get(Uri.parse(
        'http://10.0.2.2:8080/api/v1/chats/' + _getUserId.toString()));

    final results = chatsFromJson(response.body);
    if (results.length > 0) {
      for (var i = 0; i < results.length; i++) {
        if (results[i].edges!.chatUser != null) {
          for (var j = 0; j < results[i].edges!.chatUser!.length; j++) {
            var findUser = results[i].edges!.chatUser![j];
            if (findUser.id != UserId) {
              var resultsUser = await _futurePinfo(findUser.id);
              pInfoUser.add(resultsUser);
            }
          }
        }
      }
    }

    return results;
  }

  Future<PInfo> _futurePinfo(int UserId) async {
    var response = await http.get(
        Uri.parse('http://10.0.2.2:8080/api/v1/pinfos/' + UserId.toString()));
    // if (response.statusCode == 404 || response.reasonPhrase == "Not Found") {
    //   return false;
    // }
    final results = pInfoFromJson(response.body);
    return await results;
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: _futureChats,
      builder: (BuildContext context, AsyncSnapshot<List<Chats>> snapshot) {
        if (snapshot.hasError) {
          return Scaffold(
            appBar: AppBar(
              title: Text("Error"),
            ),
            body: Center(
              child: Text("Error : ${snapshot.error} => ${pInfoUser.length}"),
            ),
          );
        } else if (snapshot.connectionState == ConnectionState.done) {
          if (snapshot.data!.length < 0) {
            return Scaffold(
              backgroundColor: Color.fromARGB(255, 208, 244, 255),
              appBar: AppBar(
                title: Text(
                  "แชท",
                  style: TextStyle(
                      fontSize: 20,
                      fontWeight: FontWeight.bold,
                      color: Colors.black87),
                ),
                centerTitle: true,
                backgroundColor: Colors.transparent,
                shadowColor: Colors.transparent,
              ),
              body: Column(
                children: [
                  Expanded(
                    child: ListView.builder(
                        itemCount: snapshot.data!.length,
                        itemBuilder: (context, index) {
                          Chats chats = snapshot.data![index];
                          if (chats.isLockChat == null) {
                            chats.isLockChat = false;
                          }
                          return ChatCard(
                              isChatsLock: chats.isLockChat,
                              chats: chats,
                              PInfoUser: pInfoUser[index],
                              press: () {
                                print(Role);
                                if (Role == 3) {
                                  if (chats.isLockChat) {
                                    Fluttertoast.showToast(
                                        msg: "ห้องล็อคนี้อยู่",
                                        gravity: ToastGravity.CENTER);
                                  }
                                }
                                Navigator.push(
                                  context,
                                  MaterialPageRoute(
                                    builder: (context) => MessagesScreen(
                                      chatId: chats.id,
                                      PInfoUser: pInfoUser[index],
                                    ),
                                  ),
                                );
                              });
                        }),
                  ),
                ],
              ),
            );
          } else {
            return Scaffold(
              backgroundColor: Color.fromARGB(255, 208, 244, 255),
              appBar: AppBar(
                title: Text(
                  "แชท",
                  style: TextStyle(
                      fontSize: 20,
                      fontWeight: FontWeight.bold,
                      color: Colors.black87),
                ),
                centerTitle: true,
                backgroundColor: Colors.transparent,
                shadowColor: Colors.transparent,
              ),
              body: Center(
                child: Text(
                  "คุณยังไม่มีห้องแชท",
                  style: TextStyle(
                      fontSize: 24,
                      fontWeight: FontWeight.bold,
                      color: Colors.black54),
                ),
              ),
            );
          }
        } else {
          return Scaffold(
            body: Center(
              child: CircularProgressIndicator(),
            ),
          );
        }
      },
    );
    // Column(
    //   children: [
    // Container(
    //   padding: EdgeInsets.fromLTRB(
    //       kDefaultPadding, 0, kDefaultPadding, kDefaultPadding),
    //   color: kPrimaryColor,
    //   child: Row(
    //     children: [
    //       FillOutlineButton(press: () {}, text: "Recent Message"),
    //       SizedBox(width: kDefaultPadding),
    //       FillOutlineButton(
    //         press: () {},
    //         text: "Active",
    //         isFilled: false,
    //       ),
    //     ],
    //   ),
    // ),

    //   ],
    // );
  }
}
