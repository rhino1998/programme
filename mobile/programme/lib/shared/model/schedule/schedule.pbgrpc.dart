///
//  Generated code. Do not modify.
//  source: shared/model/schedule/schedule.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async';

import 'package:grpc/grpc.dart';

import 'schedule.pb.dart';
export 'schedule.pb.dart';

class TaskManagerClient extends Client {
  static final _$addTask = new ClientMethod<NewTaskRequest, Boolean>(
      '/shared.model.schedule.TaskManager/AddTask',
      (NewTaskRequest value) => value.writeToBuffer(),
      (List<int> value) => new Boolean.fromBuffer(value));

  TaskManagerClient(ClientChannel channel, {CallOptions options})
      : super(channel, options: options);

  ResponseFuture<Boolean> addTask(NewTaskRequest request,
      {CallOptions options}) {
    final call = $createCall(_$addTask, new Stream.fromIterable([request]),
        options: options);
    return new ResponseFuture(call);
  }
}

abstract class TaskManagerServiceBase extends Service {
  String get $name => 'shared.model.schedule.TaskManager';

  TaskManagerServiceBase() {
    $addMethod(new ServiceMethod<NewTaskRequest, Boolean>(
        'AddTask',
        addTask_Pre,
        false,
        false,
        (List<int> value) => new NewTaskRequest.fromBuffer(value),
        (Boolean value) => value.writeToBuffer()));
  }

  Future<Boolean> addTask_Pre(ServiceCall call, Future request) async {
    return addTask(call, await request);
  }

  Future<Boolean> addTask(ServiceCall call, NewTaskRequest request);
}
