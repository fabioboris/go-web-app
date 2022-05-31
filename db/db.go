package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDB() *sql.DB {
	connStr := "user=fabioboris dbname=playground sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}
