import 'dart:ui';

import 'package:flutter/material.dart';

import 'model.dart';
import 'shared/assets.dart';
import 'shared/ui/icon.dart';
import 'shared/ui/floating.dart';

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
          child: Text("Low stress day", style: TextStyle(fontWeight: FontWeight.bold)),
        ),
        Spacer(),
      ],
    );
  }

  Widget _weatherRow() {
    return Row(
      children: [
        scoreIcon(8000, radius: 14.0),
        Container(
          padding: EdgeInsets.fromLTRB(8.0, 0.0, 0.0, 0.0),
          child: Text("Weather BAD", style: TextStyle(fontWeight: FontWeight.bold)),
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
          child: Text("TRAFFIC SUX", style: TextStyle(fontWeight: FontWeight.bold)),
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
    return Scaffold(
      appBar: AppBar(
        leading: new IconButton(
          icon: new Icon(Icons.menu),
          color: Theme.of(context).primaryIconTheme.color,
          tooltip: 'Menu',
        ),
        title: Text(_day.dateName(), style: Theme.of(context).textTheme.title),
        backgroundColor: Theme.of(context).primaryColor,
      ),
      body: ListView.builder(
        padding: const EdgeInsets.all(0.0),
        itemBuilder: (context, i) {
          if (i == 0) {
            return _chipGrid();
          }
          i--;

          if (i==0){
            return Text("G");
          }

          return Text(i.toString());
        },
      ),
      floatingActionButton: floatingActionButton,
    );
  }
}

class TaskSchedule extends StatelessWidget{

  final ScheduledTask task;
  final ScheduledTask next;

  TaskSchedule(this.task, this.next);

  @override
  Widget build(BuildContext context) {
    return Text("GGGG");
  }
}