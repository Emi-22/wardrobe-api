package db

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func Connect() {
	connStr := "host=localhost port=5432 user=closet_user password=closet_pass dbname=closet_db sslmodel=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed trying to connect to the database: ", err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal("Failed making ping to the database: ", err)
	}

	fmt.Println("Connected to PostgreSQL")
}
