package user

import (
	"net/http"

	"github.com/dunky-star/ecomm-proj/payloads"
	"github.com/gorilla/mux"
)

type Handler struct {
	store payloads.UserStore
}

func NewHandler(store payloads.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request){

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request){

}
