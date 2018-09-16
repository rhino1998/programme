///
//  Generated code. Do not modify.
//  source: shared/model/weather/weather.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async';

import 'package:grpc/grpc.dart';

import '../routes/routes.pb.dart' as $0;
import 'weather.pb.dart';
export 'weather.pb.dart';

class WeatherAPIClient extends Client {
  static final _$getForecast = new ClientMethod<$0.Coords, Forecast>(
      '/shared.model.weather.WeatherAPI/GetForecast',
      ($0.Coords value) => value.writeToBuffer(),
      (List<int> value) => new Forecast.fromBuffer(value));

  WeatherAPIClient(ClientChannel channel, {CallOptions options})
      : super(channel, options: options);

  ResponseFuture<Forecast> getForecast($0.Coords request,
      {CallOptions options}) {
    final call = $createCall(_$getForecast, new Stream.fromIterable([request]),
        options: options);
    return new ResponseFuture(call);
  }
}

abstract class WeatherAPIServiceBase extends Service {
  String get $name => 'shared.model.weather.WeatherAPI';

  WeatherAPIServiceBase() {
    $addMethod(new ServiceMethod<$0.Coords, Forecast>(
        'GetForecast',
        getForecast_Pre,
        false,
        false,
        (List<int> value) => new $0.Coords.fromBuffer(value),
        (Forecast value) => value.writeToBuffer()));
  }

  Future<Forecast> getForecast_Pre(ServiceCall call, Future request) async {
    return getForecast(call, await request);
  }

  Future<Forecast> getForecast(ServiceCall call, $0.Coords request);
}
