package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
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
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY,
		uid INTEGER REFERENCES users(id),
		name varchar(256),
		duration INTERVAL NOT NULL,
		stress INTEGER NOT NULL,
		start TIMESTAMP NULL,
		location_namem varchar(256) NULL,
		location_lat INTEGER NULL,
		location_lon INTEGER NULL,
		travel_method INTEGER NULL,
		description TEXT
	)
	`)
	return err
}

// func (s *Server) AddTask(ctx context.Context, task *schedule.Task) (*schedule.Boolean, error) {
// 	s.db.Exec("")
// }

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
}
