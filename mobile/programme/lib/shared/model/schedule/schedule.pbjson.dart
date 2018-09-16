///
//  Generated code. Do not modify.
//  source: shared/model/schedule/schedule.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

const TaskType$json = const {
  '1': 'TaskType',
  '2': const [
    const {'1': 'Unknown', '2': 0},
    const {'1': 'Floating', '2': 1},
    const {'1': 'Scheduled', '2': 2},
    const {'1': 'Calendar', '2': 3},
    const {'1': 'Free', '2': 4},
    const {'1': 'Travel', '2': 5},
  ],
};

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
    const {'1': 'task_type', '3': 1, '4': 1, '5': 14, '6': '.shared.model.schedule.TaskType', '10': 'taskType'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'description', '3': 3, '4': 1, '5': 9, '10': 'description'},
    const {'1': 'duration', '3': 4, '4': 1, '5': 3, '10': 'duration'},
    const {'1': 'stress', '3': 5, '4': 1, '5': 3, '10': 'stress'},
    const {'1': 'deadline_null', '3': 8, '4': 1, '5': 8, '9': 0, '10': 'deadlineNull'},
    const {'1': 'deadline_value', '3': 9, '4': 1, '5': 3, '9': 0, '10': 'deadlineValue'},
    const {'1': 'start_null', '3': 10, '4': 1, '5': 8, '9': 1, '10': 'startNull'},
    const {'1': 'start_value', '3': 11, '4': 1, '5': 3, '9': 1, '10': 'startValue'},
    const {'1': 'travel_method_null', '3': 12, '4': 1, '5': 8, '9': 2, '10': 'travelMethodNull'},
    const {'1': 'travel_method_value', '3': 13, '4': 1, '5': 14, '6': '.shared.model.routes.TravelMethod', '9': 2, '10': 'travelMethodValue'},
    const {'1': 'location_null', '3': 14, '4': 1, '5': 8, '9': 3, '10': 'locationNull'},
    const {'1': 'location_value', '3': 15, '4': 1, '5': 11, '6': '.shared.model.routes.Location', '9': 3, '10': 'locationValue'},
  ],
  '8': const [
    const {'1': 'deadline'},
    const {'1': 'start'},
    const {'1': 'travel_method'},
    const {'1': 'location'},
  ],
};

const NewTaskRequest$json = const {
  '1': 'NewTaskRequest',
  '2': const [
    const {'1': 'user', '3': 1, '4': 1, '5': 9, '10': 'user'},
    const {'1': 'task', '3': 2, '4': 1, '5': 11, '6': '.shared.model.schedule.Task', '10': 'task'},
  ],
};

const Boolean$json = const {
  '1': 'Boolean',
  '2': const [
    const {'1': 'boolean', '3': 1, '4': 1, '5': 8, '10': 'boolean'},
  ],
};

