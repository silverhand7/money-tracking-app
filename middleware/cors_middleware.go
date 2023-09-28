package middleware

import (
	"net/http"
)

type CorsMiddleware struct {
	Handler http.Handler
}

func NewCorsMiddleware(handler http.Handler) *CorsMiddleware {
	return &CorsMiddleware{
		Handler: handler,
	}
}

func (middleware *CorsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials ", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	middleware.Handler.ServeHTTP(w, r)
}
