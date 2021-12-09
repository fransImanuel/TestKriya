package db

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "13.228.73.161"
	port     = 8976
	user     = "test"
	password = "test"
	dbname   = "kriya_test"
)

var PsqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func Connect() *sqlx.DB {
	db, err := sqlx.Open("postgres", PsqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// in nanosecond
	db.SetMaxOpenConns(10)                 //max
	db.SetConnMaxIdleTime(2 * time.Second) //max time connection may be idle
	db.SetConnMaxLifetime(5 * time.Second) // max time connection may be reuse
	return db
}
