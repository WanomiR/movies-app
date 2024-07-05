package dbrepo

import (
	"database/sql"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (db *PostgresDBRepo) Connection() *sql.DB {
	return db.DB
}
