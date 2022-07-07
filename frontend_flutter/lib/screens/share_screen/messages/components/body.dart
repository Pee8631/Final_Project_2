import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:frontend_flutter/models/messages.dart';
import 'package:frontend_flutter/models/pInfo.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';
import '../../../../constants.dart';
import 'chat_input_field.dart';
import 'message.dart';

class Body extends StatefulWidget {
  final int ChatId;
  final List<Messages>? data;
  final PInfo pInfoUser;
  const Body(this.ChatId, this.data, this.pInfoUser, {Key? key}) : super(key: key);
  //final _ChatId = chatId;
  @override
  State<Body> createState() => _BodyState();
}

class _BodyState extends State<Body> {
  late Future<int> _UserId;

  //final bool isExpired = .isBefore(now);
  //bool IsSender;
  @override
  initState() {
    _UserId = getUserId();
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

  @override
  Widget build(BuildContext context) {
    if (widget.data != null) {
      return FutureBuilder(
        future: _UserId,
        builder: (BuildContext context, AsyncSnapshot<int> snapshot) {
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
            String Profile = widget.pInfoUser.profile == null
        ? 'assets/images/Profile_Default.png'
        : widget.pInfoUser.profile!;
            return Column(
              children: [
                Expanded(
                  child: Padding(
                    padding:
                        const EdgeInsets.symmetric(horizontal: kDefaultPadding),
                    child: ListView.builder(
                      itemCount:
                          widget.data!.length,
                      itemBuilder: (context, index) => Message(
                          messages: widget.data![index] , UserId: snapshot.data!, Profile: Profile,),
                    ),
                  ),
                ),
                ChatInputField(widget.ChatId , widget.pInfoUser),
              ],
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
    } else {
      return Column(
        children: [
          Expanded(
            child: Padding(
              padding: const EdgeInsets.symmetric(horizontal: kDefaultPadding),
            ),
          ),
          ChatInputField(widget.ChatId, widget.pInfoUser),
        ],
      );
    }
  }
}
