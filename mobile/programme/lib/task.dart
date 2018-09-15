import 'package:flutter/material.dart';
import 'package:english_words/english_words.dart';
import 'package:intl/intl.dart';
import 'package:material_design_icons_flutter/material_design_icons_flutter.dart';
import 'dart:math';

var rng = new Random();

var timeFormat = DateFormat('hh:mm a');

class Day {
  List<ScheduledTask> tasks = new List();
  DateTime startTime = new DateTime(0, 0, 0, 8);

  Day() {
    var cur = startTime;
    for (int i = rng.nextInt(5) + 3; i >= 0; i--) {
      var next = cur.add(new Duration(minutes: rng.nextInt(60) + 30));
      tasks.add(new ScheduledTask(cur, next));
      cur = next;
    }
  }

  String dateName() {
    return "Today";
  }

  Duration duration() {
    return tasks.fold(new Duration(seconds: 0),
        (sum, task) => task.end.difference(task.start) + sum);
  }

  int stress() {
    return tasks.fold(0, (sum, task) => task.stress + sum);
  }

  int score() {
    var dur = duration();
    return stress() * dur.inMinutes + dur.inHours * dur.inHours;
  }
}

class ScheduledTask {
  String name = "GENERIC";
  int stress = 0;
  DateTime start;
  DateTime end;

  ScheduledTask(this.start, this.end) {
    this.name = WordPair.random().asPascalCase;
    this.stress = rng.nextInt(10);
  }
}

class TaskList extends StatefulWidget {
  @override
  TaskListState createState() => new TaskListState();
}

class TaskListState extends State<TaskList> {
  List<Day> days = new List();

  TaskListState() {
    for (int i = rng.nextInt(5) + 3; i >= 0; i--) {
      days.add(Day());
    }
  }

  @override
  Widget build(BuildContext context) {
    var rows = days.fold<List<Widget>>(
        [],
        (l, day) =>
            l +
            ((day.tasks.length == 0) ? [] : [new DayRow(day)]) +
            day.tasks.map((task) => new TaskRow(task)).toList()) + List.generate(5, (_)=>ListTile());
    return ListView.builder(
      padding: const EdgeInsets.all(0.0),
      itemCount: rows.length,
      itemBuilder: (context, i) {
        return rows[i];
      },
    );
  }
}

class DayRow extends StatelessWidget {
  final Day day;

  DayRow(this.day);

  @override
  Widget build(BuildContext context) {
    return Row(
      children: <Widget>[
        Container(
          constraints: BoxConstraints.expand(
            height: 56.0,
            width: 10.0,
          ),
          decoration: BoxDecoration(
            color: Colors.purple,
          ),
        ),
        Padding(
          padding: EdgeInsets.fromLTRB(12.0, 0.0, 0.0, 0.0),
          child: Text(
            day.dateName(),
            style: TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 18.0,
            ),
          ),
        ),
        Spacer(),
        Padding(
          padding: EdgeInsets.fromLTRB(0.0, 0.0, 16.0, 0.0),
          child: Row(
            children: <Widget>[
              scoreIcon(),
            ],
          ),
        ),
      ],
    );
  }

  CircleAvatar scoreIcon() {
    var score = day.score();
    if (score < 4000) {
      return CircleAvatar(
        backgroundColor: Colors.green,
        child: new Icon(
          MdiIcons.emoticonHappy,
          color: Colors.white,
        ),
      );
    } else if (score < 8000) {
      return CircleAvatar(
        backgroundColor: Colors.yellow,
        child: new Icon(
          MdiIcons.emoticonNeutral,
          color: Colors.white,
        ),
      );
    }
    return CircleAvatar(
      backgroundColor: Colors.red,
      child: new Icon(
        MdiIcons.emoticonDead,
        color: Colors.white,
      ),
    );
  }
}

class TaskRow extends StatefulWidget {
  final ScheduledTask _task;

  TaskRow(this._task);

  @override
  TaskRowState createState() => new TaskRowState(this._task);
}

class TaskRowState extends State<TaskRow> {
  ScheduledTask _task;

  TaskRowState(this._task);

  @override
  Widget build(BuildContext context) {
    return Row(
      children: <Widget>[
        Container(
          constraints: BoxConstraints.expand(
            height: 56.0,
            width: 10.0,
          ),
          decoration: BoxDecoration(
            color: this.taskListStressBarColor(),
          ),
        ),
        Padding(
          padding: EdgeInsets.fromLTRB(12.0, 0.0, 0.0, 0.0),
          child: Text(
            this._task.name,
            style: TextStyle(
              fontSize: 18.0,
            ),
          ),
        ),
        Spacer(),
        Padding(
          padding: EdgeInsets.fromLTRB(0.0, 0.0, 48.0, 0.0),
          child: Text(
            timeFormat.format(this._task.start),
            style: TextStyle(
              color: Theme.of(context).textTheme.body2.color,
              fontSize: 18.0,
            ),
          ),
        ),
      ],
    );
  }

  Color taskListStressBarColor() {
    if (this._task.stress <= 3) {
      return Colors.green;
    } else if (this._task.stress <= 6) {
      return Colors.orange;
    }
    return Colors.red;
  }
}
