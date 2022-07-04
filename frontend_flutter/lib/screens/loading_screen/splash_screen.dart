import 'dart:async';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:frontend_flutter/screens/signupandsignin/signin_userscreen.dart';
import 'package:video_player/video_player.dart';

class SplashScreen extends StatefulWidget {
  const SplashScreen({Key? key}) : super(key: key);

  @override
  State<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen> {
  late VideoPlayerController video;
  @override
  void initState() {
    super.initState();
    video = VideoPlayerController.asset('assets/videos/videoplayback.mp4')
      ..initialize().then((_) {
        setState(() {});
      })
      ..setVolume(0.0);

    _playVideo();
  }

  void _playVideo() async {
    Navigator.of(context).popUntil((route) => route.isFirst);
    video.setPlaybackSpeed(3);
    video.play();
    Timer(const Duration(milliseconds: 6000), () {
              Navigator.pushReplacement(
                  context,
                  MaterialPageRoute(
                      builder: (BuildContext context) => SignInScreen()));
            });
  }

  @override
  void dispose() {
    video.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: Center(
        child: Container(
          height: 200,
          width: 200,
          child: video.value.isInitialized
              ? AspectRatio(
                  aspectRatio: video.value.aspectRatio,
                  child: VideoPlayer(video),
                )
              : SizedBox(child: Text('')),
        ),
      ),
    );
  }
}
