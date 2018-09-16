///
//  Generated code. Do not modify.
//  source: shared/model/weather/weather.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

const Coord$json = const {
  '1': 'Coord',
  '2': const [
    const {'1': 'longitude', '3': 1, '4': 1, '5': 1, '10': 'longitude'},
    const {'1': 'latitude', '3': 2, '4': 1, '5': 1, '10': 'latitude'},
  ],
};

const Sys$json = const {
  '1': 'Sys',
  '2': const [
    const {'1': 'country', '3': 1, '4': 1, '5': 9, '10': 'country'},
    const {'1': 'sunrise', '3': 2, '4': 1, '5': 4, '10': 'sunrise'},
    const {'1': 'sunset', '3': 3, '4': 1, '5': 4, '10': 'sunset'},
  ],
};

const Main$json = const {
  '1': 'Main',
  '2': const [
    const {'1': 'temperature', '3': 1, '4': 1, '5': 1, '10': 'temperature'},
    const {'1': 'temp_min', '3': 2, '4': 1, '5': 1, '10': 'tempMin'},
    const {'1': 'temp_max', '3': 3, '4': 1, '5': 1, '10': 'tempMax'},
    const {'1': 'pressure', '3': 4, '4': 1, '5': 1, '10': 'pressure'},
    const {'1': 'sea_level', '3': 5, '4': 1, '5': 1, '10': 'seaLevel'},
    const {'1': 'ground_level', '3': 6, '4': 1, '5': 1, '10': 'groundLevel'},
    const {'1': 'humidity', '3': 7, '4': 1, '5': 3, '10': 'humidity'},
  ],
};

const Weather$json = const {
  '1': 'Weather',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'main', '3': 2, '4': 1, '5': 9, '10': 'main'},
    const {'1': 'description', '3': 3, '4': 1, '5': 9, '10': 'description'},
  ],
};

const Wind$json = const {
  '1': 'Wind',
  '2': const [
    const {'1': 'speed', '3': 1, '4': 1, '5': 1, '10': 'speed'},
    const {'1': 'deg', '3': 2, '4': 1, '5': 1, '10': 'deg'},
  ],
};

const Rain$json = const {
  '1': 'Rain',
  '2': const [
    const {'1': 'three_hours', '3': 1, '4': 1, '5': 1, '10': 'threeHours'},
  ],
};

const Snow$json = const {
  '1': 'Snow',
  '2': const [
    const {'1': 'three_hours', '3': 1, '4': 1, '5': 1, '10': 'threeHours'},
  ],
};

const Clouds$json = const {
  '1': 'Clouds',
  '2': const [
    const {'1': 'all', '3': 1, '4': 1, '5': 3, '10': 'all'},
  ],
};

const List_$json = const {
  '1': 'List',
  '2': const [
    const {'1': 'dt', '3': 1, '4': 1, '5': 3, '10': 'dt'},
    const {'1': 'main', '3': 2, '4': 1, '5': 11, '6': '.shared.model.weather.Main', '10': 'main'},
    const {'1': 'weather', '3': 3, '4': 3, '5': 11, '6': '.shared.model.weather.Weather', '10': 'weather'},
    const {'1': 'clouds', '3': 4, '4': 1, '5': 11, '6': '.shared.model.weather.Clouds', '10': 'clouds'},
    const {'1': 'wind', '3': 5, '4': 1, '5': 11, '6': '.shared.model.weather.Wind', '10': 'wind'},
    const {'1': 'rain', '3': 6, '4': 1, '5': 11, '6': '.shared.model.weather.Rain', '10': 'rain'},
    const {'1': 'sys', '3': 7, '4': 1, '5': 11, '6': '.shared.model.weather.Sys', '10': 'sys'},
  ],
};

const City$json = const {
  '1': 'City',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'coord', '3': 3, '4': 1, '5': 11, '6': '.shared.model.weather.Coord', '10': 'coord'},
    const {'1': 'country', '3': 4, '4': 1, '5': 9, '10': 'country'},
  ],
};

const Forecast$json = const {
  '1': 'Forecast',
  '2': const [
    const {'1': 'list', '3': 1, '4': 3, '5': 11, '6': '.shared.model.weather.List', '10': 'list'},
    const {'1': 'city', '3': 2, '4': 1, '5': 11, '6': '.shared.model.weather.City', '10': 'city'},
  ],
};

const WeatherAPI$json = const {
  '1': 'WeatherAPI',
  '2': const [
    const {'1': 'GetForecast', '2': '.shared.model.weather.Coord', '3': '.shared.model.weather.Forecast'},
  ],
};

const WeatherAPI$messageJson = const {
  '.shared.model.weather.Coord': Coord$json,
  '.shared.model.weather.Forecast': Forecast$json,
  '.shared.model.weather.List': List_$json,
  '.shared.model.weather.Main': Main$json,
  '.shared.model.weather.Weather': Weather$json,
  '.shared.model.weather.Clouds': Clouds$json,
  '.shared.model.weather.Wind': Wind$json,
  '.shared.model.weather.Rain': Rain$json,
  '.shared.model.weather.Sys': Sys$json,
  '.shared.model.weather.City': City$json,
};

