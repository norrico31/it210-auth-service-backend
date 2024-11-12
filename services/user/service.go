package user

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/norrico31/it210-auth-service-backend/entities"
	"github.com/norrico31/it210-auth-service-backend/utils"
)

type Handler struct {
	store entities.UserStore
}

func NewHandler(store entities.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	str := fmt.Sprintf("Auth Service: Hello World ")
	utils.WriteJSON(w, http.StatusOK, str)
}

func (h *Handler) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.GetUsers()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	w.Header().Set("Content-Type", "application/json")

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"data": users})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload entities.UserRegisterPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(entities.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload entities.UserLoginPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if !utils.ComparePasswords(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid credentials"))
		return
	}

	err = h.store.SetUserActive(u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"token": token, "user": u})
}

func (h *Handler) handleLogout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["userId"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user Id"))
		return
	}
	userId, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user Id"))
		return
	}
	now := time.Now()
	err = h.store.UpdateLastActiveTime(userId, now)
}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["userId"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user Id"))
		return
	}
	userId, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID"))
	}

	user, err := h.store.GetUserById(userId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["userId"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user ID"))
		return
	}

	userId, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID"))
		return
	}

	var payload entities.UserUpdatePayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user := entities.User{
		ID:        userId,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
	}

	if payload.Password != "" {
		hashedPassword, err := utils.HashPassword(payload.Password)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		user.Password = hashedPassword
	}

	err = h.store.UpdateUser(user)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["userId"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user ID"))
		return
	}

	userId, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID"))
		return
	}

	err = h.store.DeleteUser(userId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
