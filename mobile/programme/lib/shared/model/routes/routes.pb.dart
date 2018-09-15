///
//  Generated code. Do not modify.
//  source: shared/model/routes/routes.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

// ignore: UNUSED_SHOWN_NAME
import 'dart:core' show int, bool, double, String, List, override;

import 'package:fixnum/fixnum.dart';
import 'package:protobuf/protobuf.dart' as $pb;

import 'routes.pbenum.dart';

export 'routes.pbenum.dart';

class Location extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Location', package: const $pb.PackageName('shared.model.routes'))
    ..aOS(1, 'name')
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
}

class Trip extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Trip', package: const $pb.PackageName('shared.model.routes'))
    ..a<Location>(1, 'start', $pb.PbFieldType.OM, Location.getDefault, Location.create)
    ..a<Location>(2, 'end', $pb.PbFieldType.OM, Location.getDefault, Location.create)
    ..e<TravelMethod>(3, 'method', $pb.PbFieldType.OE, TravelMethod.Walking, TravelMethod.valueOf, TravelMethod.values)
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

