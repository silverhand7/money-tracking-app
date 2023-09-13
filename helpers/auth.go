package helpers

import (
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) string {
	val := headers.Get("Authorization")

	if val == "" {
		panic("no authorization info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		panic("malformed auth header")
	}

	if vals[0] != "Bearer" {
		panic("malformed first part of auth header")
	}

	return vals[1]
}
