package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/rhino1998/programme/backend/lib/model/mobile"
	"github.com/rhino1998/programme/backend/lib/model/routes"
	"github.com/rhino1998/programme/backend/lib/model/schedule"
	"github.com/rhino1998/programme/backend/lib/model/weather"
)

type Server struct {
	taskManager schedule.TaskManagerClient
	weatherAPI  weather.WeatherAPIClient
}

func (s *Server) AddTask(ctx context.Context, taskReq *schedule.NewTaskRequest) (*schedule.Boolean, error) {
	return s.taskManager.AddTask(ctx, taskReq)
}

func (s *Server) GetForecast(ctx context.Context, coords *routes.Coords) (*weather.Forecast, error) {
	return s.weatherAPI.GetForecast(ctx, coords)
}

func main() {
	conn1, err := grpc.Dial("task_manager:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	conn2, err := grpc.Dial("weather_api:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	s := &Server{
		taskManager: schedule.NewTaskManagerClient(conn1),
		weatherAPI:  weather.NewWeatherAPIClient(conn2),
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":8080"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	mobile.RegisterMobileGatewayServer(grpcServer, s)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}
