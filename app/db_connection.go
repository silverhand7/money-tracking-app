package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/silverhand7/money-tracking-app/app/data"
)

type apiConfig struct {
	DB *data.Queries
}

func NewDB() apiConfig {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the .env")
	}

	connection, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	db := data.New(connection)

	apiConfig := apiConfig{
		DB: db,
	}

	return apiConfig
}
