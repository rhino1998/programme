///
//  Generated code. Do not modify.
//  source: shared/model/routes/routes.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async';
// ignore: UNUSED_SHOWN_NAME
import 'dart:core' show int, bool, double, String, List, override;

import 'package:fixnum/fixnum.dart';
import 'package:protobuf/protobuf.dart' as $pb;

import 'routes.pbenum.dart';

export 'routes.pbenum.dart';

class Coords extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Coords', package: const $pb.PackageName('shared.model.routes'))
    ..a<double>(1, 'latitude', $pb.PbFieldType.OD)
    ..a<double>(2, 'longitude', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  Coords() : super();
  Coords.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Coords.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Coords clone() => new Coords()..mergeFromMessage(this);
  Coords copyWith(void Function(Coords) updates) => super.copyWith((message) => updates(message as Coords));
  $pb.BuilderInfo get info_ => _i;
  static Coords create() => new Coords();
  static $pb.PbList<Coords> createRepeated() => new $pb.PbList<Coords>();
  static Coords getDefault() => _defaultInstance ??= create()..freeze();
  static Coords _defaultInstance;
  static void $checkItem(Coords v) {
    if (v is! Coords) $pb.checkItemFailed(v, _i.messageName);
  }

  double get latitude => $_getN(0);
  set latitude(double v) { $_setDouble(0, v); }
  bool hasLatitude() => $_has(0);
  void clearLatitude() => clearField(1);

  double get longitude => $_getN(1);
  set longitude(double v) { $_setDouble(1, v); }
  bool hasLongitude() => $_has(1);
  void clearLongitude() => clearField(2);
}

class Location extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Location', package: const $pb.PackageName('shared.model.routes'))
    ..aOS(1, 'name')
    ..a<Coords>(2, 'coords', $pb.PbFieldType.OM, Coords.getDefault, Coords.create)
    ..hasRequiredFields = false
  ;

  Location() : super();
  Location.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Location.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Location clone() => new Location()..mergeFromMessage(this);
  Location copyWith(void Function(Location) updates) => super.copyWith((message) => updates(message as Location));
  $pb.BuilderInfo get info_ => _i;
  static Location create() => new Location();
  static $pb.PbList<Location> createRepeated() => new $pb.PbList<Location>();
  static Location getDefault() => _defaultInstance ??= create()..freeze();
  static Location _defaultInstance;
  static void $checkItem(Location v) {
    if (v is! Location) $pb.checkItemFailed(v, _i.messageName);
  }

  String get name => $_getS(0, '');
  set name(String v) { $_setString(0, v); }
  bool hasName() => $_has(0);
  void clearName() => clearField(1);

  Coords get coords => $_getN(1);
  set coords(Coords v) { setField(2, v); }
  bool hasCoords() => $_has(1);
  void clearCoords() => clearField(2);
}

class Trip extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Trip', package: const $pb.PackageName('shared.model.routes'))
    ..a<Location>(1, 'start', $pb.PbFieldType.OM, Location.getDefault, Location.create)
    ..a<Location>(2, 'end', $pb.PbFieldType.OM, Location.getDefault, Location.create)
    ..e<TravelMethod>(3, 'method', $pb.PbFieldType.OE, TravelMethod.Unknown, TravelMethod.valueOf, TravelMethod.values)
    ..aInt64(4, 'duration')
    ..hasRequiredFields = false
  ;

  Trip() : super();
  Trip.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Trip.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Trip clone() => new Trip()..mergeFromMessage(this);
  Trip copyWith(void Function(Trip) updates) => super.copyWith((message) => updates(message as Trip));
  $pb.BuilderInfo get info_ => _i;
  static Trip create() => new Trip();
  static $pb.PbList<Trip> createRepeated() => new $pb.PbList<Trip>();
  static Trip getDefault() => _defaultInstance ??= create()..freeze();
  static Trip _defaultInstance;
  static void $checkItem(Trip v) {
    if (v is! Trip) $pb.checkItemFailed(v, _i.messageName);
  }

  Location get start => $_getN(0);
  set start(Location v) { setField(1, v); }
  bool hasStart() => $_has(0);
  void clearStart() => clearField(1);

  Location get end => $_getN(1);
  set end(Location v) { setField(2, v); }
  bool hasEnd() => $_has(1);
  void clearEnd() => clearField(2);

  TravelMethod get method => $_getN(2);
  set method(TravelMethod v) { setField(3, v); }
  bool hasMethod() => $_has(2);
  void clearMethod() => clearField(3);

  Int64 get duration => $_getI64(3);
  set duration(Int64 v) { $_setInt64(3, v); }
  bool hasDuration() => $_has(3);
  void clearDuration() => clearField(4);
}

class Trips extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Trips', package: const $pb.PackageName('shared.model.routes'))
    ..pp<Trip>(1, 'trips', $pb.PbFieldType.PM, Trip.$checkItem, Trip.create)
    ..hasRequiredFields = false
  ;

  Trips() : super();
  Trips.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Trips.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Trips clone() => new Trips()..mergeFromMessage(this);
  Trips copyWith(void Function(Trips) updates) => super.copyWith((message) => updates(message as Trips));
  $pb.BuilderInfo get info_ => _i;
  static Trips create() => new Trips();
  static $pb.PbList<Trips> createRepeated() => new $pb.PbList<Trips>();
  static Trips getDefault() => _defaultInstance ??= create()..freeze();
  static Trips _defaultInstance;
  static void $checkItem(Trips v) {
    if (v is! Trips) $pb.checkItemFailed(v, _i.messageName);
  }

  List<Trip> get trips => $_getList(0);
}

class RoutesAPIApi {
  $pb.RpcClient _client;
  RoutesAPIApi(this._client);

  Future<Trips> calcTravelTime($pb.ClientContext ctx, Trip request) {
    var emptyResponse = new Trips();
    return _client.invoke<Trips>(ctx, 'RoutesAPI', 'CalcTravelTime', request, emptyResponse);
  }
}

