package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/silverhand7/money-tracking-app/app/data"
)

type ApiConfig struct {
	DB *data.Queries
}

func NewDB() ApiConfig {
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

	apiConfig := ApiConfig{
		DB: db,
	}

	return apiConfig
}

func (apiconfig *ApiConfig) GetHomePage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "hello world")
}
