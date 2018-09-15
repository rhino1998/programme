import 'package:flutter/material.dart';
import 'package:intl/intl.dart';

import 'day.dart';
import 'model.dart';

import 'shared/ui/icon.dart';
import 'shared/ui/google.dart';
import 'shared/ui/timeformat.dart';

class TaskList extends StatefulWidget {
  @override
  TaskListState createState() => new TaskListState();
}

class TaskListState extends State<TaskList> {
  List<Day> days = new List();

  TaskListState() {
    for (int i = rng.nextInt(10) + 6; i >= 0; i--) {
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
    return GestureDetector(
      onTap: () {
              Navigator.push(
                context,
                MaterialPageRoute(builder: (context) => DayPage(day)),
              );
            },
      child: Row(
      children: <Widget>[
        Container(
          constraints: BoxConstraints.expand(
            height: 56.0,
            width: 10.0,
          ),
          decoration: BoxDecoration(
            color: Theme.of(context).accentColor,
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
              scoreIcon(day.score()),
            ],
          ),
        ),
      ],
      ),
    );
  }
}

class TaskRow extends StatefulWidget {
  final ScheduledTask _task;

  TaskRow(this._task);

  @override
  TaskRowState createState() => new TaskRowState(_task);
}

class TaskRowState extends State<TaskRow> {
  ScheduledTask _task;

  TaskRowState(this._task);

  @override
  Widget build(BuildContext context) {
    return Row(
      children: <Widget>[
        rectangle(height:56.0, width:10.0, color: stressColor(_task.stress)),
        Padding(
          padding: EdgeInsets.fromLTRB(12.0, 0.0, 0.0, 0.0),
          child: Text(
            _task.name,
            style: TextStyle(
              fontSize: 18.0,
            ),
          ),
        ),
        Spacer(),
        Padding(
          padding: EdgeInsets.fromLTRB(0.0, 0.0, 48.0, 0.0),
          child: Text(
            timeFormat.format(_task.start),
            style: TextStyle(
              color: Theme.of(context).textTheme.body2.color,
              fontSize: 18.0,
            ),
          ),
        ),
      ],
    );
  }
}