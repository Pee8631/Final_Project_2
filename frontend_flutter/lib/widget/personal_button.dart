import 'package:flutter/material.dart';

class PersonalButtonWidget extends StatelessWidget {
  final String text;
  final VoidCallback onClicked;

  const PersonalButtonWidget(
      {Key? key, required this.text, required this.onClicked})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      child: ElevatedButton(
          style: ElevatedButton.styleFrom(
            shape: RoundedRectangleBorder(),
            onPrimary: Colors.transparent,
            primary: Colors.transparent,
            elevation: 0,
            
          ),
          child: Text(
            text,
            style: TextStyle(fontSize: 12, color: Colors.black),
          ),
          onPressed: this.onClicked),
    );
  }
}
