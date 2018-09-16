///
//  Generated code. Do not modify.
//  source: shared/model/routes/routes.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async';

import 'package:grpc/grpc.dart';

import 'routes.pb.dart';
export 'routes.pb.dart';

class RoutesAPIClient extends Client {
  static final _$calcTravelTime = new ClientMethod<Trip, Trips>(
      '/shared.model.routes.RoutesAPI/CalcTravelTime',
      (Trip value) => value.writeToBuffer(),
      (List<int> value) => new Trips.fromBuffer(value));

  RoutesAPIClient(ClientChannel channel, {CallOptions options})
      : super(channel, options: options);

  ResponseFuture<Trips> calcTravelTime(Trip request, {CallOptions options}) {
    final call = $createCall(
        _$calcTravelTime, new Stream.fromIterable([request]),
        options: options);
    return new ResponseFuture(call);
  }
}

abstract class RoutesAPIServiceBase extends Service {
  String get $name => 'shared.model.routes.RoutesAPI';

  RoutesAPIServiceBase() {
    $addMethod(new ServiceMethod<Trip, Trips>(
        'CalcTravelTime',
        calcTravelTime_Pre,
        false,
        false,
        (List<int> value) => new Trip.fromBuffer(value),
        (Trips value) => value.writeToBuffer()));
  }

  Future<Trips> calcTravelTime_Pre(ServiceCall call, Future request) async {
    return calcTravelTime(call, await request);
  }

  Future<Trips> calcTravelTime(ServiceCall call, Trip request);
}
