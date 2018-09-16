import "package:grpc/grpc.dart";
import 'shared/model/mobile/mobile.pbgrpc.dart';
import 'shared/model/routes/routes.pb.dart' as $0;
import 'shared/model/weather/weather.pb.dart';


final channel = new ClientChannel(
    '10.194.123.158',
      port: 8080,
      options: const ChannelOptions(
          credentials: const ChannelCredentials.insecure()));
  final stub = new MobileGatewayClient(channel);

ResponseFuture<Forecast> getWeather() {
  var c = $0.Coords();
  c.latitude = 38.9072;
  c.longitude = -77.0369;
  return stub.getForecast(c);
}