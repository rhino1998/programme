#!/bin/sh

echo "Generating schedule"
protoc shared/model/schedule/schedule.proto \
    --gogoslick_out=../../../ \
    --dart_out=mobile/programme/lib \
    -I.

echo "Generating routes"
protoc shared/model/routes/routes.proto \
    --gogoslick_out=../../../ \
    --dart_out=mobile/programme/lib \
    -I.

echo "Generating weather"
protoc shared/model/weather/weather.proto \
    --gogoslick_out=../../../ \
    --dart_out=mobile/programme/lib \
    -I.