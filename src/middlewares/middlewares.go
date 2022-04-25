package middlewares

import (
	"log"
	"net/http"
	"social-network/src/authentication"
	"social-network/src/responses"
)

func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if error := authentication.ValidateToken(r); error != nil {
			responses.Error(w, http.StatusUnauthorized, error)
			return
		}
		nextFunc(w, r)
	}
}

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}
