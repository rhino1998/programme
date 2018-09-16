import 'package:card_settings/card_settings.dart';
import 'package:flutter/material.dart';

import 'shared/ui/icon.dart';
import 'model.dart';
import "shared/ui/datetime.dart";

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

  Duration duration;
  DateTime _deadline_start;


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
              ),
              CardSettingsParagraph(
                label: 'Description',
                numberOfLines: 2,
                contentOnNewLine: false,
              ),
              CardSettingsNumberPicker(label: "Stress", min: 0, max: 10),
              CardSettingsSwitch(
                label: "Scheduled?",
                onChanged: (value) => setState(() => _scheduled = value),
              ),
              CardSettingsTimePicker(
                label: "Duration",
                initialValue: new TimeOfDay(hour: 0, minute: 0),
                onSaved:(val) => (duration=Duration(hours: val.hour, minutes:val.minute)),
              ),
              CardSettingsDatePicker(
                  label: _scheduled?"Start":"Deadline",
                  initialValue: this.day,
                  onSaved: (val) => _deadline_start=val,
              ),
              CardSettingsButton(label: "Add", onPressed: () {}),
            ],
          ),
        ));
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
