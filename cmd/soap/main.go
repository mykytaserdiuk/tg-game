package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	coins = make(map[string]int)
)

func fixContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		w.Header().Add("Content-Type", "application/json")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {

	r := mux.NewRouter()
	r.Use(fixContentType)
	r.HandleFunc("/coin", AddCoin).Methods(http.MethodPut)
	r.HandleFunc("/coin", GetCoin).Methods(http.MethodGet)
	r.HandleFunc("/", Main).Methods(http.MethodGet)
	log.Println("Server starting....")
	fmt.Print("FMT")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, r))
	log.Println("Server start")
}

func AddCoin(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("user_id")
	if id == "" {
		w.WriteHeader(400)
		return
	}
	newC := coins[id] + 1
	coins[id] = newC
	w.Write([]byte(string(newC)))
}
func GetCoin(w http.ResponseWriter, r *http.Request) {
	log.Print("GET COIN")
	id := r.URL.Query().Get("user_id")
	if id == "" {
		w.WriteHeader(400)
		return
	}
	newC := coins[id]
	w.Write([]byte(string(newC)))
}
func Main(w http.ResponseWriter, r *http.Request) {
	log.Print("Main, full data")

	data, _ := json.Marshal(&coins)
	w.Write(data)
}
