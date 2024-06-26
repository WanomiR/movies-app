package main

import (
	"database/sql"
	"log"

	// include to use drivers
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func (s *WebServer) connectToDB() (*sql.DB, error) {
	conn, err := sql.Open("pgx", s.DSN)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("connected to database")

	return conn, nil
}
