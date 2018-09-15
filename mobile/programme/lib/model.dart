

import 'dart:math';

import 'package:english_words/english_words.dart';

var rng = new Random();

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
  bool travel = rng.nextBool();

  ScheduledTask(this.start, this.end) {
    this.name = WordPair.random().asPascalCase;
    this.stress = rng.nextInt(10);
  }
}