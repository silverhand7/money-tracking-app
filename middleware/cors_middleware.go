package middleware

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CorsMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Println("get hit!!")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")

		// Continue processing the request
		next(w, r, ps)
	}
}
