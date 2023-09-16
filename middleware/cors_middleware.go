package middleware

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CorsMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Println("get hit!!")
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace with your frontend's origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Continue processing the request
		next(w, r, ps)
	}
}
