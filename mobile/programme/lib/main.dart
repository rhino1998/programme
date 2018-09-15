import 'package:flutter/material.dart';

import 'task.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'programme',
      home: RandomWords(),
      theme: ThemeData(
        brightness: Brightness.light,
        primaryColor: Colors.white,
        accentColor: Colors.cyan[600],
        primaryIconTheme: IconThemeData(
          color:Colors.black,
        ),
        textTheme: TextTheme(
          body2: TextStyle(
            color: Colors.grey
          ),
          title: TextStyle(
            color: Colors.black,
          ),
        ),
      ),
    );
  }
}

class RandomWords extends StatefulWidget {
  @override
  RandomWordsState createState() => new RandomWordsState();
}

class RandomWordsState extends State<RandomWords> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: new IconButton(
            icon: new Icon(Icons.menu),
            color: Theme.of(context).primaryIconTheme.color,
            tooltip: 'Menu',
            onPressed: () {
              Navigator.push(
                context,
                MaterialPageRoute(builder: (context) => DayPage()),
              );
            }),
        title: Text("programme", style: Theme.of(context).textTheme.title),
        backgroundColor: Theme.of(context).primaryColor,
      ),
      body: new TaskList(),
    );
  }
}

class DayPage extends StatefulWidget {
  @override
  DayPageState createState() => new DayPageState();
}

class DayPageState extends State<DayPage> {
  final String _day = "Today";

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: new IconButton(
          icon: new Icon(Icons.menu),
          color: Theme.of(context).primaryIconTheme.color,
          tooltip: 'Air it',
        ),
        title: Text(_day, style: Theme.of(context).textTheme.title),
        backgroundColor: Theme.of(context).primaryColor,
      ),
    );
  }
}

