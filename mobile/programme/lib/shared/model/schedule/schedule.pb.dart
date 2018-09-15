///
//  Generated code. Do not modify.
//  source: shared/model/schedule/schedule.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

// ignore: UNUSED_SHOWN_NAME
import 'dart:core' show int, bool, double, String, List, override;

import 'package:fixnum/fixnum.dart';
import 'package:protobuf/protobuf.dart' as $pb;

import '../routes/routes.pb.dart' as $0;

class Day extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Day', package: const $pb.PackageName('shared.model.schedule'))
    ..aInt64(1, 'datetime')
    ..pp<Task>(2, 'tasks', $pb.PbFieldType.PM, Task.$checkItem, Task.create)
    ..aOB(12, 'weatherNull')
    ..aOB(14, 'trafficNull')
    ..hasRequiredFields = false
  ;

  Day() : super();
  Day.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Day.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Day clone() => new Day()..mergeFromMessage(this);
  Day copyWith(void Function(Day) updates) => super.copyWith((message) => updates(message as Day));
  $pb.BuilderInfo get info_ => _i;
  static Day create() => new Day();
  static $pb.PbList<Day> createRepeated() => new $pb.PbList<Day>();
  static Day getDefault() => _defaultInstance ??= create()..freeze();
  static Day _defaultInstance;
  static void $checkItem(Day v) {
    if (v is! Day) $pb.checkItemFailed(v, _i.messageName);
  }

  Int64 get datetime => $_getI64(0);
  set datetime(Int64 v) { $_setInt64(0, v); }
  bool hasDatetime() => $_has(0);
  void clearDatetime() => clearField(1);

  List<Task> get tasks => $_getList(1);

  bool get weatherNull => $_get(2, false);
  set weatherNull(bool v) { $_setBool(2, v); }
  bool hasWeatherNull() => $_has(2);
  void clearWeatherNull() => clearField(12);

  bool get trafficNull => $_get(3, false);
  set trafficNull(bool v) { $_setBool(3, v); }
  bool hasTrafficNull() => $_has(3);
  void clearTrafficNull() => clearField(14);
}

class Task extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Task', package: const $pb.PackageName('shared.model.schedule'))
    ..aOS(1, 'name')
    ..aInt64(2, 'duration')
    ..aInt64(3, 'stress')
    ..aOB(14, 'locationNull')
    ..a<$0.Location>(15, 'locationValue', $pb.PbFieldType.OM, $0.Location.getDefault, $0.Location.create)
    ..hasRequiredFields = false
  ;

  Task() : super();
  Task.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Task.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Task clone() => new Task()..mergeFromMessage(this);
  Task copyWith(void Function(Task) updates) => super.copyWith((message) => updates(message as Task));
  $pb.BuilderInfo get info_ => _i;
  static Task create() => new Task();
  static $pb.PbList<Task> createRepeated() => new $pb.PbList<Task>();
  static Task getDefault() => _defaultInstance ??= create()..freeze();
  static Task _defaultInstance;
  static void $checkItem(Task v) {
    if (v is! Task) $pb.checkItemFailed(v, _i.messageName);
  }

  String get name => $_getS(0, '');
  set name(String v) { $_setString(0, v); }
  bool hasName() => $_has(0);
  void clearName() => clearField(1);

  Int64 get duration => $_getI64(1);
  set duration(Int64 v) { $_setInt64(1, v); }
  bool hasDuration() => $_has(1);
  void clearDuration() => clearField(2);

  Int64 get stress => $_getI64(2);
  set stress(Int64 v) { $_setInt64(2, v); }
  bool hasStress() => $_has(2);
  void clearStress() => clearField(3);

  bool get locationNull => $_get(3, false);
  set locationNull(bool v) { $_setBool(3, v); }
  bool hasLocationNull() => $_has(3);
  void clearLocationNull() => clearField(14);

  $0.Location get locationValue => $_getN(4);
  set locationValue($0.Location v) { setField(15, v); }
  bool hasLocationValue() => $_has(4);
  void clearLocationValue() => clearField(15);
}

