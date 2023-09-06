package main

import (
	"fmt"

	"github.com/silverhand7/money-tracking-app/app"
)

func main() {
	apiConfig := app.NewDB()
	fmt.Println(apiConfig)
}
