///
//  Generated code. Do not modify.
//  source: shared/model/schedule/schedule.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

const Day$json = const {
  '1': 'Day',
  '2': const [
    const {'1': 'datetime', '3': 1, '4': 1, '5': 3, '10': 'datetime'},
    const {'1': 'tasks', '3': 2, '4': 3, '5': 11, '6': '.shared.model.schedule.Task', '10': 'tasks'},
    const {'1': 'weather_null', '3': 12, '4': 1, '5': 8, '9': 0, '10': 'weatherNull'},
    const {'1': 'traffic_null', '3': 14, '4': 1, '5': 8, '9': 1, '10': 'trafficNull'},
  ],
  '8': const [
    const {'1': 'weather'},
    const {'1': 'traffic'},
  ],
};

const Task$json = const {
  '1': 'Task',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'duration', '3': 2, '4': 1, '5': 3, '10': 'duration'},
    const {'1': 'stress', '3': 3, '4': 1, '5': 3, '10': 'stress'},
    const {'1': 'location_null', '3': 14, '4': 1, '5': 8, '9': 0, '10': 'locationNull'},
    const {'1': 'location_value', '3': 15, '4': 1, '5': 11, '6': '.shared.model.routes.Location', '9': 0, '10': 'locationValue'},
  ],
  '8': const [
    const {'1': 'location'},
  ],
};
