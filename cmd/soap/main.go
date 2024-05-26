package main

import (
	"encoding/json"
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
	r.HandleFunc("/", Main).Methods(http.MethodGet)
	err := http.ListenAndServe("0.0.0.0:3000", r)
	log.Println("Server start")
	log.Fatal(err)
}

func AddCoin(w http.ResponseWriter, r *http.Request) {
	log.Print("ADD COIN")
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
