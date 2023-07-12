package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Instance *sql.DB

func InitDB() {
	var conn_str = fmt.Sprintf(
		"postgres://%s:%s@postgresql:5432/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	db, err := sql.Open("postgres", conn_str)

	if err != nil {
		log.Fatal("Couldn't connect to DB: ", err)
	}

	Instance = db
}

func CloseConnection() {
	Instance.Close()
}
