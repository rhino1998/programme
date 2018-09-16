package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/rhino1998/programme/backend/lib/model/mobile"
	"github.com/rhino1998/programme/backend/lib/model/schedule"
)

type Server struct {
	taskManager schedule.TaskManagerClient
}

func (s *Server) AddTask(ctx context.Context, taskReq *schedule.NewTaskRequest) (*schedule.Boolean, error) {
	return s.taskManager.AddTask(ctx, taskReq)
}

func main() {
	conn, err := grpc.Dial("task_manager:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	s := &Server{
		taskManager: schedule.NewTaskManagerClient(conn),
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
