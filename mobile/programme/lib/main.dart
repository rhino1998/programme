import 'package:flutter/material.dart';

import 'task.dart';

import "shared/ui/floating.dart";

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'programme',
      home: TaskListPage(),
      theme: ThemeData(
        brightness: Brightness.light,
        primaryColor: Colors.white,
        accentColor: Colors.deepPurple,
        primaryIconTheme: IconThemeData(
          color: Colors.black,
        ),
        textTheme: TextTheme(
          body2: TextStyle(color: Colors.grey),
          title: TextStyle(
            color: Colors.black,
          ),
        ),
      ),
    );
  }
}

class TaskListPage extends StatefulWidget {
  @override
  TaskListPageState createState() => new TaskListPageState();
}

class TaskListPageState extends State<TaskListPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: new IconButton(
          icon: new Icon(Icons.menu),
          color: Theme.of(context).primaryIconTheme.color,
          tooltip: 'Menu',
          onPressed: () => null,
        ),
        title: Text("programme", style: Theme.of(context).textTheme.title),
        backgroundColor: Theme.of(context).primaryColor,
      ),
      body: new TaskList(),
      floatingActionButton: floatingActionButton,
    );
  }
}
