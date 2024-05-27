package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

var (
	coins = make(map[string]int)
)

func New(r *mux.Router) *Handler {
	h := &Handler{}

	r.HandleFunc("/coin", h.AddCoin).Methods(http.MethodPut)
	r.HandleFunc("/coin", h.GetCoin).Methods(http.MethodGet)
	r.HandleFunc("/", h.Main).Methods(http.MethodGet)
	return h
}

func (h *Handler) AddCoin(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("user_id")
	if id == "" {
		w.WriteHeader(400)
		return
	}
	newC := coins[id] + 1
	coins[id] = newC
	w.Write([]byte(string(newC)))
}
func (h *Handler) GetCoin(w http.ResponseWriter, r *http.Request) {
	log.Print("GET COIN")
	id := r.URL.Query().Get("user_id")
	if id == "" {
		w.WriteHeader(400)
		return
	}
	newC := coins[id]
	w.Write([]byte(string(newC)))
}
func (h *Handler) Main(w http.ResponseWriter, r *http.Request) {
	log.Print("Main, full data")

	data, _ := json.Marshal(&coins)
	w.Write(data)
}
