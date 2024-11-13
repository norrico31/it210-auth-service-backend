package user

import (
	"encoding/json"
	"net/http"

	"github.com/norrico31/it210-auth-service-backend/entities"
)

type Handler struct {
	store entities.UserStore
}

func NewHandler(store entities.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload entities.UserLoginPayload

	// Decode the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if payload.Email == "" || payload.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Call the Login function
	token, user, err := h.store.Login(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Send user details and token in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user": map[string]interface{}{
			"id":           user.ID,
			"firstName":    user.FirstName,
			"lastName":     user.LastName,
			"email":        user.Email,
			"age":          user.Age,
			"lastActiveAt": user.LastActiveAt,
			"createdAt":    user.CreatedAt,
			"updatedAt":    user.UpdatedAt,
			"deletedAt":    user.DeletedAt,
		},
		"token": token,
	})
}
