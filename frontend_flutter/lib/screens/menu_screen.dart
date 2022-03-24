import 'package:flutter/material.dart';
import 'package:frontend_flutter/util/user_references.dart';
import 'package:frontend_flutter/widget/menu_button.dart';

import 'home_screen.dart';

class MenuScreen extends StatelessWidget {
  MenuScreen({Key? key}) : super(key: key);
  final profile = UserReferences.myProfile;
  final List<String> textList = [
    "He'd have you",
    "Heed not the rabble",
    "Sound of screams",
    "Who scream",
    "Revolution is",
    "Revolution, they..."
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
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
                    color: Colors.blue,
                  ),
                ),
              ),
            ),
            SliverGrid(
              gridDelegate: const SliverGridDelegateWithMaxCrossAxisExtent(
                maxCrossAxisExtent: 200.0,
                mainAxisSpacing: 20.0,
                mainAxisExtent: 150.0,
                crossAxisSpacing: 15.0,
                childAspectRatio: 2.0,
              ),
              delegate: SliverChildBuilderDelegate(
                (BuildContext context, int index) {
                  var b = textList[index];
                  return Container(
                    alignment: Alignment.center,
                    child: Column(
                      children: [
                        MenuButtonWidget(
                          imagePath:
                              'https://png.pngtree.com/png-clipart/20190705/original/pngtree-vector-cancel-icon-png-image_4254546.jpg',
                          onClicked: () async {},
                        ),
                        Text(
                          b,
                          style:
                              TextStyle(fontSize: 20, color: Colors.grey[600]),
                        ),
                      ],
                    ),
                  );
                },
                childCount: textList.length,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
