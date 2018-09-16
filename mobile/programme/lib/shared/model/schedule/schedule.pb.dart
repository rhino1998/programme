///
//  Generated code. Do not modify.
//  source: shared/model/schedule/schedule.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async';
// ignore: UNUSED_SHOWN_NAME
import 'dart:core' show int, bool, double, String, List, override;

import 'package:fixnum/fixnum.dart';
import 'package:protobuf/protobuf.dart' as $pb;

import '../routes/routes.pb.dart' as $0;

import 'schedule.pbenum.dart';
import '../routes/routes.pbenum.dart' as $0;

export 'schedule.pbenum.dart';

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
    ..e<TaskType>(1, 'taskType', $pb.PbFieldType.OE, TaskType.Unknown, TaskType.valueOf, TaskType.values)
    ..aOS(2, 'name')
    ..aOS(3, 'description')
    ..aInt64(4, 'duration')
    ..aInt64(5, 'stress')
    ..aOB(8, 'deadlineNull')
    ..aInt64(9, 'deadlineValue')
    ..aOB(10, 'startNull')
    ..aInt64(11, 'startValue')
    ..aOB(12, 'travelMethodNull')
    ..e<$0.TravelMethod>(13, 'travelMethodValue', $pb.PbFieldType.OE, $0.TravelMethod.Unknown, $0.TravelMethod.valueOf, $0.TravelMethod.values)
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

  TaskType get taskType => $_getN(0);
  set taskType(TaskType v) { setField(1, v); }
  bool hasTaskType() => $_has(0);
  void clearTaskType() => clearField(1);

  String get name => $_getS(1, '');
  set name(String v) { $_setString(1, v); }
  bool hasName() => $_has(1);
  void clearName() => clearField(2);

  String get description => $_getS(2, '');
  set description(String v) { $_setString(2, v); }
  bool hasDescription() => $_has(2);
  void clearDescription() => clearField(3);

  Int64 get duration => $_getI64(3);
  set duration(Int64 v) { $_setInt64(3, v); }
  bool hasDuration() => $_has(3);
  void clearDuration() => clearField(4);

  Int64 get stress => $_getI64(4);
  set stress(Int64 v) { $_setInt64(4, v); }
  bool hasStress() => $_has(4);
  void clearStress() => clearField(5);

  bool get deadlineNull => $_get(5, false);
  set deadlineNull(bool v) { $_setBool(5, v); }
  bool hasDeadlineNull() => $_has(5);
  void clearDeadlineNull() => clearField(8);

  Int64 get deadlineValue => $_getI64(6);
  set deadlineValue(Int64 v) { $_setInt64(6, v); }
  bool hasDeadlineValue() => $_has(6);
  void clearDeadlineValue() => clearField(9);

  bool get startNull => $_get(7, false);
  set startNull(bool v) { $_setBool(7, v); }
  bool hasStartNull() => $_has(7);
  void clearStartNull() => clearField(10);

  Int64 get startValue => $_getI64(8);
  set startValue(Int64 v) { $_setInt64(8, v); }
  bool hasStartValue() => $_has(8);
  void clearStartValue() => clearField(11);

  bool get travelMethodNull => $_get(9, false);
  set travelMethodNull(bool v) { $_setBool(9, v); }
  bool hasTravelMethodNull() => $_has(9);
  void clearTravelMethodNull() => clearField(12);

  $0.TravelMethod get travelMethodValue => $_getN(10);
  set travelMethodValue($0.TravelMethod v) { setField(13, v); }
  bool hasTravelMethodValue() => $_has(10);
  void clearTravelMethodValue() => clearField(13);

  bool get locationNull => $_get(11, false);
  set locationNull(bool v) { $_setBool(11, v); }
  bool hasLocationNull() => $_has(11);
  void clearLocationNull() => clearField(14);

  $0.Location get locationValue => $_getN(12);
  set locationValue($0.Location v) { setField(15, v); }
  bool hasLocationValue() => $_has(12);
  void clearLocationValue() => clearField(15);
}

class NewTaskRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('NewTaskRequest', package: const $pb.PackageName('shared.model.schedule'))
    ..aOS(1, 'user')
    ..a<Task>(2, 'task', $pb.PbFieldType.OM, Task.getDefault, Task.create)
    ..hasRequiredFields = false
  ;

  NewTaskRequest() : super();
  NewTaskRequest.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  NewTaskRequest.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  NewTaskRequest clone() => new NewTaskRequest()..mergeFromMessage(this);
  NewTaskRequest copyWith(void Function(NewTaskRequest) updates) => super.copyWith((message) => updates(message as NewTaskRequest));
  $pb.BuilderInfo get info_ => _i;
  static NewTaskRequest create() => new NewTaskRequest();
  static $pb.PbList<NewTaskRequest> createRepeated() => new $pb.PbList<NewTaskRequest>();
  static NewTaskRequest getDefault() => _defaultInstance ??= create()..freeze();
  static NewTaskRequest _defaultInstance;
  static void $checkItem(NewTaskRequest v) {
    if (v is! NewTaskRequest) $pb.checkItemFailed(v, _i.messageName);
  }

  String get user => $_getS(0, '');
  set user(String v) { $_setString(0, v); }
  bool hasUser() => $_has(0);
  void clearUser() => clearField(1);

  Task get task => $_getN(1);
  set task(Task v) { setField(2, v); }
  bool hasTask() => $_has(1);
  void clearTask() => clearField(2);
}

class Boolean extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Boolean', package: const $pb.PackageName('shared.model.schedule'))
    ..aOB(1, 'boolean')
    ..hasRequiredFields = false
  ;

  Boolean() : super();
  Boolean.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Boolean.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Boolean clone() => new Boolean()..mergeFromMessage(this);
  Boolean copyWith(void Function(Boolean) updates) => super.copyWith((message) => updates(message as Boolean));
  $pb.BuilderInfo get info_ => _i;
  static Boolean create() => new Boolean();
  static $pb.PbList<Boolean> createRepeated() => new $pb.PbList<Boolean>();
  static Boolean getDefault() => _defaultInstance ??= create()..freeze();
  static Boolean _defaultInstance;
  static void $checkItem(Boolean v) {
    if (v is! Boolean) $pb.checkItemFailed(v, _i.messageName);
  }

  bool get boolean => $_get(0, false);
  set boolean(bool v) { $_setBool(0, v); }
  bool hasBoolean() => $_has(0);
  void clearBoolean() => clearField(1);
}

class TaskManagerApi {
  $pb.RpcClient _client;
  TaskManagerApi(this._client);

  Future<Boolean> addTask($pb.ClientContext ctx, NewTaskRequest request) {
    var emptyResponse = new Boolean();
    return _client.invoke<Boolean>(ctx, 'TaskManager', 'AddTask', request, emptyResponse);
  }
}

