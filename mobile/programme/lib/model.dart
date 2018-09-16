import 'dart:math';

import 'package:flutter/material.dart';
import 'package:material_design_icons_flutter/material_design_icons_flutter.dart';

var rng = new Random();

class Day {
  List<ScheduledTask> tasks = new List();
  DateTime startTime = new DateTime(0, 0, 0, 8);

  Day() {
    var cur = startTime;
    var dur = new Duration(minutes: rng.nextInt(60) + 30);
    for (int i = rng.nextInt(5) + 3; i >= 0; i--) {

      switch (rng.nextInt(4)){
        case 0:
        tasks.add(new FreeTask("T", dur, rng.nextInt(11), cur,"travel"));
        break;
        case 1:
         tasks.add(new TravelTask("T", dur, rng.nextInt(11), cur,"travel", TravelMethod.car));
         break;
        default:
         tasks.add(new ScheduledTask("G", dur, rng.nextInt(11), cur, ""));
      }
      cur = cur.add(dur);
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

class Coords {
  double latitude;
  double longitude;

  Coords(this.latitude, this.longitude);
}

class Location {
  String name;
  Coords coords;

  Location(this.name, this.coords);
}

class Task {
  String name;
  String description = "";
  Duration duration;
  int stress;
  Location location;

  Task(this.name, this.duration, this.stress,this.description, {Location location});

  Task.empty() : super();
}

class ScheduledTask extends Task {
  DateTime start;
  DateTime end;

  ScheduledTask.empty() : super.empty();

  ScheduledTask(name, duration, stress, this.start, description,{Location location})
      : super(name, duration, stress, description, location: location) {
    this.end = start.add(duration);
  }
}

class TravelTask extends ScheduledTask {
  TravelTask.empty() : super.empty();
  TravelMethod travelMethod;

  Widget travelMethodIcon() {
    switch (travelMethod) {
      case TravelMethod.walking:
        return new Icon(MdiIcons.walk);
      case TravelMethod.biking:
        return new Icon(MdiIcons.bike);
      case TravelMethod.rideshare:
        return new Icon(MdiIcons.caravan);
      case TravelMethod.car:
        return new Icon(MdiIcons.car);
    }
    return Text("");
  }

  TravelTask(name, duration, stress, start,description, this.travelMethod,
      {Location location})
      : super(name, duration, stress, start, description, location: location);
}

class FreeTask extends ScheduledTask {
  FreeTask.empty() : super.empty();

  FreeTask(name, duration, stress, start,description, {Location location})
      : super(name, duration, stress, start, description,location: location);
}

class FloatingTask extends Task {
  DateTime deadline;

  FloatingTask.empty() : super.empty();

  FloatingTask(name, duration, stress, description, this.deadline, {Location location})
      : super(name, duration, stress,description, location: location);
}

enum TaskType {
  floating,
  scheduled,
  travel,
  free,
}

enum TravelMethod {
  walking,
  biking,
  rideshare,
  car,
}
