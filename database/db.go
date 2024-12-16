package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func connect() {
	var err error
	DB, err = sqlx.Connect("postgres", "user=postgres dbname=taskmanager password=golang")
	if err != nil {
		log.Fatal("Failed to connect db:", err)

	} else {
		log.Println("DB connected successfully")
	}
}
