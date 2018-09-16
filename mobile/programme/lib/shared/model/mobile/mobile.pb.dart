///
//  Generated code. Do not modify.
//  source: shared/model/mobile/mobile.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async';
// ignore: UNUSED_SHOWN_NAME
import 'dart:core' show int, bool, double, String, List, override;

import 'package:protobuf/protobuf.dart' as $pb;

import '../schedule/schedule.pb.dart' as $0;

class MobileGatewayApi {
  $pb.RpcClient _client;
  MobileGatewayApi(this._client);

  Future<$0.Boolean> addTask($pb.ClientContext ctx, $0.NewTaskRequest request) {
    var emptyResponse = new $0.Boolean();
    return _client.invoke<$0.Boolean>(ctx, 'MobileGateway', 'AddTask', request, emptyResponse);
  }
}

