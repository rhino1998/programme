///
//  Generated code. Do not modify.
//  source: shared/model/mobile/mobile.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async';

import 'package:protobuf/protobuf.dart';

import '../schedule/schedule.pb.dart' as $0;
import 'mobile.pbjson.dart';

export 'mobile.pb.dart';

abstract class MobileGatewayServiceBase extends GeneratedService {
  Future<$0.Boolean> addTask(ServerContext ctx, $0.NewTaskRequest request);

  GeneratedMessage createRequest(String method) {
    switch (method) {
      case 'AddTask': return new $0.NewTaskRequest();
      default: throw new ArgumentError('Unknown method: $method');
    }
  }

  Future<GeneratedMessage> handleCall(ServerContext ctx, String method, GeneratedMessage request) {
    switch (method) {
      case 'AddTask': return this.addTask(ctx, request);
      default: throw new ArgumentError('Unknown method: $method');
    }
  }

  Map<String, dynamic> get $json => MobileGateway$json;
  Map<String, Map<String, dynamic>> get $messageJson => MobileGateway$messageJson;
}

