package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func connectDB() {
	connStr := "user = postgres password = ****** dbname = todo_app sslmode = disable"

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

}
