import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

import '../../../../constants.dart';

import '../../../../models/messages.dart';
import 'audio_message.dart';
import 'text_message.dart';
import 'video_message.dart';

class Message extends StatefulWidget {
  final Messages messages;
  final int UserId;
  final String Profile;
  const Message({Key? key, required this.messages, required this.UserId , required this.Profile}) : super(key: key);

  @override
  State<Message> createState() => _MessageState();
}

class _MessageState extends State<Message> {

  @override
  Widget build(BuildContext context) {
    Widget messageContaint(Messages? messages) {
      return TextMessage(messages: messages! , UserId: widget.UserId);
      // switch (message.messageType) {
      //   case ChatMessageType.text:
      //     return TextMessage(messages: messages, message: message);
      //   case ChatMessageType.audio:
      //     return AudioMessage(message: message);
      //   case ChatMessageType.video:
      //     return VideoMessage();
      //   default:
      //     return SizedBox();
      // }
    }

    return Padding(
            padding: const EdgeInsets.only(top: kDefaultPadding),
            child: Row(
              mainAxisAlignment:
                  widget.messages.edges.whoSendMessages!.id == widget.UserId
                      ? MainAxisAlignment.end
                      : MainAxisAlignment.start,
              children: [
                if (widget.messages.edges.whoSendMessages!.id !=
                    widget.UserId) ...[
                  CircleAvatar(
                    radius: 12,
                    backgroundImage: AssetImage(widget.Profile),
                  ),
                ],

                SizedBox(width: kDefaultPadding / 2),
                messageContaint(widget.messages),
                // if (!message.isSender) ...[
                //   CircleAvatar(
                //     radius: 12,
                //     backgroundImage: AssetImage("assets/images/user_2.png"),
                //   ),
                //   SizedBox(width: kDefaultPadding / 2),
                // ],

                // if (message.isSender) MessageStatusDot(status: message.messageStatus)
              ],
            ),
          );
  }
}

class MessageStatusDot extends StatelessWidget {
  //final MessageStatus? status;

  const MessageStatusDot({
    Key? key,
    /*this.status*/
  }) : super(key: key);
  @override
  Widget build(BuildContext context) {
    // Color dotColor(MessageStatus status) {
    //   switch (status) {
    //     case MessageStatus.not_sent:
    //       return kErrorColor;
    //     case MessageStatus.not_view:
    //       return Theme.of(context).textTheme.bodyText1!.color!.withOpacity(0.1);
    //     case MessageStatus.viewed:
    //       return kPrimaryColor;
    //     default:
    //       return Colors.transparent;
    //   }
    // }

    return Container(
      margin: EdgeInsets.only(left: kDefaultPadding / 2),
      height: 12,
      width: 12,
      decoration: BoxDecoration(
        //color: dotColor(status!),
        shape: BoxShape.circle,
      ),
      // child: Icon(
      //   status == MessageStatus.not_sent ? Icons.close : Icons.done,
      //   size: 8,
      //   color: Theme.of(context).scaffoldBackgroundColor,
      // ),
    );
  }
}
