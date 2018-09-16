import 'dart:ui';

import 'package:flutter/material.dart';
import 'package:material_design_icons_flutter/material_design_icons_flutter.dart';

import 'google.dart';

CircleAvatar scoreIcon(int score, {double radius}) {
  if (score < 4000) {
    return CircleAvatar(
      radius: radius,
      backgroundColor: GoogleColors.green,
      child: new Icon(
        MdiIcons.emoticonHappy,
        color: Colors.white,
      ),
    );
  } else if (score < 8000) {
    return CircleAvatar(
      radius: radius,
      backgroundColor: GoogleColors.yellow,
      child: new Icon(
        MdiIcons.emoticonNeutral,
        color: Colors.white,
      ),
    );
  }
  return CircleAvatar(
    radius: radius,
    backgroundColor: GoogleColors.red,
    child: new Icon(
      MdiIcons.emoticonDead,
      color: Colors.white,
    ),
  );
}

Widget blurChip(Widget widget, {double radius}) {
  return ClipRRect(
    borderRadius: BorderRadius.all(Radius.circular(radius)),
    child: BackdropFilter(
      filter: ImageFilter.blur(sigmaX: 1.0, sigmaY: 1.0),
      child: Container(
        child: widget,
        color: Color(0x78FFFFFF),
      ),
    ),
  );
}

Widget rectangle(
    {width: 0.0, height: 0.0, Color color, BorderRadius radius, Widget child}) {
  if (radius != null) {
    return ClipRRect(
      borderRadius: radius,
      child: Container(
        constraints: BoxConstraints.expand(
          height: height,
          width: width,
        ),
        child: child,
        decoration: BoxDecoration(
          color: color,
        ),
      ),
    );
  }
  return Container(
    constraints: BoxConstraints.expand(
      height: height,
      width: width,
    ),
    child: child,
    decoration: BoxDecoration(
      color: color,
    ),
  );
}

Color stressColor(int stress) {
  if (stress <= 4) {
    return GoogleColors.green;
  } else if (stress <= 6) {
    return GoogleColors.yellow;
  }
  return GoogleColors.red;
}

Widget weatherIcon(int id, {double radius}) {
  switch (id ~/100) {
    case (2):
      return CircleAvatar(
        radius: radius,
        backgroundColor: GoogleColors.red,
        child: Icon(
          MdiIcons.weatherLightningRainy,
          color: Colors.white,
        ),
      );
    case (3):
    case (5):
      return CircleAvatar(
        radius: radius,
        backgroundColor: GoogleColors.yellow,
        child: Icon(
          MdiIcons.weatherRainy,
          color: Colors.white,
        ),
      );
    case (6):
      return CircleAvatar(
        radius: radius,
        backgroundColor: GoogleColors.red,
        child: Icon(
          MdiIcons.weatherSnowy,
          color: Colors.white,
        ),
      );
    case (8):
      return CircleAvatar(
        radius: radius,
        backgroundColor: GoogleColors.green,
        child: Icon(
          MdiIcons.weatherSunny,
          color: Colors.white,
        ),
      );
  }
  return CircleAvatar(
    radius: radius,
    backgroundColor: GoogleColors.blue,
    child: Icon(
      MdiIcons.calendarQuestion,
      color: Colors.white,
    ),
  );
}
