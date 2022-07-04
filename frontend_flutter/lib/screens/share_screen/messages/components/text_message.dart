import 'package:flutter/material.dart';
import 'package:jiffy/jiffy.dart';

import '../../../../constants.dart';
import '../../../../models/messages.dart';

class TextMessage extends StatelessWidget {
  const TextMessage({
    Key? key,
    required this.messages, required this.UserId,
  }) : super(key: key);

  final Messages messages;
  final int UserId;
  //final ChatMessage? message;

  @override
  Widget build(BuildContext context) {
    String CalculateTime() {
      var time = Jiffy(messages.sentDateTime).startOf(Units.SECOND).fromNow();
      return time;
    }

    return Column(
      children: [
        Text(
          CalculateTime().toString(),
          textAlign: messages.edges.whoSendMessages!.id == UserId ? TextAlign.end : TextAlign.start,
          style: TextStyle(
            fontSize: 12,
            color: kContentColorLightTheme.withOpacity(0.5),
          ),
        ),
        Container(
          //color:  Colors.white,
          padding: EdgeInsets.symmetric(
            horizontal: kDefaultPadding * 0.75,
            vertical: kDefaultPadding / 2,
          ),
          decoration: BoxDecoration(
            color: kPrimaryColor.withOpacity(0.1),
            borderRadius: BorderRadius.circular(30),
          ),
          child: Text(
            messages.messageText.toString(),
            style: TextStyle(
              color: Theme.of(context).textTheme.bodyText1!.color,
            ),
          ),
        ),
      ],
    );
  }
}
