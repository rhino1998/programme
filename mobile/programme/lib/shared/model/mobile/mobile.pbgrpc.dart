///
//  Generated code. Do not modify.
//  source: shared/model/mobile/mobile.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async';

import 'package:grpc/grpc.dart';

import '../schedule/schedule.pb.dart' as $1;
import '../routes/routes.pb.dart' as $0;
import '../weather/weather.pb.dart' as $2;
export 'mobile.pb.dart';

class MobileGatewayClient extends Client {
  static final _$addTask = new ClientMethod<$1.NewTaskRequest, $1.Boolean>(
      '/shared.model.mobile.MobileGateway/AddTask',
      ($1.NewTaskRequest value) => value.writeToBuffer(),
      (List<int> value) => new $1.Boolean.fromBuffer(value));
  static final _$getForecast = new ClientMethod<$0.Coords, $2.Forecast>(
      '/shared.model.mobile.MobileGateway/GetForecast',
      ($0.Coords value) => value.writeToBuffer(),
      (List<int> value) => new $2.Forecast.fromBuffer(value));

  MobileGatewayClient(ClientChannel channel, {CallOptions options})
      : super(channel, options: options);

  ResponseFuture<$1.Boolean> addTask($1.NewTaskRequest request,
      {CallOptions options}) {
    final call = $createCall(_$addTask, new Stream.fromIterable([request]),
        options: options);
    return new ResponseFuture(call);
  }

  ResponseFuture<$2.Forecast> getForecast($0.Coords request,
      {CallOptions options}) {
    final call = $createCall(_$getForecast, new Stream.fromIterable([request]),
        options: options);
    return new ResponseFuture(call);
  }
}

abstract class MobileGatewayServiceBase extends Service {
  String get $name => 'shared.model.mobile.MobileGateway';

  MobileGatewayServiceBase() {
    $addMethod(new ServiceMethod<$1.NewTaskRequest, $1.Boolean>(
        'AddTask',
        addTask_Pre,
        false,
        false,
        (List<int> value) => new $1.NewTaskRequest.fromBuffer(value),
        ($1.Boolean value) => value.writeToBuffer()));
    $addMethod(new ServiceMethod<$0.Coords, $2.Forecast>(
        'GetForecast',
        getForecast_Pre,
        false,
        false,
        (List<int> value) => new $0.Coords.fromBuffer(value),
        ($2.Forecast value) => value.writeToBuffer()));
  }

  Future<$1.Boolean> addTask_Pre(ServiceCall call, Future request) async {
    return addTask(call, await request);
  }

  Future<$2.Forecast> getForecast_Pre(ServiceCall call, Future request) async {
    return getForecast(call, await request);
  }

  Future<$1.Boolean> addTask(ServiceCall call, $1.NewTaskRequest request);
  Future<$2.Forecast> getForecast(ServiceCall call, $0.Coords request);
}
