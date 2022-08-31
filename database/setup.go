package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	HOST = "localhost"
	PORT = "5432"
)

var db *sql.DB

func Connect() {
	godotenv.Load("../../.env")

	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")

	info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, DB_USER, DB_PASS, DB_NAME)

	var err error

	db, err = sql.Open("postgres", info) // connect
	if err != nil {
		fmt.Println(err)
	}
}

func Get() *sql.DB {
	return db
}
