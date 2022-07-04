import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:frontend_flutter/provider/event_provider.dart';
import 'package:frontend_flutter/screens/loading_screen/splash_screen.dart';
import 'package:intl/intl.dart';
import 'package:provider/provider.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await SystemChrome.setPreferredOrientations([
    DeviceOrientation.portraitUp,
    DeviceOrientation.portraitDown,
  ]);

  Intl.defaultLocale = "th";
  //initializeDateFormatting();
  
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        ChangeNotifierProvider( create: (context) => EventProvider()),
      ],
      child: MaterialApp(
          title: 'Flutter Demo',
          debugShowCheckedModeBanner: false,
          theme: ThemeData(
            // This is the theme of your application.
            //
            // Try running your application with "flutter run". You'll see the
            // application has a blue toolbar. Then, without quitting the app, try
            // changing the primarySwatch below to Colors.green and then invoke
            // "hot reload" (press "r" in the console where you ran "flutter run",
            // or simply save your changes to "hot reload" in a Flutter IDE).
            // Notice that the counter didn't reset back to zero; the application
            // is not restarted.
            primaryColor: Colors.blue.shade300,
          ),
          themeMode: ThemeMode.light,
          darkTheme: ThemeData.light().copyWith(
            scaffoldBackgroundColor: Colors.transparent,
            accentColor: Colors.black,
            primaryColor: Colors.blue,
          ),
          home: SplashScreen(),
        ),
    );
  }
}
