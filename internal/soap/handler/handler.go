package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mykytaserdiuk/souptgbot/internal/soap"
)

type Handler struct {
	walletService soap.WalletService
	userService   soap.UserService
}

func New(r *mux.Router, walletService soap.WalletService, userService soap.UserService) *Handler {
	h := &Handler{walletService, userService}

	r.HandleFunc("/wallet", h.CreateWallet).Methods(http.MethodPost)
	r.HandleFunc("/user", h.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/wallet", h.GetWallet).Methods(http.MethodGet)
	r.HandleFunc("/", h.Main).Methods(http.MethodGet)
	return h
}

func (h *Handler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("user_id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("user_id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.userService.Get(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
}

func (h *Handler) GetWallet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("user_id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	walletID := r.URL.Query().Get("wallet_id")
	if walletID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	wallet, err := h.walletService.Get(r.Context(), walletID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	data, err := json.Marshal(&wallet)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
func (h *Handler) Main(w http.ResponseWriter, r *http.Request) {
	data, err := h.walletService.Admin(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json, err := json.Marshal(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
