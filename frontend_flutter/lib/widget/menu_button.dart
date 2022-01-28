import 'package:flutter/material.dart';

class MenuButtonWidget extends StatelessWidget {
  final String imagePath;
  final VoidCallback onClicked;

  const MenuButtonWidget(
      {Key? key, required this.imagePath, required this.onClicked})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ClipRRect(
      borderRadius: BorderRadius.circular(8.0),
      child: Material(
        color: Colors.transparent,
        child: Ink.image(
          image: NetworkImage(imagePath),
          fit: BoxFit.cover,
          width: 100,
          height: 100,
          child: InkWell(onTap: onClicked, highlightColor: Colors.grey,),
        ),
      ),
    ); 
  }



}
