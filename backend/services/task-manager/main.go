package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/rhino1998/programme/backend/lib/model/schedule"
)

type Server struct {
	db *sql.DB
}

func (s *Server) Setup() error {

	_, err := s.db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		name VARCHAR(256),
		email VARCHAR(256)
	)
	`)

	if err != nil {
		return err
	}

	_, err = s.db.Exec(`
	INSERT INTO users
		(id, name, email)
		SELECT 1, 'Riley Wilburn', 'jamesrileywilburn@gmail.com'
		WHERE
			NOT EXISTS (
				SELECT id FROM users WHERE id = 1
			)
	`)

	if err != nil {
		return err
	}

	_, err = s.db.Exec(`
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY,
		name varchar(256) NOT NULL,
		duration INTERVAL NOT NULL,
		stress INTEGER NOT NULL,
		start TIMESTAMP NULL,
		deadline TIMESTAMP NULL,
		location_namem varchar(256) NULL,
		location_lat DOUBLE PRECISION NULL,
		location_lon DOUBLE PRECISION NULL,
		travel_method INTEGER NULL,
		description TEXT
	)
	`)
	return err
}

func (s *Server) AddTask(ctx context.Context, taskReq *schedule.NewTaskRequest) (*schedule.Boolean, error) {
	if taskReq.Task == nil {
		return nil, fmt.Errorf("invalid task add request: no task")
	}
	task := taskReq.GetTask()

	if task.GetLocationNull() != task.GetTravelMethodNull() {
		return nil, fmt.Errorf("invalid task add request: traveling tasks with no travel method")
	}

	var err error
	switch task.GetTaskType() {
	case schedule.Floating:
		if task.GetDeadlineNull() {
			return nil, fmt.Errorf("invalid task add request: floating task with no deadline")
		}

		if !task.GetLocationNull() {
			location := task.GetLocationValue()
			_, err = s.db.ExecContext(ctx, `
				INSERT INTO tasks
					(name, description, duration, stress,
					 location_name, location_lat, location_lon,
					 travel_method,deadline
				 )
				 VALUES ((?,?,?,?,?,?,?,?,?))
			`,
				task.Name, task.Description, task.Duration, task.Stress,
				location.Name, location.Coords.Latitude, location.Coords.Longitude,
				task.GetTravelMethodValue(), task.GetDeadlineValue(),
			)
		} else {
			_, err = s.db.ExecContext(ctx, `
				INSERT INTO tasks
					(name, description,duration, stress,deadline)
					VALUES ((?,?,?,?,?))
				`,
				task.Name, task.Description, task.Duration, task.Stress, task.GetDeadlineValue(),
			)
		}
	case schedule.Scheduled:
		if task.GetStartNull() {
			return nil, fmt.Errorf("invalid task add request: scheduled task with no start")
		}

		if !task.GetLocationNull() {
			location := task.GetLocationValue()
			_, err = s.db.ExecContext(ctx, `
				INSERT INTO tasks
					(name, description, duration, stress,
					 location_name, location_lat, location_lon,
					 start, travel_method
				 )
				 VALUES ((?,?,?,?,?,?,?,?,?))
			`,
				task.Name, task.Description, task.Duration, task.Stress,
				location.Name, location.Coords.Latitude, location.Coords.Longitude,
				time.Unix(task.GetStartValue(), 0), task.GetTravelMethod(),
			)
		} else {
			_, err = s.db.ExecContext(ctx, `
				INSERT INTO tasks
					(name, description,duration, stress,start)
					VALUES ((?,?,?,?,?))
				`,
				task.Name, task.Description, task.Duration, task.Stress,
				time.Unix(task.GetStartValue(), 0),
			)
		}
	default:
		return nil, fmt.Errorf("invalid task add request: invalid task type: %s", schedule.TaskType_name[int32(task.GetTaskType())])
	}

	if err != nil {
		return nil, err
	}

	return &schedule.Boolean{Boolean: true}, nil

}

func main() {
	connStrFmt := "postgres://%s:%s@postgres/%s?sslmode=disable"

	connStr := fmt.Sprintf(connStrFmt, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	s := &Server{
		db: db,
	}
	err = s.Setup()
	if err != nil {
		log.Fatalln(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":8080"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	schedule.RegisterTaskManagerServer(grpcServer, s)
	grpcServer.Serve(lis)

}
