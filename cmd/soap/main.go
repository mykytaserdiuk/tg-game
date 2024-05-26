package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	coins = make(map[string]int)
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/coin", AddCoin).Methods(http.MethodPut)
	r.HandleFunc("/coin", GetCoin).Methods(http.MethodGet)
	err := http.ListenAndServe(":1234", r)
	log.Fatal(err)
}

func AddCoin(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["user_id"]
	if id == "" {
		w.WriteHeader(400)
		return
	}
	newC := coins[id] + 1
	coins[id] = newC
	w.Write([]byte(string(newC)))
}
func GetCoin(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["user_id"]
	if id == "" {
		w.WriteHeader(400)
		return
	}
	newC := coins[id]
	w.Write([]byte(string(newC)))
}
