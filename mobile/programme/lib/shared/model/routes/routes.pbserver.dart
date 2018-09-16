///
//  Generated code. Do not modify.
//  source: shared/model/routes/routes.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async';

import 'package:protobuf/protobuf.dart';

import 'routes.pb.dart';
import 'routes.pbjson.dart';

export 'routes.pb.dart';

abstract class RoutesAPIServiceBase extends GeneratedService {
  Future<Trips> calcTravelTime(ServerContext ctx, Trip request);

  GeneratedMessage createRequest(String method) {
    switch (method) {
      case 'CalcTravelTime': return new Trip();
      default: throw new ArgumentError('Unknown method: $method');
    }
  }

  Future<GeneratedMessage> handleCall(ServerContext ctx, String method, GeneratedMessage request) {
    switch (method) {
      case 'CalcTravelTime': return this.calcTravelTime(ctx, request);
      default: throw new ArgumentError('Unknown method: $method');
    }
  }

  Map<String, dynamic> get $json => RoutesAPI$json;
  Map<String, Map<String, dynamic>> get $messageJson => RoutesAPI$messageJson;
}

