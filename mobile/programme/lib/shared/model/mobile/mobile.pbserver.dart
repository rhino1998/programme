///
//  Generated code. Do not modify.
//  source: shared/model/mobile/mobile.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async';

import 'package:protobuf/protobuf.dart';

import '../schedule/schedule.pb.dart' as $0;
import '../routes/routes.pb.dart' as $1;
import '../weather/weather.pb.dart' as $2;
import 'mobile.pbjson.dart';

export 'mobile.pb.dart';

abstract class MobileGatewayServiceBase extends GeneratedService {
  Future<$0.Boolean> addTask(ServerContext ctx, $0.NewTaskRequest request);
  Future<$2.Forecast> getForecast(ServerContext ctx, $1.Coords request);

  GeneratedMessage createRequest(String method) {
    switch (method) {
      case 'AddTask': return new $0.NewTaskRequest();
      case 'GetForecast': return new $1.Coords();
      default: throw new ArgumentError('Unknown method: $method');
    }
  }

  Future<GeneratedMessage> handleCall(ServerContext ctx, String method, GeneratedMessage request) {
    switch (method) {
      case 'AddTask': return this.addTask(ctx, request);
      case 'GetForecast': return this.getForecast(ctx, request);
      default: throw new ArgumentError('Unknown method: $method');
    }
  }

  Map<String, dynamic> get $json => MobileGateway$json;
  Map<String, Map<String, dynamic>> get $messageJson => MobileGateway$messageJson;
}

