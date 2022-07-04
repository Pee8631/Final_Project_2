import 'package:flutter/material.dart';
import 'package:jiffy/jiffy.dart';

import '../../../../constants.dart';
import '../../../../models/chats.dart';

class ChatCard extends StatelessWidget {
  const ChatCard({
    Key? key,
    required this.isChatsLock,
    required this.press,
    required this.chats,
  }) : super(key: key);

  // final CreateChat chat;
  final Chats chats;
  final VoidCallback press;
  final bool isChatsLock;

  @override
  Widget build(BuildContext context) {
    String CalculateTime() {
      var time = Jiffy(chats.edges!
              .chatMessage![chats.edges!.chatMessage!.length - 1].sentDateTime)
          .startOf(Units.MILLISECOND)
          .fromNow();
      return time;
    }

    return InkWell(
      onTap: press,
      child: Padding(
        padding: const EdgeInsets.symmetric(
            horizontal: kDefaultPadding, vertical: kDefaultPadding * 0.75),
        child: Row(
          children: [
            Stack(
              children: [
                CircleAvatar(
                  radius: 24,
                  // backgroundImage: AssetImage(chat.image),
                ),
                // if (chat.isActive)
                //   Positioned(
                //     right: 0,
                //     bottom: 0,
                //     child: Container(
                //       height: 16,
                //       width: 16,
                //       decoration: BoxDecoration(
                //         color: kPrimaryColor,
                //         shape: BoxShape.circle,
                //         border: Border.all(
                //             color: Theme.of(context).scaffoldBackgroundColor,
                //             width: 3),
                //       ),
                //     ),
                //   )
              ],
            ),
            Expanded(
              child: Padding(
                padding:
                    const EdgeInsets.symmetric(horizontal: kDefaultPadding),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      chats.chatRoomName,
                      style:
                          TextStyle(fontSize: 16, fontWeight: FontWeight.w500),
                    ),
                    SizedBox(height: 8),
                    if (chats.edges!.chatMessage != null) ...[
                      Opacity(
                        opacity: 0.64,
                        child: Text(
                          chats
                              .edges!
                              .chatMessage![chats.edges!.chatMessage!.length - 1]
                              .messageText,
                          maxLines: 1,
                          overflow: TextOverflow.ellipsis,
                        ),
                      ),
                    ],
                  ],
                ),
              ),
            ),
            if (isChatsLock) ...{
              Icon(
                Icons.lock,
                color: Colors.pink,
                size: 20.0,
              ),
            },
            if (chats.edges!.chatMessage != null) ...{
              Opacity(
                opacity: 0.64,
                child: Text(CalculateTime()),
              ),
            },
            
          ],
        ),
      ),
    );
  }
}
