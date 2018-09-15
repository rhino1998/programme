///
//  Generated code. Do not modify.
//  source: shared/model/schedule/schedule.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async';

import 'package:protobuf/protobuf.dart';

import 'schedule.pb.dart';
import 'schedule.pbjson.dart';

export 'schedule.pb.dart';

abstract class TaskManagerServiceBase extends GeneratedService {
  Future<Boolean> addTask(ServerContext ctx, NewTaskRequest request);

  GeneratedMessage createRequest(String method) {
    switch (method) {
      case 'AddTask': return new NewTaskRequest();
      default: throw new ArgumentError('Unknown method: $method');
    }
  }

  Future<GeneratedMessage> handleCall(ServerContext ctx, String method, GeneratedMessage request) {
    switch (method) {
      case 'AddTask': return this.addTask(ctx, request);
      default: throw new ArgumentError('Unknown method: $method');
    }
  }

  Map<String, dynamic> get $json => TaskManager$json;
  Map<String, Map<String, dynamic>> get $messageJson => TaskManager$messageJson;
}

