package user

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, h *Handler) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
}
