package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // just import libpq
)

// DB ...
var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "shopping"
)

// InitDB ...
func InitDB() {
	var err error
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	fmt.Println("Successfully connected! DB")
}
