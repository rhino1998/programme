package main

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"github.com/rhino1998/programme/backend/lib/model/weather"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var apiKey = os.Getenv("OWM_API_KEY")

type WeatherServer struct {
}

func (s *WeatherServer) GetForecast(ctx context.Context, c *Coord) *Forecast {
	w, err := owm.NewForecast("5", "C", "EN", apiKey)
	if err != nil {
		log.Fatalln(err)
	}

	w.DailyByCoordinates(
		&owm.Coordinates{
			Longitude: c.Longitude,
			Latitude:  c.Latitude,
		},
		5,
	)

	data, ok := w.ForecastWeatherJson.(*owm.Forecast5WeatherData)
	if !ok {
		return nil, fmt.Errorf("Invalid forecast type")
	}

	forecast := weather.ForecastFromOWM5(data)

	return forecast, nil
}

func main() {

}
