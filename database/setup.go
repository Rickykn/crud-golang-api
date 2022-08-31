package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() {
	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

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
