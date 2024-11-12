package user

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, h *Handler) {
	router.HandleFunc("/", h.handleGetUsers).Methods("GET")
	router.HandleFunc("/helloworld", h.handleHelloWorld).Methods("GET")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/logout/{userId}", h.handleLogout).Methods("POST")
	router.HandleFunc("/{userId}", h.handleGetUser).Methods("GET")
	router.HandleFunc("/{userId}", h.HandleUpdateUser).Methods("PUT")
	router.HandleFunc("/{userId}", h.HandleDeleteUser).Methods("DELETE")
}
