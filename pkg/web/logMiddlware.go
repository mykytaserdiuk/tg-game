package web

import (
	"log"
	"net/http"
)

func FixContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
