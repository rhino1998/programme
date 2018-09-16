import 'package:card_settings/card_settings.dart';
import 'package:flutter/material.dart';

import 'model.dart';

class NewTaskDialog extends StatefulWidget {
  final DateTime day;

  NewTaskDialog(this.day);

  @override
  NewTaskDialogState createState() => new NewTaskDialogState(day);
}

class NewTaskDialogState extends State<NewTaskDialog> {
  final DateTime day;

  NewTaskDialogState(this.day);

  Task task;
  bool _scheduled = false;
  final _formKey = GlobalKey<FormState>();

  String name;
  String description = "";
  int stress;
  Duration duration;
  DateTime _deadlineStart;
  Location location;

  ScheduledTask makeScheduleldTask() {
    return ScheduledTask(name, duration, stress, _deadlineStart, description,
        location: location);
  }

  FloatingTask makeFloatingTask() {
    return FloatingTask(name, duration, stress, description, _deadlineStart,
        location: location);
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.all(32.0),
      child: Form(
        key: _formKey,
        child: CardSettings(
          padding: 0.0,
          children: <Widget>[
            CardSettingsHeader(
              label: "New Task",
            ),
            CardSettingsText(
              label: 'Name',
              validator: (value) {
                if (value == "") return 'Must enter a name';
              },
              onSaved: (val) => name = val,
            ),
            CardSettingsParagraph(
              label: 'Description',
              numberOfLines: 2,
              onSaved: (val) => description = val,
            ),
            CardSettingsNumberPicker(
              label: "Stress",
              min: 0,
              max: 11,
              onSaved: (val) => stress = val,
            ),
            CardSettingsSwitch(
              label: "Scheduled?",
              onChanged: (value) => setState(() => _scheduled = value),
            ),
            CardSettingsTimePicker(
              label: "Duration",
              initialValue: new TimeOfDay(hour: 0, minute: 0),
              onSaved: (val) =>
                  (duration = Duration(hours: val.hour, minutes: val.minute)),
            ),
            CardSettingsDatePicker(
              label: _scheduled ? "Start" : "Deadline",
              initialValue: this.day.add(Duration(hours: 2)),
              validator: (val) {
                if (val.isBefore(DateTime.now())) {
                  return "Date must be in the future";
                }
              },
              onSaved: (val) => _deadlineStart = val,
            ),
            CardSettingsButton(
                label: "Add",
                onPressed: () {
                  if (_formKey.currentState.validate()) {
                    if (_scheduled) {
                      makeScheduleldTask();
                    } else {
                      makeFloatingTask();
                    }
                  }
                }),
          ],
        ),
      ),
    );
  }
}

Widget newTaskButton(context, DateTime today) {
  return FloatingActionButton(
    child: Icon(Icons.add),
    onPressed: () => showDialog(
          context: context,
          builder: (BuildContext context) {
            return NewTaskDialog(today);
          },
        ),
  );
}
