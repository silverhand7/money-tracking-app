package middleware

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type authedHandler func(http.ResponseWriter, *http.Request, httprouter.Params)

func Auth(handler authedHandler, db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// apiKey := "hello"

		// user, err := apiConfig.DB.GetUserByAPIKey(r.Context(), apiKey)

		// if err != nil {
		// 	respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		// 	return
		// }
		fmt.Println("middleware")
		handler(w, r, nil)
	}
}
