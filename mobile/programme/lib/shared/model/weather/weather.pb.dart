///
//  Generated code. Do not modify.
//  source: shared/model/weather/weather.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async';
// ignore: UNUSED_SHOWN_NAME
import 'dart:core' show int, bool, double, String, List, override;

import 'package:fixnum/fixnum.dart';
import 'package:protobuf/protobuf.dart' as $pb;

class Coord extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Coord', package: const $pb.PackageName('shared.model.weather'))
    ..a<double>(1, 'longitude', $pb.PbFieldType.OD)
    ..a<double>(2, 'latitude', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  Coord() : super();
  Coord.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Coord.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Coord clone() => new Coord()..mergeFromMessage(this);
  Coord copyWith(void Function(Coord) updates) => super.copyWith((message) => updates(message as Coord));
  $pb.BuilderInfo get info_ => _i;
  static Coord create() => new Coord();
  static $pb.PbList<Coord> createRepeated() => new $pb.PbList<Coord>();
  static Coord getDefault() => _defaultInstance ??= create()..freeze();
  static Coord _defaultInstance;
  static void $checkItem(Coord v) {
    if (v is! Coord) $pb.checkItemFailed(v, _i.messageName);
  }

  double get longitude => $_getN(0);
  set longitude(double v) { $_setDouble(0, v); }
  bool hasLongitude() => $_has(0);
  void clearLongitude() => clearField(1);

  double get latitude => $_getN(1);
  set latitude(double v) { $_setDouble(1, v); }
  bool hasLatitude() => $_has(1);
  void clearLatitude() => clearField(2);
}

class Sys extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Sys', package: const $pb.PackageName('shared.model.weather'))
    ..aOS(1, 'country')
    ..a<Int64>(2, 'sunrise', $pb.PbFieldType.OU6, Int64.ZERO)
    ..a<Int64>(3, 'sunset', $pb.PbFieldType.OU6, Int64.ZERO)
    ..hasRequiredFields = false
  ;

  Sys() : super();
  Sys.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Sys.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Sys clone() => new Sys()..mergeFromMessage(this);
  Sys copyWith(void Function(Sys) updates) => super.copyWith((message) => updates(message as Sys));
  $pb.BuilderInfo get info_ => _i;
  static Sys create() => new Sys();
  static $pb.PbList<Sys> createRepeated() => new $pb.PbList<Sys>();
  static Sys getDefault() => _defaultInstance ??= create()..freeze();
  static Sys _defaultInstance;
  static void $checkItem(Sys v) {
    if (v is! Sys) $pb.checkItemFailed(v, _i.messageName);
  }

  String get country => $_getS(0, '');
  set country(String v) { $_setString(0, v); }
  bool hasCountry() => $_has(0);
  void clearCountry() => clearField(1);

  Int64 get sunrise => $_getI64(1);
  set sunrise(Int64 v) { $_setInt64(1, v); }
  bool hasSunrise() => $_has(1);
  void clearSunrise() => clearField(2);

  Int64 get sunset => $_getI64(2);
  set sunset(Int64 v) { $_setInt64(2, v); }
  bool hasSunset() => $_has(2);
  void clearSunset() => clearField(3);
}

class Main extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Main', package: const $pb.PackageName('shared.model.weather'))
    ..a<double>(1, 'temperature', $pb.PbFieldType.OD)
    ..a<double>(2, 'tempMin', $pb.PbFieldType.OD)
    ..a<double>(3, 'tempMax', $pb.PbFieldType.OD)
    ..a<double>(4, 'pressure', $pb.PbFieldType.OD)
    ..a<double>(5, 'seaLevel', $pb.PbFieldType.OD)
    ..a<double>(6, 'groundLevel', $pb.PbFieldType.OD)
    ..aInt64(7, 'humidity')
    ..hasRequiredFields = false
  ;

  Main() : super();
  Main.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Main.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Main clone() => new Main()..mergeFromMessage(this);
  Main copyWith(void Function(Main) updates) => super.copyWith((message) => updates(message as Main));
  $pb.BuilderInfo get info_ => _i;
  static Main create() => new Main();
  static $pb.PbList<Main> createRepeated() => new $pb.PbList<Main>();
  static Main getDefault() => _defaultInstance ??= create()..freeze();
  static Main _defaultInstance;
  static void $checkItem(Main v) {
    if (v is! Main) $pb.checkItemFailed(v, _i.messageName);
  }

  double get temperature => $_getN(0);
  set temperature(double v) { $_setDouble(0, v); }
  bool hasTemperature() => $_has(0);
  void clearTemperature() => clearField(1);

  double get tempMin => $_getN(1);
  set tempMin(double v) { $_setDouble(1, v); }
  bool hasTempMin() => $_has(1);
  void clearTempMin() => clearField(2);

  double get tempMax => $_getN(2);
  set tempMax(double v) { $_setDouble(2, v); }
  bool hasTempMax() => $_has(2);
  void clearTempMax() => clearField(3);

  double get pressure => $_getN(3);
  set pressure(double v) { $_setDouble(3, v); }
  bool hasPressure() => $_has(3);
  void clearPressure() => clearField(4);

  double get seaLevel => $_getN(4);
  set seaLevel(double v) { $_setDouble(4, v); }
  bool hasSeaLevel() => $_has(4);
  void clearSeaLevel() => clearField(5);

  double get groundLevel => $_getN(5);
  set groundLevel(double v) { $_setDouble(5, v); }
  bool hasGroundLevel() => $_has(5);
  void clearGroundLevel() => clearField(6);

  Int64 get humidity => $_getI64(6);
  set humidity(Int64 v) { $_setInt64(6, v); }
  bool hasHumidity() => $_has(6);
  void clearHumidity() => clearField(7);
}

class Weather extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Weather', package: const $pb.PackageName('shared.model.weather'))
    ..aInt64(1, 'id')
    ..aOS(2, 'main')
    ..aOS(3, 'description')
    ..hasRequiredFields = false
  ;

  Weather() : super();
  Weather.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Weather.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Weather clone() => new Weather()..mergeFromMessage(this);
  Weather copyWith(void Function(Weather) updates) => super.copyWith((message) => updates(message as Weather));
  $pb.BuilderInfo get info_ => _i;
  static Weather create() => new Weather();
  static $pb.PbList<Weather> createRepeated() => new $pb.PbList<Weather>();
  static Weather getDefault() => _defaultInstance ??= create()..freeze();
  static Weather _defaultInstance;
  static void $checkItem(Weather v) {
    if (v is! Weather) $pb.checkItemFailed(v, _i.messageName);
  }

  Int64 get id => $_getI64(0);
  set id(Int64 v) { $_setInt64(0, v); }
  bool hasId() => $_has(0);
  void clearId() => clearField(1);

  String get main => $_getS(1, '');
  set main(String v) { $_setString(1, v); }
  bool hasMain() => $_has(1);
  void clearMain() => clearField(2);

  String get description => $_getS(2, '');
  set description(String v) { $_setString(2, v); }
  bool hasDescription() => $_has(2);
  void clearDescription() => clearField(3);
}

class Wind extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Wind', package: const $pb.PackageName('shared.model.weather'))
    ..a<double>(1, 'speed', $pb.PbFieldType.OD)
    ..a<double>(2, 'deg', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  Wind() : super();
  Wind.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Wind.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Wind clone() => new Wind()..mergeFromMessage(this);
  Wind copyWith(void Function(Wind) updates) => super.copyWith((message) => updates(message as Wind));
  $pb.BuilderInfo get info_ => _i;
  static Wind create() => new Wind();
  static $pb.PbList<Wind> createRepeated() => new $pb.PbList<Wind>();
  static Wind getDefault() => _defaultInstance ??= create()..freeze();
  static Wind _defaultInstance;
  static void $checkItem(Wind v) {
    if (v is! Wind) $pb.checkItemFailed(v, _i.messageName);
  }

  double get speed => $_getN(0);
  set speed(double v) { $_setDouble(0, v); }
  bool hasSpeed() => $_has(0);
  void clearSpeed() => clearField(1);

  double get deg => $_getN(1);
  set deg(double v) { $_setDouble(1, v); }
  bool hasDeg() => $_has(1);
  void clearDeg() => clearField(2);
}

class Rain extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Rain', package: const $pb.PackageName('shared.model.weather'))
    ..a<double>(1, 'threeHours', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  Rain() : super();
  Rain.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Rain.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Rain clone() => new Rain()..mergeFromMessage(this);
  Rain copyWith(void Function(Rain) updates) => super.copyWith((message) => updates(message as Rain));
  $pb.BuilderInfo get info_ => _i;
  static Rain create() => new Rain();
  static $pb.PbList<Rain> createRepeated() => new $pb.PbList<Rain>();
  static Rain getDefault() => _defaultInstance ??= create()..freeze();
  static Rain _defaultInstance;
  static void $checkItem(Rain v) {
    if (v is! Rain) $pb.checkItemFailed(v, _i.messageName);
  }

  double get threeHours => $_getN(0);
  set threeHours(double v) { $_setDouble(0, v); }
  bool hasThreeHours() => $_has(0);
  void clearThreeHours() => clearField(1);
}

class Snow extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Snow', package: const $pb.PackageName('shared.model.weather'))
    ..a<double>(1, 'threeHours', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  Snow() : super();
  Snow.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Snow.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Snow clone() => new Snow()..mergeFromMessage(this);
  Snow copyWith(void Function(Snow) updates) => super.copyWith((message) => updates(message as Snow));
  $pb.BuilderInfo get info_ => _i;
  static Snow create() => new Snow();
  static $pb.PbList<Snow> createRepeated() => new $pb.PbList<Snow>();
  static Snow getDefault() => _defaultInstance ??= create()..freeze();
  static Snow _defaultInstance;
  static void $checkItem(Snow v) {
    if (v is! Snow) $pb.checkItemFailed(v, _i.messageName);
  }

  double get threeHours => $_getN(0);
  set threeHours(double v) { $_setDouble(0, v); }
  bool hasThreeHours() => $_has(0);
  void clearThreeHours() => clearField(1);
}

class Clouds extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Clouds', package: const $pb.PackageName('shared.model.weather'))
    ..aInt64(1, 'all')
    ..hasRequiredFields = false
  ;

  Clouds() : super();
  Clouds.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Clouds.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Clouds clone() => new Clouds()..mergeFromMessage(this);
  Clouds copyWith(void Function(Clouds) updates) => super.copyWith((message) => updates(message as Clouds));
  $pb.BuilderInfo get info_ => _i;
  static Clouds create() => new Clouds();
  static $pb.PbList<Clouds> createRepeated() => new $pb.PbList<Clouds>();
  static Clouds getDefault() => _defaultInstance ??= create()..freeze();
  static Clouds _defaultInstance;
  static void $checkItem(Clouds v) {
    if (v is! Clouds) $pb.checkItemFailed(v, _i.messageName);
  }

  Int64 get all => $_getI64(0);
  set all(Int64 v) { $_setInt64(0, v); }
  bool hasAll() => $_has(0);
  void clearAll() => clearField(1);
}

class List_ extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('List', package: const $pb.PackageName('shared.model.weather'))
    ..aInt64(1, 'dt')
    ..a<Main>(2, 'main', $pb.PbFieldType.OM, Main.getDefault, Main.create)
    ..pp<Weather>(3, 'weather', $pb.PbFieldType.PM, Weather.$checkItem, Weather.create)
    ..a<Clouds>(4, 'clouds', $pb.PbFieldType.OM, Clouds.getDefault, Clouds.create)
    ..a<Wind>(5, 'wind', $pb.PbFieldType.OM, Wind.getDefault, Wind.create)
    ..a<Rain>(6, 'rain', $pb.PbFieldType.OM, Rain.getDefault, Rain.create)
    ..a<Sys>(7, 'sys', $pb.PbFieldType.OM, Sys.getDefault, Sys.create)
    ..hasRequiredFields = false
  ;

  List_() : super();
  List_.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  List_.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  List_ clone() => new List_()..mergeFromMessage(this);
  List_ copyWith(void Function(List_) updates) => super.copyWith((message) => updates(message as List_));
  $pb.BuilderInfo get info_ => _i;
  static List_ create() => new List_();
  static $pb.PbList<List_> createRepeated() => new $pb.PbList<List_>();
  static List_ getDefault() => _defaultInstance ??= create()..freeze();
  static List_ _defaultInstance;
  static void $checkItem(List_ v) {
    if (v is! List_) $pb.checkItemFailed(v, _i.messageName);
  }

  Int64 get dt => $_getI64(0);
  set dt(Int64 v) { $_setInt64(0, v); }
  bool hasDt() => $_has(0);
  void clearDt() => clearField(1);

  Main get main => $_getN(1);
  set main(Main v) { setField(2, v); }
  bool hasMain() => $_has(1);
  void clearMain() => clearField(2);

  List<Weather> get weather => $_getList(2);

  Clouds get clouds => $_getN(3);
  set clouds(Clouds v) { setField(4, v); }
  bool hasClouds() => $_has(3);
  void clearClouds() => clearField(4);

  Wind get wind => $_getN(4);
  set wind(Wind v) { setField(5, v); }
  bool hasWind() => $_has(4);
  void clearWind() => clearField(5);

  Rain get rain => $_getN(5);
  set rain(Rain v) { setField(6, v); }
  bool hasRain() => $_has(5);
  void clearRain() => clearField(6);

  Sys get sys => $_getN(6);
  set sys(Sys v) { setField(7, v); }
  bool hasSys() => $_has(6);
  void clearSys() => clearField(7);
}

class City extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('City', package: const $pb.PackageName('shared.model.weather'))
    ..aInt64(1, 'id')
    ..aOS(2, 'name')
    ..a<Coord>(3, 'coord', $pb.PbFieldType.OM, Coord.getDefault, Coord.create)
    ..aOS(4, 'country')
    ..hasRequiredFields = false
  ;

  City() : super();
  City.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  City.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  City clone() => new City()..mergeFromMessage(this);
  City copyWith(void Function(City) updates) => super.copyWith((message) => updates(message as City));
  $pb.BuilderInfo get info_ => _i;
  static City create() => new City();
  static $pb.PbList<City> createRepeated() => new $pb.PbList<City>();
  static City getDefault() => _defaultInstance ??= create()..freeze();
  static City _defaultInstance;
  static void $checkItem(City v) {
    if (v is! City) $pb.checkItemFailed(v, _i.messageName);
  }

  Int64 get id => $_getI64(0);
  set id(Int64 v) { $_setInt64(0, v); }
  bool hasId() => $_has(0);
  void clearId() => clearField(1);

  String get name => $_getS(1, '');
  set name(String v) { $_setString(1, v); }
  bool hasName() => $_has(1);
  void clearName() => clearField(2);

  Coord get coord => $_getN(2);
  set coord(Coord v) { setField(3, v); }
  bool hasCoord() => $_has(2);
  void clearCoord() => clearField(3);

  String get country => $_getS(3, '');
  set country(String v) { $_setString(3, v); }
  bool hasCountry() => $_has(3);
  void clearCountry() => clearField(4);
}

class Forecast extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Forecast', package: const $pb.PackageName('shared.model.weather'))
    ..pp<List_>(1, 'list', $pb.PbFieldType.PM, List_.$checkItem, List_.create)
    ..a<City>(2, 'city', $pb.PbFieldType.OM, City.getDefault, City.create)
    ..hasRequiredFields = false
  ;

  Forecast() : super();
  Forecast.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Forecast.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Forecast clone() => new Forecast()..mergeFromMessage(this);
  Forecast copyWith(void Function(Forecast) updates) => super.copyWith((message) => updates(message as Forecast));
  $pb.BuilderInfo get info_ => _i;
  static Forecast create() => new Forecast();
  static $pb.PbList<Forecast> createRepeated() => new $pb.PbList<Forecast>();
  static Forecast getDefault() => _defaultInstance ??= create()..freeze();
  static Forecast _defaultInstance;
  static void $checkItem(Forecast v) {
    if (v is! Forecast) $pb.checkItemFailed(v, _i.messageName);
  }

  List<List_> get list => $_getList(0);

  City get city => $_getN(1);
  set city(City v) { setField(2, v); }
  bool hasCity() => $_has(1);
  void clearCity() => clearField(2);
}

class WeatherAPIApi {
  $pb.RpcClient _client;
  WeatherAPIApi(this._client);

  Future<Forecast> getForecast($pb.ClientContext ctx, Coord request) {
    var emptyResponse = new Forecast();
    return _client.invoke<Forecast>(ctx, 'WeatherAPI', 'GetForecast', request, emptyResponse);
  }
}

class WeatherAPIApi {
  $pb.RpcClient _client;
  WeatherAPIApi(this._client);

  Future<Forecast> getForecast($pb.ClientContext ctx, Coord request) {
    var emptyResponse = new Forecast();
    return _client.invoke<Forecast>(ctx, 'WeatherAPI', 'GetForecast', request, emptyResponse);
  }
}

