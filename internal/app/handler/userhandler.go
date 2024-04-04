package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-rest-api/internal/app/model"
	"go-rest-api/internal/app/service"

	"github.com/gorilla/mux"
)

// UserHandler represents the HTTP handler for managing users.
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler creates a new instance of UserHandler.
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUser handles POST requests to create a new user.
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	if err := h.userService.CreateUser(r.Context(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

	return nil
}

// GetUserByID handles GET requests to fetch a user by ID.
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	user, err := h.userService.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return err
	}

	json.NewEncoder(w).Encode(user)
	return nil
}

// GetAllUsers handles GET requests to fetch all users.
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := h.userService.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

	return nil
}
