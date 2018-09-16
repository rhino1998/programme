///
//  Generated code. Do not modify.
//  source: shared/model/weather/weather.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async';

import 'package:protobuf/protobuf.dart';

import '../routes/routes.pb.dart' as $0;
import 'weather.pb.dart';
import 'weather.pbjson.dart';

export 'weather.pb.dart';

abstract class WeatherAPIServiceBase extends GeneratedService {
  Future<Forecast> getForecast(ServerContext ctx, $0.Coords request);

  GeneratedMessage createRequest(String method) {
    switch (method) {
      case 'GetForecast': return new $0.Coords();
      default: throw new ArgumentError('Unknown method: $method');
    }
  }

  Future<GeneratedMessage> handleCall(ServerContext ctx, String method, GeneratedMessage request) {
    switch (method) {
      case 'GetForecast': return this.getForecast(ctx, request);
      default: throw new ArgumentError('Unknown method: $method');
    }
  }

  Map<String, dynamic> get $json => WeatherAPI$json;
  Map<String, Map<String, dynamic>> get $messageJson => WeatherAPI$messageJson;
}

