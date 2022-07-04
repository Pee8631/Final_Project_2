import 'package:flutter/material.dart';

// ignore: must_be_immutable
class HomeButton extends StatelessWidget {
  Icon icons;
  String title;
  double fontsize;
  Color color;
  Widget nextPage;

  HomeButton(this.icons, this.title, this.fontsize, this.color, this.nextPage);

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: double.infinity,
      child: ElevatedButton.icon(
        icon: Icon(icons.icon),
        label: Text(
          title,
          style: TextStyle(fontSize: fontsize),
        ),
        onPressed: () {
          Navigator.pushReplacement(context,
              MaterialPageRoute(builder: (context) {
            return nextPage;
          }));
        },
      ),
    );
  }
}
