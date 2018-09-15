import 'dart:ui';

import 'package:flutter/material.dart';
import 'package:material_design_icons_flutter/material_design_icons_flutter.dart';

CircleAvatar scoreIcon(int score, {double radius}) {
  if (score < 3000) {
    return CircleAvatar(
      radius: radius,
      backgroundColor: Colors.green,
      child: new Icon(
        MdiIcons.emoticonHappy,
        color: Colors.white,
      ),
    );
  } else if (score < 6000) {
    return CircleAvatar(
      radius: radius,
      backgroundColor: Colors.yellow,
      child: new Icon(
        MdiIcons.emoticonNeutral,
        color: Colors.white,
      ),
    );
  } else if (score < 9000) {
    return CircleAvatar(
      radius: radius,
      backgroundColor: Colors.orange,
      child: new Icon(
        MdiIcons.emoticonSad,
        color: Colors.white,
      ),
    );
  }
  return CircleAvatar(
    radius: radius,
    backgroundColor: Colors.red,
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
