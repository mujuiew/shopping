package models

import (
	"database/sql"
	"fmt"
)

// DB ...
type DB struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
}

// Connect ...
func Connect(dc DB) *sql.DB {
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	return db
}
