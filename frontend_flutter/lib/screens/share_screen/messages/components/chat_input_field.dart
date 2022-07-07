import 'dart:convert';
import 'dart:io';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:frontend_flutter/models/appointment.dart';
import 'package:frontend_flutter/models/chats.dart';
import 'package:frontend_flutter/models/pInfo.dart';
import 'package:frontend_flutter/screens/share_screen/messages/message_screen.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';
import '../../../../constants.dart';

class ChatInputField extends StatefulWidget {
  final int ChatId;
  final  PInfo pInfoUser;
  const ChatInputField(this.ChatId, this.pInfoUser, {Key? key}) : super(key: key);

  @override
  State<ChatInputField> createState() => _ChatInputFieldState();
}

class _ChatInputFieldState extends State<ChatInputField> {
  final formkey = GlobalKey<FormState>();
  bool isChatLock = true;
  var _message;
  // late Future<int> _UserId;
  @override
  initState() {
    setState(() {
      _futureChat();
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
    String authToken = await getToken();
    var response = await http
        .get(Uri.parse('http://10.0.2.2:8080/api/v1/tokens/' + authToken));
    if (response.statusCode == 200) {
      final results = jsonDecode(response.body);
      return results == null ? 0 : results;
    } else {
      throw ("User ID not found: " + authToken);
    }
  }

  Future<void> _futureChat() async {
    var response = await http.get(Uri.parse(
        'http://10.0.2.2:8080/api/v1/chats/chatroom/' + widget.ChatId.toString()));
    final results = chatFromJson(response.body);
    setState((){
      if(results.isLockChat == null ){
        results.isLockChat == false;
      }
      isChatLock = results.isLockChat;
    });
  }

  Future<void> SendMessage(String message) async {
    int UserId = await getUserId();
    // String _formatdate = new DateFormat('yyyy-MM-ddTHH:mm:ss.mmmuuuZ').format(DateTime.now());
    await http.post(
      Uri.parse('http://10.0.2.2:8080/api/v1/messages'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        "MessageText": message,
        "SentDateTime": DateTime.now().toIso8601String() + "Z",
        "ChatMessage": widget.ChatId,
        "UserMessage": UserId
      }),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.symmetric(
        horizontal: kDefaultPadding,
        vertical: kDefaultPadding / 2,
      ),
      decoration: BoxDecoration(
        color: Theme.of(context).scaffoldBackgroundColor,
        boxShadow: [
          BoxShadow(
            offset: Offset(0, 4),
            blurRadius: 32,
            color: Color(0xFF087949).withOpacity(0.08),
          ),
        ],
      ),
      child: SafeArea(
        child: Row(
          children: [
            // Icon(Icons.mic, color: kPrimaryColor),
            SizedBox(width: kDefaultPadding),
            Expanded(
              child: Container(
                padding: EdgeInsets.symmetric(
                  horizontal: kDefaultPadding * 0.75,
                ),
                decoration: BoxDecoration(
                  color: kPrimaryColor.withOpacity(0.05),
                  borderRadius: BorderRadius.circular(40),
                ),
                child: Form(
                  key: formkey,
                  child: Row(
                    children: [
                      // Icon(
                      //   Icons.sentiment_satisfied_alt_outlined,
                      //   color: Theme.of(context)
                      //       .textTheme
                      //       .bodyText1!
                      //       .color!
                      //       .withOpacity(0.64),
                      // ),
                      SizedBox(width: kDefaultPadding / 4),
                      Expanded(
                        child: TextFormField(
                          autofocus: true,
                          readOnly: isChatLock,
                          decoration: InputDecoration(
                            hintText: "Type message",
                            border: InputBorder.none,
                          ),
                          onChanged: (Message) {
                            _message = Message;
                          },
                        ),
                      ),
                      // Icon(
                      //   Icons.attach_file,
                      //   color: Theme.of(context)
                      //       .textTheme
                      //       .bodyText1!
                      //       .color!
                      //       .withOpacity(0.64),
                      // ),
                      // SizedBox(width: kDefaultPadding / 4),
                      // Icon(
                      //   Icons.camera_alt_outlined,
                      //   color: Theme.of(context)
                      //       .textTheme
                      //       .bodyText1!
                      //       .color!
                      //       .withOpacity(0.64),
                      // ),
                      ElevatedButton(
                          style: ButtonStyle(
                              shadowColor: MaterialStateProperty.all<Color>(
                                  Colors.transparent),
                              backgroundColor: MaterialStateProperty.all<Color>(
                                  Colors.transparent)),
                          child: Icon(
                            Icons.send,
                            color: Theme.of(context)
                                .textTheme
                                .bodyText1!
                                .color!
                                .withOpacity(0.64),
                          ),
                          onPressed: () async {
                            try {
                              if (formkey.currentState!.validate()) {
                                if (isChatLock) {
                                  Fluttertoast.showToast(
                                      msg: "ห้องล็อคนี้อยู่",
                                      gravity: ToastGravity.CENTER);
                                } else {
                                  if (_message != null) {
                                    SendMessage(_message).then((value) =>
                                        Navigator.pushReplacement(context,
                                            MaterialPageRoute(
                                                builder: (context) {
                                          return MessagesScreen(
                                              chatId: widget.ChatId, PInfoUser: widget.pInfoUser, );
                                        })));
                                  }
                                }
                              }
                              formkey.currentState!.reset();
                            } on HttpException catch (error) {
                              throw (error);
                            }
                          }),
                    ],
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
