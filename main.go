package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/silverhand7/money-tracking-app/app"
)

func main() {
	apiConfig := app.NewDB()

	router := httprouter.New()

	router.GET("/", apiConfig.GetHomePage)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
