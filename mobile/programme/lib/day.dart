import 'dart:ui';

import 'package:collection/collection.dart';
import 'package:flutter/material.dart';

import 'model.dart';
import 'new_task.dart';
import 'shared/assets.dart';
import 'shared/ui/icon.dart';
import "shared/ui/google.dart";
import 'shared/ui/timeformat.dart';

class DayPage extends StatefulWidget {
  final Day _day;

  DayPage(this._day);

  @override
  DayPageState createState() => new DayPageState(_day);
}

class DayPageState extends State<DayPage> {
  final Day _day;

  DayPageState(this._day);

  Widget _scoreRow() {
    return Row(
      children: [
        scoreIcon(2000, radius: 14.0),
        Container(
          padding: EdgeInsets.fromLTRB(8.0, 0.0, 0.0, 0.0),
          child: Text("Low stress day",
              style: TextStyle(fontWeight: FontWeight.bold)),
        ),
        Spacer(),
      ],
    );
  }

  Widget _weatherRow() {
    return Row(
      children: [
        scoreIcon(5000, radius: 14.0),
        Container(
          padding: EdgeInsets.fromLTRB(8.0, 0.0, 0.0, 0.0),
          child: Text("Weather not great",
              style: TextStyle(fontWeight: FontWeight.bold)),
        ),
        Spacer(),
      ],
    );
  }

  Widget _trafficRow() {
    return Row(
      children: [
        scoreIcon(10000, radius: 14.0),
        Container(
          padding: EdgeInsets.fromLTRB(8.0, 0.0, 0.0, 0.0),
          child: Text("TRAFFIC SUX",
              style: TextStyle(fontWeight: FontWeight.bold)),
        ),
        Spacer(),
      ],
    );
  }

  Widget _chipContainer(Widget w) {
    return Expanded(
      flex: 2,
      child: w,
    );
  }

  Widget _chipGrid() {
    return Container(
      constraints: BoxConstraints.expand(
        height: 56.0 * 3,
      ),
      decoration: BoxDecoration(
        image: DecorationImage(
          image: rainPNG,
          fit: BoxFit.cover,
        ),
      ),
      child: Row(
        children: [
          Spacer(),
          Expanded(
            flex: 16,
            child: Column(
              children: <Widget>[
                Spacer(flex: 2),
                _chipContainer(
                  blurChip(
                    _scoreRow(),
                    radius: 14.0,
                  ),
                ),
                Spacer(),
                _chipContainer(
                  blurChip(
                    _weatherRow(),
                    radius: 14.0,
                  ),
                ),
                Spacer(),
                _chipContainer(
                  blurChip(
                    _trafficRow(),
                    radius: 14.0,
                  ),
                ),
                Spacer(flex: 2),
              ],
            ),
          ),
          Spacer(),
        ],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    var rows = IterableZip(<List<ScheduledTask>>[
          <ScheduledTask>[null, null] + _day.tasks,
          <ScheduledTask>[null] + _day.tasks + <ScheduledTask>[null],
          _day.tasks + <ScheduledTask>[null, null],
        ])
            .map<Widget>(
                (triple) => TaskSchedule(triple[0], triple[1], triple[2]))
            .toList() +
        List.generate(5, (_) => ListTile());

    return Scaffold(
      appBar: AppBar(
        // leading: new IconButton(
        //   icon: new Icon(Icons.menu),
        //   color: Theme.of(context).primaryIconTheme.color,
        //   tooltip: 'Menu',
        //   onPressed: () => null,
        // ),
        title: Text(_day.dateName(), style: Theme.of(context).textTheme.title),
        backgroundColor: Theme.of(context).primaryColor,
      ),
      body: ListView.builder(
        padding: const EdgeInsets.all(0.0),
        itemCount: rows.length + 1,
        itemBuilder: (context, i) {
          if (i == 0) {
            return _chipGrid();
          }
          i--;

          return rows[i];
        },
      ),
      floatingActionButton: newTaskButton(context, _day.startTime),
    );
  }
}

class TaskSchedule extends StatelessWidget {
  final ScheduledTask previous;
  final ScheduledTask current;
  final ScheduledTask next;

  final halfnoblock =
      rectangle(width: 7.0, height: 3.5, color: Colors.transparent);
  final noblock = rectangle(width: 7.0, height: 7.0, color: Colors.transparent);

  final labelStyle = TextStyle(
    fontSize: 16.0,
  );
  final blockStyle = TextStyle(
    fontSize: 24.0,
    color: Colors.white,
  );

  TaskSchedule(this.previous, this.current, this.next);

  Widget _startTask(context) {
    return Row(
      children: <Widget>[
        Expanded(
          flex: 2,
          child: Row(
            children: [
              Spacer(),
              Padding(
                padding: EdgeInsets.fromLTRB(0.0, 0.0, 16.0, 0.0),
                child: Text(
                  "\n" + timeFormat.format(next.start),
                  style: labelStyle,
                  textAlign: TextAlign.right,
                ),
              ),
            ],
          ),
        ),
        Container(
          constraints: BoxConstraints.expand(
            height: 56.0,
            width: 14.0,
          ),
          child: Center(
            child: Column(children: [
              rectangle(width: 24.5, height: 24.5, color: Colors.transparent),
              rectangle(
                  width: 14.0,
                  height: 14.0,
                  color: GoogleColors.blue,
                  radius: BorderRadius.circular(7.0)),
              rectangle(width: 7.0, height: 7.0, color: Colors.transparent),
              rectangle(
                  width: 7.0, height: 7.0, color: stressColor(next.stress)),
            ]),
          ),
        ),
        Spacer(flex: 4),
        rectangle(
          width: 56.0,
          height: 56.0,
        ),
      ],
    );
  }

  Widget _endTask(context) {
    return Row(
      children: <Widget>[
        Expanded(
          flex: 2,
          child: Row(
            children: [
              Spacer(),
              Padding(
                padding: EdgeInsets.fromLTRB(0.0, 0.0, 16.0, 0.0),
                child: Text(
                  timeFormat.format(previous.end) + "\n",
                  style: labelStyle,
                  textAlign: TextAlign.right,
                ),
              ),
            ],
          ),
        ),
        Container(
          constraints: BoxConstraints.expand(
            height: 56.0,
            width: 14.0,
          ),
          child: Center(
            child: Column(children: [
              halfnoblock,
              rectangle(
                  width: 7.0, height: 7.0, color: stressColor(previous.stress)),
              rectangle(width: 7.0, height: 7.0, color: Colors.transparent),
              rectangle(
                  width: 14.0,
                  height: 14.0,
                  color: GoogleColors.blue,
                  radius: BorderRadius.circular(7.0)),
              rectangle(width: 24.5, height: 24.5, color: Colors.transparent),
            ]),
          ),
        ),
        Spacer(flex: 4),
        rectangle(
          width: 56.0,
          height: 56.0,
        ),
      ],
    );
  }

  Widget _task(context) {
    return Row(
      children: <Widget>[
        Expanded(
          flex: 2,
          child: Row(
            children: [
              Spacer(),
              Padding(
                padding: EdgeInsets.fromLTRB(0.0, 0.0, 16.0, 0.0),
                child: Text(
                  "${timeFormat.format(current.start)}\n${timeFormat.format(current.end)}",
                  style: labelStyle,
                  textAlign: TextAlign.right,
                ),
              ),
            ],
          ),
        ),
        Container(
          constraints: BoxConstraints.expand(
            height: 56.0,
            width: 14.0,
          ),
          child: Center(
            child: Column(children: [
              halfnoblock,
              rectangle(
                  width: 14.0,
                  height: 49.0,
                  color: stressColor(current.stress),
                  radius: BorderRadius.all(Radius.circular(7.0))),
              halfnoblock,
            ]),
          ),
        ),
        Spacer(),
        Expanded(
          flex: 3,
          child: Padding(
            padding: EdgeInsets.fromLTRB(12.0, 0.0, 0.0, 0.0),
            child: Text(
              current.description,
              style: TextStyle(
                color: Theme.of(context).textTheme.body2.color,
                fontSize: 18.0,
              ),
            ),
          ),
        ),
        rectangle(
            width: 56.0,
            height: 56.0,
            color: stressColor(current.stress),
            child: Center(
                child: Text(current.stress.toString(), style: blockStyle))),
      ],
    );
  }

  Widget _travelTask(context, travel) {
    var block =
        rectangle(width: 7.0, height: 7.0, color: stressColor(current.stress));

    var dur = current.end.difference(current.start);
    return Row(
      children: <Widget>[
        Expanded(
          flex: 2,
          child: Row(
            children: [
              Spacer(),
              Padding(
                padding: EdgeInsets.fromLTRB(0.0, 0.0, 16.0, 0.0),
                child: Text(
                  "(${dur.inHours}:${(dur.inMinutes % 60).toString().padLeft(2, '0')})",
                  style: labelStyle,
                  textAlign: TextAlign.right,
                ),
              ),
            ],
          ),
        ),
        Container(
          constraints: BoxConstraints.expand(
            height: 56.0,
            width: 14.0,
          ),
          child: Center(
            child: Column(children: [
              halfnoblock,
              block,
              noblock,
              block,
              noblock,
              block,
              noblock,
              block,
              halfnoblock,
            ]),
          ),
        ),
        Spacer(),
        Expanded(
          flex: 3,
          child: Padding(
            padding: EdgeInsets.fromLTRB(12.0, 0.0, 0.0, 0.0),
            child: Row(
              children: [
                travel.travelMethodIcon(),
                Text(
                  "R",
                  style: labelStyle,
                ),
                Spacer(flex:7),
              ],
            ),
          ),
        ),
        rectangle(
            width: 56.0,
            height: 56.0,
            color: stressColor(current.stress),
            child: Center(
                child: Text(current.stress.toString(), style: blockStyle))),
      ],
    );
  }

  Widget _freeTask(context) {
    var block =
        rectangle(width: 7.0, height: 7.0, color: stressColor(current.stress));

    var dur = current.end.difference(current.start);
    return Row(
      children: <Widget>[
        Expanded(
          flex: 2,
          child: Row(
            children: [
              Spacer(),
              Padding(
                padding: EdgeInsets.fromLTRB(0.0, 0.0, 16.0, 0.0),
                child: Text(
                  "(${dur.inHours}:${(dur.inMinutes % 60).toString().padLeft(2, '0')})",
                  style: labelStyle,
                  textAlign: TextAlign.right,
                ),
              ),
            ],
          ),
        ),
        Container(
          constraints: BoxConstraints.expand(
            height: 56.0,
            width: 14.0,
          ),
          child: Center(
            child: Column(children: [
              halfnoblock,
              rectangle(
                  width: 14.0,
                  height: 49.0,
                  color: GoogleColors.blue,
                  radius: BorderRadius.all(Radius.circular(7.0))),
              halfnoblock,
            ]),
          ),
        ),
        Spacer(),
        Expanded(
          flex: 3,
          child: Padding(
            padding: EdgeInsets.fromLTRB(12.0, 0.0, 0.0, 0.0),
            child: Text(
              "",
              style: labelStyle,
            ),
          ),
        ),
        rectangle(
            width: 56.0,
            height: 56.0,
            color: stressColor(current.stress),
            child: Center(
                child: Text(current.stress.toString(), style: blockStyle))),
      ],
    );
  }

  @override
  Widget build(BuildContext context) {
    if (previous == null && current == null) {
      return _startTask(context);
    }

    if (next == null && current == null) {
      return _endTask(context);
    }

    if (current is TravelTask) {
      return _travelTask(context, current as TravelTask);
    }

    if (current is FreeTask) {
      return _freeTask(context);
    }

    return _task(context);
  }
}
