///
//  Generated code. Do not modify.
//  source: shared/model/routes/routes.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

const TravelMethod$json = const {
  '1': 'TravelMethod',
  '2': const [
    const {'1': 'Unknown', '2': 0},
    const {'1': 'Walking', '2': 1},
    const {'1': 'Biking', '2': 2},
    const {'1': 'Rideshare', '2': 3},
    const {'1': 'Car', '2': 4},
  ],
};

const Coords$json = const {
  '1': 'Coords',
  '2': const [
    const {'1': 'Latitude', '3': 1, '4': 1, '5': 1, '10': 'Latitude'},
    const {'1': 'Longitude', '3': 2, '4': 1, '5': 1, '10': 'Longitude'},
  ],
};

const Location$json = const {
  '1': 'Location',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'coords', '3': 2, '4': 1, '5': 11, '6': '.shared.model.routes.Coords', '10': 'coords'},
  ],
};

const Trip$json = const {
  '1': 'Trip',
  '2': const [
    const {'1': 'start', '3': 1, '4': 1, '5': 11, '6': '.shared.model.routes.Location', '10': 'start'},
    const {'1': 'end', '3': 2, '4': 1, '5': 11, '6': '.shared.model.routes.Location', '10': 'end'},
    const {'1': 'method', '3': 3, '4': 1, '5': 14, '6': '.shared.model.routes.TravelMethod', '10': 'method'},
    const {'1': 'duration', '3': 4, '4': 1, '5': 3, '10': 'duration'},
  ],
};

const Trips$json = const {
  '1': 'Trips',
  '2': const [
    const {'1': 'trips', '3': 1, '4': 3, '5': 11, '6': '.shared.model.routes.Trip', '10': 'trips'},
  ],
};

const RoutesAPI$json = const {
  '1': 'RoutesAPI',
  '2': const [
    const {'1': 'CalcTravelTime', '2': '.shared.model.routes.Trip', '3': '.shared.model.routes.Trips'},
  ],
};

const RoutesAPI$messageJson = const {
  '.shared.model.routes.Trip': Trip$json,
  '.shared.model.routes.Location': Location$json,
  '.shared.model.routes.Coords': Coords$json,
  '.shared.model.routes.Trips': Trips$json,
};

