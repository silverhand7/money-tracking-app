package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func NewDB() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the .env")
	}

	_, err := sql.Open("postgresql", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}
}
