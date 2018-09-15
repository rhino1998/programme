///
//  Generated code. Do not modify.
//  source: shared/model/schedule/schedule.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' show int, dynamic, String, List, Map;
import 'package:protobuf/protobuf.dart' as $pb;

class TaskType extends $pb.ProtobufEnum {
  static const TaskType Unknown = const TaskType._(0, 'Unknown');
  static const TaskType Floating = const TaskType._(1, 'Floating');
  static const TaskType Scheduled = const TaskType._(2, 'Scheduled');
  static const TaskType Calendar = const TaskType._(3, 'Calendar');
  static const TaskType Free = const TaskType._(4, 'Free');
  static const TaskType Travel = const TaskType._(5, 'Travel');

  static const List<TaskType> values = const <TaskType> [
    Unknown,
    Floating,
    Scheduled,
    Calendar,
    Free,
    Travel,
  ];

  static final Map<int, dynamic> _byValue = $pb.ProtobufEnum.initByValue(values);
  static TaskType valueOf(int value) => _byValue[value] as TaskType;
  static void $checkItem(TaskType v) {
    if (v is! TaskType) $pb.checkItemFailed(v, 'TaskType');
  }

  const TaskType._(int v, String n) : super(v, n);
}

