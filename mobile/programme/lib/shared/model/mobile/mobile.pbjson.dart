///
//  Generated code. Do not modify.
//  source: shared/model/mobile/mobile.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import '../schedule/schedule.pbjson.dart' as $0;
import '../routes/routes.pbjson.dart' as $1;
import '../weather/weather.pbjson.dart' as $2;

const MobileGateway$json = const {
  '1': 'MobileGateway',
  '2': const [
    const {'1': 'AddTask', '2': '.shared.model.schedule.NewTaskRequest', '3': '.shared.model.schedule.Boolean'},
    const {'1': 'GetForecast', '2': '.shared.model.routes.Coords', '3': '.shared.model.weather.Forecast'},
  ],
};

const MobileGateway$messageJson = const {
  '.shared.model.schedule.NewTaskRequest': $0.NewTaskRequest$json,
  '.shared.model.schedule.Task': $0.Task$json,
  '.shared.model.routes.Location': $1.Location$json,
  '.shared.model.routes.Coords': $1.Coords$json,
  '.shared.model.schedule.Boolean': $0.Boolean$json,
  '.shared.model.weather.Forecast': $2.Forecast$json,
  '.shared.model.weather.List': $2.List_$json,
  '.shared.model.weather.Main': $2.Main$json,
  '.shared.model.weather.Weather': $2.Weather$json,
  '.shared.model.weather.Clouds': $2.Clouds$json,
  '.shared.model.weather.Wind': $2.Wind$json,
  '.shared.model.weather.Rain': $2.Rain$json,
  '.shared.model.weather.Sys': $2.Sys$json,
  '.shared.model.weather.City': $2.City$json,
};

