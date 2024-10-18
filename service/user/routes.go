package user

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, h *Handler) {
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/logout/{userId}", h.handleLogout).Methods("POST")
	router.HandleFunc("/user/{userId}", h.handleGetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", h.HandleUpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", h.HandleDeleteUser).Methods("DELETE")
}
