///
//  Generated code. Do not modify.
//  source: shared/model/routes/routes.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' show int, dynamic, String, List, Map;
import 'package:protobuf/protobuf.dart' as $pb;

class TravelMethod extends $pb.ProtobufEnum {
  static const TravelMethod Walking = const TravelMethod._(0, 'Walking');
  static const TravelMethod Biking = const TravelMethod._(1, 'Biking');
  static const TravelMethod Rideshare = const TravelMethod._(2, 'Rideshare');
  static const TravelMethod Car = const TravelMethod._(3, 'Car');

  static const List<TravelMethod> values = const <TravelMethod> [
    Walking,
    Biking,
    Rideshare,
    Car,
  ];

  static final Map<int, dynamic> _byValue = $pb.ProtobufEnum.initByValue(values);
  static TravelMethod valueOf(int value) => _byValue[value] as TravelMethod;
  static void $checkItem(TravelMethod v) {
    if (v is! TravelMethod) $pb.checkItemFailed(v, 'TravelMethod');
  }

  const TravelMethod._(int v, String n) : super(v, n);
}

