package weather

import (
	owm "github.com/briandowns/openweathermap"
)

func ForecastFromOWM5(w *owm.Forecast5WeatherData) *Forecast {
	forecast := &Forecast{}
	forecast.City = CityFromOWM(&w.City)

	for i := range w.List {
		forecast.List = append(forecast.List, ListFromOWM5(&w.List[i]))
	}

	return forecast
}

func ListFromOWM5(l *owm.Forecast5WeatherList) *List {
	list := &List{
		Dt:     int64(l.Dt),
		Main:   MainFromOWM(&l.Main),
		Clouds: CloudsFromOWM(&l.Clouds),
		Wind:   WindFromOWM(&l.Wind),
		Rain:   RainFromOWM(&l.Rain),
	}

	for i := range l.Weather {
		list.Weather = append(list.Weather, WeatherFromOWM(&l.Weather[i]))
	}

	return list
}

func CityFromOWM(city *owm.City) *City {
	return &City{
		Id:      int64(city.ID),
		Name:    city.Name,
		Coord:   CoordFromOWM(&city.Coord),
		Country: city.Country,
	}
}

func CoordFromOWM(coord *owm.Coordinates) *Coord {
	return &Coord{
		Latitude:  coord.Latitude,
		Longitude: coord.Longitude,
	}
}

func MainFromOWM(main *owm.Main) *Main {
	return &Main{
		Temperature: main.Temp,
		TempMin:     main.TempMin,
		TempMax:     main.TempMax,
		Pressure:    main.Pressure,
		SeaLevel:    main.SeaLevel,
		GroundLevel: main.GrndLevel,
		Humidity:    int64(main.Humidity),
	}
}

func WeatherFromOWM(weather *owm.Weather) *Weather {
	return &Weather{
		Id:          int64(weather.ID),
		Main:        weather.Main,
		Description: weather.Description,
	}
}

func SysFromOWM(sys *owm.Sys) *Sys {
	return &Sys{
		Country: sys.Country,
		Sunrise: uint64(sys.Sunrise),
		Sunset:  uint64(sys.Sunset),
	}
}

func WindFromOWM(wind *owm.Wind) *Wind {
	return &Wind{
		Speed: wind.Speed,
		Deg:   wind.Deg,
	}
}

func RainFromOWM(rain *owm.Rain) *Rain {
	return &Rain{
		ThreeHours: rain.ThreeH,
	}
}

func SnowFromOWM(snow *owm.Snow) *Snow {
	return &Snow{
		ThreeHours: snow.ThreeH,
	}
}

func CloudsFromOWM(clouds *owm.Clouds) *Clouds {
	return &Clouds{
		All: int64(clouds.All),
	}
}

func ForecastFromOWM16(w *owm.Forecast16WeatherData) *Forecast {
	forecast := &Forecast{}
	forecast.City = CityFromOWM(&w.City)

	for i := range w.List {
		forecast.List = append(forecast.List, ListFromOWM16(&w.List[i]))
	}

	return forecast
}

func ListFromOWM16(l *owm.Forecast16WeatherList) *List {
	list := &List{
		Dt: int64(l.Dt),
		Main: &Main{
			Temperature: l.Temp.Day,
			TempMin:     l.Temp.Min,
			TempMax:     l.Temp.Max,
		},
		Clouds: &Clouds{
			All: int64(l.Clouds),
		},
		Wind: &Wind{
			Speed: l.Speed,
		},
		Rain: &Rain{
			ThreeHours: l.Rain,
		},
	}

	for i := range l.Weather {
		list.Weather = append(list.Weather, WeatherFromOWM(&l.Weather[i]))
	}

	return list
}
